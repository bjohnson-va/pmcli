package marketplacepackages

import (
	"fmt"
	"time"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

type Currencies_Currency int32

const (
	Currencies_USD Currencies_Currency = 0
	Currencies_AUD Currencies_Currency = 1
	Currencies_BRL Currencies_Currency = 2
	Currencies_CAD Currencies_Currency = 3
	Currencies_CHF Currencies_Currency = 4
	Currencies_CNY Currencies_Currency = 5
	Currencies_CZK Currencies_Currency = 6
	Currencies_EUR Currencies_Currency = 7
	Currencies_GBP Currencies_Currency = 8
	Currencies_HKD Currencies_Currency = 9
	Currencies_INR Currencies_Currency = 10
	Currencies_JPY Currencies_Currency = 11
	Currencies_KHR Currencies_Currency = 12
	Currencies_KRW Currencies_Currency = 13
	Currencies_MXN Currencies_Currency = 14
	Currencies_NOK Currencies_Currency = 15
	Currencies_NZD Currencies_Currency = 16
	Currencies_RUB Currencies_Currency = 17
	Currencies_SEK Currencies_Currency = 18
	Currencies_SGD Currencies_Currency = 19
	Currencies_TRY Currencies_Currency = 20
	Currencies_ZAR Currencies_Currency = 21
)

type ServiceModel_Type int32

const (
	ServiceModel_DIY  ServiceModel_Type = 0
	ServiceModel_DIWM ServiceModel_Type = 1
	ServiceModel_DIFM ServiceModel_Type = 2
)

type Lmi_Categories int32

const (
	Lmi_ADVERTISING            Lmi_Categories = 0
	Lmi_WEBSITE                Lmi_Categories = 1
	Lmi_CONTENT_AND_EXPERIENCE Lmi_Categories = 2
	Lmi_LISTINGS               Lmi_Categories = 3
	Lmi_REPUTATION             Lmi_Categories = 4
	Lmi_SEO                    Lmi_Categories = 5
	Lmi_SOCIAL                 Lmi_Categories = 6
)

type Statuses_Status int32

const (
	Statuses_DRAFT     Statuses_Status = 0
	Statuses_PUBLISHED Statuses_Status = 1
	Statuses_ARCHIVED  Statuses_Status = 2
)

type Frequencies_Frequency int32

const (
	Frequencies_MONTHLY Frequencies_Frequency = 0
	Frequencies_DAILY   Frequencies_Frequency = 1
	Frequencies_ONCE    Frequencies_Frequency = 2
	Frequencies_YEARLY  Frequencies_Frequency = 3
	Frequencies_WEEKLY  Frequencies_Frequency = 4
	Frequencies_OTHER   Frequencies_Frequency = 5
)

// Token holds the response data from a GetOAuthToken call.
type Token struct {
	AccessToken string
	TokenType   string
	Expires     time.Time
}

type FrequentlyAskedQuestions struct {
	Question string
	Answer   string
}

type Price struct {
	// Price of the product
	Price int32
	// Billing frequency of the product
	Frequency Frequencies_Frequency
}

type Pricing struct {
	// Currency of the pricing
	Currency Currencies_Currency
	// Prices of the package
	Prices []*Price
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

type Package struct {
	Created                  *google_protobuf.Timestamp
	Updated                  *google_protobuf.Timestamp
	Archived                 *google_protobuf.Timestamp
	UpdatedBy                string
	PackageId                string
	PartnerId                string
	MarketId                 string
	Name                     string
	Icon                     string
	Status                   Statuses_Status
	HeaderImageUrl           string
	Tagline                  string
	Content                  string
	Products                 []string
	HideProductDetails       bool
	HideProductIconsAndNames bool
	Pricing                  *Pricing
	NormalizedAnnualPrice    int64
}

type Product struct {
	ProductId             string
	PartnerId             string
	Name                  string
	Tagline               string
	Description           string
	KeySellingPoints      []string
	IconUrl               string
	HeaderImageUrl        string
	ScreenshotUrls        []string
	PdfUploadUrls         []string
	Currency              Currencies_Currency
	WholesalePrice        string
	RecommendedSellPrice  string
	BillingFrequency      string
	ServiceModel          []ServiceModel_Type
	Category              string
	LmiCategories         []Lmi_Categories
	Origin                string
	UsesCustomizationForm bool
	EntryUrl              string
	SessionUrl            string
	WebsiteUrl            string
	UpdatedBy             string
	EndDate               *google_protobuf.Timestamp
	Faqs                  []*FrequentlyAskedQuestions
	Created               *google_protobuf.Timestamp
	Updated               *google_protobuf.Timestamp
	IsArchived            bool
	Restrictions          *Restrictions
}

// Client is an interface for the marketplace packages microservice.
type Client interface {
	ListPackages(ctx context.Context, PartnerId string, MarketId string, LmiCategory Lmi_Categories, Statuses []Statuses_Status, Sort bool, PageSize int32, Cursor string) ([]*Package, string, bool, error)
	ListProducts(ctx context.Context, PartnerId string, PageSize int32, Cursor string) ([]*Product, string, bool, error)
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
	config.Local: "marketplace-packages-api-test.vendasta-internal.com:443",
	config.Test:  "marketplace-packages-api-test.vendasta-internal.com:443",
	config.Demo:  "marketplace-packages-api-demo.vendasta-internal.com:443",
	config.Prod:  "marketplace-packages-api-prod.vendasta-internal.com:443",
}

var scopes = map[config.Env]string{
	config.Local: "https://marketplace-packages-api-test.vendasta-internal.com",
	config.Test:  "https://marketplace-packages-api-test.vendasta-internal.com",
	config.Demo:  "https://marketplace-packages-api-demo.vendasta-internal.com",
	config.Prod:  "https://marketplace-packages-api-prod.vendasta-internal.com",
}
