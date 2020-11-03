from .snappicommon import SnappiObject


class CaptureMacAddressFilter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def mac(self):
        """capture_macaddressfilter.mac getter

        The type of mac address filters  
        This can be either source or destination  

        Returns: Union[source, destination]
        """
        return self._properties['mac']

    @mac.setter
    def mac(self, value):
        """capture_macaddressfilter.mac setter

        The type of mac address filters  
        This can be either source or destination  

        value: Union[source, destination]
        """
        self._properties['mac'] = value

    @property
    def filter(self):
        """capture_macaddressfilter.filter getter

        The value of the mac address  

        Returns: str
        """
        return self._properties['filter']

    @filter.setter
    def filter(self, value):
        """capture_macaddressfilter.filter setter

        The value of the mac address  

        value: str
        """
        self._properties['filter'] = value

    @property
    def mask(self):
        """capture_macaddressfilter.mask getter

        The value of the mask to be applied to the mac address  

        Returns: str
        """
        return self._properties['mask']

    @mask.setter
    def mask(self, value):
        """capture_macaddressfilter.mask setter

        The value of the mask to be applied to the mac address  

        value: str
        """
        self._properties['mask'] = value

    def set(self, mac=None, filter=None, mask=None):
        self.mac = mac
        self.filter = filter
        self.mask = mask
        return self
