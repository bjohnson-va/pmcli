package vstore

import (
	"errors"

	"github.com/vendasta/gosdks/pb/vstorepb"
)

// NewSchema creates a new schema for a vStore kind.
// Ex:
//	vstore.NewSchema(
//		"vbc",
//		"AccountGroup",
//		vstore.NewFieldBuilder().StringField(
//			"account_group_id", Required(),
//		).StructField(
//			"accounts", vstore.NewFieldBuilder().StringField("account", Required()).Build(),
//		).Build(),
//	)
//
func NewSchema(namespace string, kind string, primaryKey []string, properties []*Property, secondaryIndexes []*SecondaryIndex, backupConfig *BackupConfig, opts ...SchemaOption) *Schema {
	sch := &Schema{
		Namespace:        namespace,
		Kind:             kind,
		PrimaryKey:       primaryKey,
		Properties:       properties,
		SecondaryIndexes: secondaryIndexes,
		BackupConfig:     backupConfig,
	}
	for _, o := range opts {
		o(sch)
	}
	return sch
}

// SchemaOption applies itself to a schema
type SchemaOption func(sch *Schema) *Schema

// ExtendedKeyLength option allows longer key components at the cost of restricted secondary indexes (No CloudSQL)
func ExtendedKeyLength() SchemaOption {
	return func(sch *Schema) *Schema {
		sch.ExtendedKeyLength = true
		return sch
	}
}

//NewPropertyBuilder returns a struct that stores a schema's property definitions
func NewPropertyBuilder() *propertyBuilder {
	return &propertyBuilder{}
}

//NewSecondaryIndexBuilder returns a struct that stores a schema's secondary index definitions
func NewSecondaryIndexBuilder() *secondaryIndexBuilder {
	return &secondaryIndexBuilder{}
}

//NewBackupConfigBuilder returns a struct that stores a schema's backup configuration details
func NewBackupConfigBuilder() *backupConfigBuilder {
	return &backupConfigBuilder{}
}

//FieldOption applies a customization to a field/property.
type FieldOption func(*Property)

//Required makes a property required.
func Required() FieldOption {
	return func(p *Property) {
		p.IsRequired = true
	}
}

//Repeated makes a property repeated.
func Repeated() FieldOption {
	return func(p *Property) {
		p.IsRepeated = true
	}
}

type elasticFieldOption func(*vstorepb.SecondaryIndexPropertyConfig_Elasticsearch, *Property)

func getElasticType(vstoreType FieldType) string {
	switch vstoreType {
	case StringType:
		return "string"
	case IntType:
		return "integer"
	case FloatType:
		return "double"
	case BoolType:
		return "boolean"
	case TimeType:
		return "date"
	case GeoPointType:
		return "geo_point"
	case StructType:
		return "nested"
	}
	panic(ErrUnhandledTypeConversion)
}

// ElasticsearchField is used for storing an elasticsearch property more then once with different index types and field types.
// Example: vstore.ElasticsearchField("raw", "analyzed", "string")
// https://www.elastic.co/guide/en/elasticsearch/reference/2.4/multi-fields.html
// https://www.elastic.co/guide/en/elasticsearch/reference/5.4/multi-fields.html
func ElasticsearchField(name string, indexType string, fieldType string) elasticFieldOption {
	return func(c *vstorepb.SecondaryIndexPropertyConfig_Elasticsearch, p *Property) {
		if fieldType == "" {
			fieldType = getElasticType(p.FType)
		}
		c.Fields = append(c.Fields, &vstorepb.SecondaryIndexPropertyConfig_ElasticsearchField{
			Name:  name,
			Type:  fieldType,
			Index: indexType,
		})
	}
}

// ElasticsearchExclusion will prevent the field this field option is set on from being replicated to Elasticsearch
func ElasticsearchExclusion() elasticFieldOption {
	return func(c *vstorepb.SecondaryIndexPropertyConfig_Elasticsearch, p *Property) {
		c.Exclude = true
	}
}

// ElasticsearchProperty applies customized mappings to this property in Elasticsearch.
func ElasticsearchProperty(indexName string, indexType string, fieldOptions ...elasticFieldOption) FieldOption {
	return func(p *Property) {
		if p.SecondaryIndexConfigs == nil {
			p.SecondaryIndexConfigs = map[string]*vstorepb.SecondaryIndexPropertyConfig{}
		}
		config := &vstorepb.SecondaryIndexPropertyConfig_Elasticsearch{
			Type:  getElasticType(p.FType),
			Index: indexType,
		}
		for _, opt := range fieldOptions {
			opt(config, p)
		}
		p.SecondaryIndexConfigs[indexName] = &vstorepb.SecondaryIndexPropertyConfig{
			&vstorepb.SecondaryIndexPropertyConfig_ElasticsearchPropertyConfig{
				config,
			},
		}
	}
}

type cloudSQLFieldOption func(*vstorepb.SecondaryIndexPropertyConfig_CloudSQL, *Property)

func CloudSQLFieldType(typeOverride string) cloudSQLFieldOption {
	return func(c *vstorepb.SecondaryIndexPropertyConfig_CloudSQL, p *Property) {
		c.Type = typeOverride
	}
}

// CloudSQLExclusion will prevent the field this field option is set on from being replicated to CloudSQL
func CloudSQLExclusion() cloudSQLFieldOption {
	return func(c *vstorepb.SecondaryIndexPropertyConfig_CloudSQL, p *Property) {
		c.Exclude = true
	}
}

func CloudSQLProperty(indexName string, fieldOptions ...cloudSQLFieldOption) FieldOption {
	return func(p *Property) {
		if p.SecondaryIndexConfigs == nil {
			p.SecondaryIndexConfigs = map[string]*vstorepb.SecondaryIndexPropertyConfig{}
		}
		config := &vstorepb.SecondaryIndexPropertyConfig_CloudSQL{}
		for _, opt := range fieldOptions {
			opt(config, p)
		}
		p.SecondaryIndexConfigs[indexName] = &vstorepb.SecondaryIndexPropertyConfig{
			&vstorepb.SecondaryIndexPropertyConfig_CloudsqlPropertyConfig{
				config,
			},
		}
	}
}

//Schema defines how a VStore Kind is both stored and replicated
type Schema struct {
	Namespace         string
	Kind              string
	PrimaryKey        []string
	Properties        []*Property
	SecondaryIndexes  []*SecondaryIndex
	BackupConfig      *BackupConfig
	ExtendedKeyLength bool
}

//BackupConfig defines the kind's backup strategy
type BackupConfig struct {
	BackupConfigPb *vstorepb.BackupConfig
}

type backupConfigBuilder struct {
	config *BackupConfig
}

//PeriodicBackup adds a backup policy that causes the kind to be backed up based on a defined period
func (b *backupConfigBuilder) PeriodicBackup(frequency backupFrequency) *backupConfigBuilder {
	b.config = &BackupConfig{
		&vstorepb.BackupConfig{
			Frequency: frequency(),
		},
	}
	return b
}

//Build returns the backup config
func (b *backupConfigBuilder) Build() *BackupConfig {
	return b.config
}

type backupFrequency func() vstorepb.BackupConfig_BackupFrequency

// DailyBackup will cause your kind to be backed up once every day
func DailyBackup() vstorepb.BackupConfig_BackupFrequency {
	return vstorepb.BackupConfig_DAILY
}

// WeeklyBackup will cause your kind to be backed up once every week
func WeeklyBackup() vstorepb.BackupConfig_BackupFrequency {
	return vstorepb.BackupConfig_WEEKLY
}

// MonthlyBackup will cause your kind to be backed up once every month
func MonthlyBackup() vstorepb.BackupConfig_BackupFrequency {
	return vstorepb.BackupConfig_MONTHLY
}

//SecondaryIndex is a construct that defines a replication destination for a kind, including any configuration relevant for that index type
type SecondaryIndex struct {
	SecondaryIndexPb *vstorepb.SecondaryIndex
}

type ElasticsearchIndexOption func(*vstorepb.ElasticsearchConfig)

//ElasticsearchNumberOfShards controls how many shards the index is created with.
func ElasticsearchNumberOfShards(n int64) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		c.NumberOfShards = n
	}
}

//ElasticsearchNumberOfReplicas controls how many replicas the index is created with.
func ElasticsearchNumberOfReplicas(n int64) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		c.NumberOfReplicas = n
	}
}

//ElasticsearchRefreshInterval controls how often your index is refreshed
func ElasticsearchRefreshInterval(s string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		c.RefreshInterval = s
	}
}

//ElasticsearchClusterConfig specifies an arbitrary ES Cluster that VStore will replicate this index to
//Note that Elastic-Cloud managed clusters run on version 5.X on elastic
func ElasticsearchClusterConfig(hostName, username, password string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		c.Cluster = &vstorepb.ElasticsearchCluster{
			HostName: hostName,
			Username: username,
			Password: password,
		}
	}
}

//ElasticsearchAnalyzer creates a new ElasticsearchAnalyzer
func ElasticsearchAnalyzer(name, analyzerType, tokenizer string, stemExclusion, stopWords, charFilter, filter []string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		a := &vstorepb.ElasticsearchAnalyzer{
			Name:          name,
			Type:          analyzerType,
			Tokenizer:     tokenizer,
			StemExclusion: stemExclusion,
			StopWords:     stopWords,
			CharFilter:    charFilter,
			Filter:        filter,
		}
		if c.Analysis == nil {
			c.Analysis = &vstorepb.ElasticsearchAnalysis{}
		}
		c.Analysis.Analyzers = append(c.Analysis.Analyzers, a)
	}
}

//ElasticsearchFilter creates a new ElasticsearchFilter
func ElasticsearchFilter(name, analyzerType, pattern, replacement string, synonyms []string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		f := &vstorepb.ElasticsearchFilter{
			Name:        name,
			Type:        analyzerType,
			Pattern:     pattern,
			Replacement: replacement,
			Synonyms:    synonyms,
		}
		if c.Analysis == nil {
			c.Analysis = &vstorepb.ElasticsearchAnalysis{}
		}
		c.Analysis.Filters = append(c.Analysis.Filters, f)
	}
}

//ElasticsearchCharFilter creates a new ElasticsearchCharFilter
func ElasticsearchCharFilter(name, analyzerType, pattern, replacement string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		f := &vstorepb.ElasticsearchCharFilter{
			Name:        name,
			Type:        analyzerType,
			Pattern:     pattern,
			Replacement: replacement,
		}
		if c.Analysis == nil {
			c.Analysis = &vstorepb.ElasticsearchAnalysis{}
		}
		c.Analysis.CharFilters = append(c.Analysis.CharFilters, f)
	}
}

//ElasticsearchTokenizer creates a new ElasticsearchTokenizer
func ElasticsearchTokenizer(name, analyzerType, pattern, delimiter string) ElasticsearchIndexOption {
	return func(c *vstorepb.ElasticsearchConfig) {
		f := &vstorepb.ElasticsearchTokenizer{
			Name:      name,
			Type:      analyzerType,
			Pattern:   pattern,
			Delimiter: delimiter,
		}
		if c.Analysis == nil {
			c.Analysis = &vstorepb.ElasticsearchAnalysis{}
		}
		c.Analysis.Tokenizers = append(c.Analysis.Tokenizers, f)
	}
}

type secondaryIndexBuilder struct {
	indexes []*SecondaryIndex
}

// Deprecated: Use Elasticsearch instead
// This will be removed in 1.0.0
func (s *secondaryIndexBuilder) RawElasticsearch(name string, MappingJSON string, SettingsJSON string) *secondaryIndexBuilder {
	i := &vstorepb.SecondaryIndex{
		Name: name,
		Index: &vstorepb.SecondaryIndex_EsRawConfig{
			EsRawConfig: &vstorepb.ElasticsearchRawConfig{
				MappingJson:  MappingJSON,
				SettingsJson: SettingsJSON,
			},
		},
	}
	s.indexes = append(s.indexes, &SecondaryIndex{i})
	return s
}

//Elasticsearch adds an elasticsearch index to the definition with the specified settings
func (s *secondaryIndexBuilder) Elasticsearch(name string, opts ...ElasticsearchIndexOption) *secondaryIndexBuilder {
	c := &vstorepb.ElasticsearchConfig{
		NumberOfShards:   5,
		NumberOfReplicas: 1,
		RefreshInterval:  "1s",
	}
	for _, opt := range opts {
		opt(c)
	}

	i := &vstorepb.SecondaryIndex{
		Name: name,
		Index: &vstorepb.SecondaryIndex_EsConfig{
			EsConfig: c,
		},
	}
	s.indexes = append(s.indexes, &SecondaryIndex{i})
	return s
}

//CloudSQL adds a CloudSQL index to the definition with the specified settings
func (s *secondaryIndexBuilder) CloudSQL(name, instanceIP, userName, password, projectID, instanceName string, clientKey, clientCert, serverCertificateAuthority []byte) *secondaryIndexBuilder {
	i := &vstorepb.SecondaryIndex{
		Name: name,
		Index: &vstorepb.SecondaryIndex_CloudSqlConfig{
			CloudSqlConfig: &vstorepb.CloudSQLConfig{
				InstanceIp:                 instanceIP,
				UserName:                   userName,
				Password:                   password,
				ClientKey:                  clientKey,
				ClientCert:                 clientCert,
				ServerCertificateAuthority: serverCertificateAuthority,
				ProjectId:                  projectID,
				InstanceName:               instanceName,
			},
		},
	}
	s.indexes = append(s.indexes, &SecondaryIndex{i})
	return s
}

//PubSub adds a PubSub index to the definition with the specified settings
func (s *secondaryIndexBuilder) PubSub(name string) *secondaryIndexBuilder {
	i := &vstorepb.SecondaryIndex{
		Name: name,
		Index: &vstorepb.SecondaryIndex_PubsubConfig{
			&vstorepb.PubSubConfig{},
		},
	}
	s.indexes = append(s.indexes, &SecondaryIndex{i})
	return s
}

// BigQuery adds a Big Query index to the definition with the specified settings
func (s *secondaryIndexBuilder) BigQuery(name string) *secondaryIndexBuilder {
	i := &vstorepb.SecondaryIndex{
		Name: name,
		Index: &vstorepb.SecondaryIndex_BigQueryConfig{
			&vstorepb.BigQueryConfig{},
		},
	}
	s.indexes = append(s.indexes, &SecondaryIndex{i})
	return s
}

//Build returns the secondary index definitions
func (s *secondaryIndexBuilder) Build() []*SecondaryIndex {
	return s.indexes
}

type propertyBuilder struct {
	properties []*Property
}

//StringProperty adds a string property to the list of property definitions
func (s *propertyBuilder) StringProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: StringType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}
	return s
}


//StringProperty adds a string property to the list of property definitions
func (s *propertyBuilder) BytesProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: BytesType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}
	return s
}

//IntegerProperty adds an integer property to the list of property definitions
func (s *propertyBuilder) IntegerProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: IntType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//IntegerProperty adds an integer property to the list of property definitions
func (s *propertyBuilder) FloatProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: FloatType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//BooleanProperty adds a boolean property to the list of property definitions
func (s *propertyBuilder) BooleanProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: BoolType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//TimeProperty adds a time property to the list of property definitions
func (s *propertyBuilder) TimeProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: TimeType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//GeoPointProperty adds a geo point property to the list of property definitions
func (s *propertyBuilder) GeoPointProperty(name string, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:  name,
		FType: GeoPointType,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//StructProperty adds a struct property to the list of property definitions
func (s *propertyBuilder) StructProperty(name string, properties []*Property, opts ...FieldOption) *propertyBuilder {
	f := &Property{
		Name:       name,
		FType:      StructType,
		Properties: properties,
	}
	s.properties = append(s.properties, f)
	for _, opt := range opts {
		opt(f)
	}

	return s
}

//Build returns the property definitions
func (s *propertyBuilder) Build() []*Property {
	return s.properties
}

//Property is the definition of a single field/column and how it is stored, validated and structured.
type Property struct {
	Name                  string
	FType                 FieldType
	IsRequired            bool
	IsRepeated            bool
	Properties            []*Property
	SecondaryIndexConfigs map[string]*vstorepb.SecondaryIndexPropertyConfig
}

//ToPb return the protobuf representation of a Property
func (p *Property) ToPb() (property *vstorepb.Property, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok || err != ErrUnhandledTypeConversion {
				panic(r)
			}
		}
	}()
	var properties []*vstorepb.Property
	if len(p.Properties) > 0 {
		properties, err = PropertiesToPb(p.Properties...)
		if err != nil {
			return
		}
	}
	property = &vstorepb.Property{
		Name:                  p.Name,
		Required:              p.IsRequired,
		Repeated:              p.IsRepeated,
		Type:                  p.FType.ToPb(),
		Properties:            properties,
		SecondaryIndexConfigs: p.SecondaryIndexConfigs,
	}
	return
}

//PropertiesToPb returns a list of protobuf serialized properties
func PropertiesToPb(properties ...*Property) ([]*vstorepb.Property, error) {
	var err error
	propertiesPb := make([]*vstorepb.Property, len(properties))
	for i, p := range properties {
		propertiesPb[i], err = p.ToPb()
		if err != nil {
			return nil, err
		}
	}
	return propertiesPb, nil
}

type FieldType int64

//The complete spectrum of supported field types
const (
	StringType FieldType = iota
	IntType
	FloatType
	BoolType
	TimeType
	GeoPointType
	StructType
	BytesType
)

//ErrUnhandledTypeConversion is thrown when an unsupported field type is detected
var ErrUnhandledTypeConversion = errors.New("Unhandled type conversion to pb.")

func (f FieldType) ToPb() vstorepb.Property_Type {
	switch f {
	case StringType:
		return vstorepb.Property_STRING
	case BytesType:
		return vstorepb.Property_BYTES
	case IntType:
		return vstorepb.Property_INT64
	case FloatType:
		return vstorepb.Property_DOUBLE
	case BoolType:
		return vstorepb.Property_BOOL
	case TimeType:
		return vstorepb.Property_TIMESTAMP
	case GeoPointType:
		return vstorepb.Property_GEOPOINT
	case StructType:
		return vstorepb.Property_STRUCT
	}
	panic(ErrUnhandledTypeConversion)
}

//SecondaryIndexesToPb transforms a list of secondary indexes into protobuf format
func SecondaryIndexesToPb(secondaryIndexes ...*SecondaryIndex) []*vstorepb.SecondaryIndex {
	secondaryIndexesPb := make([]*vstorepb.SecondaryIndex, len(secondaryIndexes))
	for i, secondaryIndex := range secondaryIndexes {
		secondaryIndexesPb[i] = secondaryIndex.SecondaryIndexPb
	}
	return secondaryIndexesPb
}

//BackupConfigToPb transforms a backup configuration into protobuf format
func BackupConfigToPb(backupConfig *BackupConfig) *vstorepb.BackupConfig {
	if backupConfig == nil {
		return nil
	}
	return backupConfig.BackupConfigPb
}
