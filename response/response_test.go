package response_test

import (
	"testing"
	"github.com/bjohnson-va/pmcli/random"
	"github.com/bjohnson-va/pmcli/response"
	"context"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/bjohnson-va/pmcli/parse"
	"github.com/magiconair/properties/assert"
	"github.com/bjohnson-va/pmcli/inputs"
)

type trackOverrides struct {
	overridesRequested []inputs.BreadCrumb
}

func (*trackOverrides) GetRPCInstruction(instruction string, defaultValue interface{}) interface{} {
	return defaultValue
}

// TODO: Update interface to take a breadcrumb instead of a string.
// Then, we search the overrides for either un-numbered or numbered overrides
// because the breadcrumb gives us positional info
func (m *trackOverrides) GetFieldOverride(fieldBreadcrumb inputs.BreadCrumb, defaultValue interface{}) interface{} {
	m.overridesRequested = append(m.overridesRequested, fieldBreadcrumb)
	return defaultValue
}

func (*trackOverrides) GetFieldExclusion(fieldBreadcrumb string) bool {
	return false
}

func (*trackOverrides) GetFieldInstruction(fieldBreadcrumb inputs.BreadCrumb, instructionKey string, defaultValue interface{}) interface{} {
	if fieldBreadcrumb.ToString() == "outerField.repeatedField" && instructionKey == "num" {
		return 2.0
	}
	return defaultValue
}

func TestRequestsOverridesUsingExpectedBreadcrumbs(t *testing.T) {
	ctx := context.Background()

	p, _ := protofiles.Read("./test_data/outer-repeated-inner.proto")
	service := parse.Services(p.Elements)[0]
	rpc := parse.RPCs(service.Elements)[0]
	types, _ := parse.AllFieldTypesFromProtos(".", p)
	outerMessage, _ := parse.GetMessageReturnedByRPC(rpc, *types)

	var fields random.FieldProvider
	fields = random.MockFieldProvider{}

	track := trackOverrides{}
	bc := inputs.InitialBreadCrumb().AndField("outerField")
	response.GenerateForMessage(ctx, &fields, bc, *outerMessage, types, &track)

	expectedOverrideRequests := []inputs.BreadCrumb{
		inputs.InitialBreadCrumb().AndField("outerField").AndField("repeatedField").Indexed(0),
		inputs.InitialBreadCrumb().AndField("outerField").AndField("repeatedField").Indexed(0).AndField("innerField"),
		inputs.InitialBreadCrumb().AndField("outerField").AndField("repeatedField").Indexed(1),
		inputs.InitialBreadCrumb().AndField("outerField").AndField("repeatedField").Indexed(1).AndField("innerField"),
	}
	assert.Equal(t, track.overridesRequested, expectedOverrideRequests)
}
