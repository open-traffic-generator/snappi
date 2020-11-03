from .snappicommon import SnappiList


class CaptureList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, port_names=None, choice=None, pcap=None, enable=True, overwrite=False, format=pcap):
        from .capture import Capture
        self._items.append(Capture().set(name, port_names, choice, pcap, enable, overwrite, format))
        return self
