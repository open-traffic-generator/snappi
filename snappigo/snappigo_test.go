package snappigo_test

import (
	"testing"

	snappigo "github.com/open-traffic-generator/snappi/snappigo"
	assert "github.com/stretchr/testify/assert"
)

// Basic test for the grpc mock server
func TestApi(t *testing.T) {
	api := snappigo.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetLocation("127.0.0.1:50001").SetRequestTimeout(10000)
	config := api.NewConfig()
	port := config.Ports().Add()
	port.SetName("port1")
	port.SetLocation("location1")
	state, err := api.SetConfig(config)
	assert.NotNil(t, state)
	assert.Nil(t, err)
	status, _error := api.GetConfig()
	assert.NotNil(t, status)
	assert.Nil(t, _error)
}
