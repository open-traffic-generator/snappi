import snappi


def test_get():
    config = snappi.Api().config()
    f = config.flows.flow()[-1]
    s = f.get('size')
    assert s == None
    s = f.get('size', False)
    assert s == None
    fixed = f.get('size', True).get('fixed', True)
    assert fixed == 64
    assert f._properties.get('size') == None
    r = f.get('rate', True)
    assert r != None
    assert f._properties.get('rate') == None
