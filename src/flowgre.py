from .snappicommon import SnappiObject


class FlowGre(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def checksum_present(self):
        """flow_gre.checksum_present getter

        Checksum bit  
        Set to 1 if a checksum is present  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'checksum_present' not in self._properties or self._properties['checksum_present'] is None:
            self._properties['checksum_present'] = FlowPattern()
        return self._properties['checksum_present']

    @property
    def key_present(self):
        """flow_gre.key_present getter

        Key bit  
        Set to 1 if a key is present  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'key_present' not in self._properties or self._properties['key_present'] is None:
            self._properties['key_present'] = FlowPattern()
        return self._properties['key_present']

    @property
    def seq_number_present(self):
        """flow_gre.seq_number_present getter

        Sequence number bit  
        Set to 1 if a sequence number is present  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'seq_number_present' not in self._properties or self._properties['seq_number_present'] is None:
            self._properties['seq_number_present'] = FlowPattern()
        return self._properties['seq_number_present']

    @property
    def reserved0(self):
        """flow_gre.reserved0 getter

        Reserved bits  
        Set to 0  
        9 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved0' not in self._properties or self._properties['reserved0'] is None:
            self._properties['reserved0'] = FlowPattern()
        return self._properties['reserved0']

    @property
    def version(self):
        """flow_gre.version getter

        Gre version number  
        Set to 0  
        3 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'version' not in self._properties or self._properties['version'] is None:
            self._properties['version'] = FlowPattern()
        return self._properties['version']

    @property
    def protocol(self):
        """flow_gre.protocol getter

        Indicates the ether protocol type of the encapsulated payload  
        - 0x0800 ipv4 - 0x86DD ipv6  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'protocol' not in self._properties or self._properties['protocol'] is None:
            self._properties['protocol'] = FlowPattern()
        return self._properties['protocol']

    @property
    def checksum(self):
        """flow_gre.checksum getter

        Present if the checksum_present bit is set  
        Contains the checksum for the gre header and payload  
        16 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'checksum' not in self._properties or self._properties['checksum'] is None:
            self._properties['checksum'] = FlowPattern()
        return self._properties['checksum']

    @property
    def reserved1(self):
        """flow_gre.reserved1 getter

        Reserved bits  
        Set to 0  
        16 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved1' not in self._properties or self._properties['reserved1'] is None:
            self._properties['reserved1'] = FlowPattern()
        return self._properties['reserved1']

    @property
    def key(self):
        """flow_gre.key getter

        Present if the key_present bit is set  
        Contains an application specific key value  
        32 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'key' not in self._properties or self._properties['key'] is None:
            self._properties['key'] = FlowPattern()
        return self._properties['key']

    @property
    def sequence_number(self):
        """flow_gre.sequence_number getter

        Present if the seq_number_present bit is set  
        Contains a sequence number for the gre packet  
        32 bits  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'sequence_number' not in self._properties or self._properties['sequence_number'] is None:
            self._properties['sequence_number'] = FlowPattern()
        return self._properties['sequence_number']

    def set(self):
        return self
