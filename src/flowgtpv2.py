from .snappicommon import SnappiObject


class FlowGtpv2(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def version(self):
        """flow_gtpv2.version getter

        It is a 3-bit field  
        For GTPv2, this has a value of 2  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'version' not in self._properties or self._properties['version'] is None:
            self._properties['version'] = FlowPattern()
        return self._properties['version']

    @property
    def piggybacking_flag(self):
        """flow_gtpv2.piggybacking_flag getter

        If this bit is set to 1 then another GTP-C message with its own header shall be present at the end of the current message  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'piggybacking_flag' not in self._properties or self._properties['piggybacking_flag'] is None:
            self._properties['piggybacking_flag'] = FlowPattern()
        return self._properties['piggybacking_flag']

    @property
    def teid_flag(self):
        """flow_gtpv2.teid_flag getter

        If this bit is set to 1 then the TEID field will be present between the message length and the sequence number  
        All messages except Echo and Echo reply require TEID to be present  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'teid_flag' not in self._properties or self._properties['teid_flag'] is None:
            self._properties['teid_flag'] = FlowPattern()
        return self._properties['teid_flag']

    @property
    def spare1(self):
        """flow_gtpv2.spare1 getter

        A 3-bit reserved field (must be 0)  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'spare1' not in self._properties or self._properties['spare1'] is None:
            self._properties['spare1'] = FlowPattern()
        return self._properties['spare1']

    @property
    def message_type(self):
        """flow_gtpv2.message_type getter

        An 8-bit field that indicates the type of GTP message  
        Different types of messages are defined in 3GPP TS 29.060 section 7.1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'message_type' not in self._properties or self._properties['message_type'] is None:
            self._properties['message_type'] = FlowPattern()
        return self._properties['message_type']

    @property
    def message_length(self):
        """flow_gtpv2.message_length getter

        A 16-bit field that indicates the length of the payload in bytes (excluding the mandatory GTP-c header (first 4 bytes)  
        Includes the TEID and sequence_number if they are present  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'message_length' not in self._properties or self._properties['message_length'] is None:
            self._properties['message_length'] = FlowPattern()
        return self._properties['message_length']

    @property
    def teid(self):
        """flow_gtpv2.teid getter

        Tunnel endpoint identifier  
        A 32-bit (4-octet) field used to multiplex different connections in the same GTP tunnel  
        Is present only if the teid_flag is set  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'teid' not in self._properties or self._properties['teid'] is None:
            self._properties['teid'] = FlowPattern()
        return self._properties['teid']

    @property
    def sequence_number(self):
        """flow_gtpv2.sequence_number getter

        A 24-bit field  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'sequence_number' not in self._properties or self._properties['sequence_number'] is None:
            self._properties['sequence_number'] = FlowPattern()
        return self._properties['sequence_number']

    @property
    def spare2(self):
        """flow_gtpv2.spare2 getter

        An 8-bit reserved field (must be 0)  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'spare2' not in self._properties or self._properties['spare2'] is None:
            self._properties['spare2'] = FlowPattern()
        return self._properties['spare2']

    def set(self):
        return self
