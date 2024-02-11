package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv6Version *****
type patternFlowIpv6Version struct {
	validation
	obj              *otg.PatternFlowIpv6Version
	marshaller       marshalPatternFlowIpv6Version
	unMarshaller     unMarshalPatternFlowIpv6Version
	incrementHolder  PatternFlowIpv6VersionCounter
	decrementHolder  PatternFlowIpv6VersionCounter
	metricTagsHolder PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
}

func NewPatternFlowIpv6Version() PatternFlowIpv6Version {
	obj := patternFlowIpv6Version{obj: &otg.PatternFlowIpv6Version{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv6Version) msg() *otg.PatternFlowIpv6Version {
	return obj.obj
}

func (obj *patternFlowIpv6Version) setMsg(msg *otg.PatternFlowIpv6Version) PatternFlowIpv6Version {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv6Version struct {
	obj *patternFlowIpv6Version
}

type marshalPatternFlowIpv6Version interface {
	// ToProto marshals PatternFlowIpv6Version to protobuf object *otg.PatternFlowIpv6Version
	ToProto() (*otg.PatternFlowIpv6Version, error)
	// ToPbText marshals PatternFlowIpv6Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv6Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv6Version to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv6Version struct {
	obj *patternFlowIpv6Version
}

type unMarshalPatternFlowIpv6Version interface {
	// FromProto unmarshals PatternFlowIpv6Version from protobuf object *otg.PatternFlowIpv6Version
	FromProto(msg *otg.PatternFlowIpv6Version) (PatternFlowIpv6Version, error)
	// FromPbText unmarshals PatternFlowIpv6Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv6Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv6Version from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv6Version) Marshal() marshalPatternFlowIpv6Version {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv6Version{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv6Version) Unmarshal() unMarshalPatternFlowIpv6Version {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv6Version{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv6Version) ToProto() (*otg.PatternFlowIpv6Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv6Version) FromProto(msg *otg.PatternFlowIpv6Version) (PatternFlowIpv6Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv6Version) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Version) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv6Version) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Version) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv6Version) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv6Version) FromJson(value string) error {
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

func (obj *patternFlowIpv6Version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv6Version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv6Version) Clone() (PatternFlowIpv6Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv6Version()
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

func (obj *patternFlowIpv6Version) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv6Version is version number
type PatternFlowIpv6Version interface {
	Validation
	// msg marshals PatternFlowIpv6Version to protobuf object *otg.PatternFlowIpv6Version
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv6Version
	// setMsg unmarshals PatternFlowIpv6Version from protobuf object *otg.PatternFlowIpv6Version
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv6Version) PatternFlowIpv6Version
	// provides marshal interface
	Marshal() marshalPatternFlowIpv6Version
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv6Version
	// validate validates PatternFlowIpv6Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv6Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv6VersionChoiceEnum, set in PatternFlowIpv6Version
	Choice() PatternFlowIpv6VersionChoiceEnum
	// setChoice assigns PatternFlowIpv6VersionChoiceEnum provided by user to PatternFlowIpv6Version
	setChoice(value PatternFlowIpv6VersionChoiceEnum) PatternFlowIpv6Version
	// HasChoice checks if Choice has been set in PatternFlowIpv6Version
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv6Version.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv6Version
	SetValue(value uint32) PatternFlowIpv6Version
	// HasValue checks if Value has been set in PatternFlowIpv6Version
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv6Version.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv6Version
	SetValues(value []uint32) PatternFlowIpv6Version
	// Increment returns PatternFlowIpv6VersionCounter, set in PatternFlowIpv6Version.
	// PatternFlowIpv6VersionCounter is integer counter pattern
	Increment() PatternFlowIpv6VersionCounter
	// SetIncrement assigns PatternFlowIpv6VersionCounter provided by user to PatternFlowIpv6Version.
	// PatternFlowIpv6VersionCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv6VersionCounter) PatternFlowIpv6Version
	// HasIncrement checks if Increment has been set in PatternFlowIpv6Version
	HasIncrement() bool
	// Decrement returns PatternFlowIpv6VersionCounter, set in PatternFlowIpv6Version.
	// PatternFlowIpv6VersionCounter is integer counter pattern
	Decrement() PatternFlowIpv6VersionCounter
	// SetDecrement assigns PatternFlowIpv6VersionCounter provided by user to PatternFlowIpv6Version.
	// PatternFlowIpv6VersionCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv6VersionCounter) PatternFlowIpv6Version
	// HasDecrement checks if Decrement has been set in PatternFlowIpv6Version
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIterIter, set in PatternFlowIpv6Version
	MetricTags() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	setNil()
}

type PatternFlowIpv6VersionChoiceEnum string

// Enum of Choice on PatternFlowIpv6Version
var PatternFlowIpv6VersionChoice = struct {
	VALUE     PatternFlowIpv6VersionChoiceEnum
	VALUES    PatternFlowIpv6VersionChoiceEnum
	INCREMENT PatternFlowIpv6VersionChoiceEnum
	DECREMENT PatternFlowIpv6VersionChoiceEnum
}{
	VALUE:     PatternFlowIpv6VersionChoiceEnum("value"),
	VALUES:    PatternFlowIpv6VersionChoiceEnum("values"),
	INCREMENT: PatternFlowIpv6VersionChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv6VersionChoiceEnum("decrement"),
}

func (obj *patternFlowIpv6Version) Choice() PatternFlowIpv6VersionChoiceEnum {
	return PatternFlowIpv6VersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv6Version) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv6Version) setChoice(value PatternFlowIpv6VersionChoiceEnum) PatternFlowIpv6Version {
	intValue, ok := otg.PatternFlowIpv6Version_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv6VersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv6Version_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv6VersionChoice.VALUE {
		defaultValue := uint32(6)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv6VersionChoice.VALUES {
		defaultValue := []uint32{6}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv6VersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv6VersionCounter().msg()
	}

	if value == PatternFlowIpv6VersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv6VersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6Version) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv6VersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv6Version) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv6Version object
func (obj *patternFlowIpv6Version) SetValue(value uint32) PatternFlowIpv6Version {
	obj.setChoice(PatternFlowIpv6VersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv6Version) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{6})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv6Version object
func (obj *patternFlowIpv6Version) SetValues(value []uint32) PatternFlowIpv6Version {
	obj.setChoice(PatternFlowIpv6VersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv6VersionCounter
func (obj *patternFlowIpv6Version) Increment() PatternFlowIpv6VersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv6VersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv6VersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv6VersionCounter
func (obj *patternFlowIpv6Version) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv6VersionCounter value in the PatternFlowIpv6Version object
func (obj *patternFlowIpv6Version) SetIncrement(value PatternFlowIpv6VersionCounter) PatternFlowIpv6Version {
	obj.setChoice(PatternFlowIpv6VersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv6VersionCounter
func (obj *patternFlowIpv6Version) Decrement() PatternFlowIpv6VersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv6VersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv6VersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv6VersionCounter
func (obj *patternFlowIpv6Version) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv6VersionCounter value in the PatternFlowIpv6Version object
func (obj *patternFlowIpv6Version) SetDecrement(value PatternFlowIpv6VersionCounter) PatternFlowIpv6Version {
	obj.setChoice(PatternFlowIpv6VersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv6VersionMetricTag
func (obj *patternFlowIpv6Version) MetricTags() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv6VersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter struct {
	obj                                  *patternFlowIpv6Version
	patternFlowIpv6VersionMetricTagSlice []PatternFlowIpv6VersionMetricTag
	fieldPtr                             *[]*otg.PatternFlowIpv6VersionMetricTag
}

func newPatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter(ptr *[]*otg.PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	return &patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter interface {
	setMsg(*patternFlowIpv6Version) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	Items() []PatternFlowIpv6VersionMetricTag
	Add() PatternFlowIpv6VersionMetricTag
	Append(items ...PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	Set(index int, newObj PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	Clear() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	clearHolderSlice() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
	appendHolderSlice(item PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter
}

func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) setMsg(msg *patternFlowIpv6Version) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv6VersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) Items() []PatternFlowIpv6VersionMetricTag {
	return obj.patternFlowIpv6VersionMetricTagSlice
}

func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) Add() PatternFlowIpv6VersionMetricTag {
	newObj := &otg.PatternFlowIpv6VersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv6VersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv6VersionMetricTagSlice = append(obj.patternFlowIpv6VersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) Append(items ...PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv6VersionMetricTagSlice = append(obj.patternFlowIpv6VersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) Set(index int, newObj PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv6VersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) Clear() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv6VersionMetricTag{}
		obj.patternFlowIpv6VersionMetricTagSlice = []PatternFlowIpv6VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) clearHolderSlice() PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	if len(obj.patternFlowIpv6VersionMetricTagSlice) > 0 {
		obj.patternFlowIpv6VersionMetricTagSlice = []PatternFlowIpv6VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter) appendHolderSlice(item PatternFlowIpv6VersionMetricTag) PatternFlowIpv6VersionPatternFlowIpv6VersionMetricTagIter {
	obj.patternFlowIpv6VersionMetricTagSlice = append(obj.patternFlowIpv6VersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv6Version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv6Version.Value <= 15 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 15 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv6Version.Values <= 15 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv6VersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv6Version) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv6VersionChoice.VALUE)

	}

}
