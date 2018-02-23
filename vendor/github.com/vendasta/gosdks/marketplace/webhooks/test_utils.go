package marketplacewebhooks

import (
	"encoding/json"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	invalidPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCqGKukO1De7zhZj6+H0qtjTkVxwTCpvKe4eCZ0FPqri0cb2JZfXJ/DgYSF6vUp
wmJG8wVQZKjeGcjDOL5UlsuusFncCzWBQ7RKNUSesmQRMSGkVb1/3j+skZ6UtW+5u09lHNsj6tQ5
1s1SPrCBkedbNf0Tp0GbMJDyR4e9T04ZZwIDAQABAoGAFijko56+qGyN8M0RVyaRAXz++xTqHBLh
3tx4VgMtrQ+WEgCjhoTwo23KMBAuJGSYnRmoBZM3lMfTKevIkAidPExvYCdm5dYq3XToLkkLv5L2
pIIVOFMDG+KESnAFV7l2c+cnzRMW0+b6f8mR1CJzZuxVLL6Q02fvLi55/mbSYxECQQDeAw6fiIQX
GukBI4eMZZt4nscy2o12KyYner3VpoeE+Np2q+Z3pvAMd/aNzQ/W9WaI+NRfcxUJrmfPwIGm63il
AkEAxCL5HQb2bQr4ByorcMWm/hEP2MZzROV73yF41hPsRC9m66KrheO9HPTJuo3/9s5p+sqGxOlF
L0NDt4SkosjgGwJAFklyR1uZ/wPJjj611cdBcztlPdqoxssQGnh85BzCj/u3WqBpE2vjvyyvyI5k
X6zk7S0ljKtt2jny2+00VsBerQJBAJGC1Mg5Oydo5NwD6BiROrPxGo2bpTbu/fhrT8ebHkTz2epl
U9VQQSQzY1oZMVX8i1m5WUTLPz2yLJIBQVdXqhMCQBGoiuSoSjafUhV7i1cEGpb88h5NBYZzWXGZ
37sJ5QsW+sJyoNde3xH8vdXhzU7eT82D6X/scw9RZz+/6rCJ4p0=
-----END RSA PRIVATE KEY-----`
	localPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAs79gLWQH/AGPSrkGuiNz3Ne9EGTUM1MtpSDb4sjDt58nKcDy
hi1KnOs+iHBGxu1e8Cz0byuOgfWCzRW40Qu3VO22oAb4KQodi6NmqcJjIrZva43L
tK5KoLe9qIsc5o4xZsjoy5PUnBZIYjwYTLlQ+PpMHqAvtzhAblj16qB6VER8cdN6
qL9gvzZ0DS5LP34f54kwE5zEYilylee3ZXA3KEX9O1zt2HfwHlWVqBDVcOC9DbW4
RM3Kvdn8IkH9HFoL7QvKniJRBYIDExCQ+wwCOmmWj0HBLAXwqehjD8WUyZTCJLsK
3iNehkDNR4aBpWKJMPDmGp/Q7p2yjVJQ4R6xjwIDAQABAoIBAQCJCSId+OFNczL3
WNCUIGs4txFg6j4NUyC9f6uj3kO4PhHxdzRuryDvOTnXcyHCjz2OBffkuWRQGAi8
qg7Y8QG4MfLFVgrj3t5fkxe13IkP8d5INTZUbXxfwzVIxaYm78kh/5HbgrzYsvRE
UGSGOdJNvyVBRAacvLoCZrg1RFt+2d9oabXWMLKIPCJtlrz5tSZK/KuVGNezpZ4n
CThL+Bi2gQWIIyrAcFLsRhieF93LmtHc8qrHISPJXhjWkHrpord5PMurRVADYjCi
FQNZXiqhH2X2dL3ryvXfTTXV6Ws+gLkdIVvUErH5mVIOCTo6XQ8AMRt1O2iIfqn3
487053JhAoGBAOncFacDp4F5GHYfHouL9mCgrtRkd1e9D5o7bh/Ag0voZFD0guRP
erkUnPfYJzcrjAjWgQEsquymiVvEBCTZI2v/n7jUOA2mqkfhC5S70bWSJ1YXbN+V
2YJqJu3J9J7O9Om2YGAjKuj4hfctnGcPUl1SLIO84S8x4HKgqKGm/bolAoGBAMTD
ztiqp6OhHuaJXWPsaa5t69j1pV57BLnvk9l84G8j9XNBIU2e5xq7TZFZskMx4NGw
NElGMDcWQpODuECcazAmNjRiL6Q+4U1c0xgBU2fRAzGzzkzqN9E05e7U5Nr3oeBa
YS3o8qYQLrxncyajomT7j8vm9n3Peb9zQ2asBryjAoGADRtJl1+sDqSTy2L5GkcQ
2t+GjN7IlYFCE5nLFTp74VA39xWNFHCj5yWRmo5hvETTh1dGSaOQPM0km6EDvnDw
/EeKUxiJv7IcjwS82mwgV7XGQJv1fR5v31BFuEP+Urw0m7id9CVkbUQode73egsw
yYwmdIesczGva503dj/eroUCgYBO3mGhzT2mw+Vb6MGLVO7DYQIBE+12In8xfEta
CMj5I4O17+xGbaQavesIkNh4QJzpz+QVXtHissR61nhbRZddVGOzuYt5HU83sDZY
z+c7nLbYDEdr4xH+hGKJycH11v6GrBN50H66e5AQKrMoULy+F2t7ApfYs/oL+EqS
8+87WwKBgQCIA966oTwca4189kZ7GasJplsy5i4FJxn5NaxqcxaOBbkxnk3QZ4/6
LQlGLYbE5wO0aBCNaCEHGevilMhYfs2rEC/mipl9BwbtvhdZiRP84AearPK2VCsY
GRG4ZVp9HzF+s7Bv6zVkLvy1DHUvv+NxFOgNgDfNGqytnvb6x5XFDA==
-----END RSA PRIVATE KEY-----`
)

type bodyStub struct {
	Reader io.Reader
}

func (b bodyStub) Read(p []byte) (int, error) {
	return b.Reader.Read(p)
}

func (bodyStub) Close() error {
	return nil
}

// signJWT signs the given payload (a JSON bytestring) for the current time against the given private key
func signJWT(privateKey []byte, webhookPayload []byte, overrides map[string]interface{}) string {
	var claims = map[string]interface{}{}

	json.Unmarshal(webhookPayload, &claims)

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	for k, v := range overrides {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims(claims))
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		panic(err.Error())
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err.Error())
	}

	return tokenString
}
