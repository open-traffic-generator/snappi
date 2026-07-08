package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbrTransactionId *****
type patternFlowCfmLbrTransactionId struct {
	validation
	obj             *otg.PatternFlowCfmLbrTransactionId
	marshaller      marshalPatternFlowCfmLbrTransactionId
	unMarshaller    unMarshalPatternFlowCfmLbrTransactionId
	incrementHolder PatternFlowCfmLbrTransactionIdCounter
	decrementHolder PatternFlowCfmLbrTransactionIdCounter
}

func NewPatternFlowCfmLbrTransactionId() PatternFlowCfmLbrTransactionId {
	obj := patternFlowCfmLbrTransactionId{obj: &otg.PatternFlowCfmLbrTransactionId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbrTransactionId) msg() *otg.PatternFlowCfmLbrTransactionId {
	return obj.obj
}

func (obj *patternFlowCfmLbrTransactionId) setMsg(msg *otg.PatternFlowCfmLbrTransactionId) PatternFlowCfmLbrTransactionId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbrTransactionId struct {
	obj *patternFlowCfmLbrTransactionId
}

type marshalPatternFlowCfmLbrTransactionId interface {
	// ToProto marshals PatternFlowCfmLbrTransactionId to protobuf object *otg.PatternFlowCfmLbrTransactionId
	ToProto() (*otg.PatternFlowCfmLbrTransactionId, error)
	// ToPbText marshals PatternFlowCfmLbrTransactionId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbrTransactionId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbrTransactionId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbrTransactionId struct {
	obj *patternFlowCfmLbrTransactionId
}

type unMarshalPatternFlowCfmLbrTransactionId interface {
	// FromProto unmarshals PatternFlowCfmLbrTransactionId from protobuf object *otg.PatternFlowCfmLbrTransactionId
	FromProto(msg *otg.PatternFlowCfmLbrTransactionId) (PatternFlowCfmLbrTransactionId, error)
	// FromPbText unmarshals PatternFlowCfmLbrTransactionId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbrTransactionId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbrTransactionId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbrTransactionId) Marshal() marshalPatternFlowCfmLbrTransactionId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbrTransactionId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbrTransactionId) Unmarshal() unMarshalPatternFlowCfmLbrTransactionId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbrTransactionId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbrTransactionId) ToProto() (*otg.PatternFlowCfmLbrTransactionId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbrTransactionId) FromProto(msg *otg.PatternFlowCfmLbrTransactionId) (PatternFlowCfmLbrTransactionId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbrTransactionId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionId) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbrTransactionId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionId) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbrTransactionId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbrTransactionId) FromJson(value string) error {
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

func (obj *patternFlowCfmLbrTransactionId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrTransactionId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbrTransactionId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbrTransactionId) Clone() (PatternFlowCfmLbrTransactionId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbrTransactionId()
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

func (obj *patternFlowCfmLbrTransactionId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmLbrTransactionId is transaction id or sequence number copied from the LBM to match the reply to the specific request.
type PatternFlowCfmLbrTransactionId interface {
	Validation
	// msg marshals PatternFlowCfmLbrTransactionId to protobuf object *otg.PatternFlowCfmLbrTransactionId
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbrTransactionId
	// setMsg unmarshals PatternFlowCfmLbrTransactionId from protobuf object *otg.PatternFlowCfmLbrTransactionId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbrTransactionId) PatternFlowCfmLbrTransactionId
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbrTransactionId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbrTransactionId
	// validate validates PatternFlowCfmLbrTransactionId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbrTransactionId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmLbrTransactionIdChoiceEnum, set in PatternFlowCfmLbrTransactionId
	Choice() PatternFlowCfmLbrTransactionIdChoiceEnum
	// setChoice assigns PatternFlowCfmLbrTransactionIdChoiceEnum provided by user to PatternFlowCfmLbrTransactionId
	setChoice(value PatternFlowCfmLbrTransactionIdChoiceEnum) PatternFlowCfmLbrTransactionId
	// HasChoice checks if Choice has been set in PatternFlowCfmLbrTransactionId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmLbrTransactionId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmLbrTransactionId
	SetValue(value uint32) PatternFlowCfmLbrTransactionId
	// HasValue checks if Value has been set in PatternFlowCfmLbrTransactionId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmLbrTransactionId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmLbrTransactionId
	SetValues(value []uint32) PatternFlowCfmLbrTransactionId
	// Increment returns PatternFlowCfmLbrTransactionIdCounter, set in PatternFlowCfmLbrTransactionId.
	// PatternFlowCfmLbrTransactionIdCounter is integer counter pattern
	Increment() PatternFlowCfmLbrTransactionIdCounter
	// SetIncrement assigns PatternFlowCfmLbrTransactionIdCounter provided by user to PatternFlowCfmLbrTransactionId.
	// PatternFlowCfmLbrTransactionIdCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionId
	// HasIncrement checks if Increment has been set in PatternFlowCfmLbrTransactionId
	HasIncrement() bool
	// Decrement returns PatternFlowCfmLbrTransactionIdCounter, set in PatternFlowCfmLbrTransactionId.
	// PatternFlowCfmLbrTransactionIdCounter is integer counter pattern
	Decrement() PatternFlowCfmLbrTransactionIdCounter
	// SetDecrement assigns PatternFlowCfmLbrTransactionIdCounter provided by user to PatternFlowCfmLbrTransactionId.
	// PatternFlowCfmLbrTransactionIdCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionId
	// HasDecrement checks if Decrement has been set in PatternFlowCfmLbrTransactionId
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmLbrTransactionIdChoiceEnum string

// Enum of Choice on PatternFlowCfmLbrTransactionId
var PatternFlowCfmLbrTransactionIdChoice = struct {
	VALUE     PatternFlowCfmLbrTransactionIdChoiceEnum
	VALUES    PatternFlowCfmLbrTransactionIdChoiceEnum
	INCREMENT PatternFlowCfmLbrTransactionIdChoiceEnum
	DECREMENT PatternFlowCfmLbrTransactionIdChoiceEnum
}{
	VALUE:     PatternFlowCfmLbrTransactionIdChoiceEnum("value"),
	VALUES:    PatternFlowCfmLbrTransactionIdChoiceEnum("values"),
	INCREMENT: PatternFlowCfmLbrTransactionIdChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmLbrTransactionIdChoiceEnum("decrement"),
}

func (obj *patternFlowCfmLbrTransactionId) Choice() PatternFlowCfmLbrTransactionIdChoiceEnum {
	return PatternFlowCfmLbrTransactionIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmLbrTransactionId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmLbrTransactionId) setChoice(value PatternFlowCfmLbrTransactionIdChoiceEnum) PatternFlowCfmLbrTransactionId {
	intValue, ok := otg.PatternFlowCfmLbrTransactionId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmLbrTransactionIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmLbrTransactionId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmLbrTransactionIdChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmLbrTransactionIdChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmLbrTransactionIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmLbrTransactionIdCounter().msg()
	}

	if value == PatternFlowCfmLbrTransactionIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmLbrTransactionIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbrTransactionId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbrTransactionId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmLbrTransactionId object
func (obj *patternFlowCfmLbrTransactionId) SetValue(value uint32) PatternFlowCfmLbrTransactionId {
	obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmLbrTransactionId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmLbrTransactionId object
func (obj *patternFlowCfmLbrTransactionId) SetValues(value []uint32) PatternFlowCfmLbrTransactionId {
	obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmLbrTransactionIdCounter
func (obj *patternFlowCfmLbrTransactionId) Increment() PatternFlowCfmLbrTransactionIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmLbrTransactionIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmLbrTransactionIdCounter
func (obj *patternFlowCfmLbrTransactionId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmLbrTransactionIdCounter value in the PatternFlowCfmLbrTransactionId object
func (obj *patternFlowCfmLbrTransactionId) SetIncrement(value PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionId {
	obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmLbrTransactionIdCounter
func (obj *patternFlowCfmLbrTransactionId) Decrement() PatternFlowCfmLbrTransactionIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmLbrTransactionIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmLbrTransactionIdCounter
func (obj *patternFlowCfmLbrTransactionId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmLbrTransactionIdCounter value in the PatternFlowCfmLbrTransactionId object
func (obj *patternFlowCfmLbrTransactionId) SetDecrement(value PatternFlowCfmLbrTransactionIdCounter) PatternFlowCfmLbrTransactionId {
	obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmLbrTransactionId) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmLbrTransactionId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmLbrTransactionIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrTransactionIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmLbrTransactionIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrTransactionIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmLbrTransactionIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmLbrTransactionIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmLbrTransactionId")
			}
		} else {
			intVal := otg.PatternFlowCfmLbrTransactionId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmLbrTransactionId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
