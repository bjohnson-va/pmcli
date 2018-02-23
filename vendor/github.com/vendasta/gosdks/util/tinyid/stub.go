package tinyid

import (
	"errors"

	"golang.org/x/net/context"
)

// Stub fulfills the tinyid.Interface. You may set Values to the list of string you expect to be returned when
// GenerateTinyID is called. These values will be returned in order. If there are no values, or the end of the list has
// been reached an error will be returned.
type Stub struct {
	Values    []string
	nextIndex int
}

// GenerateTinyID returns the next tiny id value from Values. If there are none an error will be returned.
func (t *Stub) GenerateTinyID(ctx context.Context) (string, error) {
	if t.nextIndex < len(t.Values) {
		v := t.Values[t.nextIndex]
		t.nextIndex++
		return v, nil
	}
	return "", errors.New("failed to generate tinyid")
}
