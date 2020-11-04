import pytest
import sys
sys.path.append('../')
import src
import json


def test_serialize_deserialize():
    api = src.api.Api()
    control_state = api.control_state.set(choice='config_state')
    config = control_state.config_state.set(state='set').config
    config.options.port_options.set(location_preemption=True)
    config.ports.append(name='Tx Port', location='10.36.74.26;02;13')
    config.ports.append(name='Rx Port', location='10.36.74.26;02;14')
    flow = config.flows.append(name='Tx -> Rx Flow')
    flow.tx_rx.set(choice='port')
    flow.tx_rx.port.set(tx_port_name=config.ports[0].name,
                        rx_port_name=config.ports[1].name)
    serialization = json.dumps(control_state, indent=2)
    snappi_object = src.snappicommon.decode(api.control_state.__class__,
                                            json.loads(serialization))
    print(snappi_object)


if __name__ == '__main__':
    pytest.main(['-s', __file__])
