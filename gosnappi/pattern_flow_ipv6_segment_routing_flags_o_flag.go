package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsOFlag *****
type patternFlowIpv6SegmentRoutingFlagsOFlag struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsOFlag
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlag
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsOFlag() PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj := patternFlowIpv6SegmentRoutingFlagsOFlag{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsOFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsOFlag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOFlag
}

type marshalPatternFlowIpv6SegmentRoutingFlagsOFlag interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsOFlag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOFlag, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsOFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsOFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsOFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag struct {
	obj *patternFlowIpv6SegmentRoutingFlagsOFlag
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlag interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag) (PatternFlowIpv6SegmentRoutingFlagsOFlag, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsOFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlag) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsOFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag) (PatternFlowIpv6SegmentRoutingFlagsOFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsOFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsOFlag) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Clone() (PatternFlowIpv6SegmentRoutingFlagsOFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsOFlag()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsOFlag is oAM (O) flag (RFC 9259 Section 2). When set, marks this packet as an active OAM probe. SRv6 endpoint nodes that process OAM packets inspect SRH TLVs such as the Ingress Node TLV (type 1) and Egress Node TLV (type 2) to perform path verification. Reference: RFC 8754 Section 2.1, RFC 9259 Section 2.
type PatternFlowIpv6SegmentRoutingFlagsOFlag interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsOFlag to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsOFlag from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsOFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsOFlag) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsOFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsOFlag
	// validate validates PatternFlowIpv6SegmentRoutingFlagsOFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsOFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsOFlag
	Choice() PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlag
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsOFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsOFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlag
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsOFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsOFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlag
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsOFlag.
	// PatternFlowIpv6SegmentRoutingFlagsOFlagCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsOFlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlag.
	// PatternFlowIpv6SegmentRoutingFlagsOFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsOFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsOFlagCounter, set in PatternFlowIpv6SegmentRoutingFlagsOFlag.
	// PatternFlowIpv6SegmentRoutingFlagsOFlagCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsOFlagCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsOFlag.
	// PatternFlowIpv6SegmentRoutingFlagsOFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsOFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsOFlag
var PatternFlowIpv6SegmentRoutingFlagsOFlagChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Choice() PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) setChoice(value PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsOFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsOFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsOFlagCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsOFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsOFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Increment() PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsOFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsOFlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsOFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) Decrement() PatternFlowIpv6SegmentRoutingFlagsOFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsOFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsOFlagCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsOFlagCounter value in the PatternFlowIpv6SegmentRoutingFlagsOFlag object
func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsOFlagCounter) PatternFlowIpv6SegmentRoutingFlagsOFlag {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsOFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsOFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsOFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsOFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsOFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsOFlag")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsOFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsOFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
