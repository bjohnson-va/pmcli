package nap

import (
	"golang.org/x/net/context"
	"github.com/vendasta/gosdks/pb/nap/v1"
)

// NewMockClient returns a mock NAP client.
func NewMockClient() Interface {
	return &napMockClient{}
}

type napMockClient struct {
}

func (nc *napMockClient) ParsePhoneNumber(ctx context.Context, phoneNumber string, opts ...ParseNumberOpt) (*PhoneNumberParseResult, error) {
	return &PhoneNumberParseResult{
		ParseResult: &ParseResult{NationalNumber: phoneNumber, Extension: "", CountryCode: "1"},
		ValidationResult: &ValidationResult{IsPossible: true, IsValid: nil, IsValidForRegion: nil},
		FormatResult: &FormatResult{National: "", International: "", E164: "", RFC3966: ""},
		MetaDataResult: &MetaDataResult{Timezones: []string{"CST"}, LocationDescription: "", Type: nap_v1.PhoneNumberType_MOBILE},
	}, nil
}

func (nc *napMockClient) ListCountries(ctx context.Context) (*ListCountriesResult, error) {
	return nil, nil
}
func (nc *napMockClient) ListStates(ctx context.Context, countryID string) (*ListStatesResult, error) {
	return nil, nil
}
func (nc *napMockClient) ValidateCountryState(ctx context.Context, countryID string, stateID string) error {
	return nil
}
func (nc *napMockClient) ValidateCountry(ctx context.Context, countryID string) error {
	return nil
}
