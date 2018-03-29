package config

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/bjohnson-va/pmcli/inputs"
)


func GetInputsForRPC(s proto.Service, r proto.RPC, config map[string]interface{}) (inputs.Provider, error) {
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
	return inputs.New(r.Name, o, i, e), nil
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
		return nil, fmt.Errorf("couldn't turn config into map: %s")
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

	return &File{
		ConfigMap:      i,
		ProtofileNames: protofiles,
		AllowedOrigin:  allowedOrigin,
		Port:           port,
	}, nil
}
