package iam

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

// SubjectIdentifier controls how a client identifies a subject. Can either specify Email or SubjectID both not both.
type SubjectIdentifier struct {
	Email     string
	SubjectID string
}

type ResourceEntityIdentifier map[string][]string

// AccessResource encapsulates the information necessary to ask IAM if a specific user has permissions to access a
// specified resource.
type AccessResource struct {
	// The context under which the subject is being identified
	Context *subjectcontext.Context

	// Indicates how the subject should be identified.
	SubjectIdentifier *SubjectIdentifier

	// SessionID is the IAM session ID of the subject - alternate way to identify who is making the request
	SessionID string

	// OwnerId is the app_id of the ResourceOwner that owns the resource_id
	OwnerID string

	// ResourceId is the resource type that the subject wants to access
	ResourceID string

	// ResourceEntityIdentifier contains the information necessary to identify a specific resource of the type specified by resource_id
	ResourceEntityIdentifier ResourceEntityIdentifier

	// ResourceEntityIdentifiers contains the information necessary to identify multiple resources of the type specified by resource_id
	ResourceEntityIdentifiers []ResourceEntityIdentifier

	// AccessScope represents the reasons that the user wants to access the resource.
	AccessScope []AccessScope

	// ResourceAttributes contains the resource attributes to use during policy evaluation.
	// Only works for registered resources without a resource_owner_service_url specified.
	ResourceAttributes *iam_v1.StructAttribute
}

// ToPB converts an AccessResource to an iam_v1.AccessResourceRequest
func (ar *AccessResource) ToPB() *iam_v1.AccessResourceRequest {
	accessScopesProtos := []iam_v1.AccessScope{}
	for _, accessScope := range ar.AccessScope {
		accessScopesProtos = append(accessScopesProtos, iam_v1.AccessScope(accessScope))
	}

	r := &iam_v1.AccessResourceRequest{
		Context:            ar.Context.ToPB(),
		OwnerId:            ar.OwnerID,
		ResourceId:         ar.ResourceID,
		AccessScope:        accessScopesProtos,
		ResourceAttributes: ar.ResourceAttributes,
	}

	if ar.ResourceEntityIdentifier != nil {
		resouceIdentifierListProto := map[string]*iam_v1.ValueList{}
		for k, vals := range ar.ResourceEntityIdentifier {
			resouceIdentifierListProto[k] = &iam_v1.ValueList{Values: vals}
		}
		r.ResourceEntityIdentifier = resouceIdentifierListProto
	} else {
		resourceEntityIdentifiers := make([]*iam_v1.ResourceIdentifier, len(ar.ResourceEntityIdentifiers))
		for n, resourceIdentifier := range ar.ResourceEntityIdentifiers {
			vl := map[string]*iam_v1.ValueList{}
			for k, vals := range resourceIdentifier {
				vl[k] = &iam_v1.ValueList{Values: vals}
			}
			resourceEntityIdentifiers[n] = &iam_v1.ResourceIdentifier{
				Identifiers: vl,
			}
		}
		r.ResourceEntityIdentifiers = resourceEntityIdentifiers
	}

	if ar.SessionID != "" {
		r.Session = ar.SessionID
	} else if ar.SubjectIdentifier != nil && ar.SubjectIdentifier.Email != "" {
		r.SubjectIdentifier = &iam_v1.AccessResourceRequest_Email{
			Email: ar.SubjectIdentifier.Email,
		}
	} else if ar.SubjectIdentifier != nil && ar.SubjectIdentifier.SubjectID != "" {
		r.SubjectIdentifier = &iam_v1.AccessResourceRequest_SubjectId{
			SubjectId: ar.SubjectIdentifier.SubjectID,
		}
	}

	return r
}
