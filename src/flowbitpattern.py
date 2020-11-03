from .snappicommon import SnappiObject


class FlowBitPattern(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_bitpattern.choice getter

        TBD  

        Returns: Union[bitlist, bitcounter]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_bitpattern.choice setter

        TBD  

        value: Union[bitlist, bitcounter]
        """
        self._properties['choice'] = value

    @property
    def bitlist(self):
        """flow_bitpattern.bitlist getter

        TBD  

        Returns: obj(snappi.FlowBitList)
        """
        from .flowbitlist import FlowBitList
        if 'bitlist' not in self._properties or self._properties['bitlist'] is None:
            self._properties['bitlist'] = FlowBitList()
        return self._properties['bitlist']

    @property
    def bitcounter(self):
        """flow_bitpattern.bitcounter getter

        TBD  

        Returns: obj(snappi.FlowBitCounter)
        """
        from .flowbitcounter import FlowBitCounter
        if 'bitcounter' not in self._properties or self._properties['bitcounter'] is None:
            self._properties['bitcounter'] = FlowBitCounter()
        return self._properties['bitcounter']

    def set(self, choice=None):
        self.choice = choice
        return self
