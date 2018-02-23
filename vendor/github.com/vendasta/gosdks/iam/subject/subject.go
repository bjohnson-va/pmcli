package subject

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/iam/attribute"
	"github.com/vendasta/gosdks/iam/structuredattributes"
	"github.com/vendasta/gosdks/iam/subjectcontext"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/util"
)

const (
	PartnerSubjectType      = "partner"
	SMBSubjectType          = "smb"
	SalesPersonSubjectType  = "sales_person"
	PartnerAppSubjectType   = "partner_app"
	VendorSubjectType       = "vendor"
	DigitalAgentSubjectType = "digital_agent"
)

func New(c *subjectcontext.Context, s *iam_v1.Subject) (Subject, error) {
	if c == nil || s == nil {
		return nil, util.Error(util.NotFound, "Subject not found.")
	}
	var subj Subject
	switch c.Type {
	case PartnerSubjectType:
		subj = ToPartner(&subject{Subject: s})
	case SMBSubjectType:
		subj = ToSMB(c, &subject{Subject: s})
	case SalesPersonSubjectType:
		subj = ToSalesPerson(c, &subject{Subject: s})
	case PartnerAppSubjectType:
		subj = ToPartnerApp(&subject{Subject: s})
	case VendorSubjectType:
		subj = ToVendor(&subject{Subject: s})
	case DigitalAgentSubjectType:
		subj = ToDigitalAgent(&subject{Subject: s})
	}
	if subj == nil {
		return nil, util.Error(util.InvalidArgument, "Invalid subject type %s.", c.Type)
	}
	err := populateAttributes(subj, s)
	if err != nil {
		return nil, err
	}
	return subj, nil
}

func populateAttributes(s interface{}, su *iam_v1.Subject) error {
	if su.Attributes != nil {
		err := structuredattributes.UnmarshalLegacy(s, buildLegacyAttributes(su))
		if err != nil {
			return err
		}
	}
	if su.StructAttributes != nil {
		err := structuredattributes.Unmarshal(s, su.StructAttributes)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildLegacyAttributes(s *iam_v1.Subject) []*attribute.Attribute {
	attrs := make([]*attribute.Attribute, len(s.Attributes))
	for i, attr := range s.Attributes {
		attrs[i] = attribute.NewLegacy(attr.Key, attr.Values)
	}
	return attrs
}

// Subject is the base interface that all user types implement.
type Subject interface {
	Context() *subjectcontext.Context
	SubjectID() string
	Email() string
	Created() time.Time
	Updated() time.Time
	LastLogin() time.Time
	Attributes() map[string]*iam_v1.Attribute
}

type subject struct {
	*iam_v1.Subject
}

func (s *subject) SubjectID() string {
	return s.Subject.SubjectId
}

func (s *subject) Email() string {
	return s.Subject.Email
}

func (s *subject) Created() time.Time {
	var t time.Time
	if s.Subject.Created == nil {
		return t
	}
	t, _ = ptypes.Timestamp(s.Subject.Created)
	return t
}

func (s *subject) Updated() time.Time {
	var t time.Time
	if s.Subject.Updated == nil {
		return t
	}
	t, _ = ptypes.Timestamp(s.Subject.Updated)
	return t
}

func (s *subject) LastLogin() time.Time {
	var t time.Time
	if s.Subject.LastLogin == nil {
		return t
	}
	t, _ = ptypes.Timestamp(s.Subject.LastLogin)
	return t
}

// Attributes should not be used to retrieve specific attributes
// You should only use this if you don't know what persona you have and you are looking for dynamic attributes
// You should be casting the subject to the specific Persona type you want
func (s *subject) Attributes() map[string]*iam_v1.Attribute {
	return s.StructAttributes.GetAttributes()
}
