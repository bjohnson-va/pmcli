package serverconfig

import (
	"fmt"
	"net"
	"runtime/debug"

	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/vendasta/gosdks/util"
)

// CreateGrpcServer creates a basic GRPC Server with the specified interceptors
func CreateGrpcServer(interceptors ...grpc.UnaryServerInterceptor) *grpc.Server {
	interceptors = append(interceptors, recoveryInterceptor)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(chainUnaryServer(interceptors...)),
	)
	return s
}

// StartGrpcServer starts a new server to handle GRPC requests
func StartGrpcServer(server *grpc.Server, port int) error {
	var lis net.Listener
	var err error

	if lis, err = net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
		logging.Errorf(context.Background(), "Error creating GRPC listening socket: %s", err.Error())
		return err
	}

	//The following call blocks until an error occurs
	return server.Serve(lis)
}

func recoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			logging.Criticalf(ctx, "Recovered from panic: %s", debug.Stack())
			err = util.ToGrpcError(util.Error(util.Internal, "An unexpected error occured"))
		}
	}()
	return handler(ctx, req)
}

// NoAuthInterceptor satisfies the GRPCInterceptor but does not actually check any auth
func NoAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

// ChainUnaryServer combines multiple grpc.UnaryServerInterceptor into a single grpc.UnaryServerInterceptor (required by GRPC)
func chainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		buildChain := func(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return current(currentCtx, currentReq, info, next)
			}
		}
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = buildChain(interceptors[i], chain)
		}
		return chain(ctx, req)
	}
}
