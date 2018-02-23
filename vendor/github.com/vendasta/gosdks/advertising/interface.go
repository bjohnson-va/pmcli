package advertising

import (
	"golang.org/x/net/context"
)

// Client to advertising microservice API
type Client interface {
	CampaignEventClient
	AdwordsClient
}

type CampaignEventClient interface {
	CreateOrderEvent(ctx context.Context, order *OrderEvent) error
}

type AdwordsClient interface {
	StoreCredentialsForBusiness(ctx context.Context, oauthRefreshToken string, businessID string) error
}
