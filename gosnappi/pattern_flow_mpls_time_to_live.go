package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsTimeToLive *****
type patternFlowMplsTimeToLive struct {
	validation
	obj              *otg.PatternFlowMplsTimeToLive
	marshaller       marshalPatternFlowMplsTimeToLive
	unMarshaller     unMarshalPatternFlowMplsTimeToLive
	incrementHolder  PatternFlowMplsTimeToLiveCounter
	decrementHolder  PatternFlowMplsTimeToLiveCounter
	metricTagsHolder PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
}

func NewPatternFlowMplsTimeToLive() PatternFlowMplsTimeToLive {
	obj := patternFlowMplsTimeToLive{obj: &otg.PatternFlowMplsTimeToLive{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTimeToLive) msg() *otg.PatternFlowMplsTimeToLive {
	return obj.obj
}

func (obj *patternFlowMplsTimeToLive) setMsg(msg *otg.PatternFlowMplsTimeToLive) PatternFlowMplsTimeToLive {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTimeToLive struct {
	obj *patternFlowMplsTimeToLive
}

type marshalPatternFlowMplsTimeToLive interface {
	// ToProto marshals PatternFlowMplsTimeToLive to protobuf object *otg.PatternFlowMplsTimeToLive
	ToProto() (*otg.PatternFlowMplsTimeToLive, error)
	// ToPbText marshals PatternFlowMplsTimeToLive to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTimeToLive to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTimeToLive to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowMplsTimeToLive struct {
	obj *patternFlowMplsTimeToLive
}

type unMarshalPatternFlowMplsTimeToLive interface {
	// FromProto unmarshals PatternFlowMplsTimeToLive from protobuf object *otg.PatternFlowMplsTimeToLive
	FromProto(msg *otg.PatternFlowMplsTimeToLive) (PatternFlowMplsTimeToLive, error)
	// FromPbText unmarshals PatternFlowMplsTimeToLive from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTimeToLive from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTimeToLive from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTimeToLive) Marshal() marshalPatternFlowMplsTimeToLive {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTimeToLive{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTimeToLive) Unmarshal() unMarshalPatternFlowMplsTimeToLive {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTimeToLive{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTimeToLive) ToProto() (*otg.PatternFlowMplsTimeToLive, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTimeToLive) FromProto(msg *otg.PatternFlowMplsTimeToLive) (PatternFlowMplsTimeToLive, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTimeToLive) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLive) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLive) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLive) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTimeToLive) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTimeToLive) FromJson(value string) error {
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

func (obj *patternFlowMplsTimeToLive) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLive) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTimeToLive) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTimeToLive) Clone() (PatternFlowMplsTimeToLive, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTimeToLive()
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

func (obj *patternFlowMplsTimeToLive) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowMplsTimeToLive is time to live
type PatternFlowMplsTimeToLive interface {
	Validation
	// msg marshals PatternFlowMplsTimeToLive to protobuf object *otg.PatternFlowMplsTimeToLive
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTimeToLive
	// setMsg unmarshals PatternFlowMplsTimeToLive from protobuf object *otg.PatternFlowMplsTimeToLive
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTimeToLive) PatternFlowMplsTimeToLive
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTimeToLive
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTimeToLive
	// validate validates PatternFlowMplsTimeToLive
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTimeToLive, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowMplsTimeToLiveChoiceEnum, set in PatternFlowMplsTimeToLive
	Choice() PatternFlowMplsTimeToLiveChoiceEnum
	// setChoice assigns PatternFlowMplsTimeToLiveChoiceEnum provided by user to PatternFlowMplsTimeToLive
	setChoice(value PatternFlowMplsTimeToLiveChoiceEnum) PatternFlowMplsTimeToLive
	// HasChoice checks if Choice has been set in PatternFlowMplsTimeToLive
	HasChoice() bool
	// Value returns uint32, set in PatternFlowMplsTimeToLive.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowMplsTimeToLive
	SetValue(value uint32) PatternFlowMplsTimeToLive
	// HasValue checks if Value has been set in PatternFlowMplsTimeToLive
	HasValue() bool
	// Values returns []uint32, set in PatternFlowMplsTimeToLive.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowMplsTimeToLive
	SetValues(value []uint32) PatternFlowMplsTimeToLive
	// Increment returns PatternFlowMplsTimeToLiveCounter, set in PatternFlowMplsTimeToLive.
	// PatternFlowMplsTimeToLiveCounter is integer counter pattern
	Increment() PatternFlowMplsTimeToLiveCounter
	// SetIncrement assigns PatternFlowMplsTimeToLiveCounter provided by user to PatternFlowMplsTimeToLive.
	// PatternFlowMplsTimeToLiveCounter is integer counter pattern
	SetIncrement(value PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLive
	// HasIncrement checks if Increment has been set in PatternFlowMplsTimeToLive
	HasIncrement() bool
	// Decrement returns PatternFlowMplsTimeToLiveCounter, set in PatternFlowMplsTimeToLive.
	// PatternFlowMplsTimeToLiveCounter is integer counter pattern
	Decrement() PatternFlowMplsTimeToLiveCounter
	// SetDecrement assigns PatternFlowMplsTimeToLiveCounter provided by user to PatternFlowMplsTimeToLive.
	// PatternFlowMplsTimeToLiveCounter is integer counter pattern
	SetDecrement(value PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLive
	// HasDecrement checks if Decrement has been set in PatternFlowMplsTimeToLive
	HasDecrement() bool
	// MetricTags returns PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIterIter, set in PatternFlowMplsTimeToLive
	MetricTags() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	setNil()
}

type PatternFlowMplsTimeToLiveChoiceEnum string

// Enum of Choice on PatternFlowMplsTimeToLive
var PatternFlowMplsTimeToLiveChoice = struct {
	VALUE     PatternFlowMplsTimeToLiveChoiceEnum
	VALUES    PatternFlowMplsTimeToLiveChoiceEnum
	INCREMENT PatternFlowMplsTimeToLiveChoiceEnum
	DECREMENT PatternFlowMplsTimeToLiveChoiceEnum
}{
	VALUE:     PatternFlowMplsTimeToLiveChoiceEnum("value"),
	VALUES:    PatternFlowMplsTimeToLiveChoiceEnum("values"),
	INCREMENT: PatternFlowMplsTimeToLiveChoiceEnum("increment"),
	DECREMENT: PatternFlowMplsTimeToLiveChoiceEnum("decrement"),
}

func (obj *patternFlowMplsTimeToLive) Choice() PatternFlowMplsTimeToLiveChoiceEnum {
	return PatternFlowMplsTimeToLiveChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowMplsTimeToLive) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowMplsTimeToLive) setChoice(value PatternFlowMplsTimeToLiveChoiceEnum) PatternFlowMplsTimeToLive {
	intValue, ok := otg.PatternFlowMplsTimeToLive_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowMplsTimeToLiveChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowMplsTimeToLive_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowMplsTimeToLiveChoice.VALUE {
		defaultValue := uint32(64)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowMplsTimeToLiveChoice.VALUES {
		defaultValue := []uint32{64}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowMplsTimeToLiveChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowMplsTimeToLiveCounter().msg()
	}

	if value == PatternFlowMplsTimeToLiveChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowMplsTimeToLiveCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsTimeToLive) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowMplsTimeToLiveChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsTimeToLive) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowMplsTimeToLive object
func (obj *patternFlowMplsTimeToLive) SetValue(value uint32) PatternFlowMplsTimeToLive {
	obj.setChoice(PatternFlowMplsTimeToLiveChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowMplsTimeToLive) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{64})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowMplsTimeToLive object
func (obj *patternFlowMplsTimeToLive) SetValues(value []uint32) PatternFlowMplsTimeToLive {
	obj.setChoice(PatternFlowMplsTimeToLiveChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowMplsTimeToLiveCounter
func (obj *patternFlowMplsTimeToLive) Increment() PatternFlowMplsTimeToLiveCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowMplsTimeToLiveChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowMplsTimeToLiveCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowMplsTimeToLiveCounter
func (obj *patternFlowMplsTimeToLive) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowMplsTimeToLiveCounter value in the PatternFlowMplsTimeToLive object
func (obj *patternFlowMplsTimeToLive) SetIncrement(value PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLive {
	obj.setChoice(PatternFlowMplsTimeToLiveChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowMplsTimeToLiveCounter
func (obj *patternFlowMplsTimeToLive) Decrement() PatternFlowMplsTimeToLiveCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowMplsTimeToLiveChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowMplsTimeToLiveCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowMplsTimeToLiveCounter
func (obj *patternFlowMplsTimeToLive) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowMplsTimeToLiveCounter value in the PatternFlowMplsTimeToLive object
func (obj *patternFlowMplsTimeToLive) SetDecrement(value PatternFlowMplsTimeToLiveCounter) PatternFlowMplsTimeToLive {
	obj.setChoice(PatternFlowMplsTimeToLiveChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowMplsTimeToLiveMetricTag
func (obj *patternFlowMplsTimeToLive) MetricTags() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowMplsTimeToLiveMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter struct {
	obj                                     *patternFlowMplsTimeToLive
	patternFlowMplsTimeToLiveMetricTagSlice []PatternFlowMplsTimeToLiveMetricTag
	fieldPtr                                *[]*otg.PatternFlowMplsTimeToLiveMetricTag
}

func newPatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter(ptr *[]*otg.PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	return &patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter{fieldPtr: ptr}
}

type PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter interface {
	setMsg(*patternFlowMplsTimeToLive) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	Items() []PatternFlowMplsTimeToLiveMetricTag
	Add() PatternFlowMplsTimeToLiveMetricTag
	Append(items ...PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	Set(index int, newObj PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	Clear() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	clearHolderSlice() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
	appendHolderSlice(item PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter
}

func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) setMsg(msg *patternFlowMplsTimeToLive) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowMplsTimeToLiveMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) Items() []PatternFlowMplsTimeToLiveMetricTag {
	return obj.patternFlowMplsTimeToLiveMetricTagSlice
}

func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) Add() PatternFlowMplsTimeToLiveMetricTag {
	newObj := &otg.PatternFlowMplsTimeToLiveMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowMplsTimeToLiveMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowMplsTimeToLiveMetricTagSlice = append(obj.patternFlowMplsTimeToLiveMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) Append(items ...PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowMplsTimeToLiveMetricTagSlice = append(obj.patternFlowMplsTimeToLiveMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) Set(index int, newObj PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowMplsTimeToLiveMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) Clear() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowMplsTimeToLiveMetricTag{}
		obj.patternFlowMplsTimeToLiveMetricTagSlice = []PatternFlowMplsTimeToLiveMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) clearHolderSlice() PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	if len(obj.patternFlowMplsTimeToLiveMetricTagSlice) > 0 {
		obj.patternFlowMplsTimeToLiveMetricTagSlice = []PatternFlowMplsTimeToLiveMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter) appendHolderSlice(item PatternFlowMplsTimeToLiveMetricTag) PatternFlowMplsTimeToLivePatternFlowMplsTimeToLiveMetricTagIter {
	obj.patternFlowMplsTimeToLiveMetricTagSlice = append(obj.patternFlowMplsTimeToLiveMetricTagSlice, item)
	return obj
}

func (obj *patternFlowMplsTimeToLive) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTimeToLive.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowMplsTimeToLive.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowMplsTimeToLiveMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowMplsTimeToLive) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowMplsTimeToLiveChoice.VALUE)

	}

}
