from .snappicommon import SnappiObject


class ControlState(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def choice(self):
        """control_state.choice getter

        TBD  

        Returns: Union[config_state, port_link_state, port_capture_state, flow_transmit_state, bgpv4_route_action, bgpv6_route_action]
        """
        return self._properties['choice']

    @choice.setter
    def choice(self, value):
        """control_state.choice setter

        TBD  

        value: Union[config_state, port_link_state, port_capture_state, flow_transmit_state, bgpv4_route_action, bgpv6_route_action]
        """
        self._properties['choice'] = value

    @property
    def config_state(self):
        """control_state.config_state getter

        TBD  

        Returns: obj(snappi.ControlConfigState)
        """
        from .controlconfigstate import ControlConfigState
        if 'config_state' not in self._properties or self._properties['config_state'] is None:
            self._properties['config_state'] = ControlConfigState()
        return self._properties['config_state']

    @property
    def port_link_state(self):
        """control_state.port_link_state getter

        TBD  

        Returns: obj(snappi.ControlPortLinkState)
        """
        from .controlportlinkstate import ControlPortLinkState
        if 'port_link_state' not in self._properties or self._properties['port_link_state'] is None:
            self._properties['port_link_state'] = ControlPortLinkState()
        return self._properties['port_link_state']

    @property
    def port_capture_state(self):
        """control_state.port_capture_state getter

        TBD  

        Returns: obj(snappi.ControlPortCaptureState)
        """
        from .controlportcapturestate import ControlPortCaptureState
        if 'port_capture_state' not in self._properties or self._properties['port_capture_state'] is None:
            self._properties['port_capture_state'] = ControlPortCaptureState()
        return self._properties['port_capture_state']

    @property
    def flow_transmit_state(self):
        """control_state.flow_transmit_state getter

        TBD  

        Returns: obj(snappi.ControlFlowTransmitState)
        """
        from .controlflowtransmitstate import ControlFlowTransmitState
        if 'flow_transmit_state' not in self._properties or self._properties['flow_transmit_state'] is None:
            self._properties['flow_transmit_state'] = ControlFlowTransmitState()
        return self._properties['flow_transmit_state']

    @property
    def bgpv4_route_action(self):
        """control_state.bgpv4_route_action getter

        TBD  

        Returns: obj(snappi.ControlBgpv4RouteAction)
        """
        from .controlbgpv4routeaction import ControlBgpv4RouteAction
        if 'bgpv4_route_action' not in self._properties or self._properties['bgpv4_route_action'] is None:
            self._properties['bgpv4_route_action'] = ControlBgpv4RouteAction()
        return self._properties['bgpv4_route_action']

    @property
    def bgpv6_route_action(self):
        """control_state.bgpv6_route_action getter

        TBD  

        Returns: obj(snappi.ControlBgpv6RouteAction)
        """
        from .controlbgpv6routeaction import ControlBgpv6RouteAction
        if 'bgpv6_route_action' not in self._properties or self._properties['bgpv6_route_action'] is None:
            self._properties['bgpv6_route_action'] = ControlBgpv6RouteAction()
        return self._properties['bgpv6_route_action']

    def set(self, choice=None):
        self.choice = choice
        return self
