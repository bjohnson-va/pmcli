package vstore

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type EnvTestSuite struct {
	suite.Suite
	currentEnvValue string
}

func TestEnvTestSuite(t *testing.T) {
	suite.Run(t, new(EnvTestSuite))
}

func (suite *EnvTestSuite) SetupTest() {
	suite.currentEnvValue = os.Getenv("ENVIRONMENT")
}

func (suite *EnvTestSuite) TearDownTest() {
	//We do this to make sure our tests aren't screwing with environment variables that are important to other test suites
	os.Setenv("ENVIRONMENT", suite.currentEnvValue)
}

func (suite *EnvTestSuite) Test_EnvReturnsLocalValForLocalEnvVar() {
	os.Setenv("ENVIRONMENT", "local")
	e := Env()
	suite.Assert().Equal(Local, *e)
}

func (suite *EnvTestSuite) Test_EnvReturnsInternalValForInternalEnvVar() {
	os.Setenv("ENVIRONMENT", "internal")
	e := Env()
	suite.Assert().Equal(Internal, *e)
}

func (suite *EnvTestSuite) Test_EnvReturnsTestValForTestEnvVar() {
	os.Setenv("ENVIRONMENT", "test")
	e := Env()
	suite.Assert().Equal(Test, *e)
}

func (suite *EnvTestSuite) Test_EnvReturnsDemoValForDemoEnvVar() {
	os.Setenv("ENVIRONMENT", "demo")
	e := Env()
	suite.Assert().Equal(Demo, *e)
}

func (suite *EnvTestSuite) Test_EnvReturnsProdValForProdEnvVar() {
	os.Setenv("ENVIRONMENT", "prod")
	e := Env()
	suite.Assert().Equal(Prod, *e)
}

func (suite *EnvTestSuite) Test_EnvPanicsOnUnknownEnvVar() {
	os.Setenv("ENVIRONMENT", "brent yates")

	// we should be scared if our environment is called brent yates
	suite.Panics(func() {
		Env()
	})
}

func (suite *EnvTestSuite) Test_EnvReturnsLocalValForMissingEnvVar() {
	os.Setenv("ENVIRONMENT", "")
	e := Env()
	suite.Assert().Equal(Local, *e)
}
