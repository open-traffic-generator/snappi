package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServReserved2 *****
type patternFlowRSVPPathSenderTspecIntServReserved2 struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServReserved2
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2
	incrementHolder PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServReserved2Counter
}

func NewPatternFlowRSVPPathSenderTspecIntServReserved2() PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj := patternFlowRSVPPathSenderTspecIntServReserved2{obj: &otg.PatternFlowRSVPPathSenderTspecIntServReserved2{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved2 {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServReserved2 struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved2
}

type marshalPatternFlowRSVPPathSenderTspecIntServReserved2 interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServReserved2 to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved2, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServReserved2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServReserved2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServReserved2 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2 struct {
	obj *patternFlowRSVPPathSenderTspecIntServReserved2
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2 interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2 from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2) (PatternFlowRSVPPathSenderTspecIntServReserved2, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServReserved2{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServReserved2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServReserved2) (PatternFlowRSVPPathSenderTspecIntServReserved2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServReserved2) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServReserved2) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Clone() (PatternFlowRSVPPathSenderTspecIntServReserved2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServReserved2()
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServReserved2 is reserved.
type PatternFlowRSVPPathSenderTspecIntServReserved2 interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServReserved2 to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServReserved2 from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServReserved2
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServReserved2) PatternFlowRSVPPathSenderTspecIntServReserved2
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServReserved2
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServReserved2
	// validate validates PatternFlowRSVPPathSenderTspecIntServReserved2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServReserved2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServReserved2
	Choice() PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2
	setChoice(value PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum) PatternFlowRSVPPathSenderTspecIntServReserved2
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServReserved2
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved2.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServReserved2
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServReserved2.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServReserved2
	// Increment returns PatternFlowRSVPPathSenderTspecIntServReserved2Counter, set in PatternFlowRSVPPathSenderTspecIntServReserved2.
	// PatternFlowRSVPPathSenderTspecIntServReserved2Counter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServReserved2Counter provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2.
	// PatternFlowRSVPPathSenderTspecIntServReserved2Counter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServReserved2
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServReserved2Counter, set in PatternFlowRSVPPathSenderTspecIntServReserved2.
	// PatternFlowRSVPPathSenderTspecIntServReserved2Counter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServReserved2Counter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServReserved2Counter provided by user to PatternFlowRSVPPathSenderTspecIntServReserved2.
	// PatternFlowRSVPPathSenderTspecIntServReserved2Counter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServReserved2
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServReserved2
var PatternFlowRSVPPathSenderTspecIntServReserved2Choice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Choice() PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) setChoice(value PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServReserved2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServReserved2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServReserved2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved2Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServReserved2Counter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServReserved2Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServReserved2Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved2 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServReserved2 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServReserved2Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Increment() PatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServReserved2Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServReserved2Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServReserved2Counter value in the PatternFlowRSVPPathSenderTspecIntServReserved2 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServReserved2Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) Decrement() PatternFlowRSVPPathSenderTspecIntServReserved2Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServReserved2Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServReserved2Counter
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServReserved2Counter value in the PatternFlowRSVPPathSenderTspecIntServReserved2 object
func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServReserved2Counter) PatternFlowRSVPPathSenderTspecIntServReserved2 {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServReserved2.Value <= 127 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 127 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServReserved2.Values <= 127 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServReserved2) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServReserved2Choice.VALUE)

	}

}
