package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingSegmentsLeft *****
type patternFlowIpv6SegmentRoutingSegmentsLeft struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	marshaller      marshalPatternFlowIpv6SegmentRoutingSegmentsLeft
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeft
	incrementHolder PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	decrementHolder PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
}

func NewPatternFlowIpv6SegmentRoutingSegmentsLeft() PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj := patternFlowIpv6SegmentRoutingSegmentsLeft{obj: &otg.PatternFlowIpv6SegmentRoutingSegmentsLeft{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) msg() *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingSegmentsLeft struct {
	obj *patternFlowIpv6SegmentRoutingSegmentsLeft
}

type marshalPatternFlowIpv6SegmentRoutingSegmentsLeft interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingSegmentsLeft to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentsLeft, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingSegmentsLeft to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingSegmentsLeft to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingSegmentsLeft to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft struct {
	obj *patternFlowIpv6SegmentRoutingSegmentsLeft
}

type unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeft interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeft from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft) (PatternFlowIpv6SegmentRoutingSegmentsLeft, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeft from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeft from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeft from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentsLeft {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingSegmentsLeft{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeft {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeft) ToProto() (*otg.PatternFlowIpv6SegmentRoutingSegmentsLeft, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft) (PatternFlowIpv6SegmentRoutingSegmentsLeft, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeft) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeft) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingSegmentsLeft) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingSegmentsLeft) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Clone() (PatternFlowIpv6SegmentRoutingSegmentsLeft, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingSegmentsLeft()
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

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingSegmentsLeft is 8-bit unsigned integer containing the number of remaining segments to be visited before the packet reaches its final destination. It is decremented at each segment endpoint. It points to the current active segment in the Segment List. This should not be more than the number of segments specified in the segment list. When auto is assigned the value is set to the number of segments specified.
type PatternFlowIpv6SegmentRoutingSegmentsLeft interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingSegmentsLeft to protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingSegmentsLeft from protobuf object *otg.PatternFlowIpv6SegmentRoutingSegmentsLeft
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingSegmentsLeft) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingSegmentsLeft
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingSegmentsLeft
	// validate validates PatternFlowIpv6SegmentRoutingSegmentsLeft
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingSegmentsLeft, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum, set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	Choice() PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeft
	setChoice(value PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeft.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeft
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeft.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeft
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingSegmentsLeft.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, set in PatternFlowIpv6SegmentRoutingSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingSegmentsLeftCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingSegmentsLeftCounter provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingSegmentsLeftCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingSegmentsLeftCounter, set in PatternFlowIpv6SegmentRoutingSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingSegmentsLeftCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingSegmentsLeftCounter provided by user to PatternFlowIpv6SegmentRoutingSegmentsLeft.
	// PatternFlowIpv6SegmentRoutingSegmentsLeftCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeft
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingSegmentsLeft
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingSegmentsLeft
var PatternFlowIpv6SegmentRoutingSegmentsLeftChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Choice() PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum {
	return PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) setChoice(value PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingSegmentsLeft_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingSegmentsLeft_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingSegmentsLeftCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingSegmentsLeftCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) SetValue(value uint32) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUES)
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
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Increment() PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingSegmentsLeftCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingSegmentsLeftCounter value in the PatternFlowIpv6SegmentRoutingSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) SetIncrement(value PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) Decrement() PatternFlowIpv6SegmentRoutingSegmentsLeftCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingSegmentsLeftCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingSegmentsLeftCounter
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingSegmentsLeftCounter value in the PatternFlowIpv6SegmentRoutingSegmentsLeft object
func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) SetDecrement(value PatternFlowIpv6SegmentRoutingSegmentsLeftCounter) PatternFlowIpv6SegmentRoutingSegmentsLeft {
	obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingSegmentsLeft.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingSegmentsLeft.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingSegmentsLeft.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingSegmentsLeft) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingSegmentsLeftChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingSegmentsLeftChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingSegmentsLeft")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingSegmentsLeft_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingSegmentsLeft_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
