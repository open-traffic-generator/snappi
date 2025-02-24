package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameter127Flag *****
type patternFlowRSVPPathSenderTspecIntServParameter127Flag struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag
	incrementHolder PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameter127Flag() PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj := patternFlowRSVPPathSenderTspecIntServParameter127Flag{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag) (PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag) (PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Flag) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameter127Flag()
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServParameter127Flag is parameter 127 flags (none set)
type PatternFlowRSVPPathSenderTspecIntServParameter127Flag interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Flag from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127Flag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	Choice() PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	setChoice(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// Increment returns PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	// PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	// PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter, set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	// PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Flag.
	// PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Flag
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServParameter127Flag
var PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Choice() PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) setChoice(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127Flag object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127Flag object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Increment() PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter value in the PatternFlowRSVPPathSenderTspecIntServParameter127Flag object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) Decrement() PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServParameter127FlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter value in the PatternFlowRSVPPathSenderTspecIntServParameter127Flag object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameter127FlagCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Flag {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127Flag.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServParameter127Flag.Values <= 255 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Flag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127FlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServParameter127Flag")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Flag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
