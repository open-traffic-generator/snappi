package gosnappi_test

import (
	"testing"

	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
)

func TestConfigIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewConfig()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLayer1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCapture()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDevice()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlow()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestEventIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewEvent()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestConfigOptionsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewConfigOptions()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestTransmitStateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewTransmitState()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLinkStateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLinkState()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureStateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureState()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowsUpdateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowsUpdate()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestRouteStateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewRouteState()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPingRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPingRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPingIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPing()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestProtocolStateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewProtocolState()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPortMetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPortMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowMetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpv4MetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpv4MetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpv6MetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpv6MetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisMetricsRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisMetricsRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestStatesRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewStatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestNeighborsv4StatesRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewNeighborsv4StatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestNeighborsv6StatesRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewNeighborsv6StatesRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureRequest()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetConfigResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestResponseWarningIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewResponseWarning()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestResponseErrorIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewResponseError()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetConfigResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewGetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetTransmitStateResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetTransmitStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetLinkStateResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetLinkStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetCaptureStateResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetCaptureStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestUpdateFlowsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewUpdateFlowsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetRouteStateResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetRouteStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSendPingResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSendPingResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPingResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPingResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetProtocolStateResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewSetProtocolStateResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetMetricsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewGetMetricsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMetricsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewMetricsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetStatesResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewGetStatesResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestStatesResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewStatesResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetCaptureResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewGetCaptureResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLagPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLagPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1AutoNegotiationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLayer1AutoNegotiation()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1FlowControlIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLayer1FlowControl()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureFilterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureFilter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceEthernetIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIpv4LoopbackIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIpv4Loopback()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIpv6LoopbackIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIpv6Loopback()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIsisRouterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIsisRouter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceBgpRouterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceBgpRouter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowTxRxIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowTxRx()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowHeaderIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowHeader()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowSizeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowSize()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowRateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowRate()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowDurationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowDuration()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowMetricsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowMetrics()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestEventLinkIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewEventLink()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestEventRxRateThresholdIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewEventRxRateThreshold()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestEventRouteAdvertiseWithdrawIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewEventRouteAdvertiseWithdraw()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPortOptionsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPortOptions()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPingIpv4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPingIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPingIpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPingIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowMetricGroupRequestIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowMetricTagFilter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPortMetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPortMetric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowMetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowMetric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpv4MetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpv4Metric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpv6MetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpv6Metric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisMetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisMetric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestNeighborsv4StateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewNeighborsv4State()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestNeighborsv6StateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewNeighborsv6State()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLagProtocolIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLagProtocol()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceEthernetBaseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceEthernetBase()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1Ieee8021QbbIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLayer1Ieee8021Qbb()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1Ieee8023XIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLayer1Ieee8023X()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureCustomIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureCustom()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureEthernetIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureVlanIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureVlan()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureIpv4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureIpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIpv4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceVlanIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceVlan()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestDeviceIsisMultiInstanceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewDeviceIsisMultiInstance()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisInterfaceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisInterface()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisBasicIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisBasic()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisAdvancedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisAuthenticationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisAuthentication()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisV4RouteRangeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisV4RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisV6RouteRangeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisV6RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV4InterfaceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV4Interface()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV6InterfaceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV6Interface()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowRouterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowRouter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowCustomIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowCustom()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowEthernetIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowEthernet()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowVlanIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowVlan()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowVxlanIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowVxlan()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIpv4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowPfcPauseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowPfcPause()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowEthernetPauseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowEthernetPause()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowTcpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowTcp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowUdpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowUdp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowGreIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowGre()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowGtpv1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowGtpv1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowGtpv2IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowGtpv2()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowArpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowArp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIcmpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIcmp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIcmpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIcmpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowPppIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowPpp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIgmpv1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIgmpv1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowSizeIncrementIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowSizeIncrement()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowSizeRandomIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowSizeRandom()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowFixedPacketsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowFixedPackets()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowFixedSecondsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowFixedSeconds()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowBurstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowBurst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowContinuousIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowContinuous()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowLatencyMetricsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowLatencyMetrics()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowMetricGroupIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowMetricGroup()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMetricTimestampIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewMetricTimestamp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMetricLatencyIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewMetricLatency()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLagLacpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLagProtocolLacp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLagStaticIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLagProtocolStatic()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestCaptureFieldIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewCaptureField()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisInterfaceLevelIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisInterfaceLevel()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisMTIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisMT()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLinkStateTEIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLinkStateTE()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisInterfaceAuthenticationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisInterfaceAuthentication()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisInterfaceAdvancedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisInterfaceAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisInterfaceLinkProtectionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisInterfaceLinkProtection()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIsisAuthenticationBaseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewIsisAuthenticationBase()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestV4RouteAddressIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewV4RouteAddress()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestV6RouteAddressIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewV6RouteAddress()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV4PeerIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV4Peer()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV6PeerIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV6Peer()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetDstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetDst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetSrcIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetSrc()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetEtherTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPfcQueueIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPfcQueue()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanPriorityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanPriority()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanCfiIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanCfi()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanIdIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanId()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanTpidIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanTpid()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanFlagsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanFlags()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanReserved0IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanReserved0()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanVniIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanVni()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanReserved1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanReserved1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4VersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Version()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4HeaderLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4HeaderLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIpv4PriorityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIpv4Priority()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TotalLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TotalLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4IdentificationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Identification()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4ReservedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Reserved()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DontFragmentIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DontFragment()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4MoreFragmentsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4MoreFragments()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4FragmentOffsetIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4FragmentOffset()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TimeToLiveIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TimeToLive()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4ProtocolIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Protocol()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4HeaderChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4HeaderChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4SrcIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Src()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4Dst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6VersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6Version()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6TrafficClassIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6TrafficClass()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6FlowLabelIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6FlowLabel()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6PayloadLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6PayloadLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6NextHeaderIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6NextHeader()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6HopLimitIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6HopLimit()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6SrcIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6Src()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6DstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6Dst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseDstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseDst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseSrcIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseSrc()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseEtherTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseControlOpCodeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseControlOpCode()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseClassEnableVectorIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseClassEnableVector()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass0IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass0()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass2IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass2()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass3IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass3()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass5IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass5()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass7IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass7()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseDstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseDst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseSrcIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseSrc()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseEtherTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseEtherType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseControlOpCodeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseControlOpCode()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseTimeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseTime()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpSrcPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpSrcPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpDstPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpDstPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpSeqNumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpSeqNum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpAckNumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpAckNum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpDataOffsetIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpDataOffset()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnNsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnNs()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnCwrIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnCwr()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnEchoIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnEcho()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlUrgIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlUrg()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlAckIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlAck()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlPshIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlPsh()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlRstIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlRst()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlSynIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlSyn()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlFinIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlFin()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpWindowIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpWindow()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpSrcPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpSrcPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpDstPortIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpDstPort()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreChecksumPresentIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreChecksumPresent()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreReserved0IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreReserved0()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreVersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreVersion()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreProtocolIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreProtocol()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreReserved1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreReserved1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1VersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1Version()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1ProtocolTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1ProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1ReservedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1Reserved()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1EFlagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1EFlag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1SFlagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1SFlag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1PnFlagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1PnFlag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1MessageTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1MessageType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1MessageLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1MessageLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1TeidIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1Teid()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1SquenceNumberIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1SquenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1NPduNumberIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1NPduNumber()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1NextExtensionHeaderTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1NextExtensionHeaderType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowGtpExtensionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowGtpExtension()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2VersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Version()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2PiggybackingFlagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2PiggybackingFlag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2TeidFlagIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2TeidFlag()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2Spare1IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Spare1()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2MessageTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2MessageType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2MessageLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2MessageLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2TeidIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Teid()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2SequenceNumberIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2SequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2Spare2IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Spare2()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpHardwareTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpHardwareType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpProtocolTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpHardwareLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpHardwareLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpProtocolLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpProtocolLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpOperationIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpOperation()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpSenderHardwareAddrIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpSenderHardwareAddr()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpSenderProtocolAddrIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpSenderProtocolAddr()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpTargetHardwareAddrIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpTargetHardwareAddr()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpTargetProtocolAddrIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpTargetProtocolAddr()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIcmpEchoIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIcmpEcho()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIcmpv6EchoIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIcmpv6Echo()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppAddressIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppAddress()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppControlIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppControl()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppProtocolTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppProtocolType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1VersionIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1Version()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1TypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1Type()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1UnusedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1Unused()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1ChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1Checksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1GroupAddressIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1GroupAddress()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowDelayIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowDelay()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowDurationInterBurstGapIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowDurationInterBurstGap()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLinkStatepriorityBandwidthsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewLinkStatepriorityBandwidths()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpAdvancedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpCapabilityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpCapability()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV4RouteRangeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV4RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV6RouteRangeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV6RouteRange()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteV4PolicyIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteV4Policy()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteV6PolicyIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteV6Policy()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpV6SegmentRoutingIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpV6SegmentRouting()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetDstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetSrcCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetEtherTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPfcQueueCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPfcQueueCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanPriorityCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanPriorityCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanCfiCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanCfiCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanIdCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanIdCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVlanTpidCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVlanTpidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanFlagsCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanFlagsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanReserved0CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanReserved0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanVniCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanVniCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowVxlanReserved1CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowVxlanReserved1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4VersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4HeaderLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4HeaderLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4PriorityRawIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4PriorityRaw()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIpv4TosIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIpv4Tos()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFlowIpv4DscpIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewFlowIpv4Dscp()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TotalLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TotalLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4IdentificationCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4IdentificationCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4ReservedCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4ReservedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DontFragmentCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DontFragmentCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4MoreFragmentsCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4MoreFragmentsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4FragmentOffsetCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4FragmentOffsetCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TimeToLiveCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TimeToLiveCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4ProtocolCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4ProtocolCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4SrcCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4SrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6VersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6TrafficClassCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6TrafficClassCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6FlowLabelCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6FlowLabelCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6PayloadLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6PayloadLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6NextHeaderCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6NextHeaderCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6HopLimitCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6HopLimitCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6SrcCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6SrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv6DstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv6DstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseDstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseSrcCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseEtherTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseControlOpCodeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseControlOpCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPauseClassEnableVectorCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPauseClassEnableVectorCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass0CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass1CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass2CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass2Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass3CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass3Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass4CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass4Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass5CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass5Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass6CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass6Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPfcPausePauseClass7CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPfcPausePauseClass7Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseDstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseDstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseSrcCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseSrcCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseEtherTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseEtherTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseControlOpCodeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseControlOpCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowEthernetPauseTimeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowEthernetPauseTimeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpSrcPortCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpSrcPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpDstPortCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpDstPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpSeqNumCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpSeqNumCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpAckNumCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpAckNumCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpDataOffsetCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpDataOffsetCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnNsCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnNsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnCwrCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnCwrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpEcnEchoCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpEcnEchoCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlUrgCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlUrgCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlAckCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlAckCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlPshCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlPshCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlRstCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlRstCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlSynCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlSynCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpCtlFinCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpCtlFinCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowTcpWindowCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowTcpWindowCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpSrcPortCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpSrcPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpDstPortCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpDstPortCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowUdpLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowUdpLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreChecksumPresentCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreChecksumPresentCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreReserved0CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreReserved0Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreVersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreVersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreProtocolCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreProtocolCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGreReserved1CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGreReserved1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1VersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1ProtocolTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1ProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1ReservedCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1ReservedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1EFlagCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1EFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1SFlagCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1SFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1PnFlagCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1PnFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1MessageTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1MessageTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1MessageLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1MessageLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1TeidCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1TeidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1SquenceNumberCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1SquenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1NPduNumberCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1NPduNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv1NextExtensionHeaderTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv1NextExtensionHeaderTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionExtensionLengthIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionExtensionLength()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionContentsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionContents()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionNextExtensionHeaderIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionNextExtensionHeader()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2VersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2PiggybackingFlagCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2PiggybackingFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2TeidFlagCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2TeidFlagCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2Spare1CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Spare1Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2MessageTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2MessageTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2MessageLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2MessageLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2TeidCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2TeidCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2SequenceNumberCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2SequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpv2Spare2CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpv2Spare2Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpHardwareTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpHardwareTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpProtocolTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpHardwareLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpHardwareLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpProtocolLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpProtocolLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpOperationCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpOperationCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpSenderHardwareAddrCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpSenderHardwareAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpSenderProtocolAddrCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpSenderProtocolAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpTargetHardwareAddrCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpTargetHardwareAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowArpTargetProtocolAddrCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowArpTargetProtocolAddrCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoCodeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoCode()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoIdentifierIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoIdentifier()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoSequenceNumberIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoSequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoTypeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoType()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoCodeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoCode()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoIdentifierIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoIdentifier()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoSequenceNumberIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoSequenceNumber()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppAddressCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppAddressCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppControlCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppControlCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowPppProtocolTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowPppProtocolTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1VersionCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1VersionCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1TypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1TypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1UnusedCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1UnusedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIgmpv1GroupAddressCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIgmpv1GroupAddressCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpRouteAdvancedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpRouteAdvanced()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpCommunityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpCommunity()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpAsPathIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpAsPath()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpAddPathIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpAddPath()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpExtCommunityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpExtCommunity()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteV4TunnelTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteV4TunnelTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteV6TunnelTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteV6TunnelTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4PriorityRawCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4PriorityRawCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosPrecedenceIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosPrecedence()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosDelayIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosDelay()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosThroughputIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosThroughput()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosReliabilityIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosReliability()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosMonetaryIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosMonetary()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosUnusedIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosUnused()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DscpPhbIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DscpPhb()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DscpEcnIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DscpEcn()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionExtensionLengthCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionExtensionLengthCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionContentsCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionContentsCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowGtpExtensionNextExtensionHeaderCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowGtpExtensionNextExtensionHeaderCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoCodeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoIdentifierCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoIdentifierCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpEchoSequenceNumberCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpEchoSequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoTypeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoTypeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoCodeCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoCodeCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoIdentifierCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoIdentifierCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIcmpv6EchoSequenceNumberCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIcmpv6EchoSequenceNumberCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpAsPathSegmentIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpAsPathSegment()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteRemoteEndpointSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteRemoteEndpointSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteColorSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteColorSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteBindingSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteBindingSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrtePreferenceSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrtePreferenceSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrtePolicyPrioritySubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrtePolicyPrioritySubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrtePolicyNameSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrtePolicyNameSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteExplicitNullLabelPolicySubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteExplicitNullLabelPolicySubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentListIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentList()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosPrecedenceCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosPrecedenceCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosDelayCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosDelayCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosThroughputCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosThroughputCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosReliabilityCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosReliabilityCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosMonetaryCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosMonetaryCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4TosUnusedCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4TosUnusedCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DscpPhbCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DscpPhbCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternFlowIpv4DscpEcnCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewPatternFlowIpv4DscpEcnCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegment()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentATypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentATypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentBTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentBTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentCTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentCTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentDTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentDTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentETypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentETypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentFTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentFTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentGTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentGTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentHTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentHTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentITypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentITypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentJTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentJTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSegmentKTypeSubTlvIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSegmentKTypeSubTlv()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSRv6SIDEndpointBehaviorAndStructureIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSRv6SIDEndpointBehaviorAndStructure()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestBgpSrteSrMplsSidIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := gosnappi.NewBgpSrteSrMplsSid()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
