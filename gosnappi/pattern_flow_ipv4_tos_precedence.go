package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosPrecedence *****
type patternFlowIpv4TosPrecedence struct {
	validation
	obj              *otg.PatternFlowIpv4TosPrecedence
	marshaller       marshalPatternFlowIpv4TosPrecedence
	unMarshaller     unMarshalPatternFlowIpv4TosPrecedence
	incrementHolder  PatternFlowIpv4TosPrecedenceCounter
	decrementHolder  PatternFlowIpv4TosPrecedenceCounter
	metricTagsHolder PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
}

func NewPatternFlowIpv4TosPrecedence() PatternFlowIpv4TosPrecedence {
	obj := patternFlowIpv4TosPrecedence{obj: &otg.PatternFlowIpv4TosPrecedence{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosPrecedence) msg() *otg.PatternFlowIpv4TosPrecedence {
	return obj.obj
}

func (obj *patternFlowIpv4TosPrecedence) setMsg(msg *otg.PatternFlowIpv4TosPrecedence) PatternFlowIpv4TosPrecedence {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosPrecedence struct {
	obj *patternFlowIpv4TosPrecedence
}

type marshalPatternFlowIpv4TosPrecedence interface {
	// ToProto marshals PatternFlowIpv4TosPrecedence to protobuf object *otg.PatternFlowIpv4TosPrecedence
	ToProto() (*otg.PatternFlowIpv4TosPrecedence, error)
	// ToPbText marshals PatternFlowIpv4TosPrecedence to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosPrecedence to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosPrecedence to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosPrecedence struct {
	obj *patternFlowIpv4TosPrecedence
}

type unMarshalPatternFlowIpv4TosPrecedence interface {
	// FromProto unmarshals PatternFlowIpv4TosPrecedence from protobuf object *otg.PatternFlowIpv4TosPrecedence
	FromProto(msg *otg.PatternFlowIpv4TosPrecedence) (PatternFlowIpv4TosPrecedence, error)
	// FromPbText unmarshals PatternFlowIpv4TosPrecedence from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosPrecedence from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosPrecedence from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosPrecedence) Marshal() marshalPatternFlowIpv4TosPrecedence {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosPrecedence{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosPrecedence) Unmarshal() unMarshalPatternFlowIpv4TosPrecedence {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosPrecedence{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosPrecedence) ToProto() (*otg.PatternFlowIpv4TosPrecedence, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosPrecedence) FromProto(msg *otg.PatternFlowIpv4TosPrecedence) (PatternFlowIpv4TosPrecedence, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosPrecedence) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedence) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedence) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedence) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosPrecedence) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosPrecedence) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosPrecedence) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedence) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosPrecedence) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosPrecedence) Clone() (PatternFlowIpv4TosPrecedence, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosPrecedence()
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

func (obj *patternFlowIpv4TosPrecedence) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosPrecedence is precedence
type PatternFlowIpv4TosPrecedence interface {
	Validation
	// msg marshals PatternFlowIpv4TosPrecedence to protobuf object *otg.PatternFlowIpv4TosPrecedence
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosPrecedence
	// setMsg unmarshals PatternFlowIpv4TosPrecedence from protobuf object *otg.PatternFlowIpv4TosPrecedence
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosPrecedence) PatternFlowIpv4TosPrecedence
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosPrecedence
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosPrecedence
	// validate validates PatternFlowIpv4TosPrecedence
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosPrecedence, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosPrecedenceChoiceEnum, set in PatternFlowIpv4TosPrecedence
	Choice() PatternFlowIpv4TosPrecedenceChoiceEnum
	// setChoice assigns PatternFlowIpv4TosPrecedenceChoiceEnum provided by user to PatternFlowIpv4TosPrecedence
	setChoice(value PatternFlowIpv4TosPrecedenceChoiceEnum) PatternFlowIpv4TosPrecedence
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosPrecedence
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosPrecedence.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosPrecedence
	SetValue(value uint32) PatternFlowIpv4TosPrecedence
	// HasValue checks if Value has been set in PatternFlowIpv4TosPrecedence
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosPrecedence.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosPrecedence
	SetValues(value []uint32) PatternFlowIpv4TosPrecedence
	// Increment returns PatternFlowIpv4TosPrecedenceCounter, set in PatternFlowIpv4TosPrecedence.
	// PatternFlowIpv4TosPrecedenceCounter is integer counter pattern
	Increment() PatternFlowIpv4TosPrecedenceCounter
	// SetIncrement assigns PatternFlowIpv4TosPrecedenceCounter provided by user to PatternFlowIpv4TosPrecedence.
	// PatternFlowIpv4TosPrecedenceCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedence
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosPrecedence
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosPrecedenceCounter, set in PatternFlowIpv4TosPrecedence.
	// PatternFlowIpv4TosPrecedenceCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosPrecedenceCounter
	// SetDecrement assigns PatternFlowIpv4TosPrecedenceCounter provided by user to PatternFlowIpv4TosPrecedence.
	// PatternFlowIpv4TosPrecedenceCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedence
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosPrecedence
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIterIter, set in PatternFlowIpv4TosPrecedence
	MetricTags() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	setNil()
}

type PatternFlowIpv4TosPrecedenceChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosPrecedence
var PatternFlowIpv4TosPrecedenceChoice = struct {
	VALUE     PatternFlowIpv4TosPrecedenceChoiceEnum
	VALUES    PatternFlowIpv4TosPrecedenceChoiceEnum
	INCREMENT PatternFlowIpv4TosPrecedenceChoiceEnum
	DECREMENT PatternFlowIpv4TosPrecedenceChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosPrecedenceChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosPrecedenceChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosPrecedenceChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosPrecedenceChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosPrecedence) Choice() PatternFlowIpv4TosPrecedenceChoiceEnum {
	return PatternFlowIpv4TosPrecedenceChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosPrecedence) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosPrecedence) setChoice(value PatternFlowIpv4TosPrecedenceChoiceEnum) PatternFlowIpv4TosPrecedence {
	intValue, ok := otg.PatternFlowIpv4TosPrecedence_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosPrecedenceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosPrecedence_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosPrecedenceChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosPrecedenceChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosPrecedenceChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosPrecedenceCounter().msg()
	}

	if value == PatternFlowIpv4TosPrecedenceChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosPrecedenceCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosPrecedence) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosPrecedence) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosPrecedence object
func (obj *patternFlowIpv4TosPrecedence) SetValue(value uint32) PatternFlowIpv4TosPrecedence {
	obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosPrecedence) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosPrecedence object
func (obj *patternFlowIpv4TosPrecedence) SetValues(value []uint32) PatternFlowIpv4TosPrecedence {
	obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosPrecedenceCounter
func (obj *patternFlowIpv4TosPrecedence) Increment() PatternFlowIpv4TosPrecedenceCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosPrecedenceCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosPrecedenceCounter
func (obj *patternFlowIpv4TosPrecedence) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosPrecedenceCounter value in the PatternFlowIpv4TosPrecedence object
func (obj *patternFlowIpv4TosPrecedence) SetIncrement(value PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedence {
	obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosPrecedenceCounter
func (obj *patternFlowIpv4TosPrecedence) Decrement() PatternFlowIpv4TosPrecedenceCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosPrecedenceCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosPrecedenceCounter
func (obj *patternFlowIpv4TosPrecedence) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosPrecedenceCounter value in the PatternFlowIpv4TosPrecedence object
func (obj *patternFlowIpv4TosPrecedence) SetDecrement(value PatternFlowIpv4TosPrecedenceCounter) PatternFlowIpv4TosPrecedence {
	obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosPrecedenceMetricTag
func (obj *patternFlowIpv4TosPrecedence) MetricTags() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosPrecedenceMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter struct {
	obj                                        *patternFlowIpv4TosPrecedence
	patternFlowIpv4TosPrecedenceMetricTagSlice []PatternFlowIpv4TosPrecedenceMetricTag
	fieldPtr                                   *[]*otg.PatternFlowIpv4TosPrecedenceMetricTag
}

func newPatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	return &patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter interface {
	setMsg(*patternFlowIpv4TosPrecedence) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	Items() []PatternFlowIpv4TosPrecedenceMetricTag
	Add() PatternFlowIpv4TosPrecedenceMetricTag
	Append(items ...PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	Clear() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter
}

func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) setMsg(msg *patternFlowIpv4TosPrecedence) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosPrecedenceMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) Items() []PatternFlowIpv4TosPrecedenceMetricTag {
	return obj.patternFlowIpv4TosPrecedenceMetricTagSlice
}

func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) Add() PatternFlowIpv4TosPrecedenceMetricTag {
	newObj := &otg.PatternFlowIpv4TosPrecedenceMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosPrecedenceMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosPrecedenceMetricTagSlice = append(obj.patternFlowIpv4TosPrecedenceMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) Append(items ...PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosPrecedenceMetricTagSlice = append(obj.patternFlowIpv4TosPrecedenceMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) Set(index int, newObj PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosPrecedenceMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) Clear() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosPrecedenceMetricTag{}
		obj.patternFlowIpv4TosPrecedenceMetricTagSlice = []PatternFlowIpv4TosPrecedenceMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) clearHolderSlice() PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	if len(obj.patternFlowIpv4TosPrecedenceMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosPrecedenceMetricTagSlice = []PatternFlowIpv4TosPrecedenceMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosPrecedenceMetricTag) PatternFlowIpv4TosPrecedencePatternFlowIpv4TosPrecedenceMetricTagIter {
	obj.patternFlowIpv4TosPrecedenceMetricTagSlice = append(obj.patternFlowIpv4TosPrecedenceMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosPrecedence) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosPrecedence.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosPrecedence.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosPrecedenceMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosPrecedence) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4TosPrecedenceChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosPrecedenceChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4TosPrecedenceChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosPrecedenceChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosPrecedenceChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4TosPrecedenceChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4TosPrecedence")
			}
		} else {
			intVal := otg.PatternFlowIpv4TosPrecedence_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4TosPrecedence_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
