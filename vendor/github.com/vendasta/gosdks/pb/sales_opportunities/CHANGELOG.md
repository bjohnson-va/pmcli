## 4.21.0
- Add actual_close_date_filter to the Filters message of ListOpportunitiesRequest message

## 4.20.0
- Added package_ids to ListOpportunitiesRequest

## 4.19.0
- Add reopened_reason to the opportunity message
- Add reason to ReopenClosedOpportunityRequest

## 4.18.0
- Add packages to CreateOpportunityRequest

## 4.17.0
- Add UpdateOpportunityMarket rpc

## 4.16.0
- Add partner_id and market_id to ListOpportunitiesRequest

## 4.15.0
- Add ProbabilityRangeFilter to ListOpportunitiesRequest

## 4.14.0
- Add ReopenClosedOpportunityRequest message
- Add ReopenClosedOpportunity rpc

## 4.13.1
- Import empty proto

## 4.13.0
- Add DeleteOpportunityRequest message
- Add DeleteOpportunity rpc

## 4.12.0
- Undeprecate UpdateLastConnectedDate rpc

## 4.11.0
- Add last_activity_date to opportunity
- Add UpdateLastActivityDate rpc
- Deprecate UpdateLastConnectedDate rpc
- Add LAST_ACTIVITY and LAST_CONNECTED to the SortField of ListOpportunitiesRequest

## 4.10.0
- Added quantity to the package message
- Added UpdateOpportunityPackage rpc

## 4.9.0
- Added total_projected_first_year_value to the ListOpportunitiesResponse

## 4.8.0
- Added total_results to the ListOpportunitiesResponse
- Added NAME to the SortField of ListOpportunitiesRequest

## 4.7.0
- Added filters and sort options to ListOpportunitiesRequest

## 4.6.0
- Add last_connected_date to the opportunity message
- Add UpdateLastConnectedDate rpc

## 4.5.0
- Include the sales_person_id on the ListOpportunitiesRequest

## 4.4.0
- Include the actual closed date on the opportunity

## 4.3.0
- Include the closed lost reason on the opportunity

## 4.2.0
- Split closing opportunities into two rpcs: CloseAsWonOpportunity and CloseAsLostOpportunity

## 4.1.0
- Add CloseOpportunity rpc
- Add pipeline stage and is closed to Opportunity

## 4.0.0
- Changed AddOpportunityPackage/RemoveOpportunityPackage to AddOpportunityPackages/RemoveOpportunityPackages

## 3.0.0
- Replace Revenue with Packages on Opportunity. Packages now contain the Revenue.
- Remove Revenue from CreateOpportunityRequest
- Remove UpdateOpportunityRevenue and UpdateOpportunityPackages
- Add AddOpportunityPackage, RemoveOpportunityPackage, and UpdateOpportunityPackageRevenue

## 2.0.0
- Remove UpdateOpportunity endpoint/messages

## 1.5.0
- Added account_group_id to individual field edit RPC request messages

## 1.4.0
- Added individual field edit RPCs: EditOpportunityName, EditOpportunityProbability, EditOpportunitySalesperson, 
EditOpportunityExpectedCloseDate, EditOpportunityNotes, EditOpportunityRevenue, EditOpportunityPackages

## 1.3.1
- Use CAD not CDN

## 1.3.0
- Update opportunity added.

## 1.2.1
- Pluralized repeated account_group_id field on GetOpportunityCountByAccountGroupsRequest for naming consistency

## 1.2.0
- Added projected_first_year_value to Opportunity

## 1.1.0
- Added GetOpportunityCountByAccountGroups to SalesOpportunities service.

## 1.0.1
- Change folder name to remove hyphens because hyphens completely break python protoc grpc generation.

## 1.0.0
- Initial commit of sales-opportunities protos
