package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowGtpExtensionContents *****
type patternFlowGtpExtensionContents struct {
	validation
	obj              *otg.PatternFlowGtpExtensionContents
	marshaller       marshalPatternFlowGtpExtensionContents
	unMarshaller     unMarshalPatternFlowGtpExtensionContents
	incrementHolder  PatternFlowGtpExtensionContentsCounter
	decrementHolder  PatternFlowGtpExtensionContentsCounter
	metricTagsHolder PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
}

func NewPatternFlowGtpExtensionContents() PatternFlowGtpExtensionContents {
	obj := patternFlowGtpExtensionContents{obj: &otg.PatternFlowGtpExtensionContents{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowGtpExtensionContents) msg() *otg.PatternFlowGtpExtensionContents {
	return obj.obj
}

func (obj *patternFlowGtpExtensionContents) setMsg(msg *otg.PatternFlowGtpExtensionContents) PatternFlowGtpExtensionContents {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowGtpExtensionContents struct {
	obj *patternFlowGtpExtensionContents
}

type marshalPatternFlowGtpExtensionContents interface {
	// ToProto marshals PatternFlowGtpExtensionContents to protobuf object *otg.PatternFlowGtpExtensionContents
	ToProto() (*otg.PatternFlowGtpExtensionContents, error)
	// ToPbText marshals PatternFlowGtpExtensionContents to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowGtpExtensionContents to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowGtpExtensionContents to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowGtpExtensionContents struct {
	obj *patternFlowGtpExtensionContents
}

type unMarshalPatternFlowGtpExtensionContents interface {
	// FromProto unmarshals PatternFlowGtpExtensionContents from protobuf object *otg.PatternFlowGtpExtensionContents
	FromProto(msg *otg.PatternFlowGtpExtensionContents) (PatternFlowGtpExtensionContents, error)
	// FromPbText unmarshals PatternFlowGtpExtensionContents from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowGtpExtensionContents from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowGtpExtensionContents from JSON text
	FromJson(value string) error
}

func (obj *patternFlowGtpExtensionContents) Marshal() marshalPatternFlowGtpExtensionContents {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowGtpExtensionContents{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowGtpExtensionContents) Unmarshal() unMarshalPatternFlowGtpExtensionContents {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowGtpExtensionContents{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowGtpExtensionContents) ToProto() (*otg.PatternFlowGtpExtensionContents, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowGtpExtensionContents) FromProto(msg *otg.PatternFlowGtpExtensionContents) (PatternFlowGtpExtensionContents, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowGtpExtensionContents) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContents) FromPbText(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContents) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContents) FromYaml(value string) error {
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

func (m *marshalpatternFlowGtpExtensionContents) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowGtpExtensionContents) FromJson(value string) error {
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

func (obj *patternFlowGtpExtensionContents) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContents) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowGtpExtensionContents) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowGtpExtensionContents) Clone() (PatternFlowGtpExtensionContents, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowGtpExtensionContents()
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

func (obj *patternFlowGtpExtensionContents) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowGtpExtensionContents is the extension header contents
type PatternFlowGtpExtensionContents interface {
	Validation
	// msg marshals PatternFlowGtpExtensionContents to protobuf object *otg.PatternFlowGtpExtensionContents
	// and doesn't set defaults
	msg() *otg.PatternFlowGtpExtensionContents
	// setMsg unmarshals PatternFlowGtpExtensionContents from protobuf object *otg.PatternFlowGtpExtensionContents
	// and doesn't set defaults
	setMsg(*otg.PatternFlowGtpExtensionContents) PatternFlowGtpExtensionContents
	// provides marshal interface
	Marshal() marshalPatternFlowGtpExtensionContents
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowGtpExtensionContents
	// validate validates PatternFlowGtpExtensionContents
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowGtpExtensionContents, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowGtpExtensionContentsChoiceEnum, set in PatternFlowGtpExtensionContents
	Choice() PatternFlowGtpExtensionContentsChoiceEnum
	// setChoice assigns PatternFlowGtpExtensionContentsChoiceEnum provided by user to PatternFlowGtpExtensionContents
	setChoice(value PatternFlowGtpExtensionContentsChoiceEnum) PatternFlowGtpExtensionContents
	// HasChoice checks if Choice has been set in PatternFlowGtpExtensionContents
	HasChoice() bool
	// Value returns uint64, set in PatternFlowGtpExtensionContents.
	Value() uint64
	// SetValue assigns uint64 provided by user to PatternFlowGtpExtensionContents
	SetValue(value uint64) PatternFlowGtpExtensionContents
	// HasValue checks if Value has been set in PatternFlowGtpExtensionContents
	HasValue() bool
	// Values returns []uint64, set in PatternFlowGtpExtensionContents.
	Values() []uint64
	// SetValues assigns []uint64 provided by user to PatternFlowGtpExtensionContents
	SetValues(value []uint64) PatternFlowGtpExtensionContents
	// Increment returns PatternFlowGtpExtensionContentsCounter, set in PatternFlowGtpExtensionContents.
	// PatternFlowGtpExtensionContentsCounter is integer counter pattern
	Increment() PatternFlowGtpExtensionContentsCounter
	// SetIncrement assigns PatternFlowGtpExtensionContentsCounter provided by user to PatternFlowGtpExtensionContents.
	// PatternFlowGtpExtensionContentsCounter is integer counter pattern
	SetIncrement(value PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContents
	// HasIncrement checks if Increment has been set in PatternFlowGtpExtensionContents
	HasIncrement() bool
	// Decrement returns PatternFlowGtpExtensionContentsCounter, set in PatternFlowGtpExtensionContents.
	// PatternFlowGtpExtensionContentsCounter is integer counter pattern
	Decrement() PatternFlowGtpExtensionContentsCounter
	// SetDecrement assigns PatternFlowGtpExtensionContentsCounter provided by user to PatternFlowGtpExtensionContents.
	// PatternFlowGtpExtensionContentsCounter is integer counter pattern
	SetDecrement(value PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContents
	// HasDecrement checks if Decrement has been set in PatternFlowGtpExtensionContents
	HasDecrement() bool
	// MetricTags returns PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIterIter, set in PatternFlowGtpExtensionContents
	MetricTags() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	setNil()
}

type PatternFlowGtpExtensionContentsChoiceEnum string

// Enum of Choice on PatternFlowGtpExtensionContents
var PatternFlowGtpExtensionContentsChoice = struct {
	VALUE     PatternFlowGtpExtensionContentsChoiceEnum
	VALUES    PatternFlowGtpExtensionContentsChoiceEnum
	INCREMENT PatternFlowGtpExtensionContentsChoiceEnum
	DECREMENT PatternFlowGtpExtensionContentsChoiceEnum
}{
	VALUE:     PatternFlowGtpExtensionContentsChoiceEnum("value"),
	VALUES:    PatternFlowGtpExtensionContentsChoiceEnum("values"),
	INCREMENT: PatternFlowGtpExtensionContentsChoiceEnum("increment"),
	DECREMENT: PatternFlowGtpExtensionContentsChoiceEnum("decrement"),
}

func (obj *patternFlowGtpExtensionContents) Choice() PatternFlowGtpExtensionContentsChoiceEnum {
	return PatternFlowGtpExtensionContentsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowGtpExtensionContents) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowGtpExtensionContents) setChoice(value PatternFlowGtpExtensionContentsChoiceEnum) PatternFlowGtpExtensionContents {
	intValue, ok := otg.PatternFlowGtpExtensionContents_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowGtpExtensionContentsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowGtpExtensionContents_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowGtpExtensionContentsChoice.VALUE {
		defaultValue := uint64(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowGtpExtensionContentsChoice.VALUES {
		defaultValue := []uint64{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowGtpExtensionContentsChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowGtpExtensionContentsCounter().msg()
	}

	if value == PatternFlowGtpExtensionContentsChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowGtpExtensionContentsCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint64
func (obj *patternFlowGtpExtensionContents) Value() uint64 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowGtpExtensionContentsChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint64
func (obj *patternFlowGtpExtensionContents) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint64 value in the PatternFlowGtpExtensionContents object
func (obj *patternFlowGtpExtensionContents) SetValue(value uint64) PatternFlowGtpExtensionContents {
	obj.setChoice(PatternFlowGtpExtensionContentsChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint64
func (obj *patternFlowGtpExtensionContents) Values() []uint64 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint64{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint64 value in the PatternFlowGtpExtensionContents object
func (obj *patternFlowGtpExtensionContents) SetValues(value []uint64) PatternFlowGtpExtensionContents {
	obj.setChoice(PatternFlowGtpExtensionContentsChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint64, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionContentsCounter
func (obj *patternFlowGtpExtensionContents) Increment() PatternFlowGtpExtensionContentsCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowGtpExtensionContentsChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowGtpExtensionContentsCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowGtpExtensionContentsCounter
func (obj *patternFlowGtpExtensionContents) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowGtpExtensionContentsCounter value in the PatternFlowGtpExtensionContents object
func (obj *patternFlowGtpExtensionContents) SetIncrement(value PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContents {
	obj.setChoice(PatternFlowGtpExtensionContentsChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionContentsCounter
func (obj *patternFlowGtpExtensionContents) Decrement() PatternFlowGtpExtensionContentsCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowGtpExtensionContentsChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowGtpExtensionContentsCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowGtpExtensionContentsCounter
func (obj *patternFlowGtpExtensionContents) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowGtpExtensionContentsCounter value in the PatternFlowGtpExtensionContents object
func (obj *patternFlowGtpExtensionContents) SetDecrement(value PatternFlowGtpExtensionContentsCounter) PatternFlowGtpExtensionContents {
	obj.setChoice(PatternFlowGtpExtensionContentsChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowGtpExtensionContentsMetricTag
func (obj *patternFlowGtpExtensionContents) MetricTags() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowGtpExtensionContentsMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter struct {
	obj                                           *patternFlowGtpExtensionContents
	patternFlowGtpExtensionContentsMetricTagSlice []PatternFlowGtpExtensionContentsMetricTag
	fieldPtr                                      *[]*otg.PatternFlowGtpExtensionContentsMetricTag
}

func newPatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter(ptr *[]*otg.PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	return &patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter{fieldPtr: ptr}
}

type PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter interface {
	setMsg(*patternFlowGtpExtensionContents) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	Items() []PatternFlowGtpExtensionContentsMetricTag
	Add() PatternFlowGtpExtensionContentsMetricTag
	Append(items ...PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	Set(index int, newObj PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	Clear() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	clearHolderSlice() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
	appendHolderSlice(item PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter
}

func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) setMsg(msg *patternFlowGtpExtensionContents) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowGtpExtensionContentsMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) Items() []PatternFlowGtpExtensionContentsMetricTag {
	return obj.patternFlowGtpExtensionContentsMetricTagSlice
}

func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) Add() PatternFlowGtpExtensionContentsMetricTag {
	newObj := &otg.PatternFlowGtpExtensionContentsMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowGtpExtensionContentsMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowGtpExtensionContentsMetricTagSlice = append(obj.patternFlowGtpExtensionContentsMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) Append(items ...PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowGtpExtensionContentsMetricTagSlice = append(obj.patternFlowGtpExtensionContentsMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) Set(index int, newObj PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowGtpExtensionContentsMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) Clear() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowGtpExtensionContentsMetricTag{}
		obj.patternFlowGtpExtensionContentsMetricTagSlice = []PatternFlowGtpExtensionContentsMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) clearHolderSlice() PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	if len(obj.patternFlowGtpExtensionContentsMetricTagSlice) > 0 {
		obj.patternFlowGtpExtensionContentsMetricTagSlice = []PatternFlowGtpExtensionContentsMetricTag{}
	}
	return obj
}
func (obj *patternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter) appendHolderSlice(item PatternFlowGtpExtensionContentsMetricTag) PatternFlowGtpExtensionContentsPatternFlowGtpExtensionContentsMetricTagIter {
	obj.patternFlowGtpExtensionContentsMetricTagSlice = append(obj.patternFlowGtpExtensionContentsMetricTagSlice, item)
	return obj
}

func (obj *patternFlowGtpExtensionContents) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 281474976710655 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternFlowGtpExtensionContents.Value <= 281474976710655 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 281474976710655 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint64) <= PatternFlowGtpExtensionContents.Values <= 281474976710655 but Got %d", item))
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
				obj.MetricTags().appendHolderSlice(&patternFlowGtpExtensionContentsMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowGtpExtensionContents) setDefault() {
	var choices_set int = 0
	var choice PatternFlowGtpExtensionContentsChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionContentsChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowGtpExtensionContentsChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionContentsChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowGtpExtensionContentsChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowGtpExtensionContentsChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowGtpExtensionContents")
			}
		} else {
			intVal := otg.PatternFlowGtpExtensionContents_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowGtpExtensionContents_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
