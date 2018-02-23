package vstore

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NewKeySetReturnsSerializableKeySetWithExpectedProperties(t *testing.T) {
	k := NewKeySet("unit-test", "TestEntity", []string{"id1", "id2"})
	r := k.ToKeySetPB()
	assert.Equal(t, r.Namespace, "unit-test")
	assert.Equal(t, r.Kind, "TestEntity")
	assert.Equal(t, r.Keys, []string{"id1", "id2"})
}
