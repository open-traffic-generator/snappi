package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmCcmFlags *****
type patternFlowCfmCcmFlags struct {
	validation
	obj             *otg.PatternFlowCfmCcmFlags
	marshaller      marshalPatternFlowCfmCcmFlags
	unMarshaller    unMarshalPatternFlowCfmCcmFlags
	incrementHolder PatternFlowCfmCcmFlagsCounter
	decrementHolder PatternFlowCfmCcmFlagsCounter
}

func NewPatternFlowCfmCcmFlags() PatternFlowCfmCcmFlags {
	obj := patternFlowCfmCcmFlags{obj: &otg.PatternFlowCfmCcmFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmCcmFlags) msg() *otg.PatternFlowCfmCcmFlags {
	return obj.obj
}

func (obj *patternFlowCfmCcmFlags) setMsg(msg *otg.PatternFlowCfmCcmFlags) PatternFlowCfmCcmFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmCcmFlags struct {
	obj *patternFlowCfmCcmFlags
}

type marshalPatternFlowCfmCcmFlags interface {
	// ToProto marshals PatternFlowCfmCcmFlags to protobuf object *otg.PatternFlowCfmCcmFlags
	ToProto() (*otg.PatternFlowCfmCcmFlags, error)
	// ToPbText marshals PatternFlowCfmCcmFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmCcmFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmCcmFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmCcmFlags struct {
	obj *patternFlowCfmCcmFlags
}

type unMarshalPatternFlowCfmCcmFlags interface {
	// FromProto unmarshals PatternFlowCfmCcmFlags from protobuf object *otg.PatternFlowCfmCcmFlags
	FromProto(msg *otg.PatternFlowCfmCcmFlags) (PatternFlowCfmCcmFlags, error)
	// FromPbText unmarshals PatternFlowCfmCcmFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmCcmFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmCcmFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmCcmFlags) Marshal() marshalPatternFlowCfmCcmFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmCcmFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmCcmFlags) Unmarshal() unMarshalPatternFlowCfmCcmFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmCcmFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmCcmFlags) ToProto() (*otg.PatternFlowCfmCcmFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmCcmFlags) FromProto(msg *otg.PatternFlowCfmCcmFlags) (PatternFlowCfmCcmFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmCcmFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmCcmFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmCcmFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmCcmFlags) FromJson(value string) error {
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

func (obj *patternFlowCfmCcmFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmCcmFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmCcmFlags) Clone() (PatternFlowCfmCcmFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmCcmFlags()
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

func (obj *patternFlowCfmCcmFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmCcmFlags is defines operational flags for CCM.
type PatternFlowCfmCcmFlags interface {
	Validation
	// msg marshals PatternFlowCfmCcmFlags to protobuf object *otg.PatternFlowCfmCcmFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmCcmFlags
	// setMsg unmarshals PatternFlowCfmCcmFlags from protobuf object *otg.PatternFlowCfmCcmFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmCcmFlags) PatternFlowCfmCcmFlags
	// provides marshal interface
	Marshal() marshalPatternFlowCfmCcmFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmCcmFlags
	// validate validates PatternFlowCfmCcmFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmCcmFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmCcmFlagsChoiceEnum, set in PatternFlowCfmCcmFlags
	Choice() PatternFlowCfmCcmFlagsChoiceEnum
	// setChoice assigns PatternFlowCfmCcmFlagsChoiceEnum provided by user to PatternFlowCfmCcmFlags
	setChoice(value PatternFlowCfmCcmFlagsChoiceEnum) PatternFlowCfmCcmFlags
	// HasChoice checks if Choice has been set in PatternFlowCfmCcmFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmCcmFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmCcmFlags
	SetValue(value uint32) PatternFlowCfmCcmFlags
	// HasValue checks if Value has been set in PatternFlowCfmCcmFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmCcmFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmCcmFlags
	SetValues(value []uint32) PatternFlowCfmCcmFlags
	// Increment returns PatternFlowCfmCcmFlagsCounter, set in PatternFlowCfmCcmFlags.
	// PatternFlowCfmCcmFlagsCounter is integer counter pattern
	Increment() PatternFlowCfmCcmFlagsCounter
	// SetIncrement assigns PatternFlowCfmCcmFlagsCounter provided by user to PatternFlowCfmCcmFlags.
	// PatternFlowCfmCcmFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlags
	// HasIncrement checks if Increment has been set in PatternFlowCfmCcmFlags
	HasIncrement() bool
	// Decrement returns PatternFlowCfmCcmFlagsCounter, set in PatternFlowCfmCcmFlags.
	// PatternFlowCfmCcmFlagsCounter is integer counter pattern
	Decrement() PatternFlowCfmCcmFlagsCounter
	// SetDecrement assigns PatternFlowCfmCcmFlagsCounter provided by user to PatternFlowCfmCcmFlags.
	// PatternFlowCfmCcmFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlags
	// HasDecrement checks if Decrement has been set in PatternFlowCfmCcmFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmCcmFlagsChoiceEnum string

// Enum of Choice on PatternFlowCfmCcmFlags
var PatternFlowCfmCcmFlagsChoice = struct {
	VALUE     PatternFlowCfmCcmFlagsChoiceEnum
	VALUES    PatternFlowCfmCcmFlagsChoiceEnum
	INCREMENT PatternFlowCfmCcmFlagsChoiceEnum
	DECREMENT PatternFlowCfmCcmFlagsChoiceEnum
}{
	VALUE:     PatternFlowCfmCcmFlagsChoiceEnum("value"),
	VALUES:    PatternFlowCfmCcmFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowCfmCcmFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmCcmFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowCfmCcmFlags) Choice() PatternFlowCfmCcmFlagsChoiceEnum {
	return PatternFlowCfmCcmFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmCcmFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmCcmFlags) setChoice(value PatternFlowCfmCcmFlagsChoiceEnum) PatternFlowCfmCcmFlags {
	intValue, ok := otg.PatternFlowCfmCcmFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmCcmFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmCcmFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmCcmFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmCcmFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmCcmFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmCcmFlagsCounter().msg()
	}

	if value == PatternFlowCfmCcmFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmCcmFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmCcmFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmCcmFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmCcmFlags object
func (obj *patternFlowCfmCcmFlags) SetValue(value uint32) PatternFlowCfmCcmFlags {
	obj.setChoice(PatternFlowCfmCcmFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmCcmFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmCcmFlags object
func (obj *patternFlowCfmCcmFlags) SetValues(value []uint32) PatternFlowCfmCcmFlags {
	obj.setChoice(PatternFlowCfmCcmFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmCcmFlagsCounter
func (obj *patternFlowCfmCcmFlags) Increment() PatternFlowCfmCcmFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmCcmFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmCcmFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmCcmFlagsCounter
func (obj *patternFlowCfmCcmFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmCcmFlagsCounter value in the PatternFlowCfmCcmFlags object
func (obj *patternFlowCfmCcmFlags) SetIncrement(value PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlags {
	obj.setChoice(PatternFlowCfmCcmFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmFlagsCounter
func (obj *patternFlowCfmCcmFlags) Decrement() PatternFlowCfmCcmFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmCcmFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmCcmFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmCcmFlagsCounter
func (obj *patternFlowCfmCcmFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmCcmFlagsCounter value in the PatternFlowCfmCcmFlags object
func (obj *patternFlowCfmCcmFlags) SetDecrement(value PatternFlowCfmCcmFlagsCounter) PatternFlowCfmCcmFlags {
	obj.setChoice(PatternFlowCfmCcmFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmCcmFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmCcmFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmCcmFlags.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmCcmFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmCcmFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmCcmFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmCcmFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmCcmFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmCcmFlags")
			}
		} else {
			intVal := otg.PatternFlowCfmCcmFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmCcmFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
