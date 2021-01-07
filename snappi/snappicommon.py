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


class SnappiSerialization(object):
    JSON = 'json'
    YAML = 'yaml'
    DICT = 'dict'

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
        if encoding == SnappiSerialization.JSON:
            return json.dumps(self._encode(), indent=2)
        elif encoding == SnappiSerialization.YAML:
            return yaml.safe_dump(self._encode())
        elif encoding == SnappiSerialization.DICT:
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


class SnappiObject(SnappiSerialization):
    """Base class for any /components/schemas object
    """
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
        package_name = __name__.split('.')[0]
        child_class_path_name = self._TYPES[property_name].split('.')
        child_class_path = '.'.join(child_class_path_name[0:-1])
        child_class_name = child_class_path_name[-1]
        module = importlib.import_module(child_class_path,
                                         package=package_name)
        object_class = getattr(module, child_class_name)
        if is_property_list is True:
            list_class = object_class
            module = importlib.import_module(child_class_path[0:-4],
                                             package=package_name)
            object_class = getattr(module, child_class_name[0:-4])
        return (list_class, object_class)

    def __str__(self):
        return self.serialize(self.YAML)


class SnappiList(SnappiSerialization):
    """Container class for SnappiObject

    Inheriting classes contain 0..n instances of an OpenAPI components/schemas 
    object.
    - flow = config.flows.append(name='test')

    The __getitem__ method allows getting an instance using ordinal or name.
    - config.flows[0]
    - config.flows['test']

    The __iter__ method allows for iterating across the encapsulated contents
    - for flow in config.flows:
    """
    def __init__(self):
        super(SnappiList, self).__init__()
        self._index = -1
        self._items = []

    def __getattr__(self, name):
        if len(self._items) == 0:
            raise IndexError('List is empty')
        elif hasattr(self._items[self._index], name):
            return getattr(self._items[self._index], name, None)
        else:
            raise AttributeError('%s.%s is not a valid attribute' %
                                 (type(self).__name__, name))

    def __len__(self):
        return len(self._items)

    def __getitem__(self, key):
        found = None
        if isinstance(key, int):
            found = self._items[key]
        elif isinstance(key, str):
            for item in self._items:
                if item.name == key:
                    found = item
        if found is None:
            raise IndexError()
        if 'choice' in found._properties:
            return found._properties[found._properties['choice']]
        return found

    def __iter__(self):
        self._index = -1
        return self

    def next(self):
        return self.__next__()

    def __next__(self):
        if self._index + 1 >= len(self._items):
            raise StopIteration
        else:
            self._index += 1
        return self[self._index]

    def _keys(self):
        return [item.name for item in self._items]

    def _add(self, item):
        self._items.append(item)
        self._index = len(self._items) - 1

    def clear(self):
        self._items.clear()
        self._index = -1

    def _encode(self):
        return [item._encode() for item in self._items]

    def _decode(self, encoded_snappi_list):
        item_class_path = self.__class__.__module__.replace('list', '')
        item_class_name = self.__class__.__name__.replace('List', '')
        module = importlib.import_module(item_class_path)
        object_class = getattr(module, item_class_name)
        self._items.clear()
        for item in encoded_snappi_list:
            self._add(object_class()._decode(item))

    def __str__(self):
        return yaml.safe_dump(self._encode())
