# ![snappi](logo/logo_color.png)

Create test case scripts once using the `snappi` client and run them using a 
traffic generator that conforms to the [Open Traffic Generator API](https://github.com/open-traffic-generator/models/releases).

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

api = snappi.Api()
config = api.config()
tx_port, rx_port = config.ports \
    .port(name='Tx Port', location='10.36.74.26;02;13')
    .port(name='Rx Port', location='10.36.74.26;02;14')
flow = config.flows.flow(name='Tx -> Rx Flow')[-1]
flow.tx_rx.port.tx_name = tx_port.name
flow.tx_rx.port.tx_name = rx_port.name
flow.packet.ethernet().vlan().ipv4().tcp()
flow.size.fixed = 128
flow.rate.pps = 1000
flow.duration.fixed_packets.packets = 10000
api.set_config(config)

tx_state = api.transmit_state(state='start')
api.set_transmit(tx_state)

while True:
    metrics = api.get_port_metrics(api.port_metrics_request())
    tx_frames = sum([metric.frames_tx for metric in metrics])
    if tx_frames >= flow.duration.fixed_packets.packets:
        break
```
