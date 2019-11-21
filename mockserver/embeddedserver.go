package mockserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

func AddCertsToServer(server *http.Server, port int64, cert string, key string) Server {
	return &embeddedServer{
		server:               server,
		port:                 port,
		webserverCertificate: cert,
		webserverKey:         key,
	}
}

type embeddedServer struct {
	server               *http.Server
	port                 int64
	webserverCertificate string
	webserverKey         string
}

func (srv *embeddedServer) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *embeddedServer) ListenAndServe() error {

	config := &tls.Config{
		MinVersion: tls.VersionTLS10,
	}
	if srv.server.TLSConfig != nil {
		*config = *srv.server.TLSConfig
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}

	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair([]byte(srv.webserverCertificate), []byte(srv.webserverKey))
	if err != nil {
		return fmt.Errorf("failed creating cert pair: %s", err.Error())
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.port))
	if err != nil {
		return fmt.Errorf("failed listening to TCP: %s", err.Error())
	}

	tlsListener := tls.NewListener(conn, config)
	err = srv.server.Serve(tlsListener)
	if err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}
	return nil
}
