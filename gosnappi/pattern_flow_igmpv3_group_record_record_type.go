package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordRecordType *****
type patternFlowIgmpv3GroupRecordRecordType struct {
	validation
	obj             *otg.PatternFlowIgmpv3GroupRecordRecordType
	marshaller      marshalPatternFlowIgmpv3GroupRecordRecordType
	unMarshaller    unMarshalPatternFlowIgmpv3GroupRecordRecordType
	incrementHolder PatternFlowIgmpv3GroupRecordRecordTypeCounter
	decrementHolder PatternFlowIgmpv3GroupRecordRecordTypeCounter
}

func NewPatternFlowIgmpv3GroupRecordRecordType() PatternFlowIgmpv3GroupRecordRecordType {
	obj := patternFlowIgmpv3GroupRecordRecordType{obj: &otg.PatternFlowIgmpv3GroupRecordRecordType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) msg() *otg.PatternFlowIgmpv3GroupRecordRecordType {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordRecordType) PatternFlowIgmpv3GroupRecordRecordType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordRecordType struct {
	obj *patternFlowIgmpv3GroupRecordRecordType
}

type marshalPatternFlowIgmpv3GroupRecordRecordType interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordRecordType to protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordType
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordRecordType, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordRecordType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordRecordType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordRecordType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordRecordType struct {
	obj *patternFlowIgmpv3GroupRecordRecordType
}

type unMarshalPatternFlowIgmpv3GroupRecordRecordType interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordRecordType from protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordType
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordRecordType) (PatternFlowIgmpv3GroupRecordRecordType, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordRecordType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordRecordType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordRecordType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) Marshal() marshalPatternFlowIgmpv3GroupRecordRecordType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordRecordType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordRecordType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordRecordType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordRecordType) ToProto() (*otg.PatternFlowIgmpv3GroupRecordRecordType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordType) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordRecordType) (PatternFlowIgmpv3GroupRecordRecordType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordRecordType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordRecordType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordRecordType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordRecordType) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordRecordType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) Clone() (PatternFlowIgmpv3GroupRecordRecordType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordRecordType()
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

func (obj *patternFlowIgmpv3GroupRecordRecordType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3GroupRecordRecordType is record Type (RFC 9776 Section 4.2.13). Current-State records, sent in reply to a Query, are mode_is_include and mode_is_exclude. Filter-Mode-Change records are change_to_include_mode and change_to_exclude_mode. Source-List-Change records are allow_new_sources and block_old_sources. A change_to_include_mode record with no sources is how a host leaves a group. Receivers silently ignore other values, so any value up to 255 may be sent for negative testing. For a group in the SSM range 232/8 a host should not send mode_is_exclude or change_to_exclude_mode (RFC 4604).
type PatternFlowIgmpv3GroupRecordRecordType interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordRecordType to protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordType
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordRecordType
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordRecordType from protobuf object *otg.PatternFlowIgmpv3GroupRecordRecordType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordRecordType) PatternFlowIgmpv3GroupRecordRecordType
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordRecordType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordRecordType
	// validate validates PatternFlowIgmpv3GroupRecordRecordType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordRecordType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum, set in PatternFlowIgmpv3GroupRecordRecordType
	Choice() PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum
	// setChoice assigns PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum provided by user to PatternFlowIgmpv3GroupRecordRecordType
	setChoice(value PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum) PatternFlowIgmpv3GroupRecordRecordType
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3GroupRecordRecordType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3GroupRecordRecordType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordRecordType
	SetValue(value uint32) PatternFlowIgmpv3GroupRecordRecordType
	// HasValue checks if Value has been set in PatternFlowIgmpv3GroupRecordRecordType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3GroupRecordRecordType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3GroupRecordRecordType
	SetValues(value []uint32) PatternFlowIgmpv3GroupRecordRecordType
	// Increment returns PatternFlowIgmpv3GroupRecordRecordTypeCounter, set in PatternFlowIgmpv3GroupRecordRecordType.
	// PatternFlowIgmpv3GroupRecordRecordTypeCounter is integer counter pattern
	Increment() PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// SetIncrement assigns PatternFlowIgmpv3GroupRecordRecordTypeCounter provided by user to PatternFlowIgmpv3GroupRecordRecordType.
	// PatternFlowIgmpv3GroupRecordRecordTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordType
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3GroupRecordRecordType
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3GroupRecordRecordTypeCounter, set in PatternFlowIgmpv3GroupRecordRecordType.
	// PatternFlowIgmpv3GroupRecordRecordTypeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3GroupRecordRecordTypeCounter
	// SetDecrement assigns PatternFlowIgmpv3GroupRecordRecordTypeCounter provided by user to PatternFlowIgmpv3GroupRecordRecordType.
	// PatternFlowIgmpv3GroupRecordRecordTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordType
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3GroupRecordRecordType
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3GroupRecordRecordType
var PatternFlowIgmpv3GroupRecordRecordTypeChoice = struct {
	VALUE     PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum
	VALUES    PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum
	INCREMENT PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum
	DECREMENT PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) Choice() PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum {
	return PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3GroupRecordRecordType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) setChoice(value PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum) PatternFlowIgmpv3GroupRecordRecordType {
	intValue, ok := otg.PatternFlowIgmpv3GroupRecordRecordType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3GroupRecordRecordType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordRecordTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3GroupRecordRecordTypeCounter().msg()
	}

	if value == PatternFlowIgmpv3GroupRecordRecordTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3GroupRecordRecordTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordRecordType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3GroupRecordRecordType object
func (obj *patternFlowIgmpv3GroupRecordRecordType) SetValue(value uint32) PatternFlowIgmpv3GroupRecordRecordType {
	obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3GroupRecordRecordType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3GroupRecordRecordType object
func (obj *patternFlowIgmpv3GroupRecordRecordType) SetValues(value []uint32) PatternFlowIgmpv3GroupRecordRecordType {
	obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordRecordTypeCounter
func (obj *patternFlowIgmpv3GroupRecordRecordType) Increment() PatternFlowIgmpv3GroupRecordRecordTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3GroupRecordRecordTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordRecordTypeCounter
func (obj *patternFlowIgmpv3GroupRecordRecordType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3GroupRecordRecordTypeCounter value in the PatternFlowIgmpv3GroupRecordRecordType object
func (obj *patternFlowIgmpv3GroupRecordRecordType) SetIncrement(value PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordType {
	obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordRecordTypeCounter
func (obj *patternFlowIgmpv3GroupRecordRecordType) Decrement() PatternFlowIgmpv3GroupRecordRecordTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3GroupRecordRecordTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordRecordTypeCounter
func (obj *patternFlowIgmpv3GroupRecordRecordType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3GroupRecordRecordTypeCounter value in the PatternFlowIgmpv3GroupRecordRecordType object
func (obj *patternFlowIgmpv3GroupRecordRecordType) SetDecrement(value PatternFlowIgmpv3GroupRecordRecordTypeCounter) PatternFlowIgmpv3GroupRecordRecordType {
	obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3GroupRecordRecordType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordRecordType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3GroupRecordRecordType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3GroupRecordRecordType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3GroupRecordRecordTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordRecordTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordRecordTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3GroupRecordRecordTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3GroupRecordRecordType")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3GroupRecordRecordType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3GroupRecordRecordType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
