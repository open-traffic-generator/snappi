package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsOam *****
type patternFlowIpv6SegmentRoutingFlagsOam struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsOam
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsOam
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsOamCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsOamCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsOam() PatternFlowIpv6SegmentRoutingFlagsOam {
	obj := patternFlowIpv6SegmentRoutingFlagsOam{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsOam{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOam {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOam) PatternFlowIpv6SegmentRoutingFlagsOam {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsOam struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOam
}

type marshalPatternFlowIpv6SegmentRoutingFlagsOam interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsOam to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOam, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsOam to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsOam to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsOam to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsOam struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOam
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsOam interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsOam from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOam) (PatternFlowIpv6SegmentRoutingFlagsOam, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsOam from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsOam from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsOam from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOam {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsOam{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOam {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsOam{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOam) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOam, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOam) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOam) (PatternFlowIpv6SegmentRoutingFlagsOam, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOam) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOam) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOam) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOam) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOam) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOam) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Clone() (PatternFlowIpv6SegmentRoutingFlagsOam, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsOam()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsOam is oAM Flag. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
type PatternFlowIpv6SegmentRoutingFlagsOam interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsOam to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsOam from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOam
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsOam) PatternFlowIpv6SegmentRoutingFlagsOam
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOam
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOam
	// validate validates PatternFlowIpv6SegmentRoutingFlagsOam
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsOam, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsOam
	Choice() PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsOam
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsOam
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsOam
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOam.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOam
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsOam
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsOam
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsOam.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOam
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsOam
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsOamCounter, set in PatternFlowIpv6SegmentRoutingFlagsOam.
	// PatternFlowIpv6SegmentRoutingFlagsOamCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsOamCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsOam.
	// PatternFlowIpv6SegmentRoutingFlagsOamCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOam
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsOam
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsOamCounter, set in PatternFlowIpv6SegmentRoutingFlagsOam.
	// PatternFlowIpv6SegmentRoutingFlagsOamCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsOamCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsOamCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsOam.
	// PatternFlowIpv6SegmentRoutingFlagsOamCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOam
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsOam
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsOam
var PatternFlowIpv6SegmentRoutingFlagsOamChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Choice() PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) setChoice(value PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsOam {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsOam_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsOam_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOamChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsOamCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOamChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsOamCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOam object
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOam object
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Increment() PatternFlowIpv6SegmentRoutingFlagsOamCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsOamCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsOamCounter value in the PatternFlowIpv6SegmentRoutingFlagsOam object
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) Decrement() PatternFlowIpv6SegmentRoutingFlagsOamCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsOamCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsOamCounter value in the PatternFlowIpv6SegmentRoutingFlagsOam object
func (obj *patternFlowIpv6SegmentRoutingFlagsOam) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsOamCounter) PatternFlowIpv6SegmentRoutingFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOam.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsOam.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOam) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsOamChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOamChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOamChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOamChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsOam")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsOam_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsOam_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
