import pytest


def test_choice(api):
    """Test choice functionality

    Accessing a property that is a choice object or its child properties 
    should not change the choice.
    The only thing that should change the choice is when choice object's 
    child property is set.
    """
    config = api.config()

    flow = config.flows.flow(name='f')[-1]
    eth, vlan1, vlan2, ip = flow.packet.ethernet().vlan().vlan().ipv4()
    assert(flow.packet[0]._parent.choice == 'ethernet')
    assert(flow.packet[1]._parent.choice == 'vlan')
    assert(flow.packet[2]._parent.choice == 'vlan')
    assert(flow.packet[3]._parent.choice == 'ipv4')

    device = config.devices.device(name='d')[-1]
    device.ipv4.name = 'i'
    assert(device.choice == 'ipv4')
    device.ethernet.name = 'e'
    assert(device.choice == 'ethernet')
    device.bgpv4.name = 'b'
    assert(device.choice == 'bgpv4')
    device.ipv4.name
    assert(device.choice == 'bgpv4')
    
    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
