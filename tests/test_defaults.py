def test_defaults(api):
    defaults = {
        'flows': [
            {
                'tx_rx': {
                    'choice': 'port',
                    'port': {
                        'tx_name': 'ptest',
                    }
                },
                'name': "f1",
                'size': {
                    'choice': 'fixed',
                    'fixed': 64
                },
                'rate': {
                    'choice': 'pps', # default choice is pps
                    'pps': "1000"
                },
                'packet': [
                    {
                        'choice': 'ethernet',
                        'ethernet': {
                            'dst': {
                                'choice': 'auto',
                                'auto': '00:00:00:00:00:00',
                            },
                            'src': {
                                'choice': 'values',
                                'values': ['00:00:00:00:00:00'],
                            },
                            'ether_type': {
                                'auto': 65535,
                                'choice': 'auto',
                            }
                        },
                    },
                    {
                        'choice': 'ipv4',
                        'ipv4': {
                            'src': {
                                'choice': 'increment',
                                'increment': {
                                    'count': 1,
                                    'start': '0.0.0.0',
                                    'step': '0.0.0.1'
                                },
                            }
                        }

                    }
                ]
            }
        ],
        'lags': [
            {
                'name': "abc",
                'protocol': {
                    'choice': 'lacp', # default choice is lacp
                    'lacp': {
                        'actor_key': 0,
                        'actor_system_id': '00:00:00:00:00:00',
                        'actor_system_priority': 0,
                    }
                },
                'ports': [
                    {
                        'port_name': 'lagport',
                        'lacp': {
                            'actor_activity': 'active',
                            'actor_port_number': 0,
                            'actor_port_priority': 1,
                            'lacpdu_periodic_time_interval': 0,    
                            'lacpdu_timeout': 0
                        },
                        'ethernet': {
                            'mac': '00:00:00:00:00:00', 'mtu': 1500, 'name': 'test'
                        }
                    }
                ],
                "min_links": 1
            }
        ],
        'layer1': [
            {
                'auto_negotiate': True,
                'ieee_media_defaults': True,
                'mtu': 1500,
                'name': "abc",
                'port_names': ["test"],
                'promiscuous': True,
                'speed': 'speed_10_gbps',
                'flow_control': {
                    'choice': 'ieee_802_1qbb',
                    'directed_address': '01:80:C2:00:00:01',
                    'ieee_802_1qbb': {
                        'pfc_class_0': 0,
                        'pfc_class_1': 1,
                        'pfc_class_2': 2,
                        'pfc_class_3': 3,
                        'pfc_class_4': 4,
                        'pfc_class_5': 5,
                        'pfc_class_6': 6,
                        'pfc_class_7': 7,
                        'pfc_delay': 0
                    }
                }
            }
        ]
    }
    config = api.config()
    layer1 = config.layer1.layer1()[-1]
    layer1.name = "abc"
    layer1.port_names = ["test"]
    layer1.ieee_media_defaults = True
    layer1.auto_negotiate = True
    layer1.flow_control
    l = config.lags.lag()[-1]
    l.name = "abc"
    p = l.ports.port()[-1]
    p.port_name = "lagport"
    p.lacp
    l.protocol
    p.ethernet.name = "test"
    p.ethernet.mac = "00:00:00:00:00:00"
    f = config.flows.flow()[-1]
    f.name = "f1"
    f.size.fixed
    f.tx_rx.port.tx_name = "ptest"
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
                'protocol': {
                    'choice': 'lacp', # default choice is lacp
                    'lacp': {
                        'actor_key': 0,
                        'actor_system_id': '00:00:00:00:00:00',
                        'actor_system_priority': 0,
                    }
                },
                'ports': [
                    {
                        'port_name': 'lagport',
                        'lacp': {
                            'actor_activity': 'active',
                            'actor_port_number': None, #default is 0
                            'actor_port_priority': 1,
                            'lacpdu_periodic_time_interval': 0,    
                            'lacpdu_timeout': 0
                        },
                        'ethernet': {
                            'mac': '00:00:00:00:00:00', 'mtu': 1500, 'name': 'test'
                        }
                    }
                ],
                "min_links": 1
            }
        ],
        'layer1': [
            {
                'auto_negotiate': True,
                'ieee_media_defaults': True,
                'mtu': 1500,
                'name': "abc",
                'port_names': ["test"],
                'promiscuous': True,
                'speed': 'speed_10_gbps',
                'flow_control': {
                    'directed_address': '01:80:C2:00:00:01'
                }
            }
        ]
    }
    config = api.config()
    config.deserialize(defaults)

    assert config.layer1[0].mtu == 1500
    assert config.layer1[0].flow_control.directed_address == '01:80:C2:00:00:01'
    assert config.lags[0].protocol.choice == 'lacp'
    assert config.lags[0].ports[0].lacp.actor_port_number == 0
    assert config.lags[0].ports[0].lacp.actor_port_priority == 1
    assert config.lags[0].protocol.lacp.actor_system_id == '00:00:00:00:00:00'
