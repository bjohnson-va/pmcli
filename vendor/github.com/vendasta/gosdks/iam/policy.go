package iam

import (
	"fmt"
	"github.com/vendasta/gosdks/pb/iam/v1"
)

// NewPolicyNode returns a new Policy Node. PolicyNode leverages a fluent API for building the tree of policies.
func NewPolicyNode() *PolicyNode {
	return &PolicyNode{}
}

// NewPolicy creates a new Policy
func NewPolicy(AppID string, ResourceID string, PolicyID string, PolicyName string, PolicyNode *PolicyNode, Operations ...AccessScope) *Policy {
	return &Policy{
		AppID:      AppID,
		ResourceID: ResourceID,
		PolicyID:   PolicyID,
		PolicyName: PolicyName,
		Policy:     PolicyNode,
		Operations: Operations,
	}
}

// AccessScope is a specific scope that applications ask for permission to do against a specific resource.
//
// Example is a Partner user A wants AccessScope READ permission to partner ABC
type AccessScope int

const (
	// READ level scope
	READ AccessScope = iota
	// WRITE level scope
	WRITE
	// DELETE level scope
	DELETE
)

var AccessScopeName = map[AccessScope]string{
	READ:   "READ",
	WRITE:  "WRITE",
	DELETE: "DELETE",
}

// BooleanOperator allows policies to be ANDd or ORd together or NOT a specific policy node.
type BooleanOperator int

const (
	// AND boolean operator
	AND BooleanOperator = iota

	// OR boolean operator
	OR

	// NOT boolean operator
	NOT
)

//ForOperator defines how the SubjectResourceFor clauses will evaluate the list of StructAttributes
type ForOperator int

const (
	//ANY struct inside of the list must match the PolicyNode to be true
	ANY ForOperator = iota

	//ALL structs inside of the list must match the PolicyNode to be true
	ALL
)

// A PolicyNode can either be a Boolean operator over one or more child PolicyNodes, or it can be an SubjectResourceIntersectionClause.
// A policy tree might look like this where C1, C2, C3 are SubjectResourceIntersectionClauses:
//          AND
//         /   \
//        C1   OR
//            /  \
//           C2  C3
// This means that the expression would be evaluated as C1 && (C2 || C3)
// Arbitrary nesting is valid.
type PolicyNode struct {
	// A policy node must either be a boolean conjunction of child clauses or an actual clause
	operator                                *Operator
	subjectResourceIntersectionClause       *SubjectResourceIntersectionClause
	subjectResourceSubsetClause             *SubjectResourceSubsetClause
	subjectValueIntersection                *SubjectValueIntersectionClause
	subjectMissingValueClause               *SubjectMissingValueClause
	subjectResourceMappedIntersectionClause *SubjectResourceMappedIntersectionClause
	subjectResourceForClause                *SubjectResourceForClause
}

func (pn *PolicyNode) isValueSet() bool {
	return pn.operator != nil || pn.subjectResourceIntersectionClause != nil || pn.subjectResourceSubsetClause != nil ||
		pn.subjectValueIntersection != nil || pn.subjectMissingValueClause != nil ||
		pn.subjectResourceMappedIntersectionClause != nil || pn.subjectResourceForClause != nil
}

// SubjectValueIntersection applies a subject value intersection node.
func (pn *PolicyNode) SubjectValueIntersection(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.subjectValueIntersection == nil {
		return fmt.Errorf("Unable to set subjectValueIntersection to a nil value.")
	}
	pn.subjectValueIntersection = p.subjectValueIntersection
	return nil
}

// SubjectResourceIntersectionClause applies a subject resource intersection node.
func (pn *PolicyNode) SubjectResourceIntersectionClause(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.subjectResourceIntersectionClause == nil {
		return fmt.Errorf("Unable to set subjectResourceIntersectionClause to a nil value.")
	}
	pn.subjectResourceIntersectionClause = p.subjectResourceIntersectionClause
	return nil
}

// SubjectResourceMappedIntersectionClause applies a subject resource mapped intersection node.
func (pn *PolicyNode) SubjectResourceMappedIntersectionClause(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.subjectResourceMappedIntersectionClause == nil {
		return fmt.Errorf("Unable to set subjectResourceIntersectionClause to a nil value.")
	}
	pn.subjectResourceMappedIntersectionClause = p.subjectResourceMappedIntersectionClause
	return nil
}

// SubjectResourceForClause applies a subject resource for node.
func (pn *PolicyNode) SubjectResourceForClause(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("policyNode has already been set with a node")
	}
	if p == nil || p.subjectResourceForClause == nil {
		return fmt.Errorf("unable to set subjectResourceForClause to a nil value")
	}
	pn.subjectResourceForClause = p.subjectResourceForClause
	return nil
}

// SubjectResourceSubsetClause applies a subject resource subset node.
func (pn *PolicyNode) SubjectResourceSubsetClause(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.subjectResourceSubsetClause == nil {
		return fmt.Errorf("Unable to set subjectResourceSubsetClause to a nil value.")
	}
	pn.subjectResourceSubsetClause = p.subjectResourceSubsetClause
	return nil
}

// MissingSubjectValue applies a missing subject value node.
func (pn *PolicyNode) MissingSubjectValue(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.subjectMissingValueClause == nil {
		return fmt.Errorf("Unable to set operator to a nil value.")
	}
	pn.subjectMissingValueClause = p.subjectMissingValueClause
	return nil
}

// Operator applies an operator level node.
func (pn *PolicyNode) Operator(p *PolicyNode) error {
	if pn.isValueSet() {
		return fmt.Errorf("PolicyNode has already been set with a node.")
	}
	if p == nil || p.operator == nil {
		return fmt.Errorf("Unable to set operator to a nil value.")
	}
	pn.operator = p.operator
	return nil
}

// GetSubjectValueIntersection returns the subject value intersection policy node.
func (pn *PolicyNode) GetSubjectValueIntersection() *SubjectValueIntersectionClause {
	return pn.subjectValueIntersection
}

// GetSubjectResourceIntersectionClause returns the subject resource intersection policy node.
// Deprecated; use SubjectResourceMappedIntersectionClause instead.
func (pn *PolicyNode) GetSubjectResourceIntersectionClause() *SubjectResourceIntersectionClause {
	return pn.subjectResourceIntersectionClause
}

// GetSubjectResourceSubsetClause returns the resource subset clause policy node.
func (pn *PolicyNode) GetSubjectResourceSubsetClause() *SubjectResourceSubsetClause {
	return pn.subjectResourceSubsetClause
}

// GetSubjectMissingValueClause returns the missing value clause policy node.
func (pn *PolicyNode) GetSubjectMissingValueClause() *SubjectMissingValueClause {
	return pn.subjectMissingValueClause
}

// GetSubjectResourceMappedIntersectionClause returns the subject resource mapped intersection clause policy node.
func (pn *PolicyNode) GetSubjectResourceMappedIntersectionClause() *SubjectResourceMappedIntersectionClause {
	return pn.subjectResourceMappedIntersectionClause
}

// GetSubjectResourceForClause returns the subject resource for clause policy node.
func (pn *PolicyNode) GetSubjectResourceForClause() *SubjectResourceForClause {
	return pn.subjectResourceForClause
}

// GetOperator returns the boolean operator
func (pn *PolicyNode) GetOperator() *Operator {
	return pn.operator
}

// SubjectValueIntersection creates a new subject value intersection policy node.
func SubjectValueIntersection(attributeName string, attributeValue *iam_v1.Attribute) *PolicyNode {
	return &PolicyNode{
		subjectValueIntersection: &SubjectValueIntersectionClause{
			AttributeName:  attributeName,
			AttributeValue: attributeValue,
		},
	}
}

// SubjectResourceIntersection creates a new subject resource intersection policy node.
func SubjectResourceIntersection(attributeName string) *PolicyNode {
	return &PolicyNode{
		subjectResourceIntersectionClause: &SubjectResourceIntersectionClause{
			AttributeName: attributeName,
		},
	}
}

// SubjectResourceMappedIntersection creates a new subject resource mapped intersection policy node.
func SubjectResourceMappedIntersection(subjectAttributeName, resourceAttributeName string) *PolicyNode {
	return &PolicyNode{
		subjectResourceMappedIntersectionClause: &SubjectResourceMappedIntersectionClause{
			SubjectAttributeName:  subjectAttributeName,
			ResourceAttributeName: resourceAttributeName,
		},
	}
}

// SubjectResourceSubset creates a new subject resource subset policy node.
func SubjectResourceSubset(attributeName string) *PolicyNode {
	return &PolicyNode{
		subjectResourceSubsetClause: &SubjectResourceSubsetClause{
			AttributeName: attributeName,
		},
	}
}

// SubjectMissingValue creates a new missing value policy node.
func SubjectMissingValue(attributeName string) *PolicyNode {
	return &PolicyNode{
		subjectMissingValueClause: &SubjectMissingValueClause{
			AttributeName: attributeName,
		},
	}
}

// SubjectResourceFor creates a new subject resource for policy node.
func SubjectResourceFor(attributeName string, operator ForOperator, policy *PolicyNode) *PolicyNode {
	return &PolicyNode{
		subjectResourceForClause: &SubjectResourceForClause{
			AttributeName: attributeName,
			Operator: operator,
			Policy: policy,
		},
	}
}

// And builds a new PolicyNode that forces that all of the child nodes must evaluate to a truthy value.
func And(policyNodes ...*PolicyNode) *PolicyNode {
	return &PolicyNode{
		operator: &Operator{
			Operator: AND,
			Children: policyNodes,
		},
	}
}

// Or builds a new PolicyNode that forces that at least one of the child nodes must evaluate to a truthy value.
func Or(policyNodes ...*PolicyNode) *PolicyNode {
	return &PolicyNode{
		operator: &Operator{
			Operator: OR,
			Children: policyNodes,
		},
	}
}

// Not build a new PolicyNode that takes ONLY an SubjectResourceIntersectionClause policy node.
func Not(policyNode *PolicyNode) *PolicyNode {
	return &PolicyNode{
		operator: &Operator{
			Operator: NOT,
			Children: []*PolicyNode{policyNode},
		},
	}
}

// SubjectValueIntersectionClause describes how an attribute's value on the subject must contain the arbitrary value specified by
// attribute_value.
// eg: with attribute_name = "subject_type" and attribute_value = "partner", then the subject MUST have an attribute named
// "subject_type" that contains the value "partner". Note that this is different from SubjectResourceIntersectionClause because we don't
// care about this attributes presence on the resource or its value, it is purely for enforcing constraints on the subject itself.
type SubjectValueIntersectionClause struct {
	AttributeName  string
	AttributeValue *iam_v1.Attribute
}

// SubjectMissingValueClause is true if the subject attribute keyed by attribute_name is either missing or has no values.
// eg: with attribute_name = "market_id", then a subject with "market_id" = [] is TRUE, a subject with "market_id" = ["something"] is FALSE,
// and if the subject simply doesn't have an attribute called "market_id" then this clause is TRUE
type SubjectMissingValueClause struct {
	AttributeName string
}

// SubjectResourceIntersectionClause describes how an attribute's value on the resource must be related to the attributes on the subject.
// eg: with and attribute_name = "account_group_id" then the subject MUST have an attribute named "account_group_id"
// that contains the value of the resource's "account_group_id" attribute for this clause to be considered TRUE.
type SubjectResourceIntersectionClause struct {
	AttributeName string
}

//SubjectResourceMappedIntersectionClause is true if the set of values keyed by resource_attribute_name on the resource
//has a non-empty intersection with the set of values keyed by subject_attribute_name on the subject.
//eg: with resource_attribute_name = "account_group_id" and subject_attribute_name "accessible_group_ids"
//then the subject's "accessible_group_ids" attribute must contain any individual value of the resource's "account_group_id"
//attribute to be considered TRUE.
type SubjectResourceMappedIntersectionClause struct {
	SubjectAttributeName  string
	ResourceAttributeName string
}

//SubjectResourceForClause will loop through a list of StructAttributes and return true if any or all match the PolicyNode defined.
//The Subject MUST have a ListAttribute that contains StructAttributes at the attribute_name key for this to be evaluated.
type SubjectResourceForClause struct {
	AttributeName string
	Operator      ForOperator
	Policy        *PolicyNode
}

// SubjectResourceSubsetClause is true if the set of values keyed by attribute_name on the resource is a subset of the set of values keyed by attribute_name on the subject.
type SubjectResourceSubsetClause struct {
	AttributeName string
}

// Operator consists of a boolean conjunction and one to two children nodes.
// NOT is the only conjunction that supports a single child node, whereas AND and OR both require exactly two child nodes.
type Operator struct {
	Operator BooleanOperator
	Children []*PolicyNode
}

// A Policy is a binary boolean expression tree of attribute-based clauses that determines whether a subject is authorized to access a resource in a specific context.
type Policy struct {
	// AppID is a unique identifier for the ResourceOwner. It must be unique across all IAM resource owners. (eg "VBC")
	AppID string

	// The unique identifier for the resource this policy will be applied to.
	ResourceID string

	// The identifier for this policy. This need only be unique within the scope of the resource.
	PolicyID string

	// PolicyName is the human readable name of the policy, often phrased as a question (eg: "Can user access account group?")
	PolicyName string

	// Operations that a user may take on a resource.
	// In particular, these will be the operations that this policy is associated with, so if this policy is associated
	// with "read", and a user attempts to "read" this resource type, this policy would be evaluated.
	Operations []AccessScope

	// policy is the boolean expression tree of boolean operators and attribute-based clauses that expresses the method by which the policy is evaluated.
	Policy *PolicyNode
}
