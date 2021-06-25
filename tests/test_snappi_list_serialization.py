import pytest


def test_snappi_list_serialization(api):
    """Test serialization and deserialization of Snappi list objects
    """
    config = api.config()
    config.ports.port(name='A').port(name='B').port(name='C')
    serialization1 = config.ports.serialize()
    config.ports.deserialize(serialization1)
    serialization2 = config.ports.serialize()
    assert(serialization1 == serialization2)
    

if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
