package billingsdk

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// ContractInfo is the most recent contract
type ContractInfo struct {
	Activated      time.Time `json:"activated_date"`
	Cost           int       `json:"cost"`
	Currency       string    `json:"currency"`
	Frequency      string    `json:"frequency"`
	Modified       time.Time `json:"modified_date"`
	Name           string    `json:"name"`
	NetDays        int       `json:"net_d"`
	OnboardingFee  int       `json:"onboarding_fee"`
	Signed         time.Time `json:"signed_date"`
	Status         string    `json:"status"`
	SubscriptionID string    `json:"subscription_id"`
	Verified       time.Time `json:"verified_datetime"`
	Discount       *Discount `json:"discount"`
}

// Discount is a discount applied to a contract
type Discount struct {
	End   time.Time `json:"endDate"`
	ID    int64     `json:"id"`
	Start time.Time `json:"startDate"`
	Type  string    `json:"type"`
	Value int       `json:"value"`
}

func contractInfoFromResponse(r io.Reader) (*ContractInfo, error) {
	type Response struct {
		Resp *ContractInfo `json:"data"`
	}
	res := Response{}
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		reason := "Failed to convert response to ContractInfo: " + err.Error()
		return nil, errors.New(reason)
	}
	return res.Resp, nil
}
