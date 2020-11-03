from .snappicommon import SnappiObject


class FlowIpv6(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def version(self):
        """flow_ipv6.version getter

        Default version number is 6 (bit sequence 0110) 4 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'version' not in self._properties or self._properties['version'] is None:
            self._properties['version'] = FlowPattern()
        return self._properties['version']

    @property
    def traffic_class(self):
        """flow_ipv6.traffic_class getter

        8 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'traffic_class' not in self._properties or self._properties['traffic_class'] is None:
            self._properties['traffic_class'] = FlowPattern()
        return self._properties['traffic_class']

    @property
    def flow_label(self):
        """flow_ipv6.flow_label getter

        20 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'flow_label' not in self._properties or self._properties['flow_label'] is None:
            self._properties['flow_label'] = FlowPattern()
        return self._properties['flow_label']

    @property
    def payload_length(self):
        """flow_ipv6.payload_length getter

        16 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'payload_length' not in self._properties or self._properties['payload_length'] is None:
            self._properties['payload_length'] = FlowPattern()
        return self._properties['payload_length']

    @property
    def next_header(self):
        """flow_ipv6.next_header getter

        8 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'next_header' not in self._properties or self._properties['next_header'] is None:
            self._properties['next_header'] = FlowPattern()
        return self._properties['next_header']

    @property
    def hop_limit(self):
        """flow_ipv6.hop_limit getter

        8 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'hop_limit' not in self._properties or self._properties['hop_limit'] is None:
            self._properties['hop_limit'] = FlowPattern()
        return self._properties['hop_limit']

    @property
    def src(self):
        """flow_ipv6.src getter

        128 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src' not in self._properties or self._properties['src'] is None:
            self._properties['src'] = FlowPattern()
        return self._properties['src']

    @property
    def dst(self):
        """flow_ipv6.dst getter

        128 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst' not in self._properties or self._properties['dst'] is None:
            self._properties['dst'] = FlowPattern()
        return self._properties['dst']

    def set(self):
        return self
