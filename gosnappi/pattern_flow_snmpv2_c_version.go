package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CVersion *****
type patternFlowSnmpv2CVersion struct {
	validation
	obj             *otg.PatternFlowSnmpv2CVersion
	marshaller      marshalPatternFlowSnmpv2CVersion
	unMarshaller    unMarshalPatternFlowSnmpv2CVersion
	incrementHolder PatternFlowSnmpv2CVersionCounter
	decrementHolder PatternFlowSnmpv2CVersionCounter
}

func NewPatternFlowSnmpv2CVersion() PatternFlowSnmpv2CVersion {
	obj := patternFlowSnmpv2CVersion{obj: &otg.PatternFlowSnmpv2CVersion{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CVersion) msg() *otg.PatternFlowSnmpv2CVersion {
	return obj.obj
}

func (obj *patternFlowSnmpv2CVersion) setMsg(msg *otg.PatternFlowSnmpv2CVersion) PatternFlowSnmpv2CVersion {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CVersion struct {
	obj *patternFlowSnmpv2CVersion
}

type marshalPatternFlowSnmpv2CVersion interface {
	// ToProto marshals PatternFlowSnmpv2CVersion to protobuf object *otg.PatternFlowSnmpv2CVersion
	ToProto() (*otg.PatternFlowSnmpv2CVersion, error)
	// ToPbText marshals PatternFlowSnmpv2CVersion to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CVersion to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CVersion to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowSnmpv2CVersion struct {
	obj *patternFlowSnmpv2CVersion
}

type unMarshalPatternFlowSnmpv2CVersion interface {
	// FromProto unmarshals PatternFlowSnmpv2CVersion from protobuf object *otg.PatternFlowSnmpv2CVersion
	FromProto(msg *otg.PatternFlowSnmpv2CVersion) (PatternFlowSnmpv2CVersion, error)
	// FromPbText unmarshals PatternFlowSnmpv2CVersion from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CVersion from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CVersion from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CVersion) Marshal() marshalPatternFlowSnmpv2CVersion {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CVersion{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CVersion) Unmarshal() unMarshalPatternFlowSnmpv2CVersion {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CVersion{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CVersion) ToProto() (*otg.PatternFlowSnmpv2CVersion, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CVersion) FromProto(msg *otg.PatternFlowSnmpv2CVersion) (PatternFlowSnmpv2CVersion, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CVersion) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersion) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVersion) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersion) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CVersion) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CVersion) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CVersion) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVersion) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CVersion) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CVersion) Clone() (PatternFlowSnmpv2CVersion, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CVersion()
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

func (obj *patternFlowSnmpv2CVersion) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CVersion is version
type PatternFlowSnmpv2CVersion interface {
	Validation
	// msg marshals PatternFlowSnmpv2CVersion to protobuf object *otg.PatternFlowSnmpv2CVersion
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CVersion
	// setMsg unmarshals PatternFlowSnmpv2CVersion from protobuf object *otg.PatternFlowSnmpv2CVersion
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CVersion) PatternFlowSnmpv2CVersion
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CVersion
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CVersion
	// validate validates PatternFlowSnmpv2CVersion
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CVersion, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CVersionChoiceEnum, set in PatternFlowSnmpv2CVersion
	Choice() PatternFlowSnmpv2CVersionChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CVersionChoiceEnum provided by user to PatternFlowSnmpv2CVersion
	setChoice(value PatternFlowSnmpv2CVersionChoiceEnum) PatternFlowSnmpv2CVersion
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CVersion
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CVersion.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CVersion
	SetValue(value uint32) PatternFlowSnmpv2CVersion
	// HasValue checks if Value has been set in PatternFlowSnmpv2CVersion
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CVersion.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CVersion
	SetValues(value []uint32) PatternFlowSnmpv2CVersion
	// Increment returns PatternFlowSnmpv2CVersionCounter, set in PatternFlowSnmpv2CVersion.
	Increment() PatternFlowSnmpv2CVersionCounter
	// SetIncrement assigns PatternFlowSnmpv2CVersionCounter provided by user to PatternFlowSnmpv2CVersion.
	SetIncrement(value PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersion
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CVersion
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CVersionCounter, set in PatternFlowSnmpv2CVersion.
	Decrement() PatternFlowSnmpv2CVersionCounter
	// SetDecrement assigns PatternFlowSnmpv2CVersionCounter provided by user to PatternFlowSnmpv2CVersion.
	SetDecrement(value PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersion
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CVersion
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CVersionChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CVersion
var PatternFlowSnmpv2CVersionChoice = struct {
	VALUE     PatternFlowSnmpv2CVersionChoiceEnum
	VALUES    PatternFlowSnmpv2CVersionChoiceEnum
	INCREMENT PatternFlowSnmpv2CVersionChoiceEnum
	DECREMENT PatternFlowSnmpv2CVersionChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CVersionChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CVersionChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CVersionChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CVersionChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CVersion) Choice() PatternFlowSnmpv2CVersionChoiceEnum {
	return PatternFlowSnmpv2CVersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CVersion) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CVersion) setChoice(value PatternFlowSnmpv2CVersionChoiceEnum) PatternFlowSnmpv2CVersion {
	intValue, ok := otg.PatternFlowSnmpv2CVersion_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CVersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CVersion_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CVersionChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CVersionChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CVersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CVersionCounter().msg()
	}

	if value == PatternFlowSnmpv2CVersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CVersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVersion) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CVersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CVersion) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CVersion object
func (obj *patternFlowSnmpv2CVersion) SetValue(value uint32) PatternFlowSnmpv2CVersion {
	obj.setChoice(PatternFlowSnmpv2CVersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CVersion) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CVersion object
func (obj *patternFlowSnmpv2CVersion) SetValues(value []uint32) PatternFlowSnmpv2CVersion {
	obj.setChoice(PatternFlowSnmpv2CVersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVersionCounter
func (obj *patternFlowSnmpv2CVersion) Increment() PatternFlowSnmpv2CVersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CVersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CVersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CVersionCounter
func (obj *patternFlowSnmpv2CVersion) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CVersionCounter value in the PatternFlowSnmpv2CVersion object
func (obj *patternFlowSnmpv2CVersion) SetIncrement(value PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersion {
	obj.setChoice(PatternFlowSnmpv2CVersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVersionCounter
func (obj *patternFlowSnmpv2CVersion) Decrement() PatternFlowSnmpv2CVersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CVersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CVersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CVersionCounter
func (obj *patternFlowSnmpv2CVersion) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CVersionCounter value in the PatternFlowSnmpv2CVersion object
func (obj *patternFlowSnmpv2CVersion) SetDecrement(value PatternFlowSnmpv2CVersionCounter) PatternFlowSnmpv2CVersion {
	obj.setChoice(PatternFlowSnmpv2CVersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CVersion) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowSnmpv2CVersion.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowSnmpv2CVersion.Values <= 255 but Got %d", item))
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

func (obj *patternFlowSnmpv2CVersion) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CVersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CVersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CVersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CVersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CVersion")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CVersion_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CVersion_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
