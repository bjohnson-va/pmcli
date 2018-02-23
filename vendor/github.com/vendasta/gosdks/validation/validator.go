package validation

import (
	"strings"
	"errors"
)

// Rule is a struct whose state can be evaluated with a single function.
// Consider adding any values that need to be investigated for correctness as properties on your own Rule
type Rule interface {
	Validate() error
}

// Validator provides a mechanism for chaining individual ValidationRules together and evaluating all of the rules easily
// Rules are validated in the order they are added to the validator.
type Validator struct {
	rules []Rule
}

// NewValidator returns a fresh Validator devoid of any rules.
// Use this to create a new set of rules to validate like this:
// err := NewValidator().
//     Rule(ListingIDRequired(listingID)).
//     Rule(PageSizeWithinBounds(pageSize)).
//     Validate()
// if err != nil {
//    ...
// }
func NewValidator() *Validator {
	return &Validator{}
}

// Rule adds a ValidationRule to the Validator
func (c *Validator) Rule(r ...Rule) *Validator {
	c.rules = append(c.rules, r...)
	return c
}

// Validate checks each rule to see if it has been validated.
// It will process the rules in order and throw the first error thrown by a rule.
func (c *Validator) Validate() error {
	var err error
	for _, r := range c.rules {
		err = r.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidateAndJoinErrors performs similarly to Validate() above, but will return all failures in a single error message.
func (c *Validator) ValidateAndJoinErrors() error {
	messages := []string{}
	var err error
	for _, r := range c.rules {
		err = r.Validate()
		if err != nil {
			messages = append(messages, err.Error())
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n"))
	}
	return nil
}