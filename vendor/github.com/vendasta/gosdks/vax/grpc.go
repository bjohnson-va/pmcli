package vax

import (
	"crypto/tls"
	"time"

	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGRPCConnection(ctx context.Context, address string, useTLS bool, audience string, enableToken bool, dialOptions ...grpc.DialOption) (*grpc.ClientConn, error) {
	dialOptions = append(
		dialOptions,
		grpc.WithBalancer(grpc.RoundRobin(NewPoolResolver(3, &DialSettings{Endpoint: address}))),
		grpc.WithBackoffConfig(grpc.DefaultBackoffConfig),
	)
	if useTLS {
		dialOptions = append(dialOptions,
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		)
		if enableToken {
			tokenSource, err := NewTokenSource(audience)
			if err != nil {
				return nil, err
			}
			dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(tokenSource))
		}
	} else {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	}

	return grpc.Dial(address, dialOptions...)
}

//TimeoutInterceptor forces a timeout
func TimeoutInterceptor(maxTimeout time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, cancel := context.WithTimeout(ctx, maxTimeout)
		defer cancel()
		return handler(newCtx, req)
	}
}

// RequestLoggingInterceptor logs the request
func RequestLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logging.Debugf(ctx, "%T.%s called with %+v", info.Server, info.FullMethod, req)
		return handler(ctx, req)
	}
}
