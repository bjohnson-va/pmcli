package marketplaceapps

import (
	"fmt"
	"time"

	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

// IdentifiedPublicKey The identified public keys are considered "identified" because they have an id/key
// associated with them that identifies a user or purpose for using that particular key instead of the default one,
// for being  able to identify the callers of the API. It could be a backend service, a particular user or admin
// for an app, etc.
type IdentifiedPublicKey struct {
	// The public key ID (corresponds to a CredentialsID in the JWT).
	PublicKeyID string
	// The public key to use when given a CredentialsID matching the PublicKeyID.
	PublicKey string
}

// PublicKeys the public keys configured for an app.
type PublicKeys struct {
	// The default public key for the app.
	PublicKey string
	// The identified public keys for the app.
	IdentifiedPublicKeys []*IdentifiedPublicKey
}

type OrderFormField struct {
	Label       string
	Id          string
	Type        string
	Options     []string
	Required    bool
	Description string
	UploadUrl   string
}

type IncludedCommonFormFields struct {
	BusinessName           bool
	BusinessAddress        bool
	BusinessPhoneNumber    bool
	BusinessAccountGroupId bool
	SalespersonName        bool
	SalespersonPhoneNumber bool
	SalespersonEmail       bool
	ContactName            bool
	ContactPhoneNumber     bool
	ContactEmail           bool
}

type OrderForm struct {
	OrderForm         []*OrderFormField
	CommonForm        *IncludedCommonFormFields
	ActivationMessage string
}

type FrequentlyAskedQuestions struct {
	Question string
	Answer   string
}

type MarketingInformation struct {
	Description      string
	KeySellingPoints []string
	Faqs             []*FrequentlyAskedQuestions
	Files            []string
}

// White and black list permissions
type PermissionLists struct {
	WhiteList []string
	BlackList []string
}

// Restrictions for certain fields
type Restrictions struct {
	Country *PermissionLists
}

// Addon is a item that can be purchased to enable additional features within a marketplace app.
// These can be either Consumables, Non-Consumables or Subscriptions.
type Addon struct {
	// The id of the addon, unique within the scope of the app the addon is for
	AddonID string
	// The id of the app which the addon is for
	AppID string
	// The wholesale price of the addon, in cents
	Price int64
	// The title of the addon
	Title string
	// Order form fields
	OrderForm *OrderForm
	// Indicates if the order form should be used during addon activation
	UsesOrderForm bool
	// Addon billing_frequency
	BillingFrequency string
	// Addon service_model
	ServiceModel []string
	// A tagline for an addon
	Tagline string
	// Icon url for the addon
	Icon string
	// The banner image of an addon
	BannerImage string
	// Reseller marketing information
	ResellerMarketing *MarketingInformation
	// End user marketing information
	EndUserMarketing *MarketingInformation
	// A timestamp of when the addon was created
	Created time.Time
	// A timestamp of when the addon was last updated
	Updated time.Time
	// Is the addon approved for marketplace
	Approved bool
	// Is the addon discoverable within marketplace
	Discoverable bool
	// The list of screenshots for an addon
	Screenshots []string
	// Can the addon support multiple activations
	MultipleActivations bool
	// What restrictions does the addon have
	Restrictions *Restrictions
}

// Integration with vendor center
type Integration struct {
	// The public key of the app
	PublicKey string
	// A list of identified public keys
	IdentifiedPublicKeys []*IdentifiedPublicKey
	// Email of contact for app
	ContactEmail string
	// The status of the application
	Approved bool
	// The name of the vendor
	Vendor string
	// A list of the users in the application
	UserIds []string
	// Webhook url: purchase_url
	PurchaseWebhookUrl string
	// Webhook url: addon_purchase_url
	AddonPurchaseWebhookUrl string
	// Webhook url: user_mod_url
	UserModWebhookUrl string
	// Webhook url: salestool_url
	SalestoolWebhookUrl string
	// Webhook url: account_mod_url
	AccountModWebhookUrl string
	// Webhook url: logout url
	LogoutWebhookUrl string
	// Email notifications for fulfillment
	FulfillmentEmail []string
	// SSO url: Entry point for the application
	EntryUrl string
	// SSO url: Session code generation url
	SessionUrl string
	// Service url: get current settings/schema, post updated data
	SettingsUrl string
	// Reserve Url: for calling out to applications to let them know a activation is coming, but not guaranteed yet
	ReserveIdUrl string
}

type MarketplaceApp_Lmi_LmiCategories int32
type ProductCategory int32

// App is a item that can be purchased or enabled to get a service
type App struct {
	// The id of the app
	AppId string
	// The version of the app
	Version int64
	// The name of the app
	Name string
	// Icon for the product
	Icon string
	// Allow white labelling
	AllowWhitelabel bool
	// The app is currently in development
	InDevelopment bool
	// Supports a trial state of the application
	TrialSupport bool
	// Order form: the product uses order form
	UsesOrderForm bool
	// Product information: tagline
	Tagline string
	// Product information: sign_in_url
	SignInUrl string
	// Product information: currency
	Currency string
	// Product information: price stored as cents
	Price int64
	// Product information: billing_frequency
	BillingFrequency string
	// Product information: service_model
	ServiceModel []string
	// Product information: screenshot
	Screenshot []string
	// Product information: website_url
	WebsiteUrl string
	// Product information: header_image
	HeaderImage string
	// Category of the app
	ProductCategory ProductCategory
	// Local marketing index category of the app
	LmiCategories []MarketplaceApp_Lmi_LmiCategories
	// Integration with vendor center
	Integration *Integration
	// Order form fields
	OrderForm *OrderForm
	// Reseller marketing information
	ResellerMarketing *MarketingInformation
	// End user marketing information
	EndUserMarketing *MarketingInformation
	// A timestamp of when the app was created
	Created *time.Time
	// A timestamp of when the app was last updated
	Updated *time.Time
	// What restrictions the app has
	Restrictions *Restrictions
}

// Client an interface for the marketplace apps microservice.
type Client interface {
	// GetPublicKeys get's the public keys for a marketplace application.
	GetPublicKeys(ctx context.Context, appID string) (*PublicKeys, error)
	// GetAddon get's a specific addon for the specified app
	GetAddon(ctx context.Context, appID string, addonID string) (*Addon, error)
	// GetSSOUrls get's the SSO Urls (entry and session URL's) for a marketplace application.
	GetSSOUrls(ctx context.Context, appID string) (entryURL string, sessionURL string, err error)
	// GetApp gets an app based off of appID
	GetApp(ctx context.Context, appID string) (app *App, err error)
	// ListApprovedAddons lists all the approved addons for an app
	ListApprovedAddons(ctx context.Context, appID string, cursor string, pageSize int64, appIDs []string) ([]*Addon, string, bool, error)
	// ListAddons lists all the addons for an app
	ListAddons(ctx context.Context, appID string, cursor string, pageSize int64) ([]*Addon, string, bool, error)
}

// NewClient returns the default concrete implementation of the client, pre-configured given an environment.
func NewClient(ctx context.Context, e config.Env) (Client, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("unable to create client with environment %d", e)
	}
	useTLS := e != config.Local
	return NewGRPCClient(ctx, address, useTLS, scopes[e])
}

var addresses = map[config.Env]string{
	config.Local: "marketplace-apps-api-test.vendasta-internal.com:443",
	config.Test:  "marketplace-apps-api-test.vendasta-internal.com:443",
	config.Demo:  "marketplace-apps-api-demo.vendasta-internal.com:443",
	config.Prod:  "marketplace-apps-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "https://marketplace-apps-api-test.vendasta-internal.com",
	config.Test:  "https://marketplace-apps-api-test.vendasta-internal.com",
	config.Demo:  "https://marketplace-apps-api-demo.vendasta-internal.com",
	config.Prod:  "https://marketplace-apps-api-prod.vendasta-internal.com",
}
