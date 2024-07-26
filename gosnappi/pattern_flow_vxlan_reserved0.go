package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowVxlanReserved0 *****
type patternFlowVxlanReserved0 struct {
	validation
	obj              *otg.PatternFlowVxlanReserved0
	marshaller       marshalPatternFlowVxlanReserved0
	unMarshaller     unMarshalPatternFlowVxlanReserved0
	incrementHolder  PatternFlowVxlanReserved0Counter
	decrementHolder  PatternFlowVxlanReserved0Counter
	metricTagsHolder PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
}

func NewPatternFlowVxlanReserved0() PatternFlowVxlanReserved0 {
	obj := patternFlowVxlanReserved0{obj: &otg.PatternFlowVxlanReserved0{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowVxlanReserved0) msg() *otg.PatternFlowVxlanReserved0 {
	return obj.obj
}

func (obj *patternFlowVxlanReserved0) setMsg(msg *otg.PatternFlowVxlanReserved0) PatternFlowVxlanReserved0 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowVxlanReserved0 struct {
	obj *patternFlowVxlanReserved0
}

type marshalPatternFlowVxlanReserved0 interface {
	// ToProto marshals PatternFlowVxlanReserved0 to protobuf object *otg.PatternFlowVxlanReserved0
	ToProto() (*otg.PatternFlowVxlanReserved0, error)
	// ToPbText marshals PatternFlowVxlanReserved0 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowVxlanReserved0 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowVxlanReserved0 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowVxlanReserved0 struct {
	obj *patternFlowVxlanReserved0
}

type unMarshalPatternFlowVxlanReserved0 interface {
	// FromProto unmarshals PatternFlowVxlanReserved0 from protobuf object *otg.PatternFlowVxlanReserved0
	FromProto(msg *otg.PatternFlowVxlanReserved0) (PatternFlowVxlanReserved0, error)
	// FromPbText unmarshals PatternFlowVxlanReserved0 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowVxlanReserved0 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowVxlanReserved0 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowVxlanReserved0) Marshal() marshalPatternFlowVxlanReserved0 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowVxlanReserved0{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowVxlanReserved0) Unmarshal() unMarshalPatternFlowVxlanReserved0 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowVxlanReserved0{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowVxlanReserved0) ToProto() (*otg.PatternFlowVxlanReserved0, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowVxlanReserved0) FromProto(msg *otg.PatternFlowVxlanReserved0) (PatternFlowVxlanReserved0, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowVxlanReserved0) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0) FromPbText(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0) FromYaml(value string) error {
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

func (m *marshalpatternFlowVxlanReserved0) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowVxlanReserved0) FromJson(value string) error {
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

func (obj *patternFlowVxlanReserved0) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowVxlanReserved0) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowVxlanReserved0) Clone() (PatternFlowVxlanReserved0, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowVxlanReserved0()
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

func (obj *patternFlowVxlanReserved0) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowVxlanReserved0 is reserved field
type PatternFlowVxlanReserved0 interface {
	Validation
	// msg marshals PatternFlowVxlanReserved0 to protobuf object *otg.PatternFlowVxlanReserved0
	// and doesn't set defaults
	msg() *otg.PatternFlowVxlanReserved0
	// setMsg unmarshals PatternFlowVxlanReserved0 from protobuf object *otg.PatternFlowVxlanReserved0
	// and doesn't set defaults
	setMsg(*otg.PatternFlowVxlanReserved0) PatternFlowVxlanReserved0
	// provides marshal interface
	Marshal() marshalPatternFlowVxlanReserved0
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowVxlanReserved0
	// validate validates PatternFlowVxlanReserved0
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowVxlanReserved0, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowVxlanReserved0ChoiceEnum, set in PatternFlowVxlanReserved0
	Choice() PatternFlowVxlanReserved0ChoiceEnum
	// setChoice assigns PatternFlowVxlanReserved0ChoiceEnum provided by user to PatternFlowVxlanReserved0
	setChoice(value PatternFlowVxlanReserved0ChoiceEnum) PatternFlowVxlanReserved0
	// HasChoice checks if Choice has been set in PatternFlowVxlanReserved0
	HasChoice() bool
	// Value returns uint32, set in PatternFlowVxlanReserved0.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowVxlanReserved0
	SetValue(value uint32) PatternFlowVxlanReserved0
	// HasValue checks if Value has been set in PatternFlowVxlanReserved0
	HasValue() bool
	// Values returns []uint32, set in PatternFlowVxlanReserved0.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowVxlanReserved0
	SetValues(value []uint32) PatternFlowVxlanReserved0
	// Increment returns PatternFlowVxlanReserved0Counter, set in PatternFlowVxlanReserved0.
	// PatternFlowVxlanReserved0Counter is integer counter pattern
	Increment() PatternFlowVxlanReserved0Counter
	// SetIncrement assigns PatternFlowVxlanReserved0Counter provided by user to PatternFlowVxlanReserved0.
	// PatternFlowVxlanReserved0Counter is integer counter pattern
	SetIncrement(value PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0
	// HasIncrement checks if Increment has been set in PatternFlowVxlanReserved0
	HasIncrement() bool
	// Decrement returns PatternFlowVxlanReserved0Counter, set in PatternFlowVxlanReserved0.
	// PatternFlowVxlanReserved0Counter is integer counter pattern
	Decrement() PatternFlowVxlanReserved0Counter
	// SetDecrement assigns PatternFlowVxlanReserved0Counter provided by user to PatternFlowVxlanReserved0.
	// PatternFlowVxlanReserved0Counter is integer counter pattern
	SetDecrement(value PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0
	// HasDecrement checks if Decrement has been set in PatternFlowVxlanReserved0
	HasDecrement() bool
	// MetricTags returns PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIterIter, set in PatternFlowVxlanReserved0
	MetricTags() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	setNil()
}

type PatternFlowVxlanReserved0ChoiceEnum string

// Enum of Choice on PatternFlowVxlanReserved0
var PatternFlowVxlanReserved0Choice = struct {
	VALUE     PatternFlowVxlanReserved0ChoiceEnum
	VALUES    PatternFlowVxlanReserved0ChoiceEnum
	INCREMENT PatternFlowVxlanReserved0ChoiceEnum
	DECREMENT PatternFlowVxlanReserved0ChoiceEnum
}{
	VALUE:     PatternFlowVxlanReserved0ChoiceEnum("value"),
	VALUES:    PatternFlowVxlanReserved0ChoiceEnum("values"),
	INCREMENT: PatternFlowVxlanReserved0ChoiceEnum("increment"),
	DECREMENT: PatternFlowVxlanReserved0ChoiceEnum("decrement"),
}

func (obj *patternFlowVxlanReserved0) Choice() PatternFlowVxlanReserved0ChoiceEnum {
	return PatternFlowVxlanReserved0ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowVxlanReserved0) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowVxlanReserved0) setChoice(value PatternFlowVxlanReserved0ChoiceEnum) PatternFlowVxlanReserved0 {
	intValue, ok := otg.PatternFlowVxlanReserved0_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowVxlanReserved0ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowVxlanReserved0_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowVxlanReserved0Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowVxlanReserved0Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowVxlanReserved0Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowVxlanReserved0Counter().msg()
	}

	if value == PatternFlowVxlanReserved0Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowVxlanReserved0Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanReserved0) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowVxlanReserved0Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowVxlanReserved0) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowVxlanReserved0 object
func (obj *patternFlowVxlanReserved0) SetValue(value uint32) PatternFlowVxlanReserved0 {
	obj.setChoice(PatternFlowVxlanReserved0Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowVxlanReserved0) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowVxlanReserved0 object
func (obj *patternFlowVxlanReserved0) SetValues(value []uint32) PatternFlowVxlanReserved0 {
	obj.setChoice(PatternFlowVxlanReserved0Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowVxlanReserved0Counter
func (obj *patternFlowVxlanReserved0) Increment() PatternFlowVxlanReserved0Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowVxlanReserved0Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowVxlanReserved0Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowVxlanReserved0Counter
func (obj *patternFlowVxlanReserved0) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowVxlanReserved0Counter value in the PatternFlowVxlanReserved0 object
func (obj *patternFlowVxlanReserved0) SetIncrement(value PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0 {
	obj.setChoice(PatternFlowVxlanReserved0Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowVxlanReserved0Counter
func (obj *patternFlowVxlanReserved0) Decrement() PatternFlowVxlanReserved0Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowVxlanReserved0Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowVxlanReserved0Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowVxlanReserved0Counter
func (obj *patternFlowVxlanReserved0) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowVxlanReserved0Counter value in the PatternFlowVxlanReserved0 object
func (obj *patternFlowVxlanReserved0) SetDecrement(value PatternFlowVxlanReserved0Counter) PatternFlowVxlanReserved0 {
	obj.setChoice(PatternFlowVxlanReserved0Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowVxlanReserved0MetricTag
func (obj *patternFlowVxlanReserved0) MetricTags() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowVxlanReserved0MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter struct {
	obj                                     *patternFlowVxlanReserved0
	patternFlowVxlanReserved0MetricTagSlice []PatternFlowVxlanReserved0MetricTag
	fieldPtr                                *[]*otg.PatternFlowVxlanReserved0MetricTag
}

func newPatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter(ptr *[]*otg.PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	return &patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter{fieldPtr: ptr}
}

type PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter interface {
	setMsg(*patternFlowVxlanReserved0) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	Items() []PatternFlowVxlanReserved0MetricTag
	Add() PatternFlowVxlanReserved0MetricTag
	Append(items ...PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	Set(index int, newObj PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	Clear() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	clearHolderSlice() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
	appendHolderSlice(item PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter
}

func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) setMsg(msg *patternFlowVxlanReserved0) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowVxlanReserved0MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) Items() []PatternFlowVxlanReserved0MetricTag {
	return obj.patternFlowVxlanReserved0MetricTagSlice
}

func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) Add() PatternFlowVxlanReserved0MetricTag {
	newObj := &otg.PatternFlowVxlanReserved0MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowVxlanReserved0MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowVxlanReserved0MetricTagSlice = append(obj.patternFlowVxlanReserved0MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) Append(items ...PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowVxlanReserved0MetricTagSlice = append(obj.patternFlowVxlanReserved0MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) Set(index int, newObj PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowVxlanReserved0MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) Clear() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowVxlanReserved0MetricTag{}
		obj.patternFlowVxlanReserved0MetricTagSlice = []PatternFlowVxlanReserved0MetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) clearHolderSlice() PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	if len(obj.patternFlowVxlanReserved0MetricTagSlice) > 0 {
		obj.patternFlowVxlanReserved0MetricTagSlice = []PatternFlowVxlanReserved0MetricTag{}
	}
	return obj
}
func (obj *patternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter) appendHolderSlice(item PatternFlowVxlanReserved0MetricTag) PatternFlowVxlanReserved0PatternFlowVxlanReserved0MetricTagIter {
	obj.patternFlowVxlanReserved0MetricTagSlice = append(obj.patternFlowVxlanReserved0MetricTagSlice, item)
	return obj
}

func (obj *patternFlowVxlanReserved0) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowVxlanReserved0.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowVxlanReserved0.Values <= 16777215 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowVxlanReserved0MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowVxlanReserved0) setDefault() {
	var choices_set int = 0
	var choice PatternFlowVxlanReserved0ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowVxlanReserved0Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowVxlanReserved0Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowVxlanReserved0Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowVxlanReserved0Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowVxlanReserved0Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowVxlanReserved0")
			}
		} else {
			intVal := otg.PatternFlowVxlanReserved0_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowVxlanReserved0_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
