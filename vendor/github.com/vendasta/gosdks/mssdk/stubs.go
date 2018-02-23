package mssdk

import "context"

// AccountClientStub implements the AccountClientInterface for the purposes of testing
type AccountClientStub struct {
	CreateResponse   *Account
	ActivateResponse string
	ErrorResponse    error
	Params           map[string]interface{}
}

//Activate returns ActivateResponse and ErrorResponse, as well as populating Params with the params this function was called with
func (s AccountClientStub) Activate(ctx context.Context, accountGroupID string, partnerID string, ssoToken string) (string, error) {
	s.Params = map[string]interface{}{
		"agid":     accountGroupID,
		"pid":      partnerID,
		"ssoToken": ssoToken,
	}
	return s.ActivateResponse, s.ErrorResponse
}

//Deactivate returns Empty and ErrorResponse, as well as populating Params with the params this function was called with
func (s AccountClientStub) Deactivate(ctx context.Context, accountID string, partnerID string) error {
	s.Params = map[string]interface{}{
		"msid": accountID,
		"pid":  partnerID,
	}
	return s.ErrorResponse
}
