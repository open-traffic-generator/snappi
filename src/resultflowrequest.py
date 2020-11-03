from .snappicommon import SnappiObject


class ResultFlowRequest(SnappiObject):
    def __init__(self):
        super().__init__()
        self.set()

    @property
    def flow_names(self):
        """result_flowrequest.flow_names getter

        The names of flow objects to return results for  
        An empty list will return results for all flows  

        Returns: list[str]
        """
        return self._properties['flow_names']

    @flow_names.setter
    def flow_names(self, value):
        """result_flowrequest.flow_names setter

        The names of flow objects to return results for  
        An empty list will return results for all flows  

        value: list[str]
        """
        self._properties['flow_names'] = value

    @property
    def column_names(self):
        """result_flowrequest.column_names getter

        The names of Result.Flow properties to return  
        If the list is empty then all properties will be returned  

        Returns: list[str]
        """
        return self._properties['column_names']

    @column_names.setter
    def column_names(self, value):
        """result_flowrequest.column_names setter

        The names of Result.Flow properties to return  
        If the list is empty then all properties will be returned  

        value: list[str]
        """
        self._properties['column_names'] = value

    @property
    def ingress_result_names(self):
        """result_flowrequest.ingress_result_names getter

        Add any configured Flow.Pattern.ingress_result_name values that are to be included in the results  
        If the name is not configured then it will be excluded from the Result.Flow.columns and Result.Flow.rows  
        The name in the Result.Flow.columns will be a combination of the ingress_result_name and any system assigned name  

        Returns: list[str]
        """
        return self._properties['ingress_result_names']

    @ingress_result_names.setter
    def ingress_result_names(self, value):
        """result_flowrequest.ingress_result_names setter

        Add any configured Flow.Pattern.ingress_result_name values that are to be included in the results  
        If the name is not configured then it will be excluded from the Result.Flow.columns and Result.Flow.rows  
        The name in the Result.Flow.columns will be a combination of the ingress_result_name and any system assigned name  

        value: list[str]
        """
        self._properties['ingress_result_names'] = value

    def set(self, flow_names=None, column_names=None, ingress_result_names=None):
        self.flow_names = flow_names
        self.column_names = column_names
        self.ingress_result_names = ingress_result_names
        return self
