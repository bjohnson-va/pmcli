package vax

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
	"golang.org/x/oauth2/jwt"
)

func parseKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKSC1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return parsed, nil
}

// Token holds the id token returned from google.
type Token struct {
	IDToken string `json:"id_token"`
}

type tokenError struct {
	errorMessage string
}

func (t *tokenError) Error() string {
	return t.errorMessage
}

// TokenError returns an error
func TokenError(f string, i ...interface{}) error {
	return &tokenError{
		errorMessage: fmt.Sprintf(f, i...),
	}
}

// NewTokenSource returns an implementation of credentials.PerRPCCredentials that can identity with Vendastas APIs backed by cloud endpoints.
func NewTokenSource(scope string) (*TokenSource, error) {
	sa, err := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		return nil, TokenError("Could not read service account file: %s", err.Error())
	}
	conf, err := google.JWTConfigFromJSON(sa)
	if err != nil {
		return nil, TokenError("Could not parse service account JSON: %s", err.Error())
	}
	rsaKey, err := parseKey(conf.PrivateKey)
	if err != nil {
		return nil, TokenError("Could not get RSA key: %s", err.Error())
	}
	return &TokenSource{
		privateKey:  rsaKey,
		config:      conf,
		scope:       scope,
		TokenSource: conf.TokenSource(context.Background()),
	}, nil
}

// TokenSource supplies PerRPCCredentials from an oauth2.TokenSource.
type TokenSource struct {
	privateKey *rsa.PrivateKey
	config     *jwt.Config
	oauth2.TokenSource
	scope       string
	accessToken string
	expires     time.Time
	sync.Mutex
}

// GetRequestMetadata gets the request metadata as a map from a TokenSource.
func (ts *TokenSource) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if ts.accessToken == "" || ts.expires.Before(time.Now().UTC()) {
		ts.Lock()
		defer ts.Unlock()
		token, err := ts.getToken()
		if err != nil {
			return nil, err
		}
		ts.accessToken = token.IDToken
		ts.expires = time.Now().UTC().Add(time.Minute * 30)
	}

	return map[string]string{
		"authorization": "Bearer" + " " + ts.accessToken,
		"x-api-key":     "AIzaSyDAfizKwYN_c4YeqaZIO-T0RWcgOzA_d2k",
	}, nil
}

func (ts *TokenSource) getToken() (*Token, error) {
	assertion, err := ts.getAssertion()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://www.googleapis.com/oauth2/v4/token", strings.NewReader("grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion="+assertion))
	if err != nil {
		return nil, TokenError("Error creating token request. %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, TokenError("Error requesting token from Google. %s", err.Error())
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, TokenError("Error reading response body. %s", err.Error())
	}
	t := Token{}
	json.Unmarshal(b, &t)
	return &t, nil
}

func (ts *TokenSource) getAssertion() (string, error) {
	iat := time.Now()
	exp := iat.Add(time.Minute)

	jwt := &jws.ClaimSet{
		Iss:   ts.config.Email,
		Aud:   "https://www.googleapis.com/oauth2/v4/token",
		Iat:   iat.Unix(),
		Exp:   exp.Unix(),
		Scope: ts.scope,
	}
	jwsHeader := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}
	return jws.Encode(jwsHeader, jwt, ts.privateKey)
}

// RequireTransportSecurity indicates whether the credentails requires transport security.
func (ts *TokenSource) RequireTransportSecurity() bool {
	return true
}
