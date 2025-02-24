package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4FragmentOffset *****
type patternFlowIpv4FragmentOffset struct {
	validation
	obj              *otg.PatternFlowIpv4FragmentOffset
	marshaller       marshalPatternFlowIpv4FragmentOffset
	unMarshaller     unMarshalPatternFlowIpv4FragmentOffset
	incrementHolder  PatternFlowIpv4FragmentOffsetCounter
	decrementHolder  PatternFlowIpv4FragmentOffsetCounter
	metricTagsHolder PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
}

func NewPatternFlowIpv4FragmentOffset() PatternFlowIpv4FragmentOffset {
	obj := patternFlowIpv4FragmentOffset{obj: &otg.PatternFlowIpv4FragmentOffset{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4FragmentOffset) msg() *otg.PatternFlowIpv4FragmentOffset {
	return obj.obj
}

func (obj *patternFlowIpv4FragmentOffset) setMsg(msg *otg.PatternFlowIpv4FragmentOffset) PatternFlowIpv4FragmentOffset {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4FragmentOffset struct {
	obj *patternFlowIpv4FragmentOffset
}

type marshalPatternFlowIpv4FragmentOffset interface {
	// ToProto marshals PatternFlowIpv4FragmentOffset to protobuf object *otg.PatternFlowIpv4FragmentOffset
	ToProto() (*otg.PatternFlowIpv4FragmentOffset, error)
	// ToPbText marshals PatternFlowIpv4FragmentOffset to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4FragmentOffset to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4FragmentOffset to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4FragmentOffset to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4FragmentOffset struct {
	obj *patternFlowIpv4FragmentOffset
}

type unMarshalPatternFlowIpv4FragmentOffset interface {
	// FromProto unmarshals PatternFlowIpv4FragmentOffset from protobuf object *otg.PatternFlowIpv4FragmentOffset
	FromProto(msg *otg.PatternFlowIpv4FragmentOffset) (PatternFlowIpv4FragmentOffset, error)
	// FromPbText unmarshals PatternFlowIpv4FragmentOffset from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4FragmentOffset from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4FragmentOffset from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4FragmentOffset) Marshal() marshalPatternFlowIpv4FragmentOffset {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4FragmentOffset{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4FragmentOffset) Unmarshal() unMarshalPatternFlowIpv4FragmentOffset {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4FragmentOffset{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4FragmentOffset) ToProto() (*otg.PatternFlowIpv4FragmentOffset, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4FragmentOffset) FromProto(msg *otg.PatternFlowIpv4FragmentOffset) (PatternFlowIpv4FragmentOffset, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4FragmentOffset) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffset) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffset) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffset) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4FragmentOffset) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4FragmentOffset) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4FragmentOffset) FromJson(value string) error {
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

func (obj *patternFlowIpv4FragmentOffset) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffset) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4FragmentOffset) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4FragmentOffset) Clone() (PatternFlowIpv4FragmentOffset, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4FragmentOffset()
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

func (obj *patternFlowIpv4FragmentOffset) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4FragmentOffset is fragment offset
type PatternFlowIpv4FragmentOffset interface {
	Validation
	// msg marshals PatternFlowIpv4FragmentOffset to protobuf object *otg.PatternFlowIpv4FragmentOffset
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4FragmentOffset
	// setMsg unmarshals PatternFlowIpv4FragmentOffset from protobuf object *otg.PatternFlowIpv4FragmentOffset
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4FragmentOffset) PatternFlowIpv4FragmentOffset
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4FragmentOffset
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4FragmentOffset
	// validate validates PatternFlowIpv4FragmentOffset
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4FragmentOffset, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4FragmentOffsetChoiceEnum, set in PatternFlowIpv4FragmentOffset
	Choice() PatternFlowIpv4FragmentOffsetChoiceEnum
	// setChoice assigns PatternFlowIpv4FragmentOffsetChoiceEnum provided by user to PatternFlowIpv4FragmentOffset
	setChoice(value PatternFlowIpv4FragmentOffsetChoiceEnum) PatternFlowIpv4FragmentOffset
	// HasChoice checks if Choice has been set in PatternFlowIpv4FragmentOffset
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4FragmentOffset.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4FragmentOffset
	SetValue(value uint32) PatternFlowIpv4FragmentOffset
	// HasValue checks if Value has been set in PatternFlowIpv4FragmentOffset
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4FragmentOffset.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4FragmentOffset
	SetValues(value []uint32) PatternFlowIpv4FragmentOffset
	// Increment returns PatternFlowIpv4FragmentOffsetCounter, set in PatternFlowIpv4FragmentOffset.
	// PatternFlowIpv4FragmentOffsetCounter is integer counter pattern
	Increment() PatternFlowIpv4FragmentOffsetCounter
	// SetIncrement assigns PatternFlowIpv4FragmentOffsetCounter provided by user to PatternFlowIpv4FragmentOffset.
	// PatternFlowIpv4FragmentOffsetCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffset
	// HasIncrement checks if Increment has been set in PatternFlowIpv4FragmentOffset
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4FragmentOffsetCounter, set in PatternFlowIpv4FragmentOffset.
	// PatternFlowIpv4FragmentOffsetCounter is integer counter pattern
	Decrement() PatternFlowIpv4FragmentOffsetCounter
	// SetDecrement assigns PatternFlowIpv4FragmentOffsetCounter provided by user to PatternFlowIpv4FragmentOffset.
	// PatternFlowIpv4FragmentOffsetCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffset
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4FragmentOffset
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIterIter, set in PatternFlowIpv4FragmentOffset
	MetricTags() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	setNil()
}

type PatternFlowIpv4FragmentOffsetChoiceEnum string

// Enum of Choice on PatternFlowIpv4FragmentOffset
var PatternFlowIpv4FragmentOffsetChoice = struct {
	VALUE     PatternFlowIpv4FragmentOffsetChoiceEnum
	VALUES    PatternFlowIpv4FragmentOffsetChoiceEnum
	INCREMENT PatternFlowIpv4FragmentOffsetChoiceEnum
	DECREMENT PatternFlowIpv4FragmentOffsetChoiceEnum
}{
	VALUE:     PatternFlowIpv4FragmentOffsetChoiceEnum("value"),
	VALUES:    PatternFlowIpv4FragmentOffsetChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4FragmentOffsetChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4FragmentOffsetChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4FragmentOffset) Choice() PatternFlowIpv4FragmentOffsetChoiceEnum {
	return PatternFlowIpv4FragmentOffsetChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4FragmentOffset) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4FragmentOffset) setChoice(value PatternFlowIpv4FragmentOffsetChoiceEnum) PatternFlowIpv4FragmentOffset {
	intValue, ok := otg.PatternFlowIpv4FragmentOffset_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4FragmentOffsetChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4FragmentOffset_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4FragmentOffsetChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4FragmentOffsetChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4FragmentOffsetChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4FragmentOffsetCounter().msg()
	}

	if value == PatternFlowIpv4FragmentOffsetChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4FragmentOffsetCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4FragmentOffset) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4FragmentOffset) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4FragmentOffset object
func (obj *patternFlowIpv4FragmentOffset) SetValue(value uint32) PatternFlowIpv4FragmentOffset {
	obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4FragmentOffset) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4FragmentOffset object
func (obj *patternFlowIpv4FragmentOffset) SetValues(value []uint32) PatternFlowIpv4FragmentOffset {
	obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4FragmentOffsetCounter
func (obj *patternFlowIpv4FragmentOffset) Increment() PatternFlowIpv4FragmentOffsetCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4FragmentOffsetCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4FragmentOffsetCounter
func (obj *patternFlowIpv4FragmentOffset) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4FragmentOffsetCounter value in the PatternFlowIpv4FragmentOffset object
func (obj *patternFlowIpv4FragmentOffset) SetIncrement(value PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffset {
	obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4FragmentOffsetCounter
func (obj *patternFlowIpv4FragmentOffset) Decrement() PatternFlowIpv4FragmentOffsetCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4FragmentOffsetCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4FragmentOffsetCounter
func (obj *patternFlowIpv4FragmentOffset) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4FragmentOffsetCounter value in the PatternFlowIpv4FragmentOffset object
func (obj *patternFlowIpv4FragmentOffset) SetDecrement(value PatternFlowIpv4FragmentOffsetCounter) PatternFlowIpv4FragmentOffset {
	obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4FragmentOffsetMetricTag
func (obj *patternFlowIpv4FragmentOffset) MetricTags() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4FragmentOffsetMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter struct {
	obj                                         *patternFlowIpv4FragmentOffset
	patternFlowIpv4FragmentOffsetMetricTagSlice []PatternFlowIpv4FragmentOffsetMetricTag
	fieldPtr                                    *[]*otg.PatternFlowIpv4FragmentOffsetMetricTag
}

func newPatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter(ptr *[]*otg.PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	return &patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter interface {
	setMsg(*patternFlowIpv4FragmentOffset) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	Items() []PatternFlowIpv4FragmentOffsetMetricTag
	Add() PatternFlowIpv4FragmentOffsetMetricTag
	Append(items ...PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	Set(index int, newObj PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	Clear() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	clearHolderSlice() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
	appendHolderSlice(item PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter
}

func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) setMsg(msg *patternFlowIpv4FragmentOffset) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4FragmentOffsetMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) Items() []PatternFlowIpv4FragmentOffsetMetricTag {
	return obj.patternFlowIpv4FragmentOffsetMetricTagSlice
}

func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) Add() PatternFlowIpv4FragmentOffsetMetricTag {
	newObj := &otg.PatternFlowIpv4FragmentOffsetMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4FragmentOffsetMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4FragmentOffsetMetricTagSlice = append(obj.patternFlowIpv4FragmentOffsetMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) Append(items ...PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4FragmentOffsetMetricTagSlice = append(obj.patternFlowIpv4FragmentOffsetMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) Set(index int, newObj PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4FragmentOffsetMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) Clear() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4FragmentOffsetMetricTag{}
		obj.patternFlowIpv4FragmentOffsetMetricTagSlice = []PatternFlowIpv4FragmentOffsetMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) clearHolderSlice() PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	if len(obj.patternFlowIpv4FragmentOffsetMetricTagSlice) > 0 {
		obj.patternFlowIpv4FragmentOffsetMetricTagSlice = []PatternFlowIpv4FragmentOffsetMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter) appendHolderSlice(item PatternFlowIpv4FragmentOffsetMetricTag) PatternFlowIpv4FragmentOffsetPatternFlowIpv4FragmentOffsetMetricTagIter {
	obj.patternFlowIpv4FragmentOffsetMetricTagSlice = append(obj.patternFlowIpv4FragmentOffsetMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4FragmentOffset) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4FragmentOffset.Value <= 31 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 31 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4FragmentOffset.Values <= 31 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4FragmentOffsetMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4FragmentOffset) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4FragmentOffsetChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4FragmentOffsetChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4FragmentOffsetChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4FragmentOffsetChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4FragmentOffsetChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4FragmentOffsetChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4FragmentOffset")
			}
		} else {
			intVal := otg.PatternFlowIpv4FragmentOffset_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4FragmentOffset_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
