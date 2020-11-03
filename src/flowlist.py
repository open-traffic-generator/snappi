from .snappicommon import SnappiList


class FlowList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None):
        from .flow import Flow
        self._items.append(Flow().set(name))
        return self
