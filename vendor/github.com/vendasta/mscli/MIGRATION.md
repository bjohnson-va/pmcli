# Code Changes

## 2.23.0 (SDKGen 0.14.0)

The hand-rolled files in the python sdk, usually called something like api.py, no longer need to do `from_proto()`. Instead of

```
response = AccountGroupAPI._get_client().Lookup(projection_filter.to_proto(), filters.to_proto(), cursor, page_size, sort_options.to_proto(), search_term)
response_vobject = PagedResponse.from_proto(response)
return response_vobject.account_groups, response_vobject.next_cursor, response_vobject.has_more, response_vobject.total_results
```

write

```
response = AccountGroupAPI._get_client().Lookup(projection_filter.to_proto(), filters.to_proto(), cursor, page_size, sort_options.to_proto(), search_term)
return response.account_groups, response.next_cursor, response.has_more, response.total_results
```
