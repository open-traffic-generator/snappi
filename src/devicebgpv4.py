from .snappicommon import SnappiObject


class DeviceBgpv4(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """device_bgpv4.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """device_bgpv4.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def router_id(self):
        """device_bgpv4.router_id getter

        specifies BGP router identifier  
        It must be the string representation of an IPv4 address   

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'router_id' not in self._properties or self._properties['router_id'] is None:
            self._properties['router_id'] = DevicePattern()
        return self._properties['router_id']

    @property
    def as_number(self):
        """device_bgpv4.as_number getter

        Autonomous system (AS) number of 4 byte  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'as_number' not in self._properties or self._properties['as_number'] is None:
            self._properties['as_number'] = DevicePattern()
        return self._properties['as_number']

    @property
    def as_type(self):
        """device_bgpv4.as_type getter

        The type of BGP autonomous system  
        External BGP (EBGP) is used for BGP links between two or more autonomous systems  
        Internal BGP (IBGP) is used within a single autonomous system  

        Returns: Union[ibgp, ebgp]
        """
        return self._properties['as_type']

    @as_type.setter
    def as_type(self, value):
        """device_bgpv4.as_type setter

        The type of BGP autonomous system  
        External BGP (EBGP) is used for BGP links between two or more autonomous systems  
        Internal BGP (IBGP) is used within a single autonomous system  

        value: Union[ibgp, ebgp]
        """
        self._properties['as_type'] = value

    @property
    def hold_time_interval(self):
        """device_bgpv4.hold_time_interval getter

        Number of seconds the sender proposes for the value of the Hold Timer  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'hold_time_interval' not in self._properties or self._properties['hold_time_interval'] is None:
            self._properties['hold_time_interval'] = DevicePattern()
        return self._properties['hold_time_interval']

    @property
    def keep_alive_interval(self):
        """device_bgpv4.keep_alive_interval getter

        Number of seconds between transmissions of Keep Alive messages by router  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'keep_alive_interval' not in self._properties or self._properties['keep_alive_interval'] is None:
            self._properties['keep_alive_interval'] = DevicePattern()
        return self._properties['keep_alive_interval']

    @property
    def dut_ipv4_address(self):
        """device_bgpv4.dut_ipv4_address getter

        IPv4 address of the BGP peer for the session  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'dut_ipv4_address' not in self._properties or self._properties['dut_ipv4_address'] is None:
            self._properties['dut_ipv4_address'] = DevicePattern()
        return self._properties['dut_ipv4_address']

    @property
    def dut_as_number(self):
        """device_bgpv4.dut_as_number getter

        Autonomous system (AS) number of the BGP peer router (DUT)  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'dut_as_number' not in self._properties or self._properties['dut_as_number'] is None:
            self._properties['dut_as_number'] = DevicePattern()
        return self._properties['dut_as_number']

    @property
    def ipv4(self):
        """device_bgpv4.ipv4 getter

        The ipv4 stack that the bgp4 protocol is implemented over  

        Returns: obj(snappi.DeviceIpv4)
        """
        from .deviceipv4 import DeviceIpv4
        if 'ipv4' not in self._properties or self._properties['ipv4'] is None:
            self._properties['ipv4'] = DeviceIpv4()
        return self._properties['ipv4']

    @property
    def bgpv4_route_range(self):
        """device_bgpv4.bgpv4_route_range getter

        Emulated BGPv4 route range  

        Returns: list[obj(snappi.DeviceBgpv4RouteRange)]
        """
        from .devicebgpv4routerangelist import DeviceBgpv4RouteRangeList
        if 'bgpv4_route_range' not in self._properties or self._properties['bgpv4_route_range'] is None:
            self._properties['bgpv4_route_range'] = DeviceBgpv4RouteRangeList()
        return self._properties['bgpv4_route_range']

    @property
    def bgpv6_route_range(self):
        """device_bgpv4.bgpv6_route_range getter

        Emulated bgpv6 route range  

        Returns: list[obj(snappi.DeviceBgpv6RouteRange)]
        """
        from .devicebgpv6routerangelist import DeviceBgpv6RouteRangeList
        if 'bgpv6_route_range' not in self._properties or self._properties['bgpv6_route_range'] is None:
            self._properties['bgpv6_route_range'] = DeviceBgpv6RouteRangeList()
        return self._properties['bgpv6_route_range']

    def set(self, name=None, as_type=None):
        self.name = name
        self.as_type = as_type
        return self
