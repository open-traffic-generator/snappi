import pytest
import snappi


def test_hints():
    """Demonstrate hinting of Snappi objects
    """
    api = snappi.Api()
    config = api.config()
    p1, p2, p3 = config.ports.port(name='P1').port(name='P2').port(name='P3') 
    p1 = config.ports[0]
    p2 = config.ports[1]
    p1, p2, p3 = config.ports 
    config.ports.port().port()
    for port in config.ports:
        print(port)
    print(p1, p2)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
