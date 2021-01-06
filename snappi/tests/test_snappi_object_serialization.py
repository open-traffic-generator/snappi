import pytest


@pytest.mark.parametrize('encoding', ['yaml', 'json', 'dict'])
def test_snappi_object_serialization(api, b2b_config, encoding):
    """Test serialization and deserialization of Snappi objects
    """
    # serialize the configuration locally
    serialization1 = b2b_config.serialize(encoding)

    # use a mock web server to push the config
    api.set_config(b2b_config)

    # use a mock web server to pull the config
    config = api.get_config()

    # serialize the pulled config
    serialization2 = config.serialize(encoding)

    # verify the original config is the same as the roundtripped configuration
    assert (serialization1 == serialization2)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
