package gosnappi_test

import (
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
)

func TestConfigIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewConfig()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLayer1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLayer1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCapture()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDevice()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlow()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestEventIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewEvent()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestConfigOptionsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewConfigOptions()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestTransmitStateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewTransmitState()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLinkStateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLinkState()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureStateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureState()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowsUpdateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowsUpdate()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestRouteStateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewRouteState()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPingRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPingRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPingIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPing()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestProtocolStateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewProtocolState()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestMetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPortMetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPortMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowMetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpv4MetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpv4MetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpv6MetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpv6MetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisMetricsRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestStatesRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewStatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestNeighborsv4StatesRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewNeighborsv4StatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestNeighborsv6StatesRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewNeighborsv6StatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetConfigResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestResponseWarningIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewResponseWarning()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestResponseErrorIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewResponseError()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestGetConfigResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewGetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetTransmitStateResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetTransmitStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetLinkStateResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetLinkStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetCaptureStateResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetCaptureStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestUpdateFlowsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewUpdateFlowsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetRouteStateResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetRouteStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSendPingResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSendPingResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPingResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPingResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestSetProtocolStateResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewSetProtocolStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestGetMetricsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewGetMetricsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestMetricsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewMetricsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestGetStatesResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewGetStatesResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestStatesResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewStatesResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestGetCaptureResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewGetCaptureResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLagPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLagPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLayer1AutoNegotiationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLayer1AutoNegotiation()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLayer1FlowControlIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLayer1FlowControl()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureFilterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureFilter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceEthernetIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIpv4LoopbackIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIpv4Loopback()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIpv6LoopbackIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIpv6Loopback()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIsisRouterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIsisRouter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceBgpRouterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceBgpRouter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowTxRxIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowTxRx()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowHeaderIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowHeader()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowSizeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowSize()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowRateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowRate()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowDurationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowDuration()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowMetricsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowMetrics()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestEventLinkIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewEventLink()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestEventRxRateThresholdIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewEventRxRateThreshold()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestEventRouteAdvertiseWithdrawIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewEventRouteAdvertiseWithdraw()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPortOptionsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPortOptions()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPingIpv4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPingIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPingIpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPingIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowMetricGroupRequestIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowMetricGroupRequest()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewResponse()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPortMetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPortMetric()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowMetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowMetric()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpv4MetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpv4Metric()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpv6MetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpv6Metric()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisMetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisMetric()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestNeighborsv4StateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewNeighborsv4State()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestNeighborsv6StateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewNeighborsv6State()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLagProtocolIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLagProtocol()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceEthernetBaseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceEthernetBase()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLayer1Ieee8021QbbIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLayer1Ieee8021Qbb()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLayer1Ieee8023XIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLayer1Ieee8023X()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureCustomIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureCustom()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureEthernetIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureVlanIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureVlan()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureIpv4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureIpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIpv4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceVlanIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceVlan()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestDeviceIsisMultiInstanceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewDeviceIsisMultiInstance()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisInterfaceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisInterface()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisBasicIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisBasic()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisAdvancedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisAuthenticationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisAuthentication()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisV4RouteRangeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisV4RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisV6RouteRangeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisV6RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV4InterfaceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV4Interface()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV6InterfaceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV6Interface()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowRouterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowRouter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowCustomIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowCustom()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowEthernetIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowVlanIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowVlan()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowVxlanIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowVxlan()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIpv4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowPfcPauseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowPfcPause()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowEthernetPauseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowEthernetPause()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowTcpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowTcp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowUdpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowUdp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowGreIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowGre()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowGtpv1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowGtpv1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowGtpv2IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowGtpv2()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowArpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowArp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIcmpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIcmp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIcmpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIcmpv6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowPppIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowPpp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIgmpv1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIgmpv1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowSizeIncrementIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowSizeIncrement()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowSizeRandomIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowSizeRandom()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowFixedPacketsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowFixedPackets()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowFixedSecondsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowFixedSeconds()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowBurstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowBurst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowContinuousIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowContinuous()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowLatencyMetricsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowLatencyMetrics()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowMetricGroupIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowMetricGroup()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestMetricTimestampIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewMetricTimestamp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestMetricLatencyIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewMetricLatency()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLagLacpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLagProtocolLacp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLagStaticIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLagProtocolStatic()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestCaptureFieldIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewCaptureField()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisInterfaceLevelIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisInterfaceLevel()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisMTIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisMT()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLinkStateTEIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLinkStateTE()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisInterfaceAuthenticationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisInterfaceAuthentication()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisInterfaceAdvancedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisInterfaceAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisInterfaceLinkProtectionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisInterfaceLinkProtection()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestIsisAuthenticationBaseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewIsisAuthenticationBase()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestV4RouteAddressIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewV4RouteAddress()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestV6RouteAddressIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewV6RouteAddress()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV4PeerIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV4Peer()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV6PeerIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV6Peer()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetDstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetDst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetSrcIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetSrc()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetEtherTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPfcQueueIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPfcQueue()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanPriorityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanPriority()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanCfiIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanCfi()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanIdIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanId()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanTpidIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanTpid()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanFlagsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanFlags()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanReserved0IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanReserved0()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanVniIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanVni()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanReserved1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanReserved1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4VersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Version()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4HeaderLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4HeaderLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIpv4PriorityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIpv4Priority()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TotalLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TotalLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4IdentificationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Identification()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4ReservedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Reserved()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DontFragmentIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DontFragment()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4MoreFragmentsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4MoreFragments()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4FragmentOffsetIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4FragmentOffset()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TimeToLiveIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TimeToLive()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4ProtocolIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Protocol()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4HeaderChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4HeaderChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4SrcIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Src()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4Dst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6VersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6Version()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6TrafficClassIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6TrafficClass()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6FlowLabelIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6FlowLabel()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6PayloadLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6PayloadLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6NextHeaderIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6NextHeader()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6HopLimitIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6HopLimit()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6SrcIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6Src()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6DstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6Dst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseDstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseDst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseSrcIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseSrc()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseEtherTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseControlOpCodeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseControlOpCode()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseClassEnableVectorIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseClassEnableVector()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass0IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass0()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass2IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass2()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass3IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass3()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass4()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass5IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass5()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass6()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass7IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass7()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseDstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseDst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseSrcIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseSrc()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseEtherTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseControlOpCodeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseControlOpCode()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseTimeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseTime()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpSrcPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpSrcPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpDstPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpDstPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpSeqNumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpSeqNum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpAckNumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpAckNum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpDataOffsetIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpDataOffset()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnNsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnNs()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnCwrIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnCwr()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnEchoIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnEcho()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlUrgIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlUrg()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlAckIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlAck()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlPshIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlPsh()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlRstIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlRst()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlSynIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlSyn()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlFinIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlFin()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpWindowIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpWindow()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpSrcPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpSrcPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpDstPortIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpDstPort()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreChecksumPresentIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreChecksumPresent()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreReserved0IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreReserved0()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreVersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreVersion()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreProtocolIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreProtocol()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreReserved1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreReserved1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1VersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1Version()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1ProtocolTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1ProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1ReservedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1Reserved()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1EFlagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1EFlag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1SFlagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1SFlag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1PnFlagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1PnFlag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1MessageTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1MessageType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1MessageLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1MessageLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1TeidIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1Teid()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1SquenceNumberIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1SquenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1NPduNumberIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1NPduNumber()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1NextExtensionHeaderTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1NextExtensionHeaderType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowGtpExtensionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowGtpExtension()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2VersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Version()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2PiggybackingFlagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2PiggybackingFlag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2TeidFlagIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2TeidFlag()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2Spare1IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Spare1()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2MessageTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2MessageType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2MessageLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2MessageLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2TeidIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Teid()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2SequenceNumberIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2SequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2Spare2IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Spare2()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpHardwareTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpHardwareType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpProtocolTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpHardwareLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpHardwareLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpProtocolLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpProtocolLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpOperationIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpOperation()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpSenderHardwareAddrIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpSenderHardwareAddr()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpSenderProtocolAddrIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpSenderProtocolAddr()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpTargetHardwareAddrIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpTargetHardwareAddr()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpTargetProtocolAddrIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpTargetProtocolAddr()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIcmpEchoIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIcmpEcho()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIcmpv6EchoIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIcmpv6Echo()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppAddressIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppAddress()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppControlIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppControl()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppProtocolTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1VersionIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1Version()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1TypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1Type()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1UnusedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1Unused()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1ChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1Checksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1GroupAddressIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1GroupAddress()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowDelayIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowDelay()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowDurationInterBurstGapIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowDurationInterBurstGap()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestLinkStatepriorityBandwidthsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewLinkStatepriorityBandwidths()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpAdvancedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpCapabilityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpCapability()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV4RouteRangeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV4RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV6RouteRangeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV6RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteV4PolicyIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteV4Policy()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteV6PolicyIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteV6Policy()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpV6SegmentRoutingIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpV6SegmentRouting()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetDstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetSrcCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetEtherTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPfcQueueCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPfcQueueCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanPriorityCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanPriorityCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanCfiCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanCfiCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanIdCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanIdCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVlanTpidCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVlanTpidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanFlagsCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanFlagsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanReserved0CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanReserved0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanVniCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanVniCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowVxlanReserved1CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowVxlanReserved1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4VersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4HeaderLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4HeaderLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4PriorityRawIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4PriorityRaw()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIpv4TosIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIpv4Tos()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestFlowIpv4DscpIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewFlowIpv4Dscp()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TotalLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TotalLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4IdentificationCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4IdentificationCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4ReservedCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4ReservedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DontFragmentCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DontFragmentCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4MoreFragmentsCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4MoreFragmentsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4FragmentOffsetCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4FragmentOffsetCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TimeToLiveCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TimeToLiveCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4ProtocolCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4ProtocolCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4SrcCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4SrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6VersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6TrafficClassCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6TrafficClassCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6FlowLabelCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6FlowLabelCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6PayloadLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6PayloadLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6NextHeaderCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6NextHeaderCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6HopLimitCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6HopLimitCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6SrcCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6SrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv6DstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv6DstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseDstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseSrcCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseEtherTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseControlOpCodeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseControlOpCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPauseClassEnableVectorCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPauseClassEnableVectorCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass0CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass1CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass2CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass2Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass3CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass3Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass4CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass4Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass5CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass5Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass6CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass6Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPfcPausePauseClass7CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass7Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseDstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseSrcCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseEtherTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseControlOpCodeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseControlOpCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowEthernetPauseTimeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowEthernetPauseTimeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpSrcPortCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpSrcPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpDstPortCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpDstPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpSeqNumCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpSeqNumCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpAckNumCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpAckNumCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpDataOffsetCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpDataOffsetCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnNsCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnNsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnCwrCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnCwrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpEcnEchoCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpEcnEchoCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlUrgCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlUrgCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlAckCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlAckCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlPshCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlPshCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlRstCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlRstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlSynCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlSynCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpCtlFinCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpCtlFinCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowTcpWindowCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowTcpWindowCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpSrcPortCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpSrcPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpDstPortCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpDstPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowUdpLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowUdpLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreChecksumPresentCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreChecksumPresentCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreReserved0CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreReserved0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreVersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreVersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreProtocolCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreProtocolCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGreReserved1CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGreReserved1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1VersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1ProtocolTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1ProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1ReservedCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1ReservedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1EFlagCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1EFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1SFlagCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1SFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1PnFlagCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1PnFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1MessageTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1MessageTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1MessageLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1MessageLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1TeidCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1TeidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1SquenceNumberCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1SquenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1NPduNumberCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1NPduNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv1NextExtensionHeaderTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv1NextExtensionHeaderTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionExtensionLengthIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionExtensionLength()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionContentsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionContents()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionNextExtensionHeaderIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionNextExtensionHeader()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2VersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2PiggybackingFlagCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2PiggybackingFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2TeidFlagCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2TeidFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2Spare1CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Spare1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2MessageTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2MessageTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2MessageLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2MessageLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2TeidCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2TeidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2SequenceNumberCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2SequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpv2Spare2CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpv2Spare2Counter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpHardwareTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpHardwareTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpProtocolTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpHardwareLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpHardwareLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpProtocolLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpProtocolLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpOperationCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpOperationCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpSenderHardwareAddrCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpSenderHardwareAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpSenderProtocolAddrCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpSenderProtocolAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpTargetHardwareAddrCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpTargetHardwareAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowArpTargetProtocolAddrCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowArpTargetProtocolAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoCodeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoCode()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoIdentifierIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoIdentifier()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoSequenceNumberIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoSequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoTypeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoType()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoCodeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoCode()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoIdentifierIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoIdentifier()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoSequenceNumberIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoSequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppAddressCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppAddressCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppControlCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppControlCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowPppProtocolTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowPppProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1VersionCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1TypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1TypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1UnusedCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1UnusedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIgmpv1GroupAddressCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIgmpv1GroupAddressCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpRouteAdvancedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpRouteAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpCommunityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpCommunity()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpAsPathIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpAsPath()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpAddPathIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpAddPath()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpExtCommunityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpExtCommunity()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteV4TunnelTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteV4TunnelTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteV6TunnelTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteV6TunnelTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4PriorityRawCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4PriorityRawCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosPrecedenceIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosPrecedence()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosDelayIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosDelay()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosThroughputIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosThroughput()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosReliabilityIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosReliability()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosMonetaryIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosMonetary()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosUnusedIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosUnused()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DscpPhbIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DscpPhb()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DscpEcnIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DscpEcn()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionExtensionLengthCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionExtensionLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionContentsCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionContentsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowGtpExtensionNextExtensionHeaderCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowGtpExtensionNextExtensionHeaderCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoCodeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoIdentifierCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoIdentifierCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpEchoSequenceNumberCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpEchoSequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoTypeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoCodeCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoIdentifierCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoIdentifierCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIcmpv6EchoSequenceNumberCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIcmpv6EchoSequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpAsPathSegmentIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpAsPathSegment()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteRemoteEndpointSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteRemoteEndpointSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteColorSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteColorSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteBindingSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteBindingSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrtePreferenceSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrtePreferenceSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrtePolicyPrioritySubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrtePolicyPrioritySubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrtePolicyNameSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrtePolicyNameSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteExplicitNullLabelPolicySubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteExplicitNullLabelPolicySubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentListIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentList()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosPrecedenceCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosPrecedenceCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosDelayCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosDelayCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosThroughputCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosThroughputCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosReliabilityCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosReliabilityCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosMonetaryCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosMonetaryCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4TosUnusedCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4TosUnusedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DscpPhbCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DscpPhbCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestPatternFlowIpv4DscpEcnCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewPatternFlowIpv4DscpEcnCounter()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegment()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentATypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentATypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentBTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentBTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentCTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentCTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentDTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentDTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentETypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentETypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentFTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentFTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentGTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentGTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentHTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentHTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentITypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentITypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentJTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentJTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSegmentKTypeSubTlvIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSegmentKTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSRv6SIDEndpointBehaviorAndStructureIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSRv6SIDEndpointBehaviorAndStructure()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
func TestBgpSrteSrMplsSidIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := gosnappi.NewBgpSrteSrMplsSid()
	assert.NotNil(t, object.FromYaml(incorrect_format))
	assert.NotNil(t, object.FromJson(incorrect_format))
	assert.NotNil(t, object.FromPbText(incorrect_format))
}
