import snappi
import pandas
import dpkt

# create the top level api object
api = snappi.Api()

# get the configuration object using a property accessor
config = api.config

# add ports to config ports list using factory method
tx_port = config.ports.port(name='Tx Port', location='10.36.74.26;02;13')
rx_port = config.ports.port(name='Rx Port', location='10.36.74.26;02;14')

# add capture to config captures list using factory method
capture = api.captures.capture(name='src capture', port_names=[rx_port.name])

# configure capture to filter all source addresses
capture.mac_address_filter.mac=capture.mac_address_filter.SOURCE
capture.mac_address_filter.filter='00:00:00:00:00:00'

# add flow to config flows list using factory method
flow = config.flows.flow(name='Tx -> Rx Flow')

# add a source port and destinatinon port to the flow
flow.tx_rx.port.tx_name = tx_port.name
flow.tx_rx.port.rx_name = rx_port.name

# add flow packet headers to the flow packet using factory methods
# and also demonstrating the unpacking of headers to relevant variables
eth, _, _, _ = flow.packet.ethernet().vlan().ipv4().tcp()

# set ethernet packet header field
eth.src.counter.start = '00:00:01:00:00:01'
eth.src.counter.step = '00:00:00:00:00:01'
eth.src.counter.count = 10
eth.src.result_group = 'eth_src'

# set flow packet size, flow rate and flow duration
flow.size.fixed = 128
flow.rate.pps = 1000
flow.duration.fixed_packets.packets = 10000

# set the configuration
api.set_config(api.config)

# start capture on all configured ports, by default if no capture names are 
# provided all captures are started which is covered in the model documentation
# to start capture on a subset of ports provide the subset of port names to the
# ports parameter
api.capture_state.ports=[]
api.capture_state.state=api.capture_state.START
api.set_capture_state(api.capture_state)

# start transmit on all configured flows
# to start transmit on a subset of flows provide the subset of flow names to the
# flows parameter
api.flow_state.flows=[]
api.flow_state=api.flow_state.START
api.set_state(api.flow_state)

# get results to determine when transmit has stopped
api.result_portrequest.columns = [
    api.result_portrequest.FRAMES_TX, api.result_portrequest.FRAMES_RX
]
while True:
    results = api.get_port_results(api.result_portrequest)
    df = pandas.DataFrame.from_dict(results)
    df_col = df[api.result_portrequest.FRAMES_TX]
    if df_col.sum() >= flow.duration.fixed_packets:
        break

# stop the capture on a port and receive the capture as a stream of bytes
api.capture_results.port_name = rx_port.name
pcap_bytes = api.get_capture_results(api.capture_results)

# write the pcap bytes to a local file
pcap_filename = '%s.pcap' % rx_port.name
with open(pcap_filename, 'wb') as fid:
    fid.write(pcap_bytes)

# process the pcap file
for ts, pkt in dpkt.pcap.Reader(open(pcap_filename, 'rb')):
    print(ts, pkt)

