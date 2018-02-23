package mockserver

import (
	"strings"
	"math"
	"fmt"
	"github.com/emicklei/proto"
	"github.com/vendasta/gosdks/util"
	"math/rand"
)

type RandomFieldProvider interface {
	newBool() bool
	newInt32() int32
	newFloat32() float32
	newFloat64() float64
	newString() string
	newBytes() []byte
}

type PseudoRandomFieldProvider struct {
}

