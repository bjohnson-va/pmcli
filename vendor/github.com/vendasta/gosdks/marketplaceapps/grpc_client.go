package marketplaceapps

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/pb/marketplace_apps/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GRPCClient a GRPC client for the marketplace apps microservice.
type GRPCClient struct {
	client marketplaceapps_v1.MarketplaceAppsClient
}

// NewGRPCClient returns an new GRPC client for the marketplace apps microservice.
func NewGRPCClient(ctx context.Context, address string, useTLS bool, scope string, dialOptions ...grpc.DialOption) (Client, error) {
	conn, err := vax.NewGRPCConnection(ctx, address, useTLS, scope, true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{client: marketplaceapps_v1.NewMarketplaceAppsClient(conn)}, nil
}

// defaultRetryCallOptions controls the errors that we will automatically retry on. This is due to the case where the
// server has given us an error that is deemed retry-able.
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

// GetPublicKeys get's the public keys for a marketplace application.
func (c *GRPCClient) GetPublicKeys(ctx context.Context, appID string) (*PublicKeys, error) {
	ctx = util.NewContext(ctx)

	var res *marketplaceapps_v1.GetPublicKeysResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetPublicKeys(ctx, &marketplaceapps_v1.GetPublicKeysRequest{AppId: appID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	ipk := make([]*IdentifiedPublicKey, len(res.IdentifiedPublicKeys))
	for i := range res.IdentifiedPublicKeys {
		ipk[i] = &IdentifiedPublicKey{
			PublicKeyID: res.IdentifiedPublicKeys[i].PublicKeyId,
			PublicKey:   res.IdentifiedPublicKeys[i].PublicKey,
		}
	}

	return &PublicKeys{
		PublicKey:            res.PublicKey,
		IdentifiedPublicKeys: ipk,
	}, nil
}

// GetSSOUrls get's the SSO Urls (entry and session URL's) for a marketplace application.
func (c *GRPCClient) GetSSOUrls(ctx context.Context, appID string) (entryURL string, sessionURL string, err error) {
	ctx = util.NewContext(ctx)
	var res *marketplaceapps_v1.GetSSOUrlsResponse
	err = vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetSSOUrls(ctx, &marketplaceapps_v1.GetSSOUrlsRequest{AppId: appID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return "", "", err
	}
	return res.EntryUrl, res.SessionUrl, nil
}

// GetAddon get's a specific addon for the specified app
func (c *GRPCClient) GetAddon(ctx context.Context, appID string, addonID string) (*Addon, error) {
	ctx = util.NewContext(ctx)

	var res *marketplaceapps_v1.AddonResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetAddon(ctx, &marketplaceapps_v1.GetAddonRequest{AppId: appID, AddonId: addonID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	addon, err := buildAddon(res.Addon)
	if err != nil {
		return nil, err
	}

	return addon, nil
}

func (c *GRPCClient) ListAddons(ctx context.Context, appID string, cursor string, pageSize int64) ([]*Addon, string, bool, error) {
	ctx = util.NewContext(ctx)

	var res *marketplaceapps_v1.AddonsListResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.ListAddons(ctx, &marketplaceapps_v1.ListAddonsRequest{AppId: appID, Cursor: cursor, PageSize: pageSize}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, "", false, err
	}

	var addonList []*Addon
	for _, addon := range res.Addons {
		a, err := buildAddon(addon)
		if err != nil {
			return nil, "", false, err
		}
		addonList = append(addonList, a)
	}

	return addonList, res.NextCursor, res.HasMore, nil
}

func (c *GRPCClient) ListApprovedAddons(ctx context.Context, appID string, cursor string, pageSize int64, appIDs []string) ([]*Addon, string, bool, error) {
	ctx = util.NewContext(ctx)

	var res *marketplaceapps_v1.AddonsListResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		// Increase timeout for ListApprovedAddons as it goes to elastic which can have some large latency spikes
		ctx, _ = context.WithTimeout(ctx, 10*time.Second)
		res, err = c.client.ListApprovedAddons(ctx, &marketplaceapps_v1.ListApprovedAddonsRequest{AppId: appID, Cursor: cursor, PageSize: pageSize, AppIds: appIDs}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, "", false, err
	}

	var addonList []*Addon
	for _, addon := range res.Addons {
		a, err := buildAddon(addon)
		if err != nil {
			return nil, "", false, err
		}
		addonList = append(addonList, a)
	}

	return addonList, res.NextCursor, res.HasMore, nil
}

func (c *GRPCClient) GetApp(ctx context.Context, appID string) (*App, error) {
	ctx = util.NewContext(ctx)
	var res *marketplaceapps_v1.GetAppResponse
	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, _ = context.WithTimeout(ctx, 3*time.Second)
		res, err = c.client.GetApp(ctx, &marketplaceapps_v1.GetAppRequest{AppId: appID}, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)
	if err != nil {
		return nil, err
	}

	created, err := ptypes.Timestamp(res.App.Created)
	if err != nil {
		return nil, err
	}

	updated, err := ptypes.Timestamp(res.App.Updated)
	if err != nil {
		return nil, err
	}

	return &App{
		AppId:             res.App.AppId,
		Version:           res.App.Version,
		Name:              res.App.Name,
		Icon:              res.App.Icon,
		AllowWhitelabel:   res.App.AllowWhitelabel,
		InDevelopment:     res.App.InDevelopment,
		TrialSupport:      res.App.TrialSupport,
		UsesOrderForm:     res.App.UsesOrderForm,
		Tagline:           res.App.Tagline,
		SignInUrl:         res.App.SignInUrl,
		Currency:          res.App.Currency,
		Price:             res.App.Price,
		BillingFrequency:  res.App.BillingFrequency,
		ServiceModel:      res.App.ServiceModel,
		Screenshot:        res.App.Screenshot,
		WebsiteUrl:        res.App.WebsiteUrl,
		HeaderImage:       res.App.HeaderImage,
		ProductCategory:   ProductCategory(res.App.ProductCategory),
		LmiCategories:     convertLmiCategories(res.App.LmiCategories),
		Integration:       buildIntegrationData(res.App.Integration),
		OrderForm:         buildOrderFormData(res.App.OrderForm),
		ResellerMarketing: buildMarketingInformation(res.App.ResellerMarketing),
		EndUserMarketing:  buildMarketingInformation(res.App.EndUserMarketing),
		Created:           &created,
		Updated:           &updated,
		Restrictions:      buildRestrictionsData(res.App.Restrictions),
	}, nil

}

func convertLmiCategories(lmiCat []marketplaceapps_v1.MarketplaceApp_Lmi_LmiCategories) []MarketplaceApp_Lmi_LmiCategories {
	lmiCategories := make([]MarketplaceApp_Lmi_LmiCategories, len(lmiCat))
	for i, v := range lmiCat {
		lmiCategories[i] = MarketplaceApp_Lmi_LmiCategories(v)
	}
	return lmiCategories
}

func buildIntegrationData(integrationData *marketplaceapps_v1.Integration) *Integration {
	publicKeys := make([]*IdentifiedPublicKey, len(integrationData.IdentifiedPublicKeys))
	for i, v := range integrationData.IdentifiedPublicKeys {
		publicKeys[i] = &IdentifiedPublicKey{
			PublicKeyID: v.PublicKeyId,
			PublicKey:   v.PublicKey,
		}
	}

	return &Integration{
		PublicKey:               integrationData.PublicKey,
		IdentifiedPublicKeys:    publicKeys,
		ContactEmail:            integrationData.ContactEmail,
		Approved:                integrationData.Approved,
		Vendor:                  integrationData.Vendor,
		UserIds:                 integrationData.UserIds,
		PurchaseWebhookUrl:      integrationData.PurchaseWebhookUrl,
		AddonPurchaseWebhookUrl: integrationData.AddonPurchaseWebhookUrl,
		UserModWebhookUrl:       integrationData.UserModWebhookUrl,
		SalestoolWebhookUrl:     integrationData.SalestoolWebhookUrl,
		AccountModWebhookUrl:    integrationData.AccountModWebhookUrl,
		LogoutWebhookUrl:        integrationData.LogoutWebhookUrl,
		FulfillmentEmail:        integrationData.FulfillmentEmail,
		EntryUrl:                integrationData.EntryUrl,
		SessionUrl:              integrationData.SessionUrl,
		SettingsUrl:             integrationData.SettingsUrl,
		ReserveIdUrl:            integrationData.ReserveIdUrl,
	}

}

func buildRestrictionsData(restrictions *marketplaceapps_v1.Restrictions) *Restrictions {
	var whiteList []string
	var blackList []string

	if restrictions != nil && restrictions.Country != nil {
		whiteList = restrictions.Country.Whitelist
		blackList = restrictions.Country.Blacklist
	}

	permissionLists := PermissionLists{
		WhiteList: whiteList,
		BlackList: blackList,
	}

	return &Restrictions{
		Country: &permissionLists,
	}
}

func buildOrderFormData(form *marketplaceapps_v1.OrderForm) *OrderForm {
	orderFormFields := []*OrderFormField{}
	includeCommonFormFields := IncludedCommonFormFields{}
	activationMessage := ""

	if form.CommonForm != nil {
		includeCommonFormFields = IncludedCommonFormFields{
			BusinessName:           form.CommonForm.BusinessName,
			BusinessAddress:        form.CommonForm.BusinessAddress,
			BusinessPhoneNumber:    form.CommonForm.BusinessPhoneNumber,
			BusinessAccountGroupId: form.CommonForm.BusinessAccountGroupId,
			SalespersonName:        form.CommonForm.SalespersonName,
			SalespersonPhoneNumber: form.CommonForm.SalespersonPhoneNumber,
			SalespersonEmail:       form.CommonForm.SalespersonEmail,
			ContactName:            form.CommonForm.ContactName,
			ContactPhoneNumber:     form.CommonForm.ContactPhoneNumber,
			ContactEmail:           form.CommonForm.ContactEmail,
		}
	}

	if form.OrderForm != nil {
		orderFormFields = make([]*OrderFormField, len(form.OrderForm))
		for i, v := range form.OrderForm {
			var f OrderFormField
			f.Description = v.Description
			f.Label = v.Label
			f.Id = v.Id
			f.Type = v.Type
			f.Options = v.Options
			f.Required = v.Required
			f.Description = v.Description
			f.UploadUrl = v.UploadUrl
			orderFormFields[i] = &f
		}
	}

	if form.ActivationMessage != "" {
		activationMessage = form.ActivationMessage
	}

	return &OrderForm{
		OrderForm:         orderFormFields,
		CommonForm:        &includeCommonFormFields,
		ActivationMessage: activationMessage,
	}
}

func buildMarketingInformation(marketInfo *marketplaceapps_v1.MarketingInformation) *MarketingInformation {
	faqs := []*FrequentlyAskedQuestions{}
	if marketInfo.Faqs != nil {
		faqs = make([]*FrequentlyAskedQuestions, len(marketInfo.Faqs))
		for i, q := range marketInfo.Faqs {
			var f FrequentlyAskedQuestions
			f.Answer = q.Answer
			f.Question = q.Question
			faqs[i] = &f
		}
	}

	return &MarketingInformation{
		Description:      marketInfo.Description,
		KeySellingPoints: marketInfo.KeySellingPoints,
		Faqs:             faqs,
		Files:            marketInfo.Files,
	}

}

func buildAddon(addon *marketplaceapps_v1.Addon) (*Addon, error) {
	created, err := ptypes.Timestamp(addon.Created)
	if err != nil {
		return nil, err
	}

	updated, err := ptypes.Timestamp(addon.Updated)
	if err != nil {
		return nil, err
	}

	//build order form structure
	orderForm := &OrderForm{
		OrderForm:         []*OrderFormField{},
		CommonForm:        &IncludedCommonFormFields{},
		ActivationMessage: "",
	}

	if addon.OrderForm != nil {
		orderForm = buildOrderFormData(addon.OrderForm)
	}

	// build reseller marketing structure
	resellerMarketing := &MarketingInformation{
		Description:      "",
		KeySellingPoints: []string{},
		Faqs:             []*FrequentlyAskedQuestions{},
		Files:            []string{},
	}

	if addon.ResellerMarketing != nil {
		resellerMarketing = buildMarketingInformation(addon.ResellerMarketing)
	}

	//build end user marketing structure
	enduserMarketing := &MarketingInformation{
		Description:      "",
		KeySellingPoints: []string{},
		Faqs:             []*FrequentlyAskedQuestions{},
		Files:            []string{},
	}

	if addon.EndUserMarketing != nil {
		enduserMarketing = buildMarketingInformation(addon.EndUserMarketing)
	}

	return &Addon{
		AppID:               addon.AppId,
		AddonID:             addon.AddonId,
		Price:               addon.Price,
		Title:               addon.Title,
		Approved:            addon.Approved,
		Discoverable:        addon.Discoverable,
		MultipleActivations: addon.MultipleActivations,
		Created:             created,
		Updated:             updated,
		OrderForm:           orderForm,
		UsesOrderForm:       addon.UsesOrderForm,
		BillingFrequency:    addon.BillingFrequency,
		ServiceModel:        addon.ServiceModel,
		Tagline:             addon.Tagline,
		Icon:                addon.Icon,
		BannerImage:         addon.BannerImage,
		Screenshots:         addon.Screenshots,
		ResellerMarketing:   resellerMarketing,
		EndUserMarketing:    enduserMarketing,
		Restrictions:        buildRestrictionsData(addon.Restrictions),
	}, nil
}
