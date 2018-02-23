package group

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/group/v1"
)

func Test_fromProto(t *testing.T) {
	now := time.Now().UTC()
	nowProto, _ := ptypes.TimestampProto(now)
	testGroup := &group_v1.Group{
		GroupId: "G-456",
		ForeignKeys: &group_v1.ForeignKeys{
			PartnerId: "ABC",
			MarketId:  "market-1",
		},
		Path: &group_v1.Path{
			Nodes: []string{"G-123", "G-456"},
		},
		Name:       "My Group",
		MemberType: "account-group",
		MembershipVersion: 4,
		Namespace:  "brands",
		Created:    nowProto,
		Updated:    nowProto,
	}

	g, _ := fromProto(testGroup)
	assert.Equal(t, "G-456", g.GroupID)
	assert.Equal(t, "ABC", g.ForeignKeys.PartnerID)
	assert.Equal(t, "market-1", g.ForeignKeys.MarketID)
	assert.Equal(t, []string{"G-123", "G-456"}, g.Path.Nodes)
	assert.Equal(t, "My Group", g.Name)
	assert.Equal(t, "account-group", g.MemberType)
	assert.Equal(t, "brands", g.Namespace)
	assert.Equal(t, int64(4), g.MembershipVersion)
	assert.Equal(t, now, g.Created)
	assert.Equal(t, now, g.Updated)
}
