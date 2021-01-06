import pytest


def test_e2e_port_flow_config():
    """Get port metrics
    """
    import snappi

    api = snappi.api.Api()

    config = api.config()

    tx_port, rx_port = config.ports \
        .port(name='Tx Port', location='10.36.74.26;02;13') \
        .port(name='Rx Port', location='10.36.74.26;02;14')

    flow = config.flows.flow(name='Tx -> Rx Flow')
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    eth, vlan, ip, tcp = flow.packet.ethernet().vlan().ipv4().tcp()

    eth.src.value = '00:00:01:00:00:01'
    eth.dst.value_list = ['00:00:02:00:00:01', '00:00:02:00:00:01']
    eth.dst.result_group = 'eth dst mac'

    ip.src.increment.start = '1.1.1.1'
    ip.src.increment.step = '0.0.0.1'
    ip.src.increment.count = 10

    ip.dst.decrement.start = '1.1.2.200'
    ip.dst.decrement.step = '0.0.0.1'
    ip.dst.decrement.count = 10

    ip.priority.dscp.phb.value_list = [8, 16, 32]
    ip.priority.dscp.ecn.value = 1
    tcp.src_port.increment.start = '10'
    tcp.dst_port.increment.start = 1

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
