package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbrFlags *****
type patternFlowCfmLbrFlags struct {
	validation
	obj             *otg.PatternFlowCfmLbrFlags
	marshaller      marshalPatternFlowCfmLbrFlags
	unMarshaller    unMarshalPatternFlowCfmLbrFlags
	incrementHolder PatternFlowCfmLbrFlagsCounter
	decrementHolder PatternFlowCfmLbrFlagsCounter
}

func NewPatternFlowCfmLbrFlags() PatternFlowCfmLbrFlags {
	obj := patternFlowCfmLbrFlags{obj: &otg.PatternFlowCfmLbrFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbrFlags) msg() *otg.PatternFlowCfmLbrFlags {
	return obj.obj
}

func (obj *patternFlowCfmLbrFlags) setMsg(msg *otg.PatternFlowCfmLbrFlags) PatternFlowCfmLbrFlags {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbrFlags struct {
	obj *patternFlowCfmLbrFlags
}

type marshalPatternFlowCfmLbrFlags interface {
	// ToProto marshals PatternFlowCfmLbrFlags to protobuf object *otg.PatternFlowCfmLbrFlags
	ToProto() (*otg.PatternFlowCfmLbrFlags, error)
	// ToPbText marshals PatternFlowCfmLbrFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbrFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbrFlags to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbrFlags struct {
	obj *patternFlowCfmLbrFlags
}

type unMarshalPatternFlowCfmLbrFlags interface {
	// FromProto unmarshals PatternFlowCfmLbrFlags from protobuf object *otg.PatternFlowCfmLbrFlags
	FromProto(msg *otg.PatternFlowCfmLbrFlags) (PatternFlowCfmLbrFlags, error)
	// FromPbText unmarshals PatternFlowCfmLbrFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbrFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbrFlags from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbrFlags) Marshal() marshalPatternFlowCfmLbrFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbrFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbrFlags) Unmarshal() unMarshalPatternFlowCfmLbrFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbrFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbrFlags) ToProto() (*otg.PatternFlowCfmLbrFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbrFlags) FromProto(msg *otg.PatternFlowCfmLbrFlags) (PatternFlowCfmLbrFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbrFlags) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlags) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbrFlags) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlags) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbrFlags) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrFlags) FromJson(value string) error {
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

func (obj *patternFlowCfmLbrFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbrFlags) Clone() (PatternFlowCfmLbrFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbrFlags()
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

func (obj *patternFlowCfmLbrFlags) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmLbrFlags is defines operational flags for LBR.
type PatternFlowCfmLbrFlags interface {
	Validation
	// msg marshals PatternFlowCfmLbrFlags to protobuf object *otg.PatternFlowCfmLbrFlags
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbrFlags
	// setMsg unmarshals PatternFlowCfmLbrFlags from protobuf object *otg.PatternFlowCfmLbrFlags
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbrFlags) PatternFlowCfmLbrFlags
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbrFlags
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbrFlags
	// validate validates PatternFlowCfmLbrFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbrFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmLbrFlagsChoiceEnum, set in PatternFlowCfmLbrFlags
	Choice() PatternFlowCfmLbrFlagsChoiceEnum
	// setChoice assigns PatternFlowCfmLbrFlagsChoiceEnum provided by user to PatternFlowCfmLbrFlags
	setChoice(value PatternFlowCfmLbrFlagsChoiceEnum) PatternFlowCfmLbrFlags
	// HasChoice checks if Choice has been set in PatternFlowCfmLbrFlags
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmLbrFlags.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmLbrFlags
	SetValue(value uint32) PatternFlowCfmLbrFlags
	// HasValue checks if Value has been set in PatternFlowCfmLbrFlags
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmLbrFlags.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmLbrFlags
	SetValues(value []uint32) PatternFlowCfmLbrFlags
	// Increment returns PatternFlowCfmLbrFlagsCounter, set in PatternFlowCfmLbrFlags.
	// PatternFlowCfmLbrFlagsCounter is integer counter pattern
	Increment() PatternFlowCfmLbrFlagsCounter
	// SetIncrement assigns PatternFlowCfmLbrFlagsCounter provided by user to PatternFlowCfmLbrFlags.
	// PatternFlowCfmLbrFlagsCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlags
	// HasIncrement checks if Increment has been set in PatternFlowCfmLbrFlags
	HasIncrement() bool
	// Decrement returns PatternFlowCfmLbrFlagsCounter, set in PatternFlowCfmLbrFlags.
	// PatternFlowCfmLbrFlagsCounter is integer counter pattern
	Decrement() PatternFlowCfmLbrFlagsCounter
	// SetDecrement assigns PatternFlowCfmLbrFlagsCounter provided by user to PatternFlowCfmLbrFlags.
	// PatternFlowCfmLbrFlagsCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlags
	// HasDecrement checks if Decrement has been set in PatternFlowCfmLbrFlags
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmLbrFlagsChoiceEnum string

// Enum of Choice on PatternFlowCfmLbrFlags
var PatternFlowCfmLbrFlagsChoice = struct {
	VALUE     PatternFlowCfmLbrFlagsChoiceEnum
	VALUES    PatternFlowCfmLbrFlagsChoiceEnum
	INCREMENT PatternFlowCfmLbrFlagsChoiceEnum
	DECREMENT PatternFlowCfmLbrFlagsChoiceEnum
}{
	VALUE:     PatternFlowCfmLbrFlagsChoiceEnum("value"),
	VALUES:    PatternFlowCfmLbrFlagsChoiceEnum("values"),
	INCREMENT: PatternFlowCfmLbrFlagsChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmLbrFlagsChoiceEnum("decrement"),
}

func (obj *patternFlowCfmLbrFlags) Choice() PatternFlowCfmLbrFlagsChoiceEnum {
	return PatternFlowCfmLbrFlagsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmLbrFlags) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmLbrFlags) setChoice(value PatternFlowCfmLbrFlagsChoiceEnum) PatternFlowCfmLbrFlags {
	intValue, ok := otg.PatternFlowCfmLbrFlags_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmLbrFlagsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmLbrFlags_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmLbrFlagsChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmLbrFlagsChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmLbrFlagsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmLbrFlagsCounter().msg()
	}

	if value == PatternFlowCfmLbrFlagsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmLbrFlagsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbrFlags) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmLbrFlagsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbrFlags) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmLbrFlags object
func (obj *patternFlowCfmLbrFlags) SetValue(value uint32) PatternFlowCfmLbrFlags {
	obj.setChoice(PatternFlowCfmLbrFlagsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmLbrFlags) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmLbrFlags object
func (obj *patternFlowCfmLbrFlags) SetValues(value []uint32) PatternFlowCfmLbrFlags {
	obj.setChoice(PatternFlowCfmLbrFlagsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmLbrFlagsCounter
func (obj *patternFlowCfmLbrFlags) Increment() PatternFlowCfmLbrFlagsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmLbrFlagsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmLbrFlagsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmLbrFlagsCounter
func (obj *patternFlowCfmLbrFlags) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmLbrFlagsCounter value in the PatternFlowCfmLbrFlags object
func (obj *patternFlowCfmLbrFlags) SetIncrement(value PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlags {
	obj.setChoice(PatternFlowCfmLbrFlagsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmLbrFlagsCounter
func (obj *patternFlowCfmLbrFlags) Decrement() PatternFlowCfmLbrFlagsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmLbrFlagsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmLbrFlagsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmLbrFlagsCounter
func (obj *patternFlowCfmLbrFlags) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmLbrFlagsCounter value in the PatternFlowCfmLbrFlags object
func (obj *patternFlowCfmLbrFlags) SetDecrement(value PatternFlowCfmLbrFlagsCounter) PatternFlowCfmLbrFlags {
	obj.setChoice(PatternFlowCfmLbrFlagsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmLbrFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmLbrFlags.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmLbrFlags.Values <= 255 but Got %d", item))
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

func (obj *patternFlowCfmLbrFlags) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmLbrFlagsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrFlagsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmLbrFlagsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrFlagsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrFlagsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmLbrFlagsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmLbrFlags")
			}
		} else {
			intVal := otg.PatternFlowCfmLbrFlags_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmLbrFlags_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
