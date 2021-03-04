import pytest
import snappi


def test_snappi_lists(api):
    """Validate SnappiList object indexing and unpacking
    """
    config = api.config()

    flows = config.flows.flow(name='1')
    assert(flows.__class__ == snappi.FlowIter)
    flow = flows[0]
    assert(flow.__class__ == snappi.Flow)
    eth, vlan, vlan1 = flow.packet.ethernet().vlan().vlan()
    vlan.id.value = 1

    flow = config.flows.flow(name='2')[-1]
    assert(flow.__class__ == snappi.Flow)
    eth, vlan, ipv4, tcp = flow.packet.ethernet().vlan().ipv4().tcp()
    assert(eth.__class__ == snappi.FlowEthernet)
    assert(vlan.__class__ == snappi.FlowVlan)
    assert(ipv4.__class__ == snappi.FlowIpv4)
    assert(tcp.__class__ == snappi.FlowTcp)
    vlan.id.value = 2

    flow = config.flows.flow(name='3')[-1]
    pkt = flow.packet.ethernet().vlan()
    assert(pkt.__class__ == snappi.FlowHeaderIter)
    vlan = pkt[-1]
    vlan.id.value = 3

    flow = config.flows.flow(name='4')[-1]
    vlan = flow.packet.ethernet().vlan()[-1]
    assert(vlan.__class__ == snappi.FlowVlan)
    vlan.id.value = 4

    print(config)
    api.set_config(config)
    
    assert (len(config.flows) == 4)
    assert (config.flows[0].name == '1')
    assert (config.flows[1].name == '2')
    assert (config.flows[2].name == '3')
    assert (config.flows[3].name == '4')
    assert (len(config.flows[0].packet) == 3)
    assert (len(config.flows[1].packet) == 4)
    assert (len(config.flows[2].packet) == 2)
    assert (len(config.flows[3].packet) == 2)
    assert (config.flows[0].packet[1].id.value == 1)
    assert (config.flows[1].packet[1].id.value == 2)
    assert (config.flows[2].packet[1].id.value == 3)
    assert (config.flows[3].packet[1].id.value == 4)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
