package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4DscpPhb *****
type patternFlowIpv4DscpPhb struct {
	validation
	obj              *otg.PatternFlowIpv4DscpPhb
	marshaller       marshalPatternFlowIpv4DscpPhb
	unMarshaller     unMarshalPatternFlowIpv4DscpPhb
	incrementHolder  PatternFlowIpv4DscpPhbCounter
	decrementHolder  PatternFlowIpv4DscpPhbCounter
	metricTagsHolder PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
}

func NewPatternFlowIpv4DscpPhb() PatternFlowIpv4DscpPhb {
	obj := patternFlowIpv4DscpPhb{obj: &otg.PatternFlowIpv4DscpPhb{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4DscpPhb) msg() *otg.PatternFlowIpv4DscpPhb {
	return obj.obj
}

func (obj *patternFlowIpv4DscpPhb) setMsg(msg *otg.PatternFlowIpv4DscpPhb) PatternFlowIpv4DscpPhb {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4DscpPhb struct {
	obj *patternFlowIpv4DscpPhb
}

type marshalPatternFlowIpv4DscpPhb interface {
	// ToProto marshals PatternFlowIpv4DscpPhb to protobuf object *otg.PatternFlowIpv4DscpPhb
	ToProto() (*otg.PatternFlowIpv4DscpPhb, error)
	// ToPbText marshals PatternFlowIpv4DscpPhb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4DscpPhb to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4DscpPhb to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4DscpPhb struct {
	obj *patternFlowIpv4DscpPhb
}

type unMarshalPatternFlowIpv4DscpPhb interface {
	// FromProto unmarshals PatternFlowIpv4DscpPhb from protobuf object *otg.PatternFlowIpv4DscpPhb
	FromProto(msg *otg.PatternFlowIpv4DscpPhb) (PatternFlowIpv4DscpPhb, error)
	// FromPbText unmarshals PatternFlowIpv4DscpPhb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4DscpPhb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4DscpPhb from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4DscpPhb) Marshal() marshalPatternFlowIpv4DscpPhb {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4DscpPhb{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4DscpPhb) Unmarshal() unMarshalPatternFlowIpv4DscpPhb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4DscpPhb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4DscpPhb) ToProto() (*otg.PatternFlowIpv4DscpPhb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4DscpPhb) FromProto(msg *otg.PatternFlowIpv4DscpPhb) (PatternFlowIpv4DscpPhb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4DscpPhb) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhb) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhb) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhb) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4DscpPhb) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4DscpPhb) FromJson(value string) error {
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

func (obj *patternFlowIpv4DscpPhb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4DscpPhb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4DscpPhb) Clone() (PatternFlowIpv4DscpPhb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4DscpPhb()
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

func (obj *patternFlowIpv4DscpPhb) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4DscpPhb is per hop behavior
type PatternFlowIpv4DscpPhb interface {
	Validation
	// msg marshals PatternFlowIpv4DscpPhb to protobuf object *otg.PatternFlowIpv4DscpPhb
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4DscpPhb
	// setMsg unmarshals PatternFlowIpv4DscpPhb from protobuf object *otg.PatternFlowIpv4DscpPhb
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4DscpPhb) PatternFlowIpv4DscpPhb
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4DscpPhb
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4DscpPhb
	// validate validates PatternFlowIpv4DscpPhb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4DscpPhb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4DscpPhbChoiceEnum, set in PatternFlowIpv4DscpPhb
	Choice() PatternFlowIpv4DscpPhbChoiceEnum
	// setChoice assigns PatternFlowIpv4DscpPhbChoiceEnum provided by user to PatternFlowIpv4DscpPhb
	setChoice(value PatternFlowIpv4DscpPhbChoiceEnum) PatternFlowIpv4DscpPhb
	// HasChoice checks if Choice has been set in PatternFlowIpv4DscpPhb
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4DscpPhb.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4DscpPhb
	SetValue(value uint32) PatternFlowIpv4DscpPhb
	// HasValue checks if Value has been set in PatternFlowIpv4DscpPhb
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4DscpPhb.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4DscpPhb
	SetValues(value []uint32) PatternFlowIpv4DscpPhb
	// Increment returns PatternFlowIpv4DscpPhbCounter, set in PatternFlowIpv4DscpPhb.
	// PatternFlowIpv4DscpPhbCounter is integer counter pattern
	Increment() PatternFlowIpv4DscpPhbCounter
	// SetIncrement assigns PatternFlowIpv4DscpPhbCounter provided by user to PatternFlowIpv4DscpPhb.
	// PatternFlowIpv4DscpPhbCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhb
	// HasIncrement checks if Increment has been set in PatternFlowIpv4DscpPhb
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4DscpPhbCounter, set in PatternFlowIpv4DscpPhb.
	// PatternFlowIpv4DscpPhbCounter is integer counter pattern
	Decrement() PatternFlowIpv4DscpPhbCounter
	// SetDecrement assigns PatternFlowIpv4DscpPhbCounter provided by user to PatternFlowIpv4DscpPhb.
	// PatternFlowIpv4DscpPhbCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhb
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4DscpPhb
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIterIter, set in PatternFlowIpv4DscpPhb
	MetricTags() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	setNil()
}

type PatternFlowIpv4DscpPhbChoiceEnum string

// Enum of Choice on PatternFlowIpv4DscpPhb
var PatternFlowIpv4DscpPhbChoice = struct {
	VALUE     PatternFlowIpv4DscpPhbChoiceEnum
	VALUES    PatternFlowIpv4DscpPhbChoiceEnum
	INCREMENT PatternFlowIpv4DscpPhbChoiceEnum
	DECREMENT PatternFlowIpv4DscpPhbChoiceEnum
}{
	VALUE:     PatternFlowIpv4DscpPhbChoiceEnum("value"),
	VALUES:    PatternFlowIpv4DscpPhbChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4DscpPhbChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4DscpPhbChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4DscpPhb) Choice() PatternFlowIpv4DscpPhbChoiceEnum {
	return PatternFlowIpv4DscpPhbChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4DscpPhb) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4DscpPhb) setChoice(value PatternFlowIpv4DscpPhbChoiceEnum) PatternFlowIpv4DscpPhb {
	intValue, ok := otg.PatternFlowIpv4DscpPhb_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4DscpPhbChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4DscpPhb_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4DscpPhbChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4DscpPhbChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4DscpPhbChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4DscpPhbCounter().msg()
	}

	if value == PatternFlowIpv4DscpPhbChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4DscpPhbCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DscpPhb) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4DscpPhbChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4DscpPhb) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4DscpPhb object
func (obj *patternFlowIpv4DscpPhb) SetValue(value uint32) PatternFlowIpv4DscpPhb {
	obj.setChoice(PatternFlowIpv4DscpPhbChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4DscpPhb) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4DscpPhb object
func (obj *patternFlowIpv4DscpPhb) SetValues(value []uint32) PatternFlowIpv4DscpPhb {
	obj.setChoice(PatternFlowIpv4DscpPhbChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4DscpPhbCounter
func (obj *patternFlowIpv4DscpPhb) Increment() PatternFlowIpv4DscpPhbCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4DscpPhbChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4DscpPhbCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4DscpPhbCounter
func (obj *patternFlowIpv4DscpPhb) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4DscpPhbCounter value in the PatternFlowIpv4DscpPhb object
func (obj *patternFlowIpv4DscpPhb) SetIncrement(value PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhb {
	obj.setChoice(PatternFlowIpv4DscpPhbChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4DscpPhbCounter
func (obj *patternFlowIpv4DscpPhb) Decrement() PatternFlowIpv4DscpPhbCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4DscpPhbChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4DscpPhbCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4DscpPhbCounter
func (obj *patternFlowIpv4DscpPhb) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4DscpPhbCounter value in the PatternFlowIpv4DscpPhb object
func (obj *patternFlowIpv4DscpPhb) SetDecrement(value PatternFlowIpv4DscpPhbCounter) PatternFlowIpv4DscpPhb {
	obj.setChoice(PatternFlowIpv4DscpPhbChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4DscpPhbMetricTag
func (obj *patternFlowIpv4DscpPhb) MetricTags() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4DscpPhbMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter struct {
	obj                                  *patternFlowIpv4DscpPhb
	patternFlowIpv4DscpPhbMetricTagSlice []PatternFlowIpv4DscpPhbMetricTag
	fieldPtr                             *[]*otg.PatternFlowIpv4DscpPhbMetricTag
}

func newPatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter(ptr *[]*otg.PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	return &patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter interface {
	setMsg(*patternFlowIpv4DscpPhb) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	Items() []PatternFlowIpv4DscpPhbMetricTag
	Add() PatternFlowIpv4DscpPhbMetricTag
	Append(items ...PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	Set(index int, newObj PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	Clear() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	clearHolderSlice() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
	appendHolderSlice(item PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter
}

func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) setMsg(msg *patternFlowIpv4DscpPhb) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4DscpPhbMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) Items() []PatternFlowIpv4DscpPhbMetricTag {
	return obj.patternFlowIpv4DscpPhbMetricTagSlice
}

func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) Add() PatternFlowIpv4DscpPhbMetricTag {
	newObj := &otg.PatternFlowIpv4DscpPhbMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4DscpPhbMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4DscpPhbMetricTagSlice = append(obj.patternFlowIpv4DscpPhbMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) Append(items ...PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4DscpPhbMetricTagSlice = append(obj.patternFlowIpv4DscpPhbMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) Set(index int, newObj PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4DscpPhbMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) Clear() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4DscpPhbMetricTag{}
		obj.patternFlowIpv4DscpPhbMetricTagSlice = []PatternFlowIpv4DscpPhbMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) clearHolderSlice() PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	if len(obj.patternFlowIpv4DscpPhbMetricTagSlice) > 0 {
		obj.patternFlowIpv4DscpPhbMetricTagSlice = []PatternFlowIpv4DscpPhbMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter) appendHolderSlice(item PatternFlowIpv4DscpPhbMetricTag) PatternFlowIpv4DscpPhbPatternFlowIpv4DscpPhbMetricTagIter {
	obj.patternFlowIpv4DscpPhbMetricTagSlice = append(obj.patternFlowIpv4DscpPhbMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4DscpPhb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4DscpPhb.Value <= 63 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 63 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4DscpPhb.Values <= 63 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4DscpPhbMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4DscpPhb) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4DscpPhbChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpPhbChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4DscpPhbChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpPhbChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4DscpPhbChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4DscpPhbChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4DscpPhb")
			}
		} else {
			intVal := otg.PatternFlowIpv4DscpPhb_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4DscpPhb_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
