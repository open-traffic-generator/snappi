from .snappicommon import SnappiObject


class DeviceRandom(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def min(self):
        """device_random.min getter

        TBD  

        Returns: str
        """
        return self._properties['min']

    @min.setter
    def min(self, value):
        """device_random.min setter

        TBD  

        value: str
        """
        self._properties['min'] = value

    @property
    def max(self):
        """device_random.max getter

        TBD  

        Returns: str
        """
        return self._properties['max']

    @max.setter
    def max(self, value):
        """device_random.max setter

        TBD  

        value: str
        """
        self._properties['max'] = value

    @property
    def step(self):
        """device_random.step getter

        TBD  

        Returns: float
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """device_random.step setter

        TBD  

        value: float
        """
        self._properties['step'] = value

    @property
    def seed(self):
        """device_random.seed getter

        TBD  

        Returns: str
        """
        return self._properties['seed']

    @seed.setter
    def seed(self, value):
        """device_random.seed setter

        TBD  

        value: str
        """
        self._properties['seed'] = value

    def set(self, min=None, max=None, step=None, seed=None):
        self.min = min
        self.max = max
        self.step = step
        self.seed = seed
        return self
