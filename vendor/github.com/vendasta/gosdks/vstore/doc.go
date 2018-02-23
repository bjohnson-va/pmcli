package vstore

/*
vStore client provides methods for creating/update/looking up entities from vStore.

Example:

package accountgroup

type AccountGroup struct {
	AccountGroupId `vstore:"account_group_id"`
}

func (a *AccountGroup) KeySet() *vstore.KeySet {
	return AccountGroupKeySet(a.AccountGroupId)
}

func AccountGroupKeySet(accountGroupId string) *vstore.KeySet {
	return vstore.NewKeySet("vbc, "AccountGroup", []string{accountGroupId})
}

vstoreClient := vstore.New(vstore.Prod)

func Create(ctx context.Context, accountGroupId string) (*AccountGroup, error) {
	ks :=  AccountGroupKeySet("AG-123")
	var accountGroup *AccountGroup
	err := vstoreClient.Transaction(ctx, ks, f func(tx Transaction, m vstore.Model) error {
		if m != nil {
			return errors.New("Account Group Already Exists")
		}
		accountGroup = &AccountGroup{AccountGroupId: accountGroupId}
		return tx.Save(account_group)
	})
	if err != nil {
		return nil, err
	}
	return accountGroup, nil
}

func UpdateCompanyName(ctx context.Context, accountGroupId string, companyName string) (*AccountGroup, error) {
	ks := AccountGroupKeySet("AG-123")
	var accountGroup *AccountGroup
	err := vstoreClient.Transaction(ctx, ks, f func(tx Transaction, m vstore.Model) error {
		if m == nil {
			return errors.New("Account Group doesnt exist")
		}
		accountGroup = m.(*AccountGroup)
		accountGroup.CompanyName = companyName
		return tx.Save(account_group)
	})
	if err != nil {
		return nil, err
	}
	return accountGroup, nil
}

 */
