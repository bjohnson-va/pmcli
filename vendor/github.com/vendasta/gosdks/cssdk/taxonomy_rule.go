package cssdk

import (
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

// validTaxIDs Implements validation.Rule
type validTaxIDs struct {
	taxIDs    []string
	errorType util.ErrorType
	message   string
	client    TaxonomyClientInterface
	ctx       context.Context
}

// Validate that the tax ids are correct
func (r *validTaxIDs) Validate() error {
	taxonomies, err := r.client.List(r.ctx)
	if err != nil {
		return err
	}
	for _, taxID := range r.taxIDs {
		if !TaxonomySet(taxonomies).IsValid(taxID) {
			return util.Error(r.errorType, r.message)
		}
	}
	return nil
}

func TaxonomyIDValidationRule(client TaxonomyClientInterface, ctx context.Context, taxonomyIDs []string, errorType util.ErrorType, message string) validation.Rule {
	return &validTaxIDs{
		taxIDs:    taxonomyIDs,
		errorType: errorType,
		message:   message,
		client:    client,
		ctx:       ctx,
	}
}
