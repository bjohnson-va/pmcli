package marketplace

import (
	"golang.org/x/net/context"
)

type UserClientStub struct {
	Method                           string
	Params                           map[string]interface{}
	GetUserResponse                  *User
	GetUserAccountPermissionResponse bool
	ErrorResponse                    error
}

func (a *UserClientStub) GetUser(ctx context.Context, userID string) (*User, error) {
	a.Method = "GetUser"
	a.Params = make(map[string]interface{})
	a.Params["userID"] = userID
	return a.GetUserResponse, a.ErrorResponse
}

func (a *UserClientStub) GetUserAccountPermission(ctx context.Context, userID string, accountID string) (bool, error) {
	a.Method = "GetUser"
	a.Params = make(map[string]interface{})
	a.Params["userID"] = userID
	a.Params["accountID"] = accountID
	return a.GetUserAccountPermissionResponse, a.ErrorResponse
}
