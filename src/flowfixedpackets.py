from .snappicommon import SnappiObject


class FlowFixedPackets(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def packets(self):
        """flow_fixedpackets.packets getter

        Stop transmit of the flow after this number of packets  

        Returns: int
        """
        return self._properties['packets']

    @packets.setter
    def packets(self, value):
        """flow_fixedpackets.packets setter

        Stop transmit of the flow after this number of packets  

        value: int
        """
        self._properties['packets'] = value

    @property
    def gap(self):
        """flow_fixedpackets.gap getter

        The minimum gap between packets expressed as bytes  

        Returns: int
        """
        return self._properties['gap']

    @gap.setter
    def gap(self, value):
        """flow_fixedpackets.gap setter

        The minimum gap between packets expressed as bytes  

        value: int
        """
        self._properties['gap'] = value

    @property
    def delay(self):
        """flow_fixedpackets.delay getter

        The delay before starting transmission of packets  

        Returns: int
        """
        return self._properties['delay']

    @delay.setter
    def delay(self, value):
        """flow_fixedpackets.delay setter

        The delay before starting transmission of packets  

        value: int
        """
        self._properties['delay'] = value

    @property
    def delay_unit(self):
        """flow_fixedpackets.delay_unit getter

        The delay expressed as a number of this value  

        Returns: Union[bytes, nanoseconds]
        """
        return self._properties['delay_unit']

    @delay_unit.setter
    def delay_unit(self, value):
        """flow_fixedpackets.delay_unit setter

        The delay expressed as a number of this value  

        value: Union[bytes, nanoseconds]
        """
        self._properties['delay_unit'] = value

    def set(self, packets=None, gap=None, delay=None, delay_unit=None):
        self.packets = packets
        self.gap = gap
        self.delay = delay
        self.delay_unit = delay_unit
        return self
