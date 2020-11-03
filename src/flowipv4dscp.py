from .snappicommon import SnappiObject


class FlowIpv4Dscp(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def phb(self):
        """flow_ipv4_dscp.phb getter

        phb (per-hop-behavior) value is 6 bits: >=0 PHB <=63  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'phb' not in self._properties or self._properties['phb'] is None:
            self._properties['phb'] = FlowPattern()
        return self._properties['phb']

    @property
    def ecn(self):
        """flow_ipv4_dscp.ecn getter

        ecn (explicit-congestion-notification) value is 2 bits: >=0 ecn <=3  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ecn' not in self._properties or self._properties['ecn'] is None:
            self._properties['ecn'] = FlowPattern()
        return self._properties['ecn']

    def set(self):
        return self
