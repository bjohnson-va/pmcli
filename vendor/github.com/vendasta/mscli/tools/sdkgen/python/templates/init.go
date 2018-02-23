package templates

//InitTemplate is the template for generating the main __init__.py file
var InitTemplate = `""" Generated """
`

// PythonGeneratedInitTemplate python file template
const GeneratedInitTemplate = `{{- $template := . -}}
""" Generated Protos. Made by SDKGen """
import sys
import os
sys.path.append(os.path.join(os.path.dirname(__file__), "grpc"))

{{ range $idx, $service := .Proto.Services }}
{{ $service.PyFileName }} = None
{{ end }}
{{ range $idx, $service := .Proto.NonServiceProtos }}
{{ $service.PyFileName }} = None
{{ end }}
try:
    {{ range $idx, $service := .Proto.Services }}{{ if eq $idx 0 }}from {{ $template.ModuleName }}.{{ $template.Version }} import {{ end }}{{ if gt $idx 0 }}, {{ end }}{{ $service.PyFileName }}{{ end }}
	{{ range $idx, $service := .Proto.NonServiceProtos }}{{ if eq $idx 0 }}from {{ $template.ModuleName }}.{{ $template.Version }} import {{ end }}{{ if gt $idx 0 }}, {{ end }}{{ $service.PyFileName }}{{ end }}
except ImportError:
	sys.path.remove(os.path.join(os.path.dirname(__file__), "grpc"))
	sys.path.append(os.path.join(os.path.dirname(__file__), "proto"))
	for key in sys.modules.keys():
		if key.startswith('{{ $template.ModuleName }}.'):
			sys.modules.pop(key)
		if key == '{{ $template.ModuleName }}':
			sys.modules.pop(key)
    {{ range $idx, $service := .Proto.Services }}{{ if eq $idx 0 }}from {{ $template.ModuleName }}.{{ $template.Version }} import {{ end }}{{ if gt $idx 0 }}, {{ end }}{{ $service.PyFileName }}{{ end }}
	{{ range $idx, $service := .Proto.NonServiceProtos }}{{ if eq $idx 0 }}from {{ $template.ModuleName }}.{{ $template.Version }} import {{ end }}{{ if gt $idx 0 }}, {{ end }}{{ $service.PyFileName }}{{ end }}
`
