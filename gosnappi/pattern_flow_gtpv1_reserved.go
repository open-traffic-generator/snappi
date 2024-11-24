package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1Reserved *****
type patternFlowGtpv1Reserved struct {
	validation
	obj              *otg.PatternFlowGtpv1Reserved
	marshaller       marshalPatternFlowGtpv1Reserved
	unMarshaller     unMarshalPatternFlowGtpv1Reserved
	incrementHolder  PatternFlowGtpv1ReservedCounter
	decrementHolder  PatternFlowGtpv1ReservedCounter
	metricTagsHolder PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
}

func NewPatternFlowGtpv1Reserved() PatternFlowGtpv1Reserved {
	obj := patternFlowGtpv1Reserved{obj: &otg.PatternFlowGtpv1Reserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1Reserved) msg() *otg.PatternFlowGtpv1Reserved {
	return obj.obj
}

func (obj *patternFlowGtpv1Reserved) setMsg(msg *otg.PatternFlowGtpv1Reserved) PatternFlowGtpv1Reserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1Reserved struct {
	obj *patternFlowGtpv1Reserved
}

type marshalPatternFlowGtpv1Reserved interface {
	// ToProto marshals PatternFlowGtpv1Reserved to protobuf object *otg.PatternFlowGtpv1Reserved
	ToProto() (*otg.PatternFlowGtpv1Reserved, error)
	// ToPbText marshals PatternFlowGtpv1Reserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1Reserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1Reserved to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1Reserved struct {
	obj *patternFlowGtpv1Reserved
}

type unMarshalPatternFlowGtpv1Reserved interface {
	// FromProto unmarshals PatternFlowGtpv1Reserved from protobuf object *otg.PatternFlowGtpv1Reserved
	FromProto(msg *otg.PatternFlowGtpv1Reserved) (PatternFlowGtpv1Reserved, error)
	// FromPbText unmarshals PatternFlowGtpv1Reserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1Reserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1Reserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1Reserved) Marshal() marshalPatternFlowGtpv1Reserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1Reserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1Reserved) Unmarshal() unMarshalPatternFlowGtpv1Reserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1Reserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1Reserved) ToProto() (*otg.PatternFlowGtpv1Reserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1Reserved) FromProto(msg *otg.PatternFlowGtpv1Reserved) (PatternFlowGtpv1Reserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1Reserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Reserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1Reserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Reserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1Reserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Reserved) FromJson(value string) error {
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

func (obj *patternFlowGtpv1Reserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Reserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Reserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1Reserved) Clone() (PatternFlowGtpv1Reserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1Reserved()
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

func (obj *patternFlowGtpv1Reserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1Reserved is reserved field
type PatternFlowGtpv1Reserved interface {
	Validation
	// msg marshals PatternFlowGtpv1Reserved to protobuf object *otg.PatternFlowGtpv1Reserved
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1Reserved
	// setMsg unmarshals PatternFlowGtpv1Reserved from protobuf object *otg.PatternFlowGtpv1Reserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1Reserved) PatternFlowGtpv1Reserved
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1Reserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1Reserved
	// validate validates PatternFlowGtpv1Reserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1Reserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1ReservedChoiceEnum, set in PatternFlowGtpv1Reserved
	Choice() PatternFlowGtpv1ReservedChoiceEnum
	// setChoice assigns PatternFlowGtpv1ReservedChoiceEnum provided by user to PatternFlowGtpv1Reserved
	setChoice(value PatternFlowGtpv1ReservedChoiceEnum) PatternFlowGtpv1Reserved
	// HasChoice checks if Choice has been set in PatternFlowGtpv1Reserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1Reserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1Reserved
	SetValue(value uint32) PatternFlowGtpv1Reserved
	// HasValue checks if Value has been set in PatternFlowGtpv1Reserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1Reserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1Reserved
	SetValues(value []uint32) PatternFlowGtpv1Reserved
	// Increment returns PatternFlowGtpv1ReservedCounter, set in PatternFlowGtpv1Reserved.
	// PatternFlowGtpv1ReservedCounter is integer counter pattern
	Increment() PatternFlowGtpv1ReservedCounter
	// SetIncrement assigns PatternFlowGtpv1ReservedCounter provided by user to PatternFlowGtpv1Reserved.
	// PatternFlowGtpv1ReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1Reserved
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1Reserved
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1ReservedCounter, set in PatternFlowGtpv1Reserved.
	// PatternFlowGtpv1ReservedCounter is integer counter pattern
	Decrement() PatternFlowGtpv1ReservedCounter
	// SetDecrement assigns PatternFlowGtpv1ReservedCounter provided by user to PatternFlowGtpv1Reserved.
	// PatternFlowGtpv1ReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1Reserved
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1Reserved
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIterIter, set in PatternFlowGtpv1Reserved
	MetricTags() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	setNil()
}

type PatternFlowGtpv1ReservedChoiceEnum string

// Enum of Choice on PatternFlowGtpv1Reserved
var PatternFlowGtpv1ReservedChoice = struct {
	VALUE     PatternFlowGtpv1ReservedChoiceEnum
	VALUES    PatternFlowGtpv1ReservedChoiceEnum
	INCREMENT PatternFlowGtpv1ReservedChoiceEnum
	DECREMENT PatternFlowGtpv1ReservedChoiceEnum
}{
	VALUE:     PatternFlowGtpv1ReservedChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1ReservedChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1ReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1ReservedChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1Reserved) Choice() PatternFlowGtpv1ReservedChoiceEnum {
	return PatternFlowGtpv1ReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1Reserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1Reserved) setChoice(value PatternFlowGtpv1ReservedChoiceEnum) PatternFlowGtpv1Reserved {
	intValue, ok := otg.PatternFlowGtpv1Reserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1ReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1Reserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1ReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1ReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1ReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1ReservedCounter().msg()
	}

	if value == PatternFlowGtpv1ReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1ReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Reserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1ReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Reserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1Reserved object
func (obj *patternFlowGtpv1Reserved) SetValue(value uint32) PatternFlowGtpv1Reserved {
	obj.setChoice(PatternFlowGtpv1ReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1Reserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1Reserved object
func (obj *patternFlowGtpv1Reserved) SetValues(value []uint32) PatternFlowGtpv1Reserved {
	obj.setChoice(PatternFlowGtpv1ReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1ReservedCounter
func (obj *patternFlowGtpv1Reserved) Increment() PatternFlowGtpv1ReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1ReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1ReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1ReservedCounter
func (obj *patternFlowGtpv1Reserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1ReservedCounter value in the PatternFlowGtpv1Reserved object
func (obj *patternFlowGtpv1Reserved) SetIncrement(value PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1Reserved {
	obj.setChoice(PatternFlowGtpv1ReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1ReservedCounter
func (obj *patternFlowGtpv1Reserved) Decrement() PatternFlowGtpv1ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1ReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1ReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1ReservedCounter
func (obj *patternFlowGtpv1Reserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1ReservedCounter value in the PatternFlowGtpv1Reserved object
func (obj *patternFlowGtpv1Reserved) SetDecrement(value PatternFlowGtpv1ReservedCounter) PatternFlowGtpv1Reserved {
	obj.setChoice(PatternFlowGtpv1ReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1ReservedMetricTag
func (obj *patternFlowGtpv1Reserved) MetricTags() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1ReservedMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter struct {
	obj                                    *patternFlowGtpv1Reserved
	patternFlowGtpv1ReservedMetricTagSlice []PatternFlowGtpv1ReservedMetricTag
	fieldPtr                               *[]*otg.PatternFlowGtpv1ReservedMetricTag
}

func newPatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter(ptr *[]*otg.PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	return &patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter interface {
	setMsg(*patternFlowGtpv1Reserved) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	Items() []PatternFlowGtpv1ReservedMetricTag
	Add() PatternFlowGtpv1ReservedMetricTag
	Append(items ...PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	Set(index int, newObj PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	Clear() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	clearHolderSlice() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter
}

func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) setMsg(msg *patternFlowGtpv1Reserved) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1ReservedMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) Items() []PatternFlowGtpv1ReservedMetricTag {
	return obj.patternFlowGtpv1ReservedMetricTagSlice
}

func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) Add() PatternFlowGtpv1ReservedMetricTag {
	newObj := &otg.PatternFlowGtpv1ReservedMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1ReservedMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1ReservedMetricTagSlice = append(obj.patternFlowGtpv1ReservedMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) Append(items ...PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1ReservedMetricTagSlice = append(obj.patternFlowGtpv1ReservedMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) Set(index int, newObj PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1ReservedMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) Clear() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1ReservedMetricTag{}
		obj.patternFlowGtpv1ReservedMetricTagSlice = []PatternFlowGtpv1ReservedMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) clearHolderSlice() PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	if len(obj.patternFlowGtpv1ReservedMetricTagSlice) > 0 {
		obj.patternFlowGtpv1ReservedMetricTagSlice = []PatternFlowGtpv1ReservedMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter) appendHolderSlice(item PatternFlowGtpv1ReservedMetricTag) PatternFlowGtpv1ReservedPatternFlowGtpv1ReservedMetricTagIter {
	obj.patternFlowGtpv1ReservedMetricTagSlice = append(obj.patternFlowGtpv1ReservedMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1Reserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1Reserved.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1Reserved.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1ReservedMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1Reserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1ReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1ReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1ReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1ReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1Reserved")
			}
		} else {
			intVal := otg.PatternFlowGtpv1Reserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1Reserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
