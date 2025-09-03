package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateDefaulted *****
type patternFlowLacpPartnerStateDefaulted struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateDefaulted
	marshaller      marshalPatternFlowLacpPartnerStateDefaulted
	unMarshaller    unMarshalPatternFlowLacpPartnerStateDefaulted
	incrementHolder PatternFlowLacpPartnerStateDefaultedCounter
	decrementHolder PatternFlowLacpPartnerStateDefaultedCounter
}

func NewPatternFlowLacpPartnerStateDefaulted() PatternFlowLacpPartnerStateDefaulted {
	obj := patternFlowLacpPartnerStateDefaulted{obj: &otg.PatternFlowLacpPartnerStateDefaulted{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateDefaulted) msg() *otg.PatternFlowLacpPartnerStateDefaulted {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateDefaulted) setMsg(msg *otg.PatternFlowLacpPartnerStateDefaulted) PatternFlowLacpPartnerStateDefaulted {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateDefaulted struct {
	obj *patternFlowLacpPartnerStateDefaulted
}

type marshalPatternFlowLacpPartnerStateDefaulted interface {
	// ToProto marshals PatternFlowLacpPartnerStateDefaulted to protobuf object *otg.PatternFlowLacpPartnerStateDefaulted
	ToProto() (*otg.PatternFlowLacpPartnerStateDefaulted, error)
	// ToPbText marshals PatternFlowLacpPartnerStateDefaulted to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateDefaulted to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateDefaulted to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateDefaulted struct {
	obj *patternFlowLacpPartnerStateDefaulted
}

type unMarshalPatternFlowLacpPartnerStateDefaulted interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateDefaulted from protobuf object *otg.PatternFlowLacpPartnerStateDefaulted
	FromProto(msg *otg.PatternFlowLacpPartnerStateDefaulted) (PatternFlowLacpPartnerStateDefaulted, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateDefaulted from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateDefaulted from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateDefaulted from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateDefaulted) Marshal() marshalPatternFlowLacpPartnerStateDefaulted {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateDefaulted{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateDefaulted) Unmarshal() unMarshalPatternFlowLacpPartnerStateDefaulted {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateDefaulted{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateDefaulted) ToProto() (*otg.PatternFlowLacpPartnerStateDefaulted, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateDefaulted) FromProto(msg *otg.PatternFlowLacpPartnerStateDefaulted) (PatternFlowLacpPartnerStateDefaulted, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateDefaulted) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaulted) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDefaulted) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaulted) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDefaulted) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDefaulted) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateDefaulted) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDefaulted) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDefaulted) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateDefaulted) Clone() (PatternFlowLacpPartnerStateDefaulted, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateDefaulted()
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

func (obj *patternFlowLacpPartnerStateDefaulted) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
type PatternFlowLacpPartnerStateDefaulted interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateDefaulted to protobuf object *otg.PatternFlowLacpPartnerStateDefaulted
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateDefaulted
	// setMsg unmarshals PatternFlowLacpPartnerStateDefaulted from protobuf object *otg.PatternFlowLacpPartnerStateDefaulted
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateDefaulted) PatternFlowLacpPartnerStateDefaulted
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateDefaulted
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateDefaulted
	// validate validates PatternFlowLacpPartnerStateDefaulted
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateDefaulted, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateDefaultedChoiceEnum, set in PatternFlowLacpPartnerStateDefaulted
	Choice() PatternFlowLacpPartnerStateDefaultedChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateDefaultedChoiceEnum provided by user to PatternFlowLacpPartnerStateDefaulted
	setChoice(value PatternFlowLacpPartnerStateDefaultedChoiceEnum) PatternFlowLacpPartnerStateDefaulted
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateDefaulted
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateDefaulted.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateDefaulted
	SetValue(value uint32) PatternFlowLacpPartnerStateDefaulted
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateDefaulted
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateDefaulted.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateDefaulted
	SetValues(value []uint32) PatternFlowLacpPartnerStateDefaulted
	// Increment returns PatternFlowLacpPartnerStateDefaultedCounter, set in PatternFlowLacpPartnerStateDefaulted.
	// PatternFlowLacpPartnerStateDefaultedCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateDefaultedCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateDefaultedCounter provided by user to PatternFlowLacpPartnerStateDefaulted.
	// PatternFlowLacpPartnerStateDefaultedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaulted
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateDefaulted
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateDefaultedCounter, set in PatternFlowLacpPartnerStateDefaulted.
	// PatternFlowLacpPartnerStateDefaultedCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateDefaultedCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateDefaultedCounter provided by user to PatternFlowLacpPartnerStateDefaulted.
	// PatternFlowLacpPartnerStateDefaultedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaulted
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateDefaulted
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateDefaultedChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateDefaulted
var PatternFlowLacpPartnerStateDefaultedChoice = struct {
	VALUE     PatternFlowLacpPartnerStateDefaultedChoiceEnum
	VALUES    PatternFlowLacpPartnerStateDefaultedChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateDefaultedChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateDefaultedChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateDefaultedChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateDefaultedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateDefaultedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateDefaultedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateDefaulted) Choice() PatternFlowLacpPartnerStateDefaultedChoiceEnum {
	return PatternFlowLacpPartnerStateDefaultedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateDefaulted) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateDefaulted) setChoice(value PatternFlowLacpPartnerStateDefaultedChoiceEnum) PatternFlowLacpPartnerStateDefaulted {
	intValue, ok := otg.PatternFlowLacpPartnerStateDefaulted_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateDefaultedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateDefaulted_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateDefaultedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateDefaultedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateDefaultedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateDefaultedCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateDefaultedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateDefaultedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateDefaulted) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateDefaulted) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateDefaulted object
func (obj *patternFlowLacpPartnerStateDefaulted) SetValue(value uint32) PatternFlowLacpPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateDefaulted) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateDefaulted object
func (obj *patternFlowLacpPartnerStateDefaulted) SetValues(value []uint32) PatternFlowLacpPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateDefaultedCounter
func (obj *patternFlowLacpPartnerStateDefaulted) Increment() PatternFlowLacpPartnerStateDefaultedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateDefaultedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateDefaultedCounter
func (obj *patternFlowLacpPartnerStateDefaulted) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateDefaultedCounter value in the PatternFlowLacpPartnerStateDefaulted object
func (obj *patternFlowLacpPartnerStateDefaulted) SetIncrement(value PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateDefaultedCounter
func (obj *patternFlowLacpPartnerStateDefaulted) Decrement() PatternFlowLacpPartnerStateDefaultedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateDefaultedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateDefaultedCounter
func (obj *patternFlowLacpPartnerStateDefaulted) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateDefaultedCounter value in the PatternFlowLacpPartnerStateDefaulted object
func (obj *patternFlowLacpPartnerStateDefaulted) SetDecrement(value PatternFlowLacpPartnerStateDefaultedCounter) PatternFlowLacpPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateDefaulted) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDefaulted.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateDefaulted.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateDefaulted) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateDefaultedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDefaultedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDefaultedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDefaultedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDefaultedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateDefaultedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateDefaulted")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateDefaulted_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateDefaulted_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
