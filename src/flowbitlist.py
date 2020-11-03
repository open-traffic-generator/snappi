from .snappicommon import SnappiObject


class FlowBitList(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def offset(self):
        """flow_bitlist.offset getter

        Bit offset in the packet at which the pattern will be applied  

        Returns: int
        """
        return self._properties['offset']

    @offset.setter
    def offset(self, value):
        """flow_bitlist.offset setter

        Bit offset in the packet at which the pattern will be applied  

        value: int
        """
        self._properties['offset'] = value

    @property
    def length(self):
        """flow_bitlist.length getter

        The number of bits in the packet that the pattern will span  

        Returns: int
        """
        return self._properties['length']

    @length.setter
    def length(self, value):
        """flow_bitlist.length setter

        The number of bits in the packet that the pattern will span  

        value: int
        """
        self._properties['length'] = value

    @property
    def count(self):
        """flow_bitlist.count getter

        The number of values to generate before repeating  

        Returns: int
        """
        return self._properties['count']

    @count.setter
    def count(self, value):
        """flow_bitlist.count setter

        The number of values to generate before repeating  

        value: int
        """
        self._properties['count'] = value

    @property
    def values(self):
        """flow_bitlist.values getter

        TBD  

        Returns: list[str]
        """
        return self._properties['values']

    @values.setter
    def values(self, value):
        """flow_bitlist.values setter

        TBD  

        value: list[str]
        """
        self._properties['values'] = value

    def set(self, offset=None, length=None, count=None, values=None):
        self.offset = offset
        self.length = length
        self.count = count
        self.values = values
        return self
