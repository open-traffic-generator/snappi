package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv2MaxResponseTime *****
type patternFlowIgmpv2MaxResponseTime struct {
	validation
	obj             *otg.PatternFlowIgmpv2MaxResponseTime
	marshaller      marshalPatternFlowIgmpv2MaxResponseTime
	unMarshaller    unMarshalPatternFlowIgmpv2MaxResponseTime
	incrementHolder PatternFlowIgmpv2MaxResponseTimeCounter
	decrementHolder PatternFlowIgmpv2MaxResponseTimeCounter
}

func NewPatternFlowIgmpv2MaxResponseTime() PatternFlowIgmpv2MaxResponseTime {
	obj := patternFlowIgmpv2MaxResponseTime{obj: &otg.PatternFlowIgmpv2MaxResponseTime{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv2MaxResponseTime) msg() *otg.PatternFlowIgmpv2MaxResponseTime {
	return obj.obj
}

func (obj *patternFlowIgmpv2MaxResponseTime) setMsg(msg *otg.PatternFlowIgmpv2MaxResponseTime) PatternFlowIgmpv2MaxResponseTime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv2MaxResponseTime struct {
	obj *patternFlowIgmpv2MaxResponseTime
}

type marshalPatternFlowIgmpv2MaxResponseTime interface {
	// ToProto marshals PatternFlowIgmpv2MaxResponseTime to protobuf object *otg.PatternFlowIgmpv2MaxResponseTime
	ToProto() (*otg.PatternFlowIgmpv2MaxResponseTime, error)
	// ToPbText marshals PatternFlowIgmpv2MaxResponseTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv2MaxResponseTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv2MaxResponseTime to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv2MaxResponseTime struct {
	obj *patternFlowIgmpv2MaxResponseTime
}

type unMarshalPatternFlowIgmpv2MaxResponseTime interface {
	// FromProto unmarshals PatternFlowIgmpv2MaxResponseTime from protobuf object *otg.PatternFlowIgmpv2MaxResponseTime
	FromProto(msg *otg.PatternFlowIgmpv2MaxResponseTime) (PatternFlowIgmpv2MaxResponseTime, error)
	// FromPbText unmarshals PatternFlowIgmpv2MaxResponseTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv2MaxResponseTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv2MaxResponseTime from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv2MaxResponseTime) Marshal() marshalPatternFlowIgmpv2MaxResponseTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv2MaxResponseTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv2MaxResponseTime) Unmarshal() unMarshalPatternFlowIgmpv2MaxResponseTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv2MaxResponseTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv2MaxResponseTime) ToProto() (*otg.PatternFlowIgmpv2MaxResponseTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv2MaxResponseTime) FromProto(msg *otg.PatternFlowIgmpv2MaxResponseTime) (PatternFlowIgmpv2MaxResponseTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv2MaxResponseTime) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTime) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv2MaxResponseTime) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTime) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv2MaxResponseTime) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv2MaxResponseTime) FromJson(value string) error {
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

func (obj *patternFlowIgmpv2MaxResponseTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2MaxResponseTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv2MaxResponseTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv2MaxResponseTime) Clone() (PatternFlowIgmpv2MaxResponseTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv2MaxResponseTime()
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

func (obj *patternFlowIgmpv2MaxResponseTime) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv2MaxResponseTime is maximum response time in units of 1/10 second. Meaningful only in Membership Query messages; ignored in Reports and Leave Group (RFC 2236 Section 2.2).
type PatternFlowIgmpv2MaxResponseTime interface {
	Validation
	// msg marshals PatternFlowIgmpv2MaxResponseTime to protobuf object *otg.PatternFlowIgmpv2MaxResponseTime
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv2MaxResponseTime
	// setMsg unmarshals PatternFlowIgmpv2MaxResponseTime from protobuf object *otg.PatternFlowIgmpv2MaxResponseTime
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv2MaxResponseTime) PatternFlowIgmpv2MaxResponseTime
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv2MaxResponseTime
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv2MaxResponseTime
	// validate validates PatternFlowIgmpv2MaxResponseTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv2MaxResponseTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv2MaxResponseTimeChoiceEnum, set in PatternFlowIgmpv2MaxResponseTime
	Choice() PatternFlowIgmpv2MaxResponseTimeChoiceEnum
	// setChoice assigns PatternFlowIgmpv2MaxResponseTimeChoiceEnum provided by user to PatternFlowIgmpv2MaxResponseTime
	setChoice(value PatternFlowIgmpv2MaxResponseTimeChoiceEnum) PatternFlowIgmpv2MaxResponseTime
	// HasChoice checks if Choice has been set in PatternFlowIgmpv2MaxResponseTime
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv2MaxResponseTime.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv2MaxResponseTime
	SetValue(value uint32) PatternFlowIgmpv2MaxResponseTime
	// HasValue checks if Value has been set in PatternFlowIgmpv2MaxResponseTime
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv2MaxResponseTime.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv2MaxResponseTime
	SetValues(value []uint32) PatternFlowIgmpv2MaxResponseTime
	// Increment returns PatternFlowIgmpv2MaxResponseTimeCounter, set in PatternFlowIgmpv2MaxResponseTime.
	// PatternFlowIgmpv2MaxResponseTimeCounter is integer counter pattern
	Increment() PatternFlowIgmpv2MaxResponseTimeCounter
	// SetIncrement assigns PatternFlowIgmpv2MaxResponseTimeCounter provided by user to PatternFlowIgmpv2MaxResponseTime.
	// PatternFlowIgmpv2MaxResponseTimeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTime
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv2MaxResponseTime
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv2MaxResponseTimeCounter, set in PatternFlowIgmpv2MaxResponseTime.
	// PatternFlowIgmpv2MaxResponseTimeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv2MaxResponseTimeCounter
	// SetDecrement assigns PatternFlowIgmpv2MaxResponseTimeCounter provided by user to PatternFlowIgmpv2MaxResponseTime.
	// PatternFlowIgmpv2MaxResponseTimeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTime
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv2MaxResponseTime
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv2MaxResponseTimeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv2MaxResponseTime
var PatternFlowIgmpv2MaxResponseTimeChoice = struct {
	VALUE     PatternFlowIgmpv2MaxResponseTimeChoiceEnum
	VALUES    PatternFlowIgmpv2MaxResponseTimeChoiceEnum
	INCREMENT PatternFlowIgmpv2MaxResponseTimeChoiceEnum
	DECREMENT PatternFlowIgmpv2MaxResponseTimeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv2MaxResponseTimeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv2MaxResponseTimeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv2MaxResponseTimeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv2MaxResponseTimeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv2MaxResponseTime) Choice() PatternFlowIgmpv2MaxResponseTimeChoiceEnum {
	return PatternFlowIgmpv2MaxResponseTimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv2MaxResponseTime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv2MaxResponseTime) setChoice(value PatternFlowIgmpv2MaxResponseTimeChoiceEnum) PatternFlowIgmpv2MaxResponseTime {
	intValue, ok := otg.PatternFlowIgmpv2MaxResponseTime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv2MaxResponseTimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv2MaxResponseTime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv2MaxResponseTimeChoice.VALUE {
		defaultValue := uint32(100)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv2MaxResponseTimeChoice.VALUES {
		defaultValue := []uint32{100}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv2MaxResponseTimeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv2MaxResponseTimeCounter().msg()
	}

	if value == PatternFlowIgmpv2MaxResponseTimeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv2MaxResponseTimeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTime) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv2MaxResponseTime) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv2MaxResponseTime object
func (obj *patternFlowIgmpv2MaxResponseTime) SetValue(value uint32) PatternFlowIgmpv2MaxResponseTime {
	obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv2MaxResponseTime) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{100})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv2MaxResponseTime object
func (obj *patternFlowIgmpv2MaxResponseTime) SetValues(value []uint32) PatternFlowIgmpv2MaxResponseTime {
	obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv2MaxResponseTimeCounter
func (obj *patternFlowIgmpv2MaxResponseTime) Increment() PatternFlowIgmpv2MaxResponseTimeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv2MaxResponseTimeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv2MaxResponseTimeCounter
func (obj *patternFlowIgmpv2MaxResponseTime) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv2MaxResponseTimeCounter value in the PatternFlowIgmpv2MaxResponseTime object
func (obj *patternFlowIgmpv2MaxResponseTime) SetIncrement(value PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTime {
	obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2MaxResponseTimeCounter
func (obj *patternFlowIgmpv2MaxResponseTime) Decrement() PatternFlowIgmpv2MaxResponseTimeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv2MaxResponseTimeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv2MaxResponseTimeCounter
func (obj *patternFlowIgmpv2MaxResponseTime) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv2MaxResponseTimeCounter value in the PatternFlowIgmpv2MaxResponseTime object
func (obj *patternFlowIgmpv2MaxResponseTime) SetDecrement(value PatternFlowIgmpv2MaxResponseTimeCounter) PatternFlowIgmpv2MaxResponseTime {
	obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv2MaxResponseTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv2MaxResponseTime.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv2MaxResponseTime.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv2MaxResponseTime) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv2MaxResponseTimeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2MaxResponseTimeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv2MaxResponseTimeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2MaxResponseTimeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv2MaxResponseTimeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv2MaxResponseTimeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv2MaxResponseTime")
			}
		} else {
			intVal := otg.PatternFlowIgmpv2MaxResponseTime_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv2MaxResponseTime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
