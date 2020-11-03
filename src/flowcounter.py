from .snappicommon import SnappiObject


class FlowCounter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def start(self):
        """flow_counter.start getter

        The value at which the pattern will start  

        Returns: str
        """
        return self._properties['start']

    @start.setter
    def start(self, value):
        """flow_counter.start setter

        The value at which the pattern will start  

        value: str
        """
        self._properties['start'] = value

    @property
    def step(self):
        """flow_counter.step getter

        The value at which the pattern will increment or decrement by  

        Returns: str
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """flow_counter.step setter

        The value at which the pattern will increment or decrement by  

        value: str
        """
        self._properties['step'] = value

    @property
    def count(self):
        """flow_counter.count getter

        The number of values in the pattern  

        Returns: float
        """
        return self._properties['count']

    @count.setter
    def count(self, value):
        """flow_counter.count setter

        The number of values in the pattern  

        value: float
        """
        self._properties['count'] = value

    @property
    def up(self):
        """flow_counter.up getter

        Determines whether the pattern will increment (true) or decrement (false)  

        Returns: boolean
        """
        return self._properties['up']

    @up.setter
    def up(self, value):
        """flow_counter.up setter

        Determines whether the pattern will increment (true) or decrement (false)  

        value: boolean
        """
        self._properties['up'] = value

    def set(self, start=None, step=None, count=None, up=None):
        self.start = start
        self.step = step
        self.count = count
        self.up = up
        return self
