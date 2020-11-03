from .snappicommon import SnappiList


class DeviceBgpv6RouteRangeList(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, route_count_per_device=1):
        from .devicebgpv6routerange import DeviceBgpv6RouteRange
        self._items.append(DeviceBgpv6RouteRange().set(name, route_count_per_device))
        return self
