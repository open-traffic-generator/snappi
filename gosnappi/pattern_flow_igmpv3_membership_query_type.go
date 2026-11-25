package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryType *****
type patternFlowIgmpv3MembershipQueryType struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryType
	marshaller      marshalPatternFlowIgmpv3MembershipQueryType
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryType
	incrementHolder PatternFlowIgmpv3MembershipQueryTypeCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryTypeCounter
}

func NewPatternFlowIgmpv3MembershipQueryType() PatternFlowIgmpv3MembershipQueryType {
	obj := patternFlowIgmpv3MembershipQueryType{obj: &otg.PatternFlowIgmpv3MembershipQueryType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryType) msg() *otg.PatternFlowIgmpv3MembershipQueryType {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryType) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryType) PatternFlowIgmpv3MembershipQueryType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryType struct {
	obj *patternFlowIgmpv3MembershipQueryType
}

type marshalPatternFlowIgmpv3MembershipQueryType interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryType to protobuf object *otg.PatternFlowIgmpv3MembershipQueryType
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryType, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryType struct {
	obj *patternFlowIgmpv3MembershipQueryType
}

type unMarshalPatternFlowIgmpv3MembershipQueryType interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryType from protobuf object *otg.PatternFlowIgmpv3MembershipQueryType
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryType) (PatternFlowIgmpv3MembershipQueryType, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryType) Marshal() marshalPatternFlowIgmpv3MembershipQueryType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryType) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryType) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryType) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryType) (PatternFlowIgmpv3MembershipQueryType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryType) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryType) Clone() (PatternFlowIgmpv3MembershipQueryType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryType()
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

func (obj *patternFlowIgmpv3MembershipQueryType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryType is iGMP message type (RFC 9776 Section 4.1). Default 17 (0x11), Membership Query. This header always sends the IGMPv3 layout, the eight octet IGMPv1 and IGMPv2 messages use the igmpv1 and igmpv2 headers. Another value sends this layout with that type for negative testing, receivers silently ignore unrecognized types.
type PatternFlowIgmpv3MembershipQueryType interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryType to protobuf object *otg.PatternFlowIgmpv3MembershipQueryType
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryType
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryType from protobuf object *otg.PatternFlowIgmpv3MembershipQueryType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryType) PatternFlowIgmpv3MembershipQueryType
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryType
	// validate validates PatternFlowIgmpv3MembershipQueryType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryTypeChoiceEnum, set in PatternFlowIgmpv3MembershipQueryType
	Choice() PatternFlowIgmpv3MembershipQueryTypeChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryTypeChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryType
	setChoice(value PatternFlowIgmpv3MembershipQueryTypeChoiceEnum) PatternFlowIgmpv3MembershipQueryType
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryType
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryType
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryType
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryType
	// Increment returns PatternFlowIgmpv3MembershipQueryTypeCounter, set in PatternFlowIgmpv3MembershipQueryType.
	// PatternFlowIgmpv3MembershipQueryTypeCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryTypeCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryTypeCounter provided by user to PatternFlowIgmpv3MembershipQueryType.
	// PatternFlowIgmpv3MembershipQueryTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryType
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryType
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryTypeCounter, set in PatternFlowIgmpv3MembershipQueryType.
	// PatternFlowIgmpv3MembershipQueryTypeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryTypeCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryTypeCounter provided by user to PatternFlowIgmpv3MembershipQueryType.
	// PatternFlowIgmpv3MembershipQueryTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryType
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryType
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryTypeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryType
var PatternFlowIgmpv3MembershipQueryTypeChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryTypeChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryTypeChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryTypeChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryTypeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryTypeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryType) Choice() PatternFlowIgmpv3MembershipQueryTypeChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryType) setChoice(value PatternFlowIgmpv3MembershipQueryTypeChoiceEnum) PatternFlowIgmpv3MembershipQueryType {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryTypeChoice.VALUE {
		defaultValue := uint32(17)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryTypeChoice.VALUES {
		defaultValue := []uint32{17}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryTypeCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryType object
func (obj *patternFlowIgmpv3MembershipQueryType) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryType {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{17})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryType object
func (obj *patternFlowIgmpv3MembershipQueryType) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryType {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryTypeCounter
func (obj *patternFlowIgmpv3MembershipQueryType) Increment() PatternFlowIgmpv3MembershipQueryTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryTypeCounter
func (obj *patternFlowIgmpv3MembershipQueryType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryTypeCounter value in the PatternFlowIgmpv3MembershipQueryType object
func (obj *patternFlowIgmpv3MembershipQueryType) SetIncrement(value PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryType {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryTypeCounter
func (obj *patternFlowIgmpv3MembershipQueryType) Decrement() PatternFlowIgmpv3MembershipQueryTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryTypeCounter
func (obj *patternFlowIgmpv3MembershipQueryType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryTypeCounter value in the PatternFlowIgmpv3MembershipQueryType object
func (obj *patternFlowIgmpv3MembershipQueryType) SetDecrement(value PatternFlowIgmpv3MembershipQueryTypeCounter) PatternFlowIgmpv3MembershipQueryType {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQueryType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryType")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
