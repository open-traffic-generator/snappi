package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmrRxTimestampF *****
type patternFlowCfmDmrRxTimestampF struct {
	validation
	obj             *otg.PatternFlowCfmDmrRxTimestampF
	marshaller      marshalPatternFlowCfmDmrRxTimestampF
	unMarshaller    unMarshalPatternFlowCfmDmrRxTimestampF
	incrementHolder PatternFlowCfmDmrRxTimestampFCounter
	decrementHolder PatternFlowCfmDmrRxTimestampFCounter
}

func NewPatternFlowCfmDmrRxTimestampF() PatternFlowCfmDmrRxTimestampF {
	obj := patternFlowCfmDmrRxTimestampF{obj: &otg.PatternFlowCfmDmrRxTimestampF{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmrRxTimestampF) msg() *otg.PatternFlowCfmDmrRxTimestampF {
	return obj.obj
}

func (obj *patternFlowCfmDmrRxTimestampF) setMsg(msg *otg.PatternFlowCfmDmrRxTimestampF) PatternFlowCfmDmrRxTimestampF {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmrRxTimestampF struct {
	obj *patternFlowCfmDmrRxTimestampF
}

type marshalPatternFlowCfmDmrRxTimestampF interface {
	// ToProto marshals PatternFlowCfmDmrRxTimestampF to protobuf object *otg.PatternFlowCfmDmrRxTimestampF
	ToProto() (*otg.PatternFlowCfmDmrRxTimestampF, error)
	// ToPbText marshals PatternFlowCfmDmrRxTimestampF to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmrRxTimestampF to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmrRxTimestampF to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmrRxTimestampF struct {
	obj *patternFlowCfmDmrRxTimestampF
}

type unMarshalPatternFlowCfmDmrRxTimestampF interface {
	// FromProto unmarshals PatternFlowCfmDmrRxTimestampF from protobuf object *otg.PatternFlowCfmDmrRxTimestampF
	FromProto(msg *otg.PatternFlowCfmDmrRxTimestampF) (PatternFlowCfmDmrRxTimestampF, error)
	// FromPbText unmarshals PatternFlowCfmDmrRxTimestampF from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmrRxTimestampF from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmrRxTimestampF from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmrRxTimestampF) Marshal() marshalPatternFlowCfmDmrRxTimestampF {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmrRxTimestampF{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmrRxTimestampF) Unmarshal() unMarshalPatternFlowCfmDmrRxTimestampF {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmrRxTimestampF{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmrRxTimestampF) ToProto() (*otg.PatternFlowCfmDmrRxTimestampF, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmrRxTimestampF) FromProto(msg *otg.PatternFlowCfmDmrRxTimestampF) (PatternFlowCfmDmrRxTimestampF, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmrRxTimestampF) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampF) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmrRxTimestampF) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampF) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmrRxTimestampF) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmrRxTimestampF) FromJson(value string) error {
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

func (obj *patternFlowCfmDmrRxTimestampF) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrRxTimestampF) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmrRxTimestampF) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmrRxTimestampF) Clone() (PatternFlowCfmDmrRxTimestampF, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmrRxTimestampF()
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

func (obj *patternFlowCfmDmrRxTimestampF) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmrRxTimestampF is rxTimeStampf (ITU-T G.8013/Y.1731 clause 9.16): the time the responder received the DMM, in the responder's clock. IEEE 1588 timestamp format as for tx_timestamp_f.
type PatternFlowCfmDmrRxTimestampF interface {
	Validation
	// msg marshals PatternFlowCfmDmrRxTimestampF to protobuf object *otg.PatternFlowCfmDmrRxTimestampF
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmrRxTimestampF
	// setMsg unmarshals PatternFlowCfmDmrRxTimestampF from protobuf object *otg.PatternFlowCfmDmrRxTimestampF
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmrRxTimestampF) PatternFlowCfmDmrRxTimestampF
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmrRxTimestampF
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmrRxTimestampF
	// validate validates PatternFlowCfmDmrRxTimestampF
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmrRxTimestampF, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmrRxTimestampFChoiceEnum, set in PatternFlowCfmDmrRxTimestampF
	Choice() PatternFlowCfmDmrRxTimestampFChoiceEnum
	// setChoice assigns PatternFlowCfmDmrRxTimestampFChoiceEnum provided by user to PatternFlowCfmDmrRxTimestampF
	setChoice(value PatternFlowCfmDmrRxTimestampFChoiceEnum) PatternFlowCfmDmrRxTimestampF
	// HasChoice checks if Choice has been set in PatternFlowCfmDmrRxTimestampF
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmrRxTimestampF.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmrRxTimestampF
	SetValue(value uint64) PatternFlowCfmDmrRxTimestampF
	// HasValue checks if Value has been set in PatternFlowCfmDmrRxTimestampF
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmrRxTimestampF.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmrRxTimestampF
	SetValues(value []uint64) PatternFlowCfmDmrRxTimestampF
	// Increment returns PatternFlowCfmDmrRxTimestampFCounter, set in PatternFlowCfmDmrRxTimestampF.
	// PatternFlowCfmDmrRxTimestampFCounter is integer counter pattern
	Increment() PatternFlowCfmDmrRxTimestampFCounter
	// SetIncrement assigns PatternFlowCfmDmrRxTimestampFCounter provided by user to PatternFlowCfmDmrRxTimestampF.
	// PatternFlowCfmDmrRxTimestampFCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampF
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmrRxTimestampF
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmrRxTimestampFCounter, set in PatternFlowCfmDmrRxTimestampF.
	// PatternFlowCfmDmrRxTimestampFCounter is integer counter pattern
	Decrement() PatternFlowCfmDmrRxTimestampFCounter
	// SetDecrement assigns PatternFlowCfmDmrRxTimestampFCounter provided by user to PatternFlowCfmDmrRxTimestampF.
	// PatternFlowCfmDmrRxTimestampFCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampF
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmrRxTimestampF
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmrRxTimestampFChoiceEnum string

// Enum of Choice on PatternFlowCfmDmrRxTimestampF
var PatternFlowCfmDmrRxTimestampFChoice = struct {
	VALUE     PatternFlowCfmDmrRxTimestampFChoiceEnum
	VALUES    PatternFlowCfmDmrRxTimestampFChoiceEnum
	INCREMENT PatternFlowCfmDmrRxTimestampFChoiceEnum
	DECREMENT PatternFlowCfmDmrRxTimestampFChoiceEnum
}{
	VALUE:     PatternFlowCfmDmrRxTimestampFChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmrRxTimestampFChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmrRxTimestampFChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmrRxTimestampFChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmrRxTimestampF) Choice() PatternFlowCfmDmrRxTimestampFChoiceEnum {
	return PatternFlowCfmDmrRxTimestampFChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmrRxTimestampF) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmrRxTimestampF) setChoice(value PatternFlowCfmDmrRxTimestampFChoiceEnum) PatternFlowCfmDmrRxTimestampF {
	intValue, ok := otg.PatternFlowCfmDmrRxTimestampF_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmrRxTimestampFChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmrRxTimestampF_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmrRxTimestampFChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmrRxTimestampFChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmrRxTimestampFChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmrRxTimestampFCounter().msg()
	}

	if value == PatternFlowCfmDmrRxTimestampFChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmrRxTimestampFCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrRxTimestampF) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmrRxTimestampF) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmrRxTimestampF object
func (obj *patternFlowCfmDmrRxTimestampF) SetValue(value uint64) PatternFlowCfmDmrRxTimestampF {
	obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmrRxTimestampF) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmrRxTimestampF object
func (obj *patternFlowCfmDmrRxTimestampF) SetValues(value []uint64) PatternFlowCfmDmrRxTimestampF {
	obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmrRxTimestampFCounter
func (obj *patternFlowCfmDmrRxTimestampF) Increment() PatternFlowCfmDmrRxTimestampFCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmrRxTimestampFCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmrRxTimestampFCounter
func (obj *patternFlowCfmDmrRxTimestampF) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmrRxTimestampFCounter value in the PatternFlowCfmDmrRxTimestampF object
func (obj *patternFlowCfmDmrRxTimestampF) SetIncrement(value PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampF {
	obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrRxTimestampFCounter
func (obj *patternFlowCfmDmrRxTimestampF) Decrement() PatternFlowCfmDmrRxTimestampFCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmrRxTimestampFCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmrRxTimestampFCounter
func (obj *patternFlowCfmDmrRxTimestampF) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmrRxTimestampFCounter value in the PatternFlowCfmDmrRxTimestampF object
func (obj *patternFlowCfmDmrRxTimestampF) SetDecrement(value PatternFlowCfmDmrRxTimestampFCounter) PatternFlowCfmDmrRxTimestampF {
	obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmrRxTimestampF) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmrRxTimestampF) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmrRxTimestampFChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrRxTimestampFChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmrRxTimestampFChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrRxTimestampFChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmrRxTimestampFChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmrRxTimestampFChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmrRxTimestampF")
			}
		} else {
			intVal := otg.PatternFlowCfmDmrRxTimestampF_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmrRxTimestampF_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
