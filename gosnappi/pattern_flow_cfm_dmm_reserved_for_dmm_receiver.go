package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmReservedForDmmReceiver *****
type patternFlowCfmDmmReservedForDmmReceiver struct {
	validation
	obj             *otg.PatternFlowCfmDmmReservedForDmmReceiver
	marshaller      marshalPatternFlowCfmDmmReservedForDmmReceiver
	unMarshaller    unMarshalPatternFlowCfmDmmReservedForDmmReceiver
	incrementHolder PatternFlowCfmDmmReservedForDmmReceiverCounter
	decrementHolder PatternFlowCfmDmmReservedForDmmReceiverCounter
}

func NewPatternFlowCfmDmmReservedForDmmReceiver() PatternFlowCfmDmmReservedForDmmReceiver {
	obj := patternFlowCfmDmmReservedForDmmReceiver{obj: &otg.PatternFlowCfmDmmReservedForDmmReceiver{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) msg() *otg.PatternFlowCfmDmmReservedForDmmReceiver {
	return obj.obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) setMsg(msg *otg.PatternFlowCfmDmmReservedForDmmReceiver) PatternFlowCfmDmmReservedForDmmReceiver {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmReservedForDmmReceiver struct {
	obj *patternFlowCfmDmmReservedForDmmReceiver
}

type marshalPatternFlowCfmDmmReservedForDmmReceiver interface {
	// ToProto marshals PatternFlowCfmDmmReservedForDmmReceiver to protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiver
	ToProto() (*otg.PatternFlowCfmDmmReservedForDmmReceiver, error)
	// ToPbText marshals PatternFlowCfmDmmReservedForDmmReceiver to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmReservedForDmmReceiver to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmReservedForDmmReceiver to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmReservedForDmmReceiver struct {
	obj *patternFlowCfmDmmReservedForDmmReceiver
}

type unMarshalPatternFlowCfmDmmReservedForDmmReceiver interface {
	// FromProto unmarshals PatternFlowCfmDmmReservedForDmmReceiver from protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiver
	FromProto(msg *otg.PatternFlowCfmDmmReservedForDmmReceiver) (PatternFlowCfmDmmReservedForDmmReceiver, error)
	// FromPbText unmarshals PatternFlowCfmDmmReservedForDmmReceiver from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmReservedForDmmReceiver from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmReservedForDmmReceiver from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) Marshal() marshalPatternFlowCfmDmmReservedForDmmReceiver {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmReservedForDmmReceiver{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmmReceiver {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmReservedForDmmReceiver{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiver) ToProto() (*otg.PatternFlowCfmDmmReservedForDmmReceiver, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiver) FromProto(msg *otg.PatternFlowCfmDmmReservedForDmmReceiver) (PatternFlowCfmDmmReservedForDmmReceiver, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiver) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiver) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiver) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiver) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmReservedForDmmReceiver) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmReservedForDmmReceiver) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmReservedForDmmReceiver) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) Clone() (PatternFlowCfmDmmReservedForDmmReceiver, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmReservedForDmmReceiver()
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

func (obj *patternFlowCfmDmmReservedForDmmReceiver) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmmReservedForDmmReceiver is reserved for the DMM receiving equipment to record its receive timestamp when forming the DMR. Set to zero by the DMM originator (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmmReservedForDmmReceiver interface {
	Validation
	// msg marshals PatternFlowCfmDmmReservedForDmmReceiver to protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiver
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmReservedForDmmReceiver
	// setMsg unmarshals PatternFlowCfmDmmReservedForDmmReceiver from protobuf object *otg.PatternFlowCfmDmmReservedForDmmReceiver
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmReservedForDmmReceiver) PatternFlowCfmDmmReservedForDmmReceiver
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmReservedForDmmReceiver
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmReservedForDmmReceiver
	// validate validates PatternFlowCfmDmmReservedForDmmReceiver
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmReservedForDmmReceiver, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum, set in PatternFlowCfmDmmReservedForDmmReceiver
	Choice() PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum
	// setChoice assigns PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum provided by user to PatternFlowCfmDmmReservedForDmmReceiver
	setChoice(value PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum) PatternFlowCfmDmmReservedForDmmReceiver
	// HasChoice checks if Choice has been set in PatternFlowCfmDmmReservedForDmmReceiver
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmmReservedForDmmReceiver.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmmReservedForDmmReceiver
	SetValue(value uint64) PatternFlowCfmDmmReservedForDmmReceiver
	// HasValue checks if Value has been set in PatternFlowCfmDmmReservedForDmmReceiver
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmmReservedForDmmReceiver.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmmReservedForDmmReceiver
	SetValues(value []uint64) PatternFlowCfmDmmReservedForDmmReceiver
	// Increment returns PatternFlowCfmDmmReservedForDmmReceiverCounter, set in PatternFlowCfmDmmReservedForDmmReceiver.
	// PatternFlowCfmDmmReservedForDmmReceiverCounter is integer counter pattern
	Increment() PatternFlowCfmDmmReservedForDmmReceiverCounter
	// SetIncrement assigns PatternFlowCfmDmmReservedForDmmReceiverCounter provided by user to PatternFlowCfmDmmReservedForDmmReceiver.
	// PatternFlowCfmDmmReservedForDmmReceiverCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiver
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmmReservedForDmmReceiver
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmmReservedForDmmReceiverCounter, set in PatternFlowCfmDmmReservedForDmmReceiver.
	// PatternFlowCfmDmmReservedForDmmReceiverCounter is integer counter pattern
	Decrement() PatternFlowCfmDmmReservedForDmmReceiverCounter
	// SetDecrement assigns PatternFlowCfmDmmReservedForDmmReceiverCounter provided by user to PatternFlowCfmDmmReservedForDmmReceiver.
	// PatternFlowCfmDmmReservedForDmmReceiverCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiver
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmmReservedForDmmReceiver
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum string

// Enum of Choice on PatternFlowCfmDmmReservedForDmmReceiver
var PatternFlowCfmDmmReservedForDmmReceiverChoice = struct {
	VALUE     PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum
	VALUES    PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum
	INCREMENT PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum
	DECREMENT PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum
}{
	VALUE:     PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) Choice() PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum {
	return PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmmReservedForDmmReceiver) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) setChoice(value PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum) PatternFlowCfmDmmReservedForDmmReceiver {
	intValue, ok := otg.PatternFlowCfmDmmReservedForDmmReceiver_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmmReservedForDmmReceiver_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmmReservedForDmmReceiverChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmmReservedForDmmReceiverCounter().msg()
	}

	if value == PatternFlowCfmDmmReservedForDmmReceiverChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmmReservedForDmmReceiverCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiver) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiver) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmmReservedForDmmReceiver object
func (obj *patternFlowCfmDmmReservedForDmmReceiver) SetValue(value uint64) PatternFlowCfmDmmReservedForDmmReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmmReservedForDmmReceiver) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmmReservedForDmmReceiver object
func (obj *patternFlowCfmDmmReservedForDmmReceiver) SetValues(value []uint64) PatternFlowCfmDmmReservedForDmmReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmmReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmmReceiver) Increment() PatternFlowCfmDmmReservedForDmmReceiverCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmmReservedForDmmReceiverCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmmReservedForDmmReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmmReceiver) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmmReservedForDmmReceiverCounter value in the PatternFlowCfmDmmReservedForDmmReceiver object
func (obj *patternFlowCfmDmmReservedForDmmReceiver) SetIncrement(value PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmmReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmmReceiver) Decrement() PatternFlowCfmDmmReservedForDmmReceiverCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmmReservedForDmmReceiverCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmReservedForDmmReceiverCounter
func (obj *patternFlowCfmDmmReservedForDmmReceiver) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmmReservedForDmmReceiverCounter value in the PatternFlowCfmDmmReservedForDmmReceiver object
func (obj *patternFlowCfmDmmReservedForDmmReceiver) SetDecrement(value PatternFlowCfmDmmReservedForDmmReceiverCounter) PatternFlowCfmDmmReservedForDmmReceiver {
	obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmmReservedForDmmReceiver) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmmReservedForDmmReceiver) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmmReservedForDmmReceiverChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmmReceiverChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmReservedForDmmReceiverChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmmReservedForDmmReceiverChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmmReservedForDmmReceiver")
			}
		} else {
			intVal := otg.PatternFlowCfmDmmReservedForDmmReceiver_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmmReservedForDmmReceiver_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
