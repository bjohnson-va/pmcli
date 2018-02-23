package concierge

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"

	"testing"
)

func Test_CreateAccountValidatesRequiredPartnerId(t *testing.T) {
	err := noHTTPConciergeClient.CreateAccount(context.Background(), "", "AG-1234")
	assert.EqualError(t, err, "partnerID is required")
}

func Test_CreateAccountValidatesRequiredBusinessId(t *testing.T) {
	err := noHTTPConciergeClient.CreateAccount(context.Background(), "ABC", "")
	assert.EqualError(t, err, "businessID is required")
}

func Test_CreateAccountReturnsNilErrorFromSuccessfulCall(t *testing.T) {
	client := conciergeClient{&basesdk.BaseClientMock{}}
	err := client.CreateAccount(context.Background(), "ABC", "AG-1234")
	assert.Nil(t, err)
}
