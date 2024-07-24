package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue *****
type patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj := patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Clone() (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue is unsigned integer value returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	Choice() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter, set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue
var PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Choice() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Increment() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) Decrement() PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter value in the PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue object
func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueCounter) PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CVariableBindingValueUnsignedIntegerValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
