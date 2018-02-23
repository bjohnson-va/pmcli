package cloudsqlclient

import (
	"fmt"
	"io/ioutil"
	"os"
)

// CloudSQL environment variable names
const (
	ipName       = "CLOUD_SQL_IP"
	instanceName = "CLOUD_SQL_INSTANCE_NAME"
	username     = "CLOUD_SQL_USERNAME"
	password     = "CLOUD_SQL_PASSWORD"
	certsPath    = "CLOUD_SQL_CERTS_PATH"
)

var certs *Certs

// Certs contains the client key, client cert and server ca
type Certs struct {
	ClientKey  []byte
	ClientCert []byte
	ServerCA   []byte
}

// GetCerts returns the certs
func GetCerts() (*Certs, error) {
	if certs != nil {
		return certs, nil
	}
	certsPath := os.Getenv(certsPath)
	if certsPath == "" {
		return nil, fmt.Errorf("environment variable %s is not set", certsPath)
	}
	clientKey, err := ioutil.ReadFile(certsPath + "/client-key.pem")
	if err != nil {
		return nil, err
	}
	clientCert, err := ioutil.ReadFile(certsPath + "/client-cert.pem")
	if err != nil {
		return nil, err
	}
	serverCA, err := ioutil.ReadFile(certsPath + "/server-ca.pem")
	if err != nil {
		return nil, err
	}
	certs = &Certs{
		ClientKey:  clientKey,
		ClientCert: clientCert,
		ServerCA:   serverCA,
	}
	return certs, nil
}

// IP returns the Cloud SQL ip address
func IP() string {
	return os.Getenv(ipName)
}

// InstanceName returns the Cloud SQL instance name
func InstanceName() string {
	return os.Getenv(instanceName)
}

// Username returns the Cloud SQL username
func Username() string {
	return os.Getenv(username)
}

// Password returns the Cloud SQL password
func Password() string {
	return os.Getenv(password)
}
