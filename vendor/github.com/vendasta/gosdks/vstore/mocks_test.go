package vstore

import "testing"

func TestVStoreMockMatchesInterface(t *testing.T) {
	var client Interface
	client = &VStoreMock{}
	_ = client // To shut up the compiler's unused var error
}
