package iam

// ResourceOwner is an application that owns one or many distinct resources.
type ResourceOwner struct {
	// AppID is a unique identifier for the registrant. It must be unique across all IAM resource owners. (eg "VBC")
	AppID string

	// AppName is the human readable name of the application registering as a Resource Owner (eg "Vendasta Business Center")
	AppName string
}

// NewResourceOwner returns a new ResourceOwner
func NewResourceOwner(AppID string, AppName string) *ResourceOwner {
	return &ResourceOwner{AppID: AppID, AppName: AppName}
}
