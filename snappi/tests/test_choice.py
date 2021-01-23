import pytest
import snappi


def test_choice(api):
    """Test choice functionality

    Accessing a property that is a choice object or its child properties 
    should not change the choice.
    The only thing that should change the choice is when choice object's 
    child property is set.
    """
    config = api.config()

    device = config.devices.device(name='d1').device(name='d2')[0]
    assert (len(config.devices) == 2)
    assert (device.name == 'd1')
    device.ipv4.name = 'i'
    assert (device.choice == 'ipv4')
    device.ethernet.name = 'e'
    assert (device.choice == 'ethernet')
    device.bgpv4.name = 'b'
    assert (device.choice == 'bgpv4')
    device.ipv4.name
    assert (device.choice == 'bgpv4')

    flow = config.flows.flow(name='f')[-1]
    flow.tx_rx.device.tx_names = [config.devices[0].name]
    flow.tx_rx.device.rx_names = [config.devices[1].name]
    flow.packet.ethernet().vlan().vlan().ipv4()
    assert (flow.packet[0].parent.choice == 'ethernet')
    assert (flow.packet[1].parent.choice == 'vlan')
    assert (flow.packet[2].parent.choice == 'vlan')
    assert (flow.packet[3].parent.choice == 'ipv4')

    eth = flow.packet[0] 
    assert (eth.__class__ == snappi.FlowEthernet)
    eth.src.value = '00:00:01:00:00:01'
    assert(eth.src.choice == 'value')
    eth.src.values = ['00:00:01:00:00:01', '00:00:01:00:00:0a']
    assert(eth.src.choice == 'values')

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
