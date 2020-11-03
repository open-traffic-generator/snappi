from .snappicommon import SnappiObject


class FlowRandom(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def min(self):
        """flow_random.min getter

        TBD  

        Returns: str
        """
        return self._properties['min']

    @min.setter
    def min(self, value):
        """flow_random.min setter

        TBD  

        value: str
        """
        self._properties['min'] = value

    @property
    def max(self):
        """flow_random.max getter

        TBD  

        Returns: str
        """
        return self._properties['max']

    @max.setter
    def max(self, value):
        """flow_random.max setter

        TBD  

        value: str
        """
        self._properties['max'] = value

    @property
    def step(self):
        """flow_random.step getter

        TBD  

        Returns: float
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """flow_random.step setter

        TBD  

        value: float
        """
        self._properties['step'] = value

    @property
    def seed(self):
        """flow_random.seed getter

        TBD  

        Returns: str
        """
        return self._properties['seed']

    @seed.setter
    def seed(self, value):
        """flow_random.seed setter

        TBD  

        value: str
        """
        self._properties['seed'] = value

    @property
    def count(self):
        """flow_random.count getter

        TBD  

        Returns: float
        """
        return self._properties['count']

    @count.setter
    def count(self, value):
        """flow_random.count setter

        TBD  

        value: float
        """
        self._properties['count'] = value

    def set(self, min=None, max=None, step=None, seed=None, count=None):
        self.min = min
        self.max = max
        self.step = step
        self.seed = seed
        self.count = count
        return self
