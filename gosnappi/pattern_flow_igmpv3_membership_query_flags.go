package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryFlags *****
type patternFlowIgmpv3MembershipQueryFlags struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryFlags
	marshaller      marshalPatternFlowIgmpv3MembershipQueryFlags
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryFlags
	incrementHolder PatternFlowIgmpv3MembershipQueryFlagsCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryFlagsCounter
}

func NewPatternFlowIgmpv3MembershipQueryFlags() PatternFlowIgmpv3MembershipQueryFlags {
	obj := patternFlowIgmpv3MembershipQueryFlags{obj: &otg.PatternFlowIgmpv3MembershipQueryFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) msg() *otg.PatternFlowIgmpv3MembershipQueryFlags {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryFlags) PatternFlowIgmpv3MembershipQueryFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryFlags struct {
	obj *patternFlowIgmpv3MembershipQueryFlags
}

type marshalPatternFlowIgmpv3MembershipQueryFlags interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryFlags to protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlags
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryFlags, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryFlags struct {
	obj *patternFlowIgmpv3MembershipQueryFlags
}

type unMarshalPatternFlowIgmpv3MembershipQueryFlags interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryFlags from protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlags
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryFlags) (PatternFlowIgmpv3MembershipQueryFlags, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) Marshal() marshalPatternFlowIgmpv3MembershipQueryFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryFlags) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlags) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryFlags) (PatternFlowIgmpv3MembershipQueryFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryFlags) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) Clone() (PatternFlowIgmpv3MembershipQueryFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryFlags()
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

func (obj *patternFlowIgmpv3MembershipQueryFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryFlags is flags field of the query, 4 bits that are managed by IANA and are set to zero by a conforming sender.
type PatternFlowIgmpv3MembershipQueryFlags interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryFlags to protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryFlags
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryFlags from protobuf object *otg.PatternFlowIgmpv3MembershipQueryFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryFlags) PatternFlowIgmpv3MembershipQueryFlags
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryFlags
	// validate validates PatternFlowIgmpv3MembershipQueryFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum, set in PatternFlowIgmpv3MembershipQueryFlags
	Choice() PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryFlags
	setChoice(value PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum) PatternFlowIgmpv3MembershipQueryFlags
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryFlags
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryFlags
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryFlags
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryFlags
	// Increment returns PatternFlowIgmpv3MembershipQueryFlagsCounter, set in PatternFlowIgmpv3MembershipQueryFlags.
	// PatternFlowIgmpv3MembershipQueryFlagsCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryFlagsCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryFlagsCounter provided by user to PatternFlowIgmpv3MembershipQueryFlags.
	// PatternFlowIgmpv3MembershipQueryFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlags
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryFlags
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryFlagsCounter, set in PatternFlowIgmpv3MembershipQueryFlags.
	// PatternFlowIgmpv3MembershipQueryFlagsCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryFlagsCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryFlagsCounter provided by user to PatternFlowIgmpv3MembershipQueryFlags.
	// PatternFlowIgmpv3MembershipQueryFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlags
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryFlags
var PatternFlowIgmpv3MembershipQueryFlagsChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) Choice() PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) setChoice(value PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum) PatternFlowIgmpv3MembershipQueryFlags {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryFlagsCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryFlags object
func (obj *patternFlowIgmpv3MembershipQueryFlags) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryFlags object
func (obj *patternFlowIgmpv3MembershipQueryFlags) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryFlagsCounter
func (obj *patternFlowIgmpv3MembershipQueryFlags) Increment() PatternFlowIgmpv3MembershipQueryFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryFlagsCounter
func (obj *patternFlowIgmpv3MembershipQueryFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryFlagsCounter value in the PatternFlowIgmpv3MembershipQueryFlags object
func (obj *patternFlowIgmpv3MembershipQueryFlags) SetIncrement(value PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryFlagsCounter
func (obj *patternFlowIgmpv3MembershipQueryFlags) Decrement() PatternFlowIgmpv3MembershipQueryFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryFlagsCounter
func (obj *patternFlowIgmpv3MembershipQueryFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryFlagsCounter value in the PatternFlowIgmpv3MembershipQueryFlags object
func (obj *patternFlowIgmpv3MembershipQueryFlags) SetDecrement(value PatternFlowIgmpv3MembershipQueryFlagsCounter) PatternFlowIgmpv3MembershipQueryFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryFlags.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryFlags.Values <= 15 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQueryFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryFlags")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
