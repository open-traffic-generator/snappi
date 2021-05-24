def validate_serialize(node, expected, ignore_assert=False):
    try:
        getattr(node, 'serialize')(getattr(node, 'DICT'))
        assert ignore_assert
    except ValueError as e:
        if expected in e.args[0]:
            assert True
        else:
            assert False


def test_required_serialize(api):

    config = api.config()
    p = config.ports.port()[-1]
    validate_serialize(p, '{}'.format(p._REQUIRED))
    l = config.layer1.layer1()[-1]
    validate_serialize(l, '{}'.format(l._REQUIRED))
    lg = config.lags.lag()[-1]
    validate_serialize(lg, '{}'.format(lg._REQUIRED))
    d = config.devices.device()[-1]
    validate_serialize(d, '{}'.format(d._REQUIRED))
    f = config.flows.flow()[-1]
    validate_serialize(f, '{}'.format(f._REQUIRED))
    s = f.size
    validate_serialize(s, '{}'.format(s._REQUIRED))
    r = f.rate
    validate_serialize(r, '', True)

    validate_serialize(config.ports, '{}'.format(p._REQUIRED))
    validate_serialize(config.layer1, '{}'.format(l._REQUIRED))
    validate_serialize(config.lags, '{}'.format(lg._REQUIRED))
    validate_serialize(config.devices, '{}'.format(d._REQUIRED))
    validate_serialize(config.flows, '{}'.format(f._REQUIRED))

    validate_serialize(f.packet, '', True)


def validate_deserialize(node, input, expected, ignore_assert=False):
    try:
        getattr(node, 'deserialize')(input)
        assert ignore_assert
    except ValueError as e:
        if expected in e.args[0]:
            assert True
        else:
            assert False


def test_required_deserialize(api):
    config = api.config()
    import snappi
    config = snappi.Api().config()

    ports = [
        {
            'name': "Test",
            'location': None
        }
    ]
    validate_deserialize(config.ports, ports, '{}'.format(snappi.Port._REQUIRED), True)

    ports = [
        {
            'name': None,
            'location': None
        }
    ]
    validate_deserialize(config.ports, ports, '{}'.format(snappi.Port._REQUIRED))

    lags = [
        {
            'name': None
        }
    ]

    validate_deserialize(config.lags, lags, '{}'.format(snappi.Lag._REQUIRED))

    l = config.lags.lag()[-1]
    lp = [
        {}
    ]
    validate_deserialize(l.ports, lp, '{}'.format(snappi.LagPort._REQUIRED))

    lp = [
        {
            'port_name': 'abcd',
            'protocol': {
                'choice': 'static'
            },
            'ethernet': {
                'name': None
            }
        }
    ]
    validate_deserialize(l.ports, lp, '{}'.format(snappi.DeviceEthernetBase._REQUIRED))