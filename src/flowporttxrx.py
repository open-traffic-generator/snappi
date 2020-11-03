from .snappicommon import SnappiObject


class FlowPortTxRx(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def tx_port_name(self):
        """flow_porttxrx.tx_port_name getter

        The unique name of a port that is the transmit port  

        Returns: str
        """
        return self._properties['tx_port_name']

    @tx_port_name.setter
    def tx_port_name(self, value):
        """flow_porttxrx.tx_port_name setter

        The unique name of a port that is the transmit port  

        value: str
        """
        self._properties['tx_port_name'] = value

    @property
    def rx_port_name(self):
        """flow_porttxrx.rx_port_name getter

        The unique name of a port that is the intended receive port  

        Returns: str
        """
        return self._properties['rx_port_name']

    @rx_port_name.setter
    def rx_port_name(self, value):
        """flow_porttxrx.rx_port_name setter

        The unique name of a port that is the intended receive port  

        value: str
        """
        self._properties['rx_port_name'] = value

    def set(self, tx_port_name=None, rx_port_name=None):
        self.tx_port_name = tx_port_name
        self.rx_port_name = rx_port_name
        return self
