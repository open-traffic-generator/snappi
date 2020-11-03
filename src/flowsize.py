from .snappicommon import SnappiObject


class FlowSize(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_size.choice getter

        TBD  

        Returns: Union[fixed, increment, random]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_size.choice setter

        TBD  

        value: Union[fixed, increment, random]
        """
        self._properties['choice'] = value

    @property
    def fixed(self):
        """flow_size.fixed getter

        TBD  

        Returns: int
        """
        return self._properties['fixed']

    @fixed.setter
    def fixed(self, value):
        """flow_size.fixed setter

        TBD  

        value: int
        """
        self._properties['fixed'] = value

    @property
    def increment(self):
        """flow_size.increment getter

        TBD  

        Returns: obj(snappi.FlowSizeIncrement)
        """
        from .flowsizeincrement import FlowSizeIncrement
        if 'increment' not in self._properties or self._properties['increment'] is None:
            self._properties['increment'] = FlowSizeIncrement()
        return self._properties['increment']

    @property
    def random(self):
        """flow_size.random getter

        TBD  

        Returns: obj(snappi.FlowSizeRandom)
        """
        from .flowsizerandom import FlowSizeRandom
        if 'random' not in self._properties or self._properties['random'] is None:
            self._properties['random'] = FlowSizeRandom()
        return self._properties['random']

    def set(self, choice=None, fixed=None):
        self.choice = choice
        self.fixed = fixed
        return self
