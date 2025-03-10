package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIcmpEchoIdentifier *****
type patternFlowIcmpEchoIdentifier struct {
	validation
	obj              *otg.PatternFlowIcmpEchoIdentifier
	marshaller       marshalPatternFlowIcmpEchoIdentifier
	unMarshaller     unMarshalPatternFlowIcmpEchoIdentifier
	incrementHolder  PatternFlowIcmpEchoIdentifierCounter
	decrementHolder  PatternFlowIcmpEchoIdentifierCounter
	metricTagsHolder PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
}

func NewPatternFlowIcmpEchoIdentifier() PatternFlowIcmpEchoIdentifier {
	obj := patternFlowIcmpEchoIdentifier{obj: &otg.PatternFlowIcmpEchoIdentifier{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIcmpEchoIdentifier) msg() *otg.PatternFlowIcmpEchoIdentifier {
	return obj.obj
}

func (obj *patternFlowIcmpEchoIdentifier) setMsg(msg *otg.PatternFlowIcmpEchoIdentifier) PatternFlowIcmpEchoIdentifier {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIcmpEchoIdentifier struct {
	obj *patternFlowIcmpEchoIdentifier
}

type marshalPatternFlowIcmpEchoIdentifier interface {
	// ToProto marshals PatternFlowIcmpEchoIdentifier to protobuf object *otg.PatternFlowIcmpEchoIdentifier
	ToProto() (*otg.PatternFlowIcmpEchoIdentifier, error)
	// ToPbText marshals PatternFlowIcmpEchoIdentifier to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIcmpEchoIdentifier to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIcmpEchoIdentifier to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIcmpEchoIdentifier to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIcmpEchoIdentifier struct {
	obj *patternFlowIcmpEchoIdentifier
}

type unMarshalPatternFlowIcmpEchoIdentifier interface {
	// FromProto unmarshals PatternFlowIcmpEchoIdentifier from protobuf object *otg.PatternFlowIcmpEchoIdentifier
	FromProto(msg *otg.PatternFlowIcmpEchoIdentifier) (PatternFlowIcmpEchoIdentifier, error)
	// FromPbText unmarshals PatternFlowIcmpEchoIdentifier from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIcmpEchoIdentifier from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIcmpEchoIdentifier from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIcmpEchoIdentifier) Marshal() marshalPatternFlowIcmpEchoIdentifier {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIcmpEchoIdentifier{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIcmpEchoIdentifier) Unmarshal() unMarshalPatternFlowIcmpEchoIdentifier {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIcmpEchoIdentifier{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIcmpEchoIdentifier) ToProto() (*otg.PatternFlowIcmpEchoIdentifier, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIcmpEchoIdentifier) FromProto(msg *otg.PatternFlowIcmpEchoIdentifier) (PatternFlowIcmpEchoIdentifier, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIcmpEchoIdentifier) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifier) FromPbText(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifier) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifier) FromYaml(value string) error {
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

func (m *marshalpatternFlowIcmpEchoIdentifier) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIcmpEchoIdentifier) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIcmpEchoIdentifier) FromJson(value string) error {
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

func (obj *patternFlowIcmpEchoIdentifier) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifier) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIcmpEchoIdentifier) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIcmpEchoIdentifier) Clone() (PatternFlowIcmpEchoIdentifier, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIcmpEchoIdentifier()
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

func (obj *patternFlowIcmpEchoIdentifier) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIcmpEchoIdentifier is iCMP identifier
type PatternFlowIcmpEchoIdentifier interface {
	Validation
	// msg marshals PatternFlowIcmpEchoIdentifier to protobuf object *otg.PatternFlowIcmpEchoIdentifier
	// and doesn't set defaults
	msg() *otg.PatternFlowIcmpEchoIdentifier
	// setMsg unmarshals PatternFlowIcmpEchoIdentifier from protobuf object *otg.PatternFlowIcmpEchoIdentifier
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIcmpEchoIdentifier) PatternFlowIcmpEchoIdentifier
	// provides marshal interface
	Marshal() marshalPatternFlowIcmpEchoIdentifier
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIcmpEchoIdentifier
	// validate validates PatternFlowIcmpEchoIdentifier
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIcmpEchoIdentifier, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIcmpEchoIdentifierChoiceEnum, set in PatternFlowIcmpEchoIdentifier
	Choice() PatternFlowIcmpEchoIdentifierChoiceEnum
	// setChoice assigns PatternFlowIcmpEchoIdentifierChoiceEnum provided by user to PatternFlowIcmpEchoIdentifier
	setChoice(value PatternFlowIcmpEchoIdentifierChoiceEnum) PatternFlowIcmpEchoIdentifier
	// HasChoice checks if Choice has been set in PatternFlowIcmpEchoIdentifier
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIcmpEchoIdentifier.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIcmpEchoIdentifier
	SetValue(value uint32) PatternFlowIcmpEchoIdentifier
	// HasValue checks if Value has been set in PatternFlowIcmpEchoIdentifier
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIcmpEchoIdentifier.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIcmpEchoIdentifier
	SetValues(value []uint32) PatternFlowIcmpEchoIdentifier
	// Increment returns PatternFlowIcmpEchoIdentifierCounter, set in PatternFlowIcmpEchoIdentifier.
	// PatternFlowIcmpEchoIdentifierCounter is integer counter pattern
	Increment() PatternFlowIcmpEchoIdentifierCounter
	// SetIncrement assigns PatternFlowIcmpEchoIdentifierCounter provided by user to PatternFlowIcmpEchoIdentifier.
	// PatternFlowIcmpEchoIdentifierCounter is integer counter pattern
	SetIncrement(value PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifier
	// HasIncrement checks if Increment has been set in PatternFlowIcmpEchoIdentifier
	HasIncrement() bool
	// Decrement returns PatternFlowIcmpEchoIdentifierCounter, set in PatternFlowIcmpEchoIdentifier.
	// PatternFlowIcmpEchoIdentifierCounter is integer counter pattern
	Decrement() PatternFlowIcmpEchoIdentifierCounter
	// SetDecrement assigns PatternFlowIcmpEchoIdentifierCounter provided by user to PatternFlowIcmpEchoIdentifier.
	// PatternFlowIcmpEchoIdentifierCounter is integer counter pattern
	SetDecrement(value PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifier
	// HasDecrement checks if Decrement has been set in PatternFlowIcmpEchoIdentifier
	HasDecrement() bool
	// MetricTags returns PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIterIter, set in PatternFlowIcmpEchoIdentifier
	MetricTags() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	setNil()
}

type PatternFlowIcmpEchoIdentifierChoiceEnum string

// Enum of Choice on PatternFlowIcmpEchoIdentifier
var PatternFlowIcmpEchoIdentifierChoice = struct {
	VALUE     PatternFlowIcmpEchoIdentifierChoiceEnum
	VALUES    PatternFlowIcmpEchoIdentifierChoiceEnum
	INCREMENT PatternFlowIcmpEchoIdentifierChoiceEnum
	DECREMENT PatternFlowIcmpEchoIdentifierChoiceEnum
}{
	VALUE:     PatternFlowIcmpEchoIdentifierChoiceEnum("value"),
	VALUES:    PatternFlowIcmpEchoIdentifierChoiceEnum("values"),
	INCREMENT: PatternFlowIcmpEchoIdentifierChoiceEnum("increment"),
	DECREMENT: PatternFlowIcmpEchoIdentifierChoiceEnum("decrement"),
}

func (obj *patternFlowIcmpEchoIdentifier) Choice() PatternFlowIcmpEchoIdentifierChoiceEnum {
	return PatternFlowIcmpEchoIdentifierChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIcmpEchoIdentifier) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIcmpEchoIdentifier) setChoice(value PatternFlowIcmpEchoIdentifierChoiceEnum) PatternFlowIcmpEchoIdentifier {
	intValue, ok := otg.PatternFlowIcmpEchoIdentifier_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIcmpEchoIdentifierChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIcmpEchoIdentifier_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIcmpEchoIdentifierChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIcmpEchoIdentifierChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIcmpEchoIdentifierChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIcmpEchoIdentifierCounter().msg()
	}

	if value == PatternFlowIcmpEchoIdentifierChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIcmpEchoIdentifierCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoIdentifier) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIcmpEchoIdentifier) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIcmpEchoIdentifier object
func (obj *patternFlowIcmpEchoIdentifier) SetValue(value uint32) PatternFlowIcmpEchoIdentifier {
	obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIcmpEchoIdentifier) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIcmpEchoIdentifier object
func (obj *patternFlowIcmpEchoIdentifier) SetValues(value []uint32) PatternFlowIcmpEchoIdentifier {
	obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoIdentifierCounter
func (obj *patternFlowIcmpEchoIdentifier) Increment() PatternFlowIcmpEchoIdentifierCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIcmpEchoIdentifierCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIcmpEchoIdentifierCounter
func (obj *patternFlowIcmpEchoIdentifier) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIcmpEchoIdentifierCounter value in the PatternFlowIcmpEchoIdentifier object
func (obj *patternFlowIcmpEchoIdentifier) SetIncrement(value PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifier {
	obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoIdentifierCounter
func (obj *patternFlowIcmpEchoIdentifier) Decrement() PatternFlowIcmpEchoIdentifierCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIcmpEchoIdentifierCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIcmpEchoIdentifierCounter
func (obj *patternFlowIcmpEchoIdentifier) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIcmpEchoIdentifierCounter value in the PatternFlowIcmpEchoIdentifier object
func (obj *patternFlowIcmpEchoIdentifier) SetDecrement(value PatternFlowIcmpEchoIdentifierCounter) PatternFlowIcmpEchoIdentifier {
	obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIcmpEchoIdentifierMetricTag
func (obj *patternFlowIcmpEchoIdentifier) MetricTags() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIcmpEchoIdentifierMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter struct {
	obj                                         *patternFlowIcmpEchoIdentifier
	patternFlowIcmpEchoIdentifierMetricTagSlice []PatternFlowIcmpEchoIdentifierMetricTag
	fieldPtr                                    *[]*otg.PatternFlowIcmpEchoIdentifierMetricTag
}

func newPatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter(ptr *[]*otg.PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	return &patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter interface {
	setMsg(*patternFlowIcmpEchoIdentifier) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	Items() []PatternFlowIcmpEchoIdentifierMetricTag
	Add() PatternFlowIcmpEchoIdentifierMetricTag
	Append(items ...PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	Set(index int, newObj PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	Clear() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	clearHolderSlice() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
	appendHolderSlice(item PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter
}

func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) setMsg(msg *patternFlowIcmpEchoIdentifier) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIcmpEchoIdentifierMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) Items() []PatternFlowIcmpEchoIdentifierMetricTag {
	return obj.patternFlowIcmpEchoIdentifierMetricTagSlice
}

func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) Add() PatternFlowIcmpEchoIdentifierMetricTag {
	newObj := &otg.PatternFlowIcmpEchoIdentifierMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIcmpEchoIdentifierMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIcmpEchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpEchoIdentifierMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) Append(items ...PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIcmpEchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpEchoIdentifierMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) Set(index int, newObj PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIcmpEchoIdentifierMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) Clear() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIcmpEchoIdentifierMetricTag{}
		obj.patternFlowIcmpEchoIdentifierMetricTagSlice = []PatternFlowIcmpEchoIdentifierMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) clearHolderSlice() PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	if len(obj.patternFlowIcmpEchoIdentifierMetricTagSlice) > 0 {
		obj.patternFlowIcmpEchoIdentifierMetricTagSlice = []PatternFlowIcmpEchoIdentifierMetricTag{}
	}
	return obj
}
func (obj *patternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter) appendHolderSlice(item PatternFlowIcmpEchoIdentifierMetricTag) PatternFlowIcmpEchoIdentifierPatternFlowIcmpEchoIdentifierMetricTagIter {
	obj.patternFlowIcmpEchoIdentifierMetricTagSlice = append(obj.patternFlowIcmpEchoIdentifierMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIcmpEchoIdentifier) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIcmpEchoIdentifier.Value <= 65535 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 65535 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIcmpEchoIdentifier.Values <= 65535 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIcmpEchoIdentifierMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIcmpEchoIdentifier) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIcmpEchoIdentifierChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoIdentifierChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIcmpEchoIdentifierChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoIdentifierChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIcmpEchoIdentifierChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIcmpEchoIdentifierChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIcmpEchoIdentifier")
			}
		} else {
			intVal := otg.PatternFlowIcmpEchoIdentifier_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIcmpEchoIdentifier_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
