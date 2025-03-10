package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPfcQueue *****
type patternFlowEthernetPfcQueue struct {
	validation
	obj              *otg.PatternFlowEthernetPfcQueue
	marshaller       marshalPatternFlowEthernetPfcQueue
	unMarshaller     unMarshalPatternFlowEthernetPfcQueue
	incrementHolder  PatternFlowEthernetPfcQueueCounter
	decrementHolder  PatternFlowEthernetPfcQueueCounter
	metricTagsHolder PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
}

func NewPatternFlowEthernetPfcQueue() PatternFlowEthernetPfcQueue {
	obj := patternFlowEthernetPfcQueue{obj: &otg.PatternFlowEthernetPfcQueue{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPfcQueue) msg() *otg.PatternFlowEthernetPfcQueue {
	return obj.obj
}

func (obj *patternFlowEthernetPfcQueue) setMsg(msg *otg.PatternFlowEthernetPfcQueue) PatternFlowEthernetPfcQueue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPfcQueue struct {
	obj *patternFlowEthernetPfcQueue
}

type marshalPatternFlowEthernetPfcQueue interface {
	// ToProto marshals PatternFlowEthernetPfcQueue to protobuf object *otg.PatternFlowEthernetPfcQueue
	ToProto() (*otg.PatternFlowEthernetPfcQueue, error)
	// ToPbText marshals PatternFlowEthernetPfcQueue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPfcQueue to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPfcQueue to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowEthernetPfcQueue to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowEthernetPfcQueue struct {
	obj *patternFlowEthernetPfcQueue
}

type unMarshalPatternFlowEthernetPfcQueue interface {
	// FromProto unmarshals PatternFlowEthernetPfcQueue from protobuf object *otg.PatternFlowEthernetPfcQueue
	FromProto(msg *otg.PatternFlowEthernetPfcQueue) (PatternFlowEthernetPfcQueue, error)
	// FromPbText unmarshals PatternFlowEthernetPfcQueue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPfcQueue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPfcQueue from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPfcQueue) Marshal() marshalPatternFlowEthernetPfcQueue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPfcQueue{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPfcQueue) Unmarshal() unMarshalPatternFlowEthernetPfcQueue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPfcQueue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPfcQueue) ToProto() (*otg.PatternFlowEthernetPfcQueue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPfcQueue) FromProto(msg *otg.PatternFlowEthernetPfcQueue) (PatternFlowEthernetPfcQueue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPfcQueue) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueue) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueue) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueue) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPfcQueue) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowEthernetPfcQueue) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPfcQueue) FromJson(value string) error {
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

func (obj *patternFlowEthernetPfcQueue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPfcQueue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPfcQueue) Clone() (PatternFlowEthernetPfcQueue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPfcQueue()
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

func (obj *patternFlowEthernetPfcQueue) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPfcQueue is priority flow control queue
type PatternFlowEthernetPfcQueue interface {
	Validation
	// msg marshals PatternFlowEthernetPfcQueue to protobuf object *otg.PatternFlowEthernetPfcQueue
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPfcQueue
	// setMsg unmarshals PatternFlowEthernetPfcQueue from protobuf object *otg.PatternFlowEthernetPfcQueue
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPfcQueue) PatternFlowEthernetPfcQueue
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPfcQueue
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPfcQueue
	// validate validates PatternFlowEthernetPfcQueue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPfcQueue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPfcQueueChoiceEnum, set in PatternFlowEthernetPfcQueue
	Choice() PatternFlowEthernetPfcQueueChoiceEnum
	// setChoice assigns PatternFlowEthernetPfcQueueChoiceEnum provided by user to PatternFlowEthernetPfcQueue
	setChoice(value PatternFlowEthernetPfcQueueChoiceEnum) PatternFlowEthernetPfcQueue
	// HasChoice checks if Choice has been set in PatternFlowEthernetPfcQueue
	HasChoice() bool
	// Value returns uint32, set in PatternFlowEthernetPfcQueue.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowEthernetPfcQueue
	SetValue(value uint32) PatternFlowEthernetPfcQueue
	// HasValue checks if Value has been set in PatternFlowEthernetPfcQueue
	HasValue() bool
	// Values returns []uint32, set in PatternFlowEthernetPfcQueue.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowEthernetPfcQueue
	SetValues(value []uint32) PatternFlowEthernetPfcQueue
	// Increment returns PatternFlowEthernetPfcQueueCounter, set in PatternFlowEthernetPfcQueue.
	// PatternFlowEthernetPfcQueueCounter is integer counter pattern
	Increment() PatternFlowEthernetPfcQueueCounter
	// SetIncrement assigns PatternFlowEthernetPfcQueueCounter provided by user to PatternFlowEthernetPfcQueue.
	// PatternFlowEthernetPfcQueueCounter is integer counter pattern
	SetIncrement(value PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueue
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPfcQueue
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPfcQueueCounter, set in PatternFlowEthernetPfcQueue.
	// PatternFlowEthernetPfcQueueCounter is integer counter pattern
	Decrement() PatternFlowEthernetPfcQueueCounter
	// SetDecrement assigns PatternFlowEthernetPfcQueueCounter provided by user to PatternFlowEthernetPfcQueue.
	// PatternFlowEthernetPfcQueueCounter is integer counter pattern
	SetDecrement(value PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueue
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPfcQueue
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIterIter, set in PatternFlowEthernetPfcQueue
	MetricTags() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	setNil()
}

type PatternFlowEthernetPfcQueueChoiceEnum string

// Enum of Choice on PatternFlowEthernetPfcQueue
var PatternFlowEthernetPfcQueueChoice = struct {
	VALUE     PatternFlowEthernetPfcQueueChoiceEnum
	VALUES    PatternFlowEthernetPfcQueueChoiceEnum
	INCREMENT PatternFlowEthernetPfcQueueChoiceEnum
	DECREMENT PatternFlowEthernetPfcQueueChoiceEnum
}{
	VALUE:     PatternFlowEthernetPfcQueueChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPfcQueueChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPfcQueueChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPfcQueueChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPfcQueue) Choice() PatternFlowEthernetPfcQueueChoiceEnum {
	return PatternFlowEthernetPfcQueueChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPfcQueue) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPfcQueue) setChoice(value PatternFlowEthernetPfcQueueChoiceEnum) PatternFlowEthernetPfcQueue {
	intValue, ok := otg.PatternFlowEthernetPfcQueue_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPfcQueueChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPfcQueue_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPfcQueueChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPfcQueueChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPfcQueueChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPfcQueueCounter().msg()
	}

	if value == PatternFlowEthernetPfcQueueChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPfcQueueCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPfcQueue) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPfcQueueChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowEthernetPfcQueue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowEthernetPfcQueue object
func (obj *patternFlowEthernetPfcQueue) SetValue(value uint32) PatternFlowEthernetPfcQueue {
	obj.setChoice(PatternFlowEthernetPfcQueueChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowEthernetPfcQueue) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowEthernetPfcQueue object
func (obj *patternFlowEthernetPfcQueue) SetValues(value []uint32) PatternFlowEthernetPfcQueue {
	obj.setChoice(PatternFlowEthernetPfcQueueChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPfcQueueCounter
func (obj *patternFlowEthernetPfcQueue) Increment() PatternFlowEthernetPfcQueueCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPfcQueueChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPfcQueueCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPfcQueueCounter
func (obj *patternFlowEthernetPfcQueue) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPfcQueueCounter value in the PatternFlowEthernetPfcQueue object
func (obj *patternFlowEthernetPfcQueue) SetIncrement(value PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueue {
	obj.setChoice(PatternFlowEthernetPfcQueueChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPfcQueueCounter
func (obj *patternFlowEthernetPfcQueue) Decrement() PatternFlowEthernetPfcQueueCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPfcQueueChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPfcQueueCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPfcQueueCounter
func (obj *patternFlowEthernetPfcQueue) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPfcQueueCounter value in the PatternFlowEthernetPfcQueue object
func (obj *patternFlowEthernetPfcQueue) SetDecrement(value PatternFlowEthernetPfcQueueCounter) PatternFlowEthernetPfcQueue {
	obj.setChoice(PatternFlowEthernetPfcQueueChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPfcQueueMetricTag
func (obj *patternFlowEthernetPfcQueue) MetricTags() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPfcQueueMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter struct {
	obj                                       *patternFlowEthernetPfcQueue
	patternFlowEthernetPfcQueueMetricTagSlice []PatternFlowEthernetPfcQueueMetricTag
	fieldPtr                                  *[]*otg.PatternFlowEthernetPfcQueueMetricTag
}

func newPatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter(ptr *[]*otg.PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	return &patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter interface {
	setMsg(*patternFlowEthernetPfcQueue) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	Items() []PatternFlowEthernetPfcQueueMetricTag
	Add() PatternFlowEthernetPfcQueueMetricTag
	Append(items ...PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	Set(index int, newObj PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	Clear() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	clearHolderSlice() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter
}

func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) setMsg(msg *patternFlowEthernetPfcQueue) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPfcQueueMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) Items() []PatternFlowEthernetPfcQueueMetricTag {
	return obj.patternFlowEthernetPfcQueueMetricTagSlice
}

func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) Add() PatternFlowEthernetPfcQueueMetricTag {
	newObj := &otg.PatternFlowEthernetPfcQueueMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPfcQueueMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPfcQueueMetricTagSlice = append(obj.patternFlowEthernetPfcQueueMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) Append(items ...PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPfcQueueMetricTagSlice = append(obj.patternFlowEthernetPfcQueueMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) Set(index int, newObj PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPfcQueueMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) Clear() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPfcQueueMetricTag{}
		obj.patternFlowEthernetPfcQueueMetricTagSlice = []PatternFlowEthernetPfcQueueMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) clearHolderSlice() PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	if len(obj.patternFlowEthernetPfcQueueMetricTagSlice) > 0 {
		obj.patternFlowEthernetPfcQueueMetricTagSlice = []PatternFlowEthernetPfcQueueMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter) appendHolderSlice(item PatternFlowEthernetPfcQueueMetricTag) PatternFlowEthernetPfcQueuePatternFlowEthernetPfcQueueMetricTagIter {
	obj.patternFlowEthernetPfcQueueMetricTagSlice = append(obj.patternFlowEthernetPfcQueueMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPfcQueue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowEthernetPfcQueue.Value <= 7 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 7 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowEthernetPfcQueue.Values <= 7 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPfcQueueMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPfcQueue) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetPfcQueueChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetPfcQueueChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetPfcQueueChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetPfcQueueChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetPfcQueueChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetPfcQueueChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetPfcQueue")
			}
		} else {
			intVal := otg.PatternFlowEthernetPfcQueue_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetPfcQueue_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
