package random

import (
	"github.com/emicklei/proto"
)

type MockFieldProvider struct {
}

func (MockFieldProvider) NewBool(breadcrumb string) bool {
	return true
}

func (MockFieldProvider) NewInt32(breadcrumb string) int32 {
	return int32(32)
}

func (MockFieldProvider) NewFloat32(breadcrumb string) float32 {
	return float32(32.32)
}

func (MockFieldProvider) NewFloat64(breadcrumb string) float64 {
	return float64(64.64)
}

func (MockFieldProvider) NewString(breadcrumb string) string {
	return "string"
}

func (MockFieldProvider) NewBytes(breadcrumb string) []byte {
	return []byte{'b','y','t','e'}
}

func (MockFieldProvider) NewTimestamp(breadcrumb string) string {
	return "2006-01-02T15:04:05.999999999Z"
}

func (MockFieldProvider) NewEnumValue(breadcrumb string, enum proto.Enum) string {
	return "some_enum"
}


