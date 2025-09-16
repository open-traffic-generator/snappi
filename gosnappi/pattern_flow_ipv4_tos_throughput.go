package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosThroughput *****
type patternFlowIpv4TosThroughput struct {
	validation
	obj              *otg.PatternFlowIpv4TosThroughput
	marshaller       marshalPatternFlowIpv4TosThroughput
	unMarshaller     unMarshalPatternFlowIpv4TosThroughput
	incrementHolder  PatternFlowIpv4TosThroughputCounter
	decrementHolder  PatternFlowIpv4TosThroughputCounter
	metricTagsHolder PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
}

func NewPatternFlowIpv4TosThroughput() PatternFlowIpv4TosThroughput {
	obj := patternFlowIpv4TosThroughput{obj: &otg.PatternFlowIpv4TosThroughput{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosThroughput) msg() *otg.PatternFlowIpv4TosThroughput {
	return obj.obj
}

func (obj *patternFlowIpv4TosThroughput) setMsg(msg *otg.PatternFlowIpv4TosThroughput) PatternFlowIpv4TosThroughput {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosThroughput struct {
	obj *patternFlowIpv4TosThroughput
}

type marshalPatternFlowIpv4TosThroughput interface {
	// ToProto marshals PatternFlowIpv4TosThroughput to protobuf object *otg.PatternFlowIpv4TosThroughput
	ToProto() (*otg.PatternFlowIpv4TosThroughput, error)
	// ToPbText marshals PatternFlowIpv4TosThroughput to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosThroughput to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosThroughput to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosThroughput struct {
	obj *patternFlowIpv4TosThroughput
}

type unMarshalPatternFlowIpv4TosThroughput interface {
	// FromProto unmarshals PatternFlowIpv4TosThroughput from protobuf object *otg.PatternFlowIpv4TosThroughput
	FromProto(msg *otg.PatternFlowIpv4TosThroughput) (PatternFlowIpv4TosThroughput, error)
	// FromPbText unmarshals PatternFlowIpv4TosThroughput from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosThroughput from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosThroughput from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosThroughput) Marshal() marshalPatternFlowIpv4TosThroughput {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosThroughput{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosThroughput) Unmarshal() unMarshalPatternFlowIpv4TosThroughput {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosThroughput{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosThroughput) ToProto() (*otg.PatternFlowIpv4TosThroughput, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosThroughput) FromProto(msg *otg.PatternFlowIpv4TosThroughput) (PatternFlowIpv4TosThroughput, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosThroughput) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughput) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughput) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughput) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosThroughput) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosThroughput) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosThroughput) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughput) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosThroughput) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosThroughput) Clone() (PatternFlowIpv4TosThroughput, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosThroughput()
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

func (obj *patternFlowIpv4TosThroughput) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosThroughput is throughput
type PatternFlowIpv4TosThroughput interface {
	Validation
	// msg marshals PatternFlowIpv4TosThroughput to protobuf object *otg.PatternFlowIpv4TosThroughput
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosThroughput
	// setMsg unmarshals PatternFlowIpv4TosThroughput from protobuf object *otg.PatternFlowIpv4TosThroughput
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosThroughput) PatternFlowIpv4TosThroughput
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosThroughput
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosThroughput
	// validate validates PatternFlowIpv4TosThroughput
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosThroughput, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosThroughputChoiceEnum, set in PatternFlowIpv4TosThroughput
	Choice() PatternFlowIpv4TosThroughputChoiceEnum
	// setChoice assigns PatternFlowIpv4TosThroughputChoiceEnum provided by user to PatternFlowIpv4TosThroughput
	setChoice(value PatternFlowIpv4TosThroughputChoiceEnum) PatternFlowIpv4TosThroughput
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosThroughput
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosThroughput.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosThroughput
	SetValue(value uint32) PatternFlowIpv4TosThroughput
	// HasValue checks if Value has been set in PatternFlowIpv4TosThroughput
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosThroughput.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosThroughput
	SetValues(value []uint32) PatternFlowIpv4TosThroughput
	// Increment returns PatternFlowIpv4TosThroughputCounter, set in PatternFlowIpv4TosThroughput.
	// PatternFlowIpv4TosThroughputCounter is integer counter pattern
	Increment() PatternFlowIpv4TosThroughputCounter
	// SetIncrement assigns PatternFlowIpv4TosThroughputCounter provided by user to PatternFlowIpv4TosThroughput.
	// PatternFlowIpv4TosThroughputCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughput
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosThroughput
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosThroughputCounter, set in PatternFlowIpv4TosThroughput.
	// PatternFlowIpv4TosThroughputCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosThroughputCounter
	// SetDecrement assigns PatternFlowIpv4TosThroughputCounter provided by user to PatternFlowIpv4TosThroughput.
	// PatternFlowIpv4TosThroughputCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughput
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosThroughput
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIterIter, set in PatternFlowIpv4TosThroughput
	MetricTags() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	setNil()
}

type PatternFlowIpv4TosThroughputChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosThroughput
var PatternFlowIpv4TosThroughputChoice = struct {
	VALUE     PatternFlowIpv4TosThroughputChoiceEnum
	VALUES    PatternFlowIpv4TosThroughputChoiceEnum
	INCREMENT PatternFlowIpv4TosThroughputChoiceEnum
	DECREMENT PatternFlowIpv4TosThroughputChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosThroughputChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosThroughputChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosThroughputChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosThroughputChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosThroughput) Choice() PatternFlowIpv4TosThroughputChoiceEnum {
	return PatternFlowIpv4TosThroughputChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosThroughput) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosThroughput) setChoice(value PatternFlowIpv4TosThroughputChoiceEnum) PatternFlowIpv4TosThroughput {
	intValue, ok := otg.PatternFlowIpv4TosThroughput_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosThroughputChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosThroughput_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosThroughputChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosThroughputChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosThroughputChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosThroughputCounter().msg()
	}

	if value == PatternFlowIpv4TosThroughputChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosThroughputCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosThroughput) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosThroughputChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosThroughput) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosThroughput object
func (obj *patternFlowIpv4TosThroughput) SetValue(value uint32) PatternFlowIpv4TosThroughput {
	obj.setChoice(PatternFlowIpv4TosThroughputChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosThroughput) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosThroughput object
func (obj *patternFlowIpv4TosThroughput) SetValues(value []uint32) PatternFlowIpv4TosThroughput {
	obj.setChoice(PatternFlowIpv4TosThroughputChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosThroughputCounter
func (obj *patternFlowIpv4TosThroughput) Increment() PatternFlowIpv4TosThroughputCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosThroughputChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosThroughputCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosThroughputCounter
func (obj *patternFlowIpv4TosThroughput) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosThroughputCounter value in the PatternFlowIpv4TosThroughput object
func (obj *patternFlowIpv4TosThroughput) SetIncrement(value PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughput {
	obj.setChoice(PatternFlowIpv4TosThroughputChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosThroughputCounter
func (obj *patternFlowIpv4TosThroughput) Decrement() PatternFlowIpv4TosThroughputCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosThroughputChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosThroughputCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosThroughputCounter
func (obj *patternFlowIpv4TosThroughput) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosThroughputCounter value in the PatternFlowIpv4TosThroughput object
func (obj *patternFlowIpv4TosThroughput) SetDecrement(value PatternFlowIpv4TosThroughputCounter) PatternFlowIpv4TosThroughput {
	obj.setChoice(PatternFlowIpv4TosThroughputChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosThroughputMetricTag
func (obj *patternFlowIpv4TosThroughput) MetricTags() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosThroughputMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter struct {
	obj                                        *patternFlowIpv4TosThroughput
	patternFlowIpv4TosThroughputMetricTagSlice []PatternFlowIpv4TosThroughputMetricTag
	fieldPtr                                   *[]*otg.PatternFlowIpv4TosThroughputMetricTag
}

func newPatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	return &patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter interface {
	setMsg(*patternFlowIpv4TosThroughput) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	Items() []PatternFlowIpv4TosThroughputMetricTag
	Add() PatternFlowIpv4TosThroughputMetricTag
	Append(items ...PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	Clear() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter
}

func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) setMsg(msg *patternFlowIpv4TosThroughput) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosThroughputMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) Items() []PatternFlowIpv4TosThroughputMetricTag {
	return obj.patternFlowIpv4TosThroughputMetricTagSlice
}

func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) Add() PatternFlowIpv4TosThroughputMetricTag {
	newObj := &otg.PatternFlowIpv4TosThroughputMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosThroughputMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosThroughputMetricTagSlice = append(obj.patternFlowIpv4TosThroughputMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) Append(items ...PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosThroughputMetricTagSlice = append(obj.patternFlowIpv4TosThroughputMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) Set(index int, newObj PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosThroughputMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) Clear() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosThroughputMetricTag{}
		obj.patternFlowIpv4TosThroughputMetricTagSlice = []PatternFlowIpv4TosThroughputMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) clearHolderSlice() PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	if len(obj.patternFlowIpv4TosThroughputMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosThroughputMetricTagSlice = []PatternFlowIpv4TosThroughputMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosThroughputMetricTag) PatternFlowIpv4TosThroughputPatternFlowIpv4TosThroughputMetricTagIter {
	obj.patternFlowIpv4TosThroughputMetricTagSlice = append(obj.patternFlowIpv4TosThroughputMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosThroughput) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosThroughput.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosThroughput.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosThroughputMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosThroughput) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4TosThroughputChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosThroughputChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4TosThroughputChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosThroughputChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4TosThroughputChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4TosThroughputChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4TosThroughput")
			}
		} else {
			intVal := otg.PatternFlowIpv4TosThroughput_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4TosThroughput_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
