import pytest
import snappi


def test_packet_factory_methods(api):
    """Test flow.packet factory methods

    Packet factory methods should preserve the order and 
    choice should be preserved.
    """
    config = api.config()

    flow = config.flows.flow(name='f')[-1]
    eth, vlan1, vlan2, ip, _ = flow.packet.ethernet().vlan().vlan().ipv4().tcp()
    tcp = flow.packet[-1]
    assert(len(flow.packet) == 5)
    assert(eth.__class__ == snappi.FlowEthernet)
    assert(flow.packet[0].parent.choice == 'ethernet')
    assert(vlan1.__class__ == snappi.FlowVlan)
    assert(flow.packet[1].parent.choice == 'vlan')
    assert(vlan2.__class__ == snappi.FlowVlan)
    assert(flow.packet[2].parent.choice == 'vlan')
    assert(ip.__class__ == snappi.FlowIpv4)
    assert(flow.packet[3].parent.choice == 'ipv4')
    assert(tcp.__class__ == snappi.FlowTcp)
    assert(flow.packet[4].parent.choice == 'tcp')
    
    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
