package nap

import (
	"fmt"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/pb/nap/v1"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/vax"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewClient returns an NAP client.
func NewClient(ctx context.Context, e config.Env, dialOptions ...grpc.DialOption) (Interface, error) {
	address := addresses[e]
	if address == "" {
		return nil, fmt.Errorf("Unable to create client with environment %d", e)
	}
	conn, err := vax.NewGRPCConnection(ctx, address, e != config.Local, scopes[e], true, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &napClient{nap_v1.NewNAPDataServiceClient(conn)}, nil
}

type napClient struct {
	nap_v1.NAPDataServiceClient
}

func (nc *napClient) ParsePhoneNumber(ctx context.Context, phoneNumber string, opts ...ParseNumberOpt) (*PhoneNumberParseResult, error) {
	var res *nap_v1.ParsePhoneNumberResponse
	req := &nap_v1.ParsePhoneNumberRequest{
		Number: phoneNumber,
	}
	reqOpts := parseNumberOpts{}
	for _, opt := range opts {
		opt(&reqOpts)
	}
	reqOpts.fillProto(req)

	err := vax.Invoke(util.NewContext(ctx), func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		res, err = nc.NAPDataServiceClient.ParsePhoneNumber(ctx, req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}
	return &PhoneNumberParseResult{
		ParseResult:      FromParseResultProto(res.GetParseResult()),
		ValidationResult: validationResultFromProto(res.ValidationResult),
		FormatResult:     FromFormatResultProto(res.GetFormatResult()),
		MetaDataResult:   FromMetaDataResultProto(res.MetadataResult),
	}, nil
}

func (nc *napClient) ListCountries(ctx context.Context) (*ListCountriesResult, error) {
	var res *nap_v1.ListCountriesResponse
	req := &nap_v1.ListCountriesRequest{}

	err := vax.Invoke(util.NewContext(ctx), func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		res, err = nc.NAPDataServiceClient.ListCountries(ctx, req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	countries := make([]Country, len(res.GetCountries()))
	for i, resCountry := range res.GetCountries() {
		countries[i] = Country{
			ID:   resCountry.GetId(),
			Name: resCountry.GetName(),
		}
	}
	return &ListCountriesResult{
		Countries: countries,
	}, nil
}

func (nc *napClient) ListStates(ctx context.Context, countryID string) (*ListStatesResult, error) {
	var res *nap_v1.ListStatesResponse
	req := &nap_v1.ListStatesRequest{
		CountryId: countryID,
	}

	err := vax.Invoke(util.NewContext(ctx), func(ctx context.Context, cs vax.CallSettings) error {
		var err error
		res, err = nc.NAPDataServiceClient.ListStates(ctx, req, grpc.FailFast(false))
		return err
	}, defaultRetryCallOptions)

	if err != nil {
		return nil, err
	}

	states := make([]State, len(res.GetStates()))
	for i, resState := range res.GetStates() {
		states[i] = State{
			ID:   resState.GetId(),
			Name: resState.GetName(),
		}
	}
	return &ListStatesResult{
		States: states,
	}, nil
}

// ValidateCountryState checks if provided country/state ids are valid as a combination
// Country is required, and if a country has states, state is required
func (nc *napClient) ValidateCountryState(ctx context.Context, countryID string, stateID string) error {
	if countryID == "" {
		return util.Error(util.InvalidArgument, "Country is required")
	}

	res, err := nc.ListStates(ctx, countryID)
	if err != nil {
		resStatus, _ := status.FromError(err)
		if resStatus.Code() == codes.InvalidArgument {
			return util.Error(util.InvalidArgument, "Country %s was invalid", countryID)
		}
		logging.Debugf(ctx, "Error from ListStates: %s, %s", resStatus.Code(), resStatus.Message())
		return util.Error(util.Internal, "Internal Error")
	}

	if len(res.States) == 0 {
		// When no states are present, there may still be valid subregions according to ISO-3166-2. Just accepting...
		return nil
	}
	if len(res.States) > 0 && stateID == "" {
		return util.Error(util.InvalidArgument, "State is required for Country %s", countryID)
	}

	stateValid := false
	for _, state := range res.States {
		if state.ID == stateID {
			stateValid = true
			break
		}
	}
	if !stateValid {
		return util.Error(util.InvalidArgument, "State %s is invalid for Country %s", stateID, countryID)
	}

	return nil
}

// ValidateCountry checks if provided country id is valid
func (nc *napClient) ValidateCountry(ctx context.Context, countryID string) error {
	if countryID == "" {
		return util.Error(util.InvalidArgument, "Country is required")
	}

	_, err := nc.ListStates(ctx, countryID)
	if err != nil {
		return util.Error(util.InvalidArgument, "Country %s was invalid", countryID)
	}

	return nil
}
