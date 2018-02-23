package advertising

import (
	"github.com/vendasta/gosdks/pb/advertising/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
)

// StoreCredentialsForBusiness associates the given refresh token with a business
func (c *client) StoreCredentialsForBusiness(ctx context.Context, oauthRefreshToken string, businessID string) error {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(oauthRefreshToken, util.InvalidArgument, "oauthRefreshToken is required"),
		validation.StringNotEmpty(businessID, util.InvalidArgument, "businessID is required"),
	).ValidateAndJoinErrors()
	if err != nil {
		return err
	}

	req := &advertising_v1.StoreCredentialsForBusinessRequest{
		BusinessId:        businessID,
		OauthRefreshToken: oauthRefreshToken,
	}

	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		_, err = c.adwordsClient.StoreCredentialsForBusiness(ctx, req)
		return err
	}, defaultRetryCallOptions)
	return err
}
