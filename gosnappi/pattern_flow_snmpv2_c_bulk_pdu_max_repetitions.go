package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CBulkPDUMaxRepetitions *****
type patternFlowSnmpv2CBulkPDUMaxRepetitions struct {
	validation
	obj             *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	marshaller      marshalPatternFlowSnmpv2CBulkPDUMaxRepetitions
	unMarshaller    unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitions
	incrementHolder PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	decrementHolder PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
}

func NewPatternFlowSnmpv2CBulkPDUMaxRepetitions() PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj := patternFlowSnmpv2CBulkPDUMaxRepetitions{obj: &otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) msg() *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	return obj.obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) setMsg(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions struct {
	obj *patternFlowSnmpv2CBulkPDUMaxRepetitions
}

type marshalPatternFlowSnmpv2CBulkPDUMaxRepetitions interface {
	// ToProto marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	ToProto() (*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions, error)
	// ToPbText marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions struct {
	obj *patternFlowSnmpv2CBulkPDUMaxRepetitions
}

type unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitions interface {
	// FromProto unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitions from protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions) (PatternFlowSnmpv2CBulkPDUMaxRepetitions, error)
	// FromPbText unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitions from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Marshal() marshalPatternFlowSnmpv2CBulkPDUMaxRepetitions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) ToProto() (*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) FromProto(msg *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions) (PatternFlowSnmpv2CBulkPDUMaxRepetitions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CBulkPDUMaxRepetitions) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Clone() (PatternFlowSnmpv2CBulkPDUMaxRepetitions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CBulkPDUMaxRepetitions()
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

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CBulkPDUMaxRepetitions is a maximum of max_repetitions variable bindings are requested in the Response-PDU for each of the remaining variable bindings in the GetBulkRequest after the non_repeaters variable bindings.
type PatternFlowSnmpv2CBulkPDUMaxRepetitions interface {
	Validation
	// msg marshals PatternFlowSnmpv2CBulkPDUMaxRepetitions to protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// setMsg unmarshals PatternFlowSnmpv2CBulkPDUMaxRepetitions from protobuf object *otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CBulkPDUMaxRepetitions
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CBulkPDUMaxRepetitions
	// validate validates PatternFlowSnmpv2CBulkPDUMaxRepetitions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CBulkPDUMaxRepetitions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum, set in PatternFlowSnmpv2CBulkPDUMaxRepetitions
	Choice() PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitions
	setChoice(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitions
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitions
	SetValue(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// HasValue checks if Value has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitions
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitions
	SetValues(value []uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// Increment returns PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, set in PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	Increment() PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// SetIncrement assigns PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	SetIncrement(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitions
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter, set in PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	Decrement() PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
	// SetDecrement assigns PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter provided by user to PatternFlowSnmpv2CBulkPDUMaxRepetitions.
	SetDecrement(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitions
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CBulkPDUMaxRepetitions
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CBulkPDUMaxRepetitions
var PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice = struct {
	VALUE     PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum
	VALUES    PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum
	INCREMENT PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum
	DECREMENT PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Choice() PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum {
	return PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) setChoice(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	intValue, ok := otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter().msg()
	}

	if value == PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CBulkPDUMaxRepetitions object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) SetValue(value uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CBulkPDUMaxRepetitions object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) SetValues(value []uint32) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Increment() PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter value in the PatternFlowSnmpv2CBulkPDUMaxRepetitions object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) SetIncrement(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) Decrement() PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CBulkPDUMaxRepetitionsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter value in the PatternFlowSnmpv2CBulkPDUMaxRepetitions object
func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) SetDecrement(value PatternFlowSnmpv2CBulkPDUMaxRepetitionsCounter) PatternFlowSnmpv2CBulkPDUMaxRepetitions {
	obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CBulkPDUMaxRepetitions) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CBulkPDUMaxRepetitionsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CBulkPDUMaxRepetitions")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CBulkPDUMaxRepetitions_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
