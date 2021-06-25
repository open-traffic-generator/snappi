import pytest


def test_port_factory_methods(api):
    """Test port factory method

    Ports factory method should populate saved structures
    """
    config = api.config()

    params = [
        ('Tx Port', '10.36.74.26;02;13'),
        ('Rx Port', '10.36.74.26;02;14'),
    ]
    for param in params:
        config.ports.port(name=param[0], location=param[1])

    for i in range(0, len(config.ports)):
        assert(config.ports[i].name == params[i][0])
        assert(config.ports[i].location == params[i][1])     

    print(config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
