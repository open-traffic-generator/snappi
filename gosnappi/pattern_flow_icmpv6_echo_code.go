package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoCode *****
type patternFlowIcmpv6EchoCode struct {
	validation
	obj              *otg.PatternFlowIcmpv6EchoCode
	marshaller       marshalPatternFlowIcmpv6EchoCode
	unMarshaller     unMarshalPatternFlowIcmpv6EchoCode
	incrementHolder  PatternFlowIcmpv6EchoCodeCounter
	decrementHolder  PatternFlowIcmpv6EchoCodeCounter
	metricTagsHolder PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
}

func NewPatternFlowIcmpv6EchoCode() PatternFlowIcmpv6EchoCode {
	obj := patternFlowIcmpv6EchoCode{obj: &otg.PatternFlowIcmpv6EchoCode{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoCode) msg() *otg.PatternFlowIcmpv6EchoCode {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoCode) setMsg(msg *otg.PatternFlowIcmpv6EchoCode) PatternFlowIcmpv6EchoCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoCode struct {
	obj *patternFlowIcmpv6EchoCode
}

type marshalPatternFlowIcmpv6EchoCode interface {
	// ToProto marshals PatternFlowIcmpv6EchoCode to protobuf object *otg.PatternFlowIcmpv6EchoCode
	ToProto() (*otg.PatternFlowIcmpv6EchoCode, error)
	// ToPbText marshals PatternFlowIcmpv6EchoCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoCode to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoCode struct {
	obj *patternFlowIcmpv6EchoCode
}

type unMarshalPatternFlowIcmpv6EchoCode interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoCode from protobuf object *otg.PatternFlowIcmpv6EchoCode
	FromProto(msg *otg.PatternFlowIcmpv6EchoCode) (PatternFlowIcmpv6EchoCode, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoCode from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoCode) Marshal() marshalPatternFlowIcmpv6EchoCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoCode) Unmarshal() unMarshalPatternFlowIcmpv6EchoCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoCode) ToProto() (*otg.PatternFlowIcmpv6EchoCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoCode) FromProto(msg *otg.PatternFlowIcmpv6EchoCode) (PatternFlowIcmpv6EchoCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoCode) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCode) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCode) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCode) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoCode) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoCode) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoCode) Clone() (PatternFlowIcmpv6EchoCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoCode()
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

func (obj *patternFlowIcmpv6EchoCode) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6EchoCode is iCMPv6 echo sub type
type PatternFlowIcmpv6EchoCode interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoCode to protobuf object *otg.PatternFlowIcmpv6EchoCode
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoCode
	// setMsg unmarshals PatternFlowIcmpv6EchoCode from protobuf object *otg.PatternFlowIcmpv6EchoCode
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoCode) PatternFlowIcmpv6EchoCode
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoCode
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoCode
	// validate validates PatternFlowIcmpv6EchoCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6EchoCodeChoiceEnum, set in PatternFlowIcmpv6EchoCode
	Choice() PatternFlowIcmpv6EchoCodeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6EchoCodeChoiceEnum provided by user to PatternFlowIcmpv6EchoCode
	setChoice(value PatternFlowIcmpv6EchoCodeChoiceEnum) PatternFlowIcmpv6EchoCode
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6EchoCode
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6EchoCode.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6EchoCode
	SetValue(value uint32) PatternFlowIcmpv6EchoCode
	// HasValue checks if Value has been set in PatternFlowIcmpv6EchoCode
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6EchoCode.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6EchoCode
	SetValues(value []uint32) PatternFlowIcmpv6EchoCode
	// Increment returns PatternFlowIcmpv6EchoCodeCounter, set in PatternFlowIcmpv6EchoCode.
	// PatternFlowIcmpv6EchoCodeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6EchoCodeCounter
	// SetIncrement assigns PatternFlowIcmpv6EchoCodeCounter provided by user to PatternFlowIcmpv6EchoCode.
	// PatternFlowIcmpv6EchoCodeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCode
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6EchoCode
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6EchoCodeCounter, set in PatternFlowIcmpv6EchoCode.
	// PatternFlowIcmpv6EchoCodeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6EchoCodeCounter
	// SetDecrement assigns PatternFlowIcmpv6EchoCodeCounter provided by user to PatternFlowIcmpv6EchoCode.
	// PatternFlowIcmpv6EchoCodeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCode
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6EchoCode
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIterIter, set in PatternFlowIcmpv6EchoCode
	MetricTags() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	setNil()
}

type PatternFlowIcmpv6EchoCodeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6EchoCode
var PatternFlowIcmpv6EchoCodeChoice = struct {
	VALUE     PatternFlowIcmpv6EchoCodeChoiceEnum
	VALUES    PatternFlowIcmpv6EchoCodeChoiceEnum
	INCREMENT PatternFlowIcmpv6EchoCodeChoiceEnum
	DECREMENT PatternFlowIcmpv6EchoCodeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6EchoCodeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6EchoCodeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6EchoCodeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6EchoCodeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6EchoCode) Choice() PatternFlowIcmpv6EchoCodeChoiceEnum {
	return PatternFlowIcmpv6EchoCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6EchoCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6EchoCode) setChoice(value PatternFlowIcmpv6EchoCodeChoiceEnum) PatternFlowIcmpv6EchoCode {
	intValue, ok := otg.PatternFlowIcmpv6EchoCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6EchoCodeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6EchoCodeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6EchoCodeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6EchoCodeCounter().msg()
	}

	if value == PatternFlowIcmpv6EchoCodeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6EchoCodeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoCode) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoCode) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6EchoCode object
func (obj *patternFlowIcmpv6EchoCode) SetValue(value uint32) PatternFlowIcmpv6EchoCode {
	obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6EchoCode) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6EchoCode object
func (obj *patternFlowIcmpv6EchoCode) SetValues(value []uint32) PatternFlowIcmpv6EchoCode {
	obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoCodeCounter
func (obj *patternFlowIcmpv6EchoCode) Increment() PatternFlowIcmpv6EchoCodeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6EchoCodeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoCodeCounter
func (obj *patternFlowIcmpv6EchoCode) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6EchoCodeCounter value in the PatternFlowIcmpv6EchoCode object
func (obj *patternFlowIcmpv6EchoCode) SetIncrement(value PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCode {
	obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoCodeCounter
func (obj *patternFlowIcmpv6EchoCode) Decrement() PatternFlowIcmpv6EchoCodeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6EchoCodeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoCodeCounter
func (obj *patternFlowIcmpv6EchoCode) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6EchoCodeCounter value in the PatternFlowIcmpv6EchoCode object
func (obj *patternFlowIcmpv6EchoCode) SetDecrement(value PatternFlowIcmpv6EchoCodeCounter) PatternFlowIcmpv6EchoCode {
	obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpv6EchoCodeMetricTag
func (obj *patternFlowIcmpv6EchoCode) MetricTags() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpv6EchoCodeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter struct {
	obj                                     *patternFlowIcmpv6EchoCode
	patternFlowIcmpv6EchoCodeMetricTagSlice []PatternFlowIcmpv6EchoCodeMetricTag
	fieldPtr                                *[]*otg.PatternFlowIcmpv6EchoCodeMetricTag
}

func newPatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter(ptr *[]*otg.PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	return &patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter interface {
	setMsg(*patternFlowIcmpv6EchoCode) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	Items() []PatternFlowIcmpv6EchoCodeMetricTag
	Add() PatternFlowIcmpv6EchoCodeMetricTag
	Append(items ...PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	Set(index int, newObj PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	Clear() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	clearHolderSlice() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
	appendHolderSlice(item PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter
}

func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) setMsg(msg *patternFlowIcmpv6EchoCode) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpv6EchoCodeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) Items() []PatternFlowIcmpv6EchoCodeMetricTag {
	return obj.patternFlowIcmpv6EchoCodeMetricTagSlice
}

func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) Add() PatternFlowIcmpv6EchoCodeMetricTag {
	newObj := &otg.PatternFlowIcmpv6EchoCodeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpv6EchoCodeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpv6EchoCodeMetricTagSlice = append(obj.patternFlowIcmpv6EchoCodeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) Append(items ...PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpv6EchoCodeMetricTagSlice = append(obj.patternFlowIcmpv6EchoCodeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) Set(index int, newObj PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpv6EchoCodeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) Clear() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpv6EchoCodeMetricTag{}
		obj.patternFlowIcmpv6EchoCodeMetricTagSlice = []PatternFlowIcmpv6EchoCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) clearHolderSlice() PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	if len(obj.patternFlowIcmpv6EchoCodeMetricTagSlice) > 0 {
		obj.patternFlowIcmpv6EchoCodeMetricTagSlice = []PatternFlowIcmpv6EchoCodeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter) appendHolderSlice(item PatternFlowIcmpv6EchoCodeMetricTag) PatternFlowIcmpv6EchoCodePatternFlowIcmpv6EchoCodeMetricTagIter {
	obj.patternFlowIcmpv6EchoCodeMetricTagSlice = append(obj.patternFlowIcmpv6EchoCodeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpv6EchoCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoCode.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6EchoCode.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpv6EchoCodeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpv6EchoCode) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6EchoCodeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoCodeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoCodeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoCodeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoCodeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6EchoCodeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6EchoCode")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6EchoCode_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6EchoCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
