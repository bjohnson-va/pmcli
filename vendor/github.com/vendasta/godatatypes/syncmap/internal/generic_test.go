package syncmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSyncKeyTypeValueTypeMap(t *testing.T) {
	inK := new(KeyType)
	inV := new(ValueType)

	o := map[KeyType]ValueType{}
	o[inK] = inV

	m := ToSyncKeyTypeValueTypeMap(o)
	r, ok := m.Get(inK)
	assert.Equal(t, inV, r)
	assert.True(t, ok)
}

func TestSyncKeyTypeValueTypeMap_Get(t *testing.T) {
	type testCase struct {
		setup       func(m *SyncKeyTypeValueTypeMap)
		expectedVal ValueType
		expectedOK  bool
		description string
	}

	cases := []*testCase{
		{
			setup:       func(m *SyncKeyTypeValueTypeMap) {},
			expectedVal: nil,
			expectedOK:  false,
			description: "Key miss",
		},
		{
			setup: func(m *SyncKeyTypeValueTypeMap) {
				m.Set("k", "v")
			},
			expectedVal: "v",
			expectedOK:  true,
			description: "Key hit",
		},
		{
			setup: func(m *SyncKeyTypeValueTypeMap) {
				m.Set("k", "v")
				m.Set("k", "t")
			},
			expectedVal: "t",
			expectedOK:  true,
			description: "Key hit (overwritten)",
		},
		{
			setup: func(m *SyncKeyTypeValueTypeMap) {
				m.Set("k", "v")
				m.Delete("k")
			},
			expectedVal: nil,
			expectedOK:  false,
			description: "Key miss (after delete)",
		},
	}

	for _, c := range cases {
		m := NewSyncKeyTypeValueTypeMap()
		c.setup(m)
		val, ok := m.Get("k")
		assert.Equal(t, c.expectedVal, val, c.description)
		assert.Equal(t, c.expectedOK, ok, c.description)
	}
}
