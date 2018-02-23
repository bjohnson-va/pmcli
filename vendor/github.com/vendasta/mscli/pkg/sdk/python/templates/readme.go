package templates

// ReadmeTemplate data for generating the readme
type ReadmeTemplate struct {
	Version string
}

// SDKReadmeTemplate is used to generate the Python SDK README.md
const SDKReadmeTemplate = `{{ .Version }}
- Initial generated sdk
`
