import pytest


@pytest.mark.skip(reason='proof of concept')
def test_capture(api):
    """Test capture functionality

    Capture should allow for multiple custom filters and multiple packet filters.
    Custom filters are used when the filter byte offset is known.
    Packet filters can be used when the packet structure 
    is known but the filter byte offset.
    Both custom filters and packet filters can be configured at the same time.
    """
    config = api.config()

    config.ports.port(name='a').port(name='b')
    capture = config.captures.capture(name='c', port_names=[port.name for port in config.ports])[-1]

    # add a list of custom filter
    custom1, custom2 = capture.custom_filters.custom().custom()
    custom1.offset = 0
    custom1.value = '000100000001'
    custom1.mask = 'FFFFFFFFFFFF'
    custom2.offset = 6
    custom2.value = '000200000001'
    custom2.mask = 'FFFFFFFFFFFF'

    # add a packet filter with multiple filters
    _, i_vlan, o_vlan = capture.packet_filter.ethernet().vlan().vlan()
    i_vlan.id.value = '01'
    i_vlan.id.mask = 'FF'
    o_vlan.id.value = '03'
    o_vlan.id.mask = 'FF'

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
