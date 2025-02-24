package gosnappi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockGrpcServerLocation = "127.0.0.1:50001"
var mockHttpServerLocation = "http://127.0.0.1:50002"

func init() {
	if err := gosnappi.StartMockGrpcServer(mockGrpcServerLocation); err != nil {
		log.Fatal("Mock Grpc server init failed")
	}
	http_path := strings.Split(mockHttpServerLocation, "//")
	gosnappi.StartMockHttpServer(http_path[1])

}

func config1(api gosnappi.Api) gosnappi.Config {
	config := gosnappi.NewConfig()
	port1 := config.Ports().Add()
	port2 := config.Ports().Add()
	port1.SetName("port1")
	port1.SetLocation("location1")
	port2.SetName("port2")
	port2.SetLocation("location2")
	f1 := config.Flows().Add().SetName("f1")
	f2 := config.Flows().Add().SetName("f2")
	f1.TxRx().Port().SetTxName(port1.Name())
	f1.TxRx().Port().SetRxName(port2.Name())
	f1.Metrics().SetEnable(true)
	f2.TxRx().Port().SetTxName(port1.Name())
	f2.TxRx().Port().SetTxName(port2.Name())
	f2.Metrics().SetEnable(true)
	f1.Rate()
	fmt.Println(config.Marshal().ToJson())
	return config
}

func config2(api gosnappi.Api) gosnappi.Config {
	config := config1(api)
	d1 := config.Devices().Add().SetName("d1")
	eth1 := d1.Ethernets().Add().
		SetName("Ethernet1").
		SetMac("00:11:01:00:00:03")
	eth1.Connection().SetPortName("port1")
	eth1.Ipv4Addresses().Add().
		SetName("IPv41").
		SetAddress("10.10.0.1").
		SetGateway("10.10.0.2")
	bgp_iP_inf := d1.Bgp().SetRouterId("1.1.1.1").
		Ipv4Interfaces().Add().
		SetIpv4Name("IPv41")
	bgp_peer := bgp_iP_inf.Peers().Add().
		SetPeerAddress("2.2.2.2").
		SetAsType(gosnappi.BgpV4PeerAsType.IBGP).
		SetName("BGP-1").
		SetAsNumber(3)
	bgp_peer.V4Routes().Add().SetName("RR-1")
	return config
}

// Basic test for the grpc mock server
func TestGrpcApi(t *testing.T) {
	api := gosnappi.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetLocation(mockGrpcServerLocation)
	config := config1(api)
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}

func TestGrpcClientConnection(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, mockGrpcServerLocation, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed grpc dialcontext due to %s", err.Error()))
	}
	api := gosnappi.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetClientConnection(conn)
	assert.NotNil(t, grpc.ClientConnection())
	config := config1(api)
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}

func TestHttpApi(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	config := config1(api)
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}

func TestGrpcGetMetricsFlowResponse(t *testing.T) {
	// Send Get Metrics request with flow name f2 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1", "f2"})
	flow_req_json, err := flow_req.Marshal().ToJson()
	assert.Nil(t, err)
	flow_req_yaml, err := flow_req.Marshal().ToYaml()
	log.Print(flow_req_json, flow_req_yaml)
	resp, err := api.GetMetrics(req)
	resp_json, err := resp.Marshal().ToYaml()
	assert.Nil(t, err)
	log.Print("grpc flow response :", resp_json)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestHttpGetMetricsFlowResponse(t *testing.T) {
	// Send Get Metrics request with flow name f2 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1", "f2"})
	resp, err := api.GetMetrics(req)
	resp_json, err := resp.Marshal().ToJson()
	assert.Nil(t, err)
	fmt.Println("HTTP flow response :", resp_json)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.FlowMetrics().Items()[0].Name(), string("f1"))
	assert.Equal(t, resp.FlowMetrics().Items()[0].BytesTx(), uint64(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[0].BytesRx(), uint64(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[1].Name(), string("f2"))
	assert.Equal(t, resp.FlowMetrics().Items()[1].BytesTx(), uint64(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[1].BytesRx(), uint64(1000))
}

func TestSetNewConfig(t *testing.T) {
	// Set new Config with flows metrics disabled, the same config will
	// be used for all the remaining tests followed
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config2(api)
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
}

func TestGetMetrics400FlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f1 where metrics is not enabled in
	// the config, validate the err as expected
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config1(api)
	config.Flows().Items()[0].Metrics().SetEnable(false)
	api.SetConfig(config)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1"})
	_, err := api.GetMetrics(req)
	assert.Contains(t, err.Error(), "metrics not enabled for all the flows")
}

func TestGetMetrics500FlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f3 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f3"})
	_, err := api.GetMetrics(req)
	result := strings.Contains(err.Error(), "requested flow is not available in configured flows")
	log.Print(result)
	assert.Equal(t, result, true)
}

func TestGrpcGetMetricsPortResponse(t *testing.T) {
	// Send Get Metrics request with port name port1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	flow_req_json, err := flow_req.Marshal().ToJson()
	assert.Nil(t, err)
	flow_req_yaml, err := flow_req.Marshal().ToYaml()
	log.Print(flow_req_json, flow_req_yaml)
	resp, err := api.GetMetrics(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestHttpGetMetricsPortResponse(t *testing.T) {
	// Send Get Metrics request with port name port1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	resp, err := api.GetMetrics(req)
	resp_json, err := resp.Marshal().ToJson()
	assert.Nil(t, err)
	fmt.Println("HTTP Port Response :", resp_json)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.PortMetrics().Items()[0].Name(), string("port1"))
	assert.Equal(t, resp.PortMetrics().Items()[0].BytesTx(), uint64(2000))
	assert.Equal(t, resp.PortMetrics().Items()[0].BytesRx(), uint64(2000))
}

func TestGetMetricsPortResponseError(t *testing.T) {
	// Send Get Metrics request with port name port2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config1(api)
	config.Flows().Items()[0].Metrics().SetEnable(false)
	api.SetConfig(config)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port3"})
	_, err := api.GetMetrics(req)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports")
	assert.Equal(t, result, true)
}

func TestGrpcGetMetricsBgpv4Response(t *testing.T) {
	// Send Get Metrics request with device name d1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config2(api)
	_, err := api.SetConfig(config)
	assert.Nil(t, err)
	req := gosnappi.NewMetricsRequest()
	req.Bgpv4()
	flow_req := req.Bgpv4()
	flow_req.SetPeerNames([]string{"BGP-1"})
	resp, err := api.GetMetrics(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestHttpGetMetricsBgpv4Response(t *testing.T) {
	// Send Get Metrics request with device name d1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetPeerNames([]string{"BGP-1"})
	resp, err := api.GetMetrics(req)
	log.Print(resp.Marshal().ToJson())
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.Bgpv4Metrics().Items()[0].Name(), string("BGP-1"))
	assert.Equal(t, resp.Bgpv4Metrics().Items()[0].RoutesAdvertised(), uint64(80))
}

func TestGetMetricsBgpv4ResponseError(t *testing.T) {
	// Send Get Metrics request with port name d2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetPeerNames([]string{"d2"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "d2 is not a valid BGPv4 device")
	assert.Equal(t, result, true)
}

func TestSetTransmitStateResponse(t *testing.T) {
	flow_names := []string{"f1", "f2"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	ft := req.Traffic().FlowTransmit()
	ft.SetFlowNames([]string{"f1", "f2"})
	ft.SetState(gosnappi.StateTrafficFlowTransmitState.START)
	assert.Equal(t, flow_names, ft.FlowNames())
	req_json, err := req.Marshal().ToJson()
	assert.Nil(t, err)
	req_yaml, err := req.Marshal().ToYaml()
	assert.Nil(t, err)
	log.Print(req_json, req_yaml)
	resp, _ := api.SetControlState(req)
	assert.NotNil(t, resp)
}

func TestSetTransmitStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	ft := req.Traffic().FlowTransmit()
	ft.SetFlowNames([]string{"f3"})
	ft.SetState(gosnappi.StateTrafficFlowTransmitState.START)
	_, err := api.SetControlState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested flow is not available in configured flows to start")
	assert.Equal(t, result, true)
}

func TestSetLinkStateResponse(t *testing.T) {
	port_names := []string{"port1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	ls := req.Port().Link()
	ls.SetPortNames([]string{"port1"})
	ls.SetState(gosnappi.StatePortLinkState.DOWN)
	req_json, err := req.Marshal().ToJson()
	assert.Nil(t, err)
	req_yaml, err := req.Marshal().ToYaml()
	assert.Nil(t, err)
	log.Print(req_json, req_yaml)
	assert.Equal(t, port_names, ls.PortNames())
	resp, _ := api.SetControlState(req)
	assert.NotNil(t, resp)
}

func TestSetLinkStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	ls := req.Port().Link()
	ls.SetPortNames([]string{"port3"})
	ls.SetState(gosnappi.StatePortLinkState.DOWN)
	_, err := api.SetControlState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports to do link down")
	assert.Equal(t, result, true)
}

func TestSetCaptureStateResponse(t *testing.T) {
	port_names := []string{"port1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	cp := req.Port().Capture()
	cp.SetPortNames([]string{"port1"})
	cp.SetState(gosnappi.StatePortCaptureState.START)
	req_json, err := req.Marshal().ToJson()
	assert.Nil(t, err)
	req_yaml, err := req.Marshal().ToYaml()
	assert.Nil(t, err)
	log.Print(req_json, req_yaml)
	assert.Equal(t, port_names, cp.PortNames())
	resp, _ := api.SetControlState(req)
	assert.NotNil(t, resp)
}

func TestSetCaptureStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	cp := req.Port().Capture()
	cp.SetPortNames([]string{"port3"})
	cp.SetState(gosnappi.StatePortCaptureState.START)
	_, err := api.SetControlState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports to start capture")
	assert.Equal(t, result, true)
}

func TestSetRouteStateResponse(t *testing.T) {
	route_names := []string{"RR-1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config2(api)
	_, err := api.SetConfig(config)
	assert.Nil(t, err)
	req := gosnappi.NewControlState()
	rs := req.Protocol().Route()
	rs.SetNames([]string{"RR-1"})
	rs.SetState(gosnappi.StateProtocolRouteState.ADVERTISE)
	req_json, err := req.Marshal().ToJson()
	assert.Nil(t, err)
	req_yaml, err := req.Marshal().ToYaml()
	assert.Nil(t, err)
	log.Print(req_json, req_yaml)
	assert.Equal(t, route_names, rs.Names())
	resp, _ := api.SetControlState(req)
	assert.NotNil(t, resp)
}

func TestSetRouteStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := gosnappi.NewControlState()
	rs := req.Protocol().Route()
	rs.SetNames([]string{"RR-2"})
	rs.SetState(gosnappi.StateProtocolRouteState.ADVERTISE)
	_, err := api.SetControlState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested route is not available in configured routes to advertise")
	assert.Equal(t, result, true)
}

func TestPorts(t *testing.T) {
	config := gosnappi.NewConfig()
	port := config.Ports().Add().SetName("p1").SetLocation("ixn-location")
	assert.Equal(t, port.Name(), "p1", "Both shall be equal")
	assert.Equal(t, port.Location(), "ixn-location", "Both shall be equal")

	port_map := []map[string]string{{"name": "p1", "location": "ixn-location"}}
	config_map := map[string][]map[string]string{"ports": port_map}
	data, err := json.Marshal(config_map)
	assert.Nil(t, err)
	config_new := gosnappi.NewConfig()
	config_new.Unmarshal().FromJson(string(data))

	config_json, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	config_new_json, err := config_new.Marshal().ToJson()
	assert.Nil(t, err)
	// assert.Equal(t, config, config_new, "Both configs shall be equal")
	assert.Equal(t, config_json, config_new_json, "Both json shall be equal")
}

func TestDevices(t *testing.T) {
	config := gosnappi.NewConfig()
	device := config.Devices().Add().SetName("d1")
	assert.Equal(t, device.Name(), "d1")
	// TODO: Add validation on Json and Yaml
	device.Marshal().ToJson()
	device.Marshal().ToYaml()
	device.Marshal().ToPbText()

	eth := device.Ethernets().Add().
		SetName("Eth").
		SetMac("00:00:11:11:00:00").
		SetMtu(1500)

	eth.Connection().SetPortName("p1")
	assert.Equal(t, eth.Connection().PortName(), "p1")
	assert.Equal(t, eth.Name(), "Eth")
	assert.Equal(t, eth.Mac(), "00:00:11:11:00:00")
	assert.Equal(t, eth.Mtu(), uint32(1500))
	// TODO: Add validation on Json and Yaml
	eth.Marshal().ToJson()
	eth.Marshal().ToYaml()
	eth.Marshal().ToPbText()

	vlan := eth.Vlans().Add().SetName("vlan1").SetId(1).SetPriority(1)
	assert.Equal(t, vlan.Name(), "vlan1")
	assert.Equal(t, vlan.Id(), uint32(1))
	assert.Equal(t, vlan.Priority(), uint32(1))
	// TODO: Add validation on Json and Yaml
	vlan.Marshal().ToJson()
	vlan.Marshal().ToYaml()
	vlan.Marshal().ToPbText()

	ip := eth.Ipv4Addresses().Add().
		SetName("ipv4").
		SetAddress("10.1.1.1").
		SetGateway("10.1.1.2").
		SetPrefix(24)
	assert.Equal(t, ip.Name(), "ipv4")
	assert.Equal(t, ip.Address(), "10.1.1.1")
	assert.Equal(t, ip.Gateway(), "10.1.1.2")
	assert.Equal(t, ip.Prefix(), uint32(24))
	// TODO: Add validation on Json and Yaml
	ip.Marshal().ToJson()
	ip.Marshal().ToYaml()
	ip.Marshal().ToPbText()

	ip6 := eth.Ipv6Addresses().Add().
		SetName("ipv6").
		SetAddress("2000::1").
		SetGateway("2000::2").
		SetPrefix(64)
	assert.Equal(t, ip6.Name(), "ipv6")
	assert.Equal(t, ip6.Address(), "2000::1")
	assert.Equal(t, ip6.Gateway(), "2000::2")
	assert.Equal(t, ip6.Prefix(), uint32(64))
	// TODO: Add validation on Json and Yaml
	ip6.Marshal().ToJson()
	ip6.Marshal().ToYaml()
	ip6.Marshal().ToPbText()

	bgp := device.Bgp().SetRouterId("192.12.0.1")
	assert.Equal(t, bgp.RouterId(), "192.12.0.1")
	// TODO: Add validation on Json and Yaml
	bgp.Marshal().ToJson()
	bgp.Marshal().ToYaml()
	bgp.Marshal().ToPbText()

	bgpv4Int := bgp.Ipv4Interfaces().Add().
		SetIpv4Name("bgpv4Int")
	assert.Equal(t, bgpv4Int.Ipv4Name(), "bgpv4Int")
	bgpv4Int.Marshal().ToJson()
	bgpv4Int.Marshal().ToYaml()
	bgpv4Int.Marshal().ToPbText()

	bgpv4Peer := bgpv4Int.Peers().Add().
		SetName("bgpv4Peer").
		SetAsNumber(3).
		SetAsNumberWidth(gosnappi.BgpV4PeerAsNumberWidth.TWO).
		SetAsType(gosnappi.BgpV4PeerAsType.EBGP).
		SetPeerAddress("10.2.2.2")
	assert.Equal(t, bgpv4Peer.Name(), "bgpv4Peer")
	assert.Equal(t, bgpv4Peer.AsNumber(), uint32(3))
	assert.Equal(t, bgpv4Peer.AsNumberWidth(), gosnappi.BgpV4PeerAsNumberWidth.TWO)
	assert.Equal(t, bgpv4Peer.AsType(), gosnappi.BgpV4PeerAsType.EBGP)
	assert.Equal(t, bgpv4Peer.PeerAddress(), "10.2.2.2")
	bgpv4Peer.Marshal().ToJson()
	bgpv4Peer.Marshal().ToYaml()
	bgpv4Peer.Marshal().ToPbText()

	bgpv6Int := bgp.Ipv6Interfaces().Add().
		SetIpv6Name("bgpv6Int")
	assert.Equal(t, bgpv6Int.Ipv6Name(), "bgpv6Int")
	bgpv6Int.Marshal().ToJson()
	bgpv6Int.Marshal().ToYaml()
	bgpv6Int.Marshal().ToPbText()

	bgpv6Peer := bgpv6Int.Peers().Add().
		SetName("bgpv6Peer").
		SetAsNumber(3).
		SetAsNumberWidth(gosnappi.BgpV6PeerAsNumberWidth.FOUR).
		SetAsType(gosnappi.BgpV6PeerAsType.IBGP).
		SetPeerAddress("2000::1")
	assert.Equal(t, bgpv6Peer.Name(), "bgpv6Peer")
	assert.Equal(t, bgpv6Peer.AsNumber(), uint32(3))
	assert.Equal(t, bgpv6Peer.AsNumberWidth(), gosnappi.BgpV6PeerAsNumberWidth.FOUR)
	assert.Equal(t, bgpv6Peer.AsType(), gosnappi.BgpV6PeerAsType.IBGP)
	assert.Equal(t, bgpv6Peer.PeerAddress(), "2000::1")
	bgpv6Peer.Marshal().ToJson()
	bgpv6Peer.Marshal().ToYaml()
	bgpv6Peer.Marshal().ToPbText()

	adv := bgpv4Peer.Advanced().
		SetHoldTimeInterval(10).
		SetKeepAliveInterval(10).
		SetMd5Key("abc").
		SetTimeToLive(10).
		SetUpdateInterval(10)
	assert.Equal(t, adv.HoldTimeInterval(), uint32(10))
	assert.Equal(t, adv.KeepAliveInterval(), uint32(10))
	assert.Equal(t, adv.Md5Key(), "abc")
	assert.Equal(t, adv.TimeToLive(), uint32(10))
	assert.Equal(t, adv.UpdateInterval(), uint32(10))
	// TODO: Add validation on Json and Yaml
	adv.Marshal().ToJson()
	adv.Marshal().ToYaml()
	adv.Marshal().ToPbText()

	adv6 := bgpv6Peer.Advanced().
		SetHoldTimeInterval(10).
		SetKeepAliveInterval(10).
		SetMd5Key("abc").
		SetTimeToLive(10).
		SetUpdateInterval(10)
	assert.Equal(t, adv6.HoldTimeInterval(), uint32(10))
	assert.Equal(t, adv6.KeepAliveInterval(), uint32(10))
	assert.Equal(t, adv6.Md5Key(), "abc")
	assert.Equal(t, adv6.TimeToLive(), uint32(10))
	assert.Equal(t, adv6.UpdateInterval(), uint32(10))
	// TODO: Add validation on Json and Yaml
	adv6.Marshal().ToJson()
	adv6.Marshal().ToYaml()
	adv6.Marshal().ToPbText()

	cap := bgpv4Peer.Capability().SetEvpn(true).SetExtendedNextHopEncoding(true).SetIpv4Mdt(true).SetIpv4MplsVpn(true).
		SetIpv4Multicast(true).SetIpv4MulticastMplsVpn(true).SetIpv4MulticastVpn(true).SetIpv4SrTePolicy(true).
		SetIpv4Unicast(true).SetIpv4UnicastAddPath(true).SetIpv4UnicastFlowSpec(true).SetIpv6Mdt(true).
		SetIpv6MplsVpn(true).SetIpv6Multicast(true).SetIpv6MulticastMplsVpn(true).SetIpv6MulticastVpn(true).
		SetIpv6SrTePolicy(true).SetIpv6Unicast(true).SetIpv6UnicastAddPath(true).SetIpv6UnicastFlowSpec(true).
		SetLinkStateNonVpn(true).SetLinkStateVpn(false).SetRouteConstraint(false).SetRouteRefresh(false).
		SetVpls(false)
	assert.Equal(t, cap.Evpn(), true)
	assert.Equal(t, cap.ExtendedNextHopEncoding(), true)
	assert.Equal(t, cap.Ipv4Mdt(), true)
	assert.Equal(t, cap.Ipv4MplsVpn(), true)
	assert.Equal(t, cap.Ipv4Multicast(), true)
	assert.Equal(t, cap.Ipv4MulticastMplsVpn(), true)
	assert.Equal(t, cap.Ipv4MulticastVpn(), true)
	assert.Equal(t, cap.Ipv4SrTePolicy(), true)
	assert.Equal(t, cap.Ipv4Unicast(), true)
	assert.Equal(t, cap.Ipv4UnicastAddPath(), true)
	assert.Equal(t, cap.Ipv4UnicastFlowSpec(), true)
	assert.Equal(t, cap.Ipv6Mdt(), true)
	assert.Equal(t, cap.Ipv6MplsVpn(), true)
	assert.Equal(t, cap.Ipv6Multicast(), true)
	assert.Equal(t, cap.Ipv6MulticastMplsVpn(), true)
	assert.Equal(t, cap.Ipv6MulticastVpn(), true)
	assert.Equal(t, cap.Ipv6SrTePolicy(), true)
	assert.Equal(t, cap.Ipv6SrTePolicy(), true)
	assert.Equal(t, cap.Ipv6Unicast(), true)
	assert.Equal(t, cap.Ipv6UnicastAddPath(), true)
	assert.Equal(t, cap.Ipv6UnicastFlowSpec(), true)
	assert.Equal(t, cap.LinkStateNonVpn(), true)
	assert.Equal(t, cap.LinkStateVpn(), false)
	assert.Equal(t, cap.RouteConstraint(), false)
	assert.Equal(t, cap.RouteRefresh(), false)
	// TODO: Add validation on Json and Yaml
	cap.Marshal().ToJson()
	cap.Marshal().ToYaml()
	cap.Marshal().ToPbText()

	cap6 := bgpv6Peer.Capability().SetEvpn(true).SetExtendedNextHopEncoding(true).SetIpv4Mdt(true).SetIpv4MplsVpn(true).
		SetIpv4Multicast(true).SetIpv4MulticastMplsVpn(true).SetIpv4MulticastVpn(true).SetIpv4SrTePolicy(true).
		SetIpv4Unicast(true).SetIpv4UnicastAddPath(true).SetIpv4UnicastFlowSpec(true).SetIpv6Mdt(true).
		SetIpv6MplsVpn(true).SetIpv6Multicast(true).SetIpv6MulticastMplsVpn(true).SetIpv6MulticastVpn(true).
		SetIpv6SrTePolicy(true).SetIpv6Unicast(true).SetIpv6UnicastAddPath(true).SetIpv6UnicastFlowSpec(true).
		SetLinkStateNonVpn(true).SetLinkStateVpn(false).SetRouteConstraint(false).SetRouteRefresh(false).
		SetVpls(false)

	assert.Equal(t, cap6.Evpn(), true)
	assert.Equal(t, cap6.ExtendedNextHopEncoding(), true)
	assert.Equal(t, cap6.Ipv4Mdt(), true)
	assert.Equal(t, cap6.Ipv4MplsVpn(), true)
	assert.Equal(t, cap6.Ipv4Multicast(), true)
	assert.Equal(t, cap6.Ipv4MulticastMplsVpn(), true)
	assert.Equal(t, cap6.Ipv4MulticastVpn(), true)
	assert.Equal(t, cap6.Ipv4SrTePolicy(), true)
	assert.Equal(t, cap6.Ipv4Unicast(), true)
	assert.Equal(t, cap6.Ipv4UnicastAddPath(), true)
	assert.Equal(t, cap6.Ipv4UnicastFlowSpec(), true)
	assert.Equal(t, cap6.Ipv6Mdt(), true)
	assert.Equal(t, cap6.Ipv6MplsVpn(), true)
	assert.Equal(t, cap6.Ipv6Multicast(), true)
	assert.Equal(t, cap6.Ipv6MulticastMplsVpn(), true)
	assert.Equal(t, cap6.Ipv6MulticastVpn(), true)
	assert.Equal(t, cap6.Ipv6SrTePolicy(), true)
	assert.Equal(t, cap6.Ipv6SrTePolicy(), true)
	assert.Equal(t, cap6.Ipv6Unicast(), true)
	assert.Equal(t, cap6.Ipv6UnicastAddPath(), true)
	assert.Equal(t, cap6.Ipv6UnicastFlowSpec(), true)
	assert.Equal(t, cap6.LinkStateNonVpn(), true)
	assert.Equal(t, cap6.LinkStateVpn(), false)
	assert.Equal(t, cap6.RouteConstraint(), false)
	assert.Equal(t, cap6.RouteRefresh(), false)
	// TODO: Add validation on Json and Yaml
	cap6.Marshal().ToJson()
	cap6.Marshal().ToYaml()
	cap6.Marshal().ToPbText()
}

func TestFlows(t *testing.T) {
	flow_names := []string{"f1", "f2"}
	port_names := []string{"p1", "p2"}
	config := gosnappi.NewConfig()
	port1 := config.Ports().Add().SetName("p1")
	port2 := config.Ports().Add().SetName("p2")
	flow1 := config.Flows().Add().SetName("f1")
	flow2 := config.Flows().Add().SetName("f2")
	flow1.TxRx().Port().SetTxName(port1.Name()).SetRxName(port2.Name())
	flow2.TxRx().Port().SetTxName(port1.Name()).SetRxName(port2.Name())
	for i, P := range config.Ports().Items() {
		assert.Equal(t, port_names[i], P.Name())
	}
	for i, F := range config.Flows().Items() {
		assert.Equal(t, flow_names[i], F.Name())
		assert.Equal(t, port_names[0], F.TxRx().Port().TxName())
		assert.Equal(t, port_names[1], F.TxRx().Port().RxName())
	}

	eth := flow1.Packet().Add().Ethernet()
	eth.Src().SetValue("a2:71:d4:59:b3:b8")
	eth.Dst().SetValue("6e:d2:47:65:aa:69")
	assert.Equal(t, "a2:71:d4:59:b3:b8", eth.Src().Value())
	assert.Equal(t, "6e:d2:47:65:aa:69", eth.Dst().Value())

	ipv4 := flow1.Packet().Add().Ipv4()
	ipv4.Src().SetValue("10.10.10.1")
	ipv4.Dst().SetValue("10.10.10.2")
	assert.Equal(t, "10.10.10.1", ipv4.Src().Value())
	assert.Equal(t, "10.10.10.2", ipv4.Dst().Value())

	udp := flow1.Packet().Add().Udp()
	udp.SrcPort().SetValue(3000)
	udp.DstPort().SetValue(4000)
	udp.Checksum().SetCustom(1)
	assert.Equal(t, uint32(3000), udp.SrcPort().Value())
	assert.Equal(t, uint32(4000), udp.DstPort().Value())
	assert.Equal(t, uint32(1), udp.Checksum().Custom())

	flow1.Duration().FixedPackets().SetPackets(10000).SetGap(2).Delay().SetBytes(8)
	flow1.Rate().SetPps(1000)
	assert.Equal(t, uint32(10000), flow1.Duration().FixedPackets().Packets())
	assert.Equal(t, uint32(2), flow1.Duration().FixedPackets().Gap())
	assert.Equal(t, float32(8), flow1.Duration().FixedPackets().Delay().Bytes())
	assert.Equal(t, uint64(1000), flow1.Rate().Pps())
	log.Print(config.Marshal().ToYaml())
}

func TestValidation(t *testing.T) {
	config := gosnappi.NewConfig()
	// To Validate required field
	p := config.Ports().Add()
	_, err1 := p.Marshal().ToYaml()
	assert.Contains(t, err1.Error(), "Name is required")
	d := config.Devices().Add().SetName("d1")

	// To Validate Formats
	//
	// Mac address Validation
	eth := d.Ethernets().Add().SetName("Eth1")
	eth.SetMac("10.1.1.1")
	_, err3 := eth.Marshal().ToYaml()
	fmt.Println(err3.Error())
	assert.Contains(t, strings.ToLower(err3.Error()), "invalid mac address")

	// Ipv4 address Validation
	ipv4 := eth.Ipv4Addresses().Add().SetName("ipv4").SetGateway("20.1.1.1").SetAddress("ff.1.1.1")
	_, err4 := ipv4.Marshal().ToYaml()
	assert.Contains(t, strings.ToLower(err4.Error()), "invalid ipv4 address")

	// Ipv6 address Validation
	ipv6 := eth.Ipv6Addresses().Add().SetName("ipv6").SetGateway("10.1.1.1").SetAddress("abcd:::abcd")
	_, err5 := ipv6.Marshal().ToYaml()
	assert.Contains(t, strings.ToLower(err5.Error()), "invalid ipv6 address")

	f := config.Flows().Add().SetName("f1")
	f.TxRx().Port().SetRxName("rx").SetTxName("tx")
	floweth := f.Packet().Add().Ethernet()
	// Mac address slice validation
	floweth.Src().SetValues([]string{"abcd:abcd", "ab:ab:ab:ab:ac:ff"})
	_, f_err := floweth.Src().Marshal().ToYaml()
	fmt.Println(f_err.Error())
	assert.Contains(t, strings.ToLower(f_err.Error()), "invalid mac addresses at")

	// Ipv4 address slice validation
	flowV4 := f.Packet().Add().Ipv4()
	flowV4.Src().SetValues([]string{"1111", "1.1.1.1"})
	_, v4_err := flowV4.Marshal().ToYaml()
	fmt.Println(v4_err)
	assert.Contains(t, strings.ToLower(v4_err.Error()), "invalid ipv4 addresses at")

	// Ipv6 address slice validation
	flowV6 := f.Packet().Add().Ipv6()
	flowV6.Dst().SetValues([]string{"abcd:abcd::1234", "::", "10.1.1.1"})
	_, v6_err := flowV6.Marshal().ToYaml()
	assert.Contains(t, strings.ToLower(v6_err.Error()), "invalid ipv6 addresses at")

	_, err := config.Marshal().ToYaml()
	fmt.Println(err)
}

var expected_device_json = `{
	"devices":  [
		{
		"ethernets":  [
			{
			"port_name":  "p1",
			"ipv4_addresses":  [
				{
				"gateway":  "10.1.1.2",
				"address":  "10.1.1.1",
				"prefix":  24,
				"name":  "ipv4"
				}
			],
			"ipv6_addresses":  [
				{
				"gateway":  "2000::2",
				"address":  "2000::1",
				"prefix":  64,
				"name":  "ipv6"
				}
			],
			"mac":  "00:00:11:11:00:00",
			"mtu":  1500,
			"vlans":  [
				{
				"tpid":  "x8100",
				"priority":  0,
				"id":  1,
				"name":  "vlan1"
				}
			],
			"name":  "Eth"
			}
		],
		"ipv4_loopbacks":  [
			{
			"eth_name":  "Eth",
			"address":  "0.0.0.0",
			"name":  "IPv4loopback"
			}
		],
		"ipv6_loopbacks":  [
			{
			"eth_name":  "Eth",
			"address":  "::0",
			"name":  "IPv6loopback"
			}
		],
		"bgp":  {
			"router_id":  "192.12.0.1",
			"ipv4_interfaces":  [
			{
				"ipv4_name":  "bgpv4Int",
				"peers":  [
				{
					"peer_address":  "10.2.2.2",
					"as_type":  "ebgp",
					"as_number":  3,
					"as_number_width":  "four",
					"advanced":  {
					"hold_time_interval":  90,
					"keep_alive_interval":  30,
					"update_interval":  0,
					"time_to_live":  64
					},
					"capability":  {
					"ipv4_unicast":  true,
					"ipv4_multicast":  false,
					"ipv6_unicast":  true,
					"ipv6_multicast":  false,
					"vpls":  false,
					"route_refresh":  true,
					"route_constraint":  false,
					"link_state_non_vpn":  false,
					"link_state_vpn":  false,
					"evpn":  false,
					"extended_next_hop_encoding":  false,
					"ipv4_multicast_vpn":  false,
					"ipv4_mpls_vpn":  false,
					"ipv4_mdt":  false,
					"ipv4_multicast_mpls_vpn":  false,
					"ipv4_unicast_flow_spec":  false,
					"ipv4_sr_te_policy":  false,
					"ipv4_unicast_add_path":  false,
					"ipv6_multicast_vpn":  false,
					"ipv6_mpls_vpn":  false,
					"ipv6_mdt":  false,
					"ipv6_multicast_mpls_vpn":  false,
					"ipv6_unicast_flow_spec":  false,
					"ipv6_sr_te_policy":  false,
					"ipv6_unicast_add_path":  false
					},
					"name":  "bgpv4Peer"
				}
				]
			}
			]
		},
		"name":  "d1"
		}
	]}`

func TestDefaultsDevice(t *testing.T) {
	config := gosnappi.NewConfig()
	device := config.Devices().Add().SetName("d1")
	eth := device.Ethernets().Add().
		SetName("Eth").SetMac("00:00:11:11:00:00")
	eth.Connection().SetPortName("p1")
	eth.Vlans().Add().SetName("vlan1")
	eth.Ipv4Addresses().Add().
		SetName("ipv4").SetAddress("10.1.1.1").SetGateway("10.1.1.2")
	eth.Ipv6Addresses().Add().
		SetName("ipv6").SetAddress("2000::1").SetGateway("2000::2")

	device.Ipv4Loopbacks().Add().
		SetEthName("Eth").SetName("IPv4loopback")
	device.Ipv6Loopbacks().Add().
		SetEthName("Eth").SetName("IPv6loopback")

	bgp := device.Bgp().SetRouterId("192.12.0.1")
	bgpv4Int := bgp.Ipv4Interfaces().Add().
		SetIpv4Name("bgpv4Int")
	bgpv4peer := bgpv4Int.Peers().Add().
		SetName("bgpv4Peer").
		SetAsNumber(3).SetAsType(gosnappi.BgpV4PeerAsType.EBGP).
		SetPeerAddress("10.2.2.2")
	bgpv4peer.Advanced()
	bgpv4peer.Capability()

	_, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	// require.JSONEq(t, expected_device_json, expected_result)
}

func TestDefaultsDeviceFromJson(t *testing.T) {
	input_str := `{
		"devices":  [
			{
			"ethernets":  [
				{
				"port_name":  "p1",
				"ipv4_addresses":  [
					{
					"gateway":  "10.1.1.2",
					"address":  "10.1.1.1",
					"name":  "ipv4"
					}
				],
				"ipv6_addresses":  [
					{
					"gateway":  "2000::2",
					"address":  "2000::1",
					"name":  "ipv6"
					}
				],
				"mac":  "00:00:11:11:00:00",
				"vlans":  [
					{
					"name":  "vlan1"
					}
				],
				"name":  "Eth"
				}
			],
			"ipv4_loopbacks":  [
				{
				"eth_name":  "Eth",
				"name":  "IPv4loopback"
				}
			],
			"ipv6_loopbacks":  [
				{
				"eth_name":  "Eth",
				"name":  "IPv6loopback"
				}
			],
			"bgp":  {
				"router_id":  "192.12.0.1",
				"ipv4_interfaces":  [
				{
					"ipv4_name":  "bgpv4Int",
					"peers":  [
					{
						"peer_address":  "10.2.2.2",
						"as_type":  "ebgp",
						"as_number":  3,
						"advanced":  {
						},
						"capability":  {
						},
						"name":  "bgpv4Peer"
					}
					]
				}
				]
			},
			"name":  "d1"
			}
		]}`
	config := gosnappi.NewConfig()
	config.Unmarshal().FromJson(input_str)
	_, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	// require.JSONEq(t, expected_device_json, expected_result)
}

var expected_flow_json = `{
	"ports":  [
		{
		"name":  "p1"
		},
		{
		"name":  "p2"
		}
	],
	"flows":  [
		{
		"tx_rx":  {
			"choice":  "port",
			"port":  {
			"tx_name":  "p1",
			"rx_name":  "p2"
			}
		},
		"packet":  [
			{
			"choice":  "ethernet",
			"ethernet":  {
				"dst":  {
				"choice":  "auto",
				"auto":  "00:00:00:00:00:00"
				},
				"src":  {
				"choice":  "value",
				"value":  "00:00:00:00:00:00"
				},
				"ether_type":  {
				"choice":  "auto",
				"auto":  65535
				}
			}
			},
			{
			"choice":  "ipv4",
			"ipv4":  {
				"version":  {
				"choice":  "value",
				"value":  4
				},
				"priority":  {
				"choice":  "dscp",
				"dscp": {}
				},
				"src":  {
				"choice":  "value",
				"value":  "0.0.0.0"
				},
				"dst":  {
				"choice":  "value",
				"value":  "0.0.0.0"
				}
			}
			},
			{
			"choice":  "udp",
			"udp":  {
				"src_port":  {
				"choice":  "value",
				"value":  0
				},
				"dst_port":  {
				"choice":  "value",
				"value":  0
				},
				"checksum":  {
				"choice":  "generated",
				"generated":  "good"
				}
			}
			}
		],
		"rate":  {
			"choice":  "pps",
			"pps":  "1000"
		},
		"duration":  {
			"choice":  "continuous",
			"continuous":  {
			"gap":  12
			}
		},
		"name":  "f1"
		}
	]}`

func TestDefaultsFlow(t *testing.T) {
	config := gosnappi.NewConfig()
	port1 := config.Ports().Add().SetName("p1")
	port2 := config.Ports().Add().SetName("p2")
	flow1 := config.Flows().Add().SetName("f1")
	flow1.TxRx().Port().SetTxName(port1.Name()).SetRxName(port2.Name())

	eth := flow1.Packet().Add().Ethernet()
	eth.Src()
	eth.Dst()
	eth.EtherType()

	ipv4 := flow1.Packet().Add().Ipv4()
	ipv4.Src()
	ipv4.Dst()
	ipv4.Version()
	ipv4.Priority()

	udp := flow1.Packet().Add().Udp()
	udp.SrcPort()
	udp.DstPort()
	udp.Checksum()

	flow1.Duration().Continuous()
	flow1.Rate()
	expected_result, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, expected_flow_json, expected_result)
}

func TestFromJsonUnmarshalError(t *testing.T) {
	input_str := `{
		"state": "run"
	  }`
	req := gosnappi.NewControlState().Protocol().All()
	err := req.Unmarshal().FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 2:12): invalid value for enum field state: "run"`)
}
