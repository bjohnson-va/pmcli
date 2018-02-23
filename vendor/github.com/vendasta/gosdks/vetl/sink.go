package vetl

import "github.com/vendasta/gosdks/pb/vetl/v1"

type sinkOption struct {
	sink *vetl_v1.DataSink
}

// Sink defines a destination for the output of a vETL transform.
type Sink func(r *sinkOption)

type vstoreSinkOption struct {
	secondaryIndexes []*vetl_v1.VStoreSink_SecondaryIndex
}

// VStoreSinkOption is an option on a VStore Sink.
type VStoreSinkOption func(o *vstoreSinkOption)

// CloudSQLIndex specifies a Cloud SQL secondary index on a VStore Sink.
func CloudSQLIndex(name, instanceIP, username, password string, clientKey, clientCert, serverCertificateAuthority []byte, projectID string, instanceName string) VStoreSinkOption {
	return func(o *vstoreSinkOption) {
		si := &vetl_v1.VStoreSink_SecondaryIndex{
			Name: name,
			Index: &vetl_v1.VStoreSink_SecondaryIndex_CloudSqlConfig{
				CloudSqlConfig: &vetl_v1.VStoreSink_SecondaryIndex_CloudSQLConfig{
					IndexName: name,
					InstanceIp: instanceIP,
					UserName: username,
					Password: password,
					ClientKey: clientKey,
					ClientCert: clientCert,
					ServerCertificateAuthority: serverCertificateAuthority,
					ProjectId: projectID,
					InstanceName: instanceName,
				},
			},
		}
		o.secondaryIndexes = append(o.secondaryIndexes, si)
	}
}

// VStoreDataSink returns a vetl Sink
func VStoreDataSink(namespace, kind string, primaryKey []string, opts ...VStoreSinkOption) Sink {
	return func(r *sinkOption) {
		o := &vstoreSinkOption{}
		for _, f := range opts {
			f(o)
		}

		r.sink = &vetl_v1.DataSink{
			Sink: &vetl_v1.DataSink_Vstore{
				Vstore: &vetl_v1.VStoreSink{
					Namespace: namespace,
					Kind: kind,
					PrimaryKey: primaryKey,
					SecondaryIndexes: o.secondaryIndexes,
				},
			},
		}
	}
}

// TesseractDataSink returns a vetl Sink
func TesseractDataSink(namespace, kind string, primaryKey []string) Sink {
	return func(r *sinkOption) {
		tesseractSink := &vetl_v1.DataSink_Tesseract{
			Tesseract: &vetl_v1.TesseractSink{
				Namespace:  namespace,
				Kind:       kind,
				PrimaryKey: primaryKey,
			},
		}
		r.sink = &vetl_v1.DataSink{
			Sink: tesseractSink,
		}
	}
}
