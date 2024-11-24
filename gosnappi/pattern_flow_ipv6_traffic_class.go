package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6TrafficClass *****
type patternFlowIpv6TrafficClass struct {
	validation
	obj              *otg.PatternFlowIpv6TrafficClass
	marshaller       marshalPatternFlowIpv6TrafficClass
	unMarshaller     unMarshalPatternFlowIpv6TrafficClass
	incrementHolder  PatternFlowIpv6TrafficClassCounter
	decrementHolder  PatternFlowIpv6TrafficClassCounter
	metricTagsHolder PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
}

func NewPatternFlowIpv6TrafficClass() PatternFlowIpv6TrafficClass {
	obj := patternFlowIpv6TrafficClass{obj: &otg.PatternFlowIpv6TrafficClass{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6TrafficClass) msg() *otg.PatternFlowIpv6TrafficClass {
	return obj.obj
}

func (obj *patternFlowIpv6TrafficClass) setMsg(msg *otg.PatternFlowIpv6TrafficClass) PatternFlowIpv6TrafficClass {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6TrafficClass struct {
	obj *patternFlowIpv6TrafficClass
}

type marshalPatternFlowIpv6TrafficClass interface {
	// ToProto marshals PatternFlowIpv6TrafficClass to protobuf object *otg.PatternFlowIpv6TrafficClass
	ToProto() (*otg.PatternFlowIpv6TrafficClass, error)
	// ToPbText marshals PatternFlowIpv6TrafficClass to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6TrafficClass to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6TrafficClass to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv6TrafficClass to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv6TrafficClass struct {
	obj *patternFlowIpv6TrafficClass
}

type unMarshalPatternFlowIpv6TrafficClass interface {
	// FromProto unmarshals PatternFlowIpv6TrafficClass from protobuf object *otg.PatternFlowIpv6TrafficClass
	FromProto(msg *otg.PatternFlowIpv6TrafficClass) (PatternFlowIpv6TrafficClass, error)
	// FromPbText unmarshals PatternFlowIpv6TrafficClass from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6TrafficClass from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6TrafficClass from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6TrafficClass) Marshal() marshalPatternFlowIpv6TrafficClass {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6TrafficClass{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6TrafficClass) Unmarshal() unMarshalPatternFlowIpv6TrafficClass {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6TrafficClass{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6TrafficClass) ToProto() (*otg.PatternFlowIpv6TrafficClass, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6TrafficClass) FromProto(msg *otg.PatternFlowIpv6TrafficClass) (PatternFlowIpv6TrafficClass, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6TrafficClass) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClass) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClass) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClass) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6TrafficClass) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv6TrafficClass) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6TrafficClass) FromJson(value string) error {
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

func (obj *patternFlowIpv6TrafficClass) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClass) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6TrafficClass) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6TrafficClass) Clone() (PatternFlowIpv6TrafficClass, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6TrafficClass()
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

func (obj *patternFlowIpv6TrafficClass) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6TrafficClass is traffic class
type PatternFlowIpv6TrafficClass interface {
	Validation
	// msg marshals PatternFlowIpv6TrafficClass to protobuf object *otg.PatternFlowIpv6TrafficClass
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6TrafficClass
	// setMsg unmarshals PatternFlowIpv6TrafficClass from protobuf object *otg.PatternFlowIpv6TrafficClass
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6TrafficClass) PatternFlowIpv6TrafficClass
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6TrafficClass
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6TrafficClass
	// validate validates PatternFlowIpv6TrafficClass
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6TrafficClass, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6TrafficClassChoiceEnum, set in PatternFlowIpv6TrafficClass
	Choice() PatternFlowIpv6TrafficClassChoiceEnum
	// setChoice assigns PatternFlowIpv6TrafficClassChoiceEnum provided by user to PatternFlowIpv6TrafficClass
	setChoice(value PatternFlowIpv6TrafficClassChoiceEnum) PatternFlowIpv6TrafficClass
	// HasChoice checks if Choice has been set in PatternFlowIpv6TrafficClass
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6TrafficClass.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6TrafficClass
	SetValue(value uint32) PatternFlowIpv6TrafficClass
	// HasValue checks if Value has been set in PatternFlowIpv6TrafficClass
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6TrafficClass.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6TrafficClass
	SetValues(value []uint32) PatternFlowIpv6TrafficClass
	// Increment returns PatternFlowIpv6TrafficClassCounter, set in PatternFlowIpv6TrafficClass.
	// PatternFlowIpv6TrafficClassCounter is integer counter pattern
	Increment() PatternFlowIpv6TrafficClassCounter
	// SetIncrement assigns PatternFlowIpv6TrafficClassCounter provided by user to PatternFlowIpv6TrafficClass.
	// PatternFlowIpv6TrafficClassCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClass
	// HasIncrement checks if Increment has been set in PatternFlowIpv6TrafficClass
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6TrafficClassCounter, set in PatternFlowIpv6TrafficClass.
	// PatternFlowIpv6TrafficClassCounter is integer counter pattern
	Decrement() PatternFlowIpv6TrafficClassCounter
	// SetDecrement assigns PatternFlowIpv6TrafficClassCounter provided by user to PatternFlowIpv6TrafficClass.
	// PatternFlowIpv6TrafficClassCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClass
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6TrafficClass
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIterIter, set in PatternFlowIpv6TrafficClass
	MetricTags() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	setNil()
}

type PatternFlowIpv6TrafficClassChoiceEnum string

// Enum of Choice on PatternFlowIpv6TrafficClass
var PatternFlowIpv6TrafficClassChoice = struct {
	VALUE     PatternFlowIpv6TrafficClassChoiceEnum
	VALUES    PatternFlowIpv6TrafficClassChoiceEnum
	INCREMENT PatternFlowIpv6TrafficClassChoiceEnum
	DECREMENT PatternFlowIpv6TrafficClassChoiceEnum
}{
	VALUE:     PatternFlowIpv6TrafficClassChoiceEnum("value"),
	VALUES:    PatternFlowIpv6TrafficClassChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6TrafficClassChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6TrafficClassChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6TrafficClass) Choice() PatternFlowIpv6TrafficClassChoiceEnum {
	return PatternFlowIpv6TrafficClassChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6TrafficClass) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6TrafficClass) setChoice(value PatternFlowIpv6TrafficClassChoiceEnum) PatternFlowIpv6TrafficClass {
	intValue, ok := otg.PatternFlowIpv6TrafficClass_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6TrafficClassChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6TrafficClass_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6TrafficClassChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6TrafficClassChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6TrafficClassChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6TrafficClassCounter().msg()
	}

	if value == PatternFlowIpv6TrafficClassChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6TrafficClassCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6TrafficClass) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6TrafficClassChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6TrafficClass) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6TrafficClass object
func (obj *patternFlowIpv6TrafficClass) SetValue(value uint32) PatternFlowIpv6TrafficClass {
	obj.setChoice(PatternFlowIpv6TrafficClassChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6TrafficClass) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6TrafficClass object
func (obj *patternFlowIpv6TrafficClass) SetValues(value []uint32) PatternFlowIpv6TrafficClass {
	obj.setChoice(PatternFlowIpv6TrafficClassChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6TrafficClassCounter
func (obj *patternFlowIpv6TrafficClass) Increment() PatternFlowIpv6TrafficClassCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6TrafficClassChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6TrafficClassCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6TrafficClassCounter
func (obj *patternFlowIpv6TrafficClass) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6TrafficClassCounter value in the PatternFlowIpv6TrafficClass object
func (obj *patternFlowIpv6TrafficClass) SetIncrement(value PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClass {
	obj.setChoice(PatternFlowIpv6TrafficClassChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6TrafficClassCounter
func (obj *patternFlowIpv6TrafficClass) Decrement() PatternFlowIpv6TrafficClassCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6TrafficClassChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6TrafficClassCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6TrafficClassCounter
func (obj *patternFlowIpv6TrafficClass) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6TrafficClassCounter value in the PatternFlowIpv6TrafficClass object
func (obj *patternFlowIpv6TrafficClass) SetDecrement(value PatternFlowIpv6TrafficClassCounter) PatternFlowIpv6TrafficClass {
	obj.setChoice(PatternFlowIpv6TrafficClassChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6TrafficClassMetricTag
func (obj *patternFlowIpv6TrafficClass) MetricTags() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6TrafficClassMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter struct {
	obj                                       *patternFlowIpv6TrafficClass
	patternFlowIpv6TrafficClassMetricTagSlice []PatternFlowIpv6TrafficClassMetricTag
	fieldPtr                                  *[]*otg.PatternFlowIpv6TrafficClassMetricTag
}

func newPatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter(ptr *[]*otg.PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	return &patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter interface {
	setMsg(*patternFlowIpv6TrafficClass) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	Items() []PatternFlowIpv6TrafficClassMetricTag
	Add() PatternFlowIpv6TrafficClassMetricTag
	Append(items ...PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	Set(index int, newObj PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	Clear() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	clearHolderSlice() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
	appendHolderSlice(item PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter
}

func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) setMsg(msg *patternFlowIpv6TrafficClass) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6TrafficClassMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) Items() []PatternFlowIpv6TrafficClassMetricTag {
	return obj.patternFlowIpv6TrafficClassMetricTagSlice
}

func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) Add() PatternFlowIpv6TrafficClassMetricTag {
	newObj := &otg.PatternFlowIpv6TrafficClassMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6TrafficClassMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6TrafficClassMetricTagSlice = append(obj.patternFlowIpv6TrafficClassMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) Append(items ...PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6TrafficClassMetricTagSlice = append(obj.patternFlowIpv6TrafficClassMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) Set(index int, newObj PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6TrafficClassMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) Clear() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6TrafficClassMetricTag{}
		obj.patternFlowIpv6TrafficClassMetricTagSlice = []PatternFlowIpv6TrafficClassMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) clearHolderSlice() PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	if len(obj.patternFlowIpv6TrafficClassMetricTagSlice) > 0 {
		obj.patternFlowIpv6TrafficClassMetricTagSlice = []PatternFlowIpv6TrafficClassMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter) appendHolderSlice(item PatternFlowIpv6TrafficClassMetricTag) PatternFlowIpv6TrafficClassPatternFlowIpv6TrafficClassMetricTagIter {
	obj.patternFlowIpv6TrafficClassMetricTagSlice = append(obj.patternFlowIpv6TrafficClassMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6TrafficClass) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6TrafficClass.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6TrafficClass.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6TrafficClassMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6TrafficClass) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv6TrafficClassChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv6TrafficClassChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv6TrafficClassChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv6TrafficClassChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv6TrafficClassChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv6TrafficClassChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv6TrafficClass")
			}
		} else {
			intVal := otg.PatternFlowIpv6TrafficClass_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv6TrafficClass_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
