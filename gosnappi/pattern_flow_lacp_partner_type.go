package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerType *****
type patternFlowLacpPartnerType struct {
	validation
	obj             *otg.PatternFlowLacpPartnerType
	marshaller      marshalPatternFlowLacpPartnerType
	unMarshaller    unMarshalPatternFlowLacpPartnerType
	incrementHolder PatternFlowLacpPartnerTypeCounter
	decrementHolder PatternFlowLacpPartnerTypeCounter
}

func NewPatternFlowLacpPartnerType() PatternFlowLacpPartnerType {
	obj := patternFlowLacpPartnerType{obj: &otg.PatternFlowLacpPartnerType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerType) msg() *otg.PatternFlowLacpPartnerType {
	return obj.obj
}

func (obj *patternFlowLacpPartnerType) setMsg(msg *otg.PatternFlowLacpPartnerType) PatternFlowLacpPartnerType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerType struct {
	obj *patternFlowLacpPartnerType
}

type marshalPatternFlowLacpPartnerType interface {
	// ToProto marshals PatternFlowLacpPartnerType to protobuf object *otg.PatternFlowLacpPartnerType
	ToProto() (*otg.PatternFlowLacpPartnerType, error)
	// ToPbText marshals PatternFlowLacpPartnerType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerType struct {
	obj *patternFlowLacpPartnerType
}

type unMarshalPatternFlowLacpPartnerType interface {
	// FromProto unmarshals PatternFlowLacpPartnerType from protobuf object *otg.PatternFlowLacpPartnerType
	FromProto(msg *otg.PatternFlowLacpPartnerType) (PatternFlowLacpPartnerType, error)
	// FromPbText unmarshals PatternFlowLacpPartnerType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerType) Marshal() marshalPatternFlowLacpPartnerType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerType) Unmarshal() unMarshalPatternFlowLacpPartnerType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerType) ToProto() (*otg.PatternFlowLacpPartnerType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerType) FromProto(msg *otg.PatternFlowLacpPartnerType) (PatternFlowLacpPartnerType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerType) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerType) Clone() (PatternFlowLacpPartnerType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerType()
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

func (obj *patternFlowLacpPartnerType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
type PatternFlowLacpPartnerType interface {
	Validation
	// msg marshals PatternFlowLacpPartnerType to protobuf object *otg.PatternFlowLacpPartnerType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerType
	// setMsg unmarshals PatternFlowLacpPartnerType from protobuf object *otg.PatternFlowLacpPartnerType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerType) PatternFlowLacpPartnerType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerType
	// validate validates PatternFlowLacpPartnerType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerTypeChoiceEnum, set in PatternFlowLacpPartnerType
	Choice() PatternFlowLacpPartnerTypeChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerTypeChoiceEnum provided by user to PatternFlowLacpPartnerType
	setChoice(value PatternFlowLacpPartnerTypeChoiceEnum) PatternFlowLacpPartnerType
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerType
	SetValue(value uint32) PatternFlowLacpPartnerType
	// HasValue checks if Value has been set in PatternFlowLacpPartnerType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerType
	SetValues(value []uint32) PatternFlowLacpPartnerType
	// Increment returns PatternFlowLacpPartnerTypeCounter, set in PatternFlowLacpPartnerType.
	// PatternFlowLacpPartnerTypeCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerTypeCounter
	// SetIncrement assigns PatternFlowLacpPartnerTypeCounter provided by user to PatternFlowLacpPartnerType.
	// PatternFlowLacpPartnerTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerType
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerTypeCounter, set in PatternFlowLacpPartnerType.
	// PatternFlowLacpPartnerTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerTypeCounter
	// SetDecrement assigns PatternFlowLacpPartnerTypeCounter provided by user to PatternFlowLacpPartnerType.
	// PatternFlowLacpPartnerTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerType
var PatternFlowLacpPartnerTypeChoice = struct {
	VALUE     PatternFlowLacpPartnerTypeChoiceEnum
	VALUES    PatternFlowLacpPartnerTypeChoiceEnum
	INCREMENT PatternFlowLacpPartnerTypeChoiceEnum
	DECREMENT PatternFlowLacpPartnerTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerType) Choice() PatternFlowLacpPartnerTypeChoiceEnum {
	return PatternFlowLacpPartnerTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerType) setChoice(value PatternFlowLacpPartnerTypeChoiceEnum) PatternFlowLacpPartnerType {
	intValue, ok := otg.PatternFlowLacpPartnerType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerTypeChoice.VALUE {
		defaultValue := uint32(2)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerTypeChoice.VALUES {
		defaultValue := []uint32{2}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerTypeCounter().msg()
	}

	if value == PatternFlowLacpPartnerTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerType object
func (obj *patternFlowLacpPartnerType) SetValue(value uint32) PatternFlowLacpPartnerType {
	obj.setChoice(PatternFlowLacpPartnerTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerType object
func (obj *patternFlowLacpPartnerType) SetValues(value []uint32) PatternFlowLacpPartnerType {
	obj.setChoice(PatternFlowLacpPartnerTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerTypeCounter
func (obj *patternFlowLacpPartnerType) Increment() PatternFlowLacpPartnerTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerTypeCounter
func (obj *patternFlowLacpPartnerType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerTypeCounter value in the PatternFlowLacpPartnerType object
func (obj *patternFlowLacpPartnerType) SetIncrement(value PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerType {
	obj.setChoice(PatternFlowLacpPartnerTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerTypeCounter
func (obj *patternFlowLacpPartnerType) Decrement() PatternFlowLacpPartnerTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerTypeCounter
func (obj *patternFlowLacpPartnerType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerTypeCounter value in the PatternFlowLacpPartnerType object
func (obj *patternFlowLacpPartnerType) SetDecrement(value PatternFlowLacpPartnerTypeCounter) PatternFlowLacpPartnerType {
	obj.setChoice(PatternFlowLacpPartnerTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpPartnerType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerType")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
