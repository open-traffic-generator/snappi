package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerKey *****
type patternFlowLacpduPartnerKey struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerKey
	marshaller      marshalPatternFlowLacpduPartnerKey
	unMarshaller    unMarshalPatternFlowLacpduPartnerKey
	incrementHolder PatternFlowLacpduPartnerKeyCounter
	decrementHolder PatternFlowLacpduPartnerKeyCounter
}

func NewPatternFlowLacpduPartnerKey() PatternFlowLacpduPartnerKey {
	obj := patternFlowLacpduPartnerKey{obj: &otg.PatternFlowLacpduPartnerKey{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerKey) msg() *otg.PatternFlowLacpduPartnerKey {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerKey) setMsg(msg *otg.PatternFlowLacpduPartnerKey) PatternFlowLacpduPartnerKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerKey struct {
	obj *patternFlowLacpduPartnerKey
}

type marshalPatternFlowLacpduPartnerKey interface {
	// ToProto marshals PatternFlowLacpduPartnerKey to protobuf object *otg.PatternFlowLacpduPartnerKey
	ToProto() (*otg.PatternFlowLacpduPartnerKey, error)
	// ToPbText marshals PatternFlowLacpduPartnerKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerKey to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerKey struct {
	obj *patternFlowLacpduPartnerKey
}

type unMarshalPatternFlowLacpduPartnerKey interface {
	// FromProto unmarshals PatternFlowLacpduPartnerKey from protobuf object *otg.PatternFlowLacpduPartnerKey
	FromProto(msg *otg.PatternFlowLacpduPartnerKey) (PatternFlowLacpduPartnerKey, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerKey from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerKey) Marshal() marshalPatternFlowLacpduPartnerKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerKey) Unmarshal() unMarshalPatternFlowLacpduPartnerKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerKey) ToProto() (*otg.PatternFlowLacpduPartnerKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerKey) FromProto(msg *otg.PatternFlowLacpduPartnerKey) (PatternFlowLacpduPartnerKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerKey) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKey) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerKey) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKey) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerKey) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerKey) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerKey) Clone() (PatternFlowLacpduPartnerKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerKey()
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

func (obj *patternFlowLacpduPartnerKey) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerKey is the operational Key value of the Partner's port, as received by the Actor.
type PatternFlowLacpduPartnerKey interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerKey to protobuf object *otg.PatternFlowLacpduPartnerKey
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerKey
	// setMsg unmarshals PatternFlowLacpduPartnerKey from protobuf object *otg.PatternFlowLacpduPartnerKey
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerKey) PatternFlowLacpduPartnerKey
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerKey
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerKey
	// validate validates PatternFlowLacpduPartnerKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerKeyChoiceEnum, set in PatternFlowLacpduPartnerKey
	Choice() PatternFlowLacpduPartnerKeyChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerKeyChoiceEnum provided by user to PatternFlowLacpduPartnerKey
	setChoice(value PatternFlowLacpduPartnerKeyChoiceEnum) PatternFlowLacpduPartnerKey
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerKey
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerKey.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerKey
	SetValue(value uint32) PatternFlowLacpduPartnerKey
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerKey
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerKey.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerKey
	SetValues(value []uint32) PatternFlowLacpduPartnerKey
	// Increment returns PatternFlowLacpduPartnerKeyCounter, set in PatternFlowLacpduPartnerKey.
	// PatternFlowLacpduPartnerKeyCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerKeyCounter
	// SetIncrement assigns PatternFlowLacpduPartnerKeyCounter provided by user to PatternFlowLacpduPartnerKey.
	// PatternFlowLacpduPartnerKeyCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKey
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerKey
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerKeyCounter, set in PatternFlowLacpduPartnerKey.
	// PatternFlowLacpduPartnerKeyCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerKeyCounter
	// SetDecrement assigns PatternFlowLacpduPartnerKeyCounter provided by user to PatternFlowLacpduPartnerKey.
	// PatternFlowLacpduPartnerKeyCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKey
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerKey
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerKeyChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerKey
var PatternFlowLacpduPartnerKeyChoice = struct {
	VALUE     PatternFlowLacpduPartnerKeyChoiceEnum
	VALUES    PatternFlowLacpduPartnerKeyChoiceEnum
	INCREMENT PatternFlowLacpduPartnerKeyChoiceEnum
	DECREMENT PatternFlowLacpduPartnerKeyChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerKeyChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerKeyChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerKeyChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerKeyChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerKey) Choice() PatternFlowLacpduPartnerKeyChoiceEnum {
	return PatternFlowLacpduPartnerKeyChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerKey) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerKey) setChoice(value PatternFlowLacpduPartnerKeyChoiceEnum) PatternFlowLacpduPartnerKey {
	intValue, ok := otg.PatternFlowLacpduPartnerKey_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerKeyChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerKey_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerKeyChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerKeyChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerKeyChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerKeyCounter().msg()
	}

	if value == PatternFlowLacpduPartnerKeyChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerKeyCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerKey) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerKeyChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerKey) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerKey object
func (obj *patternFlowLacpduPartnerKey) SetValue(value uint32) PatternFlowLacpduPartnerKey {
	obj.setChoice(PatternFlowLacpduPartnerKeyChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerKey) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerKey object
func (obj *patternFlowLacpduPartnerKey) SetValues(value []uint32) PatternFlowLacpduPartnerKey {
	obj.setChoice(PatternFlowLacpduPartnerKeyChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerKeyCounter
func (obj *patternFlowLacpduPartnerKey) Increment() PatternFlowLacpduPartnerKeyCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerKeyChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerKeyCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerKeyCounter
func (obj *patternFlowLacpduPartnerKey) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerKeyCounter value in the PatternFlowLacpduPartnerKey object
func (obj *patternFlowLacpduPartnerKey) SetIncrement(value PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKey {
	obj.setChoice(PatternFlowLacpduPartnerKeyChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerKeyCounter
func (obj *patternFlowLacpduPartnerKey) Decrement() PatternFlowLacpduPartnerKeyCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerKeyChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerKeyCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerKeyCounter
func (obj *patternFlowLacpduPartnerKey) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerKeyCounter value in the PatternFlowLacpduPartnerKey object
func (obj *patternFlowLacpduPartnerKey) SetDecrement(value PatternFlowLacpduPartnerKeyCounter) PatternFlowLacpduPartnerKey {
	obj.setChoice(PatternFlowLacpduPartnerKeyChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerKey.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerKey.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerKey) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerKeyChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerKeyChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerKeyChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerKeyChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerKeyChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerKeyChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerKey")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerKey_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerKey_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
