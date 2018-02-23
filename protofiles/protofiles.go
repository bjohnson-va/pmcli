package protofiles

import (
	"github.com/vendasta/mscli/pkg/spec"
	"fmt"
	"os"
	"github.com/emicklei/proto"
)

func GetNames(config spec.MicroserviceConfig) ([]string, error) {
	protoPaths := config.GetSDKProtoPaths()

	var absolutePaths []string
	for _, p := range protoPaths {
		absolutePaths = append(absolutePaths, p)
	}
	return absolutePaths, nil
}

func Read(absoluteFilePath string) (*proto.Proto, error) {
	f := fmt.Sprintf("%s", absoluteFilePath)
	reader, err := os.Open(f)
	defer reader.Close()
	if err != nil {
		err := fmt.Errorf("unable to read proto file: %s\n%s", f, err.Error())
		return nil, err
	}
	parser := proto.NewParser(reader)
	d, err := parser.Parse()
	if err != nil {
		err := fmt.Errorf("unable to parse proto file: %s\n%s", f, err.Error())
		return nil, err
	}
	return d, nil
}
