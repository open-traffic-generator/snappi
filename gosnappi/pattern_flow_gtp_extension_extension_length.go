package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionExtensionLength *****
type patternFlowGtpExtensionExtensionLength struct {
	validation
	obj              *otg.PatternFlowGtpExtensionExtensionLength
	marshaller       marshalPatternFlowGtpExtensionExtensionLength
	unMarshaller     unMarshalPatternFlowGtpExtensionExtensionLength
	incrementHolder  PatternFlowGtpExtensionExtensionLengthCounter
	decrementHolder  PatternFlowGtpExtensionExtensionLengthCounter
	metricTagsHolder PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
}

func NewPatternFlowGtpExtensionExtensionLength() PatternFlowGtpExtensionExtensionLength {
	obj := patternFlowGtpExtensionExtensionLength{obj: &otg.PatternFlowGtpExtensionExtensionLength{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionExtensionLength) msg() *otg.PatternFlowGtpExtensionExtensionLength {
	return obj.obj
}

func (obj *patternFlowGtpExtensionExtensionLength) setMsg(msg *otg.PatternFlowGtpExtensionExtensionLength) PatternFlowGtpExtensionExtensionLength {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionExtensionLength struct {
	obj *patternFlowGtpExtensionExtensionLength
}

type marshalPatternFlowGtpExtensionExtensionLength interface {
	// ToProto marshals PatternFlowGtpExtensionExtensionLength to protobuf object *otg.PatternFlowGtpExtensionExtensionLength
	ToProto() (*otg.PatternFlowGtpExtensionExtensionLength, error)
	// ToPbText marshals PatternFlowGtpExtensionExtensionLength to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionExtensionLength to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionExtensionLength to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionExtensionLength struct {
	obj *patternFlowGtpExtensionExtensionLength
}

type unMarshalPatternFlowGtpExtensionExtensionLength interface {
	// FromProto unmarshals PatternFlowGtpExtensionExtensionLength from protobuf object *otg.PatternFlowGtpExtensionExtensionLength
	FromProto(msg *otg.PatternFlowGtpExtensionExtensionLength) (PatternFlowGtpExtensionExtensionLength, error)
	// FromPbText unmarshals PatternFlowGtpExtensionExtensionLength from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionExtensionLength from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionExtensionLength from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionExtensionLength) Marshal() marshalPatternFlowGtpExtensionExtensionLength {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionExtensionLength{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionExtensionLength) Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLength {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionExtensionLength{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionExtensionLength) ToProto() (*otg.PatternFlowGtpExtensionExtensionLength, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionExtensionLength) FromProto(msg *otg.PatternFlowGtpExtensionExtensionLength) (PatternFlowGtpExtensionExtensionLength, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionExtensionLength) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLength) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLength) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLength) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionExtensionLength) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionExtensionLength) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionExtensionLength) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLength) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionExtensionLength) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionExtensionLength) Clone() (PatternFlowGtpExtensionExtensionLength, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionExtensionLength()
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

func (obj *patternFlowGtpExtensionExtensionLength) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpExtensionExtensionLength is this field states the length of this extension header,  including the length, the contents, and the next extension header field, in 4-octet units, so the length of the extension must  always be a multiple of 4.
type PatternFlowGtpExtensionExtensionLength interface {
	Validation
	// msg marshals PatternFlowGtpExtensionExtensionLength to protobuf object *otg.PatternFlowGtpExtensionExtensionLength
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionExtensionLength
	// setMsg unmarshals PatternFlowGtpExtensionExtensionLength from protobuf object *otg.PatternFlowGtpExtensionExtensionLength
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionExtensionLength) PatternFlowGtpExtensionExtensionLength
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionExtensionLength
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionExtensionLength
	// validate validates PatternFlowGtpExtensionExtensionLength
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionExtensionLength, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpExtensionExtensionLengthChoiceEnum, set in PatternFlowGtpExtensionExtensionLength
	Choice() PatternFlowGtpExtensionExtensionLengthChoiceEnum
	// setChoice assigns PatternFlowGtpExtensionExtensionLengthChoiceEnum provided by user to PatternFlowGtpExtensionExtensionLength
	setChoice(value PatternFlowGtpExtensionExtensionLengthChoiceEnum) PatternFlowGtpExtensionExtensionLength
	// HasChoice checks if Choice has been set in PatternFlowGtpExtensionExtensionLength
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpExtensionExtensionLength.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpExtensionExtensionLength
	SetValue(value uint32) PatternFlowGtpExtensionExtensionLength
	// HasValue checks if Value has been set in PatternFlowGtpExtensionExtensionLength
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpExtensionExtensionLength.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpExtensionExtensionLength
	SetValues(value []uint32) PatternFlowGtpExtensionExtensionLength
	// Increment returns PatternFlowGtpExtensionExtensionLengthCounter, set in PatternFlowGtpExtensionExtensionLength.
	// PatternFlowGtpExtensionExtensionLengthCounter is integer counter pattern
	Increment() PatternFlowGtpExtensionExtensionLengthCounter
	// SetIncrement assigns PatternFlowGtpExtensionExtensionLengthCounter provided by user to PatternFlowGtpExtensionExtensionLength.
	// PatternFlowGtpExtensionExtensionLengthCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLength
	// HasIncrement checks if Increment has been set in PatternFlowGtpExtensionExtensionLength
	HasIncrement() bool
	// Decrement returns PatternFlowGtpExtensionExtensionLengthCounter, set in PatternFlowGtpExtensionExtensionLength.
	// PatternFlowGtpExtensionExtensionLengthCounter is integer counter pattern
	Decrement() PatternFlowGtpExtensionExtensionLengthCounter
	// SetDecrement assigns PatternFlowGtpExtensionExtensionLengthCounter provided by user to PatternFlowGtpExtensionExtensionLength.
	// PatternFlowGtpExtensionExtensionLengthCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLength
	// HasDecrement checks if Decrement has been set in PatternFlowGtpExtensionExtensionLength
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIterIter, set in PatternFlowGtpExtensionExtensionLength
	MetricTags() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	setNil()
}

type PatternFlowGtpExtensionExtensionLengthChoiceEnum string

// Enum of Choice on PatternFlowGtpExtensionExtensionLength
var PatternFlowGtpExtensionExtensionLengthChoice = struct {
	VALUE     PatternFlowGtpExtensionExtensionLengthChoiceEnum
	VALUES    PatternFlowGtpExtensionExtensionLengthChoiceEnum
	INCREMENT PatternFlowGtpExtensionExtensionLengthChoiceEnum
	DECREMENT PatternFlowGtpExtensionExtensionLengthChoiceEnum
}{
	VALUE:     PatternFlowGtpExtensionExtensionLengthChoiceEnum("value"),
	VALUES:    PatternFlowGtpExtensionExtensionLengthChoiceEnum("values"),
	INCREMENT: PatternFlowGtpExtensionExtensionLengthChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpExtensionExtensionLengthChoiceEnum("decrement"),
}

func (obj *patternFlowGtpExtensionExtensionLength) Choice() PatternFlowGtpExtensionExtensionLengthChoiceEnum {
	return PatternFlowGtpExtensionExtensionLengthChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpExtensionExtensionLength) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpExtensionExtensionLength) setChoice(value PatternFlowGtpExtensionExtensionLengthChoiceEnum) PatternFlowGtpExtensionExtensionLength {
	intValue, ok := otg.PatternFlowGtpExtensionExtensionLength_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpExtensionExtensionLengthChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpExtensionExtensionLength_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpExtensionExtensionLengthChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpExtensionExtensionLengthChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpExtensionExtensionLengthChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpExtensionExtensionLengthCounter().msg()
	}

	if value == PatternFlowGtpExtensionExtensionLengthChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpExtensionExtensionLengthCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpExtensionExtensionLength) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpExtensionExtensionLength) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpExtensionExtensionLength object
func (obj *patternFlowGtpExtensionExtensionLength) SetValue(value uint32) PatternFlowGtpExtensionExtensionLength {
	obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpExtensionExtensionLength) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpExtensionExtensionLength object
func (obj *patternFlowGtpExtensionExtensionLength) SetValues(value []uint32) PatternFlowGtpExtensionExtensionLength {
	obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionExtensionLengthCounter
func (obj *patternFlowGtpExtensionExtensionLength) Increment() PatternFlowGtpExtensionExtensionLengthCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpExtensionExtensionLengthCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionExtensionLengthCounter
func (obj *patternFlowGtpExtensionExtensionLength) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpExtensionExtensionLengthCounter value in the PatternFlowGtpExtensionExtensionLength object
func (obj *patternFlowGtpExtensionExtensionLength) SetIncrement(value PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLength {
	obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionExtensionLengthCounter
func (obj *patternFlowGtpExtensionExtensionLength) Decrement() PatternFlowGtpExtensionExtensionLengthCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpExtensionExtensionLengthCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionExtensionLengthCounter
func (obj *patternFlowGtpExtensionExtensionLength) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpExtensionExtensionLengthCounter value in the PatternFlowGtpExtensionExtensionLength object
func (obj *patternFlowGtpExtensionExtensionLength) SetDecrement(value PatternFlowGtpExtensionExtensionLengthCounter) PatternFlowGtpExtensionExtensionLength {
	obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpExtensionExtensionLengthMetricTag
func (obj *patternFlowGtpExtensionExtensionLength) MetricTags() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpExtensionExtensionLengthMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter struct {
	obj                                                  *patternFlowGtpExtensionExtensionLength
	patternFlowGtpExtensionExtensionLengthMetricTagSlice []PatternFlowGtpExtensionExtensionLengthMetricTag
	fieldPtr                                             *[]*otg.PatternFlowGtpExtensionExtensionLengthMetricTag
}

func newPatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter(ptr *[]*otg.PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	return &patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter interface {
	setMsg(*patternFlowGtpExtensionExtensionLength) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	Items() []PatternFlowGtpExtensionExtensionLengthMetricTag
	Add() PatternFlowGtpExtensionExtensionLengthMetricTag
	Append(items ...PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	Set(index int, newObj PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	Clear() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	clearHolderSlice() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
	appendHolderSlice(item PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter
}

func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) setMsg(msg *patternFlowGtpExtensionExtensionLength) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpExtensionExtensionLengthMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) Items() []PatternFlowGtpExtensionExtensionLengthMetricTag {
	return obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice
}

func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) Add() PatternFlowGtpExtensionExtensionLengthMetricTag {
	newObj := &otg.PatternFlowGtpExtensionExtensionLengthMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpExtensionExtensionLengthMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice = append(obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) Append(items ...PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice = append(obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) Set(index int, newObj PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) Clear() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpExtensionExtensionLengthMetricTag{}
		obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice = []PatternFlowGtpExtensionExtensionLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) clearHolderSlice() PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	if len(obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice) > 0 {
		obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice = []PatternFlowGtpExtensionExtensionLengthMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter) appendHolderSlice(item PatternFlowGtpExtensionExtensionLengthMetricTag) PatternFlowGtpExtensionExtensionLengthPatternFlowGtpExtensionExtensionLengthMetricTagIter {
	obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice = append(obj.patternFlowGtpExtensionExtensionLengthMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpExtensionExtensionLength) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionExtensionLength.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpExtensionExtensionLength.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpExtensionExtensionLengthMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpExtensionExtensionLength) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpExtensionExtensionLengthChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionExtensionLengthChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpExtensionExtensionLengthChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionExtensionLengthChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionExtensionLengthChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpExtensionExtensionLengthChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpExtensionExtensionLength")
			}
		} else {
			intVal := otg.PatternFlowGtpExtensionExtensionLength_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpExtensionExtensionLength_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
