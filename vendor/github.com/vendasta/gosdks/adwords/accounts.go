package adwords

import (
	"github.com/vendasta/gosdks/pb/adwords_service/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ListAllAccounts lists all accounts associated with the given oauth token
func (c *client) ListAllAccounts(
	ctx context.Context,
	oauthRefreshToken string,
) ([]Account, error) {
	ctx = metadata.NewOutgoingContext(ctx, nil)

	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(oauthRefreshToken, util.InvalidArgument, "oauthRefreshToken is required"),
	).ValidateAndJoinErrors()
	if err != nil {
		return nil, err
	}

	request := &adwords_v1.ListAllAccountsRequest{OauthRefreshToken: oauthRefreshToken}
	var response *adwords_v1.ListAllAccountsResponse
	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		response, err = c.AccountsClient.ListAllAccounts(ctx, request, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(response.Accounts))
	for i, acc := range response.Accounts {
		accounts[i] = fromAccountProto(acc)
	}

	return accounts, nil
}

type Account struct {
	// The Account's ID; i.e. CID
	CustomerId int64
	// Display name for the account
	Name string
	// Is this a manager account (i.e. MCC)?
	IsManager bool
	// Is this a test account?
	IsTestAccount bool
	// The ISO 4217 currency code of the account
	CurrencyCode string
	// List of timezones here: https://developers.google.com/adwords/api/docs/appendix/codes-formats#timezone-ids
	TimeZone string
}

func fromAccountProto(account *adwords_v1.Account) Account {
	return Account{
		CustomerId:    account.CustomerId,
		Name:          account.Name,
		CurrencyCode:  account.CurrencyCode,
		IsManager:     account.IsManager,
		IsTestAccount: account.IsTestAccount,
		TimeZone:      account.TimeZone,
	}
}
