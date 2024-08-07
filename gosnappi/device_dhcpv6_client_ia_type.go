package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientIaType *****
type deviceDhcpv6ClientIaType struct {
	validation
	obj          *otg.DeviceDhcpv6ClientIaType
	marshaller   marshalDeviceDhcpv6ClientIaType
	unMarshaller unMarshalDeviceDhcpv6ClientIaType
	ianaHolder   DeviceDhcpv6ClientIaTimeValue
	iapdHolder   DeviceDhcpv6ClientIaTimeValue
	ianapdHolder DeviceDhcpv6ClientIaTimeValue
}

func NewDeviceDhcpv6ClientIaType() DeviceDhcpv6ClientIaType {
	obj := deviceDhcpv6ClientIaType{obj: &otg.DeviceDhcpv6ClientIaType{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientIaType) msg() *otg.DeviceDhcpv6ClientIaType {
	return obj.obj
}

func (obj *deviceDhcpv6ClientIaType) setMsg(msg *otg.DeviceDhcpv6ClientIaType) DeviceDhcpv6ClientIaType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientIaType struct {
	obj *deviceDhcpv6ClientIaType
}

type marshalDeviceDhcpv6ClientIaType interface {
	// ToProto marshals DeviceDhcpv6ClientIaType to protobuf object *otg.DeviceDhcpv6ClientIaType
	ToProto() (*otg.DeviceDhcpv6ClientIaType, error)
	// ToPbText marshals DeviceDhcpv6ClientIaType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientIaType to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientIaType to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6ClientIaType struct {
	obj *deviceDhcpv6ClientIaType
}

type unMarshalDeviceDhcpv6ClientIaType interface {
	// FromProto unmarshals DeviceDhcpv6ClientIaType from protobuf object *otg.DeviceDhcpv6ClientIaType
	FromProto(msg *otg.DeviceDhcpv6ClientIaType) (DeviceDhcpv6ClientIaType, error)
	// FromPbText unmarshals DeviceDhcpv6ClientIaType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientIaType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientIaType from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientIaType) Marshal() marshalDeviceDhcpv6ClientIaType {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientIaType{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientIaType) Unmarshal() unMarshalDeviceDhcpv6ClientIaType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientIaType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientIaType) ToProto() (*otg.DeviceDhcpv6ClientIaType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientIaType) FromProto(msg *otg.DeviceDhcpv6ClientIaType) (DeviceDhcpv6ClientIaType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientIaType) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaType) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientIaType) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaType) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientIaType) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientIaType) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientIaType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientIaType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientIaType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientIaType) Clone() (DeviceDhcpv6ClientIaType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientIaType()
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

func (obj *deviceDhcpv6ClientIaType) setNil() {
	obj.ianaHolder = nil
	obj.iapdHolder = nil
	obj.ianapdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6ClientIaType is description is TBD
type DeviceDhcpv6ClientIaType interface {
	Validation
	// msg marshals DeviceDhcpv6ClientIaType to protobuf object *otg.DeviceDhcpv6ClientIaType
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientIaType
	// setMsg unmarshals DeviceDhcpv6ClientIaType from protobuf object *otg.DeviceDhcpv6ClientIaType
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientIaType) DeviceDhcpv6ClientIaType
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientIaType
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientIaType
	// validate validates DeviceDhcpv6ClientIaType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientIaType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceDhcpv6ClientIaTypeChoiceEnum, set in DeviceDhcpv6ClientIaType
	Choice() DeviceDhcpv6ClientIaTypeChoiceEnum
	// setChoice assigns DeviceDhcpv6ClientIaTypeChoiceEnum provided by user to DeviceDhcpv6ClientIaType
	setChoice(value DeviceDhcpv6ClientIaTypeChoiceEnum) DeviceDhcpv6ClientIaType
	// HasChoice checks if Choice has been set in DeviceDhcpv6ClientIaType
	HasChoice() bool
	// getter for Iata to set choice.
	Iata()
	// Iana returns DeviceDhcpv6ClientIaTimeValue, set in DeviceDhcpv6ClientIaType.
	Iana() DeviceDhcpv6ClientIaTimeValue
	// SetIana assigns DeviceDhcpv6ClientIaTimeValue provided by user to DeviceDhcpv6ClientIaType.
	SetIana(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType
	// HasIana checks if Iana has been set in DeviceDhcpv6ClientIaType
	HasIana() bool
	// Iapd returns DeviceDhcpv6ClientIaTimeValue, set in DeviceDhcpv6ClientIaType.
	Iapd() DeviceDhcpv6ClientIaTimeValue
	// SetIapd assigns DeviceDhcpv6ClientIaTimeValue provided by user to DeviceDhcpv6ClientIaType.
	SetIapd(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType
	// HasIapd checks if Iapd has been set in DeviceDhcpv6ClientIaType
	HasIapd() bool
	// Ianapd returns DeviceDhcpv6ClientIaTimeValue, set in DeviceDhcpv6ClientIaType.
	Ianapd() DeviceDhcpv6ClientIaTimeValue
	// SetIanapd assigns DeviceDhcpv6ClientIaTimeValue provided by user to DeviceDhcpv6ClientIaType.
	SetIanapd(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType
	// HasIanapd checks if Ianapd has been set in DeviceDhcpv6ClientIaType
	HasIanapd() bool
	setNil()
}

type DeviceDhcpv6ClientIaTypeChoiceEnum string

// Enum of Choice on DeviceDhcpv6ClientIaType
var DeviceDhcpv6ClientIaTypeChoice = struct {
	IANA   DeviceDhcpv6ClientIaTypeChoiceEnum
	IATA   DeviceDhcpv6ClientIaTypeChoiceEnum
	IAPD   DeviceDhcpv6ClientIaTypeChoiceEnum
	IANAPD DeviceDhcpv6ClientIaTypeChoiceEnum
}{
	IANA:   DeviceDhcpv6ClientIaTypeChoiceEnum("iana"),
	IATA:   DeviceDhcpv6ClientIaTypeChoiceEnum("iata"),
	IAPD:   DeviceDhcpv6ClientIaTypeChoiceEnum("iapd"),
	IANAPD: DeviceDhcpv6ClientIaTypeChoiceEnum("ianapd"),
}

func (obj *deviceDhcpv6ClientIaType) Choice() DeviceDhcpv6ClientIaTypeChoiceEnum {
	return DeviceDhcpv6ClientIaTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Iata to set choice
func (obj *deviceDhcpv6ClientIaType) Iata() {
	obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IATA)
}

// Identity Association: Each IA has an associated IAID. IA_NA and IA_TA options represent different types of  IPv6 addresses and parameters accepted by DHCPv6 clients each used in different context by an IPv6 node. IA_NA  is the Identity Association for Non-temporary Addresses option. IA_TA is the Identity Association for Temporary  Addresses option. IA_PD and IA_NAPD options represent one or more IPv6 prefix and parameters. IA_PD is the Identity  Association for Prefix Delegation and IA_NAPD s the Identity Association for Temporary Prefix Delegation.
// Choice returns a string
func (obj *deviceDhcpv6ClientIaType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceDhcpv6ClientIaType) setChoice(value DeviceDhcpv6ClientIaTypeChoiceEnum) DeviceDhcpv6ClientIaType {
	intValue, ok := otg.DeviceDhcpv6ClientIaType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceDhcpv6ClientIaTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceDhcpv6ClientIaType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ianapd = nil
	obj.ianapdHolder = nil
	obj.obj.Iapd = nil
	obj.iapdHolder = nil
	obj.obj.Iana = nil
	obj.ianaHolder = nil

	if value == DeviceDhcpv6ClientIaTypeChoice.IANA {
		obj.obj.Iana = NewDeviceDhcpv6ClientIaTimeValue().msg()
	}

	if value == DeviceDhcpv6ClientIaTypeChoice.IAPD {
		obj.obj.Iapd = NewDeviceDhcpv6ClientIaTimeValue().msg()
	}

	if value == DeviceDhcpv6ClientIaTypeChoice.IANAPD {
		obj.obj.Ianapd = NewDeviceDhcpv6ClientIaTimeValue().msg()
	}

	return obj
}

// description is TBD
// Iana returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) Iana() DeviceDhcpv6ClientIaTimeValue {
	if obj.obj.Iana == nil {
		obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IANA)
	}
	if obj.ianaHolder == nil {
		obj.ianaHolder = &deviceDhcpv6ClientIaTimeValue{obj: obj.obj.Iana}
	}
	return obj.ianaHolder
}

// description is TBD
// Iana returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) HasIana() bool {
	return obj.obj.Iana != nil
}

// description is TBD
// SetIana sets the DeviceDhcpv6ClientIaTimeValue value in the DeviceDhcpv6ClientIaType object
func (obj *deviceDhcpv6ClientIaType) SetIana(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType {
	obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IANA)
	obj.ianaHolder = nil
	obj.obj.Iana = value.msg()

	return obj
}

// description is TBD
// Iapd returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) Iapd() DeviceDhcpv6ClientIaTimeValue {
	if obj.obj.Iapd == nil {
		obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IAPD)
	}
	if obj.iapdHolder == nil {
		obj.iapdHolder = &deviceDhcpv6ClientIaTimeValue{obj: obj.obj.Iapd}
	}
	return obj.iapdHolder
}

// description is TBD
// Iapd returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) HasIapd() bool {
	return obj.obj.Iapd != nil
}

// description is TBD
// SetIapd sets the DeviceDhcpv6ClientIaTimeValue value in the DeviceDhcpv6ClientIaType object
func (obj *deviceDhcpv6ClientIaType) SetIapd(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType {
	obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IAPD)
	obj.iapdHolder = nil
	obj.obj.Iapd = value.msg()

	return obj
}

// description is TBD
// Ianapd returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) Ianapd() DeviceDhcpv6ClientIaTimeValue {
	if obj.obj.Ianapd == nil {
		obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IANAPD)
	}
	if obj.ianapdHolder == nil {
		obj.ianapdHolder = &deviceDhcpv6ClientIaTimeValue{obj: obj.obj.Ianapd}
	}
	return obj.ianapdHolder
}

// description is TBD
// Ianapd returns a DeviceDhcpv6ClientIaTimeValue
func (obj *deviceDhcpv6ClientIaType) HasIanapd() bool {
	return obj.obj.Ianapd != nil
}

// description is TBD
// SetIanapd sets the DeviceDhcpv6ClientIaTimeValue value in the DeviceDhcpv6ClientIaType object
func (obj *deviceDhcpv6ClientIaType) SetIanapd(value DeviceDhcpv6ClientIaTimeValue) DeviceDhcpv6ClientIaType {
	obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IANAPD)
	obj.ianapdHolder = nil
	obj.obj.Ianapd = value.msg()

	return obj
}

func (obj *deviceDhcpv6ClientIaType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Iana != nil {

		obj.Iana().validateObj(vObj, set_default)
	}

	if obj.obj.Iapd != nil {

		obj.Iapd().validateObj(vObj, set_default)
	}

	if obj.obj.Ianapd != nil {

		obj.Ianapd().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv6ClientIaType) setDefault() {
	var choices_set int = 0
	var choice DeviceDhcpv6ClientIaTypeChoiceEnum

	if obj.obj.Iana != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientIaTypeChoice.IANA
	}

	if obj.obj.Iapd != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientIaTypeChoice.IAPD
	}

	if obj.obj.Ianapd != nil {
		choices_set += 1
		choice = DeviceDhcpv6ClientIaTypeChoice.IANAPD
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceDhcpv6ClientIaTypeChoice.IANA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceDhcpv6ClientIaType")
			}
		} else {
			intVal := otg.DeviceDhcpv6ClientIaType_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceDhcpv6ClientIaType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
