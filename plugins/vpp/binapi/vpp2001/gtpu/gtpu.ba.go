// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/gtpu.api.json

/*
Package gtpu is a generated VPP binary API for 'gtpu' module.

It consists of:
	 10 enums
	  6 aliases
	  6 types
	  1 union
	  8 messages
	  4 services
*/
package gtpu

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"

	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "gtpu"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x7046552
)

type AddressFamily = ip_types.AddressFamily

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

type AddressWithPrefix = ip_types.AddressWithPrefix

type InterfaceIndex = interface_types.InterfaceIndex

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

type IP4Prefix = ip_types.IP4Prefix

type IP6Prefix = ip_types.IP6Prefix

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

type AddressUnion = ip_types.AddressUnion

// GtpuAddDelTunnel represents VPP binary API message 'gtpu_add_del_tunnel'.
type GtpuAddDelTunnel struct {
	IsAdd          bool
	SrcAddress     Address
	DstAddress     Address
	McastSwIfIndex InterfaceIndex
	EncapVrfID     uint32
	DecapNextIndex uint32
	SrcTeid        uint32
	DstTeid        uint32
}

func (m *GtpuAddDelTunnel) Reset()                        { *m = GtpuAddDelTunnel{} }
func (*GtpuAddDelTunnel) GetMessageName() string          { return "gtpu_add_del_tunnel" }
func (*GtpuAddDelTunnel) GetCrcString() string            { return "2c4409a6" }
func (*GtpuAddDelTunnel) GetMessageType() api.MessageType { return api.RequestMessage }

// GtpuAddDelTunnelReply represents VPP binary API message 'gtpu_add_del_tunnel_reply'.
type GtpuAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *GtpuAddDelTunnelReply) Reset()                        { *m = GtpuAddDelTunnelReply{} }
func (*GtpuAddDelTunnelReply) GetMessageName() string          { return "gtpu_add_del_tunnel_reply" }
func (*GtpuAddDelTunnelReply) GetCrcString() string            { return "5383d31f" }
func (*GtpuAddDelTunnelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GtpuSetTunnelDst represents VPP binary API message 'gtpu_set_tunnel_dst'.
type GtpuSetTunnelDst struct {
	SwIfIndex      InterfaceIndex
	DstAddress     Address
	McastSwIfIndex InterfaceIndex
	DstTeid        uint32
}

func (m *GtpuSetTunnelDst) Reset()                        { *m = GtpuSetTunnelDst{} }
func (*GtpuSetTunnelDst) GetMessageName() string          { return "gtpu_set_tunnel_dst" }
func (*GtpuSetTunnelDst) GetCrcString() string            { return "26cc5d2d" }
func (*GtpuSetTunnelDst) GetMessageType() api.MessageType { return api.RequestMessage }

// GtpuSetTunnelDstReply represents VPP binary API message 'gtpu_set_tunnel_dst_reply'.
type GtpuSetTunnelDstReply struct {
	Retval int32
}

func (m *GtpuSetTunnelDstReply) Reset()                        { *m = GtpuSetTunnelDstReply{} }
func (*GtpuSetTunnelDstReply) GetMessageName() string          { return "gtpu_set_tunnel_dst_reply" }
func (*GtpuSetTunnelDstReply) GetCrcString() string            { return "e8d4e804" }
func (*GtpuSetTunnelDstReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GtpuTunnelDetails represents VPP binary API message 'gtpu_tunnel_details'.
type GtpuTunnelDetails struct {
	SwIfIndex      InterfaceIndex
	SrcAddress     Address
	DstAddress     Address
	McastSwIfIndex InterfaceIndex
	EncapVrfID     uint32
	DecapNextIndex uint32
	SrcTeid        uint32
	DstTeid        uint32
}

func (m *GtpuTunnelDetails) Reset()                        { *m = GtpuTunnelDetails{} }
func (*GtpuTunnelDetails) GetMessageName() string          { return "gtpu_tunnel_details" }
func (*GtpuTunnelDetails) GetCrcString() string            { return "ca7b71ac" }
func (*GtpuTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// GtpuTunnelDump represents VPP binary API message 'gtpu_tunnel_dump'.
type GtpuTunnelDump struct {
	SwIfIndex InterfaceIndex
}

func (m *GtpuTunnelDump) Reset()                        { *m = GtpuTunnelDump{} }
func (*GtpuTunnelDump) GetMessageName() string          { return "gtpu_tunnel_dump" }
func (*GtpuTunnelDump) GetCrcString() string            { return "f9e6675e" }
func (*GtpuTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetGtpuBypass represents VPP binary API message 'sw_interface_set_gtpu_bypass'.
type SwInterfaceSetGtpuBypass struct {
	SwIfIndex InterfaceIndex
	IsIPv6    bool
	Enable    bool
}

func (m *SwInterfaceSetGtpuBypass) Reset()                        { *m = SwInterfaceSetGtpuBypass{} }
func (*SwInterfaceSetGtpuBypass) GetMessageName() string          { return "sw_interface_set_gtpu_bypass" }
func (*SwInterfaceSetGtpuBypass) GetCrcString() string            { return "65247409" }
func (*SwInterfaceSetGtpuBypass) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetGtpuBypassReply represents VPP binary API message 'sw_interface_set_gtpu_bypass_reply'.
type SwInterfaceSetGtpuBypassReply struct {
	Retval int32
}

func (m *SwInterfaceSetGtpuBypassReply) Reset() { *m = SwInterfaceSetGtpuBypassReply{} }
func (*SwInterfaceSetGtpuBypassReply) GetMessageName() string {
	return "sw_interface_set_gtpu_bypass_reply"
}
func (*SwInterfaceSetGtpuBypassReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetGtpuBypassReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func init() {
	api.RegisterMessage((*GtpuAddDelTunnel)(nil), "gtpu.GtpuAddDelTunnel")
	api.RegisterMessage((*GtpuAddDelTunnelReply)(nil), "gtpu.GtpuAddDelTunnelReply")
	api.RegisterMessage((*GtpuSetTunnelDst)(nil), "gtpu.GtpuSetTunnelDst")
	api.RegisterMessage((*GtpuSetTunnelDstReply)(nil), "gtpu.GtpuSetTunnelDstReply")
	api.RegisterMessage((*GtpuTunnelDetails)(nil), "gtpu.GtpuTunnelDetails")
	api.RegisterMessage((*GtpuTunnelDump)(nil), "gtpu.GtpuTunnelDump")
	api.RegisterMessage((*SwInterfaceSetGtpuBypass)(nil), "gtpu.SwInterfaceSetGtpuBypass")
	api.RegisterMessage((*SwInterfaceSetGtpuBypassReply)(nil), "gtpu.SwInterfaceSetGtpuBypassReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GtpuAddDelTunnel)(nil),
		(*GtpuAddDelTunnelReply)(nil),
		(*GtpuSetTunnelDst)(nil),
		(*GtpuSetTunnelDstReply)(nil),
		(*GtpuTunnelDetails)(nil),
		(*GtpuTunnelDump)(nil),
		(*SwInterfaceSetGtpuBypass)(nil),
		(*SwInterfaceSetGtpuBypassReply)(nil),
	}
}

// RPCService represents RPC service API for gtpu module.
type RPCService interface {
	DumpGtpuTunnel(ctx context.Context, in *GtpuTunnelDump) (RPCService_DumpGtpuTunnelClient, error)
	GtpuAddDelTunnel(ctx context.Context, in *GtpuAddDelTunnel) (*GtpuAddDelTunnelReply, error)
	GtpuSetTunnelDst(ctx context.Context, in *GtpuSetTunnelDst) (*GtpuSetTunnelDstReply, error)
	SwInterfaceSetGtpuBypass(ctx context.Context, in *SwInterfaceSetGtpuBypass) (*SwInterfaceSetGtpuBypassReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpGtpuTunnel(ctx context.Context, in *GtpuTunnelDump) (RPCService_DumpGtpuTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpGtpuTunnelClient{stream}
	return x, nil
}

type RPCService_DumpGtpuTunnelClient interface {
	Recv() (*GtpuTunnelDetails, error)
}

type serviceClient_DumpGtpuTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpGtpuTunnelClient) Recv() (*GtpuTunnelDetails, error) {
	m := new(GtpuTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) GtpuAddDelTunnel(ctx context.Context, in *GtpuAddDelTunnel) (*GtpuAddDelTunnelReply, error) {
	out := new(GtpuAddDelTunnelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GtpuSetTunnelDst(ctx context.Context, in *GtpuSetTunnelDst) (*GtpuSetTunnelDstReply, error) {
	out := new(GtpuSetTunnelDstReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceSetGtpuBypass(ctx context.Context, in *SwInterfaceSetGtpuBypass) (*SwInterfaceSetGtpuBypassReply, error) {
	out := new(SwInterfaceSetGtpuBypassReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
