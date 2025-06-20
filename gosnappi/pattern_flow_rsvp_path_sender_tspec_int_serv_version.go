package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServVersion *****
type patternFlowRSVPPathSenderTspecIntServVersion struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServVersion
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServVersion
	incrementHolder PatternFlowRSVPPathSenderTspecIntServVersionCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServVersionCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServVersion() PatternFlowRSVPPathSenderTspecIntServVersion {
	obj := patternFlowRSVPPathSenderTspecIntServVersion{obj: &otg.PatternFlowRSVPPathSenderTspecIntServVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) msg() *otg.PatternFlowRSVPPathSenderTspecIntServVersion {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersion) PatternFlowRSVPPathSenderTspecIntServVersion {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServVersion struct {
	obj *patternFlowRSVPPathSenderTspecIntServVersion
}

type marshalPatternFlowRSVPPathSenderTspecIntServVersion interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServVersion to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServVersion, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServVersion to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServVersion struct {
	obj *patternFlowRSVPPathSenderTspecIntServVersion
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServVersion interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServVersion from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersion) (PatternFlowRSVPPathSenderTspecIntServVersion, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServVersion from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersion) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersion) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServVersion) (PatternFlowRSVPPathSenderTspecIntServVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersion) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersion) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersion) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersion) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServVersion) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServVersion) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Clone() (PatternFlowRSVPPathSenderTspecIntServVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServVersion()
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

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServVersion is message format version number.
type PatternFlowRSVPPathSenderTspecIntServVersion interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServVersion to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServVersion from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServVersion
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServVersion) PatternFlowRSVPPathSenderTspecIntServVersion
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServVersion
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServVersion
	// validate validates PatternFlowRSVPPathSenderTspecIntServVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServVersion
	Choice() PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServVersion
	setChoice(value PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum) PatternFlowRSVPPathSenderTspecIntServVersion
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServVersion
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServVersion.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServVersion
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServVersion
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServVersion
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServVersion.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServVersion
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServVersion
	// Increment returns PatternFlowRSVPPathSenderTspecIntServVersionCounter, set in PatternFlowRSVPPathSenderTspecIntServVersion.
	// PatternFlowRSVPPathSenderTspecIntServVersionCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServVersionCounter provided by user to PatternFlowRSVPPathSenderTspecIntServVersion.
	// PatternFlowRSVPPathSenderTspecIntServVersionCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersion
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServVersion
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServVersionCounter, set in PatternFlowRSVPPathSenderTspecIntServVersion.
	// PatternFlowRSVPPathSenderTspecIntServVersionCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServVersionCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServVersionCounter provided by user to PatternFlowRSVPPathSenderTspecIntServVersion.
	// PatternFlowRSVPPathSenderTspecIntServVersionCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersion
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServVersion
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServVersion
var PatternFlowRSVPPathSenderTspecIntServVersionChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Choice() PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) setChoice(value PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum) PatternFlowRSVPPathSenderTspecIntServVersion {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServVersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServVersionCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServVersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServVersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServVersion object
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServVersion {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServVersion object
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServVersion {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServVersionCounter
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Increment() PatternFlowRSVPPathSenderTspecIntServVersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServVersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServVersionCounter
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServVersionCounter value in the PatternFlowRSVPPathSenderTspecIntServVersion object
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersion {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServVersionCounter
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) Decrement() PatternFlowRSVPPathSenderTspecIntServVersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServVersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServVersionCounter
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServVersionCounter value in the PatternFlowRSVPPathSenderTspecIntServVersion object
func (obj *patternFlowRSVPPathSenderTspecIntServVersion) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServVersionCounter) PatternFlowRSVPPathSenderTspecIntServVersion {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServVersion.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServVersion.Values <= 15 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServVersion) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServVersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServVersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServVersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServVersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServVersion")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
