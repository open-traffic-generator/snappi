package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2MessageType *****
type patternFlowGtpv2MessageType struct {
	validation
	obj              *otg.PatternFlowGtpv2MessageType
	marshaller       marshalPatternFlowGtpv2MessageType
	unMarshaller     unMarshalPatternFlowGtpv2MessageType
	incrementHolder  PatternFlowGtpv2MessageTypeCounter
	decrementHolder  PatternFlowGtpv2MessageTypeCounter
	metricTagsHolder PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
}

func NewPatternFlowGtpv2MessageType() PatternFlowGtpv2MessageType {
	obj := patternFlowGtpv2MessageType{obj: &otg.PatternFlowGtpv2MessageType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2MessageType) msg() *otg.PatternFlowGtpv2MessageType {
	return obj.obj
}

func (obj *patternFlowGtpv2MessageType) setMsg(msg *otg.PatternFlowGtpv2MessageType) PatternFlowGtpv2MessageType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2MessageType struct {
	obj *patternFlowGtpv2MessageType
}

type marshalPatternFlowGtpv2MessageType interface {
	// ToProto marshals PatternFlowGtpv2MessageType to protobuf object *otg.PatternFlowGtpv2MessageType
	ToProto() (*otg.PatternFlowGtpv2MessageType, error)
	// ToPbText marshals PatternFlowGtpv2MessageType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2MessageType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2MessageType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2MessageType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2MessageType struct {
	obj *patternFlowGtpv2MessageType
}

type unMarshalPatternFlowGtpv2MessageType interface {
	// FromProto unmarshals PatternFlowGtpv2MessageType from protobuf object *otg.PatternFlowGtpv2MessageType
	FromProto(msg *otg.PatternFlowGtpv2MessageType) (PatternFlowGtpv2MessageType, error)
	// FromPbText unmarshals PatternFlowGtpv2MessageType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2MessageType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2MessageType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2MessageType) Marshal() marshalPatternFlowGtpv2MessageType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2MessageType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2MessageType) Unmarshal() unMarshalPatternFlowGtpv2MessageType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2MessageType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2MessageType) ToProto() (*otg.PatternFlowGtpv2MessageType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2MessageType) FromProto(msg *otg.PatternFlowGtpv2MessageType) (PatternFlowGtpv2MessageType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2MessageType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageType) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageType) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2MessageType) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2MessageType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2MessageType) FromJson(value string) error {
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

func (obj *patternFlowGtpv2MessageType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2MessageType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2MessageType) Clone() (PatternFlowGtpv2MessageType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2MessageType()
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

func (obj *patternFlowGtpv2MessageType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2MessageType is an 8-bit field that indicates the type of GTP message. Different types of messages are defined in 3GPP TS 29.060 section 7.1
type PatternFlowGtpv2MessageType interface {
	Validation
	// msg marshals PatternFlowGtpv2MessageType to protobuf object *otg.PatternFlowGtpv2MessageType
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2MessageType
	// setMsg unmarshals PatternFlowGtpv2MessageType from protobuf object *otg.PatternFlowGtpv2MessageType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2MessageType) PatternFlowGtpv2MessageType
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2MessageType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2MessageType
	// validate validates PatternFlowGtpv2MessageType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2MessageType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2MessageTypeChoiceEnum, set in PatternFlowGtpv2MessageType
	Choice() PatternFlowGtpv2MessageTypeChoiceEnum
	// setChoice assigns PatternFlowGtpv2MessageTypeChoiceEnum provided by user to PatternFlowGtpv2MessageType
	setChoice(value PatternFlowGtpv2MessageTypeChoiceEnum) PatternFlowGtpv2MessageType
	// HasChoice checks if Choice has been set in PatternFlowGtpv2MessageType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2MessageType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2MessageType
	SetValue(value uint32) PatternFlowGtpv2MessageType
	// HasValue checks if Value has been set in PatternFlowGtpv2MessageType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2MessageType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2MessageType
	SetValues(value []uint32) PatternFlowGtpv2MessageType
	// Increment returns PatternFlowGtpv2MessageTypeCounter, set in PatternFlowGtpv2MessageType.
	// PatternFlowGtpv2MessageTypeCounter is integer counter pattern
	Increment() PatternFlowGtpv2MessageTypeCounter
	// SetIncrement assigns PatternFlowGtpv2MessageTypeCounter provided by user to PatternFlowGtpv2MessageType.
	// PatternFlowGtpv2MessageTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageType
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2MessageType
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2MessageTypeCounter, set in PatternFlowGtpv2MessageType.
	// PatternFlowGtpv2MessageTypeCounter is integer counter pattern
	Decrement() PatternFlowGtpv2MessageTypeCounter
	// SetDecrement assigns PatternFlowGtpv2MessageTypeCounter provided by user to PatternFlowGtpv2MessageType.
	// PatternFlowGtpv2MessageTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageType
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2MessageType
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIterIter, set in PatternFlowGtpv2MessageType
	MetricTags() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	setNil()
}

type PatternFlowGtpv2MessageTypeChoiceEnum string

// Enum of Choice on PatternFlowGtpv2MessageType
var PatternFlowGtpv2MessageTypeChoice = struct {
	VALUE     PatternFlowGtpv2MessageTypeChoiceEnum
	VALUES    PatternFlowGtpv2MessageTypeChoiceEnum
	INCREMENT PatternFlowGtpv2MessageTypeChoiceEnum
	DECREMENT PatternFlowGtpv2MessageTypeChoiceEnum
}{
	VALUE:     PatternFlowGtpv2MessageTypeChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2MessageTypeChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2MessageTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2MessageTypeChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2MessageType) Choice() PatternFlowGtpv2MessageTypeChoiceEnum {
	return PatternFlowGtpv2MessageTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2MessageType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2MessageType) setChoice(value PatternFlowGtpv2MessageTypeChoiceEnum) PatternFlowGtpv2MessageType {
	intValue, ok := otg.PatternFlowGtpv2MessageType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2MessageTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2MessageType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2MessageTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2MessageTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2MessageTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2MessageTypeCounter().msg()
	}

	if value == PatternFlowGtpv2MessageTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2MessageTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2MessageType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2MessageTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2MessageType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2MessageType object
func (obj *patternFlowGtpv2MessageType) SetValue(value uint32) PatternFlowGtpv2MessageType {
	obj.setChoice(PatternFlowGtpv2MessageTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2MessageType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2MessageType object
func (obj *patternFlowGtpv2MessageType) SetValues(value []uint32) PatternFlowGtpv2MessageType {
	obj.setChoice(PatternFlowGtpv2MessageTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2MessageTypeCounter
func (obj *patternFlowGtpv2MessageType) Increment() PatternFlowGtpv2MessageTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2MessageTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2MessageTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2MessageTypeCounter
func (obj *patternFlowGtpv2MessageType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2MessageTypeCounter value in the PatternFlowGtpv2MessageType object
func (obj *patternFlowGtpv2MessageType) SetIncrement(value PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageType {
	obj.setChoice(PatternFlowGtpv2MessageTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2MessageTypeCounter
func (obj *patternFlowGtpv2MessageType) Decrement() PatternFlowGtpv2MessageTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2MessageTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2MessageTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2MessageTypeCounter
func (obj *patternFlowGtpv2MessageType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2MessageTypeCounter value in the PatternFlowGtpv2MessageType object
func (obj *patternFlowGtpv2MessageType) SetDecrement(value PatternFlowGtpv2MessageTypeCounter) PatternFlowGtpv2MessageType {
	obj.setChoice(PatternFlowGtpv2MessageTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2MessageTypeMetricTag
func (obj *patternFlowGtpv2MessageType) MetricTags() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2MessageTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter struct {
	obj                                       *patternFlowGtpv2MessageType
	patternFlowGtpv2MessageTypeMetricTagSlice []PatternFlowGtpv2MessageTypeMetricTag
	fieldPtr                                  *[]*otg.PatternFlowGtpv2MessageTypeMetricTag
}

func newPatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter(ptr *[]*otg.PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	return &patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter interface {
	setMsg(*patternFlowGtpv2MessageType) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	Items() []PatternFlowGtpv2MessageTypeMetricTag
	Add() PatternFlowGtpv2MessageTypeMetricTag
	Append(items ...PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	Set(index int, newObj PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	Clear() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	clearHolderSlice() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter
}

func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) setMsg(msg *patternFlowGtpv2MessageType) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2MessageTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) Items() []PatternFlowGtpv2MessageTypeMetricTag {
	return obj.patternFlowGtpv2MessageTypeMetricTagSlice
}

func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) Add() PatternFlowGtpv2MessageTypeMetricTag {
	newObj := &otg.PatternFlowGtpv2MessageTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2MessageTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2MessageTypeMetricTagSlice = append(obj.patternFlowGtpv2MessageTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) Append(items ...PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2MessageTypeMetricTagSlice = append(obj.patternFlowGtpv2MessageTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) Set(index int, newObj PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2MessageTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) Clear() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2MessageTypeMetricTag{}
		obj.patternFlowGtpv2MessageTypeMetricTagSlice = []PatternFlowGtpv2MessageTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) clearHolderSlice() PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	if len(obj.patternFlowGtpv2MessageTypeMetricTagSlice) > 0 {
		obj.patternFlowGtpv2MessageTypeMetricTagSlice = []PatternFlowGtpv2MessageTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter) appendHolderSlice(item PatternFlowGtpv2MessageTypeMetricTag) PatternFlowGtpv2MessageTypePatternFlowGtpv2MessageTypeMetricTagIter {
	obj.patternFlowGtpv2MessageTypeMetricTagSlice = append(obj.patternFlowGtpv2MessageTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2MessageType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2MessageType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2MessageType.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2MessageTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2MessageType) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2MessageTypeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageTypeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2MessageTypeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageTypeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2MessageTypeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2MessageTypeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2MessageType")
			}
		} else {
			intVal := otg.PatternFlowGtpv2MessageType_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2MessageType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
