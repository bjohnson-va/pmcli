package mockserver

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

