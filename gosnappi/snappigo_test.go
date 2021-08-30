package gosnappi_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"

	"github.com/stretchr/testify/assert"
)

var mockServerLocation = "127.0.0.1:50001"
var mockHttpServerLocation = "http://127.0.0.1:50002"

func init() {
	if err := gosnappi.StartMockServer(mockServerLocation); err != nil {
		log.Fatal("Mock Grpc server init failed")
	}
	http_path := strings.Split(mockHttpServerLocation, "//")
	gosnappi.StartMockHttpServer(http_path[1])

}

// Basic test for the grpc mock server
func TestGrpcApi(t *testing.T) {
	api := gosnappi.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetLocation(mockServerLocation).SetRequestTimeout(10000)
	config := api.NewConfig()
	port := config.Ports().Add()
	port.SetName("port1")
	port.SetLocation("location1")
	config.Flows().Add().SetName("f1")
	config.Flows().Add().SetName("f2")
	d1 := config.Devices().Add().SetName("d1")
	eth1 := d1.Ethernet().SetName("Ethernet1")
	ip1 := eth1.Ipv4().SetName("IPv41")
	ip1.Bgpv4().SetName("BGP-1")
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
	d1 := config.Devices().Add().SetName("d1")
	eth1 := d1.Ethernet().SetName("Ethernet1")
	ip1 := eth1.Ipv4().SetName("IPv41")
	ip1.Bgpv4().SetName("BGP-1")
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}

func TestGetMetricsFlowResponse(t *testing.T) {
	// Send Get Metrics request with flow name f2 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f2"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsFlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f3 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f3"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	assert.NotNil(t, err)
}

func TestGetMetricsPortResponse(t *testing.T) {
	// Send Get Metrics request with port name port1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsPortResponseError(t *testing.T) {
	// Send Get Metrics request with port name port2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port2"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	assert.NotNil(t, err)
}

func TestGetMetricsBgpv4Response(t *testing.T) {
	// Send Get Metrics request with device name d1 which is available in
	// the config, validate the response is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	req.Bgpv4()
	flow_req := req.Bgpv4()
	flow_req.SetDeviceNames([]string{"BGP-1"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsBgpv4ResponseError(t *testing.T) {
	// Send Get Metrics request with port name d2 which is not available in
	// the config, validate the err is not nil
	api := gosnappi.NewApi()
	api.NewGrpcTransport().SetLocation(mockServerLocation)
	req := api.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetDeviceNames([]string{"d2"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	assert.NotNil(t, err)
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

	assert.Equal(t, config, config_new, "Both configs shall be equal")
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
