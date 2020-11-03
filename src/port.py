from .snappicommon import SnappiObject


class Port(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def name(self):
        """port.name getter

        Unique name of an object that is the primary key for objects found in arrays  

        Returns: str
        """
        return self._properties['name']

    @name.setter
    def name(self, value):
        """port.name setter

        Unique name of an object that is the primary key for objects found in arrays  

        value: str
        """
        self._properties['name'] = value

    @property
    def location(self):
        """port.location getter

        The location of a test port  
        It is the endpoint where packets will emit from  
        Test port locations can be the following:  
        - physical appliance with multiple ports  
        - physical chassis with multiple cards and ports  
        - local interface  
        - virtual machine, docker container, kubernetes cluster  
        The test port location format is implementation specific  
        Use the /results/capabilities API to determine what formats an implementation supports for the location property  
        Get the configured location state by using the /results/port API  

        Returns: str
        """
        return self._properties['location']

    @location.setter
    def location(self, value):
        """port.location setter

        The location of a test port  
        It is the endpoint where packets will emit from  
        Test port locations can be the following:  
        - physical appliance with multiple ports  
        - physical chassis with multiple cards and ports  
        - local interface  
        - virtual machine, docker container, kubernetes cluster  
        The test port location format is implementation specific  
        Use the /results/capabilities API to determine what formats an implementation supports for the location property  
        Get the configured location state by using the /results/port API  

        value: str
        """
        self._properties['location'] = value

    def set(self, name=None, location=None):
        self.name = name
        self.location = location
        return self
