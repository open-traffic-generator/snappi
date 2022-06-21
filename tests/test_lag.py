import pytest


def test_lag(api):
    """Test LAG functionality
    """
    config = api.config()

    # setup port container
    p1, p2, p3 = (config.ports
        .port(name='p1')
        .port(name='p2')
        .port(name='p3')
    )

    # setup lag container
    l1 = config.lags.lag(name='l1')[-1]
    lacp = l1.protocol.lacp
    for port in config.ports:
        lp = l1.ports.port(port_name=port.name)[-1]
        lp.ethernet.name = 'lpe {}'.format(port.name)
        lp.ethernet.mac = '00:00:01:00:00:01'
        lacp.actor_system_id = '00:00:0A:00:00:01'
        lacp.actor_key = 1
        lp.lacp.actor_port_number = 10

    api.set_config(config)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
