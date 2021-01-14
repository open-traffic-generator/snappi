import pytest


def test_print(api, b2b_config):
    """Demonstrate output of Snappi objects
    """
    print(b2b_config.ports)
    print(b2b_config.ports[0])
    print(b2b_config.flows)
    print(b2b_config.devices)
    print(b2b_config)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
