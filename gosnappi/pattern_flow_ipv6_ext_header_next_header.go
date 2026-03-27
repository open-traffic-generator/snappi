package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6ExtHeaderNextHeader *****
type patternFlowIpv6ExtHeaderNextHeader struct {
	validation
	obj             *otg.PatternFlowIpv6ExtHeaderNextHeader
	marshaller      marshalPatternFlowIpv6ExtHeaderNextHeader
	unMarshaller    unMarshalPatternFlowIpv6ExtHeaderNextHeader
	incrementHolder PatternFlowIpv6ExtHeaderNextHeaderCounter
	decrementHolder PatternFlowIpv6ExtHeaderNextHeaderCounter
}

func NewPatternFlowIpv6ExtHeaderNextHeader() PatternFlowIpv6ExtHeaderNextHeader {
	obj := patternFlowIpv6ExtHeaderNextHeader{obj: &otg.PatternFlowIpv6ExtHeaderNextHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) msg() *otg.PatternFlowIpv6ExtHeaderNextHeader {
	return obj.obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) setMsg(msg *otg.PatternFlowIpv6ExtHeaderNextHeader) PatternFlowIpv6ExtHeaderNextHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6ExtHeaderNextHeader struct {
	obj *patternFlowIpv6ExtHeaderNextHeader
}

type marshalPatternFlowIpv6ExtHeaderNextHeader interface {
	// ToProto marshals PatternFlowIpv6ExtHeaderNextHeader to protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeader
	ToProto() (*otg.PatternFlowIpv6ExtHeaderNextHeader, error)
	// ToPbText marshals PatternFlowIpv6ExtHeaderNextHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6ExtHeaderNextHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6ExtHeaderNextHeader to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6ExtHeaderNextHeader struct {
	obj *patternFlowIpv6ExtHeaderNextHeader
}

type unMarshalPatternFlowIpv6ExtHeaderNextHeader interface {
	// FromProto unmarshals PatternFlowIpv6ExtHeaderNextHeader from protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeader
	FromProto(msg *otg.PatternFlowIpv6ExtHeaderNextHeader) (PatternFlowIpv6ExtHeaderNextHeader, error)
	// FromPbText unmarshals PatternFlowIpv6ExtHeaderNextHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6ExtHeaderNextHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6ExtHeaderNextHeader from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) Marshal() marshalPatternFlowIpv6ExtHeaderNextHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6ExtHeaderNextHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) Unmarshal() unMarshalPatternFlowIpv6ExtHeaderNextHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6ExtHeaderNextHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6ExtHeaderNextHeader) ToProto() (*otg.PatternFlowIpv6ExtHeaderNextHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeader) FromProto(msg *otg.PatternFlowIpv6ExtHeaderNextHeader) (PatternFlowIpv6ExtHeaderNextHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6ExtHeaderNextHeader) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeader) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderNextHeader) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeader) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6ExtHeaderNextHeader) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6ExtHeaderNextHeader) FromJson(value string) error {
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

func (obj *patternFlowIpv6ExtHeaderNextHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) Clone() (PatternFlowIpv6ExtHeaderNextHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6ExtHeaderNextHeader()
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

func (obj *patternFlowIpv6ExtHeaderNextHeader) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6ExtHeaderNextHeader is the next header that identifies the type of header immediately following this IPv6 Extension header. For TCP and UDP the values are 6 and 17 respectively. A value of 59 indicates that there are no further IPv6 Extension headers and that the payload will be appended after this header. The available defined values are listed in https://www.iana.org/assignments/ipv6-parameters/ipv6-parameters.xhtml#extension-header
type PatternFlowIpv6ExtHeaderNextHeader interface {
	Validation
	// msg marshals PatternFlowIpv6ExtHeaderNextHeader to protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeader
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6ExtHeaderNextHeader
	// setMsg unmarshals PatternFlowIpv6ExtHeaderNextHeader from protobuf object *otg.PatternFlowIpv6ExtHeaderNextHeader
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6ExtHeaderNextHeader) PatternFlowIpv6ExtHeaderNextHeader
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6ExtHeaderNextHeader
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6ExtHeaderNextHeader
	// validate validates PatternFlowIpv6ExtHeaderNextHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6ExtHeaderNextHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum, set in PatternFlowIpv6ExtHeaderNextHeader
	Choice() PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
	// setChoice assigns PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum provided by user to PatternFlowIpv6ExtHeaderNextHeader
	setChoice(value PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum) PatternFlowIpv6ExtHeaderNextHeader
	// HasChoice checks if Choice has been set in PatternFlowIpv6ExtHeaderNextHeader
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6ExtHeaderNextHeader.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6ExtHeaderNextHeader
	SetValue(value uint32) PatternFlowIpv6ExtHeaderNextHeader
	// HasValue checks if Value has been set in PatternFlowIpv6ExtHeaderNextHeader
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6ExtHeaderNextHeader.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6ExtHeaderNextHeader
	SetValues(value []uint32) PatternFlowIpv6ExtHeaderNextHeader
	// Auto returns uint32, set in PatternFlowIpv6ExtHeaderNextHeader.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6ExtHeaderNextHeader
	HasAuto() bool
	// Increment returns PatternFlowIpv6ExtHeaderNextHeaderCounter, set in PatternFlowIpv6ExtHeaderNextHeader.
	// PatternFlowIpv6ExtHeaderNextHeaderCounter is integer counter pattern
	Increment() PatternFlowIpv6ExtHeaderNextHeaderCounter
	// SetIncrement assigns PatternFlowIpv6ExtHeaderNextHeaderCounter provided by user to PatternFlowIpv6ExtHeaderNextHeader.
	// PatternFlowIpv6ExtHeaderNextHeaderCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeader
	// HasIncrement checks if Increment has been set in PatternFlowIpv6ExtHeaderNextHeader
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6ExtHeaderNextHeaderCounter, set in PatternFlowIpv6ExtHeaderNextHeader.
	// PatternFlowIpv6ExtHeaderNextHeaderCounter is integer counter pattern
	Decrement() PatternFlowIpv6ExtHeaderNextHeaderCounter
	// SetDecrement assigns PatternFlowIpv6ExtHeaderNextHeaderCounter provided by user to PatternFlowIpv6ExtHeaderNextHeader.
	// PatternFlowIpv6ExtHeaderNextHeaderCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeader
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6ExtHeaderNextHeader
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum string

// Enum of Choice on PatternFlowIpv6ExtHeaderNextHeader
var PatternFlowIpv6ExtHeaderNextHeaderChoice = struct {
	VALUE     PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
	VALUES    PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
	AUTO      PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
	INCREMENT PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
	DECREMENT PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum
}{
	VALUE:     PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum("value"),
	VALUES:    PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum("values"),
	AUTO:      PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) Choice() PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum {
	return PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6ExtHeaderNextHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) setChoice(value PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum) PatternFlowIpv6ExtHeaderNextHeader {
	intValue, ok := otg.PatternFlowIpv6ExtHeaderNextHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6ExtHeaderNextHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUE {
		defaultValue := uint32(59)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUES {
		defaultValue := []uint32{59}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderNextHeaderChoice.AUTO {
		defaultValue := uint32(59)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6ExtHeaderNextHeaderChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6ExtHeaderNextHeaderCounter().msg()
	}

	if value == PatternFlowIpv6ExtHeaderNextHeaderChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6ExtHeaderNextHeaderCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeader) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeader) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6ExtHeaderNextHeader object
func (obj *patternFlowIpv6ExtHeaderNextHeader) SetValue(value uint32) PatternFlowIpv6ExtHeaderNextHeader {
	obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6ExtHeaderNextHeader) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{59})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6ExtHeaderNextHeader object
func (obj *patternFlowIpv6ExtHeaderNextHeader) SetValues(value []uint32) PatternFlowIpv6ExtHeaderNextHeader {
	obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUES)
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
func (obj *patternFlowIpv6ExtHeaderNextHeader) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6ExtHeaderNextHeader) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6ExtHeaderNextHeaderCounter
func (obj *patternFlowIpv6ExtHeaderNextHeader) Increment() PatternFlowIpv6ExtHeaderNextHeaderCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6ExtHeaderNextHeaderCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6ExtHeaderNextHeaderCounter
func (obj *patternFlowIpv6ExtHeaderNextHeader) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6ExtHeaderNextHeaderCounter value in the PatternFlowIpv6ExtHeaderNextHeader object
func (obj *patternFlowIpv6ExtHeaderNextHeader) SetIncrement(value PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeader {
	obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6ExtHeaderNextHeaderCounter
func (obj *patternFlowIpv6ExtHeaderNextHeader) Decrement() PatternFlowIpv6ExtHeaderNextHeaderCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6ExtHeaderNextHeaderCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6ExtHeaderNextHeaderCounter
func (obj *patternFlowIpv6ExtHeaderNextHeader) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6ExtHeaderNextHeaderCounter value in the PatternFlowIpv6ExtHeaderNextHeader object
func (obj *patternFlowIpv6ExtHeaderNextHeader) SetDecrement(value PatternFlowIpv6ExtHeaderNextHeaderCounter) PatternFlowIpv6ExtHeaderNextHeader {
	obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6ExtHeaderNextHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderNextHeader.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6ExtHeaderNextHeader.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6ExtHeaderNextHeader.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6ExtHeaderNextHeader) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6ExtHeaderNextHeaderChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderNextHeaderChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderNextHeaderChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderNextHeaderChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6ExtHeaderNextHeaderChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6ExtHeaderNextHeaderChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6ExtHeaderNextHeader")
			}
		} else {
			intVal := otg.PatternFlowIpv6ExtHeaderNextHeader_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6ExtHeaderNextHeader_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
