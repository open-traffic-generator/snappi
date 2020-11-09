# ***snappi***
Create test case scripts once using the `snappi` client and run them using a 
traffic generator that conforms to the [Open Traffic Generator API](https://github.com/open-traffic-generator/models/releases).

## *Install the client*
```
pip install snappi
```

## *Start scripting*
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

config = api.control_state.config_state.config
config.ports.append(name='Tx Port', location='10.36.74.26;02;13')
config.ports.append(name='Rx Port', location='10.36.74.26;02;14')
flow = config.flows.append(name='Tx -> Rx Flow')
flow.tx_rx.port_tx_rx.set(tx_port_name=config.ports[0].name,
                          rx_port_name=config.ports[1].name)
flow.packet.append(choice=snappi.FLOWHEADER_ETHERNET) \
    .append(choice=snappi.FLOWHEADER_VLAN) \
    .append(choice=snappi.FLOWHEADER_IPV4) \
    .append(choice=snappi.FLOWHEADER_TCP)
flow.size.fixed = 128
flow.duration.fixed_packets = 10000
flow.rate.pps = 1000

api.control_state.config_state.state = snappi.CONTROLCONFIGSTATE_SET
api.set_state(api.control_state)

api.control_state.flow_transmit_state.state = snappi.FLOWTRANSMITSTATE_START
api.set_state(api.control_state)

while True:
    results = api.get_port_results(api.result_portrequest)
    df = pandas.DataFrame.from_dict(results)
    print(df)
    if df.frames_tx.sum() >= config.flows[0].duration.fixed_packets:
        break
```