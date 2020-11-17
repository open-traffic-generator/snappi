# snappi-test.py

# This script creates the traffic going over the N3 GTPu interface between
# the gNB and the UPF nodes of a 5G core network.
# Refer : https://medium.com/@derekcheung/5gc-part-3-user-plane-and-gtp-u-tunnels-53319463967
#
#
# N ports (representing users) and N ports (representing servers)
# M GTP tunnels originating and terminating on each port.  Each tunnel represents one UE (mobile phone)
# Assume that
#   N = M for the sake of simplicity
#   Total traffic required in each direction is 5 Gbps
#   Each port is able to send and receive 1 Gbps
#   --> Test will need 5 ports.  Each port will send traffic over 5 GTP tunnels
#   --> Each tunnel (hence each UE) will send at the rate of ~200Mbps (1Gbps/5)
#
#                   T[1]
#       P[1]  <------------------------> P[6]
#                   T[2]
#             <------------------------> P[7]
#                   T[3]
#             <------------------------> P[8]
#      ....
#                   T[5]
#             <------------------------> P[10]
#
#
#
#                   T[6]
#       P[2]  <------------------------> P[6]
#                   T[7]
#             <------------------------> P[7]
#      ....
#
#
#                   T[21]
#       P[5]  <------------------------> P[6]
#                   T[22]
#             <------------------------> P[7]
#      ....

# We represent this scenario using one flow on each port and
# use a list to represent the UE IP addresses and the GTP tunnel IDs.
# Packets have the following headers: Ethernet, IP, UDP, GTPv1, IP
# Frame size is 128

import snappi
import pyaml
import collections

# these test parameters would normally come from a test infrastructure file
d = pyaml.yaml.safe_load("""
# gnb_mac_addr This is the MAC SA for uplink packets & MAC DA for downlink packets
# n3_mac_addr This is the MAC DA for uplink packets & MAC SA for downlink packets
globals:
  gnb_mac_addr: AA BB CC DD EE FF 
  n3_mac_addr: FF EE DD CC BB AA
  gnb_ip_addr: 10.20.31.10
  upf_ip_addr: 10.20.31.3
  inner_ip_src_dst: 172.217.5.68
  gtp_tunnels: 5

users:
  - location: localuhd/1
    teid: 0
    ue_ip: 1.1.1.1
  - location: localuhd/2
    teid: 25
    ue_ip: 1.1.2.1

servers:
  - location: localuhd/1
    teid: 1000
    ue_ip: 1.1.100.1
  - location: localuhd/2
    teid: 1025
    ue_ip: 1.1.101.1
""")
test_params = collections.namedtuple('TestParams', d.keys())(*d.values())


def create_flow(config, flow_prefix, src_param, dst_param, tx_port, rx_port,
                src_mac, dst_mac, outer_ip_src, outer_ip_dst, inner_ip_src,
                inner_ip_dst, gtp_tunnels):
    flow = config.flows.flow(name='%s %s -> %s' %
                             (flow_prefix, tx_port.name, rx_port.name))
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name

    eth, outer_ip, _, gtp, inner_ip = flow.packet.ethernet().ipv4().udp().gtpv1().ipv4()

    eth.src.fixed = src_mac
    eth.dst.fixed = dst_mac
    outer_ip.src.fixed = outer_ip_src
    outer_ip.dst.fixed = outer_ip_dst
    gtp.teid.start = src_param.teid
    gtp.teid.step = 1
    gtp.teid.count = gtp_tunnels
    inner_ip.src.fixed = inner_ip_src
    inner_ip.dst.fixed = inner_ip_dst
    gtp.teid.result_group = 'gtp_teid'

    flow_uplink.size.fixed = 128
    flow_uplink.rate.gbps = 1
    flow_uplink.duration.fixed_packets.packets = 100000


api = snappi.Api()
config = api.config

for user, server in zip(test_params.users, test_params.servers):
    tx_port = config.ports.port(name='User Port %s' % user, location=user)
    rx_port = config.ports.port(name='Server Port %s' % server,
                                location=server)
    create_flow(config, 'Uplink', user, server, tx_port, rx_port,
                test_params.gnb_mac_addr, test_params.n3_mac_addr,
                test_params.gnb_ip_addr, test_params.upf_ip_addr, user.ue_ip,
                test_params.inner_ip_src_dst, test_params.gtp_tunnels)
    create_flow(config, 'Downlink', server, user, tx_port, rx_port,
                test_params.n3_mac_addr, test_params.gnb_mac_addr,
                test_params.upf_ip_addr, test_params.gnb_ip_addr,
                test_params.inner_ip_src_dst, server.ue_ip,
                test_params.gtp_tunnels)

api.set_config(api.config)
