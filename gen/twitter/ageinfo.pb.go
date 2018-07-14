// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitter/ageinfo.proto

package twitterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InferredAgeInfo struct {
	Age       []string `protobuf:"bytes,1,rep,name=age" json:"age,omitempty"`
	BirthDate string   `protobuf:"bytes,2,opt,name=birthDate" json:"birthDate,omitempty"`
}

func (m *InferredAgeInfo) Reset()                    { *m = InferredAgeInfo{} }
func (m *InferredAgeInfo) String() string            { return proto.CompactTextString(m) }
func (*InferredAgeInfo) ProtoMessage()               {}
func (*InferredAgeInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *InferredAgeInfo) GetAge() []string {
	if m != nil {
		return m.Age
	}
	return nil
}

func (m *InferredAgeInfo) GetBirthDate() string {
	if m != nil {
		return m.BirthDate
	}
	return ""
}

type AgeMeta struct {
	InferredAgeInfo *InferredAgeInfo `protobuf:"bytes,1,opt,name=inferredAgeInfo" json:"inferredAgeInfo,omitempty"`
}

func (m *AgeMeta) Reset()                    { *m = AgeMeta{} }
func (m *AgeMeta) String() string            { return proto.CompactTextString(m) }
func (*AgeMeta) ProtoMessage()               {}
func (*AgeMeta) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *AgeMeta) GetInferredAgeInfo() *InferredAgeInfo {
	if m != nil {
		return m.InferredAgeInfo
	}
	return nil
}

type AgeinfoEntry struct {
	AgeMeta *AgeMeta `protobuf:"bytes,1,opt,name=ageMeta" json:"ageMeta,omitempty"`
}

func (m *AgeinfoEntry) Reset()                    { *m = AgeinfoEntry{} }
func (m *AgeinfoEntry) String() string            { return proto.CompactTextString(m) }
func (*AgeinfoEntry) ProtoMessage()               {}
func (*AgeinfoEntry) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *AgeinfoEntry) GetAgeMeta() *AgeMeta {
	if m != nil {
		return m.AgeMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*InferredAgeInfo)(nil), "grain.twitter.InferredAgeInfo")
	proto.RegisterType((*AgeMeta)(nil), "grain.twitter.AgeMeta")
	proto.RegisterType((*AgeinfoEntry)(nil), "grain.twitter.AgeinfoEntry")
}

func init() { proto.RegisterFile("twitter/ageinfo.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0xd2, 0x4f, 0x4c, 0x4f, 0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4d, 0x2f, 0x4a, 0xcc, 0xcc, 0xd3, 0x83, 0x4a, 0x2a, 0x39, 0x72, 0xf1, 0x7b,
	0xe6, 0xa5, 0xa5, 0x16, 0x15, 0xa5, 0xa6, 0x38, 0xa6, 0xa7, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x09,
	0x70, 0x31, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x81, 0x98, 0x42, 0x32,
	0x5c, 0x9c, 0x49, 0x99, 0x45, 0x25, 0x19, 0x2e, 0x89, 0x25, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x08, 0x01, 0xa5, 0x60, 0x2e, 0x76, 0xc7, 0xf4, 0x54, 0xdf, 0xd4, 0x92, 0x44, 0x21,
	0x0f, 0x2e, 0xfe, 0x4c, 0x54, 0xd3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xe4, 0xf4, 0x50,
	0xac, 0xd5, 0x43, 0xb3, 0x33, 0x08, 0x5d, 0x9b, 0x92, 0x03, 0x17, 0x8f, 0x23, 0xc4, 0xdd, 0xae,
	0x79, 0x25, 0x45, 0x95, 0x42, 0x06, 0x5c, 0xec, 0x89, 0x10, 0x4b, 0xa0, 0x26, 0x8a, 0xa1, 0x99,
	0x08, 0x75, 0x42, 0x10, 0x4c, 0x99, 0x13, 0x77, 0x14, 0x27, 0x54, 0xae, 0x20, 0x29, 0x89, 0x0d,
	0xec, 0x79, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x68, 0xad, 0x93, 0x15, 0x01, 0x00,
	0x00,
}
