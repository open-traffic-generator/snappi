from .snappicommon import SnappiObject


class LagProtocol(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """lag_protocol.choice getter

        The type of lag protocol  

        Returns: Union[lacp, static]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """lag_protocol.choice setter

        The type of lag protocol  

        value: Union[lacp, static]
        """
        self._properties['choice'] = value

    @property
    def lacp(self):
        """lag_protocol.lacp getter

        The container for link aggregation control protocol settings  

        Returns: obj(snappi.LagLacp)
        """
        from .laglacp import LagLacp
        if 'lacp' not in self._properties or self._properties['lacp'] is None:
            self._properties['lacp'] = LagLacp()
        return self._properties['lacp']

    @property
    def static(self):
        """lag_protocol.static getter

        The container for link aggregation static protocol settings  

        Returns: obj(snappi.LagStatic)
        """
        from .lagstatic import LagStatic
        if 'static' not in self._properties or self._properties['static'] is None:
            self._properties['static'] = LagStatic()
        return self._properties['static']

    def set(self, choice=None):
        self.choice = choice
        return self
