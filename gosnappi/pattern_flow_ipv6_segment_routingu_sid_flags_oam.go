package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutinguSidFlagsOam *****
type patternFlowIpv6SegmentRoutinguSidFlagsOam struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	marshaller      marshalPatternFlowIpv6SegmentRoutinguSidFlagsOam
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOam
	incrementHolder PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	decrementHolder PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
}

func NewPatternFlowIpv6SegmentRoutinguSidFlagsOam() PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj := patternFlowIpv6SegmentRoutinguSidFlagsOam{obj: &otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) msg() *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) setMsg(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam struct {
	obj *patternFlowIpv6SegmentRoutinguSidFlagsOam
}

type marshalPatternFlowIpv6SegmentRoutinguSidFlagsOam interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutinguSidFlagsOam to protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	ToProto() (*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutinguSidFlagsOam to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutinguSidFlagsOam to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutinguSidFlagsOam to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam struct {
	obj *patternFlowIpv6SegmentRoutinguSidFlagsOam
}

type unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOam interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOam from protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam) (PatternFlowIpv6SegmentRoutinguSidFlagsOam, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOam from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOam from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOam from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Marshal() marshalPatternFlowIpv6SegmentRoutinguSidFlagsOam {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOam {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) ToProto() (*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) FromProto(msg *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam) (PatternFlowIpv6SegmentRoutinguSidFlagsOam, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutinguSidFlagsOam) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Clone() (PatternFlowIpv6SegmentRoutinguSidFlagsOam, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutinguSidFlagsOam()
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

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutinguSidFlagsOam is (RFC 9259, section 2) OAM flag or O-flag is set at bit-2. Indicates if the packet is an Operations, Administration, and Maintenance (OAM) packet.
type PatternFlowIpv6SegmentRoutinguSidFlagsOam interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutinguSidFlagsOam to protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// setMsg unmarshals PatternFlowIpv6SegmentRoutinguSidFlagsOam from protobuf object *otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutinguSidFlagsOam
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutinguSidFlagsOam
	// validate validates PatternFlowIpv6SegmentRoutinguSidFlagsOam
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutinguSidFlagsOam, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum, set in PatternFlowIpv6SegmentRoutinguSidFlagsOam
	Choice() PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOam
	setChoice(value PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOam
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOam
	SetValue(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOam
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOam
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// Increment returns PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, set in PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOam
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter, set in PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter provided by user to PatternFlowIpv6SegmentRoutinguSidFlagsOam.
	// PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOam
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutinguSidFlagsOam
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutinguSidFlagsOam
var PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Choice() PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum {
	return PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) setChoice(value PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutinguSidFlagsOamCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutinguSidFlagsOam object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) SetValue(value uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutinguSidFlagsOam object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) SetValues(value []uint32) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Increment() PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutinguSidFlagsOamCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter value in the PatternFlowIpv6SegmentRoutinguSidFlagsOam object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) SetIncrement(value PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) Decrement() PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutinguSidFlagsOamCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter value in the PatternFlowIpv6SegmentRoutinguSidFlagsOam object
func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) SetDecrement(value PatternFlowIpv6SegmentRoutinguSidFlagsOamCounter) PatternFlowIpv6SegmentRoutinguSidFlagsOam {
	obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutinguSidFlagsOam.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutinguSidFlagsOam.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutinguSidFlagsOam) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutinguSidFlagsOamChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutinguSidFlagsOamChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutinguSidFlagsOam")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutinguSidFlagsOam_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
