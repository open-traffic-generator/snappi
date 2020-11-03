from .snappicommon import SnappiObject


class Capture(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """capture.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """capture.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def port_names(self):
        """capture.port_names getter

        The unique names of ports that the capture settings will apply to  

        Returns: list[str]
        """
        return self._properties['port_names']

    @port_names.setter
    def port_names(self, value):
        """capture.port_names setter

        The unique names of ports that the capture settings will apply to  

        value: list[str]
        """
        self._properties['port_names'] = value

    @property
    def choice(self):
        """capture.choice getter

        The type of filter  

        Returns: Union[basic, pcap]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """capture.choice setter

        The type of filter  

        value: Union[basic, pcap]
        """
        self._properties['choice'] = value

    @property
    def basic(self):
        """capture.basic getter

        An array of basic filters  
        The filters supported are source address, destination address and custom  

        Returns: list[obj(snappi.CaptureBasicFilter)]
        """
        from .capturebasicfilterlist import CaptureBasicFilterList
        if 'basic' not in self._properties or self._properties['basic'] is None:
            self._properties['basic'] = CaptureBasicFilterList()
        return self._properties['basic']

    @property
    def pcap(self):
        """capture.pcap getter

        The content of this property must be of pcap filter syntax  
        https://www.tcpdump.org/manpages/pcap-filter.7.html  

        Returns: str
        """
        return self._properties['pcap']

    @pcap.setter
    def pcap(self, value):
        """capture.pcap setter

        The content of this property must be of pcap filter syntax  
        https://www.tcpdump.org/manpages/pcap-filter.7.html  

        value: str
        """
        self._properties['pcap'] = value

    @property
    def enable(self):
        """capture.enable getter

        Enable capture on the port  

        Returns: boolean
        """
        return self._properties['enable']

    @enable.setter
    def enable(self, value):
        """capture.enable setter

        Enable capture on the port  

        value: boolean
        """
        self._properties['enable'] = value

    @property
    def overwrite(self):
        """capture.overwrite getter

        Overwrite the capture buffer  

        Returns: boolean
        """
        return self._properties['overwrite']

    @overwrite.setter
    def overwrite(self, value):
        """capture.overwrite setter

        Overwrite the capture buffer  

        value: boolean
        """
        self._properties['overwrite'] = value

    @property
    def format(self):
        """capture.format getter

        The format of the capture file  

        Returns: Union[pcap, pcapng]
        """
        return self._properties['format']

    @format.setter
    def format(self, value):
        """capture.format setter

        The format of the capture file  

        value: Union[pcap, pcapng]
        """
        self._properties['format'] = value

    def set(self, name=None, port_names=None, choice=None, pcap=None, enable=None, overwrite=None, format=None):
        self.name = name
        self.port_names = port_names
        self.choice = choice
        self.pcap = pcap
        self.enable = enable
        self.overwrite = overwrite
        self.format = format
        return self
