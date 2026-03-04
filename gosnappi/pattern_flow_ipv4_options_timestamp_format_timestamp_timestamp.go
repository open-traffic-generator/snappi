package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp *****
type patternFlowIpv4OptionsTimestampFormatTimestampTimestamp struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	incrementHolder PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp() PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj := patternFlowIpv4OptionsTimestampFormatTimestampTimestamp{obj: &otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp
}

type marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) (PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) (PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp()
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
type PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// validate validates PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	Choice() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	setChoice(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// Increment returns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp
var PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Choice() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) setChoice(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Increment() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) Decrement() PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatTimestampTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampTimestamp) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampTimestampChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatTimestampTimestamp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
