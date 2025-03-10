package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpv2Teid *****
type patternFlowGtpv2Teid struct {
	validation
	obj              *otg.PatternFlowGtpv2Teid
	marshaller       marshalPatternFlowGtpv2Teid
	unMarshaller     unMarshalPatternFlowGtpv2Teid
	incrementHolder  PatternFlowGtpv2TeidCounter
	decrementHolder  PatternFlowGtpv2TeidCounter
	metricTagsHolder PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
}

func NewPatternFlowGtpv2Teid() PatternFlowGtpv2Teid {
	obj := patternFlowGtpv2Teid{obj: &otg.PatternFlowGtpv2Teid{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpv2Teid) msg() *otg.PatternFlowGtpv2Teid {
	return obj.obj
}

func (obj *patternFlowGtpv2Teid) setMsg(msg *otg.PatternFlowGtpv2Teid) PatternFlowGtpv2Teid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpv2Teid struct {
	obj *patternFlowGtpv2Teid
}

type marshalPatternFlowGtpv2Teid interface {
	// ToProto marshals PatternFlowGtpv2Teid to protobuf object *otg.PatternFlowGtpv2Teid
	ToProto() (*otg.PatternFlowGtpv2Teid, error)
	// ToPbText marshals PatternFlowGtpv2Teid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpv2Teid to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpv2Teid to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals PatternFlowGtpv2Teid to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalpatternFlowGtpv2Teid struct {
	obj *patternFlowGtpv2Teid
}

type unMarshalPatternFlowGtpv2Teid interface {
	// FromProto unmarshals PatternFlowGtpv2Teid from protobuf object *otg.PatternFlowGtpv2Teid
	FromProto(msg *otg.PatternFlowGtpv2Teid) (PatternFlowGtpv2Teid, error)
	// FromPbText unmarshals PatternFlowGtpv2Teid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpv2Teid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpv2Teid from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpv2Teid) Marshal() marshalPatternFlowGtpv2Teid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpv2Teid{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpv2Teid) Unmarshal() unMarshalPatternFlowGtpv2Teid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpv2Teid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpv2Teid) ToProto() (*otg.PatternFlowGtpv2Teid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpv2Teid) FromProto(msg *otg.PatternFlowGtpv2Teid) (PatternFlowGtpv2Teid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpv2Teid) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Teid) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpv2Teid) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Teid) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpv2Teid) ToJsonRaw() (string, error) {
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

func (m *marshalpatternFlowGtpv2Teid) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpv2Teid) FromJson(value string) error {
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

func (obj *patternFlowGtpv2Teid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Teid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpv2Teid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpv2Teid) Clone() (PatternFlowGtpv2Teid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpv2Teid()
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

func (obj *patternFlowGtpv2Teid) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpv2Teid is tunnel endpoint identifier. A 32-bit (4-octet) field used to multiplex different connections in the same GTP tunnel. Is present only if the teid_flag is set.
type PatternFlowGtpv2Teid interface {
	Validation
	// msg marshals PatternFlowGtpv2Teid to protobuf object *otg.PatternFlowGtpv2Teid
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpv2Teid
	// setMsg unmarshals PatternFlowGtpv2Teid from protobuf object *otg.PatternFlowGtpv2Teid
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpv2Teid) PatternFlowGtpv2Teid
	// provides marshal interface
	Marshal() marshalPatternFlowGtpv2Teid
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpv2Teid
	// validate validates PatternFlowGtpv2Teid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpv2Teid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpv2TeidChoiceEnum, set in PatternFlowGtpv2Teid
	Choice() PatternFlowGtpv2TeidChoiceEnum
	// setChoice assigns PatternFlowGtpv2TeidChoiceEnum provided by user to PatternFlowGtpv2Teid
	setChoice(value PatternFlowGtpv2TeidChoiceEnum) PatternFlowGtpv2Teid
	// HasChoice checks if Choice has been set in PatternFlowGtpv2Teid
	HasChoice() bool
	// Value returns uint32, set in PatternFlowGtpv2Teid.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternFlowGtpv2Teid
	SetValue(value uint32) PatternFlowGtpv2Teid
	// HasValue checks if Value has been set in PatternFlowGtpv2Teid
	HasValue() bool
	// Values returns []uint32, set in PatternFlowGtpv2Teid.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternFlowGtpv2Teid
	SetValues(value []uint32) PatternFlowGtpv2Teid
	// Increment returns PatternFlowGtpv2TeidCounter, set in PatternFlowGtpv2Teid.
	// PatternFlowGtpv2TeidCounter is integer counter pattern
	Increment() PatternFlowGtpv2TeidCounter
	// SetIncrement assigns PatternFlowGtpv2TeidCounter provided by user to PatternFlowGtpv2Teid.
	// PatternFlowGtpv2TeidCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpv2TeidCounter) PatternFlowGtpv2Teid
	// HasIncrement checks if Increment has been set in PatternFlowGtpv2Teid
	HasIncrement() bool
	// Decrement returns PatternFlowGtpv2TeidCounter, set in PatternFlowGtpv2Teid.
	// PatternFlowGtpv2TeidCounter is integer counter pattern
	Decrement() PatternFlowGtpv2TeidCounter
	// SetDecrement assigns PatternFlowGtpv2TeidCounter provided by user to PatternFlowGtpv2Teid.
	// PatternFlowGtpv2TeidCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpv2TeidCounter) PatternFlowGtpv2Teid
	// HasDecrement checks if Decrement has been set in PatternFlowGtpv2Teid
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIterIter, set in PatternFlowGtpv2Teid
	MetricTags() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	setNil()
}

type PatternFlowGtpv2TeidChoiceEnum string

// Enum of Choice on PatternFlowGtpv2Teid
var PatternFlowGtpv2TeidChoice = struct {
	VALUE     PatternFlowGtpv2TeidChoiceEnum
	VALUES    PatternFlowGtpv2TeidChoiceEnum
	INCREMENT PatternFlowGtpv2TeidChoiceEnum
	DECREMENT PatternFlowGtpv2TeidChoiceEnum
}{
	VALUE:     PatternFlowGtpv2TeidChoiceEnum("value"),
	VALUES:    PatternFlowGtpv2TeidChoiceEnum("values"),
	INCREMENT: PatternFlowGtpv2TeidChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpv2TeidChoiceEnum("decrement"),
}

func (obj *patternFlowGtpv2Teid) Choice() PatternFlowGtpv2TeidChoiceEnum {
	return PatternFlowGtpv2TeidChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpv2Teid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpv2Teid) setChoice(value PatternFlowGtpv2TeidChoiceEnum) PatternFlowGtpv2Teid {
	intValue, ok := otg.PatternFlowGtpv2Teid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpv2TeidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpv2Teid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpv2TeidChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpv2TeidChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpv2TeidChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpv2TeidCounter().msg()
	}

	if value == PatternFlowGtpv2TeidChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpv2TeidCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Teid) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpv2TeidChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternFlowGtpv2Teid) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternFlowGtpv2Teid object
func (obj *patternFlowGtpv2Teid) SetValue(value uint32) PatternFlowGtpv2Teid {
	obj.setChoice(PatternFlowGtpv2TeidChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternFlowGtpv2Teid) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternFlowGtpv2Teid object
func (obj *patternFlowGtpv2Teid) SetValues(value []uint32) PatternFlowGtpv2Teid {
	obj.setChoice(PatternFlowGtpv2TeidChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpv2TeidCounter
func (obj *patternFlowGtpv2Teid) Increment() PatternFlowGtpv2TeidCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpv2TeidChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpv2TeidCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpv2TeidCounter
func (obj *patternFlowGtpv2Teid) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpv2TeidCounter value in the PatternFlowGtpv2Teid object
func (obj *patternFlowGtpv2Teid) SetIncrement(value PatternFlowGtpv2TeidCounter) PatternFlowGtpv2Teid {
	obj.setChoice(PatternFlowGtpv2TeidChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpv2TeidCounter
func (obj *patternFlowGtpv2Teid) Decrement() PatternFlowGtpv2TeidCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpv2TeidChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpv2TeidCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpv2TeidCounter
func (obj *patternFlowGtpv2Teid) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpv2TeidCounter value in the PatternFlowGtpv2Teid object
func (obj *patternFlowGtpv2Teid) SetDecrement(value PatternFlowGtpv2TeidCounter) PatternFlowGtpv2Teid {
	obj.setChoice(PatternFlowGtpv2TeidChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpv2TeidMetricTag
func (obj *patternFlowGtpv2Teid) MetricTags() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpv2TeidMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter struct {
	obj                                *patternFlowGtpv2Teid
	patternFlowGtpv2TeidMetricTagSlice []PatternFlowGtpv2TeidMetricTag
	fieldPtr                           *[]*otg.PatternFlowGtpv2TeidMetricTag
}

func newPatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter(ptr *[]*otg.PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	return &patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter interface {
	setMsg(*patternFlowGtpv2Teid) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	Items() []PatternFlowGtpv2TeidMetricTag
	Add() PatternFlowGtpv2TeidMetricTag
	Append(items ...PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	Set(index int, newObj PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	Clear() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	clearHolderSlice() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
	appendHolderSlice(item PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter
}

func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) setMsg(msg *patternFlowGtpv2Teid) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpv2TeidMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) Items() []PatternFlowGtpv2TeidMetricTag {
	return obj.patternFlowGtpv2TeidMetricTagSlice
}

func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) Add() PatternFlowGtpv2TeidMetricTag {
	newObj := &otg.PatternFlowGtpv2TeidMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpv2TeidMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpv2TeidMetricTagSlice = append(obj.patternFlowGtpv2TeidMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) Append(items ...PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpv2TeidMetricTagSlice = append(obj.patternFlowGtpv2TeidMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) Set(index int, newObj PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpv2TeidMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) Clear() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpv2TeidMetricTag{}
		obj.patternFlowGtpv2TeidMetricTagSlice = []PatternFlowGtpv2TeidMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) clearHolderSlice() PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	if len(obj.patternFlowGtpv2TeidMetricTagSlice) > 0 {
		obj.patternFlowGtpv2TeidMetricTagSlice = []PatternFlowGtpv2TeidMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter) appendHolderSlice(item PatternFlowGtpv2TeidMetricTag) PatternFlowGtpv2TeidPatternFlowGtpv2TeidMetricTagIter {
	obj.patternFlowGtpv2TeidMetricTagSlice = append(obj.patternFlowGtpv2TeidMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpv2Teid) validateObj(vObj *validation, set_default bool) {
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpv2TeidMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpv2Teid) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpv2TeidChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpv2TeidChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpv2TeidChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpv2TeidChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpv2TeidChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpv2TeidChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpv2Teid")
			}
		} else {
			intVal := otg.PatternFlowGtpv2Teid_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpv2Teid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
