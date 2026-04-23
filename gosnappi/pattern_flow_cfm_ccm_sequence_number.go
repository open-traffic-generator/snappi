package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmSequenceNumber *****
type patternFlowCfmCcmSequenceNumber struct {
	validation
	obj             *otg.PatternFlowCfmCcmSequenceNumber
	marshaller      marshalPatternFlowCfmCcmSequenceNumber
	unMarshaller    unMarshalPatternFlowCfmCcmSequenceNumber
	incrementHolder PatternFlowCfmCcmSequenceNumberCounter
	decrementHolder PatternFlowCfmCcmSequenceNumberCounter
}

func NewPatternFlowCfmCcmSequenceNumber() PatternFlowCfmCcmSequenceNumber {
	obj := patternFlowCfmCcmSequenceNumber{obj: &otg.PatternFlowCfmCcmSequenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmSequenceNumber) msg() *otg.PatternFlowCfmCcmSequenceNumber {
	return obj.obj
}

func (obj *patternFlowCfmCcmSequenceNumber) setMsg(msg *otg.PatternFlowCfmCcmSequenceNumber) PatternFlowCfmCcmSequenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmSequenceNumber struct {
	obj *patternFlowCfmCcmSequenceNumber
}

type marshalPatternFlowCfmCcmSequenceNumber interface {
	// ToProto marshals PatternFlowCfmCcmSequenceNumber to protobuf object *otg.PatternFlowCfmCcmSequenceNumber
	ToProto() (*otg.PatternFlowCfmCcmSequenceNumber, error)
	// ToPbText marshals PatternFlowCfmCcmSequenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmSequenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmSequenceNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmSequenceNumber struct {
	obj *patternFlowCfmCcmSequenceNumber
}

type unMarshalPatternFlowCfmCcmSequenceNumber interface {
	// FromProto unmarshals PatternFlowCfmCcmSequenceNumber from protobuf object *otg.PatternFlowCfmCcmSequenceNumber
	FromProto(msg *otg.PatternFlowCfmCcmSequenceNumber) (PatternFlowCfmCcmSequenceNumber, error)
	// FromPbText unmarshals PatternFlowCfmCcmSequenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmSequenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmSequenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmSequenceNumber) Marshal() marshalPatternFlowCfmCcmSequenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmSequenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmSequenceNumber) Unmarshal() unMarshalPatternFlowCfmCcmSequenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmSequenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmSequenceNumber) ToProto() (*otg.PatternFlowCfmCcmSequenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmSequenceNumber) FromProto(msg *otg.PatternFlowCfmCcmSequenceNumber) (PatternFlowCfmCcmSequenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmSequenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmSequenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmSequenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmSequenceNumber) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmSequenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmSequenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmSequenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmSequenceNumber) Clone() (PatternFlowCfmCcmSequenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmSequenceNumber()
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

func (obj *patternFlowCfmCcmSequenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmCcmSequenceNumber is cCM unique sequence number
type PatternFlowCfmCcmSequenceNumber interface {
	Validation
	// msg marshals PatternFlowCfmCcmSequenceNumber to protobuf object *otg.PatternFlowCfmCcmSequenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmSequenceNumber
	// setMsg unmarshals PatternFlowCfmCcmSequenceNumber from protobuf object *otg.PatternFlowCfmCcmSequenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmSequenceNumber) PatternFlowCfmCcmSequenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmSequenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmSequenceNumber
	// validate validates PatternFlowCfmCcmSequenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmSequenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmCcmSequenceNumberChoiceEnum, set in PatternFlowCfmCcmSequenceNumber
	Choice() PatternFlowCfmCcmSequenceNumberChoiceEnum
	// setChoice assigns PatternFlowCfmCcmSequenceNumberChoiceEnum provided by user to PatternFlowCfmCcmSequenceNumber
	setChoice(value PatternFlowCfmCcmSequenceNumberChoiceEnum) PatternFlowCfmCcmSequenceNumber
	// HasChoice checks if Choice has been set in PatternFlowCfmCcmSequenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmCcmSequenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmCcmSequenceNumber
	SetValue(value uint32) PatternFlowCfmCcmSequenceNumber
	// HasValue checks if Value has been set in PatternFlowCfmCcmSequenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmCcmSequenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmCcmSequenceNumber
	SetValues(value []uint32) PatternFlowCfmCcmSequenceNumber
	// Increment returns PatternFlowCfmCcmSequenceNumberCounter, set in PatternFlowCfmCcmSequenceNumber.
	// PatternFlowCfmCcmSequenceNumberCounter is integer counter pattern
	Increment() PatternFlowCfmCcmSequenceNumberCounter
	// SetIncrement assigns PatternFlowCfmCcmSequenceNumberCounter provided by user to PatternFlowCfmCcmSequenceNumber.
	// PatternFlowCfmCcmSequenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowCfmCcmSequenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowCfmCcmSequenceNumberCounter, set in PatternFlowCfmCcmSequenceNumber.
	// PatternFlowCfmCcmSequenceNumberCounter is integer counter pattern
	Decrement() PatternFlowCfmCcmSequenceNumberCounter
	// SetDecrement assigns PatternFlowCfmCcmSequenceNumberCounter provided by user to PatternFlowCfmCcmSequenceNumber.
	// PatternFlowCfmCcmSequenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowCfmCcmSequenceNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmCcmSequenceNumberChoiceEnum string

// Enum of Choice on PatternFlowCfmCcmSequenceNumber
var PatternFlowCfmCcmSequenceNumberChoice = struct {
	VALUE     PatternFlowCfmCcmSequenceNumberChoiceEnum
	VALUES    PatternFlowCfmCcmSequenceNumberChoiceEnum
	INCREMENT PatternFlowCfmCcmSequenceNumberChoiceEnum
	DECREMENT PatternFlowCfmCcmSequenceNumberChoiceEnum
}{
	VALUE:     PatternFlowCfmCcmSequenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowCfmCcmSequenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowCfmCcmSequenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmCcmSequenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowCfmCcmSequenceNumber) Choice() PatternFlowCfmCcmSequenceNumberChoiceEnum {
	return PatternFlowCfmCcmSequenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmCcmSequenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmCcmSequenceNumber) setChoice(value PatternFlowCfmCcmSequenceNumberChoiceEnum) PatternFlowCfmCcmSequenceNumber {
	intValue, ok := otg.PatternFlowCfmCcmSequenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmCcmSequenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmCcmSequenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmCcmSequenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmCcmSequenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmCcmSequenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmCcmSequenceNumberCounter().msg()
	}

	if value == PatternFlowCfmCcmSequenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmCcmSequenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmSequenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmSequenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmCcmSequenceNumber object
func (obj *patternFlowCfmCcmSequenceNumber) SetValue(value uint32) PatternFlowCfmCcmSequenceNumber {
	obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmCcmSequenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmCcmSequenceNumber object
func (obj *patternFlowCfmCcmSequenceNumber) SetValues(value []uint32) PatternFlowCfmCcmSequenceNumber {
	obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmCcmSequenceNumberCounter
func (obj *patternFlowCfmCcmSequenceNumber) Increment() PatternFlowCfmCcmSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmCcmSequenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmCcmSequenceNumberCounter
func (obj *patternFlowCfmCcmSequenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmCcmSequenceNumberCounter value in the PatternFlowCfmCcmSequenceNumber object
func (obj *patternFlowCfmCcmSequenceNumber) SetIncrement(value PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumber {
	obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmSequenceNumberCounter
func (obj *patternFlowCfmCcmSequenceNumber) Decrement() PatternFlowCfmCcmSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmCcmSequenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmSequenceNumberCounter
func (obj *patternFlowCfmCcmSequenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmCcmSequenceNumberCounter value in the PatternFlowCfmCcmSequenceNumber object
func (obj *patternFlowCfmCcmSequenceNumber) SetDecrement(value PatternFlowCfmCcmSequenceNumberCounter) PatternFlowCfmCcmSequenceNumber {
	obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmCcmSequenceNumber) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmCcmSequenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmCcmSequenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmSequenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmCcmSequenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmSequenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmSequenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmCcmSequenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmCcmSequenceNumber")
			}
		} else {
			intVal := otg.PatternFlowCfmCcmSequenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmCcmSequenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
