package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorKey *****
type patternFlowLacpActorKey struct {
	validation
	obj             *otg.PatternFlowLacpActorKey
	marshaller      marshalPatternFlowLacpActorKey
	unMarshaller    unMarshalPatternFlowLacpActorKey
	incrementHolder PatternFlowLacpActorKeyCounter
	decrementHolder PatternFlowLacpActorKeyCounter
}

func NewPatternFlowLacpActorKey() PatternFlowLacpActorKey {
	obj := patternFlowLacpActorKey{obj: &otg.PatternFlowLacpActorKey{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorKey) msg() *otg.PatternFlowLacpActorKey {
	return obj.obj
}

func (obj *patternFlowLacpActorKey) setMsg(msg *otg.PatternFlowLacpActorKey) PatternFlowLacpActorKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorKey struct {
	obj *patternFlowLacpActorKey
}

type marshalPatternFlowLacpActorKey interface {
	// ToProto marshals PatternFlowLacpActorKey to protobuf object *otg.PatternFlowLacpActorKey
	ToProto() (*otg.PatternFlowLacpActorKey, error)
	// ToPbText marshals PatternFlowLacpActorKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorKey to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorKey struct {
	obj *patternFlowLacpActorKey
}

type unMarshalPatternFlowLacpActorKey interface {
	// FromProto unmarshals PatternFlowLacpActorKey from protobuf object *otg.PatternFlowLacpActorKey
	FromProto(msg *otg.PatternFlowLacpActorKey) (PatternFlowLacpActorKey, error)
	// FromPbText unmarshals PatternFlowLacpActorKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorKey from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorKey) Marshal() marshalPatternFlowLacpActorKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorKey) Unmarshal() unMarshalPatternFlowLacpActorKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorKey) ToProto() (*otg.PatternFlowLacpActorKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorKey) FromProto(msg *otg.PatternFlowLacpActorKey) (PatternFlowLacpActorKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorKey) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKey) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorKey) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKey) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorKey) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorKey) FromJson(value string) error {
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

func (obj *patternFlowLacpActorKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorKey) Clone() (PatternFlowLacpActorKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorKey()
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

func (obj *patternFlowLacpActorKey) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorKey is the operational Key value assigned to the Actor's port. The key is generated based on port configuration (e.g., speed, duplex, trunk ID) and is used to identify potential aggregation groups. Only links with the same key can be aggregated together.
type PatternFlowLacpActorKey interface {
	Validation
	// msg marshals PatternFlowLacpActorKey to protobuf object *otg.PatternFlowLacpActorKey
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorKey
	// setMsg unmarshals PatternFlowLacpActorKey from protobuf object *otg.PatternFlowLacpActorKey
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorKey) PatternFlowLacpActorKey
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorKey
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorKey
	// validate validates PatternFlowLacpActorKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorKeyChoiceEnum, set in PatternFlowLacpActorKey
	Choice() PatternFlowLacpActorKeyChoiceEnum
	// setChoice assigns PatternFlowLacpActorKeyChoiceEnum provided by user to PatternFlowLacpActorKey
	setChoice(value PatternFlowLacpActorKeyChoiceEnum) PatternFlowLacpActorKey
	// HasChoice checks if Choice has been set in PatternFlowLacpActorKey
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorKey.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorKey
	SetValue(value uint32) PatternFlowLacpActorKey
	// HasValue checks if Value has been set in PatternFlowLacpActorKey
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorKey.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorKey
	SetValues(value []uint32) PatternFlowLacpActorKey
	// Increment returns PatternFlowLacpActorKeyCounter, set in PatternFlowLacpActorKey.
	// PatternFlowLacpActorKeyCounter is integer counter pattern
	Increment() PatternFlowLacpActorKeyCounter
	// SetIncrement assigns PatternFlowLacpActorKeyCounter provided by user to PatternFlowLacpActorKey.
	// PatternFlowLacpActorKeyCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKey
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorKey
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorKeyCounter, set in PatternFlowLacpActorKey.
	// PatternFlowLacpActorKeyCounter is integer counter pattern
	Decrement() PatternFlowLacpActorKeyCounter
	// SetDecrement assigns PatternFlowLacpActorKeyCounter provided by user to PatternFlowLacpActorKey.
	// PatternFlowLacpActorKeyCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKey
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorKey
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorKeyChoiceEnum string

// Enum of Choice on PatternFlowLacpActorKey
var PatternFlowLacpActorKeyChoice = struct {
	VALUE     PatternFlowLacpActorKeyChoiceEnum
	VALUES    PatternFlowLacpActorKeyChoiceEnum
	INCREMENT PatternFlowLacpActorKeyChoiceEnum
	DECREMENT PatternFlowLacpActorKeyChoiceEnum
}{
	VALUE:     PatternFlowLacpActorKeyChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorKeyChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorKeyChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorKeyChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorKey) Choice() PatternFlowLacpActorKeyChoiceEnum {
	return PatternFlowLacpActorKeyChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorKey) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorKey) setChoice(value PatternFlowLacpActorKeyChoiceEnum) PatternFlowLacpActorKey {
	intValue, ok := otg.PatternFlowLacpActorKey_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorKeyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorKey_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorKeyChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorKeyChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorKeyChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorKeyCounter().msg()
	}

	if value == PatternFlowLacpActorKeyChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorKeyCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorKey) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorKeyChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorKey) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorKey object
func (obj *patternFlowLacpActorKey) SetValue(value uint32) PatternFlowLacpActorKey {
	obj.setChoice(PatternFlowLacpActorKeyChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorKey) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorKey object
func (obj *patternFlowLacpActorKey) SetValues(value []uint32) PatternFlowLacpActorKey {
	obj.setChoice(PatternFlowLacpActorKeyChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorKeyCounter
func (obj *patternFlowLacpActorKey) Increment() PatternFlowLacpActorKeyCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorKeyChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorKeyCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorKeyCounter
func (obj *patternFlowLacpActorKey) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorKeyCounter value in the PatternFlowLacpActorKey object
func (obj *patternFlowLacpActorKey) SetIncrement(value PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKey {
	obj.setChoice(PatternFlowLacpActorKeyChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorKeyCounter
func (obj *patternFlowLacpActorKey) Decrement() PatternFlowLacpActorKeyCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorKeyChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorKeyCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorKeyCounter
func (obj *patternFlowLacpActorKey) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorKeyCounter value in the PatternFlowLacpActorKey object
func (obj *patternFlowLacpActorKey) SetDecrement(value PatternFlowLacpActorKeyCounter) PatternFlowLacpActorKey {
	obj.setChoice(PatternFlowLacpActorKeyChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorKey.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorKey.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpActorKey) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorKeyChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorKeyChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorKeyChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorKeyChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorKeyChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorKeyChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorKey")
			}
		} else {
			intVal := otg.PatternFlowLacpActorKey_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorKey_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
