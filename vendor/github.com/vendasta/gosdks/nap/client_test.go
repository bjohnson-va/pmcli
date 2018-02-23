package nap

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/nap/v1"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type mockNapClient struct {
	ListStatesResult *nap_v1.ListStatesResponse
	ListStatesError  error
}

func (mnc *mockNapClient) ParsePhoneNumber(ctx context.Context, in *nap_v1.ParsePhoneNumberRequest, opts ...grpc.CallOption) (*nap_v1.ParsePhoneNumberResponse, error) {
	return nil, nil
}
func (mnc *mockNapClient) ListCountries(ctx context.Context, in *nap_v1.ListCountriesRequest, opts ...grpc.CallOption) (*nap_v1.ListCountriesResponse, error) {
	return nil, nil
}
func (mnc *mockNapClient) ListStates(ctx context.Context, in *nap_v1.ListStatesRequest, opts ...grpc.CallOption) (*nap_v1.ListStatesResponse, error) {
	return mnc.ListStatesResult, mnc.ListStatesError
}

func TestValidateCountryState(t *testing.T) {
	type test struct {
		countryID        string
		stateID          string
		listStatesResult *nap_v1.ListStatesResponse
		listStatesError  error
		expected         error
		description      string
	}

	canadaStateResults := &nap_v1.ListStatesResponse{
		States: []*nap_v1.State{
			{
				Id:   "SK",
				Name: "Saskatchewan",
			},
		},
	}
	unitedKingdomStateResults := &nap_v1.ListStatesResponse{
		States: []*nap_v1.State{},
	}
	invalidCountryStateError := status.Error(codes.InvalidArgument, "Country ID is invalid")
	internalCountryStateError := status.Error(codes.Internal, "broked")

	cases := []*test{
		{
			countryID:   "",
			stateID:     "",
			expected:    util.Error(util.InvalidArgument, "Country is required"),
			description: "Country required",
		},
		{
			countryID:       "JUNK",
			stateID:         "",
			listStatesError: invalidCountryStateError,
			expected:        util.Error(util.InvalidArgument, "Country JUNK was invalid"),
			description:     "Invalid country",
		},
		{
			countryID:        "GB",
			stateID:          "UKM",
			listStatesResult: unitedKingdomStateResults,
			expected:         nil,
			description:      "Country without states accept any state",
		},
		{
			countryID:        "GB",
			stateID:          "",
			listStatesResult: unitedKingdomStateResults,
			expected:         nil,
			description:      "Country without states does not require state",
		},
		{
			countryID:        "CA",
			stateID:          "",
			listStatesResult: canadaStateResults,
			expected:         util.Error(util.InvalidArgument, "State is required for Country CA"),
			description:      "Country with states require state",
		},
		{
			countryID:        "CA",
			stateID:          "SK",
			listStatesResult: canadaStateResults,
			expected:         nil,
			description:      "Valid state for country",
		},
		{
			countryID:        "CA",
			stateID:          "XX",
			listStatesResult: canadaStateResults,
			expected:         util.Error(util.InvalidArgument, "State XX is invalid for Country CA"),
			description:      "Invalid state for country",
		},
		{
			countryID:       "CA",
			stateID:         "SK",
			listStatesError: internalCountryStateError,
			expected:        util.Error(util.Internal, "Internal Error"),
			description:     "Internal error from nap service",
		},
	}

	client := napClient{
		NAPDataServiceClient: &mockNapClient{},
	}
	for _, c := range cases {
		client.NAPDataServiceClient.(*mockNapClient).ListStatesResult = c.listStatesResult
		client.NAPDataServiceClient.(*mockNapClient).ListStatesError = c.listStatesError
		err := client.ValidateCountryState(context.Background(), c.countryID, c.stateID)
		assert.Equal(t, err, c.expected, c.description)
	}
}

func TestValidateCountry(t *testing.T) {
	type test struct {
		countryID        string
		listStatesResult *nap_v1.ListStatesResponse
		listStatesError  error
		expected         error
		description      string
	}

	canadaStateResults := &nap_v1.ListStatesResponse{
		States: []*nap_v1.State{
			{
				Id:   "SK",
				Name: "Saskatchewan",
			},
		},
	}
	invalidCountryStateError := status.Error(codes.InvalidArgument, "Country ID is invalid")

	cases := []*test{
		{
			countryID:   "",
			expected:    util.Error(util.InvalidArgument, "Country is required"),
			description: "Country required",
		},
		{
			countryID:       "JUNK",
			listStatesError: invalidCountryStateError,
			expected:        util.Error(util.InvalidArgument, "Country JUNK was invalid"),
			description:     "Invalid country",
		},
		{
			countryID:        "CA",
			listStatesResult: canadaStateResults,
			expected:         nil,
			description:      "Valid country",
		},
	}

	client := napClient{
		NAPDataServiceClient: &mockNapClient{},
	}
	for _, c := range cases {
		client.NAPDataServiceClient.(*mockNapClient).ListStatesResult = c.listStatesResult
		client.NAPDataServiceClient.(*mockNapClient).ListStatesError = c.listStatesError
		err := client.ValidateCountry(context.Background(), c.countryID)
		assert.Equal(t, err, c.expected, c.description)
	}
}
