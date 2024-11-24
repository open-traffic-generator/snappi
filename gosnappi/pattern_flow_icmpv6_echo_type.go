package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoType *****
type patternFlowIcmpv6EchoType struct {
	validation
	obj              *otg.PatternFlowIcmpv6EchoType
	marshaller       marshalPatternFlowIcmpv6EchoType
	unMarshaller     unMarshalPatternFlowIcmpv6EchoType
	incrementHolder  PatternFlowIcmpv6EchoTypeCounter
	decrementHolder  PatternFlowIcmpv6EchoTypeCounter
	metricTagsHolder PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
}

func NewPatternFlowIcmpv6EchoType() PatternFlowIcmpv6EchoType {
	obj := patternFlowIcmpv6EchoType{obj: &otg.PatternFlowIcmpv6EchoType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoType) msg() *otg.PatternFlowIcmpv6EchoType {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoType) setMsg(msg *otg.PatternFlowIcmpv6EchoType) PatternFlowIcmpv6EchoType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoType struct {
	obj *patternFlowIcmpv6EchoType
}

type marshalPatternFlowIcmpv6EchoType interface {
	// ToProto marshals PatternFlowIcmpv6EchoType to protobuf object *otg.PatternFlowIcmpv6EchoType
	ToProto() (*otg.PatternFlowIcmpv6EchoType, error)
	// ToPbText marshals PatternFlowIcmpv6EchoType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpv6EchoType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoType struct {
	obj *patternFlowIcmpv6EchoType
}

type unMarshalPatternFlowIcmpv6EchoType interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoType from protobuf object *otg.PatternFlowIcmpv6EchoType
	FromProto(msg *otg.PatternFlowIcmpv6EchoType) (PatternFlowIcmpv6EchoType, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoType) Marshal() marshalPatternFlowIcmpv6EchoType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoType) Unmarshal() unMarshalPatternFlowIcmpv6EchoType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoType) ToProto() (*otg.PatternFlowIcmpv6EchoType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoType) FromProto(msg *otg.PatternFlowIcmpv6EchoType) (PatternFlowIcmpv6EchoType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoType) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoType) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoType) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpv6EchoType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoType) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoType) Clone() (PatternFlowIcmpv6EchoType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoType()
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

func (obj *patternFlowIcmpv6EchoType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6EchoType is iCMPv6 echo type
type PatternFlowIcmpv6EchoType interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoType to protobuf object *otg.PatternFlowIcmpv6EchoType
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoType
	// setMsg unmarshals PatternFlowIcmpv6EchoType from protobuf object *otg.PatternFlowIcmpv6EchoType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoType) PatternFlowIcmpv6EchoType
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoType
	// validate validates PatternFlowIcmpv6EchoType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6EchoTypeChoiceEnum, set in PatternFlowIcmpv6EchoType
	Choice() PatternFlowIcmpv6EchoTypeChoiceEnum
	// setChoice assigns PatternFlowIcmpv6EchoTypeChoiceEnum provided by user to PatternFlowIcmpv6EchoType
	setChoice(value PatternFlowIcmpv6EchoTypeChoiceEnum) PatternFlowIcmpv6EchoType
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6EchoType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6EchoType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6EchoType
	SetValue(value uint32) PatternFlowIcmpv6EchoType
	// HasValue checks if Value has been set in PatternFlowIcmpv6EchoType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6EchoType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6EchoType
	SetValues(value []uint32) PatternFlowIcmpv6EchoType
	// Increment returns PatternFlowIcmpv6EchoTypeCounter, set in PatternFlowIcmpv6EchoType.
	// PatternFlowIcmpv6EchoTypeCounter is integer counter pattern
	Increment() PatternFlowIcmpv6EchoTypeCounter
	// SetIncrement assigns PatternFlowIcmpv6EchoTypeCounter provided by user to PatternFlowIcmpv6EchoType.
	// PatternFlowIcmpv6EchoTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoType
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6EchoType
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6EchoTypeCounter, set in PatternFlowIcmpv6EchoType.
	// PatternFlowIcmpv6EchoTypeCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6EchoTypeCounter
	// SetDecrement assigns PatternFlowIcmpv6EchoTypeCounter provided by user to PatternFlowIcmpv6EchoType.
	// PatternFlowIcmpv6EchoTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoType
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6EchoType
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIterIter, set in PatternFlowIcmpv6EchoType
	MetricTags() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	setNil()
}

type PatternFlowIcmpv6EchoTypeChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6EchoType
var PatternFlowIcmpv6EchoTypeChoice = struct {
	VALUE     PatternFlowIcmpv6EchoTypeChoiceEnum
	VALUES    PatternFlowIcmpv6EchoTypeChoiceEnum
	INCREMENT PatternFlowIcmpv6EchoTypeChoiceEnum
	DECREMENT PatternFlowIcmpv6EchoTypeChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6EchoTypeChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6EchoTypeChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6EchoTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6EchoTypeChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6EchoType) Choice() PatternFlowIcmpv6EchoTypeChoiceEnum {
	return PatternFlowIcmpv6EchoTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6EchoType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6EchoType) setChoice(value PatternFlowIcmpv6EchoTypeChoiceEnum) PatternFlowIcmpv6EchoType {
	intValue, ok := otg.PatternFlowIcmpv6EchoType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6EchoTypeChoice.VALUE {
		defaultValue := uint32(128)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6EchoTypeChoice.VALUES {
		defaultValue := []uint32{128}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6EchoTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6EchoTypeCounter().msg()
	}

	if value == PatternFlowIcmpv6EchoTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6EchoTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6EchoType object
func (obj *patternFlowIcmpv6EchoType) SetValue(value uint32) PatternFlowIcmpv6EchoType {
	obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6EchoType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{128})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6EchoType object
func (obj *patternFlowIcmpv6EchoType) SetValues(value []uint32) PatternFlowIcmpv6EchoType {
	obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoTypeCounter
func (obj *patternFlowIcmpv6EchoType) Increment() PatternFlowIcmpv6EchoTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6EchoTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoTypeCounter
func (obj *patternFlowIcmpv6EchoType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6EchoTypeCounter value in the PatternFlowIcmpv6EchoType object
func (obj *patternFlowIcmpv6EchoType) SetIncrement(value PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoType {
	obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoTypeCounter
func (obj *patternFlowIcmpv6EchoType) Decrement() PatternFlowIcmpv6EchoTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6EchoTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoTypeCounter
func (obj *patternFlowIcmpv6EchoType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6EchoTypeCounter value in the PatternFlowIcmpv6EchoType object
func (obj *patternFlowIcmpv6EchoType) SetDecrement(value PatternFlowIcmpv6EchoTypeCounter) PatternFlowIcmpv6EchoType {
	obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpv6EchoTypeMetricTag
func (obj *patternFlowIcmpv6EchoType) MetricTags() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpv6EchoTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter struct {
	obj                                     *patternFlowIcmpv6EchoType
	patternFlowIcmpv6EchoTypeMetricTagSlice []PatternFlowIcmpv6EchoTypeMetricTag
	fieldPtr                                *[]*otg.PatternFlowIcmpv6EchoTypeMetricTag
}

func newPatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter(ptr *[]*otg.PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	return &patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter interface {
	setMsg(*patternFlowIcmpv6EchoType) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	Items() []PatternFlowIcmpv6EchoTypeMetricTag
	Add() PatternFlowIcmpv6EchoTypeMetricTag
	Append(items ...PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	Set(index int, newObj PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	Clear() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	clearHolderSlice() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
	appendHolderSlice(item PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter
}

func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) setMsg(msg *patternFlowIcmpv6EchoType) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpv6EchoTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) Items() []PatternFlowIcmpv6EchoTypeMetricTag {
	return obj.patternFlowIcmpv6EchoTypeMetricTagSlice
}

func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) Add() PatternFlowIcmpv6EchoTypeMetricTag {
	newObj := &otg.PatternFlowIcmpv6EchoTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpv6EchoTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpv6EchoTypeMetricTagSlice = append(obj.patternFlowIcmpv6EchoTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) Append(items ...PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpv6EchoTypeMetricTagSlice = append(obj.patternFlowIcmpv6EchoTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) Set(index int, newObj PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpv6EchoTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) Clear() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpv6EchoTypeMetricTag{}
		obj.patternFlowIcmpv6EchoTypeMetricTagSlice = []PatternFlowIcmpv6EchoTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) clearHolderSlice() PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	if len(obj.patternFlowIcmpv6EchoTypeMetricTagSlice) > 0 {
		obj.patternFlowIcmpv6EchoTypeMetricTagSlice = []PatternFlowIcmpv6EchoTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter) appendHolderSlice(item PatternFlowIcmpv6EchoTypeMetricTag) PatternFlowIcmpv6EchoTypePatternFlowIcmpv6EchoTypeMetricTagIter {
	obj.patternFlowIcmpv6EchoTypeMetricTagSlice = append(obj.patternFlowIcmpv6EchoTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpv6EchoType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6EchoType.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpv6EchoTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpv6EchoType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6EchoTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6EchoTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6EchoType")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6EchoType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6EchoType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
