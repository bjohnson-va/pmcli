package config

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/vendasta/gosdks/util"
)

type Inputs struct {
	RPCName        string
	ProtofileNames []string
	overrides      map[string]interface{}
	instructions   map[string]interface{}
	exclusions     map[string]bool
}

func (c *Inputs) GetRPCInstruction(instruction string, defaultValue interface{}) interface{} {
	var statusCode interface{}
	statusCode, ok := c.instructions[instruction]
	if ok {
		return statusCode
	}
	return defaultValue
}

func (c *Inputs) GetFieldOverride(fieldBreadcrumb string, defaultValue interface{}) interface{} {
	return getFieldConfig(c.overrides, fieldBreadcrumb, defaultValue)
}

func (c *Inputs) GetFieldInstruction(fieldBreadcrumb string, instructionKey string, defaultValue interface{}) interface{} {

	fields, ok := c.instructions["fields"].(map[string]interface{})
	if !ok {
		return defaultValue
	}
	cf := getFieldConfig(fields, fieldBreadcrumb, defaultValue).(map[string]interface{})
	instruction, ok := cf[instructionKey]
	if ok {
		return instruction
	}
	// else not specified in config file
	return defaultValue
}

func getFieldConfig(fields map[string]interface{}, fieldBreadcrumb string, defaultValue interface{}) interface{} {
	c := getConfig(fields, fieldBreadcrumb)
	if c != nil {
		return c
	}
	camelKey := util.ToCamelCase(fieldBreadcrumb)
	c = getConfig(fields, camelKey)
	if c != nil {
		return c
	}
	return defaultValue
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
func (c *Inputs) GetFieldExclusion(fieldBreadcrumb string) bool {
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

func GetInputsForRPC(s proto.Service, r proto.RPC, config map[string]interface{}) (*Inputs, error) {
	i, err := readForRPC("instructions", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read instructions: %s", err.Error())
	}
	o, err := readForRPC("overrides", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}
	e, err := readListForRPC("exclusions", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}
	return &Inputs{
		RPCName:      r.Name,
		instructions: i,
		overrides:    o,
		exclusions:   e,
	}, nil
}

func readForRPC(category string, s proto.Service, r proto.RPC, config map[string]interface{}) (map[string]interface{}, error) {
	c, ok := config[category]
	if !ok {
		return map[string]interface{}{}, nil
	}
	configMap, ok := c.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid format for config section: %s", category)
	}
	key := s.Name + "/" + r.Name
	i, ok := configMap[key]
	if !ok {
		legacyKey := s.Name + "." + r.Name
		i, ok = configMap[legacyKey]
		if !ok {
			return map[string]interface{}{}, nil
		}
	}
	a, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("bad config value for key %s: %#v", key, i)
	}
	return a, nil
}

func readListForRPC(category string, s proto.Service, r proto.RPC, config map[string]interface{}) (map[string]bool, error) {
	c, ok := config[category]
	if !ok {
		return map[string]bool{}, nil
	}
	exclusionsByRPC, ok := c.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid format for exclusion section: %s", category)
	}
	rpcKey := s.Name + "/" + r.Name
	i, ok := exclusionsByRPC[rpcKey]
	if !ok {
		legacyKey := s.Name + "." + r.Name
		i, ok = exclusionsByRPC[legacyKey]
		if !ok {
			return map[string]bool{}, nil
		}
	}
	a, ok := i.([]interface{})
	if !ok {
		return nil, fmt.Errorf("Couldnt get list for key %s: %#v", rpcKey, i)
	}
	var out = map[string]bool{}
	for _, x := range a {
		out[x.(string)] = true
	}
	return out, nil
}

type File struct {
	ConfigMap      map[string]interface{}
	ProtofileNames []string
	AllowedOrigin  string
	Port           int64
}

// TODO: Don't use this legacy fallback in v2.0.0+
var legacyFallback = File{
	ConfigMap:      make(map[string]interface{}),
	ProtofileNames: []string{"advertising/v1/api.proto"},
}

func ReadFile(filename string) (*File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//return nil, fmt.Errorf("Could not find config file at %s", filename)
		fs := "WARNING: OVERRIDES AND INSTRUCTIONS WILL NOT WORK" +
			"Using fallback of %s for protofiles. Please create a %s file.\n"
		fmt.Printf(fs, legacyFallback.ProtofileNames, filename)
		return &legacyFallback, nil
	}
	r, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %s", err.Error())
	}
	f, err := parseConfig(r)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %s", err.Error())
	}
	return f, nil
}

func parseConfig(fileContents []byte) (*File, error) {
	var data interface{}
	err := json.Unmarshal(fileContents, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %s", err.Error())
	}

	i, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("couldn't turn config into map: %s")
	}

	k := "protofiles"
	p, ok := i[k].([]interface{})
	if !ok {
		//return nil, fmt.Errorf("nothing to mock. `protofiles` missing in %s", filename)
		fs := "WARNING: OVERRIDES AND INSTRUCTIONS WILL NOT WORK" +
			"Using fallback of %s for protofiles. Please specify %s.\n"
		fmt.Printf(fs, legacyFallback.ProtofileNames, k)
		return &legacyFallback, nil
	}
	protofiles := make([]string, len(p))
	for idx, pf := range p {
		protofiles[idx] = fmt.Sprintf("%s", pf)
	}

	allowedOrigin, ok := i["allowedOrigin"].(string)
	if !ok {
		allowedOrigin = ""
	}
	pf, ok := i["port"].(float64)
	port := int64(-1)
	if ok {
		port = int64(pf)
	}

	return &File{
		ConfigMap:      i,
		ProtofileNames: protofiles,
		AllowedOrigin:  allowedOrigin,
		Port:           port,
	}, nil
}
