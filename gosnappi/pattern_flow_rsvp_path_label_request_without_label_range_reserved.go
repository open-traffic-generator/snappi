package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved *****
type patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved struct {
	validation
	obj             *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	marshaller      marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	unMarshaller    unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	incrementHolder PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	decrementHolder PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
}

func NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj := patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved{obj: &otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	return obj.obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) setMsg(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
}

type marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved interface {
	// ToProto marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error)
	// ToPbText marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved struct {
	obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
}

type unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved interface {
	// FromProto unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error)
	// FromPbText unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) ToProto() (*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FromProto(msg *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved()
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved is this field is reserved.   It MUST be set to zero on transmission and MUST be ignored on receipt.
type PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved interface {
	Validation
	// msg marshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved to protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// setMsg unmarshals PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved from protobuf object *otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// validate validates PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	Choice() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum
	// setChoice assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	setChoice(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	SetValue(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// HasValue checks if Value has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	SetValues(value []uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// Increment returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter is integer counter pattern
	Increment() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// SetIncrement assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter, set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
	// SetDecrement assigns PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter provided by user to PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.
	// PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved
var PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice = struct {
	VALUE     PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum
	VALUES    PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum
	INCREMENT PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum
	DECREMENT PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Choice() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum {
	return PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) setChoice(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	intValue, ok := otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter().msg()
	}

	if value == PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) SetValue(value uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) SetValues(value []uint32) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Increment() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) SetIncrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) Decrement() PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter value in the PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved object
func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) SetDecrement(value PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedCounter) PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved {
	obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathLabelRequestWithoutLabelRangeReserved.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathLabelRequestWithoutLabelRangeReserved) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathLabelRequestWithoutLabelRangeReservedChoice.VALUE)

	}

}
