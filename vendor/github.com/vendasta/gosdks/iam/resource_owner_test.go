package iam_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/iam"
)

func Test_NewResourceOwner(t *testing.T) {
	resourecOwner := iam.NewResourceOwner("vbc", "Business Center")
	assert.Equal(t, "vbc", resourecOwner.AppID)
	assert.Equal(t, "Business Center", resourecOwner.AppName)
}
