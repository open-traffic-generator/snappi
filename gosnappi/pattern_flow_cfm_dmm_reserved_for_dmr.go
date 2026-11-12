package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmr *****
type patternFlowCfmDmmReservedForDmr struct {
	validation
	obj             *otg.PatternFlowCfmDmmReservedForDmr
	marshaller      marshalPatternFlowCfmDmmReservedForDmr
	unMarshaller    unMarshalPatternFlowCfmDmmReservedForDmr
	incrementHolder PatternFlowCfmDmmReservedForDmrCounter
	decrementHolder PatternFlowCfmDmmReservedForDmrCounter
}

func NewPatternFlowCfmDmmReservedForDmr() PatternFlowCfmDmmReservedForDmr {
	obj := patternFlowCfmDmmReservedForDmr{obj: &otg.PatternFlowCfmDmmReservedForDmr{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmr) msg() *otg.PatternFlowCfmDmmReservedForDmr {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmr) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmr) PatternFlowCfmDmmReservedForDmr {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmr struct {
	obj *patternFlowCfmDmmReservedForDmr
}

type marshalPatternFlowCfmDmmReservedForDmr interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmr to protobuf object *otg.PatternFlowCfmDmmReservedForDmr
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmr, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmr to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmr to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmr struct {
	obj *patternFlowCfmDmmReservedForDmr
}

type unMarshalPatternFlowCfmDmmReservedForDmr interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmr from protobuf object *otg.PatternFlowCfmDmmReservedForDmr
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmr) (PatternFlowCfmDmmReservedForDmr, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmr from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmr) Marshal() marshalPatternFlowCfmDmmReservedForDmr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmr{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmr) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmr) ToProto() (*otg.PatternFlowCfmDmmReservedForDmr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmr) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmr) (PatternFlowCfmDmmReservedForDmr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmr) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmr) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmr) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmr) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmr) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmr) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmr) Clone() (PatternFlowCfmDmmReservedForDmr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmr()
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

func (obj *patternFlowCfmDmmReservedForDmr) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmmReservedForDmr is reserved for the DMM receiving equipment to record its transmit timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmmReservedForDmr interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmr to protobuf object *otg.PatternFlowCfmDmmReservedForDmr
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmr
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmr from protobuf object *otg.PatternFlowCfmDmmReservedForDmr
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmr) PatternFlowCfmDmmReservedForDmr
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmr
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmr
	// validate validates PatternFlowCfmDmmReservedForDmr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmmReservedForDmrChoiceEnum, set in PatternFlowCfmDmmReservedForDmr
	Choice() PatternFlowCfmDmmReservedForDmrChoiceEnum
	// setChoice assigns PatternFlowCfmDmmReservedForDmrChoiceEnum provided by user to PatternFlowCfmDmmReservedForDmr
	setChoice(value PatternFlowCfmDmmReservedForDmrChoiceEnum) PatternFlowCfmDmmReservedForDmr
	// HasChoice checks if Choice has been set in PatternFlowCfmDmmReservedForDmr
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmmReservedForDmr.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmr
	SetValue(value uint64) PatternFlowCfmDmmReservedForDmr
	// HasValue checks if Value has been set in PatternFlowCfmDmmReservedForDmr
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmmReservedForDmr.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmmReservedForDmr
	SetValues(value []uint64) PatternFlowCfmDmmReservedForDmr
	// Increment returns PatternFlowCfmDmmReservedForDmrCounter, set in PatternFlowCfmDmmReservedForDmr.
	// PatternFlowCfmDmmReservedForDmrCounter is integer counter pattern
	Increment() PatternFlowCfmDmmReservedForDmrCounter
	// SetIncrement assigns PatternFlowCfmDmmReservedForDmrCounter provided by user to PatternFlowCfmDmmReservedForDmr.
	// PatternFlowCfmDmmReservedForDmrCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmr
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmmReservedForDmr
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmmReservedForDmrCounter, set in PatternFlowCfmDmmReservedForDmr.
	// PatternFlowCfmDmmReservedForDmrCounter is integer counter pattern
	Decrement() PatternFlowCfmDmmReservedForDmrCounter
	// SetDecrement assigns PatternFlowCfmDmmReservedForDmrCounter provided by user to PatternFlowCfmDmmReservedForDmr.
	// PatternFlowCfmDmmReservedForDmrCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmr
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmmReservedForDmr
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmmReservedForDmrChoiceEnum string

// Enum of Choice on PatternFlowCfmDmmReservedForDmr
var PatternFlowCfmDmmReservedForDmrChoice = struct {
	VALUE     PatternFlowCfmDmmReservedForDmrChoiceEnum
	VALUES    PatternFlowCfmDmmReservedForDmrChoiceEnum
	INCREMENT PatternFlowCfmDmmReservedForDmrChoiceEnum
	DECREMENT PatternFlowCfmDmmReservedForDmrChoiceEnum
}{
	VALUE:     PatternFlowCfmDmmReservedForDmrChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmmReservedForDmrChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmmReservedForDmrChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmmReservedForDmrChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmmReservedForDmr) Choice() PatternFlowCfmDmmReservedForDmrChoiceEnum {
	return PatternFlowCfmDmmReservedForDmrChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmmReservedForDmr) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmmReservedForDmr) setChoice(value PatternFlowCfmDmmReservedForDmrChoiceEnum) PatternFlowCfmDmmReservedForDmr {
	intValue, ok := otg.PatternFlowCfmDmmReservedForDmr_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmmReservedForDmrChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmmReservedForDmr_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmmReservedForDmrChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmrChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmrChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmmReservedForDmrCounter().msg()
	}

	if value == PatternFlowCfmDmmReservedForDmrChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmmReservedForDmrCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmr) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmr) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmmReservedForDmr object
func (obj *patternFlowCfmDmmReservedForDmr) SetValue(value uint64) PatternFlowCfmDmmReservedForDmr {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmmReservedForDmr) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmmReservedForDmr object
func (obj *patternFlowCfmDmmReservedForDmr) SetValues(value []uint64) PatternFlowCfmDmmReservedForDmr {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmrCounter
func (obj *patternFlowCfmDmmReservedForDmr) Increment() PatternFlowCfmDmmReservedForDmrCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmmReservedForDmrCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmrCounter
func (obj *patternFlowCfmDmmReservedForDmr) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmmReservedForDmrCounter value in the PatternFlowCfmDmmReservedForDmr object
func (obj *patternFlowCfmDmmReservedForDmr) SetIncrement(value PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmr {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmrCounter
func (obj *patternFlowCfmDmmReservedForDmr) Decrement() PatternFlowCfmDmmReservedForDmrCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmmReservedForDmrCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmrCounter
func (obj *patternFlowCfmDmmReservedForDmr) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmmReservedForDmrCounter value in the PatternFlowCfmDmmReservedForDmr object
func (obj *patternFlowCfmDmmReservedForDmr) SetDecrement(value PatternFlowCfmDmmReservedForDmrCounter) PatternFlowCfmDmmReservedForDmr {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmmReservedForDmr) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmmReservedForDmr) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmmReservedForDmrChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmmReservedForDmrChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmmReservedForDmr")
			}
		} else {
			intVal := otg.PatternFlowCfmDmmReservedForDmr_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmmReservedForDmr_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
