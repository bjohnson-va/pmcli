package concierge

import (
	"context"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/config"
)

const (
	createCustomTaskPath = "/internalApi/v2/tasks/custom/create"
)

// CustomTask represents a custom task in Concierge
type CustomTask struct {
	PartnerID        string `json:"partnerId,omitempty"`
	BusinessID       string `json:"accountGroupId,omitempty"`
	Status           string `json:"status,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	PublicNote       string `json:"publicNote,omitempty"`
	PrivateNote      string `json:"privateNote,omitempty"`
	UniqueExternalID string `json:"uniqueExternalId, omitempty"`
}

// ConciergeClientInterface defines the methods available on the Concierge client
type ConciergeClient interface {
	// CreateCustomTask creates a custom concierge task using data on provided CustomTask struct
	// Returns newly created task's taskId.
	CreateCustomTask(ctx context.Context, task *CustomTask) (string, error)

	// CreateAccount creates a concierge account for the given businessID
	CreateAccount(ctx context.Context, partnerID string, businessID string) (error)
}

type conciergeClient struct {
	basesdk.SDKClient
}

// NewConciergeClient creates a concierge client
func NewConciergeClient(apiUser string, apiKey string, env config.Env) *conciergeClient {
	auth := basesdk.UserKey{APIUser: apiUser, APIKey: apiKey}
	baseClient := basesdk.BaseClient{Authorization: auth, RootURL: hostMap[env]}
	return &conciergeClient{baseClient}
}

var hostMap = map[config.Env]string{
	config.Local: "http://localhost:8090",
	config.Test:  "https://arm-test.appspot.com",
	config.Demo:  "https://arm-demo.appspot.com",
	config.Prod:  "https://arm-prod.appspot.com",
}
