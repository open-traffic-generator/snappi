import snappi


def test_getproperty():
    config = snappi.Api().config()
    f = config.flows.flow()[-1]
    s = f.getproperty('size')
    assert s == None
    fixed = f.getproperty('size', True).getproperty('fixed', True)
    assert fixed == 64
    r = f.getproperty('rate', True)
    assert r != None
