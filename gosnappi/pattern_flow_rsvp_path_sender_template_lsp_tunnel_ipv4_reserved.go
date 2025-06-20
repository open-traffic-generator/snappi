package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved *****
type patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	marshaller      marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	incrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	decrementHolder PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
}

func NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj := patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved{obj: &otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) setMsg(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
}

type marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved interface {
	// ToProto marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved struct {
	obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
}

type unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) ToProto() (*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FromProto(msg *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved()
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved is reserved field, MUST be zero.
type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved to protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// setMsg unmarshals PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved from protobuf object *otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// validate validates PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	SetValue(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	SetValues(value []uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// Increment returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter, set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter provided by user to PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.
	// PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved
var PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Choice() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum {
	return PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) setChoice(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	intValue, ok := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) SetValue(value uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) SetValues(value []uint32) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Increment() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) SetIncrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) Decrement() PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter value in the PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved object
func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) SetDecrement(value PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedCounter) PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved {
	obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTemplateLspTunnelIpv4ReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTemplateLspTunnelIpv4Reserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
