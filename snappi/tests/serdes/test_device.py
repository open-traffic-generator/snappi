from py import path
import pytest


def validate_args(node, params, expected_dt):
    for p in params:
        assert getattr(node, p) == expected_dt[p]


def validate_test(node, dt, default=None):
    if default is None:
        default = {k: None for k in dt}
    validate_args(node, dt.keys(), default)
    for k in dt:
        setattr(node, k, dt[k])
    validate_args(node, dt.keys(), dt)
    assert node.serialize(node.DICT) == dt


def test_device(api):
    devices = {
        'devices': [
            {
                'container_name': 'Port1',
                'name': 'myDevice1'
            },
            {
                'container_name': 'Port2',
                'name': 'myDevice2'
            }
        ]
    }
    config = api.config()
    config.devices.device(
        name='myDevice1', container_name='Port1'
    ).device(
        name='myDevice2', container_name='Port2'
    )
    assert config.serialize(config.DICT) == devices
    config = api.config()
    config.deserialize(devices)
    assert len(config.devices) == len(devices['devices'])
    validate_args(config.devices[0], ['name', 'container_name'], devices['devices'][0])
    validate_args(config.devices[1], ['name', 'container_name'], devices['devices'][1])


def test_ethernet(api):
    eth = {
        'mac': 'abcd',
        'mtu': 10,
        'name': 'EthDevice'
    }
    config = api.config()
    d1 = config.devices.device()[-1]
    ethernet = d1.ethernet
    validate_test(ethernet, eth)
    d2 = config.devices.device()[-1]
    d2.ethernet.deserialize(eth)
    validate_args(d2.ethernet, ['name', 'mac', 'mtu'], eth)


def test_ipv4(api):
    ipv4 =  {
        'name': 'ipDevice',
        'address': '10.1.1.1',
        'gateway': '10.1.1.2',
        'prefix': '24'
    }
    config = api.config()
    d1 = config.devices.device()[-1]
    ip = d1.ethernet.ipv4
    validate_test(ip, ipv4)
    config.devices.device()[-1].ethernet.ipv4.deserialize(ipv4)
    validate_args(config.devices[-1].ethernet.ipv4, ['name', 'address', 'gateway'], ipv4)


def test_ipv6(api):    
    ipv6 =  {
        'name': 'ipDevice',
        'address': 'abcd::abcd',
        'gateway': 'abcd::abcd',
        'prefix': '48'
    }
    config = api.config()
    d1 = config.devices.device()[-1]
    ip = d1.ethernet.ipv4
    validate_test(ip, ipv6)
    config.devices.device()[-1].ethernet.ipv6.deserialize(ipv6)
    validate_args(config.devices[-1].ethernet.ipv6, ['name', 'address', 'gateway'], ipv6)


def test_bgpv4(api):
    bgpv4 = {
        "active": True,
        "as_number": 10,
        "as_number_set_mode": "do_not_include_as",
        "as_number_width": "two",
        "as_type": "ibgp",
        "dut_address": "10.1.1.1",
        "local_address": "10.1.1.2",
        "name": "bgpv4",
        "router_id": "192.168.1.1"
    }
    config = api.config()
    d1 = config.devices.device()[-1]
    bgp = d1.ethernet.ipv4.bgpv4
    validate_test(bgp, bgpv4)
    bgp = config.devices.device()[-1].ethernet.ipv4.bgpv4.deserialize(bgpv4)
    validate_args(bgp, bgpv4.keys(), bgpv4)


def test_bgpv6(api):
    bgpv6 = {
        "active": True,
        "as_number": 10,
        "as_number_set_mode": "do_not_include_as",
        "as_number_width": "two",
        "as_type": "ibgp",
        "dut_address": "abcd::abcd",
        "local_address": "abcd::abcd",
        "name": "bgpv6",
        "router_id": "abcd::abcd"
    }
    config = api.config()
    d1 = config.devices.device()[-1]
    bgp = d1.ethernet.ipv6.bgpv6
    validate_test(bgp, bgpv6)
    bgp = config.devices.device()[-1].ethernet.ipv6.bgpv6.deserialize(bgpv6)
    validate_args(bgp, bgpv6.keys(), bgpv6)


@pytest.mark.parametrize('ip_version', ['4', '6'])
def test_bgp_advanced(api, ip_version):
    adv = {
        "hold_time_interval": 90,
        "keep_alive_interval": 30,
        "md5_key": True,
        "time_to_live": 64,
        "update_interval": 0,
    }
    config = api.config()
    d = config.devices.device()[-1]
    d_adv = getattr(getattr(d.ethernet, 'ipv'+ ip_version), 'bgpv'+ ip_version).advanced
    validate_test(d_adv, adv)
    d = config.devices.device()[-1]
    d_adv = getattr(
        getattr(d.ethernet, 'ipv'+ ip_version), 'bgpv'+ ip_version).advanced.deserialize(adv)
    validate_args(d_adv, adv.keys(), adv)


@pytest.mark.parametrize('ip_version', ['4', '6'])
def test_bgp_srte(api, ip_version):
    srte_default = {
        "color": 100,
        "distinguisher": 1,
        "ipv4_endpoint": "0.0.0.0",
        "ipv6_endpoint": "::0",
        "policy_type": "ipv4"
    }
    srte = {
        "color": 200,
        "distinguisher": 2,
        "ipv4_endpoint": "1.0.1.1",
        "ipv6_endpoint": "::abcd",
        "policy_type": "ipv4"
    }
    config = api.config()
    d = config.devices.device()[-1]
    bgp = getattr(getattr(d.ethernet, 'ipv'+ ip_version), 'bgpv'+ ip_version)
    b_srte = bgp.sr_te_policies
    sr = b_srte.bgpsrtepolicy()[-1]
    validate_test(sr, srte, srte_default)
    sr1 = b_srte.bgpsrtepolicy()[-1].deserialize(srte)
    validate_args(sr1, srte.keys(), srte)
    assert len(b_srte) == 2


@pytest.mark.parametrize('ip_version', ['4', '6'])
def test_bgp_srte_next_hop(api, ip_version):
    n_hop = {
        "ipv4_address": "10.1.1.1",
        "ipv6_address": "abcd::abcd",
        "next_hop_address_type": "ipv4",
        "next_hop_mode": "manual"
    }
    config = api.config()
    d = config.devices.device()[-1]
    bgp = getattr(getattr(d.ethernet, 'ipv'+ ip_version), 'bgpv'+ ip_version)
    next_hop = bgp.sr_te_policies.bgpsrtepolicy()[-1].next_hop
    validate_test(next_hop, n_hop)
    next_hop = bgp.sr_te_policies.bgpsrtepolicy()[-1].next_hop.deserialize(n_hop)
    validate_args(next_hop, n_hop.keys(), n_hop)
    add_path = {
        "path_id": 10
    }
    path = bgp.sr_te_policies.bgpsrtepolicy()[-1].add_path
    validate_test(path, add_path)
    path = bgp.sr_te_policies.bgpsrtepolicy()[-1].add_path.deserialize(add_path)
    validate_args(path, add_path.keys(), add_path)
