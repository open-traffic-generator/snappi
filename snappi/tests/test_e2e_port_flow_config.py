import pytest


def test_e2e_port_flow_config(api):
    """Demonstrates an end to end configuration
    """
    config = api.config()

    tx_port, rx_port = config.ports \
        .port(name='Tx Port', location='10.36.74.26;02;13') \
        .port(name='Rx Port', location='10.36.74.26;02;14')

    flow = config.flows.flow(name='Tx -> Rx Flow')[0]
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    _, _, ip, _ = flow.packet.ethernet().vlan().ipv4().tcp()

    eth = flow.packet[0]
    eth.src.value = '00:00:01:00:00:01'
    eth.dst.values = ['00:00:02:00:00:01', '00:00:02:00:00:01']

    # primitive choice
    # @property
    # getter raises exception if choice is not valid
    # setter sets the choice or no setter at all?
    ip.src.value = '1.1.1.1'
    ip.src.values = ['1.1.1.1']
    # complex choice
    # @staticmethod returns instance of class, sets the choice
    # getter raises exception if choice is not valid
    # ip.src.Increment(start='', step='', count=1)
    ip.src.increment.start = '1.1.1.1'
    ip.src.increment.step = '0.0.0.1'
    ip.src.increment.count = 10
    print(ip)
    ip.dst.decrement.start = '1.1.2.200'
    ip.dst.decrement.step = '0.0.0.1'
    ip.dst.decrement.count = 10

    ip.priority.dscp.ecn.value = ip.priority.dscp.ECN_CAPABLE_TRANSPORT_1
    ip.priority.dscp.ecn.metric_group = 'ip.priority.dscp.ecn'

    # set and get the configuration
    api.set_config(config)
    print(api.get_config())

    # start transmit
    transmit_state = api.transmit_state()
    transmit_state.state = 'start'
    api.set_transmit_state(transmit_state)

    # get port metrics
    port_metrics_request = api.port_metrics_request()
    port_metrics = api.get_port_metrics(port_metrics_request)
    for metric in port_metrics:
        print(metric)

    # get flow metrics
    flow_metrics_request = api.flow_metrics_request()
    flow_metrics = api.get_flow_metrics(flow_metrics_request)
    for metric in flow_metrics:
        print(metric)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
