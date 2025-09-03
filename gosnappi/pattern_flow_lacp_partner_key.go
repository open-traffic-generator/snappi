package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerKey *****
type patternFlowLacpPartnerKey struct {
	validation
	obj             *otg.PatternFlowLacpPartnerKey
	marshaller      marshalPatternFlowLacpPartnerKey
	unMarshaller    unMarshalPatternFlowLacpPartnerKey
	incrementHolder PatternFlowLacpPartnerKeyCounter
	decrementHolder PatternFlowLacpPartnerKeyCounter
}

func NewPatternFlowLacpPartnerKey() PatternFlowLacpPartnerKey {
	obj := patternFlowLacpPartnerKey{obj: &otg.PatternFlowLacpPartnerKey{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerKey) msg() *otg.PatternFlowLacpPartnerKey {
	return obj.obj
}

func (obj *patternFlowLacpPartnerKey) setMsg(msg *otg.PatternFlowLacpPartnerKey) PatternFlowLacpPartnerKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerKey struct {
	obj *patternFlowLacpPartnerKey
}

type marshalPatternFlowLacpPartnerKey interface {
	// ToProto marshals PatternFlowLacpPartnerKey to protobuf object *otg.PatternFlowLacpPartnerKey
	ToProto() (*otg.PatternFlowLacpPartnerKey, error)
	// ToPbText marshals PatternFlowLacpPartnerKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerKey to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerKey struct {
	obj *patternFlowLacpPartnerKey
}

type unMarshalPatternFlowLacpPartnerKey interface {
	// FromProto unmarshals PatternFlowLacpPartnerKey from protobuf object *otg.PatternFlowLacpPartnerKey
	FromProto(msg *otg.PatternFlowLacpPartnerKey) (PatternFlowLacpPartnerKey, error)
	// FromPbText unmarshals PatternFlowLacpPartnerKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerKey from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerKey) Marshal() marshalPatternFlowLacpPartnerKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerKey) Unmarshal() unMarshalPatternFlowLacpPartnerKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerKey) ToProto() (*otg.PatternFlowLacpPartnerKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerKey) FromProto(msg *otg.PatternFlowLacpPartnerKey) (PatternFlowLacpPartnerKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerKey) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKey) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerKey) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKey) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerKey) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerKey) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerKey) Clone() (PatternFlowLacpPartnerKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerKey()
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

func (obj *patternFlowLacpPartnerKey) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
type PatternFlowLacpPartnerKey interface {
	Validation
	// msg marshals PatternFlowLacpPartnerKey to protobuf object *otg.PatternFlowLacpPartnerKey
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerKey
	// setMsg unmarshals PatternFlowLacpPartnerKey from protobuf object *otg.PatternFlowLacpPartnerKey
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerKey) PatternFlowLacpPartnerKey
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerKey
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerKey
	// validate validates PatternFlowLacpPartnerKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerKeyChoiceEnum, set in PatternFlowLacpPartnerKey
	Choice() PatternFlowLacpPartnerKeyChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerKeyChoiceEnum provided by user to PatternFlowLacpPartnerKey
	setChoice(value PatternFlowLacpPartnerKeyChoiceEnum) PatternFlowLacpPartnerKey
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerKey
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerKey.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerKey
	SetValue(value uint32) PatternFlowLacpPartnerKey
	// HasValue checks if Value has been set in PatternFlowLacpPartnerKey
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerKey.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerKey
	SetValues(value []uint32) PatternFlowLacpPartnerKey
	// Increment returns PatternFlowLacpPartnerKeyCounter, set in PatternFlowLacpPartnerKey.
	// PatternFlowLacpPartnerKeyCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerKeyCounter
	// SetIncrement assigns PatternFlowLacpPartnerKeyCounter provided by user to PatternFlowLacpPartnerKey.
	// PatternFlowLacpPartnerKeyCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKey
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerKey
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerKeyCounter, set in PatternFlowLacpPartnerKey.
	// PatternFlowLacpPartnerKeyCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerKeyCounter
	// SetDecrement assigns PatternFlowLacpPartnerKeyCounter provided by user to PatternFlowLacpPartnerKey.
	// PatternFlowLacpPartnerKeyCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKey
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerKey
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerKeyChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerKey
var PatternFlowLacpPartnerKeyChoice = struct {
	VALUE     PatternFlowLacpPartnerKeyChoiceEnum
	VALUES    PatternFlowLacpPartnerKeyChoiceEnum
	INCREMENT PatternFlowLacpPartnerKeyChoiceEnum
	DECREMENT PatternFlowLacpPartnerKeyChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerKeyChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerKeyChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerKeyChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerKeyChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerKey) Choice() PatternFlowLacpPartnerKeyChoiceEnum {
	return PatternFlowLacpPartnerKeyChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerKey) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerKey) setChoice(value PatternFlowLacpPartnerKeyChoiceEnum) PatternFlowLacpPartnerKey {
	intValue, ok := otg.PatternFlowLacpPartnerKey_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerKeyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerKey_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerKeyChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerKeyChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerKeyChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerKeyCounter().msg()
	}

	if value == PatternFlowLacpPartnerKeyChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerKeyCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerKey) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerKeyChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerKey) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerKey object
func (obj *patternFlowLacpPartnerKey) SetValue(value uint32) PatternFlowLacpPartnerKey {
	obj.setChoice(PatternFlowLacpPartnerKeyChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerKey) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerKey object
func (obj *patternFlowLacpPartnerKey) SetValues(value []uint32) PatternFlowLacpPartnerKey {
	obj.setChoice(PatternFlowLacpPartnerKeyChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerKeyCounter
func (obj *patternFlowLacpPartnerKey) Increment() PatternFlowLacpPartnerKeyCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerKeyChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerKeyCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerKeyCounter
func (obj *patternFlowLacpPartnerKey) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerKeyCounter value in the PatternFlowLacpPartnerKey object
func (obj *patternFlowLacpPartnerKey) SetIncrement(value PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKey {
	obj.setChoice(PatternFlowLacpPartnerKeyChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerKeyCounter
func (obj *patternFlowLacpPartnerKey) Decrement() PatternFlowLacpPartnerKeyCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerKeyChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerKeyCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerKeyCounter
func (obj *patternFlowLacpPartnerKey) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerKeyCounter value in the PatternFlowLacpPartnerKey object
func (obj *patternFlowLacpPartnerKey) SetDecrement(value PatternFlowLacpPartnerKeyCounter) PatternFlowLacpPartnerKey {
	obj.setChoice(PatternFlowLacpPartnerKeyChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerKey.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerKey.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpPartnerKey) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerKeyChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerKeyChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerKeyChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerKeyChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerKeyChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerKeyChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerKey")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerKey_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerKey_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
