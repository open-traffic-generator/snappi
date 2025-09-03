package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateTimeout *****
type patternFlowLacpPartnerStateTimeout struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateTimeout
	marshaller      marshalPatternFlowLacpPartnerStateTimeout
	unMarshaller    unMarshalPatternFlowLacpPartnerStateTimeout
	incrementHolder PatternFlowLacpPartnerStateTimeoutCounter
	decrementHolder PatternFlowLacpPartnerStateTimeoutCounter
}

func NewPatternFlowLacpPartnerStateTimeout() PatternFlowLacpPartnerStateTimeout {
	obj := patternFlowLacpPartnerStateTimeout{obj: &otg.PatternFlowLacpPartnerStateTimeout{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateTimeout) msg() *otg.PatternFlowLacpPartnerStateTimeout {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateTimeout) setMsg(msg *otg.PatternFlowLacpPartnerStateTimeout) PatternFlowLacpPartnerStateTimeout {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateTimeout struct {
	obj *patternFlowLacpPartnerStateTimeout
}

type marshalPatternFlowLacpPartnerStateTimeout interface {
	// ToProto marshals PatternFlowLacpPartnerStateTimeout to protobuf object *otg.PatternFlowLacpPartnerStateTimeout
	ToProto() (*otg.PatternFlowLacpPartnerStateTimeout, error)
	// ToPbText marshals PatternFlowLacpPartnerStateTimeout to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateTimeout to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateTimeout to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateTimeout struct {
	obj *patternFlowLacpPartnerStateTimeout
}

type unMarshalPatternFlowLacpPartnerStateTimeout interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateTimeout from protobuf object *otg.PatternFlowLacpPartnerStateTimeout
	FromProto(msg *otg.PatternFlowLacpPartnerStateTimeout) (PatternFlowLacpPartnerStateTimeout, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateTimeout from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateTimeout from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateTimeout from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateTimeout) Marshal() marshalPatternFlowLacpPartnerStateTimeout {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateTimeout{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateTimeout) Unmarshal() unMarshalPatternFlowLacpPartnerStateTimeout {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateTimeout{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateTimeout) ToProto() (*otg.PatternFlowLacpPartnerStateTimeout, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateTimeout) FromProto(msg *otg.PatternFlowLacpPartnerStateTimeout) (PatternFlowLacpPartnerStateTimeout, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateTimeout) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeout) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateTimeout) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeout) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateTimeout) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateTimeout) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateTimeout) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateTimeout) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateTimeout) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateTimeout) Clone() (PatternFlowLacpPartnerStateTimeout, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateTimeout()
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

func (obj *patternFlowLacpPartnerStateTimeout) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
type PatternFlowLacpPartnerStateTimeout interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateTimeout to protobuf object *otg.PatternFlowLacpPartnerStateTimeout
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateTimeout
	// setMsg unmarshals PatternFlowLacpPartnerStateTimeout from protobuf object *otg.PatternFlowLacpPartnerStateTimeout
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateTimeout) PatternFlowLacpPartnerStateTimeout
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateTimeout
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateTimeout
	// validate validates PatternFlowLacpPartnerStateTimeout
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateTimeout, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateTimeoutChoiceEnum, set in PatternFlowLacpPartnerStateTimeout
	Choice() PatternFlowLacpPartnerStateTimeoutChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateTimeoutChoiceEnum provided by user to PatternFlowLacpPartnerStateTimeout
	setChoice(value PatternFlowLacpPartnerStateTimeoutChoiceEnum) PatternFlowLacpPartnerStateTimeout
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateTimeout
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateTimeout.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateTimeout
	SetValue(value uint32) PatternFlowLacpPartnerStateTimeout
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateTimeout
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateTimeout.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateTimeout
	SetValues(value []uint32) PatternFlowLacpPartnerStateTimeout
	// Increment returns PatternFlowLacpPartnerStateTimeoutCounter, set in PatternFlowLacpPartnerStateTimeout.
	// PatternFlowLacpPartnerStateTimeoutCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateTimeoutCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateTimeoutCounter provided by user to PatternFlowLacpPartnerStateTimeout.
	// PatternFlowLacpPartnerStateTimeoutCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeout
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateTimeout
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateTimeoutCounter, set in PatternFlowLacpPartnerStateTimeout.
	// PatternFlowLacpPartnerStateTimeoutCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateTimeoutCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateTimeoutCounter provided by user to PatternFlowLacpPartnerStateTimeout.
	// PatternFlowLacpPartnerStateTimeoutCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeout
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateTimeout
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateTimeoutChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateTimeout
var PatternFlowLacpPartnerStateTimeoutChoice = struct {
	VALUE     PatternFlowLacpPartnerStateTimeoutChoiceEnum
	VALUES    PatternFlowLacpPartnerStateTimeoutChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateTimeoutChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateTimeoutChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateTimeoutChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateTimeoutChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateTimeoutChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateTimeoutChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateTimeout) Choice() PatternFlowLacpPartnerStateTimeoutChoiceEnum {
	return PatternFlowLacpPartnerStateTimeoutChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateTimeout) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateTimeout) setChoice(value PatternFlowLacpPartnerStateTimeoutChoiceEnum) PatternFlowLacpPartnerStateTimeout {
	intValue, ok := otg.PatternFlowLacpPartnerStateTimeout_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateTimeoutChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateTimeout_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateTimeoutChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateTimeoutChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateTimeoutChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateTimeoutCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateTimeoutChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateTimeoutCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateTimeout) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateTimeout) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateTimeout object
func (obj *patternFlowLacpPartnerStateTimeout) SetValue(value uint32) PatternFlowLacpPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateTimeout) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateTimeout object
func (obj *patternFlowLacpPartnerStateTimeout) SetValues(value []uint32) PatternFlowLacpPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateTimeoutCounter
func (obj *patternFlowLacpPartnerStateTimeout) Increment() PatternFlowLacpPartnerStateTimeoutCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateTimeoutCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateTimeoutCounter
func (obj *patternFlowLacpPartnerStateTimeout) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateTimeoutCounter value in the PatternFlowLacpPartnerStateTimeout object
func (obj *patternFlowLacpPartnerStateTimeout) SetIncrement(value PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateTimeoutCounter
func (obj *patternFlowLacpPartnerStateTimeout) Decrement() PatternFlowLacpPartnerStateTimeoutCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateTimeoutCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateTimeoutCounter
func (obj *patternFlowLacpPartnerStateTimeout) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateTimeoutCounter value in the PatternFlowLacpPartnerStateTimeout object
func (obj *patternFlowLacpPartnerStateTimeout) SetDecrement(value PatternFlowLacpPartnerStateTimeoutCounter) PatternFlowLacpPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateTimeout) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateTimeout.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateTimeout.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateTimeout) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateTimeoutChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateTimeoutChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateTimeoutChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateTimeoutChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateTimeoutChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateTimeoutChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateTimeout")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateTimeout_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateTimeout_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
