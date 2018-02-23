package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"encoding/base64"

	"encoding/json"

	"errors"

	"os"

	"github.com/vendasta/gosdks/iam/subject"
	"github.com/vendasta/gosdks/iam/subjectcontext"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"github.com/vendasta/gosdks/iam/resources"
)

type subjectPolicyMock struct {
	getBySessionReturnValue subject.Subject
	getBySessionError       error

	accessPartnerMarketError error
	accessAccountGroupError  error
	accessAccountsError      error

	requestedScopes []AccessScope
}

func (s *subjectPolicyMock) GetBySessionID(ctx context.Context, iamContext *subjectcontext.Context, sessionID string) (subject.Subject, error) {
	if s.getBySessionError != nil {
		return nil, s.getBySessionError
	}
	return s.getBySessionReturnValue, nil
}

func (s *subjectPolicyMock) GetByEmail(ctx context.Context, iamContext *subjectcontext.Context, email string) (subject.Subject, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) GetBySubjectID(ctx context.Context, iamContext *subjectcontext.Context, subjectID string) (subject.Subject, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) GetSubjectContext(ctx context.Context, subjectID string) (*subjectcontext.Context, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) ListPersonas(ctx context.Context, sessionID, cursor string, pageSize int64, personaType string) ([]subject.Subject, string, bool, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) AccessResource(context.Context, *AccessResource) error {
	panic("not implemented")
}

func (s *subjectPolicyMock) Register(ctx context.Context, iamContext *subjectcontext.Context, email, password string, attributes *iam_v1.StructAttribute) (string, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) AddKey(ctx context.Context, iamContext *subjectcontext.Context, email string) (PrivateKey, string, error) {
	panic("not implemented")
}

func (s *subjectPolicyMock) RemoveKey(ctx context.Context, iamContext *subjectcontext.Context, email, keyID string) error {
	panic("not implemented")
}

func (s *subjectPolicyMock) AccessPartnerMarket(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, partnerID string, marketIDs []string, accessScopes ...AccessScope) error {
	s.requestedScopes = accessScopes
	return s.accessPartnerMarketError
}

func (s *subjectPolicyMock) AccessAccountGroup(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupID string, accessScopes ...AccessScope) error {
	s.requestedScopes = accessScopes
	return s.accessAccountGroupError
}

func (s *subjectPolicyMock) AccessAccountGroups(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupIDs []string, accessScopes ...AccessScope) error {
	s.requestedScopes = accessScopes
	return s.accessAccountGroupError
}

func (s *subjectPolicyMock) AccessAccounts(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountIdentifiers []*resources.AccountIdentifier, accessScopes ...AccessScope) error {
	s.requestedScopes = accessScopes
	return s.accessAccountsError
}

func Test_JWTMustBeBase64Encoded(t *testing.T) {
	// arrange
	md := metadata.MD{"x-endpoint-api-userinfo": []string{"test-key1"}}
	as := &authService{&subjectPolicyMock{}, nil, nil}

	// act
	ctx, err := as.validateContextHasIdentity(
		metadata.NewIncomingContext(context.TODO(), md),
	)

	// assert
	assert.Nil(t, ctx)
	assert.Equal(t, "Permission denied", err.Error())
}

func Test_JWTMustBeProperJSON(t *testing.T) {
	// arrange
	d := base64.URLEncoding.EncodeToString([]byte("not-json"))
	as := &authService{&subjectPolicyMock{}, nil, nil}

	// act
	ctx, err := as.validateContextHasIdentity(
		metadata.NewIncomingContext(context.TODO(), metadata.MD{"x-endpoint-api-userinfo": []string{d}}),
	)

	// assert
	assert.Nil(t, ctx)
	assert.Equal(t, "Permission denied", err.Error())
}

func Test_SessionIsSetOnContext(t *testing.T) {
	// arrange
	MockPublicKey()
	token, _ := CreateTestJWT("bbass@vendasta.com")
	sa, _ := json.Marshal(&ServiceAccountInfo{Issuer: "iam.vendasta-local.com", Email: "bbass@vendasta.com"})
	d := base64.URLEncoding.EncodeToString(sa)

	// act
	as := &authService{&subjectPolicyMock{}, nil, nil}
	ctx, err := as.validateContextHasIdentity(
		metadata.NewIncomingContext(context.TODO(), metadata.MD{"x-endpoint-api-userinfo": {d}, "authorization": {"Bearer " + token}}),
	)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, ctx)
	session, err := as.GetSessionFromContext(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, session)
}

func Test_ServiceAccountIsSetOnContext(t *testing.T) {
	// arrange
	sa, _ := json.Marshal(&ServiceAccountInfo{Issuer: "accounts.google.com", Email: "bbass@vendasta.com"})
	d := base64.URLEncoding.EncodeToString(sa)

	// act
	as := &authService{&subjectPolicyMock{}, nil, nil}
	ctx, err := as.validateContextHasIdentity(
		metadata.NewIncomingContext(context.TODO(), metadata.MD{"x-endpoint-api-userinfo": {d}}),
	)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, ctx)
	saInfo, err := as.GetServiceAccountFromContext(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, saInfo)
}

func Test_UnknownServiceAccountReturnsError(t *testing.T) {
	// arrange
	sa, _ := json.Marshal(&ServiceAccountInfo{Issuer: "hackers.com", Email: "bbass@vendasta.com"})
	d := base64.URLEncoding.EncodeToString(sa)
	as := &authService{&subjectPolicyMock{}, nil, nil}

	// act
	ctx, err := as.validateContextHasIdentity(
		metadata.NewIncomingContext(context.TODO(), metadata.MD{"x-endpoint-api-userinfo": {d}}),
	)

	// assert
	assert.Equal(t, "Permission denied", err.Error())
	assert.Nil(t, ctx)

}

func Test_IsContextAuthorizedToAccessAccountGroup_DefaultsRequestedScopesToReadWriteDelete(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	mock.accessAccountGroupError = errors.New("iam broked")
	as := &authService{mock, nil, nil,}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Equal(t, "iam broked", err.Error())
}

func Test_IsContextAuthorizedToAccessAccountGroup_UsesProvidedScopes(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123", READ)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessAccountGroup_DefaultsScopesToReadWriteDelete(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ, WRITE, DELETE}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessAccountGroup_ReturnsNilOnSuccessFromIAM(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessAccountGroups_ReturnsNilOnSuccessFromIAM(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccountGroups(ctx, []string{"AG-123"})

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessAccountGroup_ReturnsNilOnSuccessWithValidServiceAccount(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE, DELETE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessAccountGroup_ReturnsErrorOnInvalidServiceAccount(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Equal(t, "Service account isn't authorized.", err.Error())
}

func Test_IsContextAuthorizedToAccessAccountGroup_ReturnsNilOnSuccessWithValidServiceAccountByScopes(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE, DELETE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123", READ)

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessAccountGroup_ReturnsErrOnInvalidServiceAccountByScopes(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123", WRITE)

	//assert
	assert.Equal(t, "Service account test@test.com isn't authorized for this scope: WRITE", err.Error())
}

func Test_DoesContextHaveValidServiceAccount_ReturnsNilOnValidServiceAccountByPartialScopes(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE, DELETE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.DoesContextHaveValidServiceAccount(ctx, WRITE, DELETE)

	//assert
	assert.Nil(t, err)
}

func Test_DoesContextHaveValidServiceAccount_ReturnsErrOnInvalidServiceAccountByPartialScopes(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.DoesContextHaveValidServiceAccount(ctx, WRITE, DELETE)

	//assert
	assert.Equal(t, "Service account test@test.com isn't authorized for this scope: DELETE", err.Error())
}

func Test_DoesContextHaveValidServiceAccount_DefaultsRequestToAllScopesIfNone(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.DoesContextHaveValidServiceAccount(ctx)

	//assert
	assert.Equal(t, "Service account test@test.com isn't authorized for this scope: DELETE", err.Error())
}

func Test_IsContextAuthorizedToAccessPartnerMarket_UsesProvidedScopes(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessPartnerMarket(ctx, "ABC", nil, READ)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessPartnerMarket_DefaultsScopesToReadWriteDelete(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessPartnerMarket(ctx, "ABC", nil)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ, WRITE, DELETE}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessPartnerMarket_ReturnsErrorFromIAM(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	mock.accessPartnerMarketError = errors.New("iam broked")
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessPartnerMarket(ctx, "ABC", nil)

	//assert
	assert.Equal(t, "iam broked", err.Error())
}

func Test_IsContextAuthorizedToAccessAccounts_UsesProvidedScopes(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccounts(ctx, []*resources.AccountIdentifier{}, READ)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessAccounts_DefaultsScopesToReadWriteDelete(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccounts(ctx, []*resources.AccountIdentifier{})

	//assert
	assert.Nil(t, err)
	assert.Equal(t, []AccessScope{READ, WRITE, DELETE}, mock.requestedScopes)
}

func Test_IsContextAuthorizedToAccessAccounts_ReturnsErrorFromIAM(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	mock.accessAccountsError = errors.New("iam broked")
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessAccounts(ctx, []*resources.AccountIdentifier{})

	//assert
	assert.Equal(t, "iam broked", err.Error())
}

func Test_IsContextAuthorizedToAccessPartnerMarket_ReturnsNilOnSuccessFromIAM(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.setSessionOnContext(context.TODO(), "session-id")

	// act
	err := as.IsContextAuthorizedToAccessPartnerMarket(ctx, "ABC", nil)

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessPartnerMarket_ReturnsNilOnSuccessWithValidServiceAccount(t *testing.T) {
	// arrange
	mock := &subjectPolicyMock{}
	accessScopedServiceAccounts := ServiceAccountToScopes{
		"test@test.com": {READ, WRITE, DELETE},
	}
	as := &authService{mock, accessScopedServiceAccounts, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{Email: "test@test.com"})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Nil(t, err)
}

func Test_IsContextAuthorizedToAccessPartnerMarket_ReturnsErrorOnInvalidServiceAccount(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	defer os.Unsetenv("ENVIRONMENT")

	// arrange
	mock := &subjectPolicyMock{}
	as := &authService{mock, nil, nil}
	ctx := as.SetServiceAccountOnContext(context.TODO(), &ServiceAccountInfo{})

	// act
	err := as.IsContextAuthorizedToAccessAccountGroup(ctx, "AG-123")

	//assert
	assert.Equal(t, "Service account isn't authorized.", err.Error())
}

func Test_AllowPublicMethods_AddsMethodsToMap(t *testing.T) {
	as := &authService{nil, nil, map[string]struct{}{}}

	as.AllowPublicMethods("/method/1", "/method/2")

	assert.Contains(t, as.publicMethods, "/method/1")
	assert.Contains(t, as.publicMethods, "/method/2")
}

func Test_IsPublicMethod_ReturnsTrueIfMethodPublic(t *testing.T) {
	as := &authService{nil, nil, map[string]struct{}{"/public/method": {}}}

	assert.True(t, as.isPublicMethod("/public/method"))
}

func Test_IsPublicMethod_ReturnsFalseIfMethodNotPublic(t *testing.T) {
	as := &authService{nil, nil, nil}

	assert.False(t, as.isPublicMethod("/other/method"))
}
