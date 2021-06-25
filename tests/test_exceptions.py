import pytest


def test_exceptions(api, b2b_config):
    """Test exceptions
    """
    b2b_config.options.port_options.location_preemption = 'invalid value'
    try:
        api.set_config(b2b_config)
    except Exception as e:
        print(e)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
