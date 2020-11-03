from .snappicommon import SnappiObject


class FlowEthernet(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def dst(self):
        """flow_ethernet.dst getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst' not in self._properties or self._properties['dst'] is None:
            self._properties['dst'] = FlowPattern()
        return self._properties['dst']

    @property
    def src(self):
        """flow_ethernet.src getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src' not in self._properties or self._properties['src'] is None:
            self._properties['src'] = FlowPattern()
        return self._properties['src']

    @property
    def ether_type(self):
        """flow_ethernet.ether_type getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ether_type' not in self._properties or self._properties['ether_type'] is None:
            self._properties['ether_type'] = FlowPattern()
        return self._properties['ether_type']

    def set(self):
        return self
