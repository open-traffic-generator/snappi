package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternFlowEthernetDst *****
type patternFlowEthernetDst struct {
	validation
	obj              *otg.PatternFlowEthernetDst
	marshaller       marshalPatternFlowEthernetDst
	unMarshaller     unMarshalPatternFlowEthernetDst
	incrementHolder  PatternFlowEthernetDstCounter
	decrementHolder  PatternFlowEthernetDstCounter
	metricTagsHolder PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
}

func NewPatternFlowEthernetDst() PatternFlowEthernetDst {
	obj := patternFlowEthernetDst{obj: &otg.PatternFlowEthernetDst{}}
	obj.setDefault()
	return &obj
}

func (obj *patternFlowEthernetDst) msg() *otg.PatternFlowEthernetDst {
	return obj.obj
}

func (obj *patternFlowEthernetDst) setMsg(msg *otg.PatternFlowEthernetDst) PatternFlowEthernetDst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternFlowEthernetDst struct {
	obj *patternFlowEthernetDst
}

type marshalPatternFlowEthernetDst interface {
	// ToProto marshals PatternFlowEthernetDst to protobuf object *otg.PatternFlowEthernetDst
	ToProto() (*otg.PatternFlowEthernetDst, error)
	// ToPbText marshals PatternFlowEthernetDst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternFlowEthernetDst to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternFlowEthernetDst to JSON text
	ToJson() (string, error)
}

type unMarshalpatternFlowEthernetDst struct {
	obj *patternFlowEthernetDst
}

type unMarshalPatternFlowEthernetDst interface {
	// FromProto unmarshals PatternFlowEthernetDst from protobuf object *otg.PatternFlowEthernetDst
	FromProto(msg *otg.PatternFlowEthernetDst) (PatternFlowEthernetDst, error)
	// FromPbText unmarshals PatternFlowEthernetDst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternFlowEthernetDst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternFlowEthernetDst from JSON text
	FromJson(value string) error
}

func (obj *patternFlowEthernetDst) Marshal() marshalPatternFlowEthernetDst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternFlowEthernetDst{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternFlowEthernetDst) Unmarshal() unMarshalPatternFlowEthernetDst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternFlowEthernetDst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternFlowEthernetDst) ToProto() (*otg.PatternFlowEthernetDst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternFlowEthernetDst) FromProto(msg *otg.PatternFlowEthernetDst) (PatternFlowEthernetDst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternFlowEthernetDst) ToPbText() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDst) FromPbText(value string) error {
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

func (m *marshalpatternFlowEthernetDst) ToYaml() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDst) FromYaml(value string) error {
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

func (m *marshalpatternFlowEthernetDst) ToJson() (string, error) {
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

func (m *unMarshalpatternFlowEthernetDst) FromJson(value string) error {
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

func (obj *patternFlowEthernetDst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternFlowEthernetDst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternFlowEthernetDst) Clone() (PatternFlowEthernetDst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternFlowEthernetDst()
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

func (obj *patternFlowEthernetDst) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.metricTagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternFlowEthernetDst is destination MAC address
type PatternFlowEthernetDst interface {
	Validation
	// msg marshals PatternFlowEthernetDst to protobuf object *otg.PatternFlowEthernetDst
	// and doesn't set defaults
	msg() *otg.PatternFlowEthernetDst
	// setMsg unmarshals PatternFlowEthernetDst from protobuf object *otg.PatternFlowEthernetDst
	// and doesn't set defaults
	setMsg(*otg.PatternFlowEthernetDst) PatternFlowEthernetDst
	// provides marshal interface
	Marshal() marshalPatternFlowEthernetDst
	// provides unmarshal interface
	Unmarshal() unMarshalPatternFlowEthernetDst
	// validate validates PatternFlowEthernetDst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternFlowEthernetDst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternFlowEthernetDstChoiceEnum, set in PatternFlowEthernetDst
	Choice() PatternFlowEthernetDstChoiceEnum
	// setChoice assigns PatternFlowEthernetDstChoiceEnum provided by user to PatternFlowEthernetDst
	setChoice(value PatternFlowEthernetDstChoiceEnum) PatternFlowEthernetDst
	// HasChoice checks if Choice has been set in PatternFlowEthernetDst
	HasChoice() bool
	// Value returns string, set in PatternFlowEthernetDst.
	Value() string
	// SetValue assigns string provided by user to PatternFlowEthernetDst
	SetValue(value string) PatternFlowEthernetDst
	// HasValue checks if Value has been set in PatternFlowEthernetDst
	HasValue() bool
	// Values returns []string, set in PatternFlowEthernetDst.
	Values() []string
	// SetValues assigns []string provided by user to PatternFlowEthernetDst
	SetValues(value []string) PatternFlowEthernetDst
	// Auto returns string, set in PatternFlowEthernetDst.
	Auto() string
	// HasAuto checks if Auto has been set in PatternFlowEthernetDst
	HasAuto() bool
	// Increment returns PatternFlowEthernetDstCounter, set in PatternFlowEthernetDst.
	// PatternFlowEthernetDstCounter is mac counter pattern
	Increment() PatternFlowEthernetDstCounter
	// SetIncrement assigns PatternFlowEthernetDstCounter provided by user to PatternFlowEthernetDst.
	// PatternFlowEthernetDstCounter is mac counter pattern
	SetIncrement(value PatternFlowEthernetDstCounter) PatternFlowEthernetDst
	// HasIncrement checks if Increment has been set in PatternFlowEthernetDst
	HasIncrement() bool
	// Decrement returns PatternFlowEthernetDstCounter, set in PatternFlowEthernetDst.
	// PatternFlowEthernetDstCounter is mac counter pattern
	Decrement() PatternFlowEthernetDstCounter
	// SetDecrement assigns PatternFlowEthernetDstCounter provided by user to PatternFlowEthernetDst.
	// PatternFlowEthernetDstCounter is mac counter pattern
	SetDecrement(value PatternFlowEthernetDstCounter) PatternFlowEthernetDst
	// HasDecrement checks if Decrement has been set in PatternFlowEthernetDst
	HasDecrement() bool
	// MetricTags returns PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIterIter, set in PatternFlowEthernetDst
	MetricTags() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	setNil()
}

type PatternFlowEthernetDstChoiceEnum string

// Enum of Choice on PatternFlowEthernetDst
var PatternFlowEthernetDstChoice = struct {
	VALUE     PatternFlowEthernetDstChoiceEnum
	VALUES    PatternFlowEthernetDstChoiceEnum
	AUTO      PatternFlowEthernetDstChoiceEnum
	INCREMENT PatternFlowEthernetDstChoiceEnum
	DECREMENT PatternFlowEthernetDstChoiceEnum
}{
	VALUE:     PatternFlowEthernetDstChoiceEnum("value"),
	VALUES:    PatternFlowEthernetDstChoiceEnum("values"),
	AUTO:      PatternFlowEthernetDstChoiceEnum("auto"),
	INCREMENT: PatternFlowEthernetDstChoiceEnum("increment"),
	DECREMENT: PatternFlowEthernetDstChoiceEnum("decrement"),
}

func (obj *patternFlowEthernetDst) Choice() PatternFlowEthernetDstChoiceEnum {
	return PatternFlowEthernetDstChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternFlowEthernetDst) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternFlowEthernetDst) setChoice(value PatternFlowEthernetDstChoiceEnum) PatternFlowEthernetDst {
	intValue, ok := otg.PatternFlowEthernetDst_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternFlowEthernetDstChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PatternFlowEthernetDst_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternFlowEthernetDstChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternFlowEthernetDstChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternFlowEthernetDstChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == PatternFlowEthernetDstChoice.INCREMENT {
		obj.obj.Increment = NewPatternFlowEthernetDstCounter().msg()
	}

	if value == PatternFlowEthernetDstChoice.DECREMENT {
		obj.obj.Decrement = NewPatternFlowEthernetDstCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetDst) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternFlowEthernetDstChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternFlowEthernetDst) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternFlowEthernetDst object
func (obj *patternFlowEthernetDst) SetValue(value string) PatternFlowEthernetDst {
	obj.setChoice(PatternFlowEthernetDstChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternFlowEthernetDst) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternFlowEthernetDst object
func (obj *patternFlowEthernetDst) SetValues(value []string) PatternFlowEthernetDst {
	obj.setChoice(PatternFlowEthernetDstChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternFlowEthernetDst) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternFlowEthernetDstChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternFlowEthernetDst) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternFlowEthernetDstCounter
func (obj *patternFlowEthernetDst) Increment() PatternFlowEthernetDstCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternFlowEthernetDstChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternFlowEthernetDstCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternFlowEthernetDstCounter
func (obj *patternFlowEthernetDst) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternFlowEthernetDstCounter value in the PatternFlowEthernetDst object
func (obj *patternFlowEthernetDst) SetIncrement(value PatternFlowEthernetDstCounter) PatternFlowEthernetDst {
	obj.setChoice(PatternFlowEthernetDstChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternFlowEthernetDstCounter
func (obj *patternFlowEthernetDst) Decrement() PatternFlowEthernetDstCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternFlowEthernetDstChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternFlowEthernetDstCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternFlowEthernetDstCounter
func (obj *patternFlowEthernetDst) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternFlowEthernetDstCounter value in the PatternFlowEthernetDst object
func (obj *patternFlowEthernetDst) SetDecrement(value PatternFlowEthernetDstCounter) PatternFlowEthernetDst {
	obj.setChoice(PatternFlowEthernetDstChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

// One or more metric tags can be used to enable tracking portion of or all bits in a corresponding header field for metrics per each applicable value. These would appear as tagged metrics in corresponding flow metrics.
// MetricTags returns a []PatternFlowEthernetDstMetricTag
func (obj *patternFlowEthernetDst) MetricTags() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	if len(obj.obj.MetricTags) == 0 {
		obj.obj.MetricTags = []*otg.PatternFlowEthernetDstMetricTag{}
	}
	if obj.metricTagsHolder == nil {
		obj.metricTagsHolder = newPatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter(&obj.obj.MetricTags).setMsg(obj)
	}
	return obj.metricTagsHolder
}

type patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter struct {
	obj                                  *patternFlowEthernetDst
	patternFlowEthernetDstMetricTagSlice []PatternFlowEthernetDstMetricTag
	fieldPtr                             *[]*otg.PatternFlowEthernetDstMetricTag
}

func newPatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter(ptr *[]*otg.PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	return &patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter{fieldPtr: ptr}
}

type PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter interface {
	setMsg(*patternFlowEthernetDst) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	Items() []PatternFlowEthernetDstMetricTag
	Add() PatternFlowEthernetDstMetricTag
	Append(items ...PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	Set(index int, newObj PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	Clear() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	clearHolderSlice() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
	appendHolderSlice(item PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter
}

func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) setMsg(msg *patternFlowEthernetDst) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&patternFlowEthernetDstMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) Items() []PatternFlowEthernetDstMetricTag {
	return obj.patternFlowEthernetDstMetricTagSlice
}

func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) Add() PatternFlowEthernetDstMetricTag {
	newObj := &otg.PatternFlowEthernetDstMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &patternFlowEthernetDstMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.patternFlowEthernetDstMetricTagSlice = append(obj.patternFlowEthernetDstMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) Append(items ...PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.patternFlowEthernetDstMetricTagSlice = append(obj.patternFlowEthernetDstMetricTagSlice, item)
	}
	return obj
}

func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) Set(index int, newObj PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.patternFlowEthernetDstMetricTagSlice[index] = newObj
	return obj
}
func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) Clear() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.PatternFlowEthernetDstMetricTag{}
		obj.patternFlowEthernetDstMetricTagSlice = []PatternFlowEthernetDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) clearHolderSlice() PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	if len(obj.patternFlowEthernetDstMetricTagSlice) > 0 {
		obj.patternFlowEthernetDstMetricTagSlice = []PatternFlowEthernetDstMetricTag{}
	}
	return obj
}
func (obj *patternFlowEthernetDstPatternFlowEthernetDstMetricTagIter) appendHolderSlice(item PatternFlowEthernetDstMetricTag) PatternFlowEthernetDstPatternFlowEthernetDstMetricTagIter {
	obj.patternFlowEthernetDstMetricTagSlice = append(obj.patternFlowEthernetDstMetricTagSlice, item)
	return obj
}

func (obj *patternFlowEthernetDst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternFlowEthernetDst.Auto"))
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
				obj.MetricTags().appendHolderSlice(&patternFlowEthernetDstMetricTag{obj: item})
			}
		}
		for _, item := range obj.MetricTags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *patternFlowEthernetDst) setDefault() {
	var choices_set int = 0
	var choice PatternFlowEthernetDstChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternFlowEthernetDstChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternFlowEthernetDstChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternFlowEthernetDstChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternFlowEthernetDstChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternFlowEthernetDstChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternFlowEthernetDstChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternFlowEthernetDst")
			}
		} else {
			intVal := otg.PatternFlowEthernetDst_Choice_Enum_value[string(choice)]
			enumValue := otg.PatternFlowEthernetDst_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
