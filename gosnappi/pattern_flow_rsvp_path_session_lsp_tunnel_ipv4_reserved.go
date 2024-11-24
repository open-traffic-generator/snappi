package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionLspTunnelIpv4Reserved *****
type patternFlowRSVPPathSessionLspTunnelIpv4Reserved struct {
	validation
	obj             *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	marshaller      marshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	unMarshaller    unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	incrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	decrementHolder PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
}

func NewPatternFlowRSVPPathSessionLspTunnelIpv4Reserved() PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj := patternFlowRSVPPathSessionLspTunnelIpv4Reserved{obj: &otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) setMsg(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved
}

type marshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved interface {
	// ToProto marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error)
	// ToPbText marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved struct {
	obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved
}

type unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) (PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) ToProto() (*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FromProto(msg *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) (PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionLspTunnelIpv4Reserved) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionLspTunnelIpv4Reserved()
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

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSessionLspTunnelIpv4Reserved is reserved field, MUST be zero.
type PatternFlowRSVPPathSessionLspTunnelIpv4Reserved interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved to protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// setMsg unmarshals PatternFlowRSVPPathSessionLspTunnelIpv4Reserved from protobuf object *otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// validate validates PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionLspTunnelIpv4Reserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum, set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	Choice() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	SetValue(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// HasValue checks if Value has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	SetValues(value []uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// Increment returns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// SetIncrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter, set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
	// SetDecrement assigns PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter provided by user to PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSessionLspTunnelIpv4Reserved
var PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice = struct {
	VALUE     PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum
	VALUES    PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum
	INCREMENT PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum
	DECREMENT PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Choice() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum {
	return PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) setChoice(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	intValue, ok := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter().msg()
	}

	if value == PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) SetValue(value uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSessionLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) SetValues(value []uint32) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Increment() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) SetIncrement(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) Decrement() PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter value in the PatternFlowRSVPPathSessionLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) SetDecrement(value PatternFlowRSVPPathSessionLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSessionLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSessionLspTunnelIpv4Reserved.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSessionLspTunnelIpv4Reserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSessionLspTunnelIpv4ReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSessionLspTunnelIpv4Reserved")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSessionLspTunnelIpv4Reserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
