package salestool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

func Test_GetSalesperson_ReturnsASalesperson(t *testing.T) {

	r := GetSalespersonRequest{
		PartnerID:     "ABC",
		SalespersonID: "SAL123",
	}
	baseClient := &basesdk.BaseClientMock{}

	response := `{"data": {"SALES-ABC123": {"city": "New York", "zip": "10005", "supervisor": "Patrick", "photoUrlSecure": "https://lh4.ggpht.com", "firstName": "Robert", "photoUrl": "http://lh4.ggpht.com/6Qq0eNFBmbTFUly", "country": "US", "region": "APlace", "lastName": "Salesguy", "pid": "SRP", "state": "NY", "phoneNumber": ["5555555555", "5551234567"], "salesPersonId": "SALES-ABC123", "jobTitle": "Seller", "address": "123 Wall St", "department": "Sales", "email": "robert@partner.com"}}}`
	baseClient.JSONBody = response

	client := SalespersonClient{baseClient}

	s, _ := client.GetSalesperson(context.Background(), &r)

	assert.Equal(t, "SALES-ABC123", s.SalespersonID)
	assert.Equal(t, "robert@partner.com", s.Email)
	assert.Equal(t, "New York", s.City)
	assert.Equal(t, "Robert", s.FirstName)
	assert.Equal(t, "10005", s.Zip)
	assert.Equal(t, "Patrick", s.Supervisor)
	assert.Equal(t, "https://lh4.ggpht.com", s.PhotoURLSecure)
	assert.Equal(t, "http://lh4.ggpht.com/6Qq0eNFBmbTFUly", s.PhotoURL)
	assert.Equal(t, "US", s.Country)
	assert.Equal(t, "APlace", s.Region)
	assert.Equal(t, "Salesguy", s.LastName)
	assert.Equal(t, "SRP", s.PartnerID)
	assert.Equal(t, "NY", s.State)
	assert.Equal(t, []string{"5555555555", "5551234567"}, s.PhoneNumber)
	assert.Equal(t, "Seller", s.JobTitle)
	assert.Equal(t, "123 Wall St", s.Address)
	assert.Equal(t, "Sales", s.Department)

}

func Test_GetSalesperson_ReturnsErrorIfExclusiveRequiredFieldNotPassedIn(t *testing.T) {

	baseClient := &basesdk.BaseClientMock{}

	response := `{"data": {"SALES-ABC123": {"city": "New York", "zip": "10005", "supervisor": "Patrick", "photoUrlSecure": "https://lh4.ggpht.com", "firstName": "Robert", "photoUrl": "http://lh4.ggpht.com/6Qq0eNFBmbTFUly", "country": "US", "region": "APlace", "lastName": "Salesguy", "pid": "SRP", "state": "NY", "phoneNumber": ["5555555555", "5551234567"], "salesPersonId": "SALES-ABC123", "jobTitle": "Seller", "address": "123 Wall St", "department": "Sales", "email": "robert@partner.com"}}}`
	baseClient.JSONBody = response

	client := SalespersonClient{baseClient}

	r := GetSalespersonRequest{
		SalespersonID: "SAL123",
	}
	_, err := client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)

	r = GetSalespersonRequest{
		PartnerID:        "ABC",
		SalespersonID:    "SAL123",
		AccountGroupID:   "AG-123",
		SalespersonEmail: "email@email.com",
	}
	_, err = client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)

	r = GetSalespersonRequest{
		PartnerID:      "ABC",
		SalespersonID:  "SAL123",
		AccountGroupID: "AG-123",
	}
	_, err = client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)

	r = GetSalespersonRequest{
		PartnerID:        "ABC",
		AccountGroupID:   "AG-123",
		SalespersonEmail: "email@email.com",
	}
	_, err = client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)

	r = GetSalespersonRequest{
		PartnerID:        "ABC",
		SalespersonID:    "SAL123",
		SalespersonEmail: "email@email.com",
	}
	_, err = client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)

}

func Test_GetSalesperson_ReturnsErrorIfPostReturnsError(t *testing.T) {
	r := GetSalespersonRequest{
		PartnerID:     "ABC",
		SalespersonID: "SAL123",
	}
	baseClient := &basesdk.BaseClientMock{
		Error:    errors.New("Something went wrong"),
		JSONBody: `{"data": {"SALES-ABC123": {"city": "New York", "zip": "10005", "supervisor": "Patrick", "photoUrlSecure": "https://lh4.ggpht.com", "firstName": "Robert", "photoUrl": "http://lh4.ggpht.com/6Qq0eNFBmbTFUly", "country": "US", "region": "APlace", "lastName": "Salesguy", "pid": "SRP", "state": "NY", "phoneNumber": ["5555555555", "5551234567"], "salesPersonId": "SALES-ABC123", "jobTitle": "Seller", "address": "123 Wall St", "department": "Sales", "email": "robert@partner.com"}}}`,
	}

	client := SalespersonClient{baseClient}

	_, err := client.GetSalesperson(context.Background(), &r)
	assert.Error(t, err)
}

func Test_RoundRobin_ReturnsSalespersonId(t *testing.T) {
	r := RoundRobinRequest{
		PartnerID: "ABC",
	}
	baseClient := &basesdk.BaseClientMock{
		JSONBody: `{"data": {"salespersonId": "SP-123"}}`,
	}

	client := SalespersonClient{baseClient}

	salespersonID, _ := client.RoundRobin(context.Background(), &r)

	assert.Equal(t, "SP-123", salespersonID)
}

func Test_RoundRobin_ReturnsErrorIfPartnerIDNotPassed(t *testing.T) {
	r := RoundRobinRequest{
		MarketID: "market-1",
	}
	baseClient := &basesdk.BaseClientMock{}
	client := SalespersonClient{baseClient}

	_, err := client.RoundRobin(context.Background(), &r)

	assert.Error(t, err)
}
