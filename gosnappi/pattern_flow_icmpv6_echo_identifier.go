package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpv6EchoIdentifier *****
type patternFlowIcmpv6EchoIdentifier struct {
	validation
	obj              *otg.PatternFlowIcmpv6EchoIdentifier
	marshaller       marshalPatternFlowIcmpv6EchoIdentifier
	unMarshaller     unMarshalPatternFlowIcmpv6EchoIdentifier
	incrementHolder  PatternFlowIcmpv6EchoIdentifierCounter
	decrementHolder  PatternFlowIcmpv6EchoIdentifierCounter
	metricTagsHolder PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
}

func NewPatternFlowIcmpv6EchoIdentifier() PatternFlowIcmpv6EchoIdentifier {
	obj := patternFlowIcmpv6EchoIdentifier{obj: &otg.PatternFlowIcmpv6EchoIdentifier{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) msg() *otg.PatternFlowIcmpv6EchoIdentifier {
	return obj.obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) setMsg(msg *otg.PatternFlowIcmpv6EchoIdentifier) PatternFlowIcmpv6EchoIdentifier {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpv6EchoIdentifier struct {
	obj *patternFlowIcmpv6EchoIdentifier
}

type marshalPatternFlowIcmpv6EchoIdentifier interface {
	// ToProto marshals PatternFlowIcmpv6EchoIdentifier to protobuf object *otg.PatternFlowIcmpv6EchoIdentifier
	ToProto() (*otg.PatternFlowIcmpv6EchoIdentifier, error)
	// ToPbText marshals PatternFlowIcmpv6EchoIdentifier to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpv6EchoIdentifier to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpv6EchoIdentifier to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIcmpv6EchoIdentifier struct {
	obj *patternFlowIcmpv6EchoIdentifier
}

type unMarshalPatternFlowIcmpv6EchoIdentifier interface {
	// FromProto unmarshals PatternFlowIcmpv6EchoIdentifier from protobuf object *otg.PatternFlowIcmpv6EchoIdentifier
	FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifier) (PatternFlowIcmpv6EchoIdentifier, error)
	// FromPbText unmarshals PatternFlowIcmpv6EchoIdentifier from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpv6EchoIdentifier from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpv6EchoIdentifier from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpv6EchoIdentifier) Marshal() marshalPatternFlowIcmpv6EchoIdentifier {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpv6EchoIdentifier{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpv6EchoIdentifier) Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifier {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpv6EchoIdentifier{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpv6EchoIdentifier) ToProto() (*otg.PatternFlowIcmpv6EchoIdentifier, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpv6EchoIdentifier) FromProto(msg *otg.PatternFlowIcmpv6EchoIdentifier) (PatternFlowIcmpv6EchoIdentifier, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpv6EchoIdentifier) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifier) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifier) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifier) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpv6EchoIdentifier) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpv6EchoIdentifier) FromJson(value string) error {
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

func (obj *patternFlowIcmpv6EchoIdentifier) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifier) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpv6EchoIdentifier) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpv6EchoIdentifier) Clone() (PatternFlowIcmpv6EchoIdentifier, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpv6EchoIdentifier()
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

func (obj *patternFlowIcmpv6EchoIdentifier) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpv6EchoIdentifier is iCMPv6 echo identifier
type PatternFlowIcmpv6EchoIdentifier interface {
	Validation
	// msg marshals PatternFlowIcmpv6EchoIdentifier to protobuf object *otg.PatternFlowIcmpv6EchoIdentifier
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpv6EchoIdentifier
	// setMsg unmarshals PatternFlowIcmpv6EchoIdentifier from protobuf object *otg.PatternFlowIcmpv6EchoIdentifier
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpv6EchoIdentifier) PatternFlowIcmpv6EchoIdentifier
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpv6EchoIdentifier
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpv6EchoIdentifier
	// validate validates PatternFlowIcmpv6EchoIdentifier
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpv6EchoIdentifier, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpv6EchoIdentifierChoiceEnum, set in PatternFlowIcmpv6EchoIdentifier
	Choice() PatternFlowIcmpv6EchoIdentifierChoiceEnum
	// setChoice assigns PatternFlowIcmpv6EchoIdentifierChoiceEnum provided by user to PatternFlowIcmpv6EchoIdentifier
	setChoice(value PatternFlowIcmpv6EchoIdentifierChoiceEnum) PatternFlowIcmpv6EchoIdentifier
	// HasChoice checks if Choice has been set in PatternFlowIcmpv6EchoIdentifier
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpv6EchoIdentifier.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpv6EchoIdentifier
	SetValue(value uint32) PatternFlowIcmpv6EchoIdentifier
	// HasValue checks if Value has been set in PatternFlowIcmpv6EchoIdentifier
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpv6EchoIdentifier.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpv6EchoIdentifier
	SetValues(value []uint32) PatternFlowIcmpv6EchoIdentifier
	// Increment returns PatternFlowIcmpv6EchoIdentifierCounter, set in PatternFlowIcmpv6EchoIdentifier.
	// PatternFlowIcmpv6EchoIdentifierCounter is integer counter pattern
	Increment() PatternFlowIcmpv6EchoIdentifierCounter
	// SetIncrement assigns PatternFlowIcmpv6EchoIdentifierCounter provided by user to PatternFlowIcmpv6EchoIdentifier.
	// PatternFlowIcmpv6EchoIdentifierCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifier
	// HasIncrement checks if Increment has been set in PatternFlowIcmpv6EchoIdentifier
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpv6EchoIdentifierCounter, set in PatternFlowIcmpv6EchoIdentifier.
	// PatternFlowIcmpv6EchoIdentifierCounter is integer counter pattern
	Decrement() PatternFlowIcmpv6EchoIdentifierCounter
	// SetDecrement assigns PatternFlowIcmpv6EchoIdentifierCounter provided by user to PatternFlowIcmpv6EchoIdentifier.
	// PatternFlowIcmpv6EchoIdentifierCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifier
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpv6EchoIdentifier
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIterIter, set in PatternFlowIcmpv6EchoIdentifier
	MetricTags() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	setNil()
}

type PatternFlowIcmpv6EchoIdentifierChoiceEnum string

// Enum of Choice on PatternFlowIcmpv6EchoIdentifier
var PatternFlowIcmpv6EchoIdentifierChoice = struct {
	VALUE     PatternFlowIcmpv6EchoIdentifierChoiceEnum
	VALUES    PatternFlowIcmpv6EchoIdentifierChoiceEnum
	INCREMENT PatternFlowIcmpv6EchoIdentifierChoiceEnum
	DECREMENT PatternFlowIcmpv6EchoIdentifierChoiceEnum
}{
	VALUE:     PatternFlowIcmpv6EchoIdentifierChoiceEnum("value"),
	VALUES:    PatternFlowIcmpv6EchoIdentifierChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpv6EchoIdentifierChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpv6EchoIdentifierChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpv6EchoIdentifier) Choice() PatternFlowIcmpv6EchoIdentifierChoiceEnum {
	return PatternFlowIcmpv6EchoIdentifierChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpv6EchoIdentifier) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpv6EchoIdentifier) setChoice(value PatternFlowIcmpv6EchoIdentifierChoiceEnum) PatternFlowIcmpv6EchoIdentifier {
	intValue, ok := otg.PatternFlowIcmpv6EchoIdentifier_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpv6EchoIdentifierChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpv6EchoIdentifier_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpv6EchoIdentifierChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpv6EchoIdentifierChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpv6EchoIdentifierChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpv6EchoIdentifierCounter().msg()
	}

	if value == PatternFlowIcmpv6EchoIdentifierChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpv6EchoIdentifierCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifier) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpv6EchoIdentifier) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpv6EchoIdentifier object
func (obj *patternFlowIcmpv6EchoIdentifier) SetValue(value uint32) PatternFlowIcmpv6EchoIdentifier {
	obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpv6EchoIdentifier) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpv6EchoIdentifier object
func (obj *patternFlowIcmpv6EchoIdentifier) SetValues(value []uint32) PatternFlowIcmpv6EchoIdentifier {
	obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoIdentifierCounter
func (obj *patternFlowIcmpv6EchoIdentifier) Increment() PatternFlowIcmpv6EchoIdentifierCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpv6EchoIdentifierCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpv6EchoIdentifierCounter
func (obj *patternFlowIcmpv6EchoIdentifier) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpv6EchoIdentifierCounter value in the PatternFlowIcmpv6EchoIdentifier object
func (obj *patternFlowIcmpv6EchoIdentifier) SetIncrement(value PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifier {
	obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoIdentifierCounter
func (obj *patternFlowIcmpv6EchoIdentifier) Decrement() PatternFlowIcmpv6EchoIdentifierCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpv6EchoIdentifierCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpv6EchoIdentifierCounter
func (obj *patternFlowIcmpv6EchoIdentifier) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpv6EchoIdentifierCounter value in the PatternFlowIcmpv6EchoIdentifier object
func (obj *patternFlowIcmpv6EchoIdentifier) SetDecrement(value PatternFlowIcmpv6EchoIdentifierCounter) PatternFlowIcmpv6EchoIdentifier {
	obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpv6EchoIdentifierMetricTag
func (obj *patternFlowIcmpv6EchoIdentifier) MetricTags() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpv6EchoIdentifierMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter struct {
	obj                                           *patternFlowIcmpv6EchoIdentifier
	patternFlowIcmpv6EchoIdentifierMetricTagSlice []PatternFlowIcmpv6EchoIdentifierMetricTag
	fieldPtr                                      *[]*otg.PatternFlowIcmpv6EchoIdentifierMetricTag
}

func newPatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter(ptr *[]*otg.PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	return &patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter interface {
	setMsg(*patternFlowIcmpv6EchoIdentifier) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	Items() []PatternFlowIcmpv6EchoIdentifierMetricTag
	Add() PatternFlowIcmpv6EchoIdentifierMetricTag
	Append(items ...PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	Set(index int, newObj PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	Clear() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	clearHolderSlice() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
	appendHolderSlice(item PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter
}

func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) setMsg(msg *patternFlowIcmpv6EchoIdentifier) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpv6EchoIdentifierMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) Items() []PatternFlowIcmpv6EchoIdentifierMetricTag {
	return obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice
}

func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) Add() PatternFlowIcmpv6EchoIdentifierMetricTag {
	newObj := &otg.PatternFlowIcmpv6EchoIdentifierMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpv6EchoIdentifierMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) Append(items ...PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) Set(index int, newObj PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) Clear() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpv6EchoIdentifierMetricTag{}
		obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice = []PatternFlowIcmpv6EchoIdentifierMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) clearHolderSlice() PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	if len(obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice) > 0 {
		obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice = []PatternFlowIcmpv6EchoIdentifierMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter) appendHolderSlice(item PatternFlowIcmpv6EchoIdentifierMetricTag) PatternFlowIcmpv6EchoIdentifierPatternFlowIcmpv6EchoIdentifierMetricTagIter {
	obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpv6EchoIdentifierMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpv6EchoIdentifier) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpv6EchoIdentifier.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpv6EchoIdentifier.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpv6EchoIdentifierMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpv6EchoIdentifier) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIcmpv6EchoIdentifierChoice.VALUE)

	}

}
