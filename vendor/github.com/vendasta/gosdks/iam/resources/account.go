package resources

const (
	// AccountResourceID resource ID for the account resource
	AccountResourceID = "account"

	// AccountOwnerID owner ID of the account resource
	AccountOwnerID = "accounts-microservice"

	// AccountResourceAccountID required resource identifier for the account resource
	AccountResourceAccountID = "account_id"

	// AccountResourceAccountGroupID required resource identifier for the account resource
	AccountResourceAccountGroupID = "account_group_id"
)

// AccountIdentifier identifiers required to identify an account resource
type AccountIdentifier struct {
	AccountID string
	AccountGroupID string
}

// ToResourceIdentifier converts the AccountIdentifier to the resource entity identifier proto
func (a *AccountIdentifier) ToResourceIdentifier() map[string][]string {
	return map[string][]string{
		AccountResourceAccountID: []string{a.AccountID},
		AccountResourceAccountGroupID: []string{a.AccountGroupID},
	}
}

// NewAccountIdentifier creates a new Account Identifier
func NewAccountIdentifier(accountID string, accountGroupID string) *AccountIdentifier {
	return &AccountIdentifier{AccountID: accountID, AccountGroupID: accountGroupID}
}
