package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1Type *****
type patternFlowIgmpv1Type struct {
	validation
	obj              *otg.PatternFlowIgmpv1Type
	marshaller       marshalPatternFlowIgmpv1Type
	unMarshaller     unMarshalPatternFlowIgmpv1Type
	incrementHolder  PatternFlowIgmpv1TypeCounter
	decrementHolder  PatternFlowIgmpv1TypeCounter
	metricTagsHolder PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
}

func NewPatternFlowIgmpv1Type() PatternFlowIgmpv1Type {
	obj := patternFlowIgmpv1Type{obj: &otg.PatternFlowIgmpv1Type{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1Type) msg() *otg.PatternFlowIgmpv1Type {
	return obj.obj
}

func (obj *patternFlowIgmpv1Type) setMsg(msg *otg.PatternFlowIgmpv1Type) PatternFlowIgmpv1Type {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1Type struct {
	obj *patternFlowIgmpv1Type
}

type marshalPatternFlowIgmpv1Type interface {
	// ToProto marshals PatternFlowIgmpv1Type to protobuf object *otg.PatternFlowIgmpv1Type
	ToProto() (*otg.PatternFlowIgmpv1Type, error)
	// ToPbText marshals PatternFlowIgmpv1Type to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1Type to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1Type to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1Type struct {
	obj *patternFlowIgmpv1Type
}

type unMarshalPatternFlowIgmpv1Type interface {
	// FromProto unmarshals PatternFlowIgmpv1Type from protobuf object *otg.PatternFlowIgmpv1Type
	FromProto(msg *otg.PatternFlowIgmpv1Type) (PatternFlowIgmpv1Type, error)
	// FromPbText unmarshals PatternFlowIgmpv1Type from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1Type from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1Type from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1Type) Marshal() marshalPatternFlowIgmpv1Type {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1Type{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1Type) Unmarshal() unMarshalPatternFlowIgmpv1Type {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1Type{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1Type) ToProto() (*otg.PatternFlowIgmpv1Type, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1Type) FromProto(msg *otg.PatternFlowIgmpv1Type) (PatternFlowIgmpv1Type, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1Type) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Type) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1Type) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Type) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1Type) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Type) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1Type) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Type) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Type) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1Type) Clone() (PatternFlowIgmpv1Type, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1Type()
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

func (obj *patternFlowIgmpv1Type) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv1Type is type of message
type PatternFlowIgmpv1Type interface {
	Validation
	// msg marshals PatternFlowIgmpv1Type to protobuf object *otg.PatternFlowIgmpv1Type
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1Type
	// setMsg unmarshals PatternFlowIgmpv1Type from protobuf object *otg.PatternFlowIgmpv1Type
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1Type) PatternFlowIgmpv1Type
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1Type
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1Type
	// validate validates PatternFlowIgmpv1Type
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1Type, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv1TypeChoiceEnum, set in PatternFlowIgmpv1Type
	Choice() PatternFlowIgmpv1TypeChoiceEnum
	// setChoice assigns PatternFlowIgmpv1TypeChoiceEnum provided by user to PatternFlowIgmpv1Type
	setChoice(value PatternFlowIgmpv1TypeChoiceEnum) PatternFlowIgmpv1Type
	// HasChoice checks if Choice has been set in PatternFlowIgmpv1Type
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv1Type.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv1Type
	SetValue(value uint32) PatternFlowIgmpv1Type
	// HasValue checks if Value has been set in PatternFlowIgmpv1Type
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv1Type.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv1Type
	SetValues(value []uint32) PatternFlowIgmpv1Type
	// Increment returns PatternFlowIgmpv1TypeCounter, set in PatternFlowIgmpv1Type.
	// PatternFlowIgmpv1TypeCounter is integer counter pattern
	Increment() PatternFlowIgmpv1TypeCounter
	// SetIncrement assigns PatternFlowIgmpv1TypeCounter provided by user to PatternFlowIgmpv1Type.
	// PatternFlowIgmpv1TypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1Type
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv1Type
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv1TypeCounter, set in PatternFlowIgmpv1Type.
	// PatternFlowIgmpv1TypeCounter is integer counter pattern
	Decrement() PatternFlowIgmpv1TypeCounter
	// SetDecrement assigns PatternFlowIgmpv1TypeCounter provided by user to PatternFlowIgmpv1Type.
	// PatternFlowIgmpv1TypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1Type
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv1Type
	HasDecrement() bool
	// MetricTags returns PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIterIter, set in PatternFlowIgmpv1Type
	MetricTags() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	setNil()
}

type PatternFlowIgmpv1TypeChoiceEnum string

// Enum of Choice on PatternFlowIgmpv1Type
var PatternFlowIgmpv1TypeChoice = struct {
	VALUE     PatternFlowIgmpv1TypeChoiceEnum
	VALUES    PatternFlowIgmpv1TypeChoiceEnum
	INCREMENT PatternFlowIgmpv1TypeChoiceEnum
	DECREMENT PatternFlowIgmpv1TypeChoiceEnum
}{
	VALUE:     PatternFlowIgmpv1TypeChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv1TypeChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv1TypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv1TypeChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv1Type) Choice() PatternFlowIgmpv1TypeChoiceEnum {
	return PatternFlowIgmpv1TypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv1Type) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv1Type) setChoice(value PatternFlowIgmpv1TypeChoiceEnum) PatternFlowIgmpv1Type {
	intValue, ok := otg.PatternFlowIgmpv1Type_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1TypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1Type_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv1TypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv1TypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv1TypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv1TypeCounter().msg()
	}

	if value == PatternFlowIgmpv1TypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv1TypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Type) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv1TypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Type) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv1Type object
func (obj *patternFlowIgmpv1Type) SetValue(value uint32) PatternFlowIgmpv1Type {
	obj.setChoice(PatternFlowIgmpv1TypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv1Type) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv1Type object
func (obj *patternFlowIgmpv1Type) SetValues(value []uint32) PatternFlowIgmpv1Type {
	obj.setChoice(PatternFlowIgmpv1TypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv1TypeCounter
func (obj *patternFlowIgmpv1Type) Increment() PatternFlowIgmpv1TypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv1TypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv1TypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv1TypeCounter
func (obj *patternFlowIgmpv1Type) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv1TypeCounter value in the PatternFlowIgmpv1Type object
func (obj *patternFlowIgmpv1Type) SetIncrement(value PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1Type {
	obj.setChoice(PatternFlowIgmpv1TypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1TypeCounter
func (obj *patternFlowIgmpv1Type) Decrement() PatternFlowIgmpv1TypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv1TypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv1TypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1TypeCounter
func (obj *patternFlowIgmpv1Type) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv1TypeCounter value in the PatternFlowIgmpv1Type object
func (obj *patternFlowIgmpv1Type) SetDecrement(value PatternFlowIgmpv1TypeCounter) PatternFlowIgmpv1Type {
	obj.setChoice(PatternFlowIgmpv1TypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIgmpv1TypeMetricTag
func (obj *patternFlowIgmpv1Type) MetricTags() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIgmpv1TypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter struct {
	obj                                 *patternFlowIgmpv1Type
	patternFlowIgmpv1TypeMetricTagSlice []PatternFlowIgmpv1TypeMetricTag
	fieldPtr                            *[]*otg.PatternFlowIgmpv1TypeMetricTag
}

func newPatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter(ptr *[]*otg.PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	return &patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter interface {
	setMsg(*patternFlowIgmpv1Type) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	Items() []PatternFlowIgmpv1TypeMetricTag
	Add() PatternFlowIgmpv1TypeMetricTag
	Append(items ...PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	Set(index int, newObj PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	Clear() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	clearHolderSlice() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
	appendHolderSlice(item PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter
}

func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) setMsg(msg *patternFlowIgmpv1Type) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIgmpv1TypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) Items() []PatternFlowIgmpv1TypeMetricTag {
	return obj.patternFlowIgmpv1TypeMetricTagSlice
}

func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) Add() PatternFlowIgmpv1TypeMetricTag {
	newObj := &otg.PatternFlowIgmpv1TypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIgmpv1TypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIgmpv1TypeMetricTagSlice = append(obj.patternFlowIgmpv1TypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) Append(items ...PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIgmpv1TypeMetricTagSlice = append(obj.patternFlowIgmpv1TypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) Set(index int, newObj PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIgmpv1TypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) Clear() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIgmpv1TypeMetricTag{}
		obj.patternFlowIgmpv1TypeMetricTagSlice = []PatternFlowIgmpv1TypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) clearHolderSlice() PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	if len(obj.patternFlowIgmpv1TypeMetricTagSlice) > 0 {
		obj.patternFlowIgmpv1TypeMetricTagSlice = []PatternFlowIgmpv1TypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter) appendHolderSlice(item PatternFlowIgmpv1TypeMetricTag) PatternFlowIgmpv1TypePatternFlowIgmpv1TypeMetricTagIter {
	obj.patternFlowIgmpv1TypeMetricTagSlice = append(obj.patternFlowIgmpv1TypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIgmpv1Type) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Type.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv1Type.Values <= 15 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIgmpv1TypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIgmpv1Type) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIgmpv1TypeChoice.VALUE)

	}

}
