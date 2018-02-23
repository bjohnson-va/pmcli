package structuredattributes

import (
	"github.com/vendasta/gosdks/iam/attribute"
	"github.com/vendasta/gosdks/iam/structuredattributes/internal"
)

// Deprecated: Use structured attributes instead
// UnmarshalLegacy attributes onto the to object
func UnmarshalLegacy(to interface{}, sa []*attribute.Attribute) error {
	return internal.LegacyAttributeToModel(to, sa)
}
