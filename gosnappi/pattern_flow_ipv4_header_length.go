package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4HeaderLength *****
type patternFlowIpv4HeaderLength struct {
	validation
	obj              *otg.PatternFlowIpv4HeaderLength
	marshaller       marshalPatternFlowIpv4HeaderLength
	unMarshaller     unMarshalPatternFlowIpv4HeaderLength
	incrementHolder  PatternFlowIpv4HeaderLengthCounter
	decrementHolder  PatternFlowIpv4HeaderLengthCounter
	metricTagsHolder PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
}

func NewPatternFlowIpv4HeaderLength() PatternFlowIpv4HeaderLength {
	obj := patternFlowIpv4HeaderLength{obj: &otg.PatternFlowIpv4HeaderLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4HeaderLength) msg() *otg.PatternFlowIpv4HeaderLength {
	return obj.obj
}

func (obj *patternFlowIpv4HeaderLength) setMsg(msg *otg.PatternFlowIpv4HeaderLength) PatternFlowIpv4HeaderLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4HeaderLength struct {
	obj *patternFlowIpv4HeaderLength
}

type marshalPatternFlowIpv4HeaderLength interface {
	// ToProto marshals PatternFlowIpv4HeaderLength to protobuf object *otg.PatternFlowIpv4HeaderLength
	ToProto() (*otg.PatternFlowIpv4HeaderLength, error)
	// ToPbText marshals PatternFlowIpv4HeaderLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4HeaderLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4HeaderLength to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4HeaderLength to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4HeaderLength struct {
	obj *patternFlowIpv4HeaderLength
}

type unMarshalPatternFlowIpv4HeaderLength interface {
	// FromProto unmarshals PatternFlowIpv4HeaderLength from protobuf object *otg.PatternFlowIpv4HeaderLength
	FromProto(msg *otg.PatternFlowIpv4HeaderLength) (PatternFlowIpv4HeaderLength, error)
	// FromPbText unmarshals PatternFlowIpv4HeaderLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4HeaderLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4HeaderLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4HeaderLength) Marshal() marshalPatternFlowIpv4HeaderLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4HeaderLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4HeaderLength) Unmarshal() unMarshalPatternFlowIpv4HeaderLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4HeaderLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4HeaderLength) ToProto() (*otg.PatternFlowIpv4HeaderLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4HeaderLength) FromProto(msg *otg.PatternFlowIpv4HeaderLength) (PatternFlowIpv4HeaderLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4HeaderLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4HeaderLength) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4HeaderLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4HeaderLength) FromJson(value string) error {
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

func (obj *patternFlowIpv4HeaderLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4HeaderLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4HeaderLength) Clone() (PatternFlowIpv4HeaderLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4HeaderLength()
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

func (obj *patternFlowIpv4HeaderLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4HeaderLength is header length
type PatternFlowIpv4HeaderLength interface {
	Validation
	// msg marshals PatternFlowIpv4HeaderLength to protobuf object *otg.PatternFlowIpv4HeaderLength
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4HeaderLength
	// setMsg unmarshals PatternFlowIpv4HeaderLength from protobuf object *otg.PatternFlowIpv4HeaderLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4HeaderLength) PatternFlowIpv4HeaderLength
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4HeaderLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4HeaderLength
	// validate validates PatternFlowIpv4HeaderLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4HeaderLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4HeaderLengthChoiceEnum, set in PatternFlowIpv4HeaderLength
	Choice() PatternFlowIpv4HeaderLengthChoiceEnum
	// setChoice assigns PatternFlowIpv4HeaderLengthChoiceEnum provided by user to PatternFlowIpv4HeaderLength
	setChoice(value PatternFlowIpv4HeaderLengthChoiceEnum) PatternFlowIpv4HeaderLength
	// HasChoice checks if Choice has been set in PatternFlowIpv4HeaderLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4HeaderLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4HeaderLength
	SetValue(value uint32) PatternFlowIpv4HeaderLength
	// HasValue checks if Value has been set in PatternFlowIpv4HeaderLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4HeaderLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4HeaderLength
	SetValues(value []uint32) PatternFlowIpv4HeaderLength
	// Auto returns uint32, set in PatternFlowIpv4HeaderLength.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowIpv4HeaderLength
	HasAuto() bool
	// Increment returns PatternFlowIpv4HeaderLengthCounter, set in PatternFlowIpv4HeaderLength.
	// PatternFlowIpv4HeaderLengthCounter is integer counter pattern
	Increment() PatternFlowIpv4HeaderLengthCounter
	// SetIncrement assigns PatternFlowIpv4HeaderLengthCounter provided by user to PatternFlowIpv4HeaderLength.
	// PatternFlowIpv4HeaderLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLength
	// HasIncrement checks if Increment has been set in PatternFlowIpv4HeaderLength
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4HeaderLengthCounter, set in PatternFlowIpv4HeaderLength.
	// PatternFlowIpv4HeaderLengthCounter is integer counter pattern
	Decrement() PatternFlowIpv4HeaderLengthCounter
	// SetDecrement assigns PatternFlowIpv4HeaderLengthCounter provided by user to PatternFlowIpv4HeaderLength.
	// PatternFlowIpv4HeaderLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLength
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4HeaderLength
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIterIter, set in PatternFlowIpv4HeaderLength
	MetricTags() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	setNil()
}

type PatternFlowIpv4HeaderLengthChoiceEnum string

// Enum of Choice on PatternFlowIpv4HeaderLength
var PatternFlowIpv4HeaderLengthChoice = struct {
	VALUE     PatternFlowIpv4HeaderLengthChoiceEnum
	VALUES    PatternFlowIpv4HeaderLengthChoiceEnum
	AUTO      PatternFlowIpv4HeaderLengthChoiceEnum
	INCREMENT PatternFlowIpv4HeaderLengthChoiceEnum
	DECREMENT PatternFlowIpv4HeaderLengthChoiceEnum
}{
	VALUE:     PatternFlowIpv4HeaderLengthChoiceEnum("value"),
	VALUES:    PatternFlowIpv4HeaderLengthChoiceEnum("values"),
	AUTO:      PatternFlowIpv4HeaderLengthChoiceEnum("auto"),
	INCREMENT: PatternFlowIpv4HeaderLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4HeaderLengthChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4HeaderLength) Choice() PatternFlowIpv4HeaderLengthChoiceEnum {
	return PatternFlowIpv4HeaderLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4HeaderLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4HeaderLength) setChoice(value PatternFlowIpv4HeaderLengthChoiceEnum) PatternFlowIpv4HeaderLength {
	intValue, ok := otg.PatternFlowIpv4HeaderLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4HeaderLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4HeaderLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4HeaderLengthChoice.VALUE {
		defaultValue := uint32(5)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4HeaderLengthChoice.VALUES {
		defaultValue := []uint32{5}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4HeaderLengthChoice.AUTO {
		defaultValue := uint32(5)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowIpv4HeaderLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4HeaderLengthCounter().msg()
	}

	if value == PatternFlowIpv4HeaderLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4HeaderLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4HeaderLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4HeaderLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4HeaderLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4HeaderLength object
func (obj *patternFlowIpv4HeaderLength) SetValue(value uint32) PatternFlowIpv4HeaderLength {
	obj.setChoice(PatternFlowIpv4HeaderLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4HeaderLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{5})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4HeaderLength object
func (obj *patternFlowIpv4HeaderLength) SetValues(value []uint32) PatternFlowIpv4HeaderLength {
	obj.setChoice(PatternFlowIpv4HeaderLengthChoice.VALUES)
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
func (obj *patternFlowIpv4HeaderLength) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowIpv4HeaderLengthChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowIpv4HeaderLength) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowIpv4HeaderLengthCounter
func (obj *patternFlowIpv4HeaderLength) Increment() PatternFlowIpv4HeaderLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4HeaderLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4HeaderLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4HeaderLengthCounter
func (obj *patternFlowIpv4HeaderLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4HeaderLengthCounter value in the PatternFlowIpv4HeaderLength object
func (obj *patternFlowIpv4HeaderLength) SetIncrement(value PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLength {
	obj.setChoice(PatternFlowIpv4HeaderLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4HeaderLengthCounter
func (obj *patternFlowIpv4HeaderLength) Decrement() PatternFlowIpv4HeaderLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4HeaderLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4HeaderLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4HeaderLengthCounter
func (obj *patternFlowIpv4HeaderLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4HeaderLengthCounter value in the PatternFlowIpv4HeaderLength object
func (obj *patternFlowIpv4HeaderLength) SetDecrement(value PatternFlowIpv4HeaderLengthCounter) PatternFlowIpv4HeaderLength {
	obj.setChoice(PatternFlowIpv4HeaderLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4HeaderLengthMetricTag
func (obj *patternFlowIpv4HeaderLength) MetricTags() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4HeaderLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter struct {
	obj                                       *patternFlowIpv4HeaderLength
	patternFlowIpv4HeaderLengthMetricTagSlice []PatternFlowIpv4HeaderLengthMetricTag
	fieldPtr                                  *[]*otg.PatternFlowIpv4HeaderLengthMetricTag
}

func newPatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter(ptr *[]*otg.PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	return &patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter interface {
	setMsg(*patternFlowIpv4HeaderLength) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	Items() []PatternFlowIpv4HeaderLengthMetricTag
	Add() PatternFlowIpv4HeaderLengthMetricTag
	Append(items ...PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	Set(index int, newObj PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	Clear() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	clearHolderSlice() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
	appendHolderSlice(item PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter
}

func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) setMsg(msg *patternFlowIpv4HeaderLength) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4HeaderLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) Items() []PatternFlowIpv4HeaderLengthMetricTag {
	return obj.patternFlowIpv4HeaderLengthMetricTagSlice
}

func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) Add() PatternFlowIpv4HeaderLengthMetricTag {
	newObj := &otg.PatternFlowIpv4HeaderLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4HeaderLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4HeaderLengthMetricTagSlice = append(obj.patternFlowIpv4HeaderLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) Append(items ...PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4HeaderLengthMetricTagSlice = append(obj.patternFlowIpv4HeaderLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) Set(index int, newObj PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4HeaderLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) Clear() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4HeaderLengthMetricTag{}
		obj.patternFlowIpv4HeaderLengthMetricTagSlice = []PatternFlowIpv4HeaderLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) clearHolderSlice() PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	if len(obj.patternFlowIpv4HeaderLengthMetricTagSlice) > 0 {
		obj.patternFlowIpv4HeaderLengthMetricTagSlice = []PatternFlowIpv4HeaderLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter) appendHolderSlice(item PatternFlowIpv4HeaderLengthMetricTag) PatternFlowIpv4HeaderLengthPatternFlowIpv4HeaderLengthMetricTagIter {
	obj.patternFlowIpv4HeaderLengthMetricTagSlice = append(obj.patternFlowIpv4HeaderLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4HeaderLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLength.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4HeaderLength.Values <= 15 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4HeaderLength.Auto <= 15 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4HeaderLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4HeaderLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4HeaderLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4HeaderLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4HeaderLengthChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowIpv4HeaderLengthChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4HeaderLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4HeaderLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4HeaderLengthChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4HeaderLength")
			}
		} else {
			intVal := otg.PatternFlowIpv4HeaderLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4HeaderLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
