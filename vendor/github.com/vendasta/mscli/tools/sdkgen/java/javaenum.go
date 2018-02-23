package java

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func GenerateEnum(resolver ResolverInterface, e *descriptor.EnumDescriptorProto) []byte {
	const templateText = `/**
{{enumDocs .}}
 **/
public enum {{.GetName}} { 
	{{javaEnumValues .}};
	
	private static final Map<{{protoEnumName .}}, {{javaEnumName .}}> protoTypeMap;
	static {
		Map<{{protoEnumName .}}, {{javaEnumName .}}> aMap;
		aMap = new HashMap<{{protoEnumName .}}, {{javaEnumName .}}>();
		{{range $k, $v := .Value}}aMap.put({{protoEnumName $}}.{{$v.Name}}, {{javaEnumName $}}.{{$v.Name}});
		{{end}}protoTypeMap = Collections.unmodifiableMap(aMap);
	}

	private static final Map<{{javaEnumName .}}, {{protoEnumName .}}> javaTypeMap;
	static {
		Map<{{javaEnumName .}}, {{protoEnumName .}}> aMap;
		aMap = new HashMap<{{javaEnumName .}}, {{protoEnumName .}}>();
		{{range $k, $v := .Value}}aMap.put({{javaEnumName $}}.{{$v.Name}}, {{protoEnumName $}}.{{$v.Name}});
		{{end}}javaTypeMap = Collections.unmodifiableMap(aMap);
	}

	private int value;

	private {{.Name}}(int i) {
		value = i;
	}
	public int getValue() {
		return value;
	}

	public static {{.Name}} fromProto({{protoEnumName .}} proto) {
		return protoTypeMap.get(proto);
	}

	public static List<{{.Name}}> fromProtos(List<{{protoEnumName .}}> protos) {
		List<{{.Name}}> result = new ArrayList<{{.Name}}>();
		for({{protoEnumName .}} proto : protos) {
			result.add({{.Name}}.fromProto(proto));
		}
		return result;
	}

	public {{protoEnumName .}} toProto() {
		return javaTypeMap.get(this);
	}

	public static List<{{protoEnumName .}}> toProtos(List<{{.Name}}> objects) {
		List<{{protoEnumName .}}> result = new ArrayList<{{protoEnumName .}}>();
		for({{.Name}} obj : objects) {
			result.add(obj.toProto());
		}
		return result;
	}
}`

	funcMap := template.FuncMap{
		"protoEnumName": func(e *descriptor.EnumDescriptorProto) string {
			return resolver.GetEnumProtoName(e)
		},
		"javaEnumName": func(e *descriptor.EnumDescriptorProto) string {
			return resolver.GetEnumJavaName(e)
		},
		"javaEnumValues": func(e *descriptor.EnumDescriptorProto) string {
			rv := ""
			for i, v := range e.GetValue() {
				if i != 0 {
					rv += ",\n\t\t"
				}
				rv += fmt.Sprintf("%s(%d)", v.GetName(), v.GetNumber())
			}
			return rv
		},
		"enumDocs": func(e *descriptor.EnumDescriptorProto) string {
			doc := resolver.GetEnumDoc(e)
			return tabCommentPad(doc, 0)
		},
	}

	buffer := bytes.NewBufferString("")
	tmpl, err := template.New("enumTemplate").Funcs(funcMap).Parse(templateText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %s", err.Error())
		os.Exit(-1)
	}

	err = tmpl.Execute(buffer, e)
	if err != nil {
		fmt.Fprintf(os.Stderr, "execution: %s", err.Error())
		os.Exit(-1)
	}
	return buffer.Bytes()
}
