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
	ethernetInterfacesHolder DeviceMacsecMacsecEthernetInterfaceIter
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

// DeviceMacsec is a container of properties for a MACsec capable device.
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
	// EthernetInterfaces returns DeviceMacsecMacsecEthernetInterfaceIterIter, set in DeviceMacsec
	EthernetInterfaces() DeviceMacsecMacsecEthernetInterfaceIter
	setNil()
}

// Ethernet Interfaces
// EthernetInterfaces returns a []MacsecEthernetInterface
func (obj *deviceMacsec) EthernetInterfaces() DeviceMacsecMacsecEthernetInterfaceIter {
	if len(obj.obj.EthernetInterfaces) == 0 {
		obj.obj.EthernetInterfaces = []*otg.MacsecEthernetInterface{}
	}
	if obj.ethernetInterfacesHolder == nil {
		obj.ethernetInterfacesHolder = newDeviceMacsecMacsecEthernetInterfaceIter(&obj.obj.EthernetInterfaces).setMsg(obj)
	}
	return obj.ethernetInterfacesHolder
}

type deviceMacsecMacsecEthernetInterfaceIter struct {
	obj                          *deviceMacsec
	macsecEthernetInterfaceSlice []MacsecEthernetInterface
	fieldPtr                     *[]*otg.MacsecEthernetInterface
}

func newDeviceMacsecMacsecEthernetInterfaceIter(ptr *[]*otg.MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter {
	return &deviceMacsecMacsecEthernetInterfaceIter{fieldPtr: ptr}
}

type DeviceMacsecMacsecEthernetInterfaceIter interface {
	setMsg(*deviceMacsec) DeviceMacsecMacsecEthernetInterfaceIter
	Items() []MacsecEthernetInterface
	Add() MacsecEthernetInterface
	Append(items ...MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter
	Set(index int, newObj MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter
	Clear() DeviceMacsecMacsecEthernetInterfaceIter
	clearHolderSlice() DeviceMacsecMacsecEthernetInterfaceIter
	appendHolderSlice(item MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter
}

func (obj *deviceMacsecMacsecEthernetInterfaceIter) setMsg(msg *deviceMacsec) DeviceMacsecMacsecEthernetInterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecEthernetInterface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceMacsecMacsecEthernetInterfaceIter) Items() []MacsecEthernetInterface {
	return obj.macsecEthernetInterfaceSlice
}

func (obj *deviceMacsecMacsecEthernetInterfaceIter) Add() MacsecEthernetInterface {
	newObj := &otg.MacsecEthernetInterface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecEthernetInterface{obj: newObj}
	newLibObj.setDefault()
	obj.macsecEthernetInterfaceSlice = append(obj.macsecEthernetInterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceMacsecMacsecEthernetInterfaceIter) Append(items ...MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecEthernetInterfaceSlice = append(obj.macsecEthernetInterfaceSlice, item)
	}
	return obj
}

func (obj *deviceMacsecMacsecEthernetInterfaceIter) Set(index int, newObj MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecEthernetInterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceMacsecMacsecEthernetInterfaceIter) Clear() DeviceMacsecMacsecEthernetInterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecEthernetInterface{}
		obj.macsecEthernetInterfaceSlice = []MacsecEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecMacsecEthernetInterfaceIter) clearHolderSlice() DeviceMacsecMacsecEthernetInterfaceIter {
	if len(obj.macsecEthernetInterfaceSlice) > 0 {
		obj.macsecEthernetInterfaceSlice = []MacsecEthernetInterface{}
	}
	return obj
}
func (obj *deviceMacsecMacsecEthernetInterfaceIter) appendHolderSlice(item MacsecEthernetInterface) DeviceMacsecMacsecEthernetInterfaceIter {
	obj.macsecEthernetInterfaceSlice = append(obj.macsecEthernetInterfaceSlice, item)
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
				obj.EthernetInterfaces().appendHolderSlice(&macsecEthernetInterface{obj: item})
			}
		}
		for _, item := range obj.EthernetInterfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceMacsec) setDefault() {

}
