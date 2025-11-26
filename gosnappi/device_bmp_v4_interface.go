package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpV4Interface *****
type deviceBmpV4Interface struct {
	validation
	obj           *otg.DeviceBmpV4Interface
	marshaller    marshalDeviceBmpV4Interface
	unMarshaller  unMarshalDeviceBmpV4Interface
	serversHolder DeviceBmpV4InterfaceDeviceBmpServerV4Iter
}

func NewDeviceBmpV4Interface() DeviceBmpV4Interface {
	obj := deviceBmpV4Interface{obj: &otg.DeviceBmpV4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpV4Interface) msg() *otg.DeviceBmpV4Interface {
	return obj.obj
}

func (obj *deviceBmpV4Interface) setMsg(msg *otg.DeviceBmpV4Interface) DeviceBmpV4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpV4Interface struct {
	obj *deviceBmpV4Interface
}

type marshalDeviceBmpV4Interface interface {
	// ToProto marshals DeviceBmpV4Interface to protobuf object *otg.DeviceBmpV4Interface
	ToProto() (*otg.DeviceBmpV4Interface, error)
	// ToPbText marshals DeviceBmpV4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpV4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpV4Interface to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpV4Interface struct {
	obj *deviceBmpV4Interface
}

type unMarshalDeviceBmpV4Interface interface {
	// FromProto unmarshals DeviceBmpV4Interface from protobuf object *otg.DeviceBmpV4Interface
	FromProto(msg *otg.DeviceBmpV4Interface) (DeviceBmpV4Interface, error)
	// FromPbText unmarshals DeviceBmpV4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpV4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpV4Interface from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpV4Interface) Marshal() marshalDeviceBmpV4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpV4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpV4Interface) Unmarshal() unMarshalDeviceBmpV4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpV4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpV4Interface) ToProto() (*otg.DeviceBmpV4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpV4Interface) FromProto(msg *otg.DeviceBmpV4Interface) (DeviceBmpV4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpV4Interface) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpV4Interface) FromPbText(value string) error {
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

func (m *marshaldeviceBmpV4Interface) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpV4Interface) FromYaml(value string) error {
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

func (m *marshaldeviceBmpV4Interface) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpV4Interface) FromJson(value string) error {
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

func (obj *deviceBmpV4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpV4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpV4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpV4Interface) Clone() (DeviceBmpV4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpV4Interface()
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

func (obj *deviceBmpV4Interface) setNil() {
	obj.serversHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpV4Interface is configuration for BMP Servers on a single IPv4 interface.
type DeviceBmpV4Interface interface {
	Validation
	// msg marshals DeviceBmpV4Interface to protobuf object *otg.DeviceBmpV4Interface
	// and doesn't set defaults
	msg() *otg.DeviceBmpV4Interface
	// setMsg unmarshals DeviceBmpV4Interface from protobuf object *otg.DeviceBmpV4Interface
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpV4Interface) DeviceBmpV4Interface
	// provides marshal interface
	Marshal() marshalDeviceBmpV4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpV4Interface
	// validate validates DeviceBmpV4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpV4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in DeviceBmpV4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to DeviceBmpV4Interface
	SetIpv4Name(value string) DeviceBmpV4Interface
	// Servers returns DeviceBmpV4InterfaceDeviceBmpServerV4IterIter, set in DeviceBmpV4Interface
	Servers() DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	setNil()
}

// The unique name of the IPv4 interface used as the source IP for the BMP Server.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *deviceBmpV4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The unique name of the IPv4 interface used as the source IP for the BMP Server.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the DeviceBmpV4Interface object
func (obj *deviceBmpV4Interface) SetIpv4Name(value string) DeviceBmpV4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// This contains the configuration of BMP Servers configured on this IPv4 interface.
// Servers returns a []DeviceBmpServerV4
func (obj *deviceBmpV4Interface) Servers() DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	if len(obj.obj.Servers) == 0 {
		obj.obj.Servers = []*otg.DeviceBmpServerV4{}
	}
	if obj.serversHolder == nil {
		obj.serversHolder = newDeviceBmpV4InterfaceDeviceBmpServerV4Iter(&obj.obj.Servers).setMsg(obj)
	}
	return obj.serversHolder
}

type deviceBmpV4InterfaceDeviceBmpServerV4Iter struct {
	obj                    *deviceBmpV4Interface
	deviceBmpServerV4Slice []DeviceBmpServerV4
	fieldPtr               *[]*otg.DeviceBmpServerV4
}

func newDeviceBmpV4InterfaceDeviceBmpServerV4Iter(ptr *[]*otg.DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	return &deviceBmpV4InterfaceDeviceBmpServerV4Iter{fieldPtr: ptr}
}

type DeviceBmpV4InterfaceDeviceBmpServerV4Iter interface {
	setMsg(*deviceBmpV4Interface) DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	Items() []DeviceBmpServerV4
	Add() DeviceBmpServerV4
	Append(items ...DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	Set(index int, newObj DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	Clear() DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	clearHolderSlice() DeviceBmpV4InterfaceDeviceBmpServerV4Iter
	appendHolderSlice(item DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter
}

func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) setMsg(msg *deviceBmpV4Interface) DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpServerV4{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) Items() []DeviceBmpServerV4 {
	return obj.deviceBmpServerV4Slice
}

func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) Add() DeviceBmpServerV4 {
	newObj := &otg.DeviceBmpServerV4{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpServerV4{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpServerV4Slice = append(obj.deviceBmpServerV4Slice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) Append(items ...DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpServerV4Slice = append(obj.deviceBmpServerV4Slice, item)
	}
	return obj
}

func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) Set(index int, newObj DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpServerV4Slice[index] = newObj
	return obj
}
func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) Clear() DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpServerV4{}
		obj.deviceBmpServerV4Slice = []DeviceBmpServerV4{}
	}
	return obj
}
func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) clearHolderSlice() DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	if len(obj.deviceBmpServerV4Slice) > 0 {
		obj.deviceBmpServerV4Slice = []DeviceBmpServerV4{}
	}
	return obj
}
func (obj *deviceBmpV4InterfaceDeviceBmpServerV4Iter) appendHolderSlice(item DeviceBmpServerV4) DeviceBmpV4InterfaceDeviceBmpServerV4Iter {
	obj.deviceBmpServerV4Slice = append(obj.deviceBmpServerV4Slice, item)
	return obj
}

func (obj *deviceBmpV4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface DeviceBmpV4Interface")
	}

	if len(obj.obj.Servers) != 0 {

		if set_default {
			obj.Servers().clearHolderSlice()
			for _, item := range obj.obj.Servers {
				obj.Servers().appendHolderSlice(&deviceBmpServerV4{obj: item})
			}
		}
		for _, item := range obj.Servers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBmpV4Interface) setDefault() {

}
