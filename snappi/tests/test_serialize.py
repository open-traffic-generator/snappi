import pytest


def test_serialize_deserialize(api):
    """Demonstrate that there are no errors during serialization/deserialization
    of Snappi objects
    """
    control_state = api.control_state()
    control_state.set(choice='config_state')
    config = control_state.config_state.set(state='set').config
    config.options.port_options.set(location_preemption=True)
    config.ports.append(name='Tx Port', location='10.36.74.26;02;13')
    config.ports.append(name='Rx Port', location='10.36.74.26;02;14')
    flow = config.flows.append(name='Tx -> Rx Flow')
    flow.tx_rx.set(choice='port')
    flow.tx_rx.port.set(tx_port_name=config.ports[0].name,
                        rx_port_name=config.ports[1].name)
    flow.packet.append(choice='ethernet').append(choice='ipv4')

    # serialize as JSON and print to stdout
    serialization = control_state.encode(encoding='json')
    print(serialization)

    # serialize as YAML and print to stdout
    serialization = control_state.encode(encoding='yaml')
    print(serialization)

    # deserialize YAML to SnappiObject
    control_state = api.control_state().decode(serialization)

    # verify SnappiObject encoded as YAML is the same as initial YAML
    assert(control_state.encode(encoding='yaml') == serialization)


if __name__ == '__main__':
    pytest.main(['-s', __file__])