package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceMacsecDevice *****
type deviceMacsecDevice struct {
	validation
	obj                      *otg.DeviceMacsecDevice
	marshaller               marshalDeviceMacsecDevice
	unMarshaller             unMarshalDeviceMacsecDevice
	ethernetInterfacesHolder DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
}

func NewDeviceMacsecDevice() DeviceMacsecDevice {
	obj := deviceMacsecDevice{obj: &otg.DeviceMacsecDevice{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceMacsecDevice) msg() *otg.DeviceMacsecDevice {
	return obj.obj
}

func (obj *deviceMacsecDevice) setMsg(msg *otg.DeviceMacsecDevice) DeviceMacsecDevice {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceMacsecDevice struct {
	obj *deviceMacsecDevice
}

type marshalDeviceMacsecDevice interface {
	// ToProto marshals DeviceMacsecDevice to protobuf object *otg.DeviceMacsecDevice
	ToProto() (*otg.DeviceMacsecDevice, error)
	// ToPbText marshals DeviceMacsecDevice to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceMacsecDevice to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceMacsecDevice to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceMacsecDevice struct {
	obj *deviceMacsecDevice
}

type unMarshalDeviceMacsecDevice interface {
	// FromProto unmarshals DeviceMacsecDevice from protobuf object *otg.DeviceMacsecDevice
	FromProto(msg *otg.DeviceMacsecDevice) (DeviceMacsecDevice, error)
	// FromPbText unmarshals DeviceMacsecDevice from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceMacsecDevice from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceMacsecDevice from JSON text
	FromJson(value string) error
}

func (obj *deviceMacsecDevice) Marshal() marshalDeviceMacsecDevice {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceMacsecDevice{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceMacsecDevice) Unmarshal() unMarshalDeviceMacsecDevice {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceMacsecDevice{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceMacsecDevice) ToProto() (*otg.DeviceMacsecDevice, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceMacsecDevice) FromProto(msg *otg.DeviceMacsecDevice) (DeviceMacsecDevice, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceMacsecDevice) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshaldeviceMacsecDevice) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshaldeviceMacsecDevice) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldeviceMacsecDevice) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshaldeviceMacsecDevice) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshaldeviceMacsecDevice) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *deviceMacsecDevice) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceMacsecDevice) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceMacsecDevice) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceMacsecDevice) Clone() (DeviceMacsecDevice, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceMacsecDevice()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *deviceMacsecDevice) setNil() {
	obj.ethernetInterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceMacsecDevice is a container of properties for a MACsec capable device.
type DeviceMacsecDevice interface {
	Validation
	// msg marshals DeviceMacsecDevice to protobuf object *otg.DeviceMacsecDevice
	// and doesn't set defaults
	msg() *otg.DeviceMacsecDevice
	// setMsg unmarshals DeviceMacsecDevice from protobuf object *otg.DeviceMacsecDevice
	// and doesn't set defaults
	setMsg(*otg.DeviceMacsecDevice) DeviceMacsecDevice
	// provides marshal interface
	Marshal() marshalDeviceMacsecDevice
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceMacsecDevice
	// validate validates DeviceMacsecDevice
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceMacsecDevice, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetInterfaces returns DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIterIter, set in DeviceMacsecDevice
	EthernetInterfaces() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	setNil()
}

// Ethernet Interfaces
// EthernetInterfaces returns a []MacsecDeviceEthernetInterface
func (obj *deviceMacsecDevice) EthernetInterfaces() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	if len(obj.obj.EthernetInterfaces) == 0 {
		obj.obj.EthernetInterfaces = []*otg.MacsecDeviceEthernetInterface{}
	}
	if obj.ethernetInterfacesHolder == nil {
		obj.ethernetInterfacesHolder = newDeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter(&obj.obj.EthernetInterfaces).setMsg(obj)
	}
	return obj.ethernetInterfacesHolder
}

type deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter struct {
	obj                                *deviceMacsecDevice
	macsecDeviceEthernetInterfaceSlice []MacsecDeviceEthernetInterface
	fieldPtr                           *[]*otg.MacsecDeviceEthernetInterface
}

func newDeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter(ptr *[]*otg.MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	return &deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter{fieldPtr: ptr}
}

type DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter interface {
	setMsg(*deviceMacsecDevice) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	Items() []MacsecDeviceEthernetInterface
	Add() MacsecDeviceEthernetInterface
	Append(items ...MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	Set(index int, newObj MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	Clear() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	clearHolderSlice() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
	appendHolderSlice(item MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter
}

func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) setMsg(msg *deviceMacsecDevice) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecDeviceEthernetInterface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) Items() []MacsecDeviceEthernetInterface {
	return obj.macsecDeviceEthernetInterfaceSlice
}

func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) Add() MacsecDeviceEthernetInterface {
	newObj := &otg.MacsecDeviceEthernetInterface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecDeviceEthernetInterface{obj: newObj}
	newLibObj.setDefault()
	obj.macsecDeviceEthernetInterfaceSlice = append(obj.macsecDeviceEthernetInterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) Append(items ...MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecDeviceEthernetInterfaceSlice = append(obj.macsecDeviceEthernetInterfaceSlice, item)
	}
	return obj
}

func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) Set(index int, newObj MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecDeviceEthernetInterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) Clear() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecDeviceEthernetInterface{}
		obj.macsecDeviceEthernetInterfaceSlice = []MacsecDeviceEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) clearHolderSlice() DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	if len(obj.macsecDeviceEthernetInterfaceSlice) > 0 {
		obj.macsecDeviceEthernetInterfaceSlice = []MacsecDeviceEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecDeviceMacsecDeviceEthernetInterfaceIter) appendHolderSlice(item MacsecDeviceEthernetInterface) DeviceMacsecDeviceMacsecDeviceEthernetInterfaceIter {
	obj.macsecDeviceEthernetInterfaceSlice = append(obj.macsecDeviceEthernetInterfaceSlice, item)
	return obj
}

func (obj *deviceMacsecDevice) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.EthernetInterfaces) != 0 {

		if set_default {
			obj.EthernetInterfaces().clearHolderSlice()
			for _, item := range obj.obj.EthernetInterfaces {
				obj.EthernetInterfaces().appendHolderSlice(&macsecDeviceEthernetInterface{obj: item})
			}
		}
		for _, item := range obj.EthernetInterfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceMacsecDevice) setDefault() {

}
