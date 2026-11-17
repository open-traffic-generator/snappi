package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportFlags *****
type patternFlowIgmpv3MembershipReportFlags struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipReportFlags
	marshaller      marshalPatternFlowIgmpv3MembershipReportFlags
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipReportFlags
	incrementHolder PatternFlowIgmpv3MembershipReportFlagsCounter
	decrementHolder PatternFlowIgmpv3MembershipReportFlagsCounter
}

func NewPatternFlowIgmpv3MembershipReportFlags() PatternFlowIgmpv3MembershipReportFlags {
	obj := patternFlowIgmpv3MembershipReportFlags{obj: &otg.PatternFlowIgmpv3MembershipReportFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportFlags) msg() *otg.PatternFlowIgmpv3MembershipReportFlags {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportFlags) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportFlags) PatternFlowIgmpv3MembershipReportFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportFlags struct {
	obj *patternFlowIgmpv3MembershipReportFlags
}

type marshalPatternFlowIgmpv3MembershipReportFlags interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportFlags to protobuf object *otg.PatternFlowIgmpv3MembershipReportFlags
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportFlags, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportFlags struct {
	obj *patternFlowIgmpv3MembershipReportFlags
}

type unMarshalPatternFlowIgmpv3MembershipReportFlags interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportFlags from protobuf object *otg.PatternFlowIgmpv3MembershipReportFlags
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportFlags) (PatternFlowIgmpv3MembershipReportFlags, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportFlags) Marshal() marshalPatternFlowIgmpv3MembershipReportFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportFlags) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportFlags) ToProto() (*otg.PatternFlowIgmpv3MembershipReportFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlags) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportFlags) (PatternFlowIgmpv3MembershipReportFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportFlags) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportFlags) Clone() (PatternFlowIgmpv3MembershipReportFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportFlags()
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

func (obj *patternFlowIgmpv3MembershipReportFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipReportFlags is flags field of the report, 16 bits that are managed by IANA and are set to zero by a conforming sender.
type PatternFlowIgmpv3MembershipReportFlags interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportFlags to protobuf object *otg.PatternFlowIgmpv3MembershipReportFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportFlags
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportFlags from protobuf object *otg.PatternFlowIgmpv3MembershipReportFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportFlags) PatternFlowIgmpv3MembershipReportFlags
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportFlags
	// validate validates PatternFlowIgmpv3MembershipReportFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipReportFlagsChoiceEnum, set in PatternFlowIgmpv3MembershipReportFlags
	Choice() PatternFlowIgmpv3MembershipReportFlagsChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipReportFlagsChoiceEnum provided by user to PatternFlowIgmpv3MembershipReportFlags
	setChoice(value PatternFlowIgmpv3MembershipReportFlagsChoiceEnum) PatternFlowIgmpv3MembershipReportFlags
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipReportFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipReportFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportFlags
	SetValue(value uint32) PatternFlowIgmpv3MembershipReportFlags
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipReportFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipReportFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipReportFlags
	SetValues(value []uint32) PatternFlowIgmpv3MembershipReportFlags
	// Increment returns PatternFlowIgmpv3MembershipReportFlagsCounter, set in PatternFlowIgmpv3MembershipReportFlags.
	// PatternFlowIgmpv3MembershipReportFlagsCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipReportFlagsCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipReportFlagsCounter provided by user to PatternFlowIgmpv3MembershipReportFlags.
	// PatternFlowIgmpv3MembershipReportFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlags
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipReportFlags
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipReportFlagsCounter, set in PatternFlowIgmpv3MembershipReportFlags.
	// PatternFlowIgmpv3MembershipReportFlagsCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipReportFlagsCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipReportFlagsCounter provided by user to PatternFlowIgmpv3MembershipReportFlags.
	// PatternFlowIgmpv3MembershipReportFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlags
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipReportFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipReportFlagsChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipReportFlags
var PatternFlowIgmpv3MembershipReportFlagsChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipReportFlagsChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipReportFlagsChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipReportFlagsChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipReportFlagsChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipReportFlagsChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipReportFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipReportFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipReportFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipReportFlags) Choice() PatternFlowIgmpv3MembershipReportFlagsChoiceEnum {
	return PatternFlowIgmpv3MembershipReportFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipReportFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipReportFlags) setChoice(value PatternFlowIgmpv3MembershipReportFlagsChoiceEnum) PatternFlowIgmpv3MembershipReportFlags {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipReportFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipReportFlagsCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipReportFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipReportFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipReportFlags object
func (obj *patternFlowIgmpv3MembershipReportFlags) SetValue(value uint32) PatternFlowIgmpv3MembershipReportFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipReportFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipReportFlags object
func (obj *patternFlowIgmpv3MembershipReportFlags) SetValues(value []uint32) PatternFlowIgmpv3MembershipReportFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportFlagsCounter
func (obj *patternFlowIgmpv3MembershipReportFlags) Increment() PatternFlowIgmpv3MembershipReportFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipReportFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportFlagsCounter
func (obj *patternFlowIgmpv3MembershipReportFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipReportFlagsCounter value in the PatternFlowIgmpv3MembershipReportFlags object
func (obj *patternFlowIgmpv3MembershipReportFlags) SetIncrement(value PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportFlagsCounter
func (obj *patternFlowIgmpv3MembershipReportFlags) Decrement() PatternFlowIgmpv3MembershipReportFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipReportFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportFlagsCounter
func (obj *patternFlowIgmpv3MembershipReportFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipReportFlagsCounter value in the PatternFlowIgmpv3MembershipReportFlags object
func (obj *patternFlowIgmpv3MembershipReportFlags) SetDecrement(value PatternFlowIgmpv3MembershipReportFlagsCounter) PatternFlowIgmpv3MembershipReportFlags {
	obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipReportFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportFlags.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipReportFlags.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipReportFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipReportFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipReportFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipReportFlags")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipReportFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipReportFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
