package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeOptionClass *****
type patternFlowIpv4OptionsCustomTypeOptionClass struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	marshaller      marshalPatternFlowIpv4OptionsCustomTypeOptionClass
	unMarshaller    unMarshalPatternFlowIpv4OptionsCustomTypeOptionClass
	incrementHolder PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	decrementHolder PatternFlowIpv4OptionsCustomTypeOptionClassCounter
}

func NewPatternFlowIpv4OptionsCustomTypeOptionClass() PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj := patternFlowIpv4OptionsCustomTypeOptionClass{obj: &otg.PatternFlowIpv4OptionsCustomTypeOptionClass{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionClass {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClass) PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeOptionClass struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionClass
}

type marshalPatternFlowIpv4OptionsCustomTypeOptionClass interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeOptionClass to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionClass, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeOptionClass to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeOptionClass to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeOptionClass to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4OptionsCustomTypeOptionClass to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass struct {
	obj *patternFlowIpv4OptionsCustomTypeOptionClass
}

type unMarshalPatternFlowIpv4OptionsCustomTypeOptionClass interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeOptionClass from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClass) (PatternFlowIpv4OptionsCustomTypeOptionClass, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeOptionClass from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeOptionClass from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeOptionClass from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionClass {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeOptionClass{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionClass {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClass) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeOptionClass, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeOptionClass) (PatternFlowIpv4OptionsCustomTypeOptionClass, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClass) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClass) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClass) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeOptionClass) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeOptionClass) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Clone() (PatternFlowIpv4OptionsCustomTypeOptionClass, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeOptionClass()
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsCustomTypeOptionClass is option class [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
type PatternFlowIpv4OptionsCustomTypeOptionClass interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeOptionClass to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeOptionClass from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeOptionClass
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeOptionClass) PatternFlowIpv4OptionsCustomTypeOptionClass
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeOptionClass
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeOptionClass
	// validate validates PatternFlowIpv4OptionsCustomTypeOptionClass
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeOptionClass, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum, set in PatternFlowIpv4OptionsCustomTypeOptionClass
	Choice() PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum provided by user to PatternFlowIpv4OptionsCustomTypeOptionClass
	setChoice(value PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum) PatternFlowIpv4OptionsCustomTypeOptionClass
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsCustomTypeOptionClass
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsCustomTypeOptionClass.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionClass
	SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClass
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsCustomTypeOptionClass
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsCustomTypeOptionClass.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsCustomTypeOptionClass
	SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeOptionClass
	// Increment returns PatternFlowIpv4OptionsCustomTypeOptionClassCounter, set in PatternFlowIpv4OptionsCustomTypeOptionClass.
	// PatternFlowIpv4OptionsCustomTypeOptionClassCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// SetIncrement assigns PatternFlowIpv4OptionsCustomTypeOptionClassCounter provided by user to PatternFlowIpv4OptionsCustomTypeOptionClass.
	// PatternFlowIpv4OptionsCustomTypeOptionClassCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClass
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsCustomTypeOptionClass
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsCustomTypeOptionClassCounter, set in PatternFlowIpv4OptionsCustomTypeOptionClass.
	// PatternFlowIpv4OptionsCustomTypeOptionClassCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsCustomTypeOptionClassCounter
	// SetDecrement assigns PatternFlowIpv4OptionsCustomTypeOptionClassCounter provided by user to PatternFlowIpv4OptionsCustomTypeOptionClass.
	// PatternFlowIpv4OptionsCustomTypeOptionClassCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClass
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsCustomTypeOptionClass
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsCustomTypeOptionClass
var PatternFlowIpv4OptionsCustomTypeOptionClassChoice = struct {
	VALUE     PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum
	VALUES    PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum
	INCREMENT PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum
	DECREMENT PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Choice() PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum {
	return PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) setChoice(value PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum) PatternFlowIpv4OptionsCustomTypeOptionClass {
	intValue, ok := otg.PatternFlowIpv4OptionsCustomTypeOptionClass_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsCustomTypeOptionClass_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionClassChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsCustomTypeOptionClassCounter().msg()
	}

	if value == PatternFlowIpv4OptionsCustomTypeOptionClassChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsCustomTypeOptionClassCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionClass object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsCustomTypeOptionClass object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeOptionClassCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Increment() PatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsCustomTypeOptionClassCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeOptionClassCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsCustomTypeOptionClassCounter value in the PatternFlowIpv4OptionsCustomTypeOptionClass object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) SetIncrement(value PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeOptionClassCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) Decrement() PatternFlowIpv4OptionsCustomTypeOptionClassCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsCustomTypeOptionClassCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeOptionClassCounter
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsCustomTypeOptionClassCounter value in the PatternFlowIpv4OptionsCustomTypeOptionClass object
func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) SetDecrement(value PatternFlowIpv4OptionsCustomTypeOptionClassCounter) PatternFlowIpv4OptionsCustomTypeOptionClass {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 3 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeOptionClass.Value <= 3 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 3 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4OptionsCustomTypeOptionClass.Values <= 3 but Got %d", item))
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

func (obj *patternFlowIpv4OptionsCustomTypeOptionClass) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsCustomTypeOptionClassChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionClassChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeOptionClassChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsCustomTypeOptionClassChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsCustomTypeOptionClass")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsCustomTypeOptionClass_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsCustomTypeOptionClass_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
