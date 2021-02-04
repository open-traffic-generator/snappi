from flask import Flask, request, Response
import threading
import json
import time
import snappi

app = Flask(__name__)
CONFIG = None


@app.route('/config', methods=['POST'])
def set_config():
    global CONFIG
    config = snappi.api().config()
    config.deserialize(request.data.decode('utf-8'))
    test = config.options.port_options.location_preemption
    if test is not None and isinstance(test, bool) is False:
        return Response(status=590,
                        response=json.dumps({'detail': 'invalid data type'}),
                        headers={'Content-Type': 'application/json'})
    else:
        CONFIG = config
        return Response(status=200)


@app.route('/config', methods=['GET'])
def get_config():
    global CONFIG
    return Response(CONFIG.serialize() if CONFIG is not None else '{}',
                    mimetype='application/json',
                    status=200)


@app.route('/control/transmit', methods=['POST'])
def set_transmit_state():
    global CONFIG
    return Response(status=200)


@app.route('/results/metrics', methods=['POST'])
def get_metrics():
    global CONFIG
    api = snappi.api()

    metrics_request = api.metrics_request()
    metrics_request.deserialize(request.data.decode('utf-8'))
    metrics_response = api.metrics_response()
    if metrics_request.choice == 'port':
        for port in CONFIG.ports:
            metrics_response.port_metrics.metric(
                name=port.name, frames_tx=10000, frames_rx=10000
            )
    elif metrics_request.choice == 'flow':
        for flow in CONFIG.flows:
            metrics_response.flow_metrics.metric(
                name=flow.name, frames_tx=10000, frames_rx=10000
            )

    return Response(metrics_response.serialize(),
                    mimetype='application/json',
                    status=200)

@app.after_request
def after_request(resp):
    print(request.method, request.url, ' -> ', resp.status)
    return resp


def web_server():
    app.run(port=80, debug=True, use_reloader=False)


class SnappiServer(object):
    def __init__(self):
        self._CONFIG = None

    def start(self):
        self._web_server_thread = threading.Thread(target=web_server)
        self._web_server_thread.setDaemon(True)
        self._web_server_thread.start()
        self._wait_until_ready()
        return self

    def _wait_until_ready(self):
        api = snappi.api(host='http://127.0.0.1:80')
        while True:
            try:
                api.get_config()
                break
            except Exception as err:
                pass
            time.sleep(0.1)