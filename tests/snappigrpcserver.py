import os
import sys
import grpc
import json
import base64
import threading
import importlib
from concurrent import futures
from google.protobuf import json_format

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "snappi"))
pb2_grpc = importlib.import_module("otg_pb2_grpc")
pb2 = importlib.import_module("otg_pb2")

GRPC_PORT = 50001

class OpenapiServicer(pb2_grpc.OpenapiServicer):
    def __init__(self):
        self._config = None
        super(OpenapiServicer, self).__init__()

    def _log(self, value):
        print("gRPC Server: %s" %value)

    def SetConfig(self, request, context):
        self._log("Executing SetConfig")
        response_400 = """
            {
                "status_code_400" : {
                    "errors" : ["invalid value"]
                }
            }
            """

        response_200 = """
            {
                "warning" : {
                    "warnings" : ["no"]
                }
            }
            """

        self._config = json_format.MessageToDict(
            request.config, preserving_proto_field_name=True
        )
        res_obj = json_format.Parse(response_200, pb2.SetConfigResponse())
        return res_obj

    def StreamConfig(self, request_iterator, context):
        self._log("Executing SetConfig")
        full_str = b""
        for data in request_iterator:
            full_str += data.datum
            self._log("received ")

        self._log("received all chunks ")
        self._log(full_str)
        obj = pb2.Config()
        obj.ParseFromString(full_str)
        self._log(obj)

        response_200 = """
                   {
                       "warning" : {
                           "warnings" : ["no"]
                       }
                   }
                   """

        self._config = json_format.MessageToDict(
            obj, preserving_proto_field_name=True
        )
        res_obj = json_format.Parse(response_200, pb2.SetConfigResponse())
        return res_obj

    def GetConfig(self, request, context):
        self._log("Executing GetConfig")
        if self._config is None:
            self._config = {}
        response_200 = {
            "config": self._config
        }
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetConfigResponse()
        )
        return res_obj

def gRpcServer():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_OpenapiServicer_to_server(OpenapiServicer(), server)
    print("gRPC Server: Starting server. Listening on port %s." % GRPC_PORT)
    server.add_insecure_port("[::]:{}".format(GRPC_PORT))
    server.start()

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(5)
        print("Server shutdown gracefully")


def grpc_server():
    web_server_thread = threading.Thread(target=gRpcServer)
    web_server_thread.setDaemon(True)
    web_server_thread.start()

if __name__ == '__main__':
    gRpcServer()