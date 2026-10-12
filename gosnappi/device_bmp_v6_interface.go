package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpV6Interface *****
type deviceBmpV6Interface struct {
	validation
	obj           *otg.DeviceBmpV6Interface
	marshaller    marshalDeviceBmpV6Interface
	unMarshaller  unMarshalDeviceBmpV6Interface
	serversHolder DeviceBmpV6InterfaceDeviceBmpServerV6Iter
}

func NewDeviceBmpV6Interface() DeviceBmpV6Interface {
	obj := deviceBmpV6Interface{obj: &otg.DeviceBmpV6Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpV6Interface) msg() *otg.DeviceBmpV6Interface {
	return obj.obj
}

func (obj *deviceBmpV6Interface) setMsg(msg *otg.DeviceBmpV6Interface) DeviceBmpV6Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpV6Interface struct {
	obj *deviceBmpV6Interface
}

type marshalDeviceBmpV6Interface interface {
	// ToProto marshals DeviceBmpV6Interface to protobuf object *otg.DeviceBmpV6Interface
	ToProto() (*otg.DeviceBmpV6Interface, error)
	// ToPbText marshals DeviceBmpV6Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpV6Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpV6Interface to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpV6Interface struct {
	obj *deviceBmpV6Interface
}

type unMarshalDeviceBmpV6Interface interface {
	// FromProto unmarshals DeviceBmpV6Interface from protobuf object *otg.DeviceBmpV6Interface
	FromProto(msg *otg.DeviceBmpV6Interface) (DeviceBmpV6Interface, error)
	// FromPbText unmarshals DeviceBmpV6Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpV6Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpV6Interface from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpV6Interface) Marshal() marshalDeviceBmpV6Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpV6Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpV6Interface) Unmarshal() unMarshalDeviceBmpV6Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpV6Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpV6Interface) ToProto() (*otg.DeviceBmpV6Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpV6Interface) FromProto(msg *otg.DeviceBmpV6Interface) (DeviceBmpV6Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpV6Interface) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpV6Interface) FromPbText(value string) error {
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

func (m *marshaldeviceBmpV6Interface) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpV6Interface) FromYaml(value string) error {
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

func (m *marshaldeviceBmpV6Interface) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpV6Interface) FromJson(value string) error {
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

func (obj *deviceBmpV6Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpV6Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpV6Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpV6Interface) Clone() (DeviceBmpV6Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpV6Interface()
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

func (obj *deviceBmpV6Interface) setNil() {
	obj.serversHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpV6Interface is configuration for BMP Servers on a single IPv6 interface.
type DeviceBmpV6Interface interface {
	Validation
	// msg marshals DeviceBmpV6Interface to protobuf object *otg.DeviceBmpV6Interface
	// and doesn't set defaults
	msg() *otg.DeviceBmpV6Interface
	// setMsg unmarshals DeviceBmpV6Interface from protobuf object *otg.DeviceBmpV6Interface
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpV6Interface) DeviceBmpV6Interface
	// provides marshal interface
	Marshal() marshalDeviceBmpV6Interface
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpV6Interface
	// validate validates DeviceBmpV6Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpV6Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6Name returns string, set in DeviceBmpV6Interface.
	Ipv6Name() string
	// SetIpv6Name assigns string provided by user to DeviceBmpV6Interface
	SetIpv6Name(value string) DeviceBmpV6Interface
	// Servers returns DeviceBmpV6InterfaceDeviceBmpServerV6IterIter, set in DeviceBmpV6Interface
	Servers() DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	setNil()
}

// The unique name of the IPv6 interface used as the source IP for BMP Server.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// Ipv6Name returns a string
func (obj *deviceBmpV6Interface) Ipv6Name() string {

	return *obj.obj.Ipv6Name

}

// The unique name of the IPv6 interface used as the source IP for BMP Server.
//
// x-constraint:
// - /components/schemas/Device.Ipv6/properties/name
//
// SetIpv6Name sets the string value in the DeviceBmpV6Interface object
func (obj *deviceBmpV6Interface) SetIpv6Name(value string) DeviceBmpV6Interface {

	obj.obj.Ipv6Name = &value
	return obj
}

// This contains the configuration of BMP Servers configured on this IPv6 interface.
// Servers returns a []DeviceBmpServerV6
func (obj *deviceBmpV6Interface) Servers() DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	if len(obj.obj.Servers) == 0 {
		obj.obj.Servers = []*otg.DeviceBmpServerV6{}
	}
	if obj.serversHolder == nil {
		obj.serversHolder = newDeviceBmpV6InterfaceDeviceBmpServerV6Iter(&obj.obj.Servers).setMsg(obj)
	}
	return obj.serversHolder
}

type deviceBmpV6InterfaceDeviceBmpServerV6Iter struct {
	obj                    *deviceBmpV6Interface
	deviceBmpServerV6Slice []DeviceBmpServerV6
	fieldPtr               *[]*otg.DeviceBmpServerV6
}

func newDeviceBmpV6InterfaceDeviceBmpServerV6Iter(ptr *[]*otg.DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	return &deviceBmpV6InterfaceDeviceBmpServerV6Iter{fieldPtr: ptr}
}

type DeviceBmpV6InterfaceDeviceBmpServerV6Iter interface {
	setMsg(*deviceBmpV6Interface) DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	Items() []DeviceBmpServerV6
	Add() DeviceBmpServerV6
	Append(items ...DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	Set(index int, newObj DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	Clear() DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	clearHolderSlice() DeviceBmpV6InterfaceDeviceBmpServerV6Iter
	appendHolderSlice(item DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter
}

func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) setMsg(msg *deviceBmpV6Interface) DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpServerV6{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) Items() []DeviceBmpServerV6 {
	return obj.deviceBmpServerV6Slice
}

func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) Add() DeviceBmpServerV6 {
	newObj := &otg.DeviceBmpServerV6{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpServerV6{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpServerV6Slice = append(obj.deviceBmpServerV6Slice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) Append(items ...DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpServerV6Slice = append(obj.deviceBmpServerV6Slice, item)
	}
	return obj
}

func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) Set(index int, newObj DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpServerV6Slice[index] = newObj
	return obj
}
func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) Clear() DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpServerV6{}
		obj.deviceBmpServerV6Slice = []DeviceBmpServerV6{}
	}
	return obj
}
func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) clearHolderSlice() DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	if len(obj.deviceBmpServerV6Slice) > 0 {
		obj.deviceBmpServerV6Slice = []DeviceBmpServerV6{}
	}
	return obj
}
func (obj *deviceBmpV6InterfaceDeviceBmpServerV6Iter) appendHolderSlice(item DeviceBmpServerV6) DeviceBmpV6InterfaceDeviceBmpServerV6Iter {
	obj.deviceBmpServerV6Slice = append(obj.deviceBmpServerV6Slice, item)
	return obj
}

func (obj *deviceBmpV6Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv6Name is required
	if obj.obj.Ipv6Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6Name is required field on interface DeviceBmpV6Interface")
	}

	if len(obj.obj.Servers) != 0 {

		if set_default {
			obj.Servers().clearHolderSlice()
			for _, item := range obj.obj.Servers {
				obj.Servers().appendHolderSlice(&deviceBmpServerV6{obj: item})
			}
		}
		for _, item := range obj.Servers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBmpV6Interface) setDefault() {

}
