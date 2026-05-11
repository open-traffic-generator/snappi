package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingUsidSegmentsLeft *****
type patternFlowIpv6SegmentRoutingUsidSegmentsLeft struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	marshaller      marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	incrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	decrementHolder PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
}

func NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeft() PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj := patternFlowIpv6SegmentRoutingUsidSegmentsLeft{obj: &otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft
}

type marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft struct {
	obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft
}

type unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) (PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) ToProto() (*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) (PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingUsidSegmentsLeft) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeft()
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

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingUsidSegmentsLeft is number of SRH segment list entries remaining to be visited after the current active container (RFC 8754 Section 2.1). The active container is Segment[segments_left]; the value is decremented each time the active container is exhausted and the pointer advances to the next entry. When auto is assigned the value is set to one less than the number of containers specified (last_entry value).
type PatternFlowIpv6SegmentRoutingUsidSegmentsLeft interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft to protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingUsidSegmentsLeft from protobuf object *otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// validate validates PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingUsidSegmentsLeft, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	Choice() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter, set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter provided by user to PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingUsidSegmentsLeft
var PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Choice() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum {
	return PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) setChoice(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) SetValue(value uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Increment() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) SetIncrement(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) Decrement() PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter value in the PatternFlowIpv6SegmentRoutingUsidSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) SetDecrement(value PatternFlowIpv6SegmentRoutingUsidSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingUsidSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingUsidSegmentsLeft.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingUsidSegmentsLeft) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingUsidSegmentsLeftChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingUsidSegmentsLeft")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingUsidSegmentsLeft_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
