package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPadTlvType *****
type patternFlowIpv6SRHPadTlvType struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPadTlvType
	marshaller      marshalPatternFlowIpv6SRHPadTlvType
	unMarshaller    unMarshalPatternFlowIpv6SRHPadTlvType
	incrementHolder PatternFlowIpv6SRHPadTlvTypeCounter
	decrementHolder PatternFlowIpv6SRHPadTlvTypeCounter
}

func NewPatternFlowIpv6SRHPadTlvType() PatternFlowIpv6SRHPadTlvType {
	obj := patternFlowIpv6SRHPadTlvType{obj: &otg.PatternFlowIpv6SRHPadTlvType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPadTlvType) msg() *otg.PatternFlowIpv6SRHPadTlvType {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPadTlvType) setMsg(msg *otg.PatternFlowIpv6SRHPadTlvType) PatternFlowIpv6SRHPadTlvType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPadTlvType struct {
	obj *patternFlowIpv6SRHPadTlvType
}

type marshalPatternFlowIpv6SRHPadTlvType interface {
	// ToProto marshals PatternFlowIpv6SRHPadTlvType to protobuf object *otg.PatternFlowIpv6SRHPadTlvType
	ToProto() (*otg.PatternFlowIpv6SRHPadTlvType, error)
	// ToPbText marshals PatternFlowIpv6SRHPadTlvType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPadTlvType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPadTlvType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPadTlvType struct {
	obj *patternFlowIpv6SRHPadTlvType
}

type unMarshalPatternFlowIpv6SRHPadTlvType interface {
	// FromProto unmarshals PatternFlowIpv6SRHPadTlvType from protobuf object *otg.PatternFlowIpv6SRHPadTlvType
	FromProto(msg *otg.PatternFlowIpv6SRHPadTlvType) (PatternFlowIpv6SRHPadTlvType, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPadTlvType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPadTlvType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPadTlvType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPadTlvType) Marshal() marshalPatternFlowIpv6SRHPadTlvType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPadTlvType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPadTlvType) Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPadTlvType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPadTlvType) ToProto() (*otg.PatternFlowIpv6SRHPadTlvType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPadTlvType) FromProto(msg *otg.PatternFlowIpv6SRHPadTlvType) (PatternFlowIpv6SRHPadTlvType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPadTlvType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPadTlvType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPadTlvType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPadTlvType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPadTlvType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPadTlvType) Clone() (PatternFlowIpv6SRHPadTlvType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPadTlvType()
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

func (obj *patternFlowIpv6SRHPadTlvType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPadTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 4. Set to a different value only for negative or conformance testing.
type PatternFlowIpv6SRHPadTlvType interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPadTlvType to protobuf object *otg.PatternFlowIpv6SRHPadTlvType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPadTlvType
	// setMsg unmarshals PatternFlowIpv6SRHPadTlvType from protobuf object *otg.PatternFlowIpv6SRHPadTlvType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPadTlvType) PatternFlowIpv6SRHPadTlvType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPadTlvType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPadTlvType
	// validate validates PatternFlowIpv6SRHPadTlvType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPadTlvType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPadTlvTypeChoiceEnum, set in PatternFlowIpv6SRHPadTlvType
	Choice() PatternFlowIpv6SRHPadTlvTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPadTlvTypeChoiceEnum provided by user to PatternFlowIpv6SRHPadTlvType
	setChoice(value PatternFlowIpv6SRHPadTlvTypeChoiceEnum) PatternFlowIpv6SRHPadTlvType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPadTlvType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPadTlvType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPadTlvType
	SetValue(value uint32) PatternFlowIpv6SRHPadTlvType
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPadTlvType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPadTlvType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPadTlvType
	SetValues(value []uint32) PatternFlowIpv6SRHPadTlvType
	// Auto returns uint32, set in PatternFlowIpv6SRHPadTlvType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHPadTlvType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHPadTlvTypeCounter, set in PatternFlowIpv6SRHPadTlvType.
	// PatternFlowIpv6SRHPadTlvTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPadTlvTypeCounter
	// SetIncrement assigns PatternFlowIpv6SRHPadTlvTypeCounter provided by user to PatternFlowIpv6SRHPadTlvType.
	// PatternFlowIpv6SRHPadTlvTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPadTlvType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPadTlvTypeCounter, set in PatternFlowIpv6SRHPadTlvType.
	// PatternFlowIpv6SRHPadTlvTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPadTlvTypeCounter
	// SetDecrement assigns PatternFlowIpv6SRHPadTlvTypeCounter provided by user to PatternFlowIpv6SRHPadTlvType.
	// PatternFlowIpv6SRHPadTlvTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPadTlvType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPadTlvTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPadTlvType
var PatternFlowIpv6SRHPadTlvTypeChoice = struct {
	VALUE     PatternFlowIpv6SRHPadTlvTypeChoiceEnum
	VALUES    PatternFlowIpv6SRHPadTlvTypeChoiceEnum
	AUTO      PatternFlowIpv6SRHPadTlvTypeChoiceEnum
	INCREMENT PatternFlowIpv6SRHPadTlvTypeChoiceEnum
	DECREMENT PatternFlowIpv6SRHPadTlvTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPadTlvTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPadTlvTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHPadTlvTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHPadTlvTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPadTlvTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPadTlvType) Choice() PatternFlowIpv6SRHPadTlvTypeChoiceEnum {
	return PatternFlowIpv6SRHPadTlvTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPadTlvType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPadTlvType) setChoice(value PatternFlowIpv6SRHPadTlvTypeChoiceEnum) PatternFlowIpv6SRHPadTlvType {
	intValue, ok := otg.PatternFlowIpv6SRHPadTlvType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPadTlvTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPadTlvType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPadTlvTypeChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvTypeChoice.VALUES {
		defaultValue := []uint32{4}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvTypeChoice.AUTO {
		defaultValue := uint32(4)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHPadTlvTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPadTlvTypeCounter().msg()
	}

	if value == PatternFlowIpv6SRHPadTlvTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPadTlvTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPadTlvType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPadTlvType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPadTlvType object
func (obj *patternFlowIpv6SRHPadTlvType) SetValue(value uint32) PatternFlowIpv6SRHPadTlvType {
	obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPadTlvType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPadTlvType object
func (obj *patternFlowIpv6SRHPadTlvType) SetValues(value []uint32) PatternFlowIpv6SRHPadTlvType {
	obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SRHPadTlvType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHPadTlvType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPadTlvTypeCounter
func (obj *patternFlowIpv6SRHPadTlvType) Increment() PatternFlowIpv6SRHPadTlvTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPadTlvTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPadTlvTypeCounter
func (obj *patternFlowIpv6SRHPadTlvType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPadTlvTypeCounter value in the PatternFlowIpv6SRHPadTlvType object
func (obj *patternFlowIpv6SRHPadTlvType) SetIncrement(value PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvType {
	obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPadTlvTypeCounter
func (obj *patternFlowIpv6SRHPadTlvType) Decrement() PatternFlowIpv6SRHPadTlvTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPadTlvTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPadTlvTypeCounter
func (obj *patternFlowIpv6SRHPadTlvType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPadTlvTypeCounter value in the PatternFlowIpv6SRHPadTlvType object
func (obj *patternFlowIpv6SRHPadTlvType) SetDecrement(value PatternFlowIpv6SRHPadTlvTypeCounter) PatternFlowIpv6SRHPadTlvType {
	obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPadTlvType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHPadTlvType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPadTlvType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPadTlvType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPadTlvTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPadTlvTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPadTlvTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPadTlvType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPadTlvType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPadTlvType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
