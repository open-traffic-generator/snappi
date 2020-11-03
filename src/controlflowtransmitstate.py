from .snappicommon import SnappiObject


class ControlFlowTransmitState(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def flow_names(self):
        """control_flowtransmitstate.flow_names getter

        The names of flow objects to control  
        An empty or null list will control all flow objects  

        Returns: list[str]
        """
        return self._properties['flow_names']

    @flow_names.setter
    def flow_names(self, value):
        """control_flowtransmitstate.flow_names setter

        The names of flow objects to control  
        An empty or null list will control all flow objects  

        value: list[str]
        """
        self._properties['flow_names'] = value

    @property
    def state(self):
        """control_flowtransmitstate.state getter

        The transmit state  

        Returns: Union[start, stop, pause]
        """
        return self._properties['state']

    @state.setter
    def state(self, value):
        """control_flowtransmitstate.state setter

        The transmit state  

        value: Union[start, stop, pause]
        """
        self._properties['state'] = value

    def set(self, flow_names=None, state=None):
        self.flow_names = flow_names
        self.state = state
        return self
