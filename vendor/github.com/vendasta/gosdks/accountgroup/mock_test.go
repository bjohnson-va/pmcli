package accountgroup

import "testing"

func TestAccountGroupMockMatchesInterface(t *testing.T) {
	var client Interface
	client = &MockAccountGroupClient{}
	_ = client // To shut up the compiler's unused var error
}
