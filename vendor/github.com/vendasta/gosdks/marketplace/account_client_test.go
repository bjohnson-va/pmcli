package marketplace

import (
	"github.com/stretchr/testify/suite"
)

type BaseAccountClientTest struct {
	suite.Suite
	client accountClient
}

func (b BaseAccountClientTest) SetupTest() {
	b.client = accountClient{}
}
