package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosUnused *****
type patternFlowIpv4TosUnused struct {
	validation
	obj              *otg.PatternFlowIpv4TosUnused
	marshaller       marshalPatternFlowIpv4TosUnused
	unMarshaller     unMarshalPatternFlowIpv4TosUnused
	incrementHolder  PatternFlowIpv4TosUnusedCounter
	decrementHolder  PatternFlowIpv4TosUnusedCounter
	metricTagsHolder PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
}

func NewPatternFlowIpv4TosUnused() PatternFlowIpv4TosUnused {
	obj := patternFlowIpv4TosUnused{obj: &otg.PatternFlowIpv4TosUnused{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosUnused) msg() *otg.PatternFlowIpv4TosUnused {
	return obj.obj
}

func (obj *patternFlowIpv4TosUnused) setMsg(msg *otg.PatternFlowIpv4TosUnused) PatternFlowIpv4TosUnused {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosUnused struct {
	obj *patternFlowIpv4TosUnused
}

type marshalPatternFlowIpv4TosUnused interface {
	// ToProto marshals PatternFlowIpv4TosUnused to protobuf object *otg.PatternFlowIpv4TosUnused
	ToProto() (*otg.PatternFlowIpv4TosUnused, error)
	// ToPbText marshals PatternFlowIpv4TosUnused to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosUnused to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosUnused to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosUnused struct {
	obj *patternFlowIpv4TosUnused
}

type unMarshalPatternFlowIpv4TosUnused interface {
	// FromProto unmarshals PatternFlowIpv4TosUnused from protobuf object *otg.PatternFlowIpv4TosUnused
	FromProto(msg *otg.PatternFlowIpv4TosUnused) (PatternFlowIpv4TosUnused, error)
	// FromPbText unmarshals PatternFlowIpv4TosUnused from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosUnused from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosUnused from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosUnused) Marshal() marshalPatternFlowIpv4TosUnused {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosUnused{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosUnused) Unmarshal() unMarshalPatternFlowIpv4TosUnused {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosUnused{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosUnused) ToProto() (*otg.PatternFlowIpv4TosUnused, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosUnused) FromProto(msg *otg.PatternFlowIpv4TosUnused) (PatternFlowIpv4TosUnused, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosUnused) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnused) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnused) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnused) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosUnused) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosUnused) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosUnused) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnused) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosUnused) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosUnused) Clone() (PatternFlowIpv4TosUnused, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosUnused()
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

func (obj *patternFlowIpv4TosUnused) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosUnused is unused
type PatternFlowIpv4TosUnused interface {
	Validation
	// msg marshals PatternFlowIpv4TosUnused to protobuf object *otg.PatternFlowIpv4TosUnused
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosUnused
	// setMsg unmarshals PatternFlowIpv4TosUnused from protobuf object *otg.PatternFlowIpv4TosUnused
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosUnused) PatternFlowIpv4TosUnused
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosUnused
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosUnused
	// validate validates PatternFlowIpv4TosUnused
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosUnused, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosUnusedChoiceEnum, set in PatternFlowIpv4TosUnused
	Choice() PatternFlowIpv4TosUnusedChoiceEnum
	// setChoice assigns PatternFlowIpv4TosUnusedChoiceEnum provided by user to PatternFlowIpv4TosUnused
	setChoice(value PatternFlowIpv4TosUnusedChoiceEnum) PatternFlowIpv4TosUnused
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosUnused
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosUnused.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosUnused
	SetValue(value uint32) PatternFlowIpv4TosUnused
	// HasValue checks if Value has been set in PatternFlowIpv4TosUnused
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosUnused.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosUnused
	SetValues(value []uint32) PatternFlowIpv4TosUnused
	// Increment returns PatternFlowIpv4TosUnusedCounter, set in PatternFlowIpv4TosUnused.
	// PatternFlowIpv4TosUnusedCounter is integer counter pattern
	Increment() PatternFlowIpv4TosUnusedCounter
	// SetIncrement assigns PatternFlowIpv4TosUnusedCounter provided by user to PatternFlowIpv4TosUnused.
	// PatternFlowIpv4TosUnusedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnused
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosUnused
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosUnusedCounter, set in PatternFlowIpv4TosUnused.
	// PatternFlowIpv4TosUnusedCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosUnusedCounter
	// SetDecrement assigns PatternFlowIpv4TosUnusedCounter provided by user to PatternFlowIpv4TosUnused.
	// PatternFlowIpv4TosUnusedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnused
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosUnused
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIterIter, set in PatternFlowIpv4TosUnused
	MetricTags() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	setNil()
}

type PatternFlowIpv4TosUnusedChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosUnused
var PatternFlowIpv4TosUnusedChoice = struct {
	VALUE     PatternFlowIpv4TosUnusedChoiceEnum
	VALUES    PatternFlowIpv4TosUnusedChoiceEnum
	INCREMENT PatternFlowIpv4TosUnusedChoiceEnum
	DECREMENT PatternFlowIpv4TosUnusedChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosUnusedChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosUnusedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosUnusedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosUnusedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosUnused) Choice() PatternFlowIpv4TosUnusedChoiceEnum {
	return PatternFlowIpv4TosUnusedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosUnused) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosUnused) setChoice(value PatternFlowIpv4TosUnusedChoiceEnum) PatternFlowIpv4TosUnused {
	intValue, ok := otg.PatternFlowIpv4TosUnused_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosUnusedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosUnused_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosUnusedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosUnusedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosUnusedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosUnusedCounter().msg()
	}

	if value == PatternFlowIpv4TosUnusedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosUnusedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosUnused) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosUnusedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosUnused) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosUnused object
func (obj *patternFlowIpv4TosUnused) SetValue(value uint32) PatternFlowIpv4TosUnused {
	obj.setChoice(PatternFlowIpv4TosUnusedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosUnused) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosUnused object
func (obj *patternFlowIpv4TosUnused) SetValues(value []uint32) PatternFlowIpv4TosUnused {
	obj.setChoice(PatternFlowIpv4TosUnusedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosUnusedCounter
func (obj *patternFlowIpv4TosUnused) Increment() PatternFlowIpv4TosUnusedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosUnusedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosUnusedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosUnusedCounter
func (obj *patternFlowIpv4TosUnused) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosUnusedCounter value in the PatternFlowIpv4TosUnused object
func (obj *patternFlowIpv4TosUnused) SetIncrement(value PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnused {
	obj.setChoice(PatternFlowIpv4TosUnusedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosUnusedCounter
func (obj *patternFlowIpv4TosUnused) Decrement() PatternFlowIpv4TosUnusedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosUnusedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosUnusedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosUnusedCounter
func (obj *patternFlowIpv4TosUnused) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosUnusedCounter value in the PatternFlowIpv4TosUnused object
func (obj *patternFlowIpv4TosUnused) SetDecrement(value PatternFlowIpv4TosUnusedCounter) PatternFlowIpv4TosUnused {
	obj.setChoice(PatternFlowIpv4TosUnusedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosUnusedMetricTag
func (obj *patternFlowIpv4TosUnused) MetricTags() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosUnusedMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter struct {
	obj                                    *patternFlowIpv4TosUnused
	patternFlowIpv4TosUnusedMetricTagSlice []PatternFlowIpv4TosUnusedMetricTag
	fieldPtr                               *[]*otg.PatternFlowIpv4TosUnusedMetricTag
}

func newPatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	return &patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter interface {
	setMsg(*patternFlowIpv4TosUnused) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	Items() []PatternFlowIpv4TosUnusedMetricTag
	Add() PatternFlowIpv4TosUnusedMetricTag
	Append(items ...PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	Clear() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter
}

func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) setMsg(msg *patternFlowIpv4TosUnused) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosUnusedMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) Items() []PatternFlowIpv4TosUnusedMetricTag {
	return obj.patternFlowIpv4TosUnusedMetricTagSlice
}

func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) Add() PatternFlowIpv4TosUnusedMetricTag {
	newObj := &otg.PatternFlowIpv4TosUnusedMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosUnusedMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosUnusedMetricTagSlice = append(obj.patternFlowIpv4TosUnusedMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) Append(items ...PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosUnusedMetricTagSlice = append(obj.patternFlowIpv4TosUnusedMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) Set(index int, newObj PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosUnusedMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) Clear() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosUnusedMetricTag{}
		obj.patternFlowIpv4TosUnusedMetricTagSlice = []PatternFlowIpv4TosUnusedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) clearHolderSlice() PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	if len(obj.patternFlowIpv4TosUnusedMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosUnusedMetricTagSlice = []PatternFlowIpv4TosUnusedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosUnusedMetricTag) PatternFlowIpv4TosUnusedPatternFlowIpv4TosUnusedMetricTagIter {
	obj.patternFlowIpv4TosUnusedMetricTagSlice = append(obj.patternFlowIpv4TosUnusedMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosUnused) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosUnused.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosUnused.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosUnusedMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosUnused) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4TosUnusedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosUnusedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4TosUnusedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosUnusedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosUnusedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4TosUnusedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4TosUnused")
			}
		} else {
			intVal := otg.PatternFlowIpv4TosUnused_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4TosUnused_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
