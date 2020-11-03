from .snappicommon import SnappiObject


class FlowVxlan(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def flags(self):
        """flow_vxlan.flags getter

        RRRRIRRR Where the I flag MUST be set to 1 for a valid vxlan network id (VNI)  
        The other 7 bits (designated "R") are reserved fields and MUST be set to zero on transmission and ignored on receipt  
        8 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'flags' not in self._properties or self._properties['flags'] is None:
            self._properties['flags'] = FlowPattern()
        return self._properties['flags']

    @property
    def reserved0(self):
        """flow_vxlan.reserved0 getter

        Set to 0  
        24 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved0' not in self._properties or self._properties['reserved0'] is None:
            self._properties['reserved0'] = FlowPattern()
        return self._properties['reserved0']

    @property
    def vni(self):
        """flow_vxlan.vni getter

        Vxlan network id  
        24 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'vni' not in self._properties or self._properties['vni'] is None:
            self._properties['vni'] = FlowPattern()
        return self._properties['vni']

    @property
    def reserved1(self):
        """flow_vxlan.reserved1 getter

        Set to 0  
        8 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved1' not in self._properties or self._properties['reserved1'] is None:
            self._properties['reserved1'] = FlowPattern()
        return self._properties['reserved1']

    def set(self):
        return self
