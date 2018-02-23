package accounts

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/pb/accounts/v1"
	"github.com/vendasta/gosdks/util"
)

// Account represents an account entity from Accounts microservice
type Account struct {
	AccountID  string
	BusinessID string
	PartnerID  string

	ProductID             string
	MarketplaceAppID      string
	OrderFormSubmissionID string

	CustomEntryURL string

	Expiry time.Time
	Tags   []string
}

// MarketplaceOrProductID returns either the legacy product or marketplace app id
func (a *Account) MarketplaceOrProductID() string {
	if a.MarketplaceAppID != "" {
		return a.MarketplaceAppID
	}
	return a.ProductID
}

// IsTrial returns a bool if account is trial.
func (a *Account) IsTrial() bool {
	return util.StringInSlice("trial", a.Tags)
}

func accountFromPB(a *accounts_v1.Account) (*Account, error) {
	account := &Account{
		AccountID:             a.GetAccountId(),
		BusinessID:            a.GetBusinessId(),
		PartnerID:             a.GetPartnerId(),
		ProductID:             a.GetProductId(),
		MarketplaceAppID:      a.GetAppId(),
		CustomEntryURL:        a.GetCustomEntryUrl(),
		OrderFormSubmissionID: a.GetOrderFormSubmissionId(),
	}

	if a.GetTrial() == true {
		account.Tags = append(account.Tags, "trial")
	}
	deactivation := a.GetDeactivation()
	if deactivation != nil {
		var err error
		account.Expiry, err = ptypes.Timestamp(deactivation)
		if err != nil {
			return nil, err
		}
	}
	return account, nil
}

type Status int32

const (
	PENDING  Status = 0
	APPROVED Status = 1
	REJECTED Status = 2
)

// PendingActivation is a record of pending activations
type PendingActivation struct {
	// The unique identifier for an activation and pending activation
	ActivationID string
	// The business identifier that this pending activation is for
	BusinessID string
	// The status of the pending activation
	Status Status
	// The reason for a rejected pending activation
	RejectedReason string
	// Who dismissed the pending activation
	DismissedBy string
	// Date the pending activation was created
	CreatedDate time.Time
	// Date the pending activation was last updated
	UpdatedDate time.Time
	// The app activation for this pending activation
	AppActivation *Account
	// The addon activation for this pending activation
	AddonActivation *AddonActivation
	// Identifier for the app/product the pending activation is for
	AppID string
}

func PendingActivationFromPB(pa *accounts_v1.PendingActivation) (*PendingActivation, error) {
	pendingActivation := &PendingActivation{
		ActivationID:   pa.GetActivationId(),
		BusinessID:     pa.GetBusinessId(),
		RejectedReason: pa.GetRejectedReason(),
		DismissedBy:    pa.GetDismissedBy(),
		Status:         Status(pa.GetStatus()),
		AppID:          pa.GetAppId(),
	}

	created := pa.GetCreatedDate()
	if created != nil {
		var err error
		pendingActivation.CreatedDate, err = ptypes.Timestamp(created)
		if err != nil {
			return nil, err
		}
	}
	updated := pa.GetUpdatedDate()
	if updated != nil {
		var err error
		pendingActivation.UpdatedDate, err = ptypes.Timestamp(updated)
		if err != nil {
			return nil, err
		}
	}

	appActivation, err := accountFromPB(pa.GetAppActivation())
	if err != nil {
		return nil, err
	}
	pendingActivation.AppActivation = appActivation

	addonActivation, err := AddonActivationFromPB(pa.GetAddonActivation())
	if err != nil {
		return nil, err
	}
	pendingActivation.AddonActivation = addonActivation

	return pendingActivation, nil
}

// AddonActivation is a record of activating an addon
type AddonActivation struct {
	// A prerequisite ID representing the customer/business.
	BusinessID string
	// A prerequisite marketplace vendor's ID's of the app the addons belong to.
	AppID string
	// An ID assigned uniquely to this addon upon being activated.
	ActivationID string
	// A prerequisite marketplace vendor's ID's of the addon activated.
	AddonID string
	// UTC time the addon was activated.
	Activation time.Time
	// UTC time the addon was or will be deactivated, if ever.
	Deactivation time.Time
	// OrderFormSubmissionID is the ID of the order form submission used for this activation
	OrderFormSubmissionID string
}

func AddonActivationFromPB(aa *accounts_v1.AddonActivation) (*AddonActivation, error) {
	addonActivation := &AddonActivation{
		BusinessID:            aa.GetBusinessId(),
		AppID:                 aa.GetAppId(),
		ActivationID:          aa.GetActivationId(),
		AddonID:               aa.GetAddonId(),
		OrderFormSubmissionID: aa.GetOrderFormSubmissionId(),
	}

	activatedTime := aa.GetActivated()
	if activatedTime != nil {
		var err error
		addonActivation.Activation, err = ptypes.Timestamp(activatedTime)
		if err != nil {
			return nil, err
		}
	}
	deactivatedTime := aa.GetDeactivated()
	if deactivatedTime != nil {
		var err error
		addonActivation.Deactivation, err = ptypes.Timestamp(deactivatedTime)
		if err != nil {
			return nil, err
		}
	}
	return addonActivation, nil
}

// ActivateAppResponse is the response of product activation
type ActivateAppResponse struct {
	// A product internal unique ID that corresponds to the business.
	AccountId string
	// A unique ID assigned for this specific activation.
	ActivationId string
}
