// Code generated by protoc-gen-go. DO NOT EDIT.
// source: advertising/v1/campaign_suggestions.proto

package advertising_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Performance Stats which were estimated for a target
type PerformanceEstimation struct {
	// Estimate of average cost per click (in Micros)
	AverageCpc int64 `protobuf:"varint,1,opt,name=average_cpc,json=averageCpc" json:"average_cpc,omitempty"`
	// Estimate of average position in search results
	AveragePosition float64 `protobuf:"fixed64,2,opt,name=average_position,json=averagePosition" json:"average_position,omitempty"`
	// Estimate of possible clicks per month
	ClicksPerMonth float32 `protobuf:"fixed32,3,opt,name=clicks_per_month,json=clicksPerMonth" json:"clicks_per_month,omitempty"`
	// Estimate of possible click-throughs per month
	ClickThroughRate float64 `protobuf:"fixed64,4,opt,name=click_through_rate,json=clickThroughRate" json:"click_through_rate,omitempty"`
	// Estimate of possible impressions per month
	ImpressionsPerMonth float32 `protobuf:"fixed32,5,opt,name=impressions_per_month,json=impressionsPerMonth" json:"impressions_per_month,omitempty"`
	// Estimate of total cost per month (in Micros)
	TotalCost int64 `protobuf:"varint,6,opt,name=total_cost,json=totalCost" json:"total_cost,omitempty"`
}

func (m *PerformanceEstimation) Reset()                    { *m = PerformanceEstimation{} }
func (m *PerformanceEstimation) String() string            { return proto.CompactTextString(m) }
func (*PerformanceEstimation) ProtoMessage()               {}
func (*PerformanceEstimation) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *PerformanceEstimation) GetAverageCpc() int64 {
	if m != nil {
		return m.AverageCpc
	}
	return 0
}

func (m *PerformanceEstimation) GetAveragePosition() float64 {
	if m != nil {
		return m.AveragePosition
	}
	return 0
}

func (m *PerformanceEstimation) GetClicksPerMonth() float32 {
	if m != nil {
		return m.ClicksPerMonth
	}
	return 0
}

func (m *PerformanceEstimation) GetClickThroughRate() float64 {
	if m != nil {
		return m.ClickThroughRate
	}
	return 0
}

func (m *PerformanceEstimation) GetImpressionsPerMonth() float32 {
	if m != nil {
		return m.ImpressionsPerMonth
	}
	return 0
}

func (m *PerformanceEstimation) GetTotalCost() int64 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

// An upper and lower bounded range of performance estimations
type PerformanceEstimateRange struct {
	Min *PerformanceEstimation `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max *PerformanceEstimation `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
}

func (m *PerformanceEstimateRange) Reset()                    { *m = PerformanceEstimateRange{} }
func (m *PerformanceEstimateRange) String() string            { return proto.CompactTextString(m) }
func (*PerformanceEstimateRange) ProtoMessage()               {}
func (*PerformanceEstimateRange) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *PerformanceEstimateRange) GetMin() *PerformanceEstimation {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *PerformanceEstimateRange) GetMax() *PerformanceEstimation {
	if m != nil {
		return m.Max
	}
	return nil
}

// A performance estimation for a single keyword
type KeywordPerformanceEstimation struct {
	// The keyword as a string
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// The performance estimate as a min-max range
	Range *PerformanceEstimateRange `protobuf:"bytes,2,opt,name=range" json:"range,omitempty"`
}

func (m *KeywordPerformanceEstimation) Reset()                    { *m = KeywordPerformanceEstimation{} }
func (m *KeywordPerformanceEstimation) String() string            { return proto.CompactTextString(m) }
func (*KeywordPerformanceEstimation) ProtoMessage()               {}
func (*KeywordPerformanceEstimation) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *KeywordPerformanceEstimation) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *KeywordPerformanceEstimation) GetRange() *PerformanceEstimateRange {
	if m != nil {
		return m.Range
	}
	return nil
}

// Request to suggest keywords for a business and estimate their performance.
type SuggestCampaignWithEstimationRequest struct {
	// The identifier of a business in our system
	BusinessId string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *SuggestCampaignWithEstimationRequest) Reset()         { *m = SuggestCampaignWithEstimationRequest{} }
func (m *SuggestCampaignWithEstimationRequest) String() string { return proto.CompactTextString(m) }
func (*SuggestCampaignWithEstimationRequest) ProtoMessage()    {}
func (*SuggestCampaignWithEstimationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{3}
}

func (m *SuggestCampaignWithEstimationRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

// Response containing keyword suggestions and estimated performance for a business.
type SuggestCampaignWithEstimationResponse struct {
	KeywordPerformanceEstimations []*KeywordPerformanceEstimation `protobuf:"bytes,1,rep,name=keyword_performance_estimations,json=keywordPerformanceEstimations" json:"keyword_performance_estimations,omitempty"`
}

func (m *SuggestCampaignWithEstimationResponse) Reset()         { *m = SuggestCampaignWithEstimationResponse{} }
func (m *SuggestCampaignWithEstimationResponse) String() string { return proto.CompactTextString(m) }
func (*SuggestCampaignWithEstimationResponse) ProtoMessage()    {}
func (*SuggestCampaignWithEstimationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{4}
}

func (m *SuggestCampaignWithEstimationResponse) GetKeywordPerformanceEstimations() []*KeywordPerformanceEstimation {
	if m != nil {
		return m.KeywordPerformanceEstimations
	}
	return nil
}

func init() {
	proto.RegisterType((*PerformanceEstimation)(nil), "advertising.v1.PerformanceEstimation")
	proto.RegisterType((*PerformanceEstimateRange)(nil), "advertising.v1.PerformanceEstimateRange")
	proto.RegisterType((*KeywordPerformanceEstimation)(nil), "advertising.v1.KeywordPerformanceEstimation")
	proto.RegisterType((*SuggestCampaignWithEstimationRequest)(nil), "advertising.v1.SuggestCampaignWithEstimationRequest")
	proto.RegisterType((*SuggestCampaignWithEstimationResponse)(nil), "advertising.v1.SuggestCampaignWithEstimationResponse")
}

func init() { proto.RegisterFile("advertising/v1/campaign_suggestions.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0xe6, 0x92, 0xb6, 0xd0, 0x09, 0xd4, 0xb2, 0x52, 0xb8, 0x07, 0x4b, 0xc3, 0x61, 0xe1, 0x0a,
	0x25, 0xa5, 0xf1, 0xc1, 0x37, 0x5f, 0x82, 0x88, 0x88, 0x10, 0x56, 0xc1, 0xc7, 0x65, 0x7b, 0x19,
	0x2f, 0x4b, 0x7a, 0xbb, 0xeb, 0xce, 0x24, 0xc6, 0xff, 0x20, 0xfe, 0x03, 0xff, 0xab, 0xec, 0xde,
	0xa5, 0x9c, 0x5a, 0x5a, 0xfa, 0x76, 0x7c, 0x33, 0xdf, 0xcc, 0xf7, 0x7d, 0x3b, 0x07, 0x17, 0x7a,
	0xb1, 0xc1, 0xc0, 0x86, 0x8c, 0xad, 0xaf, 0x36, 0xd7, 0x57, 0x95, 0x6e, 0xbc, 0x36, 0xb5, 0x55,
	0xb4, 0xae, 0x6b, 0x24, 0x36, 0xce, 0xd2, 0xc4, 0x07, 0xc7, 0x4e, 0x1c, 0xf5, 0x5a, 0x27, 0x9b,
	0xeb, 0xe2, 0xd7, 0x00, 0x4e, 0xe6, 0x18, 0xbe, 0xba, 0xd0, 0x68, 0x5b, 0xe1, 0x5b, 0x62, 0xd3,
	0xe8, 0x48, 0x10, 0x67, 0x30, 0xd2, 0x1b, 0x0c, 0xba, 0x46, 0x55, 0xf9, 0x2a, 0xcf, 0xc6, 0x59,
	0x39, 0x94, 0xd0, 0x41, 0x33, 0x5f, 0x89, 0x0b, 0x38, 0xde, 0x35, 0x78, 0x47, 0x26, 0x92, 0xf2,
	0xc1, 0x38, 0x2b, 0x33, 0xf9, 0xac, 0xc3, 0xe7, 0x1d, 0x2c, 0x4a, 0x38, 0xae, 0x6e, 0x4d, 0xb5,
	0x22, 0xe5, 0x31, 0xa8, 0xc6, 0x59, 0x5e, 0xe6, 0xc3, 0x71, 0x56, 0x0e, 0xe4, 0x51, 0x8b, 0xcf,
	0x31, 0x7c, 0x8c, 0xa8, 0xb8, 0x04, 0x91, 0x10, 0xc5, 0xcb, 0xe0, 0xd6, 0xf5, 0x52, 0x05, 0xcd,
	0x98, 0xef, 0xa5, 0xb1, 0xed, 0x8c, 0xcf, 0x6d, 0x41, 0x6a, 0x46, 0x31, 0x85, 0x13, 0xd3, 0xf8,
	0x80, 0x44, 0xd1, 0x62, 0x6f, 0xf8, 0x7e, 0x1a, 0xfe, 0xbc, 0x57, 0xbc, 0xdb, 0x70, 0x0a, 0xc0,
	0x8e, 0xf5, 0xad, 0xaa, 0x1c, 0x71, 0x7e, 0x90, 0x6c, 0x1d, 0x26, 0x64, 0xe6, 0x88, 0x8b, 0x9f,
	0x19, 0xe4, 0xff, 0x07, 0x82, 0x52, 0xdb, 0x1a, 0xc5, 0x6b, 0x18, 0x36, 0xc6, 0xa6, 0x2c, 0x46,
	0xd3, 0xf3, 0xc9, 0xdf, 0x59, 0x4e, 0xee, 0xcd, 0x51, 0x46, 0x46, 0x22, 0xea, 0x6d, 0x8a, 0xe7,
	0x09, 0x44, 0xbd, 0x2d, 0x02, 0xbc, 0xf8, 0x80, 0x3f, 0xbe, 0xbb, 0xb0, 0xb8, 0xff, 0x95, 0x04,
	0xec, 0x31, 0x6e, 0x39, 0x49, 0x3a, 0x94, 0xe9, 0x5b, 0xbc, 0x81, 0xfd, 0x10, 0xe5, 0x76, 0xeb,
	0xca, 0xc7, 0xd7, 0xb5, 0xf6, 0x64, 0x4b, 0x2b, 0xde, 0xc1, 0xcb, 0x4f, 0xed, 0xe1, 0xcc, 0xba,
	0x43, 0xfa, 0x62, 0x78, 0xd9, 0x53, 0x86, 0xdf, 0xd6, 0x48, 0x1c, 0x2f, 0xe4, 0x66, 0x4d, 0xc6,
	0x22, 0x91, 0x32, 0x8b, 0x4e, 0x02, 0xec, 0xa0, 0xf7, 0x8b, 0xe2, 0x77, 0x06, 0xe7, 0x8f, 0x4c,
	0x22, 0xef, 0x2c, 0xa1, 0x60, 0x38, 0x5b, 0xb5, 0x36, 0xe3, 0x23, 0xee, 0xd4, 0x29, 0xbc, 0xeb,
	0xa4, 0x3c, 0x1b, 0x0f, 0xcb, 0xd1, 0xf4, 0xf2, 0x5f, 0x33, 0x0f, 0xa5, 0x23, 0x4f, 0x57, 0x0f,
	0x54, 0xe9, 0xe6, 0x20, 0xfd, 0x13, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x35, 0x2f,
	0x16, 0x40, 0x03, 0x00, 0x00,
}