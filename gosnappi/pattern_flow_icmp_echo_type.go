package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoType *****
type patternFlowIcmpEchoType struct {
	validation
	obj              *otg.PatternFlowIcmpEchoType
	marshaller       marshalPatternFlowIcmpEchoType
	unMarshaller     unMarshalPatternFlowIcmpEchoType
	incrementHolder  PatternFlowIcmpEchoTypeCounter
	decrementHolder  PatternFlowIcmpEchoTypeCounter
	metricTagsHolder PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
}

func NewPatternFlowIcmpEchoType() PatternFlowIcmpEchoType {
	obj := patternFlowIcmpEchoType{obj: &otg.PatternFlowIcmpEchoType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoType) msg() *otg.PatternFlowIcmpEchoType {
	return obj.obj
}

func (obj *patternFlowIcmpEchoType) setMsg(msg *otg.PatternFlowIcmpEchoType) PatternFlowIcmpEchoType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoType struct {
	obj *patternFlowIcmpEchoType
}

type marshalPatternFlowIcmpEchoType interface {
	// ToProto marshals PatternFlowIcmpEchoType to protobuf object *otg.PatternFlowIcmpEchoType
	ToProto() (*otg.PatternFlowIcmpEchoType, error)
	// ToPbText marshals PatternFlowIcmpEchoType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoType struct {
	obj *patternFlowIcmpEchoType
}

type unMarshalPatternFlowIcmpEchoType interface {
	// FromProto unmarshals PatternFlowIcmpEchoType from protobuf object *otg.PatternFlowIcmpEchoType
	FromProto(msg *otg.PatternFlowIcmpEchoType) (PatternFlowIcmpEchoType, error)
	// FromPbText unmarshals PatternFlowIcmpEchoType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoType) Marshal() marshalPatternFlowIcmpEchoType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoType) Unmarshal() unMarshalPatternFlowIcmpEchoType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoType) ToProto() (*otg.PatternFlowIcmpEchoType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoType) FromProto(msg *otg.PatternFlowIcmpEchoType) (PatternFlowIcmpEchoType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoType) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoType) Clone() (PatternFlowIcmpEchoType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoType()
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

func (obj *patternFlowIcmpEchoType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpEchoType is the type of ICMP echo packet
type PatternFlowIcmpEchoType interface {
	Validation
	// msg marshals PatternFlowIcmpEchoType to protobuf object *otg.PatternFlowIcmpEchoType
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoType
	// setMsg unmarshals PatternFlowIcmpEchoType from protobuf object *otg.PatternFlowIcmpEchoType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoType) PatternFlowIcmpEchoType
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoType
	// validate validates PatternFlowIcmpEchoType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpEchoTypeChoiceEnum, set in PatternFlowIcmpEchoType
	Choice() PatternFlowIcmpEchoTypeChoiceEnum
	// setChoice assigns PatternFlowIcmpEchoTypeChoiceEnum provided by user to PatternFlowIcmpEchoType
	setChoice(value PatternFlowIcmpEchoTypeChoiceEnum) PatternFlowIcmpEchoType
	// HasChoice checks if Choice has been set in PatternFlowIcmpEchoType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpEchoType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpEchoType
	SetValue(value uint32) PatternFlowIcmpEchoType
	// HasValue checks if Value has been set in PatternFlowIcmpEchoType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpEchoType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpEchoType
	SetValues(value []uint32) PatternFlowIcmpEchoType
	// Increment returns PatternFlowIcmpEchoTypeCounter, set in PatternFlowIcmpEchoType.
	// PatternFlowIcmpEchoTypeCounter is integer counter pattern
	Increment() PatternFlowIcmpEchoTypeCounter
	// SetIncrement assigns PatternFlowIcmpEchoTypeCounter provided by user to PatternFlowIcmpEchoType.
	// PatternFlowIcmpEchoTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoType
	// HasIncrement checks if Increment has been set in PatternFlowIcmpEchoType
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpEchoTypeCounter, set in PatternFlowIcmpEchoType.
	// PatternFlowIcmpEchoTypeCounter is integer counter pattern
	Decrement() PatternFlowIcmpEchoTypeCounter
	// SetDecrement assigns PatternFlowIcmpEchoTypeCounter provided by user to PatternFlowIcmpEchoType.
	// PatternFlowIcmpEchoTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoType
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpEchoType
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIterIter, set in PatternFlowIcmpEchoType
	MetricTags() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	setNil()
}

type PatternFlowIcmpEchoTypeChoiceEnum string

// Enum of Choice on PatternFlowIcmpEchoType
var PatternFlowIcmpEchoTypeChoice = struct {
	VALUE     PatternFlowIcmpEchoTypeChoiceEnum
	VALUES    PatternFlowIcmpEchoTypeChoiceEnum
	INCREMENT PatternFlowIcmpEchoTypeChoiceEnum
	DECREMENT PatternFlowIcmpEchoTypeChoiceEnum
}{
	VALUE:     PatternFlowIcmpEchoTypeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpEchoTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpEchoTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpEchoTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpEchoType) Choice() PatternFlowIcmpEchoTypeChoiceEnum {
	return PatternFlowIcmpEchoTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpEchoType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpEchoType) setChoice(value PatternFlowIcmpEchoTypeChoiceEnum) PatternFlowIcmpEchoType {
	intValue, ok := otg.PatternFlowIcmpEchoType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpEchoTypeChoice.VALUE {
		defaultValue := uint32(8)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpEchoTypeChoice.VALUES {
		defaultValue := []uint32{8}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpEchoTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpEchoTypeCounter().msg()
	}

	if value == PatternFlowIcmpEchoTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpEchoTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpEchoTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpEchoType object
func (obj *patternFlowIcmpEchoType) SetValue(value uint32) PatternFlowIcmpEchoType {
	obj.setChoice(PatternFlowIcmpEchoTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpEchoType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{8})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpEchoType object
func (obj *patternFlowIcmpEchoType) SetValues(value []uint32) PatternFlowIcmpEchoType {
	obj.setChoice(PatternFlowIcmpEchoTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoTypeCounter
func (obj *patternFlowIcmpEchoType) Increment() PatternFlowIcmpEchoTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpEchoTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpEchoTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoTypeCounter
func (obj *patternFlowIcmpEchoType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpEchoTypeCounter value in the PatternFlowIcmpEchoType object
func (obj *patternFlowIcmpEchoType) SetIncrement(value PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoType {
	obj.setChoice(PatternFlowIcmpEchoTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoTypeCounter
func (obj *patternFlowIcmpEchoType) Decrement() PatternFlowIcmpEchoTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpEchoTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpEchoTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoTypeCounter
func (obj *patternFlowIcmpEchoType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpEchoTypeCounter value in the PatternFlowIcmpEchoType object
func (obj *patternFlowIcmpEchoType) SetDecrement(value PatternFlowIcmpEchoTypeCounter) PatternFlowIcmpEchoType {
	obj.setChoice(PatternFlowIcmpEchoTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpEchoTypeMetricTag
func (obj *patternFlowIcmpEchoType) MetricTags() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpEchoTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter struct {
	obj                                   *patternFlowIcmpEchoType
	patternFlowIcmpEchoTypeMetricTagSlice []PatternFlowIcmpEchoTypeMetricTag
	fieldPtr                              *[]*otg.PatternFlowIcmpEchoTypeMetricTag
}

func newPatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter(ptr *[]*otg.PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	return &patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter interface {
	setMsg(*patternFlowIcmpEchoType) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	Items() []PatternFlowIcmpEchoTypeMetricTag
	Add() PatternFlowIcmpEchoTypeMetricTag
	Append(items ...PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	Set(index int, newObj PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	Clear() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	clearHolderSlice() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
	appendHolderSlice(item PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter
}

func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) setMsg(msg *patternFlowIcmpEchoType) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpEchoTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) Items() []PatternFlowIcmpEchoTypeMetricTag {
	return obj.patternFlowIcmpEchoTypeMetricTagSlice
}

func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) Add() PatternFlowIcmpEchoTypeMetricTag {
	newObj := &otg.PatternFlowIcmpEchoTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpEchoTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpEchoTypeMetricTagSlice = append(obj.patternFlowIcmpEchoTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) Append(items ...PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpEchoTypeMetricTagSlice = append(obj.patternFlowIcmpEchoTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) Set(index int, newObj PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpEchoTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) Clear() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpEchoTypeMetricTag{}
		obj.patternFlowIcmpEchoTypeMetricTagSlice = []PatternFlowIcmpEchoTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) clearHolderSlice() PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	if len(obj.patternFlowIcmpEchoTypeMetricTagSlice) > 0 {
		obj.patternFlowIcmpEchoTypeMetricTagSlice = []PatternFlowIcmpEchoTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter) appendHolderSlice(item PatternFlowIcmpEchoTypeMetricTag) PatternFlowIcmpEchoTypePatternFlowIcmpEchoTypeMetricTagIter {
	obj.patternFlowIcmpEchoTypeMetricTagSlice = append(obj.patternFlowIcmpEchoTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpEchoType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpEchoType.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpEchoTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpEchoType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIcmpEchoTypeChoice.VALUE)

	}

}
