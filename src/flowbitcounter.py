from .snappicommon import SnappiObject


class FlowBitCounter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def offset(self):
        """flow_bitcounter.offset getter

        Bit offset in the packet at which the pattern will be applied  

        Returns: int
        """
        return self._properties['offset']

    @offset.setter
    def offset(self, value):
        """flow_bitcounter.offset setter

        Bit offset in the packet at which the pattern will be applied  

        value: int
        """
        self._properties['offset'] = value

    @property
    def length(self):
        """flow_bitcounter.length getter

        The number of bits in the packet that the pattern will span  

        Returns: int
        """
        return self._properties['length']

    @length.setter
    def length(self, value):
        """flow_bitcounter.length setter

        The number of bits in the packet that the pattern will span  

        value: int
        """
        self._properties['length'] = value

    @property
    def count(self):
        """flow_bitcounter.count getter

        The number of values to generate before repeating A value of 0 means the pattern will count continuously  

        Returns: int
        """
        return self._properties['count']

    @count.setter
    def count(self, value):
        """flow_bitcounter.count setter

        The number of values to generate before repeating A value of 0 means the pattern will count continuously  

        value: int
        """
        self._properties['count'] = value

    @property
    def start(self):
        """flow_bitcounter.start getter

        The starting value of the pattern  
        If the value is greater than the length it will be truncated  

        Returns: str
        """
        return self._properties['start']

    @start.setter
    def start(self, value):
        """flow_bitcounter.start setter

        The starting value of the pattern  
        If the value is greater than the length it will be truncated  

        value: str
        """
        self._properties['start'] = value

    @property
    def step(self):
        """flow_bitcounter.step getter

        The amount the start value will be incremented by If the value is greater than the length it will be truncated  

        Returns: str
        """
        return self._properties['step']

    @step.setter
    def step(self, value):
        """flow_bitcounter.step setter

        The amount the start value will be incremented by If the value is greater than the length it will be truncated  

        value: str
        """
        self._properties['step'] = value

    def set(self, offset=None, length=None, count=None, start=None, step=None):
        self.offset = offset
        self.length = length
        self.count = count
        self.start = start
        self.step = step
        return self
