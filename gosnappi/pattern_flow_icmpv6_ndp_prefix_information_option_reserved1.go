package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 *****
type patternFlowIcmpv6NdpPrefixInformationOptionReserved1 struct {
	validation
	obj             *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	marshaller      marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	unMarshaller    unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	incrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
	decrementHolder PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
}

func NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved1() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj := patternFlowIcmpv6NdpPrefixInformationOptionReserved1{obj: &otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	return obj.obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) setMsg(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1 struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1
}

type marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1 interface {
	// ToProto marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error)
	// ToPbText marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1 struct {
	obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1
}

type unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1 interface {
	// FromProto unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) (PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error)
	// FromPbText unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) ToProto() (*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FromProto(msg *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) (PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved1()
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 is reserved. Must be initialized to zero by the sender.
type PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 interface {
	Validation
	// msg marshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 to protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// setMsg unmarshals PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 from protobuf object *otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// validate validates PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	Choice() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum
	// setChoice assigns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// HasValue checks if Value has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// Increment returns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter is integer counter pattern
	Increment() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
	// SetIncrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter, set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter is integer counter pattern
	Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
	// SetDecrement assigns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter provided by user to PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	HasDecrement() bool
	setNil()
}

type PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
var PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice = struct {
	VALUE     PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum
	VALUES    PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum
	INCREMENT PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum
	DECREMENT PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Choice() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum {
	return PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) setChoice(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	intValue, ok := otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter().msg()
	}

	if value == PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) SetValue(value uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) SetValues(value []uint32) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Increment() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) SetIncrement(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) Decrement() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter value in the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 object
func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) SetDecrement(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Counter) PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.Value <= 31 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 31 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6NdpPrefixInformationOptionReserved1.Values <= 31 but Got %d", item))
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

func (obj *patternFlowIcmpv6NdpPrefixInformationOptionReserved1) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6NdpPrefixInformationOptionReserved1ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6NdpPrefixInformationOptionReserved1Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6NdpPrefixInformationOptionReserved1")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6NdpPrefixInformationOptionReserved1_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
