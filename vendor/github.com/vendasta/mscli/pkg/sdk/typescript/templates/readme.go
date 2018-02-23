package templates

//ReadmeModel data for generating the README.md
type ReadmeModel struct {
	Name string
}

//ReadmeTemplate is used to generate README.md
const ReadmeTemplate = `# {{ .Name }} SDK
`
