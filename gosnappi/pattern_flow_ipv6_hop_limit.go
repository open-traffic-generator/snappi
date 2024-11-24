package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6HopLimit *****
type patternFlowIpv6HopLimit struct {
	validation
	obj              *otg.PatternFlowIpv6HopLimit
	marshaller       marshalPatternFlowIpv6HopLimit
	unMarshaller     unMarshalPatternFlowIpv6HopLimit
	incrementHolder  PatternFlowIpv6HopLimitCounter
	decrementHolder  PatternFlowIpv6HopLimitCounter
	metricTagsHolder PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
}

func NewPatternFlowIpv6HopLimit() PatternFlowIpv6HopLimit {
	obj := patternFlowIpv6HopLimit{obj: &otg.PatternFlowIpv6HopLimit{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6HopLimit) msg() *otg.PatternFlowIpv6HopLimit {
	return obj.obj
}

func (obj *patternFlowIpv6HopLimit) setMsg(msg *otg.PatternFlowIpv6HopLimit) PatternFlowIpv6HopLimit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6HopLimit struct {
	obj *patternFlowIpv6HopLimit
}

type marshalPatternFlowIpv6HopLimit interface {
	// ToProto marshals PatternFlowIpv6HopLimit to protobuf object *otg.PatternFlowIpv6HopLimit
	ToProto() (*otg.PatternFlowIpv6HopLimit, error)
	// ToPbText marshals PatternFlowIpv6HopLimit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6HopLimit to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6HopLimit to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6HopLimit to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6HopLimit struct {
	obj *patternFlowIpv6HopLimit
}

type unMarshalPatternFlowIpv6HopLimit interface {
	// FromProto unmarshals PatternFlowIpv6HopLimit from protobuf object *otg.PatternFlowIpv6HopLimit
	FromProto(msg *otg.PatternFlowIpv6HopLimit) (PatternFlowIpv6HopLimit, error)
	// FromPbText unmarshals PatternFlowIpv6HopLimit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6HopLimit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6HopLimit from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6HopLimit) Marshal() marshalPatternFlowIpv6HopLimit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6HopLimit{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6HopLimit) Unmarshal() unMarshalPatternFlowIpv6HopLimit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6HopLimit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6HopLimit) ToProto() (*otg.PatternFlowIpv6HopLimit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6HopLimit) FromProto(msg *otg.PatternFlowIpv6HopLimit) (PatternFlowIpv6HopLimit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6HopLimit) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimit) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimit) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimit) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6HopLimit) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6HopLimit) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6HopLimit) FromJson(value string) error {
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

func (obj *patternFlowIpv6HopLimit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6HopLimit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6HopLimit) Clone() (PatternFlowIpv6HopLimit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6HopLimit()
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

func (obj *patternFlowIpv6HopLimit) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6HopLimit is hop limit
type PatternFlowIpv6HopLimit interface {
	Validation
	// msg marshals PatternFlowIpv6HopLimit to protobuf object *otg.PatternFlowIpv6HopLimit
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6HopLimit
	// setMsg unmarshals PatternFlowIpv6HopLimit from protobuf object *otg.PatternFlowIpv6HopLimit
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6HopLimit) PatternFlowIpv6HopLimit
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6HopLimit
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6HopLimit
	// validate validates PatternFlowIpv6HopLimit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6HopLimit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6HopLimitChoiceEnum, set in PatternFlowIpv6HopLimit
	Choice() PatternFlowIpv6HopLimitChoiceEnum
	// setChoice assigns PatternFlowIpv6HopLimitChoiceEnum provided by user to PatternFlowIpv6HopLimit
	setChoice(value PatternFlowIpv6HopLimitChoiceEnum) PatternFlowIpv6HopLimit
	// HasChoice checks if Choice has been set in PatternFlowIpv6HopLimit
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6HopLimit.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6HopLimit
	SetValue(value uint32) PatternFlowIpv6HopLimit
	// HasValue checks if Value has been set in PatternFlowIpv6HopLimit
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6HopLimit.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6HopLimit
	SetValues(value []uint32) PatternFlowIpv6HopLimit
	// Increment returns PatternFlowIpv6HopLimitCounter, set in PatternFlowIpv6HopLimit.
	// PatternFlowIpv6HopLimitCounter is integer counter pattern
	Increment() PatternFlowIpv6HopLimitCounter
	// SetIncrement assigns PatternFlowIpv6HopLimitCounter provided by user to PatternFlowIpv6HopLimit.
	// PatternFlowIpv6HopLimitCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimit
	// HasIncrement checks if Increment has been set in PatternFlowIpv6HopLimit
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6HopLimitCounter, set in PatternFlowIpv6HopLimit.
	// PatternFlowIpv6HopLimitCounter is integer counter pattern
	Decrement() PatternFlowIpv6HopLimitCounter
	// SetDecrement assigns PatternFlowIpv6HopLimitCounter provided by user to PatternFlowIpv6HopLimit.
	// PatternFlowIpv6HopLimitCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimit
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6HopLimit
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIterIter, set in PatternFlowIpv6HopLimit
	MetricTags() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	setNil()
}

type PatternFlowIpv6HopLimitChoiceEnum string

// Enum of Choice on PatternFlowIpv6HopLimit
var PatternFlowIpv6HopLimitChoice = struct {
	VALUE     PatternFlowIpv6HopLimitChoiceEnum
	VALUES    PatternFlowIpv6HopLimitChoiceEnum
	INCREMENT PatternFlowIpv6HopLimitChoiceEnum
	DECREMENT PatternFlowIpv6HopLimitChoiceEnum
}{
	VALUE:     PatternFlowIpv6HopLimitChoiceEnum("value"),
	VALUES:    PatternFlowIpv6HopLimitChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6HopLimitChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6HopLimitChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6HopLimit) Choice() PatternFlowIpv6HopLimitChoiceEnum {
	return PatternFlowIpv6HopLimitChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6HopLimit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6HopLimit) setChoice(value PatternFlowIpv6HopLimitChoiceEnum) PatternFlowIpv6HopLimit {
	intValue, ok := otg.PatternFlowIpv6HopLimit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6HopLimitChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6HopLimit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6HopLimitChoice.VALUE {
		defaultValue := uint32(64)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6HopLimitChoice.VALUES {
		defaultValue := []uint32{64}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6HopLimitChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6HopLimitCounter().msg()
	}

	if value == PatternFlowIpv6HopLimitChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6HopLimitCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6HopLimit) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6HopLimitChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6HopLimit) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6HopLimit object
func (obj *patternFlowIpv6HopLimit) SetValue(value uint32) PatternFlowIpv6HopLimit {
	obj.setChoice(PatternFlowIpv6HopLimitChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6HopLimit) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{64})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6HopLimit object
func (obj *patternFlowIpv6HopLimit) SetValues(value []uint32) PatternFlowIpv6HopLimit {
	obj.setChoice(PatternFlowIpv6HopLimitChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6HopLimitCounter
func (obj *patternFlowIpv6HopLimit) Increment() PatternFlowIpv6HopLimitCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6HopLimitChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6HopLimitCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6HopLimitCounter
func (obj *patternFlowIpv6HopLimit) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6HopLimitCounter value in the PatternFlowIpv6HopLimit object
func (obj *patternFlowIpv6HopLimit) SetIncrement(value PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimit {
	obj.setChoice(PatternFlowIpv6HopLimitChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6HopLimitCounter
func (obj *patternFlowIpv6HopLimit) Decrement() PatternFlowIpv6HopLimitCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6HopLimitChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6HopLimitCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6HopLimitCounter
func (obj *patternFlowIpv6HopLimit) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6HopLimitCounter value in the PatternFlowIpv6HopLimit object
func (obj *patternFlowIpv6HopLimit) SetDecrement(value PatternFlowIpv6HopLimitCounter) PatternFlowIpv6HopLimit {
	obj.setChoice(PatternFlowIpv6HopLimitChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6HopLimitMetricTag
func (obj *patternFlowIpv6HopLimit) MetricTags() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6HopLimitMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter struct {
	obj                                   *patternFlowIpv6HopLimit
	patternFlowIpv6HopLimitMetricTagSlice []PatternFlowIpv6HopLimitMetricTag
	fieldPtr                              *[]*otg.PatternFlowIpv6HopLimitMetricTag
}

func newPatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter(ptr *[]*otg.PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	return &patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter interface {
	setMsg(*patternFlowIpv6HopLimit) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	Items() []PatternFlowIpv6HopLimitMetricTag
	Add() PatternFlowIpv6HopLimitMetricTag
	Append(items ...PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	Set(index int, newObj PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	Clear() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	clearHolderSlice() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
	appendHolderSlice(item PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter
}

func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) setMsg(msg *patternFlowIpv6HopLimit) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6HopLimitMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) Items() []PatternFlowIpv6HopLimitMetricTag {
	return obj.patternFlowIpv6HopLimitMetricTagSlice
}

func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) Add() PatternFlowIpv6HopLimitMetricTag {
	newObj := &otg.PatternFlowIpv6HopLimitMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6HopLimitMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6HopLimitMetricTagSlice = append(obj.patternFlowIpv6HopLimitMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) Append(items ...PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6HopLimitMetricTagSlice = append(obj.patternFlowIpv6HopLimitMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) Set(index int, newObj PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6HopLimitMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) Clear() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6HopLimitMetricTag{}
		obj.patternFlowIpv6HopLimitMetricTagSlice = []PatternFlowIpv6HopLimitMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) clearHolderSlice() PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	if len(obj.patternFlowIpv6HopLimitMetricTagSlice) > 0 {
		obj.patternFlowIpv6HopLimitMetricTagSlice = []PatternFlowIpv6HopLimitMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter) appendHolderSlice(item PatternFlowIpv6HopLimitMetricTag) PatternFlowIpv6HopLimitPatternFlowIpv6HopLimitMetricTagIter {
	obj.patternFlowIpv6HopLimitMetricTagSlice = append(obj.patternFlowIpv6HopLimitMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6HopLimit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6HopLimit.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6HopLimit.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6HopLimitMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6HopLimit) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6HopLimitChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6HopLimitChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6HopLimitChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6HopLimitChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6HopLimitChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6HopLimitChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6HopLimit")
			}
		} else {
			intVal := otg.PatternFlowIpv6HopLimit_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6HopLimit_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
