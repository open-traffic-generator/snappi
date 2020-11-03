from .snappicommon import SnappiObject


class FlowGtpv1(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def version(self):
        """flow_gtpv1.version getter

        It is a 3-bit field  
        For GTPv1, this has a value of 1  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'version' not in self._properties or self._properties['version'] is None:
            self._properties['version'] = FlowPattern()
        return self._properties['version']

    @property
    def protocol_type(self):
        """flow_gtpv1.protocol_type getter

        A 1-bit value that differentiates GTP (value 1) from GTP' (value 0)  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'protocol_type' not in self._properties or self._properties['protocol_type'] is None:
            self._properties['protocol_type'] = FlowPattern()
        return self._properties['protocol_type']

    @property
    def reserved(self):
        """flow_gtpv1.reserved getter

        A 1-bit reserved field (must be 0)  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'reserved' not in self._properties or self._properties['reserved'] is None:
            self._properties['reserved'] = FlowPattern()
        return self._properties['reserved']

    @property
    def e_flag(self):
        """flow_gtpv1.e_flag getter

        A 1-bit value that states whether there is an extension header optional field  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'e_flag' not in self._properties or self._properties['e_flag'] is None:
            self._properties['e_flag'] = FlowPattern()
        return self._properties['e_flag']

    @property
    def s_flag(self):
        """flow_gtpv1.s_flag getter

        A 1-bit value that states whether there is a Sequence Number optional field  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 's_flag' not in self._properties or self._properties['s_flag'] is None:
            self._properties['s_flag'] = FlowPattern()
        return self._properties['s_flag']

    @property
    def pn_flag(self):
        """flow_gtpv1.pn_flag getter

        A 1-bit value that states whether there is a N-PDU number optional field  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pn_flag' not in self._properties or self._properties['pn_flag'] is None:
            self._properties['pn_flag'] = FlowPattern()
        return self._properties['pn_flag']

    @property
    def message_type(self):
        """flow_gtpv1.message_type getter

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
        """flow_gtpv1.message_length getter

        A 16-bit field that indicates the length of the payload in bytes (rest of the packet following the mandatory 8-byte GTP header)  
        Includes the optional fields  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'message_length' not in self._properties or self._properties['message_length'] is None:
            self._properties['message_length'] = FlowPattern()
        return self._properties['message_length']

    @property
    def teid(self):
        """flow_gtpv1.teid getter

        Tunnel endpoint identifier  
        A 32-bit(4-octet) field used to multiplex different connections in the same GTP tunnel  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'teid' not in self._properties or self._properties['teid'] is None:
            self._properties['teid'] = FlowPattern()
        return self._properties['teid']

    @property
    def squence_number(self):
        """flow_gtpv1.squence_number getter

        An (optional) 16-bit field  
        This field exists if any of the e_flag, s_flag, or pn_flag bits are on  
        The field must be interpreted only if the s_flag bit is on  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'squence_number' not in self._properties or self._properties['squence_number'] is None:
            self._properties['squence_number'] = FlowPattern()
        return self._properties['squence_number']

    @property
    def n_pdu_number(self):
        """flow_gtpv1.n_pdu_number getter

        An (optional) 8-bit field  
        This field exists if any of the e_flag, s_flag, or pn_flag bits are on  
        The field must be interpreted only if the pn_flag bit is on  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'n_pdu_number' not in self._properties or self._properties['n_pdu_number'] is None:
            self._properties['n_pdu_number'] = FlowPattern()
        return self._properties['n_pdu_number']

    @property
    def next_extension_header_type(self):
        """flow_gtpv1.next_extension_header_type getter

        An (optional) 8-bit field  
        This field exists if any of the e_flag, s_flag, or pn_flag bits are on  
        The field must be interpreted only if the e_flag bit is on  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'next_extension_header_type' not in self._properties or self._properties['next_extension_header_type'] is None:
            self._properties['next_extension_header_type'] = FlowPattern()
        return self._properties['next_extension_header_type']

    @property
    def extension_headers(self):
        """flow_gtpv1.extension_headers getter

        A list of optional extension headers  

        Returns: list[obj(snappi.FlowGtpExtension)]
        """
        from .flowgtpextensionlist import FlowGtpExtensionList
        if 'extension_headers' not in self._properties or self._properties['extension_headers'] is None:
            self._properties['extension_headers'] = FlowGtpExtensionList()
        return self._properties['extension_headers']

    def set(self):
        return self
