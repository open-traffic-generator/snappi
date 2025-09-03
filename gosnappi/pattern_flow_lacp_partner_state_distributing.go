package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpPartnerStateDistributing *****
type patternFlowLacpPartnerStateDistributing struct {
	validation
	obj             *otg.PatternFlowLacpPartnerStateDistributing
	marshaller      marshalPatternFlowLacpPartnerStateDistributing
	unMarshaller    unMarshalPatternFlowLacpPartnerStateDistributing
	incrementHolder PatternFlowLacpPartnerStateDistributingCounter
	decrementHolder PatternFlowLacpPartnerStateDistributingCounter
}

func NewPatternFlowLacpPartnerStateDistributing() PatternFlowLacpPartnerStateDistributing {
	obj := patternFlowLacpPartnerStateDistributing{obj: &otg.PatternFlowLacpPartnerStateDistributing{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpPartnerStateDistributing) msg() *otg.PatternFlowLacpPartnerStateDistributing {
	return obj.obj
}

func (obj *patternFlowLacpPartnerStateDistributing) setMsg(msg *otg.PatternFlowLacpPartnerStateDistributing) PatternFlowLacpPartnerStateDistributing {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpPartnerStateDistributing struct {
	obj *patternFlowLacpPartnerStateDistributing
}

type marshalPatternFlowLacpPartnerStateDistributing interface {
	// ToProto marshals PatternFlowLacpPartnerStateDistributing to protobuf object *otg.PatternFlowLacpPartnerStateDistributing
	ToProto() (*otg.PatternFlowLacpPartnerStateDistributing, error)
	// ToPbText marshals PatternFlowLacpPartnerStateDistributing to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpPartnerStateDistributing to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpPartnerStateDistributing to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpPartnerStateDistributing struct {
	obj *patternFlowLacpPartnerStateDistributing
}

type unMarshalPatternFlowLacpPartnerStateDistributing interface {
	// FromProto unmarshals PatternFlowLacpPartnerStateDistributing from protobuf object *otg.PatternFlowLacpPartnerStateDistributing
	FromProto(msg *otg.PatternFlowLacpPartnerStateDistributing) (PatternFlowLacpPartnerStateDistributing, error)
	// FromPbText unmarshals PatternFlowLacpPartnerStateDistributing from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpPartnerStateDistributing from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpPartnerStateDistributing from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpPartnerStateDistributing) Marshal() marshalPatternFlowLacpPartnerStateDistributing {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpPartnerStateDistributing{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpPartnerStateDistributing) Unmarshal() unMarshalPatternFlowLacpPartnerStateDistributing {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpPartnerStateDistributing{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpPartnerStateDistributing) ToProto() (*otg.PatternFlowLacpPartnerStateDistributing, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpPartnerStateDistributing) FromProto(msg *otg.PatternFlowLacpPartnerStateDistributing) (PatternFlowLacpPartnerStateDistributing, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpPartnerStateDistributing) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributing) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDistributing) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributing) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpPartnerStateDistributing) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpPartnerStateDistributing) FromJson(value string) error {
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

func (obj *patternFlowLacpPartnerStateDistributing) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDistributing) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpPartnerStateDistributing) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpPartnerStateDistributing) Clone() (PatternFlowLacpPartnerStateDistributing, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpPartnerStateDistributing()
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

func (obj *patternFlowLacpPartnerStateDistributing) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
type PatternFlowLacpPartnerStateDistributing interface {
	Validation
	// msg marshals PatternFlowLacpPartnerStateDistributing to protobuf object *otg.PatternFlowLacpPartnerStateDistributing
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpPartnerStateDistributing
	// setMsg unmarshals PatternFlowLacpPartnerStateDistributing from protobuf object *otg.PatternFlowLacpPartnerStateDistributing
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpPartnerStateDistributing) PatternFlowLacpPartnerStateDistributing
	// provides marshal interface
	Marshal() marshalPatternFlowLacpPartnerStateDistributing
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpPartnerStateDistributing
	// validate validates PatternFlowLacpPartnerStateDistributing
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpPartnerStateDistributing, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpPartnerStateDistributingChoiceEnum, set in PatternFlowLacpPartnerStateDistributing
	Choice() PatternFlowLacpPartnerStateDistributingChoiceEnum
	// setChoice assigns PatternFlowLacpPartnerStateDistributingChoiceEnum provided by user to PatternFlowLacpPartnerStateDistributing
	setChoice(value PatternFlowLacpPartnerStateDistributingChoiceEnum) PatternFlowLacpPartnerStateDistributing
	// HasChoice checks if Choice has been set in PatternFlowLacpPartnerStateDistributing
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpPartnerStateDistributing.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpPartnerStateDistributing
	SetValue(value uint32) PatternFlowLacpPartnerStateDistributing
	// HasValue checks if Value has been set in PatternFlowLacpPartnerStateDistributing
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpPartnerStateDistributing.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpPartnerStateDistributing
	SetValues(value []uint32) PatternFlowLacpPartnerStateDistributing
	// Increment returns PatternFlowLacpPartnerStateDistributingCounter, set in PatternFlowLacpPartnerStateDistributing.
	// PatternFlowLacpPartnerStateDistributingCounter is integer counter pattern
	Increment() PatternFlowLacpPartnerStateDistributingCounter
	// SetIncrement assigns PatternFlowLacpPartnerStateDistributingCounter provided by user to PatternFlowLacpPartnerStateDistributing.
	// PatternFlowLacpPartnerStateDistributingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributing
	// HasIncrement checks if Increment has been set in PatternFlowLacpPartnerStateDistributing
	HasIncrement() bool
	// Decrement returns PatternFlowLacpPartnerStateDistributingCounter, set in PatternFlowLacpPartnerStateDistributing.
	// PatternFlowLacpPartnerStateDistributingCounter is integer counter pattern
	Decrement() PatternFlowLacpPartnerStateDistributingCounter
	// SetDecrement assigns PatternFlowLacpPartnerStateDistributingCounter provided by user to PatternFlowLacpPartnerStateDistributing.
	// PatternFlowLacpPartnerStateDistributingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributing
	// HasDecrement checks if Decrement has been set in PatternFlowLacpPartnerStateDistributing
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpPartnerStateDistributingChoiceEnum string

// Enum of Choice on PatternFlowLacpPartnerStateDistributing
var PatternFlowLacpPartnerStateDistributingChoice = struct {
	VALUE     PatternFlowLacpPartnerStateDistributingChoiceEnum
	VALUES    PatternFlowLacpPartnerStateDistributingChoiceEnum
	INCREMENT PatternFlowLacpPartnerStateDistributingChoiceEnum
	DECREMENT PatternFlowLacpPartnerStateDistributingChoiceEnum
}{
	VALUE:     PatternFlowLacpPartnerStateDistributingChoiceEnum("value"),
	VALUES:    PatternFlowLacpPartnerStateDistributingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpPartnerStateDistributingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpPartnerStateDistributingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpPartnerStateDistributing) Choice() PatternFlowLacpPartnerStateDistributingChoiceEnum {
	return PatternFlowLacpPartnerStateDistributingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpPartnerStateDistributing) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpPartnerStateDistributing) setChoice(value PatternFlowLacpPartnerStateDistributingChoiceEnum) PatternFlowLacpPartnerStateDistributing {
	intValue, ok := otg.PatternFlowLacpPartnerStateDistributing_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpPartnerStateDistributingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpPartnerStateDistributing_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpPartnerStateDistributingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpPartnerStateDistributingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpPartnerStateDistributingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpPartnerStateDistributingCounter().msg()
	}

	if value == PatternFlowLacpPartnerStateDistributingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpPartnerStateDistributingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateDistributing) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpPartnerStateDistributing) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpPartnerStateDistributing object
func (obj *patternFlowLacpPartnerStateDistributing) SetValue(value uint32) PatternFlowLacpPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpPartnerStateDistributing) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpPartnerStateDistributing object
func (obj *patternFlowLacpPartnerStateDistributing) SetValues(value []uint32) PatternFlowLacpPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateDistributingCounter
func (obj *patternFlowLacpPartnerStateDistributing) Increment() PatternFlowLacpPartnerStateDistributingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpPartnerStateDistributingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpPartnerStateDistributingCounter
func (obj *patternFlowLacpPartnerStateDistributing) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpPartnerStateDistributingCounter value in the PatternFlowLacpPartnerStateDistributing object
func (obj *patternFlowLacpPartnerStateDistributing) SetIncrement(value PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateDistributingCounter
func (obj *patternFlowLacpPartnerStateDistributing) Decrement() PatternFlowLacpPartnerStateDistributingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpPartnerStateDistributingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpPartnerStateDistributingCounter
func (obj *patternFlowLacpPartnerStateDistributing) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpPartnerStateDistributingCounter value in the PatternFlowLacpPartnerStateDistributing object
func (obj *patternFlowLacpPartnerStateDistributing) SetDecrement(value PatternFlowLacpPartnerStateDistributingCounter) PatternFlowLacpPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpPartnerStateDistributing) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpPartnerStateDistributing.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpPartnerStateDistributing.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpPartnerStateDistributing) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpPartnerStateDistributingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDistributingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDistributingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDistributingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpPartnerStateDistributingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpPartnerStateDistributingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpPartnerStateDistributing")
			}
		} else {
			intVal := otg.PatternFlowLacpPartnerStateDistributing_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpPartnerStateDistributing_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
