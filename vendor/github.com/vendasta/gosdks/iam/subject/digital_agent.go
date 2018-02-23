package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToDigitalAgent(s *subject) *DigitalAgent {
	digitalAgent := &DigitalAgent{
		subject: s,
	}

	return digitalAgent
}

// DigitalAgent represents Concierge user.
type DigitalAgent struct {
	*subject

	FirstName            string `attribute:"first_name"`
	LastName             string `attribute:"last_name"`
	AccessiblePartnerIDs []string `attribute:"accessible_partner_ids"`
}

func (s *DigitalAgent) Context() *subjectcontext.Context {
	return subjectcontext.New(DigitalAgentSubjectType, "")
}

// NewDigitalAgent creates a new Digital Agent struct
func NewDigitalAgent(s *subject, FirstName string, LastName string, AccessiblePartnerIDs []string) *DigitalAgent {
	return &DigitalAgent{
		subject:              s,
		AccessiblePartnerIDs: AccessiblePartnerIDs,
		FirstName:            FirstName,
		LastName:             LastName,
	}
}
