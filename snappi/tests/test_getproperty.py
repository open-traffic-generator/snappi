import snappi


def test_getproperty():
    config = snappi.Api().config()
    f = config.flows.flow()[-1]
    s = f.getproperty('size')
    assert s == None
    s = f.getproperty('size', False)
    assert s == None
    fixed = f.getproperty('size', True).getproperty('fixed', True)
    assert fixed == 64
    assert f._properties.get('size') == None
    r = f.getproperty('rate', True)
    assert r != None
    assert f._properties.get('rate') == None
