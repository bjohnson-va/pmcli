package config

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetEnv(t *testing.T) {
	var tests = []struct {
		in       string
		expected Env
	}{
		{"local", Local},
		{"prod", Prod},
		{"production", Prod},
		{"test", Test},
		{"demo", Demo},
		{"internal", Internal},
	}

	for _, tt := range tests {
		actual := GetEnv(tt.in)
		assert.Equal(t, tt.expected, actual)
	}
}
