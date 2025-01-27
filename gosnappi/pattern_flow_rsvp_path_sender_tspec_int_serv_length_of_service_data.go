package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData *****
type patternFlowRSVPPathSenderTspecIntServLengthOfServiceData struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	incrementHolder PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj := patternFlowRSVPPathSenderTspecIntServLengthOfServiceData{obj: &otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) msg() *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData struct {
	obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData
}

type marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData struct {
	obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Clone() (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData()
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

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData is length of service data, 6 words not including per-service header.
type PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// validate validates PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	Choice() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	setChoice(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// Increment returns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter, set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter provided by user to PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.
	// PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData
var PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Choice() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) setChoice(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUE {
		defaultValue := uint32(6)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUES {
		defaultValue := []uint32{6}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{6})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Increment() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) Decrement() PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter value in the PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData object
func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataCounter) PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData.Values <= 65535 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServLengthOfServiceData) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServLengthOfServiceDataChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServLengthOfServiceData_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
