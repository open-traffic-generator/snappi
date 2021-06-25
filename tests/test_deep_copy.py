import pytest
import copy


def test_deep_copy(api):
    """Test deep copy of Snappi objects
    """
    config = api.config()
    f1, f2 = config.flows.flow(name='f1').flow(name='f2')
    f1.tx_rx.port.tx_name, f2.tx_rx.port.tx_name = 'p1', 'p2'
    f1.packet.ethernet().ipv4().tcp()
    f2.packet.ethernet().ipv4().udp()
    f3 = f1.clone()
    f3.name = 'f3'
    config.flows.append(f3)
    f4 = copy.deepcopy(f2)
    f4.name = 'f4'
    config.flows.append(f4)
    # print(config)
    assert(len(config.flows) == 4)
    assert(config.flows[-2].name == f3.name)
    assert(config.flows[-1].name == f4.name)



if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
