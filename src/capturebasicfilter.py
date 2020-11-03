from .snappicommon import SnappiObject


class CaptureBasicFilter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """capture_basicfilter.choice getter

        TBD  

        Returns: Union[mac_address, custom]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """capture_basicfilter.choice setter

        TBD  

        value: Union[mac_address, custom]
        """
        self._properties['choice'] = value

    @property
    def mac_address(self):
        """capture_basicfilter.mac_address getter

        TBD  

        Returns: obj(snappi.CaptureMacAddressFilter)
        """
        from .capturemacaddressfilter import CaptureMacAddressFilter
        if 'mac_address' not in self._properties or self._properties['mac_address'] is None:
            self._properties['mac_address'] = CaptureMacAddressFilter()
        return self._properties['mac_address']

    @property
    def custom(self):
        """capture_basicfilter.custom getter

        TBD  

        Returns: obj(snappi.CaptureCustomFilter)
        """
        from .capturecustomfilter import CaptureCustomFilter
        if 'custom' not in self._properties or self._properties['custom'] is None:
            self._properties['custom'] = CaptureCustomFilter()
        return self._properties['custom']

    @property
    def and_operator(self):
        """capture_basicfilter.and_operator getter

        TBD  

        Returns: boolean
        """
        return self._properties['and_operator']

    @and_operator.setter
    def and_operator(self, value):
        """capture_basicfilter.and_operator setter

        TBD  

        value: boolean
        """
        self._properties['and_operator'] = value

    @property
    def not_operator(self):
        """capture_basicfilter.not_operator getter

        TBD  

        Returns: boolean
        """
        return self._properties['not_operator']

    @not_operator.setter
    def not_operator(self, value):
        """capture_basicfilter.not_operator setter

        TBD  

        value: boolean
        """
        self._properties['not_operator'] = value

    def set(self, choice=None, and_operator=None, not_operator=None):
        self.choice = choice
        self.and_operator = and_operator
        self.not_operator = not_operator
        return self
