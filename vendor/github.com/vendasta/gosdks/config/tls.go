package config

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
)

// GetTLSConfig returns the TLS configuration for this instance based on the environment.
func GetTLSConfig() *tls.Config {
	tlsCert := os.Getenv("TLS_CERT_FILE")
	tlsKey := os.Getenv("TLS_KEY_FILE")

	cer, err := tls.LoadX509KeyPair(tlsCert, tlsKey)
	if err != nil {
		log.Fatalf("Error loading key pair: %v", err)
	}
	var rootCAs *x509.CertPool

	_, isSet := os.LookupEnv("ROOT_CA_FILE")
	if !isSet {
		rootCAs = x509.NewCertPool()
		caFile := os.Getenv("DEV_CA_FILE")
		pemData, err := ioutil.ReadFile(caFile)
		if err != nil {
			log.Fatalf("Error loading root cert: %v", err)
		}
		rootCAs.AppendCertsFromPEM(pemData)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cer},
		ClientAuth:   tls.NoClientCert,
		RootCAs:      rootCAs,
	}
}
