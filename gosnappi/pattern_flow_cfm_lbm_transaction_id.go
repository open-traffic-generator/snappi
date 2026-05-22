package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmLbmTransactionId *****
type patternFlowCfmLbmTransactionId struct {
	validation
	obj             *otg.PatternFlowCfmLbmTransactionId
	marshaller      marshalPatternFlowCfmLbmTransactionId
	unMarshaller    unMarshalPatternFlowCfmLbmTransactionId
	incrementHolder PatternFlowCfmLbmTransactionIdCounter
	decrementHolder PatternFlowCfmLbmTransactionIdCounter
}

func NewPatternFlowCfmLbmTransactionId() PatternFlowCfmLbmTransactionId {
	obj := patternFlowCfmLbmTransactionId{obj: &otg.PatternFlowCfmLbmTransactionId{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmLbmTransactionId) msg() *otg.PatternFlowCfmLbmTransactionId {
	return obj.obj
}

func (obj *patternFlowCfmLbmTransactionId) setMsg(msg *otg.PatternFlowCfmLbmTransactionId) PatternFlowCfmLbmTransactionId {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmLbmTransactionId struct {
	obj *patternFlowCfmLbmTransactionId
}

type marshalPatternFlowCfmLbmTransactionId interface {
	// ToProto marshals PatternFlowCfmLbmTransactionId to protobuf object *otg.PatternFlowCfmLbmTransactionId
	ToProto() (*otg.PatternFlowCfmLbmTransactionId, error)
	// ToPbText marshals PatternFlowCfmLbmTransactionId to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmLbmTransactionId to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmLbmTransactionId to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmLbmTransactionId struct {
	obj *patternFlowCfmLbmTransactionId
}

type unMarshalPatternFlowCfmLbmTransactionId interface {
	// FromProto unmarshals PatternFlowCfmLbmTransactionId from protobuf object *otg.PatternFlowCfmLbmTransactionId
	FromProto(msg *otg.PatternFlowCfmLbmTransactionId) (PatternFlowCfmLbmTransactionId, error)
	// FromPbText unmarshals PatternFlowCfmLbmTransactionId from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmLbmTransactionId from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmLbmTransactionId from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmLbmTransactionId) Marshal() marshalPatternFlowCfmLbmTransactionId {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmLbmTransactionId{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmLbmTransactionId) Unmarshal() unMarshalPatternFlowCfmLbmTransactionId {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmLbmTransactionId{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmLbmTransactionId) ToProto() (*otg.PatternFlowCfmLbmTransactionId, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmLbmTransactionId) FromProto(msg *otg.PatternFlowCfmLbmTransactionId) (PatternFlowCfmLbmTransactionId, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmLbmTransactionId) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionId) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmLbmTransactionId) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionId) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmLbmTransactionId) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmLbmTransactionId) FromJson(value string) error {
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

func (obj *patternFlowCfmLbmTransactionId) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbmTransactionId) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmLbmTransactionId) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmLbmTransactionId) Clone() (PatternFlowCfmLbmTransactionId, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmLbmTransactionId()
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

func (obj *patternFlowCfmLbmTransactionId) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmLbmTransactionId is transaction identifier copied verbatim into the corresponding LBR, allowing the sender to match the reply to this request (IEEE 802.1Q-2018 Figure 21-5).
type PatternFlowCfmLbmTransactionId interface {
	Validation
	// msg marshals PatternFlowCfmLbmTransactionId to protobuf object *otg.PatternFlowCfmLbmTransactionId
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmLbmTransactionId
	// setMsg unmarshals PatternFlowCfmLbmTransactionId from protobuf object *otg.PatternFlowCfmLbmTransactionId
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmLbmTransactionId) PatternFlowCfmLbmTransactionId
	// provides marshal interface
	Marshal() marshalPatternFlowCfmLbmTransactionId
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmLbmTransactionId
	// validate validates PatternFlowCfmLbmTransactionId
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmLbmTransactionId, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmLbmTransactionIdChoiceEnum, set in PatternFlowCfmLbmTransactionId
	Choice() PatternFlowCfmLbmTransactionIdChoiceEnum
	// setChoice assigns PatternFlowCfmLbmTransactionIdChoiceEnum provided by user to PatternFlowCfmLbmTransactionId
	setChoice(value PatternFlowCfmLbmTransactionIdChoiceEnum) PatternFlowCfmLbmTransactionId
	// HasChoice checks if Choice has been set in PatternFlowCfmLbmTransactionId
	HasChoice() bool
	// Value returns uint32, set in PatternFlowCfmLbmTransactionId.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowCfmLbmTransactionId
	SetValue(value uint32) PatternFlowCfmLbmTransactionId
	// HasValue checks if Value has been set in PatternFlowCfmLbmTransactionId
	HasValue() bool
	// Values returns []uint32, set in PatternFlowCfmLbmTransactionId.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowCfmLbmTransactionId
	SetValues(value []uint32) PatternFlowCfmLbmTransactionId
	// Increment returns PatternFlowCfmLbmTransactionIdCounter, set in PatternFlowCfmLbmTransactionId.
	// PatternFlowCfmLbmTransactionIdCounter is integer counter pattern
	Increment() PatternFlowCfmLbmTransactionIdCounter
	// SetIncrement assigns PatternFlowCfmLbmTransactionIdCounter provided by user to PatternFlowCfmLbmTransactionId.
	// PatternFlowCfmLbmTransactionIdCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionId
	// HasIncrement checks if Increment has been set in PatternFlowCfmLbmTransactionId
	HasIncrement() bool
	// Decrement returns PatternFlowCfmLbmTransactionIdCounter, set in PatternFlowCfmLbmTransactionId.
	// PatternFlowCfmLbmTransactionIdCounter is integer counter pattern
	Decrement() PatternFlowCfmLbmTransactionIdCounter
	// SetDecrement assigns PatternFlowCfmLbmTransactionIdCounter provided by user to PatternFlowCfmLbmTransactionId.
	// PatternFlowCfmLbmTransactionIdCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionId
	// HasDecrement checks if Decrement has been set in PatternFlowCfmLbmTransactionId
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmLbmTransactionIdChoiceEnum string

// Enum of Choice on PatternFlowCfmLbmTransactionId
var PatternFlowCfmLbmTransactionIdChoice = struct {
	VALUE     PatternFlowCfmLbmTransactionIdChoiceEnum
	VALUES    PatternFlowCfmLbmTransactionIdChoiceEnum
	INCREMENT PatternFlowCfmLbmTransactionIdChoiceEnum
	DECREMENT PatternFlowCfmLbmTransactionIdChoiceEnum
}{
	VALUE:     PatternFlowCfmLbmTransactionIdChoiceEnum("value"),
	VALUES:    PatternFlowCfmLbmTransactionIdChoiceEnum("values"),
	INCREMENT: PatternFlowCfmLbmTransactionIdChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmLbmTransactionIdChoiceEnum("decrement"),
}

func (obj *patternFlowCfmLbmTransactionId) Choice() PatternFlowCfmLbmTransactionIdChoiceEnum {
	return PatternFlowCfmLbmTransactionIdChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmLbmTransactionId) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmLbmTransactionId) setChoice(value PatternFlowCfmLbmTransactionIdChoiceEnum) PatternFlowCfmLbmTransactionId {
	intValue, ok := otg.PatternFlowCfmLbmTransactionId_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmLbmTransactionIdChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmLbmTransactionId_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmLbmTransactionIdChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmLbmTransactionIdChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmLbmTransactionIdChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmLbmTransactionIdCounter().msg()
	}

	if value == PatternFlowCfmLbmTransactionIdChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmLbmTransactionIdCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbmTransactionId) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowCfmLbmTransactionId) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowCfmLbmTransactionId object
func (obj *patternFlowCfmLbmTransactionId) SetValue(value uint32) PatternFlowCfmLbmTransactionId {
	obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowCfmLbmTransactionId) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowCfmLbmTransactionId object
func (obj *patternFlowCfmLbmTransactionId) SetValues(value []uint32) PatternFlowCfmLbmTransactionId {
	obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmLbmTransactionIdCounter
func (obj *patternFlowCfmLbmTransactionId) Increment() PatternFlowCfmLbmTransactionIdCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmLbmTransactionIdCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmLbmTransactionIdCounter
func (obj *patternFlowCfmLbmTransactionId) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmLbmTransactionIdCounter value in the PatternFlowCfmLbmTransactionId object
func (obj *patternFlowCfmLbmTransactionId) SetIncrement(value PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionId {
	obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmLbmTransactionIdCounter
func (obj *patternFlowCfmLbmTransactionId) Decrement() PatternFlowCfmLbmTransactionIdCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmLbmTransactionIdCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmLbmTransactionIdCounter
func (obj *patternFlowCfmLbmTransactionId) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmLbmTransactionIdCounter value in the PatternFlowCfmLbmTransactionId object
func (obj *patternFlowCfmLbmTransactionId) SetDecrement(value PatternFlowCfmLbmTransactionIdCounter) PatternFlowCfmLbmTransactionId {
	obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmLbmTransactionId) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmLbmTransactionId) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmLbmTransactionIdChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmLbmTransactionIdChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmLbmTransactionIdChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmLbmTransactionIdChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmLbmTransactionIdChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmLbmTransactionIdChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmLbmTransactionId")
			}
		} else {
			intVal := otg.PatternFlowCfmLbmTransactionId_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmLbmTransactionId_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
