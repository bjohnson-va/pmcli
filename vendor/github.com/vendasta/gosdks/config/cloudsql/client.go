package cloudsqlclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/vendasta/gosdks/logging"
)

// Initialize create a new sql DB connection
func Initialize(projectID, instanceName, dbName, address, username, password string, clientCert, clientKey, serverCertificateAuthority []byte) (*sql.DB, error) {
	logging.Debugf(context.Background(), "SQL Connection Initialization: Initializing connection")
	//Create a custom TLS Config
	rootCertPool := x509.NewCertPool()
	pem := serverCertificateAuthority
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		logging.Errorf(context.Background(), "SQL Connection Initialization error: Failed to append certs from PEM: `%s`", pem)
		return nil, errors.New("failed to append PEM")
	}
	newClientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		logging.Errorf(context.Background(), "SQL Connection Initialization error: Failed to load cert and key ")
		return nil, err
	}

	newClientCert = append(newClientCert, certs)
	mysql.RegisterTLSConfig(dbName, &tls.Config{
		RootCAs:      rootCertPool,
		Certificates: newClientCert,
		ServerName:   fmt.Sprintf("%s:%s", projectID, instanceName),
	})

	tcpSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/?tls=%s", username, password, address, dbName)
	sqlConnectedClient, err := sql.Open("mysql", tcpSource)
	if err != nil {
		logging.Errorf(context.Background(), "SQL Connection Initialization error: Failed to open SQL connection")
		return nil, err
	}
	err = sqlConnectedClient.Ping()
	if err != nil {
		logging.Errorf(context.Background(), "SQL Connection Initialization error: Failed to ping SQL server %s", address)
		return nil, err
	}
	return sqlConnectedClient, nil
}
