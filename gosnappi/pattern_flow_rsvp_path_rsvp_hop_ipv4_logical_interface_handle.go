package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle *****
type patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle struct {
	validation
	obj             *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	marshaller      marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	unMarshaller    unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	incrementHolder PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	decrementHolder PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
}

func NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj := patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle{obj: &otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) msg() *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	return obj.obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) setMsg(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle struct {
	obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
}

type marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle interface {
	// ToProto marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error)
	// ToPbText marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle struct {
	obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
}

type unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle interface {
	// FromProto unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error)
	// FromPbText unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) ToProto() (*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FromProto(msg *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Clone() (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle()
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

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle is logical Interface Handle (LIH) is used to distinguish logical outgoing interfaces. A node receiving an LIH in a Path message saves its value and returns it in the HOP objects of subsequent Resv messages sent to the node that originated the LIH. The LIH should be identically zero if there is no logical interface handle.
type PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle interface {
	Validation
	// msg marshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle to protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// setMsg unmarshals PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle from protobuf object *otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// validate validates PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	Choice() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum
	// setChoice assigns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	setChoice(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	SetValue(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// HasValue checks if Value has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	SetValues(value []uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// Increment returns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter is integer counter pattern
	Increment() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// SetIncrement assigns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter, set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
	// SetDecrement assigns PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter provided by user to PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle.
	// PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle
var PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice = struct {
	VALUE     PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum
	VALUES    PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum
	INCREMENT PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum
	DECREMENT PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Choice() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum {
	return PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) setChoice(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	intValue, ok := otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter().msg()
	}

	if value == PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) SetValue(value uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) SetValues(value []uint32) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Increment() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) SetIncrement(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) Decrement() PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter value in the PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle object
func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) SetDecrement(value PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleCounter) PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle {
	obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandleChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathRsvpHopIpv4LogicalInterfaceHandle_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
