package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpMtuOptionMtu *****
type patternFlowIcmpv6NdpMtuOptionMtu struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	marshaller      marshalPatternFlowIcmpv6NdpMtuOptionMtu
	unMarshaller    unMarshalPatternFlowIcmpv6NdpMtuOptionMtu
	incrementHolder PatternFlowIcmpv6NdpMtuOptionMtuCounter
	decrementHolder PatternFlowIcmpv6NdpMtuOptionMtuCounter
}

func NewPatternFlowIcmpv6NdpMtuOptionMtu() PatternFlowIcmpv6NdpMtuOptionMtu {
	obj := patternFlowIcmpv6NdpMtuOptionMtu{obj: &otg.PatternFlowIcmpv6NdpMtuOptionMtu{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) msg() *otg.PatternFlowIcmpv6NdpMtuOptionMtu {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) setMsg(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtu) PatternFlowIcmpv6NdpMtuOptionMtu {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpMtuOptionMtu struct {
	obj *patternFlowIcmpv6NdpMtuOptionMtu
}

type marshalPatternFlowIcmpv6NdpMtuOptionMtu interface {
	// ToProto marshals PatternFlowIcmpv6NdpMtuOptionMtu to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionMtu, error)
	// ToPbText marshals PatternFlowIcmpv6NdpMtuOptionMtu to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpMtuOptionMtu to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpMtuOptionMtu to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpMtuOptionMtu struct {
	obj *patternFlowIcmpv6NdpMtuOptionMtu
}

type unMarshalPatternFlowIcmpv6NdpMtuOptionMtu interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpMtuOptionMtu from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtu) (PatternFlowIcmpv6NdpMtuOptionMtu, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpMtuOptionMtu from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpMtuOptionMtu from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpMtuOptionMtu from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Marshal() marshalPatternFlowIcmpv6NdpMtuOptionMtu {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpMtuOptionMtu{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionMtu {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpMtuOptionMtu{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtu) ToProto() (*otg.PatternFlowIcmpv6NdpMtuOptionMtu, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtu) FromProto(msg *otg.PatternFlowIcmpv6NdpMtuOptionMtu) (PatternFlowIcmpv6NdpMtuOptionMtu, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtu) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtu) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtu) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtu) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpMtuOptionMtu) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpMtuOptionMtu) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Clone() (PatternFlowIcmpv6NdpMtuOptionMtu, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpMtuOptionMtu()
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

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpMtuOptionMtu is recommended MTU for the link, in octets.
type PatternFlowIcmpv6NdpMtuOptionMtu interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpMtuOptionMtu to protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	// setMsg unmarshals PatternFlowIcmpv6NdpMtuOptionMtu from protobuf object *otg.PatternFlowIcmpv6NdpMtuOptionMtu
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpMtuOptionMtu) PatternFlowIcmpv6NdpMtuOptionMtu
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpMtuOptionMtu
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpMtuOptionMtu
	// validate validates PatternFlowIcmpv6NdpMtuOptionMtu
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpMtuOptionMtu, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum, set in PatternFlowIcmpv6NdpMtuOptionMtu
	Choice() PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum provided by user to PatternFlowIcmpv6NdpMtuOptionMtu
	setChoice(value PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum) PatternFlowIcmpv6NdpMtuOptionMtu
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpMtuOptionMtu
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpMtuOptionMtu.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionMtu
	SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionMtu
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpMtuOptionMtu
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpMtuOptionMtu.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpMtuOptionMtu
	SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionMtu
	// Increment returns PatternFlowIcmpv6NdpMtuOptionMtuCounter, set in PatternFlowIcmpv6NdpMtuOptionMtu.
	// PatternFlowIcmpv6NdpMtuOptionMtuCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpMtuOptionMtuCounter provided by user to PatternFlowIcmpv6NdpMtuOptionMtu.
	// PatternFlowIcmpv6NdpMtuOptionMtuCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtu
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpMtuOptionMtu
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpMtuOptionMtuCounter, set in PatternFlowIcmpv6NdpMtuOptionMtu.
	// PatternFlowIcmpv6NdpMtuOptionMtuCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpMtuOptionMtuCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpMtuOptionMtuCounter provided by user to PatternFlowIcmpv6NdpMtuOptionMtu.
	// PatternFlowIcmpv6NdpMtuOptionMtuCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtu
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpMtuOptionMtu
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpMtuOptionMtu
var PatternFlowIcmpv6NdpMtuOptionMtuChoice = struct {
	VALUE     PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum
	VALUES    PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Choice() PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum {
	return PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) setChoice(value PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum) PatternFlowIcmpv6NdpMtuOptionMtu {
	intValue, ok := otg.PatternFlowIcmpv6NdpMtuOptionMtu_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpMtuOptionMtu_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUE {
		defaultValue := uint32(1500)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUES {
		defaultValue := []uint32{1500}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpMtuOptionMtuChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpMtuOptionMtuCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpMtuOptionMtuChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpMtuOptionMtuCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpMtuOptionMtu object
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) SetValue(value uint32) PatternFlowIcmpv6NdpMtuOptionMtu {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1500})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpMtuOptionMtu object
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) SetValues(value []uint32) PatternFlowIcmpv6NdpMtuOptionMtu {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionMtuCounter
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Increment() PatternFlowIcmpv6NdpMtuOptionMtuCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpMtuOptionMtuCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpMtuOptionMtuCounter
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpMtuOptionMtuCounter value in the PatternFlowIcmpv6NdpMtuOptionMtu object
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) SetIncrement(value PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtu {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionMtuCounter
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) Decrement() PatternFlowIcmpv6NdpMtuOptionMtuCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpMtuOptionMtuCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpMtuOptionMtuCounter
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpMtuOptionMtuCounter value in the PatternFlowIcmpv6NdpMtuOptionMtu object
func (obj *patternFlowIcmpv6NdpMtuOptionMtu) SetDecrement(value PatternFlowIcmpv6NdpMtuOptionMtuCounter) PatternFlowIcmpv6NdpMtuOptionMtu {
	obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowIcmpv6NdpMtuOptionMtu) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpMtuOptionMtuChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionMtuChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpMtuOptionMtuChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpMtuOptionMtuChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpMtuOptionMtu")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpMtuOptionMtu_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpMtuOptionMtu_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
