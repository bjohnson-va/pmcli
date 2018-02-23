package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericSet_Add(t *testing.T) {
	cases := []struct {
		name   string
		input  []Type
		output []Type
	}{
		{
			name:   "Adding nothing returns empty slice",
			input:  []Type{},
			output: []Type{},
		},
		{
			name:   "Adding value returns a slice with the value in it",
			input:  []Type{"AG-123"},
			output: []Type{"AG-123"},
		},
		{
			name:   "Only unique values will be returned",
			input:  []Type{"AG-123", "AG-123", "AG-456"},
			output: []Type{"AG-123", "AG-456"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			set := NewTypeSet()
			for _, v := range c.input {
				set.Add(v)
			}
			assert.Equal(t, c.output, set.ToSlice())
		})
	}
}

func TestGenericSet_Remove(t *testing.T) {
	cases := []struct {
		name    string
		remove  Type
		initial []Type
		output  []Type
	}{
		{
			name:    "Removing a value from an empty set should return an empty set",
			remove:  "AG-123",
			initial: []Type{},
			output:  []Type{},
		},
		{
			name:    "Value should be removed from set",
			remove:  "AG-123",
			initial: []Type{"AG-456", "AG-123"},
			output:  []Type{"AG-456"},
		},
		{
			name:    "Removing only value should return an empty set",
			remove:  "AG-123",
			initial: []Type{"AG-123"},
			output:  []Type{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			set := NewTypeSet()
			for _, v := range c.initial {
				set.Add(v)
			}
			set.Remove(c.remove)
			assert.Equal(t, c.output, set.ToSlice())
		})
	}
}

func TestGenericSet_Contains(t *testing.T) {
	cases := []struct {
		name   string
		value  Type
		set    []Type
		output bool
	}{
		{
			name:   "Returns true if value is in the set",
			value:  "AG-123",
			set:    []Type{"AG-123", "AG-456", "AG-789"},
			output: true,
		},
		{
			name:   "Returns false if value is not in the set",
			value:  "AG-1",
			set:    []Type{"AG-123", "AG-456", "AG-789"},
			output: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			set := NewTypeSet()
			for _, v := range c.set {
				set.Add(v)
			}
			assert.Equal(t, c.output, set.Contains(c.value))
		})
	}
}

func TestGenericSet_Len(t *testing.T) {
	cases := []struct {
		name   string
		input  []Type
		output int
	}{
		{
			name:   "Returns zero if there is nothing in the set",
			input:  []Type{},
			output: 0,
		},
		{
			name:   "Returns the length of the set",
			input:  []Type{"AG-123", "AG-456", "AG-789"},
			output: 3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			set := NewTypeSet()
			for _, v := range c.input {
				set.Add(v)
			}
			assert.Equal(t, c.output, set.Len())
		})
	}
}

func TestGenericSet_FromSlice_ToSlice(t *testing.T) {
	cases := []struct {
		name     string
		slice    []Type
		outSlice []Type
	}{
		{
			name:     "Empty",
			slice:    []Type{},
			outSlice: []Type{},
		},
		{
			name:     "Single item",
			slice:    []Type{"AG-123"},
			outSlice: []Type{"AG-123"},
		},
		{
			name:     "Many items",
			slice:    []Type{"AG-123", "AG-456", "AG-789"},
			outSlice: []Type{"AG-123", "AG-456", "AG-789"},
		},
		{
			name:     "Many items with duplicates",
			slice:    []Type{"AG-123", "AG-456", "AG-789", "AG-123", "AG-456", "AG-789"},
			outSlice: []Type{"AG-123", "AG-456", "AG-789"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := NewTypeSetFromSlice(c.slice)
			assert.Equal(t, len(c.outSlice), s.Len())
			for _, i := range c.outSlice {
				assert.True(t, s.Contains(i))
			}

			outSlice := s.ToSlice()
			assert.Equal(t, len(c.outSlice), len(outSlice))
			for _, i := range c.outSlice {
				assert.Contains(t, outSlice, i)
			}
		})
	}
}
