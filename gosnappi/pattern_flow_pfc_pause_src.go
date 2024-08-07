package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseSrc *****
type patternFlowPfcPauseSrc struct {
	validation
	obj              *otg.PatternFlowPfcPauseSrc
	marshaller       marshalPatternFlowPfcPauseSrc
	unMarshaller     unMarshalPatternFlowPfcPauseSrc
	incrementHolder  PatternFlowPfcPauseSrcCounter
	decrementHolder  PatternFlowPfcPauseSrcCounter
	metricTagsHolder PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
}

func NewPatternFlowPfcPauseSrc() PatternFlowPfcPauseSrc {
	obj := patternFlowPfcPauseSrc{obj: &otg.PatternFlowPfcPauseSrc{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseSrc) msg() *otg.PatternFlowPfcPauseSrc {
	return obj.obj
}

func (obj *patternFlowPfcPauseSrc) setMsg(msg *otg.PatternFlowPfcPauseSrc) PatternFlowPfcPauseSrc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseSrc struct {
	obj *patternFlowPfcPauseSrc
}

type marshalPatternFlowPfcPauseSrc interface {
	// ToProto marshals PatternFlowPfcPauseSrc to protobuf object *otg.PatternFlowPfcPauseSrc
	ToProto() (*otg.PatternFlowPfcPauseSrc, error)
	// ToPbText marshals PatternFlowPfcPauseSrc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseSrc to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseSrc to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseSrc struct {
	obj *patternFlowPfcPauseSrc
}

type unMarshalPatternFlowPfcPauseSrc interface {
	// FromProto unmarshals PatternFlowPfcPauseSrc from protobuf object *otg.PatternFlowPfcPauseSrc
	FromProto(msg *otg.PatternFlowPfcPauseSrc) (PatternFlowPfcPauseSrc, error)
	// FromPbText unmarshals PatternFlowPfcPauseSrc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseSrc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseSrc from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseSrc) Marshal() marshalPatternFlowPfcPauseSrc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseSrc{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseSrc) Unmarshal() unMarshalPatternFlowPfcPauseSrc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseSrc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseSrc) ToProto() (*otg.PatternFlowPfcPauseSrc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseSrc) FromProto(msg *otg.PatternFlowPfcPauseSrc) (PatternFlowPfcPauseSrc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseSrc) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrc) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrc) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrc) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseSrc) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseSrc) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseSrc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseSrc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseSrc) Clone() (PatternFlowPfcPauseSrc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseSrc()
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

func (obj *patternFlowPfcPauseSrc) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPauseSrc is source MAC address
type PatternFlowPfcPauseSrc interface {
	Validation
	// msg marshals PatternFlowPfcPauseSrc to protobuf object *otg.PatternFlowPfcPauseSrc
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseSrc
	// setMsg unmarshals PatternFlowPfcPauseSrc from protobuf object *otg.PatternFlowPfcPauseSrc
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseSrc) PatternFlowPfcPauseSrc
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseSrc
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseSrc
	// validate validates PatternFlowPfcPauseSrc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseSrc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPauseSrcChoiceEnum, set in PatternFlowPfcPauseSrc
	Choice() PatternFlowPfcPauseSrcChoiceEnum
	// setChoice assigns PatternFlowPfcPauseSrcChoiceEnum provided by user to PatternFlowPfcPauseSrc
	setChoice(value PatternFlowPfcPauseSrcChoiceEnum) PatternFlowPfcPauseSrc
	// HasChoice checks if Choice has been set in PatternFlowPfcPauseSrc
	HasChoice() bool
	// Value returns string, set in PatternFlowPfcPauseSrc.
	Value() string
	// SetValue assigns string provided by user to PatternFlowPfcPauseSrc
	SetValue(value string) PatternFlowPfcPauseSrc
	// HasValue checks if Value has been set in PatternFlowPfcPauseSrc
	HasValue() bool
	// Values returns []string, set in PatternFlowPfcPauseSrc.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowPfcPauseSrc
	SetValues(value []string) PatternFlowPfcPauseSrc
	// Increment returns PatternFlowPfcPauseSrcCounter, set in PatternFlowPfcPauseSrc.
	// PatternFlowPfcPauseSrcCounter is mac counter pattern
	Increment() PatternFlowPfcPauseSrcCounter
	// SetIncrement assigns PatternFlowPfcPauseSrcCounter provided by user to PatternFlowPfcPauseSrc.
	// PatternFlowPfcPauseSrcCounter is mac counter pattern
	SetIncrement(value PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrc
	// HasIncrement checks if Increment has been set in PatternFlowPfcPauseSrc
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPauseSrcCounter, set in PatternFlowPfcPauseSrc.
	// PatternFlowPfcPauseSrcCounter is mac counter pattern
	Decrement() PatternFlowPfcPauseSrcCounter
	// SetDecrement assigns PatternFlowPfcPauseSrcCounter provided by user to PatternFlowPfcPauseSrc.
	// PatternFlowPfcPauseSrcCounter is mac counter pattern
	SetDecrement(value PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrc
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPauseSrc
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIterIter, set in PatternFlowPfcPauseSrc
	MetricTags() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	setNil()
}

type PatternFlowPfcPauseSrcChoiceEnum string

// Enum of Choice on PatternFlowPfcPauseSrc
var PatternFlowPfcPauseSrcChoice = struct {
	VALUE     PatternFlowPfcPauseSrcChoiceEnum
	VALUES    PatternFlowPfcPauseSrcChoiceEnum
	INCREMENT PatternFlowPfcPauseSrcChoiceEnum
	DECREMENT PatternFlowPfcPauseSrcChoiceEnum
}{
	VALUE:     PatternFlowPfcPauseSrcChoiceEnum("value"),
	VALUES:    PatternFlowPfcPauseSrcChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPauseSrcChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPauseSrcChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPauseSrc) Choice() PatternFlowPfcPauseSrcChoiceEnum {
	return PatternFlowPfcPauseSrcChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPauseSrc) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPauseSrc) setChoice(value PatternFlowPfcPauseSrcChoiceEnum) PatternFlowPfcPauseSrc {
	intValue, ok := otg.PatternFlowPfcPauseSrc_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPauseSrcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPauseSrc_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPauseSrcChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPauseSrcChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPauseSrcChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPauseSrcCounter().msg()
	}

	if value == PatternFlowPfcPauseSrcChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPauseSrcCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowPfcPauseSrc) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPauseSrcChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowPfcPauseSrc) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowPfcPauseSrc object
func (obj *patternFlowPfcPauseSrc) SetValue(value string) PatternFlowPfcPauseSrc {
	obj.setChoice(PatternFlowPfcPauseSrcChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowPfcPauseSrc) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowPfcPauseSrc object
func (obj *patternFlowPfcPauseSrc) SetValues(value []string) PatternFlowPfcPauseSrc {
	obj.setChoice(PatternFlowPfcPauseSrcChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPauseSrcCounter
func (obj *patternFlowPfcPauseSrc) Increment() PatternFlowPfcPauseSrcCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPauseSrcChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPauseSrcCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPauseSrcCounter
func (obj *patternFlowPfcPauseSrc) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPauseSrcCounter value in the PatternFlowPfcPauseSrc object
func (obj *patternFlowPfcPauseSrc) SetIncrement(value PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrc {
	obj.setChoice(PatternFlowPfcPauseSrcChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseSrcCounter
func (obj *patternFlowPfcPauseSrc) Decrement() PatternFlowPfcPauseSrcCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPauseSrcChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPauseSrcCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseSrcCounter
func (obj *patternFlowPfcPauseSrc) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPauseSrcCounter value in the PatternFlowPfcPauseSrc object
func (obj *patternFlowPfcPauseSrc) SetDecrement(value PatternFlowPfcPauseSrcCounter) PatternFlowPfcPauseSrc {
	obj.setChoice(PatternFlowPfcPauseSrcChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPauseSrcMetricTag
func (obj *patternFlowPfcPauseSrc) MetricTags() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPauseSrcMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter struct {
	obj                                  *patternFlowPfcPauseSrc
	patternFlowPfcPauseSrcMetricTagSlice []PatternFlowPfcPauseSrcMetricTag
	fieldPtr                             *[]*otg.PatternFlowPfcPauseSrcMetricTag
}

func newPatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter(ptr *[]*otg.PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	return &patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter interface {
	setMsg(*patternFlowPfcPauseSrc) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	Items() []PatternFlowPfcPauseSrcMetricTag
	Add() PatternFlowPfcPauseSrcMetricTag
	Append(items ...PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	Set(index int, newObj PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	Clear() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	clearHolderSlice() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
	appendHolderSlice(item PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter
}

func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) setMsg(msg *patternFlowPfcPauseSrc) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPauseSrcMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) Items() []PatternFlowPfcPauseSrcMetricTag {
	return obj.patternFlowPfcPauseSrcMetricTagSlice
}

func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) Add() PatternFlowPfcPauseSrcMetricTag {
	newObj := &otg.PatternFlowPfcPauseSrcMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPauseSrcMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPauseSrcMetricTagSlice = append(obj.patternFlowPfcPauseSrcMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) Append(items ...PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPauseSrcMetricTagSlice = append(obj.patternFlowPfcPauseSrcMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) Set(index int, newObj PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPauseSrcMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) Clear() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPauseSrcMetricTag{}
		obj.patternFlowPfcPauseSrcMetricTagSlice = []PatternFlowPfcPauseSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) clearHolderSlice() PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	if len(obj.patternFlowPfcPauseSrcMetricTagSlice) > 0 {
		obj.patternFlowPfcPauseSrcMetricTagSlice = []PatternFlowPfcPauseSrcMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter) appendHolderSlice(item PatternFlowPfcPauseSrcMetricTag) PatternFlowPfcPauseSrcPatternFlowPfcPauseSrcMetricTagIter {
	obj.patternFlowPfcPauseSrcMetricTagSlice = append(obj.patternFlowPfcPauseSrcMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPauseSrc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrc.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseSrc.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPauseSrcMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPauseSrc) setDefault() {
	var choices_set int = 0
	var choice PatternFlowPfcPauseSrcChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseSrcChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowPfcPauseSrcChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseSrcChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowPfcPauseSrcChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowPfcPauseSrcChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowPfcPauseSrc")
			}
		} else {
			intVal := otg.PatternFlowPfcPauseSrc_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowPfcPauseSrc_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
