package python

import (
	"fmt"

	"strings"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
	"github.com/vendasta/mscli/tools/sdkgen/python/templates"
	"github.com/vendasta/mscli/tools/sdkgen/util"
)

func Generate(p *internal.Proto, messageEnumDescriptors []*internal.MessageEnumDescriptors) *plugin.CodeGeneratorResponse {
	pmed := messageEnumDescriptorToPython(messageEnumDescriptors)

	output := &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			generateTransportFile(p),
			generateClientFile(p),
			generateInitFile(),
		},
	}

	output.File = append(output.File, generateMessageFiles(pmed)...)
	output.File = append(output.File, generateProtoInitFiles(p)...)

	return output
}

func messageEnumDescriptorToPython(med []*internal.MessageEnumDescriptors) []*PythonMessageEnumDescriptors {
	pmed := make([]*PythonMessageEnumDescriptors, len(med), len(med))
	for i, descriptor := range med {
		pmed[i] = &PythonMessageEnumDescriptors{descriptor}
	}
	return pmed
}

func generateMessageFiles(descriptors []*PythonMessageEnumDescriptors) []*plugin.CodeGeneratorResponse_File {
	var files []*plugin.CodeGeneratorResponse_File
	for _, descriptor := range descriptors {
		file := generateMessageFile(descriptor)
		files = append(files, file)
	}
	return files
}

func generateMessageFile(descriptor *PythonMessageEnumDescriptors) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"message",
		templates.MessageTemplate,
		fmt.Sprintf("/_internal/%s.py", descriptor.PyFileName()),
		descriptor,
	)
}

func convertFilePathToVersion(filePath string) string {
	directories := strings.Split(filePath, "/")
	if len(directories) >= 2 {
		return directories[len(directories)-2]
	}
	return "v1"
}

func generateProtoInitFiles(p *internal.Proto) []*plugin.CodeGeneratorResponse_File {
	var version string
	if len(p.Services) > 0 {
		version = convertFilePathToVersion(p.Services[0].FileName)
	} else if len(p.NonServiceProtos) > 0 {
		version = convertFilePathToVersion(p.NonServiceProtos[0].FileName)
	} else {
		version = "v1" // just default and hope for the best
	}

	var files []*plugin.CodeGeneratorResponse_File
	// Create the Init file in _generated
	file := internal.TemplateFileGenerator(
		"init",
		templates.GeneratedInitTemplate,
		fmt.Sprintf("/_generated/__init__.py"),
		struct {
			Proto      *internal.Proto
			ModuleName string
			Version    string
		}{
			Proto:      p,
			ModuleName: util.ToSnakeCase(p.MicroserviceName),
			Version:    version,
		},
	)
	files = append(files, file)
	file = internal.TemplateFileGenerator(
		"init",
		templates.InitTemplate,
		fmt.Sprintf("/_generated/grpc/%s/__init__.py", util.ToSnakeCase(p.MicroserviceName)),
		nil,
	)
	files = append(files, file)
	file = internal.TemplateFileGenerator(
		"init",
		templates.InitTemplate,
		fmt.Sprintf("/_generated/proto/%s/__init__.py", util.ToSnakeCase(p.MicroserviceName)),
		nil,
	)
	files = append(files, file)
	file = internal.TemplateFileGenerator(
		"init",
		templates.InitTemplate,
		fmt.Sprintf("/_generated/grpc/%s/%s/__init__.py", util.ToSnakeCase(p.MicroserviceName), version),
		nil,
	)
	files = append(files, file)
	file = internal.TemplateFileGenerator(
		"init",
		templates.InitTemplate,
		fmt.Sprintf("/_generated/proto/%s/%s/__init__.py", util.ToSnakeCase(p.MicroserviceName), version),
		nil,
	)
	files = append(files, file)

	return files
}

func generateTransportFile(p *internal.Proto) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"transport",
		templates.TransportTemplate,
		"/_internal/transport.py",
		p,
	)
}

func generateClientFile(p *internal.Proto) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"client",
		templates.APIClientTemplate,
		"/_internal/client.py",
		p,
	)
}

func generateInitFile() *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"init",
		templates.InitTemplate,
		"/_internal/__init__.py",
		nil,
	)
}
