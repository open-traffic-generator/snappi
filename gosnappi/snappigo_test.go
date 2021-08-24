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
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}

func TestGetMetricsResponse(t *testing.T) {
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f2"})
	resp, _ := api.GetMetrics(req)
	assert.NotNil(t, resp)
}

func TestGetMetricsError(t *testing.T) {
	api := snappigo.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testport))
	req := api.NewMetricsRequest()
	flow_req := req.Flow()
	flow_req.SetFlowNames([]string{"f3"})
	_, err := api.GetMetrics(req)
	log.Print(err)
	assert.NotNil(t, err)
}
