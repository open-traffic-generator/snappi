package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateTimeout *****
type patternFlowLacpduPartnerStateTimeout struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateTimeout
	marshaller      marshalPatternFlowLacpduPartnerStateTimeout
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateTimeout
	incrementHolder PatternFlowLacpduPartnerStateTimeoutCounter
	decrementHolder PatternFlowLacpduPartnerStateTimeoutCounter
}

func NewPatternFlowLacpduPartnerStateTimeout() PatternFlowLacpduPartnerStateTimeout {
	obj := patternFlowLacpduPartnerStateTimeout{obj: &otg.PatternFlowLacpduPartnerStateTimeout{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateTimeout) msg() *otg.PatternFlowLacpduPartnerStateTimeout {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateTimeout) setMsg(msg *otg.PatternFlowLacpduPartnerStateTimeout) PatternFlowLacpduPartnerStateTimeout {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateTimeout struct {
	obj *patternFlowLacpduPartnerStateTimeout
}

type marshalPatternFlowLacpduPartnerStateTimeout interface {
	// ToProto marshals PatternFlowLacpduPartnerStateTimeout to protobuf object *otg.PatternFlowLacpduPartnerStateTimeout
	ToProto() (*otg.PatternFlowLacpduPartnerStateTimeout, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateTimeout to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateTimeout to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateTimeout to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateTimeout struct {
	obj *patternFlowLacpduPartnerStateTimeout
}

type unMarshalPatternFlowLacpduPartnerStateTimeout interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateTimeout from protobuf object *otg.PatternFlowLacpduPartnerStateTimeout
	FromProto(msg *otg.PatternFlowLacpduPartnerStateTimeout) (PatternFlowLacpduPartnerStateTimeout, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateTimeout from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateTimeout from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateTimeout from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateTimeout) Marshal() marshalPatternFlowLacpduPartnerStateTimeout {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateTimeout{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateTimeout) Unmarshal() unMarshalPatternFlowLacpduPartnerStateTimeout {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateTimeout{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateTimeout) ToProto() (*otg.PatternFlowLacpduPartnerStateTimeout, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateTimeout) FromProto(msg *otg.PatternFlowLacpduPartnerStateTimeout) (PatternFlowLacpduPartnerStateTimeout, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateTimeout) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeout) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateTimeout) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeout) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateTimeout) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateTimeout) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateTimeout) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateTimeout) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateTimeout) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateTimeout) Clone() (PatternFlowLacpduPartnerStateTimeout, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateTimeout()
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

func (obj *patternFlowLacpduPartnerStateTimeout) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateTimeout is lACP Timeout (1=Fast timeout, 0=Slow timeout)
type PatternFlowLacpduPartnerStateTimeout interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateTimeout to protobuf object *otg.PatternFlowLacpduPartnerStateTimeout
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateTimeout
	// setMsg unmarshals PatternFlowLacpduPartnerStateTimeout from protobuf object *otg.PatternFlowLacpduPartnerStateTimeout
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateTimeout) PatternFlowLacpduPartnerStateTimeout
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateTimeout
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateTimeout
	// validate validates PatternFlowLacpduPartnerStateTimeout
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateTimeout, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateTimeoutChoiceEnum, set in PatternFlowLacpduPartnerStateTimeout
	Choice() PatternFlowLacpduPartnerStateTimeoutChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateTimeoutChoiceEnum provided by user to PatternFlowLacpduPartnerStateTimeout
	setChoice(value PatternFlowLacpduPartnerStateTimeoutChoiceEnum) PatternFlowLacpduPartnerStateTimeout
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateTimeout
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateTimeout.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateTimeout
	SetValue(value uint32) PatternFlowLacpduPartnerStateTimeout
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateTimeout
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateTimeout.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateTimeout
	SetValues(value []uint32) PatternFlowLacpduPartnerStateTimeout
	// Increment returns PatternFlowLacpduPartnerStateTimeoutCounter, set in PatternFlowLacpduPartnerStateTimeout.
	// PatternFlowLacpduPartnerStateTimeoutCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateTimeoutCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateTimeoutCounter provided by user to PatternFlowLacpduPartnerStateTimeout.
	// PatternFlowLacpduPartnerStateTimeoutCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeout
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateTimeout
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateTimeoutCounter, set in PatternFlowLacpduPartnerStateTimeout.
	// PatternFlowLacpduPartnerStateTimeoutCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateTimeoutCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateTimeoutCounter provided by user to PatternFlowLacpduPartnerStateTimeout.
	// PatternFlowLacpduPartnerStateTimeoutCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeout
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateTimeout
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateTimeoutChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateTimeout
var PatternFlowLacpduPartnerStateTimeoutChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateTimeoutChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateTimeoutChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateTimeoutChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateTimeoutChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateTimeoutChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateTimeoutChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateTimeoutChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateTimeoutChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateTimeout) Choice() PatternFlowLacpduPartnerStateTimeoutChoiceEnum {
	return PatternFlowLacpduPartnerStateTimeoutChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateTimeout) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateTimeout) setChoice(value PatternFlowLacpduPartnerStateTimeoutChoiceEnum) PatternFlowLacpduPartnerStateTimeout {
	intValue, ok := otg.PatternFlowLacpduPartnerStateTimeout_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateTimeoutChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateTimeout_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateTimeoutChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateTimeoutChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateTimeoutChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateTimeoutCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateTimeoutChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateTimeoutCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeout) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateTimeout) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateTimeout object
func (obj *patternFlowLacpduPartnerStateTimeout) SetValue(value uint32) PatternFlowLacpduPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateTimeout) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateTimeout object
func (obj *patternFlowLacpduPartnerStateTimeout) SetValues(value []uint32) PatternFlowLacpduPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateTimeoutCounter
func (obj *patternFlowLacpduPartnerStateTimeout) Increment() PatternFlowLacpduPartnerStateTimeoutCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateTimeoutCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateTimeoutCounter
func (obj *patternFlowLacpduPartnerStateTimeout) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateTimeoutCounter value in the PatternFlowLacpduPartnerStateTimeout object
func (obj *patternFlowLacpduPartnerStateTimeout) SetIncrement(value PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateTimeoutCounter
func (obj *patternFlowLacpduPartnerStateTimeout) Decrement() PatternFlowLacpduPartnerStateTimeoutCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateTimeoutCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateTimeoutCounter
func (obj *patternFlowLacpduPartnerStateTimeout) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateTimeoutCounter value in the PatternFlowLacpduPartnerStateTimeout object
func (obj *patternFlowLacpduPartnerStateTimeout) SetDecrement(value PatternFlowLacpduPartnerStateTimeoutCounter) PatternFlowLacpduPartnerStateTimeout {
	obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateTimeout) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateTimeout.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateTimeout.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateTimeout) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateTimeoutChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateTimeoutChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateTimeoutChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateTimeoutChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateTimeoutChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateTimeoutChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateTimeout")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateTimeout_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateTimeout_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
