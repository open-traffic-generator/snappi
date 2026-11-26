package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryQqic *****
type patternFlowIgmpv3MembershipQueryQqic struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryQqic
	marshaller      marshalPatternFlowIgmpv3MembershipQueryQqic
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryQqic
	incrementHolder PatternFlowIgmpv3MembershipQueryQqicCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryQqicCounter
}

func NewPatternFlowIgmpv3MembershipQueryQqic() PatternFlowIgmpv3MembershipQueryQqic {
	obj := patternFlowIgmpv3MembershipQueryQqic{obj: &otg.PatternFlowIgmpv3MembershipQueryQqic{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) msg() *otg.PatternFlowIgmpv3MembershipQueryQqic {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryQqic) PatternFlowIgmpv3MembershipQueryQqic {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryQqic struct {
	obj *patternFlowIgmpv3MembershipQueryQqic
}

type marshalPatternFlowIgmpv3MembershipQueryQqic interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryQqic to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqic
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQqic, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryQqic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryQqic to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryQqic to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryQqic struct {
	obj *patternFlowIgmpv3MembershipQueryQqic
}

type unMarshalPatternFlowIgmpv3MembershipQueryQqic interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryQqic from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqic
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQqic) (PatternFlowIgmpv3MembershipQueryQqic, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryQqic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryQqic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryQqic from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) Marshal() marshalPatternFlowIgmpv3MembershipQueryQqic {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryQqic{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQqic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryQqic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQqic) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryQqic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqic) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryQqic) (PatternFlowIgmpv3MembershipQueryQqic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryQqic) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqic) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQqic) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqic) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryQqic) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryQqic) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryQqic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) Clone() (PatternFlowIgmpv3MembershipQueryQqic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryQqic()
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

func (obj *patternFlowIgmpv3MembershipQueryQqic) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryQqic is qQIC, Querier's Query Interval Code (RFC 9776 Section 4.1.7), in seconds. Below 128 it is the interval itself, so the default 125 is the RFC default Query Interval (Section 8.2). From 128 up it uses the floating point form of max_resp_code, so 255 is 31744 seconds. Zero tells receivers to use their own default. The value is sent as configured.
type PatternFlowIgmpv3MembershipQueryQqic interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryQqic to protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqic
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryQqic
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryQqic from protobuf object *otg.PatternFlowIgmpv3MembershipQueryQqic
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryQqic) PatternFlowIgmpv3MembershipQueryQqic
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryQqic
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryQqic
	// validate validates PatternFlowIgmpv3MembershipQueryQqic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryQqic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryQqicChoiceEnum, set in PatternFlowIgmpv3MembershipQueryQqic
	Choice() PatternFlowIgmpv3MembershipQueryQqicChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryQqicChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryQqic
	setChoice(value PatternFlowIgmpv3MembershipQueryQqicChoiceEnum) PatternFlowIgmpv3MembershipQueryQqic
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryQqic
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryQqic.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryQqic
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryQqic
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryQqic
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryQqic.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryQqic
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryQqic
	// Increment returns PatternFlowIgmpv3MembershipQueryQqicCounter, set in PatternFlowIgmpv3MembershipQueryQqic.
	// PatternFlowIgmpv3MembershipQueryQqicCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryQqicCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryQqicCounter provided by user to PatternFlowIgmpv3MembershipQueryQqic.
	// PatternFlowIgmpv3MembershipQueryQqicCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqic
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryQqic
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryQqicCounter, set in PatternFlowIgmpv3MembershipQueryQqic.
	// PatternFlowIgmpv3MembershipQueryQqicCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryQqicCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryQqicCounter provided by user to PatternFlowIgmpv3MembershipQueryQqic.
	// PatternFlowIgmpv3MembershipQueryQqicCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqic
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryQqic
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryQqicChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryQqic
var PatternFlowIgmpv3MembershipQueryQqicChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryQqicChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryQqicChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryQqicChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryQqicChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryQqicChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryQqicChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryQqicChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryQqicChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) Choice() PatternFlowIgmpv3MembershipQueryQqicChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryQqicChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryQqic) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) setChoice(value PatternFlowIgmpv3MembershipQueryQqicChoiceEnum) PatternFlowIgmpv3MembershipQueryQqic {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryQqic_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryQqicChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryQqic_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryQqicChoice.VALUE {
		defaultValue := uint32(125)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryQqicChoice.VALUES {
		defaultValue := []uint32{125}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryQqicChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryQqicCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryQqicChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryQqicCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqic) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryQqic) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryQqic object
func (obj *patternFlowIgmpv3MembershipQueryQqic) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryQqic {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryQqic) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{125})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryQqic object
func (obj *patternFlowIgmpv3MembershipQueryQqic) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryQqic {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryQqicCounter
func (obj *patternFlowIgmpv3MembershipQueryQqic) Increment() PatternFlowIgmpv3MembershipQueryQqicCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryQqicCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryQqicCounter
func (obj *patternFlowIgmpv3MembershipQueryQqic) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryQqicCounter value in the PatternFlowIgmpv3MembershipQueryQqic object
func (obj *patternFlowIgmpv3MembershipQueryQqic) SetIncrement(value PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqic {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryQqicCounter
func (obj *patternFlowIgmpv3MembershipQueryQqic) Decrement() PatternFlowIgmpv3MembershipQueryQqicCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryQqicCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryQqicCounter
func (obj *patternFlowIgmpv3MembershipQueryQqic) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryQqicCounter value in the PatternFlowIgmpv3MembershipQueryQqic object
func (obj *patternFlowIgmpv3MembershipQueryQqic) SetDecrement(value PatternFlowIgmpv3MembershipQueryQqicCounter) PatternFlowIgmpv3MembershipQueryQqic {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryQqic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryQqic.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryQqic.Values <= 255 but Got %d", item))
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

func (obj *patternFlowIgmpv3MembershipQueryQqic) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryQqicChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQqicChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQqicChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQqicChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryQqicChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryQqicChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryQqic")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryQqic_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryQqic_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
