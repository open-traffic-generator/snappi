package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoSequenceNumber *****
type patternFlowIcmpEchoSequenceNumber struct {
	validation
	obj              *otg.PatternFlowIcmpEchoSequenceNumber
	marshaller       marshalPatternFlowIcmpEchoSequenceNumber
	unMarshaller     unMarshalPatternFlowIcmpEchoSequenceNumber
	incrementHolder  PatternFlowIcmpEchoSequenceNumberCounter
	decrementHolder  PatternFlowIcmpEchoSequenceNumberCounter
	metricTagsHolder PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
}

func NewPatternFlowIcmpEchoSequenceNumber() PatternFlowIcmpEchoSequenceNumber {
	obj := patternFlowIcmpEchoSequenceNumber{obj: &otg.PatternFlowIcmpEchoSequenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) msg() *otg.PatternFlowIcmpEchoSequenceNumber {
	return obj.obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) setMsg(msg *otg.PatternFlowIcmpEchoSequenceNumber) PatternFlowIcmpEchoSequenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoSequenceNumber struct {
	obj *patternFlowIcmpEchoSequenceNumber
}

type marshalPatternFlowIcmpEchoSequenceNumber interface {
	// ToProto marshals PatternFlowIcmpEchoSequenceNumber to protobuf object *otg.PatternFlowIcmpEchoSequenceNumber
	ToProto() (*otg.PatternFlowIcmpEchoSequenceNumber, error)
	// ToPbText marshals PatternFlowIcmpEchoSequenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoSequenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoSequenceNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpEchoSequenceNumber struct {
	obj *patternFlowIcmpEchoSequenceNumber
}

type unMarshalPatternFlowIcmpEchoSequenceNumber interface {
	// FromProto unmarshals PatternFlowIcmpEchoSequenceNumber from protobuf object *otg.PatternFlowIcmpEchoSequenceNumber
	FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumber) (PatternFlowIcmpEchoSequenceNumber, error)
	// FromPbText unmarshals PatternFlowIcmpEchoSequenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoSequenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoSequenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoSequenceNumber) Marshal() marshalPatternFlowIcmpEchoSequenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoSequenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoSequenceNumber) Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoSequenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoSequenceNumber) ToProto() (*otg.PatternFlowIcmpEchoSequenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoSequenceNumber) FromProto(msg *otg.PatternFlowIcmpEchoSequenceNumber) (PatternFlowIcmpEchoSequenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoSequenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoSequenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoSequenceNumber) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoSequenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoSequenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoSequenceNumber) Clone() (PatternFlowIcmpEchoSequenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoSequenceNumber()
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

func (obj *patternFlowIcmpEchoSequenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpEchoSequenceNumber is iCMP sequence number
type PatternFlowIcmpEchoSequenceNumber interface {
	Validation
	// msg marshals PatternFlowIcmpEchoSequenceNumber to protobuf object *otg.PatternFlowIcmpEchoSequenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoSequenceNumber
	// setMsg unmarshals PatternFlowIcmpEchoSequenceNumber from protobuf object *otg.PatternFlowIcmpEchoSequenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoSequenceNumber) PatternFlowIcmpEchoSequenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoSequenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoSequenceNumber
	// validate validates PatternFlowIcmpEchoSequenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoSequenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpEchoSequenceNumberChoiceEnum, set in PatternFlowIcmpEchoSequenceNumber
	Choice() PatternFlowIcmpEchoSequenceNumberChoiceEnum
	// setChoice assigns PatternFlowIcmpEchoSequenceNumberChoiceEnum provided by user to PatternFlowIcmpEchoSequenceNumber
	setChoice(value PatternFlowIcmpEchoSequenceNumberChoiceEnum) PatternFlowIcmpEchoSequenceNumber
	// HasChoice checks if Choice has been set in PatternFlowIcmpEchoSequenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpEchoSequenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpEchoSequenceNumber
	SetValue(value uint32) PatternFlowIcmpEchoSequenceNumber
	// HasValue checks if Value has been set in PatternFlowIcmpEchoSequenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpEchoSequenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpEchoSequenceNumber
	SetValues(value []uint32) PatternFlowIcmpEchoSequenceNumber
	// Increment returns PatternFlowIcmpEchoSequenceNumberCounter, set in PatternFlowIcmpEchoSequenceNumber.
	// PatternFlowIcmpEchoSequenceNumberCounter is integer counter pattern
	Increment() PatternFlowIcmpEchoSequenceNumberCounter
	// SetIncrement assigns PatternFlowIcmpEchoSequenceNumberCounter provided by user to PatternFlowIcmpEchoSequenceNumber.
	// PatternFlowIcmpEchoSequenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowIcmpEchoSequenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpEchoSequenceNumberCounter, set in PatternFlowIcmpEchoSequenceNumber.
	// PatternFlowIcmpEchoSequenceNumberCounter is integer counter pattern
	Decrement() PatternFlowIcmpEchoSequenceNumberCounter
	// SetDecrement assigns PatternFlowIcmpEchoSequenceNumberCounter provided by user to PatternFlowIcmpEchoSequenceNumber.
	// PatternFlowIcmpEchoSequenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpEchoSequenceNumber
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIterIter, set in PatternFlowIcmpEchoSequenceNumber
	MetricTags() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	setNil()
}

type PatternFlowIcmpEchoSequenceNumberChoiceEnum string

// Enum of Choice on PatternFlowIcmpEchoSequenceNumber
var PatternFlowIcmpEchoSequenceNumberChoice = struct {
	VALUE     PatternFlowIcmpEchoSequenceNumberChoiceEnum
	VALUES    PatternFlowIcmpEchoSequenceNumberChoiceEnum
	INCREMENT PatternFlowIcmpEchoSequenceNumberChoiceEnum
	DECREMENT PatternFlowIcmpEchoSequenceNumberChoiceEnum
}{
	VALUE:     PatternFlowIcmpEchoSequenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowIcmpEchoSequenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpEchoSequenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpEchoSequenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpEchoSequenceNumber) Choice() PatternFlowIcmpEchoSequenceNumberChoiceEnum {
	return PatternFlowIcmpEchoSequenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpEchoSequenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpEchoSequenceNumber) setChoice(value PatternFlowIcmpEchoSequenceNumberChoiceEnum) PatternFlowIcmpEchoSequenceNumber {
	intValue, ok := otg.PatternFlowIcmpEchoSequenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoSequenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoSequenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpEchoSequenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpEchoSequenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpEchoSequenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpEchoSequenceNumberCounter().msg()
	}

	if value == PatternFlowIcmpEchoSequenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpEchoSequenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoSequenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpEchoSequenceNumber object
func (obj *patternFlowIcmpEchoSequenceNumber) SetValue(value uint32) PatternFlowIcmpEchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpEchoSequenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpEchoSequenceNumber object
func (obj *patternFlowIcmpEchoSequenceNumber) SetValues(value []uint32) PatternFlowIcmpEchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoSequenceNumberCounter
func (obj *patternFlowIcmpEchoSequenceNumber) Increment() PatternFlowIcmpEchoSequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpEchoSequenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoSequenceNumberCounter
func (obj *patternFlowIcmpEchoSequenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpEchoSequenceNumberCounter value in the PatternFlowIcmpEchoSequenceNumber object
func (obj *patternFlowIcmpEchoSequenceNumber) SetIncrement(value PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoSequenceNumberCounter
func (obj *patternFlowIcmpEchoSequenceNumber) Decrement() PatternFlowIcmpEchoSequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpEchoSequenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoSequenceNumberCounter
func (obj *patternFlowIcmpEchoSequenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpEchoSequenceNumberCounter value in the PatternFlowIcmpEchoSequenceNumber object
func (obj *patternFlowIcmpEchoSequenceNumber) SetDecrement(value PatternFlowIcmpEchoSequenceNumberCounter) PatternFlowIcmpEchoSequenceNumber {
	obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpEchoSequenceNumberMetricTag
func (obj *patternFlowIcmpEchoSequenceNumber) MetricTags() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpEchoSequenceNumberMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter struct {
	obj                                             *patternFlowIcmpEchoSequenceNumber
	patternFlowIcmpEchoSequenceNumberMetricTagSlice []PatternFlowIcmpEchoSequenceNumberMetricTag
	fieldPtr                                        *[]*otg.PatternFlowIcmpEchoSequenceNumberMetricTag
}

func newPatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter(ptr *[]*otg.PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	return &patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter interface {
	setMsg(*patternFlowIcmpEchoSequenceNumber) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	Items() []PatternFlowIcmpEchoSequenceNumberMetricTag
	Add() PatternFlowIcmpEchoSequenceNumberMetricTag
	Append(items ...PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	Set(index int, newObj PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	Clear() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	clearHolderSlice() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
	appendHolderSlice(item PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter
}

func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) setMsg(msg *patternFlowIcmpEchoSequenceNumber) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpEchoSequenceNumberMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) Items() []PatternFlowIcmpEchoSequenceNumberMetricTag {
	return obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice
}

func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) Add() PatternFlowIcmpEchoSequenceNumberMetricTag {
	newObj := &otg.PatternFlowIcmpEchoSequenceNumberMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpEchoSequenceNumberMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) Append(items ...PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) Set(index int, newObj PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) Clear() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpEchoSequenceNumberMetricTag{}
		obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice = []PatternFlowIcmpEchoSequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) clearHolderSlice() PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	if len(obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice) > 0 {
		obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice = []PatternFlowIcmpEchoSequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter) appendHolderSlice(item PatternFlowIcmpEchoSequenceNumberMetricTag) PatternFlowIcmpEchoSequenceNumberPatternFlowIcmpEchoSequenceNumberMetricTagIter {
	obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice = append(obj.patternFlowIcmpEchoSequenceNumberMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpEchoSequenceNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoSequenceNumber.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpEchoSequenceNumber.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpEchoSequenceNumberMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpEchoSequenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpEchoSequenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoSequenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpEchoSequenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoSequenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoSequenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpEchoSequenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpEchoSequenceNumber")
			}
		} else {
			intVal := otg.PatternFlowIcmpEchoSequenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpEchoSequenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
