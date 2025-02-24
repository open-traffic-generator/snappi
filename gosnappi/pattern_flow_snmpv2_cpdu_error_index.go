package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowSnmpv2CPDUErrorIndex *****
type patternFlowSnmpv2CPDUErrorIndex struct {
	validation
	obj             *otg.PatternFlowSnmpv2CPDUErrorIndex
	marshaller      marshalPatternFlowSnmpv2CPDUErrorIndex
	unMarshaller    unMarshalPatternFlowSnmpv2CPDUErrorIndex
	incrementHolder PatternFlowSnmpv2CPDUErrorIndexCounter
	decrementHolder PatternFlowSnmpv2CPDUErrorIndexCounter
}

func NewPatternFlowSnmpv2CPDUErrorIndex() PatternFlowSnmpv2CPDUErrorIndex {
	obj := patternFlowSnmpv2CPDUErrorIndex{obj: &otg.PatternFlowSnmpv2CPDUErrorIndex{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) msg() *otg.PatternFlowSnmpv2CPDUErrorIndex {
	return obj.obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) setMsg(msg *otg.PatternFlowSnmpv2CPDUErrorIndex) PatternFlowSnmpv2CPDUErrorIndex {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowSnmpv2CPDUErrorIndex struct {
	obj *patternFlowSnmpv2CPDUErrorIndex
}

type marshalPatternFlowSnmpv2CPDUErrorIndex interface {
	// ToProto marshals PatternFlowSnmpv2CPDUErrorIndex to protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndex
	ToProto() (*otg.PatternFlowSnmpv2CPDUErrorIndex, error)
	// ToPbText marshals PatternFlowSnmpv2CPDUErrorIndex to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowSnmpv2CPDUErrorIndex to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowSnmpv2CPDUErrorIndex to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowSnmpv2CPDUErrorIndex to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowSnmpv2CPDUErrorIndex struct {
	obj *patternFlowSnmpv2CPDUErrorIndex
}

type unMarshalPatternFlowSnmpv2CPDUErrorIndex interface {
	// FromProto unmarshals PatternFlowSnmpv2CPDUErrorIndex from protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndex
	FromProto(msg *otg.PatternFlowSnmpv2CPDUErrorIndex) (PatternFlowSnmpv2CPDUErrorIndex, error)
	// FromPbText unmarshals PatternFlowSnmpv2CPDUErrorIndex from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowSnmpv2CPDUErrorIndex from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowSnmpv2CPDUErrorIndex from JSON text
	FromJson(value string) error
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) Marshal() marshalPatternFlowSnmpv2CPDUErrorIndex {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowSnmpv2CPDUErrorIndex{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) Unmarshal() unMarshalPatternFlowSnmpv2CPDUErrorIndex {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowSnmpv2CPDUErrorIndex{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowSnmpv2CPDUErrorIndex) ToProto() (*otg.PatternFlowSnmpv2CPDUErrorIndex, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndex) FromProto(msg *otg.PatternFlowSnmpv2CPDUErrorIndex) (PatternFlowSnmpv2CPDUErrorIndex, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowSnmpv2CPDUErrorIndex) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndex) FromPbText(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndex) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndex) FromYaml(value string) error {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndex) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowSnmpv2CPDUErrorIndex) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowSnmpv2CPDUErrorIndex) FromJson(value string) error {
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

func (obj *patternFlowSnmpv2CPDUErrorIndex) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) Clone() (PatternFlowSnmpv2CPDUErrorIndex, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowSnmpv2CPDUErrorIndex()
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

func (obj *patternFlowSnmpv2CPDUErrorIndex) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowSnmpv2CPDUErrorIndex is when Error Status is non-zero,  this field contains a pointer that specifies which object generated the error.  Always zero in a request.
type PatternFlowSnmpv2CPDUErrorIndex interface {
	Validation
	// msg marshals PatternFlowSnmpv2CPDUErrorIndex to protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndex
	// and doesn't set defaults
	msg() *otg.PatternFlowSnmpv2CPDUErrorIndex
	// setMsg unmarshals PatternFlowSnmpv2CPDUErrorIndex from protobuf object *otg.PatternFlowSnmpv2CPDUErrorIndex
	// and doesn't set defaults
	setMsg(*otg.PatternFlowSnmpv2CPDUErrorIndex) PatternFlowSnmpv2CPDUErrorIndex
	// provides marshal interface
	Marshal() marshalPatternFlowSnmpv2CPDUErrorIndex
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowSnmpv2CPDUErrorIndex
	// validate validates PatternFlowSnmpv2CPDUErrorIndex
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowSnmpv2CPDUErrorIndex, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowSnmpv2CPDUErrorIndexChoiceEnum, set in PatternFlowSnmpv2CPDUErrorIndex
	Choice() PatternFlowSnmpv2CPDUErrorIndexChoiceEnum
	// setChoice assigns PatternFlowSnmpv2CPDUErrorIndexChoiceEnum provided by user to PatternFlowSnmpv2CPDUErrorIndex
	setChoice(value PatternFlowSnmpv2CPDUErrorIndexChoiceEnum) PatternFlowSnmpv2CPDUErrorIndex
	// HasChoice checks if Choice has been set in PatternFlowSnmpv2CPDUErrorIndex
	HasChoice() bool
	// Value returns uint32, set in PatternFlowSnmpv2CPDUErrorIndex.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowSnmpv2CPDUErrorIndex
	SetValue(value uint32) PatternFlowSnmpv2CPDUErrorIndex
	// HasValue checks if Value has been set in PatternFlowSnmpv2CPDUErrorIndex
	HasValue() bool
	// Values returns []uint32, set in PatternFlowSnmpv2CPDUErrorIndex.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowSnmpv2CPDUErrorIndex
	SetValues(value []uint32) PatternFlowSnmpv2CPDUErrorIndex
	// Increment returns PatternFlowSnmpv2CPDUErrorIndexCounter, set in PatternFlowSnmpv2CPDUErrorIndex.
	Increment() PatternFlowSnmpv2CPDUErrorIndexCounter
	// SetIncrement assigns PatternFlowSnmpv2CPDUErrorIndexCounter provided by user to PatternFlowSnmpv2CPDUErrorIndex.
	SetIncrement(value PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndex
	// HasIncrement checks if Increment has been set in PatternFlowSnmpv2CPDUErrorIndex
	HasIncrement() bool
	// Decrement returns PatternFlowSnmpv2CPDUErrorIndexCounter, set in PatternFlowSnmpv2CPDUErrorIndex.
	Decrement() PatternFlowSnmpv2CPDUErrorIndexCounter
	// SetDecrement assigns PatternFlowSnmpv2CPDUErrorIndexCounter provided by user to PatternFlowSnmpv2CPDUErrorIndex.
	SetDecrement(value PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndex
	// HasDecrement checks if Decrement has been set in PatternFlowSnmpv2CPDUErrorIndex
	HasDecrement() bool
	setNil()
}

type PatternFlowSnmpv2CPDUErrorIndexChoiceEnum string

// Enum of Choice on PatternFlowSnmpv2CPDUErrorIndex
var PatternFlowSnmpv2CPDUErrorIndexChoice = struct {
	VALUE     PatternFlowSnmpv2CPDUErrorIndexChoiceEnum
	VALUES    PatternFlowSnmpv2CPDUErrorIndexChoiceEnum
	INCREMENT PatternFlowSnmpv2CPDUErrorIndexChoiceEnum
	DECREMENT PatternFlowSnmpv2CPDUErrorIndexChoiceEnum
}{
	VALUE:     PatternFlowSnmpv2CPDUErrorIndexChoiceEnum("value"),
	VALUES:    PatternFlowSnmpv2CPDUErrorIndexChoiceEnum("values"),
	INCREMENT: PatternFlowSnmpv2CPDUErrorIndexChoiceEnum("increment"),
	DECREMENT: PatternFlowSnmpv2CPDUErrorIndexChoiceEnum("decrement"),
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) Choice() PatternFlowSnmpv2CPDUErrorIndexChoiceEnum {
	return PatternFlowSnmpv2CPDUErrorIndexChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowSnmpv2CPDUErrorIndex) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) setChoice(value PatternFlowSnmpv2CPDUErrorIndexChoiceEnum) PatternFlowSnmpv2CPDUErrorIndex {
	intValue, ok := otg.PatternFlowSnmpv2CPDUErrorIndex_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowSnmpv2CPDUErrorIndexChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowSnmpv2CPDUErrorIndex_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowSnmpv2CPDUErrorIndexChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowSnmpv2CPDUErrorIndexChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowSnmpv2CPDUErrorIndexChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowSnmpv2CPDUErrorIndexCounter().msg()
	}

	if value == PatternFlowSnmpv2CPDUErrorIndexChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowSnmpv2CPDUErrorIndexCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndex) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowSnmpv2CPDUErrorIndex) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowSnmpv2CPDUErrorIndex object
func (obj *patternFlowSnmpv2CPDUErrorIndex) SetValue(value uint32) PatternFlowSnmpv2CPDUErrorIndex {
	obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowSnmpv2CPDUErrorIndex) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowSnmpv2CPDUErrorIndex object
func (obj *patternFlowSnmpv2CPDUErrorIndex) SetValues(value []uint32) PatternFlowSnmpv2CPDUErrorIndex {
	obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CPDUErrorIndexCounter
func (obj *patternFlowSnmpv2CPDUErrorIndex) Increment() PatternFlowSnmpv2CPDUErrorIndexCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowSnmpv2CPDUErrorIndexCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowSnmpv2CPDUErrorIndexCounter
func (obj *patternFlowSnmpv2CPDUErrorIndex) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowSnmpv2CPDUErrorIndexCounter value in the PatternFlowSnmpv2CPDUErrorIndex object
func (obj *patternFlowSnmpv2CPDUErrorIndex) SetIncrement(value PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndex {
	obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CPDUErrorIndexCounter
func (obj *patternFlowSnmpv2CPDUErrorIndex) Decrement() PatternFlowSnmpv2CPDUErrorIndexCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowSnmpv2CPDUErrorIndexCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowSnmpv2CPDUErrorIndexCounter
func (obj *patternFlowSnmpv2CPDUErrorIndex) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowSnmpv2CPDUErrorIndexCounter value in the PatternFlowSnmpv2CPDUErrorIndex object
func (obj *patternFlowSnmpv2CPDUErrorIndex) SetDecrement(value PatternFlowSnmpv2CPDUErrorIndexCounter) PatternFlowSnmpv2CPDUErrorIndex {
	obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternFlowSnmpv2CPDUErrorIndex) validateObj(vObj *validation, set_default bool) {
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

func (obj *patternFlowSnmpv2CPDUErrorIndex) setDefault() {
	var choices_set int = 0
	var choice PatternFlowSnmpv2CPDUErrorIndexChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDUErrorIndexChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDUErrorIndexChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDUErrorIndexChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowSnmpv2CPDUErrorIndexChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowSnmpv2CPDUErrorIndexChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowSnmpv2CPDUErrorIndex")
			}
		} else {
			intVal := otg.PatternFlowSnmpv2CPDUErrorIndex_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowSnmpv2CPDUErrorIndex_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
