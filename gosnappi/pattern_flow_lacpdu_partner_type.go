package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerType *****
type patternFlowLacpduPartnerType struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerType
	marshaller      marshalPatternFlowLacpduPartnerType
	unMarshaller    unMarshalPatternFlowLacpduPartnerType
	incrementHolder PatternFlowLacpduPartnerTypeCounter
	decrementHolder PatternFlowLacpduPartnerTypeCounter
}

func NewPatternFlowLacpduPartnerType() PatternFlowLacpduPartnerType {
	obj := patternFlowLacpduPartnerType{obj: &otg.PatternFlowLacpduPartnerType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerType) msg() *otg.PatternFlowLacpduPartnerType {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerType) setMsg(msg *otg.PatternFlowLacpduPartnerType) PatternFlowLacpduPartnerType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerType struct {
	obj *patternFlowLacpduPartnerType
}

type marshalPatternFlowLacpduPartnerType interface {
	// ToProto marshals PatternFlowLacpduPartnerType to protobuf object *otg.PatternFlowLacpduPartnerType
	ToProto() (*otg.PatternFlowLacpduPartnerType, error)
	// ToPbText marshals PatternFlowLacpduPartnerType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerType struct {
	obj *patternFlowLacpduPartnerType
}

type unMarshalPatternFlowLacpduPartnerType interface {
	// FromProto unmarshals PatternFlowLacpduPartnerType from protobuf object *otg.PatternFlowLacpduPartnerType
	FromProto(msg *otg.PatternFlowLacpduPartnerType) (PatternFlowLacpduPartnerType, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerType) Marshal() marshalPatternFlowLacpduPartnerType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerType) Unmarshal() unMarshalPatternFlowLacpduPartnerType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerType) ToProto() (*otg.PatternFlowLacpduPartnerType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerType) FromProto(msg *otg.PatternFlowLacpduPartnerType) (PatternFlowLacpduPartnerType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerType) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerType) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerType) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerType) Clone() (PatternFlowLacpduPartnerType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerType()
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

func (obj *patternFlowLacpduPartnerType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerType is tLV Type for Partner Information. The value 0x02 identifies this TLV as containing information about the remote device (the Partner), as understood by the Actor.
type PatternFlowLacpduPartnerType interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerType to protobuf object *otg.PatternFlowLacpduPartnerType
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerType
	// setMsg unmarshals PatternFlowLacpduPartnerType from protobuf object *otg.PatternFlowLacpduPartnerType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerType) PatternFlowLacpduPartnerType
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerType
	// validate validates PatternFlowLacpduPartnerType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerTypeChoiceEnum, set in PatternFlowLacpduPartnerType
	Choice() PatternFlowLacpduPartnerTypeChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerTypeChoiceEnum provided by user to PatternFlowLacpduPartnerType
	setChoice(value PatternFlowLacpduPartnerTypeChoiceEnum) PatternFlowLacpduPartnerType
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerType
	SetValue(value uint32) PatternFlowLacpduPartnerType
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerType
	SetValues(value []uint32) PatternFlowLacpduPartnerType
	// Increment returns PatternFlowLacpduPartnerTypeCounter, set in PatternFlowLacpduPartnerType.
	// PatternFlowLacpduPartnerTypeCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerTypeCounter
	// SetIncrement assigns PatternFlowLacpduPartnerTypeCounter provided by user to PatternFlowLacpduPartnerType.
	// PatternFlowLacpduPartnerTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerType
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerType
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerTypeCounter, set in PatternFlowLacpduPartnerType.
	// PatternFlowLacpduPartnerTypeCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerTypeCounter
	// SetDecrement assigns PatternFlowLacpduPartnerTypeCounter provided by user to PatternFlowLacpduPartnerType.
	// PatternFlowLacpduPartnerTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerType
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerType
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerTypeChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerType
var PatternFlowLacpduPartnerTypeChoice = struct {
	VALUE     PatternFlowLacpduPartnerTypeChoiceEnum
	VALUES    PatternFlowLacpduPartnerTypeChoiceEnum
	INCREMENT PatternFlowLacpduPartnerTypeChoiceEnum
	DECREMENT PatternFlowLacpduPartnerTypeChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerTypeChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerTypeChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerTypeChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerType) Choice() PatternFlowLacpduPartnerTypeChoiceEnum {
	return PatternFlowLacpduPartnerTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerType) setChoice(value PatternFlowLacpduPartnerTypeChoiceEnum) PatternFlowLacpduPartnerType {
	intValue, ok := otg.PatternFlowLacpduPartnerType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerTypeChoice.VALUE {
		defaultValue := uint32(2)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerTypeChoice.VALUES {
		defaultValue := []uint32{2}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerTypeCounter().msg()
	}

	if value == PatternFlowLacpduPartnerTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerType object
func (obj *patternFlowLacpduPartnerType) SetValue(value uint32) PatternFlowLacpduPartnerType {
	obj.setChoice(PatternFlowLacpduPartnerTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerType object
func (obj *patternFlowLacpduPartnerType) SetValues(value []uint32) PatternFlowLacpduPartnerType {
	obj.setChoice(PatternFlowLacpduPartnerTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerTypeCounter
func (obj *patternFlowLacpduPartnerType) Increment() PatternFlowLacpduPartnerTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerTypeCounter
func (obj *patternFlowLacpduPartnerType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerTypeCounter value in the PatternFlowLacpduPartnerType object
func (obj *patternFlowLacpduPartnerType) SetIncrement(value PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerType {
	obj.setChoice(PatternFlowLacpduPartnerTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerTypeCounter
func (obj *patternFlowLacpduPartnerType) Decrement() PatternFlowLacpduPartnerTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerTypeCounter
func (obj *patternFlowLacpduPartnerType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerTypeCounter value in the PatternFlowLacpduPartnerType object
func (obj *patternFlowLacpduPartnerType) SetDecrement(value PatternFlowLacpduPartnerTypeCounter) PatternFlowLacpduPartnerType {
	obj.setChoice(PatternFlowLacpduPartnerTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerType")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
