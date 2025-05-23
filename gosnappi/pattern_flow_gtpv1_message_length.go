package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageLength *****
type patternFlowGtpv1MessageLength struct {
	validation
	obj              *otg.PatternFlowGtpv1MessageLength
	marshaller       marshalPatternFlowGtpv1MessageLength
	unMarshaller     unMarshalPatternFlowGtpv1MessageLength
	incrementHolder  PatternFlowGtpv1MessageLengthCounter
	decrementHolder  PatternFlowGtpv1MessageLengthCounter
	metricTagsHolder PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
}

func NewPatternFlowGtpv1MessageLength() PatternFlowGtpv1MessageLength {
	obj := patternFlowGtpv1MessageLength{obj: &otg.PatternFlowGtpv1MessageLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageLength) msg() *otg.PatternFlowGtpv1MessageLength {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageLength) setMsg(msg *otg.PatternFlowGtpv1MessageLength) PatternFlowGtpv1MessageLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageLength struct {
	obj *patternFlowGtpv1MessageLength
}

type marshalPatternFlowGtpv1MessageLength interface {
	// ToProto marshals PatternFlowGtpv1MessageLength to protobuf object *otg.PatternFlowGtpv1MessageLength
	ToProto() (*otg.PatternFlowGtpv1MessageLength, error)
	// ToPbText marshals PatternFlowGtpv1MessageLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageLength to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv1MessageLength to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv1MessageLength struct {
	obj *patternFlowGtpv1MessageLength
}

type unMarshalPatternFlowGtpv1MessageLength interface {
	// FromProto unmarshals PatternFlowGtpv1MessageLength from protobuf object *otg.PatternFlowGtpv1MessageLength
	FromProto(msg *otg.PatternFlowGtpv1MessageLength) (PatternFlowGtpv1MessageLength, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageLength) Marshal() marshalPatternFlowGtpv1MessageLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageLength) Unmarshal() unMarshalPatternFlowGtpv1MessageLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageLength) ToProto() (*otg.PatternFlowGtpv1MessageLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageLength) FromProto(msg *otg.PatternFlowGtpv1MessageLength) (PatternFlowGtpv1MessageLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageLength) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv1MessageLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageLength) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageLength) Clone() (PatternFlowGtpv1MessageLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageLength()
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

func (obj *patternFlowGtpv1MessageLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1MessageLength is the length of the payload (the bytes following the mandatory 8-byte GTP header) in bytes that includes any optional fields
type PatternFlowGtpv1MessageLength interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageLength to protobuf object *otg.PatternFlowGtpv1MessageLength
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageLength
	// setMsg unmarshals PatternFlowGtpv1MessageLength from protobuf object *otg.PatternFlowGtpv1MessageLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageLength) PatternFlowGtpv1MessageLength
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageLength
	// validate validates PatternFlowGtpv1MessageLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1MessageLengthChoiceEnum, set in PatternFlowGtpv1MessageLength
	Choice() PatternFlowGtpv1MessageLengthChoiceEnum
	// setChoice assigns PatternFlowGtpv1MessageLengthChoiceEnum provided by user to PatternFlowGtpv1MessageLength
	setChoice(value PatternFlowGtpv1MessageLengthChoiceEnum) PatternFlowGtpv1MessageLength
	// HasChoice checks if Choice has been set in PatternFlowGtpv1MessageLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1MessageLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1MessageLength
	SetValue(value uint32) PatternFlowGtpv1MessageLength
	// HasValue checks if Value has been set in PatternFlowGtpv1MessageLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1MessageLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1MessageLength
	SetValues(value []uint32) PatternFlowGtpv1MessageLength
	// Increment returns PatternFlowGtpv1MessageLengthCounter, set in PatternFlowGtpv1MessageLength.
	// PatternFlowGtpv1MessageLengthCounter is integer counter pattern
	Increment() PatternFlowGtpv1MessageLengthCounter
	// SetIncrement assigns PatternFlowGtpv1MessageLengthCounter provided by user to PatternFlowGtpv1MessageLength.
	// PatternFlowGtpv1MessageLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLength
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1MessageLength
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1MessageLengthCounter, set in PatternFlowGtpv1MessageLength.
	// PatternFlowGtpv1MessageLengthCounter is integer counter pattern
	Decrement() PatternFlowGtpv1MessageLengthCounter
	// SetDecrement assigns PatternFlowGtpv1MessageLengthCounter provided by user to PatternFlowGtpv1MessageLength.
	// PatternFlowGtpv1MessageLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLength
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1MessageLength
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIterIter, set in PatternFlowGtpv1MessageLength
	MetricTags() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	setNil()
}

type PatternFlowGtpv1MessageLengthChoiceEnum string

// Enum of Choice on PatternFlowGtpv1MessageLength
var PatternFlowGtpv1MessageLengthChoice = struct {
	VALUE     PatternFlowGtpv1MessageLengthChoiceEnum
	VALUES    PatternFlowGtpv1MessageLengthChoiceEnum
	INCREMENT PatternFlowGtpv1MessageLengthChoiceEnum
	DECREMENT PatternFlowGtpv1MessageLengthChoiceEnum
}{
	VALUE:     PatternFlowGtpv1MessageLengthChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1MessageLengthChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1MessageLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1MessageLengthChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1MessageLength) Choice() PatternFlowGtpv1MessageLengthChoiceEnum {
	return PatternFlowGtpv1MessageLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1MessageLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1MessageLength) setChoice(value PatternFlowGtpv1MessageLengthChoiceEnum) PatternFlowGtpv1MessageLength {
	intValue, ok := otg.PatternFlowGtpv1MessageLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1MessageLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1MessageLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1MessageLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1MessageLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1MessageLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1MessageLengthCounter().msg()
	}

	if value == PatternFlowGtpv1MessageLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1MessageLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1MessageLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1MessageLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1MessageLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1MessageLength object
func (obj *patternFlowGtpv1MessageLength) SetValue(value uint32) PatternFlowGtpv1MessageLength {
	obj.setChoice(PatternFlowGtpv1MessageLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1MessageLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1MessageLength object
func (obj *patternFlowGtpv1MessageLength) SetValues(value []uint32) PatternFlowGtpv1MessageLength {
	obj.setChoice(PatternFlowGtpv1MessageLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1MessageLengthCounter
func (obj *patternFlowGtpv1MessageLength) Increment() PatternFlowGtpv1MessageLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1MessageLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1MessageLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1MessageLengthCounter
func (obj *patternFlowGtpv1MessageLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1MessageLengthCounter value in the PatternFlowGtpv1MessageLength object
func (obj *patternFlowGtpv1MessageLength) SetIncrement(value PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLength {
	obj.setChoice(PatternFlowGtpv1MessageLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1MessageLengthCounter
func (obj *patternFlowGtpv1MessageLength) Decrement() PatternFlowGtpv1MessageLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1MessageLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1MessageLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1MessageLengthCounter
func (obj *patternFlowGtpv1MessageLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1MessageLengthCounter value in the PatternFlowGtpv1MessageLength object
func (obj *patternFlowGtpv1MessageLength) SetDecrement(value PatternFlowGtpv1MessageLengthCounter) PatternFlowGtpv1MessageLength {
	obj.setChoice(PatternFlowGtpv1MessageLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1MessageLengthMetricTag
func (obj *patternFlowGtpv1MessageLength) MetricTags() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1MessageLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter struct {
	obj                                         *patternFlowGtpv1MessageLength
	patternFlowGtpv1MessageLengthMetricTagSlice []PatternFlowGtpv1MessageLengthMetricTag
	fieldPtr                                    *[]*otg.PatternFlowGtpv1MessageLengthMetricTag
}

func newPatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter(ptr *[]*otg.PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	return &patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter interface {
	setMsg(*patternFlowGtpv1MessageLength) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	Items() []PatternFlowGtpv1MessageLengthMetricTag
	Add() PatternFlowGtpv1MessageLengthMetricTag
	Append(items ...PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	Set(index int, newObj PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	Clear() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	clearHolderSlice() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter
}

func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) setMsg(msg *patternFlowGtpv1MessageLength) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1MessageLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) Items() []PatternFlowGtpv1MessageLengthMetricTag {
	return obj.patternFlowGtpv1MessageLengthMetricTagSlice
}

func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) Add() PatternFlowGtpv1MessageLengthMetricTag {
	newObj := &otg.PatternFlowGtpv1MessageLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1MessageLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1MessageLengthMetricTagSlice = append(obj.patternFlowGtpv1MessageLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) Append(items ...PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1MessageLengthMetricTagSlice = append(obj.patternFlowGtpv1MessageLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) Set(index int, newObj PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1MessageLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) Clear() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1MessageLengthMetricTag{}
		obj.patternFlowGtpv1MessageLengthMetricTagSlice = []PatternFlowGtpv1MessageLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) clearHolderSlice() PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	if len(obj.patternFlowGtpv1MessageLengthMetricTagSlice) > 0 {
		obj.patternFlowGtpv1MessageLengthMetricTagSlice = []PatternFlowGtpv1MessageLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter) appendHolderSlice(item PatternFlowGtpv1MessageLengthMetricTag) PatternFlowGtpv1MessageLengthPatternFlowGtpv1MessageLengthMetricTagIter {
	obj.patternFlowGtpv1MessageLengthMetricTagSlice = append(obj.patternFlowGtpv1MessageLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1MessageLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1MessageLength.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1MessageLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1MessageLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1MessageLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1MessageLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1MessageLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1MessageLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1MessageLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1MessageLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1MessageLength")
			}
		} else {
			intVal := otg.PatternFlowGtpv1MessageLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1MessageLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
