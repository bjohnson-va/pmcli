package iam_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/iam"
)

func Test_NewResource(t *testing.T) {
	r := iam.NewResource("vbc", "account-group", "AccountGroup", "URL", "audience", []string{"account_group_id"})
	assert.Equal(t, r.AppID, "vbc")
	assert.Equal(t, r.ResourceID, "account-group")
	assert.Equal(t, r.ResourceName, "AccountGroup")
	assert.Equal(t, r.ResourceOwnerServiceURL, "URL")
	assert.Equal(t, r.ResourceOwnerAudience, "audience")
	assert.Equal(t, r.RequiredResourceParams, []string{"account_group_id"})
}
