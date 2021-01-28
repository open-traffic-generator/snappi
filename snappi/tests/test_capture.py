import pytest
import snappi


def test_capture(api):
    """Test capture functionality

    Capture should allow for multiple filters.
    The capture.filters is a list of CaptureFilter objects which can be 
    CaptureCustom, CaptureEthernet, CaptureVlan and CaptureIpv4.
    A CaptureCustom object is used when packet header is not available.
    """
    config = api.config()

    # create a capture
    config.ports.port(name='port1').port(name='port2')
    capture = config.captures.capture()[0]
    capture.name = 'capture1'
    capture.port_names = [port.name for port in config.ports]

    # add 2 custom filters
    capture.filters.custom().custom()
    src_mac = capture.filters[0]  # type: snappi.CustomFilter
    src_mac.offset = 0
    src_mac.value = '000100000001'
    src_mac.mask = 'FFFFFFFFFFFF'
    dst_mac = capture.filters[1]  # type: snappi.CustomFilter
    dst_mac.offset = 0
    dst_mac.value = '000200000001'
    dst_mac.mask = 'FFFFFFFFFFFF'
    print(config)

    # add packet filters for inner and outer vlan
    capture.filters.clear()
    _, i_vlan, o_vlan = capture.filters.ethernet().vlan().vlan()
    i_vlan.id.value = '01'
    i_vlan.id.mask = 'FF'
    o_vlan.id.value = '03'
    o_vlan.id.mask = 'FF'
    print(config)

    # add packet filters for a custom packet after eth, ip
    capture.filters.clear()
    eth, _, custom = capture.filters.ethernet().ipv4().custom()
    eth.src.value = '000001000001'
    eth.src.mask = 'FFFFFFFFFFFF'
    custom.offset = 16
    custom.value = '010101'
    custom.mask = 'FFFFFF'
    print(config)

    api.set_config(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
