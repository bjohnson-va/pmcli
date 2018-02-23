package nap

import "strings"

// GetCountryCodeFromCountryString returns the ISO-3166-1 alpha2 code, if one can be found, from the given a string that may contain the countries alpha2, alpha3, or name
func GetCountryCodeFromCountryString(country string) string {
	countries := GetCountries()
	_, ok := countries[country]
	if ok {
		return country
	}
	for _, c := range countries {
		if c.Alpha3 == country || (strings.ToLower(c.Name) == strings.ToLower(country)) {
			return c.Alpha2
		}
	}
	return ""
}

// GetStateCodeFromStateString returns the ISO-3166-2 state code, if one can be found, from the countryCode and state string that may contain the state's code or name
func GetStateCodeFromStateString(countryCode string, state string) string {
	states := GetStatesForCountry(countryCode)
	if states == nil {
		return ""
	}
	_, ok := states[state]
	if ok {
		return state
	}
	for stateCode, s := range states {
		if strings.ToLower(s.Name) == strings.ToLower(state) {
			return stateCode
		}
	}
	return ""
}
