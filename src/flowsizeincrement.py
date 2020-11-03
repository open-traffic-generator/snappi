from .snappicommon import SnappiObject


class FlowSizeIncrement(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def start(self):
        """flow_sizeincrement.start getter

        Starting frame size in bytes  

        Returns: int
        """
        return self._properties['start']

    @start.setter
    def start(self, value):
        """flow_sizeincrement.start setter

        Starting frame size in bytes  

        value: int
        """
        self._properties['start'] = value

    @property
    def end(self):
        """flow_sizeincrement.end getter

        Ending frame size in bytes  

        Returns: int
        """
        return self._properties['end']

    @end.setter
    def end(self, value):
        """flow_sizeincrement.end setter

        Ending frame size in bytes  

        value: int
        """
        self._properties['end'] = value

    @property
    def step(self):
        """flow_sizeincrement.step getter

        Step frame size in bytes  

        Returns: int
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """flow_sizeincrement.step setter

        Step frame size in bytes  

        value: int
        """
        self._properties['step'] = value

    def set(self, start=None, end=None, step=None):
        self.start = start
        self.end = end
        self.step = step
        return self
