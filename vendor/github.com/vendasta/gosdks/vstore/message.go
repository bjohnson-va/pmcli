package vstore

import (
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/proto"
	"github.com/vendasta/gosdks/pb/vstorepb"
	"golang.org/x/net/context"
)

type Message interface {
	Model() Model
	Version() int64
}

type message struct {
	model   Model
	version int64
}

func (m *message) Model() Model {
	return m.model
}

func (m *message) Version() int64 {
	return m.version
}

func FromPubsubMessage(m *pubsub.Message) (Message, error) {
	e := &vstorepb.Entity{}
	err := proto.Unmarshal(m.Data, e)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not proto decode message: %s", err.Error()))
	}

	model, err := StructPBToModel(e.GetNamespace(), e.GetKind(), e.GetValues())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could deserialize entity into vstore model: %s", err.Error()))
	}

	return &message{model: model, version: e.GetVersion()}, nil
}

type MessageHandler func(ctx context.Context, message Message) error
