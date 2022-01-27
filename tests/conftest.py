import pytest

@pytest.fixture(scope='session')
def api():
    """Demonstrates creating a top level Api instance.
    """
    from .snappiserver import SnappiServer
    pytest.snappiserver = SnappiServer().start()
    import snappi
    yield snappi.api(location='http://127.0.0.1:80')


@pytest.fixture(scope='session')
def grpc_api():
    from .snappigrpcserver import grpc_server, GRPC_PORT
    grpc_server()
    import snappi
    yield snappi.api(
        location="127.0.0.1:%s" % GRPC_PORT, transport=snappi.Transport.GRPC
    )


@pytest.fixture(scope='function')
def b2b_config(api):
    """Demonstrates creating a back to back configuration of tx and rx 
    ports, devices and a single flow using those ports as endpoints for
    transmit and receive.
    """
    config = api.config()
    import snappi
    config = snappi.Api().config()
    config.options.port_options.location_preemption = True
    
    tx_port, rx_port = config.ports \
        .port(name='Tx Port', location='10.36.74.26;02;13') \
        .port(name='Rx Port', location='10.36.74.26;02;14')

    tx_device, rx_device = (config.devices \
        .device(name='Tx Devices')
        .device(name='Rx Devices')
    )

    tx_device.ethernets.ethernet(port_name=tx_port.name)
    rx_device.ethernets.ethernet(port_name=rx_port.name)

    tx_device.ethernets[-1].name = 'Tx Eth'
    tx_device.ethernets[-1].mac = '00:00:01:00:00:01'
    tx_device.ethernets[-1].ipv4_addresses.ipv4()
    tx_device.ethernets[-1].ipv4_addresses[-1].name = 'Tx Ipv4'
    tx_device.ethernets[-1].ipv4_addresses[-1].address = '1.1.1.1'
    tx_device.ethernets[-1].ipv4_addresses[-1].gateway = '1.1.2.1'
    tx_device.ethernets[-1].ipv4_addresses[-1].prefix = 16
    vlan1, vlan2 = tx_device.ethernets[-1].vlans.vlan(name='v1').vlan(name='v2')
    vlan1.id = 1
    vlan2.id = 2
    rx_device.ethernets[-1].name = 'Rx Eth'
    rx_device.ethernets[-1].mac = '00:00:01:00:00:02'

    flow = config.flows.flow(name='Tx -> Rx Flow')[0]
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    eth, vlan, ip, tcp = flow.packet.ethernet().vlan().ipv4().tcp()

    eth.src.value = '00:00:01:00:00:01'
    eth.dst.values = ['00:00:02:00:00:01', '00:00:02:00:00:01']
    eth.dst.metric_group = 'eth dst mac'

    ip.src.increment.start = '1.1.1.1'
    ip.src.increment.step = '0.0.0.1'
    ip.src.increment.count = 10

    ip.dst.decrement.start = '1.1.2.200'
    ip.dst.decrement.step = '0.0.0.1'
    ip.dst.decrement.count = 10

    ip.priority.dscp.phb.values = [8, 16, 32]
    ip.priority.dscp.ecn.value = 1

    tcp.src_port.increment.start = 10
    tcp.dst_port.increment.start = 1

    return config

