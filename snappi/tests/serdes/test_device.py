def test_device(api):
    devices = {
        'devices': [
            {
                'container_name': 'Port1',
                'name': 'myDevice1'
            },
            {
                'container_name': 'Port2',
                'name': 'myDevice2'
            }
        ]
    }
    config = api.config()
    config.devices.device(
        name='myDevice1', container_name='Port1'
    ).device(
        name='myDevice2', container_name='Port2'
    )
    assert config.serialize(config.DICT) == devices
    config = api.config()
    config.deserialize(devices)
    assert len(config.devices) == len(devices['devices'])
    assert config.devices[0].name == devices['devices'][0]['name']
    assert config.devices[0].container_name == devices['devices'][0]['container_name']
    assert config.devices[1].name == devices['devices'][1]['name']
    assert config.devices[1].container_name == devices['devices'][1]['container_name']


def test_ethernet(api):
    devices = {
        'devices': [
            {
                'container_name': 'Port1',
                'name': 'myDevice1',
                'ethernet': {
                    'mac': 'abcd',
                    'mtu': 10,
                    'name': 'EthDevice'
                }
            }
        ]
    }
    config = api.config()
    d1 = config.devices.device(
        name='myDevice1', container_name='Port1'
    )[-1]
    d1.ethernet.mac = 'abcd'
    d1.ethernet.mtu = 10
    d1.ethernet.name = 'EthDevice'
    assert config.serialize(config.DICT) == devices
    config = api.config()
    config.deserialize(devices)
    eth = devices['devices'][0]['ethernet']
    assert len(config.devices) == len(devices['devices'])
    assert config.devices[0].ethernet.name == eth['name']
    assert config.devices[0].ethernet.mac == eth['mac']
    assert config.devices[0].ethernet.mtu == eth['mtu']
