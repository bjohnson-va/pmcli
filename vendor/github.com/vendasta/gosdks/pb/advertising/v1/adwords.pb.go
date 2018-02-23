// Code generated by protoc-gen-go. DO NOT EDIT.
// source: advertising/v1/adwords.proto

/*
Package advertising_v1 is a generated protocol buffer package.

It is generated from these files:
	advertising/v1/adwords.proto
	advertising/v1/api.proto
	advertising/v1/campaign.proto
	advertising/v1/campaign_events.proto
	advertising/v1/campaign_suggestions.proto

It has these top-level messages:
	GetAllAccountsRequest
	AdwordsAccountInfo
	GetAllAccountsResponse
	GetStatsForBusinessRequest
	GetStatsForBusinessResponse
	StoreCredentialsForBusinessRequest
	GetAccountInfoRequest
	Account
	GetAccountInfoResponse
	CampaignStatus
	Campaign
	CampaignGetRequest
	CampaignGetResponse
	CampaignListRequest
	CampaignListResponse
	CampaignGetMultiRequest
	CampaignGetMultiResponse
	PagedResponseMetadata
	PagedRequestOptions
	CreateOrderEventRequest
	OrderEvent
	CreateLifelineUpdateEventRequest
	LifelineUpdateEvent
	PerformanceEstimation
	PerformanceEstimateRange
	KeywordPerformanceEstimation
	SuggestCampaignWithEstimationRequest
	SuggestCampaignWithEstimationResponse
*/
package advertising_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetAllAccountsRequest struct {
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *GetAllAccountsRequest) Reset()                    { *m = GetAllAccountsRequest{} }
func (m *GetAllAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllAccountsRequest) ProtoMessage()               {}
func (*GetAllAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetAllAccountsRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

type AdwordsAccountInfo struct {
	// The Account's ID; i.e. CID
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	// Display name for the account
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Is this a manager account (i.e. MCC)?
	IsManager bool `protobuf:"varint,3,opt,name=is_manager,json=isManager" json:"is_manager,omitempty"`
	// Is this a test account?
	IsTestAccount bool `protobuf:"varint,4,opt,name=is_test_account,json=isTestAccount" json:"is_test_account,omitempty"`
	// The ISO 4217 currency code of the account
	CurrencyCode string `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	// Local timezone of this account (see https://developers.google.com/adwords/api/docs/appendix/codes-formats#timezone-ids)
	Timezone string `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
}

func (m *AdwordsAccountInfo) Reset()                    { *m = AdwordsAccountInfo{} }
func (m *AdwordsAccountInfo) String() string            { return proto.CompactTextString(m) }
func (*AdwordsAccountInfo) ProtoMessage()               {}
func (*AdwordsAccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AdwordsAccountInfo) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *AdwordsAccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AdwordsAccountInfo) GetIsManager() bool {
	if m != nil {
		return m.IsManager
	}
	return false
}

func (m *AdwordsAccountInfo) GetIsTestAccount() bool {
	if m != nil {
		return m.IsTestAccount
	}
	return false
}

func (m *AdwordsAccountInfo) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *AdwordsAccountInfo) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

type GetAllAccountsResponse struct {
	Accounts []*AdwordsAccountInfo `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *GetAllAccountsResponse) Reset()                    { *m = GetAllAccountsResponse{} }
func (m *GetAllAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllAccountsResponse) ProtoMessage()               {}
func (*GetAllAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetAllAccountsResponse) GetAccounts() []*AdwordsAccountInfo {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Get stats from the AdWords API's "Account Performance" endpoint
type GetStatsForBusinessRequest struct {
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *GetStatsForBusinessRequest) Reset()                    { *m = GetStatsForBusinessRequest{} }
func (m *GetStatsForBusinessRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStatsForBusinessRequest) ProtoMessage()               {}
func (*GetStatsForBusinessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetStatsForBusinessRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

// Stats from the AdWords API's "Account Performance" endpoint
type GetStatsForBusinessResponse struct {
	Clicks           int64 `protobuf:"varint,1,opt,name=clicks" json:"clicks,omitempty"`
	CostMicrodollars int64 `protobuf:"varint,2,opt,name=cost_microdollars,json=costMicrodollars" json:"cost_microdollars,omitempty"`
	Impressions      int64 `protobuf:"varint,3,opt,name=impressions" json:"impressions,omitempty"`
	// AdWords campaigns/accounts, when built, are set up to optimize certain types of
	// conversions.  This is the metric of how many of those type of conversions
	// were achieved by the campaign.
	Conversions float64 `protobuf:"fixed64,4,opt,name=conversions" json:"conversions,omitempty"`
	// See conversions.  This is the metric of how many conversions were achived
	// on the campaign from the set of ALL conversions supported by AdWords.
	AllConversions float64 `protobuf:"fixed64,5,opt,name=all_conversions,json=allConversions" json:"all_conversions,omitempty"`
}

func (m *GetStatsForBusinessResponse) Reset()                    { *m = GetStatsForBusinessResponse{} }
func (m *GetStatsForBusinessResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStatsForBusinessResponse) ProtoMessage()               {}
func (*GetStatsForBusinessResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetStatsForBusinessResponse) GetClicks() int64 {
	if m != nil {
		return m.Clicks
	}
	return 0
}

func (m *GetStatsForBusinessResponse) GetCostMicrodollars() int64 {
	if m != nil {
		return m.CostMicrodollars
	}
	return 0
}

func (m *GetStatsForBusinessResponse) GetImpressions() int64 {
	if m != nil {
		return m.Impressions
	}
	return 0
}

func (m *GetStatsForBusinessResponse) GetConversions() float64 {
	if m != nil {
		return m.Conversions
	}
	return 0
}

func (m *GetStatsForBusinessResponse) GetAllConversions() float64 {
	if m != nil {
		return m.AllConversions
	}
	return 0
}

type StoreCredentialsForBusinessRequest struct {
	// The business to connect
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
	// An OAuth *Refresh* token for the business authorized against the advertising provider
	OauthRefreshToken string `protobuf:"bytes,2,opt,name=oauth_refresh_token,json=oauthRefreshToken" json:"oauth_refresh_token,omitempty"`
}

func (m *StoreCredentialsForBusinessRequest) Reset()         { *m = StoreCredentialsForBusinessRequest{} }
func (m *StoreCredentialsForBusinessRequest) String() string { return proto.CompactTextString(m) }
func (*StoreCredentialsForBusinessRequest) ProtoMessage()    {}
func (*StoreCredentialsForBusinessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *StoreCredentialsForBusinessRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

func (m *StoreCredentialsForBusinessRequest) GetOauthRefreshToken() string {
	if m != nil {
		return m.OauthRefreshToken
	}
	return ""
}

type GetAccountInfoRequest struct {
	// The business to get advertising account info about
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *GetAccountInfoRequest) Reset()                    { *m = GetAccountInfoRequest{} }
func (m *GetAccountInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoRequest) ProtoMessage()               {}
func (*GetAccountInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetAccountInfoRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

// An Adwords Account description
type Account struct {
	// The Account's ID; i.e. CID
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	// Display name for the account
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Is this a manager account (i.e. MCC)?
	IsManager bool `protobuf:"varint,3,opt,name=is_manager,json=isManager" json:"is_manager,omitempty"`
	// Is this a test account?
	IsTestAccount bool `protobuf:"varint,4,opt,name=is_test_account,json=isTestAccount" json:"is_test_account,omitempty"`
	// The ISO 4217 currency code of the account
	CurrencyCode string `protobuf:"bytes,5,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	// List of timezones here: https://developers.google.com/adwords/api/docs/appendix/codes-formats#timezone-ids
	TimeZone string `protobuf:"bytes,6,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Account) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetIsManager() bool {
	if m != nil {
		return m.IsManager
	}
	return false
}

func (m *Account) GetIsTestAccount() bool {
	if m != nil {
		return m.IsTestAccount
	}
	return false
}

func (m *Account) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Account) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

type GetAccountInfoResponse struct {
	// Info about an adwords account
	AccountInfo *Account `protobuf:"bytes,1,opt,name=account_info,json=accountInfo" json:"account_info,omitempty"`
}

func (m *GetAccountInfoResponse) Reset()                    { *m = GetAccountInfoResponse{} }
func (m *GetAccountInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoResponse) ProtoMessage()               {}
func (*GetAccountInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetAccountInfoResponse) GetAccountInfo() *Account {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetAllAccountsRequest)(nil), "advertising.v1.GetAllAccountsRequest")
	proto.RegisterType((*AdwordsAccountInfo)(nil), "advertising.v1.AdwordsAccountInfo")
	proto.RegisterType((*GetAllAccountsResponse)(nil), "advertising.v1.GetAllAccountsResponse")
	proto.RegisterType((*GetStatsForBusinessRequest)(nil), "advertising.v1.GetStatsForBusinessRequest")
	proto.RegisterType((*GetStatsForBusinessResponse)(nil), "advertising.v1.GetStatsForBusinessResponse")
	proto.RegisterType((*StoreCredentialsForBusinessRequest)(nil), "advertising.v1.StoreCredentialsForBusinessRequest")
	proto.RegisterType((*GetAccountInfoRequest)(nil), "advertising.v1.GetAccountInfoRequest")
	proto.RegisterType((*Account)(nil), "advertising.v1.Account")
	proto.RegisterType((*GetAccountInfoResponse)(nil), "advertising.v1.GetAccountInfoResponse")
}

func init() { proto.RegisterFile("advertising/v1/adwords.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x94, 0x49, 0x1b, 0x92, 0x97, 0x7e, 0xd0, 0x45, 0x14, 0xab, 0x05, 0x11, 0x19, 0x09, 0x22,
	0x21, 0xa5, 0x6a, 0xb9, 0x20, 0x24, 0x90, 0x42, 0x24, 0xaa, 0x1c, 0x7a, 0x71, 0x73, 0x40, 0x5c,
	0x56, 0xdb, 0xf5, 0x4b, 0xbb, 0xaa, 0xbd, 0x1b, 0xf6, 0xad, 0x83, 0xe0, 0x0f, 0x72, 0x44, 0xfc,
	0x23, 0xe4, 0xf5, 0xba, 0x35, 0x2d, 0x87, 0x8a, 0x13, 0x37, 0xef, 0xcc, 0xbc, 0x1d, 0x7b, 0xde,
	0xc8, 0xf0, 0x44, 0x64, 0x2b, 0xb4, 0x4e, 0x91, 0xd2, 0xe7, 0x07, 0xab, 0xc3, 0x03, 0x91, 0x7d,
	0x35, 0x36, 0xa3, 0xf1, 0xd2, 0x1a, 0x67, 0xd8, 0x56, 0x8b, 0x1d, 0xaf, 0x0e, 0x93, 0x37, 0xf0,
	0xe8, 0x18, 0xdd, 0x24, 0xcf, 0x27, 0x52, 0x9a, 0x52, 0x3b, 0x4a, 0xf1, 0x4b, 0x89, 0xe4, 0xd8,
	0x33, 0x18, 0x9c, 0x95, 0xa4, 0x34, 0x12, 0x71, 0x95, 0xc5, 0xd1, 0x30, 0x1a, 0xf5, 0x53, 0x68,
	0xa0, 0x59, 0x96, 0xfc, 0x8a, 0x80, 0x4d, 0xea, 0xbb, 0xc3, 0xec, 0x4c, 0x2f, 0x4c, 0x35, 0x27,
	0x4b, 0x72, 0xa6, 0x40, 0xdb, 0xcc, 0x75, 0x52, 0x68, 0xa0, 0x59, 0xc6, 0x18, 0xac, 0x69, 0x51,
	0x60, 0x7c, 0xcf, 0xdf, 0xe8, 0x9f, 0xd9, 0x53, 0x00, 0x45, 0xbc, 0x10, 0x5a, 0x9c, 0xa3, 0x8d,
	0x3b, 0xc3, 0x68, 0xd4, 0x4b, 0xfb, 0x8a, 0x4e, 0x6a, 0x80, 0xbd, 0x80, 0x6d, 0x45, 0xdc, 0x21,
	0x39, 0x2e, 0x6a, 0xab, 0x78, 0xcd, 0x6b, 0x36, 0x15, 0xcd, 0x91, 0x5c, 0xf0, 0x67, 0xcf, 0x61,
	0x53, 0x96, 0xd6, 0xa2, 0x96, 0xdf, 0xb8, 0x34, 0x19, 0xc6, 0xeb, 0xde, 0x63, 0xa3, 0x01, 0xa7,
	0x26, 0x43, 0xb6, 0x07, 0x3d, 0xa7, 0x0a, 0xfc, 0x6e, 0x34, 0xc6, 0x5d, 0xcf, 0x5f, 0x9d, 0x93,
	0x4f, 0xb0, 0x7b, 0x33, 0x0d, 0x5a, 0x1a, 0x4d, 0xc8, 0xde, 0x43, 0x2f, 0x58, 0x53, 0x1c, 0x0d,
	0x3b, 0xa3, 0xc1, 0x51, 0x32, 0xfe, 0x33, 0xca, 0xf1, 0xed, 0x30, 0xd2, 0xab, 0x99, 0xe4, 0x1d,
	0xec, 0x1d, 0xa3, 0x3b, 0x75, 0xc2, 0xd1, 0x47, 0x63, 0x3f, 0x84, 0x18, 0xef, 0x1c, 0xf6, 0xcf,
	0x08, 0xf6, 0xff, 0x3a, 0x1f, 0x5e, 0x6f, 0x17, 0xba, 0x32, 0x57, 0xf2, 0x92, 0x42, 0xe0, 0xe1,
	0xc4, 0x5e, 0xc1, 0x8e, 0x34, 0xe4, 0x78, 0xa1, 0xa4, 0x35, 0x99, 0xc9, 0x73, 0x61, 0xc9, 0x27,
	0xdf, 0x49, 0x1f, 0x54, 0xc4, 0x49, 0x0b, 0x67, 0x43, 0x18, 0xa8, 0x62, 0x69, 0x91, 0x48, 0x19,
	0x4d, 0x7e, 0x0d, 0x9d, 0xb4, 0x0d, 0x55, 0x0a, 0x69, 0xf4, 0x0a, 0x6d, 0xad, 0xa8, 0x96, 0x10,
	0xa5, 0x6d, 0x88, 0xbd, 0x84, 0x6d, 0x91, 0xe7, 0xbc, 0xad, 0x5a, 0xf7, 0xaa, 0x2d, 0x91, 0xe7,
	0xd3, 0x6b, 0x34, 0x29, 0x21, 0x39, 0x75, 0xc6, 0xe2, 0xd4, 0x62, 0x86, 0xda, 0x29, 0x91, 0xff,
	0x4b, 0x30, 0x6c, 0x0c, 0x0f, 0x8d, 0x28, 0xdd, 0x05, 0xb7, 0xb8, 0xb0, 0x48, 0x17, 0xdc, 0x99,
	0x4b, 0xd4, 0xa1, 0x5c, 0x3b, 0x9e, 0x4a, 0x6b, 0x66, 0x5e, 0x11, 0x4d, 0xdf, 0x5b, 0x3b, 0xba,
	0xeb, 0x0a, 0x7e, 0x44, 0x70, 0xbf, 0x29, 0xda, 0x7f, 0x5f, 0xf2, 0x7d, 0xe8, 0x57, 0xa5, 0xe6,
	0x37, 0x5b, 0xfe, 0xb9, 0x6a, 0xf9, 0xbc, 0x6e, 0x79, 0x3b, 0x83, 0x50, 0xa3, 0xb7, 0xb0, 0x11,
	0xbc, 0xb9, 0xd2, 0x0b, 0xe3, 0x3f, 0x6c, 0x70, 0xf4, 0xf8, 0x56, 0xd3, 0x6b, 0x4d, 0x3a, 0x10,
	0xd7, 0x77, 0x9c, 0x75, 0xfd, 0x0f, 0xe6, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xdd,
	0x34, 0x34, 0x80, 0x04, 0x00, 0x00,
}
