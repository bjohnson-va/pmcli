package accounts

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/pb/accounts/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

// Interface to accounts microservice API
type Interface interface {
	ListAddonActivations(ctx context.Context, businessID string, appID string) ([]*AddonActivation, error)
	List(ctx context.Context, businessID string, partnerID string) ([]*Account, error)
	ActivateApp(ctx context.Context, businessID string, appID string, partnerID string, orderFormSubmissionID string,
		activateOn time.Time, deactivateOn time.Time, trial bool) (*ActivateAppResponse, error)
	ResolvePendingActivation(ctx context.Context, activationID string, approved bool, rejectedReason string,
		accountID string, businessID string, appID string, addonID string) error
	DismissPendingActivation(ctx context.Context, activationID string, dismissedBy string, businessID string,
		appID string, addonID string) error
}

// NewClient returns a new accounts API client object
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	connection, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &client{accounts_v1.NewAccountsServiceClient(connection)}, nil
}

type client struct {
	accounts_v1.AccountsServiceClient
}

func (c *client) ListAddonActivations(ctx context.Context, businessID string, appID string) ([]*AddonActivation, error) {
	var resp *accounts_v1.ListAddonActivationsResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = c.AccountsServiceClient.ListAddonActivations(ctx, &accounts_v1.ListAddonActivationsRequest{
			BusinessId: businessID, AppId: appID})
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	addons := make([]*AddonActivation, len(resp.Activations))
	for i, v := range resp.Activations {
		addonActivation, err := AddonActivationFromPB(v)
		if err != nil {
			return nil, err
		}
		addons[i] = addonActivation
	}
	return addons, nil
}

// List return list of accounts of a business
func (c *client) List(ctx context.Context, businessID string, partnerID string) ([]*Account, error) {
	var resp *accounts_v1.ListResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		resp, err = c.AccountsServiceClient.List(ctx, &accounts_v1.ListRequest{
			BusinessId: businessID,
			PartnerId:  partnerID,
		})
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(resp.Accounts))
	for i, v := range resp.Accounts {
		account, err := accountFromPB(v)
		if err != nil {
			return nil, err
		}
		accounts[i] = account
	}
	return accounts, nil
}

// ResolvePendingActivation will either approve or reject a PendingActivation
func (c *client) ResolvePendingActivation(ctx context.Context, activationID string, approved bool, rejectedReason string, accountID string, businessID string, appID string, addonID string) error {
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		request := &accounts_v1.ResolvePendingActivationRequest{
			ActivationId:   activationID,
			Approved:       approved,
			RejectedReason: rejectedReason,
			AccountId:      accountID,
			BusinessId:     businessID,
			AppId:          appID,
			AddonId:        addonID,
		}
		_, err = c.AccountsServiceClient.ResolvePendingActivation(ctx, request)

		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return err
	}
	return nil
}

// DismissPendingActivation will set the `dismissed` flag on a PendingActivation
func (c *client) DismissPendingActivation(ctx context.Context, activationID string, dismissedBy string, businessID string, appID string, addonID string) error {
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		request := &accounts_v1.DismissPendingActivationRequest{
			ActivationId: activationID,
			DismissedBy:  dismissedBy,
			BusinessId:   businessID,
			AppId:        appID,
			AddonId:      addonID,
		}
		_, err = c.AccountsServiceClient.DismissPendingActivation(ctx, request)

		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return err
	}
	return nil
}

// ActivateApp activates an app/product
func (c *client) ActivateApp(ctx context.Context, businessID string, appID string, partnerID string, orderFormSubmissionID string, activateOn time.Time, deactivateOn time.Time, trial bool) (*ActivateAppResponse, error) {
	var resp *accounts_v1.ActivateAppResponse
	err := vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		request := &accounts_v1.ActivateAppRequest{
			BusinessId:            businessID,
			AppId:                 appID,
			PartnerId:             partnerID,
			OrderFormSubmissionId: orderFormSubmissionID,
			Trial: trial,
		}

		if !activateOn.IsZero() {
			activateOnTimestamp, err := ptypes.TimestampProto(activateOn)
			if err != nil {
				return err
			}
			request.ActivateOn = activateOnTimestamp
		}

		if !deactivateOn.IsZero() {
			deactivateOnTimestamp, err := ptypes.TimestampProto(deactivateOn)
			if err != nil {
				return err
			}
			request.DeactivateOn = deactivateOnTimestamp
		}

		resp, err = c.AccountsServiceClient.ActivateApp(ctx, request)
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return &ActivateAppResponse{AccountId: resp.AccountId, ActivationId: resp.ActivationId}, nil
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
