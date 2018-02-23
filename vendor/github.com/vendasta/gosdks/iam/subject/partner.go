package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToPartner(s *subject) *Partner {
	partner := &Partner{
		subject: s,
	}

	return partner
}

// Partner represents a Partner Center user.
type Partner struct {
	*subject

	PartnerID string `attribute:"partner_id"`
	RMUID     string `attribute:"rm_uid"`
	MSUID     string `attribute:"ms_uid"`
	NBUID     string `attribute:"nb_uid"`
	ARMUID    string `attribute:"arm_uid"`
	SMUID     string `attribute:"sm_uid"`

	AccessibleMarkets []string `attribute:"accessible_markets"`

	IsSuperAdmin bool `attribute:"is_super_admin"`
	IsAdmin      bool `attribute:"is_admin"`

	CanCustomizeWhitelabel bool `attribute:"can_customize_whitelabel"`
	CanAccessBilling       bool `attribute:"can_access_billing"`
	CanAccessAccounts      bool `attribute:"can_access_accounts"`
	CanAccessMarketing     bool `attribute:"can_access_marketing"`
	CanAccessSales         bool `attribute:"can_access_sales"`
	CanAccessConcierge     bool `attribute:"can_access_concierge"`
	CanAccessBrands        bool `attribute:"can_access_brands"`
	CanAccessDashboard     bool `attribute:"can_access_dashboard"`
	CanAccessOrders        bool `attribute:"can_access_orders"`
}

func (s *Partner) Context() *subjectcontext.Context {
	return subjectcontext.New("partner", "")
}

// NewPartner creates a new Partner struct
func NewPartner(s *subject, PartnerID, RMUID, MSUID, NBUID, ARMUID, SMUID string, AccessibleMarkets []string,
	IsSuperAdmin, IsAdmin, CanCustomizeWhitelabel, CanAccessBilling, CanAccessAccounts, CanAccessMarketing,
	CanAccessSales, CanAccessConcierge, CanAccessBrands, CanAccessDashboard bool) *Partner {
	return &Partner{
		subject:                s,
		PartnerID:              PartnerID,
		RMUID:                  RMUID,
		MSUID:                  MSUID,
		NBUID:                  NBUID,
		ARMUID:                 ARMUID,
		SMUID:                  SMUID,
		AccessibleMarkets:      AccessibleMarkets,
		IsSuperAdmin:           IsSuperAdmin,
		IsAdmin:                IsAdmin,
		CanCustomizeWhitelabel: CanCustomizeWhitelabel,
		CanAccessBilling:       CanAccessBilling,
		CanAccessAccounts:      CanAccessAccounts,
		CanAccessMarketing:     CanAccessMarketing,
		CanAccessSales:         CanAccessSales,
		CanAccessConcierge:     CanAccessConcierge,
		CanAccessBrands:        CanAccessBrands,
		CanAccessDashboard:     CanAccessDashboard,
	}
}
