from .snappicommon import SnappiObject


class DeviceCounter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def start(self):
        """device_counter.start getter

        TBD  

        Returns: str
        """
        return self._properties['start']

    @start.setter
    def start(self, value):
        """device_counter.start setter

        TBD  

        value: str
        """
        self._properties['start'] = value

    @property
    def step(self):
        """device_counter.step getter

        TBD  

        Returns: str
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """device_counter.step setter

        TBD  

        value: str
        """
        self._properties['step'] = value

    @property
    def up(self):
        """device_counter.up getter

        TBD  

        Returns: boolean
        """
        return self._properties['up']

    @up.setter
    def up(self, value):
        """device_counter.up setter

        TBD  

        value: boolean
        """
        self._properties['up'] = value

    def set(self, start=None, step=None, up=None):
        self.start = start
        self.step = step
        self.up = up
        return self
