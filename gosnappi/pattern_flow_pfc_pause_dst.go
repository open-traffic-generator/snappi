package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowPfcPauseDst *****
type patternFlowPfcPauseDst struct {
	validation
	obj              *otg.PatternFlowPfcPauseDst
	marshaller       marshalPatternFlowPfcPauseDst
	unMarshaller     unMarshalPatternFlowPfcPauseDst
	incrementHolder  PatternFlowPfcPauseDstCounter
	decrementHolder  PatternFlowPfcPauseDstCounter
	metricTagsHolder PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
}

func NewPatternFlowPfcPauseDst() PatternFlowPfcPauseDst {
	obj := patternFlowPfcPauseDst{obj: &otg.PatternFlowPfcPauseDst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowPfcPauseDst) msg() *otg.PatternFlowPfcPauseDst {
	return obj.obj
}

func (obj *patternFlowPfcPauseDst) setMsg(msg *otg.PatternFlowPfcPauseDst) PatternFlowPfcPauseDst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowPfcPauseDst struct {
	obj *patternFlowPfcPauseDst
}

type marshalPatternFlowPfcPauseDst interface {
	// ToProto marshals PatternFlowPfcPauseDst to protobuf object *otg.PatternFlowPfcPauseDst
	ToProto() (*otg.PatternFlowPfcPauseDst, error)
	// ToPbText marshals PatternFlowPfcPauseDst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowPfcPauseDst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowPfcPauseDst to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowPfcPauseDst struct {
	obj *patternFlowPfcPauseDst
}

type unMarshalPatternFlowPfcPauseDst interface {
	// FromProto unmarshals PatternFlowPfcPauseDst from protobuf object *otg.PatternFlowPfcPauseDst
	FromProto(msg *otg.PatternFlowPfcPauseDst) (PatternFlowPfcPauseDst, error)
	// FromPbText unmarshals PatternFlowPfcPauseDst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowPfcPauseDst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowPfcPauseDst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowPfcPauseDst) Marshal() marshalPatternFlowPfcPauseDst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowPfcPauseDst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowPfcPauseDst) Unmarshal() unMarshalPatternFlowPfcPauseDst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowPfcPauseDst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowPfcPauseDst) ToProto() (*otg.PatternFlowPfcPauseDst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowPfcPauseDst) FromProto(msg *otg.PatternFlowPfcPauseDst) (PatternFlowPfcPauseDst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowPfcPauseDst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDst) FromPbText(value string) error {
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

func (m *marshalpatternFlowPfcPauseDst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDst) FromYaml(value string) error {
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

func (m *marshalpatternFlowPfcPauseDst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowPfcPauseDst) FromJson(value string) error {
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

func (obj *patternFlowPfcPauseDst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowPfcPauseDst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowPfcPauseDst) Clone() (PatternFlowPfcPauseDst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowPfcPauseDst()
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

func (obj *patternFlowPfcPauseDst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowPfcPauseDst is destination MAC address
type PatternFlowPfcPauseDst interface {
	Validation
	// msg marshals PatternFlowPfcPauseDst to protobuf object *otg.PatternFlowPfcPauseDst
	// and doesn't set defaults
	msg() *otg.PatternFlowPfcPauseDst
	// setMsg unmarshals PatternFlowPfcPauseDst from protobuf object *otg.PatternFlowPfcPauseDst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowPfcPauseDst) PatternFlowPfcPauseDst
	// provides marshal interface
	Marshal() marshalPatternFlowPfcPauseDst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowPfcPauseDst
	// validate validates PatternFlowPfcPauseDst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowPfcPauseDst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowPfcPauseDstChoiceEnum, set in PatternFlowPfcPauseDst
	Choice() PatternFlowPfcPauseDstChoiceEnum
	// setChoice assigns PatternFlowPfcPauseDstChoiceEnum provided by user to PatternFlowPfcPauseDst
	setChoice(value PatternFlowPfcPauseDstChoiceEnum) PatternFlowPfcPauseDst
	// HasChoice checks if Choice has been set in PatternFlowPfcPauseDst
	HasChoice() bool
	// Value returns string, set in PatternFlowPfcPauseDst.
	Value() string
	// SetValue assigns string provided by user to PatternFlowPfcPauseDst
	SetValue(value string) PatternFlowPfcPauseDst
	// HasValue checks if Value has been set in PatternFlowPfcPauseDst
	HasValue() bool
	// Values returns []string, set in PatternFlowPfcPauseDst.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowPfcPauseDst
	SetValues(value []string) PatternFlowPfcPauseDst
	// Increment returns PatternFlowPfcPauseDstCounter, set in PatternFlowPfcPauseDst.
	// PatternFlowPfcPauseDstCounter is mac counter pattern
	Increment() PatternFlowPfcPauseDstCounter
	// SetIncrement assigns PatternFlowPfcPauseDstCounter provided by user to PatternFlowPfcPauseDst.
	// PatternFlowPfcPauseDstCounter is mac counter pattern
	SetIncrement(value PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDst
	// HasIncrement checks if Increment has been set in PatternFlowPfcPauseDst
	HasIncrement() bool
	// Decrement returns PatternFlowPfcPauseDstCounter, set in PatternFlowPfcPauseDst.
	// PatternFlowPfcPauseDstCounter is mac counter pattern
	Decrement() PatternFlowPfcPauseDstCounter
	// SetDecrement assigns PatternFlowPfcPauseDstCounter provided by user to PatternFlowPfcPauseDst.
	// PatternFlowPfcPauseDstCounter is mac counter pattern
	SetDecrement(value PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDst
	// HasDecrement checks if Decrement has been set in PatternFlowPfcPauseDst
	HasDecrement() bool
	// MetricTags returns PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIterIter, set in PatternFlowPfcPauseDst
	MetricTags() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	setNil()
}

type PatternFlowPfcPauseDstChoiceEnum string

// Enum of Choice on PatternFlowPfcPauseDst
var PatternFlowPfcPauseDstChoice = struct {
	VALUE     PatternFlowPfcPauseDstChoiceEnum
	VALUES    PatternFlowPfcPauseDstChoiceEnum
	INCREMENT PatternFlowPfcPauseDstChoiceEnum
	DECREMENT PatternFlowPfcPauseDstChoiceEnum
}{
	VALUE:     PatternFlowPfcPauseDstChoiceEnum("value"),
	VALUES:    PatternFlowPfcPauseDstChoiceEnum("values"),
	INCREMENT: PatternFlowPfcPauseDstChoiceEnum("increment"),
	DECREMENT: PatternFlowPfcPauseDstChoiceEnum("decrement"),
}

func (obj *patternFlowPfcPauseDst) Choice() PatternFlowPfcPauseDstChoiceEnum {
	return PatternFlowPfcPauseDstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowPfcPauseDst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowPfcPauseDst) setChoice(value PatternFlowPfcPauseDstChoiceEnum) PatternFlowPfcPauseDst {
	intValue, ok := otg.PatternFlowPfcPauseDst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowPfcPauseDstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowPfcPauseDst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowPfcPauseDstChoice.VALUE {
		defaultValue := "01:80:c2:00:00:01"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowPfcPauseDstChoice.VALUES {
		defaultValue := []string{"01:80:c2:00:00:01"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowPfcPauseDstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowPfcPauseDstCounter().msg()
	}

	if value == PatternFlowPfcPauseDstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowPfcPauseDstCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowPfcPauseDst) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowPfcPauseDstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowPfcPauseDst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowPfcPauseDst object
func (obj *patternFlowPfcPauseDst) SetValue(value string) PatternFlowPfcPauseDst {
	obj.setChoice(PatternFlowPfcPauseDstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowPfcPauseDst) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"01:80:c2:00:00:01"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowPfcPauseDst object
func (obj *patternFlowPfcPauseDst) SetValues(value []string) PatternFlowPfcPauseDst {
	obj.setChoice(PatternFlowPfcPauseDstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowPfcPauseDstCounter
func (obj *patternFlowPfcPauseDst) Increment() PatternFlowPfcPauseDstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowPfcPauseDstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowPfcPauseDstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowPfcPauseDstCounter
func (obj *patternFlowPfcPauseDst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowPfcPauseDstCounter value in the PatternFlowPfcPauseDst object
func (obj *patternFlowPfcPauseDst) SetIncrement(value PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDst {
	obj.setChoice(PatternFlowPfcPauseDstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseDstCounter
func (obj *patternFlowPfcPauseDst) Decrement() PatternFlowPfcPauseDstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowPfcPauseDstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowPfcPauseDstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowPfcPauseDstCounter
func (obj *patternFlowPfcPauseDst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowPfcPauseDstCounter value in the PatternFlowPfcPauseDst object
func (obj *patternFlowPfcPauseDst) SetDecrement(value PatternFlowPfcPauseDstCounter) PatternFlowPfcPauseDst {
	obj.setChoice(PatternFlowPfcPauseDstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowPfcPauseDstMetricTag
func (obj *patternFlowPfcPauseDst) MetricTags() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowPfcPauseDstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter struct {
	obj                                  *patternFlowPfcPauseDst
	patternFlowPfcPauseDstMetricTagSlice []PatternFlowPfcPauseDstMetricTag
	fieldPtr                             *[]*otg.PatternFlowPfcPauseDstMetricTag
}

func newPatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter(ptr *[]*otg.PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	return &patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter interface {
	setMsg(*patternFlowPfcPauseDst) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	Items() []PatternFlowPfcPauseDstMetricTag
	Add() PatternFlowPfcPauseDstMetricTag
	Append(items ...PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	Set(index int, newObj PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	Clear() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	clearHolderSlice() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
	appendHolderSlice(item PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter
}

func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) setMsg(msg *patternFlowPfcPauseDst) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowPfcPauseDstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) Items() []PatternFlowPfcPauseDstMetricTag {
	return obj.patternFlowPfcPauseDstMetricTagSlice
}

func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) Add() PatternFlowPfcPauseDstMetricTag {
	newObj := &otg.PatternFlowPfcPauseDstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowPfcPauseDstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowPfcPauseDstMetricTagSlice = append(obj.patternFlowPfcPauseDstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) Append(items ...PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowPfcPauseDstMetricTagSlice = append(obj.patternFlowPfcPauseDstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) Set(index int, newObj PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowPfcPauseDstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) Clear() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowPfcPauseDstMetricTag{}
		obj.patternFlowPfcPauseDstMetricTagSlice = []PatternFlowPfcPauseDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) clearHolderSlice() PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	if len(obj.patternFlowPfcPauseDstMetricTagSlice) > 0 {
		obj.patternFlowPfcPauseDstMetricTagSlice = []PatternFlowPfcPauseDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter) appendHolderSlice(item PatternFlowPfcPauseDstMetricTag) PatternFlowPfcPauseDstPatternFlowPfcPauseDstMetricTagIter {
	obj.patternFlowPfcPauseDstMetricTagSlice = append(obj.patternFlowPfcPauseDstMetricTagSlice, item)
	return obj
}

func (obj *patternFlowPfcPauseDst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowPfcPauseDst.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowPfcPauseDstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowPfcPauseDst) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternFlowPfcPauseDstChoice.VALUE)

	}

}
