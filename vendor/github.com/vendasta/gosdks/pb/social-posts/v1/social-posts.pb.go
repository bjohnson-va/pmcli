// Code generated by protoc-gen-go.
// source: social-posts.proto
// DO NOT EDIT!

package socialposts_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PostingStatus int32

const (
	PostingStatus_POSTING_IN_PROGRESS PostingStatus = 0
	PostingStatus_POSTING_FAILED      PostingStatus = 1
	PostingStatus_POSTING_COMPLETED   PostingStatus = 2
)

var PostingStatus_name = map[int32]string{
	0: "POSTING_IN_PROGRESS",
	1: "POSTING_FAILED",
	2: "POSTING_COMPLETED",
}
var PostingStatus_value = map[string]int32{
	"POSTING_IN_PROGRESS": 0,
	"POSTING_FAILED":      1,
	"POSTING_COMPLETED":   2,
}

func (x PostingStatus) String() string {
	return proto.EnumName(PostingStatus_name, int32(x))
}
func (PostingStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type SocialPost_DeletionStatus int32

const (
	SocialPost_NONE        SocialPost_DeletionStatus = 0
	SocialPost_FAILED      SocialPost_DeletionStatus = 1
	SocialPost_IN_PROGRESS SocialPost_DeletionStatus = 2
)

var SocialPost_DeletionStatus_name = map[int32]string{
	0: "NONE",
	1: "FAILED",
	2: "IN_PROGRESS",
}
var SocialPost_DeletionStatus_value = map[string]int32{
	"NONE":        0,
	"FAILED":      1,
	"IN_PROGRESS": 2,
}

func (x SocialPost_DeletionStatus) String() string {
	return proto.EnumName(SocialPost_DeletionStatus_name, int32(x))
}
func (SocialPost_DeletionStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type SocialPost_Service int32

const (
	SocialPost_TWITTER            SocialPost_Service = 0
	SocialPost_FACEBOOK           SocialPost_Service = 1
	SocialPost_LINKED_IN          SocialPost_Service = 2
	SocialPost_GOOGLE_PLUS        SocialPost_Service = 3
	SocialPost_GOOGLE_MY_BUSINESS SocialPost_Service = 4
	SocialPost_UNKNOWN            SocialPost_Service = 5
)

var SocialPost_Service_name = map[int32]string{
	0: "TWITTER",
	1: "FACEBOOK",
	2: "LINKED_IN",
	3: "GOOGLE_PLUS",
	4: "GOOGLE_MY_BUSINESS",
	5: "UNKNOWN",
}
var SocialPost_Service_value = map[string]int32{
	"TWITTER":            0,
	"FACEBOOK":           1,
	"LINKED_IN":          2,
	"GOOGLE_PLUS":        3,
	"GOOGLE_MY_BUSINESS": 4,
	"UNKNOWN":            5,
}

func (x SocialPost_Service) String() string {
	return proto.EnumName(SocialPost_Service_name, int32(x))
}
func (SocialPost_Service) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

// The information about a social post. It does not necessarily need to have been posted yet
type SocialPost struct {
	// The unique identifier of the business this post is related to
	BusinessId   string `protobuf:"bytes,1,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
	SocialPostId string `protobuf:"bytes,2,opt,name=social_post_id,json=socialPostId" json:"social_post_id,omitempty"`
	PostText     string `protobuf:"bytes,3,opt,name=post_text,json=postText" json:"post_text,omitempty"`
	// The date time that this post was posted to it's corresponding social account
	Posted         *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=posted" json:"posted,omitempty"`
	IsError        bool                       `protobuf:"varint,5,opt,name=is_error,json=isError" json:"is_error,omitempty"`
	DeletionStatus SocialPost_DeletionStatus  `protobuf:"varint,6,opt,name=deletion_status,json=deletionStatus,enum=socialposts.v1.SocialPost_DeletionStatus" json:"deletion_status,omitempty"`
	Service        SocialPost_Service         `protobuf:"varint,7,opt,name=service,enum=socialposts.v1.SocialPost_Service" json:"service,omitempty"`
	// The link to the post on the social page (e.g. an actual link to Facebook)
	Permalink string `protobuf:"bytes,8,opt,name=permalink" json:"permalink,omitempty"`
	// The date time that this entity was created in our system (not the time that it was posted)
	Created         *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created" json:"created,omitempty"`
	ProfileUrl      string                     `protobuf:"bytes,10,opt,name=profile_url,json=profileUrl" json:"profile_url,omitempty"`
	ProfileImageUrl string                     `protobuf:"bytes,11,opt,name=profile_image_url,json=profileImageUrl" json:"profile_image_url,omitempty"`
	// The date time that this post is scheduled to be posted. If it's in the past that means it we have or are attempting to post this
	// You can see if it was successful based on the status
	Scheduled *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=scheduled" json:"scheduled,omitempty"`
	// The status of attempting to post this social post to the social account
	Status PostingStatus `protobuf:"varint,13,opt,name=status,enum=socialposts.v1.PostingStatus" json:"status,omitempty"`
	// The image posted
	ImageUrl string `protobuf:"bytes,14,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	// The name of the social account (may not always be populated)
	Name string `protobuf:"bytes,15,opt,name=name" json:"name,omitempty"`
	// The username of the social account (may not always be populated)
	Username string `protobuf:"bytes,16,opt,name=username" json:"username,omitempty"`
	// The correlation id for posts that have been posted as a group
	ParentId string `protobuf:"bytes,17,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
}

func (m *SocialPost) Reset()                    { *m = SocialPost{} }
func (m *SocialPost) String() string            { return proto.CompactTextString(m) }
func (*SocialPost) ProtoMessage()               {}
func (*SocialPost) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SocialPost) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

func (m *SocialPost) GetSocialPostId() string {
	if m != nil {
		return m.SocialPostId
	}
	return ""
}

func (m *SocialPost) GetPostText() string {
	if m != nil {
		return m.PostText
	}
	return ""
}

func (m *SocialPost) GetPosted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Posted
	}
	return nil
}

func (m *SocialPost) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *SocialPost) GetDeletionStatus() SocialPost_DeletionStatus {
	if m != nil {
		return m.DeletionStatus
	}
	return SocialPost_NONE
}

func (m *SocialPost) GetService() SocialPost_Service {
	if m != nil {
		return m.Service
	}
	return SocialPost_TWITTER
}

func (m *SocialPost) GetPermalink() string {
	if m != nil {
		return m.Permalink
	}
	return ""
}

func (m *SocialPost) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *SocialPost) GetProfileUrl() string {
	if m != nil {
		return m.ProfileUrl
	}
	return ""
}

func (m *SocialPost) GetProfileImageUrl() string {
	if m != nil {
		return m.ProfileImageUrl
	}
	return ""
}

func (m *SocialPost) GetScheduled() *google_protobuf.Timestamp {
	if m != nil {
		return m.Scheduled
	}
	return nil
}

func (m *SocialPost) GetStatus() PostingStatus {
	if m != nil {
		return m.Status
	}
	return PostingStatus_POSTING_IN_PROGRESS
}

func (m *SocialPost) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *SocialPost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SocialPost) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SocialPost) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

type ListSocialPostsResponse struct {
	// A page of social posts
	SocialPosts []*SocialPost `protobuf:"bytes,1,rep,name=social_posts,json=socialPosts" json:"social_posts,omitempty"`
	// A cursor that can be provided to retrieve the next page of results
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	// Whether or not more results exist
	HasMore bool `protobuf:"varint,3,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
}

func (m *ListSocialPostsResponse) Reset()                    { *m = ListSocialPostsResponse{} }
func (m *ListSocialPostsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSocialPostsResponse) ProtoMessage()               {}
func (*ListSocialPostsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListSocialPostsResponse) GetSocialPosts() []*SocialPost {
	if m != nil {
		return m.SocialPosts
	}
	return nil
}

func (m *ListSocialPostsResponse) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *ListSocialPostsResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

type ListSocialPostsRequest struct {
	// If a start time is provided, a page of the posts that are newer than it will be returned
	// If a start time of the current date time is provided you will get a page of posts that are scheduled to be posted
	Start *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	// If an end time is provided, a page of the posts older than it will be returned
	// If an end time of the current date time is provided you will not get any posts that are scheduled to be posted
	End *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	// The business to list the social posts for
	BusinessId string `protobuf:"bytes,3,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
	PartnerId  string `protobuf:"bytes,4,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	// A cursor that can be provided to retrieve the next page of results
	Cursor string `protobuf:"bytes,5,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *ListSocialPostsRequest) Reset()                    { *m = ListSocialPostsRequest{} }
func (m *ListSocialPostsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSocialPostsRequest) ProtoMessage()               {}
func (*ListSocialPostsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ListSocialPostsRequest) GetStart() *google_protobuf.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ListSocialPostsRequest) GetEnd() *google_protobuf.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func (m *ListSocialPostsRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

func (m *ListSocialPostsRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *ListSocialPostsRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

type Error struct {
	Reason string `protobuf:"bytes,1,opt,name=reason" json:"reason,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type SchedulePostStatus struct {
	// The ID of this post that can be used to get it again
	SocialPostId string `protobuf:"bytes,1,opt,name=social_post_id,json=socialPostId" json:"social_post_id,omitempty"`
	// The ID of the social page this post was posted to
	SocialServiceId string `protobuf:"bytes,2,opt,name=social_service_id,json=socialServiceId" json:"social_service_id,omitempty"`
	// The Name or Username of the social page this post was posted to (may not always be populated)
	SocialServiceLabel string `protobuf:"bytes,3,opt,name=social_service_label,json=socialServiceLabel" json:"social_service_label,omitempty"`
	// The reason why there was an error scheduling the post. Error will be nil if it was successful
	Error *Error `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *SchedulePostStatus) Reset()                    { *m = SchedulePostStatus{} }
func (m *SchedulePostStatus) String() string            { return proto.CompactTextString(m) }
func (*SchedulePostStatus) ProtoMessage()               {}
func (*SchedulePostStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SchedulePostStatus) GetSocialPostId() string {
	if m != nil {
		return m.SocialPostId
	}
	return ""
}

func (m *SchedulePostStatus) GetSocialServiceId() string {
	if m != nil {
		return m.SocialServiceId
	}
	return ""
}

func (m *SchedulePostStatus) GetSocialServiceLabel() string {
	if m != nil {
		return m.SocialServiceLabel
	}
	return ""
}

func (m *SchedulePostStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SocialPostData struct {
	// The text to post, must be provided for all twitter and google plus page posts
	PostText string `protobuf:"bytes,1,opt,name=post_text,json=postText" json:"post_text,omitempty"`
	// The image to post
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	// When this social post should be posted. If not provided the social  will be posted asap
	ScheduleFor *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=schedule_for,json=scheduleFor" json:"schedule_for,omitempty"`
}

func (m *SocialPostData) Reset()                    { *m = SocialPostData{} }
func (m *SocialPostData) String() string            { return proto.CompactTextString(m) }
func (*SocialPostData) ProtoMessage()               {}
func (*SocialPostData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SocialPostData) GetPostText() string {
	if m != nil {
		return m.PostText
	}
	return ""
}

func (m *SocialPostData) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *SocialPostData) GetScheduleFor() *google_protobuf.Timestamp {
	if m != nil {
		return m.ScheduleFor
	}
	return nil
}

type ScheduleToAllPagesRequest struct {
	// The social post to post
	SocialPost *SocialPostData `protobuf:"bytes,1,opt,name=social_post,json=socialPost" json:"social_post,omitempty"`
	PartnerId  string          `protobuf:"bytes,2,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	// The business to post the social post to all of it's connected pages (twitter user, linkedin company, google plus page, facebook page)
	BusinessId string `protobuf:"bytes,3,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *ScheduleToAllPagesRequest) Reset()                    { *m = ScheduleToAllPagesRequest{} }
func (m *ScheduleToAllPagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleToAllPagesRequest) ProtoMessage()               {}
func (*ScheduleToAllPagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ScheduleToAllPagesRequest) GetSocialPost() *SocialPostData {
	if m != nil {
		return m.SocialPost
	}
	return nil
}

func (m *ScheduleToAllPagesRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *ScheduleToAllPagesRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

type ScheduleToAllPagesResponse struct {
	// The statuses of the posts that were attempted to be scheduled
	Statuses []*SchedulePostStatus `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *ScheduleToAllPagesResponse) Reset()                    { *m = ScheduleToAllPagesResponse{} }
func (m *ScheduleToAllPagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ScheduleToAllPagesResponse) ProtoMessage()               {}
func (*ScheduleToAllPagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ScheduleToAllPagesResponse) GetStatuses() []*SchedulePostStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type SchedulePostRequest struct {
	// The social post to post
	SocialPost *SocialPostData `protobuf:"bytes,1,opt,name=social_post,json=socialPost" json:"social_post,omitempty"`
	// The social accounts to post this post to
	SocialServiceIds []string `protobuf:"bytes,2,rep,name=social_service_ids,json=socialServiceIds" json:"social_service_ids,omitempty"`
	PartnerId        string   `protobuf:"bytes,3,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	// The business that has the social accounts connected to it
	BusinessId string `protobuf:"bytes,4,opt,name=business_id,json=businessId" json:"business_id,omitempty"`
}

func (m *SchedulePostRequest) Reset()                    { *m = SchedulePostRequest{} }
func (m *SchedulePostRequest) String() string            { return proto.CompactTextString(m) }
func (*SchedulePostRequest) ProtoMessage()               {}
func (*SchedulePostRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *SchedulePostRequest) GetSocialPost() *SocialPostData {
	if m != nil {
		return m.SocialPost
	}
	return nil
}

func (m *SchedulePostRequest) GetSocialServiceIds() []string {
	if m != nil {
		return m.SocialServiceIds
	}
	return nil
}

func (m *SchedulePostRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *SchedulePostRequest) GetBusinessId() string {
	if m != nil {
		return m.BusinessId
	}
	return ""
}

type SchedulePostResponse struct {
	// The statuses of the posts that were attempted to be scheduled
	Statuses []*SchedulePostStatus `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty"`
}

func (m *SchedulePostResponse) Reset()                    { *m = SchedulePostResponse{} }
func (m *SchedulePostResponse) String() string            { return proto.CompactTextString(m) }
func (*SchedulePostResponse) ProtoMessage()               {}
func (*SchedulePostResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *SchedulePostResponse) GetStatuses() []*SchedulePostStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

// Begin and end of date range
type DateRangeFilter struct {
	BeginRange *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=begin_range,json=beginRange" json:"begin_range,omitempty"`
	EndRange   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=end_range,json=endRange" json:"end_range,omitempty"`
}

func (m *DateRangeFilter) Reset()                    { *m = DateRangeFilter{} }
func (m *DateRangeFilter) String() string            { return proto.CompactTextString(m) }
func (*DateRangeFilter) ProtoMessage()               {}
func (*DateRangeFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DateRangeFilter) GetBeginRange() *google_protobuf.Timestamp {
	if m != nil {
		return m.BeginRange
	}
	return nil
}

func (m *DateRangeFilter) GetEndRange() *google_protobuf.Timestamp {
	if m != nil {
		return m.EndRange
	}
	return nil
}

type PartnerListScheduledSocialPostsRequest struct {
	// A cursor that can be provided to retrieve the next page of results
	Cursor  string                                          `protobuf:"bytes,5,opt,name=cursor" json:"cursor,omitempty"`
	Filters *PartnerListScheduledSocialPostsRequest_Filters `protobuf:"bytes,6,opt,name=filters" json:"filters,omitempty"`
}

func (m *PartnerListScheduledSocialPostsRequest) Reset() {
	*m = PartnerListScheduledSocialPostsRequest{}
}
func (m *PartnerListScheduledSocialPostsRequest) String() string { return proto.CompactTextString(m) }
func (*PartnerListScheduledSocialPostsRequest) ProtoMessage()    {}
func (*PartnerListScheduledSocialPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

func (m *PartnerListScheduledSocialPostsRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *PartnerListScheduledSocialPostsRequest) GetFilters() *PartnerListScheduledSocialPostsRequest_Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

type PartnerListScheduledSocialPostsRequest_Filters struct {
	// Scheduled posts that are scheduled to be posted between begin and end
	// Only one of scheduled_date_filter or created_date_filter can be provided
	DateRange *DateRangeFilter `protobuf:"bytes,1,opt,name=date_range,json=dateRange" json:"date_range,omitempty"`
	PartnerId string           `protobuf:"bytes,3,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
}

func (m *PartnerListScheduledSocialPostsRequest_Filters) Reset() {
	*m = PartnerListScheduledSocialPostsRequest_Filters{}
}
func (m *PartnerListScheduledSocialPostsRequest_Filters) String() string {
	return proto.CompactTextString(m)
}
func (*PartnerListScheduledSocialPostsRequest_Filters) ProtoMessage() {}
func (*PartnerListScheduledSocialPostsRequest_Filters) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11, 0}
}

func (m *PartnerListScheduledSocialPostsRequest_Filters) GetDateRange() *DateRangeFilter {
	if m != nil {
		return m.DateRange
	}
	return nil
}

func (m *PartnerListScheduledSocialPostsRequest_Filters) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type PartnerListScheduledPostsResponse struct {
	// A page of social posts
	SocialPosts []*SocialPost `protobuf:"bytes,1,rep,name=social_posts,json=socialPosts" json:"social_posts,omitempty"`
	// A cursor that can be provided to retrieve the next page of results
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
	// Whether or not more results exist
	HasMore bool `protobuf:"varint,3,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
}

func (m *PartnerListScheduledPostsResponse) Reset()         { *m = PartnerListScheduledPostsResponse{} }
func (m *PartnerListScheduledPostsResponse) String() string { return proto.CompactTextString(m) }
func (*PartnerListScheduledPostsResponse) ProtoMessage()    {}
func (*PartnerListScheduledPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12}
}

func (m *PartnerListScheduledPostsResponse) GetSocialPosts() []*SocialPost {
	if m != nil {
		return m.SocialPosts
	}
	return nil
}

func (m *PartnerListScheduledPostsResponse) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

func (m *PartnerListScheduledPostsResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func init() {
	proto.RegisterType((*SocialPost)(nil), "socialposts.v1.SocialPost")
	proto.RegisterType((*ListSocialPostsResponse)(nil), "socialposts.v1.ListSocialPostsResponse")
	proto.RegisterType((*ListSocialPostsRequest)(nil), "socialposts.v1.ListSocialPostsRequest")
	proto.RegisterType((*Error)(nil), "socialposts.v1.Error")
	proto.RegisterType((*SchedulePostStatus)(nil), "socialposts.v1.SchedulePostStatus")
	proto.RegisterType((*SocialPostData)(nil), "socialposts.v1.SocialPostData")
	proto.RegisterType((*ScheduleToAllPagesRequest)(nil), "socialposts.v1.ScheduleToAllPagesRequest")
	proto.RegisterType((*ScheduleToAllPagesResponse)(nil), "socialposts.v1.ScheduleToAllPagesResponse")
	proto.RegisterType((*SchedulePostRequest)(nil), "socialposts.v1.SchedulePostRequest")
	proto.RegisterType((*SchedulePostResponse)(nil), "socialposts.v1.SchedulePostResponse")
	proto.RegisterType((*DateRangeFilter)(nil), "socialposts.v1.DateRangeFilter")
	proto.RegisterType((*PartnerListScheduledSocialPostsRequest)(nil), "socialposts.v1.PartnerListScheduledSocialPostsRequest")
	proto.RegisterType((*PartnerListScheduledSocialPostsRequest_Filters)(nil), "socialposts.v1.PartnerListScheduledSocialPostsRequest.Filters")
	proto.RegisterType((*PartnerListScheduledPostsResponse)(nil), "socialposts.v1.PartnerListScheduledPostsResponse")
	proto.RegisterEnum("socialposts.v1.PostingStatus", PostingStatus_name, PostingStatus_value)
	proto.RegisterEnum("socialposts.v1.SocialPost_DeletionStatus", SocialPost_DeletionStatus_name, SocialPost_DeletionStatus_value)
	proto.RegisterEnum("socialposts.v1.SocialPost_Service", SocialPost_Service_name, SocialPost_Service_value)
}

func init() { proto.RegisterFile("social-posts.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x1b, 0xed, 0xfa, 0xbf, 0x3f, 0xa7, 0xf6, 0x66, 0xda, 0xa6, 0xdb, 0xf4, 0xd7, 0x5f, 0xcc, 0x0a,
	0xa1, 0x10, 0x8a, 0x53, 0x02, 0xa8, 0x48, 0x40, 0x50, 0x12, 0x3b, 0xd1, 0x2a, 0x8e, 0x6d, 0xad,
	0x1d, 0x0a, 0x12, 0xd2, 0x6a, 0xe2, 0x9d, 0x38, 0x2b, 0xd6, 0xbb, 0x66, 0x66, 0x1c, 0xe5, 0x01,
	0x90, 0x10, 0xf7, 0x5c, 0x22, 0x9e, 0x06, 0xee, 0x79, 0x15, 0xde, 0x00, 0xcd, 0x9f, 0x8d, 0xe3,
	0x75, 0x12, 0xf7, 0xa2, 0x17, 0xdc, 0x79, 0xce, 0x9c, 0x99, 0x9d, 0xef, 0xcc, 0xf9, 0xce, 0x18,
	0x10, 0x8b, 0x87, 0x01, 0x0e, 0x3f, 0x9e, 0xc4, 0x8c, 0xb3, 0xc6, 0x84, 0xc6, 0x3c, 0x46, 0x55,
	0x85, 0x29, 0xe8, 0xf2, 0x93, 0xf5, 0x8d, 0x51, 0x1c, 0x8f, 0x42, 0xb2, 0x2d, 0x67, 0xcf, 0xa6,
	0xe7, 0xdb, 0x3c, 0x18, 0x13, 0xc6, 0xf1, 0x78, 0xa2, 0x16, 0xd8, 0xff, 0x14, 0x00, 0xfa, 0x72,
	0x4d, 0x2f, 0x66, 0x1c, 0x6d, 0x40, 0xe5, 0x6c, 0xca, 0x82, 0x88, 0x30, 0xe6, 0x05, 0xbe, 0x65,
	0xd4, 0x8d, 0xcd, 0xb2, 0x0b, 0x09, 0xe4, 0xf8, 0xe8, 0x7d, 0xd0, 0x9f, 0xf0, 0xc4, 0x37, 0x04,
	0x27, 0x23, 0x39, 0x2b, 0xec, 0x7a, 0x13, 0xc7, 0x47, 0xcf, 0xa1, 0x2c, 0xa7, 0x39, 0xb9, 0xe2,
	0x56, 0x56, 0x12, 0x4a, 0x02, 0x18, 0x90, 0x2b, 0x8e, 0x76, 0xa0, 0x20, 0x7e, 0x13, 0xdf, 0xca,
	0xd5, 0x8d, 0xcd, 0xca, 0xce, 0x7a, 0x43, 0x1d, 0xb2, 0x91, 0x1c, 0xb2, 0x31, 0x48, 0x0e, 0xe9,
	0x6a, 0x26, 0x7a, 0x06, 0xa5, 0x80, 0x79, 0x84, 0xd2, 0x98, 0x5a, 0xf9, 0xba, 0xb1, 0x59, 0x72,
	0x8b, 0x01, 0x6b, 0x89, 0x21, 0x72, 0xa1, 0xe6, 0x93, 0x90, 0xf0, 0x20, 0x8e, 0x3c, 0xc6, 0x31,
	0x9f, 0x32, 0xab, 0x50, 0x37, 0x36, 0xab, 0x3b, 0x1f, 0x36, 0xe6, 0xc5, 0x68, 0xcc, 0xea, 0x6c,
	0x34, 0xf5, 0x8a, 0xbe, 0x5c, 0xe0, 0x56, 0xfd, 0xb9, 0x31, 0xfa, 0x0a, 0x8a, 0x8c, 0xd0, 0xcb,
	0x60, 0x48, 0xac, 0xa2, 0xdc, 0xcb, 0xbe, 0x67, 0xaf, 0xbe, 0x62, 0xba, 0xc9, 0x12, 0xf4, 0x3f,
	0x28, 0x4f, 0x08, 0x1d, 0xe3, 0x30, 0x88, 0x7e, 0xb4, 0x4a, 0xb2, 0xfa, 0x19, 0x80, 0x3e, 0x83,
	0xe2, 0x90, 0x12, 0x2c, 0xea, 0x2f, 0x2f, 0xad, 0x3f, 0xa1, 0x8a, 0x8b, 0x99, 0xd0, 0xf8, 0x3c,
	0x08, 0x89, 0x37, 0xa5, 0xa1, 0x05, 0xea, 0x62, 0x34, 0x74, 0x4a, 0x43, 0xb4, 0x05, 0xab, 0x09,
	0x21, 0x18, 0xe3, 0x91, 0xa2, 0x55, 0x24, 0xad, 0xa6, 0x27, 0x1c, 0x81, 0x0b, 0xee, 0x17, 0x50,
	0x66, 0xc3, 0x0b, 0xe2, 0x4f, 0x43, 0xe2, 0x5b, 0x2b, 0x4b, 0x0f, 0x31, 0x23, 0xa3, 0xcf, 0xa1,
	0xa0, 0x35, 0x7e, 0x28, 0x75, 0x79, 0x91, 0xd6, 0x45, 0x28, 0x12, 0x44, 0x23, 0xad, 0xab, 0x26,
	0x0b, 0x3f, 0xcc, 0x0e, 0x55, 0x55, 0x7e, 0x08, 0x92, 0xd3, 0x20, 0xc8, 0x45, 0x78, 0x4c, 0xac,
	0x9a, 0xc4, 0xe5, 0x6f, 0xb4, 0x0e, 0xa5, 0x29, 0x23, 0x54, 0xe2, 0xa6, 0xe2, 0x27, 0x63, 0x69,
	0x2e, 0x4c, 0x49, 0x24, 0xdd, 0xb7, 0xaa, 0xcd, 0x25, 0x01, 0xc7, 0xb7, 0x5f, 0x43, 0x75, 0xfe,
	0x6e, 0x51, 0x09, 0x72, 0x9d, 0x6e, 0xa7, 0x65, 0x3e, 0x40, 0x00, 0x85, 0xc3, 0x3d, 0xa7, 0xdd,
	0x6a, 0x9a, 0x06, 0xaa, 0x41, 0xc5, 0xe9, 0x78, 0x3d, 0xb7, 0x7b, 0xe4, 0xb6, 0xfa, 0x7d, 0x33,
	0x63, 0x07, 0x50, 0xd4, 0x17, 0x89, 0x2a, 0x50, 0x1c, 0xbc, 0x71, 0x06, 0x83, 0x96, 0x6b, 0x3e,
	0x40, 0x2b, 0x50, 0x3a, 0xdc, 0x3b, 0x68, 0xed, 0x77, 0xbb, 0xc7, 0xa6, 0x81, 0x1e, 0x42, 0xb9,
	0xed, 0x74, 0x8e, 0x5b, 0x4d, 0xcf, 0xe9, 0x98, 0x19, 0xb1, 0xcb, 0x51, 0xb7, 0x7b, 0xd4, 0x6e,
	0x79, 0xbd, 0xf6, 0x69, 0xdf, 0xcc, 0xa2, 0x35, 0x40, 0x1a, 0x38, 0xf9, 0xde, 0xdb, 0x3f, 0xed,
	0x3b, 0x1d, 0xb1, 0x7b, 0x4e, 0x6c, 0x79, 0xda, 0x39, 0xee, 0x74, 0xdf, 0x74, 0xcc, 0xbc, 0xfd,
	0x9b, 0x01, 0x4f, 0xdb, 0x01, 0xe3, 0x33, 0x0f, 0x31, 0x97, 0xb0, 0x49, 0x1c, 0x31, 0x82, 0xbe,
	0x86, 0x95, 0x1b, 0xfd, 0xc5, 0x2c, 0xa3, 0x9e, 0x95, 0xb7, 0x73, 0xa7, 0xfd, 0xdc, 0xca, 0xac,
	0xf3, 0x98, 0xb0, 0x49, 0x44, 0xae, 0xb8, 0x37, 0x9c, 0x52, 0x16, 0x53, 0xdd, 0x9b, 0x20, 0xa0,
	0x03, 0x89, 0x88, 0x46, 0xba, 0xc0, 0xcc, 0x1b, 0xc7, 0x94, 0xc8, 0xc6, 0x2c, 0xb9, 0xc5, 0x0b,
	0xcc, 0x4e, 0x62, 0x4a, 0xec, 0xbf, 0x0d, 0x58, 0x5b, 0x38, 0xd6, 0x4f, 0x53, 0xc2, 0x38, 0x7a,
	0x05, 0x79, 0xc6, 0x31, 0xe5, 0x32, 0x10, 0xee, 0x37, 0x8b, 0x22, 0xa2, 0x97, 0x90, 0x25, 0x91,
	0x0a, 0x87, 0xfb, 0xf9, 0x82, 0x96, 0x8e, 0x9d, 0xec, 0x42, 0xec, 0xbc, 0x00, 0x98, 0x60, 0xca,
	0x23, 0x42, 0xc5, 0x7c, 0x4e, 0xf7, 0x94, 0x42, 0x1c, 0x1f, 0xad, 0x41, 0x41, 0x57, 0x9c, 0x97,
	0x53, 0x7a, 0x64, 0x6f, 0x40, 0x5e, 0x85, 0xc4, 0x1a, 0x14, 0x28, 0xc1, 0x2c, 0x8e, 0x74, 0xa4,
	0xe9, 0x91, 0xfd, 0x97, 0x01, 0xa8, 0xaf, 0xdd, 0x2d, 0x2a, 0xd6, 0x9e, 0x59, 0x4c, 0x39, 0xe3,
	0x96, 0x94, 0xdb, 0x82, 0x55, 0xcd, 0xd2, 0x9d, 0x3f, 0x8b, 0xc3, 0x9a, 0x9a, 0xd0, 0x8e, 0x72,
	0x7c, 0xf4, 0x0a, 0x1e, 0xa7, 0xb8, 0x21, 0x3e, 0x23, 0xa1, 0x2e, 0x15, 0xcd, 0xd1, 0xdb, 0x62,
	0x06, 0x7d, 0x04, 0x79, 0x95, 0x77, 0x2a, 0x25, 0x9f, 0xa4, 0x2d, 0x20, 0x0b, 0x73, 0x15, 0xc7,
	0xfe, 0xd5, 0x80, 0xea, 0xec, 0xde, 0x9a, 0x98, 0xe3, 0xf9, 0x0c, 0x36, 0x52, 0x19, 0x3c, 0xd7,
	0x90, 0x99, 0x54, 0x43, 0x0a, 0x0f, 0x6a, 0x4d, 0xbc, 0xf3, 0x98, 0xca, 0x33, 0xde, 0x7f, 0x89,
	0x95, 0x84, 0x7f, 0x18, 0x53, 0xfb, 0x77, 0x03, 0x9e, 0x25, 0x9a, 0x0e, 0xe2, 0xbd, 0x30, 0xec,
	0xe1, 0x11, 0xb9, 0xb6, 0xd2, 0x37, 0x50, 0xb9, 0x21, 0xad, 0x36, 0xd4, 0xff, 0xef, 0xf6, 0xb7,
	0xa8, 0xc5, 0x85, 0x99, 0xee, 0x29, 0x2b, 0x64, 0xd2, 0x56, 0x58, 0x66, 0x25, 0xfb, 0x07, 0x58,
	0xbf, 0xed, 0x74, 0xba, 0xff, 0x76, 0xa1, 0xa4, 0x32, 0x8b, 0x24, 0xbd, 0xb7, 0x18, 0xfd, 0x0b,
	0x7e, 0x71, 0xaf, 0xd7, 0xd8, 0x7f, 0x1a, 0xf0, 0xe8, 0x26, 0xe1, 0x9d, 0x95, 0xfd, 0x32, 0x79,
	0xef, 0x6f, 0x98, 0x8d, 0x59, 0x99, 0x7a, 0x76, 0xb3, 0xec, 0x9a, 0x29, 0xb7, 0xb1, 0x94, 0x48,
	0xd9, 0x25, 0x22, 0xe5, 0x16, 0x44, 0xfa, 0x16, 0x1e, 0xcf, 0x57, 0xf1, 0x8e, 0xe4, 0xf9, 0xc5,
	0x80, 0x5a, 0x13, 0x73, 0xe2, 0xe2, 0x68, 0x44, 0x0e, 0x83, 0x90, 0x13, 0x8a, 0xbe, 0x84, 0xca,
	0x19, 0x19, 0x05, 0x91, 0x47, 0x05, 0xf8, 0x16, 0x11, 0x03, 0x92, 0x2e, 0xb7, 0x40, 0xaf, 0xa1,
	0x4c, 0x22, 0x5f, 0x2f, 0x5d, 0x9e, 0x36, 0x25, 0x12, 0xf9, 0x72, 0xa1, 0xfd, 0x73, 0x06, 0x3e,
	0xe8, 0x29, 0x41, 0x64, 0xe8, 0x25, 0x4f, 0xdc, 0x2d, 0xe9, 0x77, 0x47, 0xba, 0xa0, 0xef, 0xa0,
	0x78, 0x2e, 0x4b, 0x50, 0xff, 0x38, 0x2a, 0x3b, 0xbb, 0x0b, 0xaf, 0xe1, 0x5b, 0x7d, 0xa0, 0xa1,
	0x84, 0x60, 0x6e, 0xb2, 0xdd, 0xfa, 0x05, 0x14, 0x35, 0x86, 0x76, 0x01, 0x7c, 0xcc, 0xc9, 0x9c,
	0x38, 0x1b, 0xe9, 0xef, 0xa4, 0x24, 0x75, 0xcb, 0x7e, 0x02, 0x2c, 0x71, 0x82, 0xfd, 0x87, 0x01,
	0xef, 0xdd, 0x76, 0xca, 0xff, 0xca, 0xab, 0xb4, 0xd5, 0x87, 0x87, 0x73, 0xff, 0x29, 0xd0, 0x53,
	0x78, 0xd4, 0xeb, 0xf6, 0x07, 0x4e, 0xe7, 0xc8, 0xbb, 0xf9, 0x82, 0x3f, 0x40, 0x08, 0xaa, 0xc9,
	0xc4, 0xf5, 0x33, 0xff, 0x04, 0x56, 0x13, 0xec, 0xa0, 0x7b, 0xd2, 0x6b, 0xb7, 0x06, 0xad, 0xa6,
	0x99, 0xd9, 0xdf, 0x86, 0xe7, 0xc3, 0x78, 0xdc, 0xb8, 0x24, 0x91, 0x8f, 0x19, 0xc7, 0xa9, 0x3a,
	0xf6, 0xcd, 0x59, 0x21, 0x3d, 0xe1, 0x21, 0x76, 0x56, 0x90, 0x5e, 0xfa, 0xf4, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0x97, 0x8d, 0x13, 0x74, 0x0b, 0x00, 0x00,
}
