package iam

// Resource is an object that has specific access controls that are managed and enforced by IAM.
type Resource struct {
	// AppID from your registered resource owner.
	AppID string

	// ResourceID is an identifier for the resource that is unique within the resource owner's set of resources.
	ResourceID string

	// ResourceName is the human readable name of the resource (eg "Account Group")
	ResourceName string

	// ResourceOwnerServiceURL is the base url where a ResourceOwner grpc service is being served. IAM will use this service to ask about specific resources
	ResourceOwnerServiceURL string

	// RequiredResourceParams is the set of identifiers that must be provided to ask the resource provider about a resource from
	// IAM's point of view. These parameters will need to be supplied by any client asking about this resource.
	// eg: ["pid", "account_group_id"] means that each time IAM asks for one of these resources, it will require the client to specify
	// which "pid" and "account_group_id" they are asking about. IAM will proxy these parameters to the ResourceOwnerServiceURL.
	RequiredResourceParams []string

	//resource_owner_audience is a google id token audience. This is used by the resource owner to validate requests.
	ResourceOwnerAudience string
}

// NewResource returns a new IAM resource.
func NewResource(AppID, ResourceID, ResourceName, ResourceOwnerServiceURL, ResourceOwnerAudience string, RequiredResourceParams []string) *Resource {
	return &Resource{
		AppID:                   AppID,
		ResourceID:              ResourceID,
		ResourceName:            ResourceName,
		ResourceOwnerServiceURL: ResourceOwnerServiceURL,
		ResourceOwnerAudience:   ResourceOwnerAudience,
		RequiredResourceParams:  RequiredResourceParams,
	}
}
