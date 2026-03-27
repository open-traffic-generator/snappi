package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsProtected *****
type patternFlowIpv6SegmentRoutingFlagsProtected struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsProtected
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsProtected
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsProtected() PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj := patternFlowIpv6SegmentRoutingFlagsProtected{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsProtected{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsProtected {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtected) PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsProtected struct {
	obj *patternFlowIpv6SegmentRoutingFlagsProtected
}

type marshalPatternFlowIpv6SegmentRoutingFlagsProtected interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsProtected to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsProtected, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsProtected to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsProtected to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsProtected to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected struct {
	obj *patternFlowIpv6SegmentRoutingFlagsProtected
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsProtected interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsProtected from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtected) (PatternFlowIpv6SegmentRoutingFlagsProtected, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsProtected from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsProtected from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsProtected from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsProtected {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsProtected{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsProtected {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtected) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsProtected, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsProtected) (PatternFlowIpv6SegmentRoutingFlagsProtected, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtected) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtected) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsProtected) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsProtected) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Clone() (PatternFlowIpv6SegmentRoutingFlagsProtected, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsProtected()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsProtected is protected Flag. Indicates if the segment is protected by a Fast Re-Route (FRR) mechanism.
type PatternFlowIpv6SegmentRoutingFlagsProtected interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsProtected to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsProtected from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsProtected
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsProtected) PatternFlowIpv6SegmentRoutingFlagsProtected
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsProtected
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsProtected
	// validate validates PatternFlowIpv6SegmentRoutingFlagsProtected
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsProtected, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsProtected
	Choice() PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsProtected
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsProtected
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsProtected
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsProtected.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsProtected
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtected
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsProtected
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsProtected.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsProtected
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsProtected
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, set in PatternFlowIpv6SegmentRoutingFlagsProtected.
	// PatternFlowIpv6SegmentRoutingFlagsProtectedCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsProtectedCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsProtected.
	// PatternFlowIpv6SegmentRoutingFlagsProtectedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtected
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsProtected
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsProtectedCounter, set in PatternFlowIpv6SegmentRoutingFlagsProtected.
	// PatternFlowIpv6SegmentRoutingFlagsProtectedCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsProtectedCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsProtected.
	// PatternFlowIpv6SegmentRoutingFlagsProtectedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtected
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsProtected
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsProtected
var PatternFlowIpv6SegmentRoutingFlagsProtectedChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Choice() PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) setChoice(value PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsProtected {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsProtected_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsProtected_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsProtectedCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsProtectedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsProtected object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsProtected object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Increment() PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsProtectedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsProtectedCounter value in the PatternFlowIpv6SegmentRoutingFlagsProtected object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) Decrement() PatternFlowIpv6SegmentRoutingFlagsProtectedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsProtectedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsProtectedCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsProtectedCounter value in the PatternFlowIpv6SegmentRoutingFlagsProtected object
func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsProtectedCounter) PatternFlowIpv6SegmentRoutingFlagsProtected {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsProtected.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsProtected.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsProtected) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsProtectedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsProtectedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsProtected")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsProtected_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsProtected_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
