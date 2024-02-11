package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1Unused *****
type patternFlowIgmpv1Unused struct {
	validation
	obj              *otg.PatternFlowIgmpv1Unused
	marshaller       marshalPatternFlowIgmpv1Unused
	unMarshaller     unMarshalPatternFlowIgmpv1Unused
	incrementHolder  PatternFlowIgmpv1UnusedCounter
	decrementHolder  PatternFlowIgmpv1UnusedCounter
	metricTagsHolder PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
}

func NewPatternFlowIgmpv1Unused() PatternFlowIgmpv1Unused {
	obj := patternFlowIgmpv1Unused{obj: &otg.PatternFlowIgmpv1Unused{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1Unused) msg() *otg.PatternFlowIgmpv1Unused {
	return obj.obj
}

func (obj *patternFlowIgmpv1Unused) setMsg(msg *otg.PatternFlowIgmpv1Unused) PatternFlowIgmpv1Unused {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1Unused struct {
	obj *patternFlowIgmpv1Unused
}

type marshalPatternFlowIgmpv1Unused interface {
	// ToProto marshals PatternFlowIgmpv1Unused to protobuf object *otg.PatternFlowIgmpv1Unused
	ToProto() (*otg.PatternFlowIgmpv1Unused, error)
	// ToPbText marshals PatternFlowIgmpv1Unused to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1Unused to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1Unused to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1Unused struct {
	obj *patternFlowIgmpv1Unused
}

type unMarshalPatternFlowIgmpv1Unused interface {
	// FromProto unmarshals PatternFlowIgmpv1Unused from protobuf object *otg.PatternFlowIgmpv1Unused
	FromProto(msg *otg.PatternFlowIgmpv1Unused) (PatternFlowIgmpv1Unused, error)
	// FromPbText unmarshals PatternFlowIgmpv1Unused from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1Unused from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1Unused from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1Unused) Marshal() marshalPatternFlowIgmpv1Unused {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1Unused{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1Unused) Unmarshal() unMarshalPatternFlowIgmpv1Unused {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1Unused{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1Unused) ToProto() (*otg.PatternFlowIgmpv1Unused, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1Unused) FromProto(msg *otg.PatternFlowIgmpv1Unused) (PatternFlowIgmpv1Unused, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1Unused) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Unused) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1Unused) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Unused) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1Unused) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Unused) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1Unused) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Unused) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Unused) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1Unused) Clone() (PatternFlowIgmpv1Unused, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1Unused()
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

func (obj *patternFlowIgmpv1Unused) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv1Unused is unused
type PatternFlowIgmpv1Unused interface {
	Validation
	// msg marshals PatternFlowIgmpv1Unused to protobuf object *otg.PatternFlowIgmpv1Unused
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1Unused
	// setMsg unmarshals PatternFlowIgmpv1Unused from protobuf object *otg.PatternFlowIgmpv1Unused
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1Unused) PatternFlowIgmpv1Unused
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1Unused
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1Unused
	// validate validates PatternFlowIgmpv1Unused
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1Unused, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv1UnusedChoiceEnum, set in PatternFlowIgmpv1Unused
	Choice() PatternFlowIgmpv1UnusedChoiceEnum
	// setChoice assigns PatternFlowIgmpv1UnusedChoiceEnum provided by user to PatternFlowIgmpv1Unused
	setChoice(value PatternFlowIgmpv1UnusedChoiceEnum) PatternFlowIgmpv1Unused
	// HasChoice checks if Choice has been set in PatternFlowIgmpv1Unused
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv1Unused.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv1Unused
	SetValue(value uint32) PatternFlowIgmpv1Unused
	// HasValue checks if Value has been set in PatternFlowIgmpv1Unused
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv1Unused.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv1Unused
	SetValues(value []uint32) PatternFlowIgmpv1Unused
	// Increment returns PatternFlowIgmpv1UnusedCounter, set in PatternFlowIgmpv1Unused.
	// PatternFlowIgmpv1UnusedCounter is integer counter pattern
	Increment() PatternFlowIgmpv1UnusedCounter
	// SetIncrement assigns PatternFlowIgmpv1UnusedCounter provided by user to PatternFlowIgmpv1Unused.
	// PatternFlowIgmpv1UnusedCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1Unused
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv1Unused
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv1UnusedCounter, set in PatternFlowIgmpv1Unused.
	// PatternFlowIgmpv1UnusedCounter is integer counter pattern
	Decrement() PatternFlowIgmpv1UnusedCounter
	// SetDecrement assigns PatternFlowIgmpv1UnusedCounter provided by user to PatternFlowIgmpv1Unused.
	// PatternFlowIgmpv1UnusedCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1Unused
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv1Unused
	HasDecrement() bool
	// MetricTags returns PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIterIter, set in PatternFlowIgmpv1Unused
	MetricTags() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	setNil()
}

type PatternFlowIgmpv1UnusedChoiceEnum string

// Enum of Choice on PatternFlowIgmpv1Unused
var PatternFlowIgmpv1UnusedChoice = struct {
	VALUE     PatternFlowIgmpv1UnusedChoiceEnum
	VALUES    PatternFlowIgmpv1UnusedChoiceEnum
	INCREMENT PatternFlowIgmpv1UnusedChoiceEnum
	DECREMENT PatternFlowIgmpv1UnusedChoiceEnum
}{
	VALUE:     PatternFlowIgmpv1UnusedChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv1UnusedChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv1UnusedChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv1UnusedChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv1Unused) Choice() PatternFlowIgmpv1UnusedChoiceEnum {
	return PatternFlowIgmpv1UnusedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv1Unused) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv1Unused) setChoice(value PatternFlowIgmpv1UnusedChoiceEnum) PatternFlowIgmpv1Unused {
	intValue, ok := otg.PatternFlowIgmpv1Unused_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1UnusedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1Unused_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv1UnusedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv1UnusedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv1UnusedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv1UnusedCounter().msg()
	}

	if value == PatternFlowIgmpv1UnusedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv1UnusedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Unused) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv1UnusedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Unused) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv1Unused object
func (obj *patternFlowIgmpv1Unused) SetValue(value uint32) PatternFlowIgmpv1Unused {
	obj.setChoice(PatternFlowIgmpv1UnusedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv1Unused) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv1Unused object
func (obj *patternFlowIgmpv1Unused) SetValues(value []uint32) PatternFlowIgmpv1Unused {
	obj.setChoice(PatternFlowIgmpv1UnusedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv1UnusedCounter
func (obj *patternFlowIgmpv1Unused) Increment() PatternFlowIgmpv1UnusedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv1UnusedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv1UnusedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv1UnusedCounter
func (obj *patternFlowIgmpv1Unused) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv1UnusedCounter value in the PatternFlowIgmpv1Unused object
func (obj *patternFlowIgmpv1Unused) SetIncrement(value PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1Unused {
	obj.setChoice(PatternFlowIgmpv1UnusedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1UnusedCounter
func (obj *patternFlowIgmpv1Unused) Decrement() PatternFlowIgmpv1UnusedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv1UnusedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv1UnusedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1UnusedCounter
func (obj *patternFlowIgmpv1Unused) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv1UnusedCounter value in the PatternFlowIgmpv1Unused object
func (obj *patternFlowIgmpv1Unused) SetDecrement(value PatternFlowIgmpv1UnusedCounter) PatternFlowIgmpv1Unused {
	obj.setChoice(PatternFlowIgmpv1UnusedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIgmpv1UnusedMetricTag
func (obj *patternFlowIgmpv1Unused) MetricTags() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIgmpv1UnusedMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter struct {
	obj                                   *patternFlowIgmpv1Unused
	patternFlowIgmpv1UnusedMetricTagSlice []PatternFlowIgmpv1UnusedMetricTag
	fieldPtr                              *[]*otg.PatternFlowIgmpv1UnusedMetricTag
}

func newPatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter(ptr *[]*otg.PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	return &patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter interface {
	setMsg(*patternFlowIgmpv1Unused) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	Items() []PatternFlowIgmpv1UnusedMetricTag
	Add() PatternFlowIgmpv1UnusedMetricTag
	Append(items ...PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	Set(index int, newObj PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	Clear() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	clearHolderSlice() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
	appendHolderSlice(item PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter
}

func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) setMsg(msg *patternFlowIgmpv1Unused) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIgmpv1UnusedMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) Items() []PatternFlowIgmpv1UnusedMetricTag {
	return obj.patternFlowIgmpv1UnusedMetricTagSlice
}

func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) Add() PatternFlowIgmpv1UnusedMetricTag {
	newObj := &otg.PatternFlowIgmpv1UnusedMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIgmpv1UnusedMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIgmpv1UnusedMetricTagSlice = append(obj.patternFlowIgmpv1UnusedMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) Append(items ...PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIgmpv1UnusedMetricTagSlice = append(obj.patternFlowIgmpv1UnusedMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) Set(index int, newObj PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIgmpv1UnusedMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) Clear() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIgmpv1UnusedMetricTag{}
		obj.patternFlowIgmpv1UnusedMetricTagSlice = []PatternFlowIgmpv1UnusedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) clearHolderSlice() PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	if len(obj.patternFlowIgmpv1UnusedMetricTagSlice) > 0 {
		obj.patternFlowIgmpv1UnusedMetricTagSlice = []PatternFlowIgmpv1UnusedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter) appendHolderSlice(item PatternFlowIgmpv1UnusedMetricTag) PatternFlowIgmpv1UnusedPatternFlowIgmpv1UnusedMetricTagIter {
	obj.patternFlowIgmpv1UnusedMetricTagSlice = append(obj.patternFlowIgmpv1UnusedMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIgmpv1Unused) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Unused.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv1Unused.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIgmpv1UnusedMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIgmpv1Unused) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIgmpv1UnusedChoice.VALUE)

	}

}
