package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosMonetary *****
type patternFlowIpv4TosMonetary struct {
	validation
	obj              *otg.PatternFlowIpv4TosMonetary
	marshaller       marshalPatternFlowIpv4TosMonetary
	unMarshaller     unMarshalPatternFlowIpv4TosMonetary
	incrementHolder  PatternFlowIpv4TosMonetaryCounter
	decrementHolder  PatternFlowIpv4TosMonetaryCounter
	metricTagsHolder PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
}

func NewPatternFlowIpv4TosMonetary() PatternFlowIpv4TosMonetary {
	obj := patternFlowIpv4TosMonetary{obj: &otg.PatternFlowIpv4TosMonetary{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosMonetary) msg() *otg.PatternFlowIpv4TosMonetary {
	return obj.obj
}

func (obj *patternFlowIpv4TosMonetary) setMsg(msg *otg.PatternFlowIpv4TosMonetary) PatternFlowIpv4TosMonetary {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosMonetary struct {
	obj *patternFlowIpv4TosMonetary
}

type marshalPatternFlowIpv4TosMonetary interface {
	// ToProto marshals PatternFlowIpv4TosMonetary to protobuf object *otg.PatternFlowIpv4TosMonetary
	ToProto() (*otg.PatternFlowIpv4TosMonetary, error)
	// ToPbText marshals PatternFlowIpv4TosMonetary to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosMonetary to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosMonetary to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosMonetary struct {
	obj *patternFlowIpv4TosMonetary
}

type unMarshalPatternFlowIpv4TosMonetary interface {
	// FromProto unmarshals PatternFlowIpv4TosMonetary from protobuf object *otg.PatternFlowIpv4TosMonetary
	FromProto(msg *otg.PatternFlowIpv4TosMonetary) (PatternFlowIpv4TosMonetary, error)
	// FromPbText unmarshals PatternFlowIpv4TosMonetary from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosMonetary from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosMonetary from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosMonetary) Marshal() marshalPatternFlowIpv4TosMonetary {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosMonetary{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosMonetary) Unmarshal() unMarshalPatternFlowIpv4TosMonetary {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosMonetary{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosMonetary) ToProto() (*otg.PatternFlowIpv4TosMonetary, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosMonetary) FromProto(msg *otg.PatternFlowIpv4TosMonetary) (PatternFlowIpv4TosMonetary, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosMonetary) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetary) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetary) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetary) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosMonetary) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosMonetary) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosMonetary) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetary) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosMonetary) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosMonetary) Clone() (PatternFlowIpv4TosMonetary, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosMonetary()
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

func (obj *patternFlowIpv4TosMonetary) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosMonetary is monetary
type PatternFlowIpv4TosMonetary interface {
	Validation
	// msg marshals PatternFlowIpv4TosMonetary to protobuf object *otg.PatternFlowIpv4TosMonetary
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosMonetary
	// setMsg unmarshals PatternFlowIpv4TosMonetary from protobuf object *otg.PatternFlowIpv4TosMonetary
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosMonetary) PatternFlowIpv4TosMonetary
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosMonetary
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosMonetary
	// validate validates PatternFlowIpv4TosMonetary
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosMonetary, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosMonetaryChoiceEnum, set in PatternFlowIpv4TosMonetary
	Choice() PatternFlowIpv4TosMonetaryChoiceEnum
	// setChoice assigns PatternFlowIpv4TosMonetaryChoiceEnum provided by user to PatternFlowIpv4TosMonetary
	setChoice(value PatternFlowIpv4TosMonetaryChoiceEnum) PatternFlowIpv4TosMonetary
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosMonetary
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosMonetary.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosMonetary
	SetValue(value uint32) PatternFlowIpv4TosMonetary
	// HasValue checks if Value has been set in PatternFlowIpv4TosMonetary
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosMonetary.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosMonetary
	SetValues(value []uint32) PatternFlowIpv4TosMonetary
	// Increment returns PatternFlowIpv4TosMonetaryCounter, set in PatternFlowIpv4TosMonetary.
	// PatternFlowIpv4TosMonetaryCounter is integer counter pattern
	Increment() PatternFlowIpv4TosMonetaryCounter
	// SetIncrement assigns PatternFlowIpv4TosMonetaryCounter provided by user to PatternFlowIpv4TosMonetary.
	// PatternFlowIpv4TosMonetaryCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetary
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosMonetary
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosMonetaryCounter, set in PatternFlowIpv4TosMonetary.
	// PatternFlowIpv4TosMonetaryCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosMonetaryCounter
	// SetDecrement assigns PatternFlowIpv4TosMonetaryCounter provided by user to PatternFlowIpv4TosMonetary.
	// PatternFlowIpv4TosMonetaryCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetary
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosMonetary
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIterIter, set in PatternFlowIpv4TosMonetary
	MetricTags() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	setNil()
}

type PatternFlowIpv4TosMonetaryChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosMonetary
var PatternFlowIpv4TosMonetaryChoice = struct {
	VALUE     PatternFlowIpv4TosMonetaryChoiceEnum
	VALUES    PatternFlowIpv4TosMonetaryChoiceEnum
	INCREMENT PatternFlowIpv4TosMonetaryChoiceEnum
	DECREMENT PatternFlowIpv4TosMonetaryChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosMonetaryChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosMonetaryChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosMonetaryChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosMonetaryChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosMonetary) Choice() PatternFlowIpv4TosMonetaryChoiceEnum {
	return PatternFlowIpv4TosMonetaryChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosMonetary) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosMonetary) setChoice(value PatternFlowIpv4TosMonetaryChoiceEnum) PatternFlowIpv4TosMonetary {
	intValue, ok := otg.PatternFlowIpv4TosMonetary_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosMonetaryChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosMonetary_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosMonetaryChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosMonetaryChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosMonetaryChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosMonetaryCounter().msg()
	}

	if value == PatternFlowIpv4TosMonetaryChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosMonetaryCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosMonetary) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosMonetaryChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosMonetary) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosMonetary object
func (obj *patternFlowIpv4TosMonetary) SetValue(value uint32) PatternFlowIpv4TosMonetary {
	obj.setChoice(PatternFlowIpv4TosMonetaryChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosMonetary) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosMonetary object
func (obj *patternFlowIpv4TosMonetary) SetValues(value []uint32) PatternFlowIpv4TosMonetary {
	obj.setChoice(PatternFlowIpv4TosMonetaryChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosMonetaryCounter
func (obj *patternFlowIpv4TosMonetary) Increment() PatternFlowIpv4TosMonetaryCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosMonetaryChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosMonetaryCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosMonetaryCounter
func (obj *patternFlowIpv4TosMonetary) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosMonetaryCounter value in the PatternFlowIpv4TosMonetary object
func (obj *patternFlowIpv4TosMonetary) SetIncrement(value PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetary {
	obj.setChoice(PatternFlowIpv4TosMonetaryChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosMonetaryCounter
func (obj *patternFlowIpv4TosMonetary) Decrement() PatternFlowIpv4TosMonetaryCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosMonetaryChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosMonetaryCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosMonetaryCounter
func (obj *patternFlowIpv4TosMonetary) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosMonetaryCounter value in the PatternFlowIpv4TosMonetary object
func (obj *patternFlowIpv4TosMonetary) SetDecrement(value PatternFlowIpv4TosMonetaryCounter) PatternFlowIpv4TosMonetary {
	obj.setChoice(PatternFlowIpv4TosMonetaryChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosMonetaryMetricTag
func (obj *patternFlowIpv4TosMonetary) MetricTags() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosMonetaryMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter struct {
	obj                                      *patternFlowIpv4TosMonetary
	patternFlowIpv4TosMonetaryMetricTagSlice []PatternFlowIpv4TosMonetaryMetricTag
	fieldPtr                                 *[]*otg.PatternFlowIpv4TosMonetaryMetricTag
}

func newPatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	return &patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter interface {
	setMsg(*patternFlowIpv4TosMonetary) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	Items() []PatternFlowIpv4TosMonetaryMetricTag
	Add() PatternFlowIpv4TosMonetaryMetricTag
	Append(items ...PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	Clear() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter
}

func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) setMsg(msg *patternFlowIpv4TosMonetary) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosMonetaryMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) Items() []PatternFlowIpv4TosMonetaryMetricTag {
	return obj.patternFlowIpv4TosMonetaryMetricTagSlice
}

func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) Add() PatternFlowIpv4TosMonetaryMetricTag {
	newObj := &otg.PatternFlowIpv4TosMonetaryMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosMonetaryMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosMonetaryMetricTagSlice = append(obj.patternFlowIpv4TosMonetaryMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) Append(items ...PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosMonetaryMetricTagSlice = append(obj.patternFlowIpv4TosMonetaryMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) Set(index int, newObj PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosMonetaryMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) Clear() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosMonetaryMetricTag{}
		obj.patternFlowIpv4TosMonetaryMetricTagSlice = []PatternFlowIpv4TosMonetaryMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) clearHolderSlice() PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	if len(obj.patternFlowIpv4TosMonetaryMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosMonetaryMetricTagSlice = []PatternFlowIpv4TosMonetaryMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosMonetaryMetricTag) PatternFlowIpv4TosMonetaryPatternFlowIpv4TosMonetaryMetricTagIter {
	obj.patternFlowIpv4TosMonetaryMetricTagSlice = append(obj.patternFlowIpv4TosMonetaryMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosMonetary) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosMonetary.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosMonetary.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosMonetaryMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosMonetary) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv4TosMonetaryChoice.VALUE)

	}

}
