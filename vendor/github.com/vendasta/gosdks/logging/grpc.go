package logging

import (
	"io"
	"net/http"
	"time"

	"net/url"

	gce_metadata "cloud.google.com/go/compute/metadata"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

// Interceptor provides logging/tracing for incoming gRPC requests
func Interceptor() grpc.UnaryServerInterceptor {
	if gce_metadata.OnGCE() {
		i := &grpcInterceptor{config: configValue, logger: GetLogger()}
		return i.UnaryServerInterceptor
	}
	return PassThroughUnaryServerInterceptor
}

// ClientInterceptor should be used for outgoing gRPC requests.
//
// Should be provided as a dial option on creation of a gRPC transport with grpc.UnaryInterceptor(logging.ClientInterceptor())
func ClientInterceptor() grpc.UnaryClientInterceptor {
	if gce_metadata.OnGCE() {
		i := &grpcInterceptor{config: configValue, logger: GetLogger()}
		return i.UnaryClientInterceptor
	}
	return PassThroughUnaryClientInterceptor
}

// ClientStreamInterceptor should be used for outgoing gRPC stream requests.
func ClientStreamInterceptor() grpc.StreamClientInterceptor {
	if gce_metadata.OnGCE() {
		i := &grpcInterceptor{config: configValue, logger: GetLogger()}
		return i.StreamClientInterceptor
	}
	return PassThroughStreamClientInterceptor
}

type grpcInterceptor struct {
	config *config
	logger Logger
}

// PassThroughUnaryServerInterceptor provides an empty incoming interceptor.
func PassThroughUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

// PassThroughUnaryClientInterceptor provides an empty outgoing interceptor.
func PassThroughUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(ctx, method, req, reply, cc, opts...)
}

// PassThroughStreamClientInterceptor provides an empty outgoing stream interceptor.
func PassThroughStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(ctx, desc, cc, method, opts...)
}

// UnaryServerInterceptor provides an an incoming interceptor for logging/tracing.
func (g *grpcInterceptor) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			Errorf(ctx, "Recovered from panic, returning 500. %s", err.Error())
			err = grpc.Errorf(codes.Internal, "Unknown error occurred")
		}
	}()

	start := time.Now().UTC()
	requestID := g.logger.RequestID()
	ctx, rd := newRequest(ctx, requestID)

	span := g.config.TracingClient.SpanFromContext(ctx, info.FullMethod)
	defer span.Finish()
	ctx = NewContext(ctx, span)

	resp, err := handler(ctx, req)
	responseSize := 0
	respProto, ok := resp.(proto.Message)
	if resp != nil && ok {
		responseSize = proto.Size(respProto)
	}

	end := time.Now().UTC()
	statusCode := HTTPStatusFromCode(grpc.Code(err))
	if statusCode >= 500 {
		Errorf(ctx, "Error serving request with code %d Error: %s", grpc.Code(err), err.Error())
	}
	traceID := requestID
	if span != nil {
		traceID = span.TraceID()
	}

	rd.HTTPRequest.Request.URL = &url.URL{Path: info.FullMethod}
	rd.HTTPRequest.Request.Method = "POST"
	rd.HTTPRequest.Status = int(statusCode)
	rd.HTTPRequest.ResponseSize = int64(responseSize)
	rd.HTTPRequest.LocalIP = "127.0.0.1"
	rd.HTTPRequest.RemoteIP = "127.0.0.1"
	rd.HTTPRequest.Latency = end.Sub(start)
	rd.Trace = traceID

	fillInFromGRPCMetadata(ctx, rd)
	logRequest(ctx, rd)
	return resp, err
}

func fillInFromGRPCMetadata(ctx context.Context, r *requestData) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}
	r.HTTPRequest.Request.Header = http.Header(md)
}

// UnaryServerInterceptor provides an an outgoing interceptor for logging/tracing.
func (g *grpcInterceptor) UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	span := FromContext(ctx)
	ctx, childSpan := span.NewRemoteChild(ctx, method)
	defer childSpan.Finish()

	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

// StreamClientInterceptor provides an an outgoing stream interceptor for logging/tracing.
func (g *grpcInterceptor) StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	span := FromContext(ctx)
	ctx, childSpan := span.NewRemoteChild(ctx, method)

	cs, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		childSpan.Finish()
		return nil, err
	}

	return &monitoredClientStream{cs, childSpan}, nil
}

// monitoredClientStream wraps grpc.ClientStream allowing each Sent/Recv of message to increment counters.
type monitoredClientStream struct {
	grpc.ClientStream
	span *Span
}

func (s *monitoredClientStream) SendMsg(m interface{}) error {
	return s.ClientStream.SendMsg(m)
}

func (s *monitoredClientStream) RecvMsg(m interface{}) error {
	err := s.ClientStream.RecvMsg(m)
	if err == io.EOF {
		s.span.Finish()
	} else {
		s.span.Finish()
	}
	return err
}

// HTTPStatusFromCode returns an http status code from a gRPC code.
func HTTPStatusFromCode(code codes.Code) int32 {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusForbidden
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	grpclog.Printf("Unknown gRPC error code: %v", code)
	return http.StatusInternalServerError
}
