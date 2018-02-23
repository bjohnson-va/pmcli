// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package customerchurn_v1 is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PagedResponseMetadata
	PagedRequestOptions
	PredictionValue
	Prediction
	StringList
	FloatList
	BooleanList
	Metric
	PredictionWithMetrics
	ListResponse
	ListRequest
	PrioritizeRequest
	PrioritizeResponse
	IngestTrainingRequest
	IngestPredictionsRequest
*/
package customerchurn_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PagedResponseMetadata struct {
	// A cursor that can be provided to retrieve the next page of results
	NextCursor string `protobuf:"bytes,1,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	// Whether or not more results exist
	HasMore bool `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	// The total number of results for this query across all pages
	TotalResults int64 `protobuf:"varint,3,opt,name=total_results,json=totalResults" json:"total_results,omitempty"`
}

func (m *PagedResponseMetadata) Reset()                    { *m = PagedResponseMetadata{} }
func (m *PagedResponseMetadata) String() string            { return proto.CompactTextString(m) }
func (*PagedResponseMetadata) ProtoMessage()               {}
func (*PagedResponseMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PagedResponseMetadata) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *PagedResponseMetadata) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *PagedResponseMetadata) GetTotalResults() int64 {
	if m != nil {
		return m.TotalResults
	}
	return 0
}

type PagedRequestOptions struct {
	// cursor can be passed to retrieve the next page of results keyed by the cursor
	Cursor string `protobuf:"bytes,1,opt,name=cursor" json:"cursor,omitempty"`
	// page_size specifies the number of items to return in the next page
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *PagedRequestOptions) Reset()                    { *m = PagedRequestOptions{} }
func (m *PagedRequestOptions) String() string            { return proto.CompactTextString(m) }
func (*PagedRequestOptions) ProtoMessage()               {}
func (*PagedRequestOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PagedRequestOptions) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *PagedRequestOptions) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Predicted by the model, with metadata about the confidence of the predicted value
type PredictionValue struct {
	Value *Metric `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	// Describes how confidence of the model in this prediction, between 0 and 1.
	// The sum of all confidences across all predictions for a single frame is equal to 1.
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence" json:"confidence,omitempty"`
}

func (m *PredictionValue) Reset()                    { *m = PredictionValue{} }
func (m *PredictionValue) String() string            { return proto.CompactTextString(m) }
func (*PredictionValue) ProtoMessage()               {}
func (*PredictionValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PredictionValue) GetValue() *Metric {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PredictionValue) GetConfidence() float64 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// A Prediction is a set of predicted values
type Prediction struct {
	Values []*PredictionValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// Priority is a function of the churn likelihood and the value of the customer.
	Priority float64 `protobuf:"fixed64,2,opt,name=priority" json:"priority,omitempty"`
}

func (m *Prediction) Reset()                    { *m = Prediction{} }
func (m *Prediction) String() string            { return proto.CompactTextString(m) }
func (*Prediction) ProtoMessage()               {}
func (*Prediction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Prediction) GetValues() []*PredictionValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Prediction) GetPriority() float64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// Containers to hold repeated primitive values.
type StringList struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type FloatList struct {
	Value []float64 `protobuf:"fixed64,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *FloatList) Reset()                    { *m = FloatList{} }
func (m *FloatList) String() string            { return proto.CompactTextString(m) }
func (*FloatList) ProtoMessage()               {}
func (*FloatList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FloatList) GetValue() []float64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type BooleanList struct {
	Value []bool `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *BooleanList) Reset()                    { *m = BooleanList{} }
func (m *BooleanList) String() string            { return proto.CompactTextString(m) }
func (*BooleanList) ProtoMessage()               {}
func (*BooleanList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BooleanList) GetValue() []bool {
	if m != nil {
		return m.Value
	}
	return nil
}

type Metric struct {
	// The value of the metric
	//
	// Types that are valid to be assigned to Kind:
	//	*Metric_StringList
	//	*Metric_BoolList
	//	*Metric_FloatList
	Kind isMetric_Kind `protobuf_oneof:"kind"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isMetric_Kind interface {
	isMetric_Kind()
}

type Metric_StringList struct {
	StringList *StringList `protobuf:"bytes,1,opt,name=string_list,json=stringList,oneof"`
}
type Metric_BoolList struct {
	BoolList *BooleanList `protobuf:"bytes,2,opt,name=bool_list,json=boolList,oneof"`
}
type Metric_FloatList struct {
	FloatList *FloatList `protobuf:"bytes,3,opt,name=float_list,json=floatList,oneof"`
}

func (*Metric_StringList) isMetric_Kind() {}
func (*Metric_BoolList) isMetric_Kind()   {}
func (*Metric_FloatList) isMetric_Kind()  {}

func (m *Metric) GetKind() isMetric_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Metric) GetStringList() *StringList {
	if x, ok := m.GetKind().(*Metric_StringList); ok {
		return x.StringList
	}
	return nil
}

func (m *Metric) GetBoolList() *BooleanList {
	if x, ok := m.GetKind().(*Metric_BoolList); ok {
		return x.BoolList
	}
	return nil
}

func (m *Metric) GetFloatList() *FloatList {
	if x, ok := m.GetKind().(*Metric_FloatList); ok {
		return x.FloatList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Metric) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Metric_OneofMarshaler, _Metric_OneofUnmarshaler, _Metric_OneofSizer, []interface{}{
		(*Metric_StringList)(nil),
		(*Metric_BoolList)(nil),
		(*Metric_FloatList)(nil),
	}
}

func _Metric_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Metric)
	// kind
	switch x := m.Kind.(type) {
	case *Metric_StringList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringList); err != nil {
			return err
		}
	case *Metric_BoolList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BoolList); err != nil {
			return err
		}
	case *Metric_FloatList:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FloatList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Metric.Kind has unexpected type %T", x)
	}
	return nil
}

func _Metric_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Metric)
	switch tag {
	case 1: // kind.string_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringList)
		err := b.DecodeMessage(msg)
		m.Kind = &Metric_StringList{msg}
		return true, err
	case 2: // kind.bool_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BooleanList)
		err := b.DecodeMessage(msg)
		m.Kind = &Metric_BoolList{msg}
		return true, err
	case 3: // kind.float_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FloatList)
		err := b.DecodeMessage(msg)
		m.Kind = &Metric_FloatList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Metric_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Metric)
	// kind
	switch x := m.Kind.(type) {
	case *Metric_StringList:
		s := proto.Size(x.StringList)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Metric_BoolList:
		s := proto.Size(x.BoolList)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Metric_FloatList:
		s := proto.Size(x.FloatList)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PredictionWithMetrics associates a prediction and/or measured target value with the metrics used to make that prediction
type PredictionWithMetrics struct {
	// The result of the prediction based on metrics. This will be empty if the metrics are historical data (used exclusively as training data).
	Predictions map[string]*Prediction `protobuf:"bytes,1,rep,name=predictions" json:"predictions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The actual value for this set of metrics. This will be empty if we have not yet measure the actual value at this time.
	Actual map[string]*Metric `protobuf:"bytes,2,rep,name=actual" json:"actual,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Metrics used to make the prediction. These represent the inputs to the model.
	// The string key is the name of the metric.
	Metrics map[string]*Metric `protobuf:"bytes,3,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// When this prediction was made. All the metrics used to make the prediction were valid at this time.
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	// The customer this prediction was made for
	CustomerId string `protobuf:"bytes,5,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
}

func (m *PredictionWithMetrics) Reset()                    { *m = PredictionWithMetrics{} }
func (m *PredictionWithMetrics) String() string            { return proto.CompactTextString(m) }
func (*PredictionWithMetrics) ProtoMessage()               {}
func (*PredictionWithMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PredictionWithMetrics) GetPredictions() map[string]*Prediction {
	if m != nil {
		return m.Predictions
	}
	return nil
}

func (m *PredictionWithMetrics) GetActual() map[string]*Metric {
	if m != nil {
		return m.Actual
	}
	return nil
}

func (m *PredictionWithMetrics) GetMetrics() map[string]*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *PredictionWithMetrics) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PredictionWithMetrics) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type ListResponse struct {
	// helpers for paging over the rpc
	PagingMetadata *PagedResponseMetadata `protobuf:"bytes,1,opt,name=paging_metadata,json=pagingMetadata" json:"paging_metadata,omitempty"`
	// a list of predictions, actual values, and associated metrics ordered by time
	Frames []*PredictionWithMetrics `protobuf:"bytes,2,rep,name=frames" json:"frames,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListResponse) GetPagingMetadata() *PagedResponseMetadata {
	if m != nil {
		return m.PagingMetadata
	}
	return nil
}

func (m *ListResponse) GetFrames() []*PredictionWithMetrics {
	if m != nil {
		return m.Frames
	}
	return nil
}

type ListRequest struct {
	// model_id specifies which model you are asking for results for.
	// It's possible that predictions can be served for many different models for the same customer_id.
	ModelId string `protobuf:"bytes,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	// customer_id specifies the identifier of the customer you want to obtain a list of predictions for.
	// For a partner customer, this is a partner_id like "ABC", but for an SMB customer, it might be a user_id or email.
	CustomerId    string               `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	PagingOptions *PagedRequestOptions `protobuf:"bytes,3,opt,name=paging_options,json=pagingOptions" json:"paging_options,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *ListRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *ListRequest) GetPagingOptions() *PagedRequestOptions {
	if m != nil {
		return m.PagingOptions
	}
	return nil
}

type PrioritizeRequest struct {
	// model_id specifies which model you are asking for a prediction for.
	ModelId string `protobuf:"bytes,1,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	// customer_ids scopes the priorization to a finite set of customer_ids.
	// If this is not provided, priorization happens across all customer_ids.
	CustomerIds   *StringList          `protobuf:"bytes,2,opt,name=customer_ids,json=customerIds" json:"customer_ids,omitempty"`
	PagingOptions *PagedRequestOptions `protobuf:"bytes,3,opt,name=paging_options,json=pagingOptions" json:"paging_options,omitempty"`
}

func (m *PrioritizeRequest) Reset()                    { *m = PrioritizeRequest{} }
func (m *PrioritizeRequest) String() string            { return proto.CompactTextString(m) }
func (*PrioritizeRequest) ProtoMessage()               {}
func (*PrioritizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PrioritizeRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *PrioritizeRequest) GetCustomerIds() *StringList {
	if m != nil {
		return m.CustomerIds
	}
	return nil
}

func (m *PrioritizeRequest) GetPagingOptions() *PagedRequestOptions {
	if m != nil {
		return m.PagingOptions
	}
	return nil
}

type PrioritizeResponse struct {
	// helpers for paging over the rpc
	PagingMetadata *PagedResponseMetadata `protobuf:"bytes,1,opt,name=paging_metadata,json=pagingMetadata" json:"paging_metadata,omitempty"`
	// a list of predictions, actual values, and associated metrics ordered by priority
	Frames []*PredictionWithMetrics `protobuf:"bytes,2,rep,name=frames" json:"frames,omitempty"`
}

func (m *PrioritizeResponse) Reset()                    { *m = PrioritizeResponse{} }
func (m *PrioritizeResponse) String() string            { return proto.CompactTextString(m) }
func (*PrioritizeResponse) ProtoMessage()               {}
func (*PrioritizeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PrioritizeResponse) GetPagingMetadata() *PagedResponseMetadata {
	if m != nil {
		return m.PagingMetadata
	}
	return nil
}

func (m *PrioritizeResponse) GetFrames() []*PredictionWithMetrics {
	if m != nil {
		return m.Frames
	}
	return nil
}

type IngestTrainingRequest struct {
	// job_id specified the name of the ML Engine job that stores the training data you wish to ingest.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	// model_id specifies which model you are trying to ingest new training data from.
	ModelId string `protobuf:"bytes,2,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	// model_version specifies which version of the model you are trying to ingest new training data from.
	ModelVersion string `protobuf:"bytes,3,opt,name=model_version,json=modelVersion" json:"model_version,omitempty"`
}

func (m *IngestTrainingRequest) Reset()                    { *m = IngestTrainingRequest{} }
func (m *IngestTrainingRequest) String() string            { return proto.CompactTextString(m) }
func (*IngestTrainingRequest) ProtoMessage()               {}
func (*IngestTrainingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *IngestTrainingRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *IngestTrainingRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *IngestTrainingRequest) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

type IngestPredictionsRequest struct {
	// job_id specified the name of the ML Engine job that generated the predictions you wish to ingest.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	// model_id specifies which model you are trying to ingest new predictions from.
	ModelId string `protobuf:"bytes,2,opt,name=model_id,json=modelId" json:"model_id,omitempty"`
	// model_version specifies which version of the model you are trying to ingest new predictions from.
	ModelVersion string `protobuf:"bytes,3,opt,name=model_version,json=modelVersion" json:"model_version,omitempty"`
}

func (m *IngestPredictionsRequest) Reset()                    { *m = IngestPredictionsRequest{} }
func (m *IngestPredictionsRequest) String() string            { return proto.CompactTextString(m) }
func (*IngestPredictionsRequest) ProtoMessage()               {}
func (*IngestPredictionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *IngestPredictionsRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *IngestPredictionsRequest) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *IngestPredictionsRequest) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*PagedResponseMetadata)(nil), "customerchurn.v1.PagedResponseMetadata")
	proto.RegisterType((*PagedRequestOptions)(nil), "customerchurn.v1.PagedRequestOptions")
	proto.RegisterType((*PredictionValue)(nil), "customerchurn.v1.PredictionValue")
	proto.RegisterType((*Prediction)(nil), "customerchurn.v1.Prediction")
	proto.RegisterType((*StringList)(nil), "customerchurn.v1.StringList")
	proto.RegisterType((*FloatList)(nil), "customerchurn.v1.FloatList")
	proto.RegisterType((*BooleanList)(nil), "customerchurn.v1.BooleanList")
	proto.RegisterType((*Metric)(nil), "customerchurn.v1.Metric")
	proto.RegisterType((*PredictionWithMetrics)(nil), "customerchurn.v1.PredictionWithMetrics")
	proto.RegisterType((*ListResponse)(nil), "customerchurn.v1.ListResponse")
	proto.RegisterType((*ListRequest)(nil), "customerchurn.v1.ListRequest")
	proto.RegisterType((*PrioritizeRequest)(nil), "customerchurn.v1.PrioritizeRequest")
	proto.RegisterType((*PrioritizeResponse)(nil), "customerchurn.v1.PrioritizeResponse")
	proto.RegisterType((*IngestTrainingRequest)(nil), "customerchurn.v1.IngestTrainingRequest")
	proto.RegisterType((*IngestPredictionsRequest)(nil), "customerchurn.v1.IngestPredictionsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CustomerChurn service

type CustomerChurnClient interface {
	// List pages through both churn predictions and historical data for a specific customer, bundled with the associated metrics for each time period
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Prioritize pages through predictions across all customers, or a defined subset of customers for the latest time period, and returns those predictions ordered by priority
	Prioritize(ctx context.Context, in *PrioritizeRequest, opts ...grpc.CallOption) (*PrioritizeResponse, error)
}

type customerChurnClient struct {
	cc *grpc.ClientConn
}

func NewCustomerChurnClient(cc *grpc.ClientConn) CustomerChurnClient {
	return &customerChurnClient{cc}
}

func (c *customerChurnClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/customerchurn.v1.CustomerChurn/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerChurnClient) Prioritize(ctx context.Context, in *PrioritizeRequest, opts ...grpc.CallOption) (*PrioritizeResponse, error) {
	out := new(PrioritizeResponse)
	err := grpc.Invoke(ctx, "/customerchurn.v1.CustomerChurn/Prioritize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerChurn service

type CustomerChurnServer interface {
	// List pages through both churn predictions and historical data for a specific customer, bundled with the associated metrics for each time period
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Prioritize pages through predictions across all customers, or a defined subset of customers for the latest time period, and returns those predictions ordered by priority
	Prioritize(context.Context, *PrioritizeRequest) (*PrioritizeResponse, error)
}

func RegisterCustomerChurnServer(s *grpc.Server, srv CustomerChurnServer) {
	s.RegisterService(&_CustomerChurn_serviceDesc, srv)
}

func _CustomerChurn_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerChurnServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerchurn.v1.CustomerChurn/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerChurnServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerChurn_Prioritize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrioritizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerChurnServer).Prioritize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerchurn.v1.CustomerChurn/Prioritize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerChurnServer).Prioritize(ctx, req.(*PrioritizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerChurn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customerchurn.v1.CustomerChurn",
	HandlerType: (*CustomerChurnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CustomerChurn_List_Handler,
		},
		{
			MethodName: "Prioritize",
			Handler:    _CustomerChurn_Prioritize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for CustomerChurnAdmin service

type CustomerChurnAdminClient interface {
	// IngestTraining commands the microservice to ingest the training data for the ML Engine job and model
	IngestTraining(ctx context.Context, in *IngestTrainingRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// IngestTraining commands the microservice to ingest the prediction results for the ML Engine job and model
	IngestPredictions(ctx context.Context, in *IngestPredictionsRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type customerChurnAdminClient struct {
	cc *grpc.ClientConn
}

func NewCustomerChurnAdminClient(cc *grpc.ClientConn) CustomerChurnAdminClient {
	return &customerChurnAdminClient{cc}
}

func (c *customerChurnAdminClient) IngestTraining(ctx context.Context, in *IngestTrainingRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/customerchurn.v1.CustomerChurnAdmin/IngestTraining", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerChurnAdminClient) IngestPredictions(ctx context.Context, in *IngestPredictionsRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/customerchurn.v1.CustomerChurnAdmin/IngestPredictions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerChurnAdmin service

type CustomerChurnAdminServer interface {
	// IngestTraining commands the microservice to ingest the training data for the ML Engine job and model
	IngestTraining(context.Context, *IngestTrainingRequest) (*google_protobuf1.Empty, error)
	// IngestTraining commands the microservice to ingest the prediction results for the ML Engine job and model
	IngestPredictions(context.Context, *IngestPredictionsRequest) (*google_protobuf1.Empty, error)
}

func RegisterCustomerChurnAdminServer(s *grpc.Server, srv CustomerChurnAdminServer) {
	s.RegisterService(&_CustomerChurnAdmin_serviceDesc, srv)
}

func _CustomerChurnAdmin_IngestTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerChurnAdminServer).IngestTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerchurn.v1.CustomerChurnAdmin/IngestTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerChurnAdminServer).IngestTraining(ctx, req.(*IngestTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerChurnAdmin_IngestPredictions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngestPredictionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerChurnAdminServer).IngestPredictions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customerchurn.v1.CustomerChurnAdmin/IngestPredictions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerChurnAdminServer).IngestPredictions(ctx, req.(*IngestPredictionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerChurnAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customerchurn.v1.CustomerChurnAdmin",
	HandlerType: (*CustomerChurnAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IngestTraining",
			Handler:    _CustomerChurnAdmin_IngestTraining_Handler,
		},
		{
			MethodName: "IngestPredictions",
			Handler:    _CustomerChurnAdmin_IngestPredictions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 939 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xce, 0xda, 0x89, 0xe3, 0x7d, 0xd7, 0x69, 0xd3, 0x81, 0x54, 0xcb, 0x06, 0xda, 0xb0, 0xa1,
	0xaa, 0xc5, 0x61, 0x2b, 0x5c, 0x0e, 0x01, 0x21, 0x45, 0x4d, 0x54, 0x48, 0xa0, 0x81, 0x68, 0x12,
	0x15, 0x81, 0x90, 0xac, 0xf1, 0xee, 0xd8, 0x99, 0x76, 0x77, 0x67, 0xbb, 0x33, 0x8e, 0x9a, 0xfc,
	0x11, 0xce, 0x9c, 0xf8, 0x01, 0x9c, 0xb9, 0x71, 0xe5, 0xdf, 0xf0, 0x03, 0xd0, 0xce, 0xcc, 0x7a,
	0xd7, 0x5f, 0x25, 0x87, 0x22, 0xf5, 0x64, 0xcf, 0x3b, 0xef, 0xf3, 0xbc, 0x9f, 0xf3, 0xd8, 0x60,
	0x93, 0x8c, 0x05, 0x59, 0xce, 0x25, 0x47, 0x9b, 0xe1, 0x58, 0x48, 0x9e, 0xd0, 0x3c, 0xbc, 0x18,
	0xe7, 0x69, 0x70, 0xf9, 0x99, 0x77, 0x7f, 0xc4, 0xf9, 0x28, 0xa6, 0x8f, 0xd4, 0xfd, 0x60, 0x3c,
	0x7c, 0x24, 0x59, 0x42, 0x85, 0x24, 0x49, 0xa6, 0x21, 0xde, 0xf6, 0xac, 0x03, 0x4d, 0x32, 0x79,
	0xa5, 0x2f, 0xfd, 0xd7, 0xb0, 0x75, 0x4a, 0x46, 0x34, 0xc2, 0x54, 0x64, 0x3c, 0x15, 0xf4, 0x84,
	0x4a, 0x12, 0x11, 0x49, 0xd0, 0x7d, 0x70, 0x52, 0xfa, 0x5a, 0xf6, 0xc3, 0x71, 0x2e, 0x78, 0xee,
	0x5a, 0x3b, 0x56, 0xd7, 0xc6, 0x50, 0x98, 0x0e, 0x95, 0x05, 0x7d, 0x00, 0xed, 0x0b, 0x22, 0xfa,
	0x09, 0xcf, 0xa9, 0xdb, 0xd8, 0xb1, 0xba, 0x6d, 0xbc, 0x7e, 0x41, 0xc4, 0x09, 0xcf, 0x29, 0xda,
	0x85, 0x0d, 0xc9, 0x25, 0x89, 0xfb, 0x39, 0x15, 0xe3, 0x58, 0x0a, 0xb7, 0xb9, 0x63, 0x75, 0x9b,
	0xb8, 0xa3, 0x8c, 0x58, 0xdb, 0xfc, 0x6f, 0xe1, 0x3d, 0x13, 0xf9, 0xd5, 0x98, 0x0a, 0xf9, 0x43,
	0x26, 0x19, 0x4f, 0x05, 0xba, 0x0b, 0xad, 0xa9, 0x90, 0xe6, 0x84, 0xb6, 0xc1, 0xce, 0xc8, 0x88,
	0xf6, 0x05, 0xbb, 0xd6, 0xf1, 0x9a, 0xb8, 0x5d, 0x18, 0xce, 0xd8, 0x35, 0xf5, 0x09, 0xdc, 0x3e,
	0xcd, 0x69, 0xc4, 0xc2, 0x82, 0xe3, 0x39, 0x89, 0xc7, 0x14, 0x05, 0xb0, 0x76, 0x59, 0x7c, 0x51,
	0x34, 0x4e, 0xcf, 0x0d, 0x66, 0x1b, 0x17, 0x9c, 0x50, 0x99, 0xb3, 0x10, 0x6b, 0x37, 0x74, 0x0f,
	0x20, 0xe4, 0xe9, 0x90, 0x45, 0x34, 0x0d, 0x75, 0x00, 0x0b, 0xd7, 0x2c, 0x7e, 0x08, 0x50, 0x85,
	0x40, 0x5f, 0x40, 0x4b, 0xc1, 0x84, 0x6b, 0xed, 0x34, 0xbb, 0x4e, 0xef, 0xe3, 0x79, 0xfa, 0x99,
	0x84, 0xb0, 0x01, 0x20, 0x0f, 0xda, 0x59, 0xce, 0x78, 0xce, 0xe4, 0x95, 0x09, 0x33, 0x39, 0xfb,
	0x3e, 0xc0, 0x99, 0xcc, 0x59, 0x3a, 0x7a, 0xc6, 0x84, 0x44, 0xef, 0x57, 0x25, 0x34, 0xbb, 0xb6,
	0x49, 0xd4, 0x7f, 0x00, 0xf6, 0xd7, 0x31, 0x27, 0x52, 0xb9, 0xb8, 0x75, 0x17, 0xeb, 0xa0, 0xb1,
	0x69, 0x95, 0x6e, 0xbb, 0xe0, 0x1c, 0x70, 0x1e, 0x53, 0x92, 0xce, 0x73, 0xb5, 0x4b, 0xa7, 0xbf,
	0x2d, 0x68, 0xe9, 0x36, 0xa0, 0x7d, 0x70, 0x84, 0x0a, 0xdd, 0x8f, 0x99, 0x90, 0xa6, 0x6b, 0x1f,
	0xce, 0x97, 0x55, 0xe5, 0x77, 0xb4, 0x82, 0x41, 0x54, 0xd9, 0x7e, 0x05, 0xf6, 0x80, 0xf3, 0x58,
	0xc3, 0x1b, 0x0a, 0xfe, 0xd1, 0x3c, 0xbc, 0x96, 0xd3, 0xd1, 0x0a, 0x6e, 0x17, 0x08, 0x83, 0x86,
	0x61, 0x51, 0x95, 0x86, 0x37, 0x15, 0x7c, 0x7b, 0x1e, 0x3e, 0xa9, 0xfc, 0x68, 0x05, 0xdb, 0xc3,
	0xf2, 0x70, 0xd0, 0x82, 0xd5, 0x97, 0x2c, 0x8d, 0xfc, 0x7f, 0x56, 0x61, 0xab, 0xea, 0xfb, 0x8f,
	0x4c, 0x5e, 0xe8, 0xea, 0x04, 0xfa, 0x19, 0x9c, 0x6c, 0x72, 0x51, 0x4e, 0x6d, 0xef, 0x4d, 0x53,
	0xab, 0xa1, 0x6b, 0x56, 0xf1, 0x34, 0x95, 0xf9, 0x15, 0xae, 0x93, 0xa1, 0xef, 0xa0, 0x45, 0x42,
	0x39, 0x26, 0xb1, 0xdb, 0x50, 0xb4, 0x8f, 0x6f, 0x4a, 0xfb, 0x44, 0xa1, 0x34, 0xa3, 0xa1, 0x40,
	0xdf, 0xc3, 0x7a, 0xa2, 0xaf, 0xdd, 0xa6, 0x62, 0xfb, 0xfc, 0xa6, 0x6c, 0xe6, 0x53, 0xd3, 0x95,
	0x24, 0x68, 0x0f, 0xec, 0x89, 0x20, 0xb8, 0xab, 0xaa, 0xaf, 0x5e, 0xa0, 0x15, 0x21, 0x28, 0x15,
	0x21, 0x38, 0x2f, 0x3d, 0x70, 0xe5, 0x5c, 0x28, 0x40, 0x19, 0xb9, 0xcf, 0x22, 0x77, 0x4d, 0x2b,
	0x40, 0x69, 0x3a, 0x8e, 0xbc, 0x5f, 0x60, 0x73, 0xb6, 0x31, 0x68, 0x13, 0x9a, 0x2f, 0xe9, 0x95,
	0x79, 0xbb, 0xc5, 0x57, 0xd4, 0x2b, 0x37, 0xaf, 0xb1, 0x6c, 0xa5, 0x2a, 0x12, 0xb3, 0x97, 0x5f,
	0x36, 0xf6, 0x2c, 0xef, 0x0c, 0x9c, 0x5a, 0x7f, 0x16, 0x10, 0x07, 0xd3, 0xc4, 0xff, 0xf5, 0xc2,
	0x15, 0xe9, 0x39, 0x74, 0xea, 0x6d, 0x7a, 0x3b, 0xac, 0xfe, 0x6f, 0x16, 0x74, 0x8a, 0x3d, 0x2c,
	0x45, 0x14, 0x9d, 0xc2, 0xed, 0x8c, 0x8c, 0x8a, 0xc7, 0x94, 0x18, 0x3d, 0x35, 0x0f, 0xea, 0xe1,
	0x82, 0xea, 0x17, 0xc9, 0x2f, 0xbe, 0xa5, 0xf1, 0x13, 0x39, 0xde, 0x87, 0xd6, 0x30, 0x27, 0x09,
	0x15, 0x66, 0xc7, 0x1e, 0xde, 0x70, 0x2b, 0xb0, 0x81, 0xf9, 0xbf, 0x5a, 0xe0, 0xe8, 0x1c, 0x95,
	0xdc, 0x16, 0xf2, 0x9d, 0xf0, 0x88, 0xc6, 0xc5, 0x68, 0x75, 0xf9, 0xeb, 0xea, 0x7c, 0x1c, 0xcd,
	0x0e, 0xbe, 0x31, 0x3b, 0x78, 0xf4, 0x0c, 0x4c, 0x7a, 0x7d, 0xae, 0x55, 0xdb, 0x3c, 0xd8, 0x07,
	0x4b, 0xab, 0xab, 0x4b, 0x3c, 0xde, 0xd0, 0x60, 0x73, 0xf4, 0xff, 0xb4, 0xe0, 0xce, 0xa9, 0x56,
	0x40, 0x76, 0x4d, 0x6f, 0x90, 0xdf, 0x3e, 0x74, 0x6a, 0xf9, 0x89, 0xe5, 0x8b, 0x55, 0x69, 0x15,
	0x76, 0xaa, 0xf4, 0xc5, 0x5b, 0xce, 0xff, 0x77, 0x0b, 0x50, 0x3d, 0xff, 0x77, 0x77, 0x07, 0x52,
	0xd8, 0x3a, 0x4e, 0x47, 0x54, 0xc8, 0xf3, 0x9c, 0xb0, 0x94, 0xa5, 0xa3, 0xb2, 0xd9, 0x5b, 0xd0,
	0x7a, 0xc1, 0x07, 0x55, 0xab, 0xd7, 0x5e, 0xf0, 0xc1, 0x71, 0x34, 0x35, 0x83, 0xc6, 0xf4, 0x0c,
	0x76, 0x61, 0x43, 0x5f, 0x5d, 0xd2, 0x5c, 0x30, 0x9e, 0xaa, 0x0e, 0xda, 0xb8, 0xa3, 0x8c, 0xcf,
	0xb5, 0xcd, 0x7f, 0x05, 0xae, 0x8e, 0x57, 0x93, 0x89, 0xff, 0x37, 0x64, 0xef, 0x0f, 0x0b, 0x36,
	0x0e, 0x4d, 0x57, 0x0e, 0x8b, 0xae, 0xa0, 0x6f, 0x60, 0x55, 0xfd, 0xc2, 0x2c, 0xf8, 0x31, 0xaa,
	0xbd, 0x07, 0xef, 0xde, 0xb2, 0x6b, 0x3d, 0x14, 0x7f, 0x05, 0xfd, 0x54, 0xfc, 0x03, 0x28, 0xc7,
	0x8c, 0x76, 0x17, 0x35, 0x7f, 0x66, 0x89, 0xbd, 0x4f, 0xde, 0xec, 0x54, 0x52, 0xf7, 0xfe, 0xb2,
	0x00, 0x4d, 0x65, 0xfd, 0x24, 0x4a, 0x58, 0x8a, 0xce, 0xe0, 0xd6, 0xf4, 0xbc, 0xd0, 0x82, 0x91,
	0x2f, 0x9c, 0xa8, 0x77, 0x77, 0x4e, 0xe3, 0x9f, 0x16, 0xff, 0xfa, 0x54, 0x19, 0x77, 0xe6, 0x86,
	0x82, 0x3e, 0x5d, 0xc6, 0x3b, 0x3f, 0xb9, 0xe5, 0xd4, 0x83, 0x96, 0xb2, 0x3c, 0xfe, 0x37, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x53, 0xb8, 0x92, 0xb0, 0x0a, 0x00, 0x00,
}
