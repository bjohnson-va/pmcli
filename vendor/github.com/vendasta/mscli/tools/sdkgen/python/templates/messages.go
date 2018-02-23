package templates

//MessageTemplate is the template for generating the message layer
var MessageTemplate = `"""
# Code generated by sdkgen
# DO NOT EDIT!.

Message layer for data objects.
""" {{ $template := . }}
import vobject
import enum
import datetime
from .._generated import {{ $template.PyProtoFileName }}

{{ range $dep := $template.NonGoogleDependenciesFileNames $template.Dependencies -}}
	from {{ $dep }} import *
{{ end }}

def extract_datetime(proto_date_attr):
    """ Extract a datetime object from a proto property """
    if not proto_date_attr or proto_date_attr.ToDatetime() == datetime.datetime.utcfromtimestamp(0):
        return None
    return proto_date_attr.ToDatetime()


def extract_repeated(proto_repeated_attr):
    """ Extract a list from repeated attributes """
    return [a for a in proto_repeated_attr]


def extract_falsy_as_none(proto_descriptor):
    """
    Return any falsy value as None.
    We do this because python protobuf's metaclasses and magic methods obscure our ability to interact with these
    objects in a sane way.
    """
    if not proto_descriptor and proto_descriptor is not None:
        return None
    return proto_descriptor

{{ range $enum := $template.Enums }}
class {{ $enum.Name }}(enum.Enum):
	""" {{ $enum.Name }} """
	{{- range $enumVal := $enum.Value }}
	{{ $enumVal.Name }} = {{ $enumVal.Number -}}
	{{- end }}

{{ end }}

{{- range $message := $template.Messages }}
class {{$message.Message.Name}}(vobject.VObject):
	""" {{ $message.Message.Name }} """
	OWNER = "{{ $template.PyProtoFileName }}"
	CLASS_VERSION = "1.0.0"

	{{ range $field := $message.Message.Field -}}
		{{- $field.Name }} = vobject.{{ $template.FieldToVobjectProperty $field }}
	{{ end }}
	@classmethod
	def from_proto(cls, message):
		""" Convert from proto """
		return {{ $message.Message.Name }}(
		{{- range $i, $field := $message.Message.Field -}}
			{{- if ne $i 0 -}}, {{ end }}
			{{ $field.Name -}}={{- $template.FieldToDomainValue $field -}}
		{{- end }}
		)

	def to_proto(self):
		""" Convert to proto """
		proto = {{ $template.PyProtoFileName }}.{{ $message.FullyQualifiedName }}(
		{{- range $i, $field := ($template.SetInsideConstructorFields $message.Message) -}}
			{{- if ne $i 0 -}}, {{ end }}
			{{ $field.Name -}}={{- $template.FieldToProtoValue $field }} if self.{{ $field.Name }} is not None else None
		{{- end }}
		)
		{{ range $i, $field := ($template.DateTimeFields $message.Message) }}
		if self.{{ $field.Name }}:
			proto.{{ $field.Name }}.FromDatetime(self.{{ $field.Name }})
		{{- end }}
		{{ range $i, $field := ($template.WrapperFields $message.Message) }}
		if self.{{ $field.Name }} is not None:
			proto.{{ $field.Name }}.value = self.{{ $field.Name }}
		{{- end }}
		return proto

{{ end -}}
`