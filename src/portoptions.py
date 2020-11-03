from .snappicommon import SnappiObject


class PortOptions(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def location_preemption(self):
        """port_options.location_preemption getter

        Preempt all the test port locations as defined by the Port.Port.properties.location  
        If the test ports as defined by their location values are in use and this value is true, the test ports will be preempted  

        Returns: boolean
        """
        return self._properties['location_preemption']

    @location_preemption.setter
    def location_preemption(self, value):
        """port_options.location_preemption setter

        Preempt all the test port locations as defined by the Port.Port.properties.location  
        If the test ports as defined by their location values are in use and this value is true, the test ports will be preempted  

        value: boolean
        """
        self._properties['location_preemption'] = value

    def set(self, location_preemption=None):
        self.location_preemption = location_preemption
        return self
