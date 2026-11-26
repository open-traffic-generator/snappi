package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrTxTimestampB *****
type patternFlowCfmDmrTxTimestampB struct {
	validation
	obj             *otg.PatternFlowCfmDmrTxTimestampB
	marshaller      marshalPatternFlowCfmDmrTxTimestampB
	unMarshaller    unMarshalPatternFlowCfmDmrTxTimestampB
	incrementHolder PatternFlowCfmDmrTxTimestampBCounter
	decrementHolder PatternFlowCfmDmrTxTimestampBCounter
}

func NewPatternFlowCfmDmrTxTimestampB() PatternFlowCfmDmrTxTimestampB {
	obj := patternFlowCfmDmrTxTimestampB{obj: &otg.PatternFlowCfmDmrTxTimestampB{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrTxTimestampB) msg() *otg.PatternFlowCfmDmrTxTimestampB {
	return obj.obj
}

func (obj *patternFlowCfmDmrTxTimestampB) setMsg(msg *otg.PatternFlowCfmDmrTxTimestampB) PatternFlowCfmDmrTxTimestampB {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrTxTimestampB struct {
	obj *patternFlowCfmDmrTxTimestampB
}

type marshalPatternFlowCfmDmrTxTimestampB interface {
	// ToProto marshals PatternFlowCfmDmrTxTimestampB to protobuf object *otg.PatternFlowCfmDmrTxTimestampB
	ToProto() (*otg.PatternFlowCfmDmrTxTimestampB, error)
	// ToPbText marshals PatternFlowCfmDmrTxTimestampB to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrTxTimestampB to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrTxTimestampB to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrTxTimestampB struct {
	obj *patternFlowCfmDmrTxTimestampB
}

type unMarshalPatternFlowCfmDmrTxTimestampB interface {
	// FromProto unmarshals PatternFlowCfmDmrTxTimestampB from protobuf object *otg.PatternFlowCfmDmrTxTimestampB
	FromProto(msg *otg.PatternFlowCfmDmrTxTimestampB) (PatternFlowCfmDmrTxTimestampB, error)
	// FromPbText unmarshals PatternFlowCfmDmrTxTimestampB from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrTxTimestampB from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrTxTimestampB from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrTxTimestampB) Marshal() marshalPatternFlowCfmDmrTxTimestampB {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrTxTimestampB{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrTxTimestampB) Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampB {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrTxTimestampB{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrTxTimestampB) ToProto() (*otg.PatternFlowCfmDmrTxTimestampB, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrTxTimestampB) FromProto(msg *otg.PatternFlowCfmDmrTxTimestampB) (PatternFlowCfmDmrTxTimestampB, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrTxTimestampB) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampB) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampB) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampB) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampB) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampB) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrTxTimestampB) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampB) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampB) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrTxTimestampB) Clone() (PatternFlowCfmDmrTxTimestampB, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrTxTimestampB()
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

func (obj *patternFlowCfmDmrTxTimestampB) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmrTxTimestampB is txTimeStampb (ITU-T G.8013/Y.1731 clause 9.16): the time the responder transmits the DMR, in the responder's clock. IEEE 1588 timestamp format as for tx_timestamp_f. The receiver subtracts tx_timestamp_b - rx_timestamp_f (responder processing time) from the round trip, so a realistic reply has tx_timestamp_b not earlier than rx_timestamp_f.
type PatternFlowCfmDmrTxTimestampB interface {
	Validation
	// msg marshals PatternFlowCfmDmrTxTimestampB to protobuf object *otg.PatternFlowCfmDmrTxTimestampB
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrTxTimestampB
	// setMsg unmarshals PatternFlowCfmDmrTxTimestampB from protobuf object *otg.PatternFlowCfmDmrTxTimestampB
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrTxTimestampB) PatternFlowCfmDmrTxTimestampB
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrTxTimestampB
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampB
	// validate validates PatternFlowCfmDmrTxTimestampB
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrTxTimestampB, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmrTxTimestampBChoiceEnum, set in PatternFlowCfmDmrTxTimestampB
	Choice() PatternFlowCfmDmrTxTimestampBChoiceEnum
	// setChoice assigns PatternFlowCfmDmrTxTimestampBChoiceEnum provided by user to PatternFlowCfmDmrTxTimestampB
	setChoice(value PatternFlowCfmDmrTxTimestampBChoiceEnum) PatternFlowCfmDmrTxTimestampB
	// HasChoice checks if Choice has been set in PatternFlowCfmDmrTxTimestampB
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmrTxTimestampB.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampB
	SetValue(value uint64) PatternFlowCfmDmrTxTimestampB
	// HasValue checks if Value has been set in PatternFlowCfmDmrTxTimestampB
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmrTxTimestampB.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmrTxTimestampB
	SetValues(value []uint64) PatternFlowCfmDmrTxTimestampB
	// Increment returns PatternFlowCfmDmrTxTimestampBCounter, set in PatternFlowCfmDmrTxTimestampB.
	// PatternFlowCfmDmrTxTimestampBCounter is integer counter pattern
	Increment() PatternFlowCfmDmrTxTimestampBCounter
	// SetIncrement assigns PatternFlowCfmDmrTxTimestampBCounter provided by user to PatternFlowCfmDmrTxTimestampB.
	// PatternFlowCfmDmrTxTimestampBCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampB
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmrTxTimestampB
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmrTxTimestampBCounter, set in PatternFlowCfmDmrTxTimestampB.
	// PatternFlowCfmDmrTxTimestampBCounter is integer counter pattern
	Decrement() PatternFlowCfmDmrTxTimestampBCounter
	// SetDecrement assigns PatternFlowCfmDmrTxTimestampBCounter provided by user to PatternFlowCfmDmrTxTimestampB.
	// PatternFlowCfmDmrTxTimestampBCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampB
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmrTxTimestampB
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmrTxTimestampBChoiceEnum string

// Enum of Choice on PatternFlowCfmDmrTxTimestampB
var PatternFlowCfmDmrTxTimestampBChoice = struct {
	VALUE     PatternFlowCfmDmrTxTimestampBChoiceEnum
	VALUES    PatternFlowCfmDmrTxTimestampBChoiceEnum
	INCREMENT PatternFlowCfmDmrTxTimestampBChoiceEnum
	DECREMENT PatternFlowCfmDmrTxTimestampBChoiceEnum
}{
	VALUE:     PatternFlowCfmDmrTxTimestampBChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmrTxTimestampBChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmrTxTimestampBChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmrTxTimestampBChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmrTxTimestampB) Choice() PatternFlowCfmDmrTxTimestampBChoiceEnum {
	return PatternFlowCfmDmrTxTimestampBChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmrTxTimestampB) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmrTxTimestampB) setChoice(value PatternFlowCfmDmrTxTimestampBChoiceEnum) PatternFlowCfmDmrTxTimestampB {
	intValue, ok := otg.PatternFlowCfmDmrTxTimestampB_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmrTxTimestampBChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmrTxTimestampB_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmrTxTimestampBChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmrTxTimestampBChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmrTxTimestampBChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmrTxTimestampBCounter().msg()
	}

	if value == PatternFlowCfmDmrTxTimestampBChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmrTxTimestampBCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrTxTimestampB) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrTxTimestampB) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmrTxTimestampB object
func (obj *patternFlowCfmDmrTxTimestampB) SetValue(value uint64) PatternFlowCfmDmrTxTimestampB {
	obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmrTxTimestampB) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmrTxTimestampB object
func (obj *patternFlowCfmDmrTxTimestampB) SetValues(value []uint64) PatternFlowCfmDmrTxTimestampB {
	obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmrTxTimestampBCounter
func (obj *patternFlowCfmDmrTxTimestampB) Increment() PatternFlowCfmDmrTxTimestampBCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmrTxTimestampBCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmrTxTimestampBCounter
func (obj *patternFlowCfmDmrTxTimestampB) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmrTxTimestampBCounter value in the PatternFlowCfmDmrTxTimestampB object
func (obj *patternFlowCfmDmrTxTimestampB) SetIncrement(value PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampB {
	obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrTxTimestampBCounter
func (obj *patternFlowCfmDmrTxTimestampB) Decrement() PatternFlowCfmDmrTxTimestampBCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmrTxTimestampBCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrTxTimestampBCounter
func (obj *patternFlowCfmDmrTxTimestampB) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmrTxTimestampBCounter value in the PatternFlowCfmDmrTxTimestampB object
func (obj *patternFlowCfmDmrTxTimestampB) SetDecrement(value PatternFlowCfmDmrTxTimestampBCounter) PatternFlowCfmDmrTxTimestampB {
	obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmrTxTimestampB) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmrTxTimestampB) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmrTxTimestampBChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampBChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampBChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampBChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampBChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmrTxTimestampBChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmrTxTimestampB")
			}
		} else {
			intVal := otg.PatternFlowCfmDmrTxTimestampB_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmrTxTimestampB_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
