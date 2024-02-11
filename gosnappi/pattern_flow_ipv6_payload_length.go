package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6PayloadLength *****
type patternFlowIpv6PayloadLength struct {
	validation
	obj              *otg.PatternFlowIpv6PayloadLength
	marshaller       marshalPatternFlowIpv6PayloadLength
	unMarshaller     unMarshalPatternFlowIpv6PayloadLength
	incrementHolder  PatternFlowIpv6PayloadLengthCounter
	decrementHolder  PatternFlowIpv6PayloadLengthCounter
	metricTagsHolder PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
}

func NewPatternFlowIpv6PayloadLength() PatternFlowIpv6PayloadLength {
	obj := patternFlowIpv6PayloadLength{obj: &otg.PatternFlowIpv6PayloadLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6PayloadLength) msg() *otg.PatternFlowIpv6PayloadLength {
	return obj.obj
}

func (obj *patternFlowIpv6PayloadLength) setMsg(msg *otg.PatternFlowIpv6PayloadLength) PatternFlowIpv6PayloadLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6PayloadLength struct {
	obj *patternFlowIpv6PayloadLength
}

type marshalPatternFlowIpv6PayloadLength interface {
	// ToProto marshals PatternFlowIpv6PayloadLength to protobuf object *otg.PatternFlowIpv6PayloadLength
	ToProto() (*otg.PatternFlowIpv6PayloadLength, error)
	// ToPbText marshals PatternFlowIpv6PayloadLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6PayloadLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6PayloadLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6PayloadLength struct {
	obj *patternFlowIpv6PayloadLength
}

type unMarshalPatternFlowIpv6PayloadLength interface {
	// FromProto unmarshals PatternFlowIpv6PayloadLength from protobuf object *otg.PatternFlowIpv6PayloadLength
	FromProto(msg *otg.PatternFlowIpv6PayloadLength) (PatternFlowIpv6PayloadLength, error)
	// FromPbText unmarshals PatternFlowIpv6PayloadLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6PayloadLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6PayloadLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6PayloadLength) Marshal() marshalPatternFlowIpv6PayloadLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6PayloadLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6PayloadLength) Unmarshal() unMarshalPatternFlowIpv6PayloadLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6PayloadLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6PayloadLength) ToProto() (*otg.PatternFlowIpv6PayloadLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6PayloadLength) FromProto(msg *otg.PatternFlowIpv6PayloadLength) (PatternFlowIpv6PayloadLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6PayloadLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6PayloadLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6PayloadLength) FromJson(value string) error {
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

func (obj *patternFlowIpv6PayloadLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6PayloadLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6PayloadLength) Clone() (PatternFlowIpv6PayloadLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6PayloadLength()
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

func (obj *patternFlowIpv6PayloadLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6PayloadLength is payload length
type PatternFlowIpv6PayloadLength interface {
	Validation
	// msg marshals PatternFlowIpv6PayloadLength to protobuf object *otg.PatternFlowIpv6PayloadLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6PayloadLength
	// setMsg unmarshals PatternFlowIpv6PayloadLength from protobuf object *otg.PatternFlowIpv6PayloadLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6PayloadLength) PatternFlowIpv6PayloadLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6PayloadLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6PayloadLength
	// validate validates PatternFlowIpv6PayloadLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6PayloadLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6PayloadLengthChoiceEnum, set in PatternFlowIpv6PayloadLength
	Choice() PatternFlowIpv6PayloadLengthChoiceEnum
	// setChoice assigns PatternFlowIpv6PayloadLengthChoiceEnum provided by user to PatternFlowIpv6PayloadLength
	setChoice(value PatternFlowIpv6PayloadLengthChoiceEnum) PatternFlowIpv6PayloadLength
	// HasChoice checks if Choice has been set in PatternFlowIpv6PayloadLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6PayloadLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6PayloadLength
	SetValue(value uint32) PatternFlowIpv6PayloadLength
	// HasValue checks if Value has been set in PatternFlowIpv6PayloadLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6PayloadLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6PayloadLength
	SetValues(value []uint32) PatternFlowIpv6PayloadLength
	// Auto returns uint32, set in PatternFlowIpv6PayloadLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv6PayloadLength
	HasAuto() bool
	// Increment returns PatternFlowIpv6PayloadLengthCounter, set in PatternFlowIpv6PayloadLength.
	// PatternFlowIpv6PayloadLengthCounter is integer counter pattern
	Increment() PatternFlowIpv6PayloadLengthCounter
	// SetIncrement assigns PatternFlowIpv6PayloadLengthCounter provided by user to PatternFlowIpv6PayloadLength.
	// PatternFlowIpv6PayloadLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv6PayloadLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6PayloadLengthCounter, set in PatternFlowIpv6PayloadLength.
	// PatternFlowIpv6PayloadLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv6PayloadLengthCounter
	// SetDecrement assigns PatternFlowIpv6PayloadLengthCounter provided by user to PatternFlowIpv6PayloadLength.
	// PatternFlowIpv6PayloadLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6PayloadLength
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIterIter, set in PatternFlowIpv6PayloadLength
	MetricTags() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	setNil()
}

type PatternFlowIpv6PayloadLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv6PayloadLength
var PatternFlowIpv6PayloadLengthChoice = struct {
	VALUE     PatternFlowIpv6PayloadLengthChoiceEnum
	VALUES    PatternFlowIpv6PayloadLengthChoiceEnum
	AUTO      PatternFlowIpv6PayloadLengthChoiceEnum
	INCREMENT PatternFlowIpv6PayloadLengthChoiceEnum
	DECREMENT PatternFlowIpv6PayloadLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv6PayloadLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv6PayloadLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv6PayloadLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv6PayloadLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6PayloadLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6PayloadLength) Choice() PatternFlowIpv6PayloadLengthChoiceEnum {
	return PatternFlowIpv6PayloadLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6PayloadLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6PayloadLength) setChoice(value PatternFlowIpv6PayloadLengthChoiceEnum) PatternFlowIpv6PayloadLength {
	intValue, ok := otg.PatternFlowIpv6PayloadLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6PayloadLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6PayloadLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6PayloadLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6PayloadLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6PayloadLengthChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv6PayloadLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6PayloadLengthCounter().msg()
	}

	if value == PatternFlowIpv6PayloadLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6PayloadLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6PayloadLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6PayloadLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6PayloadLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6PayloadLength object
func (obj *patternFlowIpv6PayloadLength) SetValue(value uint32) PatternFlowIpv6PayloadLength {
	obj.setChoice(PatternFlowIpv6PayloadLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6PayloadLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6PayloadLength object
func (obj *patternFlowIpv6PayloadLength) SetValues(value []uint32) PatternFlowIpv6PayloadLength {
	obj.setChoice(PatternFlowIpv6PayloadLengthChoice.VALUES)
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
func (obj *patternFlowIpv6PayloadLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv6PayloadLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv6PayloadLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv6PayloadLengthCounter
func (obj *patternFlowIpv6PayloadLength) Increment() PatternFlowIpv6PayloadLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6PayloadLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6PayloadLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6PayloadLengthCounter
func (obj *patternFlowIpv6PayloadLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6PayloadLengthCounter value in the PatternFlowIpv6PayloadLength object
func (obj *patternFlowIpv6PayloadLength) SetIncrement(value PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLength {
	obj.setChoice(PatternFlowIpv6PayloadLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6PayloadLengthCounter
func (obj *patternFlowIpv6PayloadLength) Decrement() PatternFlowIpv6PayloadLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6PayloadLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6PayloadLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6PayloadLengthCounter
func (obj *patternFlowIpv6PayloadLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6PayloadLengthCounter value in the PatternFlowIpv6PayloadLength object
func (obj *patternFlowIpv6PayloadLength) SetDecrement(value PatternFlowIpv6PayloadLengthCounter) PatternFlowIpv6PayloadLength {
	obj.setChoice(PatternFlowIpv6PayloadLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6PayloadLengthMetricTag
func (obj *patternFlowIpv6PayloadLength) MetricTags() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6PayloadLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter struct {
	obj                                        *patternFlowIpv6PayloadLength
	patternFlowIpv6PayloadLengthMetricTagSlice []PatternFlowIpv6PayloadLengthMetricTag
	fieldPtr                                   *[]*otg.PatternFlowIpv6PayloadLengthMetricTag
}

func newPatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter(ptr *[]*otg.PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	return &patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter interface {
	setMsg(*patternFlowIpv6PayloadLength) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	Items() []PatternFlowIpv6PayloadLengthMetricTag
	Add() PatternFlowIpv6PayloadLengthMetricTag
	Append(items ...PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	Set(index int, newObj PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	Clear() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	clearHolderSlice() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
	appendHolderSlice(item PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter
}

func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) setMsg(msg *patternFlowIpv6PayloadLength) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6PayloadLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) Items() []PatternFlowIpv6PayloadLengthMetricTag {
	return obj.patternFlowIpv6PayloadLengthMetricTagSlice
}

func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) Add() PatternFlowIpv6PayloadLengthMetricTag {
	newObj := &otg.PatternFlowIpv6PayloadLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6PayloadLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6PayloadLengthMetricTagSlice = append(obj.patternFlowIpv6PayloadLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) Append(items ...PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6PayloadLengthMetricTagSlice = append(obj.patternFlowIpv6PayloadLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) Set(index int, newObj PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6PayloadLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) Clear() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6PayloadLengthMetricTag{}
		obj.patternFlowIpv6PayloadLengthMetricTagSlice = []PatternFlowIpv6PayloadLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) clearHolderSlice() PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	if len(obj.patternFlowIpv6PayloadLengthMetricTagSlice) > 0 {
		obj.patternFlowIpv6PayloadLengthMetricTagSlice = []PatternFlowIpv6PayloadLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter) appendHolderSlice(item PatternFlowIpv6PayloadLengthMetricTag) PatternFlowIpv6PayloadLengthPatternFlowIpv6PayloadLengthMetricTagIter {
	obj.patternFlowIpv6PayloadLengthMetricTagSlice = append(obj.patternFlowIpv6PayloadLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6PayloadLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLength.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6PayloadLength.Values <= 65535 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6PayloadLength.Auto <= 65535 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6PayloadLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6PayloadLength) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv6PayloadLengthChoice.AUTO)

	}

}
