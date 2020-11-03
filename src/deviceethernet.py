from .snappicommon import SnappiObject


class DeviceEthernet(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """device_ethernet.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """device_ethernet.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def mac(self):
        """device_ethernet.mac getter

        Media access control address (MAC) is a 48bit identifier for use as a network address  
        The value can be an int or a hex string with or without spaces or colons separating each byte  
        The min value is 0 or '00:00:00:00:00:00'  
        The max value is 281474976710655 or 'FF:FF:FF:FF:FF:FF'  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'mac' not in self._properties or self._properties['mac'] is None:
            self._properties['mac'] = DevicePattern()
        return self._properties['mac']

    @property
    def mtu(self):
        """device_ethernet.mtu getter

        TBD  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'mtu' not in self._properties or self._properties['mtu'] is None:
            self._properties['mtu'] = DevicePattern()
        return self._properties['mtu']

    @property
    def vlans(self):
        """device_ethernet.vlans getter

        List of vlans  

        Returns: list[obj(snappi.DeviceVlan)]
        """
        from .devicevlanlist import DeviceVlanList
        if 'vlans' not in self._properties or self._properties['vlans'] is None:
            self._properties['vlans'] = DeviceVlanList()
        return self._properties['vlans']

    def set(self, name=None):
        self.name = name
        return self
