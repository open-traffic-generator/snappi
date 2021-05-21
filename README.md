# ![snappi](snappi-logo.png)

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Build](https://github.com/open-traffic-generator/snappi/workflows/Build/badge.svg)](https://github.com/open-traffic-generator/snappi/actions)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/alerts/)
[![Language grade: Python](https://img.shields.io/lgtm/grade/python/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/context:python)
[![pypi](https://img.shields.io/pypi/v/snappi.svg)](https://pypi.org/project/snappi)
[![python](https://img.shields.io/pypi/pyversions/snappi.svg)](https://pypi.python.org/pypi/snappi)

Test scripts written in `snappi`, an auto-generated python SDK, can be executed against any traffic generator conforming to [Open Traffic Generator API](https://github.com/open-traffic-generator/models).

[Ixia-c](https://github.com/open-traffic-generator/ixia-c) is one such reference implementation of Open Traffic Generator API.

> The repository is under active development and is subject to updates. All efforts will be made to keep the updates backwards compatible.

## Install on a client

```sh
python -m pip install --upgrade snappi
```

## Start scripting

```python
import snappi
# create a new API instance where host points to controller
api = snappi.api(host='https://localhost')

# create a config object to be pushed to controller
config = api.config()
# add a port with location pointing to traffic engine
prt = config.ports.port(name='prt', location='localhost:5555')[-1]
# add a flow and assign endpoints
flw = config.flows.flow(name='flw')[-1]
flw.tx_rx.port.tx_name = prt.name

# configure 100 packets to be sent, each having a size of 128 bytes
flw.size.fixed = 128
flw.duration.fixed_packets.packets = 100

# add Ethernet, IP and TCP protocol headers with defaults
flw.packet.ethernet().ipv4().tcp()

# push configuration
api.set_config(config)

# start transmitting configured flows
ts = api.transmit_state()
ts.state = ts.START
api.set_transmit_state(ts)

# fetch & print port metrics
req = api.metrics_request()
req.port.port_names = [prt.name]
print(api.get_metrics(req))
```
