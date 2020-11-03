from .snappicommon import SnappiObject


class ControlBgpv4RouteAction(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def names(self):
        """control_bgpv4routeaction.names getter

        The names of Bgpv4RouteRange object to control  

        Returns: list[str]
        """
        return self._properties['names']

    @names.setter
    def names(self, value):
        """control_bgpv4routeaction.names setter

        The names of Bgpv4RouteRange object to control  

        value: list[str]
        """
        self._properties['names'] = value

    @property
    def action(self):
        """control_bgpv4routeaction.action getter

        BGPv4 route specific action  

        Returns: Union[withdraw_routes, advertise_routes]
        """
        return self._properties['action']

    @action.setter
    def action(self, value):
        """control_bgpv4routeaction.action setter

        BGPv4 route specific action  

        value: Union[withdraw_routes, advertise_routes]
        """
        self._properties['action'] = value

    def set(self, names=None, action=None):
        self.names = names
        self.action = action
        return self
