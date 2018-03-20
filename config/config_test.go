package config_test

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/emicklei/proto"
)

func TestParsesLegacyFormatOverrideFromFile(t *testing.T) {
	// RPCs uses to be formatted Like.This, now they're formatted Like/This
	c, _ := config.ReadFile("./test_data/override_legacy.json")
	breadcrumb := "a.field"
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
	breadcrumb := "a.field"
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

func TestParsesOverrideForFieldGivenInSnakeCase(t *testing.T) {
	c, _ := config.ReadFile("./test_data/override_snake.json")
	service := proto.Service{
		Name: "AnyService",
	}
	rpc := proto.RPC{
		Name: "SomeRPC",
	}
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	override := inputs.GetFieldOverride("a.longer_field", nil)
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
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	override := inputs.GetFieldOverride("a.longer_field", nil)
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
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldInstruction("a.field", "an_instruction", nil)
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
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldInstruction("a.longer_field", "an_instruction", nil)
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
	inputs, _ := config.GetInputsForRPC(service, rpc, c.ConfigMap)
	instruction := inputs.GetFieldInstruction("a.longer_field", "an_instruction", nil)
	assert.Equal(t, instruction, "value specified in file")
}

