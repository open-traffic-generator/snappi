package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2Version *****
type patternFlowGtpv2Version struct {
	validation
	obj              *otg.PatternFlowGtpv2Version
	marshaller       marshalPatternFlowGtpv2Version
	unMarshaller     unMarshalPatternFlowGtpv2Version
	incrementHolder  PatternFlowGtpv2VersionCounter
	decrementHolder  PatternFlowGtpv2VersionCounter
	metricTagsHolder PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
}

func NewPatternFlowGtpv2Version() PatternFlowGtpv2Version {
	obj := patternFlowGtpv2Version{obj: &otg.PatternFlowGtpv2Version{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Version) msg() *otg.PatternFlowGtpv2Version {
	return obj.obj
}

func (obj *patternFlowGtpv2Version) setMsg(msg *otg.PatternFlowGtpv2Version) PatternFlowGtpv2Version {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Version struct {
	obj *patternFlowGtpv2Version
}

type marshalPatternFlowGtpv2Version interface {
	// ToProto marshals PatternFlowGtpv2Version to protobuf object *otg.PatternFlowGtpv2Version
	ToProto() (*otg.PatternFlowGtpv2Version, error)
	// ToPbText marshals PatternFlowGtpv2Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Version to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2Version struct {
	obj *patternFlowGtpv2Version
}

type unMarshalPatternFlowGtpv2Version interface {
	// FromProto unmarshals PatternFlowGtpv2Version from protobuf object *otg.PatternFlowGtpv2Version
	FromProto(msg *otg.PatternFlowGtpv2Version) (PatternFlowGtpv2Version, error)
	// FromPbText unmarshals PatternFlowGtpv2Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Version from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Version) Marshal() marshalPatternFlowGtpv2Version {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Version{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Version) Unmarshal() unMarshalPatternFlowGtpv2Version {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Version{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Version) ToProto() (*otg.PatternFlowGtpv2Version, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Version) FromProto(msg *otg.PatternFlowGtpv2Version) (PatternFlowGtpv2Version, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Version) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Version) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Version) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Version) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Version) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Version) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Version) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Version) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Version) Clone() (PatternFlowGtpv2Version, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Version()
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

func (obj *patternFlowGtpv2Version) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2Version is version number
type PatternFlowGtpv2Version interface {
	Validation
	// msg marshals PatternFlowGtpv2Version to protobuf object *otg.PatternFlowGtpv2Version
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Version
	// setMsg unmarshals PatternFlowGtpv2Version from protobuf object *otg.PatternFlowGtpv2Version
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Version) PatternFlowGtpv2Version
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Version
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Version
	// validate validates PatternFlowGtpv2Version
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2VersionChoiceEnum, set in PatternFlowGtpv2Version
	Choice() PatternFlowGtpv2VersionChoiceEnum
	// setChoice assigns PatternFlowGtpv2VersionChoiceEnum provided by user to PatternFlowGtpv2Version
	setChoice(value PatternFlowGtpv2VersionChoiceEnum) PatternFlowGtpv2Version
	// HasChoice checks if Choice has been set in PatternFlowGtpv2Version
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2Version.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2Version
	SetValue(value uint32) PatternFlowGtpv2Version
	// HasValue checks if Value has been set in PatternFlowGtpv2Version
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2Version.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2Version
	SetValues(value []uint32) PatternFlowGtpv2Version
	// Increment returns PatternFlowGtpv2VersionCounter, set in PatternFlowGtpv2Version.
	// PatternFlowGtpv2VersionCounter is integer counter pattern
	Increment() PatternFlowGtpv2VersionCounter
	// SetIncrement assigns PatternFlowGtpv2VersionCounter provided by user to PatternFlowGtpv2Version.
	// PatternFlowGtpv2VersionCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2VersionCounter) PatternFlowGtpv2Version
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2Version
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2VersionCounter, set in PatternFlowGtpv2Version.
	// PatternFlowGtpv2VersionCounter is integer counter pattern
	Decrement() PatternFlowGtpv2VersionCounter
	// SetDecrement assigns PatternFlowGtpv2VersionCounter provided by user to PatternFlowGtpv2Version.
	// PatternFlowGtpv2VersionCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2VersionCounter) PatternFlowGtpv2Version
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2Version
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIterIter, set in PatternFlowGtpv2Version
	MetricTags() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	setNil()
}

type PatternFlowGtpv2VersionChoiceEnum string

// Enum of Choice on PatternFlowGtpv2Version
var PatternFlowGtpv2VersionChoice = struct {
	VALUE     PatternFlowGtpv2VersionChoiceEnum
	VALUES    PatternFlowGtpv2VersionChoiceEnum
	INCREMENT PatternFlowGtpv2VersionChoiceEnum
	DECREMENT PatternFlowGtpv2VersionChoiceEnum
}{
	VALUE:     PatternFlowGtpv2VersionChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2VersionChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2VersionChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2VersionChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2Version) Choice() PatternFlowGtpv2VersionChoiceEnum {
	return PatternFlowGtpv2VersionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2Version) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2Version) setChoice(value PatternFlowGtpv2VersionChoiceEnum) PatternFlowGtpv2Version {
	intValue, ok := otg.PatternFlowGtpv2Version_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2VersionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2Version_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2VersionChoice.VALUE {
		defaultValue := uint32(2)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2VersionChoice.VALUES {
		defaultValue := []uint32{2}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2VersionChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2VersionCounter().msg()
	}

	if value == PatternFlowGtpv2VersionChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2VersionCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Version) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2VersionChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Version) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2Version object
func (obj *patternFlowGtpv2Version) SetValue(value uint32) PatternFlowGtpv2Version {
	obj.setChoice(PatternFlowGtpv2VersionChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2Version) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{2})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2Version object
func (obj *patternFlowGtpv2Version) SetValues(value []uint32) PatternFlowGtpv2Version {
	obj.setChoice(PatternFlowGtpv2VersionChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2VersionCounter
func (obj *patternFlowGtpv2Version) Increment() PatternFlowGtpv2VersionCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2VersionChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2VersionCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2VersionCounter
func (obj *patternFlowGtpv2Version) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2VersionCounter value in the PatternFlowGtpv2Version object
func (obj *patternFlowGtpv2Version) SetIncrement(value PatternFlowGtpv2VersionCounter) PatternFlowGtpv2Version {
	obj.setChoice(PatternFlowGtpv2VersionChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2VersionCounter
func (obj *patternFlowGtpv2Version) Decrement() PatternFlowGtpv2VersionCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2VersionChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2VersionCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2VersionCounter
func (obj *patternFlowGtpv2Version) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2VersionCounter value in the PatternFlowGtpv2Version object
func (obj *patternFlowGtpv2Version) SetDecrement(value PatternFlowGtpv2VersionCounter) PatternFlowGtpv2Version {
	obj.setChoice(PatternFlowGtpv2VersionChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2VersionMetricTag
func (obj *patternFlowGtpv2Version) MetricTags() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2VersionMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter struct {
	obj                                   *patternFlowGtpv2Version
	patternFlowGtpv2VersionMetricTagSlice []PatternFlowGtpv2VersionMetricTag
	fieldPtr                              *[]*otg.PatternFlowGtpv2VersionMetricTag
}

func newPatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter(ptr *[]*otg.PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	return &patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter interface {
	setMsg(*patternFlowGtpv2Version) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	Items() []PatternFlowGtpv2VersionMetricTag
	Add() PatternFlowGtpv2VersionMetricTag
	Append(items ...PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	Set(index int, newObj PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	Clear() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	clearHolderSlice() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter
}

func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) setMsg(msg *patternFlowGtpv2Version) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2VersionMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) Items() []PatternFlowGtpv2VersionMetricTag {
	return obj.patternFlowGtpv2VersionMetricTagSlice
}

func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) Add() PatternFlowGtpv2VersionMetricTag {
	newObj := &otg.PatternFlowGtpv2VersionMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2VersionMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2VersionMetricTagSlice = append(obj.patternFlowGtpv2VersionMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) Append(items ...PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2VersionMetricTagSlice = append(obj.patternFlowGtpv2VersionMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) Set(index int, newObj PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2VersionMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) Clear() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2VersionMetricTag{}
		obj.patternFlowGtpv2VersionMetricTagSlice = []PatternFlowGtpv2VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) clearHolderSlice() PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	if len(obj.patternFlowGtpv2VersionMetricTagSlice) > 0 {
		obj.patternFlowGtpv2VersionMetricTagSlice = []PatternFlowGtpv2VersionMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter) appendHolderSlice(item PatternFlowGtpv2VersionMetricTag) PatternFlowGtpv2VersionPatternFlowGtpv2VersionMetricTagIter {
	obj.patternFlowGtpv2VersionMetricTagSlice = append(obj.patternFlowGtpv2VersionMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2Version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2Version.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2Version.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2VersionMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2Version) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv2VersionChoice.VALUE)

	}

}
