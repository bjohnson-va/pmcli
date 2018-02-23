package smsdk

import (
	"encoding/json"
	"errors"
	"net/http"
)

//Account is a subset of the fields of a Social Marketing Account. Social Marketing basically only has account id and sso token specific to it.
type Account struct {
	AccountGroupID string `json:"accountGroupId"`
	PartnerID      string `json:"pid"`
	AccountID      string `json:"accountId"`
	MarketID       string `json:"marketId"`
	SsoToken       string `json:"ssoToken"`
}

func accountFromResponse(r *http.Response) (*Account, error) {
	defer r.Body.Close()
	type Response struct {
		Account *Account `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to Account: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Account, nil
}

func accountIDFromResponse(r *http.Response) (string, error) {
	defer r.Body.Close()
	type Response struct {
		Data map[string]string `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		reason := "Failed to convert response to Account: " + err.Error()
		return "", errors.New(reason)
	}
	if res.Data == nil {
		return "", errors.New("Failed to convert response to Account")
	}
	return res.Data["accountId"], nil
}
