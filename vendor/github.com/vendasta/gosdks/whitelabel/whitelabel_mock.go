package whitelabel

import "golang.org/x/net/context"

type mockWhiteLabelClient struct {
	*WhitelabelData
}

// BuildMockWhiteLabelClient creates a white label client
func BuildMockWhiteLabelClient(mockResponse *WhitelabelData) WhitelabelClientInterface {
	return mockWhiteLabelClient{
		WhitelabelData: mockResponse,
	}
}

// Get returns the white label data.
func (c mockWhiteLabelClient) Get(ctx context.Context, partnerID string, marketID string, options ...GetOption) (*WhitelabelData, error) {
	return c.WhitelabelData, nil
}

type mockWhiteLabelClientWithError struct {
	Error error
}

//BuildMockWhitelabelClientThatReturnsAnError creates a white label client that always returns an error
func BuildMockWhitelabelClientThatReturnsAnError(err error) WhitelabelClientInterface {
	return mockWhiteLabelClientWithError{
		Error: err,
	}
}

// Get returns an error
func (c mockWhiteLabelClientWithError) Get(ctx context.Context, partnerID string, marketID string, option ...GetOption) (*WhitelabelData, error) {
	return nil, c.Error
}
