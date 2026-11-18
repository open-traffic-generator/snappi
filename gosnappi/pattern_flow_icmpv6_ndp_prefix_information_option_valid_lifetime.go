package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime *****
type patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime is length of time in seconds that the prefix is valid for on-link determination. A value of 0xffffffff represents infinity.
type PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
var PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUE {
		defaultValue := uint32(4294967295)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUES {
		defaultValue := []uint32{4294967295}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4294967295})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeCounter) PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetimeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
