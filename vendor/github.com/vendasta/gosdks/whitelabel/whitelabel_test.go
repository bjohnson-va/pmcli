package whitelabel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"golang.org/x/net/context"
)

func Test_GetWhiteLabelDataReturnsErrorIfNoPartnerIDIsSupplied(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{}
	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.get(context.Background(), nil, "", "", nil)
	assert.EqualError(t, err, "partnerID is required")
}

func Test_GetWhiteLabelDataReturnsErrorIfGetReturnsError(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: &basesdk.HTTPError{Body: "Get Failed", StatusCode: 400}}
	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "")
	assert.EqualError(t, err, "Get Failed")
}

func Test_GetWhiteLabelDataReturnsPartnerNotFoundErrorWhen404FromServiceWhenTryingToGetPartnerWhiteLabelData(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: &basesdk.HTTPError{Body: "Get Failed", StatusCode: 404}}
	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "")
	assert.Equal(t, err, ErrorPartnerNotFound)
}

func Test_GetWhiteLabelDataReturnsMarketNotFoundErrorWhen404FromServiceWhenTryingToGetMarketWhiteLabelData(t *testing.T) {
	baseClient := &basesdk.BaseClientMultiCallMock{
		Mocks: []*basesdk.BaseClientMock{
			&basesdk.BaseClientMock{JSONBody: `{"data": "{}"}`},
			&basesdk.BaseClientMock{Error: &basesdk.HTTPError{Body: "Get Failed", StatusCode: 404}},
		},
	}

	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "Market")
	assert.Equal(t, ErrorMarketNotFound, err)
}

func Test_GetWhiteLabelDataReturnsCorrectData(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{JSONBody: `{"version":"3.0","data":"{\"ms_always_enabled\": true, \"rm_exit_link_url\": null, \"sender_name\": null, \"st_use_custom_snapshot_header\": false, \"st_show_training_resources_field\": true, \"ui_theme\": \"dark\", \"favicon_url\": \"http://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32\", \"account_executive_email\": \"bdong+lasthope@vendasta.com\", \"suppress_state_province_field\": false, \"st_use_custom_snapshot_popup_message\": false, \"rm_send_executive_reports_flag\": false, \"st_custom_snapshot_footer\": null, \"has_ms\": true, \"has_rm\": true, \"exit_link_text\": \"\\u200bReturn to Business Center\", \"sm_display_facebook_account\": true, \"vbc_invitation_campaign_id\": \"PROSPECT-USER-ACCESS-PREMADE-VUNI\", \"st_allow_salesperson_campaign_add\": true, \"notifications_enabled\": true, \"social_profile_group_id\": \"SPG-9B63B9C33D4C42468F5D68D99A896E28\", \"listing_sync_pro_superadmin_enabled\": true, \"ld_show_expiry\": true, \"business_profile_editable\": true, \"st_show_seo_in_snapshot\": true, \"st_show_products_wholesale_price\": true, \"name\": \"#Vendasta Training & Testing\", \"st_use_custom_snapshot_footer\": false, \"has_sm\": true, \"ms_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"listing_sync_pro_enabled\": true, \"rm_competition_max_competitors\": 3, \"sm_exit_link_text\": \"\\u200bReturn to Business Center\", \"campaign_subscription_enabled\": true, \"listing_sync_pro_sell_price_monthly\": null, \"send_followup_email_flag\": true, \"rm_exit_link_text\": \"\\u200bReturn to Business Center\", \"st_bright_local_key\": \"14b455c69e4ffe7d5a770f3feb75e5d636120bcb\", \"ld_enabled\": true, \"partner_id\": \"VUNI\", \"st_salesperson_shared_account_access\": true, \"ms_exit_link_url\": null, \"ms_enabled\": true, \"nb_multi_location_posting_requires_approval\": true, \"nb_allow_export_reviews\": true, \"has_vbc\": true, \"beta_flag\": true, \"mobile_shortcut_icon_secure_url\": \"https://securemobileicon.com\", \"st_remarketing\": true, \"locked_fields\": [], \"ms_product_name\": \"Listing Builder\", \"rm_competition_max_services\": 3, \"listing_sync_pro_sell_price_annual\": null, \"nb_allow_multi_location_posting\": true, \"listing_sync_pro_show_admin_price\": false, \"sm_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"st_team_activity\": true, \"listing_sync_pro_service_provider\": \"Uberall\", \"favicon_secure_url\": \"https://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32\", \"sender_email\": null, \"logo_secure_url\": \"https://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_\", \"rm_listing_point_score_flag\": true, \"mobile_shortcut_icon_url\": \"http://mobileicon.com\", \"account_executive_name\": \"Bo the last hope Dong\", \"vbc_show_content_library\": true, \"rm_display_view_review_link\": true, \"sm_exit_link_url\": null, \"market_id\": null, \"st_ppc\": true, \"nb_product_name\": \"Brands\", \"sm_product_name\": \"Social Marketing\", \"ld_wholesale_price\": 50.0, \"vbc_product_name\": \"Business Center\", \"listing_sources_editable\": true, \"primary_color\": \"3f9b63\", \"review_generation_enabled\": true, \"subscription_level\": \"BASIC\", \"st_custom_snapshot_popup_message\": null, \"preferred_sites_review_gen_workflow_enabled\": true, \"ms_exit_link_text\": \"\\u200bReturn to Business Center\", \"ld_enabled_for_partner_admins\": true, \"vbc_show_training_resources_field\": true, \"st_custom_snapshot_header\": \"Paste your code here...\", \"listing_sync_pro_show_renew_price\": true, \"listing_sync_pro_show_sell_price\": false, \"rm_executive_report_frequency\": \"Weekly\", \"ms_overview_page_enabled\": true, \"ld_partner_billing_allowed\": false, \"st_default_campaign_id\": null, \"st_allow_claim_user\": true, \"rm_product_name\": \"Reputation Management\", \"listing_sync_pro_discount_flag\": false, \"logo_url\": \"http://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_\", \"ms_location_page_enabled\": true, \"_type\": \"Configuration\", \"ld_retail_price\": 61.229999999999997, \"business_directory_enabled\": true, \"default_acquire_campaign_id\": \"CAMPAIGN-DEFAULT-DRIP-PREMADE-VUNI\", \"rm_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"vbc_custom_introduction_message\": null, \"rm_listing_point_score_enabled\": null, \"ld_enabled_for_account_user\": true, \"vbc_use_custom_introduction_message\": false, \"exit_link_url\": null, \"st_send_welcome_email\": true, \"sidebar_enabled\": true, \"st_show_bright_local_section_in_snapshot\": true, \"_object_version\": \"2.16.0\"}","requestId":"594a8c4100ff00ff7209914c6f210001737e706172746e65722d63656e7472616c0001617069733a313137352d67626c61612d3333383900010108","responseTime":238,"statusCode":200}`}
	client := whiteLabelClient{SDKClient: baseClient}
	s, err := client.Get(context.Background(), "ABC", "")
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, int64(3), s.MaxRMCompetitors)
	assert.Equal(t, int64(3), s.MaxRMServices)
	assert.Equal(t, "SPG-9B63B9C33D4C42468F5D68D99A896E28", s.SocialProfileGroupID)
	assert.Equal(t, "dark", s.UITheme)
	assert.Equal(t, "3f9b63", s.PrimaryColor)
	assert.Equal(t, "https://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32", s.FaviconSecureURL)
	assert.Equal(t, "http://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32", s.FaviconURL)
	assert.Equal(t, "https://securemobileicon.com", s.MobileShortcutIconSecureURL)
	assert.Equal(t, "http://mobileicon.com", s.MobileShortcutIconURL)
	assert.Equal(t, "https://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_", s.LogoSecureURL)
	assert.Equal(t, "http://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_", s.LogoURL)
	assert.Equal(t, "Business Center", s.BusinessCenterProductName)
	assert.Equal(t, "Reputation Management", s.ReputationManagementProductName)
	assert.Equal(t, "Listing Builder", s.ListingBuilderProductName)
	assert.Equal(t, "Social Marketing", s.SocialMarketingProductName)
	assert.Equal(t, "Brands", s.BrandAnalyticsProductName)
}

func Test_GetWhiteLabelDataReturnsASpecificInvalidMarketErrorMessageIfTheInvalidAgumentIsMarketID(t *testing.T) {
	badMarketIDMock := &basesdk.BaseClientMock{
		JSONBody: "",
		Error: &basesdk.HTTPError{
			Body: `{
				"message": "Arg \"marketId\": Market Id GARBAGE is not valid. ",
				"version": "3.0",
				"requestId": "59bad2ff0000ff0e50faf54510a90001737e706172746e65722d63656e7472616c2d746573740001617069733a636f6e74696e756f75730001010a",
				"responseTime": 6,
				"statusCode": 400
			  }`,
			StatusCode: 400,
		},
	}
	baseClient := &basesdk.BaseClientMultiCallMock{
		Mocks: []*basesdk.BaseClientMock{
			&basesdk.BaseClientMock{JSONBody: `{"data": "{}"}`},
			badMarketIDMock,
		},
	}
	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "GARBAGE_MARKET_ID")
	assert.EqualError(t, err, "the market Id: GARBAGE_MARKET_ID is not valid")
}

func Test_GetWhiteLabelDataReturnsCorrectDataIfUnmergedOptionIsPassedIn(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{JSONBody: `{"version":"3.0","data":"{\"ms_always_enabled\": true, \"rm_exit_link_url\": null, \"sender_name\": null, \"st_use_custom_snapshot_header\": false, \"st_show_training_resources_field\": true, \"ui_theme\": \"dark\", \"favicon_url\": \"http://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32\", \"account_executive_email\": \"bdong+lasthope@vendasta.com\", \"suppress_state_province_field\": false, \"st_use_custom_snapshot_popup_message\": false, \"rm_send_executive_reports_flag\": false, \"st_custom_snapshot_footer\": null, \"has_ms\": true, \"has_rm\": true, \"exit_link_text\": \"\\u200bReturn to Business Center\", \"sm_display_facebook_account\": true, \"vbc_invitation_campaign_id\": \"PROSPECT-USER-ACCESS-PREMADE-VUNI\", \"st_allow_salesperson_campaign_add\": true, \"notifications_enabled\": true, \"social_profile_group_id\": \"SPG-9B63B9C33D4C42468F5D68D99A896E28\", \"listing_sync_pro_superadmin_enabled\": true, \"ld_show_expiry\": true, \"business_profile_editable\": true, \"st_show_seo_in_snapshot\": true, \"st_show_products_wholesale_price\": true, \"name\": \"#Vendasta Training & Testing\", \"st_use_custom_snapshot_footer\": false, \"has_sm\": true, \"ms_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"listing_sync_pro_enabled\": true, \"rm_competition_max_competitors\": 3, \"sm_exit_link_text\": \"\\u200bReturn to Business Center\", \"campaign_subscription_enabled\": true, \"listing_sync_pro_sell_price_monthly\": null, \"send_followup_email_flag\": true, \"rm_exit_link_text\": \"\\u200bReturn to Business Center\", \"st_bright_local_key\": \"14b455c69e4ffe7d5a770f3feb75e5d636120bcb\", \"ld_enabled\": true, \"partner_id\": \"VUNI\", \"st_salesperson_shared_account_access\": true, \"ms_exit_link_url\": null, \"ms_enabled\": true, \"nb_multi_location_posting_requires_approval\": true, \"nb_allow_export_reviews\": true, \"has_vbc\": true, \"beta_flag\": true, \"mobile_shortcut_icon_secure_url\": \"https://securemobileicon.com\", \"st_remarketing\": true, \"locked_fields\": [], \"ms_product_name\": \"Listing Builder\", \"rm_competition_max_services\": 3, \"listing_sync_pro_sell_price_annual\": null, \"nb_allow_multi_location_posting\": true, \"listing_sync_pro_show_admin_price\": false, \"sm_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"st_team_activity\": true, \"listing_sync_pro_service_provider\": \"Uberall\", \"favicon_secure_url\": \"https://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32\", \"sender_email\": null, \"logo_secure_url\": \"https://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_\", \"rm_listing_point_score_flag\": true, \"mobile_shortcut_icon_url\": \"http://mobileicon.com\", \"account_executive_name\": \"Bo the last hope Dong\", \"vbc_show_content_library\": true, \"rm_display_view_review_link\": true, \"sm_exit_link_url\": null, \"market_id\": null, \"st_ppc\": true, \"nb_product_name\": \"Brands\", \"sm_product_name\": \"Social Marketing\", \"ld_wholesale_price\": 50.0, \"vbc_product_name\": \"Business Center\", \"listing_sources_editable\": true, \"primary_color\": \"3f9b63\", \"review_generation_enabled\": true, \"subscription_level\": \"BASIC\", \"st_custom_snapshot_popup_message\": null, \"preferred_sites_review_gen_workflow_enabled\": true, \"ms_exit_link_text\": \"\\u200bReturn to Business Center\", \"ld_enabled_for_partner_admins\": true, \"vbc_show_training_resources_field\": true, \"st_custom_snapshot_header\": \"Paste your code here...\", \"listing_sync_pro_show_renew_price\": true, \"listing_sync_pro_show_sell_price\": false, \"rm_executive_report_frequency\": \"Weekly\", \"ms_overview_page_enabled\": true, \"ld_partner_billing_allowed\": false, \"st_default_campaign_id\": null, \"st_allow_claim_user\": true, \"rm_product_name\": \"Reputation Management\", \"listing_sync_pro_discount_flag\": false, \"logo_url\": \"http://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_\", \"ms_location_page_enabled\": true, \"_type\": \"Configuration\", \"ld_retail_price\": 61.229999999999997, \"business_directory_enabled\": true, \"default_acquire_campaign_id\": \"CAMPAIGN-DEFAULT-DRIP-PREMADE-VUNI\", \"rm_api_key\": \"qUkF1TmQ0xcyk5nPHrWWHNUMYOf13L1P1JCTezkL\", \"vbc_custom_introduction_message\": null, \"rm_listing_point_score_enabled\": null, \"ld_enabled_for_account_user\": true, \"vbc_use_custom_introduction_message\": false, \"exit_link_url\": null, \"st_send_welcome_email\": true, \"sidebar_enabled\": true, \"st_show_bright_local_section_in_snapshot\": true, \"_object_version\": \"2.16.0\"}","requestId":"594a8c4100ff00ff7209914c6f210001737e706172746e65722d63656e7472616c0001617069733a313137352d67626c61612d3333383900010108","responseTime":238,"statusCode":200}`}
	client := whiteLabelClient{SDKClient: baseClient}
	s, err := client.Get(context.Background(), "ABC", "", Unmerged())
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, int64(3), s.MaxRMCompetitors)
	assert.Equal(t, int64(3), s.MaxRMServices)
	assert.Equal(t, "SPG-9B63B9C33D4C42468F5D68D99A896E28", s.SocialProfileGroupID)
	assert.Equal(t, "dark", s.UITheme)
	assert.Equal(t, "3f9b63", s.PrimaryColor)
	assert.Equal(t, "https://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32", s.FaviconSecureURL)
	assert.Equal(t, "http://lh3.googleusercontent.com/tiMilqtqDCXL5JclYdck1JjIfUVOM-IjIoUy97qRo5so3z7j5pEN2h56AQFqatRsZGm2GWY6-NAnEo27N_ukKxSp-gA=s32", s.FaviconURL)
	assert.Equal(t, "https://securemobileicon.com", s.MobileShortcutIconSecureURL)
	assert.Equal(t, "http://mobileicon.com", s.MobileShortcutIconURL)
	assert.Equal(t, "https://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_", s.LogoSecureURL)
	assert.Equal(t, "http://lh3.googleusercontent.com/OB4j6gUPwm_RvqXbruxFxAePmxOoJCemCTkazpBwhk8Q-jG4fBmwMxF8O6cIfN_-cjVQBiGffxx-OXWz4uW5RPyAiPG_", s.LogoURL)
	assert.Equal(t, "Business Center", s.BusinessCenterProductName)
	assert.Equal(t, "Reputation Management", s.ReputationManagementProductName)
	assert.Equal(t, "Listing Builder", s.ListingBuilderProductName)
	assert.Equal(t, "Social Marketing", s.SocialMarketingProductName)
	assert.Equal(t, "Brands", s.BrandAnalyticsProductName)
}

func Test_GetWhiteLabelDataReturnsASpecificInvalidMarketErrorMessageIfTheInvalidAgumentIsMarketIDAndUnmergedIsPassedIn(t *testing.T) {
	badMarketIDMock := &basesdk.BaseClientMock{
		JSONBody: "",
		Error: &basesdk.HTTPError{
			Body: `{
				"message": "Arg \"marketId\": Market Id GARBAGE is not valid. ",
				"version": "3.0",
				"requestId": "59bad2ff0000ff0e50faf54510a90001737e706172746e65722d63656e7472616c2d746573740001617069733a636f6e74696e756f75730001010a",
				"responseTime": 6,
				"statusCode": 400
			  }`,
			StatusCode: 400,
		},
	}
	baseClient := &basesdk.BaseClientMultiCallMock{
		Mocks: []*basesdk.BaseClientMock{
			badMarketIDMock,
		},
	}
	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "GARBAGE_MARKET_ID", Unmerged())
	assert.EqualError(t, err, "the market Id: GARBAGE_MARKET_ID is not valid")
}

func Test_GetWhiteLabelDataReturnsMarketNotFoundErrorWhen404FromServiceWhenTryingToGetMarketWhiteLabelDataIfUnmergedIsPassedIn(t *testing.T) {
	baseClient := &basesdk.BaseClientMultiCallMock{
		Mocks: []*basesdk.BaseClientMock{
			{Error: &basesdk.HTTPError{Body: "Get Failed", StatusCode: 404}},
		},
	}

	client := whiteLabelClient{SDKClient: baseClient}
	_, err := client.Get(context.Background(), "ABC", "Market", Unmerged())
	assert.Equal(t, ErrorMarketNotFound, err)
}
