package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowRSVPPathSenderTspecIntServZeroBit *****
type patternFlowRSVPPathSenderTspecIntServZeroBit struct {
	validation
	obj             *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	marshaller      marshalPatternFlowRSVPPathSenderTspecIntServZeroBit
	unMarshaller    unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBit
	incrementHolder PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	decrementHolder PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
}

func NewPatternFlowRSVPPathSenderTspecIntServZeroBit() PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj := patternFlowRSVPPathSenderTspecIntServZeroBit{obj: &otg.PatternFlowRSVPPathSenderTspecIntServZeroBit{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) msg() *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit {
	return obj.obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) setMsg(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowRSVPPathSenderTspecIntServZeroBit struct {
	obj *patternFlowRSVPPathSenderTspecIntServZeroBit
}

type marshalPatternFlowRSVPPathSenderTspecIntServZeroBit interface {
	// ToProto marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServZeroBit, error)
	// ToPbText marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit struct {
	obj *patternFlowRSVPPathSenderTspecIntServZeroBit
}

type unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBit interface {
	// FromProto unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBit from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit) (PatternFlowRSVPPathSenderTspecIntServZeroBit, error)
	// FromPbText unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBit from JSON text
	FromJson(value string) error
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Marshal() marshalPatternFlowRSVPPathSenderTspecIntServZeroBit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowRSVPPathSenderTspecIntServZeroBit{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBit) ToProto() (*otg.PatternFlowRSVPPathSenderTspecIntServZeroBit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit) FromProto(msg *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit) (PatternFlowRSVPPathSenderTspecIntServZeroBit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBit) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit) FromPbText(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBit) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit) FromYaml(value string) error {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBit) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowRSVPPathSenderTspecIntServZeroBit) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowRSVPPathSenderTspecIntServZeroBit) FromJson(value string) error {
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

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Clone() (PatternFlowRSVPPathSenderTspecIntServZeroBit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowRSVPPathSenderTspecIntServZeroBit()
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

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowRSVPPathSenderTspecIntServZeroBit is mUST be 0.
type PatternFlowRSVPPathSenderTspecIntServZeroBit interface {
	Validation
	// msg marshals PatternFlowRSVPPathSenderTspecIntServZeroBit to protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	// and doesn't set defaults
	msg() *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	// setMsg unmarshals PatternFlowRSVPPathSenderTspecIntServZeroBit from protobuf object *otg.PatternFlowRSVPPathSenderTspecIntServZeroBit
	// and doesn't set defaults
	setMsg(*otg.PatternFlowRSVPPathSenderTspecIntServZeroBit) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// provides marshal interface
	Marshal() marshalPatternFlowRSVPPathSenderTspecIntServZeroBit
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowRSVPPathSenderTspecIntServZeroBit
	// validate validates PatternFlowRSVPPathSenderTspecIntServZeroBit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowRSVPPathSenderTspecIntServZeroBit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum, set in PatternFlowRSVPPathSenderTspecIntServZeroBit
	Choice() PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum
	// setChoice assigns PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBit
	setChoice(value PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// HasChoice checks if Choice has been set in PatternFlowRSVPPathSenderTspecIntServZeroBit
	HasChoice() bool
	// Value returns uint32, set in PatternFlowRSVPPathSenderTspecIntServZeroBit.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBit
	SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// HasValue checks if Value has been set in PatternFlowRSVPPathSenderTspecIntServZeroBit
	HasValue() bool
	// Values returns []uint32, set in PatternFlowRSVPPathSenderTspecIntServZeroBit.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBit
	SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// Increment returns PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, set in PatternFlowRSVPPathSenderTspecIntServZeroBit.
	// PatternFlowRSVPPathSenderTspecIntServZeroBitCounter is integer counter pattern
	Increment() PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// SetIncrement assigns PatternFlowRSVPPathSenderTspecIntServZeroBitCounter provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBit.
	// PatternFlowRSVPPathSenderTspecIntServZeroBitCounter is integer counter pattern
	SetIncrement(value PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// HasIncrement checks if Increment has been set in PatternFlowRSVPPathSenderTspecIntServZeroBit
	HasIncrement() bool
	// Decrement returns PatternFlowRSVPPathSenderTspecIntServZeroBitCounter, set in PatternFlowRSVPPathSenderTspecIntServZeroBit.
	// PatternFlowRSVPPathSenderTspecIntServZeroBitCounter is integer counter pattern
	Decrement() PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
	// SetDecrement assigns PatternFlowRSVPPathSenderTspecIntServZeroBitCounter provided by user to PatternFlowRSVPPathSenderTspecIntServZeroBit.
	// PatternFlowRSVPPathSenderTspecIntServZeroBitCounter is integer counter pattern
	SetDecrement(value PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBit
	// HasDecrement checks if Decrement has been set in PatternFlowRSVPPathSenderTspecIntServZeroBit
	HasDecrement() bool
	setNil()
}

type PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum string

// Enum of Choice on PatternFlowRSVPPathSenderTspecIntServZeroBit
var PatternFlowRSVPPathSenderTspecIntServZeroBitChoice = struct {
	VALUE     PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum
	VALUES    PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum
	INCREMENT PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum
	DECREMENT PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum
}{
	VALUE:     PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum("value"),
	VALUES:    PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum("values"),
	INCREMENT: PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum("increment"),
	DECREMENT: PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum("decrement"),
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Choice() PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum {
	return PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) setChoice(value PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	intValue, ok := otg.PatternFlowRSVPPathSenderTspecIntServZeroBit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowRSVPPathSenderTspecIntServZeroBit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowRSVPPathSenderTspecIntServZeroBitCounter().msg()
	}

	if value == PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowRSVPPathSenderTspecIntServZeroBitCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowRSVPPathSenderTspecIntServZeroBit object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) SetValue(value uint32) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowRSVPPathSenderTspecIntServZeroBit object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) SetValues(value []uint32) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Increment() PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowRSVPPathSenderTspecIntServZeroBitCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowRSVPPathSenderTspecIntServZeroBitCounter value in the PatternFlowRSVPPathSenderTspecIntServZeroBit object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) SetIncrement(value PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) Decrement() PatternFlowRSVPPathSenderTspecIntServZeroBitCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowRSVPPathSenderTspecIntServZeroBitCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowRSVPPathSenderTspecIntServZeroBitCounter
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowRSVPPathSenderTspecIntServZeroBitCounter value in the PatternFlowRSVPPathSenderTspecIntServZeroBit object
func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) SetDecrement(value PatternFlowRSVPPathSenderTspecIntServZeroBitCounter) PatternFlowRSVPPathSenderTspecIntServZeroBit {
	obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowRSVPPathSenderTspecIntServZeroBit.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowRSVPPathSenderTspecIntServZeroBit.Values <= 1 but Got %d", item))
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

func (obj *patternFlowRSVPPathSenderTspecIntServZeroBit) setDefault() {
	var choices_set int = 0
	var choice PatternFlowRSVPPathSenderTspecIntServZeroBitChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowRSVPPathSenderTspecIntServZeroBitChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowRSVPPathSenderTspecIntServZeroBit")
			}
		} else {
			intVal := otg.PatternFlowRSVPPathSenderTspecIntServZeroBit_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowRSVPPathSenderTspecIntServZeroBit_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
