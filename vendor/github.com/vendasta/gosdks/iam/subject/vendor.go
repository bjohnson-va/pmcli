package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToVendor(s *subject) *Vendor {
	vendor := &Vendor{
		subject: s,
	}

	return vendor
}

// Partner represents a Partner Center user.
type Vendor struct {
	*subject

	FirstName  string   `attribute:"first_name"`
	LastName   string   `attribute:"last_name"`
	VendorName string   `attribute:"vendor_name"`
	AppIDs     []string `attribute:"app_ids"`
}

func (s *Vendor) Context() *subjectcontext.Context {
	return subjectcontext.New("vendor", "")
}

// NewVendor creates a new Vendor struct
func NewVendor(s *subject, FirstName, LastName, VendorName string, AppIDs []string) *Vendor {
	return &Vendor{
		subject:   s,
		FirstName: FirstName,
		LastName: LastName,
		VendorName: VendorName,
		AppIDs: AppIDs,
	}
}
