package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueTimeticksValue *****
type patternFlowSnmpv2CVariableBindingValueTimeticksValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueTimeticksValue() PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj := patternFlowSnmpv2CVariableBindingValueTimeticksValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue) (PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue) (PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueTimeticksValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Clone() (PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueTimeticksValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueTimeticksValue is timeticks returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueTimeticksValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueTimeticksValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueTimeticksValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	Choice() PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter, set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueTimeticksValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueTimeticksValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueTimeticksValue
var PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Choice() PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueTimeticksValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValue object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) SetValue(value uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValue object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) SetValues(value []uint32) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Increment() PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValue object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) Decrement() PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueTimeticksValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter value in the PatternFlowSnmpv2CVariableBindingValueTimeticksValue object
func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueTimeticksValueCounter) PatternFlowSnmpv2CVariableBindingValueTimeticksValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CVariableBindingValueTimeticksValue) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueTimeticksValueChoice.VALUE)

	}

}
