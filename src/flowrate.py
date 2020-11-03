from .snappicommon import SnappiObject


class FlowRate(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def unit(self):
        """flow_rate.unit getter

        The value is a unit of this  

        Returns: Union[pps, bps, kbps, mbps, gbps, line]
        """
        return self._properties['unit']

    @unit.setter
    def unit(self, value):
        """flow_rate.unit setter

        The value is a unit of this  

        value: Union[pps, bps, kbps, mbps, gbps, line]
        """
        self._properties['unit'] = value

    @property
    def value(self):
        """flow_rate.value getter

        The actual rate  

        Returns: int
        """
        return self._properties['value']

    @value.setter
    def value(self, value):
        """flow_rate.value setter

        The actual rate  

        value: int
        """
        self._properties['value'] = value

    def set(self, unit=None, value=None):
        self.unit = unit
        self.value = value
        return self
