import pytest


def test_snappi_lists(api):
    """Demonstrate using a SnappiList object
    """
    config = api.config()

    flow = config.flows.flow(name='1')
    eth, vlan, vlan1 = flow.packet.ethernet().vlan().vlan()
    vlan.id.value = 1

    flow = config.flows.flow(name='2')
    eth, vlan, ipv4, tcp = flow.packet.ethernet().vlan().ipv4().tcp()
    vlan.id.value = 2

    flow = config.flows.flow(name='3')
    eth, vlan = flow.packet.ethernet().vlan()
    vlan.id.value = 3

    api.set_config(config)
    
    assert (len(config.flows) == 3)
    assert (config.flows[0].name == '1')
    assert (config.flows[1].name == '2')
    assert (config.flows[2].name == '3')
    assert (len(config.flows[0].packet) == 3)
    assert (len(config.flows[1].packet) == 4)
    assert (len(config.flows[2].packet) == 2)
    assert (config.flows[0].packet[1].id.value == 1)
    assert (config.flows[1].packet[1].id.value == 2)
    assert (config.flows[2].packet[1].id.value == 3)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
