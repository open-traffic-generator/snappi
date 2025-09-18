package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingFlagsHmac *****
type patternFlowIpv6SegmentRoutingFlagsHmac struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	marshaller      marshalPatternFlowIpv6SegmentRoutingFlagsHmac
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingFlagsHmac
	incrementHolder PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	decrementHolder PatternFlowIpv6SegmentRoutingFlagsHmacCounter
}

func NewPatternFlowIpv6SegmentRoutingFlagsHmac() PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj := patternFlowIpv6SegmentRoutingFlagsHmac{obj: &otg.PatternFlowIpv6SegmentRoutingFlagsHmac{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHmac {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmac) PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingFlagsHmac struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHmac
}

type marshalPatternFlowIpv6SegmentRoutingFlagsHmac interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingFlagsHmac to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHmac, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingFlagsHmac to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingFlagsHmac to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingFlagsHmac to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac struct {
	obj *patternFlowIpv6SegmentRoutingFlagsHmac
}

type unMarshalPatternFlowIpv6SegmentRoutingFlagsHmac interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingFlagsHmac from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmac) (PatternFlowIpv6SegmentRoutingFlagsHmac, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingFlagsHmac from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingFlagsHmac from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingFlagsHmac from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHmac {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingFlagsHmac{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHmac {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmac) ToProto() (*otg.PatternFlowIpv6SegmentRoutingFlagsHmac, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingFlagsHmac) (PatternFlowIpv6SegmentRoutingFlagsHmac, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmac) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmac) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingFlagsHmac) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingFlagsHmac) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Clone() (PatternFlowIpv6SegmentRoutingFlagsHmac, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingFlagsHmac()
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingFlagsHmac is hMAC Flag. Indicates the presence of the HMAC TLV for verification.
type PatternFlowIpv6SegmentRoutingFlagsHmac interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingFlagsHmac to protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingFlagsHmac from protobuf object *otg.PatternFlowIpv6SegmentRoutingFlagsHmac
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingFlagsHmac) PatternFlowIpv6SegmentRoutingFlagsHmac
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingFlagsHmac
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingFlagsHmac
	// validate validates PatternFlowIpv6SegmentRoutingFlagsHmac
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingFlagsHmac, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum, set in PatternFlowIpv6SegmentRoutingFlagsHmac
	Choice() PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingFlagsHmac
	setChoice(value PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsHmac
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingFlagsHmac
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingFlagsHmac.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHmac
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmac
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingFlagsHmac
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingFlagsHmac.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingFlagsHmac
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsHmac
	// Increment returns PatternFlowIpv6SegmentRoutingFlagsHmacCounter, set in PatternFlowIpv6SegmentRoutingFlagsHmac.
	// PatternFlowIpv6SegmentRoutingFlagsHmacCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingFlagsHmacCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsHmac.
	// PatternFlowIpv6SegmentRoutingFlagsHmacCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmac
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingFlagsHmac
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingFlagsHmacCounter, set in PatternFlowIpv6SegmentRoutingFlagsHmac.
	// PatternFlowIpv6SegmentRoutingFlagsHmacCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingFlagsHmacCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingFlagsHmacCounter provided by user to PatternFlowIpv6SegmentRoutingFlagsHmac.
	// PatternFlowIpv6SegmentRoutingFlagsHmacCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmac
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingFlagsHmac
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingFlagsHmac
var PatternFlowIpv6SegmentRoutingFlagsHmacChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Choice() PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum {
	return PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) setChoice(value PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum) PatternFlowIpv6SegmentRoutingFlagsHmac {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingFlagsHmac_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsHmac_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHmacChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingFlagsHmacCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingFlagsHmacChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingFlagsHmacCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHmac object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) SetValue(value uint32) PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingFlagsHmac object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsHmacCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Increment() PatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingFlagsHmacCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingFlagsHmacCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingFlagsHmacCounter value in the PatternFlowIpv6SegmentRoutingFlagsHmac object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) SetIncrement(value PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsHmacCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) Decrement() PatternFlowIpv6SegmentRoutingFlagsHmacCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingFlagsHmacCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingFlagsHmacCounter
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingFlagsHmacCounter value in the PatternFlowIpv6SegmentRoutingFlagsHmac object
func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) SetDecrement(value PatternFlowIpv6SegmentRoutingFlagsHmacCounter) PatternFlowIpv6SegmentRoutingFlagsHmac {
	obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingFlagsHmac.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingFlagsHmac.Values <= 1 but Got %d", item))
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

func (obj *patternFlowIpv6SegmentRoutingFlagsHmac) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingFlagsHmacChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHmacChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingFlagsHmacChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingFlagsHmacChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingFlagsHmac")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingFlagsHmac_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingFlagsHmac_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
