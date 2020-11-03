from .snappicommon import SnappiList


class DeviceList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, container_name=None, device_count=1, choice=None):
        from .device import Device
        self._items.append(Device().set(name, container_name, device_count, choice))
        return self
