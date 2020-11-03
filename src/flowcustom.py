from .snappicommon import SnappiObject


class FlowCustom(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def bytes(self):
        """flow_custom.bytes getter

        A custom packet header defined as a string of hex bytes  
        The string MUST contain valid hex characters  
        Spaces or colons can be part of the bytes but will be discarded This can be used to create a custom protocol from other inputs such as scapy, wireshark, pcap etc  
        An example of ethernet/ipv4: '00000000000200000000000108004500001400010000400066e70a0000010a000002'  

        Returns: str
        """
        return self._properties['bytes']

    @bytes.setter
    def bytes(self, value):
        """flow_custom.bytes setter

        A custom packet header defined as a string of hex bytes  
        The string MUST contain valid hex characters  
        Spaces or colons can be part of the bytes but will be discarded This can be used to create a custom protocol from other inputs such as scapy, wireshark, pcap etc  
        An example of ethernet/ipv4: '00000000000200000000000108004500001400010000400066e70a0000010a000002'  

        value: str
        """
        self._properties['bytes'] = value

    @property
    def patterns(self):
        """flow_custom.patterns getter

        Modify the bytes with bit based patterns  

        Returns: list[obj(snappi.FlowBitPattern)]
        """
        from .flowbitpatternlist import FlowBitPatternList
        if 'patterns' not in self._properties or self._properties['patterns'] is None:
            self._properties['patterns'] = FlowBitPatternList()
        return self._properties['patterns']

    def set(self, bytes=None):
        self.bytes = bytes
        return self
