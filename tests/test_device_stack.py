import pytest


def test_device_stack(api):
    """Test device dual stack functionality
    """
    config = api.config()
    p1, p2 = config.ports.port(name='p1').port(name='p2')
    d = config.devices.device(name='d')[-1]
    e = d.ethernets.ethernet()[-1]
    e.port_name = p1.name
    e.name = 'e'
    e.mac = '00:01:00:00:00:01'
    i4 = e.ipv4_addresses.ipv4()[-1]
    i4.name = 'i4'
    i4.address = '1.1.1.1'
    i4.gateway = '1.1.1.2'
    b4 = d.bgp.ipv4_interfaces.v4interface()[-1]
    b4.ipv4_name = i4.name
    b4.peers.v4peer()
    b4.peers[-1].name = "b4"
    b4.peers[-1].v4_routes.v4routerange(name="r1v4")
    b4.peers[-1].peer_address = '1.1.1.2'
    b4.peers[-1].as_type = b4.peers[-1].IBGP
    b4.peers[-1].as_number = 10
    d.bgp.router_id = "192.168.1.1"
    i6 = e.ipv6_addresses.ipv6()[-1]
    i6.name = 'i6'
    i6.address = '2001::1'
    i6.gateway = '2001::2'
    b6 = d.bgp.ipv6_interfaces.v6interface()[-1]
    b6.ipv6_name = i6.name
    b6.peers.v6peer()
    b6.peers[-1].peer_address = '2001::2'
    b6.peers[-1].as_type = 'ibgp'
    b6.peers[-1].as_number = 10
    b6.peers[-1].name = "b6"
    api.set_config(config)
    print(config)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
