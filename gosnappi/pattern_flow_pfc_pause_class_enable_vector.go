package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseClassEnableVector *****
type patternFlowPfcPauseClassEnableVector struct {
	validation
	obj              *otg.PatternFlowPfcPauseClassEnableVector
	marshaller       marshalPatternFlowPfcPauseClassEnableVector
	unMarshaller     unMarshalPatternFlowPfcPauseClassEnableVector
	incrementHolder  PatternFlowPfcPauseClassEnableVectorCounter
	decrementHolder  PatternFlowPfcPauseClassEnableVectorCounter
	metricTagsHolder PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
}

func NewPatternFlowPfcPauseClassEnableVector() PatternFlowPfcPauseClassEnableVector {
	obj := patternFlowPfcPauseClassEnableVector{obj: &otg.PatternFlowPfcPauseClassEnableVector{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseClassEnableVector) msg() *otg.PatternFlowPfcPauseClassEnableVector {
	return obj.obj
}

func (obj *patternFlowPfcPauseClassEnableVector) setMsg(msg *otg.PatternFlowPfcPauseClassEnableVector) PatternFlowPfcPauseClassEnableVector {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseClassEnableVector struct {
	obj *patternFlowPfcPauseClassEnableVector
}

type marshalPatternFlowPfcPauseClassEnableVector interface {
	// ToProto marshals PatternFlowPfcPauseClassEnableVector to protobuf object *otg.PatternFlowPfcPauseClassEnableVector
	ToProto() (*otg.PatternFlowPfcPauseClassEnableVector, error)
	// ToPbText marshals PatternFlowPfcPauseClassEnableVector to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseClassEnableVector to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseClassEnableVector to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowPfcPauseClassEnableVector to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowPfcPauseClassEnableVector struct {
	obj *patternFlowPfcPauseClassEnableVector
}

type unMarshalPatternFlowPfcPauseClassEnableVector interface {
	// FromProto unmarshals PatternFlowPfcPauseClassEnableVector from protobuf object *otg.PatternFlowPfcPauseClassEnableVector
	FromProto(msg *otg.PatternFlowPfcPauseClassEnableVector) (PatternFlowPfcPauseClassEnableVector, error)
	// FromPbText unmarshals PatternFlowPfcPauseClassEnableVector from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseClassEnableVector from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseClassEnableVector from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseClassEnableVector) Marshal() marshalPatternFlowPfcPauseClassEnableVector {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseClassEnableVector{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseClassEnableVector) Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVector {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseClassEnableVector{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseClassEnableVector) ToProto() (*otg.PatternFlowPfcPauseClassEnableVector, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseClassEnableVector) FromProto(msg *otg.PatternFlowPfcPauseClassEnableVector) (PatternFlowPfcPauseClassEnableVector, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseClassEnableVector) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVector) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVector) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVector) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseClassEnableVector) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowPfcPauseClassEnableVector) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseClassEnableVector) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseClassEnableVector) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVector) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseClassEnableVector) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseClassEnableVector) Clone() (PatternFlowPfcPauseClassEnableVector, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseClassEnableVector()
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

func (obj *patternFlowPfcPauseClassEnableVector) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPauseClassEnableVector is destination
type PatternFlowPfcPauseClassEnableVector interface {
	Validation
	// msg marshals PatternFlowPfcPauseClassEnableVector to protobuf object *otg.PatternFlowPfcPauseClassEnableVector
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseClassEnableVector
	// setMsg unmarshals PatternFlowPfcPauseClassEnableVector from protobuf object *otg.PatternFlowPfcPauseClassEnableVector
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseClassEnableVector) PatternFlowPfcPauseClassEnableVector
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseClassEnableVector
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseClassEnableVector
	// validate validates PatternFlowPfcPauseClassEnableVector
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseClassEnableVector, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPauseClassEnableVectorChoiceEnum, set in PatternFlowPfcPauseClassEnableVector
	Choice() PatternFlowPfcPauseClassEnableVectorChoiceEnum
	// setChoice assigns PatternFlowPfcPauseClassEnableVectorChoiceEnum provided by user to PatternFlowPfcPauseClassEnableVector
	setChoice(value PatternFlowPfcPauseClassEnableVectorChoiceEnum) PatternFlowPfcPauseClassEnableVector
	// HasChoice checks if Choice has been set in PatternFlowPfcPauseClassEnableVector
	HasChoice() bool
	// Value returns uint32, set in PatternFlowPfcPauseClassEnableVector.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowPfcPauseClassEnableVector
	SetValue(value uint32) PatternFlowPfcPauseClassEnableVector
	// HasValue checks if Value has been set in PatternFlowPfcPauseClassEnableVector
	HasValue() bool
	// Values returns []uint32, set in PatternFlowPfcPauseClassEnableVector.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowPfcPauseClassEnableVector
	SetValues(value []uint32) PatternFlowPfcPauseClassEnableVector
	// Increment returns PatternFlowPfcPauseClassEnableVectorCounter, set in PatternFlowPfcPauseClassEnableVector.
	// PatternFlowPfcPauseClassEnableVectorCounter is integer counter pattern
	Increment() PatternFlowPfcPauseClassEnableVectorCounter
	// SetIncrement assigns PatternFlowPfcPauseClassEnableVectorCounter provided by user to PatternFlowPfcPauseClassEnableVector.
	// PatternFlowPfcPauseClassEnableVectorCounter is integer counter pattern
	SetIncrement(value PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVector
	// HasIncrement checks if Increment has been set in PatternFlowPfcPauseClassEnableVector
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPauseClassEnableVectorCounter, set in PatternFlowPfcPauseClassEnableVector.
	// PatternFlowPfcPauseClassEnableVectorCounter is integer counter pattern
	Decrement() PatternFlowPfcPauseClassEnableVectorCounter
	// SetDecrement assigns PatternFlowPfcPauseClassEnableVectorCounter provided by user to PatternFlowPfcPauseClassEnableVector.
	// PatternFlowPfcPauseClassEnableVectorCounter is integer counter pattern
	SetDecrement(value PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVector
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPauseClassEnableVector
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIterIter, set in PatternFlowPfcPauseClassEnableVector
	MetricTags() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	setNil()
}

type PatternFlowPfcPauseClassEnableVectorChoiceEnum string

// Enum of Choice on PatternFlowPfcPauseClassEnableVector
var PatternFlowPfcPauseClassEnableVectorChoice = struct {
	VALUE     PatternFlowPfcPauseClassEnableVectorChoiceEnum
	VALUES    PatternFlowPfcPauseClassEnableVectorChoiceEnum
	INCREMENT PatternFlowPfcPauseClassEnableVectorChoiceEnum
	DECREMENT PatternFlowPfcPauseClassEnableVectorChoiceEnum
}{
	VALUE:     PatternFlowPfcPauseClassEnableVectorChoiceEnum("value"),
	VALUES:    PatternFlowPfcPauseClassEnableVectorChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPauseClassEnableVectorChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPauseClassEnableVectorChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPauseClassEnableVector) Choice() PatternFlowPfcPauseClassEnableVectorChoiceEnum {
	return PatternFlowPfcPauseClassEnableVectorChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPauseClassEnableVector) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPauseClassEnableVector) setChoice(value PatternFlowPfcPauseClassEnableVectorChoiceEnum) PatternFlowPfcPauseClassEnableVector {
	intValue, ok := otg.PatternFlowPfcPauseClassEnableVector_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPauseClassEnableVectorChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPauseClassEnableVector_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPauseClassEnableVectorChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPauseClassEnableVectorChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPauseClassEnableVectorChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPauseClassEnableVectorCounter().msg()
	}

	if value == PatternFlowPfcPauseClassEnableVectorChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPauseClassEnableVectorCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseClassEnableVector) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowPfcPauseClassEnableVector) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowPfcPauseClassEnableVector object
func (obj *patternFlowPfcPauseClassEnableVector) SetValue(value uint32) PatternFlowPfcPauseClassEnableVector {
	obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowPfcPauseClassEnableVector) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowPfcPauseClassEnableVector object
func (obj *patternFlowPfcPauseClassEnableVector) SetValues(value []uint32) PatternFlowPfcPauseClassEnableVector {
	obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPauseClassEnableVectorCounter
func (obj *patternFlowPfcPauseClassEnableVector) Increment() PatternFlowPfcPauseClassEnableVectorCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPauseClassEnableVectorCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPauseClassEnableVectorCounter
func (obj *patternFlowPfcPauseClassEnableVector) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPauseClassEnableVectorCounter value in the PatternFlowPfcPauseClassEnableVector object
func (obj *patternFlowPfcPauseClassEnableVector) SetIncrement(value PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVector {
	obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseClassEnableVectorCounter
func (obj *patternFlowPfcPauseClassEnableVector) Decrement() PatternFlowPfcPauseClassEnableVectorCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPauseClassEnableVectorCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseClassEnableVectorCounter
func (obj *patternFlowPfcPauseClassEnableVector) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPauseClassEnableVectorCounter value in the PatternFlowPfcPauseClassEnableVector object
func (obj *patternFlowPfcPauseClassEnableVector) SetDecrement(value PatternFlowPfcPauseClassEnableVectorCounter) PatternFlowPfcPauseClassEnableVector {
	obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPauseClassEnableVectorMetricTag
func (obj *patternFlowPfcPauseClassEnableVector) MetricTags() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPauseClassEnableVectorMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter struct {
	obj                                                *patternFlowPfcPauseClassEnableVector
	patternFlowPfcPauseClassEnableVectorMetricTagSlice []PatternFlowPfcPauseClassEnableVectorMetricTag
	fieldPtr                                           *[]*otg.PatternFlowPfcPauseClassEnableVectorMetricTag
}

func newPatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter(ptr *[]*otg.PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	return &patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter interface {
	setMsg(*patternFlowPfcPauseClassEnableVector) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	Items() []PatternFlowPfcPauseClassEnableVectorMetricTag
	Add() PatternFlowPfcPauseClassEnableVectorMetricTag
	Append(items ...PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	Set(index int, newObj PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	Clear() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	clearHolderSlice() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
	appendHolderSlice(item PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter
}

func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) setMsg(msg *patternFlowPfcPauseClassEnableVector) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPauseClassEnableVectorMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) Items() []PatternFlowPfcPauseClassEnableVectorMetricTag {
	return obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice
}

func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) Add() PatternFlowPfcPauseClassEnableVectorMetricTag {
	newObj := &otg.PatternFlowPfcPauseClassEnableVectorMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPauseClassEnableVectorMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice = append(obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) Append(items ...PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice = append(obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) Set(index int, newObj PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) Clear() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPauseClassEnableVectorMetricTag{}
		obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice = []PatternFlowPfcPauseClassEnableVectorMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) clearHolderSlice() PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	if len(obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice) > 0 {
		obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice = []PatternFlowPfcPauseClassEnableVectorMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter) appendHolderSlice(item PatternFlowPfcPauseClassEnableVectorMetricTag) PatternFlowPfcPauseClassEnableVectorPatternFlowPfcPauseClassEnableVectorMetricTagIter {
	obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice = append(obj.patternFlowPfcPauseClassEnableVectorMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPauseClassEnableVector) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowPfcPauseClassEnableVector.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowPfcPauseClassEnableVector.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPauseClassEnableVectorMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPauseClassEnableVector) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPauseClassEnableVectorChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseClassEnableVectorChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPauseClassEnableVectorChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseClassEnableVectorChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseClassEnableVectorChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPauseClassEnableVectorChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPauseClassEnableVector")
			}
		} else {
			intVal := otg.PatternFlowPfcPauseClassEnableVector_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPauseClassEnableVector_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
