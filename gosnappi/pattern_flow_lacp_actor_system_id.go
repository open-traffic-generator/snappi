package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorSystemId *****
type patternFlowLacpActorSystemId struct {
	validation
	obj             *otg.PatternFlowLacpActorSystemId
	marshaller      marshalPatternFlowLacpActorSystemId
	unMarshaller    unMarshalPatternFlowLacpActorSystemId
	incrementHolder PatternFlowLacpActorSystemIdCounter
	decrementHolder PatternFlowLacpActorSystemIdCounter
}

func NewPatternFlowLacpActorSystemId() PatternFlowLacpActorSystemId {
	obj := patternFlowLacpActorSystemId{obj: &otg.PatternFlowLacpActorSystemId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorSystemId) msg() *otg.PatternFlowLacpActorSystemId {
	return obj.obj
}

func (obj *patternFlowLacpActorSystemId) setMsg(msg *otg.PatternFlowLacpActorSystemId) PatternFlowLacpActorSystemId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorSystemId struct {
	obj *patternFlowLacpActorSystemId
}

type marshalPatternFlowLacpActorSystemId interface {
	// ToProto marshals PatternFlowLacpActorSystemId to protobuf object *otg.PatternFlowLacpActorSystemId
	ToProto() (*otg.PatternFlowLacpActorSystemId, error)
	// ToPbText marshals PatternFlowLacpActorSystemId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorSystemId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorSystemId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorSystemId struct {
	obj *patternFlowLacpActorSystemId
}

type unMarshalPatternFlowLacpActorSystemId interface {
	// FromProto unmarshals PatternFlowLacpActorSystemId from protobuf object *otg.PatternFlowLacpActorSystemId
	FromProto(msg *otg.PatternFlowLacpActorSystemId) (PatternFlowLacpActorSystemId, error)
	// FromPbText unmarshals PatternFlowLacpActorSystemId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorSystemId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorSystemId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorSystemId) Marshal() marshalPatternFlowLacpActorSystemId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorSystemId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorSystemId) Unmarshal() unMarshalPatternFlowLacpActorSystemId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorSystemId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorSystemId) ToProto() (*otg.PatternFlowLacpActorSystemId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorSystemId) FromProto(msg *otg.PatternFlowLacpActorSystemId) (PatternFlowLacpActorSystemId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorSystemId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemId) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemId) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorSystemId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorSystemId) FromJson(value string) error {
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

func (obj *patternFlowLacpActorSystemId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorSystemId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorSystemId) Clone() (PatternFlowLacpActorSystemId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorSystemId()
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

func (obj *patternFlowLacpActorSystemId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorSystemId is the Actor's System ID, which is a globally unique MAC address assigned to the system containing the Actor.
type PatternFlowLacpActorSystemId interface {
	Validation
	// msg marshals PatternFlowLacpActorSystemId to protobuf object *otg.PatternFlowLacpActorSystemId
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorSystemId
	// setMsg unmarshals PatternFlowLacpActorSystemId from protobuf object *otg.PatternFlowLacpActorSystemId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorSystemId) PatternFlowLacpActorSystemId
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorSystemId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorSystemId
	// validate validates PatternFlowLacpActorSystemId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorSystemId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorSystemIdChoiceEnum, set in PatternFlowLacpActorSystemId
	Choice() PatternFlowLacpActorSystemIdChoiceEnum
	// setChoice assigns PatternFlowLacpActorSystemIdChoiceEnum provided by user to PatternFlowLacpActorSystemId
	setChoice(value PatternFlowLacpActorSystemIdChoiceEnum) PatternFlowLacpActorSystemId
	// HasChoice checks if Choice has been set in PatternFlowLacpActorSystemId
	HasChoice() bool
	// Value returns string, set in PatternFlowLacpActorSystemId.
	Value() string
	// SetValue assigns string provided by user to PatternFlowLacpActorSystemId
	SetValue(value string) PatternFlowLacpActorSystemId
	// HasValue checks if Value has been set in PatternFlowLacpActorSystemId
	HasValue() bool
	// Values returns []string, set in PatternFlowLacpActorSystemId.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowLacpActorSystemId
	SetValues(value []string) PatternFlowLacpActorSystemId
	// Increment returns PatternFlowLacpActorSystemIdCounter, set in PatternFlowLacpActorSystemId.
	// PatternFlowLacpActorSystemIdCounter is mac counter pattern
	Increment() PatternFlowLacpActorSystemIdCounter
	// SetIncrement assigns PatternFlowLacpActorSystemIdCounter provided by user to PatternFlowLacpActorSystemId.
	// PatternFlowLacpActorSystemIdCounter is mac counter pattern
	SetIncrement(value PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemId
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorSystemId
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorSystemIdCounter, set in PatternFlowLacpActorSystemId.
	// PatternFlowLacpActorSystemIdCounter is mac counter pattern
	Decrement() PatternFlowLacpActorSystemIdCounter
	// SetDecrement assigns PatternFlowLacpActorSystemIdCounter provided by user to PatternFlowLacpActorSystemId.
	// PatternFlowLacpActorSystemIdCounter is mac counter pattern
	SetDecrement(value PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemId
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorSystemId
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorSystemIdChoiceEnum string

// Enum of Choice on PatternFlowLacpActorSystemId
var PatternFlowLacpActorSystemIdChoice = struct {
	VALUE     PatternFlowLacpActorSystemIdChoiceEnum
	VALUES    PatternFlowLacpActorSystemIdChoiceEnum
	INCREMENT PatternFlowLacpActorSystemIdChoiceEnum
	DECREMENT PatternFlowLacpActorSystemIdChoiceEnum
}{
	VALUE:     PatternFlowLacpActorSystemIdChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorSystemIdChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorSystemIdChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorSystemIdChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorSystemId) Choice() PatternFlowLacpActorSystemIdChoiceEnum {
	return PatternFlowLacpActorSystemIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorSystemId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorSystemId) setChoice(value PatternFlowLacpActorSystemIdChoiceEnum) PatternFlowLacpActorSystemId {
	intValue, ok := otg.PatternFlowLacpActorSystemId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorSystemIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorSystemId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorSystemIdChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorSystemIdChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorSystemIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorSystemIdCounter().msg()
	}

	if value == PatternFlowLacpActorSystemIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorSystemIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpActorSystemId) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorSystemIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpActorSystemId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowLacpActorSystemId object
func (obj *patternFlowLacpActorSystemId) SetValue(value string) PatternFlowLacpActorSystemId {
	obj.setChoice(PatternFlowLacpActorSystemIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowLacpActorSystemId) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowLacpActorSystemId object
func (obj *patternFlowLacpActorSystemId) SetValues(value []string) PatternFlowLacpActorSystemId {
	obj.setChoice(PatternFlowLacpActorSystemIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorSystemIdCounter
func (obj *patternFlowLacpActorSystemId) Increment() PatternFlowLacpActorSystemIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorSystemIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorSystemIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorSystemIdCounter
func (obj *patternFlowLacpActorSystemId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorSystemIdCounter value in the PatternFlowLacpActorSystemId object
func (obj *patternFlowLacpActorSystemId) SetIncrement(value PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemId {
	obj.setChoice(PatternFlowLacpActorSystemIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorSystemIdCounter
func (obj *patternFlowLacpActorSystemId) Decrement() PatternFlowLacpActorSystemIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorSystemIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorSystemIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorSystemIdCounter
func (obj *patternFlowLacpActorSystemId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorSystemIdCounter value in the PatternFlowLacpActorSystemId object
func (obj *patternFlowLacpActorSystemId) SetDecrement(value PatternFlowLacpActorSystemIdCounter) PatternFlowLacpActorSystemId {
	obj.setChoice(PatternFlowLacpActorSystemIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorSystemId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpActorSystemId.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpActorSystemId.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowLacpActorSystemId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorSystemIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorSystemIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorSystemIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorSystemIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorSystemId")
			}
		} else {
			intVal := otg.PatternFlowLacpActorSystemId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorSystemId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
