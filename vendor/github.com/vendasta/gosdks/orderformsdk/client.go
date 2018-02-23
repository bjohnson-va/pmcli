// IF YOU NEED TO USE THIS SDK, THE ORDER FORM SHOULD NOT BE IN VBC MOVE THEM!!
package orderformsdk

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
	"net/http"
	"time"
)

const (
	lockPath            = "internalApi/v3/orderFormSubmission/lock/"
	setActivationIdPath = "internalApi/v3/orderFormSubmission/activate/"
	createPath          = "internalApi/v3/orderFormSubmission/create/"
	getPath             = "internalApi/v3/orderFormSubmission/get/"
)

//order form client handles the calls to VBC to lock the order form and set the IDs
type OrderFormClient struct {
	basesdk.BaseClient
}

//Interface for order form service
type OrderFormInterface interface {
	Lock(ctx context.Context, orderFormSubmissionID string) error
	SetActivationID(ctx context.Context, orderFormSubmissionID string, activationID string) error
	Create(ctx context.Context, partnerID string, accountGroupID string, productID string, commonFields []OrderField, customFields []OrderField) (string, error)
	GetOrderFormSubmission(ctx context.Context, orderFormSubmissionId string) (*OrderFormSubmission, error)
}

// hopefully this will eventually go away and get the order form from somewhere else
// Build the client to call building order form
func BuildOrderFormClient(apiUser string, apiKey string, env config.Env) *OrderFormClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: rootURLFromEnv(env)}
	return &OrderFormClient{baseClient}
}

// DataFromResponse returns Data from an http Reponse
func DataFromResponse(r *http.Response) (map[string]interface{}, error) {
	defer r.Body.Close()
	type Response struct {
		Data map[string]interface{} `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to Data: " + err.Error()
		return nil, errors.New(reason)
	}
	if res.Data == nil {
		return nil, errors.New("failed to convert response to Data")
	}
	return res.Data, nil

}

// Lock calls the lock endpoint in VBC to lock the order form
func (c *OrderFormClient) Lock(ctx context.Context, orderFormSubmissionID string) error {

	params := map[string]interface{}{
		"orderFormSubmissionId": orderFormSubmissionID,
	}

	response, err := c.Post(ctx, lockPath, params)
	if err != nil {
		return basesdk.ConvertHttpErrorToGRPC(err)
	}
	defer response.Body.Close()
	return nil
}

// SetActivationID calls activate endpoint in VBC to set the activation id of the order form
func (c *OrderFormClient) SetActivationID(ctx context.Context, orderFormSubmissionID string, activationID string) error {

	params := map[string]interface{}{
		"orderFormSubmissionId": orderFormSubmissionID,
		"activationId":          activationID,
	}

	response, err := c.Post(ctx, setActivationIdPath, params)
	if err != nil {
		return basesdk.ConvertHttpErrorToGRPC(err)
	}
	defer response.Body.Close()
	return nil
}

// OrderField represents a field on an order
type OrderField struct {
	ID     		string
	Type   		string
	Answer      string
	Description string
	Label       string
}

// OrderFieldsToParams converts OrderFields to a form suitable to be passed as params into a Post request
func OrderFieldsToParams(orderFields []OrderField) []map[string]string {
	var result []map[string]string
	for _, orderField := range orderFields {
		entry := map[string]string {
			"field_id": orderField.ID,
			"field_type": orderField.Type,
			"answer": orderField.Answer,
			"description": orderField.Description,
			"label": orderField.Label,
		}
		result = append(result, entry)
	}
	return result
}

// Create calls create endpoint in VBC to create an order form
func (c *OrderFormClient) Create(ctx context.Context, partnerID string, accountGroupID string, productID string, commonFields []OrderField, customFields []OrderField) (string, error) {
	params := map[string]interface{} {
		"partnerId": partnerID,
		"accountGroupId": accountGroupID,
		"productId": productID,
		"commonFields": OrderFieldsToParams(commonFields),
		"customFields": OrderFieldsToParams(customFields),
	}

	response, err := c.Post(ctx, createPath, params)
	if err != nil {
		return "", basesdk.ConvertHttpErrorToGRPC(err)
	}

	data, err := DataFromResponse(response)
	if err != nil {
		return "", basesdk.ConvertHttpErrorToGRPC(err)
	}

	orderFormSubmissionId, ok := data["orderFormSubmissionId"]
	if !ok {
		return "", basesdk.ConvertHttpErrorToGRPC(errors.New("no orderFormSubmissionId on order create response"))
	}

	return orderFormSubmissionId.(string), nil
}

type OrderFormSubmissionField struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
	FieldId     string `json:"field_id"`
	FieldType   string `json:"field_type"`
}

type OrderFormSubmission struct {
	PartnerId             string `json:"partnerId"`
	AccountGroupId        string `json:"accountGroupId"`
	ProductId             string `json:"productId"`
	ActivationId          string `json:"activationId"`
	OrderFormSubmissionId string `json:"orderFormSubmissionId"`
	Locked                bool   `json:"locked"`

	CommonFields []*OrderFormSubmissionField `json:"commonFields"`
	CustomFields []*OrderFormSubmissionField `json:"customFields"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// Get calls VBC endpoint to get an orderFormSubmission given an order_form_submission_id
func (c *OrderFormClient) GetOrderFormSubmission(ctx context.Context, orderFormSubmissionId string) (*OrderFormSubmission, error) {
	params := map[string]interface{}{
		"orderFormSubmissionId": orderFormSubmissionId,
	}
	response, err := c.Get(ctx, getPath, params)
	if err != nil {
		return nil, basesdk.ConvertHttpErrorToGRPC(err)
	}
	orderForm, err := orderFormSubmissionFromResponse(response)
	if err != nil {
		return nil, basesdk.ConvertHttpErrorToGRPC(err)
	}
	return orderForm, nil
}

func orderFormSubmissionFromResponse(r *http.Response) (*OrderFormSubmission, error) {
	defer r.Body.Close()
	type Response struct {
		OrderFormSubmission *OrderFormSubmission `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to OrderFormSubmission: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.OrderFormSubmission, nil
}
