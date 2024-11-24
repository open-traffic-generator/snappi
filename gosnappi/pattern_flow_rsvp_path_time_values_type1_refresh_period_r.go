package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathTimeValuesType1RefreshPeriodR *****
type patternFlowRSVPPathTimeValuesType1RefreshPeriodR struct {
	validation
	obj             *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	marshaller      marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	unMarshaller    unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	incrementHolder PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	decrementHolder PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
}

func NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodR() PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj := patternFlowRSVPPathTimeValuesType1RefreshPeriodR{obj: &otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) msg() *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	return obj.obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) setMsg(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR struct {
	obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR
}

type marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR interface {
	// ToProto marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	ToProto() (*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error)
	// ToPbText marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR struct {
	obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR
}

type unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR interface {
	// FromProto unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR from protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	FromProto(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) (PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error)
	// FromPbText unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Marshal() marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Unmarshal() unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) ToProto() (*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FromProto(msg *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) (PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Clone() (PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodR()
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

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathTimeValuesType1RefreshPeriodR is the refresh timeout period R used to generate this message;in milliseconds.
type PatternFlowRSVPPathTimeValuesType1RefreshPeriodR interface {
	Validation
	// msg marshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR to protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// setMsg unmarshals PatternFlowRSVPPathTimeValuesType1RefreshPeriodR from protobuf object *otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// validate validates PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	Choice() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum
	// setChoice assigns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	setChoice(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	SetValue(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// HasValue checks if Value has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	SetValues(value []uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// Increment returns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter is integer counter pattern
	Increment() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// SetIncrement assigns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter, set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
	// SetDecrement assigns PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter provided by user to PatternFlowRSVPPathTimeValuesType1RefreshPeriodR.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
var PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice = struct {
	VALUE     PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum
	VALUES    PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum
	INCREMENT PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum
	DECREMENT PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Choice() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum {
	return PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) setChoice(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	intValue, ok := otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUE {
		defaultValue := uint32(30000)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUES {
		defaultValue := []uint32{30000}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter().msg()
	}

	if value == PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodR object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) SetValue(value uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{30000})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodR object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) SetValues(value []uint32) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Increment() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodR object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) SetIncrement(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) Decrement() PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter value in the PatternFlowRSVPPathTimeValuesType1RefreshPeriodR object
func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) SetDecrement(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodRCounter) PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowRSVPPathTimeValuesType1RefreshPeriodR) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathTimeValuesType1RefreshPeriodRChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathTimeValuesType1RefreshPeriodR")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathTimeValuesType1RefreshPeriodR_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
