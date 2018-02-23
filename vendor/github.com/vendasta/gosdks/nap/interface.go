package nap

import (
	"github.com/vendasta/gosdks/pb/nap/v1"
	"golang.org/x/net/context"
)

// PhoneNumberParseResult is the result of the parse
type PhoneNumberParseResult struct {
	*ParseResult
	*ValidationResult
	*FormatResult
	*MetaDataResult
}

// ParseResult contains parse information for a specific phone number
type ParseResult struct {
	// The unformatted national representation of the number
	NationalNumber string
	// The extension of the number
	Extension string
	// The country code of the number
	CountryCode string
	// The source of the country code
	CountryCodeSource nap_v1.CountryCodeSource
}

// FromParseResultProto converts nap_v1.ParseResult to ParseResult
func FromParseResultProto(p *nap_v1.ParseResult) *ParseResult {
	return &ParseResult{
		NationalNumber: p.GetNationalNumber(),
		Extension:      p.GetExtension(),
		CountryCode:    p.GetCountryCode(),
	}
}

// CountryCodeSource enumerates the sources from which the country code was determined
type CountryCodeSource int32

// CountryCodeSource enum
const (
	FromDefaultCountry        CountryCodeSource = 0
	FromNumberWithIdd         CountryCodeSource = 1
	FromNumberWithPlusSign    CountryCodeSource = 2
	FromNumberWithoutPlusSign CountryCodeSource = 3
)

// PhoneNumberFormat enumerates the available phone number formats
type PhoneNumberFormat int32

// PhoneNumberFormat enum
const (
	NATIONAL      PhoneNumberFormat = 0
	INTERNATIONAL PhoneNumberFormat = 1
	E164          PhoneNumberFormat = 2
	RFC3966       PhoneNumberFormat = 3
)

// PhoneNumberType enumerates the available phone number types
type PhoneNumberType int32

// PhoneNumberType enum
const (
	FixedLine         PhoneNumberType = 0
	FixedLineOrMobile PhoneNumberType = 1
	Mobile            PhoneNumberType = 2
	Pager             PhoneNumberType = 3
	PersonalNumber    PhoneNumberType = 4
	PremiumRate       PhoneNumberType = 5
	SharedCost        PhoneNumberType = 6
	TollFree          PhoneNumberType = 7
	UAN               PhoneNumberType = 8
	Unknown           PhoneNumberType = 9
	Voicemail         PhoneNumberType = 10
	VOIP              PhoneNumberType = 11
)

// ValidationResult contains validation information about a phone number.
type ValidationResult struct {
	// If the number is possible. This is the most lenient and fast validation, always performed
	IsPossible bool
	// If the number is valid, if requested
	IsValid *bool
	// If the number is valid for the specified region, if requested
	IsValidForRegion *bool
}

func validationResultFromProto(p *nap_v1.ValidationResult) *ValidationResult {
	if p == nil {
		return nil
	}
	vr := &ValidationResult{}
	vr.IsPossible = p.IsPossible
	if p.IsValid != nil {
		isValid := p.IsValid.Value
		vr.IsValid = &isValid
	}
	if p.IsValidForRegion != nil {
		isValidForRegion := p.IsValidForRegion.Value
		vr.IsValidForRegion = &isValidForRegion
	}
	return vr
}

// FormatResult contains different formatted phone numbers.
type FormatResult struct {
	National      string
	International string
	E164          string
	RFC3966       string
}

// FromFormatResultProto converts nap_v1.FormatResult to FormatResult
func FromFormatResultProto(fr *nap_v1.FormatResult) *FormatResult {
	return &FormatResult{
		National:      fr.GetNational(),
		International: fr.GetInternational(),
		E164:          fr.GetE164(),
		RFC3966:       fr.GetRFC3966(),
	}
}

// MetaDataResult contains miscellaneous information about the phone number.
type MetaDataResult struct {
	// The approximate timezones relevant to this number
	Timezones []string
	// A description of where the number is located -- may contain varying levels of specificity from country to city
	LocationDescription string
	// The type of number -- mobile, landline, etc
	Type nap_v1.PhoneNumberType
}

// FromMetaDataResultProto converts nap_v1.MetaDataResult to MetaDataResult
func FromMetaDataResultProto(mdr *nap_v1.MetaDataResult) *MetaDataResult {
	return &MetaDataResult{
		Timezones:           mdr.GetTimezones(),
		LocationDescription: mdr.GetLocationDescription(),
		Type:                mdr.GetType(),
	}
}

type parseNumberOpts struct {
	DefaultRegion string
	Formats       []PhoneNumberFormat
}

func (pno *parseNumberOpts) fillProto(req *nap_v1.ParsePhoneNumberRequest) {
	f := make([]nap_v1.PhoneNumberFormat, len(pno.Formats))
	for n, format := range pno.Formats {
		f[n] = nap_v1.PhoneNumberFormat(format)
	}
	req.DefaultRegion = pno.DefaultRegion
	req.Formats = f
}

// ParseNumberOpt provides optional arguments to ParsePhoneNumber
type ParseNumberOpt func(*parseNumberOpts)

// Formats sets the formats which should be returned
func Formats(formats ...PhoneNumberFormat) ParseNumberOpt {
	return func(p *parseNumberOpts) {
		p.Formats = formats
	}
}

// DefaultRegion provides insight into parsing/validating the phone number. Is required if doing regional validation.
// If an empty string is provided, the region "US" will be used as a sensible default
func DefaultRegion(defaultRegion string) ParseNumberOpt {
	return func(p *parseNumberOpts) {
		p.DefaultRegion = defaultRegion
		if p.DefaultRegion == "" {
			p.DefaultRegion = "US"
		}
	}
}

// Country represents a country
type Country struct {
	// ISO-3166-1 country code
	ID   string
	Name string
}

// State represents a state
type State struct {
	// ISO-3166-2 code (eg. "SK", not "CA-SK")
	ID   string
	Name string
}

// ListCountriesResult contains the list of countries
type ListCountriesResult struct {
	Countries []Country
}

// ListStatesResult contains the list of states
type ListStatesResult struct {
	States []State
}

// Interface defines the methods available for NAP validation.
type Interface interface {
	ParsePhoneNumber(ctx context.Context, phoneNumber string, opts ...ParseNumberOpt) (*PhoneNumberParseResult, error)

	ListCountries(ctx context.Context) (*ListCountriesResult, error)
	ListStates(ctx context.Context, countryID string) (*ListStatesResult, error)
	ValidateCountryState(ctx context.Context, countryID string, stateID string) error
	ValidateCountry(ctx context.Context, countryID string) error
}
