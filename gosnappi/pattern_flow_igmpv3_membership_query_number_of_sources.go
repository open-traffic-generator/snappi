package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3MembershipQueryNumberOfSources *****
type patternFlowIgmpv3MembershipQueryNumberOfSources struct {
	validation
	obj             *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	marshaller      marshalPatternFlowIgmpv3MembershipQueryNumberOfSources
	unMarshaller    unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSources
	incrementHolder PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	decrementHolder PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
}

func NewPatternFlowIgmpv3MembershipQueryNumberOfSources() PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj := patternFlowIgmpv3MembershipQueryNumberOfSources{obj: &otg.PatternFlowIgmpv3MembershipQueryNumberOfSources{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) msg() *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources {
	return obj.obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) setMsg(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3MembershipQueryNumberOfSources struct {
	obj *patternFlowIgmpv3MembershipQueryNumberOfSources
}

type marshalPatternFlowIgmpv3MembershipQueryNumberOfSources interface {
	// ToProto marshals PatternFlowIgmpv3MembershipQueryNumberOfSources to protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	ToProto() (*otg.PatternFlowIgmpv3MembershipQueryNumberOfSources, error)
	// ToPbText marshals PatternFlowIgmpv3MembershipQueryNumberOfSources to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3MembershipQueryNumberOfSources to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3MembershipQueryNumberOfSources to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources struct {
	obj *patternFlowIgmpv3MembershipQueryNumberOfSources
}

type unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSources interface {
	// FromProto unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSources from protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources) (PatternFlowIgmpv3MembershipQueryNumberOfSources, error)
	// FromPbText unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSources from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSources from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSources from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Marshal() marshalPatternFlowIgmpv3MembershipQueryNumberOfSources {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3MembershipQueryNumberOfSources{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSources {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSources) ToProto() (*otg.PatternFlowIgmpv3MembershipQueryNumberOfSources, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources) FromProto(msg *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources) (PatternFlowIgmpv3MembershipQueryNumberOfSources, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSources) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSources) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3MembershipQueryNumberOfSources) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3MembershipQueryNumberOfSources) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Clone() (PatternFlowIgmpv3MembershipQueryNumberOfSources, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3MembershipQueryNumberOfSources()
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

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3MembershipQueryNumberOfSources is number of source addresses in the query, zero for a general and a group-specific query. When auto, it is set to the number of entries in sources.
type PatternFlowIgmpv3MembershipQueryNumberOfSources interface {
	Validation
	// msg marshals PatternFlowIgmpv3MembershipQueryNumberOfSources to protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	// setMsg unmarshals PatternFlowIgmpv3MembershipQueryNumberOfSources from protobuf object *otg.PatternFlowIgmpv3MembershipQueryNumberOfSources
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3MembershipQueryNumberOfSources) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3MembershipQueryNumberOfSources
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3MembershipQueryNumberOfSources
	// validate validates PatternFlowIgmpv3MembershipQueryNumberOfSources
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3MembershipQueryNumberOfSources, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum, set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	Choice() PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
	// setChoice assigns PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSources
	setChoice(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSources.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSources
	SetValue(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// HasValue checks if Value has been set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSources.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSources
	SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// Auto returns uint32, set in PatternFlowIgmpv3MembershipQueryNumberOfSources.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	HasAuto() bool
	// Increment returns PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, set in PatternFlowIgmpv3MembershipQueryNumberOfSources.
	// PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter is integer counter pattern
	Increment() PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// SetIncrement assigns PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSources.
	// PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter, set in PatternFlowIgmpv3MembershipQueryNumberOfSources.
	// PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
	// SetDecrement assigns PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter provided by user to PatternFlowIgmpv3MembershipQueryNumberOfSources.
	// PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSources
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3MembershipQueryNumberOfSources
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3MembershipQueryNumberOfSources
var PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice = struct {
	VALUE     PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
	VALUES    PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
	AUTO      PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
	INCREMENT PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
	DECREMENT PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum("values"),
	AUTO:      PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum("auto"),
	INCREMENT: PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Choice() PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum {
	return PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) setChoice(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	intValue, ok := otg.PatternFlowIgmpv3MembershipQueryNumberOfSources_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3MembershipQueryNumberOfSources_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter().msg()
	}

	if value == PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3MembershipQueryNumberOfSources object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) SetValue(value uint32) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3MembershipQueryNumberOfSources object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) SetValues(value []uint32) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUES)
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
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Increment() PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter value in the PatternFlowIgmpv3MembershipQueryNumberOfSources object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) SetIncrement(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) Decrement() PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3MembershipQueryNumberOfSourcesCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter value in the PatternFlowIgmpv3MembershipQueryNumberOfSources object
func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) SetDecrement(value PatternFlowIgmpv3MembershipQueryNumberOfSourcesCounter) PatternFlowIgmpv3MembershipQueryNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryNumberOfSources.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3MembershipQueryNumberOfSources.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3MembershipQueryNumberOfSources.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3MembershipQueryNumberOfSources) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3MembershipQueryNumberOfSourcesChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3MembershipQueryNumberOfSources")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3MembershipQueryNumberOfSources_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3MembershipQueryNumberOfSources_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
