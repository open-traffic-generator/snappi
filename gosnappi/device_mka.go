package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceMka *****
type deviceMka struct {
	validation
	obj                      *otg.DeviceMka
	marshaller               marshalDeviceMka
	unMarshaller             unMarshalDeviceMka
	ethernetInterfacesHolder DeviceMkaDeviceMkaEthernetInterfaceIter
}

func NewDeviceMka() DeviceMka {
	obj := deviceMka{obj: &otg.DeviceMka{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceMka) msg() *otg.DeviceMka {
	return obj.obj
}

func (obj *deviceMka) setMsg(msg *otg.DeviceMka) DeviceMka {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceMka struct {
	obj *deviceMka
}

type marshalDeviceMka interface {
	// ToProto marshals DeviceMka to protobuf object *otg.DeviceMka
	ToProto() (*otg.DeviceMka, error)
	// ToPbText marshals DeviceMka to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceMka to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceMka to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceMka struct {
	obj *deviceMka
}

type unMarshalDeviceMka interface {
	// FromProto unmarshals DeviceMka from protobuf object *otg.DeviceMka
	FromProto(msg *otg.DeviceMka) (DeviceMka, error)
	// FromPbText unmarshals DeviceMka from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceMka from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceMka from JSON text
	FromJson(value string) error
}

func (obj *deviceMka) Marshal() marshalDeviceMka {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceMka{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceMka) Unmarshal() unMarshalDeviceMka {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceMka{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceMka) ToProto() (*otg.DeviceMka, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceMka) FromProto(msg *otg.DeviceMka) (DeviceMka, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceMka) ToPbText() (string, error) {
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

func (m *unMarshaldeviceMka) FromPbText(value string) error {
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

func (m *marshaldeviceMka) ToYaml() (string, error) {
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

func (m *unMarshaldeviceMka) FromYaml(value string) error {
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

func (m *marshaldeviceMka) ToJson() (string, error) {
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

func (m *unMarshaldeviceMka) FromJson(value string) error {
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

func (obj *deviceMka) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceMka) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceMka) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceMka) Clone() (DeviceMka, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceMka()
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

func (obj *deviceMka) setNil() {
	obj.ethernetInterfacesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceMka is a container of properties for a MKA supplicant.
type DeviceMka interface {
	Validation
	// msg marshals DeviceMka to protobuf object *otg.DeviceMka
	// and doesn't set defaults
	msg() *otg.DeviceMka
	// setMsg unmarshals DeviceMka from protobuf object *otg.DeviceMka
	// and doesn't set defaults
	setMsg(*otg.DeviceMka) DeviceMka
	// provides marshal interface
	Marshal() marshalDeviceMka
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceMka
	// validate validates DeviceMka
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceMka, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetInterfaces returns DeviceMkaDeviceMkaEthernetInterfaceIterIter, set in DeviceMka
	EthernetInterfaces() DeviceMkaDeviceMkaEthernetInterfaceIter
	setNil()
}

// Ethernet Interfaces
// EthernetInterfaces returns a []DeviceMkaEthernetInterface
func (obj *deviceMka) EthernetInterfaces() DeviceMkaDeviceMkaEthernetInterfaceIter {
	if len(obj.obj.EthernetInterfaces) == 0 {
		obj.obj.EthernetInterfaces = []*otg.DeviceMkaEthernetInterface{}
	}
	if obj.ethernetInterfacesHolder == nil {
		obj.ethernetInterfacesHolder = newDeviceMkaDeviceMkaEthernetInterfaceIter(&obj.obj.EthernetInterfaces).setMsg(obj)
	}
	return obj.ethernetInterfacesHolder
}

type deviceMkaDeviceMkaEthernetInterfaceIter struct {
	obj                             *deviceMka
	deviceMkaEthernetInterfaceSlice []DeviceMkaEthernetInterface
	fieldPtr                        *[]*otg.DeviceMkaEthernetInterface
}

func newDeviceMkaDeviceMkaEthernetInterfaceIter(ptr *[]*otg.DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter {
	return &deviceMkaDeviceMkaEthernetInterfaceIter{fieldPtr: ptr}
}

type DeviceMkaDeviceMkaEthernetInterfaceIter interface {
	setMsg(*deviceMka) DeviceMkaDeviceMkaEthernetInterfaceIter
	Items() []DeviceMkaEthernetInterface
	Add() DeviceMkaEthernetInterface
	Append(items ...DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter
	Set(index int, newObj DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter
	Clear() DeviceMkaDeviceMkaEthernetInterfaceIter
	clearHolderSlice() DeviceMkaDeviceMkaEthernetInterfaceIter
	appendHolderSlice(item DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter
}

func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) setMsg(msg *deviceMka) DeviceMkaDeviceMkaEthernetInterfaceIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceMkaEthernetInterface{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) Items() []DeviceMkaEthernetInterface {
	return obj.deviceMkaEthernetInterfaceSlice
}

func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) Add() DeviceMkaEthernetInterface {
	newObj := &otg.DeviceMkaEthernetInterface{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceMkaEthernetInterface{obj: newObj}
	newLibObj.setDefault()
	obj.deviceMkaEthernetInterfaceSlice = append(obj.deviceMkaEthernetInterfaceSlice, newLibObj)
	return newLibObj
}

func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) Append(items ...DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceMkaEthernetInterfaceSlice = append(obj.deviceMkaEthernetInterfaceSlice, item)
	}
	return obj
}

func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) Set(index int, newObj DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceMkaEthernetInterfaceSlice[index] = newObj
	return obj
}
func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) Clear() DeviceMkaDeviceMkaEthernetInterfaceIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceMkaEthernetInterface{}
		obj.deviceMkaEthernetInterfaceSlice = []DeviceMkaEthernetInterface{}
	}
	return obj
}
func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) clearHolderSlice() DeviceMkaDeviceMkaEthernetInterfaceIter {
	if len(obj.deviceMkaEthernetInterfaceSlice) > 0 {
		obj.deviceMkaEthernetInterfaceSlice = []DeviceMkaEthernetInterface{}
	}
	return obj
}
func (obj *deviceMkaDeviceMkaEthernetInterfaceIter) appendHolderSlice(item DeviceMkaEthernetInterface) DeviceMkaDeviceMkaEthernetInterfaceIter {
	obj.deviceMkaEthernetInterfaceSlice = append(obj.deviceMkaEthernetInterfaceSlice, item)
	return obj
}

func (obj *deviceMka) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.EthernetInterfaces) != 0 {

		if set_default {
			obj.EthernetInterfaces().clearHolderSlice()
			for _, item := range obj.obj.EthernetInterfaces {
				obj.EthernetInterfaces().appendHolderSlice(&deviceMkaEthernetInterface{obj: item})
			}
		}
		for _, item := range obj.EthernetInterfaces().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceMka) setDefault() {

}
