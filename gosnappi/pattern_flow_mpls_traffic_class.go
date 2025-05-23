package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsTrafficClass *****
type patternFlowMplsTrafficClass struct {
	validation
	obj              *otg.PatternFlowMplsTrafficClass
	marshaller       marshalPatternFlowMplsTrafficClass
	unMarshaller     unMarshalPatternFlowMplsTrafficClass
	incrementHolder  PatternFlowMplsTrafficClassCounter
	decrementHolder  PatternFlowMplsTrafficClassCounter
	metricTagsHolder PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
}

func NewPatternFlowMplsTrafficClass() PatternFlowMplsTrafficClass {
	obj := patternFlowMplsTrafficClass{obj: &otg.PatternFlowMplsTrafficClass{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsTrafficClass) msg() *otg.PatternFlowMplsTrafficClass {
	return obj.obj
}

func (obj *patternFlowMplsTrafficClass) setMsg(msg *otg.PatternFlowMplsTrafficClass) PatternFlowMplsTrafficClass {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsTrafficClass struct {
	obj *patternFlowMplsTrafficClass
}

type marshalPatternFlowMplsTrafficClass interface {
	// ToProto marshals PatternFlowMplsTrafficClass to protobuf object *otg.PatternFlowMplsTrafficClass
	ToProto() (*otg.PatternFlowMplsTrafficClass, error)
	// ToPbText marshals PatternFlowMplsTrafficClass to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsTrafficClass to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsTrafficClass to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowMplsTrafficClass to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowMplsTrafficClass struct {
	obj *patternFlowMplsTrafficClass
}

type unMarshalPatternFlowMplsTrafficClass interface {
	// FromProto unmarshals PatternFlowMplsTrafficClass from protobuf object *otg.PatternFlowMplsTrafficClass
	FromProto(msg *otg.PatternFlowMplsTrafficClass) (PatternFlowMplsTrafficClass, error)
	// FromPbText unmarshals PatternFlowMplsTrafficClass from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsTrafficClass from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsTrafficClass from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsTrafficClass) Marshal() marshalPatternFlowMplsTrafficClass {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsTrafficClass{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsTrafficClass) Unmarshal() unMarshalPatternFlowMplsTrafficClass {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsTrafficClass{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsTrafficClass) ToProto() (*otg.PatternFlowMplsTrafficClass, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsTrafficClass) FromProto(msg *otg.PatternFlowMplsTrafficClass) (PatternFlowMplsTrafficClass, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsTrafficClass) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClass) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClass) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClass) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsTrafficClass) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowMplsTrafficClass) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsTrafficClass) FromJson(value string) error {
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

func (obj *patternFlowMplsTrafficClass) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClass) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsTrafficClass) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsTrafficClass) Clone() (PatternFlowMplsTrafficClass, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsTrafficClass()
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

func (obj *patternFlowMplsTrafficClass) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowMplsTrafficClass is traffic class
type PatternFlowMplsTrafficClass interface {
	Validation
	// msg marshals PatternFlowMplsTrafficClass to protobuf object *otg.PatternFlowMplsTrafficClass
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsTrafficClass
	// setMsg unmarshals PatternFlowMplsTrafficClass from protobuf object *otg.PatternFlowMplsTrafficClass
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsTrafficClass) PatternFlowMplsTrafficClass
	// provides marshal interface
	Marshal() marshalPatternFlowMplsTrafficClass
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsTrafficClass
	// validate validates PatternFlowMplsTrafficClass
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsTrafficClass, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowMplsTrafficClassChoiceEnum, set in PatternFlowMplsTrafficClass
	Choice() PatternFlowMplsTrafficClassChoiceEnum
	// setChoice assigns PatternFlowMplsTrafficClassChoiceEnum provided by user to PatternFlowMplsTrafficClass
	setChoice(value PatternFlowMplsTrafficClassChoiceEnum) PatternFlowMplsTrafficClass
	// HasChoice checks if Choice has been set in PatternFlowMplsTrafficClass
	HasChoice() bool
	// Value returns uint32, set in PatternFlowMplsTrafficClass.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowMplsTrafficClass
	SetValue(value uint32) PatternFlowMplsTrafficClass
	// HasValue checks if Value has been set in PatternFlowMplsTrafficClass
	HasValue() bool
	// Values returns []uint32, set in PatternFlowMplsTrafficClass.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowMplsTrafficClass
	SetValues(value []uint32) PatternFlowMplsTrafficClass
	// Increment returns PatternFlowMplsTrafficClassCounter, set in PatternFlowMplsTrafficClass.
	// PatternFlowMplsTrafficClassCounter is integer counter pattern
	Increment() PatternFlowMplsTrafficClassCounter
	// SetIncrement assigns PatternFlowMplsTrafficClassCounter provided by user to PatternFlowMplsTrafficClass.
	// PatternFlowMplsTrafficClassCounter is integer counter pattern
	SetIncrement(value PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClass
	// HasIncrement checks if Increment has been set in PatternFlowMplsTrafficClass
	HasIncrement() bool
	// Decrement returns PatternFlowMplsTrafficClassCounter, set in PatternFlowMplsTrafficClass.
	// PatternFlowMplsTrafficClassCounter is integer counter pattern
	Decrement() PatternFlowMplsTrafficClassCounter
	// SetDecrement assigns PatternFlowMplsTrafficClassCounter provided by user to PatternFlowMplsTrafficClass.
	// PatternFlowMplsTrafficClassCounter is integer counter pattern
	SetDecrement(value PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClass
	// HasDecrement checks if Decrement has been set in PatternFlowMplsTrafficClass
	HasDecrement() bool
	// MetricTags returns PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIterIter, set in PatternFlowMplsTrafficClass
	MetricTags() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	setNil()
}

type PatternFlowMplsTrafficClassChoiceEnum string

// Enum of Choice on PatternFlowMplsTrafficClass
var PatternFlowMplsTrafficClassChoice = struct {
	VALUE     PatternFlowMplsTrafficClassChoiceEnum
	VALUES    PatternFlowMplsTrafficClassChoiceEnum
	INCREMENT PatternFlowMplsTrafficClassChoiceEnum
	DECREMENT PatternFlowMplsTrafficClassChoiceEnum
}{
	VALUE:     PatternFlowMplsTrafficClassChoiceEnum("value"),
	VALUES:    PatternFlowMplsTrafficClassChoiceEnum("values"),
	INCREMENT: PatternFlowMplsTrafficClassChoiceEnum("increment"),
	DECREMENT: PatternFlowMplsTrafficClassChoiceEnum("decrement"),
}

func (obj *patternFlowMplsTrafficClass) Choice() PatternFlowMplsTrafficClassChoiceEnum {
	return PatternFlowMplsTrafficClassChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowMplsTrafficClass) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowMplsTrafficClass) setChoice(value PatternFlowMplsTrafficClassChoiceEnum) PatternFlowMplsTrafficClass {
	intValue, ok := otg.PatternFlowMplsTrafficClass_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowMplsTrafficClassChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowMplsTrafficClass_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowMplsTrafficClassChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowMplsTrafficClassChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowMplsTrafficClassChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowMplsTrafficClassCounter().msg()
	}

	if value == PatternFlowMplsTrafficClassChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowMplsTrafficClassCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsTrafficClass) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowMplsTrafficClassChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsTrafficClass) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowMplsTrafficClass object
func (obj *patternFlowMplsTrafficClass) SetValue(value uint32) PatternFlowMplsTrafficClass {
	obj.setChoice(PatternFlowMplsTrafficClassChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowMplsTrafficClass) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowMplsTrafficClass object
func (obj *patternFlowMplsTrafficClass) SetValues(value []uint32) PatternFlowMplsTrafficClass {
	obj.setChoice(PatternFlowMplsTrafficClassChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowMplsTrafficClassCounter
func (obj *patternFlowMplsTrafficClass) Increment() PatternFlowMplsTrafficClassCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowMplsTrafficClassChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowMplsTrafficClassCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowMplsTrafficClassCounter
func (obj *patternFlowMplsTrafficClass) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowMplsTrafficClassCounter value in the PatternFlowMplsTrafficClass object
func (obj *patternFlowMplsTrafficClass) SetIncrement(value PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClass {
	obj.setChoice(PatternFlowMplsTrafficClassChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowMplsTrafficClassCounter
func (obj *patternFlowMplsTrafficClass) Decrement() PatternFlowMplsTrafficClassCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowMplsTrafficClassChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowMplsTrafficClassCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowMplsTrafficClassCounter
func (obj *patternFlowMplsTrafficClass) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowMplsTrafficClassCounter value in the PatternFlowMplsTrafficClass object
func (obj *patternFlowMplsTrafficClass) SetDecrement(value PatternFlowMplsTrafficClassCounter) PatternFlowMplsTrafficClass {
	obj.setChoice(PatternFlowMplsTrafficClassChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowMplsTrafficClassMetricTag
func (obj *patternFlowMplsTrafficClass) MetricTags() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowMplsTrafficClassMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter struct {
	obj                                       *patternFlowMplsTrafficClass
	patternFlowMplsTrafficClassMetricTagSlice []PatternFlowMplsTrafficClassMetricTag
	fieldPtr                                  *[]*otg.PatternFlowMplsTrafficClassMetricTag
}

func newPatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter(ptr *[]*otg.PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	return &patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter{fieldPtr: ptr}
}

type PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter interface {
	setMsg(*patternFlowMplsTrafficClass) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	Items() []PatternFlowMplsTrafficClassMetricTag
	Add() PatternFlowMplsTrafficClassMetricTag
	Append(items ...PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	Set(index int, newObj PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	Clear() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	clearHolderSlice() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
	appendHolderSlice(item PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter
}

func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) setMsg(msg *patternFlowMplsTrafficClass) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowMplsTrafficClassMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) Items() []PatternFlowMplsTrafficClassMetricTag {
	return obj.patternFlowMplsTrafficClassMetricTagSlice
}

func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) Add() PatternFlowMplsTrafficClassMetricTag {
	newObj := &otg.PatternFlowMplsTrafficClassMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowMplsTrafficClassMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowMplsTrafficClassMetricTagSlice = append(obj.patternFlowMplsTrafficClassMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) Append(items ...PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowMplsTrafficClassMetricTagSlice = append(obj.patternFlowMplsTrafficClassMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) Set(index int, newObj PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowMplsTrafficClassMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) Clear() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowMplsTrafficClassMetricTag{}
		obj.patternFlowMplsTrafficClassMetricTagSlice = []PatternFlowMplsTrafficClassMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) clearHolderSlice() PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	if len(obj.patternFlowMplsTrafficClassMetricTagSlice) > 0 {
		obj.patternFlowMplsTrafficClassMetricTagSlice = []PatternFlowMplsTrafficClassMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter) appendHolderSlice(item PatternFlowMplsTrafficClassMetricTag) PatternFlowMplsTrafficClassPatternFlowMplsTrafficClassMetricTagIter {
	obj.patternFlowMplsTrafficClassMetricTagSlice = append(obj.patternFlowMplsTrafficClassMetricTagSlice, item)
	return obj
}

func (obj *patternFlowMplsTrafficClass) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsTrafficClass.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowMplsTrafficClass.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowMplsTrafficClassMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowMplsTrafficClass) setDefault() {
	var choices_set int = 0
	var choice PatternFlowMplsTrafficClassChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowMplsTrafficClassChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowMplsTrafficClassChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowMplsTrafficClassChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowMplsTrafficClassChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowMplsTrafficClassChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowMplsTrafficClass")
			}
		} else {
			intVal := otg.PatternFlowMplsTrafficClass_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowMplsTrafficClass_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
