package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareLength *****
type patternFlowArpHardwareLength struct {
	validation
	obj              *otg.PatternFlowArpHardwareLength
	marshaller       marshalPatternFlowArpHardwareLength
	unMarshaller     unMarshalPatternFlowArpHardwareLength
	incrementHolder  PatternFlowArpHardwareLengthCounter
	decrementHolder  PatternFlowArpHardwareLengthCounter
	metricTagsHolder PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
}

func NewPatternFlowArpHardwareLength() PatternFlowArpHardwareLength {
	obj := patternFlowArpHardwareLength{obj: &otg.PatternFlowArpHardwareLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareLength) msg() *otg.PatternFlowArpHardwareLength {
	return obj.obj
}

func (obj *patternFlowArpHardwareLength) setMsg(msg *otg.PatternFlowArpHardwareLength) PatternFlowArpHardwareLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareLength struct {
	obj *patternFlowArpHardwareLength
}

type marshalPatternFlowArpHardwareLength interface {
	// ToProto marshals PatternFlowArpHardwareLength to protobuf object *otg.PatternFlowArpHardwareLength
	ToProto() (*otg.PatternFlowArpHardwareLength, error)
	// ToPbText marshals PatternFlowArpHardwareLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpHardwareLength struct {
	obj *patternFlowArpHardwareLength
}

type unMarshalPatternFlowArpHardwareLength interface {
	// FromProto unmarshals PatternFlowArpHardwareLength from protobuf object *otg.PatternFlowArpHardwareLength
	FromProto(msg *otg.PatternFlowArpHardwareLength) (PatternFlowArpHardwareLength, error)
	// FromPbText unmarshals PatternFlowArpHardwareLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareLength) Marshal() marshalPatternFlowArpHardwareLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareLength) Unmarshal() unMarshalPatternFlowArpHardwareLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareLength) ToProto() (*otg.PatternFlowArpHardwareLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareLength) FromProto(msg *otg.PatternFlowArpHardwareLength) (PatternFlowArpHardwareLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareLength) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareLength) Clone() (PatternFlowArpHardwareLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareLength()
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

func (obj *patternFlowArpHardwareLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpHardwareLength is length (in octets) of a hardware address
type PatternFlowArpHardwareLength interface {
	Validation
	// msg marshals PatternFlowArpHardwareLength to protobuf object *otg.PatternFlowArpHardwareLength
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareLength
	// setMsg unmarshals PatternFlowArpHardwareLength from protobuf object *otg.PatternFlowArpHardwareLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareLength) PatternFlowArpHardwareLength
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareLength
	// validate validates PatternFlowArpHardwareLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpHardwareLengthChoiceEnum, set in PatternFlowArpHardwareLength
	Choice() PatternFlowArpHardwareLengthChoiceEnum
	// setChoice assigns PatternFlowArpHardwareLengthChoiceEnum provided by user to PatternFlowArpHardwareLength
	setChoice(value PatternFlowArpHardwareLengthChoiceEnum) PatternFlowArpHardwareLength
	// HasChoice checks if Choice has been set in PatternFlowArpHardwareLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowArpHardwareLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowArpHardwareLength
	SetValue(value uint32) PatternFlowArpHardwareLength
	// HasValue checks if Value has been set in PatternFlowArpHardwareLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowArpHardwareLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowArpHardwareLength
	SetValues(value []uint32) PatternFlowArpHardwareLength
	// Increment returns PatternFlowArpHardwareLengthCounter, set in PatternFlowArpHardwareLength.
	// PatternFlowArpHardwareLengthCounter is integer counter pattern
	Increment() PatternFlowArpHardwareLengthCounter
	// SetIncrement assigns PatternFlowArpHardwareLengthCounter provided by user to PatternFlowArpHardwareLength.
	// PatternFlowArpHardwareLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLength
	// HasIncrement checks if Increment has been set in PatternFlowArpHardwareLength
	HasIncrement() bool
	// Decrement returns PatternFlowArpHardwareLengthCounter, set in PatternFlowArpHardwareLength.
	// PatternFlowArpHardwareLengthCounter is integer counter pattern
	Decrement() PatternFlowArpHardwareLengthCounter
	// SetDecrement assigns PatternFlowArpHardwareLengthCounter provided by user to PatternFlowArpHardwareLength.
	// PatternFlowArpHardwareLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLength
	// HasDecrement checks if Decrement has been set in PatternFlowArpHardwareLength
	HasDecrement() bool
	// MetricTags returns PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIterIter, set in PatternFlowArpHardwareLength
	MetricTags() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	setNil()
}

type PatternFlowArpHardwareLengthChoiceEnum string

// Enum of Choice on PatternFlowArpHardwareLength
var PatternFlowArpHardwareLengthChoice = struct {
	VALUE     PatternFlowArpHardwareLengthChoiceEnum
	VALUES    PatternFlowArpHardwareLengthChoiceEnum
	INCREMENT PatternFlowArpHardwareLengthChoiceEnum
	DECREMENT PatternFlowArpHardwareLengthChoiceEnum
}{
	VALUE:     PatternFlowArpHardwareLengthChoiceEnum("value"),
	VALUES:    PatternFlowArpHardwareLengthChoiceEnum("values"),
	INCREMENT: PatternFlowArpHardwareLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowArpHardwareLengthChoiceEnum("decrement"),
}

func (obj *patternFlowArpHardwareLength) Choice() PatternFlowArpHardwareLengthChoiceEnum {
	return PatternFlowArpHardwareLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpHardwareLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpHardwareLength) setChoice(value PatternFlowArpHardwareLengthChoiceEnum) PatternFlowArpHardwareLength {
	intValue, ok := otg.PatternFlowArpHardwareLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpHardwareLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpHardwareLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpHardwareLengthChoice.VALUE {
		defaultValue := uint32(6)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpHardwareLengthChoice.VALUES {
		defaultValue := []uint32{6}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpHardwareLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpHardwareLengthCounter().msg()
	}

	if value == PatternFlowArpHardwareLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpHardwareLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpHardwareLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpHardwareLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpHardwareLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowArpHardwareLength object
func (obj *patternFlowArpHardwareLength) SetValue(value uint32) PatternFlowArpHardwareLength {
	obj.setChoice(PatternFlowArpHardwareLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowArpHardwareLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{6})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowArpHardwareLength object
func (obj *patternFlowArpHardwareLength) SetValues(value []uint32) PatternFlowArpHardwareLength {
	obj.setChoice(PatternFlowArpHardwareLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpHardwareLengthCounter
func (obj *patternFlowArpHardwareLength) Increment() PatternFlowArpHardwareLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpHardwareLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpHardwareLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpHardwareLengthCounter
func (obj *patternFlowArpHardwareLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpHardwareLengthCounter value in the PatternFlowArpHardwareLength object
func (obj *patternFlowArpHardwareLength) SetIncrement(value PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLength {
	obj.setChoice(PatternFlowArpHardwareLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpHardwareLengthCounter
func (obj *patternFlowArpHardwareLength) Decrement() PatternFlowArpHardwareLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpHardwareLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpHardwareLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpHardwareLengthCounter
func (obj *patternFlowArpHardwareLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpHardwareLengthCounter value in the PatternFlowArpHardwareLength object
func (obj *patternFlowArpHardwareLength) SetDecrement(value PatternFlowArpHardwareLengthCounter) PatternFlowArpHardwareLength {
	obj.setChoice(PatternFlowArpHardwareLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpHardwareLengthMetricTag
func (obj *patternFlowArpHardwareLength) MetricTags() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpHardwareLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter struct {
	obj                                        *patternFlowArpHardwareLength
	patternFlowArpHardwareLengthMetricTagSlice []PatternFlowArpHardwareLengthMetricTag
	fieldPtr                                   *[]*otg.PatternFlowArpHardwareLengthMetricTag
}

func newPatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter(ptr *[]*otg.PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	return &patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter interface {
	setMsg(*patternFlowArpHardwareLength) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	Items() []PatternFlowArpHardwareLengthMetricTag
	Add() PatternFlowArpHardwareLengthMetricTag
	Append(items ...PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	Set(index int, newObj PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	Clear() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	clearHolderSlice() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
	appendHolderSlice(item PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter
}

func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) setMsg(msg *patternFlowArpHardwareLength) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpHardwareLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) Items() []PatternFlowArpHardwareLengthMetricTag {
	return obj.patternFlowArpHardwareLengthMetricTagSlice
}

func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) Add() PatternFlowArpHardwareLengthMetricTag {
	newObj := &otg.PatternFlowArpHardwareLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpHardwareLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpHardwareLengthMetricTagSlice = append(obj.patternFlowArpHardwareLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) Append(items ...PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpHardwareLengthMetricTagSlice = append(obj.patternFlowArpHardwareLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) Set(index int, newObj PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpHardwareLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) Clear() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpHardwareLengthMetricTag{}
		obj.patternFlowArpHardwareLengthMetricTagSlice = []PatternFlowArpHardwareLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) clearHolderSlice() PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	if len(obj.patternFlowArpHardwareLengthMetricTagSlice) > 0 {
		obj.patternFlowArpHardwareLengthMetricTagSlice = []PatternFlowArpHardwareLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter) appendHolderSlice(item PatternFlowArpHardwareLengthMetricTag) PatternFlowArpHardwareLengthPatternFlowArpHardwareLengthMetricTagIter {
	obj.patternFlowArpHardwareLengthMetricTagSlice = append(obj.patternFlowArpHardwareLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpHardwareLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowArpHardwareLength.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpHardwareLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpHardwareLength) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowArpHardwareLengthChoice.VALUE)

	}

}
