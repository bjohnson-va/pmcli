// Code generated by protoc-gen-go. DO NOT EDIT.
// source: advertising/v1/api.proto

package advertising_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CampaignSuggestions service

type CampaignSuggestionsClient interface {
	// Endpoint to suggest keywords for a business and estimate their performance.
	SuggestCampaignWithEstimation(ctx context.Context, in *SuggestCampaignWithEstimationRequest, opts ...grpc.CallOption) (*SuggestCampaignWithEstimationResponse, error)
}

type campaignSuggestionsClient struct {
	cc *grpc.ClientConn
}

func NewCampaignSuggestionsClient(cc *grpc.ClientConn) CampaignSuggestionsClient {
	return &campaignSuggestionsClient{cc}
}

func (c *campaignSuggestionsClient) SuggestCampaignWithEstimation(ctx context.Context, in *SuggestCampaignWithEstimationRequest, opts ...grpc.CallOption) (*SuggestCampaignWithEstimationResponse, error) {
	out := new(SuggestCampaignWithEstimationResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.CampaignSuggestions/SuggestCampaignWithEstimation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CampaignSuggestions service

type CampaignSuggestionsServer interface {
	// Endpoint to suggest keywords for a business and estimate their performance.
	SuggestCampaignWithEstimation(context.Context, *SuggestCampaignWithEstimationRequest) (*SuggestCampaignWithEstimationResponse, error)
}

func RegisterCampaignSuggestionsServer(s *grpc.Server, srv CampaignSuggestionsServer) {
	s.RegisterService(&_CampaignSuggestions_serviceDesc, srv)
}

func _CampaignSuggestions_SuggestCampaignWithEstimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestCampaignWithEstimationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignSuggestionsServer).SuggestCampaignWithEstimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.CampaignSuggestions/SuggestCampaignWithEstimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignSuggestionsServer).SuggestCampaignWithEstimation(ctx, req.(*SuggestCampaignWithEstimationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignSuggestions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertising.v1.CampaignSuggestions",
	HandlerType: (*CampaignSuggestionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestCampaignWithEstimation",
			Handler:    _CampaignSuggestions_SuggestCampaignWithEstimation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertising/v1/api.proto",
}

// Client API for Campaigns service

type CampaignsClient interface {
	// contract: does not return wholesale price
	// contract: returns a gRPC 403 if caller is unauthorized to access business
	// contract: if Campaign does not exist, returns gRPC 404
	// Get a campaign by campaign id
	Get(ctx context.Context, in *CampaignGetRequest, opts ...grpc.CallOption) (*CampaignGetResponse, error)
	// contract: does not return wholesale price
	// contract: returns a gRPC 403 if caller is unauthorized to access business
	// contract: return error if any of the campaign request failed
	// contract: if any Campaign does not exist, return nil in the response list
	// Get a list of campaigns by given campaign ids
	GetMulti(ctx context.Context, in *CampaignGetMultiRequest, opts ...grpc.CallOption) (*CampaignGetMultiResponse, error)
	// contract: returns campaigns by date (most recent first)
	// contract: returns gRPC 404 if business ID not found
	// contract: returns empty list of campaigns if business has no campaigns
	// contract: returns gRPC 403 if caller is unauthorized to access business
	// Get a list of campaigns for a business
	List(ctx context.Context, in *CampaignListRequest, opts ...grpc.CallOption) (*CampaignListResponse, error)
}

type campaignsClient struct {
	cc *grpc.ClientConn
}

func NewCampaignsClient(cc *grpc.ClientConn) CampaignsClient {
	return &campaignsClient{cc}
}

func (c *campaignsClient) Get(ctx context.Context, in *CampaignGetRequest, opts ...grpc.CallOption) (*CampaignGetResponse, error) {
	out := new(CampaignGetResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Campaigns/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignsClient) GetMulti(ctx context.Context, in *CampaignGetMultiRequest, opts ...grpc.CallOption) (*CampaignGetMultiResponse, error) {
	out := new(CampaignGetMultiResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Campaigns/GetMulti", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignsClient) List(ctx context.Context, in *CampaignListRequest, opts ...grpc.CallOption) (*CampaignListResponse, error) {
	out := new(CampaignListResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Campaigns/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Campaigns service

type CampaignsServer interface {
	// contract: does not return wholesale price
	// contract: returns a gRPC 403 if caller is unauthorized to access business
	// contract: if Campaign does not exist, returns gRPC 404
	// Get a campaign by campaign id
	Get(context.Context, *CampaignGetRequest) (*CampaignGetResponse, error)
	// contract: does not return wholesale price
	// contract: returns a gRPC 403 if caller is unauthorized to access business
	// contract: return error if any of the campaign request failed
	// contract: if any Campaign does not exist, return nil in the response list
	// Get a list of campaigns by given campaign ids
	GetMulti(context.Context, *CampaignGetMultiRequest) (*CampaignGetMultiResponse, error)
	// contract: returns campaigns by date (most recent first)
	// contract: returns gRPC 404 if business ID not found
	// contract: returns empty list of campaigns if business has no campaigns
	// contract: returns gRPC 403 if caller is unauthorized to access business
	// Get a list of campaigns for a business
	List(context.Context, *CampaignListRequest) (*CampaignListResponse, error)
}

func RegisterCampaignsServer(s *grpc.Server, srv CampaignsServer) {
	s.RegisterService(&_Campaigns_serviceDesc, srv)
}

func _Campaigns_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CampaignGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Campaigns/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignsServer).Get(ctx, req.(*CampaignGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campaigns_GetMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CampaignGetMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignsServer).GetMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Campaigns/GetMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignsServer).GetMulti(ctx, req.(*CampaignGetMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Campaigns_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CampaignListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Campaigns/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignsServer).List(ctx, req.(*CampaignListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Campaigns_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertising.v1.Campaigns",
	HandlerType: (*CampaignsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Campaigns_Get_Handler,
		},
		{
			MethodName: "GetMulti",
			Handler:    _Campaigns_GetMulti_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Campaigns_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertising/v1/api.proto",
}

// Client API for CampaignEvents service

type CampaignEventsClient interface {
	// Create an event signalling that an order was received
	CreateOrderEvent(ctx context.Context, in *CreateOrderEventRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Create an event signalling that someone has updated values in the lifeline
	CreateLifelineUpdateEvent(ctx context.Context, in *CreateLifelineUpdateEventRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Send an signalling to let Advertising creates campaign near end events for all valid campaigns
	RequestCampaignNearEndEventsCreation(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type campaignEventsClient struct {
	cc *grpc.ClientConn
}

func NewCampaignEventsClient(cc *grpc.ClientConn) CampaignEventsClient {
	return &campaignEventsClient{cc}
}

func (c *campaignEventsClient) CreateOrderEvent(ctx context.Context, in *CreateOrderEventRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/advertising.v1.CampaignEvents/CreateOrderEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignEventsClient) CreateLifelineUpdateEvent(ctx context.Context, in *CreateLifelineUpdateEventRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/advertising.v1.CampaignEvents/CreateLifelineUpdateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignEventsClient) RequestCampaignNearEndEventsCreation(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/advertising.v1.CampaignEvents/RequestCampaignNearEndEventsCreation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CampaignEvents service

type CampaignEventsServer interface {
	// Create an event signalling that an order was received
	CreateOrderEvent(context.Context, *CreateOrderEventRequest) (*google_protobuf.Empty, error)
	// Create an event signalling that someone has updated values in the lifeline
	CreateLifelineUpdateEvent(context.Context, *CreateLifelineUpdateEventRequest) (*google_protobuf.Empty, error)
	// Send an signalling to let Advertising creates campaign near end events for all valid campaigns
	RequestCampaignNearEndEventsCreation(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterCampaignEventsServer(s *grpc.Server, srv CampaignEventsServer) {
	s.RegisterService(&_CampaignEvents_serviceDesc, srv)
}

func _CampaignEvents_CreateOrderEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignEventsServer).CreateOrderEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.CampaignEvents/CreateOrderEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignEventsServer).CreateOrderEvent(ctx, req.(*CreateOrderEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignEvents_CreateLifelineUpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLifelineUpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignEventsServer).CreateLifelineUpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.CampaignEvents/CreateLifelineUpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignEventsServer).CreateLifelineUpdateEvent(ctx, req.(*CreateLifelineUpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignEvents_RequestCampaignNearEndEventsCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignEventsServer).RequestCampaignNearEndEventsCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.CampaignEvents/RequestCampaignNearEndEventsCreation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignEventsServer).RequestCampaignNearEndEventsCreation(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignEvents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertising.v1.CampaignEvents",
	HandlerType: (*CampaignEventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderEvent",
			Handler:    _CampaignEvents_CreateOrderEvent_Handler,
		},
		{
			MethodName: "CreateLifelineUpdateEvent",
			Handler:    _CampaignEvents_CreateLifelineUpdateEvent_Handler,
		},
		{
			MethodName: "RequestCampaignNearEndEventsCreation",
			Handler:    _CampaignEvents_RequestCampaignNearEndEventsCreation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertising/v1/api.proto",
}

// Client API for Admin service

type AdminClient interface {
	// Iterates over all campaigns and re-calculates and puts all notification events
	ReputAllNotificationEvents(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) ReputAllNotificationEvents(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/advertising.v1.Admin/ReputAllNotificationEvents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	// Iterates over all campaigns and re-calculates and puts all notification events
	ReputAllNotificationEvents(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_ReputAllNotificationEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ReputAllNotificationEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Admin/ReputAllNotificationEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ReputAllNotificationEvents(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertising.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReputAllNotificationEvents",
			Handler:    _Admin_ReputAllNotificationEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertising/v1/api.proto",
}

// Client API for Adwords service

type AdwordsClient interface {
	// Gets all accounts accessible to a business
	GetAllAccounts(ctx context.Context, in *GetAllAccountsRequest, opts ...grpc.CallOption) (*GetAllAccountsResponse, error)
	// Gets a stats overview for all accounts associated with the business
	GetStatsForBusiness(ctx context.Context, in *GetStatsForBusinessRequest, opts ...grpc.CallOption) (*GetStatsForBusinessResponse, error)
	// Associates credentials with a business
	StoreCredentialsForBusiness(ctx context.Context, in *StoreCredentialsForBusinessRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Gets info about the adwords account which is connected to a business
	GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error)
}

type adwordsClient struct {
	cc *grpc.ClientConn
}

func NewAdwordsClient(cc *grpc.ClientConn) AdwordsClient {
	return &adwordsClient{cc}
}

func (c *adwordsClient) GetAllAccounts(ctx context.Context, in *GetAllAccountsRequest, opts ...grpc.CallOption) (*GetAllAccountsResponse, error) {
	out := new(GetAllAccountsResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Adwords/GetAllAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adwordsClient) GetStatsForBusiness(ctx context.Context, in *GetStatsForBusinessRequest, opts ...grpc.CallOption) (*GetStatsForBusinessResponse, error) {
	out := new(GetStatsForBusinessResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Adwords/GetStatsForBusiness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adwordsClient) StoreCredentialsForBusiness(ctx context.Context, in *StoreCredentialsForBusinessRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/advertising.v1.Adwords/StoreCredentialsForBusiness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adwordsClient) GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error) {
	out := new(GetAccountInfoResponse)
	err := grpc.Invoke(ctx, "/advertising.v1.Adwords/GetAccountInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Adwords service

type AdwordsServer interface {
	// Gets all accounts accessible to a business
	GetAllAccounts(context.Context, *GetAllAccountsRequest) (*GetAllAccountsResponse, error)
	// Gets a stats overview for all accounts associated with the business
	GetStatsForBusiness(context.Context, *GetStatsForBusinessRequest) (*GetStatsForBusinessResponse, error)
	// Associates credentials with a business
	StoreCredentialsForBusiness(context.Context, *StoreCredentialsForBusinessRequest) (*google_protobuf.Empty, error)
	// Gets info about the adwords account which is connected to a business
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoResponse, error)
}

func RegisterAdwordsServer(s *grpc.Server, srv AdwordsServer) {
	s.RegisterService(&_Adwords_serviceDesc, srv)
}

func _Adwords_GetAllAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdwordsServer).GetAllAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Adwords/GetAllAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdwordsServer).GetAllAccounts(ctx, req.(*GetAllAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adwords_GetStatsForBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsForBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdwordsServer).GetStatsForBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Adwords/GetStatsForBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdwordsServer).GetStatsForBusiness(ctx, req.(*GetStatsForBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adwords_StoreCredentialsForBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredentialsForBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdwordsServer).StoreCredentialsForBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Adwords/StoreCredentialsForBusiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdwordsServer).StoreCredentialsForBusiness(ctx, req.(*StoreCredentialsForBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Adwords_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdwordsServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/advertising.v1.Adwords/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdwordsServer).GetAccountInfo(ctx, req.(*GetAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adwords_serviceDesc = grpc.ServiceDesc{
	ServiceName: "advertising.v1.Adwords",
	HandlerType: (*AdwordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAccounts",
			Handler:    _Adwords_GetAllAccounts_Handler,
		},
		{
			MethodName: "GetStatsForBusiness",
			Handler:    _Adwords_GetStatsForBusiness_Handler,
		},
		{
			MethodName: "StoreCredentialsForBusiness",
			Handler:    _Adwords_StoreCredentialsForBusiness_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _Adwords_GetAccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advertising/v1/api.proto",
}

func init() { proto.RegisterFile("advertising/v1/api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xfb, 0xc1, 0xe7, 0x1c, 0x22, 0xb4, 0x95, 0x10, 0xa4, 0xf4, 0x12, 0xc2, 0xb7, 0xe4,
	0xd0, 0x00, 0x0f, 0x10, 0xa2, 0x60, 0x21, 0x95, 0x80, 0x1a, 0x55, 0xe5, 0x56, 0x6d, 0xe3, 0x89,
	0x19, 0xe1, 0xec, 0x9a, 0xdd, 0x71, 0x10, 0x2f, 0xc1, 0x85, 0x77, 0xe0, 0xc0, 0x81, 0x67, 0x44,
	0xfe, 0xd8, 0xb4, 0x31, 0xb6, 0xd3, 0x1c, 0x93, 0xff, 0x6f, 0x7e, 0xe3, 0x19, 0xcd, 0xc2, 0x3d,
	0x19, 0x2c, 0xd0, 0x30, 0x59, 0x52, 0x61, 0x6f, 0x71, 0xd8, 0x93, 0x31, 0x79, 0xb1, 0xd1, 0xac,
	0x45, 0xeb, 0x52, 0xe2, 0x2d, 0x0e, 0xdb, 0xfb, 0xa1, 0xd6, 0x61, 0x84, 0xbd, 0x2c, 0x3d, 0x4f,
	0x66, 0x3d, 0x9c, 0xc7, 0xfc, 0x23, 0x87, 0xdb, 0x07, 0x25, 0xcd, 0x54, 0xce, 0x63, 0x49, 0xa1,
	0x2a, 0xe2, 0x6e, 0x4d, 0x7c, 0x86, 0x0b, 0x54, 0x6c, 0x0b, 0xea, 0x59, 0x1d, 0x65, 0x93, 0x30,
	0x44, 0xcb, 0xa4, 0x95, 0x43, 0x1f, 0x94, 0x3f, 0x3b, 0xf8, 0xae, 0x4d, 0x50, 0xa4, 0xfd, 0xdf,
	0xdb, 0xb0, 0x37, 0x2c, 0x8a, 0x27, 0x17, 0xb5, 0xe2, 0xe7, 0x36, 0x1c, 0x14, 0xbf, 0x5d, 0x7c,
	0x4a, 0xfc, 0x65, 0x64, 0x99, 0xe6, 0x32, 0x45, 0xc4, 0x6b, 0x6f, 0x75, 0x6a, 0xaf, 0x11, 0x3f,
	0xc6, 0x6f, 0x09, 0x5a, 0x6e, 0xbf, 0xd9, 0xb0, 0xca, 0xc6, 0x5a, 0x59, 0xec, 0x6c, 0xf5, 0x7f,
	0xed, 0xc0, 0x6d, 0x07, 0x59, 0xf1, 0x09, 0x76, 0x7d, 0x64, 0xd1, 0x29, 0xdb, 0x1c, 0xe1, 0x23,
	0xbb, 0x8e, 0x0f, 0x1b, 0x19, 0xe7, 0x17, 0x67, 0x70, 0xcb, 0x47, 0xfe, 0x90, 0x44, 0x4c, 0xe2,
	0x49, 0x43, 0x49, 0x46, 0x38, 0xf7, 0xd3, 0xf5, 0xe0, 0xb2, 0xc1, 0x04, 0xae, 0x1d, 0x91, 0x65,
	0x51, 0xfb, 0x3d, 0x69, 0xea, 0xc4, 0xdd, 0x66, 0x68, 0xb9, 0x95, 0x3f, 0x3b, 0xd0, 0x72, 0xd1,
	0x28, 0x3b, 0x10, 0x71, 0x02, 0x77, 0x86, 0x06, 0x25, 0xe3, 0x47, 0x13, 0xa0, 0xc9, 0xfe, 0xac,
	0x18, 0xa8, 0x44, 0xb8, 0xbe, 0x77, 0xbd, 0xfc, 0x74, 0x3d, 0x77, 0xba, 0xde, 0x28, 0x3d, 0xdd,
	0xce, 0x96, 0x08, 0xe1, 0x7e, 0x5e, 0x74, 0x44, 0x33, 0x8c, 0x48, 0xe1, 0x49, 0x1c, 0x48, 0xc6,
	0xdc, 0xff, 0xb2, 0xda, 0x5f, 0x81, 0xae, 0x6f, 0xf4, 0x19, 0xba, 0x05, 0xe4, 0x06, 0x1b, 0xa3,
	0x34, 0x23, 0x15, 0xe4, 0xf3, 0x65, 0xe6, 0xf4, 0xfe, 0x6a, 0x0c, 0xf5, 0xe6, 0xfe, 0x29, 0x5c,
	0x1f, 0x04, 0x73, 0x52, 0x62, 0x0c, 0xed, 0x63, 0x8c, 0x13, 0x1e, 0x44, 0xd1, 0x58, 0x33, 0xcd,
	0x68, 0x9a, 0x29, 0x8b, 0x05, 0x6e, 0x2e, 0xfe, 0xbb, 0x0b, 0x37, 0x07, 0xf9, 0xb3, 0x12, 0x12,
	0x5a, 0x3e, 0xa6, 0xe6, 0xc1, 0x74, 0xaa, 0x93, 0xd4, 0xf7, 0xa8, 0xbc, 0x9c, 0xd5, 0xdc, 0x6d,
	0xe4, 0xf1, 0x3a, 0x6c, 0x79, 0x49, 0x31, 0xec, 0xf9, 0xc8, 0x13, 0x96, 0x6c, 0xdf, 0x69, 0xf3,
	0x36, 0xb1, 0xa4, 0xd0, 0x5a, 0xf1, 0xbc, 0x42, 0x50, 0x86, 0x5c, 0xb3, 0x17, 0x57, 0x62, 0x97,
	0x1d, 0xbf, 0xc2, 0xfe, 0x84, 0xb5, 0xc1, 0xa1, 0xc1, 0x00, 0x15, 0x93, 0x8c, 0x56, 0x3a, 0xf7,
	0xff, 0x7b, 0xd4, 0xf5, 0xf0, 0xfa, 0x03, 0x28, 0x36, 0x98, 0xcf, 0xfd, 0x5e, 0xcd, 0x74, 0xf5,
	0x06, 0x2f, 0xf2, 0xc6, 0x0d, 0x5e, 0xc6, 0xdc, 0x3c, 0xe7, 0x37, 0xb2, 0xa6, 0xaf, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0x4c, 0xc6, 0xa5, 0xd3, 0x05, 0x00, 0x00,
}
