package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetPauseDst *****
type patternFlowEthernetPauseDst struct {
	validation
	obj              *otg.PatternFlowEthernetPauseDst
	marshaller       marshalPatternFlowEthernetPauseDst
	unMarshaller     unMarshalPatternFlowEthernetPauseDst
	incrementHolder  PatternFlowEthernetPauseDstCounter
	decrementHolder  PatternFlowEthernetPauseDstCounter
	metricTagsHolder PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
}

func NewPatternFlowEthernetPauseDst() PatternFlowEthernetPauseDst {
	obj := patternFlowEthernetPauseDst{obj: &otg.PatternFlowEthernetPauseDst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetPauseDst) msg() *otg.PatternFlowEthernetPauseDst {
	return obj.obj
}

func (obj *patternFlowEthernetPauseDst) setMsg(msg *otg.PatternFlowEthernetPauseDst) PatternFlowEthernetPauseDst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetPauseDst struct {
	obj *patternFlowEthernetPauseDst
}

type marshalPatternFlowEthernetPauseDst interface {
	// ToProto marshals PatternFlowEthernetPauseDst to protobuf object *otg.PatternFlowEthernetPauseDst
	ToProto() (*otg.PatternFlowEthernetPauseDst, error)
	// ToPbText marshals PatternFlowEthernetPauseDst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetPauseDst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetPauseDst to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetPauseDst struct {
	obj *patternFlowEthernetPauseDst
}

type unMarshalPatternFlowEthernetPauseDst interface {
	// FromProto unmarshals PatternFlowEthernetPauseDst from protobuf object *otg.PatternFlowEthernetPauseDst
	FromProto(msg *otg.PatternFlowEthernetPauseDst) (PatternFlowEthernetPauseDst, error)
	// FromPbText unmarshals PatternFlowEthernetPauseDst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetPauseDst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetPauseDst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetPauseDst) Marshal() marshalPatternFlowEthernetPauseDst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetPauseDst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetPauseDst) Unmarshal() unMarshalPatternFlowEthernetPauseDst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetPauseDst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetPauseDst) ToProto() (*otg.PatternFlowEthernetPauseDst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetPauseDst) FromProto(msg *otg.PatternFlowEthernetPauseDst) (PatternFlowEthernetPauseDst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetPauseDst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDst) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDst) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetPauseDst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetPauseDst) FromJson(value string) error {
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

func (obj *patternFlowEthernetPauseDst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetPauseDst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetPauseDst) Clone() (PatternFlowEthernetPauseDst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetPauseDst()
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

func (obj *patternFlowEthernetPauseDst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetPauseDst is destination MAC address
type PatternFlowEthernetPauseDst interface {
	Validation
	// msg marshals PatternFlowEthernetPauseDst to protobuf object *otg.PatternFlowEthernetPauseDst
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetPauseDst
	// setMsg unmarshals PatternFlowEthernetPauseDst from protobuf object *otg.PatternFlowEthernetPauseDst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetPauseDst) PatternFlowEthernetPauseDst
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetPauseDst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetPauseDst
	// validate validates PatternFlowEthernetPauseDst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetPauseDst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetPauseDstChoiceEnum, set in PatternFlowEthernetPauseDst
	Choice() PatternFlowEthernetPauseDstChoiceEnum
	// setChoice assigns PatternFlowEthernetPauseDstChoiceEnum provided by user to PatternFlowEthernetPauseDst
	setChoice(value PatternFlowEthernetPauseDstChoiceEnum) PatternFlowEthernetPauseDst
	// HasChoice checks if Choice has been set in PatternFlowEthernetPauseDst
	HasChoice() bool
	// Value returns string, set in PatternFlowEthernetPauseDst.
	Value() string
	// SetValue assigns string provided by user to PatternFlowEthernetPauseDst
	SetValue(value string) PatternFlowEthernetPauseDst
	// HasValue checks if Value has been set in PatternFlowEthernetPauseDst
	HasValue() bool
	// Values returns []string, set in PatternFlowEthernetPauseDst.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowEthernetPauseDst
	SetValues(value []string) PatternFlowEthernetPauseDst
	// Increment returns PatternFlowEthernetPauseDstCounter, set in PatternFlowEthernetPauseDst.
	// PatternFlowEthernetPauseDstCounter is mac counter pattern
	Increment() PatternFlowEthernetPauseDstCounter
	// SetIncrement assigns PatternFlowEthernetPauseDstCounter provided by user to PatternFlowEthernetPauseDst.
	// PatternFlowEthernetPauseDstCounter is mac counter pattern
	SetIncrement(value PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDst
	// HasIncrement checks if Increment has been set in PatternFlowEthernetPauseDst
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetPauseDstCounter, set in PatternFlowEthernetPauseDst.
	// PatternFlowEthernetPauseDstCounter is mac counter pattern
	Decrement() PatternFlowEthernetPauseDstCounter
	// SetDecrement assigns PatternFlowEthernetPauseDstCounter provided by user to PatternFlowEthernetPauseDst.
	// PatternFlowEthernetPauseDstCounter is mac counter pattern
	SetDecrement(value PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDst
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetPauseDst
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIterIter, set in PatternFlowEthernetPauseDst
	MetricTags() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	setNil()
}

type PatternFlowEthernetPauseDstChoiceEnum string

// Enum of Choice on PatternFlowEthernetPauseDst
var PatternFlowEthernetPauseDstChoice = struct {
	VALUE     PatternFlowEthernetPauseDstChoiceEnum
	VALUES    PatternFlowEthernetPauseDstChoiceEnum
	INCREMENT PatternFlowEthernetPauseDstChoiceEnum
	DECREMENT PatternFlowEthernetPauseDstChoiceEnum
}{
	VALUE:     PatternFlowEthernetPauseDstChoiceEnum("value"),
	VALUES:    PatternFlowEthernetPauseDstChoiceEnum("values"),
	INCREMENT: PatternFlowEthernetPauseDstChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetPauseDstChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetPauseDst) Choice() PatternFlowEthernetPauseDstChoiceEnum {
	return PatternFlowEthernetPauseDstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetPauseDst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetPauseDst) setChoice(value PatternFlowEthernetPauseDstChoiceEnum) PatternFlowEthernetPauseDst {
	intValue, ok := otg.PatternFlowEthernetPauseDst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetPauseDstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetPauseDst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetPauseDstChoice.VALUE {
		defaultValue := "01:80:c2:00:00:01"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetPauseDstChoice.VALUES {
		defaultValue := []string{"01:80:c2:00:00:01"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetPauseDstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetPauseDstCounter().msg()
	}

	if value == PatternFlowEthernetPauseDstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetPauseDstCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetPauseDst) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetPauseDstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetPauseDst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowEthernetPauseDst object
func (obj *patternFlowEthernetPauseDst) SetValue(value string) PatternFlowEthernetPauseDst {
	obj.setChoice(PatternFlowEthernetPauseDstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowEthernetPauseDst) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"01:80:c2:00:00:01"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowEthernetPauseDst object
func (obj *patternFlowEthernetPauseDst) SetValues(value []string) PatternFlowEthernetPauseDst {
	obj.setChoice(PatternFlowEthernetPauseDstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseDstCounter
func (obj *patternFlowEthernetPauseDst) Increment() PatternFlowEthernetPauseDstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetPauseDstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetPauseDstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetPauseDstCounter
func (obj *patternFlowEthernetPauseDst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetPauseDstCounter value in the PatternFlowEthernetPauseDst object
func (obj *patternFlowEthernetPauseDst) SetIncrement(value PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDst {
	obj.setChoice(PatternFlowEthernetPauseDstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseDstCounter
func (obj *patternFlowEthernetPauseDst) Decrement() PatternFlowEthernetPauseDstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetPauseDstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetPauseDstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetPauseDstCounter
func (obj *patternFlowEthernetPauseDst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetPauseDstCounter value in the PatternFlowEthernetPauseDst object
func (obj *patternFlowEthernetPauseDst) SetDecrement(value PatternFlowEthernetPauseDstCounter) PatternFlowEthernetPauseDst {
	obj.setChoice(PatternFlowEthernetPauseDstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetPauseDstMetricTag
func (obj *patternFlowEthernetPauseDst) MetricTags() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetPauseDstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter struct {
	obj                                       *patternFlowEthernetPauseDst
	patternFlowEthernetPauseDstMetricTagSlice []PatternFlowEthernetPauseDstMetricTag
	fieldPtr                                  *[]*otg.PatternFlowEthernetPauseDstMetricTag
}

func newPatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter(ptr *[]*otg.PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	return &patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter interface {
	setMsg(*patternFlowEthernetPauseDst) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	Items() []PatternFlowEthernetPauseDstMetricTag
	Add() PatternFlowEthernetPauseDstMetricTag
	Append(items ...PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	Set(index int, newObj PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	Clear() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	clearHolderSlice() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
	appendHolderSlice(item PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter
}

func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) setMsg(msg *patternFlowEthernetPauseDst) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetPauseDstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) Items() []PatternFlowEthernetPauseDstMetricTag {
	return obj.patternFlowEthernetPauseDstMetricTagSlice
}

func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) Add() PatternFlowEthernetPauseDstMetricTag {
	newObj := &otg.PatternFlowEthernetPauseDstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetPauseDstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetPauseDstMetricTagSlice = append(obj.patternFlowEthernetPauseDstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) Append(items ...PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetPauseDstMetricTagSlice = append(obj.patternFlowEthernetPauseDstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) Set(index int, newObj PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetPauseDstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) Clear() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetPauseDstMetricTag{}
		obj.patternFlowEthernetPauseDstMetricTagSlice = []PatternFlowEthernetPauseDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) clearHolderSlice() PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	if len(obj.patternFlowEthernetPauseDstMetricTagSlice) > 0 {
		obj.patternFlowEthernetPauseDstMetricTagSlice = []PatternFlowEthernetPauseDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter) appendHolderSlice(item PatternFlowEthernetPauseDstMetricTag) PatternFlowEthernetPauseDstPatternFlowEthernetPauseDstMetricTagIter {
	obj.patternFlowEthernetPauseDstMetricTagSlice = append(obj.patternFlowEthernetPauseDstMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetPauseDst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetPauseDst.Values"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetPauseDstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetPauseDst) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetPauseDstChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseDstChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetPauseDstChoice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseDstChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetPauseDstChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetPauseDstChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetPauseDst")
			}
		} else {
			intVal := otg.PatternFlowEthernetPauseDst_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetPauseDst_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
