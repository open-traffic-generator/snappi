package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceEthernetBase *****
type deviceEthernetBase struct {
	validation
	obj          *otg.DeviceEthernetBase
	marshaller   marshalDeviceEthernetBase
	unMarshaller unMarshalDeviceEthernetBase
	vlansHolder  DeviceEthernetBaseDeviceVlanIter
}

func NewDeviceEthernetBase() DeviceEthernetBase {
	obj := deviceEthernetBase{obj: &otg.DeviceEthernetBase{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceEthernetBase) msg() *otg.DeviceEthernetBase {
	return obj.obj
}

func (obj *deviceEthernetBase) setMsg(msg *otg.DeviceEthernetBase) DeviceEthernetBase {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceEthernetBase struct {
	obj *deviceEthernetBase
}

type marshalDeviceEthernetBase interface {
	// ToProto marshals DeviceEthernetBase to protobuf object *otg.DeviceEthernetBase
	ToProto() (*otg.DeviceEthernetBase, error)
	// ToPbText marshals DeviceEthernetBase to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceEthernetBase to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceEthernetBase to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceEthernetBase struct {
	obj *deviceEthernetBase
}

type unMarshalDeviceEthernetBase interface {
	// FromProto unmarshals DeviceEthernetBase from protobuf object *otg.DeviceEthernetBase
	FromProto(msg *otg.DeviceEthernetBase) (DeviceEthernetBase, error)
	// FromPbText unmarshals DeviceEthernetBase from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceEthernetBase from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceEthernetBase from JSON text
	FromJson(value string) error
}

func (obj *deviceEthernetBase) Marshal() marshalDeviceEthernetBase {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceEthernetBase{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceEthernetBase) Unmarshal() unMarshalDeviceEthernetBase {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceEthernetBase{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceEthernetBase) ToProto() (*otg.DeviceEthernetBase, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceEthernetBase) FromProto(msg *otg.DeviceEthernetBase) (DeviceEthernetBase, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceEthernetBase) ToPbText() (string, error) {
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

func (m *unMarshaldeviceEthernetBase) FromPbText(value string) error {
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

func (m *marshaldeviceEthernetBase) ToYaml() (string, error) {
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

func (m *unMarshaldeviceEthernetBase) FromYaml(value string) error {
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

func (m *marshaldeviceEthernetBase) ToJson() (string, error) {
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

func (m *unMarshaldeviceEthernetBase) FromJson(value string) error {
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

func (obj *deviceEthernetBase) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceEthernetBase) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceEthernetBase) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceEthernetBase) Clone() (DeviceEthernetBase, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceEthernetBase()
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

func (obj *deviceEthernetBase) setNil() {
	obj.vlansHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceEthernetBase is base Ethernet interface.
type DeviceEthernetBase interface {
	Validation
	// msg marshals DeviceEthernetBase to protobuf object *otg.DeviceEthernetBase
	// and doesn't set defaults
	msg() *otg.DeviceEthernetBase
	// setMsg unmarshals DeviceEthernetBase from protobuf object *otg.DeviceEthernetBase
	// and doesn't set defaults
	setMsg(*otg.DeviceEthernetBase) DeviceEthernetBase
	// provides marshal interface
	Marshal() marshalDeviceEthernetBase
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceEthernetBase
	// validate validates DeviceEthernetBase
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceEthernetBase, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mac returns string, set in DeviceEthernetBase.
	Mac() string
	// SetMac assigns string provided by user to DeviceEthernetBase
	SetMac(value string) DeviceEthernetBase
	// HasMac checks if Mac has been set in DeviceEthernetBase
	HasMac() bool
	// Mtu returns uint32, set in DeviceEthernetBase.
	Mtu() uint32
	// SetMtu assigns uint32 provided by user to DeviceEthernetBase
	SetMtu(value uint32) DeviceEthernetBase
	// HasMtu checks if Mtu has been set in DeviceEthernetBase
	HasMtu() bool
	// Vlans returns DeviceEthernetBaseDeviceVlanIterIter, set in DeviceEthernetBase
	Vlans() DeviceEthernetBaseDeviceVlanIter
	// Name returns string, set in DeviceEthernetBase.
	Name() string
	// SetName assigns string provided by user to DeviceEthernetBase
	SetName(value string) DeviceEthernetBase
	setNil()
}

// Media Access Control address.The implementation should ensure that the 'mac' field is explicitly configured by the user for  all types of interfaces as denoted by 'connection' attribute except 'simulated_link' where 'mac' is not mandatory.
// Mac returns a string
func (obj *deviceEthernetBase) Mac() string {

	return *obj.obj.Mac

}

// Media Access Control address.The implementation should ensure that the 'mac' field is explicitly configured by the user for  all types of interfaces as denoted by 'connection' attribute except 'simulated_link' where 'mac' is not mandatory.
// Mac returns a string
func (obj *deviceEthernetBase) HasMac() bool {
	return obj.obj.Mac != nil
}

// Media Access Control address.The implementation should ensure that the 'mac' field is explicitly configured by the user for  all types of interfaces as denoted by 'connection' attribute except 'simulated_link' where 'mac' is not mandatory.
// SetMac sets the string value in the DeviceEthernetBase object
func (obj *deviceEthernetBase) SetMac(value string) DeviceEthernetBase {

	obj.obj.Mac = &value
	return obj
}

// Maximum Transmission Unit.
// Mtu returns a uint32
func (obj *deviceEthernetBase) Mtu() uint32 {

	return *obj.obj.Mtu

}

// Maximum Transmission Unit.
// Mtu returns a uint32
func (obj *deviceEthernetBase) HasMtu() bool {
	return obj.obj.Mtu != nil
}

// Maximum Transmission Unit.
// SetMtu sets the uint32 value in the DeviceEthernetBase object
func (obj *deviceEthernetBase) SetMtu(value uint32) DeviceEthernetBase {

	obj.obj.Mtu = &value
	return obj
}

// List of VLANs
// Vlans returns a []DeviceVlan
func (obj *deviceEthernetBase) Vlans() DeviceEthernetBaseDeviceVlanIter {
	if len(obj.obj.Vlans) == 0 {
		obj.obj.Vlans = []*otg.DeviceVlan{}
	}
	if obj.vlansHolder == nil {
		obj.vlansHolder = newDeviceEthernetBaseDeviceVlanIter(&obj.obj.Vlans).setMsg(obj)
	}
	return obj.vlansHolder
}

type deviceEthernetBaseDeviceVlanIter struct {
	obj             *deviceEthernetBase
	deviceVlanSlice []DeviceVlan
	fieldPtr        *[]*otg.DeviceVlan
}

func newDeviceEthernetBaseDeviceVlanIter(ptr *[]*otg.DeviceVlan) DeviceEthernetBaseDeviceVlanIter {
	return &deviceEthernetBaseDeviceVlanIter{fieldPtr: ptr}
}

type DeviceEthernetBaseDeviceVlanIter interface {
	setMsg(*deviceEthernetBase) DeviceEthernetBaseDeviceVlanIter
	Items() []DeviceVlan
	Add() DeviceVlan
	Append(items ...DeviceVlan) DeviceEthernetBaseDeviceVlanIter
	Set(index int, newObj DeviceVlan) DeviceEthernetBaseDeviceVlanIter
	Clear() DeviceEthernetBaseDeviceVlanIter
	clearHolderSlice() DeviceEthernetBaseDeviceVlanIter
	appendHolderSlice(item DeviceVlan) DeviceEthernetBaseDeviceVlanIter
}

func (obj *deviceEthernetBaseDeviceVlanIter) setMsg(msg *deviceEthernetBase) DeviceEthernetBaseDeviceVlanIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceVlan{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceEthernetBaseDeviceVlanIter) Items() []DeviceVlan {
	return obj.deviceVlanSlice
}

func (obj *deviceEthernetBaseDeviceVlanIter) Add() DeviceVlan {
	newObj := &otg.DeviceVlan{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceVlan{obj: newObj}
	newLibObj.setDefault()
	obj.deviceVlanSlice = append(obj.deviceVlanSlice, newLibObj)
	return newLibObj
}

func (obj *deviceEthernetBaseDeviceVlanIter) Append(items ...DeviceVlan) DeviceEthernetBaseDeviceVlanIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceVlanSlice = append(obj.deviceVlanSlice, item)
	}
	return obj
}

func (obj *deviceEthernetBaseDeviceVlanIter) Set(index int, newObj DeviceVlan) DeviceEthernetBaseDeviceVlanIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceVlanSlice[index] = newObj
	return obj
}
func (obj *deviceEthernetBaseDeviceVlanIter) Clear() DeviceEthernetBaseDeviceVlanIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceVlan{}
		obj.deviceVlanSlice = []DeviceVlan{}
	}
	return obj
}
func (obj *deviceEthernetBaseDeviceVlanIter) clearHolderSlice() DeviceEthernetBaseDeviceVlanIter {
	if len(obj.deviceVlanSlice) > 0 {
		obj.deviceVlanSlice = []DeviceVlan{}
	}
	return obj
}
func (obj *deviceEthernetBaseDeviceVlanIter) appendHolderSlice(item DeviceVlan) DeviceEthernetBaseDeviceVlanIter {
	obj.deviceVlanSlice = append(obj.deviceVlanSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceEthernetBase) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceEthernetBase object
func (obj *deviceEthernetBase) SetName(value string) DeviceEthernetBase {

	obj.obj.Name = &value
	return obj
}

func (obj *deviceEthernetBase) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceEthernetBase.Mac"))
		}

	}

	if obj.obj.Mtu != nil {

		if *obj.obj.Mtu > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= DeviceEthernetBase.Mtu <= 65535 but Got %d", *obj.obj.Mtu))
		}

	}

	if len(obj.obj.Vlans) != 0 {

		if set_default {
			obj.Vlans().clearHolderSlice()
			for _, item := range obj.obj.Vlans {
				obj.Vlans().appendHolderSlice(&deviceVlan{obj: item})
			}
		}
		for _, item := range obj.Vlans().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceEthernetBase")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DeviceEthernetBase.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *deviceEthernetBase) setDefault() {
	if obj.obj.Mtu == nil {
		obj.SetMtu(1500)
	}

}
