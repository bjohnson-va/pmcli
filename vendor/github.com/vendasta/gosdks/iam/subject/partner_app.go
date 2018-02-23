package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToPartnerApp(s *subject) *PartnerApp {
	partnerApp := &PartnerApp{
		subject:   s,
	}

	return partnerApp
}

// PartnerApp represents a partner application service account.
type PartnerApp struct {
	*subject

	PartnerID string `attribute:"partner_id"`
}

func (s *PartnerApp) Context() *subjectcontext.Context {
	return subjectcontext.New("partner_app", s.PartnerID)
}

// NewPartnerApp creates a new PartnerApp struct
func NewPartnerApp(s *subject, PartnerID string) *PartnerApp {
	return &PartnerApp{
		subject:   s,
		PartnerID: PartnerID,
	}
}
