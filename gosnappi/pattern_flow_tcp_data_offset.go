package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowTcpDataOffset *****
type patternFlowTcpDataOffset struct {
	validation
	obj              *otg.PatternFlowTcpDataOffset
	marshaller       marshalPatternFlowTcpDataOffset
	unMarshaller     unMarshalPatternFlowTcpDataOffset
	incrementHolder  PatternFlowTcpDataOffsetCounter
	decrementHolder  PatternFlowTcpDataOffsetCounter
	metricTagsHolder PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
}

func NewPatternFlowTcpDataOffset() PatternFlowTcpDataOffset {
	obj := patternFlowTcpDataOffset{obj: &otg.PatternFlowTcpDataOffset{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowTcpDataOffset) msg() *otg.PatternFlowTcpDataOffset {
	return obj.obj
}

func (obj *patternFlowTcpDataOffset) setMsg(msg *otg.PatternFlowTcpDataOffset) PatternFlowTcpDataOffset {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowTcpDataOffset struct {
	obj *patternFlowTcpDataOffset
}

type marshalPatternFlowTcpDataOffset interface {
	// ToProto marshals PatternFlowTcpDataOffset to protobuf object *otg.PatternFlowTcpDataOffset
	ToProto() (*otg.PatternFlowTcpDataOffset, error)
	// ToPbText marshals PatternFlowTcpDataOffset to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowTcpDataOffset to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowTcpDataOffset to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowTcpDataOffset struct {
	obj *patternFlowTcpDataOffset
}

type unMarshalPatternFlowTcpDataOffset interface {
	// FromProto unmarshals PatternFlowTcpDataOffset from protobuf object *otg.PatternFlowTcpDataOffset
	FromProto(msg *otg.PatternFlowTcpDataOffset) (PatternFlowTcpDataOffset, error)
	// FromPbText unmarshals PatternFlowTcpDataOffset from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowTcpDataOffset from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowTcpDataOffset from JSON text
	FromJson(value string) error
}

func (obj *patternFlowTcpDataOffset) Marshal() marshalPatternFlowTcpDataOffset {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowTcpDataOffset{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowTcpDataOffset) Unmarshal() unMarshalPatternFlowTcpDataOffset {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowTcpDataOffset{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowTcpDataOffset) ToProto() (*otg.PatternFlowTcpDataOffset, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowTcpDataOffset) FromProto(msg *otg.PatternFlowTcpDataOffset) (PatternFlowTcpDataOffset, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowTcpDataOffset) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffset) FromPbText(value string) error {
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

func (m *marshalpatternFlowTcpDataOffset) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffset) FromYaml(value string) error {
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

func (m *marshalpatternFlowTcpDataOffset) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowTcpDataOffset) FromJson(value string) error {
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

func (obj *patternFlowTcpDataOffset) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffset) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowTcpDataOffset) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowTcpDataOffset) Clone() (PatternFlowTcpDataOffset, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowTcpDataOffset()
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

func (obj *patternFlowTcpDataOffset) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowTcpDataOffset is the number of 32 bit words in the TCP header. This indicates where the data begins.
type PatternFlowTcpDataOffset interface {
	Validation
	// msg marshals PatternFlowTcpDataOffset to protobuf object *otg.PatternFlowTcpDataOffset
	// and doesn't set defaults
	msg() *otg.PatternFlowTcpDataOffset
	// setMsg unmarshals PatternFlowTcpDataOffset from protobuf object *otg.PatternFlowTcpDataOffset
	// and doesn't set defaults
	setMsg(*otg.PatternFlowTcpDataOffset) PatternFlowTcpDataOffset
	// provides marshal interface
	Marshal() marshalPatternFlowTcpDataOffset
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowTcpDataOffset
	// validate validates PatternFlowTcpDataOffset
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowTcpDataOffset, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowTcpDataOffsetChoiceEnum, set in PatternFlowTcpDataOffset
	Choice() PatternFlowTcpDataOffsetChoiceEnum
	// setChoice assigns PatternFlowTcpDataOffsetChoiceEnum provided by user to PatternFlowTcpDataOffset
	setChoice(value PatternFlowTcpDataOffsetChoiceEnum) PatternFlowTcpDataOffset
	// HasChoice checks if Choice has been set in PatternFlowTcpDataOffset
	HasChoice() bool
	// Value returns uint32, set in PatternFlowTcpDataOffset.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowTcpDataOffset
	SetValue(value uint32) PatternFlowTcpDataOffset
	// HasValue checks if Value has been set in PatternFlowTcpDataOffset
	HasValue() bool
	// Values returns []uint32, set in PatternFlowTcpDataOffset.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowTcpDataOffset
	SetValues(value []uint32) PatternFlowTcpDataOffset
	// Increment returns PatternFlowTcpDataOffsetCounter, set in PatternFlowTcpDataOffset.
	// PatternFlowTcpDataOffsetCounter is integer counter pattern
	Increment() PatternFlowTcpDataOffsetCounter
	// SetIncrement assigns PatternFlowTcpDataOffsetCounter provided by user to PatternFlowTcpDataOffset.
	// PatternFlowTcpDataOffsetCounter is integer counter pattern
	SetIncrement(value PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffset
	// HasIncrement checks if Increment has been set in PatternFlowTcpDataOffset
	HasIncrement() bool
	// Decrement returns PatternFlowTcpDataOffsetCounter, set in PatternFlowTcpDataOffset.
	// PatternFlowTcpDataOffsetCounter is integer counter pattern
	Decrement() PatternFlowTcpDataOffsetCounter
	// SetDecrement assigns PatternFlowTcpDataOffsetCounter provided by user to PatternFlowTcpDataOffset.
	// PatternFlowTcpDataOffsetCounter is integer counter pattern
	SetDecrement(value PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffset
	// HasDecrement checks if Decrement has been set in PatternFlowTcpDataOffset
	HasDecrement() bool
	// MetricTags returns PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIterIter, set in PatternFlowTcpDataOffset
	MetricTags() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	setNil()
}

type PatternFlowTcpDataOffsetChoiceEnum string

// Enum of Choice on PatternFlowTcpDataOffset
var PatternFlowTcpDataOffsetChoice = struct {
	VALUE     PatternFlowTcpDataOffsetChoiceEnum
	VALUES    PatternFlowTcpDataOffsetChoiceEnum
	INCREMENT PatternFlowTcpDataOffsetChoiceEnum
	DECREMENT PatternFlowTcpDataOffsetChoiceEnum
}{
	VALUE:     PatternFlowTcpDataOffsetChoiceEnum("value"),
	VALUES:    PatternFlowTcpDataOffsetChoiceEnum("values"),
	INCREMENT: PatternFlowTcpDataOffsetChoiceEnum("increment"),
	DECREMENT: PatternFlowTcpDataOffsetChoiceEnum("decrement"),
}

func (obj *patternFlowTcpDataOffset) Choice() PatternFlowTcpDataOffsetChoiceEnum {
	return PatternFlowTcpDataOffsetChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowTcpDataOffset) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowTcpDataOffset) setChoice(value PatternFlowTcpDataOffsetChoiceEnum) PatternFlowTcpDataOffset {
	intValue, ok := otg.PatternFlowTcpDataOffset_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowTcpDataOffsetChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowTcpDataOffset_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowTcpDataOffsetChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowTcpDataOffsetChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowTcpDataOffsetChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowTcpDataOffsetCounter().msg()
	}

	if value == PatternFlowTcpDataOffsetChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowTcpDataOffsetCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpDataOffset) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowTcpDataOffsetChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowTcpDataOffset) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowTcpDataOffset object
func (obj *patternFlowTcpDataOffset) SetValue(value uint32) PatternFlowTcpDataOffset {
	obj.setChoice(PatternFlowTcpDataOffsetChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowTcpDataOffset) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowTcpDataOffset object
func (obj *patternFlowTcpDataOffset) SetValues(value []uint32) PatternFlowTcpDataOffset {
	obj.setChoice(PatternFlowTcpDataOffsetChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowTcpDataOffsetCounter
func (obj *patternFlowTcpDataOffset) Increment() PatternFlowTcpDataOffsetCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowTcpDataOffsetChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowTcpDataOffsetCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowTcpDataOffsetCounter
func (obj *patternFlowTcpDataOffset) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowTcpDataOffsetCounter value in the PatternFlowTcpDataOffset object
func (obj *patternFlowTcpDataOffset) SetIncrement(value PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffset {
	obj.setChoice(PatternFlowTcpDataOffsetChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowTcpDataOffsetCounter
func (obj *patternFlowTcpDataOffset) Decrement() PatternFlowTcpDataOffsetCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowTcpDataOffsetChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowTcpDataOffsetCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowTcpDataOffsetCounter
func (obj *patternFlowTcpDataOffset) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowTcpDataOffsetCounter value in the PatternFlowTcpDataOffset object
func (obj *patternFlowTcpDataOffset) SetDecrement(value PatternFlowTcpDataOffsetCounter) PatternFlowTcpDataOffset {
	obj.setChoice(PatternFlowTcpDataOffsetChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowTcpDataOffsetMetricTag
func (obj *patternFlowTcpDataOffset) MetricTags() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowTcpDataOffsetMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter struct {
	obj                                    *patternFlowTcpDataOffset
	patternFlowTcpDataOffsetMetricTagSlice []PatternFlowTcpDataOffsetMetricTag
	fieldPtr                               *[]*otg.PatternFlowTcpDataOffsetMetricTag
}

func newPatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter(ptr *[]*otg.PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	return &patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter{fieldPtr: ptr}
}

type PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter interface {
	setMsg(*patternFlowTcpDataOffset) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	Items() []PatternFlowTcpDataOffsetMetricTag
	Add() PatternFlowTcpDataOffsetMetricTag
	Append(items ...PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	Set(index int, newObj PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	Clear() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	clearHolderSlice() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
	appendHolderSlice(item PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter
}

func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) setMsg(msg *patternFlowTcpDataOffset) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowTcpDataOffsetMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) Items() []PatternFlowTcpDataOffsetMetricTag {
	return obj.patternFlowTcpDataOffsetMetricTagSlice
}

func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) Add() PatternFlowTcpDataOffsetMetricTag {
	newObj := &otg.PatternFlowTcpDataOffsetMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowTcpDataOffsetMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowTcpDataOffsetMetricTagSlice = append(obj.patternFlowTcpDataOffsetMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) Append(items ...PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowTcpDataOffsetMetricTagSlice = append(obj.patternFlowTcpDataOffsetMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) Set(index int, newObj PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowTcpDataOffsetMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) Clear() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowTcpDataOffsetMetricTag{}
		obj.patternFlowTcpDataOffsetMetricTagSlice = []PatternFlowTcpDataOffsetMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) clearHolderSlice() PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	if len(obj.patternFlowTcpDataOffsetMetricTagSlice) > 0 {
		obj.patternFlowTcpDataOffsetMetricTagSlice = []PatternFlowTcpDataOffsetMetricTag{}
	}
	return obj
}
func (obj *patternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter) appendHolderSlice(item PatternFlowTcpDataOffsetMetricTag) PatternFlowTcpDataOffsetPatternFlowTcpDataOffsetMetricTagIter {
	obj.patternFlowTcpDataOffsetMetricTagSlice = append(obj.patternFlowTcpDataOffsetMetricTagSlice, item)
	return obj
}

func (obj *patternFlowTcpDataOffset) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowTcpDataOffset.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowTcpDataOffset.Values <= 15 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowTcpDataOffsetMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowTcpDataOffset) setDefault() {
	var choices_set int = 0
	var choice PatternFlowTcpDataOffsetChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowTcpDataOffsetChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowTcpDataOffsetChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowTcpDataOffsetChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowTcpDataOffsetChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowTcpDataOffsetChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowTcpDataOffset")
			}
		} else {
			intVal := otg.PatternFlowTcpDataOffset_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowTcpDataOffset_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
