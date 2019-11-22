package configwizard

import (
	"context"
	"fmt"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/vendasta/gosdks/logging"
)

func PromptForOverrides(ctx context.Context, protofilePath string) map[string]interface{} {
	definition, err := protofiles.Read(protofilePath)
	if err != nil {
		logging.Errorf(ctx, "unable to read protofile: %s", err.Error())
		return map[string]interface{}{}
	}

	//p := *parse.Package(definition.Elements)
	//
	//t, err := parse.AllFieldTypesFromProtos(c.ProtofileRootPath, definition)
	//if err != nil {
	//	return nil, fmt.Errorf("unable to extract types: %s", err.Error())
	//}

	s := parse.Services(definition.Elements)
	fmt.Println("Select an service to override:")
	for n, svc := range(s) {
		fmt.Printf("[%d] - %s\n", n + 1, svc.Name)
	}
	return map[string]interface{}{}
}
