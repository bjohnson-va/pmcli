package iam

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/vendasta/gosdks/iam/attribute"
)

func Test_NewPolicy(t *testing.T) {
	smbPolicy := NewPolicyNode()
	smbPolicy.Operator(
		And(
			SubjectValueIntersection("subject_type", attribute.String("smb")),
			SubjectResourceIntersection("subject_id"),
			SubjectResourceIntersection("partner_id"),
		),
	)

	partnerPolicy := NewPolicyNode()
	partnerPolicy.Operator(
		And(
			SubjectValueIntersection("subject_type", attribute.String("partner")),
			SubjectResourceIntersection("partner_id"),
			SubjectResourceIntersection("market_id"),
		),
	)
	accountGroupFullPermissionsPolicy := NewPolicyNode()
	accountGroupFullPermissionsPolicy.Operator(
		Or(
			smbPolicy,
			partnerPolicy,
		),
	)
	p := NewPolicy(
		"vbc-prod",
		"account-group",
		"full-permission",
		"Full Permissions to AccountGroup.",
		accountGroupFullPermissionsPolicy,
		READ,
		WRITE,
		DELETE,
	)
	assert.Equal(t, p.AppID, "vbc-prod")
	assert.Equal(t, p.ResourceID, "account-group")
	assert.Equal(t, p.PolicyID, "full-permission")
	assert.Equal(t, p.PolicyName, "Full Permissions to AccountGroup.")
	assert.Equal(t, p.Operations, []AccessScope{READ, WRITE, DELETE})
	assert.Equal(t, p.Policy, accountGroupFullPermissionsPolicy)
}

func Test_PolicyNode_SubjectValueIntersectionClause(t *testing.T) {
	p := NewPolicyNode()
	err := p.SubjectValueIntersection(nil)
	assert.NotNil(t, err)

	err = p.SubjectValueIntersection(&PolicyNode{})
	assert.NotNil(t, err)

	p.SubjectResourceIntersectionClause(SubjectResourceIntersection("test"))

	err = p.SubjectValueIntersection(SubjectValueIntersection("test", attribute.String("value")))
	assert.NotNil(t, err)
}

func Test_PolicyNode_SubjectResourceIntersectionClause(t *testing.T) {
	p := NewPolicyNode()
	err := p.SubjectResourceIntersectionClause(nil)
	assert.NotNil(t, err)

	err = p.SubjectResourceIntersectionClause(&PolicyNode{})
	assert.NotNil(t, err)

	err = p.SubjectValueIntersection(SubjectValueIntersection("test", attribute.String("value")))
	assert.Nil(t, err)

	err = p.SubjectResourceIntersectionClause(SubjectResourceIntersection("test"))
	assert.NotNil(t, err)
}

func Test_PolicyNode_SubjectResourceMappedIntersectionClause(t *testing.T) {
	fullPolicy := NewPolicyNode()
	fullPolicy.SubjectResourceMappedIntersectionClause(
		SubjectResourceMappedIntersection("name", "rname"),
	)

	cases := []struct {
		name   string
		policy *PolicyNode
		input  *PolicyNode
		error  error
	}{
		{
			name:   "Error - Nil PolicyNode",
			policy: NewPolicyNode(),
			input:  nil,
			error:  fmt.Errorf("Unable to set subjectResourceIntersectionClause to a nil value."),
		},
		{
			name:   "Error - Empty PolicyNode",
			policy: NewPolicyNode(),
			input:  &PolicyNode{},
			error:  fmt.Errorf("Unable to set subjectResourceIntersectionClause to a nil value."),
		},
		{
			name:   "Error - PolicyNode already set",
			policy: fullPolicy,
			input:  SubjectResourceMappedIntersection("name", "rname"),
			error:  fmt.Errorf("PolicyNode has already been set with a node."),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.policy.SubjectResourceMappedIntersectionClause(c.input)
			assert.Equal(t, c.error, err)
		})
	}
}

func Test_PolicyNode_SubjectResourceForClause(t *testing.T) {
	innerPolicy := NewPolicyNode()
	innerPolicy.SubjectResourceMappedIntersectionClause(
		SubjectResourceMappedIntersection("name", "rname"),
	)
	fullPolicy := NewPolicyNode()
	fullPolicy.SubjectResourceForClause(
		SubjectResourceFor("name", ANY, innerPolicy),
	)

	cases := []struct {
		name   string
		policy *PolicyNode
		input  *PolicyNode
		error  error
	}{
		{
			name:   "Error - Nil PolicyNode",
			policy: NewPolicyNode(),
			input:  nil,
			error:  fmt.Errorf("unable to set subjectResourceForClause to a nil value"),
		},
		{
			name:   "Error - Empty PolicyNode",
			policy: NewPolicyNode(),
			input:  &PolicyNode{},
			error:  fmt.Errorf("unable to set subjectResourceForClause to a nil value"),
		},
		{
			name:   "Error - PolicyNode already set",
			policy: fullPolicy,
			input:  SubjectResourceFor("name", ANY, innerPolicy),
			error:  fmt.Errorf("policyNode has already been set with a node"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.policy.SubjectResourceForClause(c.input)
			assert.Equal(t, c.error, err)
		})
	}
}

func Test_PolicyNode_SubjectResourceSubsetClause(t *testing.T) {
	p := NewPolicyNode()
	err := p.SubjectResourceSubsetClause(nil)
	assert.NotNil(t, err)

	err = p.SubjectResourceSubsetClause(&PolicyNode{})
	assert.NotNil(t, err)

	err = p.SubjectResourceSubsetClause(SubjectResourceSubset("test"))
	assert.Nil(t, err)

	err = p.SubjectResourceSubsetClause(SubjectResourceSubset("test"))
	assert.NotNil(t, err)
}

func Test_PolicyNode_Operator(t *testing.T) {
	p := NewPolicyNode()
	err := p.Operator(nil)
	assert.NotNil(t, err)

	err = p.Operator(&PolicyNode{})
	assert.NotNil(t, err)

	p.SubjectValueIntersection(SubjectValueIntersection("test", attribute.String("value")))

	err = p.Operator(And())
	assert.NotNil(t, err)
}

func Test_PolicyNode_isValueSet(t *testing.T) {
	innerPolicy := NewPolicyNode()
	innerPolicy.SubjectResourceMappedIntersectionClause(
		SubjectResourceMappedIntersection("name", "rname"),
	)

	cases := []struct {
		name   string
		input  func() *PolicyNode
		output bool
	}{
		{
			name:  "False if none are set",
			input: func() *PolicyNode {return &PolicyNode{}},
			output: false,
		},
		{
			name:  "True if Operator is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.Operator(And())
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectResourceIntersectionClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.SubjectResourceIntersectionClause(SubjectResourceIntersection("abc"))
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectResourceSubsetClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.SubjectResourceSubsetClause(SubjectResourceSubset("abc"))
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectValueIntersectionClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.SubjectValueIntersection(SubjectValueIntersection("abc", attribute.String("value")))
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectMissingValueClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.MissingSubjectValue(SubjectMissingValue("ABC"))
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectResourceMappedIntersectionClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.SubjectResourceMappedIntersectionClause(SubjectResourceMappedIntersection("abc", "def"))
				return p
			},
			output: true,
		},
		{
			name:  "True if SubjectResourceForClause is set",
			input: func() *PolicyNode {
				p := &PolicyNode{}
				p.SubjectResourceForClause(SubjectResourceFor("name", ANY, innerPolicy))
				return p
			},
			output: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := c.input()
			assert.Equal(t, p.isValueSet(), c.output)
		})
	}
}

func Test_SubjectValueIntersection(t *testing.T) {
	se := SubjectValueIntersection("attr_name", attribute.String("attr_value"))
	assert.Equal(t, "attr_name", se.GetSubjectValueIntersection().AttributeName)
	assert.Equal(t, attribute.String("attr_value"), se.GetSubjectValueIntersection().AttributeValue)
}

func Test_SubjectResourceIntersectionClause(t *testing.T) {
	se := SubjectResourceIntersection("attr_name")
	assert.Equal(t, "attr_name", se.GetSubjectResourceIntersectionClause().AttributeName)
}

func Test_SubjectResourceMappedIntersectionClause(t *testing.T) {
	cases := []struct {
		name     string
		subjAttr string
		resAttr  string
	}{
		{
			name: "Subject attribute is set",
			subjAttr: "subjAttr",
			resAttr: "",
		},
		{
			name: "Resource attribute is set",
			subjAttr: "",
			resAttr: "resAttr",
		},
		{
			name: "Both attributes are set",
			subjAttr: "subjAttr",
			resAttr: "resAttr",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := SubjectResourceMappedIntersection(c.subjAttr, c.resAttr)
			assert.Equal(t, c.subjAttr, actual.GetSubjectResourceMappedIntersectionClause().SubjectAttributeName)
			assert.Equal(t, c.resAttr, actual.GetSubjectResourceMappedIntersectionClause().ResourceAttributeName)
		})
	}
}

func Test_SubjectResourceForClause(t *testing.T) {
	innerPolicy := NewPolicyNode()
	innerPolicy.SubjectResourceMappedIntersectionClause(
		SubjectResourceMappedIntersection("name", "rname"),
	)
	cases := []struct {
		name     string
		subjAttr string
		operator ForOperator
		policy   *PolicyNode
	}{
		{
			name: "Subject attribute is set",
			subjAttr: "subjAttr",
			operator: ANY,
			policy: nil,
		},
		{
			name: "Operator is set",
			subjAttr: "",
			operator: ANY,
			policy: nil,
		},
		{
			name: "Policy is set",
			subjAttr: "",
			operator: ANY,
			policy: SubjectResourceFor("name", ANY, innerPolicy),
		},
		{
			name: "All are set",
			subjAttr: "subjAttr",
			operator: ANY,
			policy: SubjectResourceFor("name", ANY, innerPolicy),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := SubjectResourceFor(c.subjAttr, c.operator, c.policy)
			assert.Equal(t, c.subjAttr, actual.GetSubjectResourceForClause().AttributeName)
			assert.Equal(t, c.operator, actual.GetSubjectResourceForClause().Operator)
			assert.Equal(t, c.policy, actual.GetSubjectResourceForClause().Policy)
		})
	}
}

func Test_SubjectMissingValueClause(t *testing.T) {
	se := SubjectMissingValue("attr_name")
	assert.Equal(t, "attr_name", se.GetSubjectMissingValueClause().AttributeName)
}

func Test_AndOperator(t *testing.T) {
	se := And(SubjectResourceIntersection("attr_name"))
	op := se.GetOperator()
	assert.Equal(t, "attr_name", op.Children[0].GetSubjectResourceIntersectionClause().AttributeName)
	assert.Equal(t, AND, op.Operator)
}

func Test_OrOperator(t *testing.T) {
	se := Or(SubjectResourceIntersection("attr_name"))
	op := se.GetOperator()
	assert.Equal(t, "attr_name", op.Children[0].GetSubjectResourceIntersectionClause().AttributeName)
	assert.Equal(t, OR, op.Operator)
}

func Test_NotOperator(t *testing.T) {
	se := Not(SubjectResourceIntersection("attr_name"))
	op := se.GetOperator()
	assert.Equal(t, "attr_name", op.Children[0].GetSubjectResourceIntersectionClause().AttributeName)
	assert.Equal(t, NOT, op.Operator)
}
