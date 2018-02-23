package internal

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"bytes"
	"strings"
	"text/template"
)

func TemplateFileGenerator(templateName string, htmlTemplate string, fileName string, contents interface{}) *plugin.CodeGeneratorResponse_File {
	tmpl := template.Must(template.New(templateName).Parse(htmlTemplate))
	b := new(bytes.Buffer)

	err := tmpl.Execute(b, contents)
	if err != nil {
		panic(err)
	}
	content := strings.Replace(b.String(), "\t", "    ", -1)
	return &plugin.CodeGeneratorResponse_File{
		Name:    &fileName,
		Content: &content,
	}
}
