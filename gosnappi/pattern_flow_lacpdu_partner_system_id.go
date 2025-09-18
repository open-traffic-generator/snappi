package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerSystemId *****
type patternFlowLacpduPartnerSystemId struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerSystemId
	marshaller      marshalPatternFlowLacpduPartnerSystemId
	unMarshaller    unMarshalPatternFlowLacpduPartnerSystemId
	incrementHolder PatternFlowLacpduPartnerSystemIdCounter
	decrementHolder PatternFlowLacpduPartnerSystemIdCounter
}

func NewPatternFlowLacpduPartnerSystemId() PatternFlowLacpduPartnerSystemId {
	obj := patternFlowLacpduPartnerSystemId{obj: &otg.PatternFlowLacpduPartnerSystemId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerSystemId) msg() *otg.PatternFlowLacpduPartnerSystemId {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerSystemId) setMsg(msg *otg.PatternFlowLacpduPartnerSystemId) PatternFlowLacpduPartnerSystemId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerSystemId struct {
	obj *patternFlowLacpduPartnerSystemId
}

type marshalPatternFlowLacpduPartnerSystemId interface {
	// ToProto marshals PatternFlowLacpduPartnerSystemId to protobuf object *otg.PatternFlowLacpduPartnerSystemId
	ToProto() (*otg.PatternFlowLacpduPartnerSystemId, error)
	// ToPbText marshals PatternFlowLacpduPartnerSystemId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerSystemId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerSystemId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerSystemId struct {
	obj *patternFlowLacpduPartnerSystemId
}

type unMarshalPatternFlowLacpduPartnerSystemId interface {
	// FromProto unmarshals PatternFlowLacpduPartnerSystemId from protobuf object *otg.PatternFlowLacpduPartnerSystemId
	FromProto(msg *otg.PatternFlowLacpduPartnerSystemId) (PatternFlowLacpduPartnerSystemId, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerSystemId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerSystemId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerSystemId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerSystemId) Marshal() marshalPatternFlowLacpduPartnerSystemId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerSystemId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerSystemId) Unmarshal() unMarshalPatternFlowLacpduPartnerSystemId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerSystemId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerSystemId) ToProto() (*otg.PatternFlowLacpduPartnerSystemId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerSystemId) FromProto(msg *otg.PatternFlowLacpduPartnerSystemId) (PatternFlowLacpduPartnerSystemId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerSystemId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemId) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemId) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerSystemId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerSystemId) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerSystemId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerSystemId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerSystemId) Clone() (PatternFlowLacpduPartnerSystemId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerSystemId()
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

func (obj *patternFlowLacpduPartnerSystemId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerSystemId is the Partner's System ID (MAC address), as received by the Actor.
type PatternFlowLacpduPartnerSystemId interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerSystemId to protobuf object *otg.PatternFlowLacpduPartnerSystemId
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerSystemId
	// setMsg unmarshals PatternFlowLacpduPartnerSystemId from protobuf object *otg.PatternFlowLacpduPartnerSystemId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerSystemId) PatternFlowLacpduPartnerSystemId
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerSystemId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerSystemId
	// validate validates PatternFlowLacpduPartnerSystemId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerSystemId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerSystemIdChoiceEnum, set in PatternFlowLacpduPartnerSystemId
	Choice() PatternFlowLacpduPartnerSystemIdChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerSystemIdChoiceEnum provided by user to PatternFlowLacpduPartnerSystemId
	setChoice(value PatternFlowLacpduPartnerSystemIdChoiceEnum) PatternFlowLacpduPartnerSystemId
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerSystemId
	HasChoice() bool
	// Value returns string, set in PatternFlowLacpduPartnerSystemId.
	Value() string
	// SetValue assigns string provided by user to PatternFlowLacpduPartnerSystemId
	SetValue(value string) PatternFlowLacpduPartnerSystemId
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerSystemId
	HasValue() bool
	// Values returns []string, set in PatternFlowLacpduPartnerSystemId.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowLacpduPartnerSystemId
	SetValues(value []string) PatternFlowLacpduPartnerSystemId
	// Increment returns PatternFlowLacpduPartnerSystemIdCounter, set in PatternFlowLacpduPartnerSystemId.
	// PatternFlowLacpduPartnerSystemIdCounter is mac counter pattern
	Increment() PatternFlowLacpduPartnerSystemIdCounter
	// SetIncrement assigns PatternFlowLacpduPartnerSystemIdCounter provided by user to PatternFlowLacpduPartnerSystemId.
	// PatternFlowLacpduPartnerSystemIdCounter is mac counter pattern
	SetIncrement(value PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemId
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerSystemId
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerSystemIdCounter, set in PatternFlowLacpduPartnerSystemId.
	// PatternFlowLacpduPartnerSystemIdCounter is mac counter pattern
	Decrement() PatternFlowLacpduPartnerSystemIdCounter
	// SetDecrement assigns PatternFlowLacpduPartnerSystemIdCounter provided by user to PatternFlowLacpduPartnerSystemId.
	// PatternFlowLacpduPartnerSystemIdCounter is mac counter pattern
	SetDecrement(value PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemId
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerSystemId
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerSystemIdChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerSystemId
var PatternFlowLacpduPartnerSystemIdChoice = struct {
	VALUE     PatternFlowLacpduPartnerSystemIdChoiceEnum
	VALUES    PatternFlowLacpduPartnerSystemIdChoiceEnum
	INCREMENT PatternFlowLacpduPartnerSystemIdChoiceEnum
	DECREMENT PatternFlowLacpduPartnerSystemIdChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerSystemIdChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerSystemIdChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerSystemIdChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerSystemIdChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerSystemId) Choice() PatternFlowLacpduPartnerSystemIdChoiceEnum {
	return PatternFlowLacpduPartnerSystemIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerSystemId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerSystemId) setChoice(value PatternFlowLacpduPartnerSystemIdChoiceEnum) PatternFlowLacpduPartnerSystemId {
	intValue, ok := otg.PatternFlowLacpduPartnerSystemId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerSystemIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerSystemId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerSystemIdChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerSystemIdChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerSystemIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerSystemIdCounter().msg()
	}

	if value == PatternFlowLacpduPartnerSystemIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerSystemIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpduPartnerSystemId) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowLacpduPartnerSystemId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowLacpduPartnerSystemId object
func (obj *patternFlowLacpduPartnerSystemId) SetValue(value string) PatternFlowLacpduPartnerSystemId {
	obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowLacpduPartnerSystemId) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowLacpduPartnerSystemId object
func (obj *patternFlowLacpduPartnerSystemId) SetValues(value []string) PatternFlowLacpduPartnerSystemId {
	obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerSystemIdCounter
func (obj *patternFlowLacpduPartnerSystemId) Increment() PatternFlowLacpduPartnerSystemIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerSystemIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerSystemIdCounter
func (obj *patternFlowLacpduPartnerSystemId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerSystemIdCounter value in the PatternFlowLacpduPartnerSystemId object
func (obj *patternFlowLacpduPartnerSystemId) SetIncrement(value PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemId {
	obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerSystemIdCounter
func (obj *patternFlowLacpduPartnerSystemId) Decrement() PatternFlowLacpduPartnerSystemIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerSystemIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerSystemIdCounter
func (obj *patternFlowLacpduPartnerSystemId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerSystemIdCounter value in the PatternFlowLacpduPartnerSystemId object
func (obj *patternFlowLacpduPartnerSystemId) SetDecrement(value PatternFlowLacpduPartnerSystemIdCounter) PatternFlowLacpduPartnerSystemId {
	obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerSystemId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduPartnerSystemId.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowLacpduPartnerSystemId.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowLacpduPartnerSystemId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerSystemIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerSystemIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerSystemIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerSystemId")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerSystemId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerSystemId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
