package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrReservedForDmrReceiver *****
type patternFlowCfmDmrReservedForDmrReceiver struct {
	validation
	obj             *otg.PatternFlowCfmDmrReservedForDmrReceiver
	marshaller      marshalPatternFlowCfmDmrReservedForDmrReceiver
	unMarshaller    unMarshalPatternFlowCfmDmrReservedForDmrReceiver
	incrementHolder PatternFlowCfmDmrReservedForDmrReceiverCounter
	decrementHolder PatternFlowCfmDmrReservedForDmrReceiverCounter
}

func NewPatternFlowCfmDmrReservedForDmrReceiver() PatternFlowCfmDmrReservedForDmrReceiver {
	obj := patternFlowCfmDmrReservedForDmrReceiver{obj: &otg.PatternFlowCfmDmrReservedForDmrReceiver{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) msg() *otg.PatternFlowCfmDmrReservedForDmrReceiver {
	return obj.obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) setMsg(msg *otg.PatternFlowCfmDmrReservedForDmrReceiver) PatternFlowCfmDmrReservedForDmrReceiver {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrReservedForDmrReceiver struct {
	obj *patternFlowCfmDmrReservedForDmrReceiver
}

type marshalPatternFlowCfmDmrReservedForDmrReceiver interface {
	// ToProto marshals PatternFlowCfmDmrReservedForDmrReceiver to protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiver
	ToProto() (*otg.PatternFlowCfmDmrReservedForDmrReceiver, error)
	// ToPbText marshals PatternFlowCfmDmrReservedForDmrReceiver to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrReservedForDmrReceiver to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrReservedForDmrReceiver to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrReservedForDmrReceiver struct {
	obj *patternFlowCfmDmrReservedForDmrReceiver
}

type unMarshalPatternFlowCfmDmrReservedForDmrReceiver interface {
	// FromProto unmarshals PatternFlowCfmDmrReservedForDmrReceiver from protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiver
	FromProto(msg *otg.PatternFlowCfmDmrReservedForDmrReceiver) (PatternFlowCfmDmrReservedForDmrReceiver, error)
	// FromPbText unmarshals PatternFlowCfmDmrReservedForDmrReceiver from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrReservedForDmrReceiver from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrReservedForDmrReceiver from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) Marshal() marshalPatternFlowCfmDmrReservedForDmrReceiver {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrReservedForDmrReceiver{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) Unmarshal() unMarshalPatternFlowCfmDmrReservedForDmrReceiver {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrReservedForDmrReceiver{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiver) ToProto() (*otg.PatternFlowCfmDmrReservedForDmrReceiver, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiver) FromProto(msg *otg.PatternFlowCfmDmrReservedForDmrReceiver) (PatternFlowCfmDmrReservedForDmrReceiver, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiver) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiver) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiver) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiver) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrReservedForDmrReceiver) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrReservedForDmrReceiver) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrReservedForDmrReceiver) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) Clone() (PatternFlowCfmDmrReservedForDmrReceiver, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrReservedForDmrReceiver()
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

func (obj *patternFlowCfmDmrReservedForDmrReceiver) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmrReservedForDmrReceiver is reserved for the DMR receiving equipment (the originating MEP) to record its receive timestamp when the DMR arrives. Set to zero by the DMR originator (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmrReservedForDmrReceiver interface {
	Validation
	// msg marshals PatternFlowCfmDmrReservedForDmrReceiver to protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiver
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrReservedForDmrReceiver
	// setMsg unmarshals PatternFlowCfmDmrReservedForDmrReceiver from protobuf object *otg.PatternFlowCfmDmrReservedForDmrReceiver
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrReservedForDmrReceiver) PatternFlowCfmDmrReservedForDmrReceiver
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrReservedForDmrReceiver
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrReservedForDmrReceiver
	// validate validates PatternFlowCfmDmrReservedForDmrReceiver
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrReservedForDmrReceiver, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum, set in PatternFlowCfmDmrReservedForDmrReceiver
	Choice() PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum
	// setChoice assigns PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum provided by user to PatternFlowCfmDmrReservedForDmrReceiver
	setChoice(value PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum) PatternFlowCfmDmrReservedForDmrReceiver
	// HasChoice checks if Choice has been set in PatternFlowCfmDmrReservedForDmrReceiver
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmrReservedForDmrReceiver.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmrReservedForDmrReceiver
	SetValue(value uint64) PatternFlowCfmDmrReservedForDmrReceiver
	// HasValue checks if Value has been set in PatternFlowCfmDmrReservedForDmrReceiver
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmrReservedForDmrReceiver.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmrReservedForDmrReceiver
	SetValues(value []uint64) PatternFlowCfmDmrReservedForDmrReceiver
	// Increment returns PatternFlowCfmDmrReservedForDmrReceiverCounter, set in PatternFlowCfmDmrReservedForDmrReceiver.
	// PatternFlowCfmDmrReservedForDmrReceiverCounter is integer counter pattern
	Increment() PatternFlowCfmDmrReservedForDmrReceiverCounter
	// SetIncrement assigns PatternFlowCfmDmrReservedForDmrReceiverCounter provided by user to PatternFlowCfmDmrReservedForDmrReceiver.
	// PatternFlowCfmDmrReservedForDmrReceiverCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiver
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmrReservedForDmrReceiver
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmrReservedForDmrReceiverCounter, set in PatternFlowCfmDmrReservedForDmrReceiver.
	// PatternFlowCfmDmrReservedForDmrReceiverCounter is integer counter pattern
	Decrement() PatternFlowCfmDmrReservedForDmrReceiverCounter
	// SetDecrement assigns PatternFlowCfmDmrReservedForDmrReceiverCounter provided by user to PatternFlowCfmDmrReservedForDmrReceiver.
	// PatternFlowCfmDmrReservedForDmrReceiverCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiver
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmrReservedForDmrReceiver
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum string

// Enum of Choice on PatternFlowCfmDmrReservedForDmrReceiver
var PatternFlowCfmDmrReservedForDmrReceiverChoice = struct {
	VALUE     PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum
	VALUES    PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum
	INCREMENT PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum
	DECREMENT PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum
}{
	VALUE:     PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) Choice() PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum {
	return PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmrReservedForDmrReceiver) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) setChoice(value PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum) PatternFlowCfmDmrReservedForDmrReceiver {
	intValue, ok := otg.PatternFlowCfmDmrReservedForDmrReceiver_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmrReservedForDmrReceiver_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmrReservedForDmrReceiverChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmrReservedForDmrReceiverCounter().msg()
	}

	if value == PatternFlowCfmDmrReservedForDmrReceiverChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmrReservedForDmrReceiverCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiver) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiver) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmrReservedForDmrReceiver object
func (obj *patternFlowCfmDmrReservedForDmrReceiver) SetValue(value uint64) PatternFlowCfmDmrReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmrReservedForDmrReceiver) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmrReservedForDmrReceiver object
func (obj *patternFlowCfmDmrReservedForDmrReceiver) SetValues(value []uint64) PatternFlowCfmDmrReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmrReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmrReservedForDmrReceiver) Increment() PatternFlowCfmDmrReservedForDmrReceiverCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmrReservedForDmrReceiverCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmrReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmrReservedForDmrReceiver) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmrReservedForDmrReceiverCounter value in the PatternFlowCfmDmrReservedForDmrReceiver object
func (obj *patternFlowCfmDmrReservedForDmrReceiver) SetIncrement(value PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmrReservedForDmrReceiver) Decrement() PatternFlowCfmDmrReservedForDmrReceiverCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmrReservedForDmrReceiverCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrReservedForDmrReceiverCounter
func (obj *patternFlowCfmDmrReservedForDmrReceiver) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmrReservedForDmrReceiverCounter value in the PatternFlowCfmDmrReservedForDmrReceiver object
func (obj *patternFlowCfmDmrReservedForDmrReceiver) SetDecrement(value PatternFlowCfmDmrReservedForDmrReceiverCounter) PatternFlowCfmDmrReservedForDmrReceiver {
	obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmrReservedForDmrReceiver) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmrReservedForDmrReceiver) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmrReservedForDmrReceiverChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrReservedForDmrReceiverChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrReservedForDmrReceiverChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmrReservedForDmrReceiverChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmrReservedForDmrReceiver")
			}
		} else {
			intVal := otg.PatternFlowCfmDmrReservedForDmrReceiver_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmrReservedForDmrReceiver_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
