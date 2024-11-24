package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpProtocolLength *****
type patternFlowArpProtocolLength struct {
	validation
	obj              *otg.PatternFlowArpProtocolLength
	marshaller       marshalPatternFlowArpProtocolLength
	unMarshaller     unMarshalPatternFlowArpProtocolLength
	incrementHolder  PatternFlowArpProtocolLengthCounter
	decrementHolder  PatternFlowArpProtocolLengthCounter
	metricTagsHolder PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
}

func NewPatternFlowArpProtocolLength() PatternFlowArpProtocolLength {
	obj := patternFlowArpProtocolLength{obj: &otg.PatternFlowArpProtocolLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpProtocolLength) msg() *otg.PatternFlowArpProtocolLength {
	return obj.obj
}

func (obj *patternFlowArpProtocolLength) setMsg(msg *otg.PatternFlowArpProtocolLength) PatternFlowArpProtocolLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpProtocolLength struct {
	obj *patternFlowArpProtocolLength
}

type marshalPatternFlowArpProtocolLength interface {
	// ToProto marshals PatternFlowArpProtocolLength to protobuf object *otg.PatternFlowArpProtocolLength
	ToProto() (*otg.PatternFlowArpProtocolLength, error)
	// ToPbText marshals PatternFlowArpProtocolLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpProtocolLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpProtocolLength to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowArpProtocolLength to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowArpProtocolLength struct {
	obj *patternFlowArpProtocolLength
}

type unMarshalPatternFlowArpProtocolLength interface {
	// FromProto unmarshals PatternFlowArpProtocolLength from protobuf object *otg.PatternFlowArpProtocolLength
	FromProto(msg *otg.PatternFlowArpProtocolLength) (PatternFlowArpProtocolLength, error)
	// FromPbText unmarshals PatternFlowArpProtocolLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpProtocolLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpProtocolLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpProtocolLength) Marshal() marshalPatternFlowArpProtocolLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpProtocolLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpProtocolLength) Unmarshal() unMarshalPatternFlowArpProtocolLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpProtocolLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpProtocolLength) ToProto() (*otg.PatternFlowArpProtocolLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpProtocolLength) FromProto(msg *otg.PatternFlowArpProtocolLength) (PatternFlowArpProtocolLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpProtocolLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpProtocolLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpProtocolLength) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowArpProtocolLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpProtocolLength) FromJson(value string) error {
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

func (obj *patternFlowArpProtocolLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpProtocolLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpProtocolLength) Clone() (PatternFlowArpProtocolLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpProtocolLength()
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

func (obj *patternFlowArpProtocolLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpProtocolLength is length (in octets) of internetwork addresses
type PatternFlowArpProtocolLength interface {
	Validation
	// msg marshals PatternFlowArpProtocolLength to protobuf object *otg.PatternFlowArpProtocolLength
	// and doesn't set defaults
	msg() *otg.PatternFlowArpProtocolLength
	// setMsg unmarshals PatternFlowArpProtocolLength from protobuf object *otg.PatternFlowArpProtocolLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpProtocolLength) PatternFlowArpProtocolLength
	// provides marshal interface
	Marshal() marshalPatternFlowArpProtocolLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpProtocolLength
	// validate validates PatternFlowArpProtocolLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpProtocolLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpProtocolLengthChoiceEnum, set in PatternFlowArpProtocolLength
	Choice() PatternFlowArpProtocolLengthChoiceEnum
	// setChoice assigns PatternFlowArpProtocolLengthChoiceEnum provided by user to PatternFlowArpProtocolLength
	setChoice(value PatternFlowArpProtocolLengthChoiceEnum) PatternFlowArpProtocolLength
	// HasChoice checks if Choice has been set in PatternFlowArpProtocolLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowArpProtocolLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowArpProtocolLength
	SetValue(value uint32) PatternFlowArpProtocolLength
	// HasValue checks if Value has been set in PatternFlowArpProtocolLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowArpProtocolLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowArpProtocolLength
	SetValues(value []uint32) PatternFlowArpProtocolLength
	// Increment returns PatternFlowArpProtocolLengthCounter, set in PatternFlowArpProtocolLength.
	// PatternFlowArpProtocolLengthCounter is integer counter pattern
	Increment() PatternFlowArpProtocolLengthCounter
	// SetIncrement assigns PatternFlowArpProtocolLengthCounter provided by user to PatternFlowArpProtocolLength.
	// PatternFlowArpProtocolLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLength
	// HasIncrement checks if Increment has been set in PatternFlowArpProtocolLength
	HasIncrement() bool
	// Decrement returns PatternFlowArpProtocolLengthCounter, set in PatternFlowArpProtocolLength.
	// PatternFlowArpProtocolLengthCounter is integer counter pattern
	Decrement() PatternFlowArpProtocolLengthCounter
	// SetDecrement assigns PatternFlowArpProtocolLengthCounter provided by user to PatternFlowArpProtocolLength.
	// PatternFlowArpProtocolLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLength
	// HasDecrement checks if Decrement has been set in PatternFlowArpProtocolLength
	HasDecrement() bool
	// MetricTags returns PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIterIter, set in PatternFlowArpProtocolLength
	MetricTags() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	setNil()
}

type PatternFlowArpProtocolLengthChoiceEnum string

// Enum of Choice on PatternFlowArpProtocolLength
var PatternFlowArpProtocolLengthChoice = struct {
	VALUE     PatternFlowArpProtocolLengthChoiceEnum
	VALUES    PatternFlowArpProtocolLengthChoiceEnum
	INCREMENT PatternFlowArpProtocolLengthChoiceEnum
	DECREMENT PatternFlowArpProtocolLengthChoiceEnum
}{
	VALUE:     PatternFlowArpProtocolLengthChoiceEnum("value"),
	VALUES:    PatternFlowArpProtocolLengthChoiceEnum("values"),
	INCREMENT: PatternFlowArpProtocolLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowArpProtocolLengthChoiceEnum("decrement"),
}

func (obj *patternFlowArpProtocolLength) Choice() PatternFlowArpProtocolLengthChoiceEnum {
	return PatternFlowArpProtocolLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpProtocolLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpProtocolLength) setChoice(value PatternFlowArpProtocolLengthChoiceEnum) PatternFlowArpProtocolLength {
	intValue, ok := otg.PatternFlowArpProtocolLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpProtocolLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpProtocolLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpProtocolLengthChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpProtocolLengthChoice.VALUES {
		defaultValue := []uint32{4}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpProtocolLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpProtocolLengthCounter().msg()
	}

	if value == PatternFlowArpProtocolLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpProtocolLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpProtocolLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpProtocolLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpProtocolLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowArpProtocolLength object
func (obj *patternFlowArpProtocolLength) SetValue(value uint32) PatternFlowArpProtocolLength {
	obj.setChoice(PatternFlowArpProtocolLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowArpProtocolLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowArpProtocolLength object
func (obj *patternFlowArpProtocolLength) SetValues(value []uint32) PatternFlowArpProtocolLength {
	obj.setChoice(PatternFlowArpProtocolLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpProtocolLengthCounter
func (obj *patternFlowArpProtocolLength) Increment() PatternFlowArpProtocolLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpProtocolLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpProtocolLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpProtocolLengthCounter
func (obj *patternFlowArpProtocolLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpProtocolLengthCounter value in the PatternFlowArpProtocolLength object
func (obj *patternFlowArpProtocolLength) SetIncrement(value PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLength {
	obj.setChoice(PatternFlowArpProtocolLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpProtocolLengthCounter
func (obj *patternFlowArpProtocolLength) Decrement() PatternFlowArpProtocolLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpProtocolLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpProtocolLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpProtocolLengthCounter
func (obj *patternFlowArpProtocolLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpProtocolLengthCounter value in the PatternFlowArpProtocolLength object
func (obj *patternFlowArpProtocolLength) SetDecrement(value PatternFlowArpProtocolLengthCounter) PatternFlowArpProtocolLength {
	obj.setChoice(PatternFlowArpProtocolLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpProtocolLengthMetricTag
func (obj *patternFlowArpProtocolLength) MetricTags() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpProtocolLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter struct {
	obj                                        *patternFlowArpProtocolLength
	patternFlowArpProtocolLengthMetricTagSlice []PatternFlowArpProtocolLengthMetricTag
	fieldPtr                                   *[]*otg.PatternFlowArpProtocolLengthMetricTag
}

func newPatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter(ptr *[]*otg.PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	return &patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter interface {
	setMsg(*patternFlowArpProtocolLength) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	Items() []PatternFlowArpProtocolLengthMetricTag
	Add() PatternFlowArpProtocolLengthMetricTag
	Append(items ...PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	Set(index int, newObj PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	Clear() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	clearHolderSlice() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
	appendHolderSlice(item PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter
}

func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) setMsg(msg *patternFlowArpProtocolLength) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpProtocolLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) Items() []PatternFlowArpProtocolLengthMetricTag {
	return obj.patternFlowArpProtocolLengthMetricTagSlice
}

func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) Add() PatternFlowArpProtocolLengthMetricTag {
	newObj := &otg.PatternFlowArpProtocolLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpProtocolLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpProtocolLengthMetricTagSlice = append(obj.patternFlowArpProtocolLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) Append(items ...PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpProtocolLengthMetricTagSlice = append(obj.patternFlowArpProtocolLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) Set(index int, newObj PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpProtocolLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) Clear() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpProtocolLengthMetricTag{}
		obj.patternFlowArpProtocolLengthMetricTagSlice = []PatternFlowArpProtocolLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) clearHolderSlice() PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	if len(obj.patternFlowArpProtocolLengthMetricTagSlice) > 0 {
		obj.patternFlowArpProtocolLengthMetricTagSlice = []PatternFlowArpProtocolLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter) appendHolderSlice(item PatternFlowArpProtocolLengthMetricTag) PatternFlowArpProtocolLengthPatternFlowArpProtocolLengthMetricTagIter {
	obj.patternFlowArpProtocolLengthMetricTagSlice = append(obj.patternFlowArpProtocolLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpProtocolLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpProtocolLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowArpProtocolLength.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpProtocolLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpProtocolLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowArpProtocolLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowArpProtocolLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowArpProtocolLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowArpProtocolLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowArpProtocolLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowArpProtocolLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowArpProtocolLength")
			}
		} else {
			intVal := otg.PatternFlowArpProtocolLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowArpProtocolLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
