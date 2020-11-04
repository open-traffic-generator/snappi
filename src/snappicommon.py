import importlib
from json import JSONEncoder

def _default_encoder(self, obj):
    return getattr(obj.__class__, "encode", _default_encoder.default)(obj)
_default_encoder.default = JSONEncoder().default
JSONEncoder.default = _default_encoder


class SnappiObject(object):
    def __init__(self):
        self._properties = {}

    def encode(self):
        output = {}
        for key, value in self._properties.items():
            if isinstance(value, (SnappiObject, SnappiList)):
                output[key] = value.encode()
            else:
                output[key] = value
        return output

def decode(snappi_class, obj):
    """Deserialize a json object depth first into the snappi object hierarchy
    Returns: obj(snappicommon.SnappiObject)
    """
    snappi_object = snappi_class()
    snappi_names = dir(snappi_object)
    for property_name, property_value in obj.items():
        if property_name in snappi_names:
            if isinstance(property_value, dict):
                child_list_class, child_object_class = get_child_class(snappi_object, property_name)
                property_value = decode(child_object_class, property_value) 
            elif isinstance(property_value, list):
                child_list_class, child_object_class = get_child_class(snappi_object, property_name, True)
                snappi_list = child_list_class()
                for item in property_value:
                    item = decode(child_object_class, item)
                    snappi_list._items.append(item)
                property_value = snappi_list
            snappi_object._properties[property_name] = property_value
    return snappi_object

def get_child_class(snappi_object, property_name, is_property_list=False):
    list_class = None
    package_name = __name__.split('.')[0]
    child_class_path_name = snappi_object._TYPES[property_name].split('.')
    child_class_path = '.'.join(child_class_path_name[0:-1])
    child_class_name = child_class_path_name[-1]
    module = importlib.import_module(child_class_path, package=package_name)
    object_class = getattr(module, child_class_name)
    if is_property_list is True:
        list_class = object_class
        module = importlib.import_module(child_class_path[0:-4], package=package_name)
        object_class = getattr(module, child_class_name[0:-4])
    return (list_class, object_class)
            

class SnappiList(object):
    """Container class for SnappiObject
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

    def encode(self):
        return [item.encode() for item in self._items]

