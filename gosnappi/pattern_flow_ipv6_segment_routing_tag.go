package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTag *****
type patternFlowIpv6SegmentRoutingTag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTag
	marshaller      marshalPatternFlowIpv6SegmentRoutingTag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTag
	incrementHolder PatternFlowIpv6SegmentRoutingTagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTagCounter
}

func NewPatternFlowIpv6SegmentRoutingTag() PatternFlowIpv6SegmentRoutingTag {
	obj := patternFlowIpv6SegmentRoutingTag{obj: &otg.PatternFlowIpv6SegmentRoutingTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTag) msg() *otg.PatternFlowIpv6SegmentRoutingTag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTag) PatternFlowIpv6SegmentRoutingTag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTag struct {
	obj *patternFlowIpv6SegmentRoutingTag
}

type marshalPatternFlowIpv6SegmentRoutingTag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTag to protobuf object *otg.PatternFlowIpv6SegmentRoutingTag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTag struct {
	obj *patternFlowIpv6SegmentRoutingTag
}

type unMarshalPatternFlowIpv6SegmentRoutingTag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTag from protobuf object *otg.PatternFlowIpv6SegmentRoutingTag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTag) (PatternFlowIpv6SegmentRoutingTag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTag) Marshal() marshalPatternFlowIpv6SegmentRoutingTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTag) (PatternFlowIpv6SegmentRoutingTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTag) Clone() (PatternFlowIpv6SegmentRoutingTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTag()
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

func (obj *patternFlowIpv6SegmentRoutingTag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTag is 16-bit field used to tag a packet as part of a class or group, enabling shared properties or policy treatment. If not used, it MUST be set to zero.
type PatternFlowIpv6SegmentRoutingTag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTag to protobuf object *otg.PatternFlowIpv6SegmentRoutingTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTag from protobuf object *otg.PatternFlowIpv6SegmentRoutingTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTag) PatternFlowIpv6SegmentRoutingTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTag
	// validate validates PatternFlowIpv6SegmentRoutingTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTagChoiceEnum, set in PatternFlowIpv6SegmentRoutingTag
	Choice() PatternFlowIpv6SegmentRoutingTagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTag
	setChoice(value PatternFlowIpv6SegmentRoutingTagChoiceEnum) PatternFlowIpv6SegmentRoutingTag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTag
	// Increment returns PatternFlowIpv6SegmentRoutingTagCounter, set in PatternFlowIpv6SegmentRoutingTag.
	// PatternFlowIpv6SegmentRoutingTagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTagCounter provided by user to PatternFlowIpv6SegmentRoutingTag.
	// PatternFlowIpv6SegmentRoutingTagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTagCounter, set in PatternFlowIpv6SegmentRoutingTag.
	// PatternFlowIpv6SegmentRoutingTagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTagCounter provided by user to PatternFlowIpv6SegmentRoutingTag.
	// PatternFlowIpv6SegmentRoutingTagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTag
var PatternFlowIpv6SegmentRoutingTagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTag) Choice() PatternFlowIpv6SegmentRoutingTagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTag) setChoice(value PatternFlowIpv6SegmentRoutingTagChoiceEnum) PatternFlowIpv6SegmentRoutingTag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTag object
func (obj *patternFlowIpv6SegmentRoutingTag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTag object
func (obj *patternFlowIpv6SegmentRoutingTag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTagCounter
func (obj *patternFlowIpv6SegmentRoutingTag) Increment() PatternFlowIpv6SegmentRoutingTagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTagCounter
func (obj *patternFlowIpv6SegmentRoutingTag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTagCounter value in the PatternFlowIpv6SegmentRoutingTag object
func (obj *patternFlowIpv6SegmentRoutingTag) SetIncrement(value PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTagCounter
func (obj *patternFlowIpv6SegmentRoutingTag) Decrement() PatternFlowIpv6SegmentRoutingTagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTagCounter
func (obj *patternFlowIpv6SegmentRoutingTag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTagCounter value in the PatternFlowIpv6SegmentRoutingTag object
func (obj *patternFlowIpv6SegmentRoutingTag) SetDecrement(value PatternFlowIpv6SegmentRoutingTagCounter) PatternFlowIpv6SegmentRoutingTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTag.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTag.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
