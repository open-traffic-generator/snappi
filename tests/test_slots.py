import pytest
import snappi


def test_slots(api):
    """Test that dynamically created attributes are disallowed.
    This is done by ensuring that the special variable __dict__ is not 
    present in any classes.
    The proper use of the special variable __slots__ ensures that the __dict__
    special variable will not be present unless it is specifically included
    in the __slots__.
    """
    base = snappi.snappi.OpenApiBase()
    assert('__dict__' not in dir(base))
    obj = snappi.snappi.OpenApiObject()
    assert('__dict__' not in dir(obj))
    config = api.config()
    assert('__dict__' not in dir(config))


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
