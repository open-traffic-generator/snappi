from .snappicommon import SnappiObject


class CaptureCustomFilter(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def filter(self):
        """capture_customfilter.filter getter

        The value to filter on  

        Returns: str
        """
        return self._properties['filter']

    @filter.setter
    def filter(self, value):
        """capture_customfilter.filter setter

        The value to filter on  

        value: str
        """
        self._properties['filter'] = value

    @property
    def mask(self):
        """capture_customfilter.mask getter

        The mask to be applied to the filter  

        Returns: str
        """
        return self._properties['mask']

    @mask.setter
    def mask(self, value):
        """capture_customfilter.mask setter

        The mask to be applied to the filter  

        value: str
        """
        self._properties['mask'] = value

    @property
    def offset(self):
        """capture_customfilter.offset getter

        The offset in the packet to filter at  

        Returns: int
        """
        return self._properties['offset']

    @offset.setter
    def offset(self, value):
        """capture_customfilter.offset setter

        The offset in the packet to filter at  

        value: int
        """
        self._properties['offset'] = value

    def set(self, filter=None, mask=None, offset=None):
        self.filter = filter
        self.mask = mask
        self.offset = offset
        return self
