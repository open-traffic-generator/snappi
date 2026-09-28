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

## Setup Client
 
```sh
python -m pip install --upgrade snappi 
```

## Start Testing

```python
import datetime
import time
import snappi
import pytest


@pytest.mark.example
def test_quickstart():
    # Create a new API handle to make API calls against OTG
    # with HTTP as default transport protocol
    api = snappi.api(location="https://localhost:8443")

    # Create a new traffic configuration that will be set on OTG
    config = api.config()

    # Add a test port to the configuration
    ptx = config.ports.add(name="ptx", location="veth-a")

    # Configure a flow and set previously created test port as one of endpoints
    flow = config.flows.add(name="flow")
    flow.tx_rx.port.tx_name = ptx.name
    # and enable tracking flow metrics
    flow.metrics.enable = True

    # Configure number of packets to transmit for previously configured flow
    flow.duration.fixed_packets.packets = 100
    # and fixed byte size of all packets in the flow
    flow.size.fixed = 128

    # Configure protocol headers for all packets in the flow
    eth, ip, udp, cus = flow.packet.ethernet().ipv4().udp().custom()

    eth.src.value = "00:11:22:33:44:55"
    eth.dst.value = "00:11:22:33:44:66"

    ip.src.value = "10.1.1.1"
    ip.dst.value = "20.1.1.1"

    # Configure repeating patterns for source and destination UDP ports
    udp.src_port.values = [5010, 5015, 5020, 5025, 5030]
    udp.dst_port.increment.start = 6010
    udp.dst_port.increment.step = 5
    udp.dst_port.increment.count = 5

    # Configure custom bytes (hex string) in payload
    cus.bytes = "".join([hex(c)[2:] for c in b"..QUICKSTART SNAPPI.."])

    # Optionally, print JSON representation of config
    print("Configuration: ", config.serialize(encoding=config.JSON))

    # Push traffic configuration constructed so far to OTG
    api.set_config(config)

    # Start transmitting the packets from configured flow
    ts = api.transmit_state()
    ts.state = ts.START
    api.set_transmit_state(ts)

    # Fetch metrics for configured flow
    req = api.metrics_request()
    req.flow.flow_names = [flow.name]
    # and keep polling until either expectation is met or deadline exceeds
    start = datetime.datetime.now()
    while True:
        metrics = api.get_metrics(req)
        if (datetime.datetime.now() - start).seconds > 10:
            raise Exception("deadline exceeded")
        # print YAML representation of flow metrics
        print(metrics)
        if metrics.flow_metrics[0].transmit == metrics.flow_metrics[0].STOPPED:
            break
        time.sleep(0.1)
```

