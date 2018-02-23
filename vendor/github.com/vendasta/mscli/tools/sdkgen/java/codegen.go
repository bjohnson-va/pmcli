package java

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/vendasta/mscli/tools/sdkgen/internal"
)

func GenerateFileHeader(packageName, javaOuterClassName string, packages []string) []byte {
	buffer := bytes.NewBufferString("")
	//Generated Code, do not modify
	const templateText = `package {{.JavaPackage}};
{{range .Packages}}
import {{.}};{{end}}
import {{.JavaPackage}}.{{.JavaOuterClassName}};

`

	tmpl, err := template.New("fileHeaderTemplate").Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	type templateData struct {
		JavaPackage        string
		JavaOuterClassName string
		Packages           []string
	}
	err = tmpl.Execute(buffer, templateData{
		JavaPackage:        packageName,
		JavaOuterClassName: javaOuterClassName,
		Packages:           packages})
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
	return buffer.Bytes()
}

//GenerateFiles follows the interface required for protoc plugins (easy to convert)
func GenerateFiles(req *plugin.CodeGeneratorRequest) (resp *plugin.CodeGeneratorResponse) {
	defer func() {
		if r := recover(); r != nil {
			errString := fmt.Sprintf("%#v", r)
			resp.Error = &errString
		}
	}()
	resp = &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{},
	}

	//Populate the resolver
	resolver := NewResolver(req)

	//Default enum headers
	enumHeaders := []string{
		"java.util.ArrayList",
		"java.util.List",
		"java.util.Map",
		"java.util.HashMap",
		"java.util.Collections",
	}

	//TODO: Calculate this
	messageHeaders := []string{
		"java.util.List",
		"java.util.ArrayList",
		"java.util.Date",
		"java.util.Map",
		"java.util.HashMap",
		"java.util.Collections",
		"java.util.Arrays",
		"org.apache.commons.lang3.StringUtils",
	}

	for _, file := range req.ProtoFile {
		if internal.IsGoogleProtoFileName(file.GetName()) {
			//Skip the google protobuf files
			fmt.Fprintf(os.Stderr, "Skipping proto %s\n", file.GetName())
			continue
		}
		javaPackageName := file.GetOptions().GetJavaPackage()
		javaDirectory := path.Join(strings.Split(javaPackageName, ".")...)

		//Generate the top-level enums
		for _, e := range file.GetEnumType() {
			headerBytes := GenerateFileHeader(javaPackageName, resolver.GetEnumJavaOuterClassName(e), enumHeaders)
			bodyBytes := GenerateEnum(resolver, e)
			outputString := string(headerBytes) + string(bodyBytes)
			outputName := path.Join(javaDirectory, fmt.Sprintf("%s.java", resolver.GetEnumJavaName(e)))
			outputFile := plugin.CodeGeneratorResponse_File{
				Name:    &outputName,
				Content: &outputString,
			}
			resp.File = append(resp.File, &outputFile)
		}

		//Generate a Message files
		for _, m := range file.GetMessageType() {
			headerBytes := GenerateFileHeader(javaPackageName, resolver.GetMessageJavaOuterClassName(m), messageHeaders)
			bodyBytes := GenerateClass(resolver, m)
			outputString := string(headerBytes) + string(bodyBytes)
			outputName := path.Join(javaDirectory, fmt.Sprintf("%s.java", resolver.GetMessageJavaName(m)))
			outputFile := plugin.CodeGeneratorResponse_File{
				Name:    &outputName,
				Content: &outputString,
			}
			resp.File = append(resp.File, &outputFile)
		}
	}
	return resp
}
