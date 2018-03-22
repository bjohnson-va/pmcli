package random

import (
	"math/rand"
	"github.com/emicklei/proto"
	"github.com/bjohnson-va/pmcli/parse"
	"fmt"
	"hash/fnv"
	"time"
)

type FieldProvider interface {
	NewBool(breadcrumb string) bool
	NewInt32(breadcrumb string) int32
	NewFloat32(breadcrumb string) float32
	NewFloat64(breadcrumb string) float64
	NewString(breadcrumb string) string
	NewBytes(breadcrumb string) []byte
	NewTimestamp(breadcrumb string) string
	NewEnumValue(breadcrumb string, enum proto.Enum) string
}

type seeder interface {
	Seed(breadcrumb string)
}

type breadcrumbSeeder struct {
	seeder
}

func (s breadcrumbSeeder) Seed(breadcrumb string) {
	rand.Seed(hash(breadcrumb))
}

func BreadcrumbBasedFieldProvider() FieldProvider {
	return unseededRandomFieldProvider{
		seeder: breadcrumbSeeder{},
	}
}

type timeBasedSeeder struct {
	seeder
}

func (s timeBasedSeeder) Seed(breadcrumb string) {
	rand.Seed(int64(time.Now().Unix()) + hash(breadcrumb))
}

func TimeBasedFieldProvider() FieldProvider {
	return unseededRandomFieldProvider{
		seeder: timeBasedSeeder{},
	}
}

type unseededRandomFieldProvider struct {
	seeder
}

func (p unseededRandomFieldProvider) NewEnumValue(breadcrumb string, enum proto.Enum) string {
	p.Seed(breadcrumb)
	possibleEnumValues := parse.EnumFields(enum.Elements)
	r := rand.Intn(len(possibleEnumValues))
	return fmt.Sprintf("%s", possibleEnumValues[r].Name)
}

func (p unseededRandomFieldProvider) NewTimestamp(breadcrumb string) string {
	p.Seed(breadcrumb)
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02T15:04:05.999999999Z")
}

func (p unseededRandomFieldProvider) NewInt32(breadcrumb string) int32 {
	p.Seed(breadcrumb)
	return rand.Int31()
}

func (p unseededRandomFieldProvider) NewFloat32(breadcrumb string) float32 {
	p.Seed(breadcrumb)
	return float32(p.NewFloat64(breadcrumb))
}

func (p unseededRandomFieldProvider) NewFloat64(breadcrumb string) float64 {
	p.Seed(breadcrumb)
	stdDev := 1e12
	return rand.NormFloat64() * stdDev // TODO: How to provide useful values?
}

func (p unseededRandomFieldProvider) NewString(breadcrumb string) string {
	p.Seed(breadcrumb)
	ipsumLength := len(loremIpsumString)
	startIndex := rand.Intn(ipsumLength - 1)
	nWords := rand.Intn(ipsumLength - 1)
	endIndex := startIndex + nWords
	if endIndex >= ipsumLength {
		start := loremIpsumString[startIndex : ipsumLength-1]
		return start + loremIpsumString[0:endIndex - ipsumLength]
	}
	return loremIpsumString[startIndex: endIndex]
}

func (p unseededRandomFieldProvider) NewBytes(breadcrumb string) []byte {
	p.Seed(breadcrumb)
	return []byte(p.NewString(breadcrumb))
}

func (p unseededRandomFieldProvider) NewBool(breadcrumb string) bool {
	p.Seed(breadcrumb)
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

func hash(s string) int64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int64(h.Sum32())
}
