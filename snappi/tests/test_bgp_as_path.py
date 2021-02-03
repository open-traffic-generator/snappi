import pytest


def test_bgp_as_path(api):
    """Test BGP AS Path functionality
    """
    config = api.config()
    p1, p2 = config.ports.port(name='p1').port(name='p2')
    d = config.devices.device(name='d', container_name=p1.name)[-1]

    # setup ethernet
    eth = d.ethernet
    eth.name = 'e'
    eth.mac.value = '00:01:00:00:00:01'

    # setup ipv4
    ip = eth.ipv4
    ip.name = 'i4'
    ip.address.value = '1.1.1.1'

    # setup bgp
    bgp = ip.bgpv4
    bgp.name = 'b4'

    # setup route range
    rr = bgp.bgpv4_route_ranges.bgpv4routerange()[-1]
    rr.range_count = 1
    rr.address_count = 1
    rr.address.value = '4.4.4.4'
    rr.address_step.value = '0.0.0.1'
    rr.prefix.value = '32'
    rr.as_path.as_set_mode = rr.as_path.INCLUDE_AS_SEQ  # one per route range
    seg = rr.as_path.as_path_segments.bgpaspathsegment()[-1]  
    seg.as_numbers = [1, 2, 3, 4, 5, 6]

    print(config)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
