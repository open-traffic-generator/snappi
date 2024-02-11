package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIgmpv1Version *****
type patternFlowIgmpv1Version struct {
	validation
	obj              *otg.PatternFlowIgmpv1Version
	marshaller       marshalPatternFlowIgmpv1Version
	unMarshaller     unMarshalPatternFlowIgmpv1Version
	incrementHolder  PatternFlowIgmpv1VersionCounter
	decrementHolder  PatternFlowIgmpv1VersionCounter
	metricTagsHolder PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
}

func NewPatternFlowIgmpv1Version() PatternFlowIgmpv1Version {
	obj := patternFlowIgmpv1Version{obj: &otg.PatternFlowIgmpv1Version{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIgmpv1Version) msg() *otg.PatternFlowIgmpv1Version {
	return obj.obj
}

func (obj *patternFlowIgmpv1Version) setMsg(msg *otg.PatternFlowIgmpv1Version) PatternFlowIgmpv1Version {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIgmpv1Version struct {
	obj *patternFlowIgmpv1Version
}

type marshalPatternFlowIgmpv1Version interface {
	// ToProto marshals PatternFlowIgmpv1Version to protobuf object *otg.PatternFlowIgmpv1Version
	ToProto() (*otg.PatternFlowIgmpv1Version, error)
	// ToPbText marshals PatternFlowIgmpv1Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIgmpv1Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIgmpv1Version to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIgmpv1Version struct {
	obj *patternFlowIgmpv1Version
}

type unMarshalPatternFlowIgmpv1Version interface {
	// FromProto unmarshals PatternFlowIgmpv1Version from protobuf object *otg.PatternFlowIgmpv1Version
	FromProto(msg *otg.PatternFlowIgmpv1Version) (PatternFlowIgmpv1Version, error)
	// FromPbText unmarshals PatternFlowIgmpv1Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIgmpv1Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIgmpv1Version from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIgmpv1Version) Marshal() marshalPatternFlowIgmpv1Version {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIgmpv1Version{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIgmpv1Version) Unmarshal() unMarshalPatternFlowIgmpv1Version {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIgmpv1Version{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIgmpv1Version) ToProto() (*otg.PatternFlowIgmpv1Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIgmpv1Version) FromProto(msg *otg.PatternFlowIgmpv1Version) (PatternFlowIgmpv1Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIgmpv1Version) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Version) FromPbText(value string) error {
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

func (m *marshalpatternFlowIgmpv1Version) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Version) FromYaml(value string) error {
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

func (m *marshalpatternFlowIgmpv1Version) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIgmpv1Version) FromJson(value string) error {
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

func (obj *patternFlowIgmpv1Version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIgmpv1Version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIgmpv1Version) Clone() (PatternFlowIgmpv1Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIgmpv1Version()
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

func (obj *patternFlowIgmpv1Version) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIgmpv1Version is version number
type PatternFlowIgmpv1Version interface {
	Validation
	// msg marshals PatternFlowIgmpv1Version to protobuf object *otg.PatternFlowIgmpv1Version
	// and doesn't set defaults
	msg() *otg.PatternFlowIgmpv1Version
	// setMsg unmarshals PatternFlowIgmpv1Version from protobuf object *otg.PatternFlowIgmpv1Version
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIgmpv1Version) PatternFlowIgmpv1Version
	// provides marshal interface
	Marshal() marshalPatternFlowIgmpv1Version
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIgmpv1Version
	// validate validates PatternFlowIgmpv1Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIgmpv1Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIgmpv1VersionChoiceEnum, set in PatternFlowIgmpv1Version
	Choice() PatternFlowIgmpv1VersionChoiceEnum
	// setChoice assigns PatternFlowIgmpv1VersionChoiceEnum provided by user to PatternFlowIgmpv1Version
	setChoice(value PatternFlowIgmpv1VersionChoiceEnum) PatternFlowIgmpv1Version
	// HasChoice checks if Choice has been set in PatternFlowIgmpv1Version
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIgmpv1Version.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIgmpv1Version
	SetValue(value uint32) PatternFlowIgmpv1Version
	// HasValue checks if Value has been set in PatternFlowIgmpv1Version
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIgmpv1Version.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIgmpv1Version
	SetValues(value []uint32) PatternFlowIgmpv1Version
	// Increment returns PatternFlowIgmpv1VersionCounter, set in PatternFlowIgmpv1Version.
	// PatternFlowIgmpv1VersionCounter is integer counter pattern
	Increment() PatternFlowIgmpv1VersionCounter
	// SetIncrement assigns PatternFlowIgmpv1VersionCounter provided by user to PatternFlowIgmpv1Version.
	// PatternFlowIgmpv1VersionCounter is integer counter pattern
	SetIncrement(value PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1Version
	// HasIncrement checks if Increment has been set in PatternFlowIgmpv1Version
	HasIncrement() bool
	// Decrement returns PatternFlowIgmpv1VersionCounter, set in PatternFlowIgmpv1Version.
	// PatternFlowIgmpv1VersionCounter is integer counter pattern
	Decrement() PatternFlowIgmpv1VersionCounter
	// SetDecrement assigns PatternFlowIgmpv1VersionCounter provided by user to PatternFlowIgmpv1Version.
	// PatternFlowIgmpv1VersionCounter is integer counter pattern
	SetDecrement(value PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1Version
	// HasDecrement checks if Decrement has been set in PatternFlowIgmpv1Version
	HasDecrement() bool
	// MetricTags returns PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIterIter, set in PatternFlowIgmpv1Version
	MetricTags() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	setNil()
}

type PatternFlowIgmpv1VersionChoiceEnum string

// Enum of Choice on PatternFlowIgmpv1Version
var PatternFlowIgmpv1VersionChoice = struct {
	VALUE     PatternFlowIgmpv1VersionChoiceEnum
	VALUES    PatternFlowIgmpv1VersionChoiceEnum
	INCREMENT PatternFlowIgmpv1VersionChoiceEnum
	DECREMENT PatternFlowIgmpv1VersionChoiceEnum
}{
	VALUE:     PatternFlowIgmpv1VersionChoiceEnum("value"),
	VALUES:    PatternFlowIgmpv1VersionChoiceEnum("values"),
	INCREMENT: PatternFlowIgmpv1VersionChoiceEnum("increment"),
	DECREMENT: PatternFlowIgmpv1VersionChoiceEnum("decrement"),
}

func (obj *patternFlowIgmpv1Version) Choice() PatternFlowIgmpv1VersionChoiceEnum {
	return PatternFlowIgmpv1VersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIgmpv1Version) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIgmpv1Version) setChoice(value PatternFlowIgmpv1VersionChoiceEnum) PatternFlowIgmpv1Version {
	intValue, ok := otg.PatternFlowIgmpv1Version_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIgmpv1VersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIgmpv1Version_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIgmpv1VersionChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIgmpv1VersionChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIgmpv1VersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIgmpv1VersionCounter().msg()
	}

	if value == PatternFlowIgmpv1VersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIgmpv1VersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Version) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIgmpv1VersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIgmpv1Version) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIgmpv1Version object
func (obj *patternFlowIgmpv1Version) SetValue(value uint32) PatternFlowIgmpv1Version {
	obj.setChoice(PatternFlowIgmpv1VersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIgmpv1Version) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIgmpv1Version object
func (obj *patternFlowIgmpv1Version) SetValues(value []uint32) PatternFlowIgmpv1Version {
	obj.setChoice(PatternFlowIgmpv1VersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIgmpv1VersionCounter
func (obj *patternFlowIgmpv1Version) Increment() PatternFlowIgmpv1VersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIgmpv1VersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIgmpv1VersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIgmpv1VersionCounter
func (obj *patternFlowIgmpv1Version) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIgmpv1VersionCounter value in the PatternFlowIgmpv1Version object
func (obj *patternFlowIgmpv1Version) SetIncrement(value PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1Version {
	obj.setChoice(PatternFlowIgmpv1VersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1VersionCounter
func (obj *patternFlowIgmpv1Version) Decrement() PatternFlowIgmpv1VersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIgmpv1VersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIgmpv1VersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIgmpv1VersionCounter
func (obj *patternFlowIgmpv1Version) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIgmpv1VersionCounter value in the PatternFlowIgmpv1Version object
func (obj *patternFlowIgmpv1Version) SetDecrement(value PatternFlowIgmpv1VersionCounter) PatternFlowIgmpv1Version {
	obj.setChoice(PatternFlowIgmpv1VersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIgmpv1VersionMetricTag
func (obj *patternFlowIgmpv1Version) MetricTags() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIgmpv1VersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter struct {
	obj                                    *patternFlowIgmpv1Version
	patternFlowIgmpv1VersionMetricTagSlice []PatternFlowIgmpv1VersionMetricTag
	fieldPtr                               *[]*otg.PatternFlowIgmpv1VersionMetricTag
}

func newPatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter(ptr *[]*otg.PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	return &patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter interface {
	setMsg(*patternFlowIgmpv1Version) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	Items() []PatternFlowIgmpv1VersionMetricTag
	Add() PatternFlowIgmpv1VersionMetricTag
	Append(items ...PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	Set(index int, newObj PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	Clear() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	clearHolderSlice() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
	appendHolderSlice(item PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter
}

func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) setMsg(msg *patternFlowIgmpv1Version) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIgmpv1VersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) Items() []PatternFlowIgmpv1VersionMetricTag {
	return obj.patternFlowIgmpv1VersionMetricTagSlice
}

func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) Add() PatternFlowIgmpv1VersionMetricTag {
	newObj := &otg.PatternFlowIgmpv1VersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIgmpv1VersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIgmpv1VersionMetricTagSlice = append(obj.patternFlowIgmpv1VersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) Append(items ...PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIgmpv1VersionMetricTagSlice = append(obj.patternFlowIgmpv1VersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) Set(index int, newObj PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIgmpv1VersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) Clear() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIgmpv1VersionMetricTag{}
		obj.patternFlowIgmpv1VersionMetricTagSlice = []PatternFlowIgmpv1VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) clearHolderSlice() PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	if len(obj.patternFlowIgmpv1VersionMetricTagSlice) > 0 {
		obj.patternFlowIgmpv1VersionMetricTagSlice = []PatternFlowIgmpv1VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter) appendHolderSlice(item PatternFlowIgmpv1VersionMetricTag) PatternFlowIgmpv1VersionPatternFlowIgmpv1VersionMetricTagIter {
	obj.patternFlowIgmpv1VersionMetricTagSlice = append(obj.patternFlowIgmpv1VersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIgmpv1Version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIgmpv1Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIgmpv1Version.Values <= 15 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIgmpv1VersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIgmpv1Version) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIgmpv1VersionChoice.VALUE)

	}

}
