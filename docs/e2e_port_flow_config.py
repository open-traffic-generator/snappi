import snappiserver
snappiserver.SnappiServer().start()


# # python type hinting
# from scapy.all import *
# pkt = Ether() / IP() / TCP()
# ether = pkt[Ether]  # type: scapy.all.packet
# print(ether.src)
# # alternate separate package called 'hinty' 


# TBD: put all classes into one autogenerated file snappi.py
# TBD: import all classes in __init__.py
import snappi

api = snappi.api()  # type: snappi.Api

config = api.config()

tx_port, rx_port = config.ports \
    .port(name='Tx Port', location='10.36.74.26;02;13') \
    .port(name='Rx Port', location='10.36.74.26;02;14')

flow = config.flows.flow(name='Tx -> Rx Flow')[-1]
flow.tx_rx.port.tx_name = tx_port.name
flow.tx_rx.port.rx_name = rx_port.name
flow.size.fixed = 128
flow.rate.pps = 1000
flow.duration.fixed_packets.packets = 10000

# test to demonstrate SnappiList unpacking
flow.packet.ethernet().vlan().ipv4().tcp()

eth = flow.packet[0]  # type: snappi.FlowEthernet
ip = flow.packet[2]  # type: snappi.FlowIpv4

# TBD: eth = flow.packet[FlowEthernet][0]
eth.src.value = '00:00:01:00:00:01'
eth.dst.values = ['00:00:02:00:00:01', '00:00:02:00:00:01']

ip.src.value = '1.1.1.1'
ip.src.increment.start = '1.1.1.1'
ip.src.increment.step = '0.0.0.1'
ip.src.increment.count = 10
ip.src.decrement.start = ''
print(ip)
ip.dst.decrement.start = '1.1.2.200'
ip.dst.decrement.step = '0.0.0.1'
ip.dst.decrement.count = 10

# constants
ip.priority.dscp.ecn.value = ip.priority.dscp.ECN_CAPABLE_TRANSPORT_1

# ingress metric tracking
ip.priority.dscp.ecn.metric_group = 'ip.priority.dscp.ecn'

# egress metric model proposal - add sample tests to demonstrate complexity and results
# _, _, e_vlan, e_ip = flow.incoming_packets.ethernet().vlan().vlan().ip()
# e_vlan.id.bits = 3
# e_vlan.id.metric_group = 'egress eth.vlan.vlan'
# e_ip.src.bits
# e_ip.src.metric_group = 'asdf'

# capture model proposal
# _, _, c_vlan = config.captures.capture(port_names=[tx_port.name]).ethernet().vlan().vlan()
# c_vlan.id.filter_value = 3
# c_vlan.id.filter_mask = 'ff'

# set and get the configuration
api.set_config(config)
print(api.get_config())

# start transmit
transmit_state = api.transmit_state()
transmit_state.state = 'start'
api.set_transmit_state(transmit_state)

# get port metrics
port_metrics_request = api.port_metrics_request()
port_metrics = api.get_port_metrics(port_metrics_request)
print(sum([metric.frames_tx for metric in port_metrics]))

# get flow metrics
flow_metrics_request = api.flow_metrics_request()
flow_metrics_request.metric_group_names = ''
flow_metrics = api.get_flow_metrics(flow_metrics_request)
print(sum(metric.frames_tx for metric in flow_metrics))

