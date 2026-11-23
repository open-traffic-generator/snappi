package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowCfmDmmTxTimestampF *****
type patternFlowCfmDmmTxTimestampF struct {
	validation
	obj             *otg.PatternFlowCfmDmmTxTimestampF
	marshaller      marshalPatternFlowCfmDmmTxTimestampF
	unMarshaller    unMarshalPatternFlowCfmDmmTxTimestampF
	incrementHolder PatternFlowCfmDmmTxTimestampFCounter
	decrementHolder PatternFlowCfmDmmTxTimestampFCounter
}

func NewPatternFlowCfmDmmTxTimestampF() PatternFlowCfmDmmTxTimestampF {
	obj := patternFlowCfmDmmTxTimestampF{obj: &otg.PatternFlowCfmDmmTxTimestampF{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowCfmDmmTxTimestampF) msg() *otg.PatternFlowCfmDmmTxTimestampF {
	return obj.obj
}

func (obj *patternFlowCfmDmmTxTimestampF) setMsg(msg *otg.PatternFlowCfmDmmTxTimestampF) PatternFlowCfmDmmTxTimestampF {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowCfmDmmTxTimestampF struct {
	obj *patternFlowCfmDmmTxTimestampF
}

type marshalPatternFlowCfmDmmTxTimestampF interface {
	// ToProto marshals PatternFlowCfmDmmTxTimestampF to protobuf object *otg.PatternFlowCfmDmmTxTimestampF
	ToProto() (*otg.PatternFlowCfmDmmTxTimestampF, error)
	// ToPbText marshals PatternFlowCfmDmmTxTimestampF to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowCfmDmmTxTimestampF to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowCfmDmmTxTimestampF to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowCfmDmmTxTimestampF struct {
	obj *patternFlowCfmDmmTxTimestampF
}

type unMarshalPatternFlowCfmDmmTxTimestampF interface {
	// FromProto unmarshals PatternFlowCfmDmmTxTimestampF from protobuf object *otg.PatternFlowCfmDmmTxTimestampF
	FromProto(msg *otg.PatternFlowCfmDmmTxTimestampF) (PatternFlowCfmDmmTxTimestampF, error)
	// FromPbText unmarshals PatternFlowCfmDmmTxTimestampF from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowCfmDmmTxTimestampF from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowCfmDmmTxTimestampF from JSON text
	FromJson(value string) error
}

func (obj *patternFlowCfmDmmTxTimestampF) Marshal() marshalPatternFlowCfmDmmTxTimestampF {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowCfmDmmTxTimestampF{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowCfmDmmTxTimestampF) Unmarshal() unMarshalPatternFlowCfmDmmTxTimestampF {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowCfmDmmTxTimestampF{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowCfmDmmTxTimestampF) ToProto() (*otg.PatternFlowCfmDmmTxTimestampF, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowCfmDmmTxTimestampF) FromProto(msg *otg.PatternFlowCfmDmmTxTimestampF) (PatternFlowCfmDmmTxTimestampF, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowCfmDmmTxTimestampF) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampF) FromPbText(value string) error {
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

func (m *marshalpatternFlowCfmDmmTxTimestampF) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampF) FromYaml(value string) error {
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

func (m *marshalpatternFlowCfmDmmTxTimestampF) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowCfmDmmTxTimestampF) FromJson(value string) error {
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

func (obj *patternFlowCfmDmmTxTimestampF) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmTxTimestampF) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowCfmDmmTxTimestampF) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowCfmDmmTxTimestampF) Clone() (PatternFlowCfmDmmTxTimestampF, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowCfmDmmTxTimestampF()
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

func (obj *patternFlowCfmDmmTxTimestampF) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowCfmDmmTxTimestampF is the timestamp at which the DMM frame is transmitted by the originating MEP (ITU-T Y.1731 Section 8.2).
type PatternFlowCfmDmmTxTimestampF interface {
	Validation
	// msg marshals PatternFlowCfmDmmTxTimestampF to protobuf object *otg.PatternFlowCfmDmmTxTimestampF
	// and doesn't set defaults
	msg() *otg.PatternFlowCfmDmmTxTimestampF
	// setMsg unmarshals PatternFlowCfmDmmTxTimestampF from protobuf object *otg.PatternFlowCfmDmmTxTimestampF
	// and doesn't set defaults
	setMsg(*otg.PatternFlowCfmDmmTxTimestampF) PatternFlowCfmDmmTxTimestampF
	// provides marshal interface
	Marshal() marshalPatternFlowCfmDmmTxTimestampF
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowCfmDmmTxTimestampF
	// validate validates PatternFlowCfmDmmTxTimestampF
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowCfmDmmTxTimestampF, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowCfmDmmTxTimestampFChoiceEnum, set in PatternFlowCfmDmmTxTimestampF
	Choice() PatternFlowCfmDmmTxTimestampFChoiceEnum
	// setChoice assigns PatternFlowCfmDmmTxTimestampFChoiceEnum provided by user to PatternFlowCfmDmmTxTimestampF
	setChoice(value PatternFlowCfmDmmTxTimestampFChoiceEnum) PatternFlowCfmDmmTxTimestampF
	// HasChoice checks if Choice has been set in PatternFlowCfmDmmTxTimestampF
	HasChoice() bool
	// Value returns uint64, set in PatternFlowCfmDmmTxTimestampF.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowCfmDmmTxTimestampF
	SetValue(value uint64) PatternFlowCfmDmmTxTimestampF
	// HasValue checks if Value has been set in PatternFlowCfmDmmTxTimestampF
	HasValue() bool
	// Values returns []uint64, set in PatternFlowCfmDmmTxTimestampF.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowCfmDmmTxTimestampF
	SetValues(value []uint64) PatternFlowCfmDmmTxTimestampF
	// Increment returns PatternFlowCfmDmmTxTimestampFCounter, set in PatternFlowCfmDmmTxTimestampF.
	// PatternFlowCfmDmmTxTimestampFCounter is integer counter pattern
	Increment() PatternFlowCfmDmmTxTimestampFCounter
	// SetIncrement assigns PatternFlowCfmDmmTxTimestampFCounter provided by user to PatternFlowCfmDmmTxTimestampF.
	// PatternFlowCfmDmmTxTimestampFCounter is integer counter pattern
	SetIncrement(value PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampF
	// HasIncrement checks if Increment has been set in PatternFlowCfmDmmTxTimestampF
	HasIncrement() bool
	// Decrement returns PatternFlowCfmDmmTxTimestampFCounter, set in PatternFlowCfmDmmTxTimestampF.
	// PatternFlowCfmDmmTxTimestampFCounter is integer counter pattern
	Decrement() PatternFlowCfmDmmTxTimestampFCounter
	// SetDecrement assigns PatternFlowCfmDmmTxTimestampFCounter provided by user to PatternFlowCfmDmmTxTimestampF.
	// PatternFlowCfmDmmTxTimestampFCounter is integer counter pattern
	SetDecrement(value PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampF
	// HasDecrement checks if Decrement has been set in PatternFlowCfmDmmTxTimestampF
	HasDecrement() bool
	setNil()
}

type PatternFlowCfmDmmTxTimestampFChoiceEnum string

// Enum of Choice on PatternFlowCfmDmmTxTimestampF
var PatternFlowCfmDmmTxTimestampFChoice = struct {
	VALUE     PatternFlowCfmDmmTxTimestampFChoiceEnum
	VALUES    PatternFlowCfmDmmTxTimestampFChoiceEnum
	INCREMENT PatternFlowCfmDmmTxTimestampFChoiceEnum
	DECREMENT PatternFlowCfmDmmTxTimestampFChoiceEnum
}{
	VALUE:     PatternFlowCfmDmmTxTimestampFChoiceEnum("value"),
	VALUES:    PatternFlowCfmDmmTxTimestampFChoiceEnum("values"),
	INCREMENT: PatternFlowCfmDmmTxTimestampFChoiceEnum("increment"),
	DECREMENT: PatternFlowCfmDmmTxTimestampFChoiceEnum("decrement"),
}

func (obj *patternFlowCfmDmmTxTimestampF) Choice() PatternFlowCfmDmmTxTimestampFChoiceEnum {
	return PatternFlowCfmDmmTxTimestampFChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowCfmDmmTxTimestampF) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowCfmDmmTxTimestampF) setChoice(value PatternFlowCfmDmmTxTimestampFChoiceEnum) PatternFlowCfmDmmTxTimestampF {
	intValue, ok := otg.PatternFlowCfmDmmTxTimestampF_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowCfmDmmTxTimestampFChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowCfmDmmTxTimestampF_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowCfmDmmTxTimestampFChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowCfmDmmTxTimestampFChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowCfmDmmTxTimestampFChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowCfmDmmTxTimestampFCounter().msg()
	}

	if value == PatternFlowCfmDmmTxTimestampFChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowCfmDmmTxTimestampFCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmTxTimestampF) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowCfmDmmTxTimestampF) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowCfmDmmTxTimestampF object
func (obj *patternFlowCfmDmmTxTimestampF) SetValue(value uint64) PatternFlowCfmDmmTxTimestampF {
	obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowCfmDmmTxTimestampF) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowCfmDmmTxTimestampF object
func (obj *patternFlowCfmDmmTxTimestampF) SetValues(value []uint64) PatternFlowCfmDmmTxTimestampF {
	obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowCfmDmmTxTimestampFCounter
func (obj *patternFlowCfmDmmTxTimestampF) Increment() PatternFlowCfmDmmTxTimestampFCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowCfmDmmTxTimestampFCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowCfmDmmTxTimestampFCounter
func (obj *patternFlowCfmDmmTxTimestampF) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowCfmDmmTxTimestampFCounter value in the PatternFlowCfmDmmTxTimestampF object
func (obj *patternFlowCfmDmmTxTimestampF) SetIncrement(value PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampF {
	obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmTxTimestampFCounter
func (obj *patternFlowCfmDmmTxTimestampF) Decrement() PatternFlowCfmDmmTxTimestampFCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowCfmDmmTxTimestampFCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowCfmDmmTxTimestampFCounter
func (obj *patternFlowCfmDmmTxTimestampF) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowCfmDmmTxTimestampFCounter value in the PatternFlowCfmDmmTxTimestampF object
func (obj *patternFlowCfmDmmTxTimestampF) SetDecrement(value PatternFlowCfmDmmTxTimestampFCounter) PatternFlowCfmDmmTxTimestampF {
	obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowCfmDmmTxTimestampF) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowCfmDmmTxTimestampF) setDefault() {
	var choices_set int = 0
	var choice PatternFlowCfmDmmTxTimestampFChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmTxTimestampFChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowCfmDmmTxTimestampFChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmTxTimestampFChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowCfmDmmTxTimestampFChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowCfmDmmTxTimestampFChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowCfmDmmTxTimestampF")
			}
		} else {
			intVal := otg.PatternFlowCfmDmmTxTimestampF_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowCfmDmmTxTimestampF_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
