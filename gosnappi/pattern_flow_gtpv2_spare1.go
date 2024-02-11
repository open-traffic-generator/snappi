package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2Spare1 *****
type patternFlowGtpv2Spare1 struct {
	validation
	obj              *otg.PatternFlowGtpv2Spare1
	marshaller       marshalPatternFlowGtpv2Spare1
	unMarshaller     unMarshalPatternFlowGtpv2Spare1
	incrementHolder  PatternFlowGtpv2Spare1Counter
	decrementHolder  PatternFlowGtpv2Spare1Counter
	metricTagsHolder PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
}

func NewPatternFlowGtpv2Spare1() PatternFlowGtpv2Spare1 {
	obj := patternFlowGtpv2Spare1{obj: &otg.PatternFlowGtpv2Spare1{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Spare1) msg() *otg.PatternFlowGtpv2Spare1 {
	return obj.obj
}

func (obj *patternFlowGtpv2Spare1) setMsg(msg *otg.PatternFlowGtpv2Spare1) PatternFlowGtpv2Spare1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Spare1 struct {
	obj *patternFlowGtpv2Spare1
}

type marshalPatternFlowGtpv2Spare1 interface {
	// ToProto marshals PatternFlowGtpv2Spare1 to protobuf object *otg.PatternFlowGtpv2Spare1
	ToProto() (*otg.PatternFlowGtpv2Spare1, error)
	// ToPbText marshals PatternFlowGtpv2Spare1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Spare1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Spare1 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2Spare1 struct {
	obj *patternFlowGtpv2Spare1
}

type unMarshalPatternFlowGtpv2Spare1 interface {
	// FromProto unmarshals PatternFlowGtpv2Spare1 from protobuf object *otg.PatternFlowGtpv2Spare1
	FromProto(msg *otg.PatternFlowGtpv2Spare1) (PatternFlowGtpv2Spare1, error)
	// FromPbText unmarshals PatternFlowGtpv2Spare1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Spare1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Spare1 from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Spare1) Marshal() marshalPatternFlowGtpv2Spare1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Spare1{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Spare1) Unmarshal() unMarshalPatternFlowGtpv2Spare1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Spare1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Spare1) ToProto() (*otg.PatternFlowGtpv2Spare1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Spare1) FromProto(msg *otg.PatternFlowGtpv2Spare1) (PatternFlowGtpv2Spare1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Spare1) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Spare1) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Spare1) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Spare1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Spare1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Spare1) Clone() (PatternFlowGtpv2Spare1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Spare1()
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

func (obj *patternFlowGtpv2Spare1) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2Spare1 is a 3-bit reserved field (must be 0).
type PatternFlowGtpv2Spare1 interface {
	Validation
	// msg marshals PatternFlowGtpv2Spare1 to protobuf object *otg.PatternFlowGtpv2Spare1
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Spare1
	// setMsg unmarshals PatternFlowGtpv2Spare1 from protobuf object *otg.PatternFlowGtpv2Spare1
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Spare1) PatternFlowGtpv2Spare1
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Spare1
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Spare1
	// validate validates PatternFlowGtpv2Spare1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Spare1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2Spare1ChoiceEnum, set in PatternFlowGtpv2Spare1
	Choice() PatternFlowGtpv2Spare1ChoiceEnum
	// setChoice assigns PatternFlowGtpv2Spare1ChoiceEnum provided by user to PatternFlowGtpv2Spare1
	setChoice(value PatternFlowGtpv2Spare1ChoiceEnum) PatternFlowGtpv2Spare1
	// HasChoice checks if Choice has been set in PatternFlowGtpv2Spare1
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2Spare1.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2Spare1
	SetValue(value uint32) PatternFlowGtpv2Spare1
	// HasValue checks if Value has been set in PatternFlowGtpv2Spare1
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2Spare1.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2Spare1
	SetValues(value []uint32) PatternFlowGtpv2Spare1
	// Increment returns PatternFlowGtpv2Spare1Counter, set in PatternFlowGtpv2Spare1.
	// PatternFlowGtpv2Spare1Counter is integer counter pattern
	Increment() PatternFlowGtpv2Spare1Counter
	// SetIncrement assigns PatternFlowGtpv2Spare1Counter provided by user to PatternFlowGtpv2Spare1.
	// PatternFlowGtpv2Spare1Counter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2Spare1
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2Spare1Counter, set in PatternFlowGtpv2Spare1.
	// PatternFlowGtpv2Spare1Counter is integer counter pattern
	Decrement() PatternFlowGtpv2Spare1Counter
	// SetDecrement assigns PatternFlowGtpv2Spare1Counter provided by user to PatternFlowGtpv2Spare1.
	// PatternFlowGtpv2Spare1Counter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2Spare1
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIterIter, set in PatternFlowGtpv2Spare1
	MetricTags() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	setNil()
}

type PatternFlowGtpv2Spare1ChoiceEnum string

// Enum of Choice on PatternFlowGtpv2Spare1
var PatternFlowGtpv2Spare1Choice = struct {
	VALUE     PatternFlowGtpv2Spare1ChoiceEnum
	VALUES    PatternFlowGtpv2Spare1ChoiceEnum
	INCREMENT PatternFlowGtpv2Spare1ChoiceEnum
	DECREMENT PatternFlowGtpv2Spare1ChoiceEnum
}{
	VALUE:     PatternFlowGtpv2Spare1ChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2Spare1ChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2Spare1ChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2Spare1ChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2Spare1) Choice() PatternFlowGtpv2Spare1ChoiceEnum {
	return PatternFlowGtpv2Spare1ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2Spare1) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2Spare1) setChoice(value PatternFlowGtpv2Spare1ChoiceEnum) PatternFlowGtpv2Spare1 {
	intValue, ok := otg.PatternFlowGtpv2Spare1_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2Spare1ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2Spare1_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2Spare1Choice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2Spare1Choice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2Spare1Choice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2Spare1Counter().msg()
	}

	if value == PatternFlowGtpv2Spare1Choice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2Spare1Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Spare1) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2Spare1Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Spare1) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2Spare1 object
func (obj *patternFlowGtpv2Spare1) SetValue(value uint32) PatternFlowGtpv2Spare1 {
	obj.setChoice(PatternFlowGtpv2Spare1Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2Spare1) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2Spare1 object
func (obj *patternFlowGtpv2Spare1) SetValues(value []uint32) PatternFlowGtpv2Spare1 {
	obj.setChoice(PatternFlowGtpv2Spare1Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2Spare1Counter
func (obj *patternFlowGtpv2Spare1) Increment() PatternFlowGtpv2Spare1Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2Spare1Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2Spare1Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2Spare1Counter
func (obj *patternFlowGtpv2Spare1) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2Spare1Counter value in the PatternFlowGtpv2Spare1 object
func (obj *patternFlowGtpv2Spare1) SetIncrement(value PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1 {
	obj.setChoice(PatternFlowGtpv2Spare1Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2Spare1Counter
func (obj *patternFlowGtpv2Spare1) Decrement() PatternFlowGtpv2Spare1Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2Spare1Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2Spare1Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2Spare1Counter
func (obj *patternFlowGtpv2Spare1) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2Spare1Counter value in the PatternFlowGtpv2Spare1 object
func (obj *patternFlowGtpv2Spare1) SetDecrement(value PatternFlowGtpv2Spare1Counter) PatternFlowGtpv2Spare1 {
	obj.setChoice(PatternFlowGtpv2Spare1Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2Spare1MetricTag
func (obj *patternFlowGtpv2Spare1) MetricTags() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2Spare1MetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter struct {
	obj                                  *patternFlowGtpv2Spare1
	patternFlowGtpv2Spare1MetricTagSlice []PatternFlowGtpv2Spare1MetricTag
	fieldPtr                             *[]*otg.PatternFlowGtpv2Spare1MetricTag
}

func newPatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter(ptr *[]*otg.PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	return &patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter interface {
	setMsg(*patternFlowGtpv2Spare1) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	Items() []PatternFlowGtpv2Spare1MetricTag
	Add() PatternFlowGtpv2Spare1MetricTag
	Append(items ...PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	Set(index int, newObj PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	Clear() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	clearHolderSlice() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
	appendHolderSlice(item PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter
}

func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) setMsg(msg *patternFlowGtpv2Spare1) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2Spare1MetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) Items() []PatternFlowGtpv2Spare1MetricTag {
	return obj.patternFlowGtpv2Spare1MetricTagSlice
}

func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) Add() PatternFlowGtpv2Spare1MetricTag {
	newObj := &otg.PatternFlowGtpv2Spare1MetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2Spare1MetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2Spare1MetricTagSlice = append(obj.patternFlowGtpv2Spare1MetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) Append(items ...PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2Spare1MetricTagSlice = append(obj.patternFlowGtpv2Spare1MetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) Set(index int, newObj PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2Spare1MetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) Clear() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2Spare1MetricTag{}
		obj.patternFlowGtpv2Spare1MetricTagSlice = []PatternFlowGtpv2Spare1MetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) clearHolderSlice() PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	if len(obj.patternFlowGtpv2Spare1MetricTagSlice) > 0 {
		obj.patternFlowGtpv2Spare1MetricTagSlice = []PatternFlowGtpv2Spare1MetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter) appendHolderSlice(item PatternFlowGtpv2Spare1MetricTag) PatternFlowGtpv2Spare1PatternFlowGtpv2Spare1MetricTagIter {
	obj.patternFlowGtpv2Spare1MetricTagSlice = append(obj.patternFlowGtpv2Spare1MetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2Spare1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Spare1.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2Spare1.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2Spare1MetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2Spare1) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv2Spare1Choice.VALUE)

	}

}
