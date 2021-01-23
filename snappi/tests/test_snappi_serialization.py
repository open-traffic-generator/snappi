import pytest


def test_device_default(api):
    """
    Configure the device containers with no params passed.
    Validate the configuration is being serialized as expected
    """
    config = api.config()
    config.devices.clear()
    cd = config.devices
    tv = get_expected_result()["devices"]
    devices = []
    for d in tv:
        getattr(cd, d)()
        dt = dict()
        if d in ['ethernet', 'ipv4', 'ipv6', 'bgpv4']:
            dt['choice'] = d
            dt[d] = tv[d]
            dt.update(tv["device"])
        else:
            dt.update(tv[d])
        devices.append(dt)
    assert cd.serialize(config.DICT) == devices
    assert len(tv) == len(cd)


def test_device_valid(api):
    """
    Configure the devices with valid param values,
    Validate the configuration is being serialized as expected
    """
    config = api.config()
    config.devices.clear()
    cd = config.devices
    tv = get_expected_result("valid")["devices"]
    df_dev = get_expected_result()["devices"]["device"]
    devices = []
    for d in tv:
        dt = dict()
        kwargs = get_arg_values(tv[d], d)
        getattr(cd, d)(**kwargs)
        if d in ['ethernet', 'ipv4', 'ipv6', 'bgpv4']:
            dt['choice'] = d
            dt.update(df_dev)
            dt[d] = dict(kwargs)
        else:
            dt = dict(kwargs)
        devices.append(dt)
    import pdb; pdb.set_trace()
    assert cd.serialize(config.DICT) == devices
    assert len(tv) == len(cd)


def get_arg_values(dt, key):
    kwargs = dict()
    val_map = {
        int : 10,
        str : "abc",
        list : [],
        dict : {},
        None : None
    }
    import random
    for k in dt:
        r = random.randint(0,len(dt[k]) - 1)
        val = val_map[dt[k][r]]
        kwargs[k] = val
    return kwargs


def get_expected_result(for_test_type="default"):
    n = {
        "default": None,
        "valid": (str, None),
        "invalid": (int, list, dict)
    }

    dc = {
        "default": 1,
        "valid": [int],
        "invalid": (str, list, dict)
    }

    dc_1 = {
        "default": None,
        "valid": (int, None),
        "invalid": (str, list, dict)
    }

    test_validation = {
        "ports": {
            "port": {
                "location": n[for_test_type],
                "name": n[for_test_type]
            }
        },
        "devices": {
            "device": {
                "container_name": n[for_test_type],
                "device_count": dc[for_test_type],
                "name": n[for_test_type]
            },
            "ethernet": {
                "name": n[for_test_type]
            },
            "ipv4": {
                "name": n[for_test_type],
            },
            "ipv6": {
                "name": n[for_test_type],
            },
            "bgpv4": {
                "name": n[for_test_type],
                "as_type": dc_1[for_test_type],
            },
        }
    }
    return test_validation

