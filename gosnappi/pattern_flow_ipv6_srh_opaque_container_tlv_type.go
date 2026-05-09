package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHOpaqueContainerTlvType *****
type patternFlowIpv6SRHOpaqueContainerTlvType struct {
	validation
	obj             *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	marshaller      marshalPatternFlowIpv6SRHOpaqueContainerTlvType
	unMarshaller    unMarshalPatternFlowIpv6SRHOpaqueContainerTlvType
	incrementHolder PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	decrementHolder PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
}

func NewPatternFlowIpv6SRHOpaqueContainerTlvType() PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj := patternFlowIpv6SRHOpaqueContainerTlvType{obj: &otg.PatternFlowIpv6SRHOpaqueContainerTlvType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvType {
	return obj.obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) setMsg(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvType) PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHOpaqueContainerTlvType struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvType
}

type marshalPatternFlowIpv6SRHOpaqueContainerTlvType interface {
	// ToProto marshals PatternFlowIpv6SRHOpaqueContainerTlvType to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvType, error)
	// ToPbText marshals PatternFlowIpv6SRHOpaqueContainerTlvType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHOpaqueContainerTlvType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHOpaqueContainerTlvType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType struct {
	obj *patternFlowIpv6SRHOpaqueContainerTlvType
}

type unMarshalPatternFlowIpv6SRHOpaqueContainerTlvType interface {
	// FromProto unmarshals PatternFlowIpv6SRHOpaqueContainerTlvType from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvType) (PatternFlowIpv6SRHOpaqueContainerTlvType, error)
	// FromPbText unmarshals PatternFlowIpv6SRHOpaqueContainerTlvType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHOpaqueContainerTlvType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHOpaqueContainerTlvType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHOpaqueContainerTlvType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvType) ToProto() (*otg.PatternFlowIpv6SRHOpaqueContainerTlvType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType) FromProto(msg *otg.PatternFlowIpv6SRHOpaqueContainerTlvType) (PatternFlowIpv6SRHOpaqueContainerTlvType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHOpaqueContainerTlvType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHOpaqueContainerTlvType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Clone() (PatternFlowIpv6SRHOpaqueContainerTlvType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHOpaqueContainerTlvType()
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

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHOpaqueContainerTlvType is tLV type code (RFC 8754 Section 2.1). The canonical value is 3. Set to a different value only for negative or conformance testing.
type PatternFlowIpv6SRHOpaqueContainerTlvType interface {
	Validation
	// msg marshals PatternFlowIpv6SRHOpaqueContainerTlvType to protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	// setMsg unmarshals PatternFlowIpv6SRHOpaqueContainerTlvType from protobuf object *otg.PatternFlowIpv6SRHOpaqueContainerTlvType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHOpaqueContainerTlvType) PatternFlowIpv6SRHOpaqueContainerTlvType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHOpaqueContainerTlvType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHOpaqueContainerTlvType
	// validate validates PatternFlowIpv6SRHOpaqueContainerTlvType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHOpaqueContainerTlvType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum, set in PatternFlowIpv6SRHOpaqueContainerTlvType
	Choice() PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum provided by user to PatternFlowIpv6SRHOpaqueContainerTlvType
	setChoice(value PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum) PatternFlowIpv6SRHOpaqueContainerTlvType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHOpaqueContainerTlvType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvType
	SetValue(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvType
	// HasValue checks if Value has been set in PatternFlowIpv6SRHOpaqueContainerTlvType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHOpaqueContainerTlvType
	SetValues(value []uint32) PatternFlowIpv6SRHOpaqueContainerTlvType
	// Auto returns uint32, set in PatternFlowIpv6SRHOpaqueContainerTlvType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHOpaqueContainerTlvType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, set in PatternFlowIpv6SRHOpaqueContainerTlvType.
	// PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// SetIncrement assigns PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter provided by user to PatternFlowIpv6SRHOpaqueContainerTlvType.
	// PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHOpaqueContainerTlvType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter, set in PatternFlowIpv6SRHOpaqueContainerTlvType.
	// PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
	// SetDecrement assigns PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter provided by user to PatternFlowIpv6SRHOpaqueContainerTlvType.
	// PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHOpaqueContainerTlvType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHOpaqueContainerTlvType
var PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice = struct {
	VALUE     PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
	VALUES    PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
	AUTO      PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
	INCREMENT PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
	DECREMENT PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Choice() PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum {
	return PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) setChoice(value PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum) PatternFlowIpv6SRHOpaqueContainerTlvType {
	intValue, ok := otg.PatternFlowIpv6SRHOpaqueContainerTlvType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHOpaqueContainerTlvType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUE {
		defaultValue := uint32(3)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUES {
		defaultValue := []uint32{3}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.AUTO {
		defaultValue := uint32(3)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter().msg()
	}

	if value == PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHOpaqueContainerTlvTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvType object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) SetValue(value uint32) PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{3})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHOpaqueContainerTlvType object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) SetValues(value []uint32) PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Increment() PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHOpaqueContainerTlvTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter value in the PatternFlowIpv6SRHOpaqueContainerTlvType object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) SetIncrement(value PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) Decrement() PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHOpaqueContainerTlvTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter value in the PatternFlowIpv6SRHOpaqueContainerTlvType object
func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) SetDecrement(value PatternFlowIpv6SRHOpaqueContainerTlvTypeCounter) PatternFlowIpv6SRHOpaqueContainerTlvType {
	obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHOpaqueContainerTlvType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHOpaqueContainerTlvType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHOpaqueContainerTlvType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHOpaqueContainerTlvTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHOpaqueContainerTlvTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHOpaqueContainerTlvType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHOpaqueContainerTlvType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHOpaqueContainerTlvType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
