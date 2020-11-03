from .snappicommon import SnappiList


class CaptureBasicFilterList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, choice=None, and_operator=True, not_operator=False):
        from .capturebasicfilter import CaptureBasicFilter
        self._items.append(CaptureBasicFilter().set(choice, and_operator, not_operator))
        return self
