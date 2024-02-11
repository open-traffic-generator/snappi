package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TimeToLive *****
type patternFlowIpv4TimeToLive struct {
	validation
	obj              *otg.PatternFlowIpv4TimeToLive
	marshaller       marshalPatternFlowIpv4TimeToLive
	unMarshaller     unMarshalPatternFlowIpv4TimeToLive
	incrementHolder  PatternFlowIpv4TimeToLiveCounter
	decrementHolder  PatternFlowIpv4TimeToLiveCounter
	metricTagsHolder PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
}

func NewPatternFlowIpv4TimeToLive() PatternFlowIpv4TimeToLive {
	obj := patternFlowIpv4TimeToLive{obj: &otg.PatternFlowIpv4TimeToLive{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TimeToLive) msg() *otg.PatternFlowIpv4TimeToLive {
	return obj.obj
}

func (obj *patternFlowIpv4TimeToLive) setMsg(msg *otg.PatternFlowIpv4TimeToLive) PatternFlowIpv4TimeToLive {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TimeToLive struct {
	obj *patternFlowIpv4TimeToLive
}

type marshalPatternFlowIpv4TimeToLive interface {
	// ToProto marshals PatternFlowIpv4TimeToLive to protobuf object *otg.PatternFlowIpv4TimeToLive
	ToProto() (*otg.PatternFlowIpv4TimeToLive, error)
	// ToPbText marshals PatternFlowIpv4TimeToLive to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TimeToLive to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TimeToLive to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TimeToLive struct {
	obj *patternFlowIpv4TimeToLive
}

type unMarshalPatternFlowIpv4TimeToLive interface {
	// FromProto unmarshals PatternFlowIpv4TimeToLive from protobuf object *otg.PatternFlowIpv4TimeToLive
	FromProto(msg *otg.PatternFlowIpv4TimeToLive) (PatternFlowIpv4TimeToLive, error)
	// FromPbText unmarshals PatternFlowIpv4TimeToLive from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TimeToLive from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TimeToLive from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TimeToLive) Marshal() marshalPatternFlowIpv4TimeToLive {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TimeToLive{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TimeToLive) Unmarshal() unMarshalPatternFlowIpv4TimeToLive {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TimeToLive{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TimeToLive) ToProto() (*otg.PatternFlowIpv4TimeToLive, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TimeToLive) FromProto(msg *otg.PatternFlowIpv4TimeToLive) (PatternFlowIpv4TimeToLive, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TimeToLive) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLive) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLive) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLive) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TimeToLive) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TimeToLive) FromJson(value string) error {
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

func (obj *patternFlowIpv4TimeToLive) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLive) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TimeToLive) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TimeToLive) Clone() (PatternFlowIpv4TimeToLive, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TimeToLive()
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

func (obj *patternFlowIpv4TimeToLive) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TimeToLive is time to live
type PatternFlowIpv4TimeToLive interface {
	Validation
	// msg marshals PatternFlowIpv4TimeToLive to protobuf object *otg.PatternFlowIpv4TimeToLive
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TimeToLive
	// setMsg unmarshals PatternFlowIpv4TimeToLive from protobuf object *otg.PatternFlowIpv4TimeToLive
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TimeToLive) PatternFlowIpv4TimeToLive
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TimeToLive
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TimeToLive
	// validate validates PatternFlowIpv4TimeToLive
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TimeToLive, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TimeToLiveChoiceEnum, set in PatternFlowIpv4TimeToLive
	Choice() PatternFlowIpv4TimeToLiveChoiceEnum
	// setChoice assigns PatternFlowIpv4TimeToLiveChoiceEnum provided by user to PatternFlowIpv4TimeToLive
	setChoice(value PatternFlowIpv4TimeToLiveChoiceEnum) PatternFlowIpv4TimeToLive
	// HasChoice checks if Choice has been set in PatternFlowIpv4TimeToLive
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TimeToLive.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TimeToLive
	SetValue(value uint32) PatternFlowIpv4TimeToLive
	// HasValue checks if Value has been set in PatternFlowIpv4TimeToLive
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TimeToLive.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TimeToLive
	SetValues(value []uint32) PatternFlowIpv4TimeToLive
	// Increment returns PatternFlowIpv4TimeToLiveCounter, set in PatternFlowIpv4TimeToLive.
	// PatternFlowIpv4TimeToLiveCounter is integer counter pattern
	Increment() PatternFlowIpv4TimeToLiveCounter
	// SetIncrement assigns PatternFlowIpv4TimeToLiveCounter provided by user to PatternFlowIpv4TimeToLive.
	// PatternFlowIpv4TimeToLiveCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLive
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TimeToLive
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TimeToLiveCounter, set in PatternFlowIpv4TimeToLive.
	// PatternFlowIpv4TimeToLiveCounter is integer counter pattern
	Decrement() PatternFlowIpv4TimeToLiveCounter
	// SetDecrement assigns PatternFlowIpv4TimeToLiveCounter provided by user to PatternFlowIpv4TimeToLive.
	// PatternFlowIpv4TimeToLiveCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLive
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TimeToLive
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIterIter, set in PatternFlowIpv4TimeToLive
	MetricTags() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	setNil()
}

type PatternFlowIpv4TimeToLiveChoiceEnum string

// Enum of Choice on PatternFlowIpv4TimeToLive
var PatternFlowIpv4TimeToLiveChoice = struct {
	VALUE     PatternFlowIpv4TimeToLiveChoiceEnum
	VALUES    PatternFlowIpv4TimeToLiveChoiceEnum
	INCREMENT PatternFlowIpv4TimeToLiveChoiceEnum
	DECREMENT PatternFlowIpv4TimeToLiveChoiceEnum
}{
	VALUE:     PatternFlowIpv4TimeToLiveChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TimeToLiveChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TimeToLiveChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TimeToLiveChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TimeToLive) Choice() PatternFlowIpv4TimeToLiveChoiceEnum {
	return PatternFlowIpv4TimeToLiveChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TimeToLive) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TimeToLive) setChoice(value PatternFlowIpv4TimeToLiveChoiceEnum) PatternFlowIpv4TimeToLive {
	intValue, ok := otg.PatternFlowIpv4TimeToLive_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TimeToLiveChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TimeToLive_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TimeToLiveChoice.VALUE {
		defaultValue := uint32(64)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TimeToLiveChoice.VALUES {
		defaultValue := []uint32{64}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TimeToLiveChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TimeToLiveCounter().msg()
	}

	if value == PatternFlowIpv4TimeToLiveChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TimeToLiveCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TimeToLive) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TimeToLiveChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TimeToLive) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TimeToLive object
func (obj *patternFlowIpv4TimeToLive) SetValue(value uint32) PatternFlowIpv4TimeToLive {
	obj.setChoice(PatternFlowIpv4TimeToLiveChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TimeToLive) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{64})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TimeToLive object
func (obj *patternFlowIpv4TimeToLive) SetValues(value []uint32) PatternFlowIpv4TimeToLive {
	obj.setChoice(PatternFlowIpv4TimeToLiveChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TimeToLiveCounter
func (obj *patternFlowIpv4TimeToLive) Increment() PatternFlowIpv4TimeToLiveCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TimeToLiveChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TimeToLiveCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TimeToLiveCounter
func (obj *patternFlowIpv4TimeToLive) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TimeToLiveCounter value in the PatternFlowIpv4TimeToLive object
func (obj *patternFlowIpv4TimeToLive) SetIncrement(value PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLive {
	obj.setChoice(PatternFlowIpv4TimeToLiveChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TimeToLiveCounter
func (obj *patternFlowIpv4TimeToLive) Decrement() PatternFlowIpv4TimeToLiveCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TimeToLiveChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TimeToLiveCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TimeToLiveCounter
func (obj *patternFlowIpv4TimeToLive) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TimeToLiveCounter value in the PatternFlowIpv4TimeToLive object
func (obj *patternFlowIpv4TimeToLive) SetDecrement(value PatternFlowIpv4TimeToLiveCounter) PatternFlowIpv4TimeToLive {
	obj.setChoice(PatternFlowIpv4TimeToLiveChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TimeToLiveMetricTag
func (obj *patternFlowIpv4TimeToLive) MetricTags() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TimeToLiveMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter struct {
	obj                                     *patternFlowIpv4TimeToLive
	patternFlowIpv4TimeToLiveMetricTagSlice []PatternFlowIpv4TimeToLiveMetricTag
	fieldPtr                                *[]*otg.PatternFlowIpv4TimeToLiveMetricTag
}

func newPatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter(ptr *[]*otg.PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	return &patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter interface {
	setMsg(*patternFlowIpv4TimeToLive) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	Items() []PatternFlowIpv4TimeToLiveMetricTag
	Add() PatternFlowIpv4TimeToLiveMetricTag
	Append(items ...PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	Set(index int, newObj PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	Clear() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	clearHolderSlice() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter
}

func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) setMsg(msg *patternFlowIpv4TimeToLive) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TimeToLiveMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) Items() []PatternFlowIpv4TimeToLiveMetricTag {
	return obj.patternFlowIpv4TimeToLiveMetricTagSlice
}

func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) Add() PatternFlowIpv4TimeToLiveMetricTag {
	newObj := &otg.PatternFlowIpv4TimeToLiveMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TimeToLiveMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TimeToLiveMetricTagSlice = append(obj.patternFlowIpv4TimeToLiveMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) Append(items ...PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TimeToLiveMetricTagSlice = append(obj.patternFlowIpv4TimeToLiveMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) Set(index int, newObj PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TimeToLiveMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) Clear() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TimeToLiveMetricTag{}
		obj.patternFlowIpv4TimeToLiveMetricTagSlice = []PatternFlowIpv4TimeToLiveMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) clearHolderSlice() PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	if len(obj.patternFlowIpv4TimeToLiveMetricTagSlice) > 0 {
		obj.patternFlowIpv4TimeToLiveMetricTagSlice = []PatternFlowIpv4TimeToLiveMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter) appendHolderSlice(item PatternFlowIpv4TimeToLiveMetricTag) PatternFlowIpv4TimeToLivePatternFlowIpv4TimeToLiveMetricTagIter {
	obj.patternFlowIpv4TimeToLiveMetricTagSlice = append(obj.patternFlowIpv4TimeToLiveMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TimeToLive) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TimeToLive.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TimeToLive.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TimeToLiveMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TimeToLive) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv4TimeToLiveChoice.VALUE)

	}

}
