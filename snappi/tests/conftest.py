import pytest


@pytest.fixture(scope='session')
def api():
    """Demonstrates creating a top level Api instance.
    """
    import threading
    web_server_thread = threading.Thread(target=web_server)
    web_server_thread.setDaemon(True)
    web_server_thread.start()
    from ..api import Api
    yield Api()


@pytest.fixture()
def b2b_config(api):
    """Demonstrates creating a back to back configuration of tx and rx 
    ports, devices and a single flow using those ports as endpoints for
    transmit and receive.
    """
    config = api.config()

    tx_port, rx_port = config.ports \
        .port(name='Tx Port', location='10.36.74.26;02;13') \
        .port(name='Rx Port', location='10.36.74.26;02;14')

    tx_device, rx_device = config.devices \
        .device(name='Tx Devices', container_name=tx_port.name, device_count=1) \
        .device(name='Rx Devices', container_name=tx_port.name, device_count=1)
    tx_device.ipv4.name = 'Tx Ipv4'
    tx_device.ipv4.address.value = '1.1.1.1'
    tx_device.ipv4.gateway.value = '1.1.2.1'
    tx_device.ipv4.prefix.value = 16
    tx_device.ipv4.ethernet.mac.value = '00:00:01:00:00:01'

    flow = config.flows.flow(name='Tx -> Rx Flow')
    flow.tx_rx.port.tx_name = tx_port.name
    flow.tx_rx.port.rx_name = rx_port.name
    flow.size.fixed = 128
    flow.rate.pps = 1000
    flow.duration.fixed_packets.packets = 10000

    eth, vlan, ip, tcp = flow.packet.ethernet().vlan().ipv4().tcp()

    eth.src.value = '00:00:01:00:00:01'
    eth.dst.value_list = ['00:00:02:00:00:01', '00:00:02:00:00:01']
    eth.dst.result_group = 'eth dst mac'

    ip.src.increment.start = '1.1.1.1'
    ip.src.increment.step = '0.0.0.1'
    ip.src.increment.count = 10

    ip.dst.decrement.start = '1.1.2.200'
    ip.dst.decrement.step = '0.0.0.1'
    ip.dst.decrement.count = 10

    ip.priority.dscp.phb.value_list = [8, 16, 32]
    ip.priority.dscp.ecn.value = 1
    tcp.src_port.increment.start = '10'
    tcp.dst_port.increment.start = 1

    return config


from flask import Flask, request, Response, jsonify
app = Flask(__name__)
CONFIG = None


@app.route('/config', methods=['POST'])
def set_config():
    global CONFIG
    import snappi
    config = snappi.api.Api().config()
    config.deserialize(request.data.decode('utf-8'))
    CONFIG = config
    return Response(status=200)


@app.route('/config', methods=['GET'])
def get_config():
    global CONFIG
    return Response(CONFIG.serialize(),
                    mimetype='application/json',
                    status=200)


@app.after_request
def after_request(resp):
    print(request.method, request.url, ' -> ', resp.status)
    return resp


def web_server():
    app.run(port=80, debug=True, use_reloader=False)