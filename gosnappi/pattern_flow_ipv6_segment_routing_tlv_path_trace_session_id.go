package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId *****
type patternFlowIpv6SegmentRoutingTlvPathTraceSessionId struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	incrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceSessionId{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId()
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId is identifier for the monitoring session.
type PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// Increment returns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId
var PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceSessionId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceSessionIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceSessionId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
