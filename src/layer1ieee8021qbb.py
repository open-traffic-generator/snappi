from .snappicommon import SnappiObject


class Layer1Ieee8021qbb(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def pfc_delay(self):
        """layer1_ieee8021qbb.pfc_delay getter

        The upper limit on the transmit time of a queue after receiving a message to pause a specified priority  
        A value of 0 or null indicates that pfc delay will not be enabled  

        Returns: int
        """
        return self._properties['pfc_delay']

    @pfc_delay.setter
    def pfc_delay(self, value):
        """layer1_ieee8021qbb.pfc_delay setter

        The upper limit on the transmit time of a queue after receiving a message to pause a specified priority  
        A value of 0 or null indicates that pfc delay will not be enabled  

        value: int
        """
        self._properties['pfc_delay'] = value

    @property
    def pfc_class_0(self):
        """layer1_ieee8021qbb.pfc_class_0 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_0']

    @pfc_class_0.setter
    def pfc_class_0(self, value):
        """layer1_ieee8021qbb.pfc_class_0 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_0'] = value

    @property
    def pfc_class_1(self):
        """layer1_ieee8021qbb.pfc_class_1 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_1']

    @pfc_class_1.setter
    def pfc_class_1(self, value):
        """layer1_ieee8021qbb.pfc_class_1 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_1'] = value

    @property
    def pfc_class_2(self):
        """layer1_ieee8021qbb.pfc_class_2 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_2']

    @pfc_class_2.setter
    def pfc_class_2(self, value):
        """layer1_ieee8021qbb.pfc_class_2 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_2'] = value

    @property
    def pfc_class_3(self):
        """layer1_ieee8021qbb.pfc_class_3 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_3']

    @pfc_class_3.setter
    def pfc_class_3(self, value):
        """layer1_ieee8021qbb.pfc_class_3 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_3'] = value

    @property
    def pfc_class_4(self):
        """layer1_ieee8021qbb.pfc_class_4 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_4']

    @pfc_class_4.setter
    def pfc_class_4(self, value):
        """layer1_ieee8021qbb.pfc_class_4 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_4'] = value

    @property
    def pfc_class_5(self):
        """layer1_ieee8021qbb.pfc_class_5 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_5']

    @pfc_class_5.setter
    def pfc_class_5(self, value):
        """layer1_ieee8021qbb.pfc_class_5 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_5'] = value

    @property
    def pfc_class_6(self):
        """layer1_ieee8021qbb.pfc_class_6 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_6']

    @pfc_class_6.setter
    def pfc_class_6(self, value):
        """layer1_ieee8021qbb.pfc_class_6 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_6'] = value

    @property
    def pfc_class_7(self):
        """layer1_ieee8021qbb.pfc_class_7 getter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        Returns: int
        """
        return self._properties['pfc_class_7']

    @pfc_class_7.setter
    def pfc_class_7(self, value):
        """layer1_ieee8021qbb.pfc_class_7 setter

        The valid values are null, 0 - 7  
        A null value indicates there is no setting for this pfc class  

        value: int
        """
        self._properties['pfc_class_7'] = value

    def set(self, pfc_delay=None, pfc_class_0=None, pfc_class_1=None, pfc_class_2=None, pfc_class_3=None, pfc_class_4=None, pfc_class_5=None, pfc_class_6=None, pfc_class_7=None):
        self.pfc_delay = pfc_delay
        self.pfc_class_0 = pfc_class_0
        self.pfc_class_1 = pfc_class_1
        self.pfc_class_2 = pfc_class_2
        self.pfc_class_3 = pfc_class_3
        self.pfc_class_4 = pfc_class_4
        self.pfc_class_5 = pfc_class_5
        self.pfc_class_6 = pfc_class_6
        self.pfc_class_7 = pfc_class_7
        return self
