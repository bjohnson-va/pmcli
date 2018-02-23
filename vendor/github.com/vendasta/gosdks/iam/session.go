package iam

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/gosdks/util"
)

type iamJWTClaim struct {
	jwt.StandardClaims
}

// ClientKey is a JSON secret we provide to applications that contains all the info necessary to build to a JWT that can
// be exchanged for a session with IAM.
type ClientKey struct {
	Type string `json:"type"`
	Namespace string `json:"namespace"`
	ClientEmail string `json:"client_email"`
	PrivateKey string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
}

// clientJWTClaim is a JWT that is signed by the client and exchanged for an IAM session (an iamJWTClaim)
type clientJWTClaim struct {
	KeyID            string `json:"kid"`
	SubjectType      string `json:"vendasta.com/subject-type"`
	SubjectNamespace string `json:"vendasta.com/subject-namespace"`
	jwt.StandardClaims
}

// ValidateJWT validates the JWT is valid and not expired
func ValidateJWT(session string) (*iamJWTClaim, error) {
	token, err := jwt.ParseWithClaims(session, &iamJWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if publicKey == nil {
			return nil, errors.New("public key not loaded, not able to parse session")
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, util.Error(util.PermissionDenied, "Invalid session passed.")
	}
	claims, ok := token.Claims.(*iamJWTClaim)
	if !ok || !token.Valid {
		return nil, util.Error(util.PermissionDenied, "Invalid session passed.")
	}
	return claims, nil
}

func getPublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeys[config.CurEnv()]))
	if block == nil {
		return nil, errors.New("no key found")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pk, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid public key.")
	}
	return pk, nil
}

var publicKey *rsa.PublicKey

func init() {
	var err error
	publicKey, err = getPublicKey()
	if err != nil {
		logging.Errorf(context.Background(), "Error loading public key, not going to be able to validate JWTs.")
	}
}

const (
	localPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzD4soeAeEus2NHZpTgrC
OCui/AEjShVZEzFeglGiV66CRx6HE9o3BDyi1Hs0jgguABJhjICFfeMYdr0NOIpZ
/1+0lZtvAfQ8wNjujRYsZ/+h7yuGucEL7issaBZcr7c6g7gKoroQ6pYRDxWCVBOg
Vbe1a3uIDOfT8qCC0qPkSiKTjjQ6Zm2qABxvX/NNZXeGRRzrtSDn8xXhvv8fwrTW
vFYH019ANeiSFXAjZXvS/aEaavpA2th16Iy07URGp8Jgr9+7/I/medYSb0tFOOSo
C7EwnxaqcoGy4ETGHLYu97d45dtwy7ECb7Nso5HbapKGHfG1NULRRpvplZOS0koP
6QIDAQAB
-----END PUBLIC KEY-----`

	testPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6THvHlo30RvGEmsWDB09
hq/z8n88yMVmpt3o7olEaakfFiW8Pl3MDufRryXX/vBELZPWwmjGD0fhqlLewIre
J02DnzdUYk9d+T3oKstVKowI4pW9zDe9boYF0fXxrhTYOcnEghNpTqdUddoL4rj6
Pg/8Ru//fW1BPFQzFNfO/BqeFc47zPGZlTXVKXsUEQrtGh4Qg5Kj/xuN+ymK0meK
lcN9AmZzy9JoS+TrcZit0JxKqaj//s25HbfcNXC2grAlOLM6yNuA1X3dorE9vxTz
Hx5u8VVoELiXZSMuzWzIxUESNabIJGaAWsU3z+OR2hhPmtDWX+3wTr7cmGwMsRsR
iwIDAQAB
-----END PUBLIC KEY-----`
	demoPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzdk9vv5ckgEDMGNSH49s
WH0ssjrsuvAj9kWg8MajZsc1H4+G+DxMPKKkQPpLBhPuSyb8QAyFIrF8BBd9bmIK
hAb0PLDkpXgN6kJn9zJPtFFzMZqb7d5p897aZnBH/kEsSdxViydrJzKXVAD8G84C
RdQt3OlfeQsp8+0c9k+mjSG4lFKY2qClwNIjIH4BXrXDYJ50BqQQ6Dcm9yh+fk6e
dubfodDNSAPdlh/vhXrVrwrNf6FaQIlWVJqDiqd5damkFpT4fs2lIgTRLMnLjm6f
JtPeU8b+CSjm3oavIQxEbGQxl0s5laGQzb6Eosve2+RWlsG1TFU+pbIp8Rr3W1LF
sQIDAQAB
-----END PUBLIC KEY-----`
	prodPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr4n11iE63EIc4pTv16/L
U5dN53ybcO5Vb3hUT4DQEcVqsKoPvpZD4nusiGLLvp5GE4dD2MXoSwJ02US9EDlZ
aC8LcwzOhHXlRA+RduAfuZCXqP4Ov6naL17YrNq5uO4bL75HjxmzDEBCm+R9n5/D
uOaRdUQDQognB9dDjYWdwcZ3yrnLNst2OwAq/YKg+BKxquJQsor9RXYCeYWMjLCU
Z9ljzDw1qgCFf+6RpuVj1O0l1QIRZJKJoB5Bbhhz+mVrvTXkjdImPNO209xNXiO5
RdQP7ZhlEQl5rX2jreKNEWS+6EK6MGG5Ra5TS/sPsMTxducBB3lhai6Kd90b2dPU
qQIDAQAB
-----END PUBLIC KEY-----`
)

var (
	publicKeys = map[config.Env]string{
		config.Local: localPublicKey,
		config.Test:  testPublicKey,
		config.Demo:  demoPublicKey,
		config.Prod:  prodPublicKey,
	}
)
