package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsHFlag *****
type patternFlowIpv6SegmentRoutingFlagsHFlag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsHFlag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlag
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsHFlag() PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj := patternFlowIpv6SegmentRoutingFlagsHFlag{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsHFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsHFlag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHFlag
}

type marshalPatternFlowIpv6SegmentRoutingFlagsHFlag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsHFlag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHFlag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsHFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsHFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsHFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHFlag
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag) (PatternFlowIpv6SegmentRoutingFlagsHFlag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsHFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag) (PatternFlowIpv6SegmentRoutingFlagsHFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHFlag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Clone() (PatternFlowIpv6SegmentRoutingFlagsHFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsHFlag()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsHFlag is hMAC (H) flag (RFC 8754 Section 2.1). When set, indicates that this SRH includes an HMAC TLV (type 5) providing data integrity and source authentication for the segment routing path. Reference: RFC 8754 Section 2.1.
type PatternFlowIpv6SegmentRoutingFlagsHFlag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsHFlag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsHFlag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsHFlag) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHFlag
	// validate validates PatternFlowIpv6SegmentRoutingFlagsHFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsHFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsHFlag
	Choice() PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlag
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsHFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsHFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsHFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsHFlag.
	// PatternFlowIpv6SegmentRoutingFlagsHFlagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsHFlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlag.
	// PatternFlowIpv6SegmentRoutingFlagsHFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsHFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsHFlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsHFlag.
	// PatternFlowIpv6SegmentRoutingFlagsHFlagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsHFlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsHFlag.
	// PatternFlowIpv6SegmentRoutingFlagsHFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsHFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsHFlag
var PatternFlowIpv6SegmentRoutingFlagsHFlagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Choice() PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) setChoice(value PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsHFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsHFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsHFlagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsHFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Increment() PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsHFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsHFlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsHFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) Decrement() PatternFlowIpv6SegmentRoutingFlagsHFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsHFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsHFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsHFlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsHFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsHFlagCounter) PatternFlowIpv6SegmentRoutingFlagsHFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsHFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsHFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsHFlag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsHFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsHFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
