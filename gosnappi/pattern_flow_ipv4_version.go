package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4Version *****
type patternFlowIpv4Version struct {
	validation
	obj              *otg.PatternFlowIpv4Version
	marshaller       marshalPatternFlowIpv4Version
	unMarshaller     unMarshalPatternFlowIpv4Version
	incrementHolder  PatternFlowIpv4VersionCounter
	decrementHolder  PatternFlowIpv4VersionCounter
	metricTagsHolder PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
}

func NewPatternFlowIpv4Version() PatternFlowIpv4Version {
	obj := patternFlowIpv4Version{obj: &otg.PatternFlowIpv4Version{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4Version) msg() *otg.PatternFlowIpv4Version {
	return obj.obj
}

func (obj *patternFlowIpv4Version) setMsg(msg *otg.PatternFlowIpv4Version) PatternFlowIpv4Version {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4Version struct {
	obj *patternFlowIpv4Version
}

type marshalPatternFlowIpv4Version interface {
	// ToProto marshals PatternFlowIpv4Version to protobuf object *otg.PatternFlowIpv4Version
	ToProto() (*otg.PatternFlowIpv4Version, error)
	// ToPbText marshals PatternFlowIpv4Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4Version to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4Version to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4Version struct {
	obj *patternFlowIpv4Version
}

type unMarshalPatternFlowIpv4Version interface {
	// FromProto unmarshals PatternFlowIpv4Version from protobuf object *otg.PatternFlowIpv4Version
	FromProto(msg *otg.PatternFlowIpv4Version) (PatternFlowIpv4Version, error)
	// FromPbText unmarshals PatternFlowIpv4Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4Version from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4Version) Marshal() marshalPatternFlowIpv4Version {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4Version{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4Version) Unmarshal() unMarshalPatternFlowIpv4Version {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4Version{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4Version) ToProto() (*otg.PatternFlowIpv4Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4Version) FromProto(msg *otg.PatternFlowIpv4Version) (PatternFlowIpv4Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4Version) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Version) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4Version) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Version) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4Version) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4Version) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4Version) FromJson(value string) error {
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

func (obj *patternFlowIpv4Version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4Version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4Version) Clone() (PatternFlowIpv4Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4Version()
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

func (obj *patternFlowIpv4Version) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4Version is version
type PatternFlowIpv4Version interface {
	Validation
	// msg marshals PatternFlowIpv4Version to protobuf object *otg.PatternFlowIpv4Version
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4Version
	// setMsg unmarshals PatternFlowIpv4Version from protobuf object *otg.PatternFlowIpv4Version
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4Version) PatternFlowIpv4Version
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4Version
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4Version
	// validate validates PatternFlowIpv4Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4VersionChoiceEnum, set in PatternFlowIpv4Version
	Choice() PatternFlowIpv4VersionChoiceEnum
	// setChoice assigns PatternFlowIpv4VersionChoiceEnum provided by user to PatternFlowIpv4Version
	setChoice(value PatternFlowIpv4VersionChoiceEnum) PatternFlowIpv4Version
	// HasChoice checks if Choice has been set in PatternFlowIpv4Version
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4Version.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4Version
	SetValue(value uint32) PatternFlowIpv4Version
	// HasValue checks if Value has been set in PatternFlowIpv4Version
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4Version.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4Version
	SetValues(value []uint32) PatternFlowIpv4Version
	// Increment returns PatternFlowIpv4VersionCounter, set in PatternFlowIpv4Version.
	// PatternFlowIpv4VersionCounter is integer counter pattern
	Increment() PatternFlowIpv4VersionCounter
	// SetIncrement assigns PatternFlowIpv4VersionCounter provided by user to PatternFlowIpv4Version.
	// PatternFlowIpv4VersionCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4VersionCounter) PatternFlowIpv4Version
	// HasIncrement checks if Increment has been set in PatternFlowIpv4Version
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4VersionCounter, set in PatternFlowIpv4Version.
	// PatternFlowIpv4VersionCounter is integer counter pattern
	Decrement() PatternFlowIpv4VersionCounter
	// SetDecrement assigns PatternFlowIpv4VersionCounter provided by user to PatternFlowIpv4Version.
	// PatternFlowIpv4VersionCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4VersionCounter) PatternFlowIpv4Version
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4Version
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIterIter, set in PatternFlowIpv4Version
	MetricTags() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	setNil()
}

type PatternFlowIpv4VersionChoiceEnum string

// Enum of Choice on PatternFlowIpv4Version
var PatternFlowIpv4VersionChoice = struct {
	VALUE     PatternFlowIpv4VersionChoiceEnum
	VALUES    PatternFlowIpv4VersionChoiceEnum
	INCREMENT PatternFlowIpv4VersionChoiceEnum
	DECREMENT PatternFlowIpv4VersionChoiceEnum
}{
	VALUE:     PatternFlowIpv4VersionChoiceEnum("value"),
	VALUES:    PatternFlowIpv4VersionChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4VersionChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4VersionChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4Version) Choice() PatternFlowIpv4VersionChoiceEnum {
	return PatternFlowIpv4VersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4Version) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4Version) setChoice(value PatternFlowIpv4VersionChoiceEnum) PatternFlowIpv4Version {
	intValue, ok := otg.PatternFlowIpv4Version_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4VersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4Version_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4VersionChoice.VALUE {
		defaultValue := uint32(4)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4VersionChoice.VALUES {
		defaultValue := []uint32{4}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4VersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4VersionCounter().msg()
	}

	if value == PatternFlowIpv4VersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4VersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Version) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4VersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4Version) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4Version object
func (obj *patternFlowIpv4Version) SetValue(value uint32) PatternFlowIpv4Version {
	obj.setChoice(PatternFlowIpv4VersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4Version) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{4})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4Version object
func (obj *patternFlowIpv4Version) SetValues(value []uint32) PatternFlowIpv4Version {
	obj.setChoice(PatternFlowIpv4VersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4VersionCounter
func (obj *patternFlowIpv4Version) Increment() PatternFlowIpv4VersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4VersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4VersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4VersionCounter
func (obj *patternFlowIpv4Version) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4VersionCounter value in the PatternFlowIpv4Version object
func (obj *patternFlowIpv4Version) SetIncrement(value PatternFlowIpv4VersionCounter) PatternFlowIpv4Version {
	obj.setChoice(PatternFlowIpv4VersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4VersionCounter
func (obj *patternFlowIpv4Version) Decrement() PatternFlowIpv4VersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4VersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4VersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4VersionCounter
func (obj *patternFlowIpv4Version) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4VersionCounter value in the PatternFlowIpv4Version object
func (obj *patternFlowIpv4Version) SetDecrement(value PatternFlowIpv4VersionCounter) PatternFlowIpv4Version {
	obj.setChoice(PatternFlowIpv4VersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4VersionMetricTag
func (obj *patternFlowIpv4Version) MetricTags() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4VersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter struct {
	obj                                  *patternFlowIpv4Version
	patternFlowIpv4VersionMetricTagSlice []PatternFlowIpv4VersionMetricTag
	fieldPtr                             *[]*otg.PatternFlowIpv4VersionMetricTag
}

func newPatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter(ptr *[]*otg.PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	return &patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter interface {
	setMsg(*patternFlowIpv4Version) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	Items() []PatternFlowIpv4VersionMetricTag
	Add() PatternFlowIpv4VersionMetricTag
	Append(items ...PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	Set(index int, newObj PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	Clear() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	clearHolderSlice() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
	appendHolderSlice(item PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter
}

func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) setMsg(msg *patternFlowIpv4Version) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4VersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) Items() []PatternFlowIpv4VersionMetricTag {
	return obj.patternFlowIpv4VersionMetricTagSlice
}

func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) Add() PatternFlowIpv4VersionMetricTag {
	newObj := &otg.PatternFlowIpv4VersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4VersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4VersionMetricTagSlice = append(obj.patternFlowIpv4VersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) Append(items ...PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4VersionMetricTagSlice = append(obj.patternFlowIpv4VersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) Set(index int, newObj PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4VersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) Clear() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4VersionMetricTag{}
		obj.patternFlowIpv4VersionMetricTagSlice = []PatternFlowIpv4VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) clearHolderSlice() PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	if len(obj.patternFlowIpv4VersionMetricTagSlice) > 0 {
		obj.patternFlowIpv4VersionMetricTagSlice = []PatternFlowIpv4VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter) appendHolderSlice(item PatternFlowIpv4VersionMetricTag) PatternFlowIpv4VersionPatternFlowIpv4VersionMetricTagIter {
	obj.patternFlowIpv4VersionMetricTagSlice = append(obj.patternFlowIpv4VersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4Version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4Version.Values <= 15 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4VersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4Version) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4VersionChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4VersionChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4VersionChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4VersionChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4VersionChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4VersionChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4Version")
			}
		} else {
			intVal := otg.PatternFlowIpv4Version_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4Version_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
