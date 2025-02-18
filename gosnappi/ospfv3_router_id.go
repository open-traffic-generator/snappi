package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3RouterId *****
type ospfv3RouterId struct {
	validation
	obj          *otg.Ospfv3RouterId
	marshaller   marshalOspfv3RouterId
	unMarshaller unMarshalOspfv3RouterId
}

func NewOspfv3RouterId() Ospfv3RouterId {
	obj := ospfv3RouterId{obj: &otg.Ospfv3RouterId{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3RouterId) msg() *otg.Ospfv3RouterId {
	return obj.obj
}

func (obj *ospfv3RouterId) setMsg(msg *otg.Ospfv3RouterId) Ospfv3RouterId {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3RouterId struct {
	obj *ospfv3RouterId
}

type marshalOspfv3RouterId interface {
	// ToProto marshals Ospfv3RouterId to protobuf object *otg.Ospfv3RouterId
	ToProto() (*otg.Ospfv3RouterId, error)
	// ToPbText marshals Ospfv3RouterId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3RouterId to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3RouterId to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3RouterId struct {
	obj *ospfv3RouterId
}

type unMarshalOspfv3RouterId interface {
	// FromProto unmarshals Ospfv3RouterId from protobuf object *otg.Ospfv3RouterId
	FromProto(msg *otg.Ospfv3RouterId) (Ospfv3RouterId, error)
	// FromPbText unmarshals Ospfv3RouterId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3RouterId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3RouterId from JSON text
	FromJson(value string) error
}

func (obj *ospfv3RouterId) Marshal() marshalOspfv3RouterId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3RouterId{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3RouterId) Unmarshal() unMarshalOspfv3RouterId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3RouterId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3RouterId) ToProto() (*otg.Ospfv3RouterId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3RouterId) FromProto(msg *otg.Ospfv3RouterId) (Ospfv3RouterId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3RouterId) ToPbText() (string, error) {
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

func (m *unMarshalospfv3RouterId) FromPbText(value string) error {
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

func (m *marshalospfv3RouterId) ToYaml() (string, error) {
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

func (m *unMarshalospfv3RouterId) FromYaml(value string) error {
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

func (m *marshalospfv3RouterId) ToJson() (string, error) {
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

func (m *unMarshalospfv3RouterId) FromJson(value string) error {
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

func (obj *ospfv3RouterId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3RouterId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3RouterId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3RouterId) Clone() (Ospfv3RouterId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3RouterId()
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

// Ospfv3RouterId is container for OSPFv3 Router ID configuration.
type Ospfv3RouterId interface {
	Validation
	// msg marshals Ospfv3RouterId to protobuf object *otg.Ospfv3RouterId
	// and doesn't set defaults
	msg() *otg.Ospfv3RouterId
	// setMsg unmarshals Ospfv3RouterId from protobuf object *otg.Ospfv3RouterId
	// and doesn't set defaults
	setMsg(*otg.Ospfv3RouterId) Ospfv3RouterId
	// provides marshal interface
	Marshal() marshalOspfv3RouterId
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3RouterId
	// validate validates Ospfv3RouterId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3RouterId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3RouterIdChoiceEnum, set in Ospfv3RouterId
	Choice() Ospfv3RouterIdChoiceEnum
	// setChoice assigns Ospfv3RouterIdChoiceEnum provided by user to Ospfv3RouterId
	setChoice(value Ospfv3RouterIdChoiceEnum) Ospfv3RouterId
	// HasChoice checks if Choice has been set in Ospfv3RouterId
	HasChoice() bool
	// getter for Auto to set choice.
	Auto()
	// Custom returns string, set in Ospfv3RouterId.
	Custom() string
	// SetCustom assigns string provided by user to Ospfv3RouterId
	SetCustom(value string) Ospfv3RouterId
	// HasCustom checks if Custom has been set in Ospfv3RouterId
	HasCustom() bool
}

type Ospfv3RouterIdChoiceEnum string

// Enum of Choice on Ospfv3RouterId
var Ospfv3RouterIdChoice = struct {
	AUTO   Ospfv3RouterIdChoiceEnum
	CUSTOM Ospfv3RouterIdChoiceEnum
}{
	AUTO:   Ospfv3RouterIdChoiceEnum("auto"),
	CUSTOM: Ospfv3RouterIdChoiceEnum("custom"),
}

func (obj *ospfv3RouterId) Choice() Ospfv3RouterIdChoiceEnum {
	return Ospfv3RouterIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Auto to set choice
func (obj *ospfv3RouterId) Auto() {
	obj.setChoice(Ospfv3RouterIdChoice.AUTO)
}

// IP address of Router ID for this emulated OSPFv3 router.
// - auto: When first IPv4 address of the router is attempted to be assigned as Router ID.
// - custom: When, Router ID needs to be configured different from first IPv4 address of the router.
// Choice returns a string
func (obj *ospfv3RouterId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3RouterId) setChoice(value Ospfv3RouterIdChoiceEnum) Ospfv3RouterId {
	intValue, ok := otg.Ospfv3RouterId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3RouterIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3RouterId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	return obj
}

// Router ID in IPv4 address format.
// Custom returns a string
func (obj *ospfv3RouterId) Custom() string {

	if obj.obj.Custom == nil {
		obj.setChoice(Ospfv3RouterIdChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// Router ID in IPv4 address format.
// Custom returns a string
func (obj *ospfv3RouterId) HasCustom() bool {
	return obj.obj.Custom != nil
}

// Router ID in IPv4 address format.
// SetCustom sets the string value in the Ospfv3RouterId object
func (obj *ospfv3RouterId) SetCustom(value string) Ospfv3RouterId {
	obj.setChoice(Ospfv3RouterIdChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *ospfv3RouterId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		err := obj.validateIpv4(obj.Custom())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3RouterId.Custom"))
		}

	}

}

func (obj *ospfv3RouterId) setDefault() {
	var choices_set int = 0
	var choice Ospfv3RouterIdChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = Ospfv3RouterIdChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3RouterIdChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3RouterId")
			}
		} else {
			intVal := otg.Ospfv3RouterId_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3RouterId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
