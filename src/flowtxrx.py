from .snappicommon import SnappiObject


class FlowTxRx(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_txrx.choice getter

        The type of transmit and receive container used by the flow  

        Returns: Union[port, device]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_txrx.choice setter

        The type of transmit and receive container used by the flow  

        value: Union[port, device]
        """
        self._properties['choice'] = value

    @property
    def port(self):
        """flow_txrx.port getter

        TBD  

        Returns: obj(snappi.FlowPortTxRx)
        """
        from .flowporttxrx import FlowPortTxRx
        if 'port' not in self._properties or self._properties['port'] is None:
            self._properties['port'] = FlowPortTxRx()
        return self._properties['port']

    @property
    def device(self):
        """flow_txrx.device getter

        TBD  

        Returns: obj(snappi.FlowDeviceTxRx)
        """
        from .flowdevicetxrx import FlowDeviceTxRx
        if 'device' not in self._properties or self._properties['device'] is None:
            self._properties['device'] = FlowDeviceTxRx()
        return self._properties['device']

    def set(self, choice=None):
        self.choice = choice
        return self
