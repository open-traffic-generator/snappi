from .snappicommon import SnappiObject


class FlowIpv4(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def version(self):
        """flow_ipv4.version getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'version' not in self._properties or self._properties['version'] is None:
            self._properties['version'] = FlowPattern()
        return self._properties['version']

    @property
    def header_length(self):
        """flow_ipv4.header_length getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'header_length' not in self._properties or self._properties['header_length'] is None:
            self._properties['header_length'] = FlowPattern()
        return self._properties['header_length']

    @property
    def priority(self):
        """flow_ipv4.priority getter

        TBD  

        Returns: obj(snappi.FlowIpv4Priority)
        """
        from .flowipv4priority import FlowIpv4Priority
        if 'priority' not in self._properties or self._properties['priority'] is None:
            self._properties['priority'] = FlowIpv4Priority()
        return self._properties['priority']

    @property
    def total_length(self):
        """flow_ipv4.total_length getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'total_length' not in self._properties or self._properties['total_length'] is None:
            self._properties['total_length'] = FlowPattern()
        return self._properties['total_length']

    @property
    def identification(self):
        """flow_ipv4.identification getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'identification' not in self._properties or self._properties['identification'] is None:
            self._properties['identification'] = FlowPattern()
        return self._properties['identification']

    @property
    def reserved(self):
        """flow_ipv4.reserved getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved' not in self._properties or self._properties['reserved'] is None:
            self._properties['reserved'] = FlowPattern()
        return self._properties['reserved']

    @property
    def dont_fragment(self):
        """flow_ipv4.dont_fragment getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dont_fragment' not in self._properties or self._properties['dont_fragment'] is None:
            self._properties['dont_fragment'] = FlowPattern()
        return self._properties['dont_fragment']

    @property
    def more_fragments(self):
        """flow_ipv4.more_fragments getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'more_fragments' not in self._properties or self._properties['more_fragments'] is None:
            self._properties['more_fragments'] = FlowPattern()
        return self._properties['more_fragments']

    @property
    def fragment_offset(self):
        """flow_ipv4.fragment_offset getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'fragment_offset' not in self._properties or self._properties['fragment_offset'] is None:
            self._properties['fragment_offset'] = FlowPattern()
        return self._properties['fragment_offset']

    @property
    def time_to_live(self):
        """flow_ipv4.time_to_live getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'time_to_live' not in self._properties or self._properties['time_to_live'] is None:
            self._properties['time_to_live'] = FlowPattern()
        return self._properties['time_to_live']

    @property
    def protocol(self):
        """flow_ipv4.protocol getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'protocol' not in self._properties or self._properties['protocol'] is None:
            self._properties['protocol'] = FlowPattern()
        return self._properties['protocol']

    @property
    def header_checksum(self):
        """flow_ipv4.header_checksum getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'header_checksum' not in self._properties or self._properties['header_checksum'] is None:
            self._properties['header_checksum'] = FlowPattern()
        return self._properties['header_checksum']

    @property
    def src(self):
        """flow_ipv4.src getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src' not in self._properties or self._properties['src'] is None:
            self._properties['src'] = FlowPattern()
        return self._properties['src']

    @property
    def dst(self):
        """flow_ipv4.dst getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst' not in self._properties or self._properties['dst'] is None:
            self._properties['dst'] = FlowPattern()
        return self._properties['dst']

    def set(self):
        return self
