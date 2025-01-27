package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit *****
type patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	incrementHolder PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj := patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit{obj: &otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) msg() *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit struct {
	obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
}

type marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit struct {
	obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Clone() (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit()
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

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit is the minimum policed unit parameter should generally be set equal to the size of the smallest packet generated by the application.
type PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// validate validates PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	Choice() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	setChoice(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// Increment returns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter, set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter provided by user to PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit.
	// PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit
var PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Choice() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) setChoice(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Increment() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) Decrement() PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter value in the PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit object
func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitCounter) PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnitChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServMinimumPolicedUnit_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
