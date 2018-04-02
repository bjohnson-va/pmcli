package config_test

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/emicklei/proto"
	"github.com/bjohnson-va/pmcli/inputs"
)

func TestParsesLegacyFormatOverrideFromFile(t *testing.T) {
	// RPCs uses to be formatted Like.This, now they're formatted Like/This
	c, _ := config.ReadFile("./test_data/override_legacy.json")
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("Field")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	override := inputs.GetFieldOverride(breadcrumb, nil)
	assert.Equal(t, override, "value specified in file")
}

func TestParsesKnownGoodOverrideFromFile(t *testing.T) {
	c, _ := config.ReadFile("./test_data/override.json")
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("field")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	is, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	override := is.GetFieldOverride(breadcrumb, nil)
	assert.Equal(t, override, "value specified in file")
}

func TestParsesOverrideForFieldGivenInSnakeCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/override_snake.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	is, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("longer_field")
	override := is.GetFieldOverride(breadcrumb, nil)
	assert.Equal(t, override, "value specified in file")
}

func TestParsesOverrideForFieldGivenInCamelCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/override_camel.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	is, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("longer_field")
	override := is.GetFieldOverride(breadcrumb, nil)
	assert.Equal(t, override, "value specified in file")
}

func TestParsesKnownGoodInstructionFromFile(t *testing.T) {
	c, _ := config.ReadFile("./test_data/instruction.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetRPCInstruction("an_instruction", nil)
	assert.Equal(t, instruction, "value specified in file")
}

func TestParsesKnownGoodFieldInstructionFromFile(t *testing.T) {
	c, _ := config.ReadFile("./test_data/instruction_field.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	is, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("field")
	instruction := is.GetFieldInstruction(breadcrumb, "an_instruction", nil)
	assert.Equal(t, instruction, "value specified in file")
}

func TestParsesInstructionForFieldGivenInSnakeCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/instruction_field_snake.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	is, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("longer_field")
	instruction := is.GetFieldInstruction(breadcrumb, "an_instruction", nil)
	assert.Equal(t, instruction, "value specified in file")
}

func TestParsesInstructionForFieldGivenInCamelCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/instruction_field_camel.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	ins, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	breadcrumb := inputs.InitialBreadCrumb().AndField("a").AndField("longer_field")
	instruction := ins.GetFieldInstruction(breadcrumb, "an_instruction", nil)
	assert.Equal(t, instruction, "value specified in file")
}

func TestParsesKnownGoodFieldExclusionFromFileWithLegacyRPCFormat(t *testing.T) {
	c, _ := config.ReadFile("./test_data/exclusion_legacy_single_field.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldExclusion("a.field")
	assert.Equal(t, instruction, true)
}

func TestParsesKnownGoodFieldExclusionFromFile(t *testing.T) {
	c, _ := config.ReadFile("./test_data/exclusion_single_field.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldExclusion("a.field")
	assert.Equal(t, instruction, true)
}

func TestParsesExclusionForFieldGivenInSnakeCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/exclusion_single_field_snake.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldExclusion("some.field_with_a_long_name")
	assert.Equal(t, instruction, true)
}

func TestParsesExclusionForFieldGivenInCamelCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/exclusion_single_field_camel.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldExclusion("some.fieldWithALongName")
	assert.Equal(t, instruction, true)
}

func TestReturnsFalseExclusionForFieldNotInFile(t *testing.T) {
	c, _ := config.ReadFile("./test_data/exclusion_single_field.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldExclusion("doNotSpecifyThisInTestFile")
	assert.Equal(t, instruction, false)
}


