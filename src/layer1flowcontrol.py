from .snappicommon import SnappiObject


class Layer1FlowControl(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def directed_address(self):
        """layer1_flowcontrol.directed_address getter

        The 48bit mac address that the layer1 port names will listen on for a directed pause  

        Returns: str
        """
        return self._properties['directed_address']

    @directed_address.setter
    def directed_address(self, value):
        """layer1_flowcontrol.directed_address setter

        The 48bit mac address that the layer1 port names will listen on for a directed pause  

        value: str
        """
        self._properties['directed_address'] = value

    @property
    def choice(self):
        """layer1_flowcontrol.choice getter

        The type of priority flow control  

        Returns: Union[ieee_802_1qbb, ieee_802_3x]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """layer1_flowcontrol.choice setter

        The type of priority flow control  

        value: Union[ieee_802_1qbb, ieee_802_3x]
        """
        self._properties['choice'] = value

    @property
    def ieee_802_1qbb(self):
        """layer1_flowcontrol.ieee_802_1qbb getter

        TBD  

        Returns: obj(snappi.Layer1Ieee8021qbb)
        """
        from .layer1ieee8021qbb import Layer1Ieee8021qbb
        if 'ieee_802_1qbb' not in self._properties or self._properties['ieee_802_1qbb'] is None:
            self._properties['ieee_802_1qbb'] = Layer1Ieee8021qbb()
        return self._properties['ieee_802_1qbb']

    @property
    def ieee_802_3x(self):
        """layer1_flowcontrol.ieee_802_3x getter

        TBD  

        Returns: obj(snappi.Layer1Ieee8023x)
        """
        from .layer1ieee8023x import Layer1Ieee8023x
        if 'ieee_802_3x' not in self._properties or self._properties['ieee_802_3x'] is None:
            self._properties['ieee_802_3x'] = Layer1Ieee8023x()
        return self._properties['ieee_802_3x']

    def set(self, directed_address=None, choice=None):
        self.directed_address = directed_address
        self.choice = choice
        return self
