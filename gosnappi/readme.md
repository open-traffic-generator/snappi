# ![snappi](../snappi-logo.png)

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Build](https://github.com/open-traffic-generator/snappi/workflows/Build/badge.svg)](https://github.com/open-traffic-generator/snappi/actions)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/alerts/)

Test scripts written in `gosnappi`, an auto-generated Go SDK, can be executed against any traffic generator conforming to [Open Traffic Generator API](https://github.com/open-traffic-generator/models).

[Ixia-c](https://github.com/open-traffic-generator/ixia-c) is one such reference implementation of Open Traffic Generator API.

> The repository is under active development and is subject to updates. All efforts will be made to keep the updates backwards compatible.

## Setup Client 

```sh
go get github.com/open-traffic-generator/snappi/gosnappi
```

## Start testing

```go
package examples

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/open-traffic-generator/snappi/gosnappi"
)

func TestQuickstart(t *testing.T) {
	// Create a new API handle to make API calls against OTG
	api := gosnappi.NewApi()

	// Set the transport protocol to HTTP
	api.NewHttpTransport().SetLocation("https://localhost:8443")

	// Create a new traffic configuration that will be set on OTG
	config := api.NewConfig()

	// Add a test port to the configuration
	ptx := config.Ports().Add().SetName("ptx").SetLocation("veth-a")

	// Configure a flow and set previously created test port as one of endpoints
	flow := config.Flows().Add().SetName("f1")
	flow.TxRx().Port().SetTxName(ptx.Name())
	// and enable tracking flow metrics
	flow.Metrics().SetEnable(true)

	// Configure number of packets to transmit for previously configured flow
	flow.Duration().FixedPackets().SetPackets(100)
	// and fixed byte size of all packets in the flow
	flow.Size().SetFixed(128)

	// Configure protocol headers for all packets in the flow
	pkt := flow.Packet()
	eth := pkt.Add().Ethernet()
	ipv4 := pkt.Add().Ipv4()
	udp := pkt.Add().Udp()
	cus := pkt.Add().Custom()

	eth.Dst().SetValue("00:11:22:33:44:55")
	eth.Src().SetValue("00:11:22:33:44:66")

	ipv4.Src().SetValue("10.1.1.1")
	ipv4.Dst().SetValue("20.1.1.1")

	// Configure repeating patterns for source and destination UDP ports
	udp.SrcPort().SetValues([]int32{5010, 5015, 5020, 5025, 5030})
	udp.DstPort().Increment().SetStart(6010).SetStep(5).SetCount(5)

	// Configure custom bytes (hex string) in payload
	cus.SetBytes(hex.EncodeToString([]byte("..QUICKSTART SNAPPI..")))

	// Optionally, print JSON representation of config
	if j, err := config.ToJson(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("Configuration: ", j)
	}

	// Push traffic configuration constructed so far to OTG
	if _, err := api.SetConfig(config); err != nil {
		t.Fatal(err)
	}

	// Start transmitting the packets from configured flow
	ts := api.NewTransmitState()
	ts.SetState(gosnappi.TransmitStateState.START)
	if _, err := api.SetTransmitState(ts); err != nil {
		t.Fatal(err)
	}

	// Fetch metrics for configured flow
	req := api.NewMetricsRequest()
	req.Flow().SetFlowNames([]string{flow.Name()})
	// and keep polling until either expectation is met or deadline exceeds
	deadline := time.Now().Add(10 * time.Second)
	for {
		metrics, err := api.GetMetrics(req)
		if err != nil || time.Now().After(deadline) {
			t.Fatalf("err = %v || deadline exceeded", err)
		}
		// print YAML representation of flow metrics
		t.Log(metrics)
		if metrics.FlowMetrics().Items()[0].Transmit() == gosnappi.FlowMetricTransmit.STOPPED {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
```
