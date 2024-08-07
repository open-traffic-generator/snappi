package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoSequenceNumber *****
type patternFlowIcmpv6EchoSequenceNumber struct {
	validation
	obj              *otg.PatternFlowIcmpv6EchoSequenceNumber
	marshaller       marshalPatternFlowIcmpv6EchoSequenceNumber
	unMarshaller     unMarshalPatternFlowIcmpv6EchoSequenceNumber
	incrementHolder  PatternFlowIcmpv6EchoSequenceNumberCounter
	decrementHolder  PatternFlowIcmpv6EchoSequenceNumberCounter
	metricTagsHolder PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
}

func NewPatternFlowIcmpv6EchoSequenceNumber() PatternFlowIcmpv6EchoSequenceNumber {
	obj := patternFlowIcmpv6EchoSequenceNumber{obj: &otg.PatternFlowIcmpv6EchoSequenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) msg() *otg.PatternFlowIcmpv6EchoSequenceNumber {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) setMsg(msg *otg.PatternFlowIcmpv6EchoSequenceNumber) PatternFlowIcmpv6EchoSequenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoSequenceNumber struct {
	obj *patternFlowIcmpv6EchoSequenceNumber
}

type marshalPatternFlowIcmpv6EchoSequenceNumber interface {
	// ToProto marshals PatternFlowIcmpv6EchoSequenceNumber to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumber
	ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumber, error)
	// ToPbText marshals PatternFlowIcmpv6EchoSequenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoSequenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoSequenceNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoSequenceNumber struct {
	obj *patternFlowIcmpv6EchoSequenceNumber
}

type unMarshalPatternFlowIcmpv6EchoSequenceNumber interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoSequenceNumber from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumber
	FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumber) (PatternFlowIcmpv6EchoSequenceNumber, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoSequenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoSequenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoSequenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Marshal() marshalPatternFlowIcmpv6EchoSequenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoSequenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoSequenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumber) ToProto() (*otg.PatternFlowIcmpv6EchoSequenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumber) FromProto(msg *otg.PatternFlowIcmpv6EchoSequenceNumber) (PatternFlowIcmpv6EchoSequenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoSequenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoSequenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoSequenceNumber) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoSequenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Clone() (PatternFlowIcmpv6EchoSequenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoSequenceNumber()
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

func (obj *patternFlowIcmpv6EchoSequenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6EchoSequenceNumber is iCMPv6 echo sequence number
type PatternFlowIcmpv6EchoSequenceNumber interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoSequenceNumber to protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoSequenceNumber
	// setMsg unmarshals PatternFlowIcmpv6EchoSequenceNumber from protobuf object *otg.PatternFlowIcmpv6EchoSequenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoSequenceNumber) PatternFlowIcmpv6EchoSequenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoSequenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoSequenceNumber
	// validate validates PatternFlowIcmpv6EchoSequenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoSequenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6EchoSequenceNumberChoiceEnum, set in PatternFlowIcmpv6EchoSequenceNumber
	Choice() PatternFlowIcmpv6EchoSequenceNumberChoiceEnum
	// setChoice assigns PatternFlowIcmpv6EchoSequenceNumberChoiceEnum provided by user to PatternFlowIcmpv6EchoSequenceNumber
	setChoice(value PatternFlowIcmpv6EchoSequenceNumberChoiceEnum) PatternFlowIcmpv6EchoSequenceNumber
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6EchoSequenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6EchoSequenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumber
	SetValue(value uint32) PatternFlowIcmpv6EchoSequenceNumber
	// HasValue checks if Value has been set in PatternFlowIcmpv6EchoSequenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6EchoSequenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6EchoSequenceNumber
	SetValues(value []uint32) PatternFlowIcmpv6EchoSequenceNumber
	// Increment returns PatternFlowIcmpv6EchoSequenceNumberCounter, set in PatternFlowIcmpv6EchoSequenceNumber.
	// PatternFlowIcmpv6EchoSequenceNumberCounter is integer counter pattern
	Increment() PatternFlowIcmpv6EchoSequenceNumberCounter
	// SetIncrement assigns PatternFlowIcmpv6EchoSequenceNumberCounter provided by user to PatternFlowIcmpv6EchoSequenceNumber.
	// PatternFlowIcmpv6EchoSequenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6EchoSequenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6EchoSequenceNumberCounter, set in PatternFlowIcmpv6EchoSequenceNumber.
	// PatternFlowIcmpv6EchoSequenceNumberCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6EchoSequenceNumberCounter
	// SetDecrement assigns PatternFlowIcmpv6EchoSequenceNumberCounter provided by user to PatternFlowIcmpv6EchoSequenceNumber.
	// PatternFlowIcmpv6EchoSequenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6EchoSequenceNumber
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIterIter, set in PatternFlowIcmpv6EchoSequenceNumber
	MetricTags() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	setNil()
}

type PatternFlowIcmpv6EchoSequenceNumberChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6EchoSequenceNumber
var PatternFlowIcmpv6EchoSequenceNumberChoice = struct {
	VALUE     PatternFlowIcmpv6EchoSequenceNumberChoiceEnum
	VALUES    PatternFlowIcmpv6EchoSequenceNumberChoiceEnum
	INCREMENT PatternFlowIcmpv6EchoSequenceNumberChoiceEnum
	DECREMENT PatternFlowIcmpv6EchoSequenceNumberChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6EchoSequenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6EchoSequenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6EchoSequenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6EchoSequenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) Choice() PatternFlowIcmpv6EchoSequenceNumberChoiceEnum {
	return PatternFlowIcmpv6EchoSequenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6EchoSequenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) setChoice(value PatternFlowIcmpv6EchoSequenceNumberChoiceEnum) PatternFlowIcmpv6EchoSequenceNumber {
	intValue, ok := otg.PatternFlowIcmpv6EchoSequenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoSequenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoSequenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6EchoSequenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6EchoSequenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6EchoSequenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6EchoSequenceNumberCounter().msg()
	}

	if value == PatternFlowIcmpv6EchoSequenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6EchoSequenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoSequenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6EchoSequenceNumber object
func (obj *patternFlowIcmpv6EchoSequenceNumber) SetValue(value uint32) PatternFlowIcmpv6EchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6EchoSequenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6EchoSequenceNumber object
func (obj *patternFlowIcmpv6EchoSequenceNumber) SetValues(value []uint32) PatternFlowIcmpv6EchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoSequenceNumberCounter
func (obj *patternFlowIcmpv6EchoSequenceNumber) Increment() PatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6EchoSequenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoSequenceNumberCounter
func (obj *patternFlowIcmpv6EchoSequenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6EchoSequenceNumberCounter value in the PatternFlowIcmpv6EchoSequenceNumber object
func (obj *patternFlowIcmpv6EchoSequenceNumber) SetIncrement(value PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoSequenceNumberCounter
func (obj *patternFlowIcmpv6EchoSequenceNumber) Decrement() PatternFlowIcmpv6EchoSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6EchoSequenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoSequenceNumberCounter
func (obj *patternFlowIcmpv6EchoSequenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6EchoSequenceNumberCounter value in the PatternFlowIcmpv6EchoSequenceNumber object
func (obj *patternFlowIcmpv6EchoSequenceNumber) SetDecrement(value PatternFlowIcmpv6EchoSequenceNumberCounter) PatternFlowIcmpv6EchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpv6EchoSequenceNumberMetricTag
func (obj *patternFlowIcmpv6EchoSequenceNumber) MetricTags() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter struct {
	obj                                               *patternFlowIcmpv6EchoSequenceNumber
	patternFlowIcmpv6EchoSequenceNumberMetricTagSlice []PatternFlowIcmpv6EchoSequenceNumberMetricTag
	fieldPtr                                          *[]*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag
}

func newPatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter(ptr *[]*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	return &patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter interface {
	setMsg(*patternFlowIcmpv6EchoSequenceNumber) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	Items() []PatternFlowIcmpv6EchoSequenceNumberMetricTag
	Add() PatternFlowIcmpv6EchoSequenceNumberMetricTag
	Append(items ...PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	Set(index int, newObj PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	Clear() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	clearHolderSlice() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
	appendHolderSlice(item PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter
}

func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) setMsg(msg *patternFlowIcmpv6EchoSequenceNumber) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpv6EchoSequenceNumberMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) Items() []PatternFlowIcmpv6EchoSequenceNumberMetricTag {
	return obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice
}

func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) Add() PatternFlowIcmpv6EchoSequenceNumberMetricTag {
	newObj := &otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpv6EchoSequenceNumberMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) Append(items ...PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) Set(index int, newObj PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) Clear() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpv6EchoSequenceNumberMetricTag{}
		obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice = []PatternFlowIcmpv6EchoSequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) clearHolderSlice() PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	if len(obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice) > 0 {
		obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice = []PatternFlowIcmpv6EchoSequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter) appendHolderSlice(item PatternFlowIcmpv6EchoSequenceNumberMetricTag) PatternFlowIcmpv6EchoSequenceNumberPatternFlowIcmpv6EchoSequenceNumberMetricTagIter {
	obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpv6EchoSequenceNumberMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpv6EchoSequenceNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoSequenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6EchoSequenceNumber.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpv6EchoSequenceNumberMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpv6EchoSequenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpv6EchoSequenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoSequenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoSequenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoSequenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpv6EchoSequenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpv6EchoSequenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpv6EchoSequenceNumber")
			}
		} else {
			intVal := otg.PatternFlowIcmpv6EchoSequenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpv6EchoSequenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
