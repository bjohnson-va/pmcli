package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToSMB(c *subjectcontext.Context, s *subject) *SMB {
	smb := &SMB{
		subject:   s,
		PartnerID: c.Namespace,
	}

	return smb
}

// AccountAccessPermission defines which marketplace apps the smb has access to on an account group
type AccountAccessPermission struct {
	AccountGroupID string   `attribute:"account_group_id"`
	AccountIDs     []string `attribute:"account_ids"`
}

// SMB represents a VBC user.
type SMB struct {
	*subject

	PartnerID string

	FirstName                string                     `attribute:"first_name"`
	LastName                 string                     `attribute:"last_name"`
	WorkPhone                string                     `attribute:"work_phone"`
	DefaultAccountGroup      string                     `attribute:"default_account_group"`
	AccountGroupAssociations []string                   `attribute:"account_group_associations"`
	AccountAccessPermissions []*AccountAccessPermission `attribute:"account_access_permissions"`
}

func (s *SMB) Context() *subjectcontext.Context {
	return subjectcontext.New("smb", s.PartnerID)
}

// NewSMB creates a new SMB struct
func NewSMB(s *subject, PartnerID string, FirstName, LastName, WorkPhone, DefaultAccountGroup string,
	AccountGroupAssociations []string, AccountAccessPermissions []*AccountAccessPermission) *SMB {
	return &SMB{
		subject:                  s,
		PartnerID:                PartnerID,
		FirstName:                FirstName,
		LastName:                 LastName,
		WorkPhone:                WorkPhone,
		DefaultAccountGroup:      DefaultAccountGroup,
		AccountGroupAssociations: AccountGroupAssociations,
		AccountAccessPermissions: AccountAccessPermissions,
	}
}
