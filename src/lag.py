from .snappicommon import SnappiObject


class Lag(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """lag.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """lag.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def port_names(self):
        """lag.port_names getter

        A list of unique names of port objects that will be part of the same lag  
        The value of the port_names property is the count for any child property in this hierarchy that is a container for a device pattern  

        Returns: list[str]
        """
        return self._properties['port_names']

    @port_names.setter
    def port_names(self, value):
        """lag.port_names setter

        A list of unique names of port objects that will be part of the same lag  
        The value of the port_names property is the count for any child property in this hierarchy that is a container for a device pattern  

        value: list[str]
        """
        self._properties['port_names'] = value

    @property
    def protocol(self):
        """lag.protocol getter

        Static lag or LACP protocol settings  

        Returns: obj(snappi.LagProtocol)
        """
        from .lagprotocol import LagProtocol
        if 'protocol' not in self._properties or self._properties['protocol'] is None:
            self._properties['protocol'] = LagProtocol()
        return self._properties['protocol']

    @property
    def ethernet(self):
        """lag.ethernet getter

        Per port ethernet and vlan settings  

        Returns: obj(snappi.DeviceEthernet)
        """
        from .deviceethernet import DeviceEthernet
        if 'ethernet' not in self._properties or self._properties['ethernet'] is None:
            self._properties['ethernet'] = DeviceEthernet()
        return self._properties['ethernet']

    def set(self, name=None, port_names=None):
        self.name = name
        self.port_names = port_names
        return self
