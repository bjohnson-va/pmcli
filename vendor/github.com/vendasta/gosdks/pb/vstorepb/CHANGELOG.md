## 0.11.0
- Add new property type: bytes

## 0.10.1
- Fix imports to support multi package compilation

## 0.10.0
- Add option to enable extended keys for a schema. This option will allow much longer keys, but 
prevents the addition of secondary indexes that require shorter keys (CloudSQL).

## 0.9.0
- Add support for range filters

## 0.8.0
- Add support for partial prefix filters

## 0.7.0
- Added analyzer field to Elastic Property Field definition

## 0.6.0
- Add scan RPC

## 0.5.0
- Add the ability to define a secondary index using an arbitrary Elasticsearch cluster

## 0.4.0
- Add the ability to define a secondary index using Google BigQuery with your vstore schema

## 0.3.0
- Add the ability to define a secondary index using Google Cloud Pubsub with your vstore schema

## 0.2.0
- Added the ability to define a backup configuration for a kind to `vstorepb.CreateKindRequest`, and read that configuration from `vstorepb.GetKindResponse`
- Added backup configuration to the proto for a `vstorepb.Schema`

## 0.1.3
- The folder name really mattered to protoc's importing mechanisms, avoiding breaking changes

## 0.1.2
- Do away with versioned package name to avoid breaking changes

## 0.1.1
- Fix package name to be relative to this repository to allow for easy code generation

## 0.1.0
- Minimal set of protos required to communicate with the vstore service

