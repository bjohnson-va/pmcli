package templates

// GeneratedTemplate generated python file
type GeneratedTemplate struct {
	FileNames []string
}

// InitTemplate imports at the top of the file
const InitTemplate = `""" Generated """

from . import {{ range $idx, $file := .FileNames }}{{ if gt $idx 0 }}, {{ end }}{{ $file }}{{ end }}
`
