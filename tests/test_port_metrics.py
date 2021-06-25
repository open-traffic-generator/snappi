import pytest


def test_port_metrics(api, b2b_config):
    """Get port metrics
    """
    api.set_config(b2b_config)
    req = api.metrics_request()
    req.choice = req.PORT
    res = api.get_metrics(req)
    for metric in res.port_metrics:
        print(metric)

if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
