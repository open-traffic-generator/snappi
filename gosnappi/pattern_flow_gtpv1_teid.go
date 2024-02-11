package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv1Teid *****
type patternFlowGtpv1Teid struct {
	validation
	obj              *otg.PatternFlowGtpv1Teid
	marshaller       marshalPatternFlowGtpv1Teid
	unMarshaller     unMarshalPatternFlowGtpv1Teid
	incrementHolder  PatternFlowGtpv1TeidCounter
	decrementHolder  PatternFlowGtpv1TeidCounter
	metricTagsHolder PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
}

func NewPatternFlowGtpv1Teid() PatternFlowGtpv1Teid {
	obj := patternFlowGtpv1Teid{obj: &otg.PatternFlowGtpv1Teid{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv1Teid) msg() *otg.PatternFlowGtpv1Teid {
	return obj.obj
}

func (obj *patternFlowGtpv1Teid) setMsg(msg *otg.PatternFlowGtpv1Teid) PatternFlowGtpv1Teid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv1Teid struct {
	obj *patternFlowGtpv1Teid
}

type marshalPatternFlowGtpv1Teid interface {
	// ToProto marshals PatternFlowGtpv1Teid to protobuf object *otg.PatternFlowGtpv1Teid
	ToProto() (*otg.PatternFlowGtpv1Teid, error)
	// ToPbText marshals PatternFlowGtpv1Teid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv1Teid to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv1Teid to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpv1Teid struct {
	obj *patternFlowGtpv1Teid
}

type unMarshalPatternFlowGtpv1Teid interface {
	// FromProto unmarshals PatternFlowGtpv1Teid from protobuf object *otg.PatternFlowGtpv1Teid
	FromProto(msg *otg.PatternFlowGtpv1Teid) (PatternFlowGtpv1Teid, error)
	// FromPbText unmarshals PatternFlowGtpv1Teid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv1Teid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv1Teid from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv1Teid) Marshal() marshalPatternFlowGtpv1Teid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv1Teid{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv1Teid) Unmarshal() unMarshalPatternFlowGtpv1Teid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv1Teid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv1Teid) ToProto() (*otg.PatternFlowGtpv1Teid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv1Teid) FromProto(msg *otg.PatternFlowGtpv1Teid) (PatternFlowGtpv1Teid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv1Teid) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Teid) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv1Teid) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Teid) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv1Teid) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv1Teid) FromJson(value string) error {
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

func (obj *patternFlowGtpv1Teid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Teid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv1Teid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv1Teid) Clone() (PatternFlowGtpv1Teid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv1Teid()
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

func (obj *patternFlowGtpv1Teid) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv1Teid is tunnel endpoint identifier (TEID) used to multiplex connections in the same GTP tunnel
type PatternFlowGtpv1Teid interface {
	Validation
	// msg marshals PatternFlowGtpv1Teid to protobuf object *otg.PatternFlowGtpv1Teid
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv1Teid
	// setMsg unmarshals PatternFlowGtpv1Teid from protobuf object *otg.PatternFlowGtpv1Teid
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv1Teid) PatternFlowGtpv1Teid
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv1Teid
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv1Teid
	// validate validates PatternFlowGtpv1Teid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv1Teid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv1TeidChoiceEnum, set in PatternFlowGtpv1Teid
	Choice() PatternFlowGtpv1TeidChoiceEnum
	// setChoice assigns PatternFlowGtpv1TeidChoiceEnum provided by user to PatternFlowGtpv1Teid
	setChoice(value PatternFlowGtpv1TeidChoiceEnum) PatternFlowGtpv1Teid
	// HasChoice checks if Choice has been set in PatternFlowGtpv1Teid
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv1Teid.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv1Teid
	SetValue(value uint32) PatternFlowGtpv1Teid
	// HasValue checks if Value has been set in PatternFlowGtpv1Teid
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv1Teid.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv1Teid
	SetValues(value []uint32) PatternFlowGtpv1Teid
	// Increment returns PatternFlowGtpv1TeidCounter, set in PatternFlowGtpv1Teid.
	// PatternFlowGtpv1TeidCounter is integer counter pattern
	Increment() PatternFlowGtpv1TeidCounter
	// SetIncrement assigns PatternFlowGtpv1TeidCounter provided by user to PatternFlowGtpv1Teid.
	// PatternFlowGtpv1TeidCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv1TeidCounter) PatternFlowGtpv1Teid
	// HasIncrement checks if Increment has been set in PatternFlowGtpv1Teid
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv1TeidCounter, set in PatternFlowGtpv1Teid.
	// PatternFlowGtpv1TeidCounter is integer counter pattern
	Decrement() PatternFlowGtpv1TeidCounter
	// SetDecrement assigns PatternFlowGtpv1TeidCounter provided by user to PatternFlowGtpv1Teid.
	// PatternFlowGtpv1TeidCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv1TeidCounter) PatternFlowGtpv1Teid
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv1Teid
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIterIter, set in PatternFlowGtpv1Teid
	MetricTags() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	setNil()
}

type PatternFlowGtpv1TeidChoiceEnum string

// Enum of Choice on PatternFlowGtpv1Teid
var PatternFlowGtpv1TeidChoice = struct {
	VALUE     PatternFlowGtpv1TeidChoiceEnum
	VALUES    PatternFlowGtpv1TeidChoiceEnum
	INCREMENT PatternFlowGtpv1TeidChoiceEnum
	DECREMENT PatternFlowGtpv1TeidChoiceEnum
}{
	VALUE:     PatternFlowGtpv1TeidChoiceEnum("value"),
	VALUES:    PatternFlowGtpv1TeidChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv1TeidChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv1TeidChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv1Teid) Choice() PatternFlowGtpv1TeidChoiceEnum {
	return PatternFlowGtpv1TeidChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv1Teid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv1Teid) setChoice(value PatternFlowGtpv1TeidChoiceEnum) PatternFlowGtpv1Teid {
	intValue, ok := otg.PatternFlowGtpv1Teid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv1TeidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv1Teid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv1TeidChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv1TeidChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv1TeidChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv1TeidCounter().msg()
	}

	if value == PatternFlowGtpv1TeidChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv1TeidCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Teid) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv1TeidChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv1Teid) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv1Teid object
func (obj *patternFlowGtpv1Teid) SetValue(value uint32) PatternFlowGtpv1Teid {
	obj.setChoice(PatternFlowGtpv1TeidChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv1Teid) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv1Teid object
func (obj *patternFlowGtpv1Teid) SetValues(value []uint32) PatternFlowGtpv1Teid {
	obj.setChoice(PatternFlowGtpv1TeidChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv1TeidCounter
func (obj *patternFlowGtpv1Teid) Increment() PatternFlowGtpv1TeidCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv1TeidChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv1TeidCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv1TeidCounter
func (obj *patternFlowGtpv1Teid) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv1TeidCounter value in the PatternFlowGtpv1Teid object
func (obj *patternFlowGtpv1Teid) SetIncrement(value PatternFlowGtpv1TeidCounter) PatternFlowGtpv1Teid {
	obj.setChoice(PatternFlowGtpv1TeidChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv1TeidCounter
func (obj *patternFlowGtpv1Teid) Decrement() PatternFlowGtpv1TeidCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv1TeidChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv1TeidCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv1TeidCounter
func (obj *patternFlowGtpv1Teid) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv1TeidCounter value in the PatternFlowGtpv1Teid object
func (obj *patternFlowGtpv1Teid) SetDecrement(value PatternFlowGtpv1TeidCounter) PatternFlowGtpv1Teid {
	obj.setChoice(PatternFlowGtpv1TeidChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv1TeidMetricTag
func (obj *patternFlowGtpv1Teid) MetricTags() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv1TeidMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter struct {
	obj                                *patternFlowGtpv1Teid
	patternFlowGtpv1TeidMetricTagSlice []PatternFlowGtpv1TeidMetricTag
	fieldPtr                           *[]*otg.PatternFlowGtpv1TeidMetricTag
}

func newPatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter(ptr *[]*otg.PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	return &patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter interface {
	setMsg(*patternFlowGtpv1Teid) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	Items() []PatternFlowGtpv1TeidMetricTag
	Add() PatternFlowGtpv1TeidMetricTag
	Append(items ...PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	Set(index int, newObj PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	Clear() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	clearHolderSlice() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
	appendHolderSlice(item PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter
}

func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) setMsg(msg *patternFlowGtpv1Teid) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv1TeidMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) Items() []PatternFlowGtpv1TeidMetricTag {
	return obj.patternFlowGtpv1TeidMetricTagSlice
}

func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) Add() PatternFlowGtpv1TeidMetricTag {
	newObj := &otg.PatternFlowGtpv1TeidMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv1TeidMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv1TeidMetricTagSlice = append(obj.patternFlowGtpv1TeidMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) Append(items ...PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv1TeidMetricTagSlice = append(obj.patternFlowGtpv1TeidMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) Set(index int, newObj PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv1TeidMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) Clear() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv1TeidMetricTag{}
		obj.patternFlowGtpv1TeidMetricTagSlice = []PatternFlowGtpv1TeidMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) clearHolderSlice() PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	if len(obj.patternFlowGtpv1TeidMetricTagSlice) > 0 {
		obj.patternFlowGtpv1TeidMetricTagSlice = []PatternFlowGtpv1TeidMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter) appendHolderSlice(item PatternFlowGtpv1TeidMetricTag) PatternFlowGtpv1TeidPatternFlowGtpv1TeidMetricTagIter {
	obj.patternFlowGtpv1TeidMetricTagSlice = append(obj.patternFlowGtpv1TeidMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv1Teid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv1TeidMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv1Teid) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowGtpv1TeidChoice.VALUE)

	}

}
