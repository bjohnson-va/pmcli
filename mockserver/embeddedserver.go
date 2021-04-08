package mockserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/bjohnson-va/pmcli/config"
)

func AddCertsToServer(server *InsecureServer, port int64, cert string, key string) Server {
	return &embeddedServer{
		server:               server,
		port:                 port,
		webserverCertificate: cert,
		webserverKey:         key,
	}
}

type embeddedServer struct {
	server               *InsecureServer
	port                 int64
	webserverCertificate string
	webserverKey         string
}

func (srv *embeddedServer) SetAssists(a config.AssistEnum) {
	srv.server.SetAssists(a)
}

func (srv *embeddedServer) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *embeddedServer) ListenAndServe(ctx context.Context) error {

	cfg := &tls.Config{
		MinVersion: tls.VersionTLS10,
	}
	if srv.server.GetTLSConfig() != nil {
		*cfg = *srv.server.GetTLSConfig()
	}
	if cfg.NextProtos == nil {
		cfg.NextProtos = []string{"http/1.1"}
	}

	var err error
	cfg.Certificates = make([]tls.Certificate, 1)
	cfg.Certificates[0], err = tls.X509KeyPair([]byte(srv.webserverCertificate), []byte(srv.webserverKey))
	if err != nil {
		return fmt.Errorf("failed creating cert pair: %s", err.Error())
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.port))
	if err != nil {
		return fmt.Errorf("failed listening to TCP: %s", err.Error())
	}

	tlsListener := tls.NewListener(conn, cfg)
	err = srv.server.Serve(ctx, tlsListener)
	if err != nil {
		return fmt.Errorf("failed to serve: %s", err.Error())
	}
	return nil
}
