package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SegmentRoutingRoutingType *****
type patternFlowIpv6SegmentRoutingRoutingType struct {
	validation
	obj             *otg.PatternFlowIpv6SegmentRoutingRoutingType
	marshaller      marshalPatternFlowIpv6SegmentRoutingRoutingType
	unMarshaller    unMarshalPatternFlowIpv6SegmentRoutingRoutingType
	incrementHolder PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	decrementHolder PatternFlowIpv6SegmentRoutingRoutingTypeCounter
}

func NewPatternFlowIpv6SegmentRoutingRoutingType() PatternFlowIpv6SegmentRoutingRoutingType {
	obj := patternFlowIpv6SegmentRoutingRoutingType{obj: &otg.PatternFlowIpv6SegmentRoutingRoutingType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) msg() *otg.PatternFlowIpv6SegmentRoutingRoutingType {
	return obj.obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) setMsg(msg *otg.PatternFlowIpv6SegmentRoutingRoutingType) PatternFlowIpv6SegmentRoutingRoutingType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SegmentRoutingRoutingType struct {
	obj *patternFlowIpv6SegmentRoutingRoutingType
}

type marshalPatternFlowIpv6SegmentRoutingRoutingType interface {
	// ToProto marshals PatternFlowIpv6SegmentRoutingRoutingType to protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingType
	ToProto() (*otg.PatternFlowIpv6SegmentRoutingRoutingType, error)
	// ToPbText marshals PatternFlowIpv6SegmentRoutingRoutingType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SegmentRoutingRoutingType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SegmentRoutingRoutingType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SegmentRoutingRoutingType struct {
	obj *patternFlowIpv6SegmentRoutingRoutingType
}

type unMarshalPatternFlowIpv6SegmentRoutingRoutingType interface {
	// FromProto unmarshals PatternFlowIpv6SegmentRoutingRoutingType from protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingType
	FromProto(msg *otg.PatternFlowIpv6SegmentRoutingRoutingType) (PatternFlowIpv6SegmentRoutingRoutingType, error)
	// FromPbText unmarshals PatternFlowIpv6SegmentRoutingRoutingType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SegmentRoutingRoutingType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SegmentRoutingRoutingType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) Marshal() marshalPatternFlowIpv6SegmentRoutingRoutingType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SegmentRoutingRoutingType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingRoutingType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SegmentRoutingRoutingType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingType) ToProto() (*otg.PatternFlowIpv6SegmentRoutingRoutingType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingType) FromProto(msg *otg.PatternFlowIpv6SegmentRoutingRoutingType) (PatternFlowIpv6SegmentRoutingRoutingType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SegmentRoutingRoutingType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SegmentRoutingRoutingType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SegmentRoutingRoutingType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) Clone() (PatternFlowIpv6SegmentRoutingRoutingType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SegmentRoutingRoutingType()
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

func (obj *patternFlowIpv6SegmentRoutingRoutingType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SegmentRoutingRoutingType is 8-bit Routing Type field in the SRH (RFC 8754 Section 2.1). The only defined value for Segment Routing is 4. When auto is assigned the implementation sets this to 4 automatically.
type PatternFlowIpv6SegmentRoutingRoutingType interface {
	Validation
	// msg marshals PatternFlowIpv6SegmentRoutingRoutingType to protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SegmentRoutingRoutingType
	// setMsg unmarshals PatternFlowIpv6SegmentRoutingRoutingType from protobuf object *otg.PatternFlowIpv6SegmentRoutingRoutingType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SegmentRoutingRoutingType) PatternFlowIpv6SegmentRoutingRoutingType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SegmentRoutingRoutingType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SegmentRoutingRoutingType
	// validate validates PatternFlowIpv6SegmentRoutingRoutingType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SegmentRoutingRoutingType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum, set in PatternFlowIpv6SegmentRoutingRoutingType
	Choice() PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum provided by user to PatternFlowIpv6SegmentRoutingRoutingType
	setChoice(value PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum) PatternFlowIpv6SegmentRoutingRoutingType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SegmentRoutingRoutingType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SegmentRoutingRoutingType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SegmentRoutingRoutingType
	SetValue(value uint32) PatternFlowIpv6SegmentRoutingRoutingType
	// HasValue checks if Value has been set in PatternFlowIpv6SegmentRoutingRoutingType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SegmentRoutingRoutingType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SegmentRoutingRoutingType
	SetValues(value []uint32) PatternFlowIpv6SegmentRoutingRoutingType
	// Auto returns uint32, set in PatternFlowIpv6SegmentRoutingRoutingType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SegmentRoutingRoutingType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SegmentRoutingRoutingTypeCounter, set in PatternFlowIpv6SegmentRoutingRoutingType.
	// PatternFlowIpv6SegmentRoutingRoutingTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// SetIncrement assigns PatternFlowIpv6SegmentRoutingRoutingTypeCounter provided by user to PatternFlowIpv6SegmentRoutingRoutingType.
	// PatternFlowIpv6SegmentRoutingRoutingTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SegmentRoutingRoutingType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SegmentRoutingRoutingTypeCounter, set in PatternFlowIpv6SegmentRoutingRoutingType.
	// PatternFlowIpv6SegmentRoutingRoutingTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SegmentRoutingRoutingTypeCounter
	// SetDecrement assigns PatternFlowIpv6SegmentRoutingRoutingTypeCounter provided by user to PatternFlowIpv6SegmentRoutingRoutingType.
	// PatternFlowIpv6SegmentRoutingRoutingTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SegmentRoutingRoutingType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SegmentRoutingRoutingType
var PatternFlowIpv6SegmentRoutingRoutingTypeChoice = struct {
	VALUE     PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
	VALUES    PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
	AUTO      PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
	INCREMENT PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
	DECREMENT PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) Choice() PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum {
	return PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SegmentRoutingRoutingType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) setChoice(value PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum) PatternFlowIpv6SegmentRoutingRoutingType {
	intValue, ok := otg.PatternFlowIpv6SegmentRoutingRoutingType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SegmentRoutingRoutingType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUES {
		defaultValue := []uint32{4}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingRoutingTypeChoice.AUTO {
		defaultValue := uint32(4)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SegmentRoutingRoutingTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SegmentRoutingRoutingTypeCounter().msg()
	}

	if value == PatternFlowIpv6SegmentRoutingRoutingTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SegmentRoutingRoutingTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SegmentRoutingRoutingType object
func (obj *patternFlowIpv6SegmentRoutingRoutingType) SetValue(value uint32) PatternFlowIpv6SegmentRoutingRoutingType {
	obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SegmentRoutingRoutingType object
func (obj *patternFlowIpv6SegmentRoutingRoutingType) SetValues(value []uint32) PatternFlowIpv6SegmentRoutingRoutingType {
	obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SegmentRoutingRoutingType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SegmentRoutingRoutingType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingRoutingTypeCounter
func (obj *patternFlowIpv6SegmentRoutingRoutingType) Increment() PatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SegmentRoutingRoutingTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SegmentRoutingRoutingTypeCounter
func (obj *patternFlowIpv6SegmentRoutingRoutingType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SegmentRoutingRoutingTypeCounter value in the PatternFlowIpv6SegmentRoutingRoutingType object
func (obj *patternFlowIpv6SegmentRoutingRoutingType) SetIncrement(value PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingType {
	obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingRoutingTypeCounter
func (obj *patternFlowIpv6SegmentRoutingRoutingType) Decrement() PatternFlowIpv6SegmentRoutingRoutingTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SegmentRoutingRoutingTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SegmentRoutingRoutingTypeCounter
func (obj *patternFlowIpv6SegmentRoutingRoutingType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SegmentRoutingRoutingTypeCounter value in the PatternFlowIpv6SegmentRoutingRoutingType object
func (obj *patternFlowIpv6SegmentRoutingRoutingType) SetDecrement(value PatternFlowIpv6SegmentRoutingRoutingTypeCounter) PatternFlowIpv6SegmentRoutingRoutingType {
	obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingRoutingType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SegmentRoutingRoutingType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SegmentRoutingRoutingType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SegmentRoutingRoutingType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SegmentRoutingRoutingTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingRoutingTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingRoutingTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingRoutingTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SegmentRoutingRoutingTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SegmentRoutingRoutingTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SegmentRoutingRoutingType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SegmentRoutingRoutingType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SegmentRoutingRoutingType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
