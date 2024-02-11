package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowUdpLength *****
type patternFlowUdpLength struct {
	validation
	obj              *otg.PatternFlowUdpLength
	marshaller       marshalPatternFlowUdpLength
	unMarshaller     unMarshalPatternFlowUdpLength
	incrementHolder  PatternFlowUdpLengthCounter
	decrementHolder  PatternFlowUdpLengthCounter
	metricTagsHolder PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
}

func NewPatternFlowUdpLength() PatternFlowUdpLength {
	obj := patternFlowUdpLength{obj: &otg.PatternFlowUdpLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowUdpLength) msg() *otg.PatternFlowUdpLength {
	return obj.obj
}

func (obj *patternFlowUdpLength) setMsg(msg *otg.PatternFlowUdpLength) PatternFlowUdpLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowUdpLength struct {
	obj *patternFlowUdpLength
}

type marshalPatternFlowUdpLength interface {
	// ToProto marshals PatternFlowUdpLength to protobuf object *otg.PatternFlowUdpLength
	ToProto() (*otg.PatternFlowUdpLength, error)
	// ToPbText marshals PatternFlowUdpLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowUdpLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowUdpLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowUdpLength struct {
	obj *patternFlowUdpLength
}

type unMarshalPatternFlowUdpLength interface {
	// FromProto unmarshals PatternFlowUdpLength from protobuf object *otg.PatternFlowUdpLength
	FromProto(msg *otg.PatternFlowUdpLength) (PatternFlowUdpLength, error)
	// FromPbText unmarshals PatternFlowUdpLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowUdpLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowUdpLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowUdpLength) Marshal() marshalPatternFlowUdpLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowUdpLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowUdpLength) Unmarshal() unMarshalPatternFlowUdpLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowUdpLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowUdpLength) ToProto() (*otg.PatternFlowUdpLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowUdpLength) FromProto(msg *otg.PatternFlowUdpLength) (PatternFlowUdpLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowUdpLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowUdpLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowUdpLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowUdpLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowUdpLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowUdpLength) FromJson(value string) error {
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

func (obj *patternFlowUdpLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowUdpLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowUdpLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowUdpLength) Clone() (PatternFlowUdpLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowUdpLength()
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

func (obj *patternFlowUdpLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowUdpLength is length
type PatternFlowUdpLength interface {
	Validation
	// msg marshals PatternFlowUdpLength to protobuf object *otg.PatternFlowUdpLength
	// and doesn't set defaults
	msg() *otg.PatternFlowUdpLength
	// setMsg unmarshals PatternFlowUdpLength from protobuf object *otg.PatternFlowUdpLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowUdpLength) PatternFlowUdpLength
	// provides marshal interface
	Marshal() marshalPatternFlowUdpLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowUdpLength
	// validate validates PatternFlowUdpLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowUdpLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowUdpLengthChoiceEnum, set in PatternFlowUdpLength
	Choice() PatternFlowUdpLengthChoiceEnum
	// setChoice assigns PatternFlowUdpLengthChoiceEnum provided by user to PatternFlowUdpLength
	setChoice(value PatternFlowUdpLengthChoiceEnum) PatternFlowUdpLength
	// HasChoice checks if Choice has been set in PatternFlowUdpLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowUdpLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowUdpLength
	SetValue(value uint32) PatternFlowUdpLength
	// HasValue checks if Value has been set in PatternFlowUdpLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowUdpLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowUdpLength
	SetValues(value []uint32) PatternFlowUdpLength
	// Increment returns PatternFlowUdpLengthCounter, set in PatternFlowUdpLength.
	// PatternFlowUdpLengthCounter is integer counter pattern
	Increment() PatternFlowUdpLengthCounter
	// SetIncrement assigns PatternFlowUdpLengthCounter provided by user to PatternFlowUdpLength.
	// PatternFlowUdpLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowUdpLengthCounter) PatternFlowUdpLength
	// HasIncrement checks if Increment has been set in PatternFlowUdpLength
	HasIncrement() bool
	// Decrement returns PatternFlowUdpLengthCounter, set in PatternFlowUdpLength.
	// PatternFlowUdpLengthCounter is integer counter pattern
	Decrement() PatternFlowUdpLengthCounter
	// SetDecrement assigns PatternFlowUdpLengthCounter provided by user to PatternFlowUdpLength.
	// PatternFlowUdpLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowUdpLengthCounter) PatternFlowUdpLength
	// HasDecrement checks if Decrement has been set in PatternFlowUdpLength
	HasDecrement() bool
	// MetricTags returns PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIterIter, set in PatternFlowUdpLength
	MetricTags() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	setNil()
}

type PatternFlowUdpLengthChoiceEnum string

// Enum of Choice on PatternFlowUdpLength
var PatternFlowUdpLengthChoice = struct {
	VALUE     PatternFlowUdpLengthChoiceEnum
	VALUES    PatternFlowUdpLengthChoiceEnum
	INCREMENT PatternFlowUdpLengthChoiceEnum
	DECREMENT PatternFlowUdpLengthChoiceEnum
}{
	VALUE:     PatternFlowUdpLengthChoiceEnum("value"),
	VALUES:    PatternFlowUdpLengthChoiceEnum("values"),
	INCREMENT: PatternFlowUdpLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowUdpLengthChoiceEnum("decrement"),
}

func (obj *patternFlowUdpLength) Choice() PatternFlowUdpLengthChoiceEnum {
	return PatternFlowUdpLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowUdpLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowUdpLength) setChoice(value PatternFlowUdpLengthChoiceEnum) PatternFlowUdpLength {
	intValue, ok := otg.PatternFlowUdpLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowUdpLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowUdpLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowUdpLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowUdpLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowUdpLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowUdpLengthCounter().msg()
	}

	if value == PatternFlowUdpLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowUdpLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowUdpLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowUdpLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowUdpLength object
func (obj *patternFlowUdpLength) SetValue(value uint32) PatternFlowUdpLength {
	obj.setChoice(PatternFlowUdpLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowUdpLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowUdpLength object
func (obj *patternFlowUdpLength) SetValues(value []uint32) PatternFlowUdpLength {
	obj.setChoice(PatternFlowUdpLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowUdpLengthCounter
func (obj *patternFlowUdpLength) Increment() PatternFlowUdpLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowUdpLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowUdpLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowUdpLengthCounter
func (obj *patternFlowUdpLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowUdpLengthCounter value in the PatternFlowUdpLength object
func (obj *patternFlowUdpLength) SetIncrement(value PatternFlowUdpLengthCounter) PatternFlowUdpLength {
	obj.setChoice(PatternFlowUdpLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowUdpLengthCounter
func (obj *patternFlowUdpLength) Decrement() PatternFlowUdpLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowUdpLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowUdpLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowUdpLengthCounter
func (obj *patternFlowUdpLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowUdpLengthCounter value in the PatternFlowUdpLength object
func (obj *patternFlowUdpLength) SetDecrement(value PatternFlowUdpLengthCounter) PatternFlowUdpLength {
	obj.setChoice(PatternFlowUdpLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowUdpLengthMetricTag
func (obj *patternFlowUdpLength) MetricTags() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowUdpLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter struct {
	obj                                *patternFlowUdpLength
	patternFlowUdpLengthMetricTagSlice []PatternFlowUdpLengthMetricTag
	fieldPtr                           *[]*otg.PatternFlowUdpLengthMetricTag
}

func newPatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter(ptr *[]*otg.PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	return &patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter interface {
	setMsg(*patternFlowUdpLength) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	Items() []PatternFlowUdpLengthMetricTag
	Add() PatternFlowUdpLengthMetricTag
	Append(items ...PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	Set(index int, newObj PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	Clear() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	clearHolderSlice() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
	appendHolderSlice(item PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter
}

func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) setMsg(msg *patternFlowUdpLength) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowUdpLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) Items() []PatternFlowUdpLengthMetricTag {
	return obj.patternFlowUdpLengthMetricTagSlice
}

func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) Add() PatternFlowUdpLengthMetricTag {
	newObj := &otg.PatternFlowUdpLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowUdpLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowUdpLengthMetricTagSlice = append(obj.patternFlowUdpLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) Append(items ...PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowUdpLengthMetricTagSlice = append(obj.patternFlowUdpLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) Set(index int, newObj PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowUdpLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) Clear() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowUdpLengthMetricTag{}
		obj.patternFlowUdpLengthMetricTagSlice = []PatternFlowUdpLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) clearHolderSlice() PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	if len(obj.patternFlowUdpLengthMetricTagSlice) > 0 {
		obj.patternFlowUdpLengthMetricTagSlice = []PatternFlowUdpLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowUdpLengthPatternFlowUdpLengthMetricTagIter) appendHolderSlice(item PatternFlowUdpLengthMetricTag) PatternFlowUdpLengthPatternFlowUdpLengthMetricTagIter {
	obj.patternFlowUdpLengthMetricTagSlice = append(obj.patternFlowUdpLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowUdpLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowUdpLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowUdpLength.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowUdpLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowUdpLength) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowUdpLengthChoice.VALUE)

	}

}
