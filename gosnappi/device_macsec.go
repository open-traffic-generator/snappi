package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceMacsec *****
type deviceMacsec struct {
	validation
	obj                      *otg.DeviceMacsec
	marshaller               marshalDeviceMacsec
	unMarshaller             unMarshalDeviceMacsec
	ethernetInterfacesHolder DeviceMacsecDeviceMacsecEthernetInterfaceIter
}

func NewDeviceMacsec() DeviceMacsec {
	obj := deviceMacsec{obj: &otg.DeviceMacsec{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceMacsec) msg() *otg.DeviceMacsec {
	return obj.obj
}

func (obj *deviceMacsec) setMsg(msg *otg.DeviceMacsec) DeviceMacsec {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceMacsec struct {
	obj *deviceMacsec
}

type marshalDeviceMacsec interface {
	// ToProto marshals DeviceMacsec to protobuf object *otg.DeviceMacsec
	ToProto() (*otg.DeviceMacsec, error)
	// ToPbText marshals DeviceMacsec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceMacsec to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceMacsec to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceMacsec struct {
	obj *deviceMacsec
}

type unMarshalDeviceMacsec interface {
	// FromProto unmarshals DeviceMacsec from protobuf object *otg.DeviceMacsec
	FromProto(msg *otg.DeviceMacsec) (DeviceMacsec, error)
	// FromPbText unmarshals DeviceMacsec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceMacsec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceMacsec from JSON text
	FromJson(value string) error
}

func (obj *deviceMacsec) Marshal() marshalDeviceMacsec {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceMacsec{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceMacsec) Unmarshal() unMarshalDeviceMacsec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceMacsec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceMacsec) ToProto() (*otg.DeviceMacsec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceMacsec) FromProto(msg *otg.DeviceMacsec) (DeviceMacsec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceMacsec) ToPbText() (string, error) {
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

func (m *unMarshaldeviceMacsec) FromPbText(value string) error {
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

func (m *marshaldeviceMacsec) ToYaml() (string, error) {
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

func (m *unMarshaldeviceMacsec) FromYaml(value string) error {
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

func (m *marshaldeviceMacsec) ToJson() (string, error) {
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

func (m *unMarshaldeviceMacsec) FromJson(value string) error {
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

func (obj *deviceMacsec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceMacsec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceMacsec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceMacsec) Clone() (DeviceMacsec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceMacsec()
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

func (obj *deviceMacsec) setNil() {
	obj.ethernetInterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceMacsec is a container of properties for a MACsec capable device. Reference https://1.ieee802.org/security/802-1ae/.
type DeviceMacsec interface {
	Validation
	// msg marshals DeviceMacsec to protobuf object *otg.DeviceMacsec
	// and doesn't set defaults
	msg() *otg.DeviceMacsec
	// setMsg unmarshals DeviceMacsec from protobuf object *otg.DeviceMacsec
	// and doesn't set defaults
	setMsg(*otg.DeviceMacsec) DeviceMacsec
	// provides marshal interface
	Marshal() marshalDeviceMacsec
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceMacsec
	// validate validates DeviceMacsec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceMacsec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetInterfaces returns DeviceMacsecDeviceMacsecEthernetInterfaceIterIter, set in DeviceMacsec
	EthernetInterfaces() DeviceMacsecDeviceMacsecEthernetInterfaceIter
	setNil()
}

// Ethernet Interfaces
// EthernetInterfaces returns a []DeviceMacsecEthernetInterface
func (obj *deviceMacsec) EthernetInterfaces() DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	if len(obj.obj.EthernetInterfaces) == 0 {
		obj.obj.EthernetInterfaces = []*otg.DeviceMacsecEthernetInterface{}
	}
	if obj.ethernetInterfacesHolder == nil {
		obj.ethernetInterfacesHolder = newDeviceMacsecDeviceMacsecEthernetInterfaceIter(&obj.obj.EthernetInterfaces).setMsg(obj)
	}
	return obj.ethernetInterfacesHolder
}

type deviceMacsecDeviceMacsecEthernetInterfaceIter struct {
	obj                                *deviceMacsec
	deviceMacsecEthernetInterfaceSlice []DeviceMacsecEthernetInterface
	fieldPtr                           *[]*otg.DeviceMacsecEthernetInterface
}

func newDeviceMacsecDeviceMacsecEthernetInterfaceIter(ptr *[]*otg.DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	return &deviceMacsecDeviceMacsecEthernetInterfaceIter{fieldPtr: ptr}
}

type DeviceMacsecDeviceMacsecEthernetInterfaceIter interface {
	setMsg(*deviceMacsec) DeviceMacsecDeviceMacsecEthernetInterfaceIter
	Items() []DeviceMacsecEthernetInterface
	Add() DeviceMacsecEthernetInterface
	Append(items ...DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter
	Set(index int, newObj DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter
	Clear() DeviceMacsecDeviceMacsecEthernetInterfaceIter
	clearHolderSlice() DeviceMacsecDeviceMacsecEthernetInterfaceIter
	appendHolderSlice(item DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter
}

func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) setMsg(msg *deviceMacsec) DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceMacsecEthernetInterface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) Items() []DeviceMacsecEthernetInterface {
	return obj.deviceMacsecEthernetInterfaceSlice
}

func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) Add() DeviceMacsecEthernetInterface {
	newObj := &otg.DeviceMacsecEthernetInterface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceMacsecEthernetInterface{obj: newObj}
	newLibObj.setDefault()
	obj.deviceMacsecEthernetInterfaceSlice = append(obj.deviceMacsecEthernetInterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) Append(items ...DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceMacsecEthernetInterfaceSlice = append(obj.deviceMacsecEthernetInterfaceSlice, item)
	}
	return obj
}

func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) Set(index int, newObj DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceMacsecEthernetInterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) Clear() DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceMacsecEthernetInterface{}
		obj.deviceMacsecEthernetInterfaceSlice = []DeviceMacsecEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) clearHolderSlice() DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	if len(obj.deviceMacsecEthernetInterfaceSlice) > 0 {
		obj.deviceMacsecEthernetInterfaceSlice = []DeviceMacsecEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecDeviceMacsecEthernetInterfaceIter) appendHolderSlice(item DeviceMacsecEthernetInterface) DeviceMacsecDeviceMacsecEthernetInterfaceIter {
	obj.deviceMacsecEthernetInterfaceSlice = append(obj.deviceMacsecEthernetInterfaceSlice, item)
	return obj
}

func (obj *deviceMacsec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.EthernetInterfaces) != 0 {

		if set_default {
			obj.EthernetInterfaces().clearHolderSlice()
			for _, item := range obj.obj.EthernetInterfaces {
				obj.EthernetInterfaces().appendHolderSlice(&deviceMacsecEthernetInterface{obj: item})
			}
		}
		for _, item := range obj.EthernetInterfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceMacsec) setDefault() {

}
