package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionLFlag *****
type patternFlowIcmpv6NdpPrefixInformationOptionLFlag struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlag() PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionLFlag{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlag()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionLFlag is on-link (L) flag. When set, the prefix can be used for on-link determination. RFC 4861 Section 4.6.2.
type PatternFlowIcmpv6NdpPrefixInformationOptionLFlag interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionLFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
var PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionLFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionLFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionLFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionLFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionLFlag")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionLFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
