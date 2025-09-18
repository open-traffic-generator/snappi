package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateDefaulted *****
type patternFlowLacpduPartnerStateDefaulted struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateDefaulted
	marshaller      marshalPatternFlowLacpduPartnerStateDefaulted
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateDefaulted
	incrementHolder PatternFlowLacpduPartnerStateDefaultedCounter
	decrementHolder PatternFlowLacpduPartnerStateDefaultedCounter
}

func NewPatternFlowLacpduPartnerStateDefaulted() PatternFlowLacpduPartnerStateDefaulted {
	obj := patternFlowLacpduPartnerStateDefaulted{obj: &otg.PatternFlowLacpduPartnerStateDefaulted{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateDefaulted) msg() *otg.PatternFlowLacpduPartnerStateDefaulted {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateDefaulted) setMsg(msg *otg.PatternFlowLacpduPartnerStateDefaulted) PatternFlowLacpduPartnerStateDefaulted {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateDefaulted struct {
	obj *patternFlowLacpduPartnerStateDefaulted
}

type marshalPatternFlowLacpduPartnerStateDefaulted interface {
	// ToProto marshals PatternFlowLacpduPartnerStateDefaulted to protobuf object *otg.PatternFlowLacpduPartnerStateDefaulted
	ToProto() (*otg.PatternFlowLacpduPartnerStateDefaulted, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateDefaulted to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateDefaulted to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateDefaulted to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateDefaulted struct {
	obj *patternFlowLacpduPartnerStateDefaulted
}

type unMarshalPatternFlowLacpduPartnerStateDefaulted interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateDefaulted from protobuf object *otg.PatternFlowLacpduPartnerStateDefaulted
	FromProto(msg *otg.PatternFlowLacpduPartnerStateDefaulted) (PatternFlowLacpduPartnerStateDefaulted, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateDefaulted from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateDefaulted from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateDefaulted from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateDefaulted) Marshal() marshalPatternFlowLacpduPartnerStateDefaulted {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateDefaulted{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateDefaulted) Unmarshal() unMarshalPatternFlowLacpduPartnerStateDefaulted {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateDefaulted{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateDefaulted) ToProto() (*otg.PatternFlowLacpduPartnerStateDefaulted, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateDefaulted) FromProto(msg *otg.PatternFlowLacpduPartnerStateDefaulted) (PatternFlowLacpduPartnerStateDefaulted, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateDefaulted) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaulted) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDefaulted) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaulted) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDefaulted) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDefaulted) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateDefaulted) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDefaulted) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDefaulted) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateDefaulted) Clone() (PatternFlowLacpduPartnerStateDefaulted, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateDefaulted()
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

func (obj *patternFlowLacpduPartnerStateDefaulted) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateDefaulted is defaulted (1=Using defaulted partner info, 0=Using received partner info)
type PatternFlowLacpduPartnerStateDefaulted interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateDefaulted to protobuf object *otg.PatternFlowLacpduPartnerStateDefaulted
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateDefaulted
	// setMsg unmarshals PatternFlowLacpduPartnerStateDefaulted from protobuf object *otg.PatternFlowLacpduPartnerStateDefaulted
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateDefaulted) PatternFlowLacpduPartnerStateDefaulted
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateDefaulted
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateDefaulted
	// validate validates PatternFlowLacpduPartnerStateDefaulted
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateDefaulted, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateDefaultedChoiceEnum, set in PatternFlowLacpduPartnerStateDefaulted
	Choice() PatternFlowLacpduPartnerStateDefaultedChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateDefaultedChoiceEnum provided by user to PatternFlowLacpduPartnerStateDefaulted
	setChoice(value PatternFlowLacpduPartnerStateDefaultedChoiceEnum) PatternFlowLacpduPartnerStateDefaulted
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateDefaulted
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateDefaulted.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateDefaulted
	SetValue(value uint32) PatternFlowLacpduPartnerStateDefaulted
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateDefaulted
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateDefaulted.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateDefaulted
	SetValues(value []uint32) PatternFlowLacpduPartnerStateDefaulted
	// Increment returns PatternFlowLacpduPartnerStateDefaultedCounter, set in PatternFlowLacpduPartnerStateDefaulted.
	// PatternFlowLacpduPartnerStateDefaultedCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateDefaultedCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateDefaultedCounter provided by user to PatternFlowLacpduPartnerStateDefaulted.
	// PatternFlowLacpduPartnerStateDefaultedCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaulted
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateDefaulted
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateDefaultedCounter, set in PatternFlowLacpduPartnerStateDefaulted.
	// PatternFlowLacpduPartnerStateDefaultedCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateDefaultedCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateDefaultedCounter provided by user to PatternFlowLacpduPartnerStateDefaulted.
	// PatternFlowLacpduPartnerStateDefaultedCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaulted
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateDefaulted
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateDefaultedChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateDefaulted
var PatternFlowLacpduPartnerStateDefaultedChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateDefaultedChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateDefaultedChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateDefaultedChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateDefaultedChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateDefaultedChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateDefaultedChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateDefaultedChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateDefaultedChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateDefaulted) Choice() PatternFlowLacpduPartnerStateDefaultedChoiceEnum {
	return PatternFlowLacpduPartnerStateDefaultedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateDefaulted) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateDefaulted) setChoice(value PatternFlowLacpduPartnerStateDefaultedChoiceEnum) PatternFlowLacpduPartnerStateDefaulted {
	intValue, ok := otg.PatternFlowLacpduPartnerStateDefaulted_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateDefaultedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateDefaulted_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateDefaultedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateDefaultedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateDefaultedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateDefaultedCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateDefaultedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateDefaultedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaulted) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateDefaulted) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateDefaulted object
func (obj *patternFlowLacpduPartnerStateDefaulted) SetValue(value uint32) PatternFlowLacpduPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateDefaulted) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateDefaulted object
func (obj *patternFlowLacpduPartnerStateDefaulted) SetValues(value []uint32) PatternFlowLacpduPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateDefaultedCounter
func (obj *patternFlowLacpduPartnerStateDefaulted) Increment() PatternFlowLacpduPartnerStateDefaultedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateDefaultedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateDefaultedCounter
func (obj *patternFlowLacpduPartnerStateDefaulted) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateDefaultedCounter value in the PatternFlowLacpduPartnerStateDefaulted object
func (obj *patternFlowLacpduPartnerStateDefaulted) SetIncrement(value PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateDefaultedCounter
func (obj *patternFlowLacpduPartnerStateDefaulted) Decrement() PatternFlowLacpduPartnerStateDefaultedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateDefaultedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateDefaultedCounter
func (obj *patternFlowLacpduPartnerStateDefaulted) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateDefaultedCounter value in the PatternFlowLacpduPartnerStateDefaulted object
func (obj *patternFlowLacpduPartnerStateDefaulted) SetDecrement(value PatternFlowLacpduPartnerStateDefaultedCounter) PatternFlowLacpduPartnerStateDefaulted {
	obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateDefaulted) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDefaulted.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateDefaulted.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateDefaulted) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateDefaultedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDefaultedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDefaultedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDefaultedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDefaultedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateDefaultedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateDefaulted")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateDefaulted_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateDefaulted_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
