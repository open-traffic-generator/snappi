import pytest


def test_sanity_config(api):
    """Configure port based flow with ethernet, vlan, ipv4 and tcp
    """
    config = api.control_state().config_state.config

    # TBD: support factory methods for type: array, items: $ref: ...
    config.ports.port(name='Tx Port', location='10.36.74.26;02;13')
    config.ports.port(name='Rx Port', location='10.36.74.26;02;14')

    # TBD: support factory methods for type: array, items: $ref: ...
    flow = config.flows.flow(name='Tx -> Rx Flow')

    # add a source port and destination port to the flow
    # flow.tx_rx.port.tx_name = tx_port.name
    # flow.tx_rx.port.rx_name = rx_port.name

    # TBD: support factory methods for type: array, items: $ref: ...
    # base packet type is Flow_Header, all choices should be a factory method
    # for FlowHeader with the choice set to 'method name'
    pkt = flow.packet.ethernet().vlan().ipv4().tcp()

    # different ways to unpack packet headers
    eth, _, _, _ = flow.packet
    eth = flow.packet[0]
    eth = flow.packet[flow.packet.ETHERNET]
    ip = pkt[flow.packet.IPV4]

    # set flow header field patterns
    eth.src = '00:00:01:00:00:01'
    eth.dst = ['00:00:02:00:00:01', '00:00:02:00:00:01']
    ip.src.increment(start='1.1.1.1', step='0.0.0.1', count=10)
    ip.src.increment(start='1.1.1.1', step=1, count=10)
    ip.dst.decrement(start='1.1.2.200', step='0.0.0.1', count=10)
    # x-type int, model min/max 0-65535, implementation may or may not support full max
    # for best ux - implementation should preprocess config prior to setting anything
    # and throw error - do a dry run
    tcp.src_port.increment(start='1.1.1.1')
    tcp.src_port.increment(start='1')
    tcp.src_port.increment(start=1)

    # use constants
    ip.priority.dscp.phb = ip.priority.dscp.PHB_CS1

    # set flow packet size
    flow.size.fixed = 128

    # set the flow rate
    flow.rate.pps = 1000

    # set the flow duration
    flow.duration.fixed_packets = 10000

    # TBD
    # specify that a packet header field can be viewed as a
    # column name in the flow results using a user supplied column name
    eth.src.result_group = 'eth_src_mac_addr'

    # set the configuration
    api.set_config(config)


if __name__ == '__main__':
    pytest.main(['-s', __file__])