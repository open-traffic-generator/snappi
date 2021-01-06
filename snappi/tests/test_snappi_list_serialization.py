import pytest


def test_snappi_list_serialization(api):
    """Test serialization and deserialization of Snappi list objects
    """
    port_metrics1 = api.port_metrics().portmetric().portmetric()
    serialization = port_metrics1.serialize()
    port_metrics2 = api.port_metrics()
    port_metrics2.deserialize(serialization)
    print(port_metrics2)
    

if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
