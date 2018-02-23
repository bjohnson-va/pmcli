# VStore Client Changelog

## 0.24.0
- Handle rate-limit errors on CreateKind/DeleteKind

## 0.23.0
- Added vstore.WithInternalTransport() as a ClientOption to use an experimental client side load balancing transport.

## 0.22.0
- Add support for bytes field

## 0.21.0
- Allow clients to pass a proto-level callback to Scan that will be executed in favor of the default vstore.Model callback

## 0.20.2
- Set FailFast option to false for DeleteKind

## 0.20.1
- Fix Transaction function to also retry on AlreadyExists -- this can happen when two Creates on the same entity are racing, the loser should retry as an update.

## 0.20.0
- Allow clients to exclude specific fields from Elastic/Cloud SQL secondary indexes using CloudSQLExclusion and ElasticsearchExclusion

## 0.19.1
- Fix Transaction function in vstore mock.
## 0.19.0
- The `fieldOption` type is now public, called `FieldOption`.
- Added TransactionOption type as an optional variadic arg to the Transaction api
  - `WithProtoTransaction` allows a transaction to be processed in terms of the raw VStore proto rather than requiring a struct with VStore tags

## 0.18.0
- **Breaking**: Add options for backup config builder that hide VStore protobufs from consumers
- **How To Upgrade**:
  - You used to do this: `backupConfig := vstore.NewBackupConfigBuilder().PeriodicBackup(vstorepb.BackupConfig_MONTHLY).Build()`
  - Do this instead: `backupConfig := vstore.NewBackupConfigBuilder().PeriodicBackup(vstore.MonthlyBackup).Build()`

## 0.17.0
- Add a new option for range matching on Scan: ScanRangeFilter

## 0.16.0
- Add a new option for partial prefix matching on Scan: ScanPartialFilter

## 0.15.0
- Add a new option when creating/updating a kind: ExtendedKeyLength
  - When used, key components will not be validated for length, but the CloudSQL secondary index cannot be used

## 0.14.0
- Add a new option for partial prefix matching on Lookup: PartialFilter

## 0.13.2
- Fixing a minor regression from 0.11.1, where the adapter would a ListValue with a empty list of Values rather than a nil list

## 0.13.1
- Update client SSL certs for internal environment

## 0.13.0
- Deprecate RawElasticsearch secondary index functionality
- Warn that it will be removed in a 1.0.0 release

## 0.12.0
- Add experimental API for Scan

## 0.11.1
- Performance improvements to VStore message deserialization

## 0.11.0
- Ripped out the pubsub specific code into a pubsub module

## 0.10.0
- Allow custom cluster configurations for elasticsearch secondary indexes.
  - Breaking change: You can provide the type to an ElasticsearchField option. 
    - This is necessary for Elastic 5.X to be able to make analyzed `text` multifields underneath non-analyzed `keyword` fields. 
    - Not providing the type will mean that vstore will choose a default.
    - If you already use these options, all you need to do is change `vstore.ElasticsearchField("raw", "not_analyzed")` to `vstore.ElasticsearchField("raw", "not_analyzed", "")` to retain the same behaviour.

## 0.9.0
- Add automatic retry support for Get, GetMulti, Lookup and Transaction.

## 0.8.0
- Added helper functions for converting between byte arrays and vstore models

## 0.7.0
- Added pubsub secondary indexing as well as a way to attach a callback to a subscription

## 0.6.1
- Fix a bug where the lookup stub was returning too many results relative to specified pageSize.

## 0.6.0
- Added in memory happy path stub for vstore.

## 0.5.0
* Added EnvFromEnvironmentVariable function to automatically calculate the environment and address that the VStore client should use

## 0.4.0
* Added GetSecondaryIndexName function to vstore client interface

## 0.3.0
* Added CloudSQL secondary indexes to the schema builder

## 0.2.0
* More documentation and examples. Not finished.
* Move changelog to its own file.
* Add lint task for all libs and lint VStore.

## 0.1.0
* Initial release.
* Available operations include Get, GetMulti, Lookup, and Transaction.
* Namespace and kind management is also available.
