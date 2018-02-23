package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/domain"
)

func TestReplaceDomain(t *testing.T) {
	tests := []struct {
		name             string
		url              string
		whiteLabelDomain *domain.Domain
		expected         string
		expectedError    bool
	}{
		{
			name:             "secure url",
			url:              "https://steprep-test-hrd.appspot.com/account/<accountId>/overview/",
			whiteLabelDomain: &domain.Domain{Name: "coding-with-brent.steprep.com", Secure: true},
			expected:         "https://coding-with-brent.steprep.com/account/<accountId>/overview/",
		},
		{
			name:             "insecure url",
			url:              "http://steprep-test-hrd.appspot.com/account/<accountId>/overview/",
			whiteLabelDomain: &domain.Domain{Name: "coding-with-brent.steprep.com", Secure: false},
			expected:         "http://coding-with-brent.steprep.com/account/<accountId>/overview/",
		},
		{
			name:             "missing domain",
			url:              "http://steprep-test-hrd.appspot.com/account/<accountId>/overview/",
			whiteLabelDomain: nil,
			expected:         "",
			expectedError:    true,
		},
		{
			name:             "bad url",
			url:              "%",
			whiteLabelDomain: &domain.Domain{Name: "coding-with-brent.steprep.com", Secure: false},
			expected:         "",
			expectedError:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := domain.ReplaceDomain(tt.url, tt.whiteLabelDomain)
			assert.Equal(t, tt.expected, actual)
			if err != nil && !tt.expectedError {
				t.Errorf("did not expect an error, got: %v", err)
			} else if err == nil && tt.expectedError {
				t.Error("expected an error but got nil")
			}
		})
	}
}
