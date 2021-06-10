def test_defaults(api):
    defaults = {
        'flows': [
            {
                'name': None,
                'size': {
                    'choice': 'fixed',
                    'fixed': 64
                },
                'rate': {
                    'choice': 'pps',
                    'pps': 1000
                },
                'packet': [
                    {
                        'choice': 'ethernet',
                        'ethernet': {
                            'dst': {
                                'choice': 'value',
                                'value': '00:00:00:00:00:00',
                                'metric_group': None
                            },
                            'src': {
                                'choice': 'values',
                                'values': ['00:00:00:00:00:00'],
                                'metric_group': None,
                                'value': '00:00:00:00:00:00'
                            },
                            'ether_type': {
                                'auto': 'auto',
                                'choice': 'auto',
                                'metric_group': None
                            }
                        },
                    },
                    {
                        'choice': 'ipv4',
                        'ipv4': {
                            'src': {
                                'choice': 'increment',
                                'metric_group': None,
                                'increment': {
                                    'count': 1,
                                    'start': '0.0.0.0',
                                    'step': '0.0.0.1'
                                },
                                'value': '0.0.0.0' # side effect
                            }
                        }

                    }
                ]
            }
        ],
        'lags': [
            {
                'name': None,
                'ports': [
                    {
                        'port_name': None,
                        'protocol': {
                            'choice': 'lacp',
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
                'name': None,
                'port_names': None,
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
    layer1.flow_control
    l = config.lags.lag()[-1]
    p = l.ports.port()[-1]
    p.protocol
    f = config.flows.flow()[-1]
    f.size.fixed
    f.rate
    eth, ipv4 = f.packet.ethernet().ipv4()
    eth.dst
    eth.src.values
    eth.ether_type
    ipv4.src.increment
    assert config.serialize(config.DICT) == defaults


def test_defaults_by_deserialize(api):
    defaults = {
        'flows': [
            {
                'name': None,
                'size': {
                    'choice': 'fixed',
                    'fixed': 64
                },
            }
        ],
        'lags': [
            {
                'name': None,
                'ports': [
                    {
                        'port_name': None,
                        'protocol': {
                            'choice': 'lacp',
                            'lacp': {
                                'actor_activity': 'active',
                                'actor_key': 0,
                                'actor_port_number': None, # made None
                                'actor_port_priority': None, # made None
                                'actor_system_id': None, # made None
                                'actor_system_priority': 0,
                                'lacpdu_periodic_time_interval': 0,    
                                'lacpdu_timeout': 0
                            }
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
                'mtu': None, # mtu made None
                'name': None,
                'port_names': None,
                'promiscuous': False,
                'speed': 'speed_10_gbps',
                'flow_control': { # no data in flow control
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
