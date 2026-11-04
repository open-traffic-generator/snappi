package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp *****
type patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	marshaller      marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	unMarshaller    unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	incrementHolder PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	decrementHolder PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
}

func NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj := patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp{obj: &otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) setMsg(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
}

type marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp interface {
	// ToProto marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error)
	// ToPbText marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp struct {
	obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
}

type unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp interface {
	// FromProto unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) ToProto() (*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FromProto(msg *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp()
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp is a 32-bit value representing the time of packet processing, recorded as the number of milliseconds elapsed since midnight Universal Time (UT).
type PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp to protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// setMsg unmarshals PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp from protobuf object *otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// validate validates PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	Choice() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	setChoice(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// Increment returns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// SetIncrement assigns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter, set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
	// SetDecrement assigns PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter provided by user to PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp.
	// PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp
var PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice = struct {
	VALUE     PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum
	VALUES    PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum
	INCREMENT PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum
	DECREMENT PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Choice() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum {
	return PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) setChoice(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	intValue, ok := otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter().msg()
	}

	if value == PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) SetValue(value uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) SetValues(value []uint32) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Increment() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) SetIncrement(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) Decrement() PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter value in the PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp object
func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) SetDecrement(value PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampCounter) PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp {
	obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIpv4OptionsTimestampFormatTimestampsTimestamp) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsTimestampFormatTimestampsTimestampChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsTimestampFormatTimestampsTimestamp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
