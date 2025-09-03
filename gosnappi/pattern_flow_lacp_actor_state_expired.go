package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpActorStateExpired *****
type patternFlowLacpActorStateExpired struct {
	validation
	obj             *otg.PatternFlowLacpActorStateExpired
	marshaller      marshalPatternFlowLacpActorStateExpired
	unMarshaller    unMarshalPatternFlowLacpActorStateExpired
	incrementHolder PatternFlowLacpActorStateExpiredCounter
	decrementHolder PatternFlowLacpActorStateExpiredCounter
}

func NewPatternFlowLacpActorStateExpired() PatternFlowLacpActorStateExpired {
	obj := patternFlowLacpActorStateExpired{obj: &otg.PatternFlowLacpActorStateExpired{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpActorStateExpired) msg() *otg.PatternFlowLacpActorStateExpired {
	return obj.obj
}

func (obj *patternFlowLacpActorStateExpired) setMsg(msg *otg.PatternFlowLacpActorStateExpired) PatternFlowLacpActorStateExpired {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpActorStateExpired struct {
	obj *patternFlowLacpActorStateExpired
}

type marshalPatternFlowLacpActorStateExpired interface {
	// ToProto marshals PatternFlowLacpActorStateExpired to protobuf object *otg.PatternFlowLacpActorStateExpired
	ToProto() (*otg.PatternFlowLacpActorStateExpired, error)
	// ToPbText marshals PatternFlowLacpActorStateExpired to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpActorStateExpired to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpActorStateExpired to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpActorStateExpired struct {
	obj *patternFlowLacpActorStateExpired
}

type unMarshalPatternFlowLacpActorStateExpired interface {
	// FromProto unmarshals PatternFlowLacpActorStateExpired from protobuf object *otg.PatternFlowLacpActorStateExpired
	FromProto(msg *otg.PatternFlowLacpActorStateExpired) (PatternFlowLacpActorStateExpired, error)
	// FromPbText unmarshals PatternFlowLacpActorStateExpired from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpActorStateExpired from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpActorStateExpired from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpActorStateExpired) Marshal() marshalPatternFlowLacpActorStateExpired {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpActorStateExpired{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpActorStateExpired) Unmarshal() unMarshalPatternFlowLacpActorStateExpired {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpActorStateExpired{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpActorStateExpired) ToProto() (*otg.PatternFlowLacpActorStateExpired, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpActorStateExpired) FromProto(msg *otg.PatternFlowLacpActorStateExpired) (PatternFlowLacpActorStateExpired, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpActorStateExpired) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpired) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpActorStateExpired) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpired) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpActorStateExpired) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpActorStateExpired) FromJson(value string) error {
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

func (obj *patternFlowLacpActorStateExpired) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateExpired) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpActorStateExpired) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpActorStateExpired) Clone() (PatternFlowLacpActorStateExpired, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpActorStateExpired()
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

func (obj *patternFlowLacpActorStateExpired) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpActorStateExpired is expired (1=Expired, 0=Not Expired)
type PatternFlowLacpActorStateExpired interface {
	Validation
	// msg marshals PatternFlowLacpActorStateExpired to protobuf object *otg.PatternFlowLacpActorStateExpired
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpActorStateExpired
	// setMsg unmarshals PatternFlowLacpActorStateExpired from protobuf object *otg.PatternFlowLacpActorStateExpired
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpActorStateExpired) PatternFlowLacpActorStateExpired
	// provides marshal interface
	Marshal() marshalPatternFlowLacpActorStateExpired
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpActorStateExpired
	// validate validates PatternFlowLacpActorStateExpired
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpActorStateExpired, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpActorStateExpiredChoiceEnum, set in PatternFlowLacpActorStateExpired
	Choice() PatternFlowLacpActorStateExpiredChoiceEnum
	// setChoice assigns PatternFlowLacpActorStateExpiredChoiceEnum provided by user to PatternFlowLacpActorStateExpired
	setChoice(value PatternFlowLacpActorStateExpiredChoiceEnum) PatternFlowLacpActorStateExpired
	// HasChoice checks if Choice has been set in PatternFlowLacpActorStateExpired
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpActorStateExpired.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpActorStateExpired
	SetValue(value uint32) PatternFlowLacpActorStateExpired
	// HasValue checks if Value has been set in PatternFlowLacpActorStateExpired
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpActorStateExpired.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpActorStateExpired
	SetValues(value []uint32) PatternFlowLacpActorStateExpired
	// Increment returns PatternFlowLacpActorStateExpiredCounter, set in PatternFlowLacpActorStateExpired.
	// PatternFlowLacpActorStateExpiredCounter is integer counter pattern
	Increment() PatternFlowLacpActorStateExpiredCounter
	// SetIncrement assigns PatternFlowLacpActorStateExpiredCounter provided by user to PatternFlowLacpActorStateExpired.
	// PatternFlowLacpActorStateExpiredCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpired
	// HasIncrement checks if Increment has been set in PatternFlowLacpActorStateExpired
	HasIncrement() bool
	// Decrement returns PatternFlowLacpActorStateExpiredCounter, set in PatternFlowLacpActorStateExpired.
	// PatternFlowLacpActorStateExpiredCounter is integer counter pattern
	Decrement() PatternFlowLacpActorStateExpiredCounter
	// SetDecrement assigns PatternFlowLacpActorStateExpiredCounter provided by user to PatternFlowLacpActorStateExpired.
	// PatternFlowLacpActorStateExpiredCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpired
	// HasDecrement checks if Decrement has been set in PatternFlowLacpActorStateExpired
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpActorStateExpiredChoiceEnum string

// Enum of Choice on PatternFlowLacpActorStateExpired
var PatternFlowLacpActorStateExpiredChoice = struct {
	VALUE     PatternFlowLacpActorStateExpiredChoiceEnum
	VALUES    PatternFlowLacpActorStateExpiredChoiceEnum
	INCREMENT PatternFlowLacpActorStateExpiredChoiceEnum
	DECREMENT PatternFlowLacpActorStateExpiredChoiceEnum
}{
	VALUE:     PatternFlowLacpActorStateExpiredChoiceEnum("value"),
	VALUES:    PatternFlowLacpActorStateExpiredChoiceEnum("values"),
	INCREMENT: PatternFlowLacpActorStateExpiredChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpActorStateExpiredChoiceEnum("decrement"),
}

func (obj *patternFlowLacpActorStateExpired) Choice() PatternFlowLacpActorStateExpiredChoiceEnum {
	return PatternFlowLacpActorStateExpiredChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpActorStateExpired) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpActorStateExpired) setChoice(value PatternFlowLacpActorStateExpiredChoiceEnum) PatternFlowLacpActorStateExpired {
	intValue, ok := otg.PatternFlowLacpActorStateExpired_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpActorStateExpiredChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpActorStateExpired_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpActorStateExpiredChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpActorStateExpiredChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpActorStateExpiredChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpActorStateExpiredCounter().msg()
	}

	if value == PatternFlowLacpActorStateExpiredChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpActorStateExpiredCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateExpired) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpActorStateExpiredChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpActorStateExpired) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpActorStateExpired object
func (obj *patternFlowLacpActorStateExpired) SetValue(value uint32) PatternFlowLacpActorStateExpired {
	obj.setChoice(PatternFlowLacpActorStateExpiredChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpActorStateExpired) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpActorStateExpired object
func (obj *patternFlowLacpActorStateExpired) SetValues(value []uint32) PatternFlowLacpActorStateExpired {
	obj.setChoice(PatternFlowLacpActorStateExpiredChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateExpiredCounter
func (obj *patternFlowLacpActorStateExpired) Increment() PatternFlowLacpActorStateExpiredCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpActorStateExpiredChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpActorStateExpiredCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpActorStateExpiredCounter
func (obj *patternFlowLacpActorStateExpired) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpActorStateExpiredCounter value in the PatternFlowLacpActorStateExpired object
func (obj *patternFlowLacpActorStateExpired) SetIncrement(value PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpired {
	obj.setChoice(PatternFlowLacpActorStateExpiredChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateExpiredCounter
func (obj *patternFlowLacpActorStateExpired) Decrement() PatternFlowLacpActorStateExpiredCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpActorStateExpiredChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpActorStateExpiredCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpActorStateExpiredCounter
func (obj *patternFlowLacpActorStateExpired) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpActorStateExpiredCounter value in the PatternFlowLacpActorStateExpired object
func (obj *patternFlowLacpActorStateExpired) SetDecrement(value PatternFlowLacpActorStateExpiredCounter) PatternFlowLacpActorStateExpired {
	obj.setChoice(PatternFlowLacpActorStateExpiredChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpActorStateExpired) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpActorStateExpired.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpActorStateExpired.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpActorStateExpired) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpActorStateExpiredChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateExpiredChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpActorStateExpiredChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateExpiredChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpActorStateExpiredChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpActorStateExpiredChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpActorStateExpired")
			}
		} else {
			intVal := otg.PatternFlowLacpActorStateExpired_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpActorStateExpired_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
