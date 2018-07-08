// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitter/archive.proto

/*
Package twitterpb is a generated protocol buffer package.

It is generated from these files:
	twitter/archive.proto
	twitter/dm.proto
	twitter/entity.proto
	twitter/list.proto
	twitter/tweet.proto
	twitter/user.proto

It has these top-level messages:
	Archive
	Target
	DirectMessageData
	DirectMessageCreate
	DirectMessageEvent
	Url
	Hashtag
	UserMention
	MediaSize
	MediaSizes
	Variant
	VideoInfo
	AdditionalMediaInfo
	EntityMedia
	Entities
	List
	Place
	Coordinates
	ExtendedTweet
	Geo
	Tweet
	User
*/
package twitterpb

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

type Archive struct {
	Lists          []*List               `protobuf:"bytes,1,rep,name=lists" json:"lists,omitempty"`
	Friends        []*User               `protobuf:"bytes,2,rep,name=friends" json:"friends,omitempty"`
	Followers      []*User               `protobuf:"bytes,3,rep,name=followers" json:"followers,omitempty"`
	Timeline       []*Tweet              `protobuf:"bytes,4,rep,name=timeline" json:"timeline,omitempty"`
	Favorites      []*Tweet              `protobuf:"bytes,5,rep,name=favorites" json:"favorites,omitempty"`
	DirectMessages []*DirectMessageEvent `protobuf:"bytes,6,rep,name=direct_messages,json=directMessages" json:"direct_messages,omitempty"`
}

func (m *Archive) Reset()                    { *m = Archive{} }
func (m *Archive) String() string            { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()               {}
func (*Archive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Archive) GetLists() []*List {
	if m != nil {
		return m.Lists
	}
	return nil
}

func (m *Archive) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *Archive) GetFollowers() []*User {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *Archive) GetTimeline() []*Tweet {
	if m != nil {
		return m.Timeline
	}
	return nil
}

func (m *Archive) GetFavorites() []*Tweet {
	if m != nil {
		return m.Favorites
	}
	return nil
}

func (m *Archive) GetDirectMessages() []*DirectMessageEvent {
	if m != nil {
		return m.DirectMessages
	}
	return nil
}

func init() {
	proto.RegisterType((*Archive)(nil), "grain.twitter.Archive")
}

func init() { proto.RegisterFile("twitter/archive.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc7, 0xf1, 0xd8, 0xda, 0x56, 0xb6, 0xf1, 0x4f, 0xb6, 0x9a, 0x6c, 0x7a, 0x52, 0x4f, 0x7a,
	0x10, 0xb5, 0x3e, 0x81, 0x46, 0x2f, 0x46, 0x2f, 0x44, 0x2f, 0x5e, 0x0c, 0x2d, 0x43, 0x9d, 0x04,
	0xd8, 0x66, 0x66, 0x84, 0x57, 0xf4, 0xb1, 0x4c, 0x59, 0x56, 0x94, 0xa8, 0xd7, 0x1f, 0x9f, 0x2f,
	0x13, 0x50, 0x07, 0x52, 0xa1, 0x08, 0xd0, 0x79, 0x4c, 0x8b, 0x37, 0x2c, 0x21, 0x5c, 0x91, 0x15,
	0xab, 0xb7, 0x97, 0x14, 0x63, 0x11, 0x36, 0x0f, 0xa7, 0x7b, 0x5e, 0x25, 0xb9, 0x03, 0x53, 0xed,
	0x97, 0x0c, 0x59, 0x9a, 0x6d, 0xe2, 0x37, 0xa9, 0x00, 0xa4, 0x0b, 0xdf, 0x19, 0xc8, 0x6d, 0xc7,
	0x1f, 0x3d, 0x35, 0xba, 0x76, 0xf7, 0xf4, 0xa9, 0x1a, 0xac, 0x5f, 0xc1, 0x66, 0xe3, 0xb0, 0x7f,
	0x32, 0x9e, 0x4d, 0xc2, 0x1f, 0x97, 0xc3, 0x07, 0x64, 0x89, 0x9c, 0xd0, 0x67, 0x6a, 0x94, 0x12,
	0x42, 0x91, 0xb0, 0xe9, 0xfd, 0x8a, 0x9f, 0x19, 0x28, 0xf2, 0x46, 0x5f, 0xaa, 0x20, 0xb5, 0x59,
	0x66, 0x2b, 0x20, 0x36, 0xfd, 0xbf, 0x83, 0x56, 0xe9, 0x0b, 0xb5, 0x25, 0x98, 0x43, 0x86, 0x05,
	0x98, 0xcd, 0xba, 0xd8, 0xef, 0x14, 0x4f, 0xeb, 0x4f, 0x8b, 0xbe, 0x94, 0x9e, 0xa9, 0x20, 0x8d,
	0x4b, 0x4b, 0x28, 0xc0, 0x66, 0xf0, 0x4f, 0xd2, 0x32, 0x7d, 0xaf, 0x76, 0x13, 0x24, 0x58, 0xc8,
	0x6b, 0x0e, 0xcc, 0xf1, 0x12, 0xd8, 0x0c, 0xeb, 0xf2, 0xa8, 0x53, 0xde, 0xd6, 0xea, 0xd1, 0xa1,
	0xbb, 0x12, 0x0a, 0x89, 0x76, 0x92, 0xef, 0x1b, 0xdf, 0x8c, 0x5f, 0x82, 0x46, 0xaf, 0xe6, 0xf3,
	0x61, 0xfd, 0x7b, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x31, 0xe5, 0x74, 0xd5, 0x01,
	0x00, 0x00,
}
