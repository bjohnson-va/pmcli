package group

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/pb/group/v1"
	"strings"
)

const pathSeparator = "|"

// ForeignKeys holds the keys which a group can belong to -- Partner is required
type ForeignKeys struct {
	PartnerID string
	MarketID  string
}

// Path to a group is composed of nodes which are group IDs
type Path struct {
	Nodes []string
}

// String returns a string representation of a Path
func (p Path) String() string {
	return strings.Join(p.Nodes, pathSeparator)
}

// PathFromString returns a Path object based on a string "G-123|G-456"
func PathFromString(s string) Path {
	n := strings.Split(s, pathSeparator)
	return Path{Nodes: n}
}

// Group is a representation of group data not including it's the child groups/members
type Group struct {
	GroupID           string
	ForeignKeys       ForeignKeys
	Path              Path
	Name              string
	MemberType        string
	Namespace         string
	MembershipVersion int64
	Created           time.Time
	Updated           time.Time
}

func fromProto(gp *group_v1.Group) (*Group, error) {
	if gp == nil {
		return nil, nil
	}
	g := Group{
		GroupID: gp.GetGroupId(),
		ForeignKeys: ForeignKeys{
			PartnerID: gp.GetForeignKeys().GetPartnerId(),
			MarketID:  gp.GetForeignKeys().GetMarketId(),
		},
		Path: Path{
			Nodes: gp.GetPath().GetNodes(),
		},
		Name:       gp.GetName(),
		MemberType: gp.GetMemberType(),
		Namespace:  gp.GetNamespace(),
		MembershipVersion:  gp.GetMembershipVersion(),
	}
	if gp.Created != nil {
		var err error
		g.Created, err = ptypes.Timestamp(gp.Created)
		if err != nil {
			return nil, err
		}
	}
	if gp.Updated != nil {
		var err error
		g.Updated, err = ptypes.Timestamp(gp.Updated)
		if err != nil {
			return nil, err
		}
	}
	return &g, nil
}
