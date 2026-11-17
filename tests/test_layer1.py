import pytest
import snappi


def test_layer1(api):
    """Test layer1
    """
    config = api.config()
    config.layer1.layer1()
    assert(len(config.layer1) == 1)


def test_layer1_mode():
    """Test layer1 mode default, enum values and serialization
    """
    config = snappi.Api().config()
    l1 = config.layer1.layer1(name="l1", port_names=["p1"])[-1]
    l1.speed = l1.SPEED_400_GBPS

    # mode defaults to mode_no_change
    assert l1.mode == l1.MODE_NO_CHANGE
    assert l1.serialize(l1.DICT)["mode"] == "mode_no_change"

    modes = [
        l1.MODE_NO_CHANGE,
        l1.MODE_DEFAULT,
        l1.MODE_MACSEC_128,
        l1.MODE_MACSEC_256,
        l1.MODE_HIGH_STREAM,
        l1.MODE_ROCEV2,
    ]
    for mode in modes:
        l1.mode = mode
        data = config.serialize()
        config_new = snappi.Api().config()
        config_new.deserialize(data)
        assert config_new.layer1[0].mode == mode
        assert config_new.layer1[0].speed == l1.SPEED_400_GBPS

    # invalid mode is rejected
    l1.mode = "mode_invalid"
    try:
        l1.serialize(l1.DICT)
        assert False
    except TypeError:
        assert True


if __name__ == '__main__':
    pytest.main(['-vv', '-s', __file__])
