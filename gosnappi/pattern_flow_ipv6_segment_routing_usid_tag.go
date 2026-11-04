package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidTag *****
type patternFlowIpv6SegmentRoutingUsidTag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidTag
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidTag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidTag
	incrementHolder PatternFlowIpv6SegmentRoutingUsidTagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidTagCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidTag() PatternFlowIpv6SegmentRoutingUsidTag {
	obj := patternFlowIpv6SegmentRoutingUsidTag{obj: &otg.PatternFlowIpv6SegmentRoutingUsidTag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) msg() *otg.PatternFlowIpv6SegmentRoutingUsidTag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidTag) PatternFlowIpv6SegmentRoutingUsidTag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidTag struct {
	obj *patternFlowIpv6SegmentRoutingUsidTag
}

type marshalPatternFlowIpv6SegmentRoutingUsidTag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidTag to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidTag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidTag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidTag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidTag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidTag struct {
	obj *patternFlowIpv6SegmentRoutingUsidTag
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidTag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidTag from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidTag) (PatternFlowIpv6SegmentRoutingUsidTag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidTag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidTag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidTag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidTag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidTag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidTag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidTag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidTag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidTag) (PatternFlowIpv6SegmentRoutingUsidTag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidTag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidTag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidTag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) Clone() (PatternFlowIpv6SegmentRoutingUsidTag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidTag()
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

func (obj *patternFlowIpv6SegmentRoutingUsidTag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidTag is 16-bit field used to tag a packet as part of a class or group (RFC 8754 Section 2.1). If not used it MUST be set to zero.
type PatternFlowIpv6SegmentRoutingUsidTag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidTag to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidTag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidTag from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidTag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidTag) PatternFlowIpv6SegmentRoutingUsidTag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidTag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidTag
	// validate validates PatternFlowIpv6SegmentRoutingUsidTag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidTag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidTag
	Choice() PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidTag
	setChoice(value PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum) PatternFlowIpv6SegmentRoutingUsidTag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidTag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingUsidTag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidTag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidTag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidTag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingUsidTag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidTag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidTag
	// Increment returns PatternFlowIpv6SegmentRoutingUsidTagCounter, set in PatternFlowIpv6SegmentRoutingUsidTag.
	// PatternFlowIpv6SegmentRoutingUsidTagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidTagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidTagCounter provided by user to PatternFlowIpv6SegmentRoutingUsidTag.
	// PatternFlowIpv6SegmentRoutingUsidTagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidTag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidTagCounter, set in PatternFlowIpv6SegmentRoutingUsidTag.
	// PatternFlowIpv6SegmentRoutingUsidTagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidTagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidTagCounter provided by user to PatternFlowIpv6SegmentRoutingUsidTag.
	// PatternFlowIpv6SegmentRoutingUsidTagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidTag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidTag
var PatternFlowIpv6SegmentRoutingUsidTagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) Choice() PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidTag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) setChoice(value PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum) PatternFlowIpv6SegmentRoutingUsidTag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidTag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidTag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidTagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidTagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidTagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidTagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidTag object
func (obj *patternFlowIpv6SegmentRoutingUsidTag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingUsidTag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingUsidTag object
func (obj *patternFlowIpv6SegmentRoutingUsidTag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidTagCounter
func (obj *patternFlowIpv6SegmentRoutingUsidTag) Increment() PatternFlowIpv6SegmentRoutingUsidTagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidTagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidTagCounter
func (obj *patternFlowIpv6SegmentRoutingUsidTag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidTagCounter value in the PatternFlowIpv6SegmentRoutingUsidTag object
func (obj *patternFlowIpv6SegmentRoutingUsidTag) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidTagCounter
func (obj *patternFlowIpv6SegmentRoutingUsidTag) Decrement() PatternFlowIpv6SegmentRoutingUsidTagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidTagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidTagCounter
func (obj *patternFlowIpv6SegmentRoutingUsidTag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidTagCounter value in the PatternFlowIpv6SegmentRoutingUsidTag object
func (obj *patternFlowIpv6SegmentRoutingUsidTag) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidTagCounter) PatternFlowIpv6SegmentRoutingUsidTag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidTag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidTag.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingUsidTag.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingUsidTag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidTagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidTagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidTagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidTagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidTag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidTag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidTag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
