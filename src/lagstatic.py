from .snappicommon import SnappiObject


class LagStatic(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def lag_id(self):
        """lag_static.lag_id getter

        The static lag id  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'lag_id' not in self._properties or self._properties['lag_id'] is None:
            self._properties['lag_id'] = DevicePattern()
        return self._properties['lag_id']

    def set(self):
        return self
