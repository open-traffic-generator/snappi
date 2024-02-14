package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServReserved1 *****
type patternFlowRSVPPathSenderTspecIntServReserved1 struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServReserved1
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1
	incrementHolder PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServReserved1Counter
}

func NewPatternFlowRSVPPathSenderTspecIntServReserved1() PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj := patternFlowRSVPPathSenderTspecIntServReserved1{obj: &otg.PatternFlowRSVPPathSenderTspecIntServReserved1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved1 {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServReserved1 struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved1
}

type marshalPatternFlowRSVPPathSenderTspecIntServReserved1 interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServReserved1 to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved1, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServReserved1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServReserved1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServReserved1 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1 struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved1
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1 interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1 from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1) (PatternFlowRSVPPathSenderTspecIntServReserved1, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServReserved1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved1) (PatternFlowRSVPPathSenderTspecIntServReserved1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved1) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Clone() (PatternFlowRSVPPathSenderTspecIntServReserved1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServReserved1()
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServReserved1 is reserved.
type PatternFlowRSVPPathSenderTspecIntServReserved1 interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServReserved1 to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServReserved1 from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServReserved1) PatternFlowRSVPPathSenderTspecIntServReserved1
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved1
	// validate validates PatternFlowRSVPPathSenderTspecIntServReserved1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServReserved1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServReserved1
	Choice() PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1
	setChoice(value PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum) PatternFlowRSVPPathSenderTspecIntServReserved1
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServReserved1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServReserved1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServReserved1
	// Increment returns PatternFlowRSVPPathSenderTspecIntServReserved1Counter, set in PatternFlowRSVPPathSenderTspecIntServReserved1.
	// PatternFlowRSVPPathSenderTspecIntServReserved1Counter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServReserved1Counter provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1.
	// PatternFlowRSVPPathSenderTspecIntServReserved1Counter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServReserved1
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServReserved1Counter, set in PatternFlowRSVPPathSenderTspecIntServReserved1.
	// PatternFlowRSVPPathSenderTspecIntServReserved1Counter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServReserved1Counter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServReserved1Counter provided by user to PatternFlowRSVPPathSenderTspecIntServReserved1.
	// PatternFlowRSVPPathSenderTspecIntServReserved1Counter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServReserved1
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServReserved1
var PatternFlowRSVPPathSenderTspecIntServReserved1Choice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Choice() PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) setChoice(value PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServReserved1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServReserved1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServReserved1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServReserved1Counter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServReserved1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved1 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved1 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServReserved1Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Increment() PatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServReserved1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServReserved1Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServReserved1Counter value in the PatternFlowRSVPPathSenderTspecIntServReserved1 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServReserved1Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) Decrement() PatternFlowRSVPPathSenderTspecIntServReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServReserved1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServReserved1Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServReserved1Counter value in the PatternFlowRSVPPathSenderTspecIntServReserved1 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServReserved1Counter) PatternFlowRSVPPathSenderTspecIntServReserved1 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved1.Value <= 4095 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 4095 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServReserved1.Values <= 4095 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved1) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved1Choice.VALUE)

	}

}
