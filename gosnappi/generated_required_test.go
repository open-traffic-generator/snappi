package gosnappi_test

import (
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestPortRequired(t *testing.T) {
	object := gosnappi.NewPort()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestLagRequired(t *testing.T) {
	object := gosnappi.NewLag()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestLayer1Required(t *testing.T) {
	object := gosnappi.NewLayer1()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestCaptureRequired(t *testing.T) {
	object := gosnappi.NewCapture()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestDeviceRequired(t *testing.T) {
	object := gosnappi.NewDevice()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestFlowRequired(t *testing.T) {
	object := gosnappi.NewFlow()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "TxRx", "Name")
	assert.Contains(t, err1.Error(), "TxRx", "Name")
	assert.Contains(t, err2.Error(), "TxRx", "Name")
}
func TestTransmitStateRequired(t *testing.T) {
	object := gosnappi.NewTransmitState()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "State")
	assert.Contains(t, err1.Error(), "State")
	assert.Contains(t, err2.Error(), "State")
}
func TestLinkStateRequired(t *testing.T) {
	object := gosnappi.NewLinkState()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "State")
	assert.Contains(t, err1.Error(), "State")
	assert.Contains(t, err2.Error(), "State")
}
func TestCaptureStateRequired(t *testing.T) {
	object := gosnappi.NewCaptureState()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "State")
	assert.Contains(t, err1.Error(), "State")
	assert.Contains(t, err2.Error(), "State")
}
func TestRouteStateRequired(t *testing.T) {
	object := gosnappi.NewRouteState()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "State")
	assert.Contains(t, err1.Error(), "State")
	assert.Contains(t, err2.Error(), "State")
}
func TestProtocolStateRequired(t *testing.T) {
	object := gosnappi.NewProtocolState()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "State")
	assert.Contains(t, err1.Error(), "State")
	assert.Contains(t, err2.Error(), "State")
}
func TestCaptureRequestRequired(t *testing.T) {
	object := gosnappi.NewCaptureRequest()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "PortName")
	assert.Contains(t, err1.Error(), "PortName")
	assert.Contains(t, err2.Error(), "PortName")
}
func TestLagPortRequired(t *testing.T) {
	object := gosnappi.NewLagPort()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "PortName", "Protocol", "Ethernet")
	assert.Contains(t, err1.Error(), "PortName", "Protocol", "Ethernet")
	assert.Contains(t, err2.Error(), "PortName", "Protocol", "Ethernet")
}
func TestDeviceEthernetRequired(t *testing.T) {
	object := gosnappi.NewDeviceEthernet()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Mac", "Name")
	assert.Contains(t, err1.Error(), "Mac", "Name")
	assert.Contains(t, err2.Error(), "Mac", "Name")
}
func TestDeviceIpv4LoopbackRequired(t *testing.T) {
	object := gosnappi.NewDeviceIpv4Loopback()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "EthName", "Name")
	assert.Contains(t, err1.Error(), "EthName", "Name")
	assert.Contains(t, err2.Error(), "EthName", "Name")
}
func TestDeviceIpv6LoopbackRequired(t *testing.T) {
	object := gosnappi.NewDeviceIpv6Loopback()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "EthName", "Name")
	assert.Contains(t, err1.Error(), "EthName", "Name")
	assert.Contains(t, err2.Error(), "EthName", "Name")
}
func TestDeviceIsisRouterRequired(t *testing.T) {
	object := gosnappi.NewDeviceIsisRouter()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "SystemId", "Name")
	assert.Contains(t, err1.Error(), "SystemId", "Name")
	assert.Contains(t, err2.Error(), "SystemId", "Name")
}
func TestDeviceBgpRouterRequired(t *testing.T) {
	object := gosnappi.NewDeviceBgpRouter()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "RouterId")
	assert.Contains(t, err1.Error(), "RouterId")
	assert.Contains(t, err2.Error(), "RouterId")
}
func TestNeighborsv4StateRequired(t *testing.T) {
	object := gosnappi.NewNeighborsv4State()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "EthernetName", "Ipv4Address")
	assert.Contains(t, err1.Error(), "EthernetName", "Ipv4Address")
	assert.Contains(t, err2.Error(), "EthernetName", "Ipv4Address")
}
func TestNeighborsv6StateRequired(t *testing.T) {
	object := gosnappi.NewNeighborsv6State()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "EthernetName", "Ipv6Address")
	assert.Contains(t, err1.Error(), "EthernetName", "Ipv6Address")
	assert.Contains(t, err2.Error(), "EthernetName", "Ipv6Address")
}
func TestDeviceEthernetBaseRequired(t *testing.T) {
	object := gosnappi.NewDeviceEthernetBase()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Mac", "Name")
	assert.Contains(t, err1.Error(), "Mac", "Name")
	assert.Contains(t, err2.Error(), "Mac", "Name")
}
func TestDeviceIpv4Required(t *testing.T) {
	object := gosnappi.NewDeviceIpv4()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Gateway", "Address", "Name")
	assert.Contains(t, err1.Error(), "Gateway", "Address", "Name")
	assert.Contains(t, err2.Error(), "Gateway", "Address", "Name")
}
func TestDeviceIpv6Required(t *testing.T) {
	object := gosnappi.NewDeviceIpv6()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Gateway", "Address", "Name")
	assert.Contains(t, err1.Error(), "Gateway", "Address", "Name")
	assert.Contains(t, err2.Error(), "Gateway", "Address", "Name")
}
func TestDeviceVlanRequired(t *testing.T) {
	object := gosnappi.NewDeviceVlan()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestIsisInterfaceRequired(t *testing.T) {
	object := gosnappi.NewIsisInterface()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "EthName", "Name")
	assert.Contains(t, err1.Error(), "EthName", "Name")
	assert.Contains(t, err2.Error(), "EthName", "Name")
}
func TestIsisV4RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewIsisV4RouteRange()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestIsisV6RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewIsisV6RouteRange()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpV4InterfaceRequired(t *testing.T) {
	object := gosnappi.NewBgpV4Interface()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv4Name")
	assert.Contains(t, err1.Error(), "Ipv4Name")
	assert.Contains(t, err2.Error(), "Ipv4Name")
}
func TestBgpV6InterfaceRequired(t *testing.T) {
	object := gosnappi.NewBgpV6Interface()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv6Name")
	assert.Contains(t, err1.Error(), "Ipv6Name")
	assert.Contains(t, err2.Error(), "Ipv6Name")
}
func TestFlowPortRequired(t *testing.T) {
	object := gosnappi.NewFlowPort()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "TxName")
	assert.Contains(t, err1.Error(), "TxName")
	assert.Contains(t, err2.Error(), "TxName")
}
func TestFlowCustomRequired(t *testing.T) {
	object := gosnappi.NewFlowCustom()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Bytes")
	assert.Contains(t, err1.Error(), "Bytes")
	assert.Contains(t, err2.Error(), "Bytes")
}
func TestIsisInterfaceAuthenticationRequired(t *testing.T) {
	object := gosnappi.NewIsisInterfaceAuthentication()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "AuthType")
	assert.Contains(t, err1.Error(), "AuthType")
	assert.Contains(t, err2.Error(), "AuthType")
}
func TestIsisAuthenticationBaseRequired(t *testing.T) {
	object := gosnappi.NewIsisAuthenticationBase()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "AuthType")
	assert.Contains(t, err1.Error(), "AuthType")
	assert.Contains(t, err2.Error(), "AuthType")
}
func TestV4RouteAddressRequired(t *testing.T) {
	object := gosnappi.NewV4RouteAddress()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Address")
	assert.Contains(t, err1.Error(), "Address")
	assert.Contains(t, err2.Error(), "Address")
}
func TestV6RouteAddressRequired(t *testing.T) {
	object := gosnappi.NewV6RouteAddress()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Address")
	assert.Contains(t, err1.Error(), "Address")
	assert.Contains(t, err2.Error(), "Address")
}
func TestBgpV4PeerRequired(t *testing.T) {
	object := gosnappi.NewBgpV4Peer()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
	assert.Contains(t, err1.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
	assert.Contains(t, err2.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
}
func TestBgpV6PeerRequired(t *testing.T) {
	object := gosnappi.NewBgpV6Peer()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
	assert.Contains(t, err1.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
	assert.Contains(t, err2.Error(), "PeerAddress", "AsType", "AsNumber", "Name")
}
func TestBgpV4RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewBgpV4RouteRange()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpV6RouteRangeRequired(t *testing.T) {
	object := gosnappi.NewBgpV6RouteRange()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpSrteV4PolicyRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV4Policy()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv4Endpoint", "Name")
	assert.Contains(t, err1.Error(), "Ipv4Endpoint", "Name")
	assert.Contains(t, err2.Error(), "Ipv4Endpoint", "Name")
}
func TestBgpSrteV6PolicyRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV6Policy()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv6Endpoint", "Name")
	assert.Contains(t, err1.Error(), "Ipv6Endpoint", "Name")
	assert.Contains(t, err2.Error(), "Ipv6Endpoint", "Name")
}
func TestBgpSrteV4TunnelTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV4TunnelTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpSrteV6TunnelTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteV6TunnelTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpSrteSegmentListRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentList()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name")
	assert.Contains(t, err1.Error(), "Name")
	assert.Contains(t, err2.Error(), "Name")
}
func TestBgpSrteSegmentRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegment()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "SegmentType", "Name")
	assert.Contains(t, err1.Error(), "SegmentType", "Name")
	assert.Contains(t, err2.Error(), "SegmentType", "Name")
}
func TestBgpSrteSegmentBTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentBTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Srv6Sid")
	assert.Contains(t, err1.Error(), "Srv6Sid")
	assert.Contains(t, err2.Error(), "Srv6Sid")
}
func TestBgpSrteSegmentCTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentCTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv4NodeAddress")
	assert.Contains(t, err1.Error(), "Ipv4NodeAddress")
	assert.Contains(t, err2.Error(), "Ipv4NodeAddress")
}
func TestBgpSrteSegmentDTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentDTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv6NodeAddress")
	assert.Contains(t, err1.Error(), "Ipv6NodeAddress")
	assert.Contains(t, err2.Error(), "Ipv6NodeAddress")
}
func TestBgpSrteSegmentETypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentETypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv4NodeAddress")
	assert.Contains(t, err1.Error(), "Ipv4NodeAddress")
	assert.Contains(t, err2.Error(), "Ipv4NodeAddress")
}
func TestBgpSrteSegmentFTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentFTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "LocalIpv4Address", "RemoteIpv4Address")
	assert.Contains(t, err1.Error(), "LocalIpv4Address", "RemoteIpv4Address")
	assert.Contains(t, err2.Error(), "LocalIpv4Address", "RemoteIpv4Address")
}
func TestBgpSrteSegmentGTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentGTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
	assert.Contains(t, err1.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
	assert.Contains(t, err2.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
}
func TestBgpSrteSegmentHTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentHTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "LocalIpv6Address", "RemoteIpv6Address")
	assert.Contains(t, err1.Error(), "LocalIpv6Address", "RemoteIpv6Address")
	assert.Contains(t, err2.Error(), "LocalIpv6Address", "RemoteIpv6Address")
}
func TestBgpSrteSegmentITypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentITypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Ipv6NodeAddress")
	assert.Contains(t, err1.Error(), "Ipv6NodeAddress")
	assert.Contains(t, err2.Error(), "Ipv6NodeAddress")
}
func TestBgpSrteSegmentJTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentJTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
	assert.Contains(t, err1.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
	assert.Contains(t, err2.Error(), "LocalIpv6NodeAddress", "RemoteIpv6NodeAddress")
}
func TestBgpSrteSegmentKTypeSubTlvRequired(t *testing.T) {
	object := gosnappi.NewBgpSrteSegmentKTypeSubTlv()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "LocalIpv6Address", "RemoteIpv6Address")
	assert.Contains(t, err1.Error(), "LocalIpv6Address", "RemoteIpv6Address")
	assert.Contains(t, err2.Error(), "LocalIpv6Address", "RemoteIpv6Address")
}
