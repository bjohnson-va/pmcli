package nap

import (
	"context"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
)

// validPhoneNumber Implements validation.Rule
type validPhoneNumber struct {
	phoneNumber *string
	errorType   util.ErrorType
	message     string
	opts        []ParseNumberOpt
	client      Interface
	ctx         context.Context
}

// Validate that the phone number is correct
func (r *validPhoneNumber) Validate() error {
	res, err := r.client.ParsePhoneNumber(r.ctx, *r.phoneNumber, r.opts...)
	if err != nil || !res.IsPossible {
		return util.Error(r.errorType, r.message)
	}
	*r.phoneNumber = res.ParseResult.NationalNumber
	return nil
}

// PhoneNumberValidationRule ensures that the number can be parsed
func PhoneNumberValidationRule(ctx context.Context, client Interface, phoneNumber *string, opts []ParseNumberOpt, errorType util.ErrorType, message string) validation.Rule {
	return &validPhoneNumber{
		phoneNumber: phoneNumber,
		errorType:   errorType,
		message:     message,
		opts:        opts,
		client:      client,
		ctx:         ctx,
	}
}

type countryStateRule struct {
	countryID string
	stateID   string
	client    Interface
	ctx       context.Context
	// Detailed messages provided in errors
}

func (r *countryStateRule) Validate() error {
	return r.client.ValidateCountryState(r.ctx, r.countryID, r.stateID)

}

// CountryStateRule validates provided ISO-3166 ids for countries and states -- country is required
func CountryStateRule(ctx context.Context, client Interface, countryID, stateID string) validation.Rule {
	return &countryStateRule{
		ctx:       ctx,
		client:    client,
		countryID: countryID,
		stateID:   stateID,
	}
}

type countryRule struct {
	countryID string
	client    Interface
	ctx       context.Context
	// Detailed messages provided in errors
}

func (r *countryRule) Validate() error {
	return r.client.ValidateCountry(r.ctx, r.countryID)

}

// CountryRule validates provided ISO-3166 id for country -- country is required
func CountryRule(ctx context.Context, client Interface, countryID string) validation.Rule {
	return &countryRule{
		ctx:       ctx,
		client:    client,
		countryID: countryID,
	}
}
