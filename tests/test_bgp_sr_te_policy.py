import pytest


def test_bgp_sr_te_policy(api):
    """Test BGP SR TE Policy functionality
    """
    config = api.config()

    # setup port container
    p1 = config.ports.port(name='p1')[-1]

    # setup device container
    d = config.devices.device(name='d', container_name=p1.name)[-1]

    # setup ethernet
    eth = d.ethernet
    eth.name = 'e'
    eth.mac = '00:01:00:00:00:01'

    # setup ipv6
    ip = eth.ipv6
    ip.name = 'i6'
    ip.address = '2a00:1450:f013:c03:8402:0:0:2'
    ip.gateway = '2a00:1450:f013:c03:0:0:0:1'
    ip.prefix = 64

    # setup bgp basic
    bgp = ip.bgpv6
    bgp.name = 'b6'
    bgp.router_id = '193.0.0.1'
    bgp.as_number = 65511
    bgp.as_number_set_mode = bgp.DO_NOT_INCLUDE_AS
    bgp.local_address = '2a00:1450:f013:c03:8402:0:0:2'
    bgp.dut_address = '2001:4860:0:0:0:1c:4001:ec2'

    # setup bgp advanced
    bgp.advanced.hold_time_interval = 90
    bgp.advanced.keep_alive_interval = 30

    # setup bgp sr te policy
    for i in range(1, 501):
        policy = bgp.sr_te_policies.bgpsrtepolicy()[-1]
        policy.policy_type = policy.IPV4
        policy.distinguisher = 1
        policy.color = i
        policy.ipv6_endpoint = '0:0:0:0:0:0:0:0'

        hop = policy.next_hop
        hop.next_hop_mode = hop.MANUAL
        hop.next_hop_address_type = hop.IPV6
        hop.ipv6_address = '2a00:1450:f013:c07:8402:0:0:2'

        # setup tunnel tlv
        tunnel = policy.tunnel_tlvs.bgptunneltlv(active=True)[-1]

        # setup tunnel tlv segment lists
        seglist = tunnel.segment_lists.bgpsegmentlist(active=True)[-1]
        seglist.segment_weight = 1

        # setup segment list segments
        for label in [1018001, 432999, 1048333, 1048561, 432001]:
            seg = seglist.segments.bgpsegment(active=True)[-1]
            seg.segment_type = seg.MPLS_SID
            seg.mpls_label = label

        # setup preference sub tlv
        pref_sub_tlv = tunnel.preference_sub_tlv
        pref_sub_tlv.preference = 400
    
        # setup binding sub tlv
        bind_sub_tlv = tunnel.binding_sub_tlv
        bind_sub_tlv.binding_sid_type = bind_sub_tlv.FOUR_OCTET_SID
        bind_sub_tlv.four_octet_sid = 483001
        bind_sub_tlv.bsid_as_mpls_label = True
        bind_sub_tlv.s_flag = False
        bind_sub_tlv.i_flag = False

        # setup explicit null label policy sub tlv
        enlp_sub_tlv = tunnel.explicit_null_label_policy_sub_tlv
        enlp_sub_tlv.explicit_null_label_policy = 2

        # setup bgpv4 route range
        v4rr = bgp.bgpv4_routes.bgpv4route(name='v4rr')[-1]
        v4rr.addresses.bgpv4routeaddress(address='4.4.4.4', prefix=32, count=5000, step=1)
        v4rr.as_path.as_set_mode = v4rr.as_path.INCLUDE_AS_SEQ
        v4rr.as_path.as_path_segments.bgpaspathsegment(as_numbers=[1, 2, 3, 4, 5, 6])


if __name__ == '__main__':
    pytest.main(['-sv', __file__])
