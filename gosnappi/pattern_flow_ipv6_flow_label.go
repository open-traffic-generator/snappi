package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6FlowLabel *****
type patternFlowIpv6FlowLabel struct {
	validation
	obj              *otg.PatternFlowIpv6FlowLabel
	marshaller       marshalPatternFlowIpv6FlowLabel
	unMarshaller     unMarshalPatternFlowIpv6FlowLabel
	incrementHolder  PatternFlowIpv6FlowLabelCounter
	decrementHolder  PatternFlowIpv6FlowLabelCounter
	metricTagsHolder PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
}

func NewPatternFlowIpv6FlowLabel() PatternFlowIpv6FlowLabel {
	obj := patternFlowIpv6FlowLabel{obj: &otg.PatternFlowIpv6FlowLabel{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6FlowLabel) msg() *otg.PatternFlowIpv6FlowLabel {
	return obj.obj
}

func (obj *patternFlowIpv6FlowLabel) setMsg(msg *otg.PatternFlowIpv6FlowLabel) PatternFlowIpv6FlowLabel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6FlowLabel struct {
	obj *patternFlowIpv6FlowLabel
}

type marshalPatternFlowIpv6FlowLabel interface {
	// ToProto marshals PatternFlowIpv6FlowLabel to protobuf object *otg.PatternFlowIpv6FlowLabel
	ToProto() (*otg.PatternFlowIpv6FlowLabel, error)
	// ToPbText marshals PatternFlowIpv6FlowLabel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6FlowLabel to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6FlowLabel to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6FlowLabel struct {
	obj *patternFlowIpv6FlowLabel
}

type unMarshalPatternFlowIpv6FlowLabel interface {
	// FromProto unmarshals PatternFlowIpv6FlowLabel from protobuf object *otg.PatternFlowIpv6FlowLabel
	FromProto(msg *otg.PatternFlowIpv6FlowLabel) (PatternFlowIpv6FlowLabel, error)
	// FromPbText unmarshals PatternFlowIpv6FlowLabel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6FlowLabel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6FlowLabel from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6FlowLabel) Marshal() marshalPatternFlowIpv6FlowLabel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6FlowLabel{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6FlowLabel) Unmarshal() unMarshalPatternFlowIpv6FlowLabel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6FlowLabel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6FlowLabel) ToProto() (*otg.PatternFlowIpv6FlowLabel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6FlowLabel) FromProto(msg *otg.PatternFlowIpv6FlowLabel) (PatternFlowIpv6FlowLabel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6FlowLabel) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabel) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabel) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabel) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6FlowLabel) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6FlowLabel) FromJson(value string) error {
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

func (obj *patternFlowIpv6FlowLabel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6FlowLabel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6FlowLabel) Clone() (PatternFlowIpv6FlowLabel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6FlowLabel()
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

func (obj *patternFlowIpv6FlowLabel) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6FlowLabel is flow label
type PatternFlowIpv6FlowLabel interface {
	Validation
	// msg marshals PatternFlowIpv6FlowLabel to protobuf object *otg.PatternFlowIpv6FlowLabel
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6FlowLabel
	// setMsg unmarshals PatternFlowIpv6FlowLabel from protobuf object *otg.PatternFlowIpv6FlowLabel
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6FlowLabel) PatternFlowIpv6FlowLabel
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6FlowLabel
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6FlowLabel
	// validate validates PatternFlowIpv6FlowLabel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6FlowLabel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6FlowLabelChoiceEnum, set in PatternFlowIpv6FlowLabel
	Choice() PatternFlowIpv6FlowLabelChoiceEnum
	// setChoice assigns PatternFlowIpv6FlowLabelChoiceEnum provided by user to PatternFlowIpv6FlowLabel
	setChoice(value PatternFlowIpv6FlowLabelChoiceEnum) PatternFlowIpv6FlowLabel
	// HasChoice checks if Choice has been set in PatternFlowIpv6FlowLabel
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6FlowLabel.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6FlowLabel
	SetValue(value uint32) PatternFlowIpv6FlowLabel
	// HasValue checks if Value has been set in PatternFlowIpv6FlowLabel
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6FlowLabel.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6FlowLabel
	SetValues(value []uint32) PatternFlowIpv6FlowLabel
	// Increment returns PatternFlowIpv6FlowLabelCounter, set in PatternFlowIpv6FlowLabel.
	// PatternFlowIpv6FlowLabelCounter is integer counter pattern
	Increment() PatternFlowIpv6FlowLabelCounter
	// SetIncrement assigns PatternFlowIpv6FlowLabelCounter provided by user to PatternFlowIpv6FlowLabel.
	// PatternFlowIpv6FlowLabelCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabel
	// HasIncrement checks if Increment has been set in PatternFlowIpv6FlowLabel
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6FlowLabelCounter, set in PatternFlowIpv6FlowLabel.
	// PatternFlowIpv6FlowLabelCounter is integer counter pattern
	Decrement() PatternFlowIpv6FlowLabelCounter
	// SetDecrement assigns PatternFlowIpv6FlowLabelCounter provided by user to PatternFlowIpv6FlowLabel.
	// PatternFlowIpv6FlowLabelCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabel
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6FlowLabel
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIterIter, set in PatternFlowIpv6FlowLabel
	MetricTags() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	setNil()
}

type PatternFlowIpv6FlowLabelChoiceEnum string

// Enum of Choice on PatternFlowIpv6FlowLabel
var PatternFlowIpv6FlowLabelChoice = struct {
	VALUE     PatternFlowIpv6FlowLabelChoiceEnum
	VALUES    PatternFlowIpv6FlowLabelChoiceEnum
	INCREMENT PatternFlowIpv6FlowLabelChoiceEnum
	DECREMENT PatternFlowIpv6FlowLabelChoiceEnum
}{
	VALUE:     PatternFlowIpv6FlowLabelChoiceEnum("value"),
	VALUES:    PatternFlowIpv6FlowLabelChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6FlowLabelChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6FlowLabelChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6FlowLabel) Choice() PatternFlowIpv6FlowLabelChoiceEnum {
	return PatternFlowIpv6FlowLabelChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6FlowLabel) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6FlowLabel) setChoice(value PatternFlowIpv6FlowLabelChoiceEnum) PatternFlowIpv6FlowLabel {
	intValue, ok := otg.PatternFlowIpv6FlowLabel_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6FlowLabelChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6FlowLabel_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6FlowLabelChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6FlowLabelChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6FlowLabelChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6FlowLabelCounter().msg()
	}

	if value == PatternFlowIpv6FlowLabelChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6FlowLabelCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6FlowLabel) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6FlowLabelChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6FlowLabel) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6FlowLabel object
func (obj *patternFlowIpv6FlowLabel) SetValue(value uint32) PatternFlowIpv6FlowLabel {
	obj.setChoice(PatternFlowIpv6FlowLabelChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6FlowLabel) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6FlowLabel object
func (obj *patternFlowIpv6FlowLabel) SetValues(value []uint32) PatternFlowIpv6FlowLabel {
	obj.setChoice(PatternFlowIpv6FlowLabelChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6FlowLabelCounter
func (obj *patternFlowIpv6FlowLabel) Increment() PatternFlowIpv6FlowLabelCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6FlowLabelChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6FlowLabelCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6FlowLabelCounter
func (obj *patternFlowIpv6FlowLabel) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6FlowLabelCounter value in the PatternFlowIpv6FlowLabel object
func (obj *patternFlowIpv6FlowLabel) SetIncrement(value PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabel {
	obj.setChoice(PatternFlowIpv6FlowLabelChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6FlowLabelCounter
func (obj *patternFlowIpv6FlowLabel) Decrement() PatternFlowIpv6FlowLabelCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6FlowLabelChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6FlowLabelCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6FlowLabelCounter
func (obj *patternFlowIpv6FlowLabel) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6FlowLabelCounter value in the PatternFlowIpv6FlowLabel object
func (obj *patternFlowIpv6FlowLabel) SetDecrement(value PatternFlowIpv6FlowLabelCounter) PatternFlowIpv6FlowLabel {
	obj.setChoice(PatternFlowIpv6FlowLabelChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6FlowLabelMetricTag
func (obj *patternFlowIpv6FlowLabel) MetricTags() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6FlowLabelMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter struct {
	obj                                    *patternFlowIpv6FlowLabel
	patternFlowIpv6FlowLabelMetricTagSlice []PatternFlowIpv6FlowLabelMetricTag
	fieldPtr                               *[]*otg.PatternFlowIpv6FlowLabelMetricTag
}

func newPatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter(ptr *[]*otg.PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	return &patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter interface {
	setMsg(*patternFlowIpv6FlowLabel) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	Items() []PatternFlowIpv6FlowLabelMetricTag
	Add() PatternFlowIpv6FlowLabelMetricTag
	Append(items ...PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	Set(index int, newObj PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	Clear() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	clearHolderSlice() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
	appendHolderSlice(item PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter
}

func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) setMsg(msg *patternFlowIpv6FlowLabel) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6FlowLabelMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) Items() []PatternFlowIpv6FlowLabelMetricTag {
	return obj.patternFlowIpv6FlowLabelMetricTagSlice
}

func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) Add() PatternFlowIpv6FlowLabelMetricTag {
	newObj := &otg.PatternFlowIpv6FlowLabelMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6FlowLabelMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6FlowLabelMetricTagSlice = append(obj.patternFlowIpv6FlowLabelMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) Append(items ...PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6FlowLabelMetricTagSlice = append(obj.patternFlowIpv6FlowLabelMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) Set(index int, newObj PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6FlowLabelMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) Clear() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6FlowLabelMetricTag{}
		obj.patternFlowIpv6FlowLabelMetricTagSlice = []PatternFlowIpv6FlowLabelMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) clearHolderSlice() PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	if len(obj.patternFlowIpv6FlowLabelMetricTagSlice) > 0 {
		obj.patternFlowIpv6FlowLabelMetricTagSlice = []PatternFlowIpv6FlowLabelMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter) appendHolderSlice(item PatternFlowIpv6FlowLabelMetricTag) PatternFlowIpv6FlowLabelPatternFlowIpv6FlowLabelMetricTagIter {
	obj.patternFlowIpv6FlowLabelMetricTagSlice = append(obj.patternFlowIpv6FlowLabelMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6FlowLabel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6FlowLabel.Value <= 1048575 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1048575 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6FlowLabel.Values <= 1048575 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6FlowLabelMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6FlowLabel) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6FlowLabelChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6FlowLabelChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6FlowLabelChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6FlowLabelChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6FlowLabelChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6FlowLabelChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6FlowLabel")
			}
		} else {
			intVal := otg.PatternFlowIpv6FlowLabel_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6FlowLabel_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
