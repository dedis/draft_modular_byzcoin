// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pedersen

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

type Deal struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Deal)(nil), "pedersen.Deal")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x48, 0x4d, 0x49,
	0x2d, 0x2a, 0x4e, 0xcd, 0x53, 0x92, 0xe2, 0x62, 0x71, 0x49, 0x4d, 0xcc, 0x11, 0x12, 0xe2, 0x62,
	0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x93, 0xd8, 0xc0,
	0x8a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x0d, 0xbc, 0x5a, 0x3e, 0x00, 0x00, 0x00,
}
