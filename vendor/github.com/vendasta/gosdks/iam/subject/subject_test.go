package subject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/iam/attribute"
	"github.com/vendasta/gosdks/iam/subjectcontext"
	"github.com/vendasta/gosdks/pb/iam/v1"
)

type testSubject struct {
	*subject

	PartnerID         string   `attribute:"partner_id"`
	AccessibleMarkets []string `attribute:"accessible_markets"`
	IsSuperAdmin      bool     `attribute:"is_super_admin"`
}

func (s *testSubject) Context() *subjectcontext.Context {
	return subjectcontext.New("test_subject", "")
}

func Test_BackwardsCompatibleAttributes(t *testing.T) {
	cases := []struct {
		name   string
		input  *iam_v1.Subject
		output testSubject
		error  error
	}{
		{
			name: "Uses value from Attributes if attribute is not present in StructuredAttributes",
			input: &iam_v1.Subject{
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "partner_id",
						Values: []string{"ABC"},
					},
				},
			},
			output: testSubject{
				PartnerID: "ABC",
			},
		},
		{
			name: "Uses value from StructuredAttributes over values from Attributes",
			input: &iam_v1.Subject{
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "partner_id",
						Values: []string{"ABC"},
					},
				},
				StructAttributes: &iam_v1.StructAttribute{
					Attributes: map[string]*iam_v1.Attribute{
						"partner_id": attribute.String("DEF"),
					},
				},
			},
			output: testSubject{
				PartnerID: "DEF",
			},
		},
		{
			name: "Uses mixture of values from both StructuredAttributes and Attributes",
			input: &iam_v1.Subject{
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "partner_id",
						Values: []string{"ABC"},
					},
				},
				StructAttributes: &iam_v1.StructAttribute{
					Attributes: map[string]*iam_v1.Attribute{
						"accessible_markets": {
							Kind: &iam_v1.Attribute_ListAttribute{
								ListAttribute: &iam_v1.ListAttribute{
									Attributes: []*iam_v1.Attribute{
										attribute.String("market-1"),
										attribute.String("market-2"),
									},
								},
							},
						},
					},
				},
			},
			output: testSubject{
				PartnerID:         "ABC",
				AccessibleMarkets: []string{"market-1", "market-2"},
			},
		},
		{
			name: "Converts string from Attributes to bool",
			input: &iam_v1.Subject{
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "is_super_admin",
						Values: []string{"true"},
					},
				},
			},
			output: testSubject{
				IsSuperAdmin: true,
			},
		},
		{
			name: "Attributes doesn't overwrite bools from StructuredAttributes",
			input: &iam_v1.Subject{
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "is_super_admin",
						Values: []string{"true"},
					},
				},
				StructAttributes: &iam_v1.StructAttribute{
					Attributes: map[string]*iam_v1.Attribute{
						"is_super_admin": attribute.Bool(false),
					},
				},
			},
			output: testSubject{
				IsSuperAdmin: false,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := testSubject{}

			err := populateAttributes(&actual, c.input)
			assert.Equal(t, c.error, err)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_Subject_New(t *testing.T) {
	cases := []struct {
		name    string
		context *subjectcontext.Context
		subject *iam_v1.Subject
		output  Subject
		error   error
	}{
		{
			name:    "Partner is created",
			context: subjectcontext.New(PartnerSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToPartner(&subject{
				Subject: &iam_v1.Subject{
					SubjectId: "UID-123",
					Email:     "email@email.com",
				},
			}),
		},
		{
			name:    "SMB is created",
			context: subjectcontext.New(SMBSubjectType, "ABC"),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToSMB(
				subjectcontext.New(SMBSubjectType, "ABC"),
				&subject{
					Subject: &iam_v1.Subject{
						SubjectId: "UID-123",
						Email:     "email@email.com",
					},
				},
			),
		},
		{
			name:    "Salesperson is created",
			context: subjectcontext.New(SalesPersonSubjectType, "ABC"),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToSalesPerson(
				subjectcontext.New(SalesPersonSubjectType, "ABC"),
				&subject{
					Subject: &iam_v1.Subject{
						SubjectId: "UID-123",
						Email:     "email@email.com",
					},
				},
			),
		},
		{
			name:    "PartnerApp is created",
			context: subjectcontext.New(PartnerAppSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToPartnerApp(&subject{
				Subject: &iam_v1.Subject{
					SubjectId: "UID-123",
					Email:     "email@email.com",
				},
			}),
		},
		{
			name:    "Vendor is created",
			context: subjectcontext.New(VendorSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToVendor(&subject{
				Subject: &iam_v1.Subject{
					SubjectId: "UID-123",
					Email:     "email@email.com",
				},
			}),
		},
		{
			name:    "Digital Agent is created",
			context: subjectcontext.New(DigitalAgentSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
			},
			output: ToDigitalAgent(&subject{
				Subject: &iam_v1.Subject{
					SubjectId: "UID-123",
					Email:     "email@email.com",
				},
			}),
		},
		{
			name:    "Legacy attributes are unmarshaled",
			context: subjectcontext.New(PartnerSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId: "UID-123",
				Email:     "email@email.com",
				Attributes: []*iam_v1.MultiValueAttribute{
					{
						Key:    "partner_id",
						Values: []string{"ABC"},
					},
				},
			},
			output: NewPartner(
				&subject{
					Subject: &iam_v1.Subject{
						SubjectId: "UID-123",
						Email:     "email@email.com",
						Attributes: []*iam_v1.MultiValueAttribute{
							{
								Key:    "partner_id",
								Values: []string{"ABC"},
							},
						},
					},
				},
				"ABC", "", "", "", "", "",
				nil,
				false, false, false, false, false,
				false, false, false, false, false,
			),
		},
		{
			name:    "StructAttributes are unmarshaled",
			context: subjectcontext.New(PartnerSubjectType, ""),
			subject: &iam_v1.Subject{
				SubjectId:        "UID-123",
				Email:            "email@email.com",
				StructAttributes: attribute.NewBuilder().String("partner_id", "ABC").Build(),
			},
			output: NewPartner(
				&subject{
					Subject: &iam_v1.Subject{
						SubjectId:        "UID-123",
						Email:            "email@email.com",
						StructAttributes: attribute.NewBuilder().String("partner_id", "ABC").Build(),
					},
				},
				"ABC", "", "", "", "", "",
				nil,
				false, false, false, false, false,
				false, false, false, false, false,
			),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := New(c.context, c.subject)
			assert.Equal(t, c.error, err)
			assert.Equal(t, c.output, actual)
		})
	}
}
