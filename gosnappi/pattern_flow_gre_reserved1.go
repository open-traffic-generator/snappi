package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGreReserved1 *****
type patternFlowGreReserved1 struct {
	validation
	obj              *otg.PatternFlowGreReserved1
	marshaller       marshalPatternFlowGreReserved1
	unMarshaller     unMarshalPatternFlowGreReserved1
	incrementHolder  PatternFlowGreReserved1Counter
	decrementHolder  PatternFlowGreReserved1Counter
	metricTagsHolder PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
}

func NewPatternFlowGreReserved1() PatternFlowGreReserved1 {
	obj := patternFlowGreReserved1{obj: &otg.PatternFlowGreReserved1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGreReserved1) msg() *otg.PatternFlowGreReserved1 {
	return obj.obj
}

func (obj *patternFlowGreReserved1) setMsg(msg *otg.PatternFlowGreReserved1) PatternFlowGreReserved1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGreReserved1 struct {
	obj *patternFlowGreReserved1
}

type marshalPatternFlowGreReserved1 interface {
	// ToProto marshals PatternFlowGreReserved1 to protobuf object *otg.PatternFlowGreReserved1
	ToProto() (*otg.PatternFlowGreReserved1, error)
	// ToPbText marshals PatternFlowGreReserved1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGreReserved1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGreReserved1 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGreReserved1 struct {
	obj *patternFlowGreReserved1
}

type unMarshalPatternFlowGreReserved1 interface {
	// FromProto unmarshals PatternFlowGreReserved1 from protobuf object *otg.PatternFlowGreReserved1
	FromProto(msg *otg.PatternFlowGreReserved1) (PatternFlowGreReserved1, error)
	// FromPbText unmarshals PatternFlowGreReserved1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGreReserved1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGreReserved1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGreReserved1) Marshal() marshalPatternFlowGreReserved1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGreReserved1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGreReserved1) Unmarshal() unMarshalPatternFlowGreReserved1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGreReserved1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGreReserved1) ToProto() (*otg.PatternFlowGreReserved1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGreReserved1) FromProto(msg *otg.PatternFlowGreReserved1) (PatternFlowGreReserved1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGreReserved1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1) FromPbText(value string) error {
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

func (m *marshalpatternFlowGreReserved1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1) FromYaml(value string) error {
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

func (m *marshalpatternFlowGreReserved1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGreReserved1) FromJson(value string) error {
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

func (obj *patternFlowGreReserved1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGreReserved1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGreReserved1) Clone() (PatternFlowGreReserved1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGreReserved1()
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

func (obj *patternFlowGreReserved1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGreReserved1 is optional reserved field. Only present if the checksum_present bit is set.
type PatternFlowGreReserved1 interface {
	Validation
	// msg marshals PatternFlowGreReserved1 to protobuf object *otg.PatternFlowGreReserved1
	// and doesn't set defaults
	msg() *otg.PatternFlowGreReserved1
	// setMsg unmarshals PatternFlowGreReserved1 from protobuf object *otg.PatternFlowGreReserved1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGreReserved1) PatternFlowGreReserved1
	// provides marshal interface
	Marshal() marshalPatternFlowGreReserved1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGreReserved1
	// validate validates PatternFlowGreReserved1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGreReserved1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGreReserved1ChoiceEnum, set in PatternFlowGreReserved1
	Choice() PatternFlowGreReserved1ChoiceEnum
	// setChoice assigns PatternFlowGreReserved1ChoiceEnum provided by user to PatternFlowGreReserved1
	setChoice(value PatternFlowGreReserved1ChoiceEnum) PatternFlowGreReserved1
	// HasChoice checks if Choice has been set in PatternFlowGreReserved1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGreReserved1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGreReserved1
	SetValue(value uint32) PatternFlowGreReserved1
	// HasValue checks if Value has been set in PatternFlowGreReserved1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGreReserved1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGreReserved1
	SetValues(value []uint32) PatternFlowGreReserved1
	// Increment returns PatternFlowGreReserved1Counter, set in PatternFlowGreReserved1.
	// PatternFlowGreReserved1Counter is integer counter pattern
	Increment() PatternFlowGreReserved1Counter
	// SetIncrement assigns PatternFlowGreReserved1Counter provided by user to PatternFlowGreReserved1.
	// PatternFlowGreReserved1Counter is integer counter pattern
	SetIncrement(value PatternFlowGreReserved1Counter) PatternFlowGreReserved1
	// HasIncrement checks if Increment has been set in PatternFlowGreReserved1
	HasIncrement() bool
	// Decrement returns PatternFlowGreReserved1Counter, set in PatternFlowGreReserved1.
	// PatternFlowGreReserved1Counter is integer counter pattern
	Decrement() PatternFlowGreReserved1Counter
	// SetDecrement assigns PatternFlowGreReserved1Counter provided by user to PatternFlowGreReserved1.
	// PatternFlowGreReserved1Counter is integer counter pattern
	SetDecrement(value PatternFlowGreReserved1Counter) PatternFlowGreReserved1
	// HasDecrement checks if Decrement has been set in PatternFlowGreReserved1
	HasDecrement() bool
	// MetricTags returns PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIterIter, set in PatternFlowGreReserved1
	MetricTags() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	setNil()
}

type PatternFlowGreReserved1ChoiceEnum string

// Enum of Choice on PatternFlowGreReserved1
var PatternFlowGreReserved1Choice = struct {
	VALUE     PatternFlowGreReserved1ChoiceEnum
	VALUES    PatternFlowGreReserved1ChoiceEnum
	INCREMENT PatternFlowGreReserved1ChoiceEnum
	DECREMENT PatternFlowGreReserved1ChoiceEnum
}{
	VALUE:     PatternFlowGreReserved1ChoiceEnum("value"),
	VALUES:    PatternFlowGreReserved1ChoiceEnum("values"),
	INCREMENT: PatternFlowGreReserved1ChoiceEnum("increment"),
	DECREMENT: PatternFlowGreReserved1ChoiceEnum("decrement"),
}

func (obj *patternFlowGreReserved1) Choice() PatternFlowGreReserved1ChoiceEnum {
	return PatternFlowGreReserved1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGreReserved1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGreReserved1) setChoice(value PatternFlowGreReserved1ChoiceEnum) PatternFlowGreReserved1 {
	intValue, ok := otg.PatternFlowGreReserved1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGreReserved1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGreReserved1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGreReserved1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGreReserved1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGreReserved1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGreReserved1Counter().msg()
	}

	if value == PatternFlowGreReserved1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGreReserved1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreReserved1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGreReserved1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGreReserved1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGreReserved1 object
func (obj *patternFlowGreReserved1) SetValue(value uint32) PatternFlowGreReserved1 {
	obj.setChoice(PatternFlowGreReserved1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGreReserved1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGreReserved1 object
func (obj *patternFlowGreReserved1) SetValues(value []uint32) PatternFlowGreReserved1 {
	obj.setChoice(PatternFlowGreReserved1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGreReserved1Counter
func (obj *patternFlowGreReserved1) Increment() PatternFlowGreReserved1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGreReserved1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGreReserved1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGreReserved1Counter
func (obj *patternFlowGreReserved1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGreReserved1Counter value in the PatternFlowGreReserved1 object
func (obj *patternFlowGreReserved1) SetIncrement(value PatternFlowGreReserved1Counter) PatternFlowGreReserved1 {
	obj.setChoice(PatternFlowGreReserved1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGreReserved1Counter
func (obj *patternFlowGreReserved1) Decrement() PatternFlowGreReserved1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGreReserved1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGreReserved1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGreReserved1Counter
func (obj *patternFlowGreReserved1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGreReserved1Counter value in the PatternFlowGreReserved1 object
func (obj *patternFlowGreReserved1) SetDecrement(value PatternFlowGreReserved1Counter) PatternFlowGreReserved1 {
	obj.setChoice(PatternFlowGreReserved1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGreReserved1MetricTag
func (obj *patternFlowGreReserved1) MetricTags() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGreReserved1MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter struct {
	obj                                   *patternFlowGreReserved1
	patternFlowGreReserved1MetricTagSlice []PatternFlowGreReserved1MetricTag
	fieldPtr                              *[]*otg.PatternFlowGreReserved1MetricTag
}

func newPatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter(ptr *[]*otg.PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	return &patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter{fieldPtr: ptr}
}

type PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter interface {
	setMsg(*patternFlowGreReserved1) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	Items() []PatternFlowGreReserved1MetricTag
	Add() PatternFlowGreReserved1MetricTag
	Append(items ...PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	Set(index int, newObj PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	Clear() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	clearHolderSlice() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
	appendHolderSlice(item PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter
}

func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) setMsg(msg *patternFlowGreReserved1) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGreReserved1MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) Items() []PatternFlowGreReserved1MetricTag {
	return obj.patternFlowGreReserved1MetricTagSlice
}

func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) Add() PatternFlowGreReserved1MetricTag {
	newObj := &otg.PatternFlowGreReserved1MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGreReserved1MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGreReserved1MetricTagSlice = append(obj.patternFlowGreReserved1MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) Append(items ...PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGreReserved1MetricTagSlice = append(obj.patternFlowGreReserved1MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) Set(index int, newObj PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGreReserved1MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) Clear() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGreReserved1MetricTag{}
		obj.patternFlowGreReserved1MetricTagSlice = []PatternFlowGreReserved1MetricTag{}
	}
	return obj
}
func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) clearHolderSlice() PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	if len(obj.patternFlowGreReserved1MetricTagSlice) > 0 {
		obj.patternFlowGreReserved1MetricTagSlice = []PatternFlowGreReserved1MetricTag{}
	}
	return obj
}
func (obj *patternFlowGreReserved1PatternFlowGreReserved1MetricTagIter) appendHolderSlice(item PatternFlowGreReserved1MetricTag) PatternFlowGreReserved1PatternFlowGreReserved1MetricTagIter {
	obj.patternFlowGreReserved1MetricTagSlice = append(obj.patternFlowGreReserved1MetricTagSlice, item)
	return obj
}

func (obj *patternFlowGreReserved1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGreReserved1.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGreReserved1.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGreReserved1MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGreReserved1) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGreReserved1Choice.VALUE)

	}

}
