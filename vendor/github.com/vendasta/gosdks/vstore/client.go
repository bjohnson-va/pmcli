package vstore

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/vendasta/gosdks/pb/vstorepb"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"github.com/vendasta/gosdks/vstore/lb"
)

const (
	prodAddress     = "vstore-api-production.vendasta-internal.com:443"
	testAddress     = "vstore-api-test.vendasta-internal.com:443"
	demoAddress     = "vstore-api-demo.vendasta-internal.com:443"
	localAddress    = "vstore-api-test.vendasta-internal.com:443" // Namespacing manages separation of developer-specific data
	internalAddress = "127.0.0.1:10000"
)

// Creates a new vstore client
func newClient(e env, useInternalTransport bool, dialOption ...grpc.DialOption) (vstorepb.VStoreClient, vstorepb.VStoreAdminClient, error) {
	dts, err := google.DefaultTokenSource(context.Background(), "https://www.googleapis.com/auth/userinfo.email")
	if err != nil {
		return nil, nil, err
	}

	var address string
	var rootCAs *x509.CertPool
	certificates := []tls.Certificate{}

	var namespace string
	if e == Internal {
		address = internalAddress
		cer, err := tls.X509KeyPair([]byte(LocalCert), []byte(LocalKey))
		if err != nil {
			return nil, nil, err
		}
		rootCAs = x509.NewCertPool()
		rootCAs.AppendCertsFromPEM([]byte(LocalCa))
		certificates = append(certificates, cer)
	} else if e == Local {
		address = localAddress
	} else if e == Demo {
		address = demoAddress
		namespace = "vstore-demo"
	} else if e == Test {
		address = testAddress
		namespace = "vstore-test"
	} else if e == Prod {
		address = prodAddress
		namespace = "vstore-prod"
	}

	config := &tls.Config{
		Certificates: certificates,
		RootCAs:      rootCAs,
	}
	creds := credentials.NewTLS(config)
	dialOption = append(
		dialOption,
		grpc.WithBackoffConfig(grpc.DefaultBackoffConfig),
		grpc.WithPerRPCCredentials(lb.TokenSource{dts}),
	)
	if useInternalTransport {
		ds := &lb.DialSettings{
			Namespace: namespace,
			Labels:    map[string]string{"type": "api"},
		}
		pr, err := lb.NewPoolResolver(ds)
		if err != nil {
			return nil, nil, err
		}
		dialOption = append(dialOption,
			grpc.WithBalancer(grpc.RoundRobin(pr)),
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
		)
	} else {
		dialOption = append(dialOption,
			grpc.WithBalancer(grpc.RoundRobin(NewPoolResolver(3, &DialSettings{Endpoint: address}))),
			grpc.WithTransportCredentials(creds),
		)
	}
	client, err := grpc.Dial(
		address,
		dialOption...,
	)
	if err != nil {
		return nil, nil, err
	}
	return vstorepb.NewVStoreClient(client), vstorepb.NewVStoreAdminClient(client), nil
}
