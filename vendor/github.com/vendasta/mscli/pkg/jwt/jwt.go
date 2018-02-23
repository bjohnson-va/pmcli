package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/vendasta/mscli/pkg/spec"
	"github.com/vendasta/mscli/pkg/utils"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jws"
)

// GoogleIDTokenResponse response structure we get back from google with oauth token
type GoogleIDTokenResponse struct {
	IDToken string `json:"id_token"`
}

// GenerateJWT generate a jwt
func GenerateJWT(specFile spec.MicroserviceFile, scope string, env utils.Environment) error {
	se, err := specFile.Microservice.GetEnv(env)
	if err != nil {
		return err
	}

	jwtConfigBytes := []byte{}
	jwtConfig := se.JwtConfig
	if jwtConfig == nil {
		jwtConfig = &spec.JwtConfig{}
		err = json.Unmarshal([]byte(spec.VendastaLocalJSONKey), jwtConfig)
		if err != nil {
			return fmt.Errorf("failed to unmarshal default credentials: %s", err.Error())
		}
	}
	if jwtConfigBytes, err = json.Marshal(jwtConfig); err != nil {
		return fmt.Errorf("failed to marshal JWTConfig: %s", err.Error())
	}

	conf, err := google.JWTConfigFromJSON(jwtConfigBytes, scope)
	if err != nil {
		return fmt.Errorf("could not parse service account JSON: %v", err)
	}
	rsaKey, err := parseKey(conf.PrivateKey)
	if err != nil {
		return fmt.Errorf("could not get RSA key: %v", err)
	}

	if scope == "" {
		scope = fmt.Sprintf("https://%s", se.Network.GRPCHost)
	}

	iat := time.Now()
	exp := iat.Add(time.Hour)

	jwt := &jws.ClaimSet{
		Iss:   conf.Email,
		Aud:   "https://www.googleapis.com/oauth2/v4/token",
		Scope: scope,
		Iat:   iat.Unix(),
		Exp:   exp.Unix(),
	}
	jwsHeader := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}

	msg, err := jws.Encode(jwsHeader, jwt, rsaKey)
	if err != nil {
		return fmt.Errorf("could not encode JWT: %s", err.Error())
	}
	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v4/token",
		url.Values{
			"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
			"assertion":  {string(msg[:])},
		})
	if err != nil {
		return fmt.Errorf("could not get JWT from google: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		d, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("could not read body: %s", err.Error())
		}
		return fmt.Errorf("Response: %s\n", string(d))
	}
	var idToken GoogleIDTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&idToken)
	if err != nil {
		return err
	}
	fmt.Println(idToken.IDToken)
	return nil
}

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
