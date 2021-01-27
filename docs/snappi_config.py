# single import
import snappi
import re


# create the top level api object
api = snappi.api()

# create a top level config object
config = api.config()

# add ports to config ports list using factory method
tx_port = config.ports.port(name='Tx Port', location='10.36.74.26;02;13')
rx_port = config.ports.port(name='Rx Port', location='10.36.74.26;02;14')

# add flow to config flows list using factory method
flow = config.flows.flow(name='Tx -> Rx Flow')

# add a source port and destination port to the flow
flow.tx_rx.port.tx_name = tx_port.name
flow.tx_rx.port.rx_name = rx_port.name

# add flow packet headers to the flow packet using factory methods
pkt = flow.packet.ethernet().vlan().ipv4().tcp()

# different ways to unpack packet headers
eth, _, _, _ = flow.packet
eth = flow.packet[0]
eth = flow.packet[flow.packet.ETHERNET]
ip = pkt[flow.packet.IPV4]

# set flow header field patterns
eth.src = '00:00:01:00:00:01'
eth.dst = ['00:00:02:00:00:01', '00:00:02:00:00:01']
ip.src.increment(start='1.1.1.1', step='0.0.0.1', count=10)
ip.src.increment(start='1.1.1.1', step=1, count=10)
ip.dst.decrement(start='1.1.2.200', step='0.0.0.1', count=10)
# x-type int, model min/max 0-65535, implementation may or may not support full max
# for best ux - implementation should preprocess config prior to setting anything 
# and throw error - do a dry run 
tcp.src_port.increment(start='1.1.1.1') 
tcp.src_port.increment(start='1')
tcp.src_port.increment(start=1)

# use constants
ip.priority.dscp.phb = ip.priority.dscp.PHB_CS1

# set flow packet size
flow.size.fixed = 128

# set the flow rate
flow.rate.pps = 1000

# set the flow duration
flow.duration.fixed_packets = 10000

# TBD
# specify that a packet header field can be viewed as a
# column name in the flow results using a user supplied column name
eth.src.result_group = 'eth_src_mac_addr'

# set the configuration
api.set_config(config)

# start transmit on all configured flows
flow_state = api.flow_state(flow_state=api.flow_state.START)
api.set_flow_state(flow_state)

# get all flows, returns a list of rows, each row is a flow 
query = api.flow_query(flow_names=[], metric_names=[], metric_match=[{'state': 'stopped'}])
flow_metrics = api.get_flow_metrics(query)
for flow_metric in flow_metrics:
    print(flow_metric.name, flow_metric.state, flow_metric.tx_port, flow_metric.rx_port, flow_metric.frames_tx)


# get results to determine when transmit has stopped
# extension to the SNAPPI package without overhead to the models
# in the same namespace
while len(config.flows) != len([
        m for m in (result.state == api.flow_metric.STOPPED)
                    for result in api.get_port_results(api.result_portrequest()))
]):
    continue
