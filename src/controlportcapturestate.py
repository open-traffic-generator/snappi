from .snappicommon import SnappiObject


class ControlPortCaptureState(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def port_names(self):
        """control_portcapturestate.port_names getter

        The name of ports to start capturing packets on  
        An empty or null list will control all port objects  

        Returns: list[str]
        """
        return self._properties['port_names']

    @port_names.setter
    def port_names(self, value):
        """control_portcapturestate.port_names setter

        The name of ports to start capturing packets on  
        An empty or null list will control all port objects  

        value: list[str]
        """
        self._properties['port_names'] = value

    @property
    def state(self):
        """control_portcapturestate.state getter

        The capture state  

        Returns: Union[start, stop]
        """
        return self._properties['state']

    @state.setter
    def state(self, value):
        """control_portcapturestate.state setter

        The capture state  

        value: Union[start, stop]
        """
        self._properties['state'] = value

    def set(self, port_names=None, state=None):
        self.port_names = port_names
        self.state = state
        return self
