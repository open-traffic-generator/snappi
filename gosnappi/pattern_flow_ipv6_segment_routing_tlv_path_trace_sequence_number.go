package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber *****
type patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	incrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber()
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber is sequence number for the packet within the monitoring session.
type PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// Increment returns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber
var PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSequenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
