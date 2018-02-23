package domain

import (
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"github.com/vendasta/gosdks/pb/domain/v1"
	"fmt"
)

// Domain provides name of the domain and whether it supports SSL
type Domain struct {
	Name   string
	Secure bool
}

// DomainMapping provides the domains registered for an Identifier.
type DomainMapping struct {
	Primary     *Domain
	Secondaries []*Domain
}

// LegacyProductPartner returns the Identifier for a legacy product partner domain
func LegacyProductPartner(productID, partnerID string) Identifier {
	return Identifier(fmt.Sprintf("/application/%s/partner/%s", productID, partnerID))
}

// VBCPartner returns the Identifier for a VBC partner domain
func VBCPartner(partnerID string) Identifier {
	return LegacyProductPartner("VBC", partnerID)
}

// SocialMarketingPartner returns the Identifier for a SM partner domain
func SocialMarketingPartner(partnerID string) Identifier {
	return LegacyProductPartner("SM", partnerID)
}

// ReputationIntelligencePartner returns the Identifier for a RM partner domain
func ReputationIntelligencePartner(partnerID string) Identifier {
	return LegacyProductPartner("RM", partnerID)
}

// ListingBuilderPartner returns the Identifier for a MS partner domain
func ListingBuilderPartner(partnerID string) Identifier {
	return LegacyProductPartner("MS", partnerID)
}

// Identifier is a contextual id that holds hierarchical data.
//
// Examples:
//
// Identifier{contextual_identifier: "/application/partner-center"}
// Identifier{contextual_identifier: "/product/RM/partner/ABC"}
// Identifier{contextual_identifier: "/product/SM/partner/DEF/market/my-market"}
// Identifier{contextual_identifier: "/marketplace/website-pro/account/AG-J7V5H8AV/website/49da31ebc3f34f6c97dd540e2447dca7"}
//
// These hierarchical identifiers allow for flexible re-use, as well as scanning across each contextual piece.
// For example, you could scan across all domains for `/product/RM`, which returns all custom domains for Reputation Intelligence.
// You could also scan across `/product/SM/partner/DEF` and get all of the custom market domains for the DEF partner in SM.
type Identifier string

var addresses = map[config.Env]string{
	config.Local: "domain:11000",
	config.Test:  "domain-api-test.vendasta-internal.com:443",
	config.Demo:  "domain-api-demo.vendasta-internal.com:443",
	config.Prod:  "domain-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "",
	config.Test:  "https://domain-api-test.vendasta-internal.com",
	config.Demo:  "https://domain-api-demo.vendasta-internal.com",
	config.Prod:  "https://domain-api-prod.vendasta-internal.com",
}

// NewClient returns a Domain client.
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	scope := scopes[e]
	conn, err := vax.NewGRPCConnection(ctx, address, true, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &domainClient{
		client: domain_v1.NewDomainServiceClient(conn),
	}, nil
}

// Interface defines all of the API methods available from Domain.
type Interface interface {
	GetDomainByIdentifier(ctx context.Context, identifier Identifier) (*DomainMapping, error)
	GetIdentifierByDomain(ctx context.Context, domain string) (Identifier, error)
}

type domainClient struct {
	client domain_v1.DomainServiceClient
}

// GetDomainByIdentifier returns the domain by the given Identifier
func (ic *domainClient) GetDomainByIdentifier(ctx context.Context, identifier Identifier) (*DomainMapping, error) {
	var resp *domain_v1.GetDomainByIdentifierResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		resp, err = ic.client.GetDomainByIdentifier(ctx, &domain_v1.GetDomainByIdentifierRequest{
			Identifier: &domain_v1.Identifier{
				string(identifier),
			},
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}
	return &DomainMapping{
		Primary:     domainFromProto(resp.Primary),
		Secondaries: domainsFromProto(resp.Secondaries),
	}, nil
}

// GetIdentifierByDomain returns the Identifier by the given domain
func (ic *domainClient) GetIdentifierByDomain(ctx context.Context, domain string) (Identifier, error) {
	var resp *domain_v1.GetIdentifierByDomainResponse
	err := vax.Invoke(ctx, func(context.Context, vax.CallSettings) error {
		var err error
		resp, err = ic.client.GetIdentifierByDomain(ctx, &domain_v1.GetIdentifierByDomainRequest{
			Domain: domain,
		}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", err
	}
	return Identifier(resp.Identifier.GetContextualIdentifier()), nil
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

func domainFromProto(d *domain_v1.Domain) *Domain {
	return &Domain{
		Name:   d.Domain,
		Secure: d.Secure,
	}
}

func domainsFromProto(domainsProto []*domain_v1.Domain) []*Domain {
	if len(domainsProto) == 0 {
		return nil
	}
	domains := make([]*Domain, len(domainsProto))
	for n, d := range domainsProto {
		domains[n] = domainFromProto(d)
	}
	return domains
}
