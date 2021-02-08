def test_device_default(api):
    """
    configure the devices with defaults and
    construct the config dict
    validate the serialization against the 
    constructed dict
    """
    config = api.config()
    config.devices.clear()
    
    tv = get_expected_result()["devices"]
    final_devices = list()
    final_devices.append(dict(tv["device"]))
    dev_list = ["ethernet", "ipv4", "bgpv4"]
    dev = config.devices.device()[-1]
    dev_map = final_devices[0]
    for i,d in enumerate(dev_list):
        dev = getattr(dev, d)
        dev_map[d] = dict(tv[d])
        dev_map = dev_map[d]
        assert final_devices == config.devices.serialize(config.DICT)
        assert len(config.devices) == len(final_devices)
    
    dev_list = ["ethernet", "ipv6"]
    dev = config.devices.device()[-1]
    final_devices.append(dict(tv["device"]))
    dev_map = final_devices[-1]
    for i,d in enumerate(dev_list):
        dev = getattr(dev, d)
        dev_map[d] = dict(tv[d])
        dev_map = dev_map[d]
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
    dev_list[0].update({"ethernet": tv["ethernet"]})

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
    eth = devices.device()[0].ethernet
    ipv4_dev = eth.ipv4
    dev_list.append(tv["device"])

    # dev_list[0]["choice"] = "ipv4"
    dev_list[0]["ethernet"] = dict(tv["ethernet"])
    dev_list[0]["ethernet"]["ipv4"] = dict(tv["ipv4"])
    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        common_eth(eth, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)

    for c in choice:
        common_ip(ipv4_dev, c, dev_list[0]["ethernet"])
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
    eth = devices.device()[0].ethernet
    ipv6_dev = eth.ipv6
    dev_list.append(tv["device"])
    dev_list[0]["ethernet"] = dict(tv["ethernet"])
    dev_list[0]["ethernet"]["ipv6"] = dict(tv["ipv6"])

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        common_eth(eth, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)

    for c in choice:
        common_ip(ipv6_dev, c, dev_list[0]["ethernet"], "ipv6")
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

    eth = devices.device()[0].ethernet
    ipv4 = eth.ipv4
    bgpv4 = ipv4.bgpv4
    dev_list.append(tv["device"])
    dev_list[0]["ethernet"] = dict(tv["ethernet"])
    dev_list[0]["ethernet"]["ipv4"] = dict(tv["ipv4"])
    dev_list[0]["ethernet"]["ipv4"]["bgpv4"] = dict(tv["bgpv4"])
    bgp_dt = dev_list[0]["ethernet"]["ipv4"]["bgpv4"]

    assert dev_list == config.devices.serialize(config.DICT)
    assert len(dev_list) == len(config.devices)

    choice = ["value", "increment", "decrement"]
    for c in choice:
        common_eth(eth, c, dev_list[0])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)

    for c in choice:
        common_ip(ipv4, c, dev_list[0]["ethernet"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)

    for c in choice:
        common_bgp(bgpv4, c, dev_list[0]["ethernet"]["ipv4"])
        assert dev_list == config.devices.serialize(config.DICT)
        assert len(dev_list) == len(config.devices)
    bgp_dt["bgpv4_route_ranges"] = []
    bgp_dt["bgpv6_route_ranges"] = []
    bgpv4.bgpv4_route_ranges.bgpv4routerange().bgpv4routerange()
    bgpv4.bgpv6_route_ranges.bgpv6routerange().bgpv6routerange()
    bgp_dt["bgpv4_route_ranges"].append(tv["bgpv4routerange"])
    bgp_dt["bgpv4_route_ranges"].append(tv["bgpv4routerange"])
    bgp_dt["bgpv6_route_ranges"].append(tv["bgpv6routerange"])
    bgp_dt["bgpv6_route_ranges"].append(tv["bgpv6routerange"])
    brt_name4 = "bgpv4_route_ranges"
    brt_name6 = "bgpv6_route_ranges"
    for c in choice:
        common_bgp_route_range(
            bgpv4.bgpv4_route_ranges[0], c, bgp_dt[brt_name4][0]
        )
        common_bgp_route_range(
            bgpv4.bgpv4_route_ranges[1], c, bgp_dt[brt_name4][1]
        )
        common_bgp_route_range(
            bgpv4.bgpv6_route_ranges[0], c, bgp_dt[brt_name6][0]
        )
        common_bgp_route_range(
            bgpv4.bgpv6_route_ranges[1], c, bgp_dt[brt_name6][1]
        )
        assert dev_list == config.devices.serialize(config.DICT)


def test_port(api):
    config = api.config()
    config.ports.clear()
    port = config.ports.port()
    port_list = [dict(get_expected_result()["ports"])]
    assert len(config.ports) == 1
    assert port_list == config.ports.serialize(config.DICT)
    port[-1].name, port_list[0]["name"] = "test", "test"
    port[-1].location, port_list[0]["location"] = \
    "location_string", "location_string"
    assert port_list == config.ports.serialize(config.DICT)
    config.ports.clear()
    assert len(config.ports) == 0
    

def test_layer1(api):
    config = api.config()
    config.layer1.clear()
    l = config.layer1.layer1()[-1]
    l_list = dict(get_expected_result()["layer1"])
    assert len(config.layer1) == 1
    assert [l_list] == config.layer1.serialize(config.DICT)
    l.auto_negotiate, l_list['auto_negotiate'] = False, False
    l.ieee_media_defaults, l_list['ieee_media_defaults'] = False, False
    l.media, l_list['media'] = 'fiber', 'fiber'
    l.mtu, l_list['mtu'] = 1500, 1500
    l.port_names, l_list['port_names'] = [], []
    l.promiscuous, l_list['promiscuous'] = True, True
    l.speed, l_list['speed'] = 'test', 'test'
    assert [l_list] == config.layer1.serialize(config.DICT)
    config.layer1.clear()
    assert len(config.layer1) == 0




def common_eth(eth_inst, choice, eth_dt):
    mac_addr = "10:10:10:00:00:00"
    mac_step = "10:10:10:00:00:01"
    mtu = 1234
    mtu_step = 1
    eth_dt["ethernet"].update({
        "name": None
        if eth_dt["ethernet"].get("name") is None else eth_dt["ethernet"]["name"],
        "mac": {} 
        if eth_dt["ethernet"].get("mac") is None else eth_dt["ethernet"]["mac"],
        "mtu": {}
        if eth_dt["ethernet"].get("mtu") is None else eth_dt["ethernet"]["mtu"],
    })
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
    ip_dt[ip_ver].update(
        {p : {} 
        if ip_dt[ip_ver].get(p) is None else ip_dt[ip_ver].get(p) 
        for p in params}
    )
    if choice == "value":
        for i,p in enumerate(params):
            getattr(ip_inst, p).value = val[i]
            ip_dt[ip_ver][p]["choice"] = "value"
            ip_dt[ip_ver][p]["value"] = val[i]
        return ip_dt
    if choice == "increment" or choice == "decrement":
        for i,p in enumerate(params):
            getattr(getattr(ip_inst, p), choice).start = val[i]
            getattr(getattr(ip_inst, p), choice).step = val_step[i]
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

    bgp_dt[bgp_ver].update({
        p : {} 
        if bgp_dt[bgp_ver].get(p) is None else bgp_dt[bgp_ver].get(p) 
        for p in params
    })
    # bgp_dt[bgp_ver]["name"] = None
    if choice == "value":
        for i,p in enumerate(params):
            getattr(bgp_inst, p).value = val[i]
            bgp_dt[bgp_ver][p]["choice"] = "value"
            bgp_dt[bgp_ver][p]["value"] = val[i]
        return bgp_dt
    if choice == "increment" or choice == "decrement":
        for i,p in enumerate(params):
            getattr(getattr(bgp_inst, p), choice).start = val[i]
            getattr(getattr(bgp_inst, p), choice).step = val_step[i]
            bgp_dt[bgp_ver][p]["choice"] = choice
            bgp_dt[bgp_ver][p][choice] = {
                "start": val[i],
                "step": val_step[i]
            }
        return bgp_dt


def common_bgp_route_range(bgp_range_inst, choice, bgp_range_dt, bgp_ver="bgpv4"):
    params = [
        "address", "address_step", "next_hop_address", "prefix"
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
        bgp_range_dt[p] = {} \
            if bgp_range_dt.get(p) is None else bgp_range_dt.get(p)
        if choice == "value":
            getattr(bgp_range_inst, p).value = val[i]
            bgp_range_dt[p]["choice"] = "value"
            bgp_range_dt[p]["value"] = val[i]
        if choice == "increment" or choice == "decrement":
            getattr(getattr(bgp_range_inst, p), choice).start = val[i]
            getattr(getattr(bgp_range_inst, p), choice).step = val_step[i]
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
        "invalid": (int, list, dict, bool)
    }

    dc = {
        "default": 1,
        "valid": [int],
        "invalid": (str, list, dict, bool)
    }

    dc_1 = {
        "default": None,
        "valid": (int, None),
        "invalid": (str, list, dict, bool)
    }
    b = {
        "default": True,
        "valid": [bool],
        "invalid": (str, list, dict, int)

    }
    b1 = {
        "default": False,
        "valid": [bool],
        "invalid": (str, list, dict, int)

    }
    lst = {
        "default": None,
        "valid": (list, None),
        "invalid": (str, int, dict, bool)
    }

    test_validation = {
        "layer1": {
            "auto_negotiate": b[for_test_type],
            "ieee_media_defaults": b[for_test_type],
            "media": n[for_test_type],
            "mtu": 1500,
            "name": n[for_test_type],
            "port_names": lst[for_test_type],
            "promiscuous": b1[for_test_type],
            "speed": "speed_10_gbps"
        },
        "ports": {
            "location": n[for_test_type],
            "name": n[for_test_type]
        },
        "devices": {
            "device": {
                "container_name": n[for_test_type],
                "device_count": dc[for_test_type],
                "name": n[for_test_type],
                # "choice": None
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
                "range_count": 1,
                "address_count": 1
            },
            "bgpv6routerange": {
                "name": None,
                "range_count": 1,
                "address_count": 1
            }
        }
    }
    return test_validation
