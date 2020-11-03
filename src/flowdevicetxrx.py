from .snappicommon import SnappiObject


class FlowDeviceTxRx(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def tx_device_names(self):
        """flow_devicetxrx.tx_device_names getter

        The unique names of devices that will be transmitting  

        Returns: list[str]
        """
        return self._properties['tx_device_names']

    @tx_device_names.setter
    def tx_device_names(self, value):
        """flow_devicetxrx.tx_device_names setter

        The unique names of devices that will be transmitting  

        value: list[str]
        """
        self._properties['tx_device_names'] = value

    @property
    def rx_device_names(self):
        """flow_devicetxrx.rx_device_names getter

        The unique names of emulated devices that will be receiving  

        Returns: list[str]
        """
        return self._properties['rx_device_names']

    @rx_device_names.setter
    def rx_device_names(self, value):
        """flow_devicetxrx.rx_device_names setter

        The unique names of emulated devices that will be receiving  

        value: list[str]
        """
        self._properties['rx_device_names'] = value

    def set(self, tx_device_names=None, rx_device_names=None):
        self.tx_device_names = tx_device_names
        self.rx_device_names = rx_device_names
        return self
