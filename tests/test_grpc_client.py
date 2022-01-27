import pytest


def test_grpc_client(grpc_api):
    """Test device dual stack functionality
    """
    config = grpc_api.config()
    p1, p2 = config.ports.port(name='p1').port(name='p2')

    flow = config.flows.flow(name='Tx -> Rx Flow')[0]
    flow.tx_rx.port.tx_name = p1.name
    flow.tx_rx.port.rx_name = p2.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    eth, ip = flow.packet.ethernet().ipv4()

    eth.src.value = '00:00:01:00:00:01'
    eth.dst.values = ['00:00:02:00:00:01', '00:00:02:00:00:01']
    eth.dst.metric_group = 'eth dst mac'

    ip.src.increment.start = '1.1.1.1'
    ip.src.increment.step = '0.0.0.1'
    ip.src.increment.count = 10

    ip.dst.decrement.start = '1.1.2.200'
    ip.dst.decrement.step = '0.0.0.1'
    ip.dst.decrement.count = 10

    ip.priority.dscp.phb.value = 16
    ip.priority.dscp.ecn.value = 1

    # tcp.src_port.increment.start = 10
    # tcp.dst_port.increment.start = 1

    result = grpc_api.set_config(config)
    print("Result : %s" %result)
    result1 = grpc_api.get_config()
    print(result1)


if __name__ == '__main__':
    pytest.main(['-s', __file__])

