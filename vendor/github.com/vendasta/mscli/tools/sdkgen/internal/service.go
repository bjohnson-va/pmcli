package internal

import (
	"strings"
	"fmt"
	"github.com/vendasta/mscli/tools/sdkgen/util"
)

// Service structure of service info
type Service struct {
	Name     string
	FileName string
	RPCS     []*RPC
}

func (s *Service) normalizedServiceName() string {
	return strings.TrimSuffix(s.Name, "Service")
}

// PyFileName filename of a python file
func (s *Service) PyFileName() string {
	fileWithNoExtension := strings.Split(s.FileName, ".")[0]
	pathParts := strings.Split(fileWithNoExtension, "/")
	fileWithNoPath := pathParts[len(pathParts)-1]
	return fmt.Sprintf("%s_pb2", fileWithNoPath)
}

// TsFileName filename of a typescript file
func (s *Service) TsFileName() string {
	return fmt.Sprintf("%s", util.ToKebabCase(s.normalizedServiceName()))
}

// TsName service name of the typescript api service
func (s *Service) TsName() string {
	return s.normalizedServiceName()
}

// UniqueRequestAndResponseClassNames list of all of the request and response names in the service
func (s *Service) UniqueRequestAndResponseClassNames() []string {
	var classNames []string
	for _, rpc := range s.RPCS {
		if !contains(classNames, rpc.RequestClassName()) {
			classNames = append(classNames, rpc.RequestClassName())
		}
		if !rpc.ResponseIsEmpty() && !contains(classNames, rpc.ResponseClassName()){
			classNames = append(classNames, rpc.ResponseClassName())
		}
	}
	return classNames
}

func contains(list []string, value string) bool {
    for _, val := range list {
        if val == value {
            return true
        }
    }
    return false
}