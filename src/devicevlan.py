from .snappicommon import SnappiObject


class DeviceVlan(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """device_vlan.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """device_vlan.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def tpid(self):
        """device_vlan.tpid getter

        Vlan tag protocol identifier  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'tpid' not in self._properties or self._properties['tpid'] is None:
            self._properties['tpid'] = DevicePattern()
        return self._properties['tpid']

    @property
    def priority(self):
        """device_vlan.priority getter

        Vlan priority  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'priority' not in self._properties or self._properties['priority'] is None:
            self._properties['priority'] = DevicePattern()
        return self._properties['priority']

    @property
    def id(self):
        """device_vlan.id getter

        Vlan id  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'id' not in self._properties or self._properties['id'] is None:
            self._properties['id'] = DevicePattern()
        return self._properties['id']

    def set(self, name=None):
        self.name = name
        return self
