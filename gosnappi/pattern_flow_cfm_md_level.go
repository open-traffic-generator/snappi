package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmMdLevel *****
type patternFlowCfmMdLevel struct {
	validation
	obj             *otg.PatternFlowCfmMdLevel
	marshaller      marshalPatternFlowCfmMdLevel
	unMarshaller    unMarshalPatternFlowCfmMdLevel
	incrementHolder PatternFlowCfmMdLevelCounter
	decrementHolder PatternFlowCfmMdLevelCounter
}

func NewPatternFlowCfmMdLevel() PatternFlowCfmMdLevel {
	obj := patternFlowCfmMdLevel{obj: &otg.PatternFlowCfmMdLevel{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmMdLevel) msg() *otg.PatternFlowCfmMdLevel {
	return obj.obj
}

func (obj *patternFlowCfmMdLevel) setMsg(msg *otg.PatternFlowCfmMdLevel) PatternFlowCfmMdLevel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmMdLevel struct {
	obj *patternFlowCfmMdLevel
}

type marshalPatternFlowCfmMdLevel interface {
	// ToProto marshals PatternFlowCfmMdLevel to protobuf object *otg.PatternFlowCfmMdLevel
	ToProto() (*otg.PatternFlowCfmMdLevel, error)
	// ToPbText marshals PatternFlowCfmMdLevel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmMdLevel to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmMdLevel to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmMdLevel struct {
	obj *patternFlowCfmMdLevel
}

type unMarshalPatternFlowCfmMdLevel interface {
	// FromProto unmarshals PatternFlowCfmMdLevel from protobuf object *otg.PatternFlowCfmMdLevel
	FromProto(msg *otg.PatternFlowCfmMdLevel) (PatternFlowCfmMdLevel, error)
	// FromPbText unmarshals PatternFlowCfmMdLevel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmMdLevel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmMdLevel from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmMdLevel) Marshal() marshalPatternFlowCfmMdLevel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmMdLevel{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmMdLevel) Unmarshal() unMarshalPatternFlowCfmMdLevel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmMdLevel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmMdLevel) ToProto() (*otg.PatternFlowCfmMdLevel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmMdLevel) FromProto(msg *otg.PatternFlowCfmMdLevel) (PatternFlowCfmMdLevel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmMdLevel) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevel) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmMdLevel) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevel) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmMdLevel) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmMdLevel) FromJson(value string) error {
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

func (obj *patternFlowCfmMdLevel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmMdLevel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmMdLevel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmMdLevel) Clone() (PatternFlowCfmMdLevel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmMdLevel()
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

func (obj *patternFlowCfmMdLevel) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmMdLevel is maintenance Domain Level
type PatternFlowCfmMdLevel interface {
	Validation
	// msg marshals PatternFlowCfmMdLevel to protobuf object *otg.PatternFlowCfmMdLevel
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmMdLevel
	// setMsg unmarshals PatternFlowCfmMdLevel from protobuf object *otg.PatternFlowCfmMdLevel
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmMdLevel) PatternFlowCfmMdLevel
	// provides marshal interface
	Marshal() marshalPatternFlowCfmMdLevel
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmMdLevel
	// validate validates PatternFlowCfmMdLevel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmMdLevel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmMdLevelChoiceEnum, set in PatternFlowCfmMdLevel
	Choice() PatternFlowCfmMdLevelChoiceEnum
	// setChoice assigns PatternFlowCfmMdLevelChoiceEnum provided by user to PatternFlowCfmMdLevel
	setChoice(value PatternFlowCfmMdLevelChoiceEnum) PatternFlowCfmMdLevel
	// HasChoice checks if Choice has been set in PatternFlowCfmMdLevel
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmMdLevel.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmMdLevel
	SetValue(value uint32) PatternFlowCfmMdLevel
	// HasValue checks if Value has been set in PatternFlowCfmMdLevel
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmMdLevel.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmMdLevel
	SetValues(value []uint32) PatternFlowCfmMdLevel
	// Increment returns PatternFlowCfmMdLevelCounter, set in PatternFlowCfmMdLevel.
	// PatternFlowCfmMdLevelCounter is integer counter pattern
	Increment() PatternFlowCfmMdLevelCounter
	// SetIncrement assigns PatternFlowCfmMdLevelCounter provided by user to PatternFlowCfmMdLevel.
	// PatternFlowCfmMdLevelCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevel
	// HasIncrement checks if Increment has been set in PatternFlowCfmMdLevel
	HasIncrement() bool
	// Decrement returns PatternFlowCfmMdLevelCounter, set in PatternFlowCfmMdLevel.
	// PatternFlowCfmMdLevelCounter is integer counter pattern
	Decrement() PatternFlowCfmMdLevelCounter
	// SetDecrement assigns PatternFlowCfmMdLevelCounter provided by user to PatternFlowCfmMdLevel.
	// PatternFlowCfmMdLevelCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevel
	// HasDecrement checks if Decrement has been set in PatternFlowCfmMdLevel
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmMdLevelChoiceEnum string

// Enum of Choice on PatternFlowCfmMdLevel
var PatternFlowCfmMdLevelChoice = struct {
	VALUE     PatternFlowCfmMdLevelChoiceEnum
	VALUES    PatternFlowCfmMdLevelChoiceEnum
	INCREMENT PatternFlowCfmMdLevelChoiceEnum
	DECREMENT PatternFlowCfmMdLevelChoiceEnum
}{
	VALUE:     PatternFlowCfmMdLevelChoiceEnum("value"),
	VALUES:    PatternFlowCfmMdLevelChoiceEnum("values"),
	INCREMENT: PatternFlowCfmMdLevelChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmMdLevelChoiceEnum("decrement"),
}

func (obj *patternFlowCfmMdLevel) Choice() PatternFlowCfmMdLevelChoiceEnum {
	return PatternFlowCfmMdLevelChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmMdLevel) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmMdLevel) setChoice(value PatternFlowCfmMdLevelChoiceEnum) PatternFlowCfmMdLevel {
	intValue, ok := otg.PatternFlowCfmMdLevel_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmMdLevelChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmMdLevel_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmMdLevelChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmMdLevelChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmMdLevelChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmMdLevelCounter().msg()
	}

	if value == PatternFlowCfmMdLevelChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmMdLevelCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmMdLevel) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmMdLevelChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmMdLevel) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmMdLevel object
func (obj *patternFlowCfmMdLevel) SetValue(value uint32) PatternFlowCfmMdLevel {
	obj.setChoice(PatternFlowCfmMdLevelChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmMdLevel) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmMdLevel object
func (obj *patternFlowCfmMdLevel) SetValues(value []uint32) PatternFlowCfmMdLevel {
	obj.setChoice(PatternFlowCfmMdLevelChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmMdLevelCounter
func (obj *patternFlowCfmMdLevel) Increment() PatternFlowCfmMdLevelCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmMdLevelChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmMdLevelCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmMdLevelCounter
func (obj *patternFlowCfmMdLevel) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmMdLevelCounter value in the PatternFlowCfmMdLevel object
func (obj *patternFlowCfmMdLevel) SetIncrement(value PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevel {
	obj.setChoice(PatternFlowCfmMdLevelChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmMdLevelCounter
func (obj *patternFlowCfmMdLevel) Decrement() PatternFlowCfmMdLevelCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmMdLevelChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmMdLevelCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmMdLevelCounter
func (obj *patternFlowCfmMdLevel) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmMdLevelCounter value in the PatternFlowCfmMdLevel object
func (obj *patternFlowCfmMdLevel) SetDecrement(value PatternFlowCfmMdLevelCounter) PatternFlowCfmMdLevel {
	obj.setChoice(PatternFlowCfmMdLevelChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmMdLevel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmMdLevel.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmMdLevel.Values <= 7 but Got %d", item))
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

func (obj *patternFlowCfmMdLevel) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmMdLevelChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmMdLevelChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmMdLevelChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmMdLevelChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmMdLevelChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmMdLevelChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmMdLevel")
			}
		} else {
			intVal := otg.PatternFlowCfmMdLevel_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmMdLevel_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
