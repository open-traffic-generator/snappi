package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrTxTimestampF *****
type patternFlowCfmDmrTxTimestampF struct {
	validation
	obj             *otg.PatternFlowCfmDmrTxTimestampF
	marshaller      marshalPatternFlowCfmDmrTxTimestampF
	unMarshaller    unMarshalPatternFlowCfmDmrTxTimestampF
	incrementHolder PatternFlowCfmDmrTxTimestampFCounter
	decrementHolder PatternFlowCfmDmrTxTimestampFCounter
}

func NewPatternFlowCfmDmrTxTimestampF() PatternFlowCfmDmrTxTimestampF {
	obj := patternFlowCfmDmrTxTimestampF{obj: &otg.PatternFlowCfmDmrTxTimestampF{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrTxTimestampF) msg() *otg.PatternFlowCfmDmrTxTimestampF {
	return obj.obj
}

func (obj *patternFlowCfmDmrTxTimestampF) setMsg(msg *otg.PatternFlowCfmDmrTxTimestampF) PatternFlowCfmDmrTxTimestampF {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrTxTimestampF struct {
	obj *patternFlowCfmDmrTxTimestampF
}

type marshalPatternFlowCfmDmrTxTimestampF interface {
	// ToProto marshals PatternFlowCfmDmrTxTimestampF to protobuf object *otg.PatternFlowCfmDmrTxTimestampF
	ToProto() (*otg.PatternFlowCfmDmrTxTimestampF, error)
	// ToPbText marshals PatternFlowCfmDmrTxTimestampF to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrTxTimestampF to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrTxTimestampF to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrTxTimestampF struct {
	obj *patternFlowCfmDmrTxTimestampF
}

type unMarshalPatternFlowCfmDmrTxTimestampF interface {
	// FromProto unmarshals PatternFlowCfmDmrTxTimestampF from protobuf object *otg.PatternFlowCfmDmrTxTimestampF
	FromProto(msg *otg.PatternFlowCfmDmrTxTimestampF) (PatternFlowCfmDmrTxTimestampF, error)
	// FromPbText unmarshals PatternFlowCfmDmrTxTimestampF from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrTxTimestampF from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrTxTimestampF from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrTxTimestampF) Marshal() marshalPatternFlowCfmDmrTxTimestampF {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrTxTimestampF{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrTxTimestampF) Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampF {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrTxTimestampF{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrTxTimestampF) ToProto() (*otg.PatternFlowCfmDmrTxTimestampF, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrTxTimestampF) FromProto(msg *otg.PatternFlowCfmDmrTxTimestampF) (PatternFlowCfmDmrTxTimestampF, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrTxTimestampF) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampF) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampF) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampF) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrTxTimestampF) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrTxTimestampF) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrTxTimestampF) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampF) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrTxTimestampF) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrTxTimestampF) Clone() (PatternFlowCfmDmrTxTimestampF, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrTxTimestampF()
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

func (obj *patternFlowCfmDmrTxTimestampF) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmrTxTimestampF is transmit timestamp recorded by the DMM originator, copied verbatim by the responder from the received DMM into the DMR (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmrTxTimestampF interface {
	Validation
	// msg marshals PatternFlowCfmDmrTxTimestampF to protobuf object *otg.PatternFlowCfmDmrTxTimestampF
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrTxTimestampF
	// setMsg unmarshals PatternFlowCfmDmrTxTimestampF from protobuf object *otg.PatternFlowCfmDmrTxTimestampF
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrTxTimestampF) PatternFlowCfmDmrTxTimestampF
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrTxTimestampF
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrTxTimestampF
	// validate validates PatternFlowCfmDmrTxTimestampF
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrTxTimestampF, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmrTxTimestampFChoiceEnum, set in PatternFlowCfmDmrTxTimestampF
	Choice() PatternFlowCfmDmrTxTimestampFChoiceEnum
	// setChoice assigns PatternFlowCfmDmrTxTimestampFChoiceEnum provided by user to PatternFlowCfmDmrTxTimestampF
	setChoice(value PatternFlowCfmDmrTxTimestampFChoiceEnum) PatternFlowCfmDmrTxTimestampF
	// HasChoice checks if Choice has been set in PatternFlowCfmDmrTxTimestampF
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmrTxTimestampF.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmrTxTimestampF
	SetValue(value uint64) PatternFlowCfmDmrTxTimestampF
	// HasValue checks if Value has been set in PatternFlowCfmDmrTxTimestampF
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmrTxTimestampF.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmrTxTimestampF
	SetValues(value []uint64) PatternFlowCfmDmrTxTimestampF
	// Increment returns PatternFlowCfmDmrTxTimestampFCounter, set in PatternFlowCfmDmrTxTimestampF.
	// PatternFlowCfmDmrTxTimestampFCounter is integer counter pattern
	Increment() PatternFlowCfmDmrTxTimestampFCounter
	// SetIncrement assigns PatternFlowCfmDmrTxTimestampFCounter provided by user to PatternFlowCfmDmrTxTimestampF.
	// PatternFlowCfmDmrTxTimestampFCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampF
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmrTxTimestampF
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmrTxTimestampFCounter, set in PatternFlowCfmDmrTxTimestampF.
	// PatternFlowCfmDmrTxTimestampFCounter is integer counter pattern
	Decrement() PatternFlowCfmDmrTxTimestampFCounter
	// SetDecrement assigns PatternFlowCfmDmrTxTimestampFCounter provided by user to PatternFlowCfmDmrTxTimestampF.
	// PatternFlowCfmDmrTxTimestampFCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampF
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmrTxTimestampF
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmrTxTimestampFChoiceEnum string

// Enum of Choice on PatternFlowCfmDmrTxTimestampF
var PatternFlowCfmDmrTxTimestampFChoice = struct {
	VALUE     PatternFlowCfmDmrTxTimestampFChoiceEnum
	VALUES    PatternFlowCfmDmrTxTimestampFChoiceEnum
	INCREMENT PatternFlowCfmDmrTxTimestampFChoiceEnum
	DECREMENT PatternFlowCfmDmrTxTimestampFChoiceEnum
}{
	VALUE:     PatternFlowCfmDmrTxTimestampFChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmrTxTimestampFChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmrTxTimestampFChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmrTxTimestampFChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmrTxTimestampF) Choice() PatternFlowCfmDmrTxTimestampFChoiceEnum {
	return PatternFlowCfmDmrTxTimestampFChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmrTxTimestampF) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmrTxTimestampF) setChoice(value PatternFlowCfmDmrTxTimestampFChoiceEnum) PatternFlowCfmDmrTxTimestampF {
	intValue, ok := otg.PatternFlowCfmDmrTxTimestampF_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmrTxTimestampFChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmrTxTimestampF_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmrTxTimestampFChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmrTxTimestampFChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmrTxTimestampFChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmrTxTimestampFCounter().msg()
	}

	if value == PatternFlowCfmDmrTxTimestampFChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmrTxTimestampFCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrTxTimestampF) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrTxTimestampF) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmrTxTimestampF object
func (obj *patternFlowCfmDmrTxTimestampF) SetValue(value uint64) PatternFlowCfmDmrTxTimestampF {
	obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmrTxTimestampF) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmrTxTimestampF object
func (obj *patternFlowCfmDmrTxTimestampF) SetValues(value []uint64) PatternFlowCfmDmrTxTimestampF {
	obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmrTxTimestampFCounter
func (obj *patternFlowCfmDmrTxTimestampF) Increment() PatternFlowCfmDmrTxTimestampFCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmrTxTimestampFCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmrTxTimestampFCounter
func (obj *patternFlowCfmDmrTxTimestampF) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmrTxTimestampFCounter value in the PatternFlowCfmDmrTxTimestampF object
func (obj *patternFlowCfmDmrTxTimestampF) SetIncrement(value PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampF {
	obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrTxTimestampFCounter
func (obj *patternFlowCfmDmrTxTimestampF) Decrement() PatternFlowCfmDmrTxTimestampFCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmrTxTimestampFCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrTxTimestampFCounter
func (obj *patternFlowCfmDmrTxTimestampF) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmrTxTimestampFCounter value in the PatternFlowCfmDmrTxTimestampF object
func (obj *patternFlowCfmDmrTxTimestampF) SetDecrement(value PatternFlowCfmDmrTxTimestampFCounter) PatternFlowCfmDmrTxTimestampF {
	obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmrTxTimestampF) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmrTxTimestampF) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmrTxTimestampFChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampFChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampFChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampFChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrTxTimestampFChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmrTxTimestampFChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmrTxTimestampF")
			}
		} else {
			intVal := otg.PatternFlowCfmDmrTxTimestampF_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmrTxTimestampF_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
