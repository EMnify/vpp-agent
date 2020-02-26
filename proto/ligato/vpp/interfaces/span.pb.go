// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/interfaces/span.proto

package vpp_interfaces

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

type Span_Direction int32

const (
	Span_UNKNOWN Span_Direction = 0
	Span_RX      Span_Direction = 1
	Span_TX      Span_Direction = 2
	Span_BOTH    Span_Direction = 3
)

var Span_Direction_name = map[int32]string{
	0: "UNKNOWN",
	1: "RX",
	2: "TX",
	3: "BOTH",
}

var Span_Direction_value = map[string]int32{
	"UNKNOWN": 0,
	"RX":      1,
	"TX":      2,
	"BOTH":    3,
}

func (x Span_Direction) String() string {
	return proto.EnumName(Span_Direction_name, int32(x))
}

func (Span_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3db2120d57515158, []int{0, 0}
}

type Span struct {
	InterfaceFrom        string         `protobuf:"bytes,1,opt,name=interface_from,json=interfaceFrom,proto3" json:"interface_from,omitempty"`
	InterfaceTo          string         `protobuf:"bytes,2,opt,name=interface_to,json=interfaceTo,proto3" json:"interface_to,omitempty"`
	Direction            Span_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=ligato.vpp.interfaces.Span_Direction" json:"direction,omitempty"`
	IsL2                 bool           `protobuf:"varint,4,opt,name=is_l2,json=isL2,proto3" json:"is_l2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_3db2120d57515158, []int{0}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetInterfaceFrom() string {
	if m != nil {
		return m.InterfaceFrom
	}
	return ""
}

func (m *Span) GetInterfaceTo() string {
	if m != nil {
		return m.InterfaceTo
	}
	return ""
}

func (m *Span) GetDirection() Span_Direction {
	if m != nil {
		return m.Direction
	}
	return Span_UNKNOWN
}

func (m *Span) GetIsL2() bool {
	if m != nil {
		return m.IsL2
	}
	return false
}

func init() {
	proto.RegisterEnum("ligato.vpp.interfaces.Span_Direction", Span_Direction_name, Span_Direction_value)
	proto.RegisterType((*Span)(nil), "ligato.vpp.interfaces.Span")
}

func init() { proto.RegisterFile("ligato/vpp/interfaces/span.proto", fileDescriptor_3db2120d57515158) }

var fileDescriptor_3db2120d57515158 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x06, 0x70, 0xd3, 0xd5, 0xb9, 0xbe, 0xe9, 0x28, 0x11, 0xa1, 0xc7, 0x3a, 0x18, 0xf4, 0x62,
	0x02, 0xdd, 0xd1, 0xdb, 0x1c, 0x43, 0x50, 0x3a, 0xa8, 0x15, 0x87, 0x97, 0x52, 0x67, 0x57, 0x02,
	0x5b, 0xde, 0x23, 0x0d, 0xfd, 0x73, 0xfd, 0x5b, 0x64, 0x1d, 0x6b, 0x2f, 0x3b, 0x85, 0x7c, 0xfc,
	0x1e, 0x1f, 0x7c, 0x10, 0xee, 0x55, 0x55, 0x58, 0x94, 0x0d, 0x91, 0x54, 0xda, 0x96, 0x66, 0x57,
	0x6c, 0xcb, 0x5a, 0xd6, 0x54, 0x68, 0x41, 0x06, 0x2d, 0xf2, 0x87, 0x93, 0x10, 0x0d, 0x91, 0xe8,
	0xc5, 0xf4, 0x8f, 0x81, 0xfb, 0x41, 0x85, 0xe6, 0x33, 0x98, 0x74, 0x71, 0xbe, 0x33, 0x78, 0x08,
	0x58, 0xc8, 0x22, 0x2f, 0xbd, 0xeb, 0xd2, 0x95, 0xc1, 0x03, 0x7f, 0x84, 0xdb, 0x9e, 0x59, 0x0c,
	0x9c, 0x16, 0x8d, 0xbb, 0x2c, 0x43, 0xfe, 0x02, 0xde, 0xaf, 0x32, 0xe5, 0xd6, 0x2a, 0xd4, 0xc1,
	0x20, 0x64, 0xd1, 0x24, 0x9e, 0x89, 0x8b, 0xed, 0xe2, 0xd8, 0x2c, 0x96, 0x67, 0x9c, 0xf6, 0x77,
	0xfc, 0x1e, 0xae, 0x55, 0x9d, 0xef, 0xe3, 0xc0, 0x0d, 0x59, 0x34, 0x4a, 0x5d, 0x55, 0xbf, 0xc7,
	0xd3, 0x18, 0xbc, 0x0e, 0xf3, 0x31, 0xdc, 0x7c, 0x26, 0x6f, 0xc9, 0xfa, 0x2b, 0xf1, 0xaf, 0xf8,
	0x10, 0x9c, 0x74, 0xe3, 0xb3, 0xe3, 0x9b, 0x6d, 0x7c, 0x87, 0x8f, 0xc0, 0x5d, 0xac, 0xb3, 0x57,
	0x7f, 0xb0, 0x58, 0x7d, 0x2f, 0x2b, 0x3c, 0xd7, 0xab, 0x76, 0xa1, 0xa7, 0xa2, 0x2a, 0xb5, 0x95,
	0xcd, 0x5c, 0xb6, 0xd3, 0xc8, 0x8b, 0xdb, 0x3d, 0x37, 0x44, 0x79, 0xff, 0xfd, 0x19, 0xb6, 0x76,
	0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xed, 0x56, 0xa5, 0x97, 0x6a, 0x01, 0x00, 0x00,
}
