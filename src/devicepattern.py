from .snappicommon import SnappiObject


class DevicePattern(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """device_pattern.choice getter

        TBD  

        Returns: Union[fixed, list, counter, random]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """device_pattern.choice setter

        TBD  

        value: Union[fixed, list, counter, random]
        """
        self._properties['choice'] = value

    @property
    def fixed(self):
        """device_pattern.fixed getter

        TBD  

        Returns: str
        """
        return self._properties['fixed']

    @fixed.setter
    def fixed(self, value):
        """device_pattern.fixed setter

        TBD  

        value: str
        """
        self._properties['fixed'] = value

    @property
    def list(self):
        """device_pattern.list getter

        TBD  

        Returns: list[str]
        """
        return self._properties['list']

    @list.setter
    def list(self, value):
        """device_pattern.list setter

        TBD  

        value: list[str]
        """
        self._properties['list'] = value

    @property
    def counter(self):
        """device_pattern.counter getter

        TBD  

        Returns: obj(snappi.DeviceCounter)
        """
        from .devicecounter import DeviceCounter
        if 'counter' not in self._properties or self._properties['counter'] is None:
            self._properties['counter'] = DeviceCounter()
        return self._properties['counter']

    @property
    def random(self):
        """device_pattern.random getter

        TBD  

        Returns: obj(snappi.DeviceRandom)
        """
        from .devicerandom import DeviceRandom
        if 'random' not in self._properties or self._properties['random'] is None:
            self._properties['random'] = DeviceRandom()
        return self._properties['random']

    def set(self, choice=None, fixed=None, list=None):
        self.choice = choice
        self.fixed = fixed
        self.list = list
        return self
