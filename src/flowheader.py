from .snappicommon import SnappiObject


class FlowHeader(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_header.choice getter

        TBD  

        Returns: Union[custom, ethernet, vlan, vxlan, ipv4, ipv6, pfcpause, ethernetpause, tcp, udp, gre, gtpv1, gtpv2]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_header.choice setter

        TBD  

        value: Union[custom, ethernet, vlan, vxlan, ipv4, ipv6, pfcpause, ethernetpause, tcp, udp, gre, gtpv1, gtpv2]
        """
        self._properties['choice'] = value

    @property
    def custom(self):
        """flow_header.custom getter

        TBD  

        Returns: obj(snappi.FlowCustom)
        """
        from .flowcustom import FlowCustom
        if 'custom' not in self._properties or self._properties['custom'] is None:
            self._properties['custom'] = FlowCustom()
        return self._properties['custom']

    @property
    def ethernet(self):
        """flow_header.ethernet getter

        TBD  

        Returns: obj(snappi.FlowEthernet)
        """
        from .flowethernet import FlowEthernet
        if 'ethernet' not in self._properties or self._properties['ethernet'] is None:
            self._properties['ethernet'] = FlowEthernet()
        return self._properties['ethernet']

    @property
    def vlan(self):
        """flow_header.vlan getter

        TBD  

        Returns: obj(snappi.FlowVlan)
        """
        from .flowvlan import FlowVlan
        if 'vlan' not in self._properties or self._properties['vlan'] is None:
            self._properties['vlan'] = FlowVlan()
        return self._properties['vlan']

    @property
    def vxlan(self):
        """flow_header.vxlan getter

        TBD  

        Returns: obj(snappi.FlowVxlan)
        """
        from .flowvxlan import FlowVxlan
        if 'vxlan' not in self._properties or self._properties['vxlan'] is None:
            self._properties['vxlan'] = FlowVxlan()
        return self._properties['vxlan']

    @property
    def ipv4(self):
        """flow_header.ipv4 getter

        TBD  

        Returns: obj(snappi.FlowIpv4)
        """
        from .flowipv4 import FlowIpv4
        if 'ipv4' not in self._properties or self._properties['ipv4'] is None:
            self._properties['ipv4'] = FlowIpv4()
        return self._properties['ipv4']

    @property
    def ipv6(self):
        """flow_header.ipv6 getter

        TBD  

        Returns: obj(snappi.FlowIpv6)
        """
        from .flowipv6 import FlowIpv6
        if 'ipv6' not in self._properties or self._properties['ipv6'] is None:
            self._properties['ipv6'] = FlowIpv6()
        return self._properties['ipv6']

    @property
    def pfcpause(self):
        """flow_header.pfcpause getter

        TBD  

        Returns: obj(snappi.FlowPfcPause)
        """
        from .flowpfcpause import FlowPfcPause
        if 'pfcpause' not in self._properties or self._properties['pfcpause'] is None:
            self._properties['pfcpause'] = FlowPfcPause()
        return self._properties['pfcpause']

    @property
    def ethernetpause(self):
        """flow_header.ethernetpause getter

        TBD  

        Returns: obj(snappi.FlowEthernetPause)
        """
        from .flowethernetpause import FlowEthernetPause
        if 'ethernetpause' not in self._properties or self._properties['ethernetpause'] is None:
            self._properties['ethernetpause'] = FlowEthernetPause()
        return self._properties['ethernetpause']

    @property
    def tcp(self):
        """flow_header.tcp getter

        TBD  

        Returns: obj(snappi.FlowTcp)
        """
        from .flowtcp import FlowTcp
        if 'tcp' not in self._properties or self._properties['tcp'] is None:
            self._properties['tcp'] = FlowTcp()
        return self._properties['tcp']

    @property
    def udp(self):
        """flow_header.udp getter

        TBD  

        Returns: obj(snappi.FlowUdp)
        """
        from .flowudp import FlowUdp
        if 'udp' not in self._properties or self._properties['udp'] is None:
            self._properties['udp'] = FlowUdp()
        return self._properties['udp']

    @property
    def gre(self):
        """flow_header.gre getter

        TBD  

        Returns: obj(snappi.FlowGre)
        """
        from .flowgre import FlowGre
        if 'gre' not in self._properties or self._properties['gre'] is None:
            self._properties['gre'] = FlowGre()
        return self._properties['gre']

    @property
    def gtpv1(self):
        """flow_header.gtpv1 getter

        TBD  

        Returns: obj(snappi.FlowGtpv1)
        """
        from .flowgtpv1 import FlowGtpv1
        if 'gtpv1' not in self._properties or self._properties['gtpv1'] is None:
            self._properties['gtpv1'] = FlowGtpv1()
        return self._properties['gtpv1']

    @property
    def gtpv2(self):
        """flow_header.gtpv2 getter

        TBD  

        Returns: obj(snappi.FlowGtpv2)
        """
        from .flowgtpv2 import FlowGtpv2
        if 'gtpv2' not in self._properties or self._properties['gtpv2'] is None:
            self._properties['gtpv2'] = FlowGtpv2()
        return self._properties['gtpv2']

    def set(self, choice=None):
        self.choice = choice
        return self
