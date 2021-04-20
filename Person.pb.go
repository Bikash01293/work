// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SocialFollowers struct {
	Youtube              int32    `protobuf:"varint,1,opt,name=youtube,proto3" json:"youtube,omitempty"`
	Twitter              int32    `protobuf:"varint,2,opt,name=twitter,proto3" json:"twitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocialFollowers) Reset()         { *m = SocialFollowers{} }
func (m *SocialFollowers) String() string { return proto.CompactTextString(m) }
func (*SocialFollowers) ProtoMessage()    {}
func (*SocialFollowers) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *SocialFollowers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocialFollowers.Unmarshal(m, b)
}
func (m *SocialFollowers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocialFollowers.Marshal(b, m, deterministic)
}
func (m *SocialFollowers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocialFollowers.Merge(m, src)
}
func (m *SocialFollowers) XXX_Size() int {
	return xxx_messageInfo_SocialFollowers.Size(m)
}
func (m *SocialFollowers) XXX_DiscardUnknown() {
	xxx_messageInfo_SocialFollowers.DiscardUnknown(m)
}

var xxx_messageInfo_SocialFollowers proto.InternalMessageInfo

func (m *SocialFollowers) GetYoutube() int32 {
	if m != nil {
		return m.Youtube
	}
	return 0
}

func (m *SocialFollowers) GetTwitter() int32 {
	if m != nil {
		return m.Twitter
	}
	return 0
}

type PersonDetails struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32            `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	SocialFollowers      *SocialFollowers `protobuf:"bytes,3,opt,name=socialFollowers,proto3" json:"socialFollowers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PersonDetails) Reset()         { *m = PersonDetails{} }
func (m *PersonDetails) String() string { return proto.CompactTextString(m) }
func (*PersonDetails) ProtoMessage()    {}
func (*PersonDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *PersonDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonDetails.Unmarshal(m, b)
}
func (m *PersonDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonDetails.Marshal(b, m, deterministic)
}
func (m *PersonDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonDetails.Merge(m, src)
}
func (m *PersonDetails) XXX_Size() int {
	return xxx_messageInfo_PersonDetails.Size(m)
}
func (m *PersonDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PersonDetails proto.InternalMessageInfo

func (m *PersonDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonDetails) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *PersonDetails) GetSocialFollowers() *SocialFollowers {
	if m != nil {
		return m.SocialFollowers
	}
	return nil
}

func init() {
	proto.RegisterType((*SocialFollowers)(nil), "main.SocialFollowers")
	proto.RegisterType((*PersonDetails)(nil), "main.PersonDetails")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x72,
	0xe5, 0xe2, 0x0f, 0xce, 0x4f, 0xce, 0x4c, 0xcc, 0x71, 0xcb, 0xcf, 0xc9, 0xc9, 0x2f, 0x4f, 0x2d,
	0x2a, 0x16, 0x92, 0xe0, 0x62, 0xaf, 0xcc, 0x2f, 0x2d, 0x29, 0x4d, 0x4a, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0x71, 0x41, 0x32, 0x25, 0xe5, 0x99, 0x25, 0x25, 0xa9, 0x45, 0x12, 0x4c,
	0x10, 0x19, 0x28, 0x57, 0xa9, 0x8c, 0x8b, 0x17, 0x62, 0xb8, 0x4b, 0x6a, 0x49, 0x62, 0x66, 0x4e,
	0xb1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x2e, 0xc4, 0x04, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x80,
	0x8b, 0x39, 0x31, 0x3d, 0x15, 0xaa, 0x15, 0xc4, 0x14, 0xb2, 0xe7, 0xe2, 0x2f, 0x46, 0xb5, 0x5d,
	0x82, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x54, 0x0f, 0xe4, 0x3a, 0x3d, 0x34, 0xa7, 0x05, 0xa1,
	0xab, 0x4e, 0x62, 0x03, 0xfb, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x84, 0x22, 0x3b, 0x13,
	0xdb, 0x00, 0x00, 0x00,
}