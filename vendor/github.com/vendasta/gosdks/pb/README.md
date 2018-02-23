Table of Contents
=================

   * [Vendasta API Interfaces](#vendasta-api-interfaces)
      * [Directory Structure](#directory-structure)
      * [Proto Version/Package Name](#proto-versionpackage-name)
      * [Proto Syntax Standards](#proto-syntax-standards)
         * [Naming Conventions](#naming-conventions)
            * [Services](#services)
            * [RPC's](#rpcs)
            * [Messages](#messages)
      * [Documentation](#documentation)
         * [Generated Via Cloud Endpoints](#generated-via-cloud-endpoints)
         * [Generating Sequence Diagrams](#generating-sequence-diagrams)

Vendasta API Interfaces
=================
This repo will be the source of truth for our proto files. Your work on a new microservice or API will usually start here. PR's into this repo are a great opportunity to design your API interface ahead of time and have people vet your syntax and parameter naming, without being watered down in unrelated implementation details.

You will need to have someone from https://github.com/orgs/vendasta/teams/vendasta-apis to approve and merge your changes into master, so tag your tech lead specifically or `@vendasta/vendasta-apis`.

## Directory Structure
Do yourself a huge favour and start with a versioned folder structure to begin with. This saves us from having to store commit hashes or tags when we check out a proto for compilation, but rather we can just point to the folder. Of course semantic versioning applies, no breaking changes without creating a new version.

Your folder tree will then look as follows:
```
- listings
- CHANGELOG.md
-- V1
--- api.proto
--- listing.proto
```
Your api directory structure will need to match your microservice name, there is a requirement of no `_` in our microservice names so the directory you create can't have a `_` in it
You can break your proto definitions up into multiple files, `api.proto` is the entry proto, which contains your service definitions and your request/response messages, and imports all your other protos. Put proto messages representing models (ex: listing in `listing.proto`) in their own file.

Import your protos relatively. In the example above, `api.proto` would import listings.proto using `import "listing.proto";`.

Do not use hyphens in your path. Protoc python grpc generation can not handle hyphenated filenames or filepaths at this time: https://github.com/grpc/grpc/issues/5226

## Proto Version/Package Name
ALWAYS use proto V3. 

Your package name should be in the format \<microservice\>.version, if it is an internal only proto. For external protos use vendasta.\<microservice\>.version

```
syntax = "proto3";

package datalake.v1;
```

## Proto Syntax Standards

### Naming Conventions

#### Services
Use the microservice's name, since the microservice's name will directly reflect the domain that it is encapsulating. DO NOT postfix the service with "Service", this will result in stutter in auto-generated code.
```
service Datalake {...}
```

When you have other services that represent different levels of authentication, use a postfix that describes the user.
```
service DatalakeAdmin {...}
service DatalakeInternal {...}
```
#### RPC's
Since the microservice should already be striving to work in a single domain, there should not be a large number of domain-specific services defined from a single microservice. However, there are two other reasons why you might have multiple services in a single microservice: 1) services requiring different levels of authentication (datalake, datalake admin), or 2) when it makes sense to split up your connection pooling (TODO: provide more context as to when this is needed).

There is a limited number of verbs to use when considering an RPC name. Your RPC will fall under one of the following categories:
- Get
- List
- Search
- Create
- Update
- Delete
- Intent or Specific Action
- Replace

Since we are favouring less services based on the microservice name as opposed to more services based on the microservice's models, we will need to include the model name in the RPC name.

So when you combine the verb with your model name, your service will end up looking like this:
```
service Datalake {
  rpc GetListing (...) returns (...) {};
  rpc ListListings (...) returns (...) {}; // Returned in some defined order, ex: created, alphabetical
  rpc SearchListings (...) returns (...) {}; // Returned in scored order, the better the match to the search, the higher
  rpc CreateListing (...) returns (...) {};
  rpc UpdateListing (...) returns (...) {};
  rpc DeleteListing (...) returns (...) {};
  rpc RescrapeListing (...) returns (...) {}; // Intent/specific action
  rpc ReplaceListing (...) returns (...) {}; // Replaces a listing if it exists, otherwise creates it
}
```
#### Messages
There are two types of messages:
- Messages defining generic, reuseable types/objects
- Messages defining RPC request/response schemas

In the first case of a generic type or object, the model name alone (without no prefix or postfix) should be used:
```
message Listing {...}
message ListingStats {...}
```

In the second case, when you're defining messages for the purpose of being a request or response, you WILL use both a prefix and a postfix. Prefix with the rpc verb from the RPC section, postfix with `Request` or `Response`.
```
message GetListingRequest {
  string listing_id = 1;
}

message Listing {...}

message ListListingsResponse {
  repeated Listing listings = 1;
  string next_cursor = 2; 
  bool has_more = 3;
}

service Datalake {
  rpc GetListing (GetListingRequest) returns (GetListingResponse) {};
  rpc GetMultiListings (GetMultiListingRequest) returns (GetMultiListingResponse) {};
  rpc ListListings (ListListingsRequest) returns (ListListingsResponse) {};
  rpc SearchListings (SearchListingsRequest) returns (SearchListingsResponse) {};
  rpc CreateListing (CreateListingRequest) returns (CreateListingResponse) {};
  rpc UpdateListing (UpdateListingRequest) returns (UpdateListingResponse) {};
  rpc DeleteListing (DeleteListingRequest) returns (DeleteListingResponse) {};
  rpc RescrapeListing (RescrapeListingRequest) returns (RescrapeListingResponse) {};
  rpc ReplaceListing (ReplaceListingRequest) returns (ReplaceListingResponse) {};
}
```

All Paged endpoints must include at least the following properties:
```
message ListItemsRequest {
  string cursor = 1;
  int64 page_size = 2;
}

message ListItemsResponse {
  repeated Item items = 1;
  string next_cursor = 2;
  bool has_more = 3;
}
```
Depending on your use case, it may be desirable to add such properties as `int64 total_results`.

If your `ListItemsRequest` needs more options like filters, it is recommended you encase those options with separate messages.
```
message Filters {
  string partner_id = 1;
  repeated string market_ids = 2;
  string sales_person_id = 3;
}

message ListItemsRequest {
  string cursor = 1;
  int64 page_size = 2;
  Filters filters = 3;
}
```

You should use a string cursor *even if* you are using Elasticsearch integer offsets in the background. Decoding a string into an integer is not hard, and this lets us easily switch backing technologies without being hamstrung by Elasticsearch's paging particulars.

## Documentation
All protos should have complete documentation.
This includes:
* All services
* All RPCs
* All request and response messages
* All inner messages in all request and response messages
* Each field in each message, unless documenting the field would be redundant, confusing or otherwise not helpful.

### Generated Via Cloud Endpoints
Documentation is automatically generated for us by a service that reads the protos deployed to Google Cloud Endpoints.

You can see this documentation at: 
- https://endpoints-test.vendasta-internal.com/
- https://endpoints-demo.vendasta-internal.com/
- https://endpoints-prod.vendasta-internal.com/

Simply place `//` comments above your proto definition statements and they will applied automatically by the doc viewer.

```
//Gets a listing
message GetListingRequest {
  //The unique identifier for a listing
  string listing_id = 1;
}
```

The above code would result in `GetListingRequest` being associated with the `Gets a listing` comment in the documentation,
and `listing_id` would be explained as `The unique identifier for a listing`.

### Generating Sequence Diagrams
If your RPCs intend to span multiple services, you should explain the interactions using:
* Comments on the RPCs
* A visual tool like a sequence diagram

You can write sequence diagrams as plaintext .sd files in this repository adjacent to their relevant service using this [syntax](https://www.websequencediagrams.com/examples.html)

Ex:
```
title AccessControl

Client->IAM: AccessResource: Does subject "S-123" have access to resource id "AccountGroup" with identifiers "AccountGroupID: AG-123, PartnerID: ABC" for scope "Read"?
note right of IAM: Look up policies for resource "AccountGroup". Filter policies by scope "Read". Policy "CanRead" found.
IAM->ResourceOwner: GetResourceRequest: Give me resource of type "AccountGroup" with identifiers "AccountGroupID: AG-123, PartnerID: ABC"
ResourceOwner-->IAM: GetResourceResponse: Ok. "AccountGroupID: AG-123, PartnerID: ABC, MarketID: west"
note right of IAM: Evaluate policy "CanRead" with subject "S-123" and Resource attributes from GetResourceResponse.
IAM-->Client: google.protobuf.Empty (An error code will signify access denied, a success will signify access granted)
```

Run `inv docs` to generate a .png file for each .sd file
