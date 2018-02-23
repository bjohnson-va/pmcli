package structuredattributes

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/iam/structuredattributes/internal"
)

// Unmarshal extracts values from the StructAttribute proto onto the interface specified.
// Example interface:
// type MyAttributes struct {
//     Name         string   `attribute:"name"`
//     Version      int      `attribute:"version"`
//     Associations []string `attribute:"associations"`
// }
func Unmarshal(to interface{}, sa *iam_v1.StructAttribute) error {
	return internal.StructuredAttributeToModel(to, sa)
}
