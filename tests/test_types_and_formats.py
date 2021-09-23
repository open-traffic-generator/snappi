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
    except TypeError:
        assert True

    # ip validation
    ip = f.packet.ipv4()[-1]
    ip.src.value = '1111'
    try:
        ip.src.serialize(ip.DICT)
        assert False
    except TypeError:
        assert True
    ip6 = f.packet.ipv6()[-1]
    ip6.src.value = ":a:1"
    try:
        ip6.src.serialize(ip6.DICT)
        assert False
    except TypeError:
        assert True
 
    #float validation
    f.duration.fixed_seconds.seconds = '10'
    f.duration.fixed_seconds.delay.microseconds = '20'
    try:
        f.duration.fixed_seconds.delay.serialize(f.DICT)
        assert False
    except TypeError:
        assert True
    
    # string validation
    config = snappi.Api().config()
    f = config.flows.flow()[-1]
    f.name = 100
    f.tx_rx.port.tx_name = 'test_tx'
    try:
        f.serialize(f.DICT)
        assert False
    except TypeError:
        assert True

    # integer validation
    d = config.devices.device()[-1]
    eth = d.ethernets.ethernet()[-1]
    eth.port_name = "p1"
    eth.name = 'eth'
    eth.mac = '00:00:00:00:00:00'
    eth.mtu = '1500'
    try:
        eth.serialize(eth.DICT)
        assert False
    except TypeError:
        assert True
    
    eth = f.packet.ethernet()[-1]
    # mac validation with list
    eth.src.values = ['a:b', ':0:0:0:0:0:0:1:']
    try:
        eth.src.serialize(eth.DICT)
        assert False
    except TypeError:
        assert True

    # reset existing choice
    f.size.fixed = 74
    assert f.size.serialize(f.DICT) == {'choice': 'fixed', 'fixed': 74}
    f.size.increment
    assert f.size.serialize(f.DICT) == {
        'choice': 'increment',
        'increment': {'start': 64, 'end': 1518, 'step': 1}
    }
    
