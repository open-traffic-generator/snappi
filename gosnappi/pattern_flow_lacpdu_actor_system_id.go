package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorSystemId *****
type patternFlowLacpduActorSystemId struct {
	validation
	obj             *otg.PatternFlowLacpduActorSystemId
	marshaller      marshalPatternFlowLacpduActorSystemId
	unMarshaller    unMarshalPatternFlowLacpduActorSystemId
	incrementHolder PatternFlowLacpduActorSystemIdCounter
	decrementHolder PatternFlowLacpduActorSystemIdCounter
}

func NewPatternFlowLacpduActorSystemId() PatternFlowLacpduActorSystemId {
	obj := patternFlowLacpduActorSystemId{obj: &otg.PatternFlowLacpduActorSystemId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorSystemId) msg() *otg.PatternFlowLacpduActorSystemId {
	return obj.obj
}

func (obj *patternFlowLacpduActorSystemId) setMsg(msg *otg.PatternFlowLacpduActorSystemId) PatternFlowLacpduActorSystemId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorSystemId struct {
	obj *patternFlowLacpduActorSystemId
}

type marshalPatternFlowLacpduActorSystemId interface {
	// ToProto marshals PatternFlowLacpduActorSystemId to protobuf object *otg.PatternFlowLacpduActorSystemId
	ToProto() (*otg.PatternFlowLacpduActorSystemId, error)
	// ToPbText marshals PatternFlowLacpduActorSystemId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorSystemId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorSystemId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorSystemId struct {
	obj *patternFlowLacpduActorSystemId
}

type unMarshalPatternFlowLacpduActorSystemId interface {
	// FromProto unmarshals PatternFlowLacpduActorSystemId from protobuf object *otg.PatternFlowLacpduActorSystemId
	FromProto(msg *otg.PatternFlowLacpduActorSystemId) (PatternFlowLacpduActorSystemId, error)
	// FromPbText unmarshals PatternFlowLacpduActorSystemId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorSystemId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorSystemId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorSystemId) Marshal() marshalPatternFlowLacpduActorSystemId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorSystemId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorSystemId) Unmarshal() unMarshalPatternFlowLacpduActorSystemId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorSystemId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorSystemId) ToProto() (*otg.PatternFlowLacpduActorSystemId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorSystemId) FromProto(msg *otg.PatternFlowLacpduActorSystemId) (PatternFlowLacpduActorSystemId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorSystemId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemId) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemId) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorSystemId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorSystemId) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorSystemId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorSystemId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorSystemId) Clone() (PatternFlowLacpduActorSystemId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorSystemId()
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

func (obj *patternFlowLacpduActorSystemId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
type PatternFlowLacpduActorSystemId interface {
	Validation
	// msg marshals PatternFlowLacpduActorSystemId to protobuf object *otg.PatternFlowLacpduActorSystemId
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorSystemId
	// setMsg unmarshals PatternFlowLacpduActorSystemId from protobuf object *otg.PatternFlowLacpduActorSystemId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorSystemId) PatternFlowLacpduActorSystemId
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorSystemId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorSystemId
	// validate validates PatternFlowLacpduActorSystemId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorSystemId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorSystemIdChoiceEnum, set in PatternFlowLacpduActorSystemId
	Choice() PatternFlowLacpduActorSystemIdChoiceEnum
	// setChoice assigns PatternFlowLacpduActorSystemIdChoiceEnum provided by user to PatternFlowLacpduActorSystemId
	setChoice(value PatternFlowLacpduActorSystemIdChoiceEnum) PatternFlowLacpduActorSystemId
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorSystemId
	HasChoice() bool
	// Value returns string, set in PatternFlowLacpduActorSystemId.
	Value() string
	// SetValue assigns string provided by user to PatternFlowLacpduActorSystemId
	SetValue(value string) PatternFlowLacpduActorSystemId
	// HasValue checks if Value has been set in PatternFlowLacpduActorSystemId
	HasValue() bool
	// Values returns []string, set in PatternFlowLacpduActorSystemId.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowLacpduActorSystemId
	SetValues(value []string) PatternFlowLacpduActorSystemId
	// Increment returns PatternFlowLacpduActorSystemIdCounter, set in PatternFlowLacpduActorSystemId.
	// PatternFlowLacpduActorSystemIdCounter is mac counter pattern
	Increment() PatternFlowLacpduActorSystemIdCounter
	// SetIncrement assigns PatternFlowLacpduActorSystemIdCounter provided by user to PatternFlowLacpduActorSystemId.
	// PatternFlowLacpduActorSystemIdCounter is mac counter pattern
	SetIncrement(value PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemId
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorSystemId
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorSystemIdCounter, set in PatternFlowLacpduActorSystemId.
	// PatternFlowLacpduActorSystemIdCounter is mac counter pattern
	Decrement() PatternFlowLacpduActorSystemIdCounter
	// SetDecrement assigns PatternFlowLacpduActorSystemIdCounter provided by user to PatternFlowLacpduActorSystemId.
	// PatternFlowLacpduActorSystemIdCounter is mac counter pattern
	SetDecrement(value PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemId
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorSystemId
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorSystemIdChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorSystemId
var PatternFlowLacpduActorSystemIdChoice = struct {
	VALUE     PatternFlowLacpduActorSystemIdChoiceEnum
	VALUES    PatternFlowLacpduActorSystemIdChoiceEnum
	INCREMENT PatternFlowLacpduActorSystemIdChoiceEnum
	DECREMENT PatternFlowLacpduActorSystemIdChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorSystemIdChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorSystemIdChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorSystemIdChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorSystemIdChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorSystemId) Choice() PatternFlowLacpduActorSystemIdChoiceEnum {
	return PatternFlowLacpduActorSystemIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorSystemId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorSystemId) setChoice(value PatternFlowLacpduActorSystemIdChoiceEnum) PatternFlowLacpduActorSystemId {
	intValue, ok := otg.PatternFlowLacpduActorSystemId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorSystemIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorSystemId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorSystemIdChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorSystemIdChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorSystemIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorSystemIdCounter().msg()
	}

	if value == PatternFlowLacpduActorSystemIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorSystemIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpduActorSystemId) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorSystemIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpduActorSystemId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowLacpduActorSystemId object
func (obj *patternFlowLacpduActorSystemId) SetValue(value string) PatternFlowLacpduActorSystemId {
	obj.setChoice(PatternFlowLacpduActorSystemIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowLacpduActorSystemId) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowLacpduActorSystemId object
func (obj *patternFlowLacpduActorSystemId) SetValues(value []string) PatternFlowLacpduActorSystemId {
	obj.setChoice(PatternFlowLacpduActorSystemIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorSystemIdCounter
func (obj *patternFlowLacpduActorSystemId) Increment() PatternFlowLacpduActorSystemIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorSystemIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorSystemIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorSystemIdCounter
func (obj *patternFlowLacpduActorSystemId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorSystemIdCounter value in the PatternFlowLacpduActorSystemId object
func (obj *patternFlowLacpduActorSystemId) SetIncrement(value PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemId {
	obj.setChoice(PatternFlowLacpduActorSystemIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorSystemIdCounter
func (obj *patternFlowLacpduActorSystemId) Decrement() PatternFlowLacpduActorSystemIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorSystemIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorSystemIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorSystemIdCounter
func (obj *patternFlowLacpduActorSystemId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorSystemIdCounter value in the PatternFlowLacpduActorSystemId object
func (obj *patternFlowLacpduActorSystemId) SetDecrement(value PatternFlowLacpduActorSystemIdCounter) PatternFlowLacpduActorSystemId {
	obj.setChoice(PatternFlowLacpduActorSystemIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorSystemId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduActorSystemId.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduActorSystemId.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowLacpduActorSystemId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorSystemIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorSystemIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorSystemIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorSystemId")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorSystemId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorSystemId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
