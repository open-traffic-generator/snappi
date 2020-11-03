from .snappicommon import SnappiObject


class FlowDuration(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_duration.choice getter

        TBD  

        Returns: Union[packets, seconds, burst, continuous]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_duration.choice setter

        TBD  

        value: Union[packets, seconds, burst, continuous]
        """
        self._properties['choice'] = value

    @property
    def packets(self):
        """flow_duration.packets getter

        TBD  

        Returns: obj(snappi.FlowFixedPackets)
        """
        from .flowfixedpackets import FlowFixedPackets
        if 'packets' not in self._properties or self._properties['packets'] is None:
            self._properties['packets'] = FlowFixedPackets()
        return self._properties['packets']

    @property
    def seconds(self):
        """flow_duration.seconds getter

        TBD  

        Returns: obj(snappi.FlowFixedSeconds)
        """
        from .flowfixedseconds import FlowFixedSeconds
        if 'seconds' not in self._properties or self._properties['seconds'] is None:
            self._properties['seconds'] = FlowFixedSeconds()
        return self._properties['seconds']

    @property
    def burst(self):
        """flow_duration.burst getter

        TBD  

        Returns: obj(snappi.FlowBurst)
        """
        from .flowburst import FlowBurst
        if 'burst' not in self._properties or self._properties['burst'] is None:
            self._properties['burst'] = FlowBurst()
        return self._properties['burst']

    @property
    def continuous(self):
        """flow_duration.continuous getter

        TBD  

        Returns: obj(snappi.FlowContinuous)
        """
        from .flowcontinuous import FlowContinuous
        if 'continuous' not in self._properties or self._properties['continuous'] is None:
            self._properties['continuous'] = FlowContinuous()
        return self._properties['continuous']

    def set(self, choice=None):
        self.choice = choice
        return self
