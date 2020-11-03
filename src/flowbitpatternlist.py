from .snappicommon import SnappiList


class FlowBitPatternList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, choice=None):
        from .flowbitpattern import FlowBitPattern
        self._items.append(FlowBitPattern().set(choice))
        return self
