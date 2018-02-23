// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package instagram_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A model representing an instagram user connected through oauth
type User struct {
	// The ID of the partner owning this account group
	PartnerId string `protobuf:"bytes,1,opt,name=partner_id,json=partnerId" json:"partner_id,omitempty"`
	// The ID of the account group authing this user
	AccountGroupId string `protobuf:"bytes,2,opt,name=account_group_id,json=accountGroupId" json:"account_group_id,omitempty"`
	// A unique internal ID for an instagram user
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// A unique ID for an instagram user
	InstagramUserId string `protobuf:"bytes,4,opt,name=instagram_user_id,json=instagramUserId" json:"instagram_user_id,omitempty"`
	// A name for an instagram user
	Username string `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
	// A URL to a profile picture
	ProfilePicture string `protobuf:"bytes,6,opt,name=profile_picture,json=profilePicture" json:"profile_picture,omitempty"`
	// A full name for this instagram user
	FullName string `protobuf:"bytes,7,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	// The bio set by the user in instagram
	Bio string `protobuf:"bytes,8,opt,name=bio" json:"bio,omitempty"`
	// A link to the user's website
	Website string `protobuf:"bytes,9,opt,name=website" json:"website,omitempty"`
	// Whether or not this account is a business
	IsBusiness bool `protobuf:"varint,10,opt,name=is_business,json=isBusiness" json:"is_business,omitempty"`
	// The access token for this user
	AccessToken string `protobuf:"bytes,11,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *User) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

func (m *User) GetAccountGroupId() string {
	if m != nil {
		return m.AccountGroupId
	}
	return ""
}

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetInstagramUserId() string {
	if m != nil {
		return m.InstagramUserId
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetProfilePicture() string {
	if m != nil {
		return m.ProfilePicture
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *User) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *User) GetIsBusiness() bool {
	if m != nil {
		return m.IsBusiness
	}
	return false
}

func (m *User) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "instagram.v1.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0x1f, 0x9a, 0x64, 0x52, 0xd1, 0xe2, 0x0d, 0x16, 0x08, 0x51, 0xd8, 0x10, 0xb1,
	0xa8, 0x84, 0xb8, 0x01, 0x1b, 0x94, 0x0d, 0x42, 0x15, 0x5d, 0x5b, 0x4e, 0x32, 0xad, 0x2c, 0x52,
	0x3b, 0xf2, 0xd8, 0x70, 0x63, 0xce, 0x81, 0xec, 0xa4, 0xd9, 0xf9, 0x7d, 0xfe, 0xfc, 0xc6, 0x1a,
	0x00, 0x4f, 0x68, 0xb7, 0x9d, 0x35, 0xce, 0xb0, 0xa5, 0xd2, 0xe4, 0xe4, 0xd1, 0xca, 0xd3, 0xf6,
	0xe7, 0xe5, 0xf1, 0x6f, 0x0a, 0xf3, 0x3d, 0xa1, 0x65, 0x77, 0x00, 0x9d, 0xb4, 0x4e, 0xa3, 0x15,
	0xaa, 0xe1, 0x93, 0xcd, 0xa4, 0xc8, 0x76, 0xd9, 0x40, 0xca, 0x86, 0x15, 0xb0, 0x96, 0x75, 0x6d,
	0xbc, 0x76, 0xe2, 0x68, 0x8d, 0xef, 0x82, 0x34, 0x8d, 0xd2, 0xe5, 0xc0, 0xdf, 0x03, 0x2e, 0x1b,
	0x76, 0x0d, 0x49, 0x98, 0x16, 0x84, 0x59, 0x14, 0x16, 0x21, 0x96, 0x0d, 0x7b, 0x86, 0xab, 0x71,
	0xb4, 0x38, 0x2b, 0xf3, 0xa8, 0xac, 0xc6, 0x8b, 0x7d, 0xef, 0xde, 0x40, 0x1a, 0x0c, 0x2d, 0x4f,
	0xc8, 0x2f, 0xa2, 0x32, 0x66, 0xf6, 0x04, 0xab, 0xce, 0x9a, 0x83, 0x6a, 0x51, 0x74, 0xaa, 0x76,
	0xde, 0x22, 0x5f, 0xf4, 0x3f, 0x19, 0xf0, 0x67, 0x4f, 0xd9, 0x2d, 0x64, 0x07, 0xdf, 0xb6, 0x22,
	0xb6, 0x24, 0x7d, 0x4b, 0x00, 0x1f, 0xa1, 0x65, 0x0d, 0xb3, 0x4a, 0x19, 0x9e, 0x46, 0x1c, 0x8e,
	0x8c, 0x43, 0xf2, 0x8b, 0x15, 0x29, 0x87, 0x3c, 0x8b, 0xf4, 0x1c, 0xd9, 0x3d, 0xe4, 0x8a, 0x44,
	0xe5, 0x49, 0x69, 0x24, 0xe2, 0xb0, 0x99, 0x14, 0xe9, 0x0e, 0x14, 0xbd, 0x0d, 0x84, 0x3d, 0xc0,
	0x52, 0xd6, 0x35, 0x12, 0x09, 0x67, 0xbe, 0x51, 0xf3, 0x3c, 0xbe, 0xcf, 0x7b, 0xf6, 0x15, 0x50,
	0xb5, 0x88, 0xdb, 0x7f, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xba, 0x68, 0x38, 0x8d, 0x8b, 0x01,
	0x00, 0x00,
}