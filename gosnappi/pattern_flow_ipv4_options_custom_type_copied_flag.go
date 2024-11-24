package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4OptionsCustomTypeCopiedFlag *****
type patternFlowIpv4OptionsCustomTypeCopiedFlag struct {
	validation
	obj             *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	marshaller      marshalPatternFlowIpv4OptionsCustomTypeCopiedFlag
	unMarshaller    unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlag
	incrementHolder PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	decrementHolder PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
}

func NewPatternFlowIpv4OptionsCustomTypeCopiedFlag() PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj := patternFlowIpv4OptionsCustomTypeCopiedFlag{obj: &otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) msg() *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	return obj.obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) setMsg(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag struct {
	obj *patternFlowIpv4OptionsCustomTypeCopiedFlag
}

type marshalPatternFlowIpv4OptionsCustomTypeCopiedFlag interface {
	// ToProto marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag, error)
	// ToPbText marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag struct {
	obj *patternFlowIpv4OptionsCustomTypeCopiedFlag
}

type unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlag interface {
	// FromProto unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlag from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag) (PatternFlowIpv4OptionsCustomTypeCopiedFlag, error)
	// FromPbText unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Marshal() marshalPatternFlowIpv4OptionsCustomTypeCopiedFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) ToProto() (*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) FromProto(msg *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag) (PatternFlowIpv4OptionsCustomTypeCopiedFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4OptionsCustomTypeCopiedFlag) FromJson(value string) error {
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

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Clone() (PatternFlowIpv4OptionsCustomTypeCopiedFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4OptionsCustomTypeCopiedFlag()
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

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4OptionsCustomTypeCopiedFlag is this flag indicates this option is copied to all fragments on fragmentations.
type PatternFlowIpv4OptionsCustomTypeCopiedFlag interface {
	Validation
	// msg marshals PatternFlowIpv4OptionsCustomTypeCopiedFlag to protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// setMsg unmarshals PatternFlowIpv4OptionsCustomTypeCopiedFlag from protobuf object *otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4OptionsCustomTypeCopiedFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4OptionsCustomTypeCopiedFlag
	// validate validates PatternFlowIpv4OptionsCustomTypeCopiedFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4OptionsCustomTypeCopiedFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum, set in PatternFlowIpv4OptionsCustomTypeCopiedFlag
	Choice() PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum
	// setChoice assigns PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlag
	setChoice(value PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// HasChoice checks if Choice has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlag
	SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// HasValue checks if Value has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlag
	SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// Increment returns PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, set in PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter is integer counter pattern
	Increment() PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// SetIncrement assigns PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// HasIncrement checks if Increment has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter, set in PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter is integer counter pattern
	Decrement() PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
	// SetDecrement assigns PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter provided by user to PatternFlowIpv4OptionsCustomTypeCopiedFlag.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4OptionsCustomTypeCopiedFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum string

// Enum of Choice on PatternFlowIpv4OptionsCustomTypeCopiedFlag
var PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice = struct {
	VALUE     PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum
	VALUES    PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum
	INCREMENT PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum
	DECREMENT PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum
}{
	VALUE:     PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum("value"),
	VALUES:    PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Choice() PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum {
	return PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) setChoice(value PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	intValue, ok := otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter().msg()
	}

	if value == PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4OptionsCustomTypeCopiedFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4OptionsCustomTypeCopiedFlag object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) SetValue(value uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4OptionsCustomTypeCopiedFlag object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) SetValues(value []uint32) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Increment() PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4OptionsCustomTypeCopiedFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter value in the PatternFlowIpv4OptionsCustomTypeCopiedFlag object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) SetIncrement(value PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) Decrement() PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4OptionsCustomTypeCopiedFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter value in the PatternFlowIpv4OptionsCustomTypeCopiedFlag object
func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) SetDecrement(value PatternFlowIpv4OptionsCustomTypeCopiedFlagCounter) PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4OptionsCustomTypeCopiedFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4OptionsCustomTypeCopiedFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv4OptionsCustomTypeCopiedFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4OptionsCustomTypeCopiedFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4OptionsCustomTypeCopiedFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4OptionsCustomTypeCopiedFlag")
			}
		} else {
			intVal := otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4OptionsCustomTypeCopiedFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
