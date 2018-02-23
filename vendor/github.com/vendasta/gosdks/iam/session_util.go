package iam

import (
	"crypto/x509"
	"encoding/pem"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vendasta/gosdks/config"
)

func CreateTestJWT(email string) (string, error) {
	now := time.Now()
	exp := now.Add(1 * time.Minute)
	jwtclaim := &iamJWTClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "iam.vendasta-local.com",
			Audience:  "https://iam.vendasta-local.com",
			Subject:   email,
			IssuedAt:  now.Unix(),
			ExpiresAt: exp.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtclaim)
	block, _ := pem.Decode(TEST_PRIVATE_KEY)
	private_key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	return token.SignedString(private_key)
}

func MockPublicKey() {
	publicKeys[config.CurEnv()] = TEST_PUBLIC_KEY
	publicKey, _ = getPublicKey()
}

var TEST_PRIVATE_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

const TEST_PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAu4TzTTajW7JRyjsE/+5d
cw3Ctek7MkREkX4qIjuXqGfyTG9SOZEpZK8RuQyRyz8DiNXmCW7hVD/CaJw9xQZt
82vMbPAM2q94ktF3fpxOGyACNwmCFjMp+K5vQtG/+IhBpGVLcfm4uHbQzaRRN6EL
mWup2xepoLsWuzZEKRPORA3zz1E0RgB6IbL6w0+HYrAaTfC0ghaPbm3ByKUKEb/F
Jpyjj5rQMSrtGA47oOIsAuPHvN72xQzspHxOFKLfoX8ucR2cNhMnH42v6byCtLFZ
1rkYYnatsslbn2hLNqeAyW0+fYkRpDnE9o7iDTSGIsayRiCEhMpp3E7garrthtCF
s7F4cZQefO8Or+t6eSqFqckOm4wmpBgAx1MtNSQpViHAzogbJVFkZKlZYUtkLY2U
yHlCbPs/Gpetk3j1umN2bME0t9THYFf5Wj0D0MKXMci1+0b18h1AAYlePmKnm5Pu
OEx2I8zM+AKfj4Nq82+iqfd/iuBGqd5pLW0+9KwlQ/epWcGwlcRGGb6na5k4bPud
7F3fOPujGZ/fS3kGOoG9sURj9CrQe5WhobU13tn3x2S9sc4uhYpP7A3haPal5EmA
XpvLFN7jbqKPZyfD5//bkvUH/CHm4e7dKhUYGnZcOaf33cQanCZaGQJaTW7PcN3q
5oIZKURbbWd4SCHJ9SoqssUCAwEAAQ==
-----END PUBLIC KEY-----`
