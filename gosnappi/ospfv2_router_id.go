package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2RouterId *****
type ospfv2RouterId struct {
	validation
	obj          *otg.Ospfv2RouterId
	marshaller   marshalOspfv2RouterId
	unMarshaller unMarshalOspfv2RouterId
}

func NewOspfv2RouterId() Ospfv2RouterId {
	obj := ospfv2RouterId{obj: &otg.Ospfv2RouterId{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2RouterId) msg() *otg.Ospfv2RouterId {
	return obj.obj
}

func (obj *ospfv2RouterId) setMsg(msg *otg.Ospfv2RouterId) Ospfv2RouterId {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2RouterId struct {
	obj *ospfv2RouterId
}

type marshalOspfv2RouterId interface {
	// ToProto marshals Ospfv2RouterId to protobuf object *otg.Ospfv2RouterId
	ToProto() (*otg.Ospfv2RouterId, error)
	// ToPbText marshals Ospfv2RouterId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2RouterId to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2RouterId to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2RouterId struct {
	obj *ospfv2RouterId
}

type unMarshalOspfv2RouterId interface {
	// FromProto unmarshals Ospfv2RouterId from protobuf object *otg.Ospfv2RouterId
	FromProto(msg *otg.Ospfv2RouterId) (Ospfv2RouterId, error)
	// FromPbText unmarshals Ospfv2RouterId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2RouterId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2RouterId from JSON text
	FromJson(value string) error
}

func (obj *ospfv2RouterId) Marshal() marshalOspfv2RouterId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2RouterId{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2RouterId) Unmarshal() unMarshalOspfv2RouterId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2RouterId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2RouterId) ToProto() (*otg.Ospfv2RouterId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2RouterId) FromProto(msg *otg.Ospfv2RouterId) (Ospfv2RouterId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2RouterId) ToPbText() (string, error) {
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

func (m *unMarshalospfv2RouterId) FromPbText(value string) error {
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

func (m *marshalospfv2RouterId) ToYaml() (string, error) {
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

func (m *unMarshalospfv2RouterId) FromYaml(value string) error {
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

func (m *marshalospfv2RouterId) ToJson() (string, error) {
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

func (m *unMarshalospfv2RouterId) FromJson(value string) error {
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

func (obj *ospfv2RouterId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2RouterId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2RouterId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2RouterId) Clone() (Ospfv2RouterId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2RouterId()
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

// Ospfv2RouterId is containter for OSPFv2 Router ID configuration.
type Ospfv2RouterId interface {
	Validation
	// msg marshals Ospfv2RouterId to protobuf object *otg.Ospfv2RouterId
	// and doesn't set defaults
	msg() *otg.Ospfv2RouterId
	// setMsg unmarshals Ospfv2RouterId from protobuf object *otg.Ospfv2RouterId
	// and doesn't set defaults
	setMsg(*otg.Ospfv2RouterId) Ospfv2RouterId
	// provides marshal interface
	Marshal() marshalOspfv2RouterId
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2RouterId
	// validate validates Ospfv2RouterId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2RouterId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2RouterIdChoiceEnum, set in Ospfv2RouterId
	Choice() Ospfv2RouterIdChoiceEnum
	// setChoice assigns Ospfv2RouterIdChoiceEnum provided by user to Ospfv2RouterId
	setChoice(value Ospfv2RouterIdChoiceEnum) Ospfv2RouterId
	// HasChoice checks if Choice has been set in Ospfv2RouterId
	HasChoice() bool
	// getter for InterfaceIp to set choice.
	InterfaceIp()
	// CustomRouterId returns string, set in Ospfv2RouterId.
	CustomRouterId() string
	// SetCustomRouterId assigns string provided by user to Ospfv2RouterId
	SetCustomRouterId(value string) Ospfv2RouterId
	// HasCustomRouterId checks if CustomRouterId has been set in Ospfv2RouterId
	HasCustomRouterId() bool
}

type Ospfv2RouterIdChoiceEnum string

// Enum of Choice on Ospfv2RouterId
var Ospfv2RouterIdChoice = struct {
	INTERFACE_IP     Ospfv2RouterIdChoiceEnum
	CUSTOM_ROUTER_ID Ospfv2RouterIdChoiceEnum
}{
	INTERFACE_IP:     Ospfv2RouterIdChoiceEnum("interface_ip"),
	CUSTOM_ROUTER_ID: Ospfv2RouterIdChoiceEnum("custom_router_id"),
}

func (obj *ospfv2RouterId) Choice() Ospfv2RouterIdChoiceEnum {
	return Ospfv2RouterIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for InterfaceIp to set choice
func (obj *ospfv2RouterId) InterfaceIp() {
	obj.setChoice(Ospfv2RouterIdChoice.INTERFACE_IP)
}

// IP address of Router ID for this emulated OSPFv2 router.
// - interface_ip: When IPv4 interface address to be assigned as Router ID.
// - custom_router_id: When, Router ID needs to be configured different from Interface IPv4 address.
// Choice returns a string
func (obj *ospfv2RouterId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2RouterId) setChoice(value Ospfv2RouterIdChoiceEnum) Ospfv2RouterId {
	intValue, ok := otg.Ospfv2RouterId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2RouterIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2RouterId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CustomRouterId = nil
	return obj
}

// Router ID in IPv4 address format.
// CustomRouterId returns a string
func (obj *ospfv2RouterId) CustomRouterId() string {

	if obj.obj.CustomRouterId == nil {
		obj.setChoice(Ospfv2RouterIdChoice.CUSTOM_ROUTER_ID)
	}

	return *obj.obj.CustomRouterId

}

// Router ID in IPv4 address format.
// CustomRouterId returns a string
func (obj *ospfv2RouterId) HasCustomRouterId() bool {
	return obj.obj.CustomRouterId != nil
}

// Router ID in IPv4 address format.
// SetCustomRouterId sets the string value in the Ospfv2RouterId object
func (obj *ospfv2RouterId) SetCustomRouterId(value string) Ospfv2RouterId {
	obj.setChoice(Ospfv2RouterIdChoice.CUSTOM_ROUTER_ID)
	obj.obj.CustomRouterId = &value
	return obj
}

func (obj *ospfv2RouterId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomRouterId != nil {

		err := obj.validateIpv4(obj.CustomRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2RouterId.CustomRouterId"))
		}

	}

}

func (obj *ospfv2RouterId) setDefault() {
	var choices_set int = 0
	var choice Ospfv2RouterIdChoiceEnum

	if obj.obj.CustomRouterId != nil {
		choices_set += 1
		choice = Ospfv2RouterIdChoice.CUSTOM_ROUTER_ID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2RouterIdChoice.INTERFACE_IP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2RouterId")
			}
		} else {
			intVal := otg.Ospfv2RouterId_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2RouterId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
