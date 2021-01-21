import pytest


def test_layer1(api):
    """Test layer1
    """
    config = api.config()
    config.layer1.layer1()
    assert(len(config.layer1) == 1)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
