

class Api(object):
    """Snappi Abstract API
    """

    def set_state(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_state_results(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_capability_results(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_config(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_port_results(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_capture_results(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_flow_results(self, content):
        """TBD
        """
        raise NotImplementedError

    def get_bgpv4_results(self, content):
        """TBD
        """
        raise NotImplementedError

    @property
    def control_state(self):
        from .controlstate import ControlState
        return ControlState()

    @property
    def result_portrequest(self):
        from .resultportrequest import ResultPortRequest
        return ResultPortRequest()

    @property
    def result_capturerequest(self):
        from .resultcapturerequest import ResultCaptureRequest
        return ResultCaptureRequest()

    @property
    def result_flowrequest(self):
        from .resultflowrequest import ResultFlowRequest
        return ResultFlowRequest()

    @property
    def result_bgpv4request(self):
        from .resultbgpv4request import ResultBgpv4Request
        return ResultBgpv4Request()
