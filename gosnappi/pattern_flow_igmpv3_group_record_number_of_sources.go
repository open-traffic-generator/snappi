package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv3GroupRecordNumberOfSources *****
type patternFlowIgmpv3GroupRecordNumberOfSources struct {
	validation
	obj             *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	marshaller      marshalPatternFlowIgmpv3GroupRecordNumberOfSources
	unMarshaller    unMarshalPatternFlowIgmpv3GroupRecordNumberOfSources
	incrementHolder PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	decrementHolder PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
}

func NewPatternFlowIgmpv3GroupRecordNumberOfSources() PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj := patternFlowIgmpv3GroupRecordNumberOfSources{obj: &otg.PatternFlowIgmpv3GroupRecordNumberOfSources{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) msg() *otg.PatternFlowIgmpv3GroupRecordNumberOfSources {
	return obj.obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) setMsg(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSources) PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv3GroupRecordNumberOfSources struct {
	obj *patternFlowIgmpv3GroupRecordNumberOfSources
}

type marshalPatternFlowIgmpv3GroupRecordNumberOfSources interface {
	// ToProto marshals PatternFlowIgmpv3GroupRecordNumberOfSources to protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	ToProto() (*otg.PatternFlowIgmpv3GroupRecordNumberOfSources, error)
	// ToPbText marshals PatternFlowIgmpv3GroupRecordNumberOfSources to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv3GroupRecordNumberOfSources to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv3GroupRecordNumberOfSources to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources struct {
	obj *patternFlowIgmpv3GroupRecordNumberOfSources
}

type unMarshalPatternFlowIgmpv3GroupRecordNumberOfSources interface {
	// FromProto unmarshals PatternFlowIgmpv3GroupRecordNumberOfSources from protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	FromProto(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSources) (PatternFlowIgmpv3GroupRecordNumberOfSources, error)
	// FromPbText unmarshals PatternFlowIgmpv3GroupRecordNumberOfSources from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv3GroupRecordNumberOfSources from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv3GroupRecordNumberOfSources from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Marshal() marshalPatternFlowIgmpv3GroupRecordNumberOfSources {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv3GroupRecordNumberOfSources{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordNumberOfSources {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSources) ToProto() (*otg.PatternFlowIgmpv3GroupRecordNumberOfSources, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources) FromProto(msg *otg.PatternFlowIgmpv3GroupRecordNumberOfSources) (PatternFlowIgmpv3GroupRecordNumberOfSources, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSources) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSources) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv3GroupRecordNumberOfSources) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv3GroupRecordNumberOfSources) FromJson(value string) error {
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

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Clone() (PatternFlowIgmpv3GroupRecordNumberOfSources, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv3GroupRecordNumberOfSources()
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

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv3GroupRecordNumberOfSources is number of source addresses in this record. When auto, it is set to the number of entries in sources.
type PatternFlowIgmpv3GroupRecordNumberOfSources interface {
	Validation
	// msg marshals PatternFlowIgmpv3GroupRecordNumberOfSources to protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	// setMsg unmarshals PatternFlowIgmpv3GroupRecordNumberOfSources from protobuf object *otg.PatternFlowIgmpv3GroupRecordNumberOfSources
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv3GroupRecordNumberOfSources) PatternFlowIgmpv3GroupRecordNumberOfSources
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv3GroupRecordNumberOfSources
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv3GroupRecordNumberOfSources
	// validate validates PatternFlowIgmpv3GroupRecordNumberOfSources
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv3GroupRecordNumberOfSources, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum, set in PatternFlowIgmpv3GroupRecordNumberOfSources
	Choice() PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
	// setChoice assigns PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum provided by user to PatternFlowIgmpv3GroupRecordNumberOfSources
	setChoice(value PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum) PatternFlowIgmpv3GroupRecordNumberOfSources
	// HasChoice checks if Choice has been set in PatternFlowIgmpv3GroupRecordNumberOfSources
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSources.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv3GroupRecordNumberOfSources
	SetValue(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSources
	// HasValue checks if Value has been set in PatternFlowIgmpv3GroupRecordNumberOfSources
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSources.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv3GroupRecordNumberOfSources
	SetValues(value []uint32) PatternFlowIgmpv3GroupRecordNumberOfSources
	// Auto returns uint32, set in PatternFlowIgmpv3GroupRecordNumberOfSources.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIgmpv3GroupRecordNumberOfSources
	HasAuto() bool
	// Increment returns PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, set in PatternFlowIgmpv3GroupRecordNumberOfSources.
	// PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter is integer counter pattern
	Increment() PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// SetIncrement assigns PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter provided by user to PatternFlowIgmpv3GroupRecordNumberOfSources.
	// PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSources
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv3GroupRecordNumberOfSources
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter, set in PatternFlowIgmpv3GroupRecordNumberOfSources.
	// PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter is integer counter pattern
	Decrement() PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
	// SetDecrement assigns PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter provided by user to PatternFlowIgmpv3GroupRecordNumberOfSources.
	// PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSources
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv3GroupRecordNumberOfSources
	HasDecrement() bool
	setNil()
}

type PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum string

// Enum of Choice on PatternFlowIgmpv3GroupRecordNumberOfSources
var PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice = struct {
	VALUE     PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
	VALUES    PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
	AUTO      PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
	INCREMENT PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
	DECREMENT PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum
}{
	VALUE:     PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum("values"),
	AUTO:      PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum("auto"),
	INCREMENT: PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Choice() PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum {
	return PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) setChoice(value PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum) PatternFlowIgmpv3GroupRecordNumberOfSources {
	intValue, ok := otg.PatternFlowIgmpv3GroupRecordNumberOfSources_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv3GroupRecordNumberOfSources_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter().msg()
	}

	if value == PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv3GroupRecordNumberOfSourcesCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv3GroupRecordNumberOfSources object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) SetValue(value uint32) PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv3GroupRecordNumberOfSources object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) SetValues(value []uint32) PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUES)
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
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Increment() PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv3GroupRecordNumberOfSourcesCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter value in the PatternFlowIgmpv3GroupRecordNumberOfSources object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) SetIncrement(value PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) Decrement() PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv3GroupRecordNumberOfSourcesCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter value in the PatternFlowIgmpv3GroupRecordNumberOfSources object
func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) SetDecrement(value PatternFlowIgmpv3GroupRecordNumberOfSourcesCounter) PatternFlowIgmpv3GroupRecordNumberOfSources {
	obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordNumberOfSources.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv3GroupRecordNumberOfSources.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv3GroupRecordNumberOfSources.Auto <= 65535 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternFlowIgmpv3GroupRecordNumberOfSources) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIgmpv3GroupRecordNumberOfSourcesChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIgmpv3GroupRecordNumberOfSourcesChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIgmpv3GroupRecordNumberOfSources")
			}
		} else {
			intVal := otg.PatternFlowIgmpv3GroupRecordNumberOfSources_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIgmpv3GroupRecordNumberOfSources_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
