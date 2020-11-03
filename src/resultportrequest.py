from .snappicommon import SnappiObject


class ResultPortRequest(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def port_names(self):
        """result_portrequest.port_names getter

        The names of objects to return results for  
        An empty list will return all port row results  

        Returns: list[str]
        """
        return self._properties['port_names']

    @port_names.setter
    def port_names(self, value):
        """result_portrequest.port_names setter

        The names of objects to return results for  
        An empty list will return all port row results  

        value: list[str]
        """
        self._properties['port_names'] = value

    @property
    def column_names(self):
        """result_portrequest.column_names getter

        The names of Result.Port properties to return  
        If the list is empty then all properties will be returned  

        Returns: list[str]
        """
        return self._properties['column_names']

    @column_names.setter
    def column_names(self, value):
        """result_portrequest.column_names setter

        The names of Result.Port properties to return  
        If the list is empty then all properties will be returned  

        value: list[str]
        """
        self._properties['column_names'] = value

    def set(self, port_names=None, column_names=None):
        self.port_names = port_names
        self.column_names = column_names
        return self
