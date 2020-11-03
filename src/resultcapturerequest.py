from .snappicommon import SnappiObject


class ResultCaptureRequest(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def port_name(self):
        """result_capturerequest.port_name getter

        The name of a port a capture is started on  

        Returns: str
        """
        return self._properties['port_name']

    @port_name.setter
    def port_name(self, value):
        """result_capturerequest.port_name setter

        The name of a port a capture is started on  

        value: str
        """
        self._properties['port_name'] = value

    def set(self, port_name=None):
        self.port_name = port_name
        return self
