package set

import "github.com/cheekybits/genny/generic"

// Type represents the generic type
type Type generic.Type

// NewTypeSet returns a new set. This is not thread safe.
func NewTypeSet(size ...int) TypeSet {
	if size != nil && len(size) > 0 {
		return make(map[Type]struct{}, size[0])
	}
	return make(map[Type]struct{})
}

// NewTypeSetFromSlice creates a set from a slice
func NewTypeSetFromSlice(sl []Type, size ...int) TypeSet {
	s := NewTypeSet(size...)
	for _, sli := range sl {
		s.Add(sli)
	}
	return s
}

// TypeSet stores the set values
type TypeSet map[Type]struct{}

// Add a value to the set
func (s TypeSet) Add(value Type) {
	s[value] = struct{}{}
}

// Remove a value from the set
func (s TypeSet) Remove(value Type) {
	delete(s, value)
}

// Contains returns true if the value exists in the set, false if it does not
func (s TypeSet) Contains(value Type) bool {
	_, found := s[value]
	return found
}

// ToSlice returns the set as a slice
func (s TypeSet) ToSlice() []Type {
	slice := make([]Type, len(s))
	i := 0
	for k := range s {
		slice[i] = k
		i++
	}
	return slice
}

// Len returns the number of values in the set
func (s TypeSet) Len() int {
	return len(s)
}
