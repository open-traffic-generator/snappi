package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowIpv4TosDelay *****
type patternFlowIpv4TosDelay struct {
	validation
	obj              *otg.PatternFlowIpv4TosDelay
	marshaller       marshalPatternFlowIpv4TosDelay
	unMarshaller     unMarshalPatternFlowIpv4TosDelay
	incrementHolder  PatternFlowIpv4TosDelayCounter
	decrementHolder  PatternFlowIpv4TosDelayCounter
	metricTagsHolder PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
}

func NewPatternFlowIpv4TosDelay() PatternFlowIpv4TosDelay {
	obj := patternFlowIpv4TosDelay{obj: &otg.PatternFlowIpv4TosDelay{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowIpv4TosDelay) msg() *otg.PatternFlowIpv4TosDelay {
	return obj.obj
}

func (obj *patternFlowIpv4TosDelay) setMsg(msg *otg.PatternFlowIpv4TosDelay) PatternFlowIpv4TosDelay {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowIpv4TosDelay struct {
	obj *patternFlowIpv4TosDelay
}

type marshalPatternFlowIpv4TosDelay interface {
	// ToProto marshals PatternFlowIpv4TosDelay to protobuf object *otg.PatternFlowIpv4TosDelay
	ToProto() (*otg.PatternFlowIpv4TosDelay, error)
	// ToPbText marshals PatternFlowIpv4TosDelay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowIpv4TosDelay to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowIpv4TosDelay to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowIpv4TosDelay struct {
	obj *patternFlowIpv4TosDelay
}

type unMarshalPatternFlowIpv4TosDelay interface {
	// FromProto unmarshals PatternFlowIpv4TosDelay from protobuf object *otg.PatternFlowIpv4TosDelay
	FromProto(msg *otg.PatternFlowIpv4TosDelay) (PatternFlowIpv4TosDelay, error)
	// FromPbText unmarshals PatternFlowIpv4TosDelay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowIpv4TosDelay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowIpv4TosDelay from JSON text
	FromJson(value string) error
}

func (obj *patternFlowIpv4TosDelay) Marshal() marshalPatternFlowIpv4TosDelay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowIpv4TosDelay{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowIpv4TosDelay) Unmarshal() unMarshalPatternFlowIpv4TosDelay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowIpv4TosDelay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowIpv4TosDelay) ToProto() (*otg.PatternFlowIpv4TosDelay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowIpv4TosDelay) FromProto(msg *otg.PatternFlowIpv4TosDelay) (PatternFlowIpv4TosDelay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowIpv4TosDelay) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelay) FromPbText(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelay) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelay) FromYaml(value string) error {
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

func (m *marshalpatternFlowIpv4TosDelay) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowIpv4TosDelay) FromJson(value string) error {
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

func (obj *patternFlowIpv4TosDelay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowIpv4TosDelay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowIpv4TosDelay) Clone() (PatternFlowIpv4TosDelay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowIpv4TosDelay()
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

func (obj *patternFlowIpv4TosDelay) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowIpv4TosDelay is delay
type PatternFlowIpv4TosDelay interface {
	Validation
	// msg marshals PatternFlowIpv4TosDelay to protobuf object *otg.PatternFlowIpv4TosDelay
	// and doesn't set defaults
	msg() *otg.PatternFlowIpv4TosDelay
	// setMsg unmarshals PatternFlowIpv4TosDelay from protobuf object *otg.PatternFlowIpv4TosDelay
	// and doesn't set defaults
	setMsg(*otg.PatternFlowIpv4TosDelay) PatternFlowIpv4TosDelay
	// provides marshal interface
	Marshal() marshalPatternFlowIpv4TosDelay
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowIpv4TosDelay
	// validate validates PatternFlowIpv4TosDelay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowIpv4TosDelay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowIpv4TosDelayChoiceEnum, set in PatternFlowIpv4TosDelay
	Choice() PatternFlowIpv4TosDelayChoiceEnum
	// setChoice assigns PatternFlowIpv4TosDelayChoiceEnum provided by user to PatternFlowIpv4TosDelay
	setChoice(value PatternFlowIpv4TosDelayChoiceEnum) PatternFlowIpv4TosDelay
	// HasChoice checks if Choice has been set in PatternFlowIpv4TosDelay
	HasChoice() bool
	// Value returns uint32, set in PatternFlowIpv4TosDelay.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowIpv4TosDelay
	SetValue(value uint32) PatternFlowIpv4TosDelay
	// HasValue checks if Value has been set in PatternFlowIpv4TosDelay
	HasValue() bool
	// Values returns []uint32, set in PatternFlowIpv4TosDelay.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowIpv4TosDelay
	SetValues(value []uint32) PatternFlowIpv4TosDelay
	// Increment returns PatternFlowIpv4TosDelayCounter, set in PatternFlowIpv4TosDelay.
	// PatternFlowIpv4TosDelayCounter is integer counter pattern
	Increment() PatternFlowIpv4TosDelayCounter
	// SetIncrement assigns PatternFlowIpv4TosDelayCounter provided by user to PatternFlowIpv4TosDelay.
	// PatternFlowIpv4TosDelayCounter is integer counter pattern
	SetIncrement(value PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelay
	// HasIncrement checks if Increment has been set in PatternFlowIpv4TosDelay
	HasIncrement() bool
	// Decrement returns PatternFlowIpv4TosDelayCounter, set in PatternFlowIpv4TosDelay.
	// PatternFlowIpv4TosDelayCounter is integer counter pattern
	Decrement() PatternFlowIpv4TosDelayCounter
	// SetDecrement assigns PatternFlowIpv4TosDelayCounter provided by user to PatternFlowIpv4TosDelay.
	// PatternFlowIpv4TosDelayCounter is integer counter pattern
	SetDecrement(value PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelay
	// HasDecrement checks if Decrement has been set in PatternFlowIpv4TosDelay
	HasDecrement() bool
	// MetricTags returns PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIterIter, set in PatternFlowIpv4TosDelay
	MetricTags() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	setNil()
}

type PatternFlowIpv4TosDelayChoiceEnum string

// Enum of Choice on PatternFlowIpv4TosDelay
var PatternFlowIpv4TosDelayChoice = struct {
	VALUE     PatternFlowIpv4TosDelayChoiceEnum
	VALUES    PatternFlowIpv4TosDelayChoiceEnum
	INCREMENT PatternFlowIpv4TosDelayChoiceEnum
	DECREMENT PatternFlowIpv4TosDelayChoiceEnum
}{
	VALUE:     PatternFlowIpv4TosDelayChoiceEnum("value"),
	VALUES:    PatternFlowIpv4TosDelayChoiceEnum("values"),
	INCREMENT: PatternFlowIpv4TosDelayChoiceEnum("increment"),
	DECREMENT: PatternFlowIpv4TosDelayChoiceEnum("decrement"),
}

func (obj *patternFlowIpv4TosDelay) Choice() PatternFlowIpv4TosDelayChoiceEnum {
	return PatternFlowIpv4TosDelayChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowIpv4TosDelay) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowIpv4TosDelay) setChoice(value PatternFlowIpv4TosDelayChoiceEnum) PatternFlowIpv4TosDelay {
	intValue, ok := otg.PatternFlowIpv4TosDelay_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowIpv4TosDelayChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowIpv4TosDelay_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowIpv4TosDelayChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowIpv4TosDelayChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowIpv4TosDelayChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowIpv4TosDelayCounter().msg()
	}

	if value == PatternFlowIpv4TosDelayChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowIpv4TosDelayCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosDelay) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowIpv4TosDelayChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowIpv4TosDelay) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowIpv4TosDelay object
func (obj *patternFlowIpv4TosDelay) SetValue(value uint32) PatternFlowIpv4TosDelay {
	obj.setChoice(PatternFlowIpv4TosDelayChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowIpv4TosDelay) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowIpv4TosDelay object
func (obj *patternFlowIpv4TosDelay) SetValues(value []uint32) PatternFlowIpv4TosDelay {
	obj.setChoice(PatternFlowIpv4TosDelayChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowIpv4TosDelayCounter
func (obj *patternFlowIpv4TosDelay) Increment() PatternFlowIpv4TosDelayCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowIpv4TosDelayChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowIpv4TosDelayCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowIpv4TosDelayCounter
func (obj *patternFlowIpv4TosDelay) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowIpv4TosDelayCounter value in the PatternFlowIpv4TosDelay object
func (obj *patternFlowIpv4TosDelay) SetIncrement(value PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelay {
	obj.setChoice(PatternFlowIpv4TosDelayChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosDelayCounter
func (obj *patternFlowIpv4TosDelay) Decrement() PatternFlowIpv4TosDelayCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowIpv4TosDelayChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowIpv4TosDelayCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowIpv4TosDelayCounter
func (obj *patternFlowIpv4TosDelay) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowIpv4TosDelayCounter value in the PatternFlowIpv4TosDelay object
func (obj *patternFlowIpv4TosDelay) SetDecrement(value PatternFlowIpv4TosDelayCounter) PatternFlowIpv4TosDelay {
	obj.setChoice(PatternFlowIpv4TosDelayChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowIpv4TosDelayMetricTag
func (obj *patternFlowIpv4TosDelay) MetricTags() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowIpv4TosDelayMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter struct {
	obj                                   *patternFlowIpv4TosDelay
	patternFlowIpv4TosDelayMetricTagSlice []PatternFlowIpv4TosDelayMetricTag
	fieldPtr                              *[]*otg.PatternFlowIpv4TosDelayMetricTag
}

func newPatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter(ptr *[]*otg.PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	return &patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter{fieldPtr: ptr}
}

type PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter interface {
	setMsg(*patternFlowIpv4TosDelay) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	Items() []PatternFlowIpv4TosDelayMetricTag
	Add() PatternFlowIpv4TosDelayMetricTag
	Append(items ...PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	Set(index int, newObj PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	Clear() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	clearHolderSlice() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
	appendHolderSlice(item PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter
}

func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) setMsg(msg *patternFlowIpv4TosDelay) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowIpv4TosDelayMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) Items() []PatternFlowIpv4TosDelayMetricTag {
	return obj.patternFlowIpv4TosDelayMetricTagSlice
}

func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) Add() PatternFlowIpv4TosDelayMetricTag {
	newObj := &otg.PatternFlowIpv4TosDelayMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowIpv4TosDelayMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowIpv4TosDelayMetricTagSlice = append(obj.patternFlowIpv4TosDelayMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) Append(items ...PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowIpv4TosDelayMetricTagSlice = append(obj.patternFlowIpv4TosDelayMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) Set(index int, newObj PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowIpv4TosDelayMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) Clear() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowIpv4TosDelayMetricTag{}
		obj.patternFlowIpv4TosDelayMetricTagSlice = []PatternFlowIpv4TosDelayMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) clearHolderSlice() PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	if len(obj.patternFlowIpv4TosDelayMetricTagSlice) > 0 {
		obj.patternFlowIpv4TosDelayMetricTagSlice = []PatternFlowIpv4TosDelayMetricTag{}
	}
	return obj
}
func (obj *patternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter) appendHolderSlice(item PatternFlowIpv4TosDelayMetricTag) PatternFlowIpv4TosDelayPatternFlowIpv4TosDelayMetricTagIter {
	obj.patternFlowIpv4TosDelayMetricTagSlice = append(obj.patternFlowIpv4TosDelayMetricTagSlice, item)
	return obj
}

func (obj *patternFlowIpv4TosDelay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowIpv4TosDelay.Value <= 1 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 1 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternFlowIpv4TosDelay.Values <= 1 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowIpv4TosDelayMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowIpv4TosDelay) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowIpv4TosDelayChoice.VALUE)

	}

}
