package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQuerySFlag *****
type patternFlowIgmpv3MembershipQuerySFlag struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQuerySFlag
	marshaller      marshalPatternFlowIgmpv3MembershipQuerySFlag
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQuerySFlag
	incrementHolder PatternFlowIgmpv3MembershipQuerySFlagCounter
	decrementHolder PatternFlowIgmpv3MembershipQuerySFlagCounter
}

func NewPatternFlowIgmpv3MembershipQuerySFlag() PatternFlowIgmpv3MembershipQuerySFlag {
	obj := patternFlowIgmpv3MembershipQuerySFlag{obj: &otg.PatternFlowIgmpv3MembershipQuerySFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) msg() *otg.PatternFlowIgmpv3MembershipQuerySFlag {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) setMsg(msg *otg.PatternFlowIgmpv3MembershipQuerySFlag) PatternFlowIgmpv3MembershipQuerySFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQuerySFlag struct {
	obj *patternFlowIgmpv3MembershipQuerySFlag
}

type marshalPatternFlowIgmpv3MembershipQuerySFlag interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQuerySFlag to protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlag
	ToProto() (*otg.PatternFlowIgmpv3MembershipQuerySFlag, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQuerySFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQuerySFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQuerySFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQuerySFlag struct {
	obj *patternFlowIgmpv3MembershipQuerySFlag
}

type unMarshalPatternFlowIgmpv3MembershipQuerySFlag interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQuerySFlag from protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlag
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQuerySFlag) (PatternFlowIgmpv3MembershipQuerySFlag, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQuerySFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQuerySFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQuerySFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) Marshal() marshalPatternFlowIgmpv3MembershipQuerySFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQuerySFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQuerySFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQuerySFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlag) ToProto() (*otg.PatternFlowIgmpv3MembershipQuerySFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlag) FromProto(msg *otg.PatternFlowIgmpv3MembershipQuerySFlag) (PatternFlowIgmpv3MembershipQuerySFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQuerySFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQuerySFlag) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQuerySFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) Clone() (PatternFlowIgmpv3MembershipQuerySFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQuerySFlag()
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

func (obj *patternFlowIgmpv3MembershipQuerySFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQuerySFlag is suppress router-side processing flag, when set receiving routers do not update their timers.
type PatternFlowIgmpv3MembershipQuerySFlag interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQuerySFlag to protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQuerySFlag
	// setMsg unmarshals PatternFlowIgmpv3MembershipQuerySFlag from protobuf object *otg.PatternFlowIgmpv3MembershipQuerySFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQuerySFlag) PatternFlowIgmpv3MembershipQuerySFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQuerySFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQuerySFlag
	// validate validates PatternFlowIgmpv3MembershipQuerySFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQuerySFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum, set in PatternFlowIgmpv3MembershipQuerySFlag
	Choice() PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum provided by user to PatternFlowIgmpv3MembershipQuerySFlag
	setChoice(value PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum) PatternFlowIgmpv3MembershipQuerySFlag
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQuerySFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQuerySFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQuerySFlag
	SetValue(value uint32) PatternFlowIgmpv3MembershipQuerySFlag
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQuerySFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQuerySFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQuerySFlag
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQuerySFlag
	// Increment returns PatternFlowIgmpv3MembershipQuerySFlagCounter, set in PatternFlowIgmpv3MembershipQuerySFlag.
	// PatternFlowIgmpv3MembershipQuerySFlagCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQuerySFlagCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQuerySFlagCounter provided by user to PatternFlowIgmpv3MembershipQuerySFlag.
	// PatternFlowIgmpv3MembershipQuerySFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlag
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQuerySFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQuerySFlagCounter, set in PatternFlowIgmpv3MembershipQuerySFlag.
	// PatternFlowIgmpv3MembershipQuerySFlagCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQuerySFlagCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQuerySFlagCounter provided by user to PatternFlowIgmpv3MembershipQuerySFlag.
	// PatternFlowIgmpv3MembershipQuerySFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQuerySFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQuerySFlag
var PatternFlowIgmpv3MembershipQuerySFlagChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) Choice() PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum {
	return PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQuerySFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) setChoice(value PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum) PatternFlowIgmpv3MembershipQuerySFlag {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQuerySFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQuerySFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQuerySFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQuerySFlagCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQuerySFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQuerySFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQuerySFlag object
func (obj *patternFlowIgmpv3MembershipQuerySFlag) SetValue(value uint32) PatternFlowIgmpv3MembershipQuerySFlag {
	obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQuerySFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQuerySFlag object
func (obj *patternFlowIgmpv3MembershipQuerySFlag) SetValues(value []uint32) PatternFlowIgmpv3MembershipQuerySFlag {
	obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQuerySFlagCounter
func (obj *patternFlowIgmpv3MembershipQuerySFlag) Increment() PatternFlowIgmpv3MembershipQuerySFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQuerySFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQuerySFlagCounter
func (obj *patternFlowIgmpv3MembershipQuerySFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQuerySFlagCounter value in the PatternFlowIgmpv3MembershipQuerySFlag object
func (obj *patternFlowIgmpv3MembershipQuerySFlag) SetIncrement(value PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlag {
	obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQuerySFlagCounter
func (obj *patternFlowIgmpv3MembershipQuerySFlag) Decrement() PatternFlowIgmpv3MembershipQuerySFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQuerySFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQuerySFlagCounter
func (obj *patternFlowIgmpv3MembershipQuerySFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQuerySFlagCounter value in the PatternFlowIgmpv3MembershipQuerySFlag object
func (obj *patternFlowIgmpv3MembershipQuerySFlag) SetDecrement(value PatternFlowIgmpv3MembershipQuerySFlagCounter) PatternFlowIgmpv3MembershipQuerySFlag {
	obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQuerySFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQuerySFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQuerySFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQuerySFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQuerySFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQuerySFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQuerySFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQuerySFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQuerySFlag")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQuerySFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQuerySFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
