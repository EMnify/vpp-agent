// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: punt.proto

package punt

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolBased_Protocol int32

const (
	ProtocolBased_TCP ProtocolBased_Protocol = 0
	ProtocolBased_UDP ProtocolBased_Protocol = 1
)

var ProtocolBased_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var ProtocolBased_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x ProtocolBased_Protocol) String() string {
	return proto.EnumName(ProtocolBased_Protocol_name, int32(x))
}
func (ProtocolBased_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_punt_ef1ca0e2d0e28ea3, []int{1, 0}
}

// Proxy allows to listen on network socket or unix domain socket, and resend to another network/unix domain socket
type Proxy struct {
	// Types that are valid to be assigned to Rx:
	//	*Proxy_RxProtocol
	//	*Proxy_RxSocket
	Rx isProxy_Rx `protobuf_oneof:"rx"`
	// Types that are valid to be assigned to Tx:
	//	*Proxy_TxProtocol
	//	*Proxy_TxSocket
	Tx                   isProxy_Tx `protobuf_oneof:"tx"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_ef1ca0e2d0e28ea3, []int{0}
}
func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (dst *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(dst, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

type isProxy_Rx interface {
	isProxy_Rx()
}
type isProxy_Tx interface {
	isProxy_Tx()
}

type Proxy_RxProtocol struct {
	RxProtocol *ProtocolBased `protobuf:"bytes,1,opt,name=rx_protocol,json=rxProtocol,oneof"`
}
type Proxy_RxSocket struct {
	RxSocket *SocketBased `protobuf:"bytes,2,opt,name=rx_socket,json=rxSocket,oneof"`
}
type Proxy_TxProtocol struct {
	TxProtocol *ProtocolBased `protobuf:"bytes,3,opt,name=tx_protocol,json=txProtocol,oneof"`
}
type Proxy_TxSocket struct {
	TxSocket *SocketBased `protobuf:"bytes,4,opt,name=tx_socket,json=txSocket,oneof"`
}

func (*Proxy_RxProtocol) isProxy_Rx() {}
func (*Proxy_RxSocket) isProxy_Rx()   {}
func (*Proxy_TxProtocol) isProxy_Tx() {}
func (*Proxy_TxSocket) isProxy_Tx()   {}

func (m *Proxy) GetRx() isProxy_Rx {
	if m != nil {
		return m.Rx
	}
	return nil
}
func (m *Proxy) GetTx() isProxy_Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *Proxy) GetRxProtocol() *ProtocolBased {
	if x, ok := m.GetRx().(*Proxy_RxProtocol); ok {
		return x.RxProtocol
	}
	return nil
}

func (m *Proxy) GetRxSocket() *SocketBased {
	if x, ok := m.GetRx().(*Proxy_RxSocket); ok {
		return x.RxSocket
	}
	return nil
}

func (m *Proxy) GetTxProtocol() *ProtocolBased {
	if x, ok := m.GetTx().(*Proxy_TxProtocol); ok {
		return x.TxProtocol
	}
	return nil
}

func (m *Proxy) GetTxSocket() *SocketBased {
	if x, ok := m.GetTx().(*Proxy_TxSocket); ok {
		return x.TxSocket
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Proxy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Proxy_OneofMarshaler, _Proxy_OneofUnmarshaler, _Proxy_OneofSizer, []interface{}{
		(*Proxy_RxProtocol)(nil),
		(*Proxy_RxSocket)(nil),
		(*Proxy_TxProtocol)(nil),
		(*Proxy_TxSocket)(nil),
	}
}

func _Proxy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Proxy)
	// rx
	switch x := m.Rx.(type) {
	case *Proxy_RxProtocol:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RxProtocol); err != nil {
			return err
		}
	case *Proxy_RxSocket:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RxSocket); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Proxy.Rx has unexpected type %T", x)
	}
	// tx
	switch x := m.Tx.(type) {
	case *Proxy_TxProtocol:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TxProtocol); err != nil {
			return err
		}
	case *Proxy_TxSocket:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TxSocket); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Proxy.Tx has unexpected type %T", x)
	}
	return nil
}

func _Proxy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Proxy)
	switch tag {
	case 1: // rx.rx_protocol
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProtocolBased)
		err := b.DecodeMessage(msg)
		m.Rx = &Proxy_RxProtocol{msg}
		return true, err
	case 2: // rx.rx_socket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketBased)
		err := b.DecodeMessage(msg)
		m.Rx = &Proxy_RxSocket{msg}
		return true, err
	case 3: // tx.tx_protocol
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProtocolBased)
		err := b.DecodeMessage(msg)
		m.Tx = &Proxy_TxProtocol{msg}
		return true, err
	case 4: // tx.tx_socket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketBased)
		err := b.DecodeMessage(msg)
		m.Tx = &Proxy_TxSocket{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Proxy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Proxy)
	// rx
	switch x := m.Rx.(type) {
	case *Proxy_RxProtocol:
		s := proto.Size(x.RxProtocol)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Proxy_RxSocket:
		s := proto.Size(x.RxSocket)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// tx
	switch x := m.Tx.(type) {
	case *Proxy_TxProtocol:
		s := proto.Size(x.TxProtocol)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Proxy_TxSocket:
		s := proto.Size(x.TxSocket)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Define network socket type
type ProtocolBased struct {
	Protocol             ProtocolBased_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=punt.ProtocolBased_Protocol" json:"protocol,omitempty"`
	IsIpv6               bool                   `protobuf:"varint,2,opt,name=is_ipv6,json=isIpv6,proto3" json:"is_ipv6,omitempty"`
	Port                 uint32                 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ProtocolBased) Reset()         { *m = ProtocolBased{} }
func (m *ProtocolBased) String() string { return proto.CompactTextString(m) }
func (*ProtocolBased) ProtoMessage()    {}
func (*ProtocolBased) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_ef1ca0e2d0e28ea3, []int{1}
}
func (m *ProtocolBased) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolBased.Unmarshal(m, b)
}
func (m *ProtocolBased) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolBased.Marshal(b, m, deterministic)
}
func (dst *ProtocolBased) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolBased.Merge(dst, src)
}
func (m *ProtocolBased) XXX_Size() int {
	return xxx_messageInfo_ProtocolBased.Size(m)
}
func (m *ProtocolBased) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolBased.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolBased proto.InternalMessageInfo

func (m *ProtocolBased) GetProtocol() ProtocolBased_Protocol {
	if m != nil {
		return m.Protocol
	}
	return ProtocolBased_TCP
}

func (m *ProtocolBased) GetIsIpv6() bool {
	if m != nil {
		return m.IsIpv6
	}
	return false
}

func (m *ProtocolBased) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Define unix domain socket type for IPC
type SocketBased struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocketBased) Reset()         { *m = SocketBased{} }
func (m *SocketBased) String() string { return proto.CompactTextString(m) }
func (*SocketBased) ProtoMessage()    {}
func (*SocketBased) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_ef1ca0e2d0e28ea3, []int{2}
}
func (m *SocketBased) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketBased.Unmarshal(m, b)
}
func (m *SocketBased) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketBased.Marshal(b, m, deterministic)
}
func (dst *SocketBased) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketBased.Merge(dst, src)
}
func (m *SocketBased) XXX_Size() int {
	return xxx_messageInfo_SocketBased.Size(m)
}
func (m *SocketBased) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketBased.DiscardUnknown(m)
}

var xxx_messageInfo_SocketBased proto.InternalMessageInfo

func (m *SocketBased) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Proxy)(nil), "punt.Proxy")
	proto.RegisterType((*ProtocolBased)(nil), "punt.ProtocolBased")
	proto.RegisterType((*SocketBased)(nil), "punt.SocketBased")
	proto.RegisterEnum("punt.ProtocolBased_Protocol", ProtocolBased_Protocol_name, ProtocolBased_Protocol_value)
}

func init() { proto.RegisterFile("punt.proto", fileDescriptor_punt_ef1ca0e2d0e28ea3) }

var fileDescriptor_punt_ef1ca0e2d0e28ea3 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x29, 0xac, 0x50, 0x66, 0x83, 0xc1, 0x7a, 0x90, 0x03, 0x07, 0xdd, 0x93, 0x27, 0x62,
	0x30, 0xd9, 0x78, 0x46, 0x0f, 0x7a, 0x6b, 0xaa, 0x9e, 0x37, 0x88, 0x24, 0x6e, 0x34, 0xb6, 0x99,
	0x8e, 0xa4, 0xfe, 0x0f, 0xff, 0xa3, 0x7f, 0xc3, 0x74, 0x80, 0x86, 0x4d, 0xdc, 0x4b, 0xf3, 0xfa,
	0xda, 0x37, 0xdf, 0x3c, 0x00, 0xf7, 0xf5, 0x49, 0x33, 0x87, 0x96, 0xac, 0xca, 0xa2, 0x2e, 0x7e,
	0x05, 0x1c, 0x69, 0xb4, 0xe1, 0x5b, 0x95, 0x90, 0x63, 0xa8, 0xf8, 0x6d, 0x65, 0x3f, 0x26, 0xe2,
	0x5c, 0x5c, 0xe6, 0xf3, 0xd3, 0x19, 0x27, 0xf4, 0xce, 0x5d, 0x2c, 0xfd, 0xfa, 0xf5, 0xbe, 0x63,
	0x00, 0xc3, 0xde, 0x52, 0x57, 0x30, 0xc4, 0x50, 0x79, 0xbb, 0x7a, 0x5f, 0xd3, 0xa4, 0xcb, 0xa9,
	0x93, 0x6d, 0xea, 0x91, 0xbd, 0x7d, 0x46, 0x62, 0xd8, 0x1a, 0x91, 0x44, 0x07, 0xa4, 0x5e, 0x3b,
	0x49, 0x18, 0xa0, 0x06, 0x89, 0x12, 0x29, 0x6b, 0x23, 0x09, 0x23, 0x69, 0x47, 0x5a, 0x64, 0xd0,
	0xc5, 0x10, 0x4f, 0x0a, 0xc5, 0x8f, 0x80, 0x51, 0x63, 0xba, 0xba, 0x01, 0xd9, 0xa8, 0x7b, 0x3c,
	0x9f, 0xfe, 0xb3, 0x44, 0xba, 0x99, 0xf4, 0x5b, 0x9d, 0xc1, 0xa0, 0xf6, 0x55, 0xed, 0x36, 0x25,
	0x37, 0x96, 0xa6, 0x5f, 0xfb, 0x07, 0xb7, 0x29, 0x95, 0x82, 0xcc, 0x59, 0x24, 0xee, 0x34, 0x32,
	0xac, 0x8b, 0x29, 0xc8, 0x54, 0x61, 0x00, 0xbd, 0xa7, 0x5b, 0x3d, 0xee, 0x44, 0xf1, 0x7c, 0xa7,
	0xc7, 0xa2, 0xb8, 0x80, 0xfc, 0x60, 0x7b, 0x1e, 0xb0, 0xa4, 0x37, 0xde, 0x67, 0x68, 0x58, 0xbf,
	0xf4, 0x99, 0x7b, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x99, 0x6e, 0x1d, 0x7a, 0xbe, 0x01, 0x00,
	0x00,
}
