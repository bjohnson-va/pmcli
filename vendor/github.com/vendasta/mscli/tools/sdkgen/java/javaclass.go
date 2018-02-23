package java

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

//----------------------------------------------------------------------

func tabPad(input string, depth int) string {
	lines := strings.Split(input, "\n")
	rv := ""
	for i, line := range lines {
		if i > 0 {
			rv += "\n"
		}
		rv += fmt.Sprintf("%s%s", strings.Repeat("\t", depth), line)
	}
	return rv
}

func tabCommentPad(input string, depth int) string {
	lines := strings.Split(input, "\n")
	rv := ""
	for i, line := range lines {
		if i > 0 {
			rv += "\n"
		}
		rv += fmt.Sprintf("%s * %s", strings.Repeat("\t", depth), line)
	}
	return rv
}

func isMessage(field *descriptor.FieldDescriptorProto) bool {
	return field.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE
}

func GenerateClass(resolver ResolverInterface, message *descriptor.DescriptorProto) []byte {
	const templateText = `/**
{{classDocs . 0}}
 **/
public {{classExposure .}} class {{.GetName}} {
{{range .GetEnumType}}
{{generateEnum .}}{{end}}

{{range .GetNestedType}}
{{generateClass .}}{{end}}

	{{range .Field}}private final {{getJavaType .}} {{.GetJsonName}};
	{{end}}

	private {{getClassName .}} ({{range $i, $f := .Field}}{{if $i}},{{end}}
		final {{getJavaType .}} {{.GetJsonName}}{{end}})
		
	{
		{{range .Field}}this.{{.GetJsonName}} = {{.GetJsonName}};
		{{end}}
	}
	{{range .Field}}
	/**
{{fieldDocs . 1}}
      * @return The final value of {{.GetJsonName}} on the object
	 **/
	public {{getJavaType .}} get{{title .GetJsonName}}() {
		return this.{{.GetJsonName}};
	}
	{{end}}

	public static class Builder {
		{{range .Field}}private {{getJavaType .}} {{.GetJsonName}};
		{{end}}
		public Builder() {
			{{range .Field}}this.{{.GetJsonName}} = {{getDefault .}};
			{{end}}
		}
		{{range .Field}}
		/**
		  * Adds a value to the builder for {{.GetJsonName}}
		  * @param {{.GetJsonName}} Value to assign to the mutable Builder
		  * @return The Builder instance so that call chaining works
		 **/
		public Builder set{{title .GetJsonName}}({{getJavaType .}} {{.GetJsonName}}) {
			this.{{.GetJsonName}} = {{.GetJsonName}};
			return this;
		}
		{{end}}
		/**
		  * Takes the configuration in the mutable Builder and uses it to instantiate a final instance
		  * of the {{getClassName .}} class
		  * @return The instantiated final {{getClassName .}}
		 **/
		public {{getClassName .}} build() {
			return new {{getClassName .}}({{range $i, $f := .Field}}{{if $i}},{{end}}
				this.{{$f.GetJsonName}}{{end}});
		}
	}

	/**
	 * Returns a Builder for {{getClassName .}}, which is a mutable representation of the object.  Once the
	 * client has built up an object they can then create an immutable {{getClassName .}} object using the
	 * build function.
	 * @return A fresh Builder instance with no values set
	 **/
	public static Builder newBuilder() {
		return new Builder();
	}

	/**
	 * Provides a human-readable representation of this object.  Useful for debugging.
	 * @return A string representation of the {{getClassName .}} instance
	 **/
	 public String toString() {
		 String result = "{{getClassName .}}\n";
		 {{range .Field}}result += "-> {{.GetJsonName}}: ({{getJavaType .}})"
		     + StringUtils.join("\n  ", Arrays.asList(String.valueOf(this.{{.GetJsonName}}).split("\n"))) + "\n"; 
		 {{end}}
		 return result;
	 }
	/**
	* Allows for simple conversion between the low-level generated protobuf object to
	* {{getClassName .}}, which is much more usable.
	* @return An instance of {{getClassName .}} representing the input proto object
	**/
	public static {{getClassName .}} fromProto({{getProtoClassName .}} proto) {
		{{getClassName .}} out = null;
		if (proto != null) {
			{{getClassName .}}.Builder outBuilder = {{getClassName .}}.newBuilder(){{range .Field}}
			{{assign .}}{{end}};
			out = outBuilder.build();
		}
		return out;
	}

	/**
	* Convenience method for handling lists of proto objects.  It calls .fromProto on each one
	* and returns a list of the converted results.
	* @return A list of {{getClassName .}} instances representing the input proto objects
	**/
	public static List<{{getClassName .}}> fromProtos(List<{{getProtoClassName .}}> protos) {
		List<{{getClassName .}}> out = new ArrayList<{{getClassName .}}>();
		for({{getProtoClassName .}} proto : protos) {
			out.add({{getClassName .}}.fromProto(proto));
		}
		return out;
	}

	/**
	 * Allows for simple conversion of an object to the low-level generated protobuf object.
	 * @return An instance of {{getProtoClassName .}} which is a proto object ready for wire transmission
	 **/
	 public {{getProtoClassName .}} toProto() {
		 {{getClassName .}} obj = this;
		 {{getProtoClassName .}}.Builder outBuilder = {{getProtoClassName .}}.newBuilder();{{range .Field}}
		 {{assignProto .}}{{end}}
		 return outBuilder.build();
	 }

	 /**
	  * Convenience method for handling lists of objects.  It calls .toProto on each one and
	  * returns a list of the converted results.
	  * @return A list of {{getProtoClassName .}} instances representing the input objects.
	  */
	public static List<{{getProtoClassName .}}> toProtos(List<{{getClassName .}}> objects) {
		List<{{getProtoClassName .}}> out = new ArrayList<{{getProtoClassName .}}>();
		if(objects != null) {
			for ({{getClassName .}} obj : objects) {
				out.add(obj!=null?obj.toProto():null);
			}
		}
		return out;
	}
}`

	funcMap := template.FuncMap{
		"title": strings.Title,
		"classDocs": func(f *descriptor.DescriptorProto, depth int) string {
			doc := resolver.GetMessageDoc(f)
			return tabCommentPad(doc, depth)
		},
		"fieldDocs": func(f *descriptor.FieldDescriptorProto, depth int) string {
			doc := resolver.GetFieldDoc(f)
			return tabCommentPad(doc, depth)
		},
		"getProtoClassName": func(m *descriptor.DescriptorProto) string {
			return resolver.GetMessageProtoName(m)
		},
		"getClassName": func(m *descriptor.DescriptorProto) string {
			return strings.Title(m.GetName())
		},
		"getJavaType": func(f *descriptor.FieldDescriptorProto) string {
			return resolver.GetFieldJavaType(f)
		},
		"classExposure": func(m *descriptor.DescriptorProto) string {
			if resolver.IsMessageNested(m) {
				return "static final"
			}
			return "final"
		},
		"isMessage": isMessage,
		"assign": func(f *descriptor.FieldDescriptorProto) string {
			isRepeated := f.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
			isEnum := f.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM
			isTimestamp := f.GetTypeName() == ".google.protobuf.Timestamp"
			isBoolValue := f.GetTypeName() == ".google.protobuf.BoolValue"
			if isTimestamp {
				if isRepeated {
					//Date type, repeated
					return fmt.Sprintf(".set%s(/* Not Supported */)",
						strings.Title((f.GetJsonName())))
				}
				//Date type
				return fmt.Sprintf(".set%s(proto.has%s()?new Date(proto.get%s().getSeconds() * 1000):null)",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			} else if isBoolValue {
				if isRepeated {
					//Built-in type, repeated
					return fmt.Sprintf(".set%s(proto.get%sList())",
						strings.Title(f.GetJsonName()),
						strings.Title(f.GetJsonName()))
				}
				//Built-in type
				return fmt.Sprintf(".set%s(proto.get%s())",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			} else if isMessage(f) {
				if isRepeated {
					//Custom type, repeated
					return fmt.Sprintf(".set%s(%s.fromProtos(proto.get%sList()))",
						strings.Title(f.GetJsonName()),
						resolver.GetFieldJavaBaseType(f),
						strings.Title(f.GetJsonName()))
				}
				//Custom type
				return fmt.Sprintf(".set%s(%s.fromProto(proto.get%s()))",
					strings.Title(f.GetJsonName()),
					resolver.GetFieldJavaType(f),
					strings.Title(f.GetJsonName()))
			}
			if isEnum {
				//Lookup the enum type
				e := resolver.GetEnumByProtoName(f.GetTypeName())
				if isRepeated {
					//Enum type, repeated
					return fmt.Sprintf(".set%s(%s.fromProtos(proto.get%sList()))",
						strings.Title(f.GetJsonName()),
						resolver.GetEnumJavaName(e),
						strings.Title(f.GetJsonName()))
				}
				//Enum type
				return fmt.Sprintf(".set%s(%s.fromProto(proto.get%s()))",
					strings.Title(f.GetJsonName()),
					resolver.GetEnumJavaName(e),
					strings.Title(f.GetJsonName()))
			}
			if isRepeated {
				//Built-in type, repeated
				return fmt.Sprintf(".set%s(proto.get%sList())",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			}
			//Built-in type
			return fmt.Sprintf(".set%s(proto.get%s())",
				strings.Title(f.GetJsonName()),
				strings.Title(f.GetJsonName()))
		},
		"assignProto": func(f *descriptor.FieldDescriptorProto) string {
			isRepeated := f.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED
			isEnum := f.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM
			isTimestamp := f.GetTypeName() == ".google.protobuf.Timestamp"
			isBoolValue := f.GetTypeName() == ".google.protobuf.BoolValue"
			if isTimestamp {
				if isRepeated {
					//Date type, repeated
					return fmt.Sprintf("outBuilder.set%s(/* Not Supported */);",
						strings.Title(f.GetJsonName()))
				}
				//Date type
				return fmt.Sprintf("if(obj.get%s()!=null){outBuilder.set%s(com.google.protobuf.Timestamp.newBuilder().setSeconds(obj.get%s().getTime() / 1000).build());}",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			} else if isBoolValue {
				if isRepeated {
					//Built-in type, repeated
					return fmt.Sprintf("outBuilder.addAll%s(obj.get%s());",
						strings.Title(f.GetJsonName()),
						strings.Title(f.GetJsonName()))
				}
				//Built-in type
				return fmt.Sprintf("outBuilder.set%s(obj.get%s()!=null?obj.get%s():null);",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			} else if isMessage(f) {
				if isRepeated {
					//Custom type, repeated
					return fmt.Sprintf("outBuilder.addAll%s(%s.toProtos(obj.get%s()));",
						strings.Title(f.GetJsonName()),
						resolver.GetFieldJavaBaseType(f),
						strings.Title(f.GetJsonName()))
				}
				//Custom type
				return fmt.Sprintf("if(obj.get%s() != null){outBuilder.set%s(obj.get%s().toProto());}",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			}
			if isEnum {
				//Lookup the enum type
				e := resolver.GetEnumByProtoName(f.GetTypeName())
				if isRepeated {
					//Enum type, repeated
					return fmt.Sprintf("outBuilder.addAll%s(%s.toProtos(obj.get%s()));",
						strings.Title(f.GetJsonName()),
						resolver.GetEnumJavaName(e),
						strings.Title(f.GetJsonName()))
				}
				//Enum type
				return fmt.Sprintf("outBuilder.set%s(obj.get%s() != null?obj.get%s().toProto():null);",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			}
			if isRepeated {
				//Built-in type, repeated
				return fmt.Sprintf("outBuilder.addAll%s(obj.get%s());",
					strings.Title(f.GetJsonName()),
					strings.Title(f.GetJsonName()))
			}
			//Built-in type
			return fmt.Sprintf("outBuilder.set%s(obj.get%s());",
				strings.Title(f.GetJsonName()),
				strings.Title(f.GetJsonName()))
		},
		"getDefault": func(f *descriptor.FieldDescriptorProto) string {
			return resolver.GetFieldJavaDefault(f)
		},
		"generateEnum": func(e *descriptor.EnumDescriptorProto) string {
			output := GenerateEnum(resolver, e)
			return tabPad(string(output), 1)
		},
		"generateClass": func(m *descriptor.DescriptorProto) string {
			output := GenerateClass(resolver, m)
			return tabPad(string(output), 1)
		},
	}

	buffer := bytes.NewBufferString("")
	tmpl, err := template.New("classTemplate").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(buffer, message)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
	return buffer.Bytes()
}
