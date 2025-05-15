package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpOperation *****
type patternFlowArpOperation struct {
	validation
	obj              *otg.PatternFlowArpOperation
	marshaller       marshalPatternFlowArpOperation
	unMarshaller     unMarshalPatternFlowArpOperation
	incrementHolder  PatternFlowArpOperationCounter
	decrementHolder  PatternFlowArpOperationCounter
	metricTagsHolder PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
}

func NewPatternFlowArpOperation() PatternFlowArpOperation {
	obj := patternFlowArpOperation{obj: &otg.PatternFlowArpOperation{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpOperation) msg() *otg.PatternFlowArpOperation {
	return obj.obj
}

func (obj *patternFlowArpOperation) setMsg(msg *otg.PatternFlowArpOperation) PatternFlowArpOperation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpOperation struct {
	obj *patternFlowArpOperation
}

type marshalPatternFlowArpOperation interface {
	// ToProto marshals PatternFlowArpOperation to protobuf object *otg.PatternFlowArpOperation
	ToProto() (*otg.PatternFlowArpOperation, error)
	// ToPbText marshals PatternFlowArpOperation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpOperation to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpOperation to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpOperation struct {
	obj *patternFlowArpOperation
}

type unMarshalPatternFlowArpOperation interface {
	// FromProto unmarshals PatternFlowArpOperation from protobuf object *otg.PatternFlowArpOperation
	FromProto(msg *otg.PatternFlowArpOperation) (PatternFlowArpOperation, error)
	// FromPbText unmarshals PatternFlowArpOperation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpOperation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpOperation from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpOperation) Marshal() marshalPatternFlowArpOperation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpOperation{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpOperation) Unmarshal() unMarshalPatternFlowArpOperation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpOperation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpOperation) ToProto() (*otg.PatternFlowArpOperation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpOperation) FromProto(msg *otg.PatternFlowArpOperation) (PatternFlowArpOperation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpOperation) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpOperation) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpOperation) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpOperation) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpOperation) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpOperation) FromJson(value string) error {
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

func (obj *patternFlowArpOperation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpOperation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpOperation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpOperation) Clone() (PatternFlowArpOperation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpOperation()
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

func (obj *patternFlowArpOperation) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpOperation is the operation that the sender is performing
type PatternFlowArpOperation interface {
	Validation
	// msg marshals PatternFlowArpOperation to protobuf object *otg.PatternFlowArpOperation
	// and doesn't set defaults
	msg() *otg.PatternFlowArpOperation
	// setMsg unmarshals PatternFlowArpOperation from protobuf object *otg.PatternFlowArpOperation
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpOperation) PatternFlowArpOperation
	// provides marshal interface
	Marshal() marshalPatternFlowArpOperation
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpOperation
	// validate validates PatternFlowArpOperation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpOperation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpOperationChoiceEnum, set in PatternFlowArpOperation
	Choice() PatternFlowArpOperationChoiceEnum
	// setChoice assigns PatternFlowArpOperationChoiceEnum provided by user to PatternFlowArpOperation
	setChoice(value PatternFlowArpOperationChoiceEnum) PatternFlowArpOperation
	// HasChoice checks if Choice has been set in PatternFlowArpOperation
	HasChoice() bool
	// Value returns uint32, set in PatternFlowArpOperation.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowArpOperation
	SetValue(value uint32) PatternFlowArpOperation
	// HasValue checks if Value has been set in PatternFlowArpOperation
	HasValue() bool
	// Values returns []uint32, set in PatternFlowArpOperation.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowArpOperation
	SetValues(value []uint32) PatternFlowArpOperation
	// Increment returns PatternFlowArpOperationCounter, set in PatternFlowArpOperation.
	// PatternFlowArpOperationCounter is integer counter pattern
	Increment() PatternFlowArpOperationCounter
	// SetIncrement assigns PatternFlowArpOperationCounter provided by user to PatternFlowArpOperation.
	// PatternFlowArpOperationCounter is integer counter pattern
	SetIncrement(value PatternFlowArpOperationCounter) PatternFlowArpOperation
	// HasIncrement checks if Increment has been set in PatternFlowArpOperation
	HasIncrement() bool
	// Decrement returns PatternFlowArpOperationCounter, set in PatternFlowArpOperation.
	// PatternFlowArpOperationCounter is integer counter pattern
	Decrement() PatternFlowArpOperationCounter
	// SetDecrement assigns PatternFlowArpOperationCounter provided by user to PatternFlowArpOperation.
	// PatternFlowArpOperationCounter is integer counter pattern
	SetDecrement(value PatternFlowArpOperationCounter) PatternFlowArpOperation
	// HasDecrement checks if Decrement has been set in PatternFlowArpOperation
	HasDecrement() bool
	// MetricTags returns PatternFlowArpOperationPatternFlowArpOperationMetricTagIterIter, set in PatternFlowArpOperation
	MetricTags() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	setNil()
}

type PatternFlowArpOperationChoiceEnum string

// Enum of Choice on PatternFlowArpOperation
var PatternFlowArpOperationChoice = struct {
	VALUE     PatternFlowArpOperationChoiceEnum
	VALUES    PatternFlowArpOperationChoiceEnum
	INCREMENT PatternFlowArpOperationChoiceEnum
	DECREMENT PatternFlowArpOperationChoiceEnum
}{
	VALUE:     PatternFlowArpOperationChoiceEnum("value"),
	VALUES:    PatternFlowArpOperationChoiceEnum("values"),
	INCREMENT: PatternFlowArpOperationChoiceEnum("increment"),
	DECREMENT: PatternFlowArpOperationChoiceEnum("decrement"),
}

func (obj *patternFlowArpOperation) Choice() PatternFlowArpOperationChoiceEnum {
	return PatternFlowArpOperationChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpOperation) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpOperation) setChoice(value PatternFlowArpOperationChoiceEnum) PatternFlowArpOperation {
	intValue, ok := otg.PatternFlowArpOperation_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpOperationChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpOperation_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpOperationChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpOperationChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpOperationChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpOperationCounter().msg()
	}

	if value == PatternFlowArpOperationChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpOperationCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpOperation) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpOperationChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpOperation) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowArpOperation object
func (obj *patternFlowArpOperation) SetValue(value uint32) PatternFlowArpOperation {
	obj.setChoice(PatternFlowArpOperationChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowArpOperation) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowArpOperation object
func (obj *patternFlowArpOperation) SetValues(value []uint32) PatternFlowArpOperation {
	obj.setChoice(PatternFlowArpOperationChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpOperationCounter
func (obj *patternFlowArpOperation) Increment() PatternFlowArpOperationCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpOperationChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpOperationCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpOperationCounter
func (obj *patternFlowArpOperation) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpOperationCounter value in the PatternFlowArpOperation object
func (obj *patternFlowArpOperation) SetIncrement(value PatternFlowArpOperationCounter) PatternFlowArpOperation {
	obj.setChoice(PatternFlowArpOperationChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpOperationCounter
func (obj *patternFlowArpOperation) Decrement() PatternFlowArpOperationCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpOperationChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpOperationCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpOperationCounter
func (obj *patternFlowArpOperation) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpOperationCounter value in the PatternFlowArpOperation object
func (obj *patternFlowArpOperation) SetDecrement(value PatternFlowArpOperationCounter) PatternFlowArpOperation {
	obj.setChoice(PatternFlowArpOperationChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpOperationMetricTag
func (obj *patternFlowArpOperation) MetricTags() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpOperationMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpOperationPatternFlowArpOperationMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpOperationPatternFlowArpOperationMetricTagIter struct {
	obj                                   *patternFlowArpOperation
	patternFlowArpOperationMetricTagSlice []PatternFlowArpOperationMetricTag
	fieldPtr                              *[]*otg.PatternFlowArpOperationMetricTag
}

func newPatternFlowArpOperationPatternFlowArpOperationMetricTagIter(ptr *[]*otg.PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	return &patternFlowArpOperationPatternFlowArpOperationMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpOperationPatternFlowArpOperationMetricTagIter interface {
	setMsg(*patternFlowArpOperation) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	Items() []PatternFlowArpOperationMetricTag
	Add() PatternFlowArpOperationMetricTag
	Append(items ...PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	Set(index int, newObj PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	Clear() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	clearHolderSlice() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
	appendHolderSlice(item PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter
}

func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) setMsg(msg *patternFlowArpOperation) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpOperationMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) Items() []PatternFlowArpOperationMetricTag {
	return obj.patternFlowArpOperationMetricTagSlice
}

func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) Add() PatternFlowArpOperationMetricTag {
	newObj := &otg.PatternFlowArpOperationMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpOperationMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpOperationMetricTagSlice = append(obj.patternFlowArpOperationMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) Append(items ...PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpOperationMetricTagSlice = append(obj.patternFlowArpOperationMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) Set(index int, newObj PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpOperationMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) Clear() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpOperationMetricTag{}
		obj.patternFlowArpOperationMetricTagSlice = []PatternFlowArpOperationMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) clearHolderSlice() PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	if len(obj.patternFlowArpOperationMetricTagSlice) > 0 {
		obj.patternFlowArpOperationMetricTagSlice = []PatternFlowArpOperationMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpOperationPatternFlowArpOperationMetricTagIter) appendHolderSlice(item PatternFlowArpOperationMetricTag) PatternFlowArpOperationPatternFlowArpOperationMetricTagIter {
	obj.patternFlowArpOperationMetricTagSlice = append(obj.patternFlowArpOperationMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpOperation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpOperation.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowArpOperation.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

	if len(obj.obj.MetricTags) != 0 {

		if set_default {
			obj.MetricTags().clearHolderSlice()
			for _, item := range obj.obj.MetricTags {
				obj.MetricTags().appendHolderSlice(&patternFlowArpOperationMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpOperation) setDefault() {
	var choices_set int = 0
	var choice PatternFlowArpOperationChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowArpOperationChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowArpOperationChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowArpOperationChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowArpOperationChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowArpOperationChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowArpOperation")
			}
		} else {
			intVal := otg.PatternFlowArpOperation_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowArpOperation_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
