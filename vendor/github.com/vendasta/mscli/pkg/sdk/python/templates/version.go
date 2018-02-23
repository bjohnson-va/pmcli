package templates

// VersionTemplate Stores the version of the sdk
type VersionTemplate struct {
	Version string
}

// SDKVersionTemplate is used to generate the Python SDK's version file
const SDKVersionTemplate = `{{ .Version }}
`
