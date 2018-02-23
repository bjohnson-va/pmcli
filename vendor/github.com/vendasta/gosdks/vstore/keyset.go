package vstore

import "github.com/vendasta/gosdks/pb/vstorepb"


//NewKeySet builds and returns a new vstore.KeySet
func NewKeySet(namespace string, kind string, keys []string) *KeySet {
	return &KeySet{namespace: namespace, kind: kind, keys: keys}
}

//KeySet contains all of the information necessary to target a specific entity in VStore
type KeySet struct {
	namespace string
	kind      string
	keys      []string
}

//ToKeySetPB serializes a KeySet into protobuf format
func (k *KeySet) ToKeySetPB() *vstorepb.KeySet {
	return &vstorepb.KeySet{
		Namespace: k.namespace,
		Kind: k.kind,
		Keys: k.keys,
	}
}
