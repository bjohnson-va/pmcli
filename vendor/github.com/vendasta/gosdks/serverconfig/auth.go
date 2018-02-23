package serverconfig

import (
	"github.com/vendasta/gosdks/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/vendasta/gosdks/util"
)

type userInfoKey struct{}

// ErrNotAuthenticated is thrown when the user is deemed unauthenticated due to missing information on the request metadata
var ErrNotAuthenticated = util.Error(util.PermissionDenied, "unable to determine user from context")

// UserInfo holds the user info deserialized from the request metadata
type UserInfo struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

// SetUserOnContext returns a context with a user set on it
func SetUserOnContext(ctx context.Context, userInfo *UserInfo) context.Context {
	logging.Infof(ctx, "Attaching user to context: %v", userInfo)
	return context.WithValue(ctx, userInfoKey{}, userInfo)
}

// GetUserInfoFromContext retrieves user info from the context's metadata
func GetUserInfoFromContext(ctx context.Context) (*UserInfo, error) {
	s, ok := ctx.Value(userInfoKey{}).(*UserInfo)
	if !ok {
		return nil, ErrNotAuthenticated
	}
	return s, nil
}

// AuthInterceptor handles auth through GRPC, bypassing it on local environments
// TODO: Remove this as it is only for backward compatibility with mscli services prior to 1.28.0
func AuthInterceptor(authorizedAccounts map[string][]string) grpc.UnaryServerInterceptor {
	return NewJwtAuthInterceptor(authorizedAccounts).Interceptor()
}

//GRPCInterceptor is implemented by interceptors
type GRPCInterceptor interface {
	// Interceptor intercepts requests to check authorization
	Interceptor() grpc.UnaryServerInterceptor
}
