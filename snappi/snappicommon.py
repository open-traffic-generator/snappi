import importlib
import json
import yaml
import requests
import io
import sys

if sys.version_info[0] == 3:
    unicode = str


class SnappiHttpTransport(object):
    def __init__(self, host=None):
        self.host = host if host else 'https://localhost'
        self._session = requests.Session()

    def send_recv(self, method, relative_url, payload=None, return_object=None):
        url = '%s%s' % (self.host, relative_url)
        data = None
        if payload is not None:
            if isinstance(payload, (str, unicode)):
                data = yaml.safe_load(payload)
            elif isinstance(payload, SnappiBase):
                data = payload.serialize()
            else:
                raise Exception('Type of payload provided is unknown')

        response = self._session.request(
            method=method,
            url=url,
            data=data,
            verify=False,
            allow_redirects=True,
            # TODO: add a timeout here
            headers={'Content-Type': 'application/json'}
        )
        if response.ok:
            if 'application/json' in response.headers['content-type']:
                # TODO: we might want to check for utf-8 charset and decode
                # accordingly, but current impl works for now
                response_dict = yaml.safe_load(response.text)
                if return_object is None:
                    # if response type is not provided, return dictionary
                    # instead of python object
                    return response_dict
                else:
                    return return_object.deserialize(response_dict)
            elif 'application/octet-stream' in response.headers['content-type']:
                return io.BytesIO(response.content)
            else:
                # TODO: for now, return bare response object for unknown
                # content types
                return response
        else:
            raise Exception(response.status_code, yaml.safe_load(response.text))


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

    Every SnappiObject is reuseable within the schema so it can 
    exist in multiple locations within the hierarchy.
    That means it can exist in multiple locations as a 
    leaf, parent/choice or parent.
    """
    __slots__ = ('_properties', '_parent', '_choice')

    def __init__(self, parent=None, choice=None):
        super(SnappiObject, self).__init__()
        self._parent = parent
        self._choice = choice
        self._properties = {}

    @property
    def parent(self):
        return self._parent
        
    def _get_property(self, name, default_value=None, parent=None, choice=None):
        if name not in self._properties or self._properties[name] is None:
            if isinstance(default_value, type) is True:
                self._properties[name] = default_value(parent=parent, choice=choice)
            else:
                self._properties[name] = default_value
        return self._properties[name]

    def _set_property(self, name, value, choice=None):
        self._properties[name] = value
        if choice is not None:
            self._properties['choice'] = choice
        elif self._parent is not None and self._choice is not None and value is not None:
            self._parent._set_property('choice', self._choice)
            
    def _encode(self):
        """Helper method for serialization
        """
        output = {}
        for key, value in self._properties.items():
            if isinstance(value, (SnappiObject, SnappiIter)):
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
                    if '_choice' in dir(child[1]) and '_parent' in dir(child[1]):
                        property_value = child[1](self, property_name)._decode(property_value)
                    else:
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
        return self.serialize(encoding=self.YAML)

    def __deepcopy__(self, memo):
        """Creates a deep copy of the current object
        """
        return self.__class__().deserialize(self.serialize())

    def __copy__(self):
        """Creates a deep copy of the current object
        """
        return self.__deepcopy__(None)

    def __eq__(self, other):
        return self.__str__() == other.__str__()

    def clone(self):
        """Creates a deep copy of the current object
        """
        return self.__deepcopy__(None)


class SnappiIter(SnappiBase):
    """Container class for SnappiObject

    Inheriting classes contain 0..n instances of an OpenAPI components/schemas 
    object.
    - config.flows.flow(name='1').flow(name='2').flow(name='3')

    The __getitem__ method allows getting an instance using ordinal.
    - config.flows[0]
    - config.flows[1:]
    - config.flows[0:1]
    - f1, f2, f3 = config.flows

    The __iter__ method allows for iterating across the encapsulated contents
    - for flow in config.flows:
    """
    __slots__ = ('_index', '_items')

    def __init__(self):
        super(SnappiIter, self).__init__()
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
        if 'choice' in found._properties and found._properties.get("choice") is not None:
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
        """Append an item to the end of the SnappiIter
        TBD: type check, raise error on mismatch
        """
        self._add(item)

    def clear(self):
        del self._items[:]
        self._index = -1

    def _encode(self):
        return [item._encode() for item in self._items]

    def _decode(self, encoded_snappi_list):
        item_class_name = self.__class__.__name__.replace('Iter', '')
        module = importlib.import_module(self.__module__)
        object_class = getattr(module, item_class_name)
        self.clear()
        for item in encoded_snappi_list:
            self._add(object_class()._decode(item))

    def __copy__(self):
        raise NotImplementedError('Shallow copy of SnappiIter objects is not supported')

    def __deepcopy__(self, memo):
        raise NotImplementedError('Deep copy of SnappiIter objects is not supported')

    def __str__(self):
        return yaml.safe_dump(self._encode())

    def __eq__(self, other):
        return self.__str__() == other.__str__()
