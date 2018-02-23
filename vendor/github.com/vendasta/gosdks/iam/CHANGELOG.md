# IAM Client Changelog
## 3.7.0
- Expose Attributes on Subject interface
## 3.6.0
- Add DigitalAgent as a subject type
## 3.5.0
- Add IsContextAuthorizedToAccessResource to authClient
- Can now pass custom resource attributes to evaluate
## 3.4.0
- Add register SMB subject property
## 3.3.0
- Add a property to Partner Subject: Can access marketing.
## 3.2.1
- Expose IsContextAuthorizedToAccessAccounts to the interface
## 3.2.0
- Added IsContextAuthorizedToAccessAccounts to authorize a user to accounts.
## 3.1.0
- Added a new attribute to the SMB persona: AccountAccessPermissions
## 3.0.0
- BREAKING CHANGE: Allow restricting service accounts to specified access scopes, replaces existing allowed service accounts parameter to NewAuthService.
This parameter is now specific to the current environment.
## 2.0.0
- SubjectValueIntersection must now specify an attribute instead of just a string.
## 1.2.0
- Back IsContextAuthorizedToAccessAccountGroups by the multi RPC.
## 1.1.1
- Fix subject.Add unmarshaling legacy attributes onto the struct
## 1.1.0
- Added CanAccessOrders to Partner
## 1.0.0
- Refactored Personas into their own files
- Added tags to unmarshal attributes onto peronas
- Added support for SubjectResourceFor policy clause
- Added an AttributeBuilder to build the new dynamic Attributes protos
## 0.6.0
- Added Vendor subject
## 0.5.0
- Remove list policy access scope
## 0.4.0
- Allow scopes to be specified on IsContextAuthorizedToAccessAccountGroup and IsContextAuthorizedToAccessPartnerMarket
## 0.3.1
- Move CreateTestJWT and MockPublicKey to session_util.go
## 0.3.0
- Get salesperson PartnerID from context namespace
## 0.2.0
- Added grpc.FailFast(false) to outgoing RPCs
- Added automatic retries
- rename iam.Context to iam.SubjectContext
## 0.1.0
- Initial release.
