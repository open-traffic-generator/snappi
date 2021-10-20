# ![snappi](../snappi-logo.png)

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Build](https://github.com/open-traffic-generator/snappi/workflows/Build/badge.svg)](https://github.com/open-traffic-generator/snappi/actions)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/alerts/)

Test scripts written in `gosnappi`, an auto-generated Go SDK, can be executed against any traffic generator conforming to [Open Traffic Generator API](https://github.com/open-traffic-generator/models).

[Ixia-c](https://github.com/open-traffic-generator/ixia-c) is one such reference implementation of Open Traffic Generator API.

> The repository is under active development and is subject to updates. All efforts will be made to keep the updates backwards compatible.

## Install on a client

```sh
go get github.com/open-traffic-generator/snappi/gosnappi
```

## Start scripting

```go
package main

import (
	"fmt"

	"github.com/open-traffic-generator/snappi/gosnappi"
)

func main() {
	// Create a new API object, A hook to create and apply the configuration on a traffic generator
	api := gosnappi.NewApi()

	// Set the communication medium to either GRPC or HTTP
	grpc := api.NewGrpcTransport()
	grpc.SetLocation("<grpc server ip>:<port>")

	// Create a new configuration object that will be applied on traffic generator.
	config := api.NewConfig()

	// Add port locations to the configuration
	p1 := config.Ports().Add().SetName("Port1").SetLocation("localhost:5555")
	p2 := config.Ports().Add().SetName("Port2").SetLocation("localhost:5556")

	// Configure the flow and set the endpoints
	flow := config.Flows().Add().SetName("Flow1")
	flow.TxRx().Port().SetTxName(p1.Name()).SetRxName(p2.Name())

	// Configure the size of a packet and the number of packets to transmit
	flow.Size().SetFixed(128)
	flow.Duration().FixedPackets().SetPackets(1000)

	// Configure the header stack
	eth := flow.Packet().Add().Ethernet()
	ipv4 := flow.Packet().Add().Ipv4()
	tcp := flow.Packet().Add().Tcp()

	eth.Dst().SetValue("00:11:22:33:44:55")
	eth.Src().SetValue("00:11:22:33:44:66")

	ipv4.Src().SetValue("10.1.1.1")
	ipv4.Dst().SetValue("20.1.1.1")

	tcp.SrcPort().SetValue(5000)
	tcp.DstPort().SetValue(6000)

	// The configuration object is set with ports and flows
	// Now it shall be pushed to the traffic generator.
	api.SetConfig(config)

	// Start transmitting the configured flows
	ts := api.NewTransmitState()
	ts.SetState(gosnappi.TransmitStateState.START)
	api.SetTransmitState(ts)

	// Fetch and the port metrics
	req := api.NewMetricsRequest()
	req.Port().SetPortNames([]string{p1.Name(), p2.Name()})
	fmt.Println(api.GetMetrics(req))
}
```
