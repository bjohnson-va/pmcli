package iam

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSessionTestSuite(t *testing.T) {
	suite.Run(t, new(sessionTestSuite))
}

type sessionTestSuite struct {
	suite.Suite
}

func (suite *sessionTestSuite) SetupSuite() {
	MockPublicKey()
}

func (suite *sessionTestSuite) Test_JWTMustBeValid() {
	token, err := CreateTestJWT("hhill@koth.com")
	suite.Assert().Nil(err)

	claim, err := ValidateJWT(token)

	suite.Assert().NotNil(claim)
	suite.Assert().Nil(err)
}
