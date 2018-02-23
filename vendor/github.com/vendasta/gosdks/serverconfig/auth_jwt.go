package serverconfig

import (
	"encoding/base64"
	"encoding/json"

	"golang.org/x/net/context"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"github.com/vendasta/gosdks/util"
)

// JWTAuthInterceptor uses JWT tokens to authorize requests
type JWTAuthInterceptor struct {
	authorizedAccounts map[string][]string
}

// NewJwtAuthInterceptor creats a new JWTAuthInterceptor
func NewJwtAuthInterceptor(authorizedAccounts map[string][]string) *JWTAuthInterceptor {
	rv := JWTAuthInterceptor{authorizedAccounts: authorizedAccounts}
	return &rv
}

// Interceptor intercepts requests to check for authorization.
func (m JWTAuthInterceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !config.IsLocal() {
			var err error
			ctx, err = m.attachUserInfoToContext(ctx)
			if err != nil {
				return nil, grpc.Errorf(codes.PermissionDenied, "Call isn't authenticated.")
			}
			if !m.isContextAuthorized(ctx) {
				return nil, grpc.Errorf(codes.PermissionDenied, "Not authorized.")
			}
		}
		return handler(ctx, req)
	}
}


func (m JWTAuthInterceptor) attachUserInfoToContext(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logging.Errorf(ctx, "No metadata associated with this request.")
		return nil, ErrNotAuthenticated
	}
	data, ok := md["x-endpoint-api-userinfo"]
	if !ok || len(data) != 1 {
		logging.Errorf(ctx, "Didn't get user info from cloud endpoints.")
		return nil, ErrNotAuthenticated
	}
	b, err := base64.URLEncoding.DecodeString(data[0])
	if err != nil {
		logging.Errorf(ctx, "Error decoding user info string. %s %s", err.Error(), data[0])
		return nil, err
	}
	userInfo := &UserInfo{}
	err = json.Unmarshal(b, userInfo)
	if err != nil {
		logging.Errorf(ctx, "Error decoding json for user info.")
		return nil, err
	}

	return SetUserOnContext(ctx, userInfo), nil
}

func (m JWTAuthInterceptor) isContextAuthorized(ctx context.Context) bool {
	if config.IsLocal() {
		return true
	}
	user, err := GetUserInfoFromContext(ctx)
	if err != nil {
		return false
	}
	if util.StringInSlice(user.Email, m.authorizedAccounts[config.Getenvironment()]) {
		return true
	}
	return false
}
