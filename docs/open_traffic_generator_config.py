from abstract_open_traffic_generator.port import port, flow, config, control, result
from ixnetwork_open_traffic_generator.ixnetworkapi import IxNetworkApi
import pandas

tx_port = port.Port(name='Tx Port', location='10.36.74.26;02;13')
rx_port = port.Port(name='Rx Port', location='10.36.74.26;02;14')

packet = [
    flow.Header(flow.Ethernet()),
    flow.Header(flow.Vlan()),
    flow.Header(flow.Ipv4()),
    flow.Header(flow.Tcp())
]
eth = packet[0].ethernet
eth.src = flow.Pattern(flow.Fixed('00:00:00:00:00:01'),
                       result_group='eth_src_mac_addr')
eth.dst = flow.Pattern(
    flow.ValueList(['00:00:00:00:00:02', '00:00:00:00:00:03']))
ip = packet[2].ipv4
ip.src = flow.Pattern(
    flow.Counter(start='1.1.1.1', step='0.0.0.1', count=10, up=True))
ip.src = flow.Pattern(
    flow.Counter(start='1.1.2.1', step='0.0.0.1', count=10, up=True))

flow = flow.Flow(name='Tx Port to Rx Port Flow',
                 tx_rx=flow.TxRx(
                     flow.PortTxRx(tx_port_name=tx_port.name,
                                   rx_port_name=rx_port.name)),
                 packet=packet,
                 size=flow.Size(128),
                 rate=flow.Rate(unit='pps', value=1000),
                 duration=flow.Duration(flow.FixedPackets(10000)))
config = config.Config(ports=[tx_port, rx_port],
                       flows=[flow],
                       options=config.Options(port_options=port.Options(
                           location_preemption=True)))

api = IxNetworkApi(address='10.36.66.49', port=11009)
api.set_state(control.State(control.ConfigState(config=config, state='set')))
state = control.State(control.FlowTransmitState(state='start'))
api.set_state(state)

request = result.PortRequest(
    column_names=['name', 'location', 'frames_tx', 'frames_rx'])
while True:
    results = api.get_port_results(request)
    df = pandas.DataFrame.from_dict(results)
    print(df)
    if df.frames_tx.sum() >= config.flows[0].duration.packets.packets:
        break
