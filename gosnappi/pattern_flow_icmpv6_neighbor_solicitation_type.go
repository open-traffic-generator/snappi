package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NeighborSolicitationType *****
type patternFlowIcmpv6NeighborSolicitationType struct {
	validation
	obj             *otg.PatternFlowIcmpv6NeighborSolicitationType
	marshaller      marshalPatternFlowIcmpv6NeighborSolicitationType
	unMarshaller    unMarshalPatternFlowIcmpv6NeighborSolicitationType
	incrementHolder PatternFlowIcmpv6NeighborSolicitationTypeCounter
	decrementHolder PatternFlowIcmpv6NeighborSolicitationTypeCounter
}

func NewPatternFlowIcmpv6NeighborSolicitationType() PatternFlowIcmpv6NeighborSolicitationType {
	obj := patternFlowIcmpv6NeighborSolicitationType{obj: &otg.PatternFlowIcmpv6NeighborSolicitationType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) msg() *otg.PatternFlowIcmpv6NeighborSolicitationType {
	return obj.obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) setMsg(msg *otg.PatternFlowIcmpv6NeighborSolicitationType) PatternFlowIcmpv6NeighborSolicitationType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NeighborSolicitationType struct {
	obj *patternFlowIcmpv6NeighborSolicitationType
}

type marshalPatternFlowIcmpv6NeighborSolicitationType interface {
	// ToProto marshals PatternFlowIcmpv6NeighborSolicitationType to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationType
	ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationType, error)
	// ToPbText marshals PatternFlowIcmpv6NeighborSolicitationType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NeighborSolicitationType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NeighborSolicitationType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NeighborSolicitationType struct {
	obj *patternFlowIcmpv6NeighborSolicitationType
}

type unMarshalPatternFlowIcmpv6NeighborSolicitationType interface {
	// FromProto unmarshals PatternFlowIcmpv6NeighborSolicitationType from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationType
	FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationType) (PatternFlowIcmpv6NeighborSolicitationType, error)
	// FromPbText unmarshals PatternFlowIcmpv6NeighborSolicitationType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NeighborSolicitationType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NeighborSolicitationType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) Marshal() marshalPatternFlowIcmpv6NeighborSolicitationType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NeighborSolicitationType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NeighborSolicitationType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationType) ToProto() (*otg.PatternFlowIcmpv6NeighborSolicitationType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationType) FromProto(msg *otg.PatternFlowIcmpv6NeighborSolicitationType) (PatternFlowIcmpv6NeighborSolicitationType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NeighborSolicitationType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NeighborSolicitationType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NeighborSolicitationType) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NeighborSolicitationType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) Clone() (PatternFlowIcmpv6NeighborSolicitationType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NeighborSolicitationType()
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

func (obj *patternFlowIcmpv6NeighborSolicitationType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NeighborSolicitationType is iCMPv6 message type. Fixed at 135 for Neighbor Solicitation.
type PatternFlowIcmpv6NeighborSolicitationType interface {
	Validation
	// msg marshals PatternFlowIcmpv6NeighborSolicitationType to protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationType
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NeighborSolicitationType
	// setMsg unmarshals PatternFlowIcmpv6NeighborSolicitationType from protobuf object *otg.PatternFlowIcmpv6NeighborSolicitationType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NeighborSolicitationType) PatternFlowIcmpv6NeighborSolicitationType
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NeighborSolicitationType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NeighborSolicitationType
	// validate validates PatternFlowIcmpv6NeighborSolicitationType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NeighborSolicitationType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum, set in PatternFlowIcmpv6NeighborSolicitationType
	Choice() PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum provided by user to PatternFlowIcmpv6NeighborSolicitationType
	setChoice(value PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum) PatternFlowIcmpv6NeighborSolicitationType
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NeighborSolicitationType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NeighborSolicitationType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationType
	SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationType
	// HasValue checks if Value has been set in PatternFlowIcmpv6NeighborSolicitationType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NeighborSolicitationType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NeighborSolicitationType
	SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationType
	// Increment returns PatternFlowIcmpv6NeighborSolicitationTypeCounter, set in PatternFlowIcmpv6NeighborSolicitationType.
	// PatternFlowIcmpv6NeighborSolicitationTypeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// SetIncrement assigns PatternFlowIcmpv6NeighborSolicitationTypeCounter provided by user to PatternFlowIcmpv6NeighborSolicitationType.
	// PatternFlowIcmpv6NeighborSolicitationTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationType
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NeighborSolicitationType
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NeighborSolicitationTypeCounter, set in PatternFlowIcmpv6NeighborSolicitationType.
	// PatternFlowIcmpv6NeighborSolicitationTypeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NeighborSolicitationTypeCounter
	// SetDecrement assigns PatternFlowIcmpv6NeighborSolicitationTypeCounter provided by user to PatternFlowIcmpv6NeighborSolicitationType.
	// PatternFlowIcmpv6NeighborSolicitationTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationType
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NeighborSolicitationType
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NeighborSolicitationType
var PatternFlowIcmpv6NeighborSolicitationTypeChoice = struct {
	VALUE     PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum
	VALUES    PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum
	INCREMENT PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum
	DECREMENT PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) Choice() PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum {
	return PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NeighborSolicitationType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) setChoice(value PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum) PatternFlowIcmpv6NeighborSolicitationType {
	intValue, ok := otg.PatternFlowIcmpv6NeighborSolicitationType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NeighborSolicitationType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUE {
		defaultValue := uint32(135)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUES {
		defaultValue := []uint32{135}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NeighborSolicitationTypeCounter().msg()
	}

	if value == PatternFlowIcmpv6NeighborSolicitationTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NeighborSolicitationTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NeighborSolicitationType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NeighborSolicitationType object
func (obj *patternFlowIcmpv6NeighborSolicitationType) SetValue(value uint32) PatternFlowIcmpv6NeighborSolicitationType {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NeighborSolicitationType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{135})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NeighborSolicitationType object
func (obj *patternFlowIcmpv6NeighborSolicitationType) SetValues(value []uint32) PatternFlowIcmpv6NeighborSolicitationType {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationTypeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationType) Increment() PatternFlowIcmpv6NeighborSolicitationTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NeighborSolicitationTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NeighborSolicitationTypeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NeighborSolicitationTypeCounter value in the PatternFlowIcmpv6NeighborSolicitationType object
func (obj *patternFlowIcmpv6NeighborSolicitationType) SetIncrement(value PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationType {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationTypeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationType) Decrement() PatternFlowIcmpv6NeighborSolicitationTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NeighborSolicitationTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NeighborSolicitationTypeCounter
func (obj *patternFlowIcmpv6NeighborSolicitationType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NeighborSolicitationTypeCounter value in the PatternFlowIcmpv6NeighborSolicitationType object
func (obj *patternFlowIcmpv6NeighborSolicitationType) SetDecrement(value PatternFlowIcmpv6NeighborSolicitationTypeCounter) PatternFlowIcmpv6NeighborSolicitationType {
	obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NeighborSolicitationType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NeighborSolicitationType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NeighborSolicitationType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIcmpv6NeighborSolicitationType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NeighborSolicitationTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NeighborSolicitationTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NeighborSolicitationTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NeighborSolicitationType")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NeighborSolicitationType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NeighborSolicitationType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
