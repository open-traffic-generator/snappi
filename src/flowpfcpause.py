from .snappicommon import SnappiObject


class FlowPfcPause(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def dst(self):
        """flow_pfcpause.dst getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'dst' not in self._properties or self._properties['dst'] is None:
            self._properties['dst'] = FlowPattern()
        return self._properties['dst']

    @property
    def src(self):
        """flow_pfcpause.src getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'src' not in self._properties or self._properties['src'] is None:
            self._properties['src'] = FlowPattern()
        return self._properties['src']

    @property
    def ether_type(self):
        """flow_pfcpause.ether_type getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'ether_type' not in self._properties or self._properties['ether_type'] is None:
            self._properties['ether_type'] = FlowPattern()
        return self._properties['ether_type']

    @property
    def control_op_code(self):
        """flow_pfcpause.control_op_code getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'control_op_code' not in self._properties or self._properties['control_op_code'] is None:
            self._properties['control_op_code'] = FlowPattern()
        return self._properties['control_op_code']

    @property
    def class_enable_vector(self):
        """flow_pfcpause.class_enable_vector getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'class_enable_vector' not in self._properties or self._properties['class_enable_vector'] is None:
            self._properties['class_enable_vector'] = FlowPattern()
        return self._properties['class_enable_vector']

    @property
    def pause_class_0(self):
        """flow_pfcpause.pause_class_0 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_0' not in self._properties or self._properties['pause_class_0'] is None:
            self._properties['pause_class_0'] = FlowPattern()
        return self._properties['pause_class_0']

    @property
    def pause_class_1(self):
        """flow_pfcpause.pause_class_1 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_1' not in self._properties or self._properties['pause_class_1'] is None:
            self._properties['pause_class_1'] = FlowPattern()
        return self._properties['pause_class_1']

    @property
    def pause_class_2(self):
        """flow_pfcpause.pause_class_2 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_2' not in self._properties or self._properties['pause_class_2'] is None:
            self._properties['pause_class_2'] = FlowPattern()
        return self._properties['pause_class_2']

    @property
    def pause_class_3(self):
        """flow_pfcpause.pause_class_3 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_3' not in self._properties or self._properties['pause_class_3'] is None:
            self._properties['pause_class_3'] = FlowPattern()
        return self._properties['pause_class_3']

    @property
    def pause_class_4(self):
        """flow_pfcpause.pause_class_4 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_4' not in self._properties or self._properties['pause_class_4'] is None:
            self._properties['pause_class_4'] = FlowPattern()
        return self._properties['pause_class_4']

    @property
    def pause_class_5(self):
        """flow_pfcpause.pause_class_5 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_5' not in self._properties or self._properties['pause_class_5'] is None:
            self._properties['pause_class_5'] = FlowPattern()
        return self._properties['pause_class_5']

    @property
    def pause_class_6(self):
        """flow_pfcpause.pause_class_6 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_6' not in self._properties or self._properties['pause_class_6'] is None:
            self._properties['pause_class_6'] = FlowPattern()
        return self._properties['pause_class_6']

    @property
    def pause_class_7(self):
        """flow_pfcpause.pause_class_7 getter

        TBD  

        Returns: obj(snappi.FlowPattern)
        """
        from .flowpattern import FlowPattern
        if 'pause_class_7' not in self._properties or self._properties['pause_class_7'] is None:
            self._properties['pause_class_7'] = FlowPattern()
        return self._properties['pause_class_7']

    def set(self):
        return self
