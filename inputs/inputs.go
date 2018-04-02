package inputs

import (
	"github.com/vendasta/gosdks/util"
	"fmt"
	"strings"
)

type Provider interface {
	GetRPCInstruction(instruction string, defaultValue interface{}) interface{}
	GetFieldOverride(fieldBreadcrumb BreadCrumb, defaultValue interface{}) interface{}
	GetFieldExclusion(fieldBreadcrumb string) bool
	GetFieldInstruction(fieldBreadcrumb BreadCrumb, instructionKey string, defaultValue interface{}) interface{}
}

func New(rpcName string, overrides map[string]interface{}, instructions map[string]interface{}, exclusions map[string]bool) Provider {
	var out Provider
	out = &inputs{
		RPCName: rpcName,
		overrides: overrides,
		instructions: instructions,
		exclusions: exclusions,
	}
	return out
}

type inputs struct {
	RPCName        string
	overrides      map[string]interface{}
	instructions   map[string]interface{}
	exclusions     map[string]bool
}

func (c *inputs) GetRPCInstruction(instruction string, defaultValue interface{}) interface{} {
	var statusCode interface{}
	statusCode, ok := c.instructions[instruction]
	if ok {
		return statusCode
	}
	return defaultValue
}

func (c *inputs) GetFieldOverride(fieldBreadcrumb BreadCrumb, defaultValue interface{}) interface{} {
	return getFieldConfig(c.overrides, fieldBreadcrumb, defaultValue)
}

func (c *inputs) GetFieldInstruction(fieldBreadcrumb BreadCrumb, instructionKey string, defaultValue interface{}) interface{} {

	fields, ok := c.instructions["fields"].(map[string]interface{})
	if !ok {
		return defaultValue
	}
	cf, ok := getFieldConfig(fields, fieldBreadcrumb, defaultValue).(map[string]interface{})
	if !ok {
		return defaultValue
	}
	instruction, ok := cf[instructionKey]
	if ok {
		return instruction
	}
	// else not specified in config file
	return defaultValue
}

func getFieldConfig(fields map[string]interface{}, b BreadCrumb, defaultValue interface{}) interface{} {

	indexed := buildIndexedKey(b)
	c := getConfig(fields, indexed)
	if c != nil {
		return c
	}
	c = getConfig(fields, b.Unindexed())
	if c != nil {
		return c
	}
	camelKey := util.ToCamelCase(b.Unindexed())
	c = getConfig(fields, camelKey)
	if c != nil {
		return c
	}
	return defaultValue
}

func buildIndexedKey(fieldBreadCrumb BreadCrumb) string {
	keyParts := strings.Split(fieldBreadCrumb.Unindexed(), ".")
	iz := fieldBreadCrumb.GetIndividualizer()
	var outParts []string
	for i := 0; i < len(keyParts); i++ {
		i2 := iz % 10
		if i > 0 {
			i2 = iz / (10 * i)
		}
		lastDigit := i2 % 10
		iPart := len(keyParts) - 1 - i
		part := keyParts[iPart]
		if lastDigit > 0 {
			part = fmt.Sprintf("%s#%d", keyParts[iPart], lastDigit)
		}
		outParts = append([]string{part}, outParts...)
	}
	return strings.Join(outParts, ".")
}

func getConfig(fields map[string]interface{}, fieldBreadcrumb string) interface{} {
	i, ok := fields[fieldBreadcrumb]
	if ok {
		if i != nil {
			return i
		}
	}
	return nil
}

// GetFieldExclusion returns true if the given breadcrumb has been excluded
func (c *inputs) GetFieldExclusion(fieldBreadcrumb string) bool {
	i, ok := c.exclusions[fieldBreadcrumb]
	if ok {
		return i
	}
	//i, ok = c.exclusions[util.ToCamelCase(fieldBreadcrumb)]
	//if ok {
	//	return i
	//}
	return false
}

