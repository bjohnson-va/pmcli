package templates

//InterfacesIndexTemplate is the template for the index.ts for the interface layer
var InterfacesIndexTemplate = `{{- $descriptors := . -}}
// *********************************
// Code generated by sdkgen
// DO NOT EDIT!.
//
// Interfaces Index.
// *********************************

{{- range $descriptor := $descriptors }}
export {
	{{- range $message := $descriptor.Messages }}
	{{ $message.InterfaceName }},
	{{- end }}
} from './{{ $descriptor.TsInterfaceFileName }}';
{{ end }}
`
