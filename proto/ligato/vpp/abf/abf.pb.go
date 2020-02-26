// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/abf/abf.proto

package vpp_abf

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

// ABF defines ACL based forwarding.
type ABF struct {
	Index                uint32                   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	AclName              string                   `protobuf:"bytes,2,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	AttachedInterfaces   []*ABF_AttachedInterface `protobuf:"bytes,3,rep,name=attached_interfaces,json=attachedInterfaces,proto3" json:"attached_interfaces,omitempty"`
	ForwardingPaths      []*ABF_ForwardingPath    `protobuf:"bytes,4,rep,name=forwarding_paths,json=forwardingPaths,proto3" json:"forwarding_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ABF) Reset()         { *m = ABF{} }
func (m *ABF) String() string { return proto.CompactTextString(m) }
func (*ABF) ProtoMessage()    {}
func (*ABF) Descriptor() ([]byte, []int) {
	return fileDescriptor_1535b8ce2775e25e, []int{0}
}

func (m *ABF) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABF.Unmarshal(m, b)
}
func (m *ABF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABF.Marshal(b, m, deterministic)
}
func (m *ABF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABF.Merge(m, src)
}
func (m *ABF) XXX_Size() int {
	return xxx_messageInfo_ABF.Size(m)
}
func (m *ABF) XXX_DiscardUnknown() {
	xxx_messageInfo_ABF.DiscardUnknown(m)
}

var xxx_messageInfo_ABF proto.InternalMessageInfo

func (m *ABF) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ABF) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *ABF) GetAttachedInterfaces() []*ABF_AttachedInterface {
	if m != nil {
		return m.AttachedInterfaces
	}
	return nil
}

func (m *ABF) GetForwardingPaths() []*ABF_ForwardingPath {
	if m != nil {
		return m.ForwardingPaths
	}
	return nil
}

// List of interfaces attached to the ABF
type ABF_AttachedInterface struct {
	InputInterface       string   `protobuf:"bytes,1,opt,name=input_interface,json=inputInterface,proto3" json:"input_interface,omitempty"`
	Priority             uint32   `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	IsIpv6               bool     `protobuf:"varint,3,opt,name=is_ipv6,json=isIpv6,proto3" json:"is_ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ABF_AttachedInterface) Reset()         { *m = ABF_AttachedInterface{} }
func (m *ABF_AttachedInterface) String() string { return proto.CompactTextString(m) }
func (*ABF_AttachedInterface) ProtoMessage()    {}
func (*ABF_AttachedInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_1535b8ce2775e25e, []int{0, 0}
}

func (m *ABF_AttachedInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABF_AttachedInterface.Unmarshal(m, b)
}
func (m *ABF_AttachedInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABF_AttachedInterface.Marshal(b, m, deterministic)
}
func (m *ABF_AttachedInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABF_AttachedInterface.Merge(m, src)
}
func (m *ABF_AttachedInterface) XXX_Size() int {
	return xxx_messageInfo_ABF_AttachedInterface.Size(m)
}
func (m *ABF_AttachedInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_ABF_AttachedInterface.DiscardUnknown(m)
}

var xxx_messageInfo_ABF_AttachedInterface proto.InternalMessageInfo

func (m *ABF_AttachedInterface) GetInputInterface() string {
	if m != nil {
		return m.InputInterface
	}
	return ""
}

func (m *ABF_AttachedInterface) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *ABF_AttachedInterface) GetIsIpv6() bool {
	if m != nil {
		return m.IsIpv6
	}
	return false
}

// List of forwarding paths added to the ABF policy (via)
type ABF_ForwardingPath struct {
	NextHopIp            string   `protobuf:"bytes,1,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Weight               uint32   `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference           uint32   `protobuf:"varint,4,opt,name=preference,proto3" json:"preference,omitempty"`
	Dvr                  bool     `protobuf:"varint,5,opt,name=dvr,proto3" json:"dvr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ABF_ForwardingPath) Reset()         { *m = ABF_ForwardingPath{} }
func (m *ABF_ForwardingPath) String() string { return proto.CompactTextString(m) }
func (*ABF_ForwardingPath) ProtoMessage()    {}
func (*ABF_ForwardingPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_1535b8ce2775e25e, []int{0, 1}
}

func (m *ABF_ForwardingPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABF_ForwardingPath.Unmarshal(m, b)
}
func (m *ABF_ForwardingPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABF_ForwardingPath.Marshal(b, m, deterministic)
}
func (m *ABF_ForwardingPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABF_ForwardingPath.Merge(m, src)
}
func (m *ABF_ForwardingPath) XXX_Size() int {
	return xxx_messageInfo_ABF_ForwardingPath.Size(m)
}
func (m *ABF_ForwardingPath) XXX_DiscardUnknown() {
	xxx_messageInfo_ABF_ForwardingPath.DiscardUnknown(m)
}

var xxx_messageInfo_ABF_ForwardingPath proto.InternalMessageInfo

func (m *ABF_ForwardingPath) GetNextHopIp() string {
	if m != nil {
		return m.NextHopIp
	}
	return ""
}

func (m *ABF_ForwardingPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *ABF_ForwardingPath) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ABF_ForwardingPath) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *ABF_ForwardingPath) GetDvr() bool {
	if m != nil {
		return m.Dvr
	}
	return false
}

func init() {
	proto.RegisterType((*ABF)(nil), "ligato.vpp.abf.ABF")
	proto.RegisterType((*ABF_AttachedInterface)(nil), "ligato.vpp.abf.ABF.AttachedInterface")
	proto.RegisterType((*ABF_ForwardingPath)(nil), "ligato.vpp.abf.ABF.ForwardingPath")
}

func init() { proto.RegisterFile("ligato/vpp/abf/abf.proto", fileDescriptor_1535b8ce2775e25e) }

var fileDescriptor_1535b8ce2775e25e = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x71, 0x9b, 0x26, 0x53, 0x25, 0x2d, 0x0b, 0x82, 0x25, 0x87, 0xca, 0xaa, 0x54, 0xe1,
	0x0b, 0xb6, 0x44, 0xa5, 0x08, 0x89, 0x53, 0x72, 0x88, 0xc8, 0x01, 0x84, 0xf6, 0xc0, 0x81, 0x8b,
	0x35, 0xb6, 0xd7, 0xf6, 0x4a, 0xe9, 0xee, 0xb0, 0x5e, 0xdc, 0xf2, 0x39, 0xfc, 0x23, 0x1f, 0x80,
	0xb2, 0x4d, 0xdd, 0x18, 0x7a, 0x58, 0x69, 0xdf, 0x9b, 0xa7, 0x79, 0x6f, 0x46, 0x03, 0x7c, 0xab,
	0x6a, 0x74, 0x26, 0xed, 0x88, 0x52, 0xcc, 0xab, 0xdd, 0x4b, 0xc8, 0x1a, 0x67, 0xd8, 0xec, 0xbe,
	0x92, 0x74, 0x44, 0x09, 0xe6, 0xd5, 0xe5, 0x9f, 0x10, 0xc2, 0xe5, 0x6a, 0xcd, 0x5e, 0xc2, 0xb1,
	0xd2, 0xa5, 0xbc, 0xe3, 0x41, 0x14, 0xc4, 0x53, 0x71, 0x0f, 0xd8, 0x1b, 0x18, 0x63, 0xb1, 0xcd,
	0x34, 0xde, 0x48, 0xfe, 0x2c, 0x0a, 0xe2, 0x89, 0x38, 0xc1, 0x62, 0xfb, 0x05, 0x6f, 0x24, 0xfb,
	0x06, 0x2f, 0xd0, 0x39, 0x2c, 0x1a, 0x59, 0x66, 0x4a, 0x3b, 0x69, 0x2b, 0x2c, 0x64, 0xcb, 0xc3,
	0x28, 0x8c, 0x4f, 0xdf, 0x5f, 0x25, 0x43, 0x9b, 0x64, 0xb9, 0x5a, 0x27, 0xcb, 0xbd, 0x7c, 0xf3,
	0xa0, 0x16, 0x0c, 0xff, 0xa5, 0x5a, 0xf6, 0x19, 0xce, 0x2b, 0x63, 0x6f, 0xd1, 0x96, 0x4a, 0xd7,
	0x19, 0xa1, 0x6b, 0x5a, 0x7e, 0xe4, 0x9b, 0x5e, 0x3e, 0xd5, 0x74, 0xdd, 0x6b, 0xbf, 0xa2, 0x6b,
	0xc4, 0x59, 0x35, 0xc0, 0xed, 0xfc, 0x07, 0x3c, 0xff, 0xcf, 0x97, 0xbd, 0x85, 0x33, 0xa5, 0xe9,
	0xa7, 0x7b, 0x0c, 0xee, 0xc7, 0x9e, 0x88, 0x99, 0xa7, 0x1f, 0x85, 0x73, 0x18, 0x93, 0x55, 0xc6,
	0x2a, 0xf7, 0xcb, 0xcf, 0x3f, 0x15, 0x3d, 0x66, 0xaf, 0xe1, 0x44, 0xb5, 0x99, 0xa2, 0x6e, 0xc1,
	0xc3, 0x28, 0x88, 0xc7, 0x62, 0xa4, 0xda, 0x0d, 0x75, 0x8b, 0xf9, 0xef, 0x00, 0x66, 0xc3, 0x58,
	0xec, 0x02, 0x4e, 0xb5, 0xbc, 0x73, 0x59, 0x63, 0x28, 0x53, 0xb4, 0x37, 0x9b, 0xec, 0xa8, 0x4f,
	0x86, 0x36, 0xc4, 0xae, 0x60, 0xd6, 0x47, 0x39, 0xdc, 0xf6, 0xb4, 0x67, 0xfd, 0xce, 0x5f, 0xc1,
	0xe8, 0x56, 0xaa, 0xba, 0x71, 0xde, 0x71, 0x2a, 0xf6, 0x88, 0x5d, 0x00, 0x90, 0x95, 0x95, 0xb4,
	0x52, 0x17, 0x92, 0x1f, 0xf9, 0xda, 0x01, 0xc3, 0xce, 0x21, 0x2c, 0x3b, 0xcb, 0x8f, 0x7d, 0xcc,
	0xdd, 0x77, 0xf5, 0xe1, 0xfb, 0xa2, 0x36, 0x0f, 0xfb, 0x54, 0xfe, 0x50, 0xde, 0x61, 0x2d, 0xb5,
	0x4b, 0xbb, 0xeb, 0xd4, 0x5f, 0x4a, 0x3a, 0x3c, 0xa1, 0x8f, 0x1d, 0x51, 0x86, 0x79, 0x95, 0x8f,
	0x7c, 0xf5, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xfc, 0xc7, 0xd3, 0x63, 0x02, 0x00,
	0x00,
}
