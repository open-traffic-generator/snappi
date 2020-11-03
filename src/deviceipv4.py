from .snappicommon import SnappiObject


class DeviceIpv4(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """device_ipv4.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """device_ipv4.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def address(self):
        """device_ipv4.address getter

        The ipv4 address  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'address' not in self._properties or self._properties['address'] is None:
            self._properties['address'] = DevicePattern()
        return self._properties['address']

    @property
    def gateway(self):
        """device_ipv4.gateway getter

        The ipv4 address of the gateway  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'gateway' not in self._properties or self._properties['gateway'] is None:
            self._properties['gateway'] = DevicePattern()
        return self._properties['gateway']

    @property
    def prefix(self):
        """device_ipv4.prefix getter

        The prefix of the ipv4 address  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'prefix' not in self._properties or self._properties['prefix'] is None:
            self._properties['prefix'] = DevicePattern()
        return self._properties['prefix']

    @property
    def ethernet(self):
        """device_ipv4.ethernet getter

        TBD  

        Returns: obj(snappi.DeviceEthernet)
        """
        from .deviceethernet import DeviceEthernet
        if 'ethernet' not in self._properties or self._properties['ethernet'] is None:
            self._properties['ethernet'] = DeviceEthernet()
        return self._properties['ethernet']

    def set(self, name=None):
        self.name = name
        return self
