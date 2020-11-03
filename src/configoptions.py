from .snappicommon import SnappiObject


class ConfigOptions(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def port_options(self):
        """config_options.port_options getter

        TBD  

        Returns: obj(snappi.PortOptions)
        """
        from .portoptions import PortOptions
        if 'port_options' not in self._properties or self._properties['port_options'] is None:
            self._properties['port_options'] = PortOptions()
        return self._properties['port_options']

    def set(self):
        return self
