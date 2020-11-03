from .snappicommon import SnappiObject


class Config(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def ports(self):
        """config.ports getter

        The ports that will be configured on the traffic generator  

        Returns: list[obj(snappi.Port)]
        """
        from .portlist import PortList
        if 'ports' not in self._properties or self._properties['ports'] is None:
            self._properties['ports'] = PortList()
        return self._properties['ports']

    @property
    def lags(self):
        """config.lags getter

        The lags that will be configured on the traffic generator  

        Returns: list[obj(snappi.Lag)]
        """
        from .laglist import LagList
        if 'lags' not in self._properties or self._properties['lags'] is None:
            self._properties['lags'] = LagList()
        return self._properties['lags']

    @property
    def layer1(self):
        """config.layer1 getter

        The layer1 settings that will be configured on the traffic generator  

        Returns: list[obj(snappi.Layer1)]
        """
        from .layer1list import Layer1List
        if 'layer1' not in self._properties or self._properties['layer1'] is None:
            self._properties['layer1'] = Layer1List()
        return self._properties['layer1']

    @property
    def captures(self):
        """config.captures getter

        The capture settings that will be configured on the traffic generator  

        Returns: list[obj(snappi.Capture)]
        """
        from .capturelist import CaptureList
        if 'captures' not in self._properties or self._properties['captures'] is None:
            self._properties['captures'] = CaptureList()
        return self._properties['captures']

    @property
    def devices(self):
        """config.devices getter

        The emulated device settings that will be configured on the traffic generator  

        Returns: list[obj(snappi.Device)]
        """
        from .devicelist import DeviceList
        if 'devices' not in self._properties or self._properties['devices'] is None:
            self._properties['devices'] = DeviceList()
        return self._properties['devices']

    @property
    def flows(self):
        """config.flows getter

        The flows that will be configured on the traffic generator  

        Returns: list[obj(snappi.Flow)]
        """
        from .flowlist import FlowList
        if 'flows' not in self._properties or self._properties['flows'] is None:
            self._properties['flows'] = FlowList()
        return self._properties['flows']

    @property
    def options(self):
        """config.options getter

        TBD  

        Returns: obj(snappi.ConfigOptions)
        """
        from .configoptions import ConfigOptions
        if 'options' not in self._properties or self._properties['options'] is None:
            self._properties['options'] = ConfigOptions()
        return self._properties['options']

    def set(self):
        return self
