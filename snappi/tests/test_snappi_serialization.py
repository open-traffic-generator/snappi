def test_device_default(api):
    """
    configure the devices with defaults and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    config.devices.clear()
    
    dt = get_expected_result()["devices"]["device"]
    final_devices = list()
    dev_list = [None, "ethernet", "ipv4", "ipv6", "bgpv4"]
    for i,d in enumerate(dev_list):
        config.devices.device()
        final_devices.append(dict(dt))
        if d:
            getattr(config.devices[i], d)
            final_devices[i][d] = {"name": None}
            final_devices[i]["choice"] = d
        if d == "bgpv4":
            final_devices[i][d] = {"name": None, "as_type": None}
        assert final_devices == config.devices.serialize(config.DICT)
        assert len(config.devices) == len(final_devices)

    config.devices.clear()
    final_devices = list()
    for i,d in enumerate(dev_list):
        config.devices.device()
        final_devices.append(dict(dt))
        if d:
            getattr(config.devices[i], d).name = "{}_device".format(d)
            final_devices[i][d] = {"name": "{}_device".format(d)}
            final_devices[i]["choice"] = d
        if d == "bgpv4":
            final_devices[i][d] = {"name": "{}_device".format(d), "as_type": None}
        assert final_devices == config.devices.serialize(config.DICT)
        assert len(config.devices) == len(final_devices)


def test_device_ethernet(api):
    """
    configure the ethernet device and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    config.devices.clear()
    devices = config.devices
    device = devices.device()[0]
    eth = device.ethernet
    tv = get_expected_result()["devices"]
    dev_list = list()
    dev_list.append(tv["device"])
    dev_list[0].update({"choice": "ethernet", "ethernet": tv["ethernet"]})

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)
    
    choice_list = ["value", "increment", "decrement"]
    for c in choice_list:
        dev_list[0] = common_eth(eth, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)


def test_device_ipv4(api):
    """
    configure the ipv4 device and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    devices = config.devices
    devices.clear()

    dev_list = list()
    tv = get_expected_result()["devices"]
    ipv4_dev = devices.device()[0].ipv4
    dev_list.append(tv["device"])

    dev_list[0]["choice"] = "ipv4"
    dev_list[0]["ipv4"] = dict(tv["ipv4"])

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        dev_list[0] = common_ip(ipv4_dev, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    eth = ipv4_dev.ethernet
    for c in choice:
        dev_list[0]["ipv4"] = common_ip(eth, c, dev_list[0]["ipv4"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)



def test_device_ipv6(api):
    """
    configure the ipv6 device and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    devices = config.devices
    devices.clear()

    dev_list = list()
    tv = get_expected_result()["devices"]

    ipv6_dev = devices.device()[0].ipv6
    dev_list.append(tv["device"])
    dev_list[0]["choice"] = "ipv6"
    dev_list[0]["ipv6"] = dict(tv["ipv6"])

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        dev_list[0] = common_ip(ipv6_dev, c, dev_list[0], "ipv6")
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    eth = ipv6_dev.ethernet
    for c in choice:
        dev_list[0]["ipv6"] = common_ip(eth, c, dev_list[0]["ipv6"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)


def test_device_bgpv4(api):
    """
    configure the bgpv4 device and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    devices = config.devices
    devices.clear()

    dev_list = list()
    tv = get_expected_result()["devices"]

    bgp_dev = devices.device()[0].bgpv4
    dev_list.append(tv["device"])
    dev_list[0]["choice"] = "bgpv4"
    dev_list[0]["bgpv4"] = dict(tv["bgpv4"])

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        dev_list[0] = common_ip(bgp_dev, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    ipv4_dev = bgp_dev.ipv4
    for c in choice:
        dev_list[0]["bgpv4"] = common_ip(ipv4_dev, c, dev_list[0]["bgpv4"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    eth_dev = ipv4_dev.ethernet
    for c in choice:
        dev_list[0]["bgpv4"]["ipv4"] = common_ip(eth_dev, c, dev_list[0]["bgpv4"]["ipv4"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    dev_list[0]["bgpv4_route_ranges"] = []
    bgp_dev.bgpv4_route_ranges.bgpv4_route_range().bgpv4routerange()
    bgp_dev.bgpv6_route_ranges.bgpv6_route_range().bgpv6routerange()
    dev_list[0]["bgpv4_route_ranges"].append(tv["bgpv4routerange"])
    dev_list[0]["bgpv4_route_ranges"].append(tv["bgpv4routerange"])
    dev_list[0]["bgpv6_route_ranges"].append(tv["bgpv6routerange"])
    dev_list[0]["bgpv6_route_ranges"].append(tv["bgpv6routerange"])
    brt_name4 = "bgpv4_route_ranges"
    brt_name6 = "bgpv6_route_ranges"
    for c in choice:
        dev_list[0]["bgpv4"][brt_name4][0] = common_bgp_route_range(
            bgp_dev.bgpv4_route_ranges[0], c, dev_list[0]["bgpv4"][brt_name4][0]
        )
        dev_list[0]["bgpv4"][brt_name4][1] = common_bgp_route_range(
            bgp_dev.bgpv4_route_ranges[1], c, dev_list[0]["bgpv4"][brt_name4][1]
        )
        dev_list[0]["bgpv4"][brt_name6][0] = common_bgp_route_range(
            bgp_dev.bgpv4_route_ranges[0], c, dev_list[0]["bgpv4"][brt_name6][0]
        )
        dev_list[0]["bgpv4"][brt_name6][1] = common_bgp_route_range(
            bgp_dev.bgpv4_route_ranges[1], c, dev_list[0]["bgpv4"][brt_name6][1]
        )
        assert dev_list == config.devices.serialize(config.DICT)


def test_port(api):
    pass


def test_layer1(api):
    pass


def common_eth(eth_inst, choice, eth_dt):
    mac_addr = "10:10:10:00:00:00"
    mac_step = "10:10:10:00:00:01"
    mtu = 1234
    mtu_step = 1
    eth_dt["ethernet"] = {
        "name": None,
        "mac": {},
        "mtu": {}
    }
    if choice == "value":
        eth_inst.mac.value = mac_addr
        eth_inst.mtu.value = mtu
        eth_dt["ethernet"]["mac"]["choice"] = choice
        eth_dt["ethernet"]["mtu"]["choice"] = choice
        eth_dt["ethernet"]["mac"]["value"] = mac_addr
        eth_dt["ethernet"]["mtu"]["value"] = mtu
        return eth_dt
    if choice == "increment" or choice == "decrement":
        getattr(eth_inst.mac, choice).start = mac_addr
        getattr(eth_inst.mtu, choice).start = mtu
        getattr(eth_inst.mac, choice).step = mac_step
        getattr(eth_inst.mtu, choice).step = mtu_step
        eth_dt["ethernet"]["mac"]["choice"] = choice
        eth_dt["ethernet"]["mtu"]["choice"] = choice
        eth_dt["ethernet"]["mac"][choice] = {
            "start": mac_addr,
            "step": mac_step
        }
        eth_dt["ethernet"]["mtu"]["choice"] = choice
        eth_dt["ethernet"]["mtu"][choice] = {
            "start": mtu,
            "step": mtu_step
        }
        return eth_dt


def common_ip(ip_inst, choice, ip_dt, ip_ver="ipv4"):
    params = ["gateway", "address", "prefix"]
    if ip_ver == "ipv4":
        val = ["10.1.1.1", "10.1.1.2", "24"]
        val_step = ["0.1.0.0", "0.1.0.0", "1"]
    else:
        val = ["abcd::abcd", "1234::1234", "48"]
        val_step = ["::1", "::1", "1"]
    ip_dt[ip_ver] = {p : {} for p in params}
    ip_dt[ip_ver]["name"] = None
    if choice == "value":
        for i,p in enumerate(params):
            getattr(ip_inst, p).value = val[i]
            ip_dt[ip_ver][p]["choice"] = "value"
            ip_dt[ip_ver][p]["value"] = val[i]
        return ip_dt
    if choice == "increment" or choice == "decrement":
        for i,p in enumerate(params):
            getattr(getattr(ip_inst, p), choice).start = val[i]
            getattr(getattr(ip_inst, p), choice).start = val_step[i]
            ip_dt[ip_ver][p]["choice"] = choice
            ip_dt[ip_ver][p][choice] = {
                "start": val[i],
                "step": val_step[i]
            }
        return ip_dt


def common_bgp(bgp_inst, choice, bgp_dt, bgp_ver="bgpv4"):
    params = [
        "dut_as_number", "dut_ipv4_address", "hold_time_interval",
        "keep_alive_interval", "router_id"
    ]
    val = [
        1014, "10.1.1.1", 30, 30, "192.168.1.1"
    ]

    val_step = [
        1, "0.0.1.0", 5, 5, "0.0.1.1"
    ]

    bgp_dt[bgp_ver] = {p: {} for p in params}
    bgp_dt[bgp_ver]["name"] = None
    if choice == "value":
        for i,p in enumerate(params):
            getattr(bgp_inst, p).value = val[i]
            bgp_dt[bgp_ver][p]["choice"] = "value"
            bgp_dt[bgp_ver][p]["value"] = val[i]
        return bgp_dt
    if choice == "increment" or choice == "decrement":
        for i,p in enumerate(params):
            getattr(getattr(bgp_inst, p), choice).start = val[i]
            getattr(getattr(bgp_inst, p), choice).start = val_step[i]
            bgp_dt[bgp_ver][p]["choice"] = choice
            bgp_dt[bgp_ver][p][choice] = {
                "start": val[i],
                "step": val_step[i]
            }
        return bgp_dt


def common_bgp_route_range(bgp_range_inst, choice, bgp_range_dt, bgp_ver="bgpv4"):
    params = [
        "address", "as_path", "community", "next_hop_address",
        "prefix", "route_count_per_device"
    ]

    val = [
        "10.1.1.1" if bgp_ver=="bgpv4" else "22aa::22aa",
        900, "test", 
        "20.1.1.1" if bgp_ver=="bgpv4" else "33aa::22aa",
        16, 200
    ]
    val_step = [
        "0.1.1.0" if bgp_ver=="bgpv4" else "::1",
        10, "1",
        "0.1.0.1" if bgp_ver=="bgpv4" else "::1",
        10, 2
    ]
    for i,p in enumerate(params):
        bgp_range_dt[p] = {}
        if choice == "value":
            getattr(bgp_range_inst, p).value = val[i]
            bgp_range_dt[p]["choice"] = "value"
            bgp_range_dt[p]["value"] = val[i]
        if choice == "increment" or choice == "decrement":
            getattr(getattr(bgp_range_inst, p), choice).start = val[i]
            getattr(getattr(bgp_range_inst, p), choice).start = val_step[i]
            bgp_range_dt[p]["choice"] = choice
            bgp_range_dt[p][choice] = {
                "start": val[i],
                "step": val_step[i]
            }
    return bgp_range_dt


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
                "name": n[for_test_type],
                "choice": None
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
            "bgpv4routerange": {
                "name": None,
                "route_count_per_device": 1
            },
            "bgpv6routerange": {
                "name": None,
                "route_count_per_device": 1
            }
        }
    }
    return test_validation
