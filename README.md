# ![snappi](snappi-logo.png)

[![license](https://img.shields.io/badge/license-MIT-green.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Build](https://github.com/open-traffic-generator/snappi/workflows/Build/badge.svg)](https://github.com/open-traffic-generator/snappi/actions)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/alerts/)
[![Language grade: Python](https://img.shields.io/lgtm/grade/python/g/open-traffic-generator/snappi.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/open-traffic-generator/snappi/context:python)
[![pypi](https://img.shields.io/pypi/v/snappi.svg)](https://pypi.org/project/snappi)
[![python](https://img.shields.io/pypi/pyversions/snappi.svg)](https://pypi.python.org/pypi/snappi)

Test scripts written using `snappi` (a python library) can be executed against  
any traffic generator conforming to [Open Traffic Generator API](https://github.com/open-traffic-generator/models/releases).

> The repository is under active development and is subject to updates. All efforts will be made to keep the updates backwards compatible.

## Install on a client

```sh
python -m pip install --upgrade snappi
# or install along with ixnetwork extension
python -m pip install --upgrade "snappi[ixnetwork]"
```

## Start scripting

```python
"""
Configure a raw TCP flow with,
- tx port as source to rx port as destination
- frame count 10000, each of size 128 bytes
- transmit rate of 1000 packets per second
Validate,
- frames transmitted and received for configured flow is as expected
"""

import snappi
# when using ixnetwork extension, host is IxNetwork API Server
api = snappi.api(host='https://localhost:443', ext='ixnetwork')
# new config
config = api.config()
# when using ixnetwork extension, port location is chassis-ip;card-id;port-id
tx, rx = (
    config.ports
    .port(name='tx', location='192.168.0.1;2;1')
    .port(name='rx', location='192.168.0.1;2;2')
)
# configure layer 1 properties
ly, = config.layer1.layer1(name='ly')
ly.port_names = [tx.name, rx.name]
ly.speed = ly.SPEED_10_GBPS
ly.media = ly.FIBER
# configure flow properties
flw, = config.flows.flow(name='flw')
# flow endpoints
flw.tx_rx.port.tx_name = tx.name
flw.tx_rx.port.rx_name = rx.name
# configure rate, size, frame count
flw.size.fixed = 128
flw.rate.pps = 1000
flw.duration.fixed_packets.packets = 10000
# configure protocol headers with defaults fields
flw.packet.ethernet().vlan().ipv4().tcp()
# push configuration
api.set_config(config)
# start transmitting configured flows
ts = api.transmit_state()
ts.state = ts.START
api.set_transmit_state(ts)
# create a query for flow metrics
req = api.metrics_request()
req.flow.flow_names = [flw.name]
# wait for flow metrics to be as expected
while True:
    res = api.get_metrics(req)
    if all([m.frames_tx == 10000 == m.frames_rx for m in res.flow_metrics]):
        break
```
