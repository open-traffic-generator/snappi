from .snappicommon import SnappiObject


class FlowSizeRandom(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def min(self):
        """flow_sizerandom.min getter

        TBD  

        Returns: int
        """
        return self._properties['min']

    @min.setter
    def min(self, value):
        """flow_sizerandom.min setter

        TBD  

        value: int
        """
        self._properties['min'] = value

    @property
    def max(self):
        """flow_sizerandom.max getter

        TBD  

        Returns: int
        """
        return self._properties['max']

    @max.setter
    def max(self, value):
        """flow_sizerandom.max setter

        TBD  

        value: int
        """
        self._properties['max'] = value

    def set(self, min=None, max=None):
        self.min = min
        self.max = max
        return self
