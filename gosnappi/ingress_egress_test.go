package gosnappi_test

import (
	"fmt"
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/stretchr/testify/assert"
)

func TestIngressEgress(t *testing.T) {
	api := gosnappi.NewApi()
	config := api.NewConfig()
	port1 := config.Ports().Add()
	port2 := config.Ports().Add()
	port1.SetName("port1")
	port1.SetLocation("location1")
	port2.SetName("port2")
	port2.SetLocation("location2")

	f1 := config.Flows().Add().SetName("f1")
	f1.TxRx().Port().
		SetTxName(port1.Name()).
		SetRxName(port2.Name())
	f1.Duration().FixedPackets().SetPackets(100)
	f1.Metrics().SetEnable(true)

	Ing := f1.Packet()
	ethIng := Ing.Add().Ethernet()
	ethIng.Src().SetValue("00:00:01:01:01:01")

	ip4Ing := f1.Packet().Add().Ipv4()
	ip4Ing.Src().SetValue("1.1.1.1")

	ethEgr := f1.EgressPacket().Add().Ethernet()
	ethEgr.Src().SetValue("00:00:01:01:01:44")

	ip4Egr := f1.EgressPacket().Add().Ipv4()
	ip4Egr.Src().SetValue("1.1.2.2")

	assert.Equal(t, len(f1.EgressPacket().Items()), 2)
	assert.Equal(t, len(f1.Packet().Items()), 2)

	fmt.Println(config)
}
