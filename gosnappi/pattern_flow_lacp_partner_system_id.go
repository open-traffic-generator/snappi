package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerSystemId *****
type patternFlowLacpPartnerSystemId struct {
	validation
	obj             *otg.PatternFlowLacpPartnerSystemId
	marshaller      marshalPatternFlowLacpPartnerSystemId
	unMarshaller    unMarshalPatternFlowLacpPartnerSystemId
	incrementHolder PatternFlowLacpPartnerSystemIdCounter
	decrementHolder PatternFlowLacpPartnerSystemIdCounter
}

func NewPatternFlowLacpPartnerSystemId() PatternFlowLacpPartnerSystemId {
	obj := patternFlowLacpPartnerSystemId{obj: &otg.PatternFlowLacpPartnerSystemId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerSystemId) msg() *otg.PatternFlowLacpPartnerSystemId {
	return obj.obj
}

func (obj *patternFlowLacpPartnerSystemId) setMsg(msg *otg.PatternFlowLacpPartnerSystemId) PatternFlowLacpPartnerSystemId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerSystemId struct {
	obj *patternFlowLacpPartnerSystemId
}

type marshalPatternFlowLacpPartnerSystemId interface {
	// ToProto marshals PatternFlowLacpPartnerSystemId to protobuf object *otg.PatternFlowLacpPartnerSystemId
	ToProto() (*otg.PatternFlowLacpPartnerSystemId, error)
	// ToPbText marshals PatternFlowLacpPartnerSystemId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerSystemId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerSystemId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerSystemId struct {
	obj *patternFlowLacpPartnerSystemId
}

type unMarshalPatternFlowLacpPartnerSystemId interface {
	// FromProto unmarshals PatternFlowLacpPartnerSystemId from protobuf object *otg.PatternFlowLacpPartnerSystemId
	FromProto(msg *otg.PatternFlowLacpPartnerSystemId) (PatternFlowLacpPartnerSystemId, error)
	// FromPbText unmarshals PatternFlowLacpPartnerSystemId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerSystemId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerSystemId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerSystemId) Marshal() marshalPatternFlowLacpPartnerSystemId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerSystemId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerSystemId) Unmarshal() unMarshalPatternFlowLacpPartnerSystemId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerSystemId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerSystemId) ToProto() (*otg.PatternFlowLacpPartnerSystemId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerSystemId) FromProto(msg *otg.PatternFlowLacpPartnerSystemId) (PatternFlowLacpPartnerSystemId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerSystemId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemId) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemId) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerSystemId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerSystemId) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerSystemId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerSystemId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerSystemId) Clone() (PatternFlowLacpPartnerSystemId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerSystemId()
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

func (obj *patternFlowLacpPartnerSystemId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
type PatternFlowLacpPartnerSystemId interface {
	Validation
	// msg marshals PatternFlowLacpPartnerSystemId to protobuf object *otg.PatternFlowLacpPartnerSystemId
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerSystemId
	// setMsg unmarshals PatternFlowLacpPartnerSystemId from protobuf object *otg.PatternFlowLacpPartnerSystemId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerSystemId) PatternFlowLacpPartnerSystemId
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerSystemId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerSystemId
	// validate validates PatternFlowLacpPartnerSystemId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerSystemId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerSystemIdChoiceEnum, set in PatternFlowLacpPartnerSystemId
	Choice() PatternFlowLacpPartnerSystemIdChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerSystemIdChoiceEnum provided by user to PatternFlowLacpPartnerSystemId
	setChoice(value PatternFlowLacpPartnerSystemIdChoiceEnum) PatternFlowLacpPartnerSystemId
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerSystemId
	HasChoice() bool
	// Value returns string, set in PatternFlowLacpPartnerSystemId.
	Value() string
	// SetValue assigns string provided by user to PatternFlowLacpPartnerSystemId
	SetValue(value string) PatternFlowLacpPartnerSystemId
	// HasValue checks if Value has been set in PatternFlowLacpPartnerSystemId
	HasValue() bool
	// Values returns []string, set in PatternFlowLacpPartnerSystemId.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowLacpPartnerSystemId
	SetValues(value []string) PatternFlowLacpPartnerSystemId
	// Increment returns PatternFlowLacpPartnerSystemIdCounter, set in PatternFlowLacpPartnerSystemId.
	// PatternFlowLacpPartnerSystemIdCounter is mac counter pattern
	Increment() PatternFlowLacpPartnerSystemIdCounter
	// SetIncrement assigns PatternFlowLacpPartnerSystemIdCounter provided by user to PatternFlowLacpPartnerSystemId.
	// PatternFlowLacpPartnerSystemIdCounter is mac counter pattern
	SetIncrement(value PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemId
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerSystemId
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerSystemIdCounter, set in PatternFlowLacpPartnerSystemId.
	// PatternFlowLacpPartnerSystemIdCounter is mac counter pattern
	Decrement() PatternFlowLacpPartnerSystemIdCounter
	// SetDecrement assigns PatternFlowLacpPartnerSystemIdCounter provided by user to PatternFlowLacpPartnerSystemId.
	// PatternFlowLacpPartnerSystemIdCounter is mac counter pattern
	SetDecrement(value PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemId
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerSystemId
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerSystemIdChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerSystemId
var PatternFlowLacpPartnerSystemIdChoice = struct {
	VALUE     PatternFlowLacpPartnerSystemIdChoiceEnum
	VALUES    PatternFlowLacpPartnerSystemIdChoiceEnum
	INCREMENT PatternFlowLacpPartnerSystemIdChoiceEnum
	DECREMENT PatternFlowLacpPartnerSystemIdChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerSystemIdChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerSystemIdChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerSystemIdChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerSystemIdChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerSystemId) Choice() PatternFlowLacpPartnerSystemIdChoiceEnum {
	return PatternFlowLacpPartnerSystemIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerSystemId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerSystemId) setChoice(value PatternFlowLacpPartnerSystemIdChoiceEnum) PatternFlowLacpPartnerSystemId {
	intValue, ok := otg.PatternFlowLacpPartnerSystemId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerSystemIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerSystemId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerSystemIdChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerSystemIdChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerSystemIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerSystemIdCounter().msg()
	}

	if value == PatternFlowLacpPartnerSystemIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerSystemIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpPartnerSystemId) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpPartnerSystemId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowLacpPartnerSystemId object
func (obj *patternFlowLacpPartnerSystemId) SetValue(value string) PatternFlowLacpPartnerSystemId {
	obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowLacpPartnerSystemId) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowLacpPartnerSystemId object
func (obj *patternFlowLacpPartnerSystemId) SetValues(value []string) PatternFlowLacpPartnerSystemId {
	obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerSystemIdCounter
func (obj *patternFlowLacpPartnerSystemId) Increment() PatternFlowLacpPartnerSystemIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerSystemIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerSystemIdCounter
func (obj *patternFlowLacpPartnerSystemId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerSystemIdCounter value in the PatternFlowLacpPartnerSystemId object
func (obj *patternFlowLacpPartnerSystemId) SetIncrement(value PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemId {
	obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerSystemIdCounter
func (obj *patternFlowLacpPartnerSystemId) Decrement() PatternFlowLacpPartnerSystemIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerSystemIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerSystemIdCounter
func (obj *patternFlowLacpPartnerSystemId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerSystemIdCounter value in the PatternFlowLacpPartnerSystemId object
func (obj *patternFlowLacpPartnerSystemId) SetDecrement(value PatternFlowLacpPartnerSystemIdCounter) PatternFlowLacpPartnerSystemId {
	obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerSystemId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpPartnerSystemId.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpPartnerSystemId.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowLacpPartnerSystemId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerSystemIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerSystemIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerSystemIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerSystemId")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerSystemId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerSystemId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
