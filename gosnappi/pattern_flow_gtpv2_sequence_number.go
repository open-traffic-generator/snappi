package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2SequenceNumber *****
type patternFlowGtpv2SequenceNumber struct {
	validation
	obj              *otg.PatternFlowGtpv2SequenceNumber
	marshaller       marshalPatternFlowGtpv2SequenceNumber
	unMarshaller     unMarshalPatternFlowGtpv2SequenceNumber
	incrementHolder  PatternFlowGtpv2SequenceNumberCounter
	decrementHolder  PatternFlowGtpv2SequenceNumberCounter
	metricTagsHolder PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
}

func NewPatternFlowGtpv2SequenceNumber() PatternFlowGtpv2SequenceNumber {
	obj := patternFlowGtpv2SequenceNumber{obj: &otg.PatternFlowGtpv2SequenceNumber{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2SequenceNumber) msg() *otg.PatternFlowGtpv2SequenceNumber {
	return obj.obj
}

func (obj *patternFlowGtpv2SequenceNumber) setMsg(msg *otg.PatternFlowGtpv2SequenceNumber) PatternFlowGtpv2SequenceNumber {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2SequenceNumber struct {
	obj *patternFlowGtpv2SequenceNumber
}

type marshalPatternFlowGtpv2SequenceNumber interface {
	// ToProto marshals PatternFlowGtpv2SequenceNumber to protobuf object *otg.PatternFlowGtpv2SequenceNumber
	ToProto() (*otg.PatternFlowGtpv2SequenceNumber, error)
	// ToPbText marshals PatternFlowGtpv2SequenceNumber to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2SequenceNumber to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2SequenceNumber to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv2SequenceNumber struct {
	obj *patternFlowGtpv2SequenceNumber
}

type unMarshalPatternFlowGtpv2SequenceNumber interface {
	// FromProto unmarshals PatternFlowGtpv2SequenceNumber from protobuf object *otg.PatternFlowGtpv2SequenceNumber
	FromProto(msg *otg.PatternFlowGtpv2SequenceNumber) (PatternFlowGtpv2SequenceNumber, error)
	// FromPbText unmarshals PatternFlowGtpv2SequenceNumber from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2SequenceNumber from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2SequenceNumber from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2SequenceNumber) Marshal() marshalPatternFlowGtpv2SequenceNumber {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2SequenceNumber{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2SequenceNumber) Unmarshal() unMarshalPatternFlowGtpv2SequenceNumber {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2SequenceNumber{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2SequenceNumber) ToProto() (*otg.PatternFlowGtpv2SequenceNumber, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2SequenceNumber) FromProto(msg *otg.PatternFlowGtpv2SequenceNumber) (PatternFlowGtpv2SequenceNumber, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2SequenceNumber) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumber) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumber) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumber) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2SequenceNumber) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2SequenceNumber) FromJson(value string) error {
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

func (obj *patternFlowGtpv2SequenceNumber) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumber) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2SequenceNumber) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2SequenceNumber) Clone() (PatternFlowGtpv2SequenceNumber, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2SequenceNumber()
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

func (obj *patternFlowGtpv2SequenceNumber) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2SequenceNumber is the sequence number
type PatternFlowGtpv2SequenceNumber interface {
	Validation
	// msg marshals PatternFlowGtpv2SequenceNumber to protobuf object *otg.PatternFlowGtpv2SequenceNumber
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2SequenceNumber
	// setMsg unmarshals PatternFlowGtpv2SequenceNumber from protobuf object *otg.PatternFlowGtpv2SequenceNumber
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2SequenceNumber) PatternFlowGtpv2SequenceNumber
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2SequenceNumber
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2SequenceNumber
	// validate validates PatternFlowGtpv2SequenceNumber
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2SequenceNumber, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2SequenceNumberChoiceEnum, set in PatternFlowGtpv2SequenceNumber
	Choice() PatternFlowGtpv2SequenceNumberChoiceEnum
	// setChoice assigns PatternFlowGtpv2SequenceNumberChoiceEnum provided by user to PatternFlowGtpv2SequenceNumber
	setChoice(value PatternFlowGtpv2SequenceNumberChoiceEnum) PatternFlowGtpv2SequenceNumber
	// HasChoice checks if Choice has been set in PatternFlowGtpv2SequenceNumber
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2SequenceNumber.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2SequenceNumber
	SetValue(value uint32) PatternFlowGtpv2SequenceNumber
	// HasValue checks if Value has been set in PatternFlowGtpv2SequenceNumber
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2SequenceNumber.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2SequenceNumber
	SetValues(value []uint32) PatternFlowGtpv2SequenceNumber
	// Increment returns PatternFlowGtpv2SequenceNumberCounter, set in PatternFlowGtpv2SequenceNumber.
	// PatternFlowGtpv2SequenceNumberCounter is integer counter pattern
	Increment() PatternFlowGtpv2SequenceNumberCounter
	// SetIncrement assigns PatternFlowGtpv2SequenceNumberCounter provided by user to PatternFlowGtpv2SequenceNumber.
	// PatternFlowGtpv2SequenceNumberCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumber
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2SequenceNumber
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2SequenceNumberCounter, set in PatternFlowGtpv2SequenceNumber.
	// PatternFlowGtpv2SequenceNumberCounter is integer counter pattern
	Decrement() PatternFlowGtpv2SequenceNumberCounter
	// SetDecrement assigns PatternFlowGtpv2SequenceNumberCounter provided by user to PatternFlowGtpv2SequenceNumber.
	// PatternFlowGtpv2SequenceNumberCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumber
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2SequenceNumber
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIterIter, set in PatternFlowGtpv2SequenceNumber
	MetricTags() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	setNil()
}

type PatternFlowGtpv2SequenceNumberChoiceEnum string

// Enum of Choice on PatternFlowGtpv2SequenceNumber
var PatternFlowGtpv2SequenceNumberChoice = struct {
	VALUE     PatternFlowGtpv2SequenceNumberChoiceEnum
	VALUES    PatternFlowGtpv2SequenceNumberChoiceEnum
	INCREMENT PatternFlowGtpv2SequenceNumberChoiceEnum
	DECREMENT PatternFlowGtpv2SequenceNumberChoiceEnum
}{
	VALUE:     PatternFlowGtpv2SequenceNumberChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2SequenceNumberChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2SequenceNumberChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2SequenceNumberChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2SequenceNumber) Choice() PatternFlowGtpv2SequenceNumberChoiceEnum {
	return PatternFlowGtpv2SequenceNumberChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2SequenceNumber) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2SequenceNumber) setChoice(value PatternFlowGtpv2SequenceNumberChoiceEnum) PatternFlowGtpv2SequenceNumber {
	intValue, ok := otg.PatternFlowGtpv2SequenceNumber_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2SequenceNumberChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2SequenceNumber_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2SequenceNumberChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2SequenceNumberChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2SequenceNumberChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2SequenceNumberCounter().msg()
	}

	if value == PatternFlowGtpv2SequenceNumberChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2SequenceNumberCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2SequenceNumber) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2SequenceNumber) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2SequenceNumber object
func (obj *patternFlowGtpv2SequenceNumber) SetValue(value uint32) PatternFlowGtpv2SequenceNumber {
	obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2SequenceNumber) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2SequenceNumber object
func (obj *patternFlowGtpv2SequenceNumber) SetValues(value []uint32) PatternFlowGtpv2SequenceNumber {
	obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2SequenceNumberCounter
func (obj *patternFlowGtpv2SequenceNumber) Increment() PatternFlowGtpv2SequenceNumberCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2SequenceNumberCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2SequenceNumberCounter
func (obj *patternFlowGtpv2SequenceNumber) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2SequenceNumberCounter value in the PatternFlowGtpv2SequenceNumber object
func (obj *patternFlowGtpv2SequenceNumber) SetIncrement(value PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumber {
	obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2SequenceNumberCounter
func (obj *patternFlowGtpv2SequenceNumber) Decrement() PatternFlowGtpv2SequenceNumberCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2SequenceNumberCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2SequenceNumberCounter
func (obj *patternFlowGtpv2SequenceNumber) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2SequenceNumberCounter value in the PatternFlowGtpv2SequenceNumber object
func (obj *patternFlowGtpv2SequenceNumber) SetDecrement(value PatternFlowGtpv2SequenceNumberCounter) PatternFlowGtpv2SequenceNumber {
	obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2SequenceNumberMetricTag
func (obj *patternFlowGtpv2SequenceNumber) MetricTags() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2SequenceNumberMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter struct {
	obj                                          *patternFlowGtpv2SequenceNumber
	patternFlowGtpv2SequenceNumberMetricTagSlice []PatternFlowGtpv2SequenceNumberMetricTag
	fieldPtr                                     *[]*otg.PatternFlowGtpv2SequenceNumberMetricTag
}

func newPatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter(ptr *[]*otg.PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	return &patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter interface {
	setMsg(*patternFlowGtpv2SequenceNumber) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	Items() []PatternFlowGtpv2SequenceNumberMetricTag
	Add() PatternFlowGtpv2SequenceNumberMetricTag
	Append(items ...PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	Set(index int, newObj PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	Clear() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	clearHolderSlice() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter
}

func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) setMsg(msg *patternFlowGtpv2SequenceNumber) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2SequenceNumberMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) Items() []PatternFlowGtpv2SequenceNumberMetricTag {
	return obj.patternFlowGtpv2SequenceNumberMetricTagSlice
}

func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) Add() PatternFlowGtpv2SequenceNumberMetricTag {
	newObj := &otg.PatternFlowGtpv2SequenceNumberMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2SequenceNumberMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2SequenceNumberMetricTagSlice = append(obj.patternFlowGtpv2SequenceNumberMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) Append(items ...PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2SequenceNumberMetricTagSlice = append(obj.patternFlowGtpv2SequenceNumberMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) Set(index int, newObj PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2SequenceNumberMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) Clear() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2SequenceNumberMetricTag{}
		obj.patternFlowGtpv2SequenceNumberMetricTagSlice = []PatternFlowGtpv2SequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) clearHolderSlice() PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	if len(obj.patternFlowGtpv2SequenceNumberMetricTagSlice) > 0 {
		obj.patternFlowGtpv2SequenceNumberMetricTagSlice = []PatternFlowGtpv2SequenceNumberMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter) appendHolderSlice(item PatternFlowGtpv2SequenceNumberMetricTag) PatternFlowGtpv2SequenceNumberPatternFlowGtpv2SequenceNumberMetricTagIter {
	obj.patternFlowGtpv2SequenceNumberMetricTagSlice = append(obj.patternFlowGtpv2SequenceNumberMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2SequenceNumber) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpv2SequenceNumber.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 16777215 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowGtpv2SequenceNumber.Values <= 16777215 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2SequenceNumberMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2SequenceNumber) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2SequenceNumberChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2SequenceNumberChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2SequenceNumberChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2SequenceNumberChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2SequenceNumberChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2SequenceNumberChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2SequenceNumber")
			}
		} else {
			intVal := otg.PatternFlowGtpv2SequenceNumber_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2SequenceNumber_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
