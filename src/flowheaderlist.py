from .snappicommon import SnappiList


class FlowHeaderList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, choice=None):
        from .flowheader import FlowHeader
        self._items.append(FlowHeader().set(choice))
        return self
