import pytest


def test_snappi_object_equality(api, b2b_config):
    """Test equality of Snappi objects
    """
    # create a serialized yaml config of the b2b_config
    yaml_content = b2b_config.serialize(b2b_config.YAML)

    # create a new config object that uses the serialized yaml
    config = api.config()
    config = config.deserialize(yaml_content)
    
    # verify the original config is the same as the config created from the yaml
    assert (b2b_config == config)
    assert (b2b_config.ports == config.ports)

    config.ports.port()
    assert (b2b_config != config)
    assert (b2b_config.ports != config.ports)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
