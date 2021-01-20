import pytest


def test_port_metrics(api, b2b_config):
    """Get port metrics
    """
    api.set_config(b2b_config)
    port_metrics_request = api.port_metrics_request()
    port_metrics = api.get_port_metrics(port_metrics_request)
    for metric in port_metrics:
        print(metric)


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
