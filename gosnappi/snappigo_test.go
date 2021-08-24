package gosnappi_test

import (
	"fmt"
	log "log"
	"testing"

	snappigo "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := StartMockServer(); err != nil {
		log.Fatal("Mock server init failed")
	}
}

// Basic test for the grpc mock server
func TestApi(t *testing.T) {
	api := snappigo.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetLocation("127.0.0.1:50001").SetRequestTimeout(10000)
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
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f2"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsFlowResponseError(t *testing.T) {
	// Send Get Metrics request with flow name f3 which is not available in
	// the config, validate the err is not nil
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
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
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Port()
	flow_req.SetPortNames([]string{"port1"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsPortResponseError(t *testing.T) {
	// Send Get Metrics request with port name port2 which is not available in
	// the config, validate the err is not nil
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
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
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetDeviceNames([]string{"d1"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsBgpv4ResponseError(t *testing.T) {
	// Send Get Metrics request with port name d2 which is not available in
	// the config, validate the err is not nil
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Bgpv4()
	flow_req.SetDeviceNames([]string{"d2"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	assert.NotNil(t, err)
}
