package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionRFlag *****
type patternFlowIcmpv6NdpPrefixInformationOptionRFlag struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlag() PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionRFlag{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlag()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionRFlag is router address (R) flag. When set, the Prefix field contains a complete IP address assigned to the sending router. Used by Mobile IPv6 (RFC 6275 Section 7.2).
type PatternFlowIcmpv6NdpPrefixInformationOptionRFlag interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionRFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
var PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionRFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionRFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionRFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionRFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionRFlag")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionRFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
