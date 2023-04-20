import pytest


def get_mock_config(grpc_api):
    """Test Metric Tag"""
    # create a new config object
    test_const = {
        "mac1": ["00:00:01:01:01:01", "00:00:01:01:01:02"],
        "mac2": ["00:00:01:01:01:03", "00:00:01:01:01:04"],
        "mac3": ["00:00:01:01:01:05", "00:00:01:01:01:06"],
        "ip1": ["1.1.1.1", "1.1.1.2"],
        "ip2": ["1.1.1.3", "1.1.1.4"],
        "ip3": ["1.1.1.5", "1.1.1.6"],
        "tcpPortValue": 80,
    }

    config = grpc_api.config()

    tx_port, rx_port = config.ports.port(
        name="Tx Port", location="10.36.74.26;02;13"
    ).port(name="Rx Port", location="10.36.74.26;02;14")

    flow = config.flows.flow(name="Tx -> Rx Flow")[0]
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    ethIngress, vlanIngress, ipIngress, tcpIngress = (
        flow.packet.ethernet().vlan().ipv4().tcp()
    )

    ethIngress.src.values = test_const["mac1"]
    ethIngress.dst.values = test_const["mac2"]
    ipIngress.src.values = test_const["ip1"]
    ipIngress.dst.values = test_const["ip2"]
    tcpIngress.src_port.value = test_const["tcpPortValue"]
    tcpIngress.dst_port.value = test_const["tcpPortValue"]
    ethIngress.dst.metric_tags.add(name="eth_dst_ingress", offset=40, length=8)
    ipIngress.dst.metric_tags.add(name="ip_dst_ingress", offset=24, length=8)

    ethEgress, ipEgress, tcpEgress = flow.egress_packet.ethernet().ipv4().tcp()

    ethEgress.src.values = test_const["mac2"]
    ethEgress.dst.values = test_const["mac3"]
    ipEgress.src.values = test_const["ip2"]
    ipEgress.dst.values = test_const["ip3"]
    tcpEgress.src_port.value = test_const["tcpPortValue"]
    tcpEgress.dst_port.value = test_const["tcpPortValue"]
    ethEgress.dst.metric_tags.add(name="eth_dst_egress", offset=40, length=8)
    ipEgress.dst.metric_tags.add(name="ip_dst_egress", offset=24, length=8)

    return config


def test_grpc_set_config(grpc_api):
    config = get_mock_config(grpc_api)
    response = grpc_api.set_config(config)
    assert len(response.warnings) == 1
    assert response.warnings[0] == "no"


def test_grpc_get_config(grpc_api):
    config = get_mock_config(grpc_api)
    response = grpc_api.get_config()
    assert config.serialize() == response.serialize()


if __name__ == "__main__":
    pytest.main(["-vv", "-s", __file__])
