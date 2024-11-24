package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowMplsBottomOfStack *****
type patternFlowMplsBottomOfStack struct {
	validation
	obj              *otg.PatternFlowMplsBottomOfStack
	marshaller       marshalPatternFlowMplsBottomOfStack
	unMarshaller     unMarshalPatternFlowMplsBottomOfStack
	incrementHolder  PatternFlowMplsBottomOfStackCounter
	decrementHolder  PatternFlowMplsBottomOfStackCounter
	metricTagsHolder PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
}

func NewPatternFlowMplsBottomOfStack() PatternFlowMplsBottomOfStack {
	obj := patternFlowMplsBottomOfStack{obj: &otg.PatternFlowMplsBottomOfStack{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowMplsBottomOfStack) msg() *otg.PatternFlowMplsBottomOfStack {
	return obj.obj
}

func (obj *patternFlowMplsBottomOfStack) setMsg(msg *otg.PatternFlowMplsBottomOfStack) PatternFlowMplsBottomOfStack {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowMplsBottomOfStack struct {
	obj *patternFlowMplsBottomOfStack
}

type marshalPatternFlowMplsBottomOfStack interface {
	// ToProto marshals PatternFlowMplsBottomOfStack to protobuf object *otg.PatternFlowMplsBottomOfStack
	ToProto() (*otg.PatternFlowMplsBottomOfStack, error)
	// ToPbText marshals PatternFlowMplsBottomOfStack to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowMplsBottomOfStack to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowMplsBottomOfStack to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowMplsBottomOfStack to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowMplsBottomOfStack struct {
	obj *patternFlowMplsBottomOfStack
}

type unMarshalPatternFlowMplsBottomOfStack interface {
	// FromProto unmarshals PatternFlowMplsBottomOfStack from protobuf object *otg.PatternFlowMplsBottomOfStack
	FromProto(msg *otg.PatternFlowMplsBottomOfStack) (PatternFlowMplsBottomOfStack, error)
	// FromPbText unmarshals PatternFlowMplsBottomOfStack from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowMplsBottomOfStack from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowMplsBottomOfStack from JSON text
	FromJson(value string) error
}

func (obj *patternFlowMplsBottomOfStack) Marshal() marshalPatternFlowMplsBottomOfStack {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowMplsBottomOfStack{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowMplsBottomOfStack) Unmarshal() unMarshalPatternFlowMplsBottomOfStack {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowMplsBottomOfStack{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowMplsBottomOfStack) ToProto() (*otg.PatternFlowMplsBottomOfStack, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowMplsBottomOfStack) FromProto(msg *otg.PatternFlowMplsBottomOfStack) (PatternFlowMplsBottomOfStack, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowMplsBottomOfStack) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStack) FromPbText(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStack) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStack) FromYaml(value string) error {
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

func (m *marshalpatternFlowMplsBottomOfStack) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowMplsBottomOfStack) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowMplsBottomOfStack) FromJson(value string) error {
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

func (obj *patternFlowMplsBottomOfStack) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStack) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowMplsBottomOfStack) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowMplsBottomOfStack) Clone() (PatternFlowMplsBottomOfStack, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowMplsBottomOfStack()
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

func (obj *patternFlowMplsBottomOfStack) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowMplsBottomOfStack is bottom of stack
type PatternFlowMplsBottomOfStack interface {
	Validation
	// msg marshals PatternFlowMplsBottomOfStack to protobuf object *otg.PatternFlowMplsBottomOfStack
	// and doesn't set defaults
	msg() *otg.PatternFlowMplsBottomOfStack
	// setMsg unmarshals PatternFlowMplsBottomOfStack from protobuf object *otg.PatternFlowMplsBottomOfStack
	// and doesn't set defaults
	setMsg(*otg.PatternFlowMplsBottomOfStack) PatternFlowMplsBottomOfStack
	// provides marshal interface
	Marshal() marshalPatternFlowMplsBottomOfStack
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowMplsBottomOfStack
	// validate validates PatternFlowMplsBottomOfStack
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowMplsBottomOfStack, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowMplsBottomOfStackChoiceEnum, set in PatternFlowMplsBottomOfStack
	Choice() PatternFlowMplsBottomOfStackChoiceEnum
	// setChoice assigns PatternFlowMplsBottomOfStackChoiceEnum provided by user to PatternFlowMplsBottomOfStack
	setChoice(value PatternFlowMplsBottomOfStackChoiceEnum) PatternFlowMplsBottomOfStack
	// HasChoice checks if Choice has been set in PatternFlowMplsBottomOfStack
	HasChoice() bool
	// Value returns uint32, set in PatternFlowMplsBottomOfStack.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowMplsBottomOfStack
	SetValue(value uint32) PatternFlowMplsBottomOfStack
	// HasValue checks if Value has been set in PatternFlowMplsBottomOfStack
	HasValue() bool
	// Values returns []uint32, set in PatternFlowMplsBottomOfStack.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowMplsBottomOfStack
	SetValues(value []uint32) PatternFlowMplsBottomOfStack
	// Auto returns uint32, set in PatternFlowMplsBottomOfStack.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternFlowMplsBottomOfStack
	HasAuto() bool
	// Increment returns PatternFlowMplsBottomOfStackCounter, set in PatternFlowMplsBottomOfStack.
	// PatternFlowMplsBottomOfStackCounter is integer counter pattern
	Increment() PatternFlowMplsBottomOfStackCounter
	// SetIncrement assigns PatternFlowMplsBottomOfStackCounter provided by user to PatternFlowMplsBottomOfStack.
	// PatternFlowMplsBottomOfStackCounter is integer counter pattern
	SetIncrement(value PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStack
	// HasIncrement checks if Increment has been set in PatternFlowMplsBottomOfStack
	HasIncrement() bool
	// Decrement returns PatternFlowMplsBottomOfStackCounter, set in PatternFlowMplsBottomOfStack.
	// PatternFlowMplsBottomOfStackCounter is integer counter pattern
	Decrement() PatternFlowMplsBottomOfStackCounter
	// SetDecrement assigns PatternFlowMplsBottomOfStackCounter provided by user to PatternFlowMplsBottomOfStack.
	// PatternFlowMplsBottomOfStackCounter is integer counter pattern
	SetDecrement(value PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStack
	// HasDecrement checks if Decrement has been set in PatternFlowMplsBottomOfStack
	HasDecrement() bool
	// MetricTags returns PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIterIter, set in PatternFlowMplsBottomOfStack
	MetricTags() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	setNil()
}

type PatternFlowMplsBottomOfStackChoiceEnum string

// Enum of Choice on PatternFlowMplsBottomOfStack
var PatternFlowMplsBottomOfStackChoice = struct {
	VALUE     PatternFlowMplsBottomOfStackChoiceEnum
	VALUES    PatternFlowMplsBottomOfStackChoiceEnum
	AUTO      PatternFlowMplsBottomOfStackChoiceEnum
	INCREMENT PatternFlowMplsBottomOfStackChoiceEnum
	DECREMENT PatternFlowMplsBottomOfStackChoiceEnum
}{
	VALUE:     PatternFlowMplsBottomOfStackChoiceEnum("value"),
	VALUES:    PatternFlowMplsBottomOfStackChoiceEnum("values"),
	AUTO:      PatternFlowMplsBottomOfStackChoiceEnum("auto"),
	INCREMENT: PatternFlowMplsBottomOfStackChoiceEnum("increment"),
	DECREMENT: PatternFlowMplsBottomOfStackChoiceEnum("decrement"),
}

func (obj *patternFlowMplsBottomOfStack) Choice() PatternFlowMplsBottomOfStackChoiceEnum {
	return PatternFlowMplsBottomOfStackChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowMplsBottomOfStack) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowMplsBottomOfStack) setChoice(value PatternFlowMplsBottomOfStackChoiceEnum) PatternFlowMplsBottomOfStack {
	intValue, ok := otg.PatternFlowMplsBottomOfStack_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowMplsBottomOfStackChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowMplsBottomOfStack_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowMplsBottomOfStackChoice.VALUE {
		defaultValue := uint32(1)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowMplsBottomOfStackChoice.VALUES {
		defaultValue := []uint32{1}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowMplsBottomOfStackChoice.AUTO {
		defaultValue := uint32(1)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowMplsBottomOfStackChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowMplsBottomOfStackCounter().msg()
	}

	if value == PatternFlowMplsBottomOfStackChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowMplsBottomOfStackCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsBottomOfStack) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowMplsBottomOfStackChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowMplsBottomOfStack) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowMplsBottomOfStack object
func (obj *patternFlowMplsBottomOfStack) SetValue(value uint32) PatternFlowMplsBottomOfStack {
	obj.setChoice(PatternFlowMplsBottomOfStackChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowMplsBottomOfStack) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{1})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowMplsBottomOfStack object
func (obj *patternFlowMplsBottomOfStack) SetValues(value []uint32) PatternFlowMplsBottomOfStack {
	obj.setChoice(PatternFlowMplsBottomOfStackChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowMplsBottomOfStack) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowMplsBottomOfStackChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternFlowMplsBottomOfStack) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowMplsBottomOfStackCounter
func (obj *patternFlowMplsBottomOfStack) Increment() PatternFlowMplsBottomOfStackCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowMplsBottomOfStackChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowMplsBottomOfStackCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowMplsBottomOfStackCounter
func (obj *patternFlowMplsBottomOfStack) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowMplsBottomOfStackCounter value in the PatternFlowMplsBottomOfStack object
func (obj *patternFlowMplsBottomOfStack) SetIncrement(value PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStack {
	obj.setChoice(PatternFlowMplsBottomOfStackChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowMplsBottomOfStackCounter
func (obj *patternFlowMplsBottomOfStack) Decrement() PatternFlowMplsBottomOfStackCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowMplsBottomOfStackChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowMplsBottomOfStackCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowMplsBottomOfStackCounter
func (obj *patternFlowMplsBottomOfStack) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowMplsBottomOfStackCounter value in the PatternFlowMplsBottomOfStack object
func (obj *patternFlowMplsBottomOfStack) SetDecrement(value PatternFlowMplsBottomOfStackCounter) PatternFlowMplsBottomOfStack {
	obj.setChoice(PatternFlowMplsBottomOfStackChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowMplsBottomOfStackMetricTag
func (obj *patternFlowMplsBottomOfStack) MetricTags() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowMplsBottomOfStackMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter struct {
	obj                                        *patternFlowMplsBottomOfStack
	patternFlowMplsBottomOfStackMetricTagSlice []PatternFlowMplsBottomOfStackMetricTag
	fieldPtr                                   *[]*otg.PatternFlowMplsBottomOfStackMetricTag
}

func newPatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter(ptr *[]*otg.PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	return &patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter{fieldPtr: ptr}
}

type PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter interface {
	setMsg(*patternFlowMplsBottomOfStack) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	Items() []PatternFlowMplsBottomOfStackMetricTag
	Add() PatternFlowMplsBottomOfStackMetricTag
	Append(items ...PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	Set(index int, newObj PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	Clear() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	clearHolderSlice() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
	appendHolderSlice(item PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter
}

func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) setMsg(msg *patternFlowMplsBottomOfStack) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowMplsBottomOfStackMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) Items() []PatternFlowMplsBottomOfStackMetricTag {
	return obj.patternFlowMplsBottomOfStackMetricTagSlice
}

func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) Add() PatternFlowMplsBottomOfStackMetricTag {
	newObj := &otg.PatternFlowMplsBottomOfStackMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowMplsBottomOfStackMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowMplsBottomOfStackMetricTagSlice = append(obj.patternFlowMplsBottomOfStackMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) Append(items ...PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowMplsBottomOfStackMetricTagSlice = append(obj.patternFlowMplsBottomOfStackMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) Set(index int, newObj PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowMplsBottomOfStackMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) Clear() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowMplsBottomOfStackMetricTag{}
		obj.patternFlowMplsBottomOfStackMetricTagSlice = []PatternFlowMplsBottomOfStackMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) clearHolderSlice() PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	if len(obj.patternFlowMplsBottomOfStackMetricTagSlice) > 0 {
		obj.patternFlowMplsBottomOfStackMetricTagSlice = []PatternFlowMplsBottomOfStackMetricTag{}
	}
	return obj
}
func (obj *patternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter) appendHolderSlice(item PatternFlowMplsBottomOfStackMetricTag) PatternFlowMplsBottomOfStackPatternFlowMplsBottomOfStackMetricTagIter {
	obj.patternFlowMplsBottomOfStackMetricTagSlice = append(obj.patternFlowMplsBottomOfStackMetricTagSlice, item)
	return obj
}

func (obj *patternFlowMplsBottomOfStack) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStack.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowMplsBottomOfStack.Values <= 1 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowMplsBottomOfStack.Auto <= 1 but Got %d", *obj.obj.Auto))
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
				obj.MetricTags().appendHolderSlice(&patternFlowMplsBottomOfStackMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowMplsBottomOfStack) setDefault() {
	var choices_set int = 0
	var choice PatternFlowMplsBottomOfStackChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowMplsBottomOfStackChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowMplsBottomOfStackChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowMplsBottomOfStackChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowMplsBottomOfStackChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowMplsBottomOfStackChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowMplsBottomOfStackChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowMplsBottomOfStack")
			}
		} else {
			intVal := otg.PatternFlowMplsBottomOfStack_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowMplsBottomOfStack_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
