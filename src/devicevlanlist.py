from .snappicommon import SnappiList


class DeviceVlanList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None):
        from .devicevlan import DeviceVlan
        self._items.append(DeviceVlan().set(name))
        return self
