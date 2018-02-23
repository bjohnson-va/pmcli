package templates

//ChangelogModel data for generating the CHANGELOG.md
type ChangelogModel struct {
	Version string
}

//PythonSDKReadmeTemplate is used to generate the CHANGELOG.md
const ChangelogTemplate = `{{ .Version }}
- Initial generated sdk
`
