from .snappicommon import SnappiList


class DeviceBgpv4RouteRangeList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, route_count_per_device=1):
        from .devicebgpv4routerange import DeviceBgpv4RouteRange
        self._items.append(DeviceBgpv4RouteRange().set(name, route_count_per_device))
        return self
