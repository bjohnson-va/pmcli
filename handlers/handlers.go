package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/emicklei/proto"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/bjohnson-va/pmcli/config"
	"context"
	"github.com/vendasta/gosdks/logging"
	"time"
	"github.com/bjohnson-va/pmcli/random"
	"github.com/bjohnson-va/pmcli/response"
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
	rand := hbc.RandomProvider
	for _, s := range services {
		rpcs := parse.RPCs(s.Elements)
		for _, r := range rpcs {
			p := "/" + packageName + "." + s.Name + "/" + r.Name
			c, err := config.GetInputsForRPC(s, r, hbc.AllConfig)
			if err != nil {
				return nil, fmt.Errorf("problem reading config: %s", err.Error())
			}
			newHandler, err := fakeHandler(hbc.AllowedOrigin, p, r, t, rand, c)
			if err != nil {
				msg := "unable to build mock handler for %s: %s"
				return nil, fmt.Errorf(msg, p, err.Error())
			}
			handlers = append(handlers, *newHandler)
		}
	}
	return handlers, nil
}

func fakeHandler(allowedOrigin string, path string, rpc proto.RPC, t *parse.FieldTypes, p *random.FieldProvider, c config.InputsProvider) (*HTTPHandler, error) {
	ctx := context.Background() // New Handler -> new Context
	// json unmarshal defaults to float64
	statusCode := int(c.GetRPCInstruction("statusCode", 200.0).(float64))
	delaySeconds := c.GetRPCInstruction("delaySeconds", 0.0).(float64)
	emptyBody := c.GetRPCInstruction("emptyBody", false).(bool)

	if emptyBody || rpc.ReturnsType == "google.protobuf.Empty" {
		return &HTTPHandler{
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
		}, nil
	}

	foundMessage, err := parse.GetMessageReturnedByRPC(rpc, *t)
	if err != nil {
		return nil, fmt.Errorf("problem with mock handler inputs: %s", err.Error())
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
	return &HTTPHandler{
		Path:        path,
		HandlerFunc: fn,
	}, nil
}

func delay(ctx context.Context, seconds float64) {
	if seconds <= 0 {
		return
	}
	logging.Infof(ctx, "Delaying %.2f seconds as instructed in config", seconds)
	time.Sleep(time.Second * time.Duration(seconds))
}

func GenerateRandomFieldsForMessage(ctx context.Context, p *random.FieldProvider,
	message proto.Message, t *parse.FieldTypes, c config.InputsProvider) interface{} {
	return response.GenerateForMessage(ctx, p, response.Initial(), message, t, c)
}
