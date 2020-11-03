from .snappicommon import SnappiObject


class Flow(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """flow.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """flow.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def tx_rx(self):
        """flow.tx_rx getter

        TBD  

        Returns: obj(snappi.FlowTxRx)
        """
        from .flowtxrx import FlowTxRx
        if 'tx_rx' not in self._properties or self._properties['tx_rx'] is None:
            self._properties['tx_rx'] = FlowTxRx()
        return self._properties['tx_rx']

    @property
    def packet(self):
        """flow.packet getter

        The packet is a list of traffic protocol headers  
        The order of traffic protocol headers assigned to the list is the order they will appear on the wire  

        Returns: list[obj(snappi.FlowHeader)]
        """
        from .flowheaderlist import FlowHeaderList
        if 'packet' not in self._properties or self._properties['packet'] is None:
            self._properties['packet'] = FlowHeaderList()
        return self._properties['packet']

    @property
    def size(self):
        """flow.size getter

        TBD  

        Returns: obj(snappi.FlowSize)
        """
        from .flowsize import FlowSize
        if 'size' not in self._properties or self._properties['size'] is None:
            self._properties['size'] = FlowSize()
        return self._properties['size']

    @property
    def rate(self):
        """flow.rate getter

        TBD  

        Returns: obj(snappi.FlowRate)
        """
        from .flowrate import FlowRate
        if 'rate' not in self._properties or self._properties['rate'] is None:
            self._properties['rate'] = FlowRate()
        return self._properties['rate']

    @property
    def duration(self):
        """flow.duration getter

        TBD  

        Returns: obj(snappi.FlowDuration)
        """
        from .flowduration import FlowDuration
        if 'duration' not in self._properties or self._properties['duration'] is None:
            self._properties['duration'] = FlowDuration()
        return self._properties['duration']

    def set(self, name=None):
        self.name = name
        return self
