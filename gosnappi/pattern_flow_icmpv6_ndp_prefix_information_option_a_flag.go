package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionAFlag *****
type patternFlowIcmpv6NdpPrefixInformationOptionAFlag struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlag() PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionAFlag{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) (PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlag()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionAFlag is autonomous address-configuration (A) flag. When set, the prefix can be used for stateless address autoconfiguration. RFC 4861 Section 4.6.2.
type PatternFlowIcmpv6NdpPrefixInformationOptionAFlag interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionAFlag from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter, set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
var PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter value in the PatternFlowIcmpv6NdpPrefixInformationOptionAFlag object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlagCounter) PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionAFlag.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionAFlag) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionAFlagChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionAFlag")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionAFlag_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
