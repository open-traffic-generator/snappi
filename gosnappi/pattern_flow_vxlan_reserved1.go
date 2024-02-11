package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanReserved1 *****
type patternFlowVxlanReserved1 struct {
	validation
	obj              *otg.PatternFlowVxlanReserved1
	marshaller       marshalPatternFlowVxlanReserved1
	unMarshaller     unMarshalPatternFlowVxlanReserved1
	incrementHolder  PatternFlowVxlanReserved1Counter
	decrementHolder  PatternFlowVxlanReserved1Counter
	metricTagsHolder PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
}

func NewPatternFlowVxlanReserved1() PatternFlowVxlanReserved1 {
	obj := patternFlowVxlanReserved1{obj: &otg.PatternFlowVxlanReserved1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved1) msg() *otg.PatternFlowVxlanReserved1 {
	return obj.obj
}

func (obj *patternFlowVxlanReserved1) setMsg(msg *otg.PatternFlowVxlanReserved1) PatternFlowVxlanReserved1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved1 struct {
	obj *patternFlowVxlanReserved1
}

type marshalPatternFlowVxlanReserved1 interface {
	// ToProto marshals PatternFlowVxlanReserved1 to protobuf object *otg.PatternFlowVxlanReserved1
	ToProto() (*otg.PatternFlowVxlanReserved1, error)
	// ToPbText marshals PatternFlowVxlanReserved1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved1 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanReserved1 struct {
	obj *patternFlowVxlanReserved1
}

type unMarshalPatternFlowVxlanReserved1 interface {
	// FromProto unmarshals PatternFlowVxlanReserved1 from protobuf object *otg.PatternFlowVxlanReserved1
	FromProto(msg *otg.PatternFlowVxlanReserved1) (PatternFlowVxlanReserved1, error)
	// FromPbText unmarshals PatternFlowVxlanReserved1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved1) Marshal() marshalPatternFlowVxlanReserved1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved1) Unmarshal() unMarshalPatternFlowVxlanReserved1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved1) ToProto() (*otg.PatternFlowVxlanReserved1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved1) FromProto(msg *otg.PatternFlowVxlanReserved1) (PatternFlowVxlanReserved1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved1) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved1) Clone() (PatternFlowVxlanReserved1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved1()
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

func (obj *patternFlowVxlanReserved1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVxlanReserved1 is reserved field
type PatternFlowVxlanReserved1 interface {
	Validation
	// msg marshals PatternFlowVxlanReserved1 to protobuf object *otg.PatternFlowVxlanReserved1
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved1
	// setMsg unmarshals PatternFlowVxlanReserved1 from protobuf object *otg.PatternFlowVxlanReserved1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved1) PatternFlowVxlanReserved1
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved1
	// validate validates PatternFlowVxlanReserved1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVxlanReserved1ChoiceEnum, set in PatternFlowVxlanReserved1
	Choice() PatternFlowVxlanReserved1ChoiceEnum
	// setChoice assigns PatternFlowVxlanReserved1ChoiceEnum provided by user to PatternFlowVxlanReserved1
	setChoice(value PatternFlowVxlanReserved1ChoiceEnum) PatternFlowVxlanReserved1
	// HasChoice checks if Choice has been set in PatternFlowVxlanReserved1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVxlanReserved1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVxlanReserved1
	SetValue(value uint32) PatternFlowVxlanReserved1
	// HasValue checks if Value has been set in PatternFlowVxlanReserved1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVxlanReserved1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVxlanReserved1
	SetValues(value []uint32) PatternFlowVxlanReserved1
	// Increment returns PatternFlowVxlanReserved1Counter, set in PatternFlowVxlanReserved1.
	// PatternFlowVxlanReserved1Counter is integer counter pattern
	Increment() PatternFlowVxlanReserved1Counter
	// SetIncrement assigns PatternFlowVxlanReserved1Counter provided by user to PatternFlowVxlanReserved1.
	// PatternFlowVxlanReserved1Counter is integer counter pattern
	SetIncrement(value PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1
	// HasIncrement checks if Increment has been set in PatternFlowVxlanReserved1
	HasIncrement() bool
	// Decrement returns PatternFlowVxlanReserved1Counter, set in PatternFlowVxlanReserved1.
	// PatternFlowVxlanReserved1Counter is integer counter pattern
	Decrement() PatternFlowVxlanReserved1Counter
	// SetDecrement assigns PatternFlowVxlanReserved1Counter provided by user to PatternFlowVxlanReserved1.
	// PatternFlowVxlanReserved1Counter is integer counter pattern
	SetDecrement(value PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1
	// HasDecrement checks if Decrement has been set in PatternFlowVxlanReserved1
	HasDecrement() bool
	// MetricTags returns PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIterIter, set in PatternFlowVxlanReserved1
	MetricTags() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	setNil()
}

type PatternFlowVxlanReserved1ChoiceEnum string

// Enum of Choice on PatternFlowVxlanReserved1
var PatternFlowVxlanReserved1Choice = struct {
	VALUE     PatternFlowVxlanReserved1ChoiceEnum
	VALUES    PatternFlowVxlanReserved1ChoiceEnum
	INCREMENT PatternFlowVxlanReserved1ChoiceEnum
	DECREMENT PatternFlowVxlanReserved1ChoiceEnum
}{
	VALUE:     PatternFlowVxlanReserved1ChoiceEnum("value"),
	VALUES:    PatternFlowVxlanReserved1ChoiceEnum("values"),
	INCREMENT: PatternFlowVxlanReserved1ChoiceEnum("increment"),
	DECREMENT: PatternFlowVxlanReserved1ChoiceEnum("decrement"),
}

func (obj *patternFlowVxlanReserved1) Choice() PatternFlowVxlanReserved1ChoiceEnum {
	return PatternFlowVxlanReserved1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVxlanReserved1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVxlanReserved1) setChoice(value PatternFlowVxlanReserved1ChoiceEnum) PatternFlowVxlanReserved1 {
	intValue, ok := otg.PatternFlowVxlanReserved1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVxlanReserved1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVxlanReserved1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVxlanReserved1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVxlanReserved1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVxlanReserved1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVxlanReserved1Counter().msg()
	}

	if value == PatternFlowVxlanReserved1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVxlanReserved1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanReserved1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVxlanReserved1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanReserved1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVxlanReserved1 object
func (obj *patternFlowVxlanReserved1) SetValue(value uint32) PatternFlowVxlanReserved1 {
	obj.setChoice(PatternFlowVxlanReserved1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVxlanReserved1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVxlanReserved1 object
func (obj *patternFlowVxlanReserved1) SetValues(value []uint32) PatternFlowVxlanReserved1 {
	obj.setChoice(PatternFlowVxlanReserved1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVxlanReserved1Counter
func (obj *patternFlowVxlanReserved1) Increment() PatternFlowVxlanReserved1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVxlanReserved1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVxlanReserved1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVxlanReserved1Counter
func (obj *patternFlowVxlanReserved1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVxlanReserved1Counter value in the PatternFlowVxlanReserved1 object
func (obj *patternFlowVxlanReserved1) SetIncrement(value PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1 {
	obj.setChoice(PatternFlowVxlanReserved1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVxlanReserved1Counter
func (obj *patternFlowVxlanReserved1) Decrement() PatternFlowVxlanReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVxlanReserved1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVxlanReserved1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVxlanReserved1Counter
func (obj *patternFlowVxlanReserved1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVxlanReserved1Counter value in the PatternFlowVxlanReserved1 object
func (obj *patternFlowVxlanReserved1) SetDecrement(value PatternFlowVxlanReserved1Counter) PatternFlowVxlanReserved1 {
	obj.setChoice(PatternFlowVxlanReserved1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVxlanReserved1MetricTag
func (obj *patternFlowVxlanReserved1) MetricTags() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVxlanReserved1MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter struct {
	obj                                     *patternFlowVxlanReserved1
	patternFlowVxlanReserved1MetricTagSlice []PatternFlowVxlanReserved1MetricTag
	fieldPtr                                *[]*otg.PatternFlowVxlanReserved1MetricTag
}

func newPatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter(ptr *[]*otg.PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	return &patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter{fieldPtr: ptr}
}

type PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter interface {
	setMsg(*patternFlowVxlanReserved1) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	Items() []PatternFlowVxlanReserved1MetricTag
	Add() PatternFlowVxlanReserved1MetricTag
	Append(items ...PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	Set(index int, newObj PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	Clear() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	clearHolderSlice() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
	appendHolderSlice(item PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter
}

func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) setMsg(msg *patternFlowVxlanReserved1) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVxlanReserved1MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) Items() []PatternFlowVxlanReserved1MetricTag {
	return obj.patternFlowVxlanReserved1MetricTagSlice
}

func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) Add() PatternFlowVxlanReserved1MetricTag {
	newObj := &otg.PatternFlowVxlanReserved1MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVxlanReserved1MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVxlanReserved1MetricTagSlice = append(obj.patternFlowVxlanReserved1MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) Append(items ...PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVxlanReserved1MetricTagSlice = append(obj.patternFlowVxlanReserved1MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) Set(index int, newObj PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVxlanReserved1MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) Clear() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVxlanReserved1MetricTag{}
		obj.patternFlowVxlanReserved1MetricTagSlice = []PatternFlowVxlanReserved1MetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) clearHolderSlice() PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	if len(obj.patternFlowVxlanReserved1MetricTagSlice) > 0 {
		obj.patternFlowVxlanReserved1MetricTagSlice = []PatternFlowVxlanReserved1MetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter) appendHolderSlice(item PatternFlowVxlanReserved1MetricTag) PatternFlowVxlanReserved1PatternFlowVxlanReserved1MetricTagIter {
	obj.patternFlowVxlanReserved1MetricTagSlice = append(obj.patternFlowVxlanReserved1MetricTagSlice, item)
	return obj
}

func (obj *patternFlowVxlanReserved1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved1.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVxlanReserved1.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVxlanReserved1MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVxlanReserved1) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowVxlanReserved1Choice.VALUE)

	}

}
