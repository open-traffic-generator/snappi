package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvType *****
type patternFlowIpv6SRHPathTraceTlvType struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvType
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvType
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvType
	incrementHolder PatternFlowIpv6SRHPathTraceTlvTypeCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvTypeCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvType() PatternFlowIpv6SRHPathTraceTlvType {
	obj := patternFlowIpv6SRHPathTraceTlvType{obj: &otg.PatternFlowIpv6SRHPathTraceTlvType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) msg() *otg.PatternFlowIpv6SRHPathTraceTlvType {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvType) PatternFlowIpv6SRHPathTraceTlvType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvType struct {
	obj *patternFlowIpv6SRHPathTraceTlvType
}

type marshalPatternFlowIpv6SRHPathTraceTlvType interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvType to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvType
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvType, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvType struct {
	obj *patternFlowIpv6SRHPathTraceTlvType
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvType interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvType from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvType
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvType) (PatternFlowIpv6SRHPathTraceTlvType, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvType) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvType) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvType) (PatternFlowIpv6SRHPathTraceTlvType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvType) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) Clone() (PatternFlowIpv6SRHPathTraceTlvType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvType()
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

func (obj *patternFlowIpv6SRHPathTraceTlvType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvType is tLV type code (draft-ietf-spring-srv6-path-tracing). The canonical value is 128. Set to a different value only for negative or conformance testing.
type PatternFlowIpv6SRHPathTraceTlvType interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvType to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvType
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvType
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvType from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvType) PatternFlowIpv6SRHPathTraceTlvType
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvType
	// validate validates PatternFlowIpv6SRHPathTraceTlvType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvType
	Choice() PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvType
	setChoice(value PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum) PatternFlowIpv6SRHPathTraceTlvType
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvType
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvType
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvType
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvType
	// Auto returns uint32, set in PatternFlowIpv6SRHPathTraceTlvType.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6SRHPathTraceTlvType
	HasAuto() bool
	// Increment returns PatternFlowIpv6SRHPathTraceTlvTypeCounter, set in PatternFlowIpv6SRHPathTraceTlvType.
	// PatternFlowIpv6SRHPathTraceTlvTypeCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvTypeCounter provided by user to PatternFlowIpv6SRHPathTraceTlvType.
	// PatternFlowIpv6SRHPathTraceTlvTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvType
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvType
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvTypeCounter, set in PatternFlowIpv6SRHPathTraceTlvType.
	// PatternFlowIpv6SRHPathTraceTlvTypeCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvTypeCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvTypeCounter provided by user to PatternFlowIpv6SRHPathTraceTlvType.
	// PatternFlowIpv6SRHPathTraceTlvTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvType
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvType
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvType
var PatternFlowIpv6SRHPathTraceTlvTypeChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
	AUTO      PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum("values"),
	AUTO:      PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) Choice() PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) setChoice(value PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum) PatternFlowIpv6SRHPathTraceTlvType {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUE {
		defaultValue := uint32(128)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUES {
		defaultValue := []uint32{128}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTypeChoice.AUTO {
		defaultValue := uint32(128)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvTypeCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvType object
func (obj *patternFlowIpv6SRHPathTraceTlvType) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvType {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{128})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvType object
func (obj *patternFlowIpv6SRHPathTraceTlvType) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvType {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUES)
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
func (obj *patternFlowIpv6SRHPathTraceTlvType) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvType) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvTypeCounter
func (obj *patternFlowIpv6SRHPathTraceTlvType) Increment() PatternFlowIpv6SRHPathTraceTlvTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvTypeCounter
func (obj *patternFlowIpv6SRHPathTraceTlvType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvTypeCounter value in the PatternFlowIpv6SRHPathTraceTlvType object
func (obj *patternFlowIpv6SRHPathTraceTlvType) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvType {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvTypeCounter
func (obj *patternFlowIpv6SRHPathTraceTlvType) Decrement() PatternFlowIpv6SRHPathTraceTlvTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvTypeCounter
func (obj *patternFlowIpv6SRHPathTraceTlvType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvTypeCounter value in the PatternFlowIpv6SRHPathTraceTlvType object
func (obj *patternFlowIpv6SRHPathTraceTlvType) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvTypeCounter) PatternFlowIpv6SRHPathTraceTlvType {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6SRHPathTraceTlvType.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6SRHPathTraceTlvType.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTypeChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTypeChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvTypeChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvType")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
