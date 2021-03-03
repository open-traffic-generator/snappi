import pytest


def test_device_factory_methods(api):
    """Test device factory methods

    Device factory methods should populate saved structures
    """
    config = api.config()

    param = ('name', 'container name')
    device = config.devices.device(name=param[0],
                                   container_name=param[1])[-1]
    assert (device.name == param[0])
    assert (device.container_name == param[1])

    name = 'eth name'
    eth = device.ethernet
    eth.name = name
    assert (eth.name == name)

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
