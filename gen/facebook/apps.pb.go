// Code generated by protoc-gen-go. DO NOT EDIT.
// source: facebook/apps.proto

package grain_facebook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InstalledApps struct {
	InstalledApps []*InstalledApps_App `protobuf:"bytes,1,rep,name=installed_apps,json=installedApps" json:"installed_apps,omitempty"`
}

func (m *InstalledApps) Reset()                    { *m = InstalledApps{} }
func (m *InstalledApps) String() string            { return proto.CompactTextString(m) }
func (*InstalledApps) ProtoMessage()               {}
func (*InstalledApps) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *InstalledApps) GetInstalledApps() []*InstalledApps_App {
	if m != nil {
		return m.InstalledApps
	}
	return nil
}

type InstalledApps_App struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	TimeAdded int64  `protobuf:"varint,2,opt,name=time_added,json=timeAdded" json:"time_added,omitempty"`
}

func (m *InstalledApps_App) Reset()                    { *m = InstalledApps_App{} }
func (m *InstalledApps_App) String() string            { return proto.CompactTextString(m) }
func (*InstalledApps_App) ProtoMessage()               {}
func (*InstalledApps_App) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *InstalledApps_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstalledApps_App) GetTimeAdded() int64 {
	if m != nil {
		return m.TimeAdded
	}
	return 0
}

// Stored as an array, not a top level message
type PostsFromApps struct {
	Posts []*PostsFromApps_Post `protobuf:"bytes,1,rep,name=posts" json:"posts,omitempty"`
}

func (m *PostsFromApps) Reset()                    { *m = PostsFromApps{} }
func (m *PostsFromApps) String() string            { return proto.CompactTextString(m) }
func (*PostsFromApps) ProtoMessage()               {}
func (*PostsFromApps) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PostsFromApps) GetPosts() []*PostsFromApps_Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type PostsFromApps_ExternalContext struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *PostsFromApps_ExternalContext) Reset()         { *m = PostsFromApps_ExternalContext{} }
func (m *PostsFromApps_ExternalContext) String() string { return proto.CompactTextString(m) }
func (*PostsFromApps_ExternalContext) ProtoMessage()    {}
func (*PostsFromApps_ExternalContext) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{1, 0}
}

func (m *PostsFromApps_ExternalContext) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostsFromApps_ExternalContext) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type PostsFromApps_Data struct {
	ExternalContext *PostsFromApps_ExternalContext `protobuf:"bytes,1,opt,name=external_context,json=externalContext" json:"external_context,omitempty"`
}

func (m *PostsFromApps_Data) Reset()                    { *m = PostsFromApps_Data{} }
func (m *PostsFromApps_Data) String() string            { return proto.CompactTextString(m) }
func (*PostsFromApps_Data) ProtoMessage()               {}
func (*PostsFromApps_Data) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 1} }

func (m *PostsFromApps_Data) GetExternalContext() *PostsFromApps_ExternalContext {
	if m != nil {
		return m.ExternalContext
	}
	return nil
}

type PostsFromApps_Attachment struct {
	Data []*PostsFromApps_Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *PostsFromApps_Attachment) Reset()                    { *m = PostsFromApps_Attachment{} }
func (m *PostsFromApps_Attachment) String() string            { return proto.CompactTextString(m) }
func (*PostsFromApps_Attachment) ProtoMessage()               {}
func (*PostsFromApps_Attachment) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 2} }

func (m *PostsFromApps_Attachment) GetData() []*PostsFromApps_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type PostsFromApps_Post struct {
	Timestamp   int64                       `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Attachments []*PostsFromApps_Attachment `protobuf:"bytes,2,rep,name=attachments" json:"attachments,omitempty"`
}

func (m *PostsFromApps_Post) Reset()                    { *m = PostsFromApps_Post{} }
func (m *PostsFromApps_Post) String() string            { return proto.CompactTextString(m) }
func (*PostsFromApps_Post) ProtoMessage()               {}
func (*PostsFromApps_Post) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 3} }

func (m *PostsFromApps_Post) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PostsFromApps_Post) GetAttachments() []*PostsFromApps_Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

type YourApps struct {
	AdminedApps []*YourApps_App `protobuf:"bytes,1,rep,name=admined_apps,json=adminedApps" json:"admined_apps,omitempty"`
}

func (m *YourApps) Reset()                    { *m = YourApps{} }
func (m *YourApps) String() string            { return proto.CompactTextString(m) }
func (*YourApps) ProtoMessage()               {}
func (*YourApps) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *YourApps) GetAdminedApps() []*YourApps_App {
	if m != nil {
		return m.AdminedApps
	}
	return nil
}

type YourApps_App struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	TimeAdded int64  `protobuf:"varint,2,opt,name=time_added,json=timeAdded" json:"time_added,omitempty"`
}

func (m *YourApps_App) Reset()                    { *m = YourApps_App{} }
func (m *YourApps_App) String() string            { return proto.CompactTextString(m) }
func (*YourApps_App) ProtoMessage()               {}
func (*YourApps_App) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2, 0} }

func (m *YourApps_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *YourApps_App) GetTimeAdded() int64 {
	if m != nil {
		return m.TimeAdded
	}
	return 0
}

func init() {
	proto.RegisterType((*InstalledApps)(nil), "grain.facebook.InstalledApps")
	proto.RegisterType((*InstalledApps_App)(nil), "grain.facebook.InstalledApps.App")
	proto.RegisterType((*PostsFromApps)(nil), "grain.facebook.PostsFromApps")
	proto.RegisterType((*PostsFromApps_ExternalContext)(nil), "grain.facebook.PostsFromApps.ExternalContext")
	proto.RegisterType((*PostsFromApps_Data)(nil), "grain.facebook.PostsFromApps.Data")
	proto.RegisterType((*PostsFromApps_Attachment)(nil), "grain.facebook.PostsFromApps.Attachment")
	proto.RegisterType((*PostsFromApps_Post)(nil), "grain.facebook.PostsFromApps.Post")
	proto.RegisterType((*YourApps)(nil), "grain.facebook.YourApps")
	proto.RegisterType((*YourApps_App)(nil), "grain.facebook.YourApps.App")
}

func init() { proto.RegisterFile("facebook/apps.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x49, 0x93, 0x7b, 0xb9, 0x3d, 0xb9, 0xfd, 0xc3, 0xb8, 0x09, 0xa1, 0x42, 0xcd, 0x2a,
	0x1b, 0x23, 0x54, 0xd0, 0xee, 0x24, 0x58, 0x45, 0x5d, 0xc9, 0xac, 0x74, 0x55, 0x4f, 0x9b, 0x51,
	0x83, 0xc9, 0xcc, 0x90, 0x99, 0x42, 0x5f, 0xc0, 0x27, 0x70, 0xeb, 0xc3, 0xca, 0x4c, 0x53, 0x6d,
	0x8a, 0x14, 0xc1, 0xdd, 0xe4, 0x9b, 0xf3, 0x9d, 0xf3, 0xcb, 0x99, 0x0f, 0xf6, 0x1e, 0x71, 0xce,
	0x66, 0x42, 0xbc, 0x1c, 0xa1, 0x94, 0x2a, 0x91, 0x95, 0xd0, 0x82, 0x74, 0x9f, 0x2a, 0xcc, 0x79,
	0xb2, 0xbe, 0x8a, 0xde, 0x1c, 0xe8, 0x5c, 0x73, 0xa5, 0xb1, 0x28, 0x58, 0x96, 0x4a, 0xa9, 0xc8,
	0x15, 0x74, 0xf3, 0xb5, 0x30, 0x35, 0xce, 0xc0, 0x19, 0xba, 0xb1, 0x3f, 0x3a, 0x48, 0x9a, 0xd6,
	0xa4, 0x61, 0x4b, 0x52, 0x29, 0x69, 0x27, 0xdf, 0x94, 0xc2, 0x31, 0xb8, 0xa9, 0x94, 0x84, 0x80,
	0xc7, 0xb1, 0x64, 0x81, 0x33, 0x74, 0xe2, 0x36, 0xb5, 0x67, 0xb2, 0x0f, 0xa0, 0xf3, 0x92, 0x4d,
	0x31, 0xcb, 0x58, 0x16, 0xb4, 0x86, 0x4e, 0xec, 0xd2, 0xb6, 0x51, 0x52, 0x23, 0x44, 0xef, 0x2e,
	0x74, 0x6e, 0x85, 0xd2, 0xea, 0xb2, 0x12, 0xa5, 0xa5, 0x1a, 0xc3, 0x1f, 0x69, 0x84, 0x1a, 0x26,
	0xda, 0x86, 0x69, 0x54, 0xdb, 0x2f, 0xba, 0x32, 0x84, 0xa7, 0xd0, 0xbb, 0x58, 0x6a, 0x56, 0x71,
	0x2c, 0xce, 0x05, 0xd7, 0x6c, 0xa9, 0xbf, 0x25, 0xea, 0x83, 0xbb, 0xa8, 0x0a, 0x8b, 0xd2, 0xa6,
	0xe6, 0x18, 0x3e, 0x80, 0x37, 0x41, 0x8d, 0xe4, 0x0e, 0xfa, 0xac, 0x6e, 0x30, 0x9d, 0xaf, 0x3a,
	0x58, 0xa7, 0x3f, 0x3a, 0xdc, 0x4d, 0xb1, 0x35, 0x96, 0xf6, 0x58, 0x53, 0x08, 0x27, 0x00, 0xa9,
	0xd6, 0x38, 0x7f, 0x2e, 0x19, 0xd7, 0xe4, 0x04, 0xbc, 0x0c, 0x35, 0xfe, 0xec, 0x0f, 0x0d, 0x19,
	0xb5, 0xf5, 0xa1, 0x04, 0xcf, 0xdc, 0x91, 0x01, 0xd8, 0x0d, 0x2a, 0x8d, 0xa5, 0xb4, 0x80, 0xf5,
	0x4a, 0xad, 0x40, 0x6e, 0xc0, 0xc7, 0xcf, 0x59, 0x2a, 0x68, 0xd9, 0x21, 0xf1, 0xee, 0x21, 0x5f,
	0x70, 0x74, 0xd3, 0x1c, 0xbd, 0x3a, 0xf0, 0xef, 0x5e, 0x2c, 0x2a, 0xfb, 0x32, 0x67, 0xf0, 0x1f,
	0xb3, 0x32, 0xe7, 0xcd, 0xb4, 0x0c, 0xb6, 0x3b, 0xaf, 0xeb, 0x6d, 0x50, 0xfc, 0xda, 0xf1, 0xbb,
	0x98, 0xcc, 0xfe, 0xda, 0x4c, 0x1f, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xec, 0xc1, 0x50,
	0xea, 0x02, 0x00, 0x00,
}
