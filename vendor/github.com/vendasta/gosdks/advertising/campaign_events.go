package advertising

import (
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/pb/advertising/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// OrderEvent contains all information pertaining to a specific ordering of a Digital Ads Product.
// BusinessID, PartnerID and OrderID are required.
type OrderEvent struct {
	BusinessID            string
	PartnerID             string
	OrderID               string
	PackageType           string
	AdvertisingPackage    string
	BusinessName          string
	BusinessAddress       string
	BusinessPhone         string
	ContactEmail          string
	ContactName           string
	ContactPhone          string
	SalespersonEmail      string
	SalespersonName       string
	SalespersonPhone      string
	CreativeLocation      []string
	CustomerValue         string
	FacebookPageURL       string
	Notes                 string
	Objective             string
	Promo                 string
	RetailPrice           string
	Specialties           string
	Targeting             string
	Term                  string
	WebsiteURL            string
	RawMarketplacePayload string
	CreatedAt             time.Time
}

// Validation returns a grpc error representing missing or invalid data or nil if the OrderEvent is valid.
func (o *OrderEvent) Validation(ctx context.Context) error {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(o.BusinessID, util.InvalidArgument, "BusinessID is required"),
		validation.StringNotEmpty(o.OrderID, util.InvalidArgument, "OrderID is required"),
		validation.TimeNot0(o.CreatedAt, util.InvalidArgument, "CreatedAt time cannot be zero"),
	).Validate()

	return err
}

func (c *client) CreateOrderEvent(ctx context.Context, order *OrderEvent) error {

	ctx = metadata.NewOutgoingContext(ctx, nil)
	err := order.Validation(ctx)
	if err != nil {
		return err
	}

	createdAtProtoTime, err := ptypes.TimestampProto(order.CreatedAt)
	if err != nil {
		return err
	}
	orderEvent := &advertising_v1.OrderEvent{
		BusinessId:            order.BusinessID,
		PartnerId:             order.PartnerID,
		OrderId:               order.OrderID,
		PackageType:           order.PackageType,
		AdvertisingPackage:    order.AdvertisingPackage,
		BusinessName:          order.BusinessName,
		BusinessAddress:       order.BusinessAddress,
		BusinessPhone:         order.BusinessPhone,
		ContactEmail:          order.ContactEmail,
		ContactName:           order.ContactName,
		ContactPhone:          order.ContactPhone,
		SalespersonEmail:      order.SalespersonEmail,
		SalespersonName:       order.SalespersonName,
		SalespersonPhone:      order.SalespersonPhone,
		CreativeLocation:      order.CreativeLocation,
		CustomerValue:         order.CustomerValue,
		FacebookPageUrl:       order.FacebookPageURL,
		Notes:                 order.Notes,
		Objective:             order.Objective,
		Promo:                 order.Promo,
		RetailPrice:           order.RetailPrice,
		Specialties:           order.Specialties,
		Targeting:             order.Targeting,
		Term:                  order.Term,
		WebsiteUrl:            order.WebsiteURL,
		RawMarketplacePayload: order.RawMarketplacePayload,
		CreatedAt:             createdAtProtoTime,
	}

	request := &advertising_v1.CreateOrderEventRequest{OrderEvent: orderEvent}
	err = vax.Invoke(ctx, func(ctx context.Context, cs vax.CallSettings) error {
		_, createOrderEventErr := c.campaignEventsClient.CreateOrderEvent(ctx,
			request,
			grpc.FailFast(false),
		)
		return createOrderEventErr
	}, defaultRetryCallOptions)

	if err != nil {
		if strings.Contains(err.Error(), "rpc error: code = AlreadyExists") {
			logging.Infof(ctx, "OrderEvent %s:%s already exists", order.BusinessID, order.OrderID)
		} else {
			return err
		}
	}
	return nil
}
