Marketplace SDK
=================
An SDK to communicate with Marketplace's APIs (not marketplace-go) via Go.

## Session Transfer URL
Docs: https://www.vendasta.com/developers/vendors/session-transfer

Here's some sample code for getting a session transfer URL:

```
client := marketplace.NewAccountClient(ctx, appID, privateKey, config.CurEnv())
url, err := client.GetSessionTransferURL(ctx, accountGroupID)
```

`appID` is your marketplace application ID, probably loaded from an env
var

`privateKey` is what you set up during your marketplace integration
and should probably be loaded from a k8s secret - talk to SRE
if you need help getting that set up.

