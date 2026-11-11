package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 *****
type patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	marshaller      marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	unMarshaller    unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	incrementHolder PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
	decrementHolder PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
}

func NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj := patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2{obj: &otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) setMsg(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
}

type marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 interface {
	// ToProto marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error)
	// ToPbText marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 struct {
	obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
}

type unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) ToProto() (*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FromProto(msg *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2()
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 is reserved. Must be initialized to zero by the sender.
type PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 to protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// setMsg unmarshals PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 from protobuf object *otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// validate validates PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	Choice() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	setChoice(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	SetValue(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	SetValues(value []uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// Increment returns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
	// SetIncrement assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter, set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
	// SetDecrement assigns PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter provided by user to PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2.
	// PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2
var PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice = struct {
	VALUE     PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum
	VALUES    PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Choice() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum {
	return PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) setChoice(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	intValue, ok := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter().msg()
	}

	if value == PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) SetValue(value uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) SetValues(value []uint32) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Increment() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) SetIncrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) Decrement() PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter value in the PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 object
func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) SetDecrement(value PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Counter) PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2 {
	obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIcmpv6NdpRedirectedHeaderOptionReserved2) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpRedirectedHeaderOptionReserved2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
