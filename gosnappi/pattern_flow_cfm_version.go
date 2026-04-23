package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmVersion *****
type patternFlowCfmVersion struct {
	validation
	obj             *otg.PatternFlowCfmVersion
	marshaller      marshalPatternFlowCfmVersion
	unMarshaller    unMarshalPatternFlowCfmVersion
	incrementHolder PatternFlowCfmVersionCounter
	decrementHolder PatternFlowCfmVersionCounter
}

func NewPatternFlowCfmVersion() PatternFlowCfmVersion {
	obj := patternFlowCfmVersion{obj: &otg.PatternFlowCfmVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmVersion) msg() *otg.PatternFlowCfmVersion {
	return obj.obj
}

func (obj *patternFlowCfmVersion) setMsg(msg *otg.PatternFlowCfmVersion) PatternFlowCfmVersion {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmVersion struct {
	obj *patternFlowCfmVersion
}

type marshalPatternFlowCfmVersion interface {
	// ToProto marshals PatternFlowCfmVersion to protobuf object *otg.PatternFlowCfmVersion
	ToProto() (*otg.PatternFlowCfmVersion, error)
	// ToPbText marshals PatternFlowCfmVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmVersion to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmVersion struct {
	obj *patternFlowCfmVersion
}

type unMarshalPatternFlowCfmVersion interface {
	// FromProto unmarshals PatternFlowCfmVersion from protobuf object *otg.PatternFlowCfmVersion
	FromProto(msg *otg.PatternFlowCfmVersion) (PatternFlowCfmVersion, error)
	// FromPbText unmarshals PatternFlowCfmVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmVersion from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmVersion) Marshal() marshalPatternFlowCfmVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmVersion) Unmarshal() unMarshalPatternFlowCfmVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmVersion) ToProto() (*otg.PatternFlowCfmVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmVersion) FromProto(msg *otg.PatternFlowCfmVersion) (PatternFlowCfmVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmVersion) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersion) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmVersion) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersion) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmVersion) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmVersion) FromJson(value string) error {
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

func (obj *patternFlowCfmVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmVersion) Clone() (PatternFlowCfmVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmVersion()
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

func (obj *patternFlowCfmVersion) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmVersion is indicates the version as specified in 8021Q
type PatternFlowCfmVersion interface {
	Validation
	// msg marshals PatternFlowCfmVersion to protobuf object *otg.PatternFlowCfmVersion
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmVersion
	// setMsg unmarshals PatternFlowCfmVersion from protobuf object *otg.PatternFlowCfmVersion
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmVersion) PatternFlowCfmVersion
	// provides marshal interface
	Marshal() marshalPatternFlowCfmVersion
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmVersion
	// validate validates PatternFlowCfmVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmVersionChoiceEnum, set in PatternFlowCfmVersion
	Choice() PatternFlowCfmVersionChoiceEnum
	// setChoice assigns PatternFlowCfmVersionChoiceEnum provided by user to PatternFlowCfmVersion
	setChoice(value PatternFlowCfmVersionChoiceEnum) PatternFlowCfmVersion
	// HasChoice checks if Choice has been set in PatternFlowCfmVersion
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmVersion.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmVersion
	SetValue(value uint32) PatternFlowCfmVersion
	// HasValue checks if Value has been set in PatternFlowCfmVersion
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmVersion.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmVersion
	SetValues(value []uint32) PatternFlowCfmVersion
	// Increment returns PatternFlowCfmVersionCounter, set in PatternFlowCfmVersion.
	// PatternFlowCfmVersionCounter is integer counter pattern
	Increment() PatternFlowCfmVersionCounter
	// SetIncrement assigns PatternFlowCfmVersionCounter provided by user to PatternFlowCfmVersion.
	// PatternFlowCfmVersionCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmVersionCounter) PatternFlowCfmVersion
	// HasIncrement checks if Increment has been set in PatternFlowCfmVersion
	HasIncrement() bool
	// Decrement returns PatternFlowCfmVersionCounter, set in PatternFlowCfmVersion.
	// PatternFlowCfmVersionCounter is integer counter pattern
	Decrement() PatternFlowCfmVersionCounter
	// SetDecrement assigns PatternFlowCfmVersionCounter provided by user to PatternFlowCfmVersion.
	// PatternFlowCfmVersionCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmVersionCounter) PatternFlowCfmVersion
	// HasDecrement checks if Decrement has been set in PatternFlowCfmVersion
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmVersionChoiceEnum string

// Enum of Choice on PatternFlowCfmVersion
var PatternFlowCfmVersionChoice = struct {
	VALUE     PatternFlowCfmVersionChoiceEnum
	VALUES    PatternFlowCfmVersionChoiceEnum
	INCREMENT PatternFlowCfmVersionChoiceEnum
	DECREMENT PatternFlowCfmVersionChoiceEnum
}{
	VALUE:     PatternFlowCfmVersionChoiceEnum("value"),
	VALUES:    PatternFlowCfmVersionChoiceEnum("values"),
	INCREMENT: PatternFlowCfmVersionChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmVersionChoiceEnum("decrement"),
}

func (obj *patternFlowCfmVersion) Choice() PatternFlowCfmVersionChoiceEnum {
	return PatternFlowCfmVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmVersion) setChoice(value PatternFlowCfmVersionChoiceEnum) PatternFlowCfmVersion {
	intValue, ok := otg.PatternFlowCfmVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmVersionChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmVersionChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmVersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmVersionCounter().msg()
	}

	if value == PatternFlowCfmVersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmVersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmVersion) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmVersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmVersion) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmVersion object
func (obj *patternFlowCfmVersion) SetValue(value uint32) PatternFlowCfmVersion {
	obj.setChoice(PatternFlowCfmVersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmVersion) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmVersion object
func (obj *patternFlowCfmVersion) SetValues(value []uint32) PatternFlowCfmVersion {
	obj.setChoice(PatternFlowCfmVersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmVersionCounter
func (obj *patternFlowCfmVersion) Increment() PatternFlowCfmVersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmVersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmVersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmVersionCounter
func (obj *patternFlowCfmVersion) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmVersionCounter value in the PatternFlowCfmVersion object
func (obj *patternFlowCfmVersion) SetIncrement(value PatternFlowCfmVersionCounter) PatternFlowCfmVersion {
	obj.setChoice(PatternFlowCfmVersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmVersionCounter
func (obj *patternFlowCfmVersion) Decrement() PatternFlowCfmVersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmVersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmVersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmVersionCounter
func (obj *patternFlowCfmVersion) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmVersionCounter value in the PatternFlowCfmVersion object
func (obj *patternFlowCfmVersion) SetDecrement(value PatternFlowCfmVersionCounter) PatternFlowCfmVersion {
	obj.setChoice(PatternFlowCfmVersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowCfmVersion.Value <= 31 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 31 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowCfmVersion.Values <= 31 but Got %d", item))
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

func (obj *patternFlowCfmVersion) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmVersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmVersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmVersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmVersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmVersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmVersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmVersion")
			}
		} else {
			intVal := otg.PatternFlowCfmVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
