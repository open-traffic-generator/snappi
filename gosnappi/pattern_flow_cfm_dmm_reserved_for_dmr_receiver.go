package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmrReceiver *****
type patternFlowCfmDmmReservedForDmrReceiver struct {
	validation
	obj             *otg.PatternFlowCfmDmmReservedForDmrReceiver
	marshaller      marshalPatternFlowCfmDmmReservedForDmrReceiver
	unMarshaller    unMarshalPatternFlowCfmDmmReservedForDmrReceiver
	incrementHolder PatternFlowCfmDmmReservedForDmrReceiverCounter
	decrementHolder PatternFlowCfmDmmReservedForDmrReceiverCounter
}

func NewPatternFlowCfmDmmReservedForDmrReceiver() PatternFlowCfmDmmReservedForDmrReceiver {
	obj := patternFlowCfmDmmReservedForDmrReceiver{obj: &otg.PatternFlowCfmDmmReservedForDmrReceiver{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) msg() *otg.PatternFlowCfmDmmReservedForDmrReceiver {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmrReceiver) PatternFlowCfmDmmReservedForDmrReceiver {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmrReceiver struct {
	obj *patternFlowCfmDmmReservedForDmrReceiver
}

type marshalPatternFlowCfmDmmReservedForDmrReceiver interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmrReceiver to protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiver
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmrReceiver, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmrReceiver to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmrReceiver to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmrReceiver to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmrReceiver struct {
	obj *patternFlowCfmDmmReservedForDmrReceiver
}

type unMarshalPatternFlowCfmDmmReservedForDmrReceiver interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmrReceiver from protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiver
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrReceiver) (PatternFlowCfmDmmReservedForDmrReceiver, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmrReceiver from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmrReceiver from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmrReceiver from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) Marshal() marshalPatternFlowCfmDmmReservedForDmrReceiver {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmrReceiver{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrReceiver {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmrReceiver{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiver) ToProto() (*otg.PatternFlowCfmDmmReservedForDmrReceiver, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiver) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmrReceiver) (PatternFlowCfmDmmReservedForDmrReceiver, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiver) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiver) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiver) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiver) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmrReceiver) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmrReceiver) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmrReceiver) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) Clone() (PatternFlowCfmDmmReservedForDmrReceiver, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmrReceiver()
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

func (obj *patternFlowCfmDmmReservedForDmrReceiver) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmmReservedForDmrReceiver is reserved for the DMR receiving equipment to record its receive timestamp when the DMR arrives back at the originating MEP. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmmReservedForDmrReceiver interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmrReceiver to protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiver
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmrReceiver
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmrReceiver from protobuf object *otg.PatternFlowCfmDmmReservedForDmrReceiver
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmrReceiver) PatternFlowCfmDmmReservedForDmrReceiver
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmrReceiver
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmrReceiver
	// validate validates PatternFlowCfmDmmReservedForDmrReceiver
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmrReceiver, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum, set in PatternFlowCfmDmmReservedForDmrReceiver
	Choice() PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum
	// setChoice assigns PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum provided by user to PatternFlowCfmDmmReservedForDmrReceiver
	setChoice(value PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum) PatternFlowCfmDmmReservedForDmrReceiver
	// HasChoice checks if Choice has been set in PatternFlowCfmDmmReservedForDmrReceiver
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmmReservedForDmrReceiver.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmrReceiver
	SetValue(value uint64) PatternFlowCfmDmmReservedForDmrReceiver
	// HasValue checks if Value has been set in PatternFlowCfmDmmReservedForDmrReceiver
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmmReservedForDmrReceiver.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmmReservedForDmrReceiver
	SetValues(value []uint64) PatternFlowCfmDmmReservedForDmrReceiver
	// Increment returns PatternFlowCfmDmmReservedForDmrReceiverCounter, set in PatternFlowCfmDmmReservedForDmrReceiver.
	// PatternFlowCfmDmmReservedForDmrReceiverCounter is integer counter pattern
	Increment() PatternFlowCfmDmmReservedForDmrReceiverCounter
	// SetIncrement assigns PatternFlowCfmDmmReservedForDmrReceiverCounter provided by user to PatternFlowCfmDmmReservedForDmrReceiver.
	// PatternFlowCfmDmmReservedForDmrReceiverCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiver
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmmReservedForDmrReceiver
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmmReservedForDmrReceiverCounter, set in PatternFlowCfmDmmReservedForDmrReceiver.
	// PatternFlowCfmDmmReservedForDmrReceiverCounter is integer counter pattern
	Decrement() PatternFlowCfmDmmReservedForDmrReceiverCounter
	// SetDecrement assigns PatternFlowCfmDmmReservedForDmrReceiverCounter provided by user to PatternFlowCfmDmmReservedForDmrReceiver.
	// PatternFlowCfmDmmReservedForDmrReceiverCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiver
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmmReservedForDmrReceiver
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum string

// Enum of Choice on PatternFlowCfmDmmReservedForDmrReceiver
var PatternFlowCfmDmmReservedForDmrReceiverChoice = struct {
	VALUE     PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum
	VALUES    PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum
	INCREMENT PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum
	DECREMENT PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum
}{
	VALUE:     PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) Choice() PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum {
	return PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmmReservedForDmrReceiver) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) setChoice(value PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum) PatternFlowCfmDmmReservedForDmrReceiver {
	intValue, ok := otg.PatternFlowCfmDmmReservedForDmrReceiver_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmmReservedForDmrReceiver_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmrReceiverChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmmReservedForDmrReceiverCounter().msg()
	}

	if value == PatternFlowCfmDmmReservedForDmrReceiverChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmmReservedForDmrReceiverCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiver) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiver) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmmReservedForDmrReceiver object
func (obj *patternFlowCfmDmmReservedForDmrReceiver) SetValue(value uint64) PatternFlowCfmDmmReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmmReservedForDmrReceiver) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmmReservedForDmrReceiver object
func (obj *patternFlowCfmDmmReservedForDmrReceiver) SetValues(value []uint64) PatternFlowCfmDmmReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmrReceiver) Increment() PatternFlowCfmDmmReservedForDmrReceiverCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmmReservedForDmrReceiverCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmrReceiver) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmmReservedForDmrReceiverCounter value in the PatternFlowCfmDmmReservedForDmrReceiver object
func (obj *patternFlowCfmDmmReservedForDmrReceiver) SetIncrement(value PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmrReceiver) Decrement() PatternFlowCfmDmmReservedForDmrReceiverCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmmReservedForDmrReceiverCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmrReceiver) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmmReservedForDmrReceiverCounter value in the PatternFlowCfmDmmReservedForDmrReceiver object
func (obj *patternFlowCfmDmmReservedForDmrReceiver) SetDecrement(value PatternFlowCfmDmmReservedForDmrReceiverCounter) PatternFlowCfmDmmReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmmReservedForDmrReceiver) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmmReservedForDmrReceiver) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmmReservedForDmrReceiverChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrReceiverChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmrReceiverChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmmReservedForDmrReceiverChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmmReservedForDmrReceiver")
			}
		} else {
			intVal := otg.PatternFlowCfmDmmReservedForDmrReceiver_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmmReservedForDmrReceiver_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
