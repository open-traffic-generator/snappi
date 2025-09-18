package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorKey *****
type patternFlowLacpduActorKey struct {
	validation
	obj             *otg.PatternFlowLacpduActorKey
	marshaller      marshalPatternFlowLacpduActorKey
	unMarshaller    unMarshalPatternFlowLacpduActorKey
	incrementHolder PatternFlowLacpduActorKeyCounter
	decrementHolder PatternFlowLacpduActorKeyCounter
}

func NewPatternFlowLacpduActorKey() PatternFlowLacpduActorKey {
	obj := patternFlowLacpduActorKey{obj: &otg.PatternFlowLacpduActorKey{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorKey) msg() *otg.PatternFlowLacpduActorKey {
	return obj.obj
}

func (obj *patternFlowLacpduActorKey) setMsg(msg *otg.PatternFlowLacpduActorKey) PatternFlowLacpduActorKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorKey struct {
	obj *patternFlowLacpduActorKey
}

type marshalPatternFlowLacpduActorKey interface {
	// ToProto marshals PatternFlowLacpduActorKey to protobuf object *otg.PatternFlowLacpduActorKey
	ToProto() (*otg.PatternFlowLacpduActorKey, error)
	// ToPbText marshals PatternFlowLacpduActorKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorKey to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorKey struct {
	obj *patternFlowLacpduActorKey
}

type unMarshalPatternFlowLacpduActorKey interface {
	// FromProto unmarshals PatternFlowLacpduActorKey from protobuf object *otg.PatternFlowLacpduActorKey
	FromProto(msg *otg.PatternFlowLacpduActorKey) (PatternFlowLacpduActorKey, error)
	// FromPbText unmarshals PatternFlowLacpduActorKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorKey from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorKey) Marshal() marshalPatternFlowLacpduActorKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorKey) Unmarshal() unMarshalPatternFlowLacpduActorKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorKey) ToProto() (*otg.PatternFlowLacpduActorKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorKey) FromProto(msg *otg.PatternFlowLacpduActorKey) (PatternFlowLacpduActorKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorKey) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKey) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorKey) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKey) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorKey) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorKey) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorKey) Clone() (PatternFlowLacpduActorKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorKey()
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

func (obj *patternFlowLacpduActorKey) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
type PatternFlowLacpduActorKey interface {
	Validation
	// msg marshals PatternFlowLacpduActorKey to protobuf object *otg.PatternFlowLacpduActorKey
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorKey
	// setMsg unmarshals PatternFlowLacpduActorKey from protobuf object *otg.PatternFlowLacpduActorKey
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorKey) PatternFlowLacpduActorKey
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorKey
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorKey
	// validate validates PatternFlowLacpduActorKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorKeyChoiceEnum, set in PatternFlowLacpduActorKey
	Choice() PatternFlowLacpduActorKeyChoiceEnum
	// setChoice assigns PatternFlowLacpduActorKeyChoiceEnum provided by user to PatternFlowLacpduActorKey
	setChoice(value PatternFlowLacpduActorKeyChoiceEnum) PatternFlowLacpduActorKey
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorKey
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorKey.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorKey
	SetValue(value uint32) PatternFlowLacpduActorKey
	// HasValue checks if Value has been set in PatternFlowLacpduActorKey
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorKey.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorKey
	SetValues(value []uint32) PatternFlowLacpduActorKey
	// Increment returns PatternFlowLacpduActorKeyCounter, set in PatternFlowLacpduActorKey.
	// PatternFlowLacpduActorKeyCounter is integer counter pattern
	Increment() PatternFlowLacpduActorKeyCounter
	// SetIncrement assigns PatternFlowLacpduActorKeyCounter provided by user to PatternFlowLacpduActorKey.
	// PatternFlowLacpduActorKeyCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKey
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorKey
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorKeyCounter, set in PatternFlowLacpduActorKey.
	// PatternFlowLacpduActorKeyCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorKeyCounter
	// SetDecrement assigns PatternFlowLacpduActorKeyCounter provided by user to PatternFlowLacpduActorKey.
	// PatternFlowLacpduActorKeyCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKey
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorKey
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorKeyChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorKey
var PatternFlowLacpduActorKeyChoice = struct {
	VALUE     PatternFlowLacpduActorKeyChoiceEnum
	VALUES    PatternFlowLacpduActorKeyChoiceEnum
	INCREMENT PatternFlowLacpduActorKeyChoiceEnum
	DECREMENT PatternFlowLacpduActorKeyChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorKeyChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorKeyChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorKeyChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorKeyChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorKey) Choice() PatternFlowLacpduActorKeyChoiceEnum {
	return PatternFlowLacpduActorKeyChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorKey) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorKey) setChoice(value PatternFlowLacpduActorKeyChoiceEnum) PatternFlowLacpduActorKey {
	intValue, ok := otg.PatternFlowLacpduActorKey_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorKeyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorKey_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorKeyChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorKeyChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorKeyChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorKeyCounter().msg()
	}

	if value == PatternFlowLacpduActorKeyChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorKeyCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorKey) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorKeyChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorKey) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorKey object
func (obj *patternFlowLacpduActorKey) SetValue(value uint32) PatternFlowLacpduActorKey {
	obj.setChoice(PatternFlowLacpduActorKeyChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorKey) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorKey object
func (obj *patternFlowLacpduActorKey) SetValues(value []uint32) PatternFlowLacpduActorKey {
	obj.setChoice(PatternFlowLacpduActorKeyChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorKeyCounter
func (obj *patternFlowLacpduActorKey) Increment() PatternFlowLacpduActorKeyCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorKeyChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorKeyCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorKeyCounter
func (obj *patternFlowLacpduActorKey) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorKeyCounter value in the PatternFlowLacpduActorKey object
func (obj *patternFlowLacpduActorKey) SetIncrement(value PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKey {
	obj.setChoice(PatternFlowLacpduActorKeyChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorKeyCounter
func (obj *patternFlowLacpduActorKey) Decrement() PatternFlowLacpduActorKeyCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorKeyChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorKeyCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorKeyCounter
func (obj *patternFlowLacpduActorKey) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorKeyCounter value in the PatternFlowLacpduActorKey object
func (obj *patternFlowLacpduActorKey) SetDecrement(value PatternFlowLacpduActorKeyCounter) PatternFlowLacpduActorKey {
	obj.setChoice(PatternFlowLacpduActorKeyChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorKey.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorKey.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduActorKey) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorKeyChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorKeyChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorKeyChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorKeyChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorKeyChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorKeyChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorKey")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorKey_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorKey_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
