import snappi


def test_types_and_formats():
    config = snappi.Api().config()
    f = config.flows.flow()[-1]
    eth = f.packet.ethernet()[-1]

    # mac validation
    eth.src.value = "11:22:33:aa:bb:cc:22"
    try:
        eth.src.serialize(eth.DICT)
        assert False
    except Exception as e:
        if "Invalid 11:22:33:aa:bb:cc:22 format" not in str(e):
            assert False

    # ip validation
    ip = f.packet.ipv4()[-1]
    ip.src.value = '1111'
    try:
        ip.src.serialize(ip.DICT)
        assert False
    except Exception as e:
        if "Invalid 1111 format" not in str(e):
            assert False
    ip6 = f.packet.ipv6()[-1]
    ip6.src.value = ":a:1"
    try:
        ip6.src.serialize(ip6.DICT)
        assert False
    except Exception as e:
        if "Invalid :a:1 format" not in str(e):
            assert False
 
    #float validation
    f.duration.fixed_seconds.seconds = '10'
    f.duration.fixed_seconds.delay.microseconds = '20'
    try:
        f.duration.fixed_seconds.delay.serialize(f.DICT)
        assert False
    except Exception as e:
        if "Invalid 20 format" not in str(e):
            assert False
    
    # string validation
    config = snappi.Api().config()
    config.ports.port(name="test_tx")
    f = config.flows.flow()[-1]
    f.name = 100
    f.tx_rx.port.tx_name = 'test_tx'
    try:
        f.serialize(f.DICT)
        assert False
    except Exception as e:
        if (
            "property name shall be of type <class 'str'>"
            " at <class 'snappi.snappi.Flow'>" not in str(e)
        ):
            print(e)
            assert False

    # integer validation
    d = config.devices.device()[-1]
    eth = d.ethernets.ethernet()[-1]
    eth.port_name = "test_tx"
    eth.name = 'eth'
    eth.mac = '00:00:00:00:00:00'
    eth.mtu = '1500'
    try:
        eth.serialize(eth.DICT)
        assert False
    except Exception as e:
        if "got 1500 of type <class 'str'>" not in str(e):
            assert False
    
    eth = f.packet.ethernet()[-1]
    # mac validation with list
    eth.src.values = ['a:b', ':0:0:0:0:0:0:1:']
    try:
        eth.src.serialize(eth.DICT)
        assert False
    except Exception as e:
        if "Invalid ['a:b', ':0:0:0:0:0:0:1:'] format" not in str(e):
            assert False

    # reset existing choice
    f.size.fixed = 74
    assert f.size.serialize(f.DICT) == {'choice': 'fixed', 'fixed': 74}
    f.size.increment
    assert f.size.serialize(f.DICT) == {
        'choice': 'increment',
        'increment': {'start': 64, 'end': 1518, 'step': 1}
    }
    
