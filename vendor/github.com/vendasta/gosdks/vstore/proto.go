package vstore

import (
	"math"
	"sync"

	"github.com/golang/protobuf/proto"
)

// protoCodec is a Codec implementation with protobuf. It is the default codec for gRPC.
type protoCodec struct {
}

type cachedProtoBuffer struct {
	lastMarshaledSize uint32
	proto.Buffer
}

func capToMaxInt32(val int) uint32 {
	if val > math.MaxInt32 {
		return uint32(math.MaxInt32)
	}
	return uint32(val)
}

func (p protoCodec) marshal(v interface{}, cb *cachedProtoBuffer) ([]byte, error) {
	protoMsg := v.(proto.Message)
	newSlice := make([]byte, 0, cb.lastMarshaledSize)

	cb.SetBuf(newSlice)
	cb.Reset()
	if err := cb.Marshal(protoMsg); err != nil {
		return nil, err
	}
	out := cb.Bytes()
	cb.lastMarshaledSize = capToMaxInt32(len(out))
	return out, nil
}

func (p protoCodec) Marshal(v interface{}) ([]byte, error) {
	cb := protoBufferPool.Get().(*cachedProtoBuffer)
	out, err := p.marshal(v, cb)

	// put back buffer and lose the ref to the slice
	cb.SetBuf(nil)
	protoBufferPool.Put(cb)
	return out, err
}

func (p protoCodec) Unmarshal(data []byte, v interface{}) error {
	cb := protoBufferPool.Get().(*cachedProtoBuffer)
	cb.SetBuf(data)
	v.(proto.Message).Reset()
	err := cb.Unmarshal(v.(proto.Message))
	cb.SetBuf(nil)
	protoBufferPool.Put(cb)
	return err
}

func (protoCodec) String() string {
	return "proto"
}

func Marshal(v interface{}) ([]byte, error) {
	return codec.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return codec.Unmarshal(data, v)
}

var (
	protoBufferPool = &sync.Pool{
		New: func() interface{} {
			return &cachedProtoBuffer{
				Buffer:            proto.Buffer{},
				lastMarshaledSize: 16,
			}
		},
	}
)

var (
	codec *protoCodec
)

func init() {
	codec = &protoCodec{}
}
