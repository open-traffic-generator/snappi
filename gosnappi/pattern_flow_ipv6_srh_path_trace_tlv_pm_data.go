package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6SRHPathTraceTlvPmData *****
type patternFlowIpv6SRHPathTraceTlvPmData struct {
	validation
	obj             *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	marshaller      marshalPatternFlowIpv6SRHPathTraceTlvPmData
	unMarshaller    unMarshalPatternFlowIpv6SRHPathTraceTlvPmData
	incrementHolder PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	decrementHolder PatternFlowIpv6SRHPathTraceTlvPmDataCounter
}

func NewPatternFlowIpv6SRHPathTraceTlvPmData() PatternFlowIpv6SRHPathTraceTlvPmData {
	obj := patternFlowIpv6SRHPathTraceTlvPmData{obj: &otg.PatternFlowIpv6SRHPathTraceTlvPmData{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) msg() *otg.PatternFlowIpv6SRHPathTraceTlvPmData {
	return obj.obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) setMsg(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmData) PatternFlowIpv6SRHPathTraceTlvPmData {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6SRHPathTraceTlvPmData struct {
	obj *patternFlowIpv6SRHPathTraceTlvPmData
}

type marshalPatternFlowIpv6SRHPathTraceTlvPmData interface {
	// ToProto marshals PatternFlowIpv6SRHPathTraceTlvPmData to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvPmData, error)
	// ToPbText marshals PatternFlowIpv6SRHPathTraceTlvPmData to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6SRHPathTraceTlvPmData to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6SRHPathTraceTlvPmData to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6SRHPathTraceTlvPmData struct {
	obj *patternFlowIpv6SRHPathTraceTlvPmData
}

type unMarshalPatternFlowIpv6SRHPathTraceTlvPmData interface {
	// FromProto unmarshals PatternFlowIpv6SRHPathTraceTlvPmData from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmData) (PatternFlowIpv6SRHPathTraceTlvPmData, error)
	// FromPbText unmarshals PatternFlowIpv6SRHPathTraceTlvPmData from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6SRHPathTraceTlvPmData from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6SRHPathTraceTlvPmData from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Marshal() marshalPatternFlowIpv6SRHPathTraceTlvPmData {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6SRHPathTraceTlvPmData{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvPmData {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6SRHPathTraceTlvPmData{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmData) ToProto() (*otg.PatternFlowIpv6SRHPathTraceTlvPmData, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmData) FromProto(msg *otg.PatternFlowIpv6SRHPathTraceTlvPmData) (PatternFlowIpv6SRHPathTraceTlvPmData, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmData) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmData) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmData) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmData) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6SRHPathTraceTlvPmData) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6SRHPathTraceTlvPmData) FromJson(value string) error {
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

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Clone() (PatternFlowIpv6SRHPathTraceTlvPmData, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6SRHPathTraceTlvPmData()
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

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6SRHPathTraceTlvPmData is performance Monitoring (PM) data inserted by the ingress node (draft-ietf-spring-srv6-path-tracing). Carries interface and load information encoded as a 32-bit value.
type PatternFlowIpv6SRHPathTraceTlvPmData interface {
	Validation
	// msg marshals PatternFlowIpv6SRHPathTraceTlvPmData to protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	// setMsg unmarshals PatternFlowIpv6SRHPathTraceTlvPmData from protobuf object *otg.PatternFlowIpv6SRHPathTraceTlvPmData
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6SRHPathTraceTlvPmData) PatternFlowIpv6SRHPathTraceTlvPmData
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6SRHPathTraceTlvPmData
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6SRHPathTraceTlvPmData
	// validate validates PatternFlowIpv6SRHPathTraceTlvPmData
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6SRHPathTraceTlvPmData, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum, set in PatternFlowIpv6SRHPathTraceTlvPmData
	Choice() PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum
	// setChoice assigns PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum provided by user to PatternFlowIpv6SRHPathTraceTlvPmData
	setChoice(value PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum) PatternFlowIpv6SRHPathTraceTlvPmData
	// HasChoice checks if Choice has been set in PatternFlowIpv6SRHPathTraceTlvPmData
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6SRHPathTraceTlvPmData.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvPmData
	SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvPmData
	// HasValue checks if Value has been set in PatternFlowIpv6SRHPathTraceTlvPmData
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6SRHPathTraceTlvPmData.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6SRHPathTraceTlvPmData
	SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvPmData
	// Increment returns PatternFlowIpv6SRHPathTraceTlvPmDataCounter, set in PatternFlowIpv6SRHPathTraceTlvPmData.
	// PatternFlowIpv6SRHPathTraceTlvPmDataCounter is integer counter pattern
	Increment() PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// SetIncrement assigns PatternFlowIpv6SRHPathTraceTlvPmDataCounter provided by user to PatternFlowIpv6SRHPathTraceTlvPmData.
	// PatternFlowIpv6SRHPathTraceTlvPmDataCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmData
	// HasIncrement checks if Increment has been set in PatternFlowIpv6SRHPathTraceTlvPmData
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6SRHPathTraceTlvPmDataCounter, set in PatternFlowIpv6SRHPathTraceTlvPmData.
	// PatternFlowIpv6SRHPathTraceTlvPmDataCounter is integer counter pattern
	Decrement() PatternFlowIpv6SRHPathTraceTlvPmDataCounter
	// SetDecrement assigns PatternFlowIpv6SRHPathTraceTlvPmDataCounter provided by user to PatternFlowIpv6SRHPathTraceTlvPmData.
	// PatternFlowIpv6SRHPathTraceTlvPmDataCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmData
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6SRHPathTraceTlvPmData
	HasDecrement() bool
	setNil()
}

type PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum string

// Enum of Choice on PatternFlowIpv6SRHPathTraceTlvPmData
var PatternFlowIpv6SRHPathTraceTlvPmDataChoice = struct {
	VALUE     PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum
	VALUES    PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum
	INCREMENT PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum
	DECREMENT PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum
}{
	VALUE:     PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum("value"),
	VALUES:    PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Choice() PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum {
	return PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) setChoice(value PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum) PatternFlowIpv6SRHPathTraceTlvPmData {
	intValue, ok := otg.PatternFlowIpv6SRHPathTraceTlvPmData_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6SRHPathTraceTlvPmData_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6SRHPathTraceTlvPmDataChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6SRHPathTraceTlvPmDataCounter().msg()
	}

	if value == PatternFlowIpv6SRHPathTraceTlvPmDataChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6SRHPathTraceTlvPmDataCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6SRHPathTraceTlvPmData object
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) SetValue(value uint32) PatternFlowIpv6SRHPathTraceTlvPmData {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6SRHPathTraceTlvPmData object
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) SetValues(value []uint32) PatternFlowIpv6SRHPathTraceTlvPmData {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvPmDataCounter
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Increment() PatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6SRHPathTraceTlvPmDataCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6SRHPathTraceTlvPmDataCounter
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6SRHPathTraceTlvPmDataCounter value in the PatternFlowIpv6SRHPathTraceTlvPmData object
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) SetIncrement(value PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmData {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvPmDataCounter
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) Decrement() PatternFlowIpv6SRHPathTraceTlvPmDataCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6SRHPathTraceTlvPmDataCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6SRHPathTraceTlvPmDataCounter
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6SRHPathTraceTlvPmDataCounter value in the PatternFlowIpv6SRHPathTraceTlvPmData object
func (obj *patternFlowIpv6SRHPathTraceTlvPmData) SetDecrement(value PatternFlowIpv6SRHPathTraceTlvPmDataCounter) PatternFlowIpv6SRHPathTraceTlvPmData {
	obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIpv6SRHPathTraceTlvPmData) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6SRHPathTraceTlvPmDataChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvPmDataChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6SRHPathTraceTlvPmDataChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6SRHPathTraceTlvPmDataChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6SRHPathTraceTlvPmData")
			}
		} else {
			intVal := otg.PatternFlowIpv6SRHPathTraceTlvPmData_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6SRHPathTraceTlvPmData_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
