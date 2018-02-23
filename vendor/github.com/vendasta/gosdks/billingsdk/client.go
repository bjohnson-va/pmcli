package billingsdk

import (
	"strconv"
	"time"

	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"golang.org/x/net/context"
)

const (
	productCreatePath     = "/internalApi/v2/product/create/"
	productUpdatePath     = "/internalApi/v2/product/update/"
	activateProductPath   = "/internalApi/v1/app/activate/"
	deactivateProductPath = "/internalApi/v1/app/deactivate/"
	getLineItemsPath      = "/internalApi/v1/invoice/line-items/partner/"
	getContractInfoPath   = "/internalApi/v2/contract/"
)

// BillingClient is a client which handles calls to billing's product apis.
// Implements the BillingClientInterface
type BillingClient struct {
	basesdk.SDKClient
}

// BillingClientInterface exposes methods for interacting with products in billing
type BillingClientInterface interface {
	ProductCreate(ctx context.Context, product Product, userID string, version int64) error
	ProductUpdate(ctx context.Context, product Product, userID string, version int64) error
	ActivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error
	DeactivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error
	GetLineItems(ctx context.Context, partnerID string, start time.Time, end time.Time, offset int, limit int) (*LineItemResponse, error)
	GetContractInfo(ctx context.Context, partnerID string) (*ContractInfo, error)
}

// BuildBillingClient creates a billing client.
func BuildBillingClient(apiUser string, apiKey string, env config.Env) *BillingClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &BillingClient{baseClient}
}

// ProductCreate creates a product in billing
func (c BillingClient) ProductCreate(ctx context.Context, product Product, userID string, version int64) error {
	params := map[string]interface{}{
		"userId":  userID,
		"version": version,
		"icon":    product.Icon,
		"price":   product.Price,
	}
	if product.ProductName != "" {
		params["productName"] = product.ProductName
	}
	if product.ProductID != "" {
		params["productId"] = product.ProductID
	}
	if product.BillingFrequency.String() != "" {
		params["billingFrequency"] = product.BillingFrequency.String()
	}

	response, err := c.Post(ctx, productCreatePath, params)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// ProductUpdate updates a product in billing
func (c BillingClient) ProductUpdate(ctx context.Context, product Product, userID string, version int64) error {
	params := map[string]interface{}{
		"userId":  userID,
		"version": version,
		"icon":    product.Icon,
		"price":   product.Price,
	}
	if product.ProductName != "" {
		params["productName"] = product.ProductName
	}
	if product.ProductID != "" {
		params["productId"] = product.ProductID
	}
	if product.BillingFrequency.String() != "" {
		params["billingFrequency"] = product.BillingFrequency.String()
	}

	response, err := c.Post(ctx, productUpdatePath, params)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// PurchaseOptions are options that can be specified for a purchase.
type PurchaseOptions struct {
	Expiry     time.Time
	InstanceID string
}

// ActivateProduct notifies billing of a product activation
func (c BillingClient) ActivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error {
	params := map[string]interface{}{
		"accountGroupId": accountGroupID,
		"appId":          productID,
		"initiatorId":    userID,
	}
	if !time.Time.IsZero(opts.Expiry) {
		params["expiryDateTime"] = opts.Expiry.UTC().Format(time.RFC3339)
	}
	if opts.InstanceID != "" {
		params["itemId"] = opts.InstanceID
	}
	response, err := c.Post(ctx, activateProductPath, params)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// DeactivateProduct notifies billing of a product deactivation
func (c BillingClient) DeactivateProduct(ctx context.Context, accountGroupID string, productID string, userID string, opts PurchaseOptions) error {
	params := map[string]interface{}{
		"accountGroupId": accountGroupID,
		"appId":          productID,
		"initiatorId":    userID,
	}
	if !time.Time.IsZero(opts.Expiry) {
		params["expiryDateTime"] = opts.Expiry.UTC().Format(time.RFC3339)
	}
	if opts.InstanceID != "" {
		params["itemId"] = opts.InstanceID
	}
	response, err := c.Post(ctx, deactivateProductPath, params)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// GetLineItems gets the lineitems of the invoice
func (c BillingClient) GetLineItems(ctx context.Context, partnerID string, start time.Time, end time.Time, offset int, limit int) (*LineItemResponse, error) {
	params := map[string]interface{}{
		"limit":     strconv.Itoa(limit),
		"offset":    strconv.Itoa(offset),
		"partnerId": partnerID,
	}

	if !time.Time.IsZero(end) {
		params["endDateTime"] = basesdk.ConvertTimeToVAPITimestamp(end)
	}

	if !time.Time.IsZero(start) {
		params["startDateTime"] = basesdk.ConvertTimeToVAPITimestamp(start)
	}

	response, err := c.Get(ctx, getLineItemsPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return lineItemsFromResponse(response.Body)
}

// GetContractInfo the contract for the partner
func (c BillingClient) GetContractInfo(ctx context.Context, partnerID string) (*ContractInfo, error) {
	params := map[string]interface{}{
		"partnerId": partnerID,
	}

	response, err := c.Get(ctx, getContractInfoPath, params, basesdk.Idempotent())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return contractInfoFromResponse(response.Body)
}
