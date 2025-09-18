package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6ExtHeaderHdrExtLen *****
type patternFlowIpv6ExtHeaderHdrExtLen struct {
	validation
	obj             *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	marshaller      marshalPatternFlowIpv6ExtHeaderHdrExtLen
	unMarshaller    unMarshalPatternFlowIpv6ExtHeaderHdrExtLen
	incrementHolder PatternFlowIpv6ExtHeaderHdrExtLenCounter
	decrementHolder PatternFlowIpv6ExtHeaderHdrExtLenCounter
}

func NewPatternFlowIpv6ExtHeaderHdrExtLen() PatternFlowIpv6ExtHeaderHdrExtLen {
	obj := patternFlowIpv6ExtHeaderHdrExtLen{obj: &otg.PatternFlowIpv6ExtHeaderHdrExtLen{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) msg() *otg.PatternFlowIpv6ExtHeaderHdrExtLen {
	return obj.obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) setMsg(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLen) PatternFlowIpv6ExtHeaderHdrExtLen {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6ExtHeaderHdrExtLen struct {
	obj *patternFlowIpv6ExtHeaderHdrExtLen
}

type marshalPatternFlowIpv6ExtHeaderHdrExtLen interface {
	// ToProto marshals PatternFlowIpv6ExtHeaderHdrExtLen to protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	ToProto() (*otg.PatternFlowIpv6ExtHeaderHdrExtLen, error)
	// ToPbText marshals PatternFlowIpv6ExtHeaderHdrExtLen to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6ExtHeaderHdrExtLen to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6ExtHeaderHdrExtLen to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6ExtHeaderHdrExtLen struct {
	obj *patternFlowIpv6ExtHeaderHdrExtLen
}

type unMarshalPatternFlowIpv6ExtHeaderHdrExtLen interface {
	// FromProto unmarshals PatternFlowIpv6ExtHeaderHdrExtLen from protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	FromProto(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLen) (PatternFlowIpv6ExtHeaderHdrExtLen, error)
	// FromPbText unmarshals PatternFlowIpv6ExtHeaderHdrExtLen from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6ExtHeaderHdrExtLen from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6ExtHeaderHdrExtLen from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Marshal() marshalPatternFlowIpv6ExtHeaderHdrExtLen {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6ExtHeaderHdrExtLen{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Unmarshal() unMarshalPatternFlowIpv6ExtHeaderHdrExtLen {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6ExtHeaderHdrExtLen{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLen) ToProto() (*otg.PatternFlowIpv6ExtHeaderHdrExtLen, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLen) FromProto(msg *otg.PatternFlowIpv6ExtHeaderHdrExtLen) (PatternFlowIpv6ExtHeaderHdrExtLen, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLen) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLen) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLen) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLen) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderHdrExtLen) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderHdrExtLen) FromJson(value string) error {
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

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Clone() (PatternFlowIpv6ExtHeaderHdrExtLen, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6ExtHeaderHdrExtLen()
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

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6ExtHeaderHdrExtLen is 8-bit unsigned integer specifying the length of the SRH in 8-octet units, not including the first 8 octets of the SRH itself. When AUTO is assigned it is set to the correct value.
type PatternFlowIpv6ExtHeaderHdrExtLen interface {
	Validation
	// msg marshals PatternFlowIpv6ExtHeaderHdrExtLen to protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	// setMsg unmarshals PatternFlowIpv6ExtHeaderHdrExtLen from protobuf object *otg.PatternFlowIpv6ExtHeaderHdrExtLen
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6ExtHeaderHdrExtLen) PatternFlowIpv6ExtHeaderHdrExtLen
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6ExtHeaderHdrExtLen
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6ExtHeaderHdrExtLen
	// validate validates PatternFlowIpv6ExtHeaderHdrExtLen
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6ExtHeaderHdrExtLen, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum, set in PatternFlowIpv6ExtHeaderHdrExtLen
	Choice() PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
	// setChoice assigns PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum provided by user to PatternFlowIpv6ExtHeaderHdrExtLen
	setChoice(value PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum) PatternFlowIpv6ExtHeaderHdrExtLen
	// HasChoice checks if Choice has been set in PatternFlowIpv6ExtHeaderHdrExtLen
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6ExtHeaderHdrExtLen.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6ExtHeaderHdrExtLen
	SetValue(value uint32) PatternFlowIpv6ExtHeaderHdrExtLen
	// HasValue checks if Value has been set in PatternFlowIpv6ExtHeaderHdrExtLen
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6ExtHeaderHdrExtLen.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6ExtHeaderHdrExtLen
	SetValues(value []uint32) PatternFlowIpv6ExtHeaderHdrExtLen
	// Auto returns uint32, set in PatternFlowIpv6ExtHeaderHdrExtLen.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6ExtHeaderHdrExtLen
	HasAuto() bool
	// Increment returns PatternFlowIpv6ExtHeaderHdrExtLenCounter, set in PatternFlowIpv6ExtHeaderHdrExtLen.
	// PatternFlowIpv6ExtHeaderHdrExtLenCounter is integer counter pattern
	Increment() PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// SetIncrement assigns PatternFlowIpv6ExtHeaderHdrExtLenCounter provided by user to PatternFlowIpv6ExtHeaderHdrExtLen.
	// PatternFlowIpv6ExtHeaderHdrExtLenCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLen
	// HasIncrement checks if Increment has been set in PatternFlowIpv6ExtHeaderHdrExtLen
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6ExtHeaderHdrExtLenCounter, set in PatternFlowIpv6ExtHeaderHdrExtLen.
	// PatternFlowIpv6ExtHeaderHdrExtLenCounter is integer counter pattern
	Decrement() PatternFlowIpv6ExtHeaderHdrExtLenCounter
	// SetDecrement assigns PatternFlowIpv6ExtHeaderHdrExtLenCounter provided by user to PatternFlowIpv6ExtHeaderHdrExtLen.
	// PatternFlowIpv6ExtHeaderHdrExtLenCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLen
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6ExtHeaderHdrExtLen
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum string

// Enum of Choice on PatternFlowIpv6ExtHeaderHdrExtLen
var PatternFlowIpv6ExtHeaderHdrExtLenChoice = struct {
	VALUE     PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
	VALUES    PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
	AUTO      PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
	INCREMENT PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
	DECREMENT PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum
}{
	VALUE:     PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum("value"),
	VALUES:    PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum("values"),
	AUTO:      PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Choice() PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum {
	return PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) setChoice(value PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum) PatternFlowIpv6ExtHeaderHdrExtLen {
	intValue, ok := otg.PatternFlowIpv6ExtHeaderHdrExtLen_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6ExtHeaderHdrExtLen_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderHdrExtLenChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderHdrExtLenChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6ExtHeaderHdrExtLenCounter().msg()
	}

	if value == PatternFlowIpv6ExtHeaderHdrExtLenChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6ExtHeaderHdrExtLenCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6ExtHeaderHdrExtLen object
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) SetValue(value uint32) PatternFlowIpv6ExtHeaderHdrExtLen {
	obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6ExtHeaderHdrExtLen object
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) SetValues(value []uint32) PatternFlowIpv6ExtHeaderHdrExtLen {
	obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6ExtHeaderHdrExtLenCounter
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Increment() PatternFlowIpv6ExtHeaderHdrExtLenCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6ExtHeaderHdrExtLenCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6ExtHeaderHdrExtLenCounter
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6ExtHeaderHdrExtLenCounter value in the PatternFlowIpv6ExtHeaderHdrExtLen object
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) SetIncrement(value PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLen {
	obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6ExtHeaderHdrExtLenCounter
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) Decrement() PatternFlowIpv6ExtHeaderHdrExtLenCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6ExtHeaderHdrExtLenCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6ExtHeaderHdrExtLenCounter
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6ExtHeaderHdrExtLenCounter value in the PatternFlowIpv6ExtHeaderHdrExtLen object
func (obj *patternFlowIpv6ExtHeaderHdrExtLen) SetDecrement(value PatternFlowIpv6ExtHeaderHdrExtLenCounter) PatternFlowIpv6ExtHeaderHdrExtLen {
	obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderHdrExtLen.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6ExtHeaderHdrExtLen.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderHdrExtLen.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6ExtHeaderHdrExtLen) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6ExtHeaderHdrExtLenChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderHdrExtLenChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderHdrExtLenChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderHdrExtLenChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderHdrExtLenChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6ExtHeaderHdrExtLenChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6ExtHeaderHdrExtLen")
			}
		} else {
			intVal := otg.PatternFlowIpv6ExtHeaderHdrExtLen_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6ExtHeaderHdrExtLen_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
