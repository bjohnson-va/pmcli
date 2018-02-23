package spec

import (
	"path"
)

//GetProtoImportPaths returns the path to the protos to use as a go import
func (s *MicroserviceConfig) GetProtoImportPaths() []string {
	var ip []string
	for _, p := range s.GetProtoPaths() {
		ip = append(ip, "github.com/vendasta/gosdks/pb/"+path.Dir(p.Path))
	}
	return ip
}

// GetProtoPaths returns a list of proto paths. The value of the deprecated ProtoPath property will be included
func (s *MicroserviceConfig) GetProtoPaths() []*ProtoPath {
	p := s.ProtoPaths
	if s.ProtoPath != "" {
		depPath := &ProtoPath{Path: s.ProtoPath}
		p = append(p, depPath)
	}
	return p
}

// GetSDKProtoPaths returns a list of proto paths that should be used to generate an SDK
func (s *MicroserviceConfig) GetSDKProtoPaths() []string {
	protoPaths := []string{}
	for _, p := range s.GetProtoPaths() {
		if p.ExcludeFromSDK == false {
			protoPaths = append(protoPaths, p.Path)
		}
	}
	return protoPaths
}

//GetPathToServerDir returns the path to the server directory
func (s *MicroserviceConfig) GetPathToServerDir() string {
	return "./server"
}

//GetPathToCodeImplementationDir returns the path to where the microservice code should be placed
func (s *MicroserviceConfig) GetPathToCodeImplementationDir() string {
	if !s.UseInternalPackage {
		return "./pkg"
	}
	return "./internal"
}

//GetCodeImplementationPackageName returns the package where the microservice code should be placed
func (s *MicroserviceConfig) GetCodeImplementationPackageName() string {
	if !s.UseInternalPackage {
		return s.GoPackageName + "/pkg"
	}
	return s.GoPackageName + "/internal"
}

//GetPathToAPIDir returns the path to the api directory
func (s *MicroserviceConfig) GetPathToAPIDir() string {
	return s.GetPathToCodeImplementationDir() + "/api"
}

//GetPathToMainGo returns the path to the main.go file
func (s *MicroserviceConfig) GetPathToMainGo() string {
	return s.GetPathToServerDir() + "/main.go"
}
