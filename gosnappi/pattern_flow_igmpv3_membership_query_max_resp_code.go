package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryMaxRespCode *****
type patternFlowIgmpv3MembershipQueryMaxRespCode struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	marshaller      marshalPatternFlowIgmpv3MembershipQueryMaxRespCode
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCode
	incrementHolder PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
}

func NewPatternFlowIgmpv3MembershipQueryMaxRespCode() PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj := patternFlowIgmpv3MembershipQueryMaxRespCode{obj: &otg.PatternFlowIgmpv3MembershipQueryMaxRespCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) msg() *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryMaxRespCode struct {
	obj *patternFlowIgmpv3MembershipQueryMaxRespCode
}

type marshalPatternFlowIgmpv3MembershipQueryMaxRespCode interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryMaxRespCode to protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryMaxRespCode, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryMaxRespCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryMaxRespCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryMaxRespCode to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode struct {
	obj *patternFlowIgmpv3MembershipQueryMaxRespCode
}

type unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCode interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCode from protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode) (PatternFlowIgmpv3MembershipQueryMaxRespCode, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Marshal() marshalPatternFlowIgmpv3MembershipQueryMaxRespCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryMaxRespCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCode) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryMaxRespCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode) (PatternFlowIgmpv3MembershipQueryMaxRespCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryMaxRespCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryMaxRespCode) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Clone() (PatternFlowIgmpv3MembershipQueryMaxRespCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryMaxRespCode()
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

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryMaxRespCode is maximum time in units of 1/10 second within which hosts must respond to the query, so the default 100 is 10 seconds. Values of 128 or more use the exponent and mantissa encoding of RFC 9776 Section 4.1.1.
type PatternFlowIgmpv3MembershipQueryMaxRespCode interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryMaxRespCode to protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryMaxRespCode from protobuf object *otg.PatternFlowIgmpv3MembershipQueryMaxRespCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryMaxRespCode) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryMaxRespCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryMaxRespCode
	// validate validates PatternFlowIgmpv3MembershipQueryMaxRespCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryMaxRespCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum, set in PatternFlowIgmpv3MembershipQueryMaxRespCode
	Choice() PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCode
	setChoice(value PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryMaxRespCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryMaxRespCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCode
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryMaxRespCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryMaxRespCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCode
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// Increment returns PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, set in PatternFlowIgmpv3MembershipQueryMaxRespCode.
	// PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCode.
	// PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryMaxRespCode
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter, set in PatternFlowIgmpv3MembershipQueryMaxRespCode.
	// PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter provided by user to PatternFlowIgmpv3MembershipQueryMaxRespCode.
	// PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCode
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryMaxRespCode
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryMaxRespCode
var PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Choice() PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) setChoice(value PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryMaxRespCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryMaxRespCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUE {
		defaultValue := uint32(100)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUES {
		defaultValue := []uint32{100}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryMaxRespCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryMaxRespCode object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{100})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryMaxRespCode object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Increment() PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryMaxRespCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter value in the PatternFlowIgmpv3MembershipQueryMaxRespCode object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) SetIncrement(value PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) Decrement() PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryMaxRespCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter value in the PatternFlowIgmpv3MembershipQueryMaxRespCode object
func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) SetDecrement(value PatternFlowIgmpv3MembershipQueryMaxRespCodeCounter) PatternFlowIgmpv3MembershipQueryMaxRespCode {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryMaxRespCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryMaxRespCode.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQueryMaxRespCode) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryMaxRespCodeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryMaxRespCodeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryMaxRespCode")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryMaxRespCode_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryMaxRespCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
