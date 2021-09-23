import pytest


def test_device_factory_methods(api):
    """Test device factory methods

    Device factory methods should populate saved structures
    """
    config = api.config()

    param = ('name', 'container name')
    device = config.devices.device(name=param[0])[-1]
    assert (device.name == param[0])

    name = 'eth name'
    eth = device.ethernets.ethernet()[-1]
    eth.name = name
    eth.port_name = "p1"
    assert (eth.name == name)
    eth.mac = '00:00:00:00:00:00'

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
