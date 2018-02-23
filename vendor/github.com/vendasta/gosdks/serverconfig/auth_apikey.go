package serverconfig

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// APIKeyUser contains a user ID and an API key
type APIKeyUser struct {
	Key string
	UID string
}

// APIKeyAuthInterceptor uses an API user, API key pair to authorize requests
type APIKeyAuthInterceptor struct {
	authorizedKeys map[string][]APIKeyUser
}

// NewAPIKeyAuthInterceptor creates a new APIKeyAuthInterceptor
func NewAPIKeyAuthInterceptor(authorizedKeys map[string][]APIKeyUser) *APIKeyAuthInterceptor {
	rv := APIKeyAuthInterceptor{authorizedKeys}
	return &rv
}

// Interceptor intercepts requests to check for authorization.
func (m APIKeyAuthInterceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		ctx, err = m.attachUserInfoToContext(ctx)
		if err != nil {
			return nil, grpc.Errorf(codes.PermissionDenied, "Call isn't authenticated.")
		}
		return handler(ctx, req)
	}
}

func (m APIKeyAuthInterceptor) getUIDFromToken(token string) (string, error) {
	authedKeys := m.authorizedKeys[config.Getenvironment()]
	for _, k := range authedKeys {
		if k.Key == token {
			return k.UID, nil
		}
	}
	return "", fmt.Errorf("Token %s invalid", token)
}

// attachUserInfoToContext adds a UserInfo object to the provided context
func (m APIKeyAuthInterceptor) attachUserInfoToContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logging.Errorf(ctx, "No metadata associated with this request.")
		return nil, ErrNotAuthenticated
	}
	data, ok := md["authorization"]
	if !ok || len(data) != 1 {
		logging.Errorf(ctx, "Didn't get an authorization header.")
		return nil, ErrNotAuthenticated
	}
	chunks := strings.Split(data[0], " ")
	if len(chunks) != 2 {
		logging.Errorf(ctx, "Error, Authorization header value malformed %s", data)
		return nil, ErrNotAuthenticated
	}
	token := chunks[1]
	uid, err := m.getUIDFromToken(token)
	if err != nil {
		logging.Errorf(ctx, "Invalid token provided: %s", err.Error())
		return nil, ErrNotAuthenticated
	}
	userInfo := UserInfo{Email: "", ID: uid}
	return SetUserOnContext(ctx, &userInfo), nil
}
