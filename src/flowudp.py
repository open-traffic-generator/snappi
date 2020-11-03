from .snappicommon import SnappiObject


class FlowUdp(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def src_port(self):
        """flow_udp.src_port getter

        Udp source port  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src_port' not in self._properties or self._properties['src_port'] is None:
            self._properties['src_port'] = FlowPattern()
        return self._properties['src_port']

    @property
    def dst_port(self):
        """flow_udp.dst_port getter

        Tcp destination port  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst_port' not in self._properties or self._properties['dst_port'] is None:
            self._properties['dst_port'] = FlowPattern()
        return self._properties['dst_port']

    @property
    def length(self):
        """flow_udp.length getter

        Length in bytes of the udp header and yudp data  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'length' not in self._properties or self._properties['length'] is None:
            self._properties['length'] = FlowPattern()
        return self._properties['length']

    @property
    def checksum(self):
        """flow_udp.checksum getter

        Checksum field used for error checking of header and data  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'checksum' not in self._properties or self._properties['checksum'] is None:
            self._properties['checksum'] = FlowPattern()
        return self._properties['checksum']

    def set(self):
        return self
