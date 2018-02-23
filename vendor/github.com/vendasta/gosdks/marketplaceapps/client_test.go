package marketplaceapps

import (
	"github.com/vendasta/gosdks/pb/marketplace_apps/v1"
	"testing"
)

func TestGrpcClient_buildIntegration(t *testing.T) {
	publicKeys := []*marketplaceapps_v1.IdentifiedPublicKey{
		{"foo", "bar"},
	}

	i := marketplaceapps_v1.Integration{
		"bar",
		publicKeys,
		"test",
		true,
		"test2",
		[]string{"test3"},
		"test4",
		"test5",
		"test6",
		"test7",
		"test8",
		"test9",
		[]string{"test10"},
		"test11",
		"test12",
		"test12",
		"test13",
		"test14",
	}

	result := buildIntegrationData(&i)
	if result.PublicKey != i.PublicKey {
		t.Errorf("Error, expected %v, got %v", i.PublicKey, result.PublicKey)
	}

	if result.IdentifiedPublicKeys[0].PublicKey != i.IdentifiedPublicKeys[0].PublicKey {
		t.Errorf("Error, expected %v, got %v", i.IdentifiedPublicKeys[0].PublicKey, result.IdentifiedPublicKeys[0].PublicKey)
	}

	if result.EntryUrl != i.EntryUrl {
		t.Errorf("Error, expected %v, got %v", i.EntryUrl, result.EntryUrl)
	}

	if result.ReserveIdUrl != i.ReserveIdUrl {
		t.Errorf("Error, expected %v, got %v", i.ReserveIdUrl, result.ReserveIdUrl)
	}
}

func TestGRPCClient_buildRestrictionsData(t *testing.T) {
	whiteList := []string{"ABC", "123"}
	blackList := []string{"ABC", "123"}

	restrictions := &marketplaceapps_v1.Restrictions{
		Country: &marketplaceapps_v1.PermissionLists{
			Whitelist: whiteList,
			Blacklist: blackList,
		},
	}

	r := buildRestrictionsData(restrictions)

	if !SlicesEqual(r.Country.BlackList, blackList) {
		t.Errorf("Error, expected %v, got %v", blackList, r.Country.BlackList)
	}

	if !SlicesEqual(r.Country.WhiteList, whiteList) {
		t.Errorf("Error, expected %v, got %v", whiteList, r.Country.WhiteList)
	}
}

func SlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestGRPCClient_buildOrderForm(t *testing.T) {
	orderFormFields := make([]*marketplaceapps_v1.OrderFormField, 1)
	orderFormFields[0] = &marketplaceapps_v1.OrderFormField{
		Label:       "Test",
		Id:          "Test1",
		Type:        "Test2",
		Options:     []string{"Test3"},
		Description: "Test4",
		Required:    true,
		UploadUrl:   "www.example.com",
	}
	of := &marketplaceapps_v1.OrderForm{
		OrderForm: orderFormFields,
		CommonForm: &marketplaceapps_v1.IncludedCommonFormFields{
			BusinessAccountGroupId: false,
			BusinessAddress:        false,
			BusinessName:           false,
			BusinessPhoneNumber:    false,
			ContactEmail:           false,
			ContactName:            false,
			ContactPhoneNumber:     false,
			SalespersonEmail:       false,
			SalespersonName:        false,
			SalespersonPhoneNumber: false,
		},
		ActivationMessage: "AAAA",
	}

	r := buildOrderFormData(of)

	if r.ActivationMessage != "AAAA" {
		t.Errorf("Error, expected %v, got %v", of.ActivationMessage, r.ActivationMessage)
	}

	if r.OrderForm[0].Required != of.OrderForm[0].Required {
		t.Errorf("Error, expected %v, got %v", of.OrderForm[0].Required, r.OrderForm[0].Required)
	}

	if r.OrderForm[0].Label != of.OrderForm[0].Label {
		t.Errorf("Error, expected %v, got %v", of.OrderForm[0].Label, r.OrderForm[0].Label)
	}

	if r.OrderForm[0].Id != of.OrderForm[0].Id {
		t.Errorf("Error, expected %v, got %v", of.OrderForm[0].Id, r.OrderForm[0].Id)
	}

	if r.OrderForm[0].UploadUrl != of.OrderForm[0].UploadUrl {
		t.Errorf("Error, expected %v, got %v", of.OrderForm[0].UploadUrl, r.OrderForm[0].UploadUrl)
	}
}

func TestGRPCClient_buildMarketingInformation(t *testing.T) {
	m := &marketplaceapps_v1.MarketingInformation{
		Description:      "reseller info",
		KeySellingPoints: []string{"Its cool", "Its Great"},
		Faqs: []*marketplaceapps_v1.FrequentlyAskedQuestions{
			{Question: "this is a question", Answer: "this is an answer"},
		},
		Files: []string{"YAY"},
	}

	r := buildMarketingInformation(m)

	if r.Description != m.Description {
		t.Errorf("Error, expected %v, got %v", m.Description, r.Description)
	}

	if r.KeySellingPoints[0] != m.KeySellingPoints[0] {
		t.Errorf("Error, expected %v, got %v", m.KeySellingPoints[0], r.KeySellingPoints[0])
	}
}
