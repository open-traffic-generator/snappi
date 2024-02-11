package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowArpHardwareType *****
type patternFlowArpHardwareType struct {
	validation
	obj              *otg.PatternFlowArpHardwareType
	marshaller       marshalPatternFlowArpHardwareType
	unMarshaller     unMarshalPatternFlowArpHardwareType
	incrementHolder  PatternFlowArpHardwareTypeCounter
	decrementHolder  PatternFlowArpHardwareTypeCounter
	metricTagsHolder PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
}

func NewPatternFlowArpHardwareType() PatternFlowArpHardwareType {
	obj := patternFlowArpHardwareType{obj: &otg.PatternFlowArpHardwareType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowArpHardwareType) msg() *otg.PatternFlowArpHardwareType {
	return obj.obj
}

func (obj *patternFlowArpHardwareType) setMsg(msg *otg.PatternFlowArpHardwareType) PatternFlowArpHardwareType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowArpHardwareType struct {
	obj *patternFlowArpHardwareType
}

type marshalPatternFlowArpHardwareType interface {
	// ToProto marshals PatternFlowArpHardwareType to protobuf object *otg.PatternFlowArpHardwareType
	ToProto() (*otg.PatternFlowArpHardwareType, error)
	// ToPbText marshals PatternFlowArpHardwareType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowArpHardwareType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowArpHardwareType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowArpHardwareType struct {
	obj *patternFlowArpHardwareType
}

type unMarshalPatternFlowArpHardwareType interface {
	// FromProto unmarshals PatternFlowArpHardwareType from protobuf object *otg.PatternFlowArpHardwareType
	FromProto(msg *otg.PatternFlowArpHardwareType) (PatternFlowArpHardwareType, error)
	// FromPbText unmarshals PatternFlowArpHardwareType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowArpHardwareType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowArpHardwareType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowArpHardwareType) Marshal() marshalPatternFlowArpHardwareType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowArpHardwareType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowArpHardwareType) Unmarshal() unMarshalPatternFlowArpHardwareType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowArpHardwareType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowArpHardwareType) ToProto() (*otg.PatternFlowArpHardwareType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowArpHardwareType) FromProto(msg *otg.PatternFlowArpHardwareType) (PatternFlowArpHardwareType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowArpHardwareType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareType) FromPbText(value string) error {
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

func (m *marshalpatternFlowArpHardwareType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareType) FromYaml(value string) error {
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

func (m *marshalpatternFlowArpHardwareType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowArpHardwareType) FromJson(value string) error {
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

func (obj *patternFlowArpHardwareType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowArpHardwareType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowArpHardwareType) Clone() (PatternFlowArpHardwareType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowArpHardwareType()
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

func (obj *patternFlowArpHardwareType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowArpHardwareType is network link protocol type
type PatternFlowArpHardwareType interface {
	Validation
	// msg marshals PatternFlowArpHardwareType to protobuf object *otg.PatternFlowArpHardwareType
	// and doesn't set defaults
	msg() *otg.PatternFlowArpHardwareType
	// setMsg unmarshals PatternFlowArpHardwareType from protobuf object *otg.PatternFlowArpHardwareType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowArpHardwareType) PatternFlowArpHardwareType
	// provides marshal interface
	Marshal() marshalPatternFlowArpHardwareType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowArpHardwareType
	// validate validates PatternFlowArpHardwareType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowArpHardwareType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowArpHardwareTypeChoiceEnum, set in PatternFlowArpHardwareType
	Choice() PatternFlowArpHardwareTypeChoiceEnum
	// setChoice assigns PatternFlowArpHardwareTypeChoiceEnum provided by user to PatternFlowArpHardwareType
	setChoice(value PatternFlowArpHardwareTypeChoiceEnum) PatternFlowArpHardwareType
	// HasChoice checks if Choice has been set in PatternFlowArpHardwareType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowArpHardwareType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowArpHardwareType
	SetValue(value uint32) PatternFlowArpHardwareType
	// HasValue checks if Value has been set in PatternFlowArpHardwareType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowArpHardwareType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowArpHardwareType
	SetValues(value []uint32) PatternFlowArpHardwareType
	// Increment returns PatternFlowArpHardwareTypeCounter, set in PatternFlowArpHardwareType.
	// PatternFlowArpHardwareTypeCounter is integer counter pattern
	Increment() PatternFlowArpHardwareTypeCounter
	// SetIncrement assigns PatternFlowArpHardwareTypeCounter provided by user to PatternFlowArpHardwareType.
	// PatternFlowArpHardwareTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareType
	// HasIncrement checks if Increment has been set in PatternFlowArpHardwareType
	HasIncrement() bool
	// Decrement returns PatternFlowArpHardwareTypeCounter, set in PatternFlowArpHardwareType.
	// PatternFlowArpHardwareTypeCounter is integer counter pattern
	Decrement() PatternFlowArpHardwareTypeCounter
	// SetDecrement assigns PatternFlowArpHardwareTypeCounter provided by user to PatternFlowArpHardwareType.
	// PatternFlowArpHardwareTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareType
	// HasDecrement checks if Decrement has been set in PatternFlowArpHardwareType
	HasDecrement() bool
	// MetricTags returns PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIterIter, set in PatternFlowArpHardwareType
	MetricTags() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	setNil()
}

type PatternFlowArpHardwareTypeChoiceEnum string

// Enum of Choice on PatternFlowArpHardwareType
var PatternFlowArpHardwareTypeChoice = struct {
	VALUE     PatternFlowArpHardwareTypeChoiceEnum
	VALUES    PatternFlowArpHardwareTypeChoiceEnum
	INCREMENT PatternFlowArpHardwareTypeChoiceEnum
	DECREMENT PatternFlowArpHardwareTypeChoiceEnum
}{
	VALUE:     PatternFlowArpHardwareTypeChoiceEnum("value"),
	VALUES:    PatternFlowArpHardwareTypeChoiceEnum("values"),
	INCREMENT: PatternFlowArpHardwareTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowArpHardwareTypeChoiceEnum("decrement"),
}

func (obj *patternFlowArpHardwareType) Choice() PatternFlowArpHardwareTypeChoiceEnum {
	return PatternFlowArpHardwareTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowArpHardwareType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowArpHardwareType) setChoice(value PatternFlowArpHardwareTypeChoiceEnum) PatternFlowArpHardwareType {
	intValue, ok := otg.PatternFlowArpHardwareType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowArpHardwareTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowArpHardwareType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowArpHardwareTypeChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowArpHardwareTypeChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowArpHardwareTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowArpHardwareTypeCounter().msg()
	}

	if value == PatternFlowArpHardwareTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowArpHardwareTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpHardwareType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowArpHardwareTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowArpHardwareType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowArpHardwareType object
func (obj *patternFlowArpHardwareType) SetValue(value uint32) PatternFlowArpHardwareType {
	obj.setChoice(PatternFlowArpHardwareTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowArpHardwareType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowArpHardwareType object
func (obj *patternFlowArpHardwareType) SetValues(value []uint32) PatternFlowArpHardwareType {
	obj.setChoice(PatternFlowArpHardwareTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowArpHardwareTypeCounter
func (obj *patternFlowArpHardwareType) Increment() PatternFlowArpHardwareTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowArpHardwareTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowArpHardwareTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowArpHardwareTypeCounter
func (obj *patternFlowArpHardwareType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowArpHardwareTypeCounter value in the PatternFlowArpHardwareType object
func (obj *patternFlowArpHardwareType) SetIncrement(value PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareType {
	obj.setChoice(PatternFlowArpHardwareTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowArpHardwareTypeCounter
func (obj *patternFlowArpHardwareType) Decrement() PatternFlowArpHardwareTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowArpHardwareTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowArpHardwareTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowArpHardwareTypeCounter
func (obj *patternFlowArpHardwareType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowArpHardwareTypeCounter value in the PatternFlowArpHardwareType object
func (obj *patternFlowArpHardwareType) SetDecrement(value PatternFlowArpHardwareTypeCounter) PatternFlowArpHardwareType {
	obj.setChoice(PatternFlowArpHardwareTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowArpHardwareTypeMetricTag
func (obj *patternFlowArpHardwareType) MetricTags() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowArpHardwareTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter struct {
	obj                                      *patternFlowArpHardwareType
	patternFlowArpHardwareTypeMetricTagSlice []PatternFlowArpHardwareTypeMetricTag
	fieldPtr                                 *[]*otg.PatternFlowArpHardwareTypeMetricTag
}

func newPatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter(ptr *[]*otg.PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	return &patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter interface {
	setMsg(*patternFlowArpHardwareType) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	Items() []PatternFlowArpHardwareTypeMetricTag
	Add() PatternFlowArpHardwareTypeMetricTag
	Append(items ...PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	Set(index int, newObj PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	Clear() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	clearHolderSlice() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
	appendHolderSlice(item PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter
}

func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) setMsg(msg *patternFlowArpHardwareType) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowArpHardwareTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) Items() []PatternFlowArpHardwareTypeMetricTag {
	return obj.patternFlowArpHardwareTypeMetricTagSlice
}

func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) Add() PatternFlowArpHardwareTypeMetricTag {
	newObj := &otg.PatternFlowArpHardwareTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowArpHardwareTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowArpHardwareTypeMetricTagSlice = append(obj.patternFlowArpHardwareTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) Append(items ...PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowArpHardwareTypeMetricTagSlice = append(obj.patternFlowArpHardwareTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) Set(index int, newObj PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowArpHardwareTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) Clear() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowArpHardwareTypeMetricTag{}
		obj.patternFlowArpHardwareTypeMetricTagSlice = []PatternFlowArpHardwareTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) clearHolderSlice() PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	if len(obj.patternFlowArpHardwareTypeMetricTagSlice) > 0 {
		obj.patternFlowArpHardwareTypeMetricTagSlice = []PatternFlowArpHardwareTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter) appendHolderSlice(item PatternFlowArpHardwareTypeMetricTag) PatternFlowArpHardwareTypePatternFlowArpHardwareTypeMetricTagIter {
	obj.patternFlowArpHardwareTypeMetricTagSlice = append(obj.patternFlowArpHardwareTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowArpHardwareType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowArpHardwareType.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowArpHardwareType.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowArpHardwareTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowArpHardwareType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowArpHardwareTypeChoice.VALUE)

	}

}
