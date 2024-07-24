package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1Version *****
type patternFlowGtpv1Version struct {
	validation
	obj              *otg.PatternFlowGtpv1Version
	marshaller       marshalPatternFlowGtpv1Version
	unMarshaller     unMarshalPatternFlowGtpv1Version
	incrementHolder  PatternFlowGtpv1VersionCounter
	decrementHolder  PatternFlowGtpv1VersionCounter
	metricTagsHolder PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
}

func NewPatternFlowGtpv1Version() PatternFlowGtpv1Version {
	obj := patternFlowGtpv1Version{obj: &otg.PatternFlowGtpv1Version{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1Version) msg() *otg.PatternFlowGtpv1Version {
	return obj.obj
}

func (obj *patternFlowGtpv1Version) setMsg(msg *otg.PatternFlowGtpv1Version) PatternFlowGtpv1Version {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1Version struct {
	obj *patternFlowGtpv1Version
}

type marshalPatternFlowGtpv1Version interface {
	// ToProto marshals PatternFlowGtpv1Version to protobuf object *otg.PatternFlowGtpv1Version
	ToProto() (*otg.PatternFlowGtpv1Version, error)
	// ToPbText marshals PatternFlowGtpv1Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1Version to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1Version struct {
	obj *patternFlowGtpv1Version
}

type unMarshalPatternFlowGtpv1Version interface {
	// FromProto unmarshals PatternFlowGtpv1Version from protobuf object *otg.PatternFlowGtpv1Version
	FromProto(msg *otg.PatternFlowGtpv1Version) (PatternFlowGtpv1Version, error)
	// FromPbText unmarshals PatternFlowGtpv1Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1Version from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1Version) Marshal() marshalPatternFlowGtpv1Version {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1Version{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1Version) Unmarshal() unMarshalPatternFlowGtpv1Version {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1Version{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1Version) ToProto() (*otg.PatternFlowGtpv1Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1Version) FromProto(msg *otg.PatternFlowGtpv1Version) (PatternFlowGtpv1Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1Version) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Version) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1Version) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Version) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1Version) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Version) FromJson(value string) error {
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

func (obj *patternFlowGtpv1Version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1Version) Clone() (PatternFlowGtpv1Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1Version()
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

func (obj *patternFlowGtpv1Version) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1Version is gTPv1 version
type PatternFlowGtpv1Version interface {
	Validation
	// msg marshals PatternFlowGtpv1Version to protobuf object *otg.PatternFlowGtpv1Version
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1Version
	// setMsg unmarshals PatternFlowGtpv1Version from protobuf object *otg.PatternFlowGtpv1Version
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1Version) PatternFlowGtpv1Version
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1Version
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1Version
	// validate validates PatternFlowGtpv1Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1VersionChoiceEnum, set in PatternFlowGtpv1Version
	Choice() PatternFlowGtpv1VersionChoiceEnum
	// setChoice assigns PatternFlowGtpv1VersionChoiceEnum provided by user to PatternFlowGtpv1Version
	setChoice(value PatternFlowGtpv1VersionChoiceEnum) PatternFlowGtpv1Version
	// HasChoice checks if Choice has been set in PatternFlowGtpv1Version
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1Version.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1Version
	SetValue(value uint32) PatternFlowGtpv1Version
	// HasValue checks if Value has been set in PatternFlowGtpv1Version
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1Version.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1Version
	SetValues(value []uint32) PatternFlowGtpv1Version
	// Increment returns PatternFlowGtpv1VersionCounter, set in PatternFlowGtpv1Version.
	// PatternFlowGtpv1VersionCounter is integer counter pattern
	Increment() PatternFlowGtpv1VersionCounter
	// SetIncrement assigns PatternFlowGtpv1VersionCounter provided by user to PatternFlowGtpv1Version.
	// PatternFlowGtpv1VersionCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1VersionCounter) PatternFlowGtpv1Version
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1Version
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1VersionCounter, set in PatternFlowGtpv1Version.
	// PatternFlowGtpv1VersionCounter is integer counter pattern
	Decrement() PatternFlowGtpv1VersionCounter
	// SetDecrement assigns PatternFlowGtpv1VersionCounter provided by user to PatternFlowGtpv1Version.
	// PatternFlowGtpv1VersionCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1VersionCounter) PatternFlowGtpv1Version
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1Version
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIterIter, set in PatternFlowGtpv1Version
	MetricTags() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	setNil()
}

type PatternFlowGtpv1VersionChoiceEnum string

// Enum of Choice on PatternFlowGtpv1Version
var PatternFlowGtpv1VersionChoice = struct {
	VALUE     PatternFlowGtpv1VersionChoiceEnum
	VALUES    PatternFlowGtpv1VersionChoiceEnum
	INCREMENT PatternFlowGtpv1VersionChoiceEnum
	DECREMENT PatternFlowGtpv1VersionChoiceEnum
}{
	VALUE:     PatternFlowGtpv1VersionChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1VersionChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1VersionChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1VersionChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1Version) Choice() PatternFlowGtpv1VersionChoiceEnum {
	return PatternFlowGtpv1VersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1Version) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1Version) setChoice(value PatternFlowGtpv1VersionChoiceEnum) PatternFlowGtpv1Version {
	intValue, ok := otg.PatternFlowGtpv1Version_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1VersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1Version_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1VersionChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1VersionChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1VersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1VersionCounter().msg()
	}

	if value == PatternFlowGtpv1VersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1VersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Version) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1VersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Version) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1Version object
func (obj *patternFlowGtpv1Version) SetValue(value uint32) PatternFlowGtpv1Version {
	obj.setChoice(PatternFlowGtpv1VersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1Version) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1Version object
func (obj *patternFlowGtpv1Version) SetValues(value []uint32) PatternFlowGtpv1Version {
	obj.setChoice(PatternFlowGtpv1VersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1VersionCounter
func (obj *patternFlowGtpv1Version) Increment() PatternFlowGtpv1VersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1VersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1VersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1VersionCounter
func (obj *patternFlowGtpv1Version) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1VersionCounter value in the PatternFlowGtpv1Version object
func (obj *patternFlowGtpv1Version) SetIncrement(value PatternFlowGtpv1VersionCounter) PatternFlowGtpv1Version {
	obj.setChoice(PatternFlowGtpv1VersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1VersionCounter
func (obj *patternFlowGtpv1Version) Decrement() PatternFlowGtpv1VersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1VersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1VersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1VersionCounter
func (obj *patternFlowGtpv1Version) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1VersionCounter value in the PatternFlowGtpv1Version object
func (obj *patternFlowGtpv1Version) SetDecrement(value PatternFlowGtpv1VersionCounter) PatternFlowGtpv1Version {
	obj.setChoice(PatternFlowGtpv1VersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1VersionMetricTag
func (obj *patternFlowGtpv1Version) MetricTags() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1VersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter struct {
	obj                                   *patternFlowGtpv1Version
	patternFlowGtpv1VersionMetricTagSlice []PatternFlowGtpv1VersionMetricTag
	fieldPtr                              *[]*otg.PatternFlowGtpv1VersionMetricTag
}

func newPatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter(ptr *[]*otg.PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	return &patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter interface {
	setMsg(*patternFlowGtpv1Version) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	Items() []PatternFlowGtpv1VersionMetricTag
	Add() PatternFlowGtpv1VersionMetricTag
	Append(items ...PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	Set(index int, newObj PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	Clear() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	clearHolderSlice() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter
}

func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) setMsg(msg *patternFlowGtpv1Version) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1VersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) Items() []PatternFlowGtpv1VersionMetricTag {
	return obj.patternFlowGtpv1VersionMetricTagSlice
}

func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) Add() PatternFlowGtpv1VersionMetricTag {
	newObj := &otg.PatternFlowGtpv1VersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1VersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1VersionMetricTagSlice = append(obj.patternFlowGtpv1VersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) Append(items ...PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1VersionMetricTagSlice = append(obj.patternFlowGtpv1VersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) Set(index int, newObj PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1VersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) Clear() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1VersionMetricTag{}
		obj.patternFlowGtpv1VersionMetricTagSlice = []PatternFlowGtpv1VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) clearHolderSlice() PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	if len(obj.patternFlowGtpv1VersionMetricTagSlice) > 0 {
		obj.patternFlowGtpv1VersionMetricTagSlice = []PatternFlowGtpv1VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter) appendHolderSlice(item PatternFlowGtpv1VersionMetricTag) PatternFlowGtpv1VersionPatternFlowGtpv1VersionMetricTagIter {
	obj.patternFlowGtpv1VersionMetricTagSlice = append(obj.patternFlowGtpv1VersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1Version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1Version.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1Version.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1VersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1Version) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv1VersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv1VersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv1VersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv1VersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv1VersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv1VersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv1Version")
			}
		} else {
			intVal := otg.PatternFlowGtpv1Version_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv1Version_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
