from .snappicommon import SnappiList


class FlowGtpExtensionList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self):
        from .flowgtpextension import FlowGtpExtension
        self._items.append(FlowGtpExtension().set())
        return self
