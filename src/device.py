from .snappicommon import SnappiObject


class Device(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """device.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """device.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def container_name(self):
        """device.container_name getter

        The unique name of a Port or Lag object that will contain the emulated interfaces and/or devices  

        Returns: str
        """
        return self._properties['container_name']

    @container_name.setter
    def container_name(self, value):
        """device.container_name setter

        The unique name of a Port or Lag object that will contain the emulated interfaces and/or devices  

        value: str
        """
        self._properties['container_name'] = value

    @property
    def device_count(self):
        """device.device_count getter

        The number of emulated protocol devices or interfaces per port  
        Example: If the device_count is 10 and the choice property value is ethernet then an implementation MUST create 10 ethernet interfaces  
        The ethernet property is a container for src, dst and eth_type properties with each on of those properties being a pattern container for 10 possible values  
        If an implementation is unable to support the maximum device_count it MUST indicate what the maximum device_count is using the /results/capabilities API  
        The device_count is also used by the individual child properties that are a container for a /components/schemas/Device.Pattern  

        Returns: int
        """
        return self._properties['device_count']

    @device_count.setter
    def device_count(self, value):
        """device.device_count setter

        The number of emulated protocol devices or interfaces per port  
        Example: If the device_count is 10 and the choice property value is ethernet then an implementation MUST create 10 ethernet interfaces  
        The ethernet property is a container for src, dst and eth_type properties with each on of those properties being a pattern container for 10 possible values  
        If an implementation is unable to support the maximum device_count it MUST indicate what the maximum device_count is using the /results/capabilities API  
        The device_count is also used by the individual child properties that are a container for a /components/schemas/Device.Pattern  

        value: int
        """
        self._properties['device_count'] = value

    @property
    def choice(self):
        """device.choice getter

        The type of emulated protocol interface or device  

        Returns: Union[ethernet, ipv4, ipv6, bgpv4]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """device.choice setter

        The type of emulated protocol interface or device  

        value: Union[ethernet, ipv4, ipv6, bgpv4]
        """
        self._properties['choice'] = value

    @property
    def ethernet(self):
        """device.ethernet getter

        TBD  

        Returns: obj(snappi.DeviceEthernet)
        """
        from .deviceethernet import DeviceEthernet
        if 'ethernet' not in self._properties or self._properties['ethernet'] is None:
            self._properties['ethernet'] = DeviceEthernet()
        return self._properties['ethernet']

    @property
    def ipv4(self):
        """device.ipv4 getter

        TBD  

        Returns: obj(snappi.DeviceIpv4)
        """
        from .deviceipv4 import DeviceIpv4
        if 'ipv4' not in self._properties or self._properties['ipv4'] is None:
            self._properties['ipv4'] = DeviceIpv4()
        return self._properties['ipv4']

    @property
    def ipv6(self):
        """device.ipv6 getter

        TBD  

        Returns: obj(snappi.DeviceIpv6)
        """
        from .deviceipv6 import DeviceIpv6
        if 'ipv6' not in self._properties or self._properties['ipv6'] is None:
            self._properties['ipv6'] = DeviceIpv6()
        return self._properties['ipv6']

    @property
    def bgpv4(self):
        """device.bgpv4 getter

        TBD  

        Returns: obj(snappi.DeviceBgpv4)
        """
        from .devicebgpv4 import DeviceBgpv4
        if 'bgpv4' not in self._properties or self._properties['bgpv4'] is None:
            self._properties['bgpv4'] = DeviceBgpv4()
        return self._properties['bgpv4']

    def set(self, name=None, container_name=None, device_count=None, choice=None):
        self.name = name
        self.container_name = container_name
        self.device_count = device_count
        self.choice = choice
        return self
