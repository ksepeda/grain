// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitter/dm.proto

package twitterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Target struct {
	RecipientId string `protobuf:"bytes,1,opt,name=recipient_id,json=recipientId" json:"recipient_id,omitempty"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Target) GetRecipientId() string {
	if m != nil {
		return m.RecipientId
	}
	return ""
}

type DirectMessageData struct {
	Text     string    `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Entities *Entities `protobuf:"bytes,2,opt,name=entities" json:"entities,omitempty"`
}

func (m *DirectMessageData) Reset()                    { *m = DirectMessageData{} }
func (m *DirectMessageData) String() string            { return proto.CompactTextString(m) }
func (*DirectMessageData) ProtoMessage()               {}
func (*DirectMessageData) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *DirectMessageData) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *DirectMessageData) GetEntities() *Entities {
	if m != nil {
		return m.Entities
	}
	return nil
}

type DirectMessageCreate struct {
	Target      *Target            `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	SenderId    string             `protobuf:"bytes,2,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	SourceAppId string             `protobuf:"bytes,3,opt,name=source_app_id,json=sourceAppId" json:"source_app_id,omitempty"`
	MessageData *DirectMessageData `protobuf:"bytes,4,opt,name=message_data,json=messageData" json:"message_data,omitempty"`
}

func (m *DirectMessageCreate) Reset()                    { *m = DirectMessageCreate{} }
func (m *DirectMessageCreate) String() string            { return proto.CompactTextString(m) }
func (*DirectMessageCreate) ProtoMessage()               {}
func (*DirectMessageCreate) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *DirectMessageCreate) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DirectMessageCreate) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *DirectMessageCreate) GetSourceAppId() string {
	if m != nil {
		return m.SourceAppId
	}
	return ""
}

func (m *DirectMessageCreate) GetMessageData() *DirectMessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

type DirectMessageEvent struct {
	Type             string               `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Id               string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	CreatedTimestamp string               `protobuf:"bytes,3,opt,name=created_timestamp,json=createdTimestamp" json:"created_timestamp,omitempty"`
	MessageCreate    *DirectMessageCreate `protobuf:"bytes,4,opt,name=message_create,json=messageCreate" json:"message_create,omitempty"`
}

func (m *DirectMessageEvent) Reset()                    { *m = DirectMessageEvent{} }
func (m *DirectMessageEvent) String() string            { return proto.CompactTextString(m) }
func (*DirectMessageEvent) ProtoMessage()               {}
func (*DirectMessageEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *DirectMessageEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DirectMessageEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DirectMessageEvent) GetCreatedTimestamp() string {
	if m != nil {
		return m.CreatedTimestamp
	}
	return ""
}

func (m *DirectMessageEvent) GetMessageCreate() *DirectMessageCreate {
	if m != nil {
		return m.MessageCreate
	}
	return nil
}

func init() {
	proto.RegisterType((*Target)(nil), "grain.twitter.Target")
	proto.RegisterType((*DirectMessageData)(nil), "grain.twitter.DirectMessageData")
	proto.RegisterType((*DirectMessageCreate)(nil), "grain.twitter.DirectMessageCreate")
	proto.RegisterType((*DirectMessageEvent)(nil), "grain.twitter.DirectMessageEvent")
}

func init() { proto.RegisterFile("twitter/dm.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xc2, 0x40,
	0x10, 0x87, 0x49, 0x2a, 0xa2, 0x13, 0x15, 0xdd, 0xb6, 0x54, 0xda, 0x8b, 0xcd, 0x49, 0x90, 0xa6,
	0xa0, 0x4f, 0xd0, 0xaa, 0x87, 0x1c, 0x7a, 0x09, 0x9e, 0x4a, 0x21, 0xac, 0xd9, 0x41, 0xf6, 0x90,
	0xb8, 0xec, 0x4e, 0xff, 0xf8, 0x56, 0x7d, 0x8e, 0x3e, 0x55, 0x71, 0x77, 0x4d, 0xab, 0x85, 0xde,
	0xb2, 0xbf, 0xf9, 0x98, 0xf9, 0x66, 0x08, 0xf4, 0xe9, 0x5d, 0x12, 0xa1, 0xbe, 0x17, 0x65, 0xa2,
	0xf4, 0x96, 0xb6, 0xac, 0xbb, 0xd1, 0x5c, 0x56, 0x89, 0xcf, 0xaf, 0x2f, 0x0e, 0x00, 0x56, 0x24,
	0x69, 0xe7, 0xa0, 0x78, 0x02, 0xcd, 0x15, 0xd7, 0x1b, 0x24, 0x76, 0x0b, 0x1d, 0x8d, 0x85, 0x54,
	0x12, 0x2b, 0xca, 0xa5, 0x18, 0x06, 0xa3, 0x60, 0xdc, 0xce, 0xa2, 0x3a, 0x4b, 0x45, 0xfc, 0x02,
	0x83, 0x85, 0xd4, 0x58, 0xd0, 0x13, 0x1a, 0xc3, 0x37, 0xb8, 0xe0, 0xc4, 0x19, 0x83, 0x06, 0xe1,
	0x07, 0x79, 0xde, 0x7e, 0xb3, 0x19, 0xb4, 0xec, 0x14, 0x89, 0x66, 0x18, 0x8e, 0x82, 0x71, 0x34,
	0xbd, 0x4a, 0x8e, 0x6c, 0x92, 0xa5, 0x2f, 0x67, 0x35, 0x18, 0x7f, 0x05, 0x70, 0x7e, 0xd4, 0x7e,
	0xae, 0x91, 0x13, 0xb2, 0x3b, 0x68, 0x92, 0x55, 0xb4, 0x23, 0xa2, 0xe9, 0xe5, 0x49, 0x2b, 0xe7,
	0x9f, 0x79, 0x88, 0xdd, 0x40, 0xdb, 0x60, 0x25, 0x50, 0xef, 0x97, 0x08, 0xad, 0x54, 0xcb, 0x05,
	0xa9, 0x60, 0x31, 0x74, 0xcd, 0xf6, 0x55, 0x17, 0x98, 0x73, 0xa5, 0xf6, 0xc0, 0x99, 0xdb, 0xd2,
	0x85, 0x0f, 0x4a, 0xa5, 0x82, 0xcd, 0xa1, 0x53, 0x3a, 0x81, 0x5c, 0x70, 0xe2, 0xc3, 0x86, 0x9d,
	0x3a, 0x3a, 0x99, 0xfa, 0xe7, 0x10, 0x59, 0x54, 0xfe, 0x3c, 0xe2, 0xcf, 0x00, 0xd8, 0x11, 0xb2,
	0x7c, 0xc3, 0x8a, 0xec, 0xb1, 0x76, 0x0a, 0xeb, 0x63, 0xed, 0x14, 0xb2, 0x1e, 0x84, 0xb5, 0x69,
	0x28, 0x05, 0x9b, 0xc0, 0xa0, 0xb0, 0x9b, 0x8b, 0x9c, 0x64, 0x89, 0x86, 0x78, 0xa9, 0xbc, 0x67,
	0xdf, 0x17, 0x56, 0x87, 0x9c, 0xa5, 0xd0, 0x3b, 0xc8, 0xba, 0x9a, 0xd7, 0x8d, 0xff, 0xd3, 0x75,
	0x87, 0xcd, 0xba, 0xe5, 0xef, 0xe7, 0x63, 0xf4, 0xdc, 0xf6, 0xb4, 0x5a, 0xaf, 0x9b, 0xf6, 0xf7,
	0x98, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xab, 0xe9, 0x16, 0x97, 0x57, 0x02, 0x00, 0x00,
}
