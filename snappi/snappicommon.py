import importlib
import json
import yaml


class SnappiObject(object):
    JSON = 'json'
    YAML = 'yaml'
    DICT = 'dict'

    def __init__(self):
        self._properties = {}

    def encode(self, encoding=JSON):
        """Helper method for serialization
        """
        if encoding == SnappiObject.JSON:
            return json.dumps(self._encode(), indent=2)
        elif encoding == SnappiObject.YAML:
            return yaml.safe_dump(self._encode())
        elif encoding == SnappiObject.DICT:
            return self._encode()
        else:
            raise NotImplementedError('Encoding %s not supported' % encoding)

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

    def decode(self, obj):
        """Deserialize a json object depth first into the snappi object hierarchy
        Returns: obj(snappicommon.SnappiObject)
        """
        if isinstance(obj, str):
            obj = yaml.safe_load(obj)
        self._decode(obj)
        return self

    def _decode(self, obj):
        snappi_names = dir(self)
        for property_name, property_value in obj.items():
            if property_name in snappi_names:
                if isinstance(property_value, dict):
                    child = self._get_child_class(property_name)
                    property_value = child[1]()._decode(property_value)
                elif isinstance(property_value, list):
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


class SnappiList(object):
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
        self._index = -1
        self._items = []

    def __getattr__(self, name):
        if len(self._items) == 0:
            raise Exception('no items')
        elif hasattr(self._items[0], name):
            return getattr(self._items[0], name, None)
        else:
            raise AttributeError('%s is not a valid attribute')

    def __len__(self):
        return len(self._items)

    def __getitem__(self, key):
        if isinstance(key, int):
            return self._items[key]
        elif isinstance(key, str):
            for item in self._items:
                if item.name == key:
                    return item
            raise IndexError()

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

    def keys(self):
        return [item.name for item in self._items]

    def _encode(self):
        return [item._encode() for item in self._items]
