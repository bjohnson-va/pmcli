package accountgroup

import (
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/account_group/v1"
)

func Test_AccountGroupFromPB(t *testing.T) {
	// arrange
	agData := `{"accountGroupId":"AG-Z8Z4TT5M","created":"2017-04-27T16:50:35.617910Z","updated":"2017-04-27T17:13:14.591700Z","version":"5","accounts":{"accounts":[{"marketplaceAppId":"RM","accountId":"K2S2XFWG"},{"marketplaceAppId":"MS","accountId":"MS-HXDCM6G6"}]},"accountGroupExternalIdentifiers":{"origin":"api-product","customerIdentifier":"K2S2XFWG","socialProfileId":"SCP-6EADC1A5B8BF422FA7F98F4012EE747F","partnerId":"SRP","taxIds":["other"]},"socialUrls":{},"hoursOfOperation": {
		"hoursOfOperation": [{
				"dayOfWeek": [
					"Monday"
				],
				"opens": "9:00",
				"closes": "17:00",
				"description": "Open weekdays"
			}
		]
	},"contactDetails":{},"snapshotReports":{"snapshots":[{"snapshotId":"SNAPSHOT-869ec1d21a034462b088bababa1cf956","created":"2017-04-27T16:50:42.280970Z","expiry":"2017-05-04T16:50:42.280970Z"}]},"legacyProductDetails":{},"richData":{},"napData":{"companyName":"Dale's House","address":"2119 Lansdowne","city":"Saskatoon","state":"SK","zip":"S7J1G6","country":"CA","location":{"latitude":52.1044583,"longitude":-106.6509008},"timezone":"America/Regina"},"status":{}}`
	ag := new(accountgroup_v1.AccountGroup)
	err := jsonpb.UnmarshalString(agData, ag)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// act
	accountGroup, err := FromPB(ag, &accountgroup_v1.ProjectionFilter{
		NapData:                         true,
		Accounts:                        true,
		ListingDistribution:             true,
		ListingSyncPro:                  true,
		Associations:                    true,
		AccountGroupExternalIdentifiers: true,
		SocialUrls:                      true,
		HoursOfOperation:                true,
		ContactDetails:                  true,
		SnapshotReports:                 true,
		LegacyProductDetails:            true,
		RichData:                        true,
		Status:                          true,
	})

	// assert
	assert.NotNil(t, accountGroup)
	assert.Nil(t, err)
	assert.Equal(t, "AG-Z8Z4TT5M", accountGroup.AccountGroupID)

	assert.Equal(t, "Dale's House", accountGroup.NAPData.CompanyName)
	assert.Equal(t, "2119 Lansdowne", accountGroup.NAPData.Address)
	assert.Equal(t, "Saskatoon", accountGroup.NAPData.City)
	assert.Equal(t, "SK", accountGroup.NAPData.State)
	assert.Equal(t, "S7J1G6", accountGroup.NAPData.Zip)
	assert.Equal(t, "America/Regina", accountGroup.NAPData.Timezone)
	assert.Equal(t, 52.1044583, accountGroup.NAPData.Location.Latitude)
	assert.Equal(t, -106.6509008, accountGroup.NAPData.Location.Longitude)
	assert.Equal(t, "", accountGroup.NAPData.Website)
	assert.Equal(t, 0, len(accountGroup.NAPData.WorkNumber))
	assert.Equal(t, 0, len(accountGroup.NAPData.CallTrackingNumber))

	assert.Equal(t, 2, len(accountGroup.Accounts))
	assert.True(t, accountGroup.Accounts[0].Expiry.IsZero())
	assert.Equal(t, ([]string)(nil), accountGroup.Accounts[0].Tags)
	assert.Equal(t, "K2S2XFWG", accountGroup.Accounts[0].AccountID)
	assert.Equal(t, "RM", accountGroup.Accounts[0].MarketplaceAppID)
	assert.Equal(t, false, accountGroup.Accounts[0].IsTrial)

	assert.True(t, accountGroup.Accounts[1].Expiry.IsZero())
	assert.Equal(t, ([]string)(nil), accountGroup.Accounts[1].Tags)
	assert.Equal(t, "MS-HXDCM6G6", accountGroup.Accounts[1].AccountID)
	assert.Equal(t, "MS", accountGroup.Accounts[1].MarketplaceAppID)
	assert.Equal(t, false, accountGroup.Accounts[1].IsTrial)

	assert.Nil(t, accountGroup.ListingDistribution)
	assert.Nil(t, accountGroup.ListingSyncPro)
	assert.Nil(t, accountGroup.Associations)

	assert.Equal(t, "SRP", accountGroup.ExternalIdentifiers.PartnerID)
	assert.Equal(t, "", accountGroup.ExternalIdentifiers.MarketID)
	assert.Equal(t, []string{"other"}, accountGroup.ExternalIdentifiers.TaxIDs)
	assert.Equal(t, "", accountGroup.ExternalIdentifiers.SalesPersonID)
	assert.Equal(t, ([]string)(nil), accountGroup.ExternalIdentifiers.Tags)
	assert.Equal(t, "api-product", accountGroup.ExternalIdentifiers.Origin)
	assert.Equal(t, ([]string)(nil), accountGroup.ExternalIdentifiers.JobID)
	assert.Equal(t, "K2S2XFWG", accountGroup.ExternalIdentifiers.CustomerIdentifier)
	assert.Equal(t, ([]string)(nil), accountGroup.ExternalIdentifiers.ActionLists)
	assert.Equal(t, "SCP-6EADC1A5B8BF422FA7F98F4012EE747F", accountGroup.ExternalIdentifiers.SocialProfileID)

	hooSpan := []*Span{
		&Span{
			DayOfWeek:   []string{"Monday"},
			Opens:       "9:00",
			Closes:      "17:00",
			Description: "Open weekdays",
		},
	}
	assert.Equal(t, hooSpan, accountGroup.HoursOfOperation.HoursOfOperation)

}
