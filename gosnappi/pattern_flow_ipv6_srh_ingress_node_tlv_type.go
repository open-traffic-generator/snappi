package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHIngressNodeTlvType *****
type patternFlowIpv6SRHIngressNodeTlvType struct {
	validation
	obj             *otg.PatternFlowIpv6SRHIngressNodeTlvType
	marshaller      marshalPatternFlowIpv6SRHIngressNodeTlvType
	unMarshaller    unMarshalPatternFlowIpv6SRHIngressNodeTlvType
	incrementHolder PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	decrementHolder PatternFlowIpv6SRHIngressNodeTlvTypeCounter
}

func NewPatternFlowIpv6SRHIngressNodeTlvType() PatternFlowIpv6SRHIngressNodeTlvType {
	obj := patternFlowIpv6SRHIngressNodeTlvType{obj: &otg.PatternFlowIpv6SRHIngressNodeTlvType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) msg() *otg.PatternFlowIpv6SRHIngressNodeTlvType {
	return obj.obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) setMsg(msg *otg.PatternFlowIpv6SRHIngressNodeTlvType) PatternFlowIpv6SRHIngressNodeTlvType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHIngressNodeTlvType struct {
	obj *patternFlowIpv6SRHIngressNodeTlvType
}

type marshalPatternFlowIpv6SRHIngressNodeTlvType interface {
	// ToProto marshals PatternFlowIpv6SRHIngressNodeTlvType to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvType
	ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvType, error)
	// ToPbText marshals PatternFlowIpv6SRHIngressNodeTlvType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHIngressNodeTlvType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHIngressNodeTlvType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHIngressNodeTlvType struct {
	obj *patternFlowIpv6SRHIngressNodeTlvType
}

type unMarshalPatternFlowIpv6SRHIngressNodeTlvType interface {
	// FromProto unmarshals PatternFlowIpv6SRHIngressNodeTlvType from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvType
	FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvType) (PatternFlowIpv6SRHIngressNodeTlvType, error)
	// FromPbText unmarshals PatternFlowIpv6SRHIngressNodeTlvType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHIngressNodeTlvType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHIngressNodeTlvType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHIngressNodeTlvType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHIngressNodeTlvType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvType) ToProto() (*otg.PatternFlowIpv6SRHIngressNodeTlvType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvType) FromProto(msg *otg.PatternFlowIpv6SRHIngressNodeTlvType) (PatternFlowIpv6SRHIngressNodeTlvType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHIngressNodeTlvType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHIngressNodeTlvType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHIngressNodeTlvType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) Clone() (PatternFlowIpv6SRHIngressNodeTlvType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHIngressNodeTlvType()
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

func (obj *patternFlowIpv6SRHIngressNodeTlvType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHIngressNodeTlvType is tLV type code (RFC 9259 Section 3.1). The canonical value is 1. Set to a different value only for negative or conformance testing.
type PatternFlowIpv6SRHIngressNodeTlvType interface {
	Validation
	// msg marshals PatternFlowIpv6SRHIngressNodeTlvType to protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHIngressNodeTlvType
	// setMsg unmarshals PatternFlowIpv6SRHIngressNodeTlvType from protobuf object *otg.PatternFlowIpv6SRHIngressNodeTlvType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHIngressNodeTlvType) PatternFlowIpv6SRHIngressNodeTlvType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHIngressNodeTlvType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHIngressNodeTlvType
	// validate validates PatternFlowIpv6SRHIngressNodeTlvType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHIngressNodeTlvType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum, set in PatternFlowIpv6SRHIngressNodeTlvType
	Choice() PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum provided by user to PatternFlowIpv6SRHIngressNodeTlvType
	setChoice(value PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHIngressNodeTlvType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvType
	SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvType
	// HasValue checks if Value has been set in PatternFlowIpv6SRHIngressNodeTlvType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHIngressNodeTlvType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHIngressNodeTlvType
	SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvType
	// Auto returns uint32, set in PatternFlowIpv6SRHIngressNodeTlvType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHIngressNodeTlvType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHIngressNodeTlvTypeCounter, set in PatternFlowIpv6SRHIngressNodeTlvType.
	// PatternFlowIpv6SRHIngressNodeTlvTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// SetIncrement assigns PatternFlowIpv6SRHIngressNodeTlvTypeCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvType.
	// PatternFlowIpv6SRHIngressNodeTlvTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHIngressNodeTlvType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHIngressNodeTlvTypeCounter, set in PatternFlowIpv6SRHIngressNodeTlvType.
	// PatternFlowIpv6SRHIngressNodeTlvTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHIngressNodeTlvTypeCounter
	// SetDecrement assigns PatternFlowIpv6SRHIngressNodeTlvTypeCounter provided by user to PatternFlowIpv6SRHIngressNodeTlvType.
	// PatternFlowIpv6SRHIngressNodeTlvTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHIngressNodeTlvType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHIngressNodeTlvType
var PatternFlowIpv6SRHIngressNodeTlvTypeChoice = struct {
	VALUE     PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
	VALUES    PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
	AUTO      PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
	INCREMENT PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
	DECREMENT PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) Choice() PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum {
	return PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHIngressNodeTlvType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) setChoice(value PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum) PatternFlowIpv6SRHIngressNodeTlvType {
	intValue, ok := otg.PatternFlowIpv6SRHIngressNodeTlvType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvTypeChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHIngressNodeTlvTypeCounter().msg()
	}

	if value == PatternFlowIpv6SRHIngressNodeTlvTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHIngressNodeTlvTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHIngressNodeTlvType object
func (obj *patternFlowIpv6SRHIngressNodeTlvType) SetValue(value uint32) PatternFlowIpv6SRHIngressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHIngressNodeTlvType object
func (obj *patternFlowIpv6SRHIngressNodeTlvType) SetValues(value []uint32) PatternFlowIpv6SRHIngressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SRHIngressNodeTlvType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHIngressNodeTlvType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvType) Increment() PatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHIngressNodeTlvTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHIngressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHIngressNodeTlvTypeCounter value in the PatternFlowIpv6SRHIngressNodeTlvType object
func (obj *patternFlowIpv6SRHIngressNodeTlvType) SetIncrement(value PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvType) Decrement() PatternFlowIpv6SRHIngressNodeTlvTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHIngressNodeTlvTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHIngressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHIngressNodeTlvType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHIngressNodeTlvTypeCounter value in the PatternFlowIpv6SRHIngressNodeTlvType object
func (obj *patternFlowIpv6SRHIngressNodeTlvType) SetDecrement(value PatternFlowIpv6SRHIngressNodeTlvTypeCounter) PatternFlowIpv6SRHIngressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHIngressNodeTlvType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHIngressNodeTlvType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHIngressNodeTlvType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHIngressNodeTlvTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHIngressNodeTlvTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHIngressNodeTlvTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHIngressNodeTlvType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHIngressNodeTlvType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHIngressNodeTlvType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
