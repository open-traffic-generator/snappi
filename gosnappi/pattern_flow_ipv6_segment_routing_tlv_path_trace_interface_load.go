package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad *****
type patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	marshaller      marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	incrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	decrementHolder PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
}

func NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj := patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad{obj: &otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
}

type marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad struct {
	obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
}

type unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) ToProto() (*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad()
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad is load on the interface at the time of measurement.
type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad to protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad from protobuf object *otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// validate validates PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// Increment returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter, set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter provided by user to PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.
	// PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad
var PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Choice() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum {
	return PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) setChoice(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) SetValue(value uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Increment() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) SetIncrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) Decrement() PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter value in the PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad object
func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) SetDecrement(value PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadCounter) PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad {
	obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad.Values <= 15 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoadChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingTlvPathTraceInterfaceLoad_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
