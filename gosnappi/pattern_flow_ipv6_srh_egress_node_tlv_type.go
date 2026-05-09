package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHEgressNodeTlvType *****
type patternFlowIpv6SRHEgressNodeTlvType struct {
	validation
	obj             *otg.PatternFlowIpv6SRHEgressNodeTlvType
	marshaller      marshalPatternFlowIpv6SRHEgressNodeTlvType
	unMarshaller    unMarshalPatternFlowIpv6SRHEgressNodeTlvType
	incrementHolder PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	decrementHolder PatternFlowIpv6SRHEgressNodeTlvTypeCounter
}

func NewPatternFlowIpv6SRHEgressNodeTlvType() PatternFlowIpv6SRHEgressNodeTlvType {
	obj := patternFlowIpv6SRHEgressNodeTlvType{obj: &otg.PatternFlowIpv6SRHEgressNodeTlvType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) msg() *otg.PatternFlowIpv6SRHEgressNodeTlvType {
	return obj.obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) setMsg(msg *otg.PatternFlowIpv6SRHEgressNodeTlvType) PatternFlowIpv6SRHEgressNodeTlvType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHEgressNodeTlvType struct {
	obj *patternFlowIpv6SRHEgressNodeTlvType
}

type marshalPatternFlowIpv6SRHEgressNodeTlvType interface {
	// ToProto marshals PatternFlowIpv6SRHEgressNodeTlvType to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvType
	ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvType, error)
	// ToPbText marshals PatternFlowIpv6SRHEgressNodeTlvType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHEgressNodeTlvType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHEgressNodeTlvType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHEgressNodeTlvType struct {
	obj *patternFlowIpv6SRHEgressNodeTlvType
}

type unMarshalPatternFlowIpv6SRHEgressNodeTlvType interface {
	// FromProto unmarshals PatternFlowIpv6SRHEgressNodeTlvType from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvType
	FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvType) (PatternFlowIpv6SRHEgressNodeTlvType, error)
	// FromPbText unmarshals PatternFlowIpv6SRHEgressNodeTlvType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHEgressNodeTlvType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHEgressNodeTlvType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHEgressNodeTlvType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHEgressNodeTlvType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvType) ToProto() (*otg.PatternFlowIpv6SRHEgressNodeTlvType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvType) FromProto(msg *otg.PatternFlowIpv6SRHEgressNodeTlvType) (PatternFlowIpv6SRHEgressNodeTlvType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHEgressNodeTlvType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHEgressNodeTlvType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHEgressNodeTlvType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) Clone() (PatternFlowIpv6SRHEgressNodeTlvType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHEgressNodeTlvType()
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

func (obj *patternFlowIpv6SRHEgressNodeTlvType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHEgressNodeTlvType is tLV type code (RFC 9259 Section 3.2). The canonical value is 2. Set to a different value only for negative or conformance testing.
type PatternFlowIpv6SRHEgressNodeTlvType interface {
	Validation
	// msg marshals PatternFlowIpv6SRHEgressNodeTlvType to protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHEgressNodeTlvType
	// setMsg unmarshals PatternFlowIpv6SRHEgressNodeTlvType from protobuf object *otg.PatternFlowIpv6SRHEgressNodeTlvType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHEgressNodeTlvType) PatternFlowIpv6SRHEgressNodeTlvType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHEgressNodeTlvType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHEgressNodeTlvType
	// validate validates PatternFlowIpv6SRHEgressNodeTlvType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHEgressNodeTlvType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum, set in PatternFlowIpv6SRHEgressNodeTlvType
	Choice() PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum provided by user to PatternFlowIpv6SRHEgressNodeTlvType
	setChoice(value PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHEgressNodeTlvType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvType
	SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvType
	// HasValue checks if Value has been set in PatternFlowIpv6SRHEgressNodeTlvType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHEgressNodeTlvType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHEgressNodeTlvType
	SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvType
	// Auto returns uint32, set in PatternFlowIpv6SRHEgressNodeTlvType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHEgressNodeTlvType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHEgressNodeTlvTypeCounter, set in PatternFlowIpv6SRHEgressNodeTlvType.
	// PatternFlowIpv6SRHEgressNodeTlvTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// SetIncrement assigns PatternFlowIpv6SRHEgressNodeTlvTypeCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvType.
	// PatternFlowIpv6SRHEgressNodeTlvTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHEgressNodeTlvType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHEgressNodeTlvTypeCounter, set in PatternFlowIpv6SRHEgressNodeTlvType.
	// PatternFlowIpv6SRHEgressNodeTlvTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHEgressNodeTlvTypeCounter
	// SetDecrement assigns PatternFlowIpv6SRHEgressNodeTlvTypeCounter provided by user to PatternFlowIpv6SRHEgressNodeTlvType.
	// PatternFlowIpv6SRHEgressNodeTlvTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHEgressNodeTlvType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHEgressNodeTlvType
var PatternFlowIpv6SRHEgressNodeTlvTypeChoice = struct {
	VALUE     PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
	VALUES    PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
	AUTO      PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
	INCREMENT PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
	DECREMENT PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) Choice() PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum {
	return PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHEgressNodeTlvType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) setChoice(value PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum) PatternFlowIpv6SRHEgressNodeTlvType {
	intValue, ok := otg.PatternFlowIpv6SRHEgressNodeTlvType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUE {
		defaultValue := uint32(2)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUES {
		defaultValue := []uint32{2}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvTypeChoice.AUTO {
		defaultValue := uint32(2)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHEgressNodeTlvTypeCounter().msg()
	}

	if value == PatternFlowIpv6SRHEgressNodeTlvTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHEgressNodeTlvTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHEgressNodeTlvType object
func (obj *patternFlowIpv6SRHEgressNodeTlvType) SetValue(value uint32) PatternFlowIpv6SRHEgressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHEgressNodeTlvType object
func (obj *patternFlowIpv6SRHEgressNodeTlvType) SetValues(value []uint32) PatternFlowIpv6SRHEgressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SRHEgressNodeTlvType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHEgressNodeTlvType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvType) Increment() PatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHEgressNodeTlvTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHEgressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHEgressNodeTlvTypeCounter value in the PatternFlowIpv6SRHEgressNodeTlvType object
func (obj *patternFlowIpv6SRHEgressNodeTlvType) SetIncrement(value PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvType) Decrement() PatternFlowIpv6SRHEgressNodeTlvTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHEgressNodeTlvTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHEgressNodeTlvTypeCounter
func (obj *patternFlowIpv6SRHEgressNodeTlvType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHEgressNodeTlvTypeCounter value in the PatternFlowIpv6SRHEgressNodeTlvType object
func (obj *patternFlowIpv6SRHEgressNodeTlvType) SetDecrement(value PatternFlowIpv6SRHEgressNodeTlvTypeCounter) PatternFlowIpv6SRHEgressNodeTlvType {
	obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHEgressNodeTlvType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHEgressNodeTlvType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHEgressNodeTlvType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHEgressNodeTlvTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHEgressNodeTlvTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHEgressNodeTlvTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHEgressNodeTlvType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHEgressNodeTlvType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHEgressNodeTlvType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
