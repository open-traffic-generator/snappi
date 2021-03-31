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
    p1, p2 = config.ports.port(name='p1').port(name='p2')
    config.devices.device(name='d1').device(name='d2')
    flow = config.flows.flow(name='f')[-1]

    flow.tx_rx.port.tx_name = p1.name
    flow.tx_rx.port.rx_name = p2.name
    assert (flow.tx_rx.choice == 'port')

    flow.tx_rx.device.tx_names = [config.devices[0].name]
    flow.tx_rx.device.rx_names = [config.devices[1].name]
    flow.tx_rx.device.mode = flow.tx_rx.device.ONE_TO_ONE

    assert (flow.tx_rx.choice == 'device')

    flow.packet.ethernet().vlan().vlan().ipv4()
    assert (flow.packet[0].parent.choice == 'ethernet')
    assert (flow.packet[1].parent.choice == 'vlan')
    assert (flow.packet[2].parent.choice == 'vlan')
    assert (flow.packet[3].parent.choice == 'ipv4')

    eth = flow.packet[0]
    assert (eth.__class__ == snappi.FlowEthernet)
    eth_type = eth.ether_type
    assert (eth_type.__class__ == snappi.PatternFlowEthernetEtherType)
    eth.src.value = '00:00:01:00:00:01'
    assert (eth.src.choice == 'value')
    eth.src.values = ['00:00:01:00:00:01', '00:00:01:00:00:0a']
    assert (eth.src.choice == 'values')

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
