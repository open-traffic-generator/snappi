from .snappicommon import SnappiObject


class ResultBgpv4Request(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def names(self):
        """result_bgpv4request.names getter

        The names of BGP objects to return results for  
        An empty list will return results for all BGP  

        Returns: list[str]
        """
        return self._properties['names']

    @names.setter
    def names(self, value):
        """result_bgpv4request.names setter

        The names of BGP objects to return results for  
        An empty list will return results for all BGP  

        value: list[str]
        """
        self._properties['names'] = value

    def set(self, names=None):
        self.names = names
        return self
