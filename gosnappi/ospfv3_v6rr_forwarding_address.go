package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3V6RRForwardingAddress *****
type ospfv3V6RRForwardingAddress struct {
	validation
	obj          *otg.Ospfv3V6RRForwardingAddress
	marshaller   marshalOspfv3V6RRForwardingAddress
	unMarshaller unMarshalOspfv3V6RRForwardingAddress
}

func NewOspfv3V6RRForwardingAddress() Ospfv3V6RRForwardingAddress {
	obj := ospfv3V6RRForwardingAddress{obj: &otg.Ospfv3V6RRForwardingAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3V6RRForwardingAddress) msg() *otg.Ospfv3V6RRForwardingAddress {
	return obj.obj
}

func (obj *ospfv3V6RRForwardingAddress) setMsg(msg *otg.Ospfv3V6RRForwardingAddress) Ospfv3V6RRForwardingAddress {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3V6RRForwardingAddress struct {
	obj *ospfv3V6RRForwardingAddress
}

type marshalOspfv3V6RRForwardingAddress interface {
	// ToProto marshals Ospfv3V6RRForwardingAddress to protobuf object *otg.Ospfv3V6RRForwardingAddress
	ToProto() (*otg.Ospfv3V6RRForwardingAddress, error)
	// ToPbText marshals Ospfv3V6RRForwardingAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3V6RRForwardingAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3V6RRForwardingAddress to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3V6RRForwardingAddress struct {
	obj *ospfv3V6RRForwardingAddress
}

type unMarshalOspfv3V6RRForwardingAddress interface {
	// FromProto unmarshals Ospfv3V6RRForwardingAddress from protobuf object *otg.Ospfv3V6RRForwardingAddress
	FromProto(msg *otg.Ospfv3V6RRForwardingAddress) (Ospfv3V6RRForwardingAddress, error)
	// FromPbText unmarshals Ospfv3V6RRForwardingAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3V6RRForwardingAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3V6RRForwardingAddress from JSON text
	FromJson(value string) error
}

func (obj *ospfv3V6RRForwardingAddress) Marshal() marshalOspfv3V6RRForwardingAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3V6RRForwardingAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3V6RRForwardingAddress) Unmarshal() unMarshalOspfv3V6RRForwardingAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3V6RRForwardingAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3V6RRForwardingAddress) ToProto() (*otg.Ospfv3V6RRForwardingAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3V6RRForwardingAddress) FromProto(msg *otg.Ospfv3V6RRForwardingAddress) (Ospfv3V6RRForwardingAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3V6RRForwardingAddress) ToPbText() (string, error) {
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

func (m *unMarshalospfv3V6RRForwardingAddress) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv3V6RRForwardingAddress) ToYaml() (string, error) {
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

func (m *unMarshalospfv3V6RRForwardingAddress) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv3V6RRForwardingAddress) ToJson() (string, error) {
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

func (m *unMarshalospfv3V6RRForwardingAddress) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv3V6RRForwardingAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3V6RRForwardingAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3V6RRForwardingAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3V6RRForwardingAddress) Clone() (Ospfv3V6RRForwardingAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3V6RRForwardingAddress()
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

// Ospfv3V6RRForwardingAddress is container for the forwarding address of NSSA External route origin configuration.
type Ospfv3V6RRForwardingAddress interface {
	Validation
	// msg marshals Ospfv3V6RRForwardingAddress to protobuf object *otg.Ospfv3V6RRForwardingAddress
	// and doesn't set defaults
	msg() *otg.Ospfv3V6RRForwardingAddress
	// setMsg unmarshals Ospfv3V6RRForwardingAddress from protobuf object *otg.Ospfv3V6RRForwardingAddress
	// and doesn't set defaults
	setMsg(*otg.Ospfv3V6RRForwardingAddress) Ospfv3V6RRForwardingAddress
	// provides marshal interface
	Marshal() marshalOspfv3V6RRForwardingAddress
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3V6RRForwardingAddress
	// validate validates Ospfv3V6RRForwardingAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3V6RRForwardingAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3V6RRForwardingAddressChoiceEnum, set in Ospfv3V6RRForwardingAddress
	Choice() Ospfv3V6RRForwardingAddressChoiceEnum
	// setChoice assigns Ospfv3V6RRForwardingAddressChoiceEnum provided by user to Ospfv3V6RRForwardingAddress
	setChoice(value Ospfv3V6RRForwardingAddressChoiceEnum) Ospfv3V6RRForwardingAddress
	// HasChoice checks if Choice has been set in Ospfv3V6RRForwardingAddress
	HasChoice() bool
	// getter for InterfaceIp to set choice.
	InterfaceIp()
	// Custom returns string, set in Ospfv3V6RRForwardingAddress.
	Custom() string
	// SetCustom assigns string provided by user to Ospfv3V6RRForwardingAddress
	SetCustom(value string) Ospfv3V6RRForwardingAddress
	// HasCustom checks if Custom has been set in Ospfv3V6RRForwardingAddress
	HasCustom() bool
}

type Ospfv3V6RRForwardingAddressChoiceEnum string

// Enum of Choice on Ospfv3V6RRForwardingAddress
var Ospfv3V6RRForwardingAddressChoice = struct {
	INTERFACE_IP Ospfv3V6RRForwardingAddressChoiceEnum
	CUSTOM       Ospfv3V6RRForwardingAddressChoiceEnum
}{
	INTERFACE_IP: Ospfv3V6RRForwardingAddressChoiceEnum("interface_ip"),
	CUSTOM:       Ospfv3V6RRForwardingAddressChoiceEnum("custom"),
}

func (obj *ospfv3V6RRForwardingAddress) Choice() Ospfv3V6RRForwardingAddressChoiceEnum {
	return Ospfv3V6RRForwardingAddressChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InterfaceIp to set choice
func (obj *ospfv3V6RRForwardingAddress) InterfaceIp() {
	obj.setChoice(Ospfv3V6RRForwardingAddressChoice.INTERFACE_IP)
}

// IPv6 forwarding address of Type 7 LSA Not-So-Stubby Area (NSSA) External.
// - interface_ip: if set, forwarding address is set with Interface IP6 address.
// - custom: if set, forwarding address is set with a custom IPv6 address.
// Choice returns a string
func (obj *ospfv3V6RRForwardingAddress) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3V6RRForwardingAddress) setChoice(value Ospfv3V6RRForwardingAddressChoiceEnum) Ospfv3V6RRForwardingAddress {
	intValue, ok := otg.Ospfv3V6RRForwardingAddress_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3V6RRForwardingAddressChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3V6RRForwardingAddress_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	return obj
}

// Forwarding address in IPv6 format.
// Custom returns a string
func (obj *ospfv3V6RRForwardingAddress) Custom() string {

	if obj.obj.Custom == nil {
		obj.setChoice(Ospfv3V6RRForwardingAddressChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// Forwarding address in IPv6 format.
// Custom returns a string
func (obj *ospfv3V6RRForwardingAddress) HasCustom() bool {
	return obj.obj.Custom != nil
}

// Forwarding address in IPv6 format.
// SetCustom sets the string value in the Ospfv3V6RRForwardingAddress object
func (obj *ospfv3V6RRForwardingAddress) SetCustom(value string) Ospfv3V6RRForwardingAddress {
	obj.setChoice(Ospfv3V6RRForwardingAddressChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *ospfv3V6RRForwardingAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		err := obj.validateIpv6(obj.Custom())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3V6RRForwardingAddress.Custom"))
		}

	}

}

func (obj *ospfv3V6RRForwardingAddress) setDefault() {
	var choices_set int = 0
	var choice Ospfv3V6RRForwardingAddressChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = Ospfv3V6RRForwardingAddressChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3V6RRForwardingAddressChoice.INTERFACE_IP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3V6RRForwardingAddress")
			}
		} else {
			intVal := otg.Ospfv3V6RRForwardingAddress_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3V6RRForwardingAddress_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
