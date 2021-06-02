def test_defaults(api):
    defaults = {
        'flows': [
            {
                'tx_rx': {
                    'choice': 'port',
                    'port': {
                        'tx_name': 'ptest',
                        'rx_name': None
                    }
                },
                'name': "f1",
                'size': {
                    'choice': 'fixed',
                    'fixed': 64
                },
                'rate': {
                    'choice': 'pps', # default choice is pps
                    'pps': 1000
                }
            }
        ],
        'lags': [
            {
                'name': "abc",
                'ports': [
                    {
                        'port_name': 'lagport',
                        'protocol': {
                            'choice': 'lacp', # default choice is lacp
                            'lacp': {
                                'actor_activity': 'active',
                                'actor_key': 0,
                                'actor_port_number': 0,
                                'actor_port_priority': 1,
                                'actor_system_id': '00:00:00:00:00:00',
                                'actor_system_priority': 0,
                                'lacpdu_periodic_time_interval': 0,    
                                'lacpdu_timeout': 0
                            }
                        },
                        'ethernet': {
                            'mac': '00:00:00:00:00:00', 'mtu': 1500, 'name': 'test'
                        }
                    }
                ]
            }
        ],
        'layer1': [
            {
                'auto_negotiate': True,
                'ieee_media_defaults': True,
                'media': None,
                'mtu': 1500,
                'name': "abc",
                'port_names': ["test"],
                'promiscuous': True,
                'speed': 'speed_10_gbps',
                'flow_control': {
                    'directed_address': '0180C2000001'
                }
            }
        ]
    }
    config = api.config()
    layer1 = config.layer1.layer1()[-1]
    layer1.name = "abc"
    layer1.port_names = ["test"]
    layer1.flow_control
    l = config.lags.lag()[-1]
    l.name = "abc"
    p = l.ports.port()[-1]
    p.port_name = "lagport"
    p.protocol
    p.ethernet.name = "test"
    f = config.flows.flow()[-1]
    f.name = "f1"
    f.size.fixed
    f.tx_rx.port.tx_name = "ptest"
    f.rate

    assert config.serialize(config.DICT) == defaults


def test_defaults_by_deserialize(api):
    defaults = {
        'flows': [
            {
                'tx_rx': {
                    'choice': 'port',
                    'port': {
                        'tx_name': 'ptest',
                        'rx_name': None
                    }
                },
                'name': "f1",
                'size': {
                    'choice': 'fixed',
                    'fixed': 64
                },
                'rate': {
                    'choice': 'pps',
                    'pps': None # defaults to a value
                }
            }
        ],
        'lags': [
            {
                'name': "abc",
                'ports': [
                    {
                        'port_name': 'lagport',
                        'protocol': {
                            'choice': 'lacp',
                            'lacp': {
                                'actor_activity': 'active',
                                'actor_key': 0,
                                'actor_port_number': None, # default is 0
                                'actor_port_priority': 1,
                                'actor_system_id': '00:00:00:00:00:00',
                                'actor_system_priority': 0,
                                'lacpdu_periodic_time_interval': 0,    
                                'lacpdu_timeout': 0
                            }
                        },
                        'ethernet': {
                            'mac': '00:00:00:00:00:00', 'mtu': 1500, 'name': 'test'
                        }
                    }
                ]
            }
        ],
        'layer1': [
            {
                'auto_negotiate': True,
                'ieee_media_defaults': True,
                'media': None,
                'mtu': 1500,
                'name': "abc",
                'port_names': ["test"],
                'promiscuous': True,
                'speed': 'speed_10_gbps',
                'flow_control': {
                    'directed_address': '0180C2000001'
                }
            }
        ]
    }
    config = api.config()
    config.deserialize(defaults)

    assert config.layer1[0].mtu == 1500
    assert config.layer1[0].flow_control.directed_address == '0180C2000001'
    assert config.lags[0].ports[0].protocol.choice == 'lacp'
    assert config.lags[0].ports[0].protocol.lacp.actor_port_number == 0
    assert config.lags[0].ports[0].protocol.lacp.actor_port_priority == 1

    assert config.lags[0].ports[0].protocol.lacp.actor_system_id == '00:00:00:00:00:00'
