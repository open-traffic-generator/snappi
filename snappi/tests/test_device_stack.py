import pytest


def test_device_stack(api):
    """Test device dual stack functionality
    """
    config = api.config()
    p1, p2 = config.ports.port(name='p1').port(name='p2')
    d = config.devices.device(name='d', container_name=p1.name)[-1]
    e = d.ethernet
    e.name = 'e'
    e.mac = '00:01:00:00:00:01'
    i4 = e.ipv4
    i4.name = 'i4'
    i4.address = '1.1.1.1'
    b4 = i4.bgpv4
    b4.name = 'b4'
    b4.bgpv4_routes.bgpv4route()
    i6 = e.ipv6
    i6.name = 'i6'
    i6.address = '2001::1'
    b6 = i6.bgpv6
    b6.name = 'b6'
    api.set_config(config)
    print(config)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
