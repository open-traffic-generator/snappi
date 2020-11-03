from .snappicommon import SnappiList


class Layer1List(SnappiList):
    def __init__(self):
        super().__init__()

    def append(self, name=None, port_names=None, speed=speed_10_gbps, media=None, promiscuous=False, mtu=1500, ieee_media_defaults=True, auto_negotiate=True):
        from .layer1 import Layer1
        self._items.append(Layer1().set(name, port_names, speed, media, promiscuous, mtu, ieee_media_defaults, auto_negotiate))
        return self
