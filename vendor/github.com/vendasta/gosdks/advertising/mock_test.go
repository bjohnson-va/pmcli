package advertising

import "testing"

func TestMockMatchesClientInterface(t *testing.T) {
	var client Client
	client = &MockAdvertisingClient{}
	_ = client // shut up golint, I know what I'm doing
}
