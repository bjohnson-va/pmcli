package config

import (
	"fmt"
	"github.com/emicklei/proto"
)



type Inputs struct {
	RPCName string
	Overrides map[string]interface{}
	Instructions map[string]interface{}
}

func (c *Inputs) GetRPCInstruction(instruction string, defaultValue interface{}) interface{} {
	var statusCode interface{}
	statusCode, ok := c.Instructions[instruction]
	if ok {
		return statusCode
	}
	return defaultValue
}

func (c *Inputs) GetFieldInstruction(fieldBreadcrumb string, instruction string, defaultValue interface{}) interface{} {
	i, ok := c.Instructions[fieldBreadcrumb].(map[string]interface{})
	if ok {
		if i != nil {
			num, ok := i[instruction]
			if ok {
				return num
			} // else not specified in config file
		}
	}
	return defaultValue
}

func GetInputsForRPC(s proto.Service, r proto.RPC, config map[string]interface{}) (*Inputs, error) {
	i, err := readForRPC("instructions", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read instructions: %s", err.Error())
	}
	o, err := readForRPC("overrides", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}
	return &Inputs{
		RPCName: r.Name,
		Instructions: i,
		Overrides: o,
	}, nil
}

func readForRPC(category string, s proto.Service, r proto.RPC, config map[string]interface{}) (map[string]interface{}, error) {
	key := s.Name + "." + r.Name
	c, ok := config[category]
	if !ok {
		return map[string]interface{}{}, nil
	}
	configMap, ok := c.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid format for config section: %s", category)
	}
	i, ok := configMap[key]
	if !ok {
		return map[string]interface{}{}, nil
	}
	a, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("bad config value for key %s: %#v", key, i)
	}
	return a, nil
}
