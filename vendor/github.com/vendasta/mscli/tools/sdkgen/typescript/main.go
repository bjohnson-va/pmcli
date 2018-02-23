package typescript

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
	"fmt"
	"github.com/vendasta/mscli/tools/sdkgen/typescript/templates"
)

func Generate(p *internal.Proto, messageEnumDescriptors []*internal.MessageEnumDescriptors) (*plugin.CodeGeneratorResponse) {
	descriptors := messageEnumDescriptorToTypescript(messageEnumDescriptors)

	output := &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			generateInterfaceIndexFile(descriptors),
			generateObjectIndexFile(descriptors),
		},
	}

	apiServiceFiles := generateAPIServiceFiles(p)
	output.File = append(output.File, apiServiceFiles...)

	enumFiles := generateEnumFiles(descriptors)
	includeEnums := len(enumFiles) > 0
	if includeEnums {
		output.File = append(output.File, enumFiles...)
		output.File = append(output.File, generateEnumIndexFile(descriptors))
	}
	output.File = append(output.File, generateIndexFile(p, includeEnums))
	output.File = append(output.File, generateObjectFiles(descriptors, includeEnums)...)
	output.File = append(output.File, generateInterfaceFiles(descriptors, includeEnums)...)
	output.File = append(output.File, generateModuleFile(p))

	return output
}

func messageEnumDescriptorToTypescript(med []*internal.MessageEnumDescriptors) []*TypescriptMessageEnumDescriptors {
	tmed := make([]*TypescriptMessageEnumDescriptors, len(med), len(med))
	for i, descriptor := range med {
		tmed[i] = &TypescriptMessageEnumDescriptors{descriptor}
	}
	return tmed
}

// ********************************************
// Interfaces
// ********************************************
func generateInterfaceFiles(descriptors []*TypescriptMessageEnumDescriptors, includeEnums bool) []*plugin.CodeGeneratorResponse_File {
	var files []*plugin.CodeGeneratorResponse_File
	for _, descriptor := range descriptors {
		files = append(files, generateInterfaceFile(descriptor, includeEnums))
	}
	return files
}

func generateInterfaceFile(descriptor *TypescriptMessageEnumDescriptors, includeEnums bool) *plugin.CodeGeneratorResponse_File {
	//InterfacesModel data for generating the interface files
	return internal.TemplateFileGenerator(
		"interface",
		templates.InterfacesTemplate,
		fmt.Sprintf("/_internal/interfaces/%s.ts", descriptor.TsInterfaceFileName()),
		struct {
			Descriptor *TypescriptMessageEnumDescriptors
			IncludeEnums bool
		}{
			Descriptor: descriptor,
			IncludeEnums: includeEnums,
		},
	)
}

func generateInterfaceIndexFile(descriptors []*TypescriptMessageEnumDescriptors) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"interfaceIndex",
		templates.InterfacesIndexTemplate,
		"/_internal/interfaces/index.ts",
		descriptors,
	)
}

// ********************************************
// Enums
// ********************************************
func generateEnumFiles(descriptors []*TypescriptMessageEnumDescriptors) []*plugin.CodeGeneratorResponse_File {
	var files []*plugin.CodeGeneratorResponse_File
	for _, descriptor := range descriptors {
		file := generateEnumFile(descriptor)
		if file != nil {
			files = append(files, file)
		}
	}
	return files
}

func generateEnumFile(descriptor *TypescriptMessageEnumDescriptors) *plugin.CodeGeneratorResponse_File {
	if len(descriptor.Enums) == 0 {
		return nil
	}
	return internal.TemplateFileGenerator(
		"enum",
		templates.EnumsTemplate,
		fmt.Sprintf("/_internal/enums/%s.ts", descriptor.TsEnumFileName()),
		descriptor,
	)
}

func generateEnumIndexFile(descriptors []*TypescriptMessageEnumDescriptors) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"enumIndex",
		templates.EnumsIndexTemplate,
		"/_internal/enums/index.ts",
		descriptors,
	)
}

// ********************************************
// Objects
// ********************************************
func generateObjectFiles(descriptors []*TypescriptMessageEnumDescriptors, includeEnums bool) []*plugin.CodeGeneratorResponse_File {
	var files []*plugin.CodeGeneratorResponse_File
	for _, descriptor := range descriptors {
		files = append(files, generateObjectFile(descriptor, includeEnums))
	}
	return files
}

func generateObjectFile(descriptor *TypescriptMessageEnumDescriptors, includeEnums bool) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"object",
		templates.ObjectsTemplate,
		fmt.Sprintf("/_internal/objects/%s.ts", descriptor.TsFileName()),
		struct {
			Descriptor *TypescriptMessageEnumDescriptors
			IncludeEnums bool
		}{
			Descriptor: descriptor,
			IncludeEnums: includeEnums,
		},
	)
}

func generateObjectIndexFile(descriptors []*TypescriptMessageEnumDescriptors) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"objectIndex",
		templates.ObjectIndexTemplate,
		"/_internal/objects/index.ts",
		descriptors,
	)
}

// ********************************************
// API Service
// ********************************************
func generateAPIServiceFiles(p *internal.Proto) []*plugin.CodeGeneratorResponse_File {
	services := make([]*plugin.CodeGeneratorResponse_File, len(p.Services))
	for i, s := range p.Services {
		services[i] = generateAPIServiceFile(s)
	}
	return services
}

func generateAPIServiceFile(s *internal.Service) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"apiService",
		templates.APIServiceTemplate,
		fmt.Sprintf("/_internal/%s.api.service.ts", s.TsFileName()),
		s,
	)
}

// ********************************************
// Main Index
// ********************************************
func generateIndexFile(p *internal.Proto, includeEnums bool) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"index",
		templates.IndexTemplate,
		"/_internal/index.ts",
		struct {
			Proto *internal.Proto
			IncludeEnums bool
		}{
			Proto: p,
			IncludeEnums: includeEnums,
		},
	)
}

// ********************************************
// Module
// ********************************************
func generateModuleFile(p *internal.Proto) *plugin.CodeGeneratorResponse_File {
	return internal.TemplateFileGenerator(
		"module",
		templates.ModuleTemplate,
		fmt.Sprintf("%s.module.ts", p.TsFileName()),
		p,
	)
}
