package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduActorStateExpired *****
type patternFlowLacpduActorStateExpired struct {
	validation
	obj             *otg.PatternFlowLacpduActorStateExpired
	marshaller      marshalPatternFlowLacpduActorStateExpired
	unMarshaller    unMarshalPatternFlowLacpduActorStateExpired
	incrementHolder PatternFlowLacpduActorStateExpiredCounter
	decrementHolder PatternFlowLacpduActorStateExpiredCounter
}

func NewPatternFlowLacpduActorStateExpired() PatternFlowLacpduActorStateExpired {
	obj := patternFlowLacpduActorStateExpired{obj: &otg.PatternFlowLacpduActorStateExpired{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduActorStateExpired) msg() *otg.PatternFlowLacpduActorStateExpired {
	return obj.obj
}

func (obj *patternFlowLacpduActorStateExpired) setMsg(msg *otg.PatternFlowLacpduActorStateExpired) PatternFlowLacpduActorStateExpired {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduActorStateExpired struct {
	obj *patternFlowLacpduActorStateExpired
}

type marshalPatternFlowLacpduActorStateExpired interface {
	// ToProto marshals PatternFlowLacpduActorStateExpired to protobuf object *otg.PatternFlowLacpduActorStateExpired
	ToProto() (*otg.PatternFlowLacpduActorStateExpired, error)
	// ToPbText marshals PatternFlowLacpduActorStateExpired to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduActorStateExpired to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduActorStateExpired to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduActorStateExpired struct {
	obj *patternFlowLacpduActorStateExpired
}

type unMarshalPatternFlowLacpduActorStateExpired interface {
	// FromProto unmarshals PatternFlowLacpduActorStateExpired from protobuf object *otg.PatternFlowLacpduActorStateExpired
	FromProto(msg *otg.PatternFlowLacpduActorStateExpired) (PatternFlowLacpduActorStateExpired, error)
	// FromPbText unmarshals PatternFlowLacpduActorStateExpired from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduActorStateExpired from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduActorStateExpired from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduActorStateExpired) Marshal() marshalPatternFlowLacpduActorStateExpired {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduActorStateExpired{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduActorStateExpired) Unmarshal() unMarshalPatternFlowLacpduActorStateExpired {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduActorStateExpired{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduActorStateExpired) ToProto() (*otg.PatternFlowLacpduActorStateExpired, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduActorStateExpired) FromProto(msg *otg.PatternFlowLacpduActorStateExpired) (PatternFlowLacpduActorStateExpired, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduActorStateExpired) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpired) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateExpired) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpired) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduActorStateExpired) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduActorStateExpired) FromJson(value string) error {
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

func (obj *patternFlowLacpduActorStateExpired) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateExpired) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduActorStateExpired) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduActorStateExpired) Clone() (PatternFlowLacpduActorStateExpired, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduActorStateExpired()
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

func (obj *patternFlowLacpduActorStateExpired) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduActorStateExpired is expired (1=Expired, 0=Not Expired)
type PatternFlowLacpduActorStateExpired interface {
	Validation
	// msg marshals PatternFlowLacpduActorStateExpired to protobuf object *otg.PatternFlowLacpduActorStateExpired
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduActorStateExpired
	// setMsg unmarshals PatternFlowLacpduActorStateExpired from protobuf object *otg.PatternFlowLacpduActorStateExpired
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduActorStateExpired) PatternFlowLacpduActorStateExpired
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduActorStateExpired
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduActorStateExpired
	// validate validates PatternFlowLacpduActorStateExpired
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduActorStateExpired, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduActorStateExpiredChoiceEnum, set in PatternFlowLacpduActorStateExpired
	Choice() PatternFlowLacpduActorStateExpiredChoiceEnum
	// setChoice assigns PatternFlowLacpduActorStateExpiredChoiceEnum provided by user to PatternFlowLacpduActorStateExpired
	setChoice(value PatternFlowLacpduActorStateExpiredChoiceEnum) PatternFlowLacpduActorStateExpired
	// HasChoice checks if Choice has been set in PatternFlowLacpduActorStateExpired
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduActorStateExpired.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduActorStateExpired
	SetValue(value uint32) PatternFlowLacpduActorStateExpired
	// HasValue checks if Value has been set in PatternFlowLacpduActorStateExpired
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduActorStateExpired.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduActorStateExpired
	SetValues(value []uint32) PatternFlowLacpduActorStateExpired
	// Increment returns PatternFlowLacpduActorStateExpiredCounter, set in PatternFlowLacpduActorStateExpired.
	// PatternFlowLacpduActorStateExpiredCounter is integer counter pattern
	Increment() PatternFlowLacpduActorStateExpiredCounter
	// SetIncrement assigns PatternFlowLacpduActorStateExpiredCounter provided by user to PatternFlowLacpduActorStateExpired.
	// PatternFlowLacpduActorStateExpiredCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpired
	// HasIncrement checks if Increment has been set in PatternFlowLacpduActorStateExpired
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduActorStateExpiredCounter, set in PatternFlowLacpduActorStateExpired.
	// PatternFlowLacpduActorStateExpiredCounter is integer counter pattern
	Decrement() PatternFlowLacpduActorStateExpiredCounter
	// SetDecrement assigns PatternFlowLacpduActorStateExpiredCounter provided by user to PatternFlowLacpduActorStateExpired.
	// PatternFlowLacpduActorStateExpiredCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpired
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduActorStateExpired
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduActorStateExpiredChoiceEnum string

// Enum of Choice on PatternFlowLacpduActorStateExpired
var PatternFlowLacpduActorStateExpiredChoice = struct {
	VALUE     PatternFlowLacpduActorStateExpiredChoiceEnum
	VALUES    PatternFlowLacpduActorStateExpiredChoiceEnum
	INCREMENT PatternFlowLacpduActorStateExpiredChoiceEnum
	DECREMENT PatternFlowLacpduActorStateExpiredChoiceEnum
}{
	VALUE:     PatternFlowLacpduActorStateExpiredChoiceEnum("value"),
	VALUES:    PatternFlowLacpduActorStateExpiredChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduActorStateExpiredChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduActorStateExpiredChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduActorStateExpired) Choice() PatternFlowLacpduActorStateExpiredChoiceEnum {
	return PatternFlowLacpduActorStateExpiredChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduActorStateExpired) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduActorStateExpired) setChoice(value PatternFlowLacpduActorStateExpiredChoiceEnum) PatternFlowLacpduActorStateExpired {
	intValue, ok := otg.PatternFlowLacpduActorStateExpired_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduActorStateExpiredChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduActorStateExpired_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduActorStateExpiredChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduActorStateExpiredChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduActorStateExpiredChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduActorStateExpiredCounter().msg()
	}

	if value == PatternFlowLacpduActorStateExpiredChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduActorStateExpiredCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateExpired) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduActorStateExpired) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduActorStateExpired object
func (obj *patternFlowLacpduActorStateExpired) SetValue(value uint32) PatternFlowLacpduActorStateExpired {
	obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduActorStateExpired) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduActorStateExpired object
func (obj *patternFlowLacpduActorStateExpired) SetValues(value []uint32) PatternFlowLacpduActorStateExpired {
	obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateExpiredCounter
func (obj *patternFlowLacpduActorStateExpired) Increment() PatternFlowLacpduActorStateExpiredCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduActorStateExpiredCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduActorStateExpiredCounter
func (obj *patternFlowLacpduActorStateExpired) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduActorStateExpiredCounter value in the PatternFlowLacpduActorStateExpired object
func (obj *patternFlowLacpduActorStateExpired) SetIncrement(value PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpired {
	obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateExpiredCounter
func (obj *patternFlowLacpduActorStateExpired) Decrement() PatternFlowLacpduActorStateExpiredCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduActorStateExpiredCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduActorStateExpiredCounter
func (obj *patternFlowLacpduActorStateExpired) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduActorStateExpiredCounter value in the PatternFlowLacpduActorStateExpired object
func (obj *patternFlowLacpduActorStateExpired) SetDecrement(value PatternFlowLacpduActorStateExpiredCounter) PatternFlowLacpduActorStateExpired {
	obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduActorStateExpired) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduActorStateExpired.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduActorStateExpired.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduActorStateExpired) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduActorStateExpiredChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateExpiredChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduActorStateExpiredChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateExpiredChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduActorStateExpiredChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduActorStateExpiredChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduActorStateExpired")
			}
		} else {
			intVal := otg.PatternFlowLacpduActorStateExpired_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduActorStateExpired_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
