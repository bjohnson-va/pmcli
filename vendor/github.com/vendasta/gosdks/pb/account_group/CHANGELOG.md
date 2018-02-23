## 2.6.0
- Add new message ActionListAppend to UpdateOperation

## 2.5.1
- Fix imports to support multi package compilation

## 2.5.0
- Add salesforce_id to the AccountGroupExternalIdentifiers

## 2.4.0
- Change location filter to not have repeated fields for country/state. Remove city from the location filter.

## 2.3.0
- Add filters for taxonomy, location, has website, has social media pages

## 2.2.0
- Add Additional Salesperson Ids to 'AccountGroupExternalIdentifiers' message

## 2.1.0
- Add ReadFilter to the GetMulti request

## 2.0.0
- Remove is_taking_patients_deprecated and reserve it's tag number

## 1.9.0
- Add CompanyName sort option

## 1.8.0
- Removed unused RPCs
- Removed unused fields

## 1.7.1
- Fix FieldMask import

## 1.7.0
- Add FieldMasks to BulkUpdateOperations

## 1.6.1
- Add missing Billing Frequency Enum value for One time purchase

## 1.6.0
- Remove Search RPC.

## 1.5.0
- Deprecate Search RPC.
- Add search_term to Lookup request.

## 1.4.0
- Added social profile group id to create request message
- Added suspended to account group location message

## 1.3.0
- Added methods for creating/updating account groups.

## 1.2.2
- Fix the GetMulti rpc annotation to point at a `get-multi` uri rather than `get` to match our existing setup.

## 1.2.1
- Put total_results into the response object

## 1.2.0
- Use SortOptions to wrap up direction and sort_by field

## 1.1.0
- Add total_results to both SearchRequest and LookupRequest
- Add sort direction to LookupRequest

## 1.0.0
- Initial commit of account_group protos
