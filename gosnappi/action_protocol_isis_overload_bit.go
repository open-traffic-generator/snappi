package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisOverloadBit *****
type actionProtocolIsisOverloadBit struct {
	validation
	obj          *otg.ActionProtocolIsisOverloadBit
	marshaller   marshalActionProtocolIsisOverloadBit
	unMarshaller unMarshalActionProtocolIsisOverloadBit
}

func NewActionProtocolIsisOverloadBit() ActionProtocolIsisOverloadBit {
	obj := actionProtocolIsisOverloadBit{obj: &otg.ActionProtocolIsisOverloadBit{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisOverloadBit) msg() *otg.ActionProtocolIsisOverloadBit {
	return obj.obj
}

func (obj *actionProtocolIsisOverloadBit) setMsg(msg *otg.ActionProtocolIsisOverloadBit) ActionProtocolIsisOverloadBit {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisOverloadBit struct {
	obj *actionProtocolIsisOverloadBit
}

type marshalActionProtocolIsisOverloadBit interface {
	// ToProto marshals ActionProtocolIsisOverloadBit to protobuf object *otg.ActionProtocolIsisOverloadBit
	ToProto() (*otg.ActionProtocolIsisOverloadBit, error)
	// ToPbText marshals ActionProtocolIsisOverloadBit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisOverloadBit to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisOverloadBit to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisOverloadBit struct {
	obj *actionProtocolIsisOverloadBit
}

type unMarshalActionProtocolIsisOverloadBit interface {
	// FromProto unmarshals ActionProtocolIsisOverloadBit from protobuf object *otg.ActionProtocolIsisOverloadBit
	FromProto(msg *otg.ActionProtocolIsisOverloadBit) (ActionProtocolIsisOverloadBit, error)
	// FromPbText unmarshals ActionProtocolIsisOverloadBit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisOverloadBit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisOverloadBit from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisOverloadBit) Marshal() marshalActionProtocolIsisOverloadBit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisOverloadBit{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisOverloadBit) Unmarshal() unMarshalActionProtocolIsisOverloadBit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisOverloadBit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisOverloadBit) ToProto() (*otg.ActionProtocolIsisOverloadBit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisOverloadBit) FromProto(msg *otg.ActionProtocolIsisOverloadBit) (ActionProtocolIsisOverloadBit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisOverloadBit) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisOverloadBit) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisOverloadBit) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisOverloadBit) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisOverloadBit) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisOverloadBit) FromJson(value string) error {
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

func (obj *actionProtocolIsisOverloadBit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisOverloadBit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisOverloadBit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisOverloadBit) Clone() (ActionProtocolIsisOverloadBit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisOverloadBit()
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

// ActionProtocolIsisOverloadBit is sets or clears the LSP Database Overload (OL) bit (Bit 3 of the
// P/ATT/LSPDBOL/IS-Type octet, ISO/IEC 10589:2002 Section 9.8) in
// LSP Number zero of the selected routers without restarting the IS-IS
// session (Section 7.2.5).
// When set, other routers exclude this router from transit IS SPF
// calculations; End System reachability through it is preserved
// (Section 7.2.8.1, RFC 3787 Section 4).
// The router re-originates LSP Number zero with an incremented sequence
// number; IS-IS floods it domain-wide so all routers recompute SPF.
// Typical uses: planned maintenance; simulating an overloaded router.
type ActionProtocolIsisOverloadBit interface {
	Validation
	// msg marshals ActionProtocolIsisOverloadBit to protobuf object *otg.ActionProtocolIsisOverloadBit
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisOverloadBit
	// setMsg unmarshals ActionProtocolIsisOverloadBit from protobuf object *otg.ActionProtocolIsisOverloadBit
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisOverloadBit) ActionProtocolIsisOverloadBit
	// provides marshal interface
	Marshal() marshalActionProtocolIsisOverloadBit
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisOverloadBit
	// validate validates ActionProtocolIsisOverloadBit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisOverloadBit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolIsisOverloadBit.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolIsisOverloadBit
	SetRouterNames(value []string) ActionProtocolIsisOverloadBit
	// Choice returns ActionProtocolIsisOverloadBitChoiceEnum, set in ActionProtocolIsisOverloadBit
	Choice() ActionProtocolIsisOverloadBitChoiceEnum
	// setChoice assigns ActionProtocolIsisOverloadBitChoiceEnum provided by user to ActionProtocolIsisOverloadBit
	setChoice(value ActionProtocolIsisOverloadBitChoiceEnum) ActionProtocolIsisOverloadBit
	// HasChoice checks if Choice has been set in ActionProtocolIsisOverloadBit
	HasChoice() bool
	// getter for Unset to set choice.
	Unset()
	// getter for Set to set choice.
	Set()
}

// The names of IS-IS routers on which to apply the overload bit state. If no names are specified then the action is applied to all configured IS-IS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolIsisOverloadBit) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of IS-IS routers on which to apply the overload bit state. If no names are specified then the action is applied to all configured IS-IS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the ActionProtocolIsisOverloadBit object
func (obj *actionProtocolIsisOverloadBit) SetRouterNames(value []string) ActionProtocolIsisOverloadBit {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type ActionProtocolIsisOverloadBitChoiceEnum string

// Enum of Choice on ActionProtocolIsisOverloadBit
var ActionProtocolIsisOverloadBitChoice = struct {
	SET   ActionProtocolIsisOverloadBitChoiceEnum
	UNSET ActionProtocolIsisOverloadBitChoiceEnum
}{
	SET:   ActionProtocolIsisOverloadBitChoiceEnum("set"),
	UNSET: ActionProtocolIsisOverloadBitChoiceEnum("unset"),
}

func (obj *actionProtocolIsisOverloadBit) Choice() ActionProtocolIsisOverloadBitChoiceEnum {
	return ActionProtocolIsisOverloadBitChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Unset to set choice
func (obj *actionProtocolIsisOverloadBit) Unset() {
	obj.setChoice(ActionProtocolIsisOverloadBitChoice.UNSET)
}

// getter for Set to set choice
func (obj *actionProtocolIsisOverloadBit) Set() {
	obj.setChoice(ActionProtocolIsisOverloadBitChoice.SET)
}

// set: The router sets the OL bit in its LSP Number zero. Other routers
// exclude this router from transit SPF path calculations for Intermediate
// System neighbours. End System reachability through this router is
// preserved.
// unset: The router clears the OL bit from its LSP Number zero. Normal
// transit path calculations resume for all routers.
// Choice returns a string
func (obj *actionProtocolIsisOverloadBit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *actionProtocolIsisOverloadBit) setChoice(value ActionProtocolIsisOverloadBitChoiceEnum) ActionProtocolIsisOverloadBit {
	intValue, ok := otg.ActionProtocolIsisOverloadBit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisOverloadBitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsisOverloadBit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *actionProtocolIsisOverloadBit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *actionProtocolIsisOverloadBit) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisOverloadBitChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(ActionProtocolIsisOverloadBitChoice.SET)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIsisOverloadBit")
			}
		} else {
			intVal := otg.ActionProtocolIsisOverloadBit_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIsisOverloadBit_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
