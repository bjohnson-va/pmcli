package nap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountryCodeFromCountryString_ReturnsCountryCodeForAlpha2(t *testing.T) {
	code := GetCountryCodeFromCountryString("US")
	assert.Equal(t, "US", code)
}

func TestGetCountryCodeFromCountryString_ReturnsCountryCodeForName(t *testing.T) {
	code := GetCountryCodeFromCountryString("United States")
	assert.Equal(t, "US", code)
}

func TestGetCountryCodeFromCountryString_ReturnsCountryCodeForAllCaps(t *testing.T) {
	code := GetCountryCodeFromCountryString("UNITED STATES")
	assert.Equal(t, "US", code)
}

func TestGetCountryCodeFromCountryString_ReturnsCountryCodeForUnicodeName(t *testing.T) {
	code := GetCountryCodeFromCountryString("Saint Barthélemy")
	assert.Equal(t, "BL", code)
}

func TestGetCountryCodeFromCountryString_ReturnsCountryCodeForAlpha3(t *testing.T) {
	code := GetCountryCodeFromCountryString("USA")
	assert.Equal(t, "US", code)
}

func TestGetCountryCodeFromCountryString_ReturnsEmptyForNonCountry(t *testing.T) {
	code := GetCountryCodeFromCountryString("Garbage")
	assert.Equal(t, "", code)
}

func TestGetStateCodeFromStateString_ReturnsStateCodeForCountryWithThatStateCode(t *testing.T) {
	code := GetStateCodeFromStateString("CA", "SK")
	assert.Equal(t, "SK", code)
}

func TestGetStateCodeFromStateString_ReturnsStateCodeForCountryWithThatStateName(t *testing.T) {
	code := GetStateCodeFromStateString("CA", "Saskatchewan")
	assert.Equal(t, "SK", code)
}

func TestGetStateCodeFromStateString_ReturnsStateCodeForCountryWithThatStateNameAllCaps(t *testing.T) {
	code := GetStateCodeFromStateString("CA", "SASKATCHEWAN")
	assert.Equal(t, "SK", code)
}


func TestGetStateCodeFromStateString_ReturnsEmptyForCountryWithoutThatState(t *testing.T) {
	code := GetStateCodeFromStateString("CA", "Garbage")
	assert.Equal(t, "", code)
}

func TestGetStateCodeFromStateString_ReturnsEmptyForNonCountry(t *testing.T) {
	code := GetStateCodeFromStateString("Garbage", "Fake")
	assert.Equal(t, "", code)
}
