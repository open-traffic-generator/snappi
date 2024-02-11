package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVariableBindingValueBigCounterValue *****
type patternFlowSnmpv2CVariableBindingValueBigCounterValue struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	marshaller      marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue
	unMarshaller    unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue
	incrementHolder PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	decrementHolder PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
}

func NewPatternFlowSnmpv2CVariableBindingValueBigCounterValue() PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj := patternFlowSnmpv2CVariableBindingValueBigCounterValue{obj: &otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) msg() *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) setMsg(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue
}

type marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue interface {
	// ToProto marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error)
	// ToPbText marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue struct {
	obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue
}

type unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue interface {
	// FromProto unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue) (PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Marshal() marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) ToProto() (*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) FromProto(msg *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue) (PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVariableBindingValueBigCounterValue) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Clone() (PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVariableBindingValueBigCounterValue()
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

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVariableBindingValueBigCounterValue is big counter returned for the requested OID.
type PatternFlowSnmpv2CVariableBindingValueBigCounterValue interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue to protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// setMsg unmarshals PatternFlowSnmpv2CVariableBindingValueBigCounterValue from protobuf object *otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// validate validates PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVariableBindingValueBigCounterValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	Choice() PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	setChoice(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	HasChoice() bool
	// Value returns uint64, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	SetValue(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	HasValue() bool
	// Values returns []uint64, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	SetValues(value []uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// Increment returns PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	Increment() PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// SetIncrement assigns PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	SetIncrement(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter, set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	Decrement() PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
	// SetDecrement assigns PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter provided by user to PatternFlowSnmpv2CVariableBindingValueBigCounterValue.
	SetDecrement(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVariableBindingValueBigCounterValue
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVariableBindingValueBigCounterValue
var PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice = struct {
	VALUE     PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum
	VALUES    PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum
	INCREMENT PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum
	DECREMENT PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Choice() PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum {
	return PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) setChoice(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	intValue, ok := otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVariableBindingValueBigCounterValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter().msg()
	}

	if value == PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) SetValue(value uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) SetValues(value []uint64) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Increment() PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) SetIncrement(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) Decrement() PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVariableBindingValueBigCounterValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter value in the PatternFlowSnmpv2CVariableBindingValueBigCounterValue object
func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) SetDecrement(value PatternFlowSnmpv2CVariableBindingValueBigCounterValueCounter) PatternFlowSnmpv2CVariableBindingValueBigCounterValue {
	obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CVariableBindingValueBigCounterValue) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowSnmpv2CVariableBindingValueBigCounterValueChoice.VALUE)

	}

}
