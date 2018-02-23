package concierge

import (
	"context"
)

var (
	DefaultMockClient = &MockClient{}
)

type MockClient struct {
	createCustomTaskFunc func(ctx context.Context, task *CustomTask) (string, error)
	createAccountFunc    func(ctx context.Context, partnerID string, businessID string) error
}

// NewMockClient creates a mock concierge client.
//
// Pass in a set of functions to mock the interface.
// Passing nil in place of a function uses a default happy path implementation
func NewMockClient(
	createCustomTaskFunc func(ctx context.Context, task *CustomTask) (string, error),
	createAccountFunc func(ctx context.Context, partnerID string, businessID string) error,
) ConciergeClient {
	return &MockClient{
		createCustomTaskFunc: createCustomTaskFunc,
		createAccountFunc:    createAccountFunc,
	}
}

// CreateCustomTask creates a custom concierge task using data on provided CustomTask struct
// Returns newly created task's taskId.
func (c *MockClient) CreateCustomTask(ctx context.Context, task *CustomTask) (string, error) {
	if c.createCustomTaskFunc == nil {
		return "TK-1234", nil
	}
	return c.createCustomTaskFunc(ctx, task)
}

// CreateAccount attempts to create an account for the given partnerID and businessID
// Returns the status code
func (c *MockClient) CreateAccount(ctx context.Context, partnerID string, businessID string) error {
	if c.createAccountFunc == nil {
		return nil
	}
	return c.createAccountFunc(ctx, partnerID, businessID)
}
