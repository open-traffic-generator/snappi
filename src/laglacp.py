from .snappicommon import SnappiObject


class LagLacp(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def actor_key(self):
        """lag_lacp.actor_key getter

        The actor key  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'actor_key' not in self._properties or self._properties['actor_key'] is None:
            self._properties['actor_key'] = DevicePattern()
        return self._properties['actor_key']

    @property
    def actor_port_number(self):
        """lag_lacp.actor_port_number getter

        The actor port number  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'actor_port_number' not in self._properties or self._properties['actor_port_number'] is None:
            self._properties['actor_port_number'] = DevicePattern()
        return self._properties['actor_port_number']

    @property
    def actor_port_priority(self):
        """lag_lacp.actor_port_priority getter

        The actor port priority  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'actor_port_priority' not in self._properties or self._properties['actor_port_priority'] is None:
            self._properties['actor_port_priority'] = DevicePattern()
        return self._properties['actor_port_priority']

    @property
    def actor_system_id(self):
        """lag_lacp.actor_system_id getter

        The actor system id  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'actor_system_id' not in self._properties or self._properties['actor_system_id'] is None:
            self._properties['actor_system_id'] = DevicePattern()
        return self._properties['actor_system_id']

    @property
    def actor_system_priority(self):
        """lag_lacp.actor_system_priority getter

        The actor system priority  

        Returns: obj(snappi.DevicePattern)
        """
        from .devicepattern import DevicePattern
        if 'actor_system_priority' not in self._properties or self._properties['actor_system_priority'] is None:
            self._properties['actor_system_priority'] = DevicePattern()
        return self._properties['actor_system_priority']

    def set(self):
        return self
