"""Send 10000 packets of tcp traffic on a single port"""
import snappi
import pandas

api = snappi.Api()

config = api.config
port = config.ports.port(name='Tx Port', location='10.36.74.26;02;13')
flow = config.flows.flow(name='Tcp Flow')
flow.tx_rx.port(tx_name=port.name)
_, _, ip, _ = flow.packet.ethernet().vlan().ipv4().tcp()
ip.src.fixed = '1.1.1.2'    # the implementation must handle conversion
ip.dst.fixed = 16843010     # the implementation must handle conversion
flow.duration.fixed_packets.packets = 10000
api.set_config(api.config)

api.flow_transmit_state.state = api.flow_transmit_state.START
api.set_transmit_state(api.flow_transmit_state)

while True:
    results = api.get_port_results(api.result_portrequest)
    df = pandas.DataFrame.from_dict(results)
    if df.frames_tx.sum() >= config.flows[0].duration.fixed_packets:
        break
