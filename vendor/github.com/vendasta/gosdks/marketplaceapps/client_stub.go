package marketplaceapps

import (
	"golang.org/x/net/context"
)

// Stub implements the MarketplaceApps Client for the purposes of testing
type Stub struct {
}

// GetPublicKeys get's the public keys for a marketplace application.
func (s Stub) GetPublicKeys(ctx context.Context, appID string) (*PublicKeys, error) {
	return nil, nil
}

// GetAddon get's a specific addon for the specified app
func (s Stub) GetAddon(ctx context.Context, appID string, addonID string) (*Addon, error) {
	orderFormFields := make([]*OrderFormField, 1)
	addon := &Addon{
		AddonID:             addonID,
		AppID:               appID,
		Title:               "Rep Man",
		Approved:            true,
		MultipleActivations: false,
		Tagline:             "This is the Best",
		BannerImage:         "http://bannerImage.com",
		UsesOrderForm:       true,
		OrderForm: &OrderForm{
			OrderForm: orderFormFields,
			CommonForm: &IncludedCommonFormFields{
				BusinessAccountGroupId: false,
				BusinessAddress:        false,
				BusinessName:           false,
				BusinessPhoneNumber:    false,
				ContactEmail:           false,
				ContactName:            false,
				ContactPhoneNumber:     false,
				SalespersonEmail:       false,
				SalespersonName:        false,
				SalespersonPhoneNumber: false,
			},
			ActivationMessage: "",
		},
		Restrictions: &Restrictions{Country: nil},
	}
	if addonID == "supportMultiple" {
		addon.MultipleActivations = true
	}
	if addonID == "noTagline" {
		addon.Tagline = ""
	}
	if addonID == "noBannerImage" {
		addon.BannerImage = ""
	}
	if addonID == "noBannerOrTagline" {
		addon.BannerImage = ""
		addon.Tagline = ""
	}
	if addonID == "noUsesOrderForm" {
		addon.UsesOrderForm = false
	}
	return addon, nil
}

// GetSSOUrls get's the SSO Urls (entry and session URL's) for a marketplace application.
func (s Stub) GetSSOUrls(ctx context.Context, appID string) (entryURL string, sessionURL string, err error) {
	return "", "", nil
}

// GetApp returns a marketplace application
func (s Stub) GetApp(ctx context.Context, appID string) (*App, error) {
	orderFormFields := make([]*OrderFormField, 1)
	App := &App{
		AppId:            appID,
		Version:          123143128438194,
		Name:             "Rep Man",
		Icon:             "https://thisisitheICon.com",
		AllowWhitelabel:  true,
		InDevelopment:    false,
		TrialSupport:     true,
		UsesOrderForm:    true,
		Tagline:          "They are Great",
		SignInUrl:        "http://YAY.com",
		Currency:         "USD",
		Price:            1000,
		BillingFrequency: "Monthly",
		ServiceModel:     []string{"DIFM"},
		Screenshot:       []string{"http://www.jajajaja.com"},
		WebsiteUrl:       "http://www.ohyea.org",
		HeaderImage:      "www.HeaderImage.com",
		ProductCategory:  ProductCategory(3),
		Integration: &Integration{
			PublicKey:               "publicKey",
			IdentifiedPublicKeys:    nil,
			ContactEmail:            "cool@test.com",
			Approved:                true,
			Vendor:                  "Bacon man",
			UserIds:                 []string{"abc123"},
			PurchaseWebhookUrl:      "www.test1.com",
			AddonPurchaseWebhookUrl: "www.addon.com",
			UserModWebhookUrl:       "www.user.com",
			SalestoolWebhookUrl:     "www.ST_url.com",
			AccountModWebhookUrl:    "www.accountURL.com",
			LogoutWebhookUrl:        "www.logout.com",
			FulfillmentEmail:        []string{"beans@testing.com"},
			EntryUrl:                "www.yes.com",
			SessionUrl:              "www.session.com",
			SettingsUrl:             "www.Setting.com",
		},
		OrderForm: &OrderForm{
			OrderForm: orderFormFields,
			CommonForm: &IncludedCommonFormFields{
				BusinessAccountGroupId: false,
				BusinessAddress:        false,
				BusinessName:           false,
				BusinessPhoneNumber:    false,
				ContactEmail:           false,
				ContactName:            false,
				ContactPhoneNumber:     false,
				SalespersonEmail:       false,
				SalespersonName:        false,
				SalespersonPhoneNumber: false,
			},
			ActivationMessage: "",
		},
		ResellerMarketing: &MarketingInformation{
			Description:      "reseller info",
			KeySellingPoints: []string{"Its cool", "Its Great"},
			Faqs: []*FrequentlyAskedQuestions{
				{Question: "this is a question", Answer: "this is an answer"},
			},
			Files: []string{"YAY"},
		},
		EndUserMarketing: &MarketingInformation{
			Description:      "enduser info",
			KeySellingPoints: []string{"Its so cool", "Its so Great"},
			Faqs: []*FrequentlyAskedQuestions{
				{Question: "question", Answer: "answer"},
			},
			Files: []string{"YAY"},
		},
		Restrictions: &Restrictions{Country: nil},
	}

	if appID == "MP-NoOrderForm" {
		App.UsesOrderForm = false
	}

	return App, nil
}

func (s Stub) ListAddons(ctx context.Context, appID string, cursor string, pageSize int64) ([]*Addon, string, bool, error) {
	return nil, "", true, nil
}

func (s Stub) ListApprovedAddons(ctx context.Context, appID string, cursor string, pageSize int64, appIDs []string) ([]*Addon, string, bool, error) {
	return nil, "", false, nil
}
