package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"github.com/emicklei/proto"
	"github.com/oliveagle/jsonpath"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/vendasta/gosdks/util"
)

type HTTPHandler struct {
	Path        string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}


func FromProtofile(allowedOrigin string, rootPath string, protofilename string, config interface{}) ([]HTTPHandler, error) {

	definition, err := protofiles.Read(fmt.Sprintf("%s/%s", rootPath, protofilename))
	if err != nil {
		return nil, fmt.Errorf("unable to read protofile: %s", err.Error())
	}

	var handlers []HTTPHandler

	pakkage := *parse.Package(definition.Elements)
	t, err := parse.AllFieldTypesFromProtos(rootPath, definition)
	if err != nil {
		return nil, fmt.Errorf("unable to extract types: %s", err.Error())
	}

	srv := parse.Services(definition.Elements)

	for _, s := range srv {
		rpcs := parse.RPCs(s.Elements)
		for _, r := range rpcs {
			p := "/" + pakkage.Name + "." + s.Name + "/" + r.Name
			c, err := readConfigForRPC(s, r, config)
			if err != nil {
				return nil, fmt.Errorf("problem reading config: %s", err.Error())
			}
			newHandler := fakeHandler(allowedOrigin, p, r, t, c)
			handlers = append(handlers, newHandler)
		}
	}
	return handlers, nil
}

func readConfigForRPC(s proto.Service, r proto.RPC, config interface{}) (*configs, error) {
	i, err := readForRPC("instructions", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read instructions: %s", err.Error())
	}
	o, err := readForRPC("overrides", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}
	return &configs{
		instructions: i,
		overrides: o,
	}, nil
}

func readForRPC(category string, s proto.Service, r proto.RPC, config interface{}) (map[string]interface{}, error) {
	jp := "$." + category + "." + s.Name + "_" + r.Name + ""
	i, err := jsonpath.JsonPathLookup(config, jp)
	if err != nil {
		// Have to assume the value just wasn't there.  Can't distinguish other errors.
		var empty map[string]interface{}
		return empty, nil
	}
	a, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("bad config value for key %s: %#v", jp, i)
	}
	return a, nil
}

type configs struct {
	overrides map[string]interface{}
	instructions map[string]interface{}
}


func fakeHandler(allowedOrigin string, path string, rpc proto.RPC, t *parse.FieldTypes, c *configs) HTTPHandler {
	if rpc.ReturnsType == "google.protobuf.Empty" {
		return HTTPHandler{
			Path: path,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
			},
		}
	}

	foundMessage := (*proto.Message)(nil)
	for _, m := range t.Messages {
		if m.Name == rpc.ReturnsType {
			foundMessage = &m;
			break;
		}
	}
	if foundMessage == nil {
		panic(fmt.Sprintf("Did not find definition for message %s", rpc.ReturnsType))
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		obj := GenerateRandomFieldsForMessage(*foundMessage, t, c)
		marshaled, _ := json.Marshal(obj)
		w.Write(([]byte)(marshaled))
	}
	return HTTPHandler{
		Path:        path,
		HandlerFunc: fn,
	}
}

func randomEnum(enum proto.Enum) proto.EnumField {
	possibleEnumValues := parse.EnumFields(enum.Elements)
	return possibleEnumValues[0] // TODO: Randomize
}

func GenerateRandomFieldsForMessage(message proto.Message, t *parse.FieldTypes, c *configs) interface{} {
	return randomFieldsForMessage("", message, t, c)
}

func randomFieldsForMessage(breadcrumb string, message proto.Message, t *parse.FieldTypes, c *configs) interface{} {
	obj := make(map[string]interface{})
	fieldz := parse.Fields(message.Elements)
	for _, f := range fieldz {

		fbreadcrumb := f.Name
		if breadcrumb != "" {
			fbreadcrumb = breadcrumb + "." + f.Name
		}
		value, err := randomFieldValue(fbreadcrumb, *f.Field, t, c)
		if err != nil {
			value = err.Error() // Expose the error to the user of the API
		}
		if f.Repeated {
			i, ok := c.instructions[fbreadcrumb].(map[string]interface{})
			if !ok {
				// Override wasn't specified for this breadcrumb
			}
			length := 1
			if i != nil {
				num, ok := i["num"]
				if ok {
					length = int(num.(float64))
				} // else not specified in config file
			}

			var vlist []interface{}
			for x := 0; x < length; x++ {
				z, err := randomFieldValue(fbreadcrumb, *f.Field, t, c)
				if err != nil {
					value = err.Error()
					continue;
				}
				vlist = append(vlist, z)
			}
			value = vlist
		}
		obj[util.ToCamelCase(f.Name)] = value
	}
	return obj
}


func randomFieldValue(breadcrumb string, field proto.Field, t *parse.FieldTypes, c *configs) (interface{}, error) {
	// TODO: randomFieldValue should product consistent pseudorandom values that dont repeat
	override, ok := c.overrides[breadcrumb] // TODO: Decide to use snake (from proto) or camel (from endpoints)
	fmt.Printf("breadcrumb is: %s %s\n", breadcrumb, util.ToCamelCase(breadcrumb))
	fmt.Printf("overrides are %#v\n", c.overrides)
	if !ok {
		// Override wasn't specified for this breadcrumb
	}
	if override != nil {
		fmt.Printf("using override %f\n", override)
		return override, nil
	}
	if field.Type == "string" || field.Type == "bytes" {
		return fmt.Sprintf("some %s %d", field.Type, rand.Intn(1000)), nil
	}
	if strings.Contains(field.Type, "int") {
		return 1234567890, nil
	}
	if strings.Contains(field.Type, "float") {
		return math.Phi, nil
	}
	if strings.Contains(field.Type, "double") {
		return math.Phi, nil
	}
	if strings.Contains(field.Type, "bool") {
		return true, nil
	}
	if strings.Contains(field.Type, "google.protobuf.Timestamp") {
		return "2016-01-01", nil // TODO: Use correct format
	}

	var isEnum bool
	fieldType := field.Type
	if strings.Contains(field.Type, ".") {
		// Probably an enum.  Eg: CampaignStatus.Status
		parts := strings.Split(field.Type, ".")
		fieldType = parts[0]
		isEnum = true;
	}

	for _, e := range t.Enums {
		if fieldType == e.Name {
			return randomEnum(e).Name, nil
		}
	}

	for _, m := range t.Messages {
		if m.Name == fieldType {
			if isEnum {
				for _, e := range parse.Enums(m.Elements) {
					return randomEnum(e).Name, nil
				}
			}
			return randomFieldsForMessage(breadcrumb, m, t, c), nil
		}
	}
	return "", fmt.Errorf("unexpected field type %s", field.Type)
}
