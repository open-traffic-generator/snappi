from .snappicommon import SnappiObject


class FlowIpv4Priority(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_ipv4_priority.choice getter

        TBD  

        Returns: Union[raw, tos, dscp]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_ipv4_priority.choice setter

        TBD  

        value: Union[raw, tos, dscp]
        """
        self._properties['choice'] = value

    @property
    def raw(self):
        """flow_ipv4_priority.raw getter

        Raw priority is 8 bits: >=0 raw <=255  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'raw' not in self._properties or self._properties['raw'] is None:
            self._properties['raw'] = FlowPattern()
        return self._properties['raw']

    @property
    def tos(self):
        """flow_ipv4_priority.tos getter

        TBD  

        Returns: obj(snappi.FlowIpv4Tos)
        """
        from .flowipv4tos import FlowIpv4Tos
        if 'tos' not in self._properties or self._properties['tos'] is None:
            self._properties['tos'] = FlowIpv4Tos()
        return self._properties['tos']

    @property
    def dscp(self):
        """flow_ipv4_priority.dscp getter

        TBD  

        Returns: obj(snappi.FlowIpv4Dscp)
        """
        from .flowipv4dscp import FlowIpv4Dscp
        if 'dscp' not in self._properties or self._properties['dscp'] is None:
            self._properties['dscp'] = FlowIpv4Dscp()
        return self._properties['dscp']

    def set(self, choice=None):
        self.choice = choice
        return self
