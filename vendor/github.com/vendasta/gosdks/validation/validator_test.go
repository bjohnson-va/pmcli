package validation

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/util"
)

func PartnerIDNotEmpty(partnerID string) Rule {
	return StringNotEmpty(
		partnerID,
		util.InvalidArgument,
		"partnerID cannot be empty",
	)
}

func DescriptionRule(value string) Rule {
	return NewValidator().
		Rule(StringNotEmpty(value, util.InvalidArgument, "description is required")).
		Rule(StringMaxLength(value, 30, util.InvalidArgument, "Description must be less than 200 characters"))
}

type ValidatorValidateTestSuite struct {
	suite.Suite
}

func TestValidatorValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorValidateTestSuite))
}

func (suite *ValidatorValidateTestSuite) TestReturnsNilForSingleRuleThatPasses() {
	suite.Assert().Nil(NewValidator().Rule(PartnerIDNotEmpty("PID")).Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsErrorForSingleRuleThatFails() {
	suite.Assert().NotNil(NewValidator().Rule(PartnerIDNotEmpty("")).Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsNilForMultipleRulesThatPass() {
	suite.Assert().Nil(NewValidator().
		Rule(PartnerIDNotEmpty("PID")).
		Rule(PartnerIDNotEmpty("PID2")).
		Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsErrorForMultipleRulesWhenOneFails() {
	suite.Assert().NotNil(NewValidator().
		Rule(PartnerIDNotEmpty("PID")).
		Rule(PartnerIDNotEmpty("")).
		Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsErrorForMultipleRulesWhenAll() {
	suite.Assert().NotNil(NewValidator().
		Rule(PartnerIDNotEmpty("")).
		Rule(PartnerIDNotEmpty("")).
		Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsNilForRulesWithCompositionWhenAllSubRulesPass() {
	suite.Assert().Nil(NewValidator().
		Rule(DescriptionRule("This is the description")).
		Validate())
}

func (suite *ValidatorValidateTestSuite) TestReturnsErrorForRulesWithCompositionWhenASubRulesFails() {
	suite.Assert().NotNil(NewValidator().
		Rule(DescriptionRule("This is the description which is too long")).
		Validate())
}

type ValidatorValidateAndJoinErrorsTestSuite struct {
	suite.Suite
}

func TestValidatorValidateAndJoinErrorsTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorValidateAndJoinErrorsTestSuite))
}

func (suite *ValidatorValidateAndJoinErrorsTestSuite) TestReturnsSingleFailureMessage() {
	err := NewValidator().
		Rule(DescriptionRule("")).
		Rule(DescriptionRule("This description is tooooooooooo long")).
		Rule(PartnerIDNotEmpty("")).
		ValidateAndJoinErrors()

	suite.Assert().Equal("description is required\nDescription must be less than 200 characters\npartnerID cannot be empty", err.Error())
}

func (suite *ValidatorValidateAndJoinErrorsTestSuite) TestReturnsNilForValidInput() {
	err := NewValidator().
		Rule(DescriptionRule("This is a reasonable desc")).
		Rule(PartnerIDNotEmpty("PID")).
		ValidateAndJoinErrors()

	suite.Assert().Nil(err)
}

func (suite *ValidatorValidateAndJoinErrorsTestSuite) TestReturnsMultipleFailureMessages() {
	err := NewValidator().
		Rule(DescriptionRule("")).
		Rule(DescriptionRule("This description is tooooooooooo long")).
		Rule(PartnerIDNotEmpty("")).
		ValidateAndJoinErrors()

	suite.Assert().Equal("description is required\nDescription must be less than 200 characters\npartnerID cannot be empty", err.Error())
}