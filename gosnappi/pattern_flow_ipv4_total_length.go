package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TotalLength *****
type patternFlowIpv4TotalLength struct {
	validation
	obj              *otg.PatternFlowIpv4TotalLength
	marshaller       marshalPatternFlowIpv4TotalLength
	unMarshaller     unMarshalPatternFlowIpv4TotalLength
	incrementHolder  PatternFlowIpv4TotalLengthCounter
	decrementHolder  PatternFlowIpv4TotalLengthCounter
	metricTagsHolder PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
}

func NewPatternFlowIpv4TotalLength() PatternFlowIpv4TotalLength {
	obj := patternFlowIpv4TotalLength{obj: &otg.PatternFlowIpv4TotalLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TotalLength) msg() *otg.PatternFlowIpv4TotalLength {
	return obj.obj
}

func (obj *patternFlowIpv4TotalLength) setMsg(msg *otg.PatternFlowIpv4TotalLength) PatternFlowIpv4TotalLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TotalLength struct {
	obj *patternFlowIpv4TotalLength
}

type marshalPatternFlowIpv4TotalLength interface {
	// ToProto marshals PatternFlowIpv4TotalLength to protobuf object *otg.PatternFlowIpv4TotalLength
	ToProto() (*otg.PatternFlowIpv4TotalLength, error)
	// ToPbText marshals PatternFlowIpv4TotalLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TotalLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TotalLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TotalLength struct {
	obj *patternFlowIpv4TotalLength
}

type unMarshalPatternFlowIpv4TotalLength interface {
	// FromProto unmarshals PatternFlowIpv4TotalLength from protobuf object *otg.PatternFlowIpv4TotalLength
	FromProto(msg *otg.PatternFlowIpv4TotalLength) (PatternFlowIpv4TotalLength, error)
	// FromPbText unmarshals PatternFlowIpv4TotalLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TotalLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TotalLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TotalLength) Marshal() marshalPatternFlowIpv4TotalLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TotalLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TotalLength) Unmarshal() unMarshalPatternFlowIpv4TotalLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TotalLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TotalLength) ToProto() (*otg.PatternFlowIpv4TotalLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TotalLength) FromProto(msg *otg.PatternFlowIpv4TotalLength) (PatternFlowIpv4TotalLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TotalLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TotalLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TotalLength) FromJson(value string) error {
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

func (obj *patternFlowIpv4TotalLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TotalLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TotalLength) Clone() (PatternFlowIpv4TotalLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TotalLength()
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

func (obj *patternFlowIpv4TotalLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TotalLength is total length
type PatternFlowIpv4TotalLength interface {
	Validation
	// msg marshals PatternFlowIpv4TotalLength to protobuf object *otg.PatternFlowIpv4TotalLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TotalLength
	// setMsg unmarshals PatternFlowIpv4TotalLength from protobuf object *otg.PatternFlowIpv4TotalLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TotalLength) PatternFlowIpv4TotalLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TotalLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TotalLength
	// validate validates PatternFlowIpv4TotalLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TotalLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TotalLengthChoiceEnum, set in PatternFlowIpv4TotalLength
	Choice() PatternFlowIpv4TotalLengthChoiceEnum
	// setChoice assigns PatternFlowIpv4TotalLengthChoiceEnum provided by user to PatternFlowIpv4TotalLength
	setChoice(value PatternFlowIpv4TotalLengthChoiceEnum) PatternFlowIpv4TotalLength
	// HasChoice checks if Choice has been set in PatternFlowIpv4TotalLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TotalLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TotalLength
	SetValue(value uint32) PatternFlowIpv4TotalLength
	// HasValue checks if Value has been set in PatternFlowIpv4TotalLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TotalLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TotalLength
	SetValues(value []uint32) PatternFlowIpv4TotalLength
	// Auto returns uint32, set in PatternFlowIpv4TotalLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv4TotalLength
	HasAuto() bool
	// Increment returns PatternFlowIpv4TotalLengthCounter, set in PatternFlowIpv4TotalLength.
	// PatternFlowIpv4TotalLengthCounter is integer counter pattern
	Increment() PatternFlowIpv4TotalLengthCounter
	// SetIncrement assigns PatternFlowIpv4TotalLengthCounter provided by user to PatternFlowIpv4TotalLength.
	// PatternFlowIpv4TotalLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TotalLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TotalLengthCounter, set in PatternFlowIpv4TotalLength.
	// PatternFlowIpv4TotalLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv4TotalLengthCounter
	// SetDecrement assigns PatternFlowIpv4TotalLengthCounter provided by user to PatternFlowIpv4TotalLength.
	// PatternFlowIpv4TotalLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TotalLength
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIterIter, set in PatternFlowIpv4TotalLength
	MetricTags() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	setNil()
}

type PatternFlowIpv4TotalLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv4TotalLength
var PatternFlowIpv4TotalLengthChoice = struct {
	VALUE     PatternFlowIpv4TotalLengthChoiceEnum
	VALUES    PatternFlowIpv4TotalLengthChoiceEnum
	AUTO      PatternFlowIpv4TotalLengthChoiceEnum
	INCREMENT PatternFlowIpv4TotalLengthChoiceEnum
	DECREMENT PatternFlowIpv4TotalLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv4TotalLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TotalLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv4TotalLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv4TotalLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TotalLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TotalLength) Choice() PatternFlowIpv4TotalLengthChoiceEnum {
	return PatternFlowIpv4TotalLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TotalLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TotalLength) setChoice(value PatternFlowIpv4TotalLengthChoiceEnum) PatternFlowIpv4TotalLength {
	intValue, ok := otg.PatternFlowIpv4TotalLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TotalLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TotalLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TotalLengthChoice.VALUE {
		defaultValue := uint32(46)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TotalLengthChoice.VALUES {
		defaultValue := []uint32{46}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TotalLengthChoice.AUTO {
		defaultValue := uint32(46)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv4TotalLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TotalLengthCounter().msg()
	}

	if value == PatternFlowIpv4TotalLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TotalLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TotalLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TotalLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TotalLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TotalLength object
func (obj *patternFlowIpv4TotalLength) SetValue(value uint32) PatternFlowIpv4TotalLength {
	obj.setChoice(PatternFlowIpv4TotalLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TotalLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{46})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TotalLength object
func (obj *patternFlowIpv4TotalLength) SetValues(value []uint32) PatternFlowIpv4TotalLength {
	obj.setChoice(PatternFlowIpv4TotalLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv4TotalLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv4TotalLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv4TotalLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv4TotalLengthCounter
func (obj *patternFlowIpv4TotalLength) Increment() PatternFlowIpv4TotalLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TotalLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TotalLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TotalLengthCounter
func (obj *patternFlowIpv4TotalLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TotalLengthCounter value in the PatternFlowIpv4TotalLength object
func (obj *patternFlowIpv4TotalLength) SetIncrement(value PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLength {
	obj.setChoice(PatternFlowIpv4TotalLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TotalLengthCounter
func (obj *patternFlowIpv4TotalLength) Decrement() PatternFlowIpv4TotalLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TotalLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TotalLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TotalLengthCounter
func (obj *patternFlowIpv4TotalLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TotalLengthCounter value in the PatternFlowIpv4TotalLength object
func (obj *patternFlowIpv4TotalLength) SetDecrement(value PatternFlowIpv4TotalLengthCounter) PatternFlowIpv4TotalLength {
	obj.setChoice(PatternFlowIpv4TotalLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TotalLengthMetricTag
func (obj *patternFlowIpv4TotalLength) MetricTags() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TotalLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter struct {
	obj                                      *patternFlowIpv4TotalLength
	patternFlowIpv4TotalLengthMetricTagSlice []PatternFlowIpv4TotalLengthMetricTag
	fieldPtr                                 *[]*otg.PatternFlowIpv4TotalLengthMetricTag
}

func newPatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter(ptr *[]*otg.PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	return &patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter interface {
	setMsg(*patternFlowIpv4TotalLength) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	Items() []PatternFlowIpv4TotalLengthMetricTag
	Add() PatternFlowIpv4TotalLengthMetricTag
	Append(items ...PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	Set(index int, newObj PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	Clear() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	clearHolderSlice() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter
}

func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) setMsg(msg *patternFlowIpv4TotalLength) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TotalLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) Items() []PatternFlowIpv4TotalLengthMetricTag {
	return obj.patternFlowIpv4TotalLengthMetricTagSlice
}

func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) Add() PatternFlowIpv4TotalLengthMetricTag {
	newObj := &otg.PatternFlowIpv4TotalLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TotalLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TotalLengthMetricTagSlice = append(obj.patternFlowIpv4TotalLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) Append(items ...PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TotalLengthMetricTagSlice = append(obj.patternFlowIpv4TotalLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) Set(index int, newObj PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TotalLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) Clear() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TotalLengthMetricTag{}
		obj.patternFlowIpv4TotalLengthMetricTagSlice = []PatternFlowIpv4TotalLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) clearHolderSlice() PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	if len(obj.patternFlowIpv4TotalLengthMetricTagSlice) > 0 {
		obj.patternFlowIpv4TotalLengthMetricTagSlice = []PatternFlowIpv4TotalLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter) appendHolderSlice(item PatternFlowIpv4TotalLengthMetricTag) PatternFlowIpv4TotalLengthPatternFlowIpv4TotalLengthMetricTagIter {
	obj.patternFlowIpv4TotalLengthMetricTagSlice = append(obj.patternFlowIpv4TotalLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TotalLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TotalLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TotalLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TotalLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TotalLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4TotalLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4TotalLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4TotalLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv4TotalLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4TotalLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4TotalLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4TotalLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4TotalLength")
			}
		} else {
			intVal := otg.PatternFlowIpv4TotalLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4TotalLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
