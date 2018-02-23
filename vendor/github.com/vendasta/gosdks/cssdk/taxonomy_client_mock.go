package cssdk

import (
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"golang.org/x/net/context"
)

// BuildMockTaxonomyClient creates a mock taxonomy client.
func BuildMockTaxonomyClient() TaxonomyClientInterface {
	return &MockTaxonomyClient{}
}

// MockTaxonomyClient Implements the TaxonomyClientInterface
type MockTaxonomyClient struct {
}

//List returns a list of all the taxonomies that Mock Core Services knows about
func (c *MockTaxonomyClient) List(ctx context.Context) ([]*Taxonomy, error) {
	return []*Taxonomy{}, nil
}

func (nc *MockTaxonomyClient) TaxonomyIDValidationRule(ctx context.Context, taxonomyIDs []string, errorType util.ErrorType, message string) validation.Rule {
	return &mockValidTaxIDs{}
}

// mockValidTaxIDs Implements validation.Rule
type mockValidTaxIDs struct {
}

// Validate that the tax ids are correct
func (r *mockValidTaxIDs) Validate() error {
	return nil
}
