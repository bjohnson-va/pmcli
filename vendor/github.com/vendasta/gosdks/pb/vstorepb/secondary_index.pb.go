// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vstorepb/secondary_index.proto

package vstorepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SecondaryIndexPropertyConfig struct {
	// Types that are valid to be assigned to Config:
	//	*SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig
	//	*SecondaryIndexPropertyConfig_CloudsqlPropertyConfig
	Config isSecondaryIndexPropertyConfig_Config `protobuf_oneof:"config"`
}

func (m *SecondaryIndexPropertyConfig) Reset()                    { *m = SecondaryIndexPropertyConfig{} }
func (m *SecondaryIndexPropertyConfig) String() string            { return proto.CompactTextString(m) }
func (*SecondaryIndexPropertyConfig) ProtoMessage()               {}
func (*SecondaryIndexPropertyConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isSecondaryIndexPropertyConfig_Config interface {
	isSecondaryIndexPropertyConfig_Config()
}

type SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig struct {
	ElasticsearchPropertyConfig *SecondaryIndexPropertyConfig_Elasticsearch `protobuf:"bytes,1,opt,name=elasticsearch_property_config,json=elasticsearchPropertyConfig,oneof"`
}
type SecondaryIndexPropertyConfig_CloudsqlPropertyConfig struct {
	CloudsqlPropertyConfig *SecondaryIndexPropertyConfig_CloudSQL `protobuf:"bytes,2,opt,name=cloudsql_property_config,json=cloudsqlPropertyConfig,oneof"`
}

func (*SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig) isSecondaryIndexPropertyConfig_Config() {
}
func (*SecondaryIndexPropertyConfig_CloudsqlPropertyConfig) isSecondaryIndexPropertyConfig_Config() {}

func (m *SecondaryIndexPropertyConfig) GetConfig() isSecondaryIndexPropertyConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *SecondaryIndexPropertyConfig) GetElasticsearchPropertyConfig() *SecondaryIndexPropertyConfig_Elasticsearch {
	if x, ok := m.GetConfig().(*SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig); ok {
		return x.ElasticsearchPropertyConfig
	}
	return nil
}

func (m *SecondaryIndexPropertyConfig) GetCloudsqlPropertyConfig() *SecondaryIndexPropertyConfig_CloudSQL {
	if x, ok := m.GetConfig().(*SecondaryIndexPropertyConfig_CloudsqlPropertyConfig); ok {
		return x.CloudsqlPropertyConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SecondaryIndexPropertyConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SecondaryIndexPropertyConfig_OneofMarshaler, _SecondaryIndexPropertyConfig_OneofUnmarshaler, _SecondaryIndexPropertyConfig_OneofSizer, []interface{}{
		(*SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig)(nil),
		(*SecondaryIndexPropertyConfig_CloudsqlPropertyConfig)(nil),
	}
}

func _SecondaryIndexPropertyConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SecondaryIndexPropertyConfig)
	// config
	switch x := m.Config.(type) {
	case *SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ElasticsearchPropertyConfig); err != nil {
			return err
		}
	case *SecondaryIndexPropertyConfig_CloudsqlPropertyConfig:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudsqlPropertyConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SecondaryIndexPropertyConfig.Config has unexpected type %T", x)
	}
	return nil
}

func _SecondaryIndexPropertyConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SecondaryIndexPropertyConfig)
	switch tag {
	case 1: // config.elasticsearch_property_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SecondaryIndexPropertyConfig_Elasticsearch)
		err := b.DecodeMessage(msg)
		m.Config = &SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig{msg}
		return true, err
	case 2: // config.cloudsql_property_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SecondaryIndexPropertyConfig_CloudSQL)
		err := b.DecodeMessage(msg)
		m.Config = &SecondaryIndexPropertyConfig_CloudsqlPropertyConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SecondaryIndexPropertyConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SecondaryIndexPropertyConfig)
	// config
	switch x := m.Config.(type) {
	case *SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig:
		s := proto.Size(x.ElasticsearchPropertyConfig)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SecondaryIndexPropertyConfig_CloudsqlPropertyConfig:
		s := proto.Size(x.CloudsqlPropertyConfig)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SecondaryIndexPropertyConfig_ElasticsearchField struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Index    string `protobuf:"bytes,3,opt,name=index" json:"index,omitempty"`
	Analyzer string `protobuf:"bytes,4,opt,name=analyzer" json:"analyzer,omitempty"`
}

func (m *SecondaryIndexPropertyConfig_ElasticsearchField) Reset() {
	*m = SecondaryIndexPropertyConfig_ElasticsearchField{}
}
func (m *SecondaryIndexPropertyConfig_ElasticsearchField) String() string {
	return proto.CompactTextString(m)
}
func (*SecondaryIndexPropertyConfig_ElasticsearchField) ProtoMessage() {}
func (*SecondaryIndexPropertyConfig_ElasticsearchField) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

func (m *SecondaryIndexPropertyConfig_ElasticsearchField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_ElasticsearchField) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_ElasticsearchField) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_ElasticsearchField) GetAnalyzer() string {
	if m != nil {
		return m.Analyzer
	}
	return ""
}

// Elasticsearch Property Config
type SecondaryIndexPropertyConfig_Elasticsearch struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index" json:"index,omitempty"`
	Exclude bool   `protobuf:"varint,4,opt,name=exclude" json:"exclude,omitempty"`
	// Allows fields to be stored multiple times with different analyzers
	Fields []*SecondaryIndexPropertyConfig_ElasticsearchField `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
}

func (m *SecondaryIndexPropertyConfig_Elasticsearch) Reset() {
	*m = SecondaryIndexPropertyConfig_Elasticsearch{}
}
func (m *SecondaryIndexPropertyConfig_Elasticsearch) String() string {
	return proto.CompactTextString(m)
}
func (*SecondaryIndexPropertyConfig_Elasticsearch) ProtoMessage() {}
func (*SecondaryIndexPropertyConfig_Elasticsearch) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 1}
}

func (m *SecondaryIndexPropertyConfig_Elasticsearch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_Elasticsearch) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_Elasticsearch) GetExclude() bool {
	if m != nil {
		return m.Exclude
	}
	return false
}

func (m *SecondaryIndexPropertyConfig_Elasticsearch) GetFields() []*SecondaryIndexPropertyConfig_ElasticsearchField {
	if m != nil {
		return m.Fields
	}
	return nil
}

type SecondaryIndexPropertyConfig_CloudSQL struct {
	Type    string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Exclude bool   `protobuf:"varint,2,opt,name=exclude" json:"exclude,omitempty"`
}

func (m *SecondaryIndexPropertyConfig_CloudSQL) Reset()         { *m = SecondaryIndexPropertyConfig_CloudSQL{} }
func (m *SecondaryIndexPropertyConfig_CloudSQL) String() string { return proto.CompactTextString(m) }
func (*SecondaryIndexPropertyConfig_CloudSQL) ProtoMessage()    {}
func (*SecondaryIndexPropertyConfig_CloudSQL) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 2}
}

func (m *SecondaryIndexPropertyConfig_CloudSQL) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SecondaryIndexPropertyConfig_CloudSQL) GetExclude() bool {
	if m != nil {
		return m.Exclude
	}
	return false
}

type SecondaryIndex struct {
	// Name of the secondary index, this name must be unique from other secondary indexes
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Index configuration and denotes the type of the secondary index as well.
	//
	// Types that are valid to be assigned to Index:
	//	*SecondaryIndex_EsRawConfig
	//	*SecondaryIndex_EsConfig
	//	*SecondaryIndex_CloudSqlConfig
	//	*SecondaryIndex_PubsubConfig
	//	*SecondaryIndex_BigQueryConfig
	Index isSecondaryIndex_Index `protobuf_oneof:"index"`
}

func (m *SecondaryIndex) Reset()                    { *m = SecondaryIndex{} }
func (m *SecondaryIndex) String() string            { return proto.CompactTextString(m) }
func (*SecondaryIndex) ProtoMessage()               {}
func (*SecondaryIndex) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type isSecondaryIndex_Index interface {
	isSecondaryIndex_Index()
}

type SecondaryIndex_EsRawConfig struct {
	EsRawConfig *ElasticsearchRawConfig `protobuf:"bytes,2,opt,name=es_raw_config,json=esRawConfig,oneof"`
}
type SecondaryIndex_EsConfig struct {
	EsConfig *ElasticsearchConfig `protobuf:"bytes,3,opt,name=es_config,json=esConfig,oneof"`
}
type SecondaryIndex_CloudSqlConfig struct {
	CloudSqlConfig *CloudSQLConfig `protobuf:"bytes,4,opt,name=cloud_sql_config,json=cloudSqlConfig,oneof"`
}
type SecondaryIndex_PubsubConfig struct {
	PubsubConfig *PubSubConfig `protobuf:"bytes,5,opt,name=pubsub_config,json=pubsubConfig,oneof"`
}
type SecondaryIndex_BigQueryConfig struct {
	BigQueryConfig *BigQueryConfig `protobuf:"bytes,6,opt,name=big_query_config,json=bigQueryConfig,oneof"`
}

func (*SecondaryIndex_EsRawConfig) isSecondaryIndex_Index()    {}
func (*SecondaryIndex_EsConfig) isSecondaryIndex_Index()       {}
func (*SecondaryIndex_CloudSqlConfig) isSecondaryIndex_Index() {}
func (*SecondaryIndex_PubsubConfig) isSecondaryIndex_Index()   {}
func (*SecondaryIndex_BigQueryConfig) isSecondaryIndex_Index() {}

func (m *SecondaryIndex) GetIndex() isSecondaryIndex_Index {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *SecondaryIndex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecondaryIndex) GetEsRawConfig() *ElasticsearchRawConfig {
	if x, ok := m.GetIndex().(*SecondaryIndex_EsRawConfig); ok {
		return x.EsRawConfig
	}
	return nil
}

func (m *SecondaryIndex) GetEsConfig() *ElasticsearchConfig {
	if x, ok := m.GetIndex().(*SecondaryIndex_EsConfig); ok {
		return x.EsConfig
	}
	return nil
}

func (m *SecondaryIndex) GetCloudSqlConfig() *CloudSQLConfig {
	if x, ok := m.GetIndex().(*SecondaryIndex_CloudSqlConfig); ok {
		return x.CloudSqlConfig
	}
	return nil
}

func (m *SecondaryIndex) GetPubsubConfig() *PubSubConfig {
	if x, ok := m.GetIndex().(*SecondaryIndex_PubsubConfig); ok {
		return x.PubsubConfig
	}
	return nil
}

func (m *SecondaryIndex) GetBigQueryConfig() *BigQueryConfig {
	if x, ok := m.GetIndex().(*SecondaryIndex_BigQueryConfig); ok {
		return x.BigQueryConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SecondaryIndex) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SecondaryIndex_OneofMarshaler, _SecondaryIndex_OneofUnmarshaler, _SecondaryIndex_OneofSizer, []interface{}{
		(*SecondaryIndex_EsRawConfig)(nil),
		(*SecondaryIndex_EsConfig)(nil),
		(*SecondaryIndex_CloudSqlConfig)(nil),
		(*SecondaryIndex_PubsubConfig)(nil),
		(*SecondaryIndex_BigQueryConfig)(nil),
	}
}

func _SecondaryIndex_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SecondaryIndex)
	// index
	switch x := m.Index.(type) {
	case *SecondaryIndex_EsRawConfig:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EsRawConfig); err != nil {
			return err
		}
	case *SecondaryIndex_EsConfig:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EsConfig); err != nil {
			return err
		}
	case *SecondaryIndex_CloudSqlConfig:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudSqlConfig); err != nil {
			return err
		}
	case *SecondaryIndex_PubsubConfig:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PubsubConfig); err != nil {
			return err
		}
	case *SecondaryIndex_BigQueryConfig:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BigQueryConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SecondaryIndex.Index has unexpected type %T", x)
	}
	return nil
}

func _SecondaryIndex_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SecondaryIndex)
	switch tag {
	case 2: // index.es_raw_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ElasticsearchRawConfig)
		err := b.DecodeMessage(msg)
		m.Index = &SecondaryIndex_EsRawConfig{msg}
		return true, err
	case 3: // index.es_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ElasticsearchConfig)
		err := b.DecodeMessage(msg)
		m.Index = &SecondaryIndex_EsConfig{msg}
		return true, err
	case 4: // index.cloud_sql_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CloudSQLConfig)
		err := b.DecodeMessage(msg)
		m.Index = &SecondaryIndex_CloudSqlConfig{msg}
		return true, err
	case 5: // index.pubsub_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PubSubConfig)
		err := b.DecodeMessage(msg)
		m.Index = &SecondaryIndex_PubsubConfig{msg}
		return true, err
	case 6: // index.big_query_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BigQueryConfig)
		err := b.DecodeMessage(msg)
		m.Index = &SecondaryIndex_BigQueryConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SecondaryIndex_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SecondaryIndex)
	// index
	switch x := m.Index.(type) {
	case *SecondaryIndex_EsRawConfig:
		s := proto.Size(x.EsRawConfig)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SecondaryIndex_EsConfig:
		s := proto.Size(x.EsConfig)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SecondaryIndex_CloudSqlConfig:
		s := proto.Size(x.CloudSqlConfig)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SecondaryIndex_PubsubConfig:
		s := proto.Size(x.PubsubConfig)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SecondaryIndex_BigQueryConfig:
		s := proto.Size(x.BigQueryConfig)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ElasticsearchRawConfig is deprecated and will be removed in a future release.
type ElasticsearchRawConfig struct {
	MappingJson  string `protobuf:"bytes,1,opt,name=mapping_json,json=mappingJson" json:"mapping_json,omitempty"`
	SettingsJson string `protobuf:"bytes,2,opt,name=settings_json,json=settingsJson" json:"settings_json,omitempty"`
	IndexName    string `protobuf:"bytes,3,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *ElasticsearchRawConfig) Reset()                    { *m = ElasticsearchRawConfig{} }
func (m *ElasticsearchRawConfig) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchRawConfig) ProtoMessage()               {}
func (*ElasticsearchRawConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ElasticsearchRawConfig) GetMappingJson() string {
	if m != nil {
		return m.MappingJson
	}
	return ""
}

func (m *ElasticsearchRawConfig) GetSettingsJson() string {
	if m != nil {
		return m.SettingsJson
	}
	return ""
}

func (m *ElasticsearchRawConfig) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

// ElasticsearchConfig uses our proprietary clusters as a destination for your indices.
// If you specify an ElasticsearchCluster, VStore will instead use that cluster as a destination.
type ElasticsearchConfig struct {
	NumberOfShards   int64                  `protobuf:"varint,1,opt,name=number_of_shards,json=numberOfShards" json:"number_of_shards,omitempty"`
	NumberOfReplicas int64                  `protobuf:"varint,2,opt,name=number_of_replicas,json=numberOfReplicas" json:"number_of_replicas,omitempty"`
	RefreshInterval  string                 `protobuf:"bytes,3,opt,name=refresh_interval,json=refreshInterval" json:"refresh_interval,omitempty"`
	Analysis         *ElasticsearchAnalysis `protobuf:"bytes,4,opt,name=analysis" json:"analysis,omitempty"`
	IndexName        string                 `protobuf:"bytes,5,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	Cluster          *ElasticsearchCluster  `protobuf:"bytes,6,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *ElasticsearchConfig) Reset()                    { *m = ElasticsearchConfig{} }
func (m *ElasticsearchConfig) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchConfig) ProtoMessage()               {}
func (*ElasticsearchConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ElasticsearchConfig) GetNumberOfShards() int64 {
	if m != nil {
		return m.NumberOfShards
	}
	return 0
}

func (m *ElasticsearchConfig) GetNumberOfReplicas() int64 {
	if m != nil {
		return m.NumberOfReplicas
	}
	return 0
}

func (m *ElasticsearchConfig) GetRefreshInterval() string {
	if m != nil {
		return m.RefreshInterval
	}
	return ""
}

func (m *ElasticsearchConfig) GetAnalysis() *ElasticsearchAnalysis {
	if m != nil {
		return m.Analysis
	}
	return nil
}

func (m *ElasticsearchConfig) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *ElasticsearchConfig) GetCluster() *ElasticsearchCluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type CloudSQLConfig struct {
	IndexName                  string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
	InstanceIp                 string `protobuf:"bytes,2,opt,name=instance_ip,json=instanceIp" json:"instance_ip,omitempty"`
	UserName                   string `protobuf:"bytes,3,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password                   string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	ClientKey                  []byte `protobuf:"bytes,5,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	ClientCert                 []byte `protobuf:"bytes,6,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	ServerCertificateAuthority []byte `protobuf:"bytes,7,opt,name=server_certificate_authority,json=serverCertificateAuthority,proto3" json:"server_certificate_authority,omitempty"`
	ProjectId                  string `protobuf:"bytes,8,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	InstanceName               string `protobuf:"bytes,9,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *CloudSQLConfig) Reset()                    { *m = CloudSQLConfig{} }
func (m *CloudSQLConfig) String() string            { return proto.CompactTextString(m) }
func (*CloudSQLConfig) ProtoMessage()               {}
func (*CloudSQLConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CloudSQLConfig) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

func (m *CloudSQLConfig) GetInstanceIp() string {
	if m != nil {
		return m.InstanceIp
	}
	return ""
}

func (m *CloudSQLConfig) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CloudSQLConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CloudSQLConfig) GetClientKey() []byte {
	if m != nil {
		return m.ClientKey
	}
	return nil
}

func (m *CloudSQLConfig) GetClientCert() []byte {
	if m != nil {
		return m.ClientCert
	}
	return nil
}

func (m *CloudSQLConfig) GetServerCertificateAuthority() []byte {
	if m != nil {
		return m.ServerCertificateAuthority
	}
	return nil
}

func (m *CloudSQLConfig) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CloudSQLConfig) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type PubSubConfig struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *PubSubConfig) Reset()                    { *m = PubSubConfig{} }
func (m *PubSubConfig) String() string            { return proto.CompactTextString(m) }
func (*PubSubConfig) ProtoMessage()               {}
func (*PubSubConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PubSubConfig) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

type BigQueryConfig struct {
	IndexName string `protobuf:"bytes,1,opt,name=index_name,json=indexName" json:"index_name,omitempty"`
}

func (m *BigQueryConfig) Reset()                    { *m = BigQueryConfig{} }
func (m *BigQueryConfig) String() string            { return proto.CompactTextString(m) }
func (*BigQueryConfig) ProtoMessage()               {}
func (*BigQueryConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *BigQueryConfig) GetIndexName() string {
	if m != nil {
		return m.IndexName
	}
	return ""
}

// https://www.elastic.co/guide/en/elasticsearch/guide/current/custom-analyzers.html
type ElasticsearchAnalysis struct {
	Analyzers   []*ElasticsearchAnalyzer   `protobuf:"bytes,1,rep,name=analyzers" json:"analyzers,omitempty"`
	Filters     []*ElasticsearchFilter     `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
	CharFilters []*ElasticsearchCharFilter `protobuf:"bytes,3,rep,name=char_filters,json=charFilters" json:"char_filters,omitempty"`
	Tokenizers  []*ElasticsearchTokenizer  `protobuf:"bytes,4,rep,name=tokenizers" json:"tokenizers,omitempty"`
}

func (m *ElasticsearchAnalysis) Reset()                    { *m = ElasticsearchAnalysis{} }
func (m *ElasticsearchAnalysis) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchAnalysis) ProtoMessage()               {}
func (*ElasticsearchAnalysis) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ElasticsearchAnalysis) GetAnalyzers() []*ElasticsearchAnalyzer {
	if m != nil {
		return m.Analyzers
	}
	return nil
}

func (m *ElasticsearchAnalysis) GetFilters() []*ElasticsearchFilter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *ElasticsearchAnalysis) GetCharFilters() []*ElasticsearchCharFilter {
	if m != nil {
		return m.CharFilters
	}
	return nil
}

func (m *ElasticsearchAnalysis) GetTokenizers() []*ElasticsearchTokenizer {
	if m != nil {
		return m.Tokenizers
	}
	return nil
}

// ElasticsearchAnalyzer configures a custom analyzer that can be built to transform your data into a
// configuration that suites your particular needs.
type ElasticsearchAnalyzer struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type          string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	StemExclusion []string `protobuf:"bytes,3,rep,name=stem_exclusion,json=stemExclusion" json:"stem_exclusion,omitempty"`
	StopWords     []string `protobuf:"bytes,4,rep,name=stop_words,json=stopWords" json:"stop_words,omitempty"`
	CharFilter    []string `protobuf:"bytes,5,rep,name=char_filter,json=charFilter" json:"char_filter,omitempty"`
	Tokenizer     string   `protobuf:"bytes,6,opt,name=tokenizer" json:"tokenizer,omitempty"`
	Filter        []string `protobuf:"bytes,7,rep,name=filter" json:"filter,omitempty"`
}

func (m *ElasticsearchAnalyzer) Reset()                    { *m = ElasticsearchAnalyzer{} }
func (m *ElasticsearchAnalyzer) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchAnalyzer) ProtoMessage()               {}
func (*ElasticsearchAnalyzer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ElasticsearchAnalyzer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ElasticsearchAnalyzer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ElasticsearchAnalyzer) GetStemExclusion() []string {
	if m != nil {
		return m.StemExclusion
	}
	return nil
}

func (m *ElasticsearchAnalyzer) GetStopWords() []string {
	if m != nil {
		return m.StopWords
	}
	return nil
}

func (m *ElasticsearchAnalyzer) GetCharFilter() []string {
	if m != nil {
		return m.CharFilter
	}
	return nil
}

func (m *ElasticsearchAnalyzer) GetTokenizer() string {
	if m != nil {
		return m.Tokenizer
	}
	return ""
}

func (m *ElasticsearchAnalyzer) GetFilter() []string {
	if m != nil {
		return m.Filter
	}
	return nil
}

// Token filters may change, add, or remove tokens.
type ElasticsearchFilter struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type        string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Pattern     string   `protobuf:"bytes,3,opt,name=pattern" json:"pattern,omitempty"`
	Replacement string   `protobuf:"bytes,4,opt,name=replacement" json:"replacement,omitempty"`
	Synonyms    []string `protobuf:"bytes,5,rep,name=synonyms" json:"synonyms,omitempty"`
}

func (m *ElasticsearchFilter) Reset()                    { *m = ElasticsearchFilter{} }
func (m *ElasticsearchFilter) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchFilter) ProtoMessage()               {}
func (*ElasticsearchFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ElasticsearchFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ElasticsearchFilter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ElasticsearchFilter) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ElasticsearchFilter) GetReplacement() string {
	if m != nil {
		return m.Replacement
	}
	return ""
}

func (m *ElasticsearchFilter) GetSynonyms() []string {
	if m != nil {
		return m.Synonyms
	}
	return nil
}

// Character filters are used to “tidy up” a string before it is tokenized.
type ElasticsearchCharFilter struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Pattern     string `protobuf:"bytes,3,opt,name=pattern" json:"pattern,omitempty"`
	Replacement string `protobuf:"bytes,4,opt,name=replacement" json:"replacement,omitempty"`
}

func (m *ElasticsearchCharFilter) Reset()                    { *m = ElasticsearchCharFilter{} }
func (m *ElasticsearchCharFilter) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchCharFilter) ProtoMessage()               {}
func (*ElasticsearchCharFilter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ElasticsearchCharFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ElasticsearchCharFilter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ElasticsearchCharFilter) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *ElasticsearchCharFilter) GetReplacement() string {
	if m != nil {
		return m.Replacement
	}
	return ""
}

// The tokenizer breaks up the string into individual terms or tokens.
type ElasticsearchTokenizer struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Delimiter string `protobuf:"bytes,3,opt,name=delimiter" json:"delimiter,omitempty"`
	Pattern   string `protobuf:"bytes,4,opt,name=pattern" json:"pattern,omitempty"`
}

func (m *ElasticsearchTokenizer) Reset()                    { *m = ElasticsearchTokenizer{} }
func (m *ElasticsearchTokenizer) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchTokenizer) ProtoMessage()               {}
func (*ElasticsearchTokenizer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ElasticsearchTokenizer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ElasticsearchTokenizer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ElasticsearchTokenizer) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

func (m *ElasticsearchTokenizer) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

// ElasticsearchCluster contains information necessary for VStore to communicate with an arbitrary ES Cluster
// VStore needs to be authed as a user with full CRUD permissions
type ElasticsearchCluster struct {
	HostName string `protobuf:"bytes,1,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *ElasticsearchCluster) Reset()                    { *m = ElasticsearchCluster{} }
func (m *ElasticsearchCluster) String() string            { return proto.CompactTextString(m) }
func (*ElasticsearchCluster) ProtoMessage()               {}
func (*ElasticsearchCluster) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *ElasticsearchCluster) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ElasticsearchCluster) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ElasticsearchCluster) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*SecondaryIndexPropertyConfig)(nil), "vstorepb.SecondaryIndexPropertyConfig")
	proto.RegisterType((*SecondaryIndexPropertyConfig_ElasticsearchField)(nil), "vstorepb.SecondaryIndexPropertyConfig.ElasticsearchField")
	proto.RegisterType((*SecondaryIndexPropertyConfig_Elasticsearch)(nil), "vstorepb.SecondaryIndexPropertyConfig.Elasticsearch")
	proto.RegisterType((*SecondaryIndexPropertyConfig_CloudSQL)(nil), "vstorepb.SecondaryIndexPropertyConfig.CloudSQL")
	proto.RegisterType((*SecondaryIndex)(nil), "vstorepb.SecondaryIndex")
	proto.RegisterType((*ElasticsearchRawConfig)(nil), "vstorepb.ElasticsearchRawConfig")
	proto.RegisterType((*ElasticsearchConfig)(nil), "vstorepb.ElasticsearchConfig")
	proto.RegisterType((*CloudSQLConfig)(nil), "vstorepb.CloudSQLConfig")
	proto.RegisterType((*PubSubConfig)(nil), "vstorepb.PubSubConfig")
	proto.RegisterType((*BigQueryConfig)(nil), "vstorepb.BigQueryConfig")
	proto.RegisterType((*ElasticsearchAnalysis)(nil), "vstorepb.ElasticsearchAnalysis")
	proto.RegisterType((*ElasticsearchAnalyzer)(nil), "vstorepb.ElasticsearchAnalyzer")
	proto.RegisterType((*ElasticsearchFilter)(nil), "vstorepb.ElasticsearchFilter")
	proto.RegisterType((*ElasticsearchCharFilter)(nil), "vstorepb.ElasticsearchCharFilter")
	proto.RegisterType((*ElasticsearchTokenizer)(nil), "vstorepb.ElasticsearchTokenizer")
	proto.RegisterType((*ElasticsearchCluster)(nil), "vstorepb.ElasticsearchCluster")
}

func init() { proto.RegisterFile("vstorepb/secondary_index.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1079 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0x6e, 0xfe, 0xb4, 0x89, 0x4f, 0xd2, 0xfc, 0xaa, 0xf9, 0x2d, 0xc5, 0xca, 0xb6, 0xbb, 0x5d,
	0x23, 0xa4, 0x22, 0x41, 0x2b, 0x2d, 0x48, 0x2c, 0x82, 0x95, 0x76, 0xdb, 0xdd, 0xaa, 0x05, 0x04,
	0x5b, 0x17, 0x89, 0x4b, 0xcb, 0x71, 0x4e, 0x9a, 0xd9, 0x3a, 0x63, 0x77, 0x66, 0xdc, 0x36, 0x95,
	0x10, 0x4f, 0x81, 0x78, 0x03, 0x2e, 0x79, 0x15, 0x5e, 0x80, 0x4b, 0x9e, 0x03, 0xa1, 0x19, 0xcf,
	0xd8, 0x71, 0x48, 0x68, 0xc5, 0x05, 0x77, 0x99, 0xcf, 0xe7, 0x7c, 0xe7, 0x3b, 0x7f, 0xe6, 0x4c,
	0xe0, 0xd1, 0x95, 0x90, 0x09, 0xc7, 0x74, 0xb0, 0x2f, 0x30, 0x4a, 0xd8, 0x30, 0xe4, 0xd3, 0x80,
	0xb2, 0x21, 0xde, 0xec, 0xa5, 0x3c, 0x91, 0x09, 0x69, 0xdb, 0xef, 0xde, 0x1f, 0x4d, 0xd8, 0x3a,
	0xb3, 0x36, 0x27, 0xca, 0xe4, 0x0d, 0x4f, 0x52, 0xe4, 0x72, 0x7a, 0x98, 0xb0, 0x11, 0x3d, 0x27,
	0xb7, 0xb0, 0x8d, 0x71, 0x28, 0x24, 0x8d, 0x04, 0x86, 0x3c, 0x1a, 0x07, 0xa9, 0xf9, 0x1e, 0x44,
	0xda, 0xc0, 0xad, 0xed, 0xd4, 0x76, 0x3b, 0x4f, 0x3f, 0xd9, 0xb3, 0x94, 0x7b, 0xff, 0x44, 0xb7,
	0xf7, 0x7a, 0x96, 0xeb, 0x78, 0xc5, 0x7f, 0x58, 0x21, 0x9f, 0x8b, 0x7d, 0x01, 0x6e, 0x14, 0x27,
	0xd9, 0x50, 0x5c, 0xc6, 0x7f, 0x0b, 0x5b, 0xd7, 0x61, 0xf7, 0xef, 0x19, 0xf6, 0x50, 0xd1, 0x9c,
	0x9d, 0x7e, 0x7d, 0xbc, 0xe2, 0x6f, 0x5a, 0xca, 0xaa, 0x49, 0x9f, 0x01, 0xa9, 0x88, 0x3b, 0xa2,
	0x18, 0x0f, 0x09, 0x81, 0x26, 0x0b, 0x27, 0xa8, 0xb3, 0x74, 0x7c, 0xfd, 0x5b, 0x61, 0x72, 0x9a,
	0xa2, 0x96, 0xe0, 0xf8, 0xfa, 0x37, 0x79, 0x00, 0xab, 0xba, 0xc0, 0x6e, 0x43, 0x83, 0xf9, 0x81,
	0xf4, 0xa1, 0x1d, 0xb2, 0x30, 0x9e, 0xde, 0x22, 0x77, 0x9b, 0xfa, 0x43, 0x71, 0xee, 0xff, 0x52,
	0x83, 0xf5, 0x4a, 0xc0, 0x82, 0xb7, 0xb6, 0x88, 0xb7, 0x3e, 0xcb, 0xeb, 0x42, 0x0b, 0x6f, 0xa2,
	0x38, 0x1b, 0xa2, 0xa6, 0x6d, 0xfb, 0xf6, 0x48, 0x4e, 0x61, 0x6d, 0xa4, 0x84, 0x0b, 0xb7, 0xb1,
	0xd3, 0xd8, 0xed, 0x3c, 0xfd, 0xec, 0xdf, 0xf4, 0x45, 0xa7, 0xee, 0x1b, 0xa2, 0xfe, 0x33, 0x68,
	0xdb, 0xf2, 0x2d, 0x94, 0x38, 0x23, 0xa6, 0x5e, 0x11, 0x73, 0xd0, 0x86, 0xb5, 0xbc, 0x5b, 0xde,
	0x9f, 0x75, 0xe8, 0x55, 0xe3, 0x2f, 0xac, 0xec, 0x11, 0xac, 0xa3, 0x08, 0x78, 0x78, 0x5d, 0xed,
	0xf2, 0x4e, 0x99, 0x44, 0x45, 0xa7, 0x1f, 0x5e, 0xe7, 0xf2, 0x8f, 0x57, 0xfc, 0x0e, 0x8a, 0xe2,
	0x48, 0xbe, 0x00, 0x07, 0x85, 0xe5, 0x68, 0x68, 0x8e, 0xed, 0x25, 0x1c, 0x05, 0x41, 0x1b, 0x85,
	0xf1, 0x7e, 0x05, 0x1b, 0x7a, 0x46, 0x02, 0x35, 0x77, 0x86, 0xa4, 0xa9, 0x49, 0xdc, 0x92, 0xc4,
	0x96, 0xa4, 0xf0, 0xef, 0x69, 0x9f, 0xb3, 0xcb, 0xd8, 0xb0, 0x3c, 0x87, 0xf5, 0x34, 0x1b, 0x88,
	0x6c, 0x60, 0x29, 0x56, 0x35, 0xc5, 0x66, 0x49, 0xf1, 0x26, 0x1b, 0x9c, 0x65, 0x83, 0x82, 0xa0,
	0x9b, 0x9b, 0x97, 0x22, 0x06, 0xf4, 0x3c, 0xb8, 0xcc, 0x90, 0x17, 0x33, 0xbf, 0x36, 0x2f, 0xe2,
	0x80, 0x9e, 0x9f, 0x2a, 0x83, 0x52, 0xc4, 0xa0, 0x82, 0x1c, 0xb4, 0xcc, 0xf8, 0x78, 0x3f, 0xc2,
	0xe6, 0xe2, 0xd2, 0x91, 0x27, 0xd0, 0x9d, 0x84, 0x69, 0x4a, 0xd9, 0x79, 0xf0, 0x56, 0x24, 0xcc,
	0xf4, 0xa3, 0x63, 0xb0, 0x2f, 0x45, 0xc2, 0xc8, 0x7b, 0xb0, 0x2e, 0x50, 0x4a, 0xca, 0xce, 0x45,
	0x6e, 0x93, 0x0f, 0x63, 0xd7, 0x82, 0xda, 0x68, 0x1b, 0x40, 0x87, 0x0a, 0x74, 0x57, 0xf3, 0x6b,
	0xe0, 0x68, 0xe4, 0x9b, 0x70, 0x82, 0xde, 0xaf, 0x75, 0xf8, 0xff, 0x82, 0xc2, 0x93, 0x5d, 0xd8,
	0x60, 0xd9, 0x64, 0x80, 0x3c, 0x48, 0x46, 0x81, 0x18, 0x87, 0x7c, 0x28, 0xb4, 0x84, 0x86, 0xdf,
	0xcb, 0xf1, 0x6f, 0x47, 0x67, 0x1a, 0x25, 0x1f, 0x02, 0x29, 0x2d, 0x39, 0xa6, 0x31, 0x8d, 0x42,
	0xa1, 0xa5, 0x34, 0xfc, 0x0d, 0x6b, 0xeb, 0x1b, 0x9c, 0x7c, 0x00, 0x1b, 0x1c, 0x47, 0x1c, 0xc5,
	0x38, 0xa0, 0x4c, 0x22, 0xbf, 0x0a, 0x63, 0x23, 0xea, 0x7f, 0x06, 0x3f, 0x31, 0x30, 0xf9, 0xdc,
	0xdc, 0x52, 0x41, 0x85, 0xe9, 0xf3, 0xe3, 0x25, 0xc3, 0xf2, 0xd2, 0x98, 0xf9, 0x85, 0xc3, 0x5c,
	0xda, 0xab, 0x73, 0x69, 0x93, 0x67, 0xd0, 0x8a, 0xe2, 0x4c, 0x48, 0xe4, 0xa6, 0x7b, 0x8f, 0x96,
	0xcd, 0x61, 0x6e, 0xe5, 0x5b, 0x73, 0xef, 0xb7, 0x3a, 0xf4, 0xaa, 0x43, 0x36, 0x17, 0xab, 0x36,
	0x1f, 0xeb, 0x31, 0x74, 0x28, 0x13, 0x32, 0x64, 0x11, 0x06, 0x34, 0x35, 0x4d, 0x02, 0x0b, 0x9d,
	0xa4, 0xe4, 0x21, 0x38, 0x99, 0x40, 0x3e, 0xdb, 0xa1, 0xb6, 0x02, 0xb4, 0x77, 0x1f, 0xda, 0x69,
	0x28, 0xc4, 0x75, 0xc2, 0x87, 0x76, 0x57, 0xd9, 0xb3, 0x0a, 0x1c, 0xc5, 0x14, 0x99, 0x0c, 0x2e,
	0x70, 0xaa, 0x93, 0xec, 0xfa, 0x4e, 0x8e, 0x7c, 0x85, 0x53, 0x15, 0xd8, 0x7c, 0x8e, 0x90, 0x4b,
	0x9d, 0x68, 0xd7, 0x37, 0x1e, 0x87, 0xc8, 0x25, 0x79, 0x01, 0x5b, 0x02, 0xf9, 0x15, 0x72, 0x6d,
	0x40, 0x47, 0x34, 0x0a, 0x25, 0x06, 0x61, 0x26, 0xc7, 0x09, 0xa7, 0x72, 0xea, 0xb6, 0xb4, 0x47,
	0x3f, 0xb7, 0x39, 0x2c, 0x4d, 0x5e, 0x5a, 0x0b, 0xa5, 0x20, 0xe5, 0xc9, 0x5b, 0x8c, 0x64, 0x40,
	0x87, 0x6e, 0x3b, 0x4f, 0xdd, 0x20, 0x27, 0x43, 0x35, 0xa1, 0x45, 0xea, 0x3a, 0x3b, 0x27, 0x9f,
	0x50, 0x0b, 0xea, 0x11, 0xfc, 0x08, 0xba, 0xb3, 0x57, 0xee, 0x8e, 0x72, 0x7a, 0xfb, 0xd0, 0xab,
	0xde, 0xaf, 0xbb, 0x1c, 0x7e, 0xae, 0xc3, 0x3b, 0x0b, 0xc7, 0x85, 0x3c, 0x07, 0xc7, 0xee, 0x7d,
	0x35, 0xdd, 0x8d, 0xbb, 0x46, 0xec, 0x16, 0xb9, 0x5f, 0x7a, 0x90, 0x4f, 0xa1, 0x35, 0xa2, 0xb1,
	0x54, 0xce, 0x75, 0xed, 0xbc, 0x6c, 0x99, 0x1d, 0x69, 0x2b, 0xdf, 0x5a, 0x93, 0x57, 0xd0, 0x8d,
	0xc6, 0x21, 0x0f, 0xac, 0x77, 0xfe, 0x26, 0x3c, 0x59, 0x36, 0x82, 0xe3, 0x90, 0x1b, 0x86, 0x4e,
	0x54, 0xfc, 0x16, 0xe4, 0x05, 0x80, 0x4c, 0x2e, 0x90, 0x51, 0x2d, 0xbf, 0xa9, 0x39, 0x96, 0xad,
	0xe4, 0xef, 0xac, 0xa1, 0x3f, 0xe3, 0xe3, 0xfd, 0x5e, 0x5b, 0x54, 0x99, 0x5b, 0xe4, 0xf7, 0x7e,
	0x5f, 0xdf, 0x87, 0x9e, 0x90, 0x38, 0x09, 0xf4, 0xd3, 0x22, 0x68, 0xc2, 0x74, 0x2e, 0x8e, 0xbf,
	0xae, 0xd0, 0xd7, 0x16, 0x54, 0x1d, 0x12, 0x32, 0x49, 0x03, 0x35, 0xb5, 0xb9, 0x54, 0xc7, 0x77,
	0x14, 0xf2, 0xbd, 0x02, 0xf4, 0xa0, 0x96, 0xf5, 0x70, 0x57, 0xf5, 0x77, 0x28, 0x73, 0x25, 0x5b,
	0xe0, 0x14, 0xb2, 0xf5, 0x1c, 0x3b, 0x7e, 0x09, 0x90, 0x4d, 0xf5, 0xb8, 0x6a, 0xcf, 0x96, 0xf6,
	0x34, 0x27, 0xef, 0xa7, 0xda, 0xdc, 0x6e, 0x33, 0x6c, 0xf7, 0x4d, 0xce, 0x85, 0x56, 0x1a, 0x4a,
	0x89, 0x9c, 0x99, 0x5b, 0x69, 0x8f, 0x64, 0x07, 0x3a, 0x6a, 0xd3, 0x85, 0x11, 0x4e, 0x90, 0x49,
	0x73, 0x2f, 0x67, 0x21, 0x75, 0x6d, 0xc5, 0x94, 0x25, 0x6c, 0x3a, 0x11, 0x26, 0x9f, 0xe2, 0xec,
	0xfd, 0x00, 0xef, 0x2e, 0x69, 0xf0, 0x7f, 0x21, 0xcd, 0xbb, 0x99, 0x7b, 0x73, 0x8a, 0xd9, 0xb8,
	0x77, 0xf4, 0x2d, 0x70, 0x86, 0x18, 0xd3, 0x09, 0x55, 0x35, 0x37, 0x4f, 0x4a, 0x01, 0xcc, 0x6a,
	0x6b, 0x56, 0xb4, 0x79, 0x17, 0xf0, 0x60, 0xd1, 0x72, 0x55, 0x0b, 0x70, 0x9c, 0x08, 0x39, 0x7b,
	0x7f, 0xdb, 0x0a, 0xb0, 0x0b, 0x50, 0x2d, 0x43, 0xfd, 0xad, 0x5e, 0x2e, 0x47, 0x36, 0xbf, 0x1c,
	0x1b, 0xd5, 0xe5, 0x38, 0x58, 0xd3, 0xff, 0xa9, 0x3f, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x90, 0xe8, 0xbe, 0x75, 0x0b, 0x00, 0x00,
}