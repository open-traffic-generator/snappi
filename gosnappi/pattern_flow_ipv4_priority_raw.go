package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4PriorityRaw *****
type patternFlowIpv4PriorityRaw struct {
	validation
	obj              *otg.PatternFlowIpv4PriorityRaw
	marshaller       marshalPatternFlowIpv4PriorityRaw
	unMarshaller     unMarshalPatternFlowIpv4PriorityRaw
	incrementHolder  PatternFlowIpv4PriorityRawCounter
	decrementHolder  PatternFlowIpv4PriorityRawCounter
	metricTagsHolder PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
}

func NewPatternFlowIpv4PriorityRaw() PatternFlowIpv4PriorityRaw {
	obj := patternFlowIpv4PriorityRaw{obj: &otg.PatternFlowIpv4PriorityRaw{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4PriorityRaw) msg() *otg.PatternFlowIpv4PriorityRaw {
	return obj.obj
}

func (obj *patternFlowIpv4PriorityRaw) setMsg(msg *otg.PatternFlowIpv4PriorityRaw) PatternFlowIpv4PriorityRaw {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4PriorityRaw struct {
	obj *patternFlowIpv4PriorityRaw
}

type marshalPatternFlowIpv4PriorityRaw interface {
	// ToProto marshals PatternFlowIpv4PriorityRaw to protobuf object *otg.PatternFlowIpv4PriorityRaw
	ToProto() (*otg.PatternFlowIpv4PriorityRaw, error)
	// ToPbText marshals PatternFlowIpv4PriorityRaw to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4PriorityRaw to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4PriorityRaw to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowIpv4PriorityRaw to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowIpv4PriorityRaw struct {
	obj *patternFlowIpv4PriorityRaw
}

type unMarshalPatternFlowIpv4PriorityRaw interface {
	// FromProto unmarshals PatternFlowIpv4PriorityRaw from protobuf object *otg.PatternFlowIpv4PriorityRaw
	FromProto(msg *otg.PatternFlowIpv4PriorityRaw) (PatternFlowIpv4PriorityRaw, error)
	// FromPbText unmarshals PatternFlowIpv4PriorityRaw from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4PriorityRaw from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4PriorityRaw from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4PriorityRaw) Marshal() marshalPatternFlowIpv4PriorityRaw {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4PriorityRaw{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4PriorityRaw) Unmarshal() unMarshalPatternFlowIpv4PriorityRaw {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4PriorityRaw{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4PriorityRaw) ToProto() (*otg.PatternFlowIpv4PriorityRaw, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4PriorityRaw) FromProto(msg *otg.PatternFlowIpv4PriorityRaw) (PatternFlowIpv4PriorityRaw, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4PriorityRaw) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRaw) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRaw) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRaw) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4PriorityRaw) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowIpv4PriorityRaw) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4PriorityRaw) FromJson(value string) error {
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

func (obj *patternFlowIpv4PriorityRaw) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRaw) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4PriorityRaw) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4PriorityRaw) Clone() (PatternFlowIpv4PriorityRaw, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4PriorityRaw()
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

func (obj *patternFlowIpv4PriorityRaw) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4PriorityRaw is raw priority
type PatternFlowIpv4PriorityRaw interface {
	Validation
	// msg marshals PatternFlowIpv4PriorityRaw to protobuf object *otg.PatternFlowIpv4PriorityRaw
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4PriorityRaw
	// setMsg unmarshals PatternFlowIpv4PriorityRaw from protobuf object *otg.PatternFlowIpv4PriorityRaw
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4PriorityRaw) PatternFlowIpv4PriorityRaw
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4PriorityRaw
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4PriorityRaw
	// validate validates PatternFlowIpv4PriorityRaw
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4PriorityRaw, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4PriorityRawChoiceEnum, set in PatternFlowIpv4PriorityRaw
	Choice() PatternFlowIpv4PriorityRawChoiceEnum
	// setChoice assigns PatternFlowIpv4PriorityRawChoiceEnum provided by user to PatternFlowIpv4PriorityRaw
	setChoice(value PatternFlowIpv4PriorityRawChoiceEnum) PatternFlowIpv4PriorityRaw
	// HasChoice checks if Choice has been set in PatternFlowIpv4PriorityRaw
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4PriorityRaw.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4PriorityRaw
	SetValue(value uint32) PatternFlowIpv4PriorityRaw
	// HasValue checks if Value has been set in PatternFlowIpv4PriorityRaw
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4PriorityRaw.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4PriorityRaw
	SetValues(value []uint32) PatternFlowIpv4PriorityRaw
	// Increment returns PatternFlowIpv4PriorityRawCounter, set in PatternFlowIpv4PriorityRaw.
	// PatternFlowIpv4PriorityRawCounter is integer counter pattern
	Increment() PatternFlowIpv4PriorityRawCounter
	// SetIncrement assigns PatternFlowIpv4PriorityRawCounter provided by user to PatternFlowIpv4PriorityRaw.
	// PatternFlowIpv4PriorityRawCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRaw
	// HasIncrement checks if Increment has been set in PatternFlowIpv4PriorityRaw
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4PriorityRawCounter, set in PatternFlowIpv4PriorityRaw.
	// PatternFlowIpv4PriorityRawCounter is integer counter pattern
	Decrement() PatternFlowIpv4PriorityRawCounter
	// SetDecrement assigns PatternFlowIpv4PriorityRawCounter provided by user to PatternFlowIpv4PriorityRaw.
	// PatternFlowIpv4PriorityRawCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRaw
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4PriorityRaw
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIterIter, set in PatternFlowIpv4PriorityRaw
	MetricTags() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	setNil()
}

type PatternFlowIpv4PriorityRawChoiceEnum string

// Enum of Choice on PatternFlowIpv4PriorityRaw
var PatternFlowIpv4PriorityRawChoice = struct {
	VALUE     PatternFlowIpv4PriorityRawChoiceEnum
	VALUES    PatternFlowIpv4PriorityRawChoiceEnum
	INCREMENT PatternFlowIpv4PriorityRawChoiceEnum
	DECREMENT PatternFlowIpv4PriorityRawChoiceEnum
}{
	VALUE:     PatternFlowIpv4PriorityRawChoiceEnum("value"),
	VALUES:    PatternFlowIpv4PriorityRawChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4PriorityRawChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4PriorityRawChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4PriorityRaw) Choice() PatternFlowIpv4PriorityRawChoiceEnum {
	return PatternFlowIpv4PriorityRawChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4PriorityRaw) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4PriorityRaw) setChoice(value PatternFlowIpv4PriorityRawChoiceEnum) PatternFlowIpv4PriorityRaw {
	intValue, ok := otg.PatternFlowIpv4PriorityRaw_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4PriorityRawChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4PriorityRaw_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4PriorityRawChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4PriorityRawChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4PriorityRawChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4PriorityRawCounter().msg()
	}

	if value == PatternFlowIpv4PriorityRawChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4PriorityRawCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4PriorityRaw) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4PriorityRawChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4PriorityRaw) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4PriorityRaw object
func (obj *patternFlowIpv4PriorityRaw) SetValue(value uint32) PatternFlowIpv4PriorityRaw {
	obj.setChoice(PatternFlowIpv4PriorityRawChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4PriorityRaw) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4PriorityRaw object
func (obj *patternFlowIpv4PriorityRaw) SetValues(value []uint32) PatternFlowIpv4PriorityRaw {
	obj.setChoice(PatternFlowIpv4PriorityRawChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4PriorityRawCounter
func (obj *patternFlowIpv4PriorityRaw) Increment() PatternFlowIpv4PriorityRawCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4PriorityRawChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4PriorityRawCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4PriorityRawCounter
func (obj *patternFlowIpv4PriorityRaw) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4PriorityRawCounter value in the PatternFlowIpv4PriorityRaw object
func (obj *patternFlowIpv4PriorityRaw) SetIncrement(value PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRaw {
	obj.setChoice(PatternFlowIpv4PriorityRawChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4PriorityRawCounter
func (obj *patternFlowIpv4PriorityRaw) Decrement() PatternFlowIpv4PriorityRawCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4PriorityRawChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4PriorityRawCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4PriorityRawCounter
func (obj *patternFlowIpv4PriorityRaw) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4PriorityRawCounter value in the PatternFlowIpv4PriorityRaw object
func (obj *patternFlowIpv4PriorityRaw) SetDecrement(value PatternFlowIpv4PriorityRawCounter) PatternFlowIpv4PriorityRaw {
	obj.setChoice(PatternFlowIpv4PriorityRawChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4PriorityRawMetricTag
func (obj *patternFlowIpv4PriorityRaw) MetricTags() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4PriorityRawMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter struct {
	obj                                      *patternFlowIpv4PriorityRaw
	patternFlowIpv4PriorityRawMetricTagSlice []PatternFlowIpv4PriorityRawMetricTag
	fieldPtr                                 *[]*otg.PatternFlowIpv4PriorityRawMetricTag
}

func newPatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter(ptr *[]*otg.PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	return &patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter interface {
	setMsg(*patternFlowIpv4PriorityRaw) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	Items() []PatternFlowIpv4PriorityRawMetricTag
	Add() PatternFlowIpv4PriorityRawMetricTag
	Append(items ...PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	Set(index int, newObj PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	Clear() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	clearHolderSlice() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
	appendHolderSlice(item PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter
}

func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) setMsg(msg *patternFlowIpv4PriorityRaw) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4PriorityRawMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) Items() []PatternFlowIpv4PriorityRawMetricTag {
	return obj.patternFlowIpv4PriorityRawMetricTagSlice
}

func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) Add() PatternFlowIpv4PriorityRawMetricTag {
	newObj := &otg.PatternFlowIpv4PriorityRawMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4PriorityRawMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4PriorityRawMetricTagSlice = append(obj.patternFlowIpv4PriorityRawMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) Append(items ...PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4PriorityRawMetricTagSlice = append(obj.patternFlowIpv4PriorityRawMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) Set(index int, newObj PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4PriorityRawMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) Clear() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4PriorityRawMetricTag{}
		obj.patternFlowIpv4PriorityRawMetricTagSlice = []PatternFlowIpv4PriorityRawMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) clearHolderSlice() PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	if len(obj.patternFlowIpv4PriorityRawMetricTagSlice) > 0 {
		obj.patternFlowIpv4PriorityRawMetricTagSlice = []PatternFlowIpv4PriorityRawMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter) appendHolderSlice(item PatternFlowIpv4PriorityRawMetricTag) PatternFlowIpv4PriorityRawPatternFlowIpv4PriorityRawMetricTagIter {
	obj.patternFlowIpv4PriorityRawMetricTagSlice = append(obj.patternFlowIpv4PriorityRawMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4PriorityRaw) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4PriorityRaw.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4PriorityRaw.Values <= 255 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4PriorityRawMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4PriorityRaw) setDefault() {
	var choices_set int = 0
	var choice PatternFlowIpv4PriorityRawChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowIpv4PriorityRawChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowIpv4PriorityRawChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowIpv4PriorityRawChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowIpv4PriorityRawChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowIpv4PriorityRawChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowIpv4PriorityRaw")
			}
		} else {
			intVal := otg.PatternFlowIpv4PriorityRaw_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowIpv4PriorityRaw_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
