from .snappicommon import SnappiObject


class FlowVlan(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def priority(self):
        """flow_vlan.priority getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'priority' not in self._properties or self._properties['priority'] is None:
            self._properties['priority'] = FlowPattern()
        return self._properties['priority']

    @property
    def cfi(self):
        """flow_vlan.cfi getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'cfi' not in self._properties or self._properties['cfi'] is None:
            self._properties['cfi'] = FlowPattern()
        return self._properties['cfi']

    @property
    def id(self):
        """flow_vlan.id getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'id' not in self._properties or self._properties['id'] is None:
            self._properties['id'] = FlowPattern()
        return self._properties['id']

    @property
    def protocol(self):
        """flow_vlan.protocol getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'protocol' not in self._properties or self._properties['protocol'] is None:
            self._properties['protocol'] = FlowPattern()
        return self._properties['protocol']

    def set(self):
        return self
