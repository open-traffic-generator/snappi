from .snappicommon import SnappiObject


class FlowEthernetPause(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def dst(self):
        """flow_ethernetpause.dst getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst' not in self._properties or self._properties['dst'] is None:
            self._properties['dst'] = FlowPattern()
        return self._properties['dst']

    @property
    def src(self):
        """flow_ethernetpause.src getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src' not in self._properties or self._properties['src'] is None:
            self._properties['src'] = FlowPattern()
        return self._properties['src']

    @property
    def ether_type(self):
        """flow_ethernetpause.ether_type getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ether_type' not in self._properties or self._properties['ether_type'] is None:
            self._properties['ether_type'] = FlowPattern()
        return self._properties['ether_type']

    @property
    def control_op_code(self):
        """flow_ethernetpause.control_op_code getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'control_op_code' not in self._properties or self._properties['control_op_code'] is None:
            self._properties['control_op_code'] = FlowPattern()
        return self._properties['control_op_code']

    @property
    def time(self):
        """flow_ethernetpause.time getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'time' not in self._properties or self._properties['time'] is None:
            self._properties['time'] = FlowPattern()
        return self._properties['time']

    def set(self):
        return self
