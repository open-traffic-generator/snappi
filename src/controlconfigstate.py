from .snappicommon import SnappiObject


class ControlConfigState(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def config(self):
        """control_configstate.config getter

        TBD  

        Returns: obj(snappi.Config)
        """
        from .config import Config
        if 'config' not in self._properties or self._properties['config'] is None:
            self._properties['config'] = Config()
        return self._properties['config']

    @property
    def state(self):
        """control_configstate.state getter

        Set the configuration state on the traffic generator  
        - set is used to submit a complete running configuration on the traffic generator  
        - update is used to submit a partial configuration which will be merged with the current running configuration on the traffic generator  

        Returns: Union[set, update]
        """
        return self._properties['state']

    @state.setter
    def state(self, value):
        """control_configstate.state setter

        Set the configuration state on the traffic generator  
        - set is used to submit a complete running configuration on the traffic generator  
        - update is used to submit a partial configuration which will be merged with the current running configuration on the traffic generator  

        value: Union[set, update]
        """
        self._properties['state'] = value

    def set(self, state=None):
        self.state = state
        return self
