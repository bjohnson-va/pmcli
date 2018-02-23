package vstore

import (
	"fmt"
	"encoding/json"
	"golang.org/x/oauth2/google"
	"golang.org/x/net/context"
	"net/http"
	"errors"
)

//UserInfo holds information necessary for identifying a user
type UserInfo struct {
	name string
}

//Namespace calculates a Namespace based off UserInfo
func (u *UserInfo) Namespace(original string) string {
	if u.name != "" {
		return fmt.Sprintf("%s-%s", original, u.name)
	}
	return original
}

// NewUserInfo returns a new UserInfo struct, handling necessary calculation based on environment.
// Providing `n` here will override the calculated username regardless of environment.
func NewUserInfo(e env, n string) (*UserInfo, error) {
	if (e == Local || e == Internal) && n == "" {
		var err error
		n, err = getLocalUserEmail(); if err != nil {
			return nil, err
		}
	}
	return &UserInfo{name: n}, nil
}

//GoogleAPIResponse represents the format of a response from Google's userinfo api
type GoogleAPIResponse struct {
	Email string `json:"email"`
}

func getLocalUserEmail() (string, error) {
	dts, err := google.DefaultTokenSource(context.Background(), "https://www.googleapis.com/auth/userinfo.email"); if err != nil {
		return "", err
	}
	token, err := dts.Token(); if err != nil {
		return "", err
	}

	url := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + token.AccessToken
	resp, err := http.Get(url); if err != nil {
		return "", errors.New("Could not determine local user email.")
	}
	if resp.StatusCode != 200 {
		return "", errors.New("Could not determine local user email.")
	}
	defer resp.Body.Close()

	data := GoogleAPIResponse{}
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&data)

	return data.Email, nil
}
