from .snappicommon import SnappiObject


class FlowGtpExtension(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def extension_length(self):
        """flow_gtpextension.extension_length getter

        An 8-bit field  
        This field states the length of this extension header, including the length, the contents, and the next extension header field, in 4-octet units, so the length of the extension must always be a multiple of 4  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'extension_length' not in self._properties or self._properties['extension_length'] is None:
            self._properties['extension_length'] = FlowPattern()
        return self._properties['extension_length']

    @property
    def contents(self):
        """flow_gtpextension.contents getter

        The extension header contents  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'contents' not in self._properties or self._properties['contents'] is None:
            self._properties['contents'] = FlowPattern()
        return self._properties['contents']

    @property
    def next_extension_header(self):
        """flow_gtpextension.next_extension_header getter

        An 8-bit field  
        It states the type of the next extension, or 0 if no next extension exists  
        This permits chaining several next extension headers  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'next_extension_header' not in self._properties or self._properties['next_extension_header'] is None:
            self._properties['next_extension_header'] = FlowPattern()
        return self._properties['next_extension_header']

    def set(self):
        return self
