package gosnappi_test

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"

	"github.com/stretchr/testify/assert"
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

func config1(api gosnappi.GosnappiApi) gosnappi.Config {
	config := api.NewConfig()
	port := config.Ports().Add()
	port.SetName("port1")
	port.SetLocation("location1")
	f1 := config.Flows().Add().SetName("f1")
	f2 := config.Flows().Add().SetName("f2")
	f1.Metrics().SetEnable(true)
	f2.Metrics().SetEnable(true)
	return config
}

func config2(api gosnappi.GosnappiApi) gosnappi.Config {
	config := api.NewConfig()
	port := config.Ports().Add()
	port.SetName("port1")
	port.SetLocation("location1")
	config.Flows().Add().SetName("f1")
	config.Flows().Add().SetName("f2")
	d1 := config.Devices().Add().SetName("d1")
	eth1 := d1.Ethernet().SetName("Ethernet1")
	ip1 := eth1.Ipv4().SetName("IPv41")
	bgp1 := ip1.Bgpv4().SetName("BGP-1")
	bgp1.Bgpv4Routes().Add().SetName("RR-1")
	return config
}

// Basic test for the grpc mock server
func TestGrpcApi(t *testing.T) {
	api := gosnappi.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetLocation(mockGrpcServerLocation).SetRequestTimeout(10000)
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
	config := api.NewConfig()
	port := config.Ports().Add()
	port.SetName("port1")
	port.SetLocation("location1")
	config.Flows().Add().SetName("f1")
	config.Flows().Add().SetName("f2")
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
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1", "f2"})
	resp, err := api.GetMetrics(req)
	fmt.Println("grpc flow response :", resp.ToYaml())
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestHttpGetMetricsFlowResponse(t *testing.T) {
	// Send Get Metrics request with flow name f2 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1", "f2"})
	resp, err := api.GetMetrics(req)
	fmt.Println("HTTP flow response :", resp.ToJson())
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.FlowMetrics().Items()[0].Name(), string("f1"))
	assert.Equal(t, resp.FlowMetrics().Items()[0].BytesTx(), int32(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[0].BytesRx(), int32(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[1].Name(), string("f2"))
	assert.Equal(t, resp.FlowMetrics().Items()[1].BytesTx(), int32(1000))
	assert.Equal(t, resp.FlowMetrics().Items()[1].BytesRx(), int32(1000))
}

func TestSetNewConfig(t *testing.T) {
	// Set new Config with flows metrics disabled, the same config will
	// be used for all the remaining tests followed
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	config := config2(api)
	api.SetConfig(config)
}

func TestGetMetrics400FlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f1 where metrics is not enabled in
	// the config, validate the err as expected
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f1"})
	_, err := api.GetMetrics(req)
	result := strings.Contains(err.Error(), "metrics not enabled for all the flows")
	assert.Equal(t, result, true)
}

func TestGetMetrics500FlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f3 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f3"})
	_, err := api.GetMetrics(req)
	result := strings.Contains(err.Error(), "requested flow is not available in configured flows")
	log.Print(result)
	// TODO: uncomment this once 500 response is working
	assert.Equal(t, result, true)
}

func TestGrpcGetMetricsPortResponse(t *testing.T) {
	// Send Get Metrics request with port name port1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	resp, err := api.GetMetrics(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestHttpGetMetricsPortResponse(t *testing.T) {
	// Send Get Metrics request with port name port1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewHttpTransport().SetLocation(mockHttpServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	resp, err := api.GetMetrics(req)
	fmt.Println("HTTP Port Response :", resp.ToJson())
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.PortMetrics().Items()[0].Name(), string("port1"))
	assert.Equal(t, resp.PortMetrics().Items()[0].BytesTx(), int32(2000))
	assert.Equal(t, resp.PortMetrics().Items()[0].BytesRx(), int32(2000))
}

func TestGetMetricsPortResponseError(t *testing.T) {
	// Send Get Metrics request with port name port2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port2"})
	_, err := api.GetMetrics(req)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports")
	assert.Equal(t, result, true)
}

func TestGrpcGetMetricsBgpv4Response(t *testing.T) {
	// Send Get Metrics request with device name d1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
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
	req := api.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetPeerNames([]string{"BGP-1"})
	resp, err := api.GetMetrics(req)
	fmt.Println(resp.ToJson())
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, resp.Bgpv4Metrics().Items()[0].Name(), string("BGP-1"))
	assert.Equal(t, resp.Bgpv4Metrics().Items()[0].RoutesAdvertised(), int32(80))
}

func TestGetMetricsBgpv4ResponseError(t *testing.T) {
	// Send Get Metrics request with port name d2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewMetricsRequest()
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
	req := api.NewTransmitState()
	req.SetFlowNames([]string{"f1", "f2"})
	req.SetState(gosnappi.TransmitStateState.START)
	assert.Equal(t, flow_names, req.FlowNames())
	resp, _ := api.SetTransmitState(req)
	assert.NotNil(t, resp)
}

func TestSetTransmitStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewTransmitState()
	req.SetFlowNames([]string{"f3"})
	req.SetState(gosnappi.TransmitStateState.START)
	_, err := api.SetTransmitState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested flow is not available in configured flows to start")
	assert.Equal(t, result, true)
}

func TestSetLinkStateResponse(t *testing.T) {
	port_names := []string{"port1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewLinkState()
	req.SetPortNames([]string{"port1"})
	req.SetState(gosnappi.LinkStateState.DOWN)
	assert.Equal(t, port_names, req.PortNames())
	resp, _ := api.SetLinkState(req)
	assert.NotNil(t, resp)
}

func TestSetLinkStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewLinkState()
	req.SetPortNames([]string{"port3"})
	req.SetState(gosnappi.LinkStateState.DOWN)
	_, err := api.SetLinkState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports to do link down")
	assert.Equal(t, result, true)
}

func TestSetCaptureStateResponse(t *testing.T) {
	port_names := []string{"port1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewCaptureState()
	req.SetPortNames([]string{"port1"})
	req.SetState(gosnappi.CaptureStateState.START)
	assert.Equal(t, port_names, req.PortNames())
	resp, _ := api.SetCaptureState(req)
	assert.NotNil(t, resp)
}

func TestSetCaptureStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewCaptureState()
	req.SetPortNames([]string{"port3"})
	req.SetState(gosnappi.CaptureStateState.START)
	_, err := api.SetCaptureState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested port is not available in configured ports to start capture")
	assert.Equal(t, result, true)
}

func TestSetRouteStateResponse(t *testing.T) {
	route_names := []string{"RR-1"}
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewRouteState()
	req.SetNames([]string{"RR-1"})
	req.SetState(gosnappi.RouteStateState.ADVERTISE)
	assert.Equal(t, route_names, req.Names())
	resp, _ := api.SetRouteState(req)
	assert.NotNil(t, resp)
}

func TestSetRouteStateResponseError(t *testing.T) {
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockGrpcServerLocation)
	req := api.NewRouteState()
	req.SetNames([]string{"RR-2"})
	req.SetState(gosnappi.RouteStateState.ADVERTISE)
	_, err := api.SetRouteState(req)
	log.Print(err)
	result := strings.Contains(err.Error(), "requested route is not available in configured routes to advertise")
	assert.Equal(t, result, true)
}

func TestPorts(t *testing.T) {
	api := gosnappi.NewApi()
	config := api.NewConfig()
	port := config.Ports().Add().SetName("p1").SetLocation("ixn-location")
	assert.Equal(t, port.Name(), "p1", "Both shall be equal")
	assert.Equal(t, port.Location(), "ixn-location", "Both shall be equal")

	port_map := []map[string]string{{"name": "p1", "location": "ixn-location"}}
	config_map := map[string][]map[string]string{"ports": port_map}
	data, err := json.Marshal(config_map)
	assert.Nil(t, err)
	config_new := api.NewConfig()
	config_new.FromJson(string(data))

	// assert.Equal(t, config, config_new, "Both configs shall be equal")
	assert.Equal(t, config.ToJson(), config_new.ToJson(), "Both json shall be equal")
}

func TestDevices(t *testing.T) {

	api := gosnappi.NewApi()
	config := api.NewConfig()
	device := config.Devices().Add().SetName("d1").SetContainerName("p1")
	assert.Equal(t, device.Name(), "d1")
	assert.Equal(t, device.ContainerName(), "p1")
	// TODO: Add validation on Json and Yaml
	device.ToJson()
	device.ToYaml()

	eth := device.Ethernet().
		SetName("Eth").
		SetMac("00:00:11:11:00:00").
		SetMtu(1500)

	assert.Equal(t, eth.Name(), "Eth")
	assert.Equal(t, eth.Mac(), "00:00:11:11:00:00")
	assert.Equal(t, eth.Mtu(), int32(1500))
	// TODO: Add validation on Json and Yaml
	eth.ToJson()
	eth.ToYaml()

	vlan := eth.Vlans().Add().SetName("vlan1").SetId(1).SetPriority(1)
	assert.Equal(t, vlan.Name(), "vlan1")
	assert.Equal(t, vlan.Id(), int32(1))
	assert.Equal(t, vlan.Priority(), int32(1))
	// TODO: Add validation on Json and Yaml
	vlan.ToJson()
	vlan.ToYaml()

	ip := eth.Ipv4().SetName("ip").SetAddress("10.1.1.1").SetGateway("10.1.1.2").SetPrefix(24)
	assert.Equal(t, ip.Name(), "ip")
	assert.Equal(t, ip.Address(), "10.1.1.1")
	assert.Equal(t, ip.Gateway(), "10.1.1.2")
	assert.Equal(t, ip.Prefix(), int32(24))
	// TODO: Add validation on Json and Yaml
	ip.ToJson()
	ip.ToYaml()

	ip6 := eth.Ipv6().SetName("ip").SetAddress("2000::1").SetGateway("2000::2").SetPrefix(64)
	assert.Equal(t, ip6.Name(), "ip")
	assert.Equal(t, ip6.Address(), "2000::1")
	assert.Equal(t, ip6.Gateway(), "2000::2")
	assert.Equal(t, ip6.Prefix(), int32(64))
	// TODO: Add validation on Json and Yaml
	ip6.ToJson()
	ip6.ToYaml()

	bgpv4 := ip.Bgpv4().SetName("bgp").SetRouterId("192.12.0.1").SetLocalAddress("10.1.1.1").SetDutAddress("10.2.2.2").
		SetAsNumber(2).SetActive(true)
	assert.Equal(t, "bgp", bgpv4.Name())
	assert.Equal(t, "192.12.0.1", bgpv4.RouterId())
	assert.Equal(t, "10.1.1.1", bgpv4.LocalAddress())
	assert.Equal(t, "10.2.2.2", bgpv4.DutAddress())
	assert.Equal(t, int32(2), bgpv4.AsNumber())
	assert.Equal(t, true, bgpv4.Active())
	// TODO: Add validation on Json and Yaml
	bgpv4.ToJson()
	bgpv4.ToYaml()

	bgpv6 := ip6.Bgpv6().SetName("bgp").SetRouterId("6000::").SetLocalAddress("2000::1").SetDutAddress("2000::2").
		SetAsNumber(2).SetActive(true)
	assert.Equal(t, "bgp", bgpv6.Name())
	assert.Equal(t, "6000::", bgpv6.RouterId())
	assert.Equal(t, "2000::1", bgpv6.LocalAddress())
	assert.Equal(t, "2000::2", bgpv6.DutAddress())
	assert.Equal(t, int32(2), bgpv6.AsNumber())
	assert.Equal(t, true, bgpv6.Active())
	// TODO: Add validation on Json and Yaml
	bgpv6.ToJson()
	bgpv6.ToYaml()

	adv := bgpv4.Advanced().SetHoldTimeInterval(10).SetKeepAliveInterval(10).SetMd5Key("abc").SetTimeToLive(10).
		SetUpdateInterval(10)
	assert.Equal(t, adv.HoldTimeInterval(), int32(10))
	assert.Equal(t, adv.KeepAliveInterval(), int32(10))
	assert.Equal(t, adv.Md5Key(), "abc")
	assert.Equal(t, adv.TimeToLive(), int32(10))
	assert.Equal(t, adv.UpdateInterval(), int32(10))
	// TODO: Add validation on Json and Yaml
	adv.ToJson()
	adv.ToYaml()

	adv6 := bgpv6.Advanced().SetHoldTimeInterval(10).SetKeepAliveInterval(10).SetMd5Key("abc").SetTimeToLive(10).
		SetUpdateInterval(10)
	assert.Equal(t, adv6.HoldTimeInterval(), int32(10))
	assert.Equal(t, adv6.KeepAliveInterval(), int32(10))
	assert.Equal(t, adv6.Md5Key(), "abc")
	assert.Equal(t, adv6.TimeToLive(), int32(10))
	assert.Equal(t, adv6.UpdateInterval(), int32(10))
	// TODO: Add validation on Json and Yaml
	adv6.ToJson()
	adv6.ToYaml()

	cap := bgpv4.Capability().SetEvpn(true).SetExtendedNextHopEncoding(true).SetIpv4Mdt(true).SetIpv4MplsVpn(true).
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
	cap.ToJson()
	cap.ToYaml()

	cap6 := bgpv6.Capability().SetEvpn(true).SetExtendedNextHopEncoding(true).SetIpv4Mdt(true).SetIpv4MplsVpn(true).
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
	cap6.ToJson()
	cap6.ToYaml()
}
