from .snappicommon import SnappiObject


class FlowTcp(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def src_port(self):
        """flow_tcp.src_port getter

        Tcp source port  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src_port' not in self._properties or self._properties['src_port'] is None:
            self._properties['src_port'] = FlowPattern()
        return self._properties['src_port']

    @property
    def dst_port(self):
        """flow_tcp.dst_port getter

        Tcp destination port  
        Max length is 2 bytes  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst_port' not in self._properties or self._properties['dst_port'] is None:
            self._properties['dst_port'] = FlowPattern()
        return self._properties['dst_port']

    @property
    def ecn_ns(self):
        """flow_tcp.ecn_ns getter

        Explicit congestion notification, concealment protection  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ecn_ns' not in self._properties or self._properties['ecn_ns'] is None:
            self._properties['ecn_ns'] = FlowPattern()
        return self._properties['ecn_ns']

    @property
    def ecn_cwr(self):
        """flow_tcp.ecn_cwr getter

        Explicit congestion notification, congestion window reduced  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ecn_cwr' not in self._properties or self._properties['ecn_cwr'] is None:
            self._properties['ecn_cwr'] = FlowPattern()
        return self._properties['ecn_cwr']

    @property
    def ecn_echo(self):
        """flow_tcp.ecn_echo getter

        Explicit congestion notification, echo  
        1 indicates the peer is ecn capable  
        0 indicates that a packet with ipv4.ecn = 11 in the ip header was received during normal transmission  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ecn_echo' not in self._properties or self._properties['ecn_echo'] is None:
            self._properties['ecn_echo'] = FlowPattern()
        return self._properties['ecn_echo']

    @property
    def ctl_urg(self):
        """flow_tcp.ctl_urg getter

        The urgent pointer field is significant  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_urg' not in self._properties or self._properties['ctl_urg'] is None:
            self._properties['ctl_urg'] = FlowPattern()
        return self._properties['ctl_urg']

    @property
    def ctl_ack(self):
        """flow_tcp.ctl_ack getter

        The ackknowledgment field is significant  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_ack' not in self._properties or self._properties['ctl_ack'] is None:
            self._properties['ctl_ack'] = FlowPattern()
        return self._properties['ctl_ack']

    @property
    def ctl_psh(self):
        """flow_tcp.ctl_psh getter

        Asks to push the buffered data to the receiving application  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_psh' not in self._properties or self._properties['ctl_psh'] is None:
            self._properties['ctl_psh'] = FlowPattern()
        return self._properties['ctl_psh']

    @property
    def ctl_rst(self):
        """flow_tcp.ctl_rst getter

        Reset the connection  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_rst' not in self._properties or self._properties['ctl_rst'] is None:
            self._properties['ctl_rst'] = FlowPattern()
        return self._properties['ctl_rst']

    @property
    def ctl_syn(self):
        """flow_tcp.ctl_syn getter

        Synchronize sequenece numbers  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_syn' not in self._properties or self._properties['ctl_syn'] is None:
            self._properties['ctl_syn'] = FlowPattern()
        return self._properties['ctl_syn']

    @property
    def ctl_fin(self):
        """flow_tcp.ctl_fin getter

        Last packet from the sender  
        Max length is 1 bit  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ctl_fin' not in self._properties or self._properties['ctl_fin'] is None:
            self._properties['ctl_fin'] = FlowPattern()
        return self._properties['ctl_fin']

    def set(self):
        return self
