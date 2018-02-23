package adwords

import "testing"

func TestClientMatchesInterface(_ *testing.T) {
	var interf Interface
	interf = &client{}
	_ = interf
}

func TestMockMatchesInterface(_ *testing.T) {
	var interf Interface
	interf = &MockAdwordsClient{}
	_ = interf
}
