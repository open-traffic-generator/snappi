# snappi
The open traffic generator client.

## Install the client
```
pip install snappi
```

## Start scripting
```python
"""A simple test that demonstrates the following:
- A port that transmits an ethernet/vlan/ipv4/tcp flow 
for a specified duration and a port that receives the packets.
- While the flow is transmitting the script prints out tx and rx statistics.
- Once all the packets have been transmitted the script will end. 
"""
import snappi
import pandas

api = snappi.Api()
control_state = api.control_state.set(choice='config_state')
config = control_state.config_state.set(state='set').config
config.ports.append(name='Tx Port', location='10.36.74.26;02;13')
config.ports.append(name='Rx Port', location='10.36.74.26;02;14')
flow = config.flows.append(name='Tx -> Rx Flow')
flow.tx_rx.set(choice='port').port.set(tx_port_name=config.ports[0].name,
                                       rx_port_name=config.ports[1].name)

# use flowheader.choice and create property accessors
flow.packet.ethernet().vlan().ipv4().tcp() 

flow.packet.append(choice='ethernet').append(choice='vlan').append(
    choice='ipv4').append(choice='tcp')

flow.packet[0].ethernet

flow.packet.append().ethernet
flow.packet.append().ipv4
flow.packet.append().tcp

flow.packet = [ixsnappi.flowethernet.FlowEthernet(), FlowIpv4()]

flow.size.set(choice='fixed').fixed = 128
flow.duration.set(choice='packets').packets.set(packets=100000)
flow.rate.set(units='pps', value=1000)
api.set_state(api.control_state)

control_state.set(choice='flow_transmit_state')
control_state.flow_transmit_state.state = 'start'
api.set_state(control_state)

request = result.PortRequest(
    column_names=['name', 'location', 'frames_tx', 'frames_rx'])
while True:
    results = api.get_port_results(request)
    df = pandas.DataFrame.from_dict(results)
    print(df)
    if df.frames_tx.sum() >= config.flows[0].duration.packets.packets:
        break
```