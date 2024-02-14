package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServParameter127Length *****
type patternFlowRSVPPathSenderTspecIntServParameter127Length struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServParameter127Length
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Length
	incrementHolder PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServParameter127Length() PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj := patternFlowRSVPPathSenderTspecIntServParameter127Length{obj: &otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127Length
}

type marshalPatternFlowRSVPPathSenderTspecIntServParameter127Length interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServParameter127Length to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServParameter127Length to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServParameter127Length to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServParameter127Length to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length struct {
	obj *patternFlowRSVPPathSenderTspecIntServParameter127Length
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Length interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Length from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length) (PatternFlowRSVPPathSenderTspecIntServParameter127Length, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Length from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Length from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Length from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127Length {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Length {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length) (PatternFlowRSVPPathSenderTspecIntServParameter127Length, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServParameter127Length) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127Length, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServParameter127Length()
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServParameter127Length is parameter 127 length, 5 words not including per-service header
type PatternFlowRSVPPathSenderTspecIntServParameter127Length interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServParameter127Length to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServParameter127Length from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServParameter127Length
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServParameter127Length
	// validate validates PatternFlowRSVPPathSenderTspecIntServParameter127Length
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServParameter127Length, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServParameter127Length
	Choice() PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Length
	setChoice(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Length
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Length
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Length
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Length
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// Increment returns PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, set in PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	// PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	// PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Length
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter, set in PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	// PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter provided by user to PatternFlowRSVPPathSenderTspecIntServParameter127Length.
	// PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Length
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServParameter127Length
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServParameter127Length
var PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Choice() PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) setChoice(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServParameter127Length_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUE {
		defaultValue := uint32(5)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUES {
		defaultValue := []uint32{5}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127Length object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{5})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServParameter127Length object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Increment() PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter value in the PatternFlowRSVPPathSenderTspecIntServParameter127Length object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) Decrement() PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServParameter127LengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter value in the PatternFlowRSVPPathSenderTspecIntServParameter127Length object
func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServParameter127LengthCounter) PatternFlowRSVPPathSenderTspecIntServParameter127Length {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServParameter127Length.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServParameter127Length.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServParameter127Length) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServParameter127LengthChoice.VALUE)

	}

}
