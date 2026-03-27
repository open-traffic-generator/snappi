package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpVersion *****
type patternFlowLacpVersion struct {
	validation
	obj             *otg.PatternFlowLacpVersion
	marshaller      marshalPatternFlowLacpVersion
	unMarshaller    unMarshalPatternFlowLacpVersion
	incrementHolder PatternFlowLacpVersionCounter
	decrementHolder PatternFlowLacpVersionCounter
}

func NewPatternFlowLacpVersion() PatternFlowLacpVersion {
	obj := patternFlowLacpVersion{obj: &otg.PatternFlowLacpVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpVersion) msg() *otg.PatternFlowLacpVersion {
	return obj.obj
}

func (obj *patternFlowLacpVersion) setMsg(msg *otg.PatternFlowLacpVersion) PatternFlowLacpVersion {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpVersion struct {
	obj *patternFlowLacpVersion
}

type marshalPatternFlowLacpVersion interface {
	// ToProto marshals PatternFlowLacpVersion to protobuf object *otg.PatternFlowLacpVersion
	ToProto() (*otg.PatternFlowLacpVersion, error)
	// ToPbText marshals PatternFlowLacpVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpVersion to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpVersion struct {
	obj *patternFlowLacpVersion
}

type unMarshalPatternFlowLacpVersion interface {
	// FromProto unmarshals PatternFlowLacpVersion from protobuf object *otg.PatternFlowLacpVersion
	FromProto(msg *otg.PatternFlowLacpVersion) (PatternFlowLacpVersion, error)
	// FromPbText unmarshals PatternFlowLacpVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpVersion from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpVersion) Marshal() marshalPatternFlowLacpVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpVersion) Unmarshal() unMarshalPatternFlowLacpVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpVersion) ToProto() (*otg.PatternFlowLacpVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpVersion) FromProto(msg *otg.PatternFlowLacpVersion) (PatternFlowLacpVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpVersion) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersion) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpVersion) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersion) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpVersion) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpVersion) FromJson(value string) error {
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

func (obj *patternFlowLacpVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpVersion) Clone() (PatternFlowLacpVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpVersion()
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

func (obj *patternFlowLacpVersion) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpVersion is the LACP version number.
type PatternFlowLacpVersion interface {
	Validation
	// msg marshals PatternFlowLacpVersion to protobuf object *otg.PatternFlowLacpVersion
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpVersion
	// setMsg unmarshals PatternFlowLacpVersion from protobuf object *otg.PatternFlowLacpVersion
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpVersion) PatternFlowLacpVersion
	// provides marshal interface
	Marshal() marshalPatternFlowLacpVersion
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpVersion
	// validate validates PatternFlowLacpVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpVersionChoiceEnum, set in PatternFlowLacpVersion
	Choice() PatternFlowLacpVersionChoiceEnum
	// setChoice assigns PatternFlowLacpVersionChoiceEnum provided by user to PatternFlowLacpVersion
	setChoice(value PatternFlowLacpVersionChoiceEnum) PatternFlowLacpVersion
	// HasChoice checks if Choice has been set in PatternFlowLacpVersion
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpVersion.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpVersion
	SetValue(value uint32) PatternFlowLacpVersion
	// HasValue checks if Value has been set in PatternFlowLacpVersion
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpVersion.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpVersion
	SetValues(value []uint32) PatternFlowLacpVersion
	// Increment returns PatternFlowLacpVersionCounter, set in PatternFlowLacpVersion.
	// PatternFlowLacpVersionCounter is integer counter pattern
	Increment() PatternFlowLacpVersionCounter
	// SetIncrement assigns PatternFlowLacpVersionCounter provided by user to PatternFlowLacpVersion.
	// PatternFlowLacpVersionCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpVersionCounter) PatternFlowLacpVersion
	// HasIncrement checks if Increment has been set in PatternFlowLacpVersion
	HasIncrement() bool
	// Decrement returns PatternFlowLacpVersionCounter, set in PatternFlowLacpVersion.
	// PatternFlowLacpVersionCounter is integer counter pattern
	Decrement() PatternFlowLacpVersionCounter
	// SetDecrement assigns PatternFlowLacpVersionCounter provided by user to PatternFlowLacpVersion.
	// PatternFlowLacpVersionCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpVersionCounter) PatternFlowLacpVersion
	// HasDecrement checks if Decrement has been set in PatternFlowLacpVersion
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpVersionChoiceEnum string

// Enum of Choice on PatternFlowLacpVersion
var PatternFlowLacpVersionChoice = struct {
	VALUE     PatternFlowLacpVersionChoiceEnum
	VALUES    PatternFlowLacpVersionChoiceEnum
	INCREMENT PatternFlowLacpVersionChoiceEnum
	DECREMENT PatternFlowLacpVersionChoiceEnum
}{
	VALUE:     PatternFlowLacpVersionChoiceEnum("value"),
	VALUES:    PatternFlowLacpVersionChoiceEnum("values"),
	INCREMENT: PatternFlowLacpVersionChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpVersionChoiceEnum("decrement"),
}

func (obj *patternFlowLacpVersion) Choice() PatternFlowLacpVersionChoiceEnum {
	return PatternFlowLacpVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpVersion) setChoice(value PatternFlowLacpVersionChoiceEnum) PatternFlowLacpVersion {
	intValue, ok := otg.PatternFlowLacpVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpVersionChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpVersionChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpVersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpVersionCounter().msg()
	}

	if value == PatternFlowLacpVersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpVersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpVersion) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpVersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpVersion) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpVersion object
func (obj *patternFlowLacpVersion) SetValue(value uint32) PatternFlowLacpVersion {
	obj.setChoice(PatternFlowLacpVersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpVersion) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpVersion object
func (obj *patternFlowLacpVersion) SetValues(value []uint32) PatternFlowLacpVersion {
	obj.setChoice(PatternFlowLacpVersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpVersionCounter
func (obj *patternFlowLacpVersion) Increment() PatternFlowLacpVersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpVersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpVersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpVersionCounter
func (obj *patternFlowLacpVersion) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpVersionCounter value in the PatternFlowLacpVersion object
func (obj *patternFlowLacpVersion) SetIncrement(value PatternFlowLacpVersionCounter) PatternFlowLacpVersion {
	obj.setChoice(PatternFlowLacpVersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpVersionCounter
func (obj *patternFlowLacpVersion) Decrement() PatternFlowLacpVersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpVersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpVersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpVersionCounter
func (obj *patternFlowLacpVersion) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpVersionCounter value in the PatternFlowLacpVersion object
func (obj *patternFlowLacpVersion) SetDecrement(value PatternFlowLacpVersionCounter) PatternFlowLacpVersion {
	obj.setChoice(PatternFlowLacpVersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpVersion.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpVersion.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpVersion) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpVersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpVersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpVersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpVersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpVersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpVersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpVersion")
			}
		} else {
			intVal := otg.PatternFlowLacpVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
