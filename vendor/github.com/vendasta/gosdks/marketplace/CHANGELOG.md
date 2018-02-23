### 6.3.0
- Added Activity client

### 6.2.0
- Added GetUserAccountPermission to UserClient

### 6.1.0
- Added userClient to get user information from marketplace
### 6.0.0
- Changed interfaces of `NewAccountClient`, `NewOAuthClient` and `BuildSolutionClient` to accept an env instead of a rootURL (for convenience)

### 5.1.0
- Added `AccountClient` with `GetSessionTransferURL`
- Added `OAuthClient` with `GetOAuthToken`

### 5.0.0 [2017-06-01]
- Require the root url to be passed in to the clients. The environment switch for which url to pass in should be in the environment variables for the µs

## 4.0.0
- Price Frequency is now a BillingFrequency
- Deprecate billing frequency

## 3.0.1
- Pricing currency will now be an enum instead of string
- Deprecate solution currency in favor of the pricing currency

### 3.0.0
- Marketplace SDK now implements the basesdk

### 2.0.0
- SolutionId, PartnerId, MarketId, IconUrl in the Solution struct are now SolutionID, PartnerID, MarketID, IconURL
- Documented all public methods

### 1.0.0
Initial Release
