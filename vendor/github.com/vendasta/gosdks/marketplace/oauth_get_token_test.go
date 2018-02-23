package marketplace

import (
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/config"

	"testing"
	"time"
)

func TestOAuthClient_CreateJWTAssertion(t *testing.T) {
	suite.Run(t, new(CreateJWTAssertionTestSuite))
}

type CreateJWTAssertionTestSuite struct {
	suite.Suite
	client OAuthClientInterface
}

func (s *CreateJWTAssertionTestSuite) SetupTest() {
	s.client = NewOAuthClient(config.Local, "")
}

const minimalValidRsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MQ==
-----END RSA PRIVATE KEY-----`

func (s *CreateJWTAssertionTestSuite) TestShouldCreateAnAssertionStringWhenGivenValidInputs() {
	assertion, _ := s.client.CreateJWTAssertion("MP-123", []byte(minimalValidRsaPrivateKey),
		time.Now(), time.Now().Add(time.Hour))
	s.NotNil(assertion)
}

func (s *CreateJWTAssertionTestSuite) TestShouldErrorOnMissingAppID() {
	_, err := s.client.CreateJWTAssertion("", []byte(minimalValidRsaPrivateKey),
		time.Now(), time.Now().Add(time.Hour))
	s.EqualError(err, "appID is required")
}
