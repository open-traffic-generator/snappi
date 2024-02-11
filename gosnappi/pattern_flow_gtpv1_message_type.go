package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1MessageType *****
type patternFlowGtpv1MessageType struct {
	validation
	obj              *otg.PatternFlowGtpv1MessageType
	marshaller       marshalPatternFlowGtpv1MessageType
	unMarshaller     unMarshalPatternFlowGtpv1MessageType
	incrementHolder  PatternFlowGtpv1MessageTypeCounter
	decrementHolder  PatternFlowGtpv1MessageTypeCounter
	metricTagsHolder PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
}

func NewPatternFlowGtpv1MessageType() PatternFlowGtpv1MessageType {
	obj := patternFlowGtpv1MessageType{obj: &otg.PatternFlowGtpv1MessageType{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1MessageType) msg() *otg.PatternFlowGtpv1MessageType {
	return obj.obj
}

func (obj *patternFlowGtpv1MessageType) setMsg(msg *otg.PatternFlowGtpv1MessageType) PatternFlowGtpv1MessageType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1MessageType struct {
	obj *patternFlowGtpv1MessageType
}

type marshalPatternFlowGtpv1MessageType interface {
	// ToProto marshals PatternFlowGtpv1MessageType to protobuf object *otg.PatternFlowGtpv1MessageType
	ToProto() (*otg.PatternFlowGtpv1MessageType, error)
	// ToPbText marshals PatternFlowGtpv1MessageType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1MessageType to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1MessageType to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1MessageType struct {
	obj *patternFlowGtpv1MessageType
}

type unMarshalPatternFlowGtpv1MessageType interface {
	// FromProto unmarshals PatternFlowGtpv1MessageType from protobuf object *otg.PatternFlowGtpv1MessageType
	FromProto(msg *otg.PatternFlowGtpv1MessageType) (PatternFlowGtpv1MessageType, error)
	// FromPbText unmarshals PatternFlowGtpv1MessageType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1MessageType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1MessageType from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1MessageType) Marshal() marshalPatternFlowGtpv1MessageType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1MessageType{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1MessageType) Unmarshal() unMarshalPatternFlowGtpv1MessageType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1MessageType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1MessageType) ToProto() (*otg.PatternFlowGtpv1MessageType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1MessageType) FromProto(msg *otg.PatternFlowGtpv1MessageType) (PatternFlowGtpv1MessageType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1MessageType) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageType) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageType) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageType) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1MessageType) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1MessageType) FromJson(value string) error {
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

func (obj *patternFlowGtpv1MessageType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1MessageType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1MessageType) Clone() (PatternFlowGtpv1MessageType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1MessageType()
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

func (obj *patternFlowGtpv1MessageType) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1MessageType is the type of GTP message Different types of messages are defined in 3GPP TS 29.060 section 7.1
type PatternFlowGtpv1MessageType interface {
	Validation
	// msg marshals PatternFlowGtpv1MessageType to protobuf object *otg.PatternFlowGtpv1MessageType
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1MessageType
	// setMsg unmarshals PatternFlowGtpv1MessageType from protobuf object *otg.PatternFlowGtpv1MessageType
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1MessageType) PatternFlowGtpv1MessageType
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1MessageType
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1MessageType
	// validate validates PatternFlowGtpv1MessageType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1MessageType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1MessageTypeChoiceEnum, set in PatternFlowGtpv1MessageType
	Choice() PatternFlowGtpv1MessageTypeChoiceEnum
	// setChoice assigns PatternFlowGtpv1MessageTypeChoiceEnum provided by user to PatternFlowGtpv1MessageType
	setChoice(value PatternFlowGtpv1MessageTypeChoiceEnum) PatternFlowGtpv1MessageType
	// HasChoice checks if Choice has been set in PatternFlowGtpv1MessageType
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1MessageType.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1MessageType
	SetValue(value uint32) PatternFlowGtpv1MessageType
	// HasValue checks if Value has been set in PatternFlowGtpv1MessageType
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1MessageType.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1MessageType
	SetValues(value []uint32) PatternFlowGtpv1MessageType
	// Increment returns PatternFlowGtpv1MessageTypeCounter, set in PatternFlowGtpv1MessageType.
	// PatternFlowGtpv1MessageTypeCounter is integer counter pattern
	Increment() PatternFlowGtpv1MessageTypeCounter
	// SetIncrement assigns PatternFlowGtpv1MessageTypeCounter provided by user to PatternFlowGtpv1MessageType.
	// PatternFlowGtpv1MessageTypeCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageType
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1MessageType
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1MessageTypeCounter, set in PatternFlowGtpv1MessageType.
	// PatternFlowGtpv1MessageTypeCounter is integer counter pattern
	Decrement() PatternFlowGtpv1MessageTypeCounter
	// SetDecrement assigns PatternFlowGtpv1MessageTypeCounter provided by user to PatternFlowGtpv1MessageType.
	// PatternFlowGtpv1MessageTypeCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageType
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1MessageType
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIterIter, set in PatternFlowGtpv1MessageType
	MetricTags() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	setNil()
}

type PatternFlowGtpv1MessageTypeChoiceEnum string

// Enum of Choice on PatternFlowGtpv1MessageType
var PatternFlowGtpv1MessageTypeChoice = struct {
	VALUE     PatternFlowGtpv1MessageTypeChoiceEnum
	VALUES    PatternFlowGtpv1MessageTypeChoiceEnum
	INCREMENT PatternFlowGtpv1MessageTypeChoiceEnum
	DECREMENT PatternFlowGtpv1MessageTypeChoiceEnum
}{
	VALUE:     PatternFlowGtpv1MessageTypeChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1MessageTypeChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1MessageTypeChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1MessageTypeChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1MessageType) Choice() PatternFlowGtpv1MessageTypeChoiceEnum {
	return PatternFlowGtpv1MessageTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1MessageType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1MessageType) setChoice(value PatternFlowGtpv1MessageTypeChoiceEnum) PatternFlowGtpv1MessageType {
	intValue, ok := otg.PatternFlowGtpv1MessageType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1MessageTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1MessageType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1MessageTypeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1MessageTypeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1MessageTypeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1MessageTypeCounter().msg()
	}

	if value == PatternFlowGtpv1MessageTypeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1MessageTypeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1MessageType) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1MessageTypeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1MessageType) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1MessageType object
func (obj *patternFlowGtpv1MessageType) SetValue(value uint32) PatternFlowGtpv1MessageType {
	obj.setChoice(PatternFlowGtpv1MessageTypeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1MessageType) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1MessageType object
func (obj *patternFlowGtpv1MessageType) SetValues(value []uint32) PatternFlowGtpv1MessageType {
	obj.setChoice(PatternFlowGtpv1MessageTypeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1MessageTypeCounter
func (obj *patternFlowGtpv1MessageType) Increment() PatternFlowGtpv1MessageTypeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1MessageTypeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1MessageTypeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1MessageTypeCounter
func (obj *patternFlowGtpv1MessageType) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1MessageTypeCounter value in the PatternFlowGtpv1MessageType object
func (obj *patternFlowGtpv1MessageType) SetIncrement(value PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageType {
	obj.setChoice(PatternFlowGtpv1MessageTypeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1MessageTypeCounter
func (obj *patternFlowGtpv1MessageType) Decrement() PatternFlowGtpv1MessageTypeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1MessageTypeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1MessageTypeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1MessageTypeCounter
func (obj *patternFlowGtpv1MessageType) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1MessageTypeCounter value in the PatternFlowGtpv1MessageType object
func (obj *patternFlowGtpv1MessageType) SetDecrement(value PatternFlowGtpv1MessageTypeCounter) PatternFlowGtpv1MessageType {
	obj.setChoice(PatternFlowGtpv1MessageTypeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1MessageTypeMetricTag
func (obj *patternFlowGtpv1MessageType) MetricTags() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1MessageTypeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter struct {
	obj                                       *patternFlowGtpv1MessageType
	patternFlowGtpv1MessageTypeMetricTagSlice []PatternFlowGtpv1MessageTypeMetricTag
	fieldPtr                                  *[]*otg.PatternFlowGtpv1MessageTypeMetricTag
}

func newPatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter(ptr *[]*otg.PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	return &patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter interface {
	setMsg(*patternFlowGtpv1MessageType) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	Items() []PatternFlowGtpv1MessageTypeMetricTag
	Add() PatternFlowGtpv1MessageTypeMetricTag
	Append(items ...PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	Set(index int, newObj PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	Clear() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	clearHolderSlice() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter
}

func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) setMsg(msg *patternFlowGtpv1MessageType) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1MessageTypeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) Items() []PatternFlowGtpv1MessageTypeMetricTag {
	return obj.patternFlowGtpv1MessageTypeMetricTagSlice
}

func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) Add() PatternFlowGtpv1MessageTypeMetricTag {
	newObj := &otg.PatternFlowGtpv1MessageTypeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1MessageTypeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1MessageTypeMetricTagSlice = append(obj.patternFlowGtpv1MessageTypeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) Append(items ...PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1MessageTypeMetricTagSlice = append(obj.patternFlowGtpv1MessageTypeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) Set(index int, newObj PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1MessageTypeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) Clear() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1MessageTypeMetricTag{}
		obj.patternFlowGtpv1MessageTypeMetricTagSlice = []PatternFlowGtpv1MessageTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) clearHolderSlice() PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	if len(obj.patternFlowGtpv1MessageTypeMetricTagSlice) > 0 {
		obj.patternFlowGtpv1MessageTypeMetricTagSlice = []PatternFlowGtpv1MessageTypeMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter) appendHolderSlice(item PatternFlowGtpv1MessageTypeMetricTag) PatternFlowGtpv1MessageTypePatternFlowGtpv1MessageTypeMetricTagIter {
	obj.patternFlowGtpv1MessageTypeMetricTagSlice = append(obj.patternFlowGtpv1MessageTypeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1MessageType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv1MessageType.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv1MessageType.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1MessageTypeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1MessageType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv1MessageTypeChoice.VALUE)

	}

}
