package cssdk

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

func handleSocialProfileResponseMock(_ *http.Response) (*SocialProfile, error) {
	return &SocialProfile{}, nil
}

func handleSocialProfileResponseErrorMock(_ *http.Response) (*SocialProfile, error) {
	return nil, errors.New("Failed to handle response")
}

func handleSocialProfileRegistrationResponseMock(_ *http.Response) (*SocialProfileRegistration, error) {
	return &SocialProfileRegistration{}, nil
}

func handleSocialProfileRegistrationResponseErrorMock(_ *http.Response) (*SocialProfileRegistration, error) {
	return nil, errors.New("Failed to handle response")
}

func handleSocialProfileRegistrationsResponseMock(_ *http.Response) ([]*SocialProfileRegistration, error) {
	return []*SocialProfileRegistration{{}}, nil
}

func handleSocialProfileRegistrationsResponseErrorMock(_ *http.Response) ([]*SocialProfileRegistration, error) {
	return nil, errors.New("Failed to handle response")
}

// createSocialProfile tests
func TestCreateSocialProfileReturnsAnErrorIfNoAccountGroupIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.createSocialProfile(context.Background(), nil, "", "")
	assert.EqualError(t, err, "account group id is required")
}

func TestCreateSocialProfileReturnsAnErrorIfTheResponseErrors(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.createSocialProfile(context.Background(), handleSocialProfileResponseErrorMock, "AG-123", "SPG-444")
	assert.EqualError(t, err, "Failed to handle response")
}

func TestCreateSocialProfileReturnsASolution(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	s, err := client.createSocialProfile(context.Background(), handleSocialProfileResponseMock, "AG-123", "SPG-444")
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

func TestCreateSocialProfileReturnsErrorIfGetFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Post Failed")}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.createSocialProfile(context.Background(), nil, "AG-123", "SPG-444")
	assert.EqualError(t, err, "Post Failed")
}

// registerSocialProfile tests
func TestRegisterSocialProfileReturnsAnErrorIfNoAccountGroupIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.registerSocialProfile(context.Background(), nil, "", "")
	assert.EqualError(t, err, "account group id is required")
}

func TestRegisterSocialProfileReturnsAnErrorIfNoSocialProfileIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.registerSocialProfile(context.Background(), nil, "AG-123", "")
	assert.EqualError(t, err, "social profile id is required")
}

func TestRegisterSocialProfileReturnsAnErrorIfTheResponseErrors(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.registerSocialProfile(context.Background(), handleSocialProfileRegistrationResponseErrorMock, "AG-123", "SCP-444")
	assert.EqualError(t, err, "Failed to handle response")
}

func TestRegisterSocialProfileReturnsASolution(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	s, err := client.registerSocialProfile(context.Background(), handleSocialProfileRegistrationResponseMock, "AG-123", "SCP-444")
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

func TestRegisterSocialProfileReturnsErrorIfGetFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Post Failed")}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.registerSocialProfile(context.Background(), nil, "AG-123", "SCP-444")
	assert.EqualError(t, err, "Post Failed")
}

// lookupSocialProfilesRegistrations
func TestLookupSocialProfileRegistrationsReturnsAnErrorIfNoSocialProfileIdIsProvided(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.lookupSocialProfilesRegistrations(context.Background(), nil, "")
	assert.EqualError(t, err, "social profile id is required")
}

func TestLookupSocialProfilesReturnsAnErrorIfTheResponseErrors(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.lookupSocialProfilesRegistrations(context.Background(), handleSocialProfileRegistrationsResponseErrorMock, "SCP-444")
	assert.EqualError(t, err, "Failed to handle response")
}

func TestLookupSocialProfilesReturnsASolution(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := SocialProfileClient{SDKClient: baseClient}
	s, err := client.lookupSocialProfilesRegistrations(context.Background(), handleSocialProfileRegistrationsResponseMock, "SCP-444")
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

func TestLookupSocialProfilesReturnsErrorIfGetFails(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: errors.New("Post Failed")}
	client := SocialProfileClient{SDKClient: baseClient}
	_, err := client.lookupSocialProfilesRegistrations(context.Background(), nil, "SCP-444")
	assert.EqualError(t, err, "Post Failed")
}
