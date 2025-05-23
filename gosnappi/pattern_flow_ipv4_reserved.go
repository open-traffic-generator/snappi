package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Reserved *****
type patternFlowIpv4Reserved struct {
	validation
	obj              *otg.PatternFlowIpv4Reserved
	marshaller       marshalPatternFlowIpv4Reserved
	unMarshaller     unMarshalPatternFlowIpv4Reserved
	incrementHolder  PatternFlowIpv4ReservedCounter
	decrementHolder  PatternFlowIpv4ReservedCounter
	metricTagsHolder PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
}

func NewPatternFlowIpv4Reserved() PatternFlowIpv4Reserved {
	obj := patternFlowIpv4Reserved{obj: &otg.PatternFlowIpv4Reserved{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Reserved) msg() *otg.PatternFlowIpv4Reserved {
	return obj.obj
}

func (obj *patternFlowIpv4Reserved) setMsg(msg *otg.PatternFlowIpv4Reserved) PatternFlowIpv4Reserved {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Reserved struct {
	obj *patternFlowIpv4Reserved
}

type marshalPatternFlowIpv4Reserved interface {
	// ToProto marshals PatternFlowIpv4Reserved to protobuf object *otg.PatternFlowIpv4Reserved
	ToProto() (*otg.PatternFlowIpv4Reserved, error)
	// ToPbText marshals PatternFlowIpv4Reserved to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Reserved to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Reserved to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4Reserved to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4Reserved struct {
	obj *patternFlowIpv4Reserved
}

type unMarshalPatternFlowIpv4Reserved interface {
	// FromProto unmarshals PatternFlowIpv4Reserved from protobuf object *otg.PatternFlowIpv4Reserved
	FromProto(msg *otg.PatternFlowIpv4Reserved) (PatternFlowIpv4Reserved, error)
	// FromPbText unmarshals PatternFlowIpv4Reserved from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Reserved from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Reserved from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Reserved) Marshal() marshalPatternFlowIpv4Reserved {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Reserved{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Reserved) Unmarshal() unMarshalPatternFlowIpv4Reserved {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Reserved{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Reserved) ToProto() (*otg.PatternFlowIpv4Reserved, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Reserved) FromProto(msg *otg.PatternFlowIpv4Reserved) (PatternFlowIpv4Reserved, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Reserved) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Reserved) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Reserved) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Reserved) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Reserved) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4Reserved) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Reserved) FromJson(value string) error {
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

func (obj *patternFlowIpv4Reserved) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Reserved) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Reserved) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Reserved) Clone() (PatternFlowIpv4Reserved, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Reserved()
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

func (obj *patternFlowIpv4Reserved) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Reserved is reserved flag.
type PatternFlowIpv4Reserved interface {
	Validation
	// msg marshals PatternFlowIpv4Reserved to protobuf object *otg.PatternFlowIpv4Reserved
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Reserved
	// setMsg unmarshals PatternFlowIpv4Reserved from protobuf object *otg.PatternFlowIpv4Reserved
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Reserved) PatternFlowIpv4Reserved
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Reserved
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Reserved
	// validate validates PatternFlowIpv4Reserved
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Reserved, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4ReservedChoiceEnum, set in PatternFlowIpv4Reserved
	Choice() PatternFlowIpv4ReservedChoiceEnum
	// setChoice assigns PatternFlowIpv4ReservedChoiceEnum provided by user to PatternFlowIpv4Reserved
	setChoice(value PatternFlowIpv4ReservedChoiceEnum) PatternFlowIpv4Reserved
	// HasChoice checks if Choice has been set in PatternFlowIpv4Reserved
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4Reserved.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4Reserved
	SetValue(value uint32) PatternFlowIpv4Reserved
	// HasValue checks if Value has been set in PatternFlowIpv4Reserved
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4Reserved.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4Reserved
	SetValues(value []uint32) PatternFlowIpv4Reserved
	// Increment returns PatternFlowIpv4ReservedCounter, set in PatternFlowIpv4Reserved.
	// PatternFlowIpv4ReservedCounter is integer counter pattern
	Increment() PatternFlowIpv4ReservedCounter
	// SetIncrement assigns PatternFlowIpv4ReservedCounter provided by user to PatternFlowIpv4Reserved.
	// PatternFlowIpv4ReservedCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4ReservedCounter) PatternFlowIpv4Reserved
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Reserved
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4ReservedCounter, set in PatternFlowIpv4Reserved.
	// PatternFlowIpv4ReservedCounter is integer counter pattern
	Decrement() PatternFlowIpv4ReservedCounter
	// SetDecrement assigns PatternFlowIpv4ReservedCounter provided by user to PatternFlowIpv4Reserved.
	// PatternFlowIpv4ReservedCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4ReservedCounter) PatternFlowIpv4Reserved
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Reserved
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIterIter, set in PatternFlowIpv4Reserved
	MetricTags() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	setNil()
}

type PatternFlowIpv4ReservedChoiceEnum string

// Enum of Choice on PatternFlowIpv4Reserved
var PatternFlowIpv4ReservedChoice = struct {
	VALUE     PatternFlowIpv4ReservedChoiceEnum
	VALUES    PatternFlowIpv4ReservedChoiceEnum
	INCREMENT PatternFlowIpv4ReservedChoiceEnum
	DECREMENT PatternFlowIpv4ReservedChoiceEnum
}{
	VALUE:     PatternFlowIpv4ReservedChoiceEnum("value"),
	VALUES:    PatternFlowIpv4ReservedChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4ReservedChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4ReservedChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4Reserved) Choice() PatternFlowIpv4ReservedChoiceEnum {
	return PatternFlowIpv4ReservedChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Reserved) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Reserved) setChoice(value PatternFlowIpv4ReservedChoiceEnum) PatternFlowIpv4Reserved {
	intValue, ok := otg.PatternFlowIpv4Reserved_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4ReservedChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Reserved_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4ReservedChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4ReservedChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4ReservedChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4ReservedCounter().msg()
	}

	if value == PatternFlowIpv4ReservedChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4ReservedCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Reserved) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4ReservedChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Reserved) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4Reserved object
func (obj *patternFlowIpv4Reserved) SetValue(value uint32) PatternFlowIpv4Reserved {
	obj.setChoice(PatternFlowIpv4ReservedChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4Reserved) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4Reserved object
func (obj *patternFlowIpv4Reserved) SetValues(value []uint32) PatternFlowIpv4Reserved {
	obj.setChoice(PatternFlowIpv4ReservedChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4ReservedCounter
func (obj *patternFlowIpv4Reserved) Increment() PatternFlowIpv4ReservedCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4ReservedChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4ReservedCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4ReservedCounter
func (obj *patternFlowIpv4Reserved) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4ReservedCounter value in the PatternFlowIpv4Reserved object
func (obj *patternFlowIpv4Reserved) SetIncrement(value PatternFlowIpv4ReservedCounter) PatternFlowIpv4Reserved {
	obj.setChoice(PatternFlowIpv4ReservedChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4ReservedCounter
func (obj *patternFlowIpv4Reserved) Decrement() PatternFlowIpv4ReservedCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4ReservedChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4ReservedCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4ReservedCounter
func (obj *patternFlowIpv4Reserved) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4ReservedCounter value in the PatternFlowIpv4Reserved object
func (obj *patternFlowIpv4Reserved) SetDecrement(value PatternFlowIpv4ReservedCounter) PatternFlowIpv4Reserved {
	obj.setChoice(PatternFlowIpv4ReservedChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4ReservedMetricTag
func (obj *patternFlowIpv4Reserved) MetricTags() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4ReservedMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter struct {
	obj                                   *patternFlowIpv4Reserved
	patternFlowIpv4ReservedMetricTagSlice []PatternFlowIpv4ReservedMetricTag
	fieldPtr                              *[]*otg.PatternFlowIpv4ReservedMetricTag
}

func newPatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter(ptr *[]*otg.PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	return &patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter interface {
	setMsg(*patternFlowIpv4Reserved) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	Items() []PatternFlowIpv4ReservedMetricTag
	Add() PatternFlowIpv4ReservedMetricTag
	Append(items ...PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	Set(index int, newObj PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	Clear() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	clearHolderSlice() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
	appendHolderSlice(item PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter
}

func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) setMsg(msg *patternFlowIpv4Reserved) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4ReservedMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) Items() []PatternFlowIpv4ReservedMetricTag {
	return obj.patternFlowIpv4ReservedMetricTagSlice
}

func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) Add() PatternFlowIpv4ReservedMetricTag {
	newObj := &otg.PatternFlowIpv4ReservedMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4ReservedMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4ReservedMetricTagSlice = append(obj.patternFlowIpv4ReservedMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) Append(items ...PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4ReservedMetricTagSlice = append(obj.patternFlowIpv4ReservedMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) Set(index int, newObj PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4ReservedMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) Clear() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4ReservedMetricTag{}
		obj.patternFlowIpv4ReservedMetricTagSlice = []PatternFlowIpv4ReservedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) clearHolderSlice() PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	if len(obj.patternFlowIpv4ReservedMetricTagSlice) > 0 {
		obj.patternFlowIpv4ReservedMetricTagSlice = []PatternFlowIpv4ReservedMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter) appendHolderSlice(item PatternFlowIpv4ReservedMetricTag) PatternFlowIpv4ReservedPatternFlowIpv4ReservedMetricTagIter {
	obj.patternFlowIpv4ReservedMetricTagSlice = append(obj.patternFlowIpv4ReservedMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4Reserved) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4Reserved.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4Reserved.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4ReservedMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4Reserved) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4ReservedChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4ReservedChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4ReservedChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4ReservedChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4ReservedChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4ReservedChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4Reserved")
			}
		} else {
			intVal := otg.PatternFlowIpv4Reserved_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4Reserved_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
