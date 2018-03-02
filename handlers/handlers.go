package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/emicklei/proto"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/vendasta/gosdks/util"
	"context"
	"github.com/vendasta/gosdks/logging"
	"time"
	"github.com/bjohnson-va/pmcli/random"
)

type HTTPHandler struct {
	Path        string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

type HandlerBuildingConfig struct {
	AllowedOrigin string
	ProtofileRootPath string
	AllConfig map[string]interface{}
	RandomProvider *random.FieldProvider
}

func FromProtofile(c HandlerBuildingConfig, protofileName string) ([]HTTPHandler, error) {

	a := fmt.Sprintf("%s/%s", c.ProtofileRootPath, protofileName)
	definition, err := protofiles.Read(a)
	if err != nil {
		return nil, fmt.Errorf("unable to read protofile: %s", err.Error())
	}

	p := *parse.Package(definition.Elements)

	t, err := parse.AllFieldTypesFromProtos(c.ProtofileRootPath, definition)
	if err != nil {
		return nil, fmt.Errorf("unable to extract types: %s", err.Error())
	}

	s := parse.Services(definition.Elements)
	h, err := buildHandlersForServices(c, s, p.Name, t)
	if err != nil {
		return nil, fmt.Errorf("unable to build handlers: %s", err.Error())
	}
	return h, nil
}

func buildHandlersForServices(hbc HandlerBuildingConfig, services []proto.Service, packageName string, t *parse.FieldTypes) ([]HTTPHandler, error) {
	var handlers []HTTPHandler
	for _, s := range services {
		rpcs := parse.RPCs(s.Elements)
		for _, r := range rpcs {
			p := "/" + packageName + "." + s.Name + "/" + r.Name // TODO: . should be /
			c, err := config.GetInputsForRPC(s, r, hbc.AllConfig)
			if err != nil {
				return nil, fmt.Errorf("problem reading config: %s", err.Error())
			}
			newHandler := fakeHandler(hbc.AllowedOrigin, p, r, t, hbc.RandomProvider, c)
			handlers = append(handlers, newHandler)
		}
	}
	return handlers, nil
}




func fakeHandler(allowedOrigin string, path string, rpc proto.RPC, t *parse.FieldTypes, p *random.FieldProvider, c *config.Inputs) HTTPHandler {
	ctx := context.Background() // New Handler -> new Context
	// json unmarshal defaults to float64
	statusCode := int(c.GetRPCInstruction("statusCode", 200.0).(float64))
	delaySeconds := c.GetRPCInstruction("delaySeconds", 0.0).(float64)
	emptyBody := c.GetRPCInstruction("emptyBody", false).(bool)

	if emptyBody || rpc.ReturnsType == "google.protobuf.Empty" {
		return HTTPHandler{
			Path: path,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				if r.Method == "OPTIONS" {
					w.WriteHeader(200)
					return
				}
				delay(ctx, delaySeconds)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(statusCode)
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
		w.WriteHeader(statusCode)

		delay(ctx, delaySeconds)
		obj := GenerateRandomFieldsForMessage(ctx, p, *foundMessage, t, c)
		marshaled, _ := json.Marshal(obj)
		w.Write(([]byte)(marshaled))
	}
	return HTTPHandler{
		Path:        path,
		HandlerFunc: fn,
	}
}

func delay(ctx context.Context, seconds float64) {
	if seconds <= 0 {
		return
	}
	logging.Infof(ctx, "Delaying %.2f seconds as instructed in config", seconds)
	time.Sleep(time.Second * time.Duration(seconds))
}

func randomEnum(enum proto.Enum) proto.EnumField {
	possibleEnumValues := parse.EnumFields(enum.Elements)
	return possibleEnumValues[0] // TODO: Randomize
}

func GenerateRandomFieldsForMessage(ctx context.Context, p *random.FieldProvider,
	message proto.Message, t *parse.FieldTypes, c *config.Inputs) interface{} {
	return randomFieldsForMessage(ctx, p, "", message, t, c)
}

func randomFieldsForMessage(ctx context.Context, p *random.FieldProvider, breadcrumb string, message proto.Message,
	t *parse.FieldTypes, c *config.Inputs) interface{} {
	obj := make(map[string]interface{})
	fieldz := parse.Fields(message.Elements)
	for _, f := range fieldz {

		fBreadcrumb := f.Name
		if breadcrumb != "" {
			fBreadcrumb = breadcrumb + "." + f.Name
		}
		var value interface{}
		if f.Repeated {
			// json unmarshal defaults to float64
			length := int(c.GetFieldInstruction(fBreadcrumb, "num", 1.0).(float64))
			var list []interface{}
			for x := 0; x < length; x++ {
				z, err := randomFieldValue(ctx, *p, fBreadcrumb, *f.Field, t, c)
				if err != nil {
					value = err.Error()
					break
				}
				list = append(list, z)
			}
			value = list
		} else {
			var err error
			value, err = randomFieldValue(ctx, *p, fBreadcrumb, *f.Field, t, c)
			if err != nil {
				value = err.Error() // Expose the error to the user of the API
			}
		}
		obj[util.ToCamelCase(f.Name)] = value
	}
	return obj
}


func randomFieldValue(ctx context.Context, p random.FieldProvider, breadcrumb string, field proto.Field, t *parse.FieldTypes, c *config.Inputs) (interface{}, error) {
	// TODO: randomFieldValue should produce consistent pseudorandom values that dont repeat
	override, ok := c.Overrides[breadcrumb] // TODO: Decide to use snake (from proto) or camel (from endpoints)
	if !ok {
		// Override wasn't specified for this breadcrumb
	}
	if override != nil {
		logging.Infof(ctx, "Using override for %s: %s", breadcrumb, override)
		return override, nil
	}
	if field.Type == "string" || field.Type == "bytes" {
		return p.NewString(breadcrumb), nil
	}
	if strings.Contains(field.Type, "int") {
		return p.NewInt32(breadcrumb), nil
	}
	if strings.Contains(field.Type, "float") {
		return p.NewFloat32(breadcrumb), nil
	}
	if strings.Contains(field.Type, "double") {
		return p.NewFloat64(breadcrumb), nil
	}
	if strings.Contains(field.Type, "bool") {
		return p.NewBool(breadcrumb), nil
	}
	if strings.Contains(field.Type, "google.protobuf.Timestamp") {
		return p.NewTimestamp(breadcrumb), nil // TODO: Use correct format
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
			return randomFieldsForMessage(ctx, &p, breadcrumb, m, t, c), nil
		}
	}
	return "", fmt.Errorf("unexpected field type %s", field.Type)
}
