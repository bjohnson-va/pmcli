package cssdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSocialServiceIDToServiceType(t *testing.T) {
	type test struct {
		ssid     string
		expected ServiceType
	}

	cases := []*test{
		{
			ssid:     "FBP-XXX",
			expected: ServiceTypeFacebookPage,
		},
		{
			ssid:     "FBU-XXX",
			expected: ServiceTypeFacebookUser,
		},
		{
			ssid:     "TWU-XXX",
			expected: ServiceTypeTwitterUser,
		},
		{
			ssid:     "GPP-XXX",
			expected: ServiceTypeGooglePlusPage,
		},
		{
			ssid:     "GPU-XXX",
			expected: ServiceTypeGooglePlusUser,
		},
		{
			ssid:     "LIU-XXX",
			expected: ServiceTypeLinkedinUser,
		},
		{
			ssid:     "LIC-XXX",
			expected: ServiceTypeLinkedinCompany,
		},
		{
			ssid:     "NOT A VALID SSID",
			expected: "",
		},
	}

	for _, c := range cases {
		result := SocialServiceIDToServiceType(c.ssid)
		assert.Equal(t, result, c.expected)
	}
}
