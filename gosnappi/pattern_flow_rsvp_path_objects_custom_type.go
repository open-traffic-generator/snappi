package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathObjectsCustomType *****
type patternFlowRSVPPathObjectsCustomType struct {
	validation
	obj             *otg.PatternFlowRSVPPathObjectsCustomType
	marshaller      marshalPatternFlowRSVPPathObjectsCustomType
	unMarshaller    unMarshalPatternFlowRSVPPathObjectsCustomType
	incrementHolder PatternFlowRSVPPathObjectsCustomTypeCounter
	decrementHolder PatternFlowRSVPPathObjectsCustomTypeCounter
}

func NewPatternFlowRSVPPathObjectsCustomType() PatternFlowRSVPPathObjectsCustomType {
	obj := patternFlowRSVPPathObjectsCustomType{obj: &otg.PatternFlowRSVPPathObjectsCustomType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathObjectsCustomType) msg() *otg.PatternFlowRSVPPathObjectsCustomType {
	return obj.obj
}

func (obj *patternFlowRSVPPathObjectsCustomType) setMsg(msg *otg.PatternFlowRSVPPathObjectsCustomType) PatternFlowRSVPPathObjectsCustomType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathObjectsCustomType struct {
	obj *patternFlowRSVPPathObjectsCustomType
}

type marshalPatternFlowRSVPPathObjectsCustomType interface {
	// ToProto marshals PatternFlowRSVPPathObjectsCustomType to protobuf object *otg.PatternFlowRSVPPathObjectsCustomType
	ToProto() (*otg.PatternFlowRSVPPathObjectsCustomType, error)
	// ToPbText marshals PatternFlowRSVPPathObjectsCustomType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathObjectsCustomType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathObjectsCustomType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowRSVPPathObjectsCustomType struct {
	obj *patternFlowRSVPPathObjectsCustomType
}

type unMarshalPatternFlowRSVPPathObjectsCustomType interface {
	// FromProto unmarshals PatternFlowRSVPPathObjectsCustomType from protobuf object *otg.PatternFlowRSVPPathObjectsCustomType
	FromProto(msg *otg.PatternFlowRSVPPathObjectsCustomType) (PatternFlowRSVPPathObjectsCustomType, error)
	// FromPbText unmarshals PatternFlowRSVPPathObjectsCustomType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathObjectsCustomType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathObjectsCustomType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathObjectsCustomType) Marshal() marshalPatternFlowRSVPPathObjectsCustomType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathObjectsCustomType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathObjectsCustomType) Unmarshal() unMarshalPatternFlowRSVPPathObjectsCustomType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathObjectsCustomType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathObjectsCustomType) ToProto() (*otg.PatternFlowRSVPPathObjectsCustomType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathObjectsCustomType) FromProto(msg *otg.PatternFlowRSVPPathObjectsCustomType) (PatternFlowRSVPPathObjectsCustomType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathObjectsCustomType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomType) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathObjectsCustomType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomType) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathObjectsCustomType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathObjectsCustomType) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathObjectsCustomType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathObjectsCustomType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathObjectsCustomType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathObjectsCustomType) Clone() (PatternFlowRSVPPathObjectsCustomType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathObjectsCustomType()
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

func (obj *patternFlowRSVPPathObjectsCustomType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathObjectsCustomType is user defined object type.
type PatternFlowRSVPPathObjectsCustomType interface {
	Validation
	// msg marshals PatternFlowRSVPPathObjectsCustomType to protobuf object *otg.PatternFlowRSVPPathObjectsCustomType
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathObjectsCustomType
	// setMsg unmarshals PatternFlowRSVPPathObjectsCustomType from protobuf object *otg.PatternFlowRSVPPathObjectsCustomType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathObjectsCustomType) PatternFlowRSVPPathObjectsCustomType
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathObjectsCustomType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathObjectsCustomType
	// validate validates PatternFlowRSVPPathObjectsCustomType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathObjectsCustomType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathObjectsCustomTypeChoiceEnum, set in PatternFlowRSVPPathObjectsCustomType
	Choice() PatternFlowRSVPPathObjectsCustomTypeChoiceEnum
	// setChoice assigns PatternFlowRSVPPathObjectsCustomTypeChoiceEnum provided by user to PatternFlowRSVPPathObjectsCustomType
	setChoice(value PatternFlowRSVPPathObjectsCustomTypeChoiceEnum) PatternFlowRSVPPathObjectsCustomType
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathObjectsCustomType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathObjectsCustomType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathObjectsCustomType
	SetValue(value uint32) PatternFlowRSVPPathObjectsCustomType
	// HasValue checks if Value has been set in PatternFlowRSVPPathObjectsCustomType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathObjectsCustomType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathObjectsCustomType
	SetValues(value []uint32) PatternFlowRSVPPathObjectsCustomType
	// Increment returns PatternFlowRSVPPathObjectsCustomTypeCounter, set in PatternFlowRSVPPathObjectsCustomType.
	// PatternFlowRSVPPathObjectsCustomTypeCounter is integer counter pattern
	Increment() PatternFlowRSVPPathObjectsCustomTypeCounter
	// SetIncrement assigns PatternFlowRSVPPathObjectsCustomTypeCounter provided by user to PatternFlowRSVPPathObjectsCustomType.
	// PatternFlowRSVPPathObjectsCustomTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomType
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathObjectsCustomType
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathObjectsCustomTypeCounter, set in PatternFlowRSVPPathObjectsCustomType.
	// PatternFlowRSVPPathObjectsCustomTypeCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathObjectsCustomTypeCounter
	// SetDecrement assigns PatternFlowRSVPPathObjectsCustomTypeCounter provided by user to PatternFlowRSVPPathObjectsCustomType.
	// PatternFlowRSVPPathObjectsCustomTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomType
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathObjectsCustomType
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathObjectsCustomTypeChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathObjectsCustomType
var PatternFlowRSVPPathObjectsCustomTypeChoice = struct {
	VALUE     PatternFlowRSVPPathObjectsCustomTypeChoiceEnum
	VALUES    PatternFlowRSVPPathObjectsCustomTypeChoiceEnum
	INCREMENT PatternFlowRSVPPathObjectsCustomTypeChoiceEnum
	DECREMENT PatternFlowRSVPPathObjectsCustomTypeChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathObjectsCustomTypeChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathObjectsCustomTypeChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathObjectsCustomTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathObjectsCustomTypeChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathObjectsCustomType) Choice() PatternFlowRSVPPathObjectsCustomTypeChoiceEnum {
	return PatternFlowRSVPPathObjectsCustomTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathObjectsCustomType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathObjectsCustomType) setChoice(value PatternFlowRSVPPathObjectsCustomTypeChoiceEnum) PatternFlowRSVPPathObjectsCustomType {
	intValue, ok := otg.PatternFlowRSVPPathObjectsCustomType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathObjectsCustomTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathObjectsCustomType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathObjectsCustomTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathObjectsCustomTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathObjectsCustomTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathObjectsCustomTypeCounter().msg()
	}

	if value == PatternFlowRSVPPathObjectsCustomTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathObjectsCustomTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathObjectsCustomType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathObjectsCustomType object
func (obj *patternFlowRSVPPathObjectsCustomType) SetValue(value uint32) PatternFlowRSVPPathObjectsCustomType {
	obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathObjectsCustomType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathObjectsCustomType object
func (obj *patternFlowRSVPPathObjectsCustomType) SetValues(value []uint32) PatternFlowRSVPPathObjectsCustomType {
	obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathObjectsCustomTypeCounter
func (obj *patternFlowRSVPPathObjectsCustomType) Increment() PatternFlowRSVPPathObjectsCustomTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathObjectsCustomTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathObjectsCustomTypeCounter
func (obj *patternFlowRSVPPathObjectsCustomType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathObjectsCustomTypeCounter value in the PatternFlowRSVPPathObjectsCustomType object
func (obj *patternFlowRSVPPathObjectsCustomType) SetIncrement(value PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomType {
	obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathObjectsCustomTypeCounter
func (obj *patternFlowRSVPPathObjectsCustomType) Decrement() PatternFlowRSVPPathObjectsCustomTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathObjectsCustomTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathObjectsCustomTypeCounter
func (obj *patternFlowRSVPPathObjectsCustomType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathObjectsCustomTypeCounter value in the PatternFlowRSVPPathObjectsCustomType object
func (obj *patternFlowRSVPPathObjectsCustomType) SetDecrement(value PatternFlowRSVPPathObjectsCustomTypeCounter) PatternFlowRSVPPathObjectsCustomType {
	obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathObjectsCustomType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathObjectsCustomType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathObjectsCustomType.Values <= 255 but Got %d", item))
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

func (obj *patternFlowRSVPPathObjectsCustomType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowRSVPPathObjectsCustomTypeChoice.VALUE)

	}

}
