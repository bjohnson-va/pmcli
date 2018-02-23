package marketplacepackages

import (
	"time"

	"github.com/vendasta/gosdks/pb/marketplace_packages/v1"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GRPCClient a GRPC client for the marketplace apps microservice.
type GRPCClient struct {
	client marketplace_packages_v1.MarketplacePackagesClient
}

// NewGRPCClient returns an new GRPC client for the marketplace packages microservice.
func NewGRPCClient(ctx context.Context, address string, useTLS bool, scope string, dialOptions ...grpc.DialOption) (Client, error) {
	conn, err := vax.NewGRPCConnection(ctx, address, useTLS, scope, false, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &GRPCClient{client: marketplace_packages_v1.NewMarketplacePackagesClient(conn)}, nil
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

// ListPackages is a grpc client side function that call the ListPackages in the grpc sever side to get the Packages from the given info
func (c *GRPCClient) ListPackages(ctx context.Context, PartnerId string, MarketId string, LmiCategory Lmi_Categories, Statuses []Statuses_Status, Sort bool, PageSize int32, Cursor string) ([]*Package, string, bool, error) {

	statuses := make([]marketplace_packages_v1.Statuses_Status, len(Statuses))
	for i, v := range Statuses {
		statuses[i] = marketplace_packages_v1.Statuses_Status(v)

	}

	req := marketplace_packages_v1.ListPackagesRequest{
		PartnerId:   PartnerId,
		MarketId:    MarketId,
		LmiCategory: marketplace_packages_v1.Lmi_Categories(LmiCategory),
		Statuses:    statuses,
		Sort:        Sort,
		PageSize:    PageSize,
		Cursor:      Cursor,
	}

	var res *marketplace_packages_v1.ListPackagesResponse

	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		res, err = c.client.ListPackages(ctx, &req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, "", false, err
	}

	packages := make([]*Package, len(res.Packages))

	for i, p := range res.Packages {
		packages[i] = toPackageProto(*p)
	}

	return packages, res.NextCursor, res.HasMore, nil
}

// ListProducts is a grpc client side function that call the ListProducts in the grpc sever side to get the Products from the given info
func (c *GRPCClient) ListProducts(ctx context.Context, PartnerId string, PageSize int32, Cursor string) ([]*Product, string, bool, error) {

	req := marketplace_packages_v1.ListProductsRequest{
		PartnerId: PartnerId,
		PageSize:  PageSize,
		Cursor:    Cursor,
	}

	var res *marketplace_packages_v1.ListProductsResponse

	err := vax.Invoke(ctx, func(ctx context.Context, settings vax.CallSettings) error {
		var err error
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		res, err = c.client.ListProducts(ctx, &req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, "", false, err
	}

	products := make([]*Product, len(res.Products))

	for i, p := range res.Products {
		products[i] = toProductProto(*p)
	}

	return products, res.NextCursor, res.HasMore, nil
}

func toProductProto(p marketplace_packages_v1.Product) *Product {

	serviceModel := make([]ServiceModel_Type, len(p.ServiceModel))
	for i, s := range p.ServiceModel {
		serviceModel[i] = ServiceModel_Type(s)
	}

	lmiCategories := make([]Lmi_Categories, len(p.LmiCategories))
	for i, c := range p.LmiCategories {
		lmiCategories[i] = Lmi_Categories(c)
	}

	faqs := make([]*FrequentlyAskedQuestions, len(p.Faqs))
	for i, q := range p.Faqs {
		faqs[i] = &FrequentlyAskedQuestions{
			Question: q.Question,
			Answer:   q.Answer,
		}
	}

	return &Product{
		ProductId:             p.ProductId,
		PartnerId:             p.PartnerId,
		Name:                  p.Name,
		Tagline:               p.Tagline,
		Description:           p.Description,
		KeySellingPoints:      p.KeySellingPoints,
		IconUrl:               p.IconUrl,
		HeaderImageUrl:        p.HeaderImageUrl,
		ScreenshotUrls:        p.ScreenshotUrls,
		PdfUploadUrls:         p.PdfUploadUrls,
		Currency:              Currencies_Currency(p.Currency),
		WholesalePrice:        p.WholesalePrice,
		RecommendedSellPrice:  p.RecommendedSellPrice,
		BillingFrequency:      p.BillingFrequency,
		ServiceModel:          serviceModel,
		Category:              p.Category,
		LmiCategories:         lmiCategories,
		Origin:                p.Origin,
		UsesCustomizationForm: p.UsesCustomizationForm,
		EntryUrl:              p.EntryUrl,
		SessionUrl:            p.SessionUrl,
		WebsiteUrl:            p.WebsiteUrl,
		UpdatedBy:             p.UpdatedBy,
		EndDate:               p.EndDate,
		Faqs:                  faqs,
		Created:               p.Created,
		Updated:               p.Updated,
		IsArchived:            p.IsArchived,
		Restrictions:          buildRestrictionsData(p.Restrictions),
	}
}

func buildRestrictionsData(restrictions *marketplace_packages_v1.Restrictions) *Restrictions {
	var whiteList []string
	var blackList []string

	if restrictions != nil && restrictions.Country != nil {
		whiteList = restrictions.Country.Whitelist
		blackList = restrictions.Country.Blacklist
	}

	countryPermissionLists := PermissionLists{
		WhiteList: whiteList,
		BlackList: blackList,
	}

	return &Restrictions{
		Country: &countryPermissionLists,
	}
}

func toPackageProto(p marketplace_packages_v1.Package) *Package {

	prices := make([]*Price, len(p.Pricing.Prices))
	for i, p := range p.Pricing.Prices {
		prices[i] = &Price{
			Price:     int32(p.Price),
			Frequency: Frequencies_Frequency(p.Frequency),
		}
	}
	pricing := &Pricing{
		Prices:   prices,
		Currency: Currencies_Currency(p.Pricing.Currency),
	}

	return &Package{
		Created:                  p.Created,
		Updated:                  p.Updated,
		Archived:                 p.Archived,
		UpdatedBy:                p.UpdatedBy,
		PackageId:                p.PackageId,
		PartnerId:                p.PartnerId,
		MarketId:                 p.MarketId,
		Name:                     p.Name,
		Icon:                     p.Icon,
		Status:                   Statuses_Status(p.Status),
		HeaderImageUrl:           p.HeaderImageUrl,
		Tagline:                  p.Tagline,
		Content:                  p.Content,
		Products:                 p.Products,
		HideProductDetails:       p.HideProductDetails,
		HideProductIconsAndNames: p.HideProductIconsAndNames,
		Pricing:                  pricing,
		NormalizedAnnualPrice:    p.NormalizedAnnualPrice,
	}
}
