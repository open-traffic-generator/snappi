from .snappicommon import SnappiObject


class FlowPattern(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """flow_pattern.choice getter

        TBD  

        Returns: Union[fixed, list, counter, random]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """flow_pattern.choice setter

        TBD  

        value: Union[fixed, list, counter, random]
        """
        self._properties['choice'] = value

    @property
    def fixed(self):
        """flow_pattern.fixed getter

        TBD  

        Returns: str
        """
        return self._properties['fixed']

    @fixed.setter
    def fixed(self, value):
        """flow_pattern.fixed setter

        TBD  

        value: str
        """
        self._properties['fixed'] = value

    @property
    def list(self):
        """flow_pattern.list getter

        TBD  

        Returns: list[str]
        """
        return self._properties['list']

    @list.setter
    def list(self, value):
        """flow_pattern.list setter

        TBD  

        value: list[str]
        """
        self._properties['list'] = value

    @property
    def counter(self):
        """flow_pattern.counter getter

        TBD  

        Returns: obj(snappi.FlowCounter)
        """
        from .flowcounter import FlowCounter
        if 'counter' not in self._properties or self._properties['counter'] is None:
            self._properties['counter'] = FlowCounter()
        return self._properties['counter']

    @property
    def random(self):
        """flow_pattern.random getter

        TBD  

        Returns: obj(snappi.FlowRandom)
        """
        from .flowrandom import FlowRandom
        if 'random' not in self._properties or self._properties['random'] is None:
            self._properties['random'] = FlowRandom()
        return self._properties['random']

    @property
    def ingress_result_name(self):
        """flow_pattern.ingress_result_name getter

        A unique name is used to indicate to the system that the field may extend the result row key and create an aggregate result row for every unique ingress value  
        To have ingress columns appear in flow result rows the flow result request allows for the ingress_result_name value to be specified as part of the request  

        Returns: str
        """
        return self._properties['ingress_result_name']

    @ingress_result_name.setter
    def ingress_result_name(self, value):
        """flow_pattern.ingress_result_name setter

        A unique name is used to indicate to the system that the field may extend the result row key and create an aggregate result row for every unique ingress value  
        To have ingress columns appear in flow result rows the flow result request allows for the ingress_result_name value to be specified as part of the request  

        value: str
        """
        self._properties['ingress_result_name'] = value

    def set(self, choice=None, fixed=None, list=None, ingress_result_name=None):
        self.choice = choice
        self.fixed = fixed
        self.list = list
        self.ingress_result_name = ingress_result_name
        return self
