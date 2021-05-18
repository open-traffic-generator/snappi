def test_defaults(api):
    defaults = {
        'flows': [{'name': None, 'size': {'choice': 'fixed', 'fixed': 64}}],
        'lags': [{'name': None,
           'ports': [{'port_name': None,
                      'protocol': {'choice': 'lacp',
                                   'lacp': {'actor_activity': 'active',
                                            'actor_key': 0,
                                            'actor_port_number': 0,
                                            'actor_port_priority': 1,
                                            'actor_system_id': '00:00:00:00:00:00',
                                            'actor_system_priority': 0,
                                            'lacpdu_periodic_time_interval': 0,    
                                            'lacpdu_timeout': 0}}}]}],
        'layer1': [
            {
                'auto_negotiate': True,
                'ieee_media_defaults': True,
                'media': None,
                'mtu': 1500,
                'name': None,
                'port_names': None,
                'promiscuous': False,
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
    p.protocol.lacp
    f = config.flows.flow()[-1]
    f.size.fixed
    assert config.serialize(config.DICT) == defaults
