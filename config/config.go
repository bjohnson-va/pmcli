package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/emicklei/proto"
	"github.com/vendasta/gosdks/util"
)

const FILENAME = "mockserver.json"
const UseHTTPSToken = "useHttps"

type Inputs struct {
	RPCName        string
	ProtofileNames []string
	overrides      map[string]RPCFieldOverrides
	Instructions   RPCInstructions
	exclusions     map[string]bool
}

func (c *Inputs) GetFieldOverride(fieldBreadcrumb string, defaultValue interface{}) interface{} {
	v, ok := c.overrides[fieldBreadcrumb]
	if !ok {
		return defaultValue
	}
	return v
}

func (c *Inputs) GetFieldInstruction(
	fieldBreadcrumb string, instructionKey string, defaultValue interface{},
) interface{} {
	fields := c.Instructions.Fields
	if len(fields) == 0 {
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
	// i, ok = c.exclusions[util.ToCamelCase(fieldBreadcrumb)]
	// if ok {
	//	return i
	// }
	return false
}

func GetInputsForRPC(
	packageName string, s proto.Service, r proto.RPC, config Map,
) (*Inputs, error) {
	i, err := readForRPC("instructions", packageName, s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read instructions: %s", err.Error())
	}
	ic, ok := i.(RPCInstructions)
	if !ok {
		return nil, fmt.Errorf("failed to assert type for instructions")
	}
	o, err := readForRPC("overrides", packageName, s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}

	oc, ok := o.(RPCOverrides)
	if !ok {
		fmt.Printf("\n\noc ====> %+v\n\n", oc)
		return nil, fmt.Errorf("failed to assert type for overrides")
	}
	e, err := readExclusionsForRPC("exclusions", s, r, config)
	if err != nil {
		return nil, fmt.Errorf("failed to read overrides: %s", err.Error())
	}
	return &Inputs{
		RPCName:      r.Name,
		Instructions: ic,
		overrides:    oc,
		exclusions:   e,
	}, nil
}

func readForRPC(
	category string, packageName string, s proto.Service, r proto.RPC, config Map,
) (interface{}, error) {
	possibleKeys := []string{
		fmt.Sprintf("%s.%s/%s", packageName, s.Name, r.Name),
		fmt.Sprintf("/%s.%s/%s", packageName, s.Name, r.Name),
		s.Name + "/" + r.Name,
		s.Name + "." + r.Name,
	}

	var i interface{}
	var ok bool
	switch category {
	case "instructions":
		for _, k := range possibleKeys {
			i, ok = config.Instructions[k]
			if ok {
				break
			}
			fmt.Printf("No instructions found for %s\n", k)
		}
	case "overrides":
		for _, k := range possibleKeys {
			i, ok = config.Overrides[k]
			if ok {
				break
			}
		}
	case "exclusions":
		for _, k := range possibleKeys {
			i, ok = config.Exclusions[k]
			if ok {
				break
			}
		}
	}
	if i == nil {
		return map[string]interface{}{}, nil
	}
	return i, nil
}

func readExclusionsForRPC(category string, s proto.Service, r proto.RPC, config Map) (map[string]bool, error) {
	exclusionsByRPC := config.Exclusions
	rpcKey := s.Name + "/" + r.Name
	a, ok := exclusionsByRPC[rpcKey]
	if !ok {
		legacyKey := s.Name + "." + r.Name
		a, ok = exclusionsByRPC[legacyKey]
		if !ok {
			return map[string]bool{}, nil
		}
	}
	var out = map[string]bool{}
	for _, x := range a {
		out[x.(string)] = true
	}
	return out, nil
}

type EmptyBody int

func EmptyBodyFromBool(eb bool) EmptyBody {
	if eb {
		return EmptyBody_True
	}
	return EmptyBody_False
}

const (
	EmptyBody_Unset = iota
	EmptyBody_True
	EmptyBody_False
)

type RPCInstructions struct {
	DelaySecs  int
	StatusCode int
	Fields     map[string]interface{}
	EmptyBody  EmptyBody
}

func (r RPCInstructions) String() string {
	return fmt.Sprintf("RPCInstructions[Delay: %d, Code: %d, Empty: %t]", r.DelaySecs, r.StatusCode, r.EmptyBody)
}

func (i RPCInstructions) GetStatusCode() int {
	if i.StatusCode == 0 {
		return 200
	}
	return i.StatusCode
}

type RPCFieldOverrides struct {
}
type RPCOverrides map[string]RPCFieldOverrides
type RPCExclusions []interface{}

type Map struct {
	Instructions map[string]RPCInstructions
	Overrides    map[string]RPCOverrides
	Exclusions   map[string]RPCExclusions
}

type File struct {
	ConfigMap      Map // TODO: Can this not be an interface{}?
	ProtofileNames []string
	AllowedOrigin  string
	Port           int64
	Https          bool
}

type Mutation struct {
	ConfigMap Map
}

func ReadFile(filename string) (*File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Could not find config file at %s", filename)
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
		return nil, fmt.Errorf("couldn't turn config into map")
	}

	k := "protofiles"
	p, ok := i[k].([]interface{})
	if !ok {
		return nil, fmt.Errorf("nothing to mock. `protofiles` missing in config file")
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
	https, ok := i[UseHTTPSToken].(bool)
	if !ok {
		https = true
	}

	instructions := parseInstructions(i["instructions"])

	return &File{
		ConfigMap: Map{
			Instructions: instructions,
		},
		ProtofileNames: protofiles,
		AllowedOrigin:  allowedOrigin,
		Port:           port,
		Https:          https,
	}, nil
}

func parseInstructions(i interface{}) map[string]RPCInstructions {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	out := make(map[string]RPCInstructions, len(m))
	for rpc, inst := range m {
		vs, ok := inst.(map[string]interface{})
		if !ok {
			continue
		}
		ins := RPCInstructions{}
		ds, ok := vs["delaySeconds"].(int)
		if ok {
			ins.DelaySecs = ds
		}
		sc, ok := vs["statusCode"].(int)
		if ok {
			ins.StatusCode = sc
		}
		eb, ok := vs["emptyBody"].(bool)
		if ok {
			ins.EmptyBody = EmptyBodyFromBool(eb)
		}
		out[rpc] = ins
	}
	return out
}
