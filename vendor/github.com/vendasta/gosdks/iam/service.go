package iam

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"github.com/vendasta/gosdks/iam/resources"
	"github.com/vendasta/gosdks/pb/iam/v1"
)

type iamKey struct{}
type sessionKey struct{}
type serviceAccountInfoKey struct{}

// AuthService provides a gRPC interceptor that validates the caller is calling with a valid identity.  It also saves
// the callers identity onto the context.  Valid identities are an IAM Subject or a Google service account.
// The application must provide the set of valid Google service accounts that are allowed to call this API.
//
// AuthService also provides functions to easily manage permissions for account groups and partner data.  It is up to
// the application to ask if the caller has permissions dependent on the level of resource check and the specific
// resource data such as Partner ID, Market ID or account group ID.
type AuthService interface {
	Interceptor() grpc.UnaryServerInterceptor
	DoesContextHaveValidServiceAccount(ctx context.Context, scopes ...AccessScope) error
	IsContextAuthorizedToAccessResource(ctx context.Context, resourceID string, resourceOwnerID string, resourceAttributes *iam_v1.StructAttribute, scopes ...AccessScope) error
	IsContextAuthorizedToAccessAccountGroups(ctx context.Context, accountGroupIDs []string, scopes ...AccessScope) error
	IsContextAuthorizedToAccessAccountGroup(ctx context.Context, accountGroupID string, scopes ...AccessScope) error
	IsContextAuthorizedToAccessPartnerMarket(ctx context.Context, partnerID string, marketIDs []string, scopes ...AccessScope) error
	IsContextAuthorizedToAccessAccounts(ctx context.Context, accountIdentifiers []*resources.AccountIdentifier, scopes ...AccessScope) error

	GetServiceAccountFromContext(ctx context.Context) (*ServiceAccountInfo, error)
	GetSessionFromContext(ctx context.Context) (string, error)

	SetServiceAccountOnContext(ctx context.Context, serviceAccountInfo *ServiceAccountInfo) context.Context

	AllowPublicMethods(publicMethods ...string)
}

type ServiceAccountToScopes map[string][]AccessScope

// NewAuthService returns a new AuthService service.
func NewAuthService(iamAPI SubjectPolicyAPI, accessScopedServiceAccounts ServiceAccountToScopes) AuthService {
	return &authService{iamAPI: iamAPI, accessScopedServiceAccounts: accessScopedServiceAccounts, publicMethods: map[string]struct{}{}}
}

// authService provides an interceptor for verifying a callers identity and methods for validating permissions.
type authService struct {
	iamAPI SubjectPolicyAPI
	// The allowed service accounts, scoped by access type (read/write/delete)
	accessScopedServiceAccounts ServiceAccountToScopes
	// full RPC method strings, i.e., "/accountgroup.v1.AccountGroupService/Create"
	publicMethods map[string]struct{}
}

// ServiceAccountInfo holds the user info deserialized from the request metadata
type ServiceAccountInfo struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

var errNotAuthenticated = util.Error(util.Unauthenticated, "Permission denied")

// Interceptor validates incoming requests are validated through a service account or an IAM session.
func (i *authService) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if i.isPublicMethod(info.FullMethod) {
			logging.Infof(ctx, "Allowing call on public method %s", info.FullMethod)
			return handler(ctx, req)
		}

		newCtx, err := i.validateContextHasIdentity(ctx)
		if err != nil {
			if config.IsLocal() && err == errNotAuthenticated {
				newCtx = ctx
				logging.Infof(ctx, "Allowing unauthenticated call on local: %s.", err.Error())
			} else {
				logging.Infof(ctx, "Validate context encountered error %s for path %s", err.Error(), info.FullMethod)
				return nil, util.ToGrpcError(err)
			}
		}
		return handler(newCtx, req)
	}
}

func (i *authService) AllowPublicMethods(publicMethods ...string) {
	for _, method := range publicMethods {
		i.publicMethods[method] = struct{}{}
	}
}

func (i *authService) isPublicMethod(fullMethod string) bool {
	_, ok := i.publicMethods[fullMethod]
	return ok
}

// validateContextHasIdentity ensures identity on the context
func (i *authService) validateContextHasIdentity(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logging.Errorf(ctx, "No metadata associated with this request.")
		return nil, errNotAuthenticated
	}

	data, ok := md["x-endpoint-api-userinfo"]
	if !ok || len(data) != 1 {
		logging.Errorf(ctx, "Didn't get user info from cloud endpoints.")
		return nil, errNotAuthenticated
	}
	b, err := base64.URLEncoding.DecodeString(data[0])
	if err != nil {
		logging.Errorf(ctx, "Error decoding user info string. %s %s", err.Error(), data)
		return nil, errNotAuthenticated
	}
	serviceAccountInfo := &ServiceAccountInfo{}
	err = json.Unmarshal(b, serviceAccountInfo)
	if err != nil {
		logging.Errorf(ctx, "Error decoding json for user info.")
		return nil, errNotAuthenticated
	}

	if i.isIssuedByIAM(serviceAccountInfo) {
		pieces := strings.Split(md["authorization"][0], " ")
		if len(pieces) != 2 || pieces[0] != "Bearer" {
			return nil, errNotAuthenticated
		}
		iamSession := pieces[1]
		_, err := ValidateJWT(iamSession)
		if err != nil {
			return nil, errNotAuthenticated
		}
		ctx = i.setSessionOnContext(ctx, iamSession)
		return ctx, nil
	}
	if !(serviceAccountInfo.Issuer == "accounts.google.com" || serviceAccountInfo.Issuer == "https://accounts.google.com") {
		return nil, errNotAuthenticated
	}
	return i.SetServiceAccountOnContext(ctx, serviceAccountInfo), nil
}

// DoesContextHaveValidServiceAccount determines if the given context has an authorized service account.
func (i *authService) DoesContextHaveValidServiceAccount(ctx context.Context, scopes ...AccessScope) error {
	if config.IsLocal() {
		return nil
	}
	if i.accessScopedServiceAccounts != nil {
		serviceAccount, err := i.GetServiceAccountFromContext(ctx)
		if err != nil {
			return err
		}
		accountScopes, ok := i.accessScopedServiceAccounts[serviceAccount.Email]
		if !ok {
			return util.Error(util.PermissionDenied, "Service account isn't authorized.")
		}

		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		for _, s := range scopes {
			found := false
			for _, vs := range accountScopes {
				if s == vs {
					found = true
					break
				}
			}
			if found == false {
				return util.Error(util.PermissionDenied, "Service account %s isn't authorized for this scope: %s", serviceAccount.Email, AccessScopeName[s])
			}
		}
		return nil
	}
	return util.Error(util.PermissionDenied, "Service account isn't authorized.")
}

// IsContextAuthorizedToAccessResource returns if the context is authorized to access the requested resource
func (i *authService) IsContextAuthorizedToAccessResource(ctx context.Context, resourceID string, resourceOwnerID string, resourceAttributes *iam_v1.StructAttribute, scopes ...AccessScope) error {
	sessionID, _ := i.GetSessionFromContext(ctx)
	if sessionID != "" {
		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		err := i.iamAPI.AccessResource(
			util.NewContext(ctx),
			&AccessResource{
				SessionID:          sessionID,
				ResourceID:         resourceID,
				OwnerID:            resourceOwnerID,
				ResourceAttributes: resourceAttributes,
				AccessScope:        scopes,
			},
		)
		if err != nil {
			if util.IsError(util.PermissionDenied, err) {
				logging.Debugf(ctx, "Session %s not authorized to access resource %s: %s", sessionID, resourceID, err.Error())
			} else {
				logging.Debugf(ctx, "Unknown error calling IAM access resource %s", err.Error())
			}
			return err
		}
		return nil
	}
	return i.DoesContextHaveValidServiceAccount(ctx, scopes...)
}

// IsContextAuthorizedToAccessAccountGroups returns if the context is authorized to access the list of account group ids
func (i *authService) IsContextAuthorizedToAccessAccountGroups(ctx context.Context, accountGroupIDs []string, scopes ...AccessScope) error {
	sessionID, _ := i.GetSessionFromContext(ctx)
	if sessionID != "" {
		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		err := i.iamAPI.AccessAccountGroups(
			util.NewContext(ctx),
			nil, nil,
			sessionID,
			accountGroupIDs,
			scopes...,
		)
		if err != nil {
			if util.IsError(util.PermissionDenied, err) {
				logging.Debugf(ctx, "Session %s not authorized to access account groups %s %s", sessionID, accountGroupIDs, err.Error())
			} else {
				logging.Debugf(ctx, "Unknown error calling IAM access resource %s", err.Error())
			}
			return err
		}
		return nil
	}
	return i.DoesContextHaveValidServiceAccount(ctx, scopes...)
}

// IsContextAuthorizedToAccessAccountGroup returns if the context is authorized to access the specified account group
func (i *authService) IsContextAuthorizedToAccessAccountGroup(ctx context.Context, accountGroupID string, scopes ...AccessScope) error {
	sessionID, _ := i.GetSessionFromContext(ctx)
	if sessionID != "" {
		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		err := i.iamAPI.AccessAccountGroup(
			util.NewContext(ctx),
			nil, nil,
			sessionID,
			accountGroupID,
			scopes...,
		)
		if err != nil {
			if util.IsError(util.PermissionDenied, err) {
				logging.Debugf(ctx, "Session %s not authorized to access account group %s %s", sessionID, accountGroupID, err.Error())
			} else {
				logging.Debugf(ctx, "Unknown error calling IAM access resource %s", err.Error())
			}
			return err
		}
		return nil
	}
	return i.DoesContextHaveValidServiceAccount(ctx, scopes...)
}

// IsContextAuthorizedToAccessPartnerMarket returns if the context is authorized to access the specified partner or partner/market combo.
func (i *authService) IsContextAuthorizedToAccessPartnerMarket(ctx context.Context, partnerID string, marketIDs []string, scopes ...AccessScope) error {
	sessionID, _ := i.GetSessionFromContext(ctx)
	if sessionID != "" {
		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		err := i.iamAPI.AccessPartnerMarket(
			util.NewContext(ctx),
			nil, nil,
			sessionID,
			partnerID,
			marketIDs,
			scopes...,
		)
		if err != nil {
			if util.IsError(util.PermissionDenied, err) {
				logging.Debugf(ctx, "Session %s not authorized to access partner %s %#v %s", sessionID, partnerID, marketIDs, err.Error())
			} else {
				logging.Debugf(ctx, "Unknown error calling IAM access resource %s", err.Error())
			}
			return err
		}
		return nil
	}
	return i.DoesContextHaveValidServiceAccount(ctx, scopes...)
}

// IsContextAuthorizedToAccessAccounts returns if the context is authorized to access the specified accounts.
func (i *authService) IsContextAuthorizedToAccessAccounts(ctx context.Context, accountIdentifiers []*resources.AccountIdentifier, scopes ...AccessScope) error {
	sessionID, _ := i.GetSessionFromContext(ctx)
	if sessionID != "" {
		if len(scopes) == 0 {
			scopes = []AccessScope{READ, WRITE, DELETE}
		}
		err := i.iamAPI.AccessAccounts(
			util.NewContext(ctx),
			nil, nil,
			sessionID,
			accountIdentifiers,
			scopes...,
		)
		if err != nil {
			if util.IsError(util.PermissionDenied, err) {
				logging.Debugf(ctx, "Session %s not authorized to access accounts %#v %s", accountIdentifiers, err.Error())
			} else {
				logging.Debugf(ctx, "Unknown error calling IAM access resource %s", err.Error())
			}
			return err
		}
		return nil
	}
	return i.DoesContextHaveValidServiceAccount(ctx, scopes...)
}

// setSessionOnContext returns a context with a session set on it
func (i *authService) setSessionOnContext(ctx context.Context, session string) context.Context {
	return context.WithValue(ctx, sessionKey{}, session)
}

// GetSessionFromContext retrieves a subject from the context's metadata
func (i *authService) GetSessionFromContext(ctx context.Context) (string, error) {
	s, ok := ctx.Value(sessionKey{}).(string)
	if !ok {
		return "", nil
	}
	return s, nil
}

// SetServiceAccountOnContext sets the given service account info on the context
func (i *authService) SetServiceAccountOnContext(ctx context.Context, serviceAccountInfo *ServiceAccountInfo) context.Context {
	return context.WithValue(ctx, serviceAccountInfoKey{}, serviceAccountInfo)
}

// GetServiceAccountFromContext retrieves the service account from the context's metadata
func (i *authService) GetServiceAccountFromContext(ctx context.Context) (*ServiceAccountInfo, error) {
	s, ok := ctx.Value(serviceAccountInfoKey{}).(*ServiceAccountInfo)
	if !ok {
		return nil, errNotAuthenticated
	}
	return s, nil
}

var iamIssuerMap = map[config.Env]string{
	config.Local: "iam.vendasta-local.com",
	config.Test:  "iam-test.vendasta-internal.com",
	config.Demo:  "iam-demo.vendasta-internal.com",
	config.Prod:  "iam-prod.vendasta-internal.com",
}

func (i *authService) isIssuedByIAM(s *ServiceAccountInfo) bool {
	return iamIssuerMap[config.CurEnv()] == s.Issuer
}
