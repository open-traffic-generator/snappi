from json import JSONEncoder, JSONDecoder

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

    def decode(self, s):
        pass

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

# class DeviceList(SnappiList):
#     def __init__(self):
#         super().__init__()
    
#     def append(self, name=None, device_count=1):
#         from snappi.device import Device
#         self._items.append(Device().set(name, device_count))
#         return self


if __name__ == '__main__':
    import sys
    sys.path.append('./snappi')
    import snappi

    dl = DeviceList()
    dl.append(name='a').append(name='b')
    print(dl.keys())
