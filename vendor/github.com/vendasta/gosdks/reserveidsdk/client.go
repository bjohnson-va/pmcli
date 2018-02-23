package reserveidsdk

import (
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

const (
	reserveIDPath = "/api/v1/webhook/reserveId/"
)

// Reserve ID client that handles the calls to VDC to reserve an vendor internal ID for product activation
type ReserveIDClient struct {
	basesdk.SDKClient
}

//Interface for reserve ID service
type ReserveIDInterface interface {
	ReserveID(ctx context.Context, partnerID string, businessID string, customerID string, appID string) (*ReserveIDResponse, error)
}

// Build the client to call building order form
func BuildReserveIDClient(apiUser string, apiKey string, env config.Env) *ReserveIDClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &ReserveIDClient{baseClient}
}

// ReserveID calls the reserve ID endpoint in VDC to reserve an vendor internal ID
func (c *ReserveIDClient) ReserveID(ctx context.Context, partnerID string, businessID string, customerID string, appID string) (*ReserveIDResponse, error) {
	err := validation.NewValidator().
		Rule(validation.StringNotEmpty(partnerID, util.InvalidArgument, "partner id is required")).
		Rule(validation.StringNotEmpty(businessID, util.InvalidArgument, "business id is required")).
		Rule(validation.StringNotEmpty(appID, util.InvalidArgument, "product/addon id is required")).
		Validate()
	if err != nil {
		return nil, util.ToGrpcError(err)
	}

	params := map[string]interface{}{
		"partner_id":  partnerID,
		"business_id": businessID,
		"customer_id": customerID,
		"app_id":      appID,
	}

	response, err := c.Post(ctx, reserveIDPath, params)
	if err != nil {
		return nil, basesdk.ConvertHttpErrorToGRPC(err)
	}
	defer response.Body.Close()
	reserveIDResponse, err := reserveIDResponseFromResponse(response)
	if err != nil {
		return nil, util.ToGrpcError(err)
	}
	return reserveIDResponse, nil
}
