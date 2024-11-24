package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseTime *****
type patternFlowEthernetPauseTime struct {
	validation
	obj              *otg.PatternFlowEthernetPauseTime
	marshaller       marshalPatternFlowEthernetPauseTime
	unMarshaller     unMarshalPatternFlowEthernetPauseTime
	incrementHolder  PatternFlowEthernetPauseTimeCounter
	decrementHolder  PatternFlowEthernetPauseTimeCounter
	metricTagsHolder PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
}

func NewPatternFlowEthernetPauseTime() PatternFlowEthernetPauseTime {
	obj := patternFlowEthernetPauseTime{obj: &otg.PatternFlowEthernetPauseTime{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseTime) msg() *otg.PatternFlowEthernetPauseTime {
	return obj.obj
}

func (obj *patternFlowEthernetPauseTime) setMsg(msg *otg.PatternFlowEthernetPauseTime) PatternFlowEthernetPauseTime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseTime struct {
	obj *patternFlowEthernetPauseTime
}

type marshalPatternFlowEthernetPauseTime interface {
	// ToProto marshals PatternFlowEthernetPauseTime to protobuf object *otg.PatternFlowEthernetPauseTime
	ToProto() (*otg.PatternFlowEthernetPauseTime, error)
	// ToPbText marshals PatternFlowEthernetPauseTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseTime to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPauseTime to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPauseTime struct {
	obj *patternFlowEthernetPauseTime
}

type unMarshalPatternFlowEthernetPauseTime interface {
	// FromProto unmarshals PatternFlowEthernetPauseTime from protobuf object *otg.PatternFlowEthernetPauseTime
	FromProto(msg *otg.PatternFlowEthernetPauseTime) (PatternFlowEthernetPauseTime, error)
	// FromPbText unmarshals PatternFlowEthernetPauseTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseTime from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseTime) Marshal() marshalPatternFlowEthernetPauseTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseTime) Unmarshal() unMarshalPatternFlowEthernetPauseTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseTime) ToProto() (*otg.PatternFlowEthernetPauseTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseTime) FromProto(msg *otg.PatternFlowEthernetPauseTime) (PatternFlowEthernetPauseTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseTime) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTime) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTime) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTime) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseTime) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPauseTime) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseTime) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseTime) Clone() (PatternFlowEthernetPauseTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseTime()
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

func (obj *patternFlowEthernetPauseTime) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPauseTime is time
type PatternFlowEthernetPauseTime interface {
	Validation
	// msg marshals PatternFlowEthernetPauseTime to protobuf object *otg.PatternFlowEthernetPauseTime
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseTime
	// setMsg unmarshals PatternFlowEthernetPauseTime from protobuf object *otg.PatternFlowEthernetPauseTime
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseTime) PatternFlowEthernetPauseTime
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseTime
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseTime
	// validate validates PatternFlowEthernetPauseTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPauseTimeChoiceEnum, set in PatternFlowEthernetPauseTime
	Choice() PatternFlowEthernetPauseTimeChoiceEnum
	// setChoice assigns PatternFlowEthernetPauseTimeChoiceEnum provided by user to PatternFlowEthernetPauseTime
	setChoice(value PatternFlowEthernetPauseTimeChoiceEnum) PatternFlowEthernetPauseTime
	// HasChoice checks if Choice has been set in PatternFlowEthernetPauseTime
	HasChoice() bool
	// Value returns uint32, set in PatternFlowEthernetPauseTime.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowEthernetPauseTime
	SetValue(value uint32) PatternFlowEthernetPauseTime
	// HasValue checks if Value has been set in PatternFlowEthernetPauseTime
	HasValue() bool
	// Values returns []uint32, set in PatternFlowEthernetPauseTime.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowEthernetPauseTime
	SetValues(value []uint32) PatternFlowEthernetPauseTime
	// Increment returns PatternFlowEthernetPauseTimeCounter, set in PatternFlowEthernetPauseTime.
	// PatternFlowEthernetPauseTimeCounter is integer counter pattern
	Increment() PatternFlowEthernetPauseTimeCounter
	// SetIncrement assigns PatternFlowEthernetPauseTimeCounter provided by user to PatternFlowEthernetPauseTime.
	// PatternFlowEthernetPauseTimeCounter is integer counter pattern
	SetIncrement(value PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTime
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPauseTime
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPauseTimeCounter, set in PatternFlowEthernetPauseTime.
	// PatternFlowEthernetPauseTimeCounter is integer counter pattern
	Decrement() PatternFlowEthernetPauseTimeCounter
	// SetDecrement assigns PatternFlowEthernetPauseTimeCounter provided by user to PatternFlowEthernetPauseTime.
	// PatternFlowEthernetPauseTimeCounter is integer counter pattern
	SetDecrement(value PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTime
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPauseTime
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIterIter, set in PatternFlowEthernetPauseTime
	MetricTags() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	setNil()
}

type PatternFlowEthernetPauseTimeChoiceEnum string

// Enum of Choice on PatternFlowEthernetPauseTime
var PatternFlowEthernetPauseTimeChoice = struct {
	VALUE     PatternFlowEthernetPauseTimeChoiceEnum
	VALUES    PatternFlowEthernetPauseTimeChoiceEnum
	INCREMENT PatternFlowEthernetPauseTimeChoiceEnum
	DECREMENT PatternFlowEthernetPauseTimeChoiceEnum
}{
	VALUE:     PatternFlowEthernetPauseTimeChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPauseTimeChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPauseTimeChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPauseTimeChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPauseTime) Choice() PatternFlowEthernetPauseTimeChoiceEnum {
	return PatternFlowEthernetPauseTimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPauseTime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPauseTime) setChoice(value PatternFlowEthernetPauseTimeChoiceEnum) PatternFlowEthernetPauseTime {
	intValue, ok := otg.PatternFlowEthernetPauseTime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPauseTimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPauseTime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPauseTimeChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPauseTimeChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPauseTimeChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPauseTimeCounter().msg()
	}

	if value == PatternFlowEthernetPauseTimeChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPauseTimeCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseTime) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPauseTimeChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPauseTime) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowEthernetPauseTime object
func (obj *patternFlowEthernetPauseTime) SetValue(value uint32) PatternFlowEthernetPauseTime {
	obj.setChoice(PatternFlowEthernetPauseTimeChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowEthernetPauseTime) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowEthernetPauseTime object
func (obj *patternFlowEthernetPauseTime) SetValues(value []uint32) PatternFlowEthernetPauseTime {
	obj.setChoice(PatternFlowEthernetPauseTimeChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseTimeCounter
func (obj *patternFlowEthernetPauseTime) Increment() PatternFlowEthernetPauseTimeCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPauseTimeChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPauseTimeCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseTimeCounter
func (obj *patternFlowEthernetPauseTime) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPauseTimeCounter value in the PatternFlowEthernetPauseTime object
func (obj *patternFlowEthernetPauseTime) SetIncrement(value PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTime {
	obj.setChoice(PatternFlowEthernetPauseTimeChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseTimeCounter
func (obj *patternFlowEthernetPauseTime) Decrement() PatternFlowEthernetPauseTimeCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPauseTimeChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPauseTimeCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseTimeCounter
func (obj *patternFlowEthernetPauseTime) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPauseTimeCounter value in the PatternFlowEthernetPauseTime object
func (obj *patternFlowEthernetPauseTime) SetDecrement(value PatternFlowEthernetPauseTimeCounter) PatternFlowEthernetPauseTime {
	obj.setChoice(PatternFlowEthernetPauseTimeChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPauseTimeMetricTag
func (obj *patternFlowEthernetPauseTime) MetricTags() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPauseTimeMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter struct {
	obj                                        *patternFlowEthernetPauseTime
	patternFlowEthernetPauseTimeMetricTagSlice []PatternFlowEthernetPauseTimeMetricTag
	fieldPtr                                   *[]*otg.PatternFlowEthernetPauseTimeMetricTag
}

func newPatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter(ptr *[]*otg.PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	return &patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter interface {
	setMsg(*patternFlowEthernetPauseTime) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	Items() []PatternFlowEthernetPauseTimeMetricTag
	Add() PatternFlowEthernetPauseTimeMetricTag
	Append(items ...PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	Set(index int, newObj PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	Clear() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	clearHolderSlice() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter
}

func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) setMsg(msg *patternFlowEthernetPauseTime) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPauseTimeMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) Items() []PatternFlowEthernetPauseTimeMetricTag {
	return obj.patternFlowEthernetPauseTimeMetricTagSlice
}

func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) Add() PatternFlowEthernetPauseTimeMetricTag {
	newObj := &otg.PatternFlowEthernetPauseTimeMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPauseTimeMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPauseTimeMetricTagSlice = append(obj.patternFlowEthernetPauseTimeMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) Append(items ...PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPauseTimeMetricTagSlice = append(obj.patternFlowEthernetPauseTimeMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) Set(index int, newObj PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPauseTimeMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) Clear() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPauseTimeMetricTag{}
		obj.patternFlowEthernetPauseTimeMetricTagSlice = []PatternFlowEthernetPauseTimeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) clearHolderSlice() PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	if len(obj.patternFlowEthernetPauseTimeMetricTagSlice) > 0 {
		obj.patternFlowEthernetPauseTimeMetricTagSlice = []PatternFlowEthernetPauseTimeMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter) appendHolderSlice(item PatternFlowEthernetPauseTimeMetricTag) PatternFlowEthernetPauseTimePatternFlowEthernetPauseTimeMetricTagIter {
	obj.patternFlowEthernetPauseTimeMetricTagSlice = append(obj.patternFlowEthernetPauseTimeMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPauseTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPauseTime.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowEthernetPauseTime.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPauseTimeMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPauseTime) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetPauseTimeChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseTimeChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetPauseTimeChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseTimeChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseTimeChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetPauseTimeChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetPauseTime")
			}
		} else {
			intVal := otg.PatternFlowEthernetPauseTime_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetPauseTime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
