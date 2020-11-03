from .snappicommon import SnappiList


class LagList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, port_names=None):
        from .lag import Lag
        self._items.append(Lag().set(name, port_names))
        return self
