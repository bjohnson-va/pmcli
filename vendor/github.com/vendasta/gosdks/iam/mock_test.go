package iam

import "testing"

func TestMockClientMatchesInterface(t *testing.T) {
	var iamClient Interface
	iamClient = &MockInterface{}
	_ = iamClient // satisfy compiler
}

func TestMockAuthServiceMatchesInterface(t *testing.T) {
	var authService AuthService
	authService = &MockAuthService{}
	_ = authService // satisfy compiler
}
