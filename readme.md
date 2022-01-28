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
# create a new API instance where location points to controller
api = snappi.api(location='https://localhost')

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

## Arguments of API Handle

The first step in any snappi script is to import the `snappi` package and instantiate an `api` object. These are possible arguments:

| API Arguments               |                      Description                                                     |
|-----------------------------|--------------------------------------------------------------------------------------|
| location                    |  Takes the HTTPS or gRPC address and port of the controller                          |
| transport                   |  Accepting `snappi.Transport.HTTP` (Default) or `snappi.Transport.GRPC`              |
| ext                         |  Translation layer used by not supported OTG model. `ext` and `transport` are not mutually exclusive.                                   |
| verify                      |  Turn off insecure certificate warning                                               |
| logger                      |  A user defined `logging.logger`, if none is provided then a default logger with a stdout handler will be provided  |
| loglevel                    |  The logging package log level. The default loglevel is `logging.INFO`  |

* Instantiate HTTPS Transport

Please specify non-default TCP port using (default is 443).

```python
import snappi
api = snappi.api(location='https://localhost', verify=False)
# or with non-default TCP port
api = snappi.api(location='https://localhost:8080', verify=False)
# or user explicitly specify transport type 
api = snappi.api(location='https://localhost:8080', verify=False, transport=snappi.Transport.HTTP)
```

* Instantiate gRPC Transport

Location should contain `address:port`.As well as explicitly specify transport type as `snappi.Transport.GRPC`. We also have a flexibility to configure `keep_alive_timeout` and `request_timeout`
```python
import snappi
api = snappi.api(location='localhost:50001', transport=snappi.Transport.GRPC)
api.keep_alive_timeout = 2000
api.request_timeout = 20
```

* Optional (`ext`) parameter

If a traffic generator doesn't natively support [Open Traffic Generator API](https://github.com/open-traffic-generator/models), snappi can be extended to write a translation layer to bridge the gap. An example is [snappi extension for IxNetwork](https://pypi.org/project/snappi-ixnetwork/) which can be installed using `python -m pip install --upgrade snappi[ixnetwork]`. 

Note: `ext` and `transport` are not mutually exclusive.

```python
import snappi
# location here refers to HTTPS address of IxNetwork API Server
api = snappi.api(location="https://localhost", ext='ixnetwork', verify=False)
```
