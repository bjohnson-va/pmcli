package iam

import (
	"errors"
	"fmt"

	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/iam/attribute"
	"github.com/vendasta/gosdks/iam/resources"
	"github.com/vendasta/gosdks/iam/subject"
	"github.com/vendasta/gosdks/iam/subjectcontext"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var addresses = map[config.Env]string{
	config.Local: "IAM:11000",
	config.Test:  "iam-api-test.vendasta-internal.com:443",
	config.Demo:  "iam-api-demo.vendasta-internal.com:443",
	config.Prod:  "iam-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://iam-api-test.vendasta-internal.com",
	config.Demo:  "https://iam-api-demo.vendasta-internal.com",
	config.Prod:  "https://iam-api-prod.vendasta-internal.com",
}

var certProviderUrls = map[config.Env]string{
	config.Local: "https://iam-test.vendasta-internal.com/oauth2/v1/certs",
	config.Test:  "https://iam-test.vendasta-internal.com/oauth2/v1/certs",
	config.Demo:  "https://iam-demo.vendasta-internal.com/oauth2/v1/certs",
	config.Prod:  "https://iam-prod.vendasta-internal.com/oauth2/v1/certs",
}

// NewClient returns an IAM client.
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	providerURL := certProviderUrls[e]
	if address == "" || providerURL == "" {
		return nil, fmt.Errorf("unable to create client with environment %d", e)
	}
	conn, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	noTokenConnection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], false, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &iamClient{iam_v1.NewIAMClient(conn), iam_v1.NewUserIAMClient(noTokenConnection), providerURL}, nil
}

// PartnerAPI defines all of the methods available to interface with Partner subjects. These are convenience methods for Partners wrapping the SubjectAPI.
type PartnerAPI interface {
	PartnerBySessionID(ctx context.Context, sessionID string) (*subject.Partner, error)
	PartnerByEmail(ctx context.Context, email string) (*subject.Partner, error)
	PartnerBySubjectID(ctx context.Context, subjectID string) (*subject.Partner, error)
}

// SMBAPI defines all of the methods available to interface with SMB subjects. These are convenience methods for SMBs wrapping the SubjectAPI.
type SMBAPI interface {
	RegisterSMBSubject(ctx context.Context, partnerID, email string, accountGroupIDs []string) (string, error)
	SMBBySessionID(ctx context.Context, partnerID string, sessionID string) (*subject.SMB, error)
	SMBByEmail(ctx context.Context, partnerID string, email string) (*subject.SMB, error)
	SMBBySubjectID(ctx context.Context, partnerID string, subjectID string) (*subject.SMB, error)
}

// SalesPersonAPI defines all of the methods available to interface with SalesPerson subjects. These are convenience methods for SalesPersons wrapping the SubjectAPI.
type SalesPersonAPI interface {
	SalesPersonBySessionID(ctx context.Context, partnerID string, sessionID string) (*subject.SalesPerson, error)
	SalesPersonByEmail(ctx context.Context, partnerID string, email string) (*subject.SalesPerson, error)
	SalesPersonBySubjectID(ctx context.Context, partnerID string, subjectID string) (*subject.SalesPerson, error)
}

// PartnerAppAPI defines all of the methods available to interface with PartnerApp subjects. These are convenience methods for PartnerApps wrapping the SubjectAPI.
type PartnerAppAPI interface {
	RegisterPartnerApp(ctx context.Context, partnerID string, email string) (string, error)
	PartnerAppByEmail(ctx context.Context, partnerID string, email string) (*subject.PartnerApp, error)
	PartnerAppAddKey(ctx context.Context, partnerID string, subjectID string) (PrivateKey, string, error)
	PartnerAppRemoveKey(ctx context.Context, partnerID, email, keyID string) error
	PartnerAppGenerateClientKey(ctx context.Context, partnerID string, email string) (*ClientKey, error)
}

// AuthAPI defines methods for interacting with IAM's authentication APIs
type AuthAPI interface {
	// ExchangeToken asks IAM to exchange the caller-provided token for a IAM session token.
	ExchangeToken(ctx context.Context, key *ClientKey) (string, error)
	// GenerateClientKey adds a new key to the specified subject and returns a secret containing all the information needed to be able to exchange tokens for IAM sessions.
	// This secret can only be retrieved once per key. Keep it secret, keep it safe.
	GenerateClientKey(ctx context.Context, subCtx *subjectcontext.Context, email string) (*ClientKey, error)

	// GetShortLivedToken retrieves a shortlived IAM session token for a subject
	GetShortLivedToken(ctx context.Context, iamContext *subjectcontext.Context, email string) (string, error)
}

// PrivateKey is encoded as a PEM
type PrivateKey string

// SubjectAPI defines all of the methods available to interact with IAM subjects.
type SubjectAPI interface {
	// GetBySessionID returns a subject associated to the given IAM session.
	GetBySessionID(ctx context.Context, iamContext *subjectcontext.Context, sessionID string) (subject.Subject, error)

	// GetByEmail returns a subject with the given email.
	GetByEmail(ctx context.Context, iamContext *subjectcontext.Context, email string) (subject.Subject, error)

	// GetBySubjectID returns a subject with the given subject id.
	GetBySubjectID(ctx context.Context, iamContext *subjectcontext.Context, subjectID string) (subject.Subject, error)

	// Register registers a new subject with IAM and returns the SubjectID provisioned by IAM
	Register(ctx context.Context, iamContext *subjectcontext.Context, email, password string, attributes *iam_v1.StructAttribute) (string, error)

	// AddKey generates a new key pair for the subject in IAM, returning the private key and its key id.
	// Note that the private key can not be retrieved after this call.
	AddKey(ctx context.Context, iamContext *subjectcontext.Context, email string) (PrivateKey, string, error)

	// RemoveKey removes the key with the given keyID from the specified subject.
	RemoveKey(ctx context.Context, iamContext *subjectcontext.Context, email, keyID string) error

	// GetSubjectContext returns the IAM subject context from the given subject id.
	GetSubjectContext(ctx context.Context, subjectID string) (*subjectcontext.Context, error)

	// ListPersonas lists all personas available to a session
	ListPersonas(ctx context.Context, sessionID, cursor string, pageSize int64, personaType string) ([]subject.Subject, string, bool, error)
}

// ResourceAPI defines all of the methods available to interact with IAM resources.
type ResourceAPI interface {
	// Registers a resource owner with IAM. This sets an application up for integration with IAM access control system.
	RegisterResourceOwner(context.Context, *ResourceOwner) error

	// Registers a specific resource with IAM. This allows IAM to ask the resource owner about its resources.
	RegisterResource(context.Context, *Resource) error

	// Registers a policy associated with a resource with IAM that is evaluated whenever a subject requests access to that type of resource.
	RegisterPolicy(context.Context, *Policy) error
}

// PolicyAPI defines all of the methods available to interact with IAM policies.
type PolicyAPI interface {
	// Asks IAM whether a specified subject has access to a specified resource
	AccessResource(context.Context, *AccessResource) error

	// AccessPartnerMarket asks IAM if a given user has permissions to the partner or partner/market combo.
	AccessPartnerMarket(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, partnerID string, marketIDs []string, accessScopes ...AccessScope) error

	// AccessAccountGroup asks IAM if a given user has permissions to the given account group.
	AccessAccountGroup(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupID string, accessScopes ...AccessScope) error

	// AccessAccountGroups asks IAM if a given user has permissions to all of the given account groups.
	AccessAccountGroups(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupIDs []string, accessScopes ...AccessScope) error

	// AccessAccounts asks IAM if a given user has permissions to all of the given accounts.
	AccessAccounts(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountIdentifiers []*resources.AccountIdentifier, accessScopes ...AccessScope) error
}

// Interface defines all of the API methods available from IAM.
type Interface interface {
	SubjectAPI
	ResourceAPI
	PolicyAPI

	// Subject type specific APIs.
	PartnerAPI
	SMBAPI
	SalesPersonAPI
	PartnerAppAPI

	AuthAPI
}

// SubjectPolicyAPI provides subject and policy methods for the IAM permissions service.
type SubjectPolicyAPI interface {
	SubjectAPI
	PolicyAPI
}

type iamClient struct {
	client      iam_v1.IAMClient
	userClient  iam_v1.UserIAMClient
	providerURL string
}

var (
	// ErrInvalidSubjectType is returned when a specific subject type API is passed a subject with the invalid ype.
	ErrInvalidSubjectType = errors.New("Invalid subject type has been found.")
)

func (ic *iamClient) GetBySessionID(ctx context.Context, iamContext *subjectcontext.Context, sessionID string) (subject.Subject, error) {
	var subjectResp *iam_v1.GetSubjectResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		subjectResp, err = ic.userClient.GetSubjectBySession(ctx, &iam_v1.GetSubjectBySessionRequest{
			Context: iamContext.ToPB(), Session: sessionID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	s, err := subject.New(iamContext, subjectResp.Subject.Subject)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ic *iamClient) GetByEmail(ctx context.Context, iamContext *subjectcontext.Context, email string) (subject.Subject, error) {
	var subjectResp *iam_v1.GetSubjectsResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		subjectResp, err = ic.client.GetSubjectsByEmail(ctx, &iam_v1.GetSubjectsByEmailRequest{
			Context: iamContext.ToPB(),
			Emails:  []string{email},
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	s, err := subject.New(iamContext, subjectResp.Subjects[0].Subject)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ic *iamClient) GetBySubjectID(ctx context.Context, iamContext *subjectcontext.Context, subjectID string) (subject.Subject, error) {
	var subjectResp *iam_v1.GetSubjectsResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		subjectResp, err = ic.client.GetSubjects(ctx, &iam_v1.GetSubjectsRequest{
			Context:    iamContext.ToPB(),
			SubjectIds: []string{subjectID},
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	s, err := subject.New(iamContext, subjectResp.Subjects[0].Subject)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Register a subject with IAM. Use attribute.NewBuilder() to construct the attributes
func (ic *iamClient) Register(ctx context.Context, iamContext *subjectcontext.Context, email, password string, attributes *iam_v1.StructAttribute) (string, error) {
	var r *iam_v1.RegisterSubjectResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		r, err = ic.client.RegisterSubject(ctx, &iam_v1.RegisterSubjectRequest{
			Context:          iamContext.ToPB(),
			Email:            email,
			Password:         password,
			StructAttributes: attributes,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}

	return r.SubjectId, nil
}

func (ic *iamClient) AddKey(ctx context.Context, iamContext *subjectcontext.Context, email string) (PrivateKey, string, error) {
	var r *iam_v1.AddKeyResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		r, err = ic.client.AddKey(ctx, &iam_v1.AddKeyRequest{
			Context: iamContext.ToPB(),
			Email:   email,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", "", err
	}

	return PrivateKey(r.PrivateKey), r.KeyId, nil
}

func (ic *iamClient) RemoveKey(ctx context.Context, iamContext *subjectcontext.Context, email, keyID string) error {
	return vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		_, err := ic.client.RemoveKey(ctx, &iam_v1.RemoveKeyRequest{
			Context: iamContext.ToPB(),
			Email:   email,
			KeyId:   keyID,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
}

func (ic *iamClient) PartnerBySessionID(ctx context.Context, sessionID string) (*subject.Partner, error) {
	s, err := ic.GetBySessionID(ctx, subjectcontext.New("partner", ""), sessionID)
	if err != nil {
		return nil, err
	}
	return toPartner(s)
}

func (ic *iamClient) PartnerByEmail(ctx context.Context, email string) (*subject.Partner, error) {
	s, err := ic.GetByEmail(ctx, subjectcontext.New("partner", ""), email)
	if err != nil {
		return nil, err
	}
	return toPartner(s)
}

func (ic *iamClient) PartnerBySubjectID(ctx context.Context, subjectID string) (*subject.Partner, error) {
	s, err := ic.GetBySubjectID(ctx, subjectcontext.New("partner", ""), subjectID)
	if err != nil {
		return nil, err
	}
	return toPartner(s)
}

func (ic *iamClient) SalesPersonBySessionID(ctx context.Context, partnerID string, sessionID string) (*subject.SalesPerson, error) {
	s, err := ic.GetBySessionID(ctx, subjectcontext.New("sales_person", partnerID), sessionID)
	if err != nil {
		return nil, err
	}
	return toSalesPerson(s)
}

func (ic *iamClient) SalesPersonByEmail(ctx context.Context, partnerID string, email string) (*subject.SalesPerson, error) {
	s, err := ic.GetByEmail(ctx, subjectcontext.New("sales_person", partnerID), email)
	if err != nil {
		return nil, err
	}
	return toSalesPerson(s)
}

func (ic *iamClient) SalesPersonBySubjectID(ctx context.Context, partnerID string, subjectID string) (*subject.SalesPerson, error) {
	subj, err := ic.GetBySubjectID(ctx, subjectcontext.New("sales_person", partnerID), subjectID)
	if err != nil {
		return nil, err
	}
	return toSalesPerson(subj)
}

func (ic *iamClient) SMBBySessionID(ctx context.Context, partnerID string, sessionID string) (*subject.SMB, error) {
	s, err := ic.GetBySessionID(ctx, subjectcontext.New("smb", partnerID), sessionID)
	if err != nil {
		return nil, err
	}
	return toSMB(s)
}

func (ic *iamClient) SMBByEmail(ctx context.Context, partnerID string, email string) (*subject.SMB, error) {
	s, err := ic.GetByEmail(ctx, subjectcontext.New("smb", partnerID), email)
	if err != nil {
		return nil, err
	}
	return toSMB(s)
}

func (ic *iamClient) SMBBySubjectID(ctx context.Context, partnerID string, subjectID string) (*subject.SMB, error) {
	s, err := ic.GetBySubjectID(ctx, subjectcontext.New("smb", partnerID), subjectID)
	if err != nil {
		return nil, err
	}
	return toSMB(s)
}

func (ic *iamClient) RegisterSMBSubject(ctx context.Context, partnerID, email string, accountGroupIDs []string) (string, error) {
	pass := uuid.New()
	smbAttr := attribute.NewBuilder().Strings("account_group_associations", accountGroupIDs).Build()

	return ic.Register(ctx, subjectcontext.New("smb", partnerID), email, pass, smbAttr)
}

func (ic *iamClient) PartnerAppByEmail(ctx context.Context, partnerID string, email string) (*subject.PartnerApp, error) {
	s, err := ic.GetByEmail(ctx, subjectcontext.New("partner_app", partnerID), email)
	if err != nil {
		return nil, err
	}
	return toPartnerApp(s)
}

func (ic *iamClient) PartnerAppGenerateClientKey(ctx context.Context, partnerID, email string) (*ClientKey, error) {
	return ic.GenerateClientKey(ctx, subjectcontext.New("partner_app", partnerID), email)
}

func (ic *iamClient) RegisterPartnerApp(ctx context.Context, partnerID string, email string) (string, error) {
	//generate a random password
	pass := uuid.New()
	partnerAttr := attribute.NewBuilder().String("partner_id", partnerID).Build()

	return ic.Register(ctx, subjectcontext.New("partner_app", partnerID), email, pass, partnerAttr)
}

func (ic *iamClient) PartnerAppAddKey(ctx context.Context, partnerID, email string) (PrivateKey, string, error) {
	return ic.AddKey(ctx, subjectcontext.New("partner_app", partnerID), email)
}

func (ic *iamClient) PartnerAppRemoveKey(ctx context.Context, partnerID, email, keyID string) error {
	return ic.RemoveKey(ctx, subjectcontext.New("partner_app", partnerID), email, keyID)
}

func (ic *iamClient) ExchangeToken(ctx context.Context, key *ClientKey) (string, error) {
	claims := &clientJWTClaim{
		KeyID:            key.PrivateKeyID,
		SubjectNamespace: key.Namespace,
		SubjectType:      key.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * 5).Unix(), // 5 minutes to allow for clock skew
			Issuer:    key.Namespace,
			Audience:  "vendasta.com",
			Id:        uuid.New(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   key.ClientEmail,
		},
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	block, _ := pem.Decode([]byte(key.PrivateKey))
	x509Encoded := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(x509Encoded)
	if err != nil {
		return "", err
	}

	signedTkn, err := tkn.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	var r *iam_v1.GetTokenResponse
	err = vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		r, err = ic.client.GetSessionToken(ctx, &iam_v1.GetSessionTokenRequest{
			Token: signedTkn,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}
	return r.GetToken(), err
}

func (ic *iamClient) GenerateClientKey(ctx context.Context, iamContext *subjectcontext.Context, email string) (*ClientKey, error) {
	pk, pkID, err := ic.AddKey(ctx, iamContext, email)
	if err != nil {
		return nil, err
	}
	k := &ClientKey{
		Type:                iamContext.Type,
		Namespace:           iamContext.Namespace,
		ClientEmail:         email,
		PrivateKey:          string(pk),
		PrivateKeyID:        pkID,
		AuthProviderCertURL: ic.providerURL,
	}
	return k, nil
}

func (ic *iamClient) GetShortLivedToken(ctx context.Context, iamContext *subjectcontext.Context, email string) (string, error) {
	var tokenResponse *iam_v1.GetTokenResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		tokenResponse, err = ic.client.GetShortLivedToken(ctx, &iam_v1.GetShortLivedTokenRequest{
			Context: iamContext.ToPB(),
			Email:   email,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}
	return tokenResponse.Token, nil
}

func (ic *iamClient) GetSubjectContext(ctx context.Context, subjectID string) (*subjectcontext.Context, error) {
	var subjectContext *iam_v1.GetSubjectContextResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		subjectContext, err = ic.client.GetSubjectContext(ctx, &iam_v1.GetSubjectContextRequest{
			SubjectId: subjectID,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	return subjectcontext.FromProto(subjectContext.Context), nil
}

// ListPersonas list all available subjects for a session
func (ic *iamClient) ListPersonas(ctx context.Context, sessionID, cursor string, pageSize int64, personaType string) ([]subject.Subject, string, bool, error) {
	var list *iam_v1.ListPersonasResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		list, err = ic.userClient.ListPersonas(ctx, &iam_v1.ListPersonasRequest{
			Session:  sessionID,
			Cursor:   cursor,
			PageSize: pageSize,
			Type:     personaType,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, "", false, err
	}

	subjects := []subject.Subject{}
	for _, persona := range list.GetPersonas() {
		s, err := FromPersona(persona)
		if err != nil {
			logging.Errorf(ctx, "Could not create subject from persona (%v): %s", persona, err)
		}
		subjects = append(subjects, s)
	}

	return subjects, list.GetNextCursor(), list.GetHasMore(), nil
}

func (ic *iamClient) RegisterResourceOwner(ctx context.Context, resourceOwner *ResourceOwner) error {
	return vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		_, err := ic.client.RegisterResourceOwner(ctx, &iam_v1.RegisterResourceOwnerRequest{
			Owner: &iam_v1.ResourceOwner{
				AppId:   resourceOwner.AppID,
				AppName: resourceOwner.AppName,
			},
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
}

func (ic *iamClient) RegisterResource(ctx context.Context, resource *Resource) error {
	return vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		_, err := ic.client.RegisterResource(ctx, &iam_v1.RegisterResourceRequest{
			AppId:                   resource.AppID,
			ResourceId:              resource.ResourceID,
			ResourceName:            resource.ResourceName,
			ResourceOwnerServiceUrl: resource.ResourceOwnerServiceURL,
			ResourceOwnerAudience:   resource.ResourceOwnerAudience,
			RequiredResourceParams:  resource.RequiredResourceParams,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

}

func (ic *iamClient) RegisterPolicy(ctx context.Context, policy *Policy) error {
	return vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		_, err := ic.client.RegisterPolicy(ctx, &iam_v1.RegisterPolicyRequest{Policy: PolicyProto(policy)}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

}

func (ic *iamClient) AccessResource(ctx context.Context, accessResource *AccessResource) error {
	return vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		_, err := ic.client.AccessResource(ctx, accessResource.ToPB(), grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

}

func (ic *iamClient) AccessPartnerMarket(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, partnerID string, marketIDs []string, accessScopes ...AccessScope) error {
	ar := &AccessResource{
		Context:           iamContext,
		SubjectIdentifier: subjectIdentifier,
		SessionID:         sessionID,
		ResourceID:        resources.PartnerResourceID,
		OwnerID:           resources.PartnerOwnerID,
		ResourceEntityIdentifier: map[string][]string{
			resources.PartnerResourcePartnerID: {partnerID},
			resources.PartnerResourceMarketIDs: marketIDs,
		},
		AccessScope: accessScopes,
	}
	return ic.AccessResource(ctx, ar)
}

func (ic *iamClient) AccessAccountGroup(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupID string, accessScopes ...AccessScope) error {
	ar := &AccessResource{
		Context:                  iamContext,
		SubjectIdentifier:        subjectIdentifier,
		SessionID:                sessionID,
		ResourceID:               resources.AccountGroupResourceID,
		OwnerID:                  resources.AccountGroupOwnerID,
		ResourceEntityIdentifier: ResourceEntityIdentifier{resources.AccountGroupResourceAccountGroupID: {accountGroupID}},
		AccessScope:              accessScopes,
	}
	return ic.AccessResource(ctx, ar)
}

func (ic *iamClient) AccessAccountGroups(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountGroupIDs []string, accessScopes ...AccessScope) error {
	entityIdentifiers := make([]ResourceEntityIdentifier, len(accountGroupIDs))
	for n, accountGroupID := range accountGroupIDs {
		entityIdentifiers[n] = ResourceEntityIdentifier{resources.AccountGroupResourceAccountGroupID: {accountGroupID}}
	}
	ar := &AccessResource{
		Context:                   iamContext,
		SubjectIdentifier:         subjectIdentifier,
		SessionID:                 sessionID,
		ResourceID:                resources.AccountGroupResourceID,
		OwnerID:                   resources.AccountGroupOwnerID,
		ResourceEntityIdentifiers: entityIdentifiers,
		AccessScope:               accessScopes,
	}
	return ic.AccessResource(ctx, ar)
}

func (ic *iamClient) AccessAccounts(ctx context.Context, iamContext *subjectcontext.Context, subjectIdentifier *SubjectIdentifier, sessionID string, accountIdentifiers []*resources.AccountIdentifier, accessScopes ...AccessScope) error {
	entityIdentifiers := make([]ResourceEntityIdentifier, len(accountIdentifiers))
	for n, accountIdentifier := range accountIdentifiers {
		entityIdentifiers[n] = accountIdentifier.ToResourceIdentifier()
	}
	ar := &AccessResource{
		Context:                   iamContext,
		SubjectIdentifier:         subjectIdentifier,
		SessionID:                 sessionID,
		ResourceID:                resources.AccountResourceID,
		OwnerID:                   resources.AccountOwnerID,
		ResourceEntityIdentifiers: entityIdentifiers,
		AccessScope:               accessScopes,
	}
	return ic.AccessResource(ctx, ar)
}

// PolicyProto converts a Policy to iam_v1.Policy
func PolicyProto(policy *Policy) *iam_v1.Policy {
	return &iam_v1.Policy{
		AppId:      policy.AppID,
		ResourceId: policy.ResourceID,
		PolicyId:   policy.PolicyID,
		PolicyName: policy.PolicyName,
		Operations: OperationsProto(policy.Operations),
		Policy:     PolicyNodeProto(policy.Policy),
	}
}

// OperationsProto converts a []AccessScope to []iam_v1.AccessScope
func OperationsProto(scopes []AccessScope) []iam_v1.AccessScope {
	scopesProto := []iam_v1.AccessScope{}
	for _, scope := range scopes {
		scopesProto = append(scopesProto, iam_v1.AccessScope(scope))
	}
	return scopesProto
}

// PolicyNodeProto converts a PolicyNode to a iam_v1.PolicyNode
func PolicyNodeProto(policyNode *PolicyNode) *iam_v1.PolicyNode {
	policyNodeProto := &iam_v1.PolicyNode{}
	attrClause := policyNode.GetSubjectResourceIntersectionClause()
	attrMappedClause := policyNode.GetSubjectResourceMappedIntersectionClause()
	subsetClause := policyNode.GetSubjectResourceSubsetClause()
	subjectValueIntersectionClause := policyNode.GetSubjectValueIntersection()
	subjectMissingValueClause := policyNode.GetSubjectMissingValueClause()
	subjectResourceForClause := policyNode.GetSubjectResourceForClause()
	operator := policyNode.GetOperator()
	if attrClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectResourceIntersection{
			SubjectResourceIntersection: &iam_v1.SubjectResourceIntersectionClause{
				AttributeName: attrClause.AttributeName,
			},
		}
	} else if subsetClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectResourceSubset{
			SubjectResourceSubset: &iam_v1.SubjectResourceSubsetClause{
				AttributeName: subsetClause.AttributeName,
			},
		}
	} else if subjectValueIntersectionClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectValueIntersection{
			SubjectValueIntersection: &iam_v1.SubjectValueIntersectionClause{
				AttributeName:            subjectValueIntersectionClause.AttributeName,
				StructuredAttributeValue: subjectValueIntersectionClause.AttributeValue,
			},
		}
	} else if operator != nil {
		children := []*iam_v1.PolicyNode{}
		for _, child := range operator.Children {
			children = append(children, PolicyNodeProto(child))
		}
		policyNodeProto.Value = &iam_v1.PolicyNode_Operator{
			Operator: &iam_v1.Operator{
				Operator: iam_v1.BooleanOperator(operator.Operator),
				Children: children,
			},
		}
	} else if subjectMissingValueClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectMissingValue{
			SubjectMissingValue: &iam_v1.SubjectMissingValueClause{
				AttributeName: subjectMissingValueClause.AttributeName,
			},
		}
	} else if attrMappedClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectResourceIntersection{
			SubjectResourceIntersection: &iam_v1.SubjectResourceIntersectionClause{
				AttributeName:         attrMappedClause.SubjectAttributeName,
				ResourceAttributeName: attrMappedClause.ResourceAttributeName,
			},
		}
	} else if subjectResourceForClause != nil {
		policyNodeProto.Value = &iam_v1.PolicyNode_SubjectResourceFor{
			SubjectResourceFor: &iam_v1.SubjectResourceForClause{
				AttributeName: subjectResourceForClause.AttributeName,
				Operator:      iam_v1.ForOperator(subjectResourceForClause.Operator),
				Rules:         PolicyNodeProto(subjectResourceForClause.Policy),
			},
		}
	}
	return policyNodeProto
}

func toPartner(s subject.Subject) (*subject.Partner, error) {
	partner, ok := s.(*subject.Partner)
	if !ok {
		return nil, ErrInvalidSubjectType
	}
	return partner, nil
}

func toSalesPerson(s subject.Subject) (*subject.SalesPerson, error) {
	salesPerson, ok := s.(*subject.SalesPerson)
	if !ok {
		return nil, ErrInvalidSubjectType
	}
	return salesPerson, nil
}

func toSMB(s subject.Subject) (*subject.SMB, error) {
	smb, ok := s.(*subject.SMB)
	if !ok {
		return nil, ErrInvalidSubjectType
	}
	return smb, nil
}

func toPartnerApp(s subject.Subject) (*subject.PartnerApp, error) {
	partnerApp, ok := s.(*subject.PartnerApp)
	if !ok {
		return nil, ErrInvalidSubjectType
	}
	return partnerApp, nil
}

var defaultRetryCallOptions = vax.WithRetry(func() vax.Retryer {
	return vax.OnCodes([]codes.Code{
		codes.DeadlineExceeded,
		codes.Unavailable,
		codes.Unknown,
	}, vax.Backoff{
		Initial:    10 * time.Millisecond,
		Max:        300 * time.Millisecond,
		Multiplier: 3,
	})
})
