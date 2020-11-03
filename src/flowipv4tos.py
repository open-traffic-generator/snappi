from .snappicommon import SnappiObject


class FlowIpv4Tos(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def precedence(self):
        """flow_ipv4_tos.precedence getter

        Precedence value is 3 bits: >=0 precedence <=3  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'precedence' not in self._properties or self._properties['precedence'] is None:
            self._properties['precedence'] = FlowPattern()
        return self._properties['precedence']

    @property
    def delay(self):
        """flow_ipv4_tos.delay getter

        Delay value is 1 bit: >=0 delay <=1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'delay' not in self._properties or self._properties['delay'] is None:
            self._properties['delay'] = FlowPattern()
        return self._properties['delay']

    @property
    def throughput(self):
        """flow_ipv4_tos.throughput getter

        Throughput value is 1 bit: >=0 throughput <=3  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'throughput' not in self._properties or self._properties['throughput'] is None:
            self._properties['throughput'] = FlowPattern()
        return self._properties['throughput']

    @property
    def reliability(self):
        """flow_ipv4_tos.reliability getter

        Reliability value is 1 bit: >=0 reliability <=1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reliability' not in self._properties or self._properties['reliability'] is None:
            self._properties['reliability'] = FlowPattern()
        return self._properties['reliability']

    @property
    def monetary(self):
        """flow_ipv4_tos.monetary getter

        Monetary value is 1 bit: >=0 monetary <=1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'monetary' not in self._properties or self._properties['monetary'] is None:
            self._properties['monetary'] = FlowPattern()
        return self._properties['monetary']

    @property
    def unused(self):
        """flow_ipv4_tos.unused getter

        Unused value is 1 bit: >=0 unused <=1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'unused' not in self._properties or self._properties['unused'] is None:
            self._properties['unused'] = FlowPattern()
        return self._properties['unused']

    def set(self):
        return self
