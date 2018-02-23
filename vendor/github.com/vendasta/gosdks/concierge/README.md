Concierge SDK
=================
An sdk to communicate with Concierge via Go.

## Versions

### 2.0.0 [2017-11-30]
Create Account and Create Custom task both take partnerId and accountGroupId rather than pid and agid

### 1.1.0 [2017-09-11]
Adds:
// CreateAccount creates a concierge account for the given businessID
CreateAccount(ctx context.Context, partnerID string, businessID string) (int, error)

### 1.0.0 [2017-09-11]
Initial Release
