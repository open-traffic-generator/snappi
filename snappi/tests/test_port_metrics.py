import pytest


def test_port_metrics(api):
    """Get port metrics
    """
    request = api.port_metrics_request()
    port_metrics = api.get_port_metrics(request)
    print(port_metrics)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
