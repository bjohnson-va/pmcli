# IAM
- [Introduction](https://github.com/vendasta/gosdks/tree/master/iam#installation)
- [Usage](https://github.com/vendasta/gosdks/tree/master/iam#usage)
  - [Introduction](https://github.com/vendasta/gosdks/tree/master/iam#introduction)

## Installation

* YOLO development? `go get github.com/vendasta/gosdks`
* Using [glide](https://github.com/Masterminds/glide) package manager? `glide get -s -v github.com/vendasta/gosdks`

## Usage

### Introduction

```golang
iam, err := iam.NewClient(ctx, config.CurEnv(), grpc.WithUnaryInterceptor(logging.ClientInterceptor()))
if err != nil {
    logging.Errorf(ctx, "Error initializing IAM client %s", err.Error())
    os.Exit(-1)
}

// This should be retrieved from configuration in practice
accessScopedServiceAccounts := ServiceAccountToScopes{
    "partner-central-test@appspot.gserviceaccount.com": {READ, WRITE, DELETE},
    "developers-vendasta-test@appspot.gserviceaccount.com": {READ},
    ...
}
var iamAuthService = iam.NewAuthService(iamAPI, accessScopedServiceAccounts)

partner, err := iam.PartnerBySessionID(ctx, iamSession)
if err != nil {
    logging.Infof(ctx, "Error getting partner by session id %s %s", iamSession, err.Error())
    return nil, err
}

// Check if context has read access to the requested account group
err := iamAuthService.IsContextAuthorizedToAccessAccountGroup(ctx, request.AccountGroupId, iam.READ)
if err != nil {
    return nil, util.ToGrpcError(err)
}
```

### Partner Application Auth

This section describes how to set a Partner Application up for auth with IAM.

1. Create a service account for their application using `RegisterPartnerApp(ctx context.Context, partnerID string, email string) (string, error)`. 
    - The string return value is the service account's subjectID, which is not important for this flow.

2. Call `PartnerAppGenerateClientKey(ctx context.Context, partnerID, email string) (*ClientKey, error)` with the same values for partnerID and email.
    - The return value is their secret containing their private key. 
    - Save this! It will be attached to the Partner's application and is used to generate the information necessary for the token exchange to work properly.

3. The partner app can now call `ExchangeToken(ctx context.Context, key *ClientKey) (string, error)` where `key` is the secret generated in the previous step.
    - The return value string is their bearer token that they can provide on subsequent calls as an IAM session.
    - The logic for `ExchangeToken` can be replicated in any language, as it is only an exchange of JWTs.
    
`PartnerAppAddKey` and `PartnerAppRemoveKey` provide more fine-grained control around the keys of a service account but should not be directly necessary in this workflow.
