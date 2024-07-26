package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoCode *****
type patternFlowIcmpEchoCode struct {
	validation
	obj              *otg.PatternFlowIcmpEchoCode
	marshaller       marshalPatternFlowIcmpEchoCode
	unMarshaller     unMarshalPatternFlowIcmpEchoCode
	incrementHolder  PatternFlowIcmpEchoCodeCounter
	decrementHolder  PatternFlowIcmpEchoCodeCounter
	metricTagsHolder PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
}

func NewPatternFlowIcmpEchoCode() PatternFlowIcmpEchoCode {
	obj := patternFlowIcmpEchoCode{obj: &otg.PatternFlowIcmpEchoCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoCode) msg() *otg.PatternFlowIcmpEchoCode {
	return obj.obj
}

func (obj *patternFlowIcmpEchoCode) setMsg(msg *otg.PatternFlowIcmpEchoCode) PatternFlowIcmpEchoCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoCode struct {
	obj *patternFlowIcmpEchoCode
}

type marshalPatternFlowIcmpEchoCode interface {
	// ToProto marshals PatternFlowIcmpEchoCode to protobuf object *otg.PatternFlowIcmpEchoCode
	ToProto() (*otg.PatternFlowIcmpEchoCode, error)
	// ToPbText marshals PatternFlowIcmpEchoCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoCode to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoCode struct {
	obj *patternFlowIcmpEchoCode
}

type unMarshalPatternFlowIcmpEchoCode interface {
	// FromProto unmarshals PatternFlowIcmpEchoCode from protobuf object *otg.PatternFlowIcmpEchoCode
	FromProto(msg *otg.PatternFlowIcmpEchoCode) (PatternFlowIcmpEchoCode, error)
	// FromPbText unmarshals PatternFlowIcmpEchoCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoCode) Marshal() marshalPatternFlowIcmpEchoCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoCode) Unmarshal() unMarshalPatternFlowIcmpEchoCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoCode) ToProto() (*otg.PatternFlowIcmpEchoCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoCode) FromProto(msg *otg.PatternFlowIcmpEchoCode) (PatternFlowIcmpEchoCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoCode) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoCode) Clone() (PatternFlowIcmpEchoCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoCode()
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

func (obj *patternFlowIcmpEchoCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpEchoCode is the ICMP subtype.  The default code for ICMP echo request and reply is 0.
type PatternFlowIcmpEchoCode interface {
	Validation
	// msg marshals PatternFlowIcmpEchoCode to protobuf object *otg.PatternFlowIcmpEchoCode
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoCode
	// setMsg unmarshals PatternFlowIcmpEchoCode from protobuf object *otg.PatternFlowIcmpEchoCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoCode) PatternFlowIcmpEchoCode
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoCode
	// validate validates PatternFlowIcmpEchoCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpEchoCodeChoiceEnum, set in PatternFlowIcmpEchoCode
	Choice() PatternFlowIcmpEchoCodeChoiceEnum
	// setChoice assigns PatternFlowIcmpEchoCodeChoiceEnum provided by user to PatternFlowIcmpEchoCode
	setChoice(value PatternFlowIcmpEchoCodeChoiceEnum) PatternFlowIcmpEchoCode
	// HasChoice checks if Choice has been set in PatternFlowIcmpEchoCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpEchoCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpEchoCode
	SetValue(value uint32) PatternFlowIcmpEchoCode
	// HasValue checks if Value has been set in PatternFlowIcmpEchoCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpEchoCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpEchoCode
	SetValues(value []uint32) PatternFlowIcmpEchoCode
	// Increment returns PatternFlowIcmpEchoCodeCounter, set in PatternFlowIcmpEchoCode.
	// PatternFlowIcmpEchoCodeCounter is integer counter pattern
	Increment() PatternFlowIcmpEchoCodeCounter
	// SetIncrement assigns PatternFlowIcmpEchoCodeCounter provided by user to PatternFlowIcmpEchoCode.
	// PatternFlowIcmpEchoCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCode
	// HasIncrement checks if Increment has been set in PatternFlowIcmpEchoCode
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpEchoCodeCounter, set in PatternFlowIcmpEchoCode.
	// PatternFlowIcmpEchoCodeCounter is integer counter pattern
	Decrement() PatternFlowIcmpEchoCodeCounter
	// SetDecrement assigns PatternFlowIcmpEchoCodeCounter provided by user to PatternFlowIcmpEchoCode.
	// PatternFlowIcmpEchoCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCode
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpEchoCode
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIterIter, set in PatternFlowIcmpEchoCode
	MetricTags() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	setNil()
}

type PatternFlowIcmpEchoCodeChoiceEnum string

// Enum of Choice on PatternFlowIcmpEchoCode
var PatternFlowIcmpEchoCodeChoice = struct {
	VALUE     PatternFlowIcmpEchoCodeChoiceEnum
	VALUES    PatternFlowIcmpEchoCodeChoiceEnum
	INCREMENT PatternFlowIcmpEchoCodeChoiceEnum
	DECREMENT PatternFlowIcmpEchoCodeChoiceEnum
}{
	VALUE:     PatternFlowIcmpEchoCodeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpEchoCodeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpEchoCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpEchoCodeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpEchoCode) Choice() PatternFlowIcmpEchoCodeChoiceEnum {
	return PatternFlowIcmpEchoCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpEchoCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpEchoCode) setChoice(value PatternFlowIcmpEchoCodeChoiceEnum) PatternFlowIcmpEchoCode {
	intValue, ok := otg.PatternFlowIcmpEchoCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpEchoCodeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpEchoCodeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpEchoCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpEchoCodeCounter().msg()
	}

	if value == PatternFlowIcmpEchoCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpEchoCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpEchoCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpEchoCode object
func (obj *patternFlowIcmpEchoCode) SetValue(value uint32) PatternFlowIcmpEchoCode {
	obj.setChoice(PatternFlowIcmpEchoCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpEchoCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpEchoCode object
func (obj *patternFlowIcmpEchoCode) SetValues(value []uint32) PatternFlowIcmpEchoCode {
	obj.setChoice(PatternFlowIcmpEchoCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoCodeCounter
func (obj *patternFlowIcmpEchoCode) Increment() PatternFlowIcmpEchoCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpEchoCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpEchoCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoCodeCounter
func (obj *patternFlowIcmpEchoCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpEchoCodeCounter value in the PatternFlowIcmpEchoCode object
func (obj *patternFlowIcmpEchoCode) SetIncrement(value PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCode {
	obj.setChoice(PatternFlowIcmpEchoCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoCodeCounter
func (obj *patternFlowIcmpEchoCode) Decrement() PatternFlowIcmpEchoCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpEchoCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpEchoCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoCodeCounter
func (obj *patternFlowIcmpEchoCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpEchoCodeCounter value in the PatternFlowIcmpEchoCode object
func (obj *patternFlowIcmpEchoCode) SetDecrement(value PatternFlowIcmpEchoCodeCounter) PatternFlowIcmpEchoCode {
	obj.setChoice(PatternFlowIcmpEchoCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpEchoCodeMetricTag
func (obj *patternFlowIcmpEchoCode) MetricTags() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpEchoCodeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter struct {
	obj                                   *patternFlowIcmpEchoCode
	patternFlowIcmpEchoCodeMetricTagSlice []PatternFlowIcmpEchoCodeMetricTag
	fieldPtr                              *[]*otg.PatternFlowIcmpEchoCodeMetricTag
}

func newPatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter(ptr *[]*otg.PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	return &patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter interface {
	setMsg(*patternFlowIcmpEchoCode) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	Items() []PatternFlowIcmpEchoCodeMetricTag
	Add() PatternFlowIcmpEchoCodeMetricTag
	Append(items ...PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	Set(index int, newObj PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	Clear() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	clearHolderSlice() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
	appendHolderSlice(item PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter
}

func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) setMsg(msg *patternFlowIcmpEchoCode) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpEchoCodeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) Items() []PatternFlowIcmpEchoCodeMetricTag {
	return obj.patternFlowIcmpEchoCodeMetricTagSlice
}

func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) Add() PatternFlowIcmpEchoCodeMetricTag {
	newObj := &otg.PatternFlowIcmpEchoCodeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpEchoCodeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpEchoCodeMetricTagSlice = append(obj.patternFlowIcmpEchoCodeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) Append(items ...PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpEchoCodeMetricTagSlice = append(obj.patternFlowIcmpEchoCodeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) Set(index int, newObj PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpEchoCodeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) Clear() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpEchoCodeMetricTag{}
		obj.patternFlowIcmpEchoCodeMetricTagSlice = []PatternFlowIcmpEchoCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) clearHolderSlice() PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	if len(obj.patternFlowIcmpEchoCodeMetricTagSlice) > 0 {
		obj.patternFlowIcmpEchoCodeMetricTagSlice = []PatternFlowIcmpEchoCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter) appendHolderSlice(item PatternFlowIcmpEchoCodeMetricTag) PatternFlowIcmpEchoCodePatternFlowIcmpEchoCodeMetricTagIter {
	obj.patternFlowIcmpEchoCodeMetricTagSlice = append(obj.patternFlowIcmpEchoCodeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpEchoCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpEchoCode.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpEchoCodeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpEchoCode) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpEchoCodeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoCodeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpEchoCodeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoCodeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoCodeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpEchoCodeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpEchoCode")
			}
		} else {
			intVal := otg.PatternFlowIcmpEchoCode_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpEchoCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
