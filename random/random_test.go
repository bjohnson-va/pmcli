package random_test

import (
	"testing"
	"github.com/bjohnson-va/pmcli/random"
	"github.com/magiconair/properties/assert"
)

func  TestProvidesConsistentValuesForSameBreadcrumb(t *testing.T) {
	provider := random.BreadcrumbBasedFieldProvider()
	breadcrumb := "a.b.c"
	v1 := provider.NewString(breadcrumb)
	v2 := provider.NewString(breadcrumb)
	assert.Equal(t, v1, v2)
}

func TestProvidesDifferentValuesForDifferentBreadcrumbs(t *testing.T) {
	provider := random.BreadcrumbBasedFieldProvider()
	breadcrumb := "a.b.c"
	breadcrumb2 := "a.b.d"
	v1 := provider.NewString(breadcrumb)
	v2 := provider.NewString(breadcrumb2)
	assert.Equal(t, v1 == v2, false)
}
