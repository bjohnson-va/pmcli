package salestool

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

//Salesperson is a salesperson
type Salesperson struct {
	// Address is the address of the salesperson
	Address string `json:"address"`

	// City is the city of the salesperson
	City string `json:"city"`

	//Country is the country of the salesperson
	Country string `json:"country"`

	//Department is the Department of the salesperson
	Department string `json:"department"`

	//Email is the Email of the salesperson
	Email string `json:"email"`

	//FirstName is the First Name of the salesperson
	FirstName string `json:"firstName"`

	//JobTitle is the Job Title of the salesperson
	JobTitle string `json:"jobTitle"`

	//LastName is the Last Name of the salesperson
	LastName string `json:"lastName"`

	//PhoneNumber is the Phone Number of the salesperson
	PhoneNumber []string `json:"phoneNumber"`

	//PhotoURL is the Photo URL of the salesperson
	PhotoURL string `json:"photoUrl"`

	//PhotoURLSecure is the Photo URL secure of the salesperson
	PhotoURLSecure string `json:"photoUrlSecure"`

	//SalespersonID is the salesperson id of the salesperson
	SalespersonID string `json:"salesPersonId"`

	//MarketID is the market id of the salesperson
	MarketID string `json:"marketId"`

	//PID is the pid of the salesperson
	PartnerID string `json:"pid"`

	//Region is the region of the salesperson
	Region string `json:"region"`

	//State is the state of the salesperson
	State string `json:"state"`

	//Supervisor is the supervisor of the salesperson
	Supervisor string `json:"supervisor"`

	//Zip is the zip of the salesperson
	Zip string `json:"zip"`

	//CoverPageTitle is the cover page title of the salesperson
	CoverPageTitle string `json:"coverPageTitle"`
}

// SalespersonFromResponse converts an http response from marketplace to a Salesperson
func salespersonFromResponse(r *http.Response) (*Salesperson, error) {
	defer r.Body.Close()
	var err error

	type Data struct {
		Data map[string]interface{} `json:"data"`
	}

	data := Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		reason := "Failed to convert response to Salesperson: " + err.Error()
		return nil, errors.New(reason)
	}

	s := Salesperson{}
	for _, val := range data.Data {
		spJSON, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(bytes.NewReader(spJSON)).Decode(&s)
		break
	}
	return &s, err
}
