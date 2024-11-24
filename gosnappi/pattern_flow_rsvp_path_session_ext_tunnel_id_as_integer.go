package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSessionExtTunnelIdAsInteger *****
type patternFlowRSVPPathSessionExtTunnelIdAsInteger struct {
	validation
	obj             *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	marshaller      marshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger
	unMarshaller    unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger
	incrementHolder PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	decrementHolder PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
}

func NewPatternFlowRSVPPathSessionExtTunnelIdAsInteger() PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj := patternFlowRSVPPathSessionExtTunnelIdAsInteger{obj: &otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	return obj.obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) setMsg(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger
}

type marshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger interface {
	// ToProto marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error)
	// ToPbText marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger struct {
	obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger
}

type unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger interface {
	// FromProto unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger) (PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error)
	// FromPbText unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) ToProto() (*otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) FromProto(msg *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger) (PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSessionExtTunnelIdAsInteger) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSessionExtTunnelIdAsInteger()
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

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSessionExtTunnelIdAsInteger is tBD
type PatternFlowRSVPPathSessionExtTunnelIdAsInteger interface {
	Validation
	// msg marshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger to protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// setMsg unmarshals PatternFlowRSVPPathSessionExtTunnelIdAsInteger from protobuf object *otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// validate validates PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSessionExtTunnelIdAsInteger, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum, set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	Choice() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	setChoice(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	SetValue(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// HasValue checks if Value has been set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	SetValues(value []uint32) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// Increment returns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// SetIncrement assigns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter, set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
	// SetDecrement assigns PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter provided by user to PatternFlowRSVPPathSessionExtTunnelIdAsInteger.
	// PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSessionExtTunnelIdAsInteger
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSessionExtTunnelIdAsInteger
var PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice = struct {
	VALUE     PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum
	VALUES    PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum
	INCREMENT PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum
	DECREMENT PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Choice() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum {
	return PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) setChoice(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	intValue, ok := otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter().msg()
	}

	if value == PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsInteger object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) SetValue(value uint32) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSessionExtTunnelIdAsInteger object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) SetValues(value []uint32) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Increment() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter value in the PatternFlowRSVPPathSessionExtTunnelIdAsInteger object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) SetIncrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) Decrement() PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter value in the PatternFlowRSVPPathSessionExtTunnelIdAsInteger object
func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) SetDecrement(value PatternFlowRSVPPathSessionExtTunnelIdAsIntegerCounter) PatternFlowRSVPPathSessionExtTunnelIdAsInteger {
	obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathSessionExtTunnelIdAsInteger) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSessionExtTunnelIdAsIntegerChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSessionExtTunnelIdAsInteger")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSessionExtTunnelIdAsInteger_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
