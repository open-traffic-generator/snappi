from .snappicommon import SnappiList


class PortList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, location=None):
        from .port import Port
        self._items.append(Port().set(name, location))
        return self
