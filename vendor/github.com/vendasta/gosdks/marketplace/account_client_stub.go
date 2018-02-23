package marketplace

import (
	"golang.org/x/net/context"
)

type AccountClientStub struct {
	Method                        string
	Params                        map[string]interface{}
	GetSessionTransferURLResponse string
	ErrorResponse                 error
}

func (a *AccountClientStub) GetSessionTransferURL(ctx context.Context, accountID string) (string, error) {
	a.Method = "GetSessionTransferURL"
	a.Params = make(map[string]interface{})
	a.Params["accountID"] = accountID
	return a.GetSessionTransferURLResponse, a.ErrorResponse
}
