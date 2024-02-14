package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathExplicitRouteType1ASNumberLBit *****
type patternFlowRSVPPathExplicitRouteType1ASNumberLBit struct {
	validation
	obj             *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	marshaller      marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	unMarshaller    unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	incrementHolder PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	decrementHolder PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
}

func NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBit() PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj := patternFlowRSVPPathExplicitRouteType1ASNumberLBit{obj: &otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) msg() *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	return obj.obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) setMsg(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit struct {
	obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit
}

type marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit interface {
	// ToProto marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error)
	// ToPbText marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit struct {
	obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit
}

type unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit interface {
	// FromProto unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) (PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error)
	// FromPbText unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Marshal() marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) ToProto() (*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FromProto(msg *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) (PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathExplicitRouteType1ASNumberLBit) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Clone() (PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBit()
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

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathExplicitRouteType1ASNumberLBit is the L bit is an attribute of the subobject. The L bit is set if the subobject represents a loose hop in the explicit route. If the bit is not set, the subobject represents a strict hop in the explicit route.
type PatternFlowRSVPPathExplicitRouteType1ASNumberLBit interface {
	Validation
	// msg marshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit to protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// setMsg unmarshals PatternFlowRSVPPathExplicitRouteType1ASNumberLBit from protobuf object *otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// validate validates PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathExplicitRouteType1ASNumberLBit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	Choice() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum
	// setChoice assigns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	setChoice(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	SetValue(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// HasValue checks if Value has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	SetValues(value []uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// Increment returns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter is integer counter pattern
	Increment() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// SetIncrement assigns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter, set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
	// SetDecrement assigns PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter provided by user to PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.
	// PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathExplicitRouteType1ASNumberLBit
var PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice = struct {
	VALUE     PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum
	VALUES    PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum
	INCREMENT PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum
	DECREMENT PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Choice() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum {
	return PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) setChoice(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	intValue, ok := otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathExplicitRouteType1ASNumberLBit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter().msg()
	}

	if value == PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) SetValue(value uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) SetValues(value []uint32) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Increment() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) SetIncrement(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) Decrement() PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter value in the PatternFlowRSVPPathExplicitRouteType1ASNumberLBit object
func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) SetDecrement(value PatternFlowRSVPPathExplicitRouteType1ASNumberLBitCounter) PatternFlowRSVPPathExplicitRouteType1ASNumberLBit {
	obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathExplicitRouteType1ASNumberLBit.Values <= 1 but Got %d", item))
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

func (obj *patternFlowRSVPPathExplicitRouteType1ASNumberLBit) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathExplicitRouteType1ASNumberLBitChoice.VALUE)

	}

}
