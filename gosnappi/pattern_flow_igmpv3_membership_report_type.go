package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipReportType *****
type patternFlowIgmpv3MembershipReportType struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipReportType
	marshaller      marshalPatternFlowIgmpv3MembershipReportType
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipReportType
	incrementHolder PatternFlowIgmpv3MembershipReportTypeCounter
	decrementHolder PatternFlowIgmpv3MembershipReportTypeCounter
}

func NewPatternFlowIgmpv3MembershipReportType() PatternFlowIgmpv3MembershipReportType {
	obj := patternFlowIgmpv3MembershipReportType{obj: &otg.PatternFlowIgmpv3MembershipReportType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipReportType) msg() *otg.PatternFlowIgmpv3MembershipReportType {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipReportType) setMsg(msg *otg.PatternFlowIgmpv3MembershipReportType) PatternFlowIgmpv3MembershipReportType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipReportType struct {
	obj *patternFlowIgmpv3MembershipReportType
}

type marshalPatternFlowIgmpv3MembershipReportType interface {
	// ToProto marshals PatternFlowIgmpv3MembershipReportType to protobuf object *otg.PatternFlowIgmpv3MembershipReportType
	ToProto() (*otg.PatternFlowIgmpv3MembershipReportType, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipReportType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipReportType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipReportType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipReportType struct {
	obj *patternFlowIgmpv3MembershipReportType
}

type unMarshalPatternFlowIgmpv3MembershipReportType interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipReportType from protobuf object *otg.PatternFlowIgmpv3MembershipReportType
	FromProto(msg *otg.PatternFlowIgmpv3MembershipReportType) (PatternFlowIgmpv3MembershipReportType, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipReportType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipReportType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipReportType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipReportType) Marshal() marshalPatternFlowIgmpv3MembershipReportType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipReportType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipReportType) Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipReportType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipReportType) ToProto() (*otg.PatternFlowIgmpv3MembershipReportType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipReportType) FromProto(msg *otg.PatternFlowIgmpv3MembershipReportType) (PatternFlowIgmpv3MembershipReportType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipReportType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipReportType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipReportType) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipReportType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipReportType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipReportType) Clone() (PatternFlowIgmpv3MembershipReportType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipReportType()
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

func (obj *patternFlowIgmpv3MembershipReportType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipReportType is iGMP message type, 34 (0x22) for a Version 3 Membership Report.
type PatternFlowIgmpv3MembershipReportType interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipReportType to protobuf object *otg.PatternFlowIgmpv3MembershipReportType
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipReportType
	// setMsg unmarshals PatternFlowIgmpv3MembershipReportType from protobuf object *otg.PatternFlowIgmpv3MembershipReportType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipReportType) PatternFlowIgmpv3MembershipReportType
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipReportType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipReportType
	// validate validates PatternFlowIgmpv3MembershipReportType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipReportType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipReportTypeChoiceEnum, set in PatternFlowIgmpv3MembershipReportType
	Choice() PatternFlowIgmpv3MembershipReportTypeChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipReportTypeChoiceEnum provided by user to PatternFlowIgmpv3MembershipReportType
	setChoice(value PatternFlowIgmpv3MembershipReportTypeChoiceEnum) PatternFlowIgmpv3MembershipReportType
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipReportType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipReportType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipReportType
	SetValue(value uint32) PatternFlowIgmpv3MembershipReportType
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipReportType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipReportType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipReportType
	SetValues(value []uint32) PatternFlowIgmpv3MembershipReportType
	// Increment returns PatternFlowIgmpv3MembershipReportTypeCounter, set in PatternFlowIgmpv3MembershipReportType.
	// PatternFlowIgmpv3MembershipReportTypeCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipReportTypeCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipReportTypeCounter provided by user to PatternFlowIgmpv3MembershipReportType.
	// PatternFlowIgmpv3MembershipReportTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportType
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipReportType
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipReportTypeCounter, set in PatternFlowIgmpv3MembershipReportType.
	// PatternFlowIgmpv3MembershipReportTypeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipReportTypeCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipReportTypeCounter provided by user to PatternFlowIgmpv3MembershipReportType.
	// PatternFlowIgmpv3MembershipReportTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportType
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipReportType
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipReportTypeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipReportType
var PatternFlowIgmpv3MembershipReportTypeChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipReportTypeChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipReportTypeChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipReportTypeChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipReportTypeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipReportTypeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipReportTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipReportTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipReportTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipReportType) Choice() PatternFlowIgmpv3MembershipReportTypeChoiceEnum {
	return PatternFlowIgmpv3MembershipReportTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipReportType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipReportType) setChoice(value PatternFlowIgmpv3MembershipReportTypeChoiceEnum) PatternFlowIgmpv3MembershipReportType {
	intValue, ok := otg.PatternFlowIgmpv3MembershipReportType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipReportTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipReportType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipReportTypeChoice.VALUE {
		defaultValue := uint32(34)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportTypeChoice.VALUES {
		defaultValue := []uint32{34}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipReportTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipReportTypeCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipReportTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipReportTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipReportType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipReportType object
func (obj *patternFlowIgmpv3MembershipReportType) SetValue(value uint32) PatternFlowIgmpv3MembershipReportType {
	obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipReportType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{34})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipReportType object
func (obj *patternFlowIgmpv3MembershipReportType) SetValues(value []uint32) PatternFlowIgmpv3MembershipReportType {
	obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportTypeCounter
func (obj *patternFlowIgmpv3MembershipReportType) Increment() PatternFlowIgmpv3MembershipReportTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipReportTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipReportTypeCounter
func (obj *patternFlowIgmpv3MembershipReportType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipReportTypeCounter value in the PatternFlowIgmpv3MembershipReportType object
func (obj *patternFlowIgmpv3MembershipReportType) SetIncrement(value PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportType {
	obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportTypeCounter
func (obj *patternFlowIgmpv3MembershipReportType) Decrement() PatternFlowIgmpv3MembershipReportTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipReportTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipReportTypeCounter
func (obj *patternFlowIgmpv3MembershipReportType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipReportTypeCounter value in the PatternFlowIgmpv3MembershipReportType object
func (obj *patternFlowIgmpv3MembershipReportType) SetDecrement(value PatternFlowIgmpv3MembershipReportTypeCounter) PatternFlowIgmpv3MembershipReportType {
	obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipReportType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipReportType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipReportType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipReportType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipReportTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipReportTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipReportTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipReportType")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipReportType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipReportType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
