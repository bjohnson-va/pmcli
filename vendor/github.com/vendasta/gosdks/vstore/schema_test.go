package vstore

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/vstorepb"
)

func getDummyFields() []*Property {
	return NewPropertyBuilder().StringProperty("artist_id").Build()
}

func TestElasticsearchNumberOfShardsSetsShardNumbersOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchNumberOfShards(12)).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, int64(12), schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().NumberOfShards)
}

func TestElasticsearchNumberOfReplicasSetsReplicasOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchNumberOfReplicas(3)).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, int64(3), schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().NumberOfReplicas)
}

func TestElasticsearchRefreshIntervalSetsRefreshIntervalOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchRefreshInterval("5s")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "5s", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().RefreshInterval)
}

func TestElasticsearchDefaultsAreSetAsExpectedInAbsenceOfOptions(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch").Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "1s", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().RefreshInterval)
	assert.Equal(t, int64(5), schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().NumberOfShards)
	assert.Equal(t, int64(1), schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().NumberOfReplicas)
}

func TestElasticsearchAnalyzerSetsAnalyzerOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchAnalyzer("artist_id", "analyzed", "tokenizer", []string{"stemInclusion1", "stemInclusion2"}, []string{"stopWord1", "stopWord2"}, []string{"charFilter1", "charFilter2"}, []string{"filter1", "filter2"})).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "artist_id", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].Name)
	assert.Equal(t, "analyzed", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].Type)
	assert.Equal(t, "tokenizer", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].Tokenizer)
	assert.Equal(t, []string{"stemInclusion1", "stemInclusion2"}, schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].StemExclusion)
	assert.Equal(t, []string{"stopWord1", "stopWord2"}, schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].StopWords)
	assert.Equal(t, []string{"charFilter1", "charFilter2"}, schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].CharFilter)
	assert.Equal(t, []string{"filter1", "filter2"}, schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].Filter)
}

func TestMultipleElasticsearchAnalyzerOptionsSetsMultipleAnalyzerOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchAnalyzer("analyzer1", "", "", nil, nil, nil, nil), ElasticsearchAnalyzer("analyzer2", "", "", nil, nil, nil, nil)).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "analyzer1", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[0].Name)
	assert.Equal(t, "analyzer2", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Analyzers[1].Name)
}

func TestElasticsearchFilterSetsFilterOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchFilter("artist_id", "analyzed", "pattern", "replacement", []string{"synonym1", "synonym1"})).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "artist_id", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Name)
	assert.Equal(t, "analyzed", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Type)
	assert.Equal(t, "pattern", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Pattern)
	assert.Equal(t, "replacement", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Replacement)
	assert.Equal(t, []string{"synonym1", "synonym1"}, schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Synonyms)
}

func TestMultipleElasticsearchFilterOptionsSetsMultipleFilterOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchFilter("filter1", "", "", "", nil), ElasticsearchFilter("filter2", "", "", "", nil)).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "filter1", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[0].Name)
	assert.Equal(t, "filter2", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Filters[1].Name)
}

func TestElasticsearchCharFilterSetsCharFilterOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchCharFilter("artist_id", "analyzed", "pattern", "replacement")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "artist_id", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[0].Name)
	assert.Equal(t, "analyzed", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[0].Type)
	assert.Equal(t, "pattern", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[0].Pattern)
	assert.Equal(t, "replacement", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[0].Replacement)
}

func TestMultipleElasticsearchCharFilterOptionsSetsMultipleCharFilterOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchCharFilter("charFilter1", "", "", ""), ElasticsearchCharFilter("charFilter2", "", "", "")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "charFilter1", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[0].Name)
	assert.Equal(t, "charFilter2", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.CharFilters[1].Name)
}

func TestElasticsearchTokenizerSetsTokenizerOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchTokenizer("artist_id", "analyzed", "pattern", "delimiter")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "artist_id", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[0].Name)
	assert.Equal(t, "analyzed", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[0].Type)
	assert.Equal(t, "pattern", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[0].Pattern)
	assert.Equal(t, "delimiter", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[0].Delimiter)
}

func TestMultipleElasticsearchTokenizerOptionsSetsMultipleTokenizerOptionsOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchTokenizer("tokenizer1", "", "", ""), ElasticsearchTokenizer("tokenizer2", "", "", "")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "tokenizer1", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[0].Name)
	assert.Equal(t, "tokenizer2", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Analysis.Tokenizers[1].Name)
}

func TestElasticsearchClusterConfigSetsElasticsearchClusterOnSchema(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchClusterConfig("10.10.0.1:8080", "username", "password")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)
	assert.Equal(t, "10.10.0.1:8080", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Cluster.HostName)
	assert.Equal(t, "username", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Cluster.Username)
	assert.Equal(t, "password", schema.SecondaryIndexes[0].SecondaryIndexPb.GetEsConfig().Cluster.Password)
}

func TestCloudSQLBuildsSetsExpectedIndexPropertiesOnSchema(t *testing.T) {
	fields := getDummyFields()
	name := "cloudSql"
	instanceIP := "104.154.165.235:3306"
	userName := "cbass"
	password := "yeehawbuddy"
	projectID := "repcore-prod"
	instanceName := "my-cool-instance"
	clientKey := []byte("big ass client key")
	clientCert := []byte("big ass client cert")
	serverCertificateAuthority := []byte("ginormous CA")
	secondaryIndexes := NewSecondaryIndexBuilder().CloudSQL(name, instanceIP, userName, password, projectID, instanceName, clientKey, clientCert, serverCertificateAuthority).Build()
	schema := NewSchema("myname", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)

	assert.Equal(t, name, schema.SecondaryIndexes[0].SecondaryIndexPb.Name)
	assert.Equal(t, instanceIP, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().InstanceIp)
	assert.Equal(t, userName, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().UserName)
	assert.Equal(t, password, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().Password)
	assert.Equal(t, projectID, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().ProjectId)
	assert.Equal(t, instanceName, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().InstanceName)
	assert.Equal(t, clientKey, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().ClientKey)
	assert.Equal(t, clientCert, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().ClientCert)
	assert.Equal(t, serverCertificateAuthority, schema.SecondaryIndexes[0].SecondaryIndexPb.GetCloudSqlConfig().ServerCertificateAuthority)
}

func TestCloudSQLFieldTypeOverridesDefaultFieldType(t *testing.T) {
	fields := NewPropertyBuilder().BooleanProperty("went_gold", CloudSQLProperty("cloudsql", CloudSQLFieldType("BIGINT"))).Build()
	secondaryIndexes := NewSecondaryIndexBuilder().CloudSQL("cloudsql", "104.154.165.235:3306", "cbass", "yeehawbuddy", "repcore-local", "instanceName", []byte("big ass client key"), []byte("big ass client cert"), []byte("ginormous CA")).Build()
	schema := NewSchema("cbass", "Song", []string{"artist_id"}, fields, secondaryIndexes, nil)

	assert.Equal(t, "BIGINT", schema.Properties[0].SecondaryIndexConfigs["cloudsql"].GetCloudsqlPropertyConfig().Type)
}

func Test_NewSchema_UsesExtendedKeyLengthOption(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchNumberOfShards(12)).Build()

	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil, ExtendedKeyLength())

	assert.True(t, schema.ExtendedKeyLength)
}

func Test_NewSchema_DefaultsExtendedKeyLength(t *testing.T) {
	fields := getDummyFields()
	secondaryIndexes := NewSecondaryIndexBuilder().Elasticsearch("elasticsearch", ElasticsearchNumberOfShards(12)).Build()

	schema := NewSchema("cbass", "Song", []string{"artist_id", "album_id", "track_number"}, fields, secondaryIndexes, nil)

	assert.False(t, schema.ExtendedKeyLength)
}

func TestBackupConfigBuilder_PeriodicBackup(t *testing.T) {
	type test struct {
		expected vstorepb.BackupConfig_BackupFrequency
		option backupFrequency
		name string
	}

	cases := []*test{
		{
			expected: vstorepb.BackupConfig_DAILY,
			option: DailyBackup,
			name: "Using Daily Backup should set the backup config object to daily",
		},
		{
			expected: vstorepb.BackupConfig_WEEKLY,
			option: WeeklyBackup,
			name: "Using Weekly Backup should set the backup config object to weekly",
		},
		{
			expected: vstorepb.BackupConfig_MONTHLY,
			option: MonthlyBackup,
			name: "Using Monthly Backup should set the backup config object to monthly",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := NewBackupConfigBuilder().PeriodicBackup(c.option).Build()
			assert.Equal(t, c.expected, b.BackupConfigPb.Frequency)
		})
	}
}
