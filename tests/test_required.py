def validate_serialize(node, expected, ignore_assert=False):
    try:
        getattr(node, 'serialize')(getattr(node, 'DICT'))
        assert ignore_assert
    except ValueError as e:
        assert any([exp in e.args[0] for exp in expected])


def test_required_serialize(api):

    config = api.config()
    p = config.ports.port()[-1]
    validate_serialize(p, p._REQUIRED)
    l = config.layer1.layer1()[-1]
    validate_serialize(l, l._REQUIRED)
    lg = config.lags.lag()[-1]
    validate_serialize(lg, lg._REQUIRED)
    d = config.devices.device()[-1]
    validate_serialize(d, d._REQUIRED)
    f = config.flows.flow()[-1]
    validate_serialize(f, f._REQUIRED)
    s = f.size
    validate_serialize(s, s._REQUIRED, True)
    r = f.rate
    validate_serialize(r, '', True)

    validate_serialize(config.ports, p._REQUIRED)
    validate_serialize(config.layer1, l._REQUIRED)
    validate_serialize(config.lags, lg._REQUIRED)
    validate_serialize(config.devices, d._REQUIRED)
    validate_serialize(config.flows, f._REQUIRED)

    validate_serialize(f.packet, '', True)


def validate_deserialize(node, input, expected, ignore_assert=False):
    try:
        getattr(node, 'deserialize')(input)
        assert ignore_assert
    except Exception as e:
        assert any([exp in e.args[0] for exp in expected])


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
    validate_deserialize(config.ports, ports, snappi.Port._REQUIRED, True)

    ports = [
        {
            'name': None,
            'location': None
        }
    ]
    validate_deserialize(config.ports, ports, snappi.Port._REQUIRED)

    lags = [
        {
            'name': None
        }
    ]

    validate_deserialize(config.lags, lags, snappi.Lag._REQUIRED)

    l = config.lags.lag()[-1]
    lp = [
        {}
    ]
    validate_deserialize(l.ports, lp, snappi.LagPort._REQUIRED)
