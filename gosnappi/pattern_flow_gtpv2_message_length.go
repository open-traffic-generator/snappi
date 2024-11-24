package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageLength *****
type patternFlowGtpv2MessageLength struct {
	validation
	obj              *otg.PatternFlowGtpv2MessageLength
	marshaller       marshalPatternFlowGtpv2MessageLength
	unMarshaller     unMarshalPatternFlowGtpv2MessageLength
	incrementHolder  PatternFlowGtpv2MessageLengthCounter
	decrementHolder  PatternFlowGtpv2MessageLengthCounter
	metricTagsHolder PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
}

func NewPatternFlowGtpv2MessageLength() PatternFlowGtpv2MessageLength {
	obj := patternFlowGtpv2MessageLength{obj: &otg.PatternFlowGtpv2MessageLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageLength) msg() *otg.PatternFlowGtpv2MessageLength {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageLength) setMsg(msg *otg.PatternFlowGtpv2MessageLength) PatternFlowGtpv2MessageLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageLength struct {
	obj *patternFlowGtpv2MessageLength
}

type marshalPatternFlowGtpv2MessageLength interface {
	// ToProto marshals PatternFlowGtpv2MessageLength to protobuf object *otg.PatternFlowGtpv2MessageLength
	ToProto() (*otg.PatternFlowGtpv2MessageLength, error)
	// ToPbText marshals PatternFlowGtpv2MessageLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageLength to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2MessageLength to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2MessageLength struct {
	obj *patternFlowGtpv2MessageLength
}

type unMarshalPatternFlowGtpv2MessageLength interface {
	// FromProto unmarshals PatternFlowGtpv2MessageLength from protobuf object *otg.PatternFlowGtpv2MessageLength
	FromProto(msg *otg.PatternFlowGtpv2MessageLength) (PatternFlowGtpv2MessageLength, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageLength) Marshal() marshalPatternFlowGtpv2MessageLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageLength) Unmarshal() unMarshalPatternFlowGtpv2MessageLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageLength) ToProto() (*otg.PatternFlowGtpv2MessageLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageLength) FromProto(msg *otg.PatternFlowGtpv2MessageLength) (PatternFlowGtpv2MessageLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageLength) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2MessageLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageLength) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageLength) Clone() (PatternFlowGtpv2MessageLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageLength()
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

func (obj *patternFlowGtpv2MessageLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2MessageLength is a 16-bit field that indicates the length of the payload in bytes, excluding the mandatory GTP-c header (first 4 bytes). Includes the TEID and sequence_number if they are present.
type PatternFlowGtpv2MessageLength interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageLength to protobuf object *otg.PatternFlowGtpv2MessageLength
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageLength
	// setMsg unmarshals PatternFlowGtpv2MessageLength from protobuf object *otg.PatternFlowGtpv2MessageLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageLength) PatternFlowGtpv2MessageLength
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageLength
	// validate validates PatternFlowGtpv2MessageLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2MessageLengthChoiceEnum, set in PatternFlowGtpv2MessageLength
	Choice() PatternFlowGtpv2MessageLengthChoiceEnum
	// setChoice assigns PatternFlowGtpv2MessageLengthChoiceEnum provided by user to PatternFlowGtpv2MessageLength
	setChoice(value PatternFlowGtpv2MessageLengthChoiceEnum) PatternFlowGtpv2MessageLength
	// HasChoice checks if Choice has been set in PatternFlowGtpv2MessageLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2MessageLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2MessageLength
	SetValue(value uint32) PatternFlowGtpv2MessageLength
	// HasValue checks if Value has been set in PatternFlowGtpv2MessageLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2MessageLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2MessageLength
	SetValues(value []uint32) PatternFlowGtpv2MessageLength
	// Increment returns PatternFlowGtpv2MessageLengthCounter, set in PatternFlowGtpv2MessageLength.
	// PatternFlowGtpv2MessageLengthCounter is integer counter pattern
	Increment() PatternFlowGtpv2MessageLengthCounter
	// SetIncrement assigns PatternFlowGtpv2MessageLengthCounter provided by user to PatternFlowGtpv2MessageLength.
	// PatternFlowGtpv2MessageLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLength
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2MessageLength
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2MessageLengthCounter, set in PatternFlowGtpv2MessageLength.
	// PatternFlowGtpv2MessageLengthCounter is integer counter pattern
	Decrement() PatternFlowGtpv2MessageLengthCounter
	// SetDecrement assigns PatternFlowGtpv2MessageLengthCounter provided by user to PatternFlowGtpv2MessageLength.
	// PatternFlowGtpv2MessageLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLength
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2MessageLength
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIterIter, set in PatternFlowGtpv2MessageLength
	MetricTags() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	setNil()
}

type PatternFlowGtpv2MessageLengthChoiceEnum string

// Enum of Choice on PatternFlowGtpv2MessageLength
var PatternFlowGtpv2MessageLengthChoice = struct {
	VALUE     PatternFlowGtpv2MessageLengthChoiceEnum
	VALUES    PatternFlowGtpv2MessageLengthChoiceEnum
	INCREMENT PatternFlowGtpv2MessageLengthChoiceEnum
	DECREMENT PatternFlowGtpv2MessageLengthChoiceEnum
}{
	VALUE:     PatternFlowGtpv2MessageLengthChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2MessageLengthChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2MessageLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2MessageLengthChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2MessageLength) Choice() PatternFlowGtpv2MessageLengthChoiceEnum {
	return PatternFlowGtpv2MessageLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2MessageLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2MessageLength) setChoice(value PatternFlowGtpv2MessageLengthChoiceEnum) PatternFlowGtpv2MessageLength {
	intValue, ok := otg.PatternFlowGtpv2MessageLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2MessageLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2MessageLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2MessageLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2MessageLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2MessageLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2MessageLengthCounter().msg()
	}

	if value == PatternFlowGtpv2MessageLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2MessageLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2MessageLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2MessageLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2MessageLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2MessageLength object
func (obj *patternFlowGtpv2MessageLength) SetValue(value uint32) PatternFlowGtpv2MessageLength {
	obj.setChoice(PatternFlowGtpv2MessageLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2MessageLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2MessageLength object
func (obj *patternFlowGtpv2MessageLength) SetValues(value []uint32) PatternFlowGtpv2MessageLength {
	obj.setChoice(PatternFlowGtpv2MessageLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2MessageLengthCounter
func (obj *patternFlowGtpv2MessageLength) Increment() PatternFlowGtpv2MessageLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2MessageLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2MessageLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2MessageLengthCounter
func (obj *patternFlowGtpv2MessageLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2MessageLengthCounter value in the PatternFlowGtpv2MessageLength object
func (obj *patternFlowGtpv2MessageLength) SetIncrement(value PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLength {
	obj.setChoice(PatternFlowGtpv2MessageLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2MessageLengthCounter
func (obj *patternFlowGtpv2MessageLength) Decrement() PatternFlowGtpv2MessageLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2MessageLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2MessageLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2MessageLengthCounter
func (obj *patternFlowGtpv2MessageLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2MessageLengthCounter value in the PatternFlowGtpv2MessageLength object
func (obj *patternFlowGtpv2MessageLength) SetDecrement(value PatternFlowGtpv2MessageLengthCounter) PatternFlowGtpv2MessageLength {
	obj.setChoice(PatternFlowGtpv2MessageLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2MessageLengthMetricTag
func (obj *patternFlowGtpv2MessageLength) MetricTags() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2MessageLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter struct {
	obj                                         *patternFlowGtpv2MessageLength
	patternFlowGtpv2MessageLengthMetricTagSlice []PatternFlowGtpv2MessageLengthMetricTag
	fieldPtr                                    *[]*otg.PatternFlowGtpv2MessageLengthMetricTag
}

func newPatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter(ptr *[]*otg.PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	return &patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter interface {
	setMsg(*patternFlowGtpv2MessageLength) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	Items() []PatternFlowGtpv2MessageLengthMetricTag
	Add() PatternFlowGtpv2MessageLengthMetricTag
	Append(items ...PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	Set(index int, newObj PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	Clear() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	clearHolderSlice() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter
}

func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) setMsg(msg *patternFlowGtpv2MessageLength) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2MessageLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) Items() []PatternFlowGtpv2MessageLengthMetricTag {
	return obj.patternFlowGtpv2MessageLengthMetricTagSlice
}

func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) Add() PatternFlowGtpv2MessageLengthMetricTag {
	newObj := &otg.PatternFlowGtpv2MessageLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2MessageLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2MessageLengthMetricTagSlice = append(obj.patternFlowGtpv2MessageLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) Append(items ...PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2MessageLengthMetricTagSlice = append(obj.patternFlowGtpv2MessageLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) Set(index int, newObj PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2MessageLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) Clear() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2MessageLengthMetricTag{}
		obj.patternFlowGtpv2MessageLengthMetricTagSlice = []PatternFlowGtpv2MessageLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) clearHolderSlice() PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	if len(obj.patternFlowGtpv2MessageLengthMetricTagSlice) > 0 {
		obj.patternFlowGtpv2MessageLengthMetricTagSlice = []PatternFlowGtpv2MessageLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter) appendHolderSlice(item PatternFlowGtpv2MessageLengthMetricTag) PatternFlowGtpv2MessageLengthPatternFlowGtpv2MessageLengthMetricTagIter {
	obj.patternFlowGtpv2MessageLengthMetricTagSlice = append(obj.patternFlowGtpv2MessageLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2MessageLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2MessageLength.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2MessageLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2MessageLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2MessageLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2MessageLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2MessageLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2MessageLength")
			}
		} else {
			intVal := otg.PatternFlowGtpv2MessageLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2MessageLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
