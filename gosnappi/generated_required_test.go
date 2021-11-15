package gosnappi_test

import (
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
)

func TestPortRequired(t *testing.T) {
	object := gosnappi.NewPort()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestLagRequired(t *testing.T) {
	object := gosnappi.NewLag()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestLayer1Required(t *testing.T) {
	object := gosnappi.NewLayer1()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestCaptureRequired(t *testing.T) {
	object := gosnappi.NewCapture()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestDeviceRequired(t *testing.T) {
	object := gosnappi.NewDevice()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestFlowRequired(t *testing.T) {
	object := gosnappi.NewFlow()
	err := object.Validate()
	assert.Contains(t, err.Error(), "TxRx", "Name")
}
func TestTransmitStateRequired(t *testing.T) {
	object := gosnappi.NewTransmitState()
	err := object.Validate()
	assert.Contains(t, err.Error(), "State")
}
func TestLinkStateRequired(t *testing.T) {
	object := gosnappi.NewLinkState()
	err := object.Validate()
	assert.Contains(t, err.Error(), "State")
}
func TestCaptureStateRequired(t *testing.T) {
	object := gosnappi.NewCaptureState()
	err := object.Validate()
	assert.Contains(t, err.Error(), "State")
}
func TestRouteStateRequired(t *testing.T) {
	object := gosnappi.NewRouteState()
	err := object.Validate()
	assert.Contains(t, err.Error(), "State")
}
func TestProtocolStateRequired(t *testing.T) {
	object := gosnappi.NewProtocolState()
	err := object.Validate()
	assert.Contains(t, err.Error(), "State")
}
func TestCaptureRequestRequired(t *testing.T) {
	object := gosnappi.NewCaptureRequest()
	err := object.Validate()
	assert.Contains(t, err.Error(), "PortName")
}
func TestLagPortRequired(t *testing.T) {
	object := gosnappi.NewLagPort()
	err := object.Validate()
	assert.Contains(t, err.Error(), "PortName", "Protocol", "Ethernet")
}
func TestDeviceEthernetRequired(t *testing.T) {
	object := gosnappi.NewDeviceEthernet()
	err := object.Validate()
	assert.Contains(t, err.Error(), "PortName", "Mac", "Name")
}
func TestDeviceIpv4LoopbackRequired(t *testing.T) {
	object := gosnappi.NewDeviceIpv4Loopback()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EthName", "Name")
}
func TestDeviceIpv6LoopbackRequired(t *testing.T) {
	object := gosnappi.NewDeviceIpv6Loopback()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EthName", "Name")
}
func TestDeviceIsisRouterRequired(t *testing.T) {
	object := gosnappi.NewDeviceIsisRouter()
	err := object.Validate()
	assert.Contains(t, err.Error(), "SystemId", "Name")
}
func TestDeviceBgpRouterRequired(t *testing.T) {
	object := gosnappi.NewDeviceBgpRouter()
	err := object.Validate()
	assert.Contains(t, err.Error(), "RouterId")
}
func TestNeighborsv4StateRequired(t *testing.T) {
	object := gosnappi.NewNeighborsv4State()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EthernetName", "Ipv4Address")
}
func TestNeighborsv6StateRequired(t *testing.T) {
	object := gosnappi.NewNeighborsv6State()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EthernetName", "Ipv6Address")
}
func TestDeviceEthernetBaseRequired(t *testing.T) {
	object := gosnappi.NewDeviceEthernetBase()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Mac", "Name")
}
func TestDeviceIpv4Required(t *testing.T) {
	object := gosnappi.NewDeviceIpv4()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Gateway", "Address", "Name")
}
func TestDeviceIpv6Required(t *testing.T) {
	object := gosnappi.NewDeviceIpv6()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Gateway", "Address", "Name")
}
func TestDeviceVlanRequired(t *testing.T) {
	object := gosnappi.NewDeviceVlan()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestIsisInterfaceRequired(t *testing.T) {
	object := gosnappi.NewIsisInterface()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EthName", "Name")
}
func TestIsisV4RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewIsisV4RouteRange()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestIsisV6RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewIsisV6RouteRange()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpV4InterfaceRequired(t *testing.T) {
	object := gosnappi.NewBgpV4Interface()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv4Name")
}
func TestBgpV6InterfaceRequired(t *testing.T) {
	object := gosnappi.NewBgpV6Interface()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv6Name")
}
func TestFlowPortRequired(t *testing.T) {
	object := gosnappi.NewFlowPort()
	err := object.Validate()
	assert.Contains(t, err.Error(), "TxName")
}
func TestFlowCustomRequired(t *testing.T) {
	object := gosnappi.NewFlowCustom()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Bytes")
}
func TestIsisInterfaceAuthenticationRequired(t *testing.T) {
	object := gosnappi.NewIsisInterfaceAuthentication()
	err := object.Validate()
	assert.Contains(t, err.Error(), "AuthType")
}
func TestIsisAuthenticationBaseRequired(t *testing.T) {
	object := gosnappi.NewIsisAuthenticationBase()
	err := object.Validate()
	assert.Contains(t, err.Error(), "AuthType")
}
func TestV4RouteAddressRequired(t *testing.T) {
	object := gosnappi.NewV4RouteAddress()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Address")
}
func TestV6RouteAddressRequired(t *testing.T) {
	object := gosnappi.NewV6RouteAddress()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Address")
}
func TestBgpV4PeerRequired(t *testing.T) {
	object := gosnappi.NewBgpV4Peer()
	err := object.Validate()
	assert.Contains(t, err.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
}
func TestBgpV6PeerRequired(t *testing.T) {
	object := gosnappi.NewBgpV6Peer()
	err := object.Validate()
	assert.Contains(t, err.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
}
func TestBgpV4RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewBgpV4RouteRange()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpV6RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewBgpV6RouteRange()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpSrteV4PolicyRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV4Policy()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv4Endpoint", "Name")
}
func TestBgpSrteV6PolicyRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV6Policy()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv6Endpoint", "Name")
}
func TestBgpSrteV4TunnelTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV4TunnelTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpSrteV6TunnelTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV6TunnelTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpSrteSegmentListRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentList()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name")
}
func TestBgpSrteSegmentRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegment()
	err := object.Validate()
	assert.Contains(t, err.Error(), "SegmentType", "Name")
}
func TestBgpSrteSegmentBTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentBTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Srv6Sid")
}
func TestBgpSrteSegmentCTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentCTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv4NodeAddress")
}
func TestBgpSrteSegmentDTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentDTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv6NodeAddress")
}
func TestBgpSrteSegmentETypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentETypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv4NodeAddress")
}
func TestBgpSrteSegmentFTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentFTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "LocalIpv4Address", "RemoteIpv4Address")
}
func TestBgpSrteSegmentGTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentGTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
}
func TestBgpSrteSegmentHTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentHTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "LocalIpv6Address", "RemoteIpv6Address")
}
func TestBgpSrteSegmentITypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentITypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Ipv6NodeAddress")
}
func TestBgpSrteSegmentJTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentJTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
}
func TestBgpSrteSegmentKTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentKTypeSubTlv()
	err := object.Validate()
	assert.Contains(t, err.Error(), "LocalIpv6Address", "RemoteIpv6Address")
}
