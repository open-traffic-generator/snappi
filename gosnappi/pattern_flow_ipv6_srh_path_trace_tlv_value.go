package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvValue *****
type patternFlowIpv6SRHPathTraceTlvValue struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvValue
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvValue
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvValue
	incrementHolder PatternFlowIpv6SRHPathTraceTlvValueCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvValueCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvValue() PatternFlowIpv6SRHPathTraceTlvValue {
	obj := patternFlowIpv6SRHPathTraceTlvValue{obj: &otg.PatternFlowIpv6SRHPathTraceTlvValue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) msg() *otg.PatternFlowIpv6SRHPathTraceTlvValue {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvValue) PatternFlowIpv6SRHPathTraceTlvValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvValue struct {
	obj *patternFlowIpv6SRHPathTraceTlvValue
}

type marshalPatternFlowIpv6SRHPathTraceTlvValue interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvValue to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValue
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvValue, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvValue to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvValue struct {
	obj *patternFlowIpv6SRHPathTraceTlvValue
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvValue interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvValue from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValue
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvValue) (PatternFlowIpv6SRHPathTraceTlvValue, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvValue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValue) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValue) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvValue) (PatternFlowIpv6SRHPathTraceTlvValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValue) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValue) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvValue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvValue) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) Clone() (PatternFlowIpv6SRHPathTraceTlvValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvValue()
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

func (obj *patternFlowIpv6SRHPathTraceTlvValue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvValue is initial node identification value inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Identifies the ingress node within the path trace record.
type PatternFlowIpv6SRHPathTraceTlvValue interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvValue to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValue
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvValue
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvValue from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvValue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvValue) PatternFlowIpv6SRHPathTraceTlvValue
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvValue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvValue
	// validate validates PatternFlowIpv6SRHPathTraceTlvValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvValue
	Choice() PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvValue
	setChoice(value PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum) PatternFlowIpv6SRHPathTraceTlvValue
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvValue
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvValue
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvValue
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvValue
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvValue.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvValue
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvValue
	// Increment returns PatternFlowIpv6SRHPathTraceTlvValueCounter, set in PatternFlowIpv6SRHPathTraceTlvValue.
	// PatternFlowIpv6SRHPathTraceTlvValueCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvValueCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvValueCounter provided by user to PatternFlowIpv6SRHPathTraceTlvValue.
	// PatternFlowIpv6SRHPathTraceTlvValueCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValue
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvValue
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvValueCounter, set in PatternFlowIpv6SRHPathTraceTlvValue.
	// PatternFlowIpv6SRHPathTraceTlvValueCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvValueCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvValueCounter provided by user to PatternFlowIpv6SRHPathTraceTlvValue.
	// PatternFlowIpv6SRHPathTraceTlvValueCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValue
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvValue
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvValue
var PatternFlowIpv6SRHPathTraceTlvValueChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) Choice() PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvValue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) setChoice(value PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum) PatternFlowIpv6SRHPathTraceTlvValue {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvValue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvValue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvValueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvValueCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvValueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvValueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValue) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvValue object
func (obj *patternFlowIpv6SRHPathTraceTlvValue) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvValue {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvValue) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvValue object
func (obj *patternFlowIpv6SRHPathTraceTlvValue) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvValue {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvValueCounter
func (obj *patternFlowIpv6SRHPathTraceTlvValue) Increment() PatternFlowIpv6SRHPathTraceTlvValueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvValueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvValueCounter
func (obj *patternFlowIpv6SRHPathTraceTlvValue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvValueCounter value in the PatternFlowIpv6SRHPathTraceTlvValue object
func (obj *patternFlowIpv6SRHPathTraceTlvValue) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValue {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvValueCounter
func (obj *patternFlowIpv6SRHPathTraceTlvValue) Decrement() PatternFlowIpv6SRHPathTraceTlvValueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvValueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvValueCounter
func (obj *patternFlowIpv6SRHPathTraceTlvValue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvValueCounter value in the PatternFlowIpv6SRHPathTraceTlvValue object
func (obj *patternFlowIpv6SRHPathTraceTlvValue) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvValueCounter) PatternFlowIpv6SRHPathTraceTlvValue {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvValue) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIpv6SRHPathTraceTlvValue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvValueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvValueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvValueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvValueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvValue")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvValue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvValue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
