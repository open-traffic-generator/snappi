import importlib
import json
import yaml
import requests
import sys

if sys.version_info[0] == 3:
    unicode = str


class SnappiRestTransport(object):
    def __init__(self):
        self.base_url = 'http://127.0.0.1:80'
        self._session = requests.Session()

    def send_recv(self, method, relative_url, payload=None, return_object=None):
        url = '%s%s' % (self.base_url, relative_url)
        data = None
        if payload is not None:
            data = payload.serialize()
        response = self._session.request(method=method,
                                         url=url,
                                         data=data,
                                         verify=False,
                                         allow_redirects=True)
        if response.ok is not True:
            raise Exception(response.status_code, yaml.safe_load(response.text))
        elif response.headers['content-type'] == 'application/json':
            return return_object.deserialize(yaml.safe_load(response.text))
        else:
            return None


class SnappiBase(object):
    """Base class for all Snappi classes
    """
    JSON = 'json'
    YAML = 'yaml'
    DICT = 'dict'

    __slots__ = ()

    def __init__(self):
        pass

    def serialize(self, encoding=JSON):
        """Serialize the current snappi object according to a specified encoding.

        Args
        ----
        - encoding (str[json, yaml, dict]): The object will be recursively
            serialized according to the specified encoding.
            The supported encodings are json, yaml and python dict. 

        Returns
        -------
        - obj(Union[str, dict]): A str or dict object depending on the specified
            encoding. The json and yaml encodings will return a str object and
            the dict encoding will return a python dict object.
        """
        if encoding == SnappiBase.JSON:
            return json.dumps(self._encode(), indent=2)
        elif encoding == SnappiBase.YAML:
            return yaml.safe_dump(self._encode())
        elif encoding == SnappiBase.DICT:
            return self._encode()
        else:
            raise NotImplementedError('Encoding %s not supported' % encoding)

    def _encode(self):
        raise NotImplementedError()

    def deserialize(self, serialized_object):
        """Deserialize a python object into the current snappi object.

        If the input `serialized_object` does not match the current 
        snappi object an exception will be raised. 
        
        Args
        ----
        - serialized_object (Union[str, dict]): The object to deserialize.
            If the serialized_object is of type str then the internal encoding 
            of the serialized_object must be json or yaml. 

        Returns
        -------
        - obj(snappicommon.SnappiObject): This snappi object with all the 
            serialized_object deserialized within.
        """
        if isinstance(serialized_object, (str, unicode)):
            serialized_object = yaml.safe_load(serialized_object)
        self._decode(serialized_object)
        return self
        
    def _decode(self, dict_object):
        raise NotImplementedError()


class SnappiObject(SnappiBase):
    """Base class for any /components/schemas object
    """
    __slots__ = ('_properties')

    def __init__(self):
        super(SnappiObject, self).__init__()
        self._properties = {}

    def _encode(self):
        """Helper method for serialization
        """
        output = {}
        for key, value in self._properties.items():
            if isinstance(value, (SnappiObject, SnappiList)):
                output[key] = value._encode()
            else:
                output[key] = value
        return output

    def _decode(self, obj):
        snappi_names = dir(self)
        for property_name, property_value in obj.items():
            if property_name in snappi_names:
                if isinstance(property_value, dict):
                    child = self._get_child_class(property_name)
                    property_value = child[1]()._decode(property_value)
                elif isinstance(property_value,
                                list) and property_name in self._TYPES:
                    child = self._get_child_class(property_name, True)
                    snappi_list = child[0]()
                    for item in property_value:
                        item = child[1]()._decode(item)
                        snappi_list._items.append(item)
                    property_value = snappi_list
                self._properties[property_name] = property_value
        return self

    def _get_child_class(self, property_name, is_property_list=False):
        list_class = None
        class_name = self._TYPES[property_name]
        module = importlib.import_module(self.__module__)
        object_class = getattr(module, class_name)
        if is_property_list is True:
            list_class = object_class
            object_class = getattr(module, class_name[0:-4])
        return (list_class, object_class)

    def __str__(self):
        return self.serialize(self.YAML)

    def __deepcopy__(self, memo):
        """Creates a deep copy of the current object
        """
        return self.__class__().deserialize(self.serialize())

    def __copy__(self):
        """Creates a deep copy of the current object
        """
        return self.__deepcopy__(None)

    def clone(self):
        """Creates a deep copy of the current object
        """
        return self.__deepcopy__(None)


class SnappiList(SnappiBase):
    """Container class for SnappiObject

    Inheriting classes contain 0..n instances of an OpenAPI components/schemas 
    object.
    - flow = config.flows.flow(name='test')

    The __getitem__ method allows getting an instance using ordinal or name.
    - config.flows[0]
    - config.flows['test']

    The __iter__ method allows for iterating across the encapsulated contents
    - for flow in config.flows:
    """
    __slots__ = ('_index', '_items')

    def __init__(self):
        super(SnappiList, self).__init__()
        self._index = -1
        self._items = []

    def __len__(self):
        return len(self._items)

    def _getitem(self, key):
        found = None
        if isinstance(key, int):
            found = self._items[key]
        elif isinstance(key, slice) is True:
            start, stop, step = key.indices(len(self))
            sliced = self.__class__()
            for i in range(start, stop, step):
                sliced._items.append(self._items[i])
            return sliced
        elif isinstance(key, str):
            for item in self._items:
                if item.name == key:
                    found = item
        if found is None:
            raise IndexError()
        if 'choice' in found._properties:
            return found._properties[found._properties['choice']]
        return found

    def _iter(self):
        self._index = -1
        return self

    def _next(self):
        if self._index + 1 >= len(self._items):
            raise StopIteration
        else:
            self._index += 1
        return self._getitem(self._index)

    def _add(self, item):
        self._items.append(item)
        self._index = len(self._items) - 1

    def append(self, item):
        """Append an item to the end of the SnappiList
        TBD: type check, raise error on mismatch
        """
        self._add(item)

    def clear(self):
        self._items.clear()
        self._index = -1

    def _encode(self):
        return [item._encode() for item in self._items]

    def _decode(self, encoded_snappi_list):
        item_class_name = self.__class__.__name__.replace('List', '')
        module = importlib.import_module(self.__module__)
        object_class = getattr(module, item_class_name)
        self.clear()
        for item in encoded_snappi_list:
            self._add(object_class()._decode(item))

    def __copy__(self):
        raise NotImplementedError('Shallow copy of SnappiList objects is not supported')

    def __deepcopy__(self, memo):
        raise NotImplementedError('Deep copy of SnappiList objects is not supported')

    def __str__(self):
        return yaml.safe_dump(self._encode())
