package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId *****
type patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	incrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId()
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId is identifier of the interface where the measurement was taken.
type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// Increment returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId
var PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 4095 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId.Values <= 4095 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
