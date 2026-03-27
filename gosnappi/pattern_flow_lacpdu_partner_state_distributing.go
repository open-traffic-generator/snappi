package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowLacpduPartnerStateDistributing *****
type patternFlowLacpduPartnerStateDistributing struct {
	validation
	obj             *otg.PatternFlowLacpduPartnerStateDistributing
	marshaller      marshalPatternFlowLacpduPartnerStateDistributing
	unMarshaller    unMarshalPatternFlowLacpduPartnerStateDistributing
	incrementHolder PatternFlowLacpduPartnerStateDistributingCounter
	decrementHolder PatternFlowLacpduPartnerStateDistributingCounter
}

func NewPatternFlowLacpduPartnerStateDistributing() PatternFlowLacpduPartnerStateDistributing {
	obj := patternFlowLacpduPartnerStateDistributing{obj: &otg.PatternFlowLacpduPartnerStateDistributing{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowLacpduPartnerStateDistributing) msg() *otg.PatternFlowLacpduPartnerStateDistributing {
	return obj.obj
}

func (obj *patternFlowLacpduPartnerStateDistributing) setMsg(msg *otg.PatternFlowLacpduPartnerStateDistributing) PatternFlowLacpduPartnerStateDistributing {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowLacpduPartnerStateDistributing struct {
	obj *patternFlowLacpduPartnerStateDistributing
}

type marshalPatternFlowLacpduPartnerStateDistributing interface {
	// ToProto marshals PatternFlowLacpduPartnerStateDistributing to protobuf object *otg.PatternFlowLacpduPartnerStateDistributing
	ToProto() (*otg.PatternFlowLacpduPartnerStateDistributing, error)
	// ToPbText marshals PatternFlowLacpduPartnerStateDistributing to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowLacpduPartnerStateDistributing to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowLacpduPartnerStateDistributing to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowLacpduPartnerStateDistributing struct {
	obj *patternFlowLacpduPartnerStateDistributing
}

type unMarshalPatternFlowLacpduPartnerStateDistributing interface {
	// FromProto unmarshals PatternFlowLacpduPartnerStateDistributing from protobuf object *otg.PatternFlowLacpduPartnerStateDistributing
	FromProto(msg *otg.PatternFlowLacpduPartnerStateDistributing) (PatternFlowLacpduPartnerStateDistributing, error)
	// FromPbText unmarshals PatternFlowLacpduPartnerStateDistributing from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowLacpduPartnerStateDistributing from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowLacpduPartnerStateDistributing from JSON text
	FromJson(value string) error
}

func (obj *patternFlowLacpduPartnerStateDistributing) Marshal() marshalPatternFlowLacpduPartnerStateDistributing {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowLacpduPartnerStateDistributing{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowLacpduPartnerStateDistributing) Unmarshal() unMarshalPatternFlowLacpduPartnerStateDistributing {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowLacpduPartnerStateDistributing{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowLacpduPartnerStateDistributing) ToProto() (*otg.PatternFlowLacpduPartnerStateDistributing, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowLacpduPartnerStateDistributing) FromProto(msg *otg.PatternFlowLacpduPartnerStateDistributing) (PatternFlowLacpduPartnerStateDistributing, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowLacpduPartnerStateDistributing) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributing) FromPbText(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDistributing) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributing) FromYaml(value string) error {
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

func (m *marshalpatternFlowLacpduPartnerStateDistributing) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowLacpduPartnerStateDistributing) FromJson(value string) error {
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

func (obj *patternFlowLacpduPartnerStateDistributing) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDistributing) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowLacpduPartnerStateDistributing) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowLacpduPartnerStateDistributing) Clone() (PatternFlowLacpduPartnerStateDistributing, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowLacpduPartnerStateDistributing()
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

func (obj *patternFlowLacpduPartnerStateDistributing) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowLacpduPartnerStateDistributing is distributing (1=Enabled, 0=Disabled)
type PatternFlowLacpduPartnerStateDistributing interface {
	Validation
	// msg marshals PatternFlowLacpduPartnerStateDistributing to protobuf object *otg.PatternFlowLacpduPartnerStateDistributing
	// and doesn't set defaults
	msg() *otg.PatternFlowLacpduPartnerStateDistributing
	// setMsg unmarshals PatternFlowLacpduPartnerStateDistributing from protobuf object *otg.PatternFlowLacpduPartnerStateDistributing
	// and doesn't set defaults
	setMsg(*otg.PatternFlowLacpduPartnerStateDistributing) PatternFlowLacpduPartnerStateDistributing
	// provides marshal interface
	Marshal() marshalPatternFlowLacpduPartnerStateDistributing
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowLacpduPartnerStateDistributing
	// validate validates PatternFlowLacpduPartnerStateDistributing
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowLacpduPartnerStateDistributing, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowLacpduPartnerStateDistributingChoiceEnum, set in PatternFlowLacpduPartnerStateDistributing
	Choice() PatternFlowLacpduPartnerStateDistributingChoiceEnum
	// setChoice assigns PatternFlowLacpduPartnerStateDistributingChoiceEnum provided by user to PatternFlowLacpduPartnerStateDistributing
	setChoice(value PatternFlowLacpduPartnerStateDistributingChoiceEnum) PatternFlowLacpduPartnerStateDistributing
	// HasChoice checks if Choice has been set in PatternFlowLacpduPartnerStateDistributing
	HasChoice() bool
	// Value returns uint32, set in PatternFlowLacpduPartnerStateDistributing.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowLacpduPartnerStateDistributing
	SetValue(value uint32) PatternFlowLacpduPartnerStateDistributing
	// HasValue checks if Value has been set in PatternFlowLacpduPartnerStateDistributing
	HasValue() bool
	// Values returns []uint32, set in PatternFlowLacpduPartnerStateDistributing.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowLacpduPartnerStateDistributing
	SetValues(value []uint32) PatternFlowLacpduPartnerStateDistributing
	// Increment returns PatternFlowLacpduPartnerStateDistributingCounter, set in PatternFlowLacpduPartnerStateDistributing.
	// PatternFlowLacpduPartnerStateDistributingCounter is integer counter pattern
	Increment() PatternFlowLacpduPartnerStateDistributingCounter
	// SetIncrement assigns PatternFlowLacpduPartnerStateDistributingCounter provided by user to PatternFlowLacpduPartnerStateDistributing.
	// PatternFlowLacpduPartnerStateDistributingCounter is integer counter pattern
	SetIncrement(value PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributing
	// HasIncrement checks if Increment has been set in PatternFlowLacpduPartnerStateDistributing
	HasIncrement() bool
	// Decrement returns PatternFlowLacpduPartnerStateDistributingCounter, set in PatternFlowLacpduPartnerStateDistributing.
	// PatternFlowLacpduPartnerStateDistributingCounter is integer counter pattern
	Decrement() PatternFlowLacpduPartnerStateDistributingCounter
	// SetDecrement assigns PatternFlowLacpduPartnerStateDistributingCounter provided by user to PatternFlowLacpduPartnerStateDistributing.
	// PatternFlowLacpduPartnerStateDistributingCounter is integer counter pattern
	SetDecrement(value PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributing
	// HasDecrement checks if Decrement has been set in PatternFlowLacpduPartnerStateDistributing
	HasDecrement() bool
	setNil()
}

type PatternFlowLacpduPartnerStateDistributingChoiceEnum string

// Enum of Choice on PatternFlowLacpduPartnerStateDistributing
var PatternFlowLacpduPartnerStateDistributingChoice = struct {
	VALUE     PatternFlowLacpduPartnerStateDistributingChoiceEnum
	VALUES    PatternFlowLacpduPartnerStateDistributingChoiceEnum
	INCREMENT PatternFlowLacpduPartnerStateDistributingChoiceEnum
	DECREMENT PatternFlowLacpduPartnerStateDistributingChoiceEnum
}{
	VALUE:     PatternFlowLacpduPartnerStateDistributingChoiceEnum("value"),
	VALUES:    PatternFlowLacpduPartnerStateDistributingChoiceEnum("values"),
	INCREMENT: PatternFlowLacpduPartnerStateDistributingChoiceEnum("increment"),
	DECREMENT: PatternFlowLacpduPartnerStateDistributingChoiceEnum("decrement"),
}

func (obj *patternFlowLacpduPartnerStateDistributing) Choice() PatternFlowLacpduPartnerStateDistributingChoiceEnum {
	return PatternFlowLacpduPartnerStateDistributingChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowLacpduPartnerStateDistributing) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowLacpduPartnerStateDistributing) setChoice(value PatternFlowLacpduPartnerStateDistributingChoiceEnum) PatternFlowLacpduPartnerStateDistributing {
	intValue, ok := otg.PatternFlowLacpduPartnerStateDistributing_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowLacpduPartnerStateDistributingChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowLacpduPartnerStateDistributing_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowLacpduPartnerStateDistributingChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowLacpduPartnerStateDistributingChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowLacpduPartnerStateDistributingChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowLacpduPartnerStateDistributingCounter().msg()
	}

	if value == PatternFlowLacpduPartnerStateDistributingChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowLacpduPartnerStateDistributingCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributing) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowLacpduPartnerStateDistributing) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowLacpduPartnerStateDistributing object
func (obj *patternFlowLacpduPartnerStateDistributing) SetValue(value uint32) PatternFlowLacpduPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowLacpduPartnerStateDistributing) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowLacpduPartnerStateDistributing object
func (obj *patternFlowLacpduPartnerStateDistributing) SetValues(value []uint32) PatternFlowLacpduPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateDistributingCounter
func (obj *patternFlowLacpduPartnerStateDistributing) Increment() PatternFlowLacpduPartnerStateDistributingCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowLacpduPartnerStateDistributingCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowLacpduPartnerStateDistributingCounter
func (obj *patternFlowLacpduPartnerStateDistributing) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowLacpduPartnerStateDistributingCounter value in the PatternFlowLacpduPartnerStateDistributing object
func (obj *patternFlowLacpduPartnerStateDistributing) SetIncrement(value PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateDistributingCounter
func (obj *patternFlowLacpduPartnerStateDistributing) Decrement() PatternFlowLacpduPartnerStateDistributingCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowLacpduPartnerStateDistributingCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowLacpduPartnerStateDistributingCounter
func (obj *patternFlowLacpduPartnerStateDistributing) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowLacpduPartnerStateDistributingCounter value in the PatternFlowLacpduPartnerStateDistributing object
func (obj *patternFlowLacpduPartnerStateDistributing) SetDecrement(value PatternFlowLacpduPartnerStateDistributingCounter) PatternFlowLacpduPartnerStateDistributing {
	obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowLacpduPartnerStateDistributing) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowLacpduPartnerStateDistributing.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowLacpduPartnerStateDistributing.Values <= 1 but Got %d", item))
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

func (obj *patternFlowLacpduPartnerStateDistributing) setDefault() {
	var choices_set int = 0
	var choice PatternFlowLacpduPartnerStateDistributingChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDistributingChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDistributingChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDistributingChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowLacpduPartnerStateDistributingChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowLacpduPartnerStateDistributingChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowLacpduPartnerStateDistributing")
			}
		} else {
			intVal := otg.PatternFlowLacpduPartnerStateDistributing_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowLacpduPartnerStateDistributing_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
