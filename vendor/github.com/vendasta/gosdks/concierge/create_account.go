package concierge

import (
	"context"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
)

const (
	createAccountPath = "/api/v2/account/provision/"
)

// CreateAccount creates a concierge account
// returns a basesdk.VAPIError on a failing request.
func (c *conciergeClient) CreateAccount(ctx context.Context, partnerID string, businessID string) (error) {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(partnerID, util.InvalidArgument, "partnerID is required"),
		validation.StringNotEmpty(businessID, util.InvalidArgument, "businessID is required"),
	).Validate()
	if err != nil {
		return err
	}

	params := map[string]interface{}{}
	params["partnerId"] = partnerID
	params["accountGroupId"] = businessID
	_, err = c.SDKClient.Post(ctx, createAccountPath, params, basesdk.Idempotent())
	if err != nil {
		return err
	}
	return nil
}
